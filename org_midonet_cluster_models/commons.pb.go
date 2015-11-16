// Code generated by protoc-gen-go.
// source: commons.proto
// DO NOT EDIT!

/*
Package org_midonet_cluster_models is a generated protocol buffer package.

It is generated from these files:
	commons.proto
	topology.proto

It has these top-level messages:
	UUID
	IPAddress
	IPSubnet
	Int32Range
	Condition
	BgpNetwork
	BgpPeer
	Chain
	Dhcp
	DhcpV6
	HealthMonitor
	Host
	IPAddrGroup
	L2Insertion
	LoadBalancer
	Mirror
	Network
	Pool
	PoolMember
	Port
	PortGroup
	Route
	Router
	Rule
	TraceRequest
	TunnelZone
	Vip
	Vtep
*/
package org_midonet_cluster_models

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type IPVersion int32

const (
	IPVersion_V4 IPVersion = 1
	IPVersion_V6 IPVersion = 2
)

var IPVersion_name = map[int32]string{
	1: "V4",
	2: "V6",
}
var IPVersion_value = map[string]int32{
	"V4": 1,
	"V6": 2,
}

func (x IPVersion) Enum() *IPVersion {
	p := new(IPVersion)
	*p = x
	return p
}
func (x IPVersion) String() string {
	return proto.EnumName(IPVersion_name, int32(x))
}
func (x *IPVersion) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(IPVersion_value, data, "IPVersion")
	if err != nil {
		return err
	}
	*x = IPVersion(value)
	return nil
}
func (IPVersion) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RuleDirection int32

const (
	RuleDirection_EGRESS  RuleDirection = 0
	RuleDirection_INGRESS RuleDirection = 1
)

var RuleDirection_name = map[int32]string{
	0: "EGRESS",
	1: "INGRESS",
}
var RuleDirection_value = map[string]int32{
	"EGRESS":  0,
	"INGRESS": 1,
}

func (x RuleDirection) Enum() *RuleDirection {
	p := new(RuleDirection)
	*p = x
	return p
}
func (x RuleDirection) String() string {
	return proto.EnumName(RuleDirection_name, int32(x))
}
func (x *RuleDirection) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RuleDirection_value, data, "RuleDirection")
	if err != nil {
		return err
	}
	*x = RuleDirection(value)
	return nil
}
func (RuleDirection) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type EtherType int32

const (
	EtherType_ARP  EtherType = 2054
	EtherType_IPV4 EtherType = 2048
	EtherType_IPV6 EtherType = 34525
)

var EtherType_name = map[int32]string{
	2054:  "ARP",
	2048:  "IPV4",
	34525: "IPV6",
}
var EtherType_value = map[string]int32{
	"ARP":  2054,
	"IPV4": 2048,
	"IPV6": 34525,
}

func (x EtherType) Enum() *EtherType {
	p := new(EtherType)
	*p = x
	return p
}
func (x EtherType) String() string {
	return proto.EnumName(EtherType_name, int32(x))
}
func (x *EtherType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EtherType_value, data, "EtherType")
	if err != nil {
		return err
	}
	*x = EtherType(value)
	return nil
}
func (EtherType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Protocol int32

const (
	Protocol_TCP    Protocol = 6
	Protocol_UDP    Protocol = 17
	Protocol_ICMP   Protocol = 1
	Protocol_ICMPV6 Protocol = 58
)

var Protocol_name = map[int32]string{
	6:  "TCP",
	17: "UDP",
	1:  "ICMP",
	58: "ICMPV6",
}
var Protocol_value = map[string]int32{
	"TCP":    6,
	"UDP":    17,
	"ICMP":   1,
	"ICMPV6": 58,
}

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}
func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}
func (x *Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Protocol_value, data, "Protocol")
	if err != nil {
		return err
	}
	*x = Protocol(value)
	return nil
}
func (Protocol) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type LBStatus int32

const (
	LBStatus_ACTIVE   LBStatus = 1
	LBStatus_INACTIVE LBStatus = 2
)

var LBStatus_name = map[int32]string{
	1: "ACTIVE",
	2: "INACTIVE",
}
var LBStatus_value = map[string]int32{
	"ACTIVE":   1,
	"INACTIVE": 2,
}

