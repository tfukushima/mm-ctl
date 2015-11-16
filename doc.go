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

/*
mm-ctl provides the capability for binding a Neutron port to a virutal
interface and unbinding the Neutron port from the virtual interface.

You can see how to use mm-ctl as follow:

    $ mm-ctl
    usage: mm-ctl [--version] [--help] <command> [<args>]

    Available commands are:
        bind      bind <port-uuid> <interface-name>
        unbind    unbind <port-uuid>

For instance, the following command binds the Neutron port which UUID is
2738e706-6ec7-4b22-b43f-be20f7ad85cb to the virtual interface which name
is b3f9a34d-veth.

    $ mm-ctl bind 2738e706-6ec7-4b22-b43f-be20f7ad85cb b3f9a34d-veth

And the following command unbinds the Neutron port.

    $ mm-ctl unbind 2738e706-6ec7-4b22-b43f-be20f7ad85cb

By default mm-ctl sees the configuration file /etc/midolman/midolman.conf to
lookup the values of zookeeper_hosts and roo_key under zookeeper section in the
config file. /etc/midolman/midolman.conf should follow .ini style and contain
the key-value configurations as follow.

    [zookeeper]
    zookeeper_hosts = "127.0.0.1:2181"
    root_key = /midonet

Or you can specify the location of the configuraton file with -config option as
follow.

    $ mm-ctl -config ./midolman.conf bind 2738e706-6ec7-4b22-b43f-be20f7ad85cb b3f9a34d-veth

mm-ctl also look at /etc/midonet_host_id.properties to retrieve the host UUID.
/etc/midonet_host_id.properties typically contains the key-value configuration
of the host UUID as follow.

    host_uuid=b53b82f3-ebe5-4e8d-9088-8b3ae4ade76e

Or you can specify the location of the configuration file with -host-config
option as follow.

    $ mm-ctl -host-config ./host_uuid.properties bind 2738e706-6ec7-4b22-b43f-be20f7ad85cb b3f9a34d-veth

Please not -config and -host-config should be placed before "bind" or "unbind"
subcommands.
*/
package main
