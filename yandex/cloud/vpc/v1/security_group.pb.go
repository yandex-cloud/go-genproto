// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/vpc/v1/security_group.proto

package vpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type SecurityGroup_Status int32

const (
	SecurityGroup_STATUS_UNSPECIFIED SecurityGroup_Status = 0
	SecurityGroup_CREATING           SecurityGroup_Status = 1
	SecurityGroup_ACTIVE             SecurityGroup_Status = 2
	// updating is a long operation because we must update all instances in SG
	SecurityGroup_UPDATING SecurityGroup_Status = 3
	SecurityGroup_DELETING SecurityGroup_Status = 4
)

var SecurityGroup_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "CREATING",
	2: "ACTIVE",
	3: "UPDATING",
	4: "DELETING",
}

var SecurityGroup_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"CREATING":           1,
	"ACTIVE":             2,
	"UPDATING":           3,
	"DELETING":           4,
}

func (x SecurityGroup_Status) String() string {
	return proto.EnumName(SecurityGroup_Status_name, int32(x))
}

func (SecurityGroup_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{0, 0}
}

type SecurityGroupRule_Direction int32

const (
	SecurityGroupRule_DIRECTION_UNSPECIFIED SecurityGroupRule_Direction = 0
	SecurityGroupRule_INGRESS               SecurityGroupRule_Direction = 1
	SecurityGroupRule_EGRESS                SecurityGroupRule_Direction = 2
)

var SecurityGroupRule_Direction_name = map[int32]string{
	0: "DIRECTION_UNSPECIFIED",
	1: "INGRESS",
	2: "EGRESS",
}

var SecurityGroupRule_Direction_value = map[string]int32{
	"DIRECTION_UNSPECIFIED": 0,
	"INGRESS":               1,
	"EGRESS":                2,
}

func (x SecurityGroupRule_Direction) String() string {
	return proto.EnumName(SecurityGroupRule_Direction_name, int32(x))
}

func (SecurityGroupRule_Direction) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{1, 0}
}