func (x LBStatus) Enum() *LBStatus {
	p := new(LBStatus)
	*p = x
	return p
}
func (x LBStatus) String() string {
	return proto.EnumName(LBStatus_name, int32(x))
}
func (x *LBStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LBStatus_value, data, "LBStatus")
	if err != nil {
		return err
	}
	*x = LBStatus(value)
	return nil
}
func (LBStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Condition_FragmentPolicy int32

const (
	Condition_ANY          Condition_FragmentPolicy = 1
	Condition_NONHEADER    Condition_FragmentPolicy = 2
	Condition_HEADER       Condition_FragmentPolicy = 3
	Condition_UNFRAGMENTED Condition_FragmentPolicy = 4
)

var Condition_FragmentPolicy_name = map[int32]string{
	1: "ANY",
	2: "NONHEADER",
	3: "HEADER",
	4: "UNFRAGMENTED",
}
var Condition_FragmentPolicy_value = map[string]int32{
	"ANY":          1,
	"NONHEADER":    2,
	"HEADER":       3,
	"UNFRAGMENTED": 4,
}

func (x Condition_FragmentPolicy) Enum() *Condition_FragmentPolicy {
	p := new(Condition_FragmentPolicy)
	*p = x
	return p
}
func (x Condition_FragmentPolicy) String() string {
	return proto.EnumName(Condition_FragmentPolicy_name, int32(x))
}
func (x *Condition_FragmentPolicy) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Condition_FragmentPolicy_value, data, "Condition_FragmentPolicy")
	if err != nil {
		return err
	}
	*x = Condition_FragmentPolicy(value)
	return nil
}
func (Condition_FragmentPolicy) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

type UUID struct {
	Msb              *uint64 `protobuf:"varint,1,req,name=msb" json:"msb,omitempty"`
	Lsb              *uint64 `protobuf:"varint,2,req,name=lsb" json:"lsb,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *UUID) Reset()                    { *m = UUID{} }
func (m *UUID) String() string            { return proto.CompactTextString(m) }
func (*UUID) ProtoMessage()               {}
func (*UUID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *UUID) GetMsb() uint64 {
	if m != nil && m.Msb != nil {
		return *m.Msb
	}
	return 0
}

func (m *UUID) GetLsb() uint64 {
	if m != nil && m.Lsb != nil {
		return *m.Lsb
	}
	return 0
}

type IPAddress struct {
	Version          *IPVersion `protobuf:"varint,1,req,name=version,enum=org.midonet.cluster.models.IPVersion" json:"version,omitempty"`
	Address          *string    `protobuf:"bytes,2,req,name=address" json:"address,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *IPAddress) Reset()                    { *m = IPAddress{} }
func (m *IPAddress) String() string            { return proto.CompactTextString(m) }
func (*IPAddress) ProtoMessage()               {}
func (*IPAddress) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IPAddress) GetVersion() IPVersion {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return IPVersion_V4
}

func (m *IPAddress) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

type IPSubnet struct {
	Version          *IPVersion `protobuf:"varint,1,req,name=version,enum=org.midonet.cluster.models.IPVersion" json:"version,omitempty"`
	Address          *string    `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
	PrefixLength     *uint32    `protobuf:"varint,3,opt,name=prefix_length" json:"prefix_length,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *IPSubnet) Reset()                    { *m = IPSubnet{} }
func (m *IPSubnet) String() string            { return proto.CompactTextString(m) }
func (*IPSubnet) ProtoMessage()               {}
func (*IPSubnet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *IPSubnet) GetVersion() IPVersion {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return IPVersion_V4
}

func (m *IPSubnet) GetAddress() string {
	if m != nil && m.Address != nil {
		return *m.Address
	}
	return ""
}

func (m *IPSubnet) GetPrefixLength() uint32 {
	if m != nil && m.PrefixLength != nil {
		return *m.PrefixLength
	}
	return 0
}

type Int32Range struct {
	Start            *int32 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End              *int32 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Int32Range) Reset()                    { *m = Int32Range{} }
func (m *Int32Range) String() string            { return proto.CompactTextString(m) }
func (*Int32Range) ProtoMessage()               {}
func (*Int32Range) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Int32Range) GetStart() int32 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *Int32Range) GetEnd() int32 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

