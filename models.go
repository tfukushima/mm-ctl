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

// The MidoNet Port model store in NSDB.
type Port struct {
	Type           string            `json:"type"`
	Properties     map[string]string `json:"properties"`
	DeviceId       string            `json:"device_id"`
	AdminStateUp   bool              `json:"adminStateUp"`
	InboundFilter  string            `json:"inboundFilter"`
	OutboundFilter string            `json:"outboundFilter"`
	PortGroupIDs   []string          `json:"portGroupIDs"`
	TunnelKey      int               `json:"tunnelKey"`
	HostId         string            `json:"hostId"`
	InterfaceName  string            `json:"interfaceName"`
	PeerId         *string           `json:"peerId"`
	V1ApiType      string            `json:"v1ApiType"`
	VlanId         *int              `json:"vlanId"`
}

// The wrapped Port model in the JSON representation.
type WrappedPort struct {
	Data    *Port  `json:"data"`
	Version string `json:"version"`
}

// The MidoNet VrnMapping model stored in NSDB.
type VrnMapping struct {
	VirtualPortId   string `json:"virtualPortId"`
	LocalDeviceName string `json:"localDeviceName"`
}

// The wrapped VrnMapping model in the JSON representation.
type WrappedVrnMapping struct {
	Data    *VrnMapping `json:"data"`
	Version string      `json:"version"`
}