type SecurityGroup struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId             string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name                 string               `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description          string               `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels               map[string]string    `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NetworkId            string               `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Status               SecurityGroup_Status `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.vpc.v1.SecurityGroup_Status" json:"status,omitempty"`
	Rules                []*SecurityGroupRule `protobuf:"bytes,9,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SecurityGroup) Reset()         { *m = SecurityGroup{} }
func (m *SecurityGroup) String() string { return proto.CompactTextString(m) }
func (*SecurityGroup) ProtoMessage()    {}
func (*SecurityGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{0}
}

func (m *SecurityGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityGroup.Unmarshal(m, b)
}
func (m *SecurityGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityGroup.Marshal(b, m, deterministic)
}
func (m *SecurityGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityGroup.Merge(m, src)
}
func (m *SecurityGroup) XXX_Size() int {
	return xxx_messageInfo_SecurityGroup.Size(m)
}
func (m *SecurityGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityGroup.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityGroup proto.InternalMessageInfo

func (m *SecurityGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SecurityGroup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *SecurityGroup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *SecurityGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SecurityGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SecurityGroup) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *SecurityGroup) GetNetworkId() string {
	if m != nil {
		return m.NetworkId
	}
	return ""
}

func (m *SecurityGroup) GetStatus() SecurityGroup_Status {
	if m != nil {
		return m.Status
	}
	return SecurityGroup_STATUS_UNSPECIFIED
}

func (m *SecurityGroup) GetRules() []*SecurityGroupRule {
	if m != nil {
		return m.Rules
	}
	return nil
}

type SecurityGroupRule struct {
	Id          string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Description string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string           `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Direction   SecurityGroupRule_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=yandex.cloud.vpc.v1.SecurityGroupRule_Direction" json:"direction,omitempty"`
	Ports       *PortRange                  `protobuf:"bytes,5,opt,name=ports,proto3" json:"ports,omitempty"`
	// null value means any protocol
	// values from https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ProtocolName   string `protobuf:"bytes,6,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	ProtocolNumber int64  `protobuf:"varint,7,opt,name=protocol_number,json=protocolNumber,proto3" json:"protocol_number,omitempty"`
	// Types that are valid to be assigned to Target:
	//	*SecurityGroupRule_CidrBlocks
	Target               isSecurityGroupRule_Target `protobuf_oneof:"target"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SecurityGroupRule) Reset()         { *m = SecurityGroupRule{} }
func (m *SecurityGroupRule) String() string { return proto.CompactTextString(m) }
func (*SecurityGroupRule) ProtoMessage()    {}
func (*SecurityGroupRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{1}
}

func (m *SecurityGroupRule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecurityGroupRule.Unmarshal(m, b)
}
func (m *SecurityGroupRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecurityGroupRule.Marshal(b, m, deterministic)
}
func (m *SecurityGroupRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecurityGroupRule.Merge(m, src)
}
func (m *SecurityGroupRule) XXX_Size() int {
	return xxx_messageInfo_SecurityGroupRule.Size(m)
}
func (m *SecurityGroupRule) XXX_DiscardUnknown() {
	xxx_messageInfo_SecurityGroupRule.DiscardUnknown(m)
}

var xxx_messageInfo_SecurityGroupRule proto.InternalMessageInfo

func (m *SecurityGroupRule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SecurityGroupRule) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *SecurityGroupRule) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *SecurityGroupRule) GetDirection() SecurityGroupRule_Direction {
	if m != nil {
		return m.Direction
	}
	return SecurityGroupRule_DIRECTION_UNSPECIFIED
}

func (m *SecurityGroupRule) GetPorts() *PortRange {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *SecurityGroupRule) GetProtocolName() string {
	if m != nil {
		return m.ProtocolName
	}
	return ""
}

func (m *SecurityGroupRule) GetProtocolNumber() int64 {
	if m != nil {
		return m.ProtocolNumber
	}
	return 0
}

type isSecurityGroupRule_Target interface {
	isSecurityGroupRule_Target()
}

type SecurityGroupRule_CidrBlocks struct {
	CidrBlocks *CidrBlocks `protobuf:"bytes,8,opt,name=cidr_blocks,json=cidrBlocks,proto3,oneof"`
}

func (*SecurityGroupRule_CidrBlocks) isSecurityGroupRule_Target() {}

func (m *SecurityGroupRule) GetTarget() isSecurityGroupRule_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *SecurityGroupRule) GetCidrBlocks() *CidrBlocks {
	if x, ok := m.GetTarget().(*SecurityGroupRule_CidrBlocks); ok {
		return x.CidrBlocks
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SecurityGroupRule) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SecurityGroupRule_CidrBlocks)(nil),
	}
}

type PortRange struct {
	FromPort             int64    `protobuf:"varint,1,opt,name=from_port,json=fromPort,proto3" json:"from_port,omitempty"`
	ToPort               int64    `protobuf:"varint,2,opt,name=to_port,json=toPort,proto3" json:"to_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PortRange) Reset()         { *m = PortRange{} }
func (m *PortRange) String() string { return proto.CompactTextString(m) }
func (*PortRange) ProtoMessage()    {}
func (*PortRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{2}
}

func (m *PortRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PortRange.Unmarshal(m, b)
}
func (m *PortRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PortRange.Marshal(b, m, deterministic)
}
func (m *PortRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PortRange.Merge(m, src)
}
func (m *PortRange) XXX_Size() int {
	return xxx_messageInfo_PortRange.Size(m)
}
func (m *PortRange) XXX_DiscardUnknown() {
	xxx_messageInfo_PortRange.DiscardUnknown(m)
}

var xxx_messageInfo_PortRange proto.InternalMessageInfo

func (m *PortRange) GetFromPort() int64 {
	if m != nil {
		return m.FromPort
	}
	return 0
}

func (m *PortRange) GetToPort() int64 {
	if m != nil {
		return m.ToPort
	}
	return 0
}