type Condition struct {
	// Condition fields
	ConjunctionInv      *bool                     `protobuf:"varint,1,opt,name=conjunction_inv" json:"conjunction_inv,omitempty"`
	MatchForwardFlow    *bool                     `protobuf:"varint,2,opt,name=match_forward_flow" json:"match_forward_flow,omitempty"`
	MatchReturnFlow     *bool                     `protobuf:"varint,3,opt,name=match_return_flow" json:"match_return_flow,omitempty"`
	InPortIds           []*UUID                   `protobuf:"bytes,4,rep,name=in_port_ids" json:"in_port_ids,omitempty"`
	InPortInv           *bool                     `protobuf:"varint,5,opt,name=in_port_inv" json:"in_port_inv,omitempty"`
	OutPortIds          []*UUID                   `protobuf:"bytes,6,rep,name=out_port_ids" json:"out_port_ids,omitempty"`
	OutPortInv          *bool                     `protobuf:"varint,7,opt,name=out_port_inv" json:"out_port_inv,omitempty"`
	PortGroupId         *UUID                     `protobuf:"bytes,8,opt,name=port_group_id" json:"port_group_id,omitempty"`
	InvPortGroup        *bool                     `protobuf:"varint,9,opt,name=inv_port_group" json:"inv_port_group,omitempty"`
	IpAddrGroupIdSrc    *UUID                     `protobuf:"bytes,10,opt,name=ip_addr_group_id_src" json:"ip_addr_group_id_src,omitempty"`
	InvIpAddrGroupIdSrc *bool                     `protobuf:"varint,11,opt,name=inv_ip_addr_group_id_src" json:"inv_ip_addr_group_id_src,omitempty"`
	IpAddrGroupIdDst    *UUID                     `protobuf:"bytes,12,opt,name=ip_addr_group_id_dst" json:"ip_addr_group_id_dst,omitempty"`
	InvIpAddrGroupIdDst *bool                     `protobuf:"varint,13,opt,name=inv_ip_addr_group_id_dst" json:"inv_ip_addr_group_id_dst,omitempty"`
	DlType              *int32                    `protobuf:"varint,14,opt,name=dl_type" json:"dl_type,omitempty"`
	InvDlType           *bool                     `protobuf:"varint,15,opt,name=inv_dl_type" json:"inv_dl_type,omitempty"`
	DlSrc               *string                   `protobuf:"bytes,16,opt,name=dl_src" json:"dl_src,omitempty"`
	DlSrcMask           *int64                    `protobuf:"varint,17,opt,name=dl_src_mask" json:"dl_src_mask,omitempty"`
	InvDlSrc            *bool                     `protobuf:"varint,18,opt,name=inv_dl_src" json:"inv_dl_src,omitempty"`
	DlDst               *string                   `protobuf:"bytes,19,opt,name=dl_dst" json:"dl_dst,omitempty"`
	DlDstMask           *int64                    `protobuf:"varint,20,opt,name=dl_dst_mask" json:"dl_dst_mask,omitempty"`
	InvDlDst            *bool                     `protobuf:"varint,21,opt,name=inv_dl_dst" json:"inv_dl_dst,omitempty"`
	NwTos               *int32                    `protobuf:"varint,22,opt,name=nw_tos" json:"nw_tos,omitempty"`
	NwTosInv            *bool                     `protobuf:"varint,23,opt,name=nw_tos_inv" json:"nw_tos_inv,omitempty"`
	NwProto             *int32                    `protobuf:"varint,24,opt,name=nw_proto" json:"nw_proto,omitempty"`
	NwProtoInv          *bool                     `protobuf:"varint,25,opt,name=nw_proto_inv" json:"nw_proto_inv,omitempty"`
	NwSrcIp             *IPSubnet                 `protobuf:"bytes,26,opt,name=nw_src_ip" json:"nw_src_ip,omitempty"`
	NwDstIp             *IPSubnet                 `protobuf:"bytes,27,opt,name=nw_dst_ip" json:"nw_dst_ip,omitempty"`
	TpSrc               *Int32Range               `protobuf:"bytes,28,opt,name=tp_src" json:"tp_src,omitempty"`
	TpDst               *Int32Range               `protobuf:"bytes,29,opt,name=tp_dst" json:"tp_dst,omitempty"`
	NwSrcInv            *bool                     `protobuf:"varint,30,opt,name=nw_src_inv" json:"nw_src_inv,omitempty"`
	NwDstInv            *bool                     `protobuf:"varint,31,opt,name=nw_dst_inv" json:"nw_dst_inv,omitempty"`
	TpSrcInv            *bool                     `protobuf:"varint,32,opt,name=tp_src_inv" json:"tp_src_inv,omitempty"`
	TpDstInv            *bool                     `protobuf:"varint,33,opt,name=tp_dst_inv" json:"tp_dst_inv,omitempty"`
	TraversedDevice     *UUID                     `protobuf:"bytes,34,opt,name=traversed_device" json:"traversed_device,omitempty"`
	TraversedDeviceInv  *bool                     `protobuf:"varint,35,opt,name=traversed_device_inv" json:"traversed_device_inv,omitempty"`
	FragmentPolicy      *Condition_FragmentPolicy `protobuf:"varint,36,opt,name=fragment_policy,enum=org.midonet.cluster.models.Condition_FragmentPolicy" json:"fragment_policy,omitempty"`
	NoVlan              *bool                     `protobuf:"varint,60,opt,name=no_vlan" json:"no_vlan,omitempty"`
	Vlan                *uint32                   `protobuf:"varint,61,opt,name=vlan" json:"vlan,omitempty"`
	XXX_unrecognized    []byte                    `json:"-"`
}

