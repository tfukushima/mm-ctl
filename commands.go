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
	"fmt"
	"os"

	"github.com/mitchellh/cli"
	ini "gopkg.in/ini.v1"
)

var hostUuid = flag.String("host_uuid", "", "The UUID of the host.")

// the name of the command.
const CommandName = "mm-ctl"

func loadConfig() error {
	midolmanConfigPath := flag.String("config", "/etc/midolman/midolman.conf",
		"The location of the Midolamn config file")
	hostConfigPath := flag.String("host-config", "/etc/midonet_host_id.properties",
		"The location of the host UUID config file")
	flag.Parse()

	cfg, err := ini.Load(*midolmanConfigPath, *hostConfigPath)
	if err != nil {
		return err
	}
	zkSection, err := cfg.GetSection("zookeeper")
	if err != nil {
		return err
	}
	*rootPath = zkSection.Key("root_key").String()
	*zookeeperAddresses = zkSection.Key("zookeeper_hosts").String()
	*hostUuid = cfg.Section("").Key("host_uuid").String()

	return nil
}

// The "binding" command.
type Binding struct{}

func (b *Binding) Help() string {
	return fmt.Sprintf(`Usage: %s bind c5951b28-0d72-41fa-9f29-650bbeae2ed3 8a1c0540-veth"

binding takes the Neutron port UUID and the interface name and binds the port to
the interface.
`, CommandName)
}

func (b *Binding) Run(args []string) int {
	bindCmdFlag := flag.NewFlagSet("bind", flag.ContinueOnError)
	if *hostUuid == "" {
		bindCmdFlag.StringVar(hostUuid, "host_uuid", "", "The UUID of the host (required).")
	}
	// bindCmdFlag.StringVar(nsdbAddresses, "zookeeper_hosts", "127.0.0.1:2181",
	// 	"The Addresses of ZooKeeper nodes separated by commas")
	// bindCmdFlag.StringVar(rootPath, "root_key", "/midonet/v1", "The root path of the NSDB")

	bindCmdFlag.Parse(args)

	if bindCmdFlag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "bind takes exactly 2 args but %d args are given.\n", bindCmdFlag.NArg())
		return 1
	}
	portUuid := bindCmdFlag.Arg(0)
	if *hostUuid == "" {
		fmt.Fprintf(os.Stderr, "Host UUID is required.\n")
		return 1
	}
	interfaceName := bindCmdFlag.Arg(1)

	if err := binding(portUuid, *hostUuid, interfaceName); err != nil {
		fmt.Fprintf(os.Stderr, "Error on binding the port: %s\n", err.Error())
		return 1
	}

	return 0
}

func (b *Binding) Synopsis() string {
	return "bind <port-uuid> <interface-name>"
}

// The "unbinding" command.
type Unbinding struct{}

func (u *Unbinding) Help() string {
	return fmt.Sprintf(`Usage: %s unbind c5951b28-0d72-41fa-9f29-650bbeae2ed3

unbinding takes the Neutron port UUID and unbinds the port.
`, CommandName)
}

func (u *Unbinding) Run(args []string) int {
	unbindCmdFlag := flag.NewFlagSet("unbind", flag.ContinueOnError)
	if *hostUuid == "" {
		unbindCmdFlag.StringVar(hostUuid, "host_uuid", "", "The UUID of the host (required).")
	}

	unbindCmdFlag.Parse(args)

	if unbindCmdFlag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "unbind takes exactly 1 arg but %d args are given.\n", unbindCmdFlag.NArg())
		return 1
	}
	portUuid := unbindCmdFlag.Arg(0)

	if err := unbinding(portUuid, *hostUuid); err != nil {
		fmt.Fprintf(os.Stderr, "Error on unbinding the port: %s\n", err.Error())
		return 1
	}

	return 0
}

func (u *Unbinding) Synopsis() string {
	return "unbind <port-uuid>"
}

func bindingCommand() (cli.Command, error) {
	return &Binding{}, nil
}

func unbindingCommand() (cli.Command, error) {
	return &Unbinding{}, nil
}
