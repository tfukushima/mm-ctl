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
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	curator "github.com/flier/curator.go"
	recipes "github.com/flier/curator.go/recipes"
	"github.com/golang/protobuf/proto"
	models "github.com/tfukushima/mm-ctl/org_midonet_cluster_models"
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

func convertToProtobufUuid(original string) (*models.UUID, error) {
	b := make([]byte, 16)
	var s1, s2, s3, s4 string
	fmt.Sscanf(original, "%8x-%4x-%4x-%4x-%12x", &b, &s1, &s2, &s3, &s4)
	b = append(b, s1...)
	b = append(b, s2...)
	b = append(b, s3...)
	b = append(b, s4...)
	msb := binary.BigEndian.Uint64(b[:8])
	lsb := binary.BigEndian.Uint64(b[8:])
	u := models.UUID{}
	u.Msb = &msb
	u.Lsb = &lsb
	return &u, nil
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
		fmt.Fprintf(os.Stderr, "Error on getting port: %s\n", err.Error())
		return err
	}
	n := len(data)
	protoText := string(data[:n])
	port := &models.Port{}
	if err = proto.UnmarshalText(protoText, port); err != nil {
		fmt.Fprintf(os.Stderr, "Error on deserializing port: %s\n", err.Error())
		return err
	}

	protoUuid, err := convertToProtobufUuid(hostUuid)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on serializing host UUID: %s\n", err.Error())
		return err
	}
	port.HostId = protoUuid
	port.InterfaceName = &interfaceName

	updatedPortString := proto.MarshalTextString(port)
	if _, err = client.SetData().ForPathWithData(portPath, []byte(updatedPortString)); err != nil {
		fmt.Fprintf(os.Stderr, "Error on setting port: %s\n", err.Error())
		return err
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
	if _, err := lock.Acquire(); err != nil {
		fmt.Fprintf(os.Stderr, "Error on locking: %s\n", err.Error())
		return err
	}
	defer lock.Release()

	portPath := GetPortPath(portUuid)
	var data []byte
	if data, err = client.GetData().ForPath(portPath); err != nil {
		fmt.Fprintf(os.Stderr, "Error on getting  port: %s\n", err.Error())
		return err
	}

	n := len(data)
	protoText := string(data[:n])
	port := &models.Port{}
	if err = proto.UnmarshalText(protoText, port); err != nil {
		fmt.Fprintf(os.Stderr, "Error on deserializing port: %s\n", err.Error())
		return err
	}

	port.InterfaceName = nil

	updatedPortString := proto.MarshalTextString(port)
	if _, err = client.SetData().ForPathWithData(portPath, []byte(updatedPortString)); err != nil {
		return err
	}

	log.Println("Succeded to unbind the port")

	return nil
}