func (m *Condition) Reset()                    { *m = Condition{} }
func (m *Condition) String() string            { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()               {}
func (*Condition) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Condition) GetConjunctionInv() bool {
	if m != nil && m.ConjunctionInv != nil {
		return *m.ConjunctionInv
	}
	return false
}

func (m *Condition) GetMatchForwardFlow() bool {
	if m != nil && m.MatchForwardFlow != nil {
		return *m.MatchForwardFlow
	}
	return false
}

func (m *Condition) GetMatchReturnFlow() bool {
	if m != nil && m.MatchReturnFlow != nil {
		return *m.MatchReturnFlow
	}
	return false
}

func (m *Condition) GetInPortIds() []*UUID {
	if m != nil {
		return m.InPortIds
	}
	return nil
}

func (m *Condition) GetInPortInv() bool {
	if m != nil && m.InPortInv != nil {
		return *m.InPortInv
	}
	return false
}

func (m *Condition) GetOutPortIds() []*UUID {
	if m != nil {
		return m.OutPortIds
	}
	return nil
}

func (m *Condition) GetOutPortInv() bool {
	if m != nil && m.OutPortInv != nil {
		return *m.OutPortInv
	}
	return false
}

func (m *Condition) GetPortGroupId() *UUID {
	if m != nil {
		return m.PortGroupId
	}
	return nil
}

func (m *Condition) GetInvPortGroup() bool {
	if m != nil && m.InvPortGroup != nil {
		return *m.InvPortGroup
	}
	return false
}

func (m *Condition) GetIpAddrGroupIdSrc() *UUID {
	if m != nil {
		return m.IpAddrGroupIdSrc
	}
	return nil
}

func (m *Condition) GetInvIpAddrGroupIdSrc() bool {
	if m != nil && m.InvIpAddrGroupIdSrc != nil {
		return *m.InvIpAddrGroupIdSrc
	}
	return false
}

func (m *Condition) GetIpAddrGroupIdDst() *UUID {
	if m != nil {
		return m.IpAddrGroupIdDst
	}
	return nil
}

func (m *Condition) GetInvIpAddrGroupIdDst() bool {
	if m != nil && m.InvIpAddrGroupIdDst != nil {
		return *m.InvIpAddrGroupIdDst
	}
	return false
}

func (m *Condition) GetDlType() int32 {
	if m != nil && m.DlType != nil {
		return *m.DlType
	}
	return 0
}

func (m *Condition) GetInvDlType() bool {
	if m != nil && m.InvDlType != nil {
		return *m.InvDlType
	}
	return false
}

