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
)

func runCommand() int {
	// iniflags.Parse()
	if err := loadConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load the config files: %s\n", err.Error())
		return 1
	}

	c := cli.NewCLI(CommandName, Version)
	c.Args = flag.Args()

	c.Commands = map[string]cli.CommandFactory{
		"bind":   bindingCommand,
		"unbind": unbindingCommand,
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err.Error())
		return 1
	}

	return exitStatus
}

func main() {
	os.Exit(runCommand())
}