type CidrBlocks struct {
	V4CidrBlocks         []string `protobuf:"bytes,1,rep,name=v4_cidr_blocks,json=v4CidrBlocks,proto3" json:"v4_cidr_blocks,omitempty"`
	V6CidrBlocks         []string `protobuf:"bytes,2,rep,name=v6_cidr_blocks,json=v6CidrBlocks,proto3" json:"v6_cidr_blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CidrBlocks) Reset()         { *m = CidrBlocks{} }
func (m *CidrBlocks) String() string { return proto.CompactTextString(m) }
func (*CidrBlocks) ProtoMessage()    {}
func (*CidrBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d1de6b05354e7cc, []int{3}
}

func (m *CidrBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CidrBlocks.Unmarshal(m, b)
}
func (m *CidrBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CidrBlocks.Marshal(b, m, deterministic)
}
func (m *CidrBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CidrBlocks.Merge(m, src)
}
func (m *CidrBlocks) XXX_Size() int {
	return xxx_messageInfo_CidrBlocks.Size(m)
}
func (m *CidrBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_CidrBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_CidrBlocks proto.InternalMessageInfo

func (m *CidrBlocks) GetV4CidrBlocks() []string {
	if m != nil {
		return m.V4CidrBlocks
	}
	return nil
}

func (m *CidrBlocks) GetV6CidrBlocks() []string {
	if m != nil {
		return m.V6CidrBlocks
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.vpc.v1.SecurityGroup_Status", SecurityGroup_Status_name, SecurityGroup_Status_value)
	proto.RegisterEnum("yandex.cloud.vpc.v1.SecurityGroupRule_Direction", SecurityGroupRule_Direction_name, SecurityGroupRule_Direction_value)
	proto.RegisterType((*SecurityGroup)(nil), "yandex.cloud.vpc.v1.SecurityGroup")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.SecurityGroup.LabelsEntry")
	proto.RegisterType((*SecurityGroupRule)(nil), "yandex.cloud.vpc.v1.SecurityGroupRule")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.vpc.v1.SecurityGroupRule.LabelsEntry")
	proto.RegisterType((*PortRange)(nil), "yandex.cloud.vpc.v1.PortRange")
	proto.RegisterType((*CidrBlocks)(nil), "yandex.cloud.vpc.v1.CidrBlocks")
}

func init() {
	proto.RegisterFile("yandex/cloud/vpc/v1/security_group.proto", fileDescriptor_7d1de6b05354e7cc)
}

var fileDescriptor_7d1de6b05354e7cc = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x5e, 0xff, 0xd4, 0x89, 0x8f, 0x77, 0x43, 0x18, 0xfe, 0x4c, 0xd0, 0xb2, 0x51, 0x58, 0x41,
	0xb8, 0xa8, 0xbd, 0xc9, 0xb6, 0x15, 0xcb, 0x22, 0xa1, 0xfc, 0x78, 0x8b, 0xd1, 0x2a, 0x54, 0x4e,
	0x5a, 0x21, 0xb8, 0xb0, 0x1c, 0x7b, 0x6a, 0xac, 0x3a, 0x19, 0x6b, 0x3c, 0x36, 0xe4, 0x8d, 0x78,
	0x07, 0x6e, 0xe0, 0xaa, 0xaf, 0xc2, 0x33, 0x70, 0x85, 0x3c, 0x76, 0x92, 0xa6, 0x89, 0x44, 0x25,
	0xee, 0x66, 0xce, 0xf9, 0xce, 0x39, 0xdf, 0x7c, 0xe7, 0xd3, 0x40, 0x77, 0xe5, 0x2d, 0x03, 0xfc,
	0x9b, 0xe9, 0xc7, 0x24, 0x0b, 0xcc, 0x3c, 0xf1, 0xcd, 0xbc, 0x67, 0xa6, 0xd8, 0xcf, 0x68, 0xc4,
	0x56, 0x6e, 0x48, 0x49, 0x96, 0x18, 0x09, 0x25, 0x8c, 0xa0, 0xf7, 0x4a, 0xa4, 0xc1, 0x91, 0x46,
	0x9e, 0xf8, 0x46, 0xde, 0x6b, 0x3d, 0x0b, 0x09, 0x09, 0x63, 0x6c, 0x72, 0xc8, 0x3c, 0xbb, 0x36,
	0x59, 0xb4, 0xc0, 0x29, 0xf3, 0x16, 0x55, 0x55, 0xeb, 0xe9, 0x6e, 0x7f, 0x2f, 0x8e, 0x02, 0x8f,
	0x45, 0x64, 0x59, 0xa6, 0x3b, 0xbf, 0xcb, 0xf0, 0x64, 0x5a, 0x4d, 0x3b, 0x2f, 0x86, 0xa1, 0x06,
	0x88, 0x51, 0xa0, 0x0b, 0x6d, 0xa1, 0xab, 0x3a, 0x62, 0x14, 0xa0, 0x4f, 0x40, 0xbd, 0x26, 0x71,
	0x80, 0xa9, 0x1b, 0x05, 0xba, 0xc8, 0xc3, 0xf5, 0x32, 0x60, 0x07, 0xe8, 0x15, 0x80, 0x4f, 0xb1,
	0xc7, 0x70, 0xe0, 0x7a, 0x4c, 0x97, 0xda, 0x42, 0x57, 0xeb, 0xb7, 0x8c, 0x92, 0x93, 0xb1, 0xe6,
	0x64, 0xcc, 0xd6, 0x9c, 0x1c, 0xb5, 0x42, 0x0f, 0x18, 0x42, 0x20, 0x2f, 0xbd, 0x05, 0xd6, 0x65,
	0xde, 0x92, 0x9f, 0x51, 0x1b, 0xb4, 0x00, 0xa7, 0x3e, 0x8d, 0x92, 0x82, 0xa2, 0x7e, 0xc4, 0x53,
	0x77, 0x43, 0xe8, 0x0d, 0x28, 0xb1, 0x37, 0xc7, 0x71, 0xaa, 0x2b, 0x6d, 0xa9, 0xab, 0xf5, 0x0d,
	0xe3, 0x80, 0x2a, 0xc6, 0xce, 0x8b, 0x8c, 0xb7, 0xbc, 0xc0, 0x5a, 0x32, 0xba, 0x72, 0xaa, 0x6a,
	0xf4, 0x14, 0x60, 0x89, 0xd9, 0xaf, 0x84, 0xde, 0x14, 0xcf, 0xaa, 0xf1, 0x41, 0x6a, 0x15, 0xb1,
	0x03, 0x34, 0x00, 0x25, 0x65, 0x1e, 0xcb, 0x52, 0xbd, 0xde, 0x16, 0xba, 0x8d, 0xfe, 0x97, 0x0f,
	0x18, 0x33, 0xe5, 0x05, 0x4e, 0x55, 0x88, 0xbe, 0x81, 0x23, 0x9a, 0xc5, 0x38, 0xd5, 0x55, 0x4e,
	0xf4, 0xf3, 0xff, 0xee, 0xe0, 0x64, 0x31, 0x76, 0xca, 0xa2, 0xd6, 0x2b, 0xd0, 0xee, 0xd0, 0x46,
	0x4d, 0x90, 0x6e, 0xf0, 0xaa, 0xda, 0x4a, 0x71, 0x44, 0xef, 0xc3, 0x51, 0xee, 0xc5, 0x19, 0xae,
	0x56, 0x52, 0x5e, 0xbe, 0x16, 0xbf, 0x12, 0x3a, 0x57, 0xa0, 0x94, 0x54, 0xd0, 0x87, 0x80, 0xa6,
	0xb3, 0xc1, 0xec, 0x72, 0xea, 0x5e, 0x4e, 0xa6, 0x17, 0xd6, 0xc8, 0x7e, 0x63, 0x5b, 0xe3, 0xe6,
	0x23, 0xf4, 0x18, 0xea, 0x23, 0xc7, 0x1a, 0xcc, 0xec, 0xc9, 0x79, 0x53, 0x40, 0x00, 0xca, 0x60,
	0x34, 0xb3, 0xaf, 0xac, 0xa6, 0x58, 0x64, 0x2e, 0x2f, 0xc6, 0x65, 0x46, 0x2a, 0x6e, 0x63, 0xeb,
	0xad, 0xc5, 0x6f, 0x72, 0xe7, 0x0f, 0x19, 0xde, 0xdd, 0xe3, 0xbb, 0x67, 0x97, 0x7b, 0x2b, 0x14,
	0xf7, 0x57, 0xf8, 0xfd, 0x66, 0x85, 0x12, 0x57, 0xa6, 0xff, 0x30, 0x65, 0x0e, 0xae, 0x71, 0x06,
	0x6a, 0x10, 0x51, 0xec, 0xf3, 0x59, 0x32, 0x5f, 0xd5, 0x8b, 0x07, 0xb6, 0x1b, 0xaf, 0xeb, 0x86,
	0xf2, 0xdf, 0xb7, 0x3d, 0xc1, 0xd9, 0x36, 0x42, 0x27, 0x70, 0x94, 0x10, 0xca, 0x52, 0x6e, 0x40,
	0xad, 0xff, 0xe9, 0xc1, 0x8e, 0x17, 0x84, 0x32, 0xc7, 0x5b, 0x86, 0xd8, 0x29, 0xc1, 0xe8, 0x33,
	0x78, 0xc2, 0x1d, 0xef, 0x93, 0xd8, 0xe5, 0xce, 0x56, 0xf8, 0xdb, 0x1f, 0xaf, 0x83, 0x93, 0xc2,
	0xe1, 0x5f, 0xc0, 0x3b, 0x5b, 0x50, 0xb6, 0x98, 0x63, 0xca, 0xcd, 0x27, 0x39, 0x8d, 0x0d, 0x8c,
	0x47, 0xd1, 0x10, 0x34, 0x3f, 0x0a, 0xa8, 0x3b, 0x8f, 0x89, 0x7f, 0x53, 0xda, 0x50, 0xeb, 0x3f,
	0x3b, 0xc8, 0x64, 0x14, 0x05, 0x74, 0xc8, 0x61, 0xdf, 0x3d, 0x72, 0xc0, 0xdf, 0xdc, 0xfe, 0x8f,
	0x89, 0xbe, 0x05, 0x75, 0x23, 0x10, 0xfa, 0x18, 0x3e, 0x18, 0xdb, 0x8e, 0x35, 0x9a, 0xd9, 0x3f,
	0x4c, 0xee, 0x59, 0x49, 0x83, 0x9a, 0x3d, 0x39, 0x77, 0xac, 0xe9, 0xb4, 0x74, 0x92, 0x55, 0x9e,
	0xc5, 0x61, 0x03, 0x14, 0xe6, 0xd1, 0x10, 0x33, 0x24, 0xff, 0xf9, 0x57, 0x4f, 0xe8, 0xfc, 0x0c,
	0xea, 0x46, 0x31, 0xd4, 0x05, 0xf5, 0x9a, 0x92, 0x85, 0x5b, 0x08, 0xc7, 0xf9, 0x48, 0x43, 0xed,
	0x9f, 0xdb, 0x5e, 0xed, 0xc5, 0xf1, 0xd9, 0xe9, 0xe9, 0xcb, 0x53, 0xa7, 0x5e, 0x64, 0x0b, 0x38,
	0x7a, 0x0e, 0x35, 0x46, 0x4a, 0x9c, 0xb8, 0x8f, 0x53, 0x18, 0x29, 0x50, 0x9d, 0x1f, 0x01, 0xb6,
	0x22, 0xa0, 0xe7, 0xd0, 0xc8, 0x4f, 0xdc, 0xbb, 0xea, 0x09, 0x6d, 0xa9, 0xd8, 0x44, 0x7e, 0x72,
	0x0f, 0x75, 0xb6, 0x83, 0x12, 0x2b, 0xd4, 0xd9, 0x16, 0x35, 0xbc, 0x82, 0x8f, 0x76, 0x24, 0xf7,
	0x92, 0xa8, 0x92, 0xfd, 0xa7, 0xd7, 0x61, 0xc4, 0x7e, 0xc9, 0xe6, 0x86, 0x4f, 0x16, 0x66, 0x89,
	0x39, 0x2e, 0x3f, 0xd9, 0x90, 0x1c, 0x87, 0x78, 0xc9, 0xf7, 0x69, 0x1e, 0xf8, 0xdd, 0x5f, 0xe7,
	0x89, 0x3f, 0x57, 0x78, 0xfa, 0xe5, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x36, 0x23, 0xd5, 0x27,
	0xff, 0x05, 0x00, 0x00,
}