func (m *Condition) GetDlSrc() string {
	if m != nil && m.DlSrc != nil {
		return *m.DlSrc
	}
	return ""
}

func (m *Condition) GetDlSrcMask() int64 {
	if m != nil && m.DlSrcMask != nil {
		return *m.DlSrcMask
	}
	return 0
}

func (m *Condition) GetInvDlSrc() bool {
	if m != nil && m.InvDlSrc != nil {
		return *m.InvDlSrc
	}
	return false
}

func (m *Condition) GetDlDst() string {
	if m != nil && m.DlDst != nil {
		return *m.DlDst
	}
	return ""
}

func (m *Condition) GetDlDstMask() int64 {
	if m != nil && m.DlDstMask != nil {
		return *m.DlDstMask
	}
	return 0
}

func (m *Condition) GetInvDlDst() bool {
	if m != nil && m.InvDlDst != nil {
		return *m.InvDlDst
	}
	return false
}

func (m *Condition) GetNwTos() int32 {
	if m != nil && m.NwTos != nil {
		return *m.NwTos
	}
	return 0
}

func (m *Condition) GetNwTosInv() bool {
	if m != nil && m.NwTosInv != nil {
		return *m.NwTosInv
	}
	return false
}

func (m *Condition) GetNwProto() int32 {
	if m != nil && m.NwProto != nil {
		return *m.NwProto
	}
	return 0
}

func (m *Condition) GetNwProtoInv() bool {
	if m != nil && m.NwProtoInv != nil {
		return *m.NwProtoInv
	}
	return false
}

func (m *Condition) GetNwSrcIp() *IPSubnet {
	if m != nil {
		return m.NwSrcIp
	}
	return nil
}

func (m *Condition) GetNwDstIp() *IPSubnet {
	if m != nil {
		return m.NwDstIp
	}
	return nil
}

func (m *Condition) GetTpSrc() *Int32Range {
	if m != nil {
		return m.TpSrc
	}
	return nil
}

func (m *Condition) GetTpDst() *Int32Range {
	if m != nil {
		return m.TpDst
	}
	return nil
}

func (m *Condition) GetNwSrcInv() bool {
	if m != nil && m.NwSrcInv != nil {
		return *m.NwSrcInv
	}
	return false
}

func (m *Condition) GetNwDstInv() bool {
	if m != nil && m.NwDstInv != nil {
		return *m.NwDstInv
	}
	return false
}

func (m *Condition) GetTpSrcInv() bool {
	if m != nil && m.TpSrcInv != nil {
		return *m.TpSrcInv
	}
	return false
}

func (m *Condition) GetTpDstInv() bool {
	if m != nil && m.TpDstInv != nil {
		return *m.TpDstInv
	}
	return false
}

func (m *Condition) GetTraversedDevice() *UUID {
	if m != nil {
		return m.TraversedDevice
	}
	return nil
}

func (m *Condition) GetTraversedDeviceInv() bool {
	if m != nil && m.TraversedDeviceInv != nil {
		return *m.TraversedDeviceInv
	}
	return false
}

func (m *Condition) GetFragmentPolicy() Condition_FragmentPolicy {
	if m != nil && m.FragmentPolicy != nil {
		return *m.FragmentPolicy
	}
	return Condition_ANY
}

func (m *Condition) GetNoVlan() bool {
	if m != nil && m.NoVlan != nil {
		return *m.NoVlan
	}
	return false
}

func (m *Condition) GetVlan() uint32 {
	if m != nil && m.Vlan != nil {
		return *m.Vlan
	}
	return 0
}

