// Copyright 2015 Midokura SARL
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/samuel/go-zookeeper/zk"
)

// The address of the NSDB.
var nsdbAddresses = flag.String("zookeeper_hosts", "127.0.0.1:2181",
	"The Addresses of ZooKeeper nodes separated by commas")

// The timout for the NSDB session in seconds.
const sessionTimeoutSec = 10

// The key for the ZOOM topology lock path.
const lockKey = "zoom-topology"

func connect() (*zk.Conn, <-chan zk.Event, error) {
	addresses := strings.Split(*nsdbAddresses, ",")
	for i := range addresses {
		addresses[i] = strings.TrimSpace(addresses[i])
	}
	return zk.Connect(addresses,
		time.Duration(sessionTimeoutSec)*time.Second)
}

func binding(portUuid, hostUuid, interfaceName string) error {
	log.Println("binding port " + portUuid + " to " + interfaceName)

	conn, _, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	lock := zk.NewLock(conn, GetLockPath(lockKey), zk.WorldACL(zk.PermAll))
	if err = lock.Lock(); err != nil {
		return err
	}
	defer lock.Unlock()

	portPath := GetPortPath(portUuid)
	var data []byte
	if data, _, err = conn.Get(portPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on getting port %s: %s\n",
			portPath, err.Error())
		return err
	}
	port := &WrappedPort{}
	if err = json.Unmarshal(data, port); err != nil {
		fmt.Fprintf(os.Stderr, "Error on deserializing port %s: %s\n",
			portPath, err.Error())
		return err
	}

	port.Data.HostId = hostUuid
	port.Data.InterfaceName = interfaceName

	updatedPort, err := json.Marshal(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on serializing port %s: %s\n",
			portPath, err.Error())
		return err
	}

	if _, err = conn.Set(portPath, updatedPort, -1); err != nil {
		fmt.Fprintf(os.Stderr, "Error on setting port %s: %s\n",
			portPath, err.Error())
		return err
	}

	vrnMappingPath := GetVrnMappingPath(hostUuid, portUuid)
	var exists bool
	if exists, _, err = conn.Exists(vrnMappingPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on examining vrnMapping %s: %s\n",
			vrnMappingPath, err.Error())
		return err
	}
	var vrnMappingData []byte
	vrnMapping := &WrappedVrnMapping{}
	if exists {
		if vrnMappingData, _, err = conn.Get(vrnMappingPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error on getting vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
		log.Println(fmt.Sprintf("Got vrnMapping data: %s", vrnMappingData))
		if err = json.Unmarshal(vrnMappingData, vrnMapping); err != nil {
			fmt.Fprintf(os.Stderr, "Error on deserializing vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
	} else {
		data := &VrnMapping{}
		data.VirtualPortId = portUuid
		data.LocalDeviceName = interfaceName
		vrnMapping.Data = data
		vrnMapping.Version = port.Version
	}
	updatedVrnMapping, err := json.Marshal(vrnMapping)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on deserializing vrnMapping %s: %s\n",
			vrnMappingPath, err.Error())
		return err
	}
	if exists {
		if _, err = conn.Set(vrnMappingPath, updatedVrnMapping, -1); err != nil {
			fmt.Fprintf(os.Stderr, "Error on setting vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
	} else {
		if _, err = conn.Create(vrnMappingPath, updatedVrnMapping, 0, zk.WorldACL(zk.PermAll)); err != nil {
			fmt.Fprintf(os.Stderr, "Error on creating a new vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
	}

	log.Println("Succeded to bind the port")

	return nil
}

func unbinding(portUuid, hostUuid string) error {
	log.Println("unbinding port " + portUuid)

	conn, _, err := connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	lock := zk.NewLock(conn, GetLockPath(lockKey), zk.WorldACL(zk.PermAll))
	if err = lock.Lock(); err != nil {
		return err
	}
	defer lock.Unlock()

	portPath := GetPortPath(portUuid)
	var data []byte
	if data, _, err = conn.Get(portPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on getting  port %s: %s\n",
			portPath, err.Error())
		return err
	}
	port := &WrappedPort{}
	if err = json.Unmarshal(data, port); err != nil {
		fmt.Fprintf(os.Stderr, "Error on deserializing port %s: %s\n",
			portPath, err.Error())
		return err
	}

	if port.Data.HostId != hostUuid {
		return errors.New("The given host ID didn't match with one in NSDB")
	}
	port.Data.InterfaceName = ""

	updatedPort, err := json.Marshal(port)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on serializing port %s: %s\n",
			portPath, err.Error())
		return err
	}

	if _, err = conn.Set(portPath, updatedPort, -1); err != nil {
		return err
	}

	vrnMappingPath := GetVrnMappingPath(hostUuid, portUuid)
	var exists bool
	if exists, _, err = conn.Exists(vrnMappingPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on examining vrnMapping %s: %s\n",
			vrnMappingPath, err.Error())
		return err
	}
	if exists {
		if err = conn.Delete(vrnMappingPath, -1); err != nil {
			fmt.Fprintf(os.Stderr, "Error on deleging vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
	}
	log.Println("Succeded to unbind the port")

	return nil
}
