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
	"time"

	curator "github.com/flier/curator.go"
	recipes "github.com/flier/curator.go/recipes"
	"github.com/samuel/go-zookeeper/zk"
)

const (
	// The base time to sleep for curator.go connection.
	BaseSleepTime = time.Second
	// The maximum number of the retry for curator.go.
	MaxRetries = 3
	// The maximum time to sleep for curator.go connection.
	MaxSleep = 15 * time.Second
)

// The address of the NSDB.
var zookeeperAddresses = flag.String("zookeeper_hosts", "127.0.0.1:2181",
	"The Addresses of ZooKeeper nodes separated by commas")

// The timout for the NSDB session in seconds.
const sessionTimeoutSec = 10

// The key for the ZOOM topology lock path.
const lockKey = "zoom-topology"

func newClient() curator.CuratorFramework {
	retryPolicy := curator.NewExponentialBackoffRetry(
		BaseSleepTime, MaxRetries, MaxSleep)
	client := curator.NewClient(*zookeeperAddresses, retryPolicy)
	client.Start()

	return client

}

func binding(portUuid, hostUuid, interfaceName string) error {
	log.Println("binding port " + portUuid + " to " + interfaceName)
	client := newClient()
	defer client.Close()

	lock, err := recipes.NewInterProcessMutex(client, GetLockPath(lockKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on instantiating a lock: %s\n", err.Error())
		return err
	}

	if _, err := lock.Acquire(); err != nil {
		fmt.Fprintf(os.Stderr, "Error on locking NSDB: %s\n", err.Error())
		return err
	}
	defer lock.Release()

	portPath := GetPortPath(portUuid)
	var data []byte
	if data, err = client.GetData().ForPath(portPath); err != nil {
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

	if _, err = client.SetData().ForPathWithData(portPath, []byte(updatedPort)); err != nil {
		fmt.Fprintf(os.Stderr, "Error on setting port %s: %s\n",
			portPath, err.Error())
		return err
	}

	vrnMappingPath := GetVrnMappingPath(hostUuid, portUuid)
	var exists bool
	var stat *zk.Stat
	if stat, err = client.CheckExists().ForPath(vrnMappingPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on examining vrnMapping %s: %s\n",
			vrnMappingPath, err.Error())
		return err
	}
	if stat != nil {
		exists = true
	} else {
		exists = false
	}
	var vrnMappingData []byte
	vrnMapping := &WrappedVrnMapping{}
	if exists {
		if vrnMappingData, err = client.GetData().ForPath(vrnMappingPath); err != nil {
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
		if _, err = client.SetData().ForPathWithData(vrnMappingPath, updatedVrnMapping); err != nil {
			fmt.Fprintf(os.Stderr, "Error on setting vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
	} else {
		if _, err = client.Create().WithACL(zk.WorldACL(zk.PermAll)...).ForPathWithData(vrnMappingPath, updatedVrnMapping); err != nil {
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

	client := newClient()
	defer client.Close()

	lock, err := recipes.NewInterProcessMutex(client, GetLockPath(lockKey))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on instantiating a lock: %s\n", err.Error())
		return err
	}
	defer lock.Release()

	portPath := GetPortPath(portUuid)
	var data []byte
	if data, err = client.GetData().ForPath(portPath); err != nil {
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

	if _, err = client.SetData().ForPathWithData(portPath, updatedPort); err != nil {
		return err
	}

	vrnMappingPath := GetVrnMappingPath(hostUuid, portUuid)
	var exists bool
	var stat *zk.Stat
	if stat, err = client.CheckExists().ForPath(vrnMappingPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on examining vrnMapping %s: %s\n",
			vrnMappingPath, err.Error())
		return err
	}
	if stat != nil {
		exists = true
	} else {
		exists = false
	}
	if exists {
		if err = client.Delete().ForPath(vrnMappingPath); err != nil {
			fmt.Fprintf(os.Stderr, "Error on deleging vrnMapping %s: %s\n",
				vrnMappingPath, err.Error())
			return err
		}
	}
	log.Println("Succeded to unbind the port")

	return nil
}