func init() {
	proto.RegisterType((*UUID)(nil), "org.midonet.cluster.models.UUID")
	proto.RegisterType((*IPAddress)(nil), "org.midonet.cluster.models.IPAddress")
	proto.RegisterType((*IPSubnet)(nil), "org.midonet.cluster.models.IPSubnet")
	proto.RegisterType((*Int32Range)(nil), "org.midonet.cluster.models.Int32Range")
	proto.RegisterType((*Condition)(nil), "org.midonet.cluster.models.Condition")
	proto.RegisterEnum("org.midonet.cluster.models.IPVersion", IPVersion_name, IPVersion_value)
	proto.RegisterEnum("org.midonet.cluster.models.RuleDirection", RuleDirection_name, RuleDirection_value)
	proto.RegisterEnum("org.midonet.cluster.models.EtherType", EtherType_name, EtherType_value)
	proto.RegisterEnum("org.midonet.cluster.models.Protocol", Protocol_name, Protocol_value)
	proto.RegisterEnum("org.midonet.cluster.models.LBStatus", LBStatus_name, LBStatus_value)
	proto.RegisterEnum("org.midonet.cluster.models.Condition_FragmentPolicy", Condition_FragmentPolicy_name, Condition_FragmentPolicy_value)
}

var fileDescriptor0 = []byte{
	// 853 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x95, 0xdf, 0x6e, 0x1a, 0x47,
	0x14, 0xc6, 0xcb, 0x1f, 0x03, 0x7b, 0x0c, 0x78, 0xbd, 0x71, 0x92, 0x89, 0x93, 0xb6, 0x94, 0x3a,
	0x15, 0x42, 0x15, 0x17, 0xd4, 0xa5, 0x52, 0xd4, 0x56, 0xc2, 0x80, 0x53, 0xa4, 0x9a, 0xae, 0x30,
	0xb6, 0xd4, 0xab, 0xd5, 0x66, 0x77, 0x8c, 0x37, 0x5d, 0x66, 0xd0, 0xec, 0x80, 0x9b, 0xbb, 0x5e,
	0xf5, 0xae, 0xef, 0xd2, 0x17, 0xea, 0xbb, 0xf4, 0xcc, 0x59, 0x88, 0x91, 0xd5, 0x20, 0x5b, 0xcd,
	0xd5, 0xee, 0xf9, 0xf6, 0x7c, 0xbf, 0x39, 0x33, 0x73, 0x66, 0x16, 0x2a, 0x81, 0x9c, 0xcd, 0xa4,
	0x48, 0x5a, 0x73, 0x25, 0xb5, 0x74, 0x0e, 0xa5, 0x9a, 0xb6, 0x66, 0x51, 0x28, 0x05, 0xd7, 0xad,
	0x20, 0x5e, 0x24, 0x9a, 0xab, 0xd6, 0x4c, 0x86, 0x3c, 0x4e, 0xea, 0x35, 0xc8, 0x5f, 0x5c, 0x0c,
	0xfb, 0xce, 0x2e, 0xe4, 0x66, 0xc9, 0x1b, 0x96, 0xa9, 0x65, 0x1b, 0x79, 0x13, 0xc4, 0x18, 0x64,
	0x4d, 0x50, 0x9f, 0x80, 0x35, 0x74, 0xbb, 0x61, 0xa8, 0x78, 0x92, 0x38, 0x1d, 0x28, 0x2e, 0xb9,
	0x4a, 0x22, 0x29, 0x28, 0xb5, 0xda, 0x7e, 0xd9, 0xfa, 0x30, 0xbc, 0x35, 0x74, 0x2f, 0xd3, 0x64,
	0x67, 0x0f, 0x8a, 0x7e, 0x8a, 0x20, 0xaa, 0x55, 0x7f, 0x0b, 0xa5, 0xa1, 0x7b, 0xbe, 0x78, 0x83,
	0xae, 0x8f, 0x03, 0xcd, 0x34, 0x2c, 0xe7, 0x31, 0x54, 0xe6, 0x8a, 0x5f, 0x45, 0xbf, 0x7b, 0x31,
	0x17, 0x53, 0x7d, 0xcd, 0x72, 0x28, 0x57, 0xea, 0x0d, 0x80, 0xa1, 0xd0, 0xdf, 0xb4, 0xc7, 0xbe,
	0x98, 0x72, 0xa7, 0x02, 0x3b, 0x89, 0xf6, 0x95, 0xc6, 0xb1, 0x32, 0x8d, 0x1d, 0x33, 0x57, 0x2e,
	0x42, 0x02, 0xec, 0xd4, 0xff, 0x06, 0xb0, 0x7a, 0x52, 0x84, 0x91, 0x36, 0xfc, 0xa7, 0xb0, 0x17,
	0x48, 0xf1, 0x76, 0x21, 0x02, 0x13, 0x7a, 0x91, 0x58, 0x92, 0xa7, 0xe4, 0x1c, 0x82, 0x33, 0xf3,
	0x75, 0x70, 0xed, 0x5d, 0x49, 0x75, 0xe3, 0xab, 0xd0, 0xbb, 0x8a, 0xe5, 0x0d, 0x21, 0x4a, 0xce,
	0x33, 0xd8, 0x4f, 0xbf, 0x29, 0xae, 0x17, 0x4a, 0xa4, 0x9f, 0x72, 0xf4, 0xe9, 0x5b, 0xd8, 0x8d,
	0x84, 0x37, 0x97, 0x4a, 0x7b, 0x51, 0x98, 0xb0, 0x7c, 0x2d, 0xd7, 0xd8, 0x6d, 0xd7, 0xb6, 0xcd,
	0x95, 0xb6, 0xe6, 0xd1, 0x86, 0x0d, 0x4b, 0xd8, 0x21, 0x56, 0x07, 0xca, 0x72, 0xa1, 0x6f, 0x61,
	0x85, 0x7b, 0xc2, 0x0e, 0x36, 0x7d, 0x48, 0x2b, 0x12, 0xed, 0x3b, 0x5c, 0x38, 0xa3, 0x4c, 0x95,
	0x5c, 0xcc, 0x91, 0xc7, 0x4a, 0x28, 0xdf, 0x07, 0xf7, 0x04, 0xaa, 0x48, 0xf1, 0x6e, 0xcd, 0xcc,
	0x22, 0xe0, 0x8f, 0x70, 0x10, 0xcd, 0x3d, 0xb3, 0x3b, 0xef, 0x99, 0x5e, 0xa2, 0x02, 0x06, 0xf7,
	0xe4, 0xd6, 0x80, 0x19, 0xee, 0x7f, 0x32, 0x76, 0x3f, 0x38, 0x42, 0x98, 0x68, 0x56, 0xfe, 0x9f,
	0x23, 0x18, 0x46, 0x85, 0x46, 0xc0, 0xf6, 0x0a, 0x63, 0x4f, 0xbf, 0x9b, 0x73, 0x56, 0xa5, 0x56,
	0xa1, 0x8d, 0x58, 0x7a, 0x6b, 0x71, 0x8f, 0xb2, 0xaa, 0x50, 0x40, 0xc1, 0xd4, 0x65, 0x53, 0x0f,
	0x62, 0x52, 0x1a, 0x7b, 0x33, 0x3f, 0xf9, 0x8d, 0xed, 0xa3, 0x98, 0x73, 0x1c, 0x80, 0x95, 0xd3,
	0x24, 0x3a, 0x1b, 0x46, 0x33, 0xdc, 0xa3, 0x0d, 0x23, 0xc6, 0xa9, 0xf1, 0xe0, 0x8e, 0xd1, 0x24,
	0x3e, 0x5e, 0x1b, 0xc5, 0x8d, 0xa7, 0x65, 0xc2, 0x9e, 0x50, 0x59, 0x98, 0x93, 0xc6, 0xb4, 0xa1,
	0x4f, 0x29, 0xc7, 0x86, 0x12, 0x6a, 0x74, 0xfc, 0x19, 0xa3, 0x2c, 0xdc, 0xf8, 0xb5, 0x42, 0x79,
	0xcf, 0x56, 0x1b, 0x6f, 0xa1, 0x6a, 0xaa, 0x8d, 0xe6, 0xec, 0x90, 0x96, 0xee, 0x68, 0xfb, 0xe1,
	0x5b, 0x9d, 0xd9, 0xd4, 0x68, 0xaa, 0x45, 0xe3, 0xf3, 0x07, 0x18, 0x3b, 0x50, 0xd0, 0x73, 0x5a,
	0x86, 0x17, 0xe4, 0xfa, 0x6a, 0xab, 0xeb, 0xf6, 0xd8, 0xa6, 0x3e, 0xb3, 0x0a, 0x9f, 0x3e, 0xc8,
	0x97, 0xae, 0x0e, 0xcd, 0x10, 0x67, 0xfd, 0x19, 0xcd, 0x3a, 0xd5, 0xa8, 0x78, 0xd4, 0x3e, 0x5f,
	0x6b, 0x69, 0x5d, 0xa4, 0xd5, 0x36, 0xb4, 0x75, 0xde, 0x17, 0xa4, 0xbd, 0x02, 0x5b, 0x2b, 0xdf,
	0xdc, 0x57, 0x1c, 0x9b, 0x85, 0x2f, 0xa3, 0x80, 0xb3, 0xfa, 0x3d, 0x7b, 0xee, 0x05, 0x1c, 0xdc,
	0xf5, 0x12, 0xf9, 0x4b, 0x22, 0x9f, 0xc1, 0xde, 0x95, 0xf2, 0xa7, 0x33, 0x2e, 0xcc, 0xf9, 0x8c,
	0xa3, 0xe0, 0x1d, 0x3b, 0xc2, 0x0f, 0xd5, 0xf6, 0xf1, 0x36, 0xf0, 0xfb, 0xeb, 0xaa, 0x75, 0xba,
	0x32, 0xbb, 0xe4, 0x35, 0xed, 0x2b, 0xa4, 0xb7, 0x8c, 0x7d, 0xc1, 0xbe, 0x27, 0x7e, 0x19, 0xf2,
	0x14, 0xfd, 0x40, 0x97, 0xe2, 0x29, 0x54, 0xef, 0x18, 0x8a, 0x90, 0xeb, 0x8e, 0x7e, 0xb5, 0x33,
	0x78, 0x43, 0x5a, 0xa3, 0x5f, 0x46, 0x3f, 0x0d, 0xba, 0xfd, 0xc1, 0xd8, 0xce, 0xe2, 0x22, 0x14,
	0x56, 0xef, 0x39, 0xec, 0xab, 0xf2, 0xc5, 0xe8, 0x74, 0xdc, 0x7d, 0x7d, 0x36, 0x18, 0x4d, 0x06,
	0x7d, 0x3b, 0xdf, 0x7c, 0x6e, 0x7e, 0x0f, 0xeb, 0x1b, 0xb9, 0x00, 0xd9, 0xcb, 0x63, 0x24, 0x98,
	0x67, 0xc7, 0xce, 0x36, 0x1b, 0x50, 0x19, 0x2f, 0x62, 0xde, 0x8f, 0x14, 0xa7, 0x3b, 0xd4, 0xb0,
	0x06, 0xaf, 0xc7, 0x83, 0xf3, 0x73, 0xfb, 0x13, 0xbc, 0x79, 0x8b, 0xc3, 0x51, 0x1a, 0x64, 0x9a,
	0x5f, 0x83, 0x35, 0xd0, 0xd7, 0x5c, 0x4d, 0xf0, 0x64, 0x39, 0x25, 0xac, 0x64, 0xec, 0xda, 0x7f,
	0xda, 0x8e, 0x05, 0x79, 0xa4, 0x1f, 0xdb, 0x7f, 0xd8, 0x68, 0x35, 0xaf, 0x1d, 0xfb, 0x9f, 0xbf,
	0xb2, 0xcd, 0x36, 0x94, 0x5c, 0xd3, 0xc9, 0x81, 0x8c, 0x4d, 0xd9, 0x93, 0x9e, 0x6b, 0x17, 0xcc,
	0xcb, 0x45, 0xdf, 0xb5, 0xf7, 0xd1, 0x9e, 0x1f, 0xf6, 0xce, 0x5c, 0xac, 0x03, 0x87, 0x33, 0x6f,
	0xe8, 0x7a, 0xd5, 0x3c, 0x82, 0xd2, 0xcf, 0x27, 0xe7, 0xda, 0xd7, 0x8b, 0xc4, 0xe8, 0xdd, 0xde,
	0x64, 0x78, 0x39, 0xc0, 0x9c, 0x32, 0xfe, 0x89, 0x46, 0xab, 0x28, 0x7b, 0xf2, 0x12, 0xb6, 0xfc,
	0x2d, 0x4f, 0x8a, 0xbd, 0xf4, 0xc7, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x03, 0x45, 0x8d,
	0x5b, 0x62, 0x07, 0x00, 0x00,
}
