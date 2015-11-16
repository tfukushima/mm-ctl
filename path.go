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
	"flag"
	"strings"
)

// The base path for MidoNet in the NSDB.
var rootPath = flag.String("root_key", "/midonet/v1", "The root path of the NSDB")

// The path to the locks in the NSDB.
var locksPath = *rootPath + "/locks"

// The path to the ports in the NSDB.
var portsPath = *rootPath + "/ports"

var vrnMappingsPath = *rootPath + "hosts"

// Get the path for the lock which name is given as an argument.
func GetLockPath(lockName string) string {
	return locksPath + "/" + lockName
}

// Get the path for the port which UUID is given as an argument.
func GetPortPath(portUuid string) string {
	return portsPath + "/" + portUuid
}

// Get the path for the vrn mapping for the port which UUID is given as the
// second argument on the host which UUID is given as the first argument.
func GetVrnMappingPath(hostUuid, portUuid string) string {
	return *rootPath + "/hosts/" + hostUuid + "/vrnMappings/ports/" + portUuid
}

// Normalize the path string, which should not end with "/".
func Normalize(path string) string {
	return strings.TrimSuffix(path, "/")
}
