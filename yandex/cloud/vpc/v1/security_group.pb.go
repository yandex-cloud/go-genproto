// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/vpc/v1/security_group.proto

package vpc

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SecurityGroup_Status int32

const (
	SecurityGroup_STATUS_UNSPECIFIED SecurityGroup_Status = 0
	SecurityGroup_CREATING           SecurityGroup_Status = 1
	SecurityGroup_ACTIVE             SecurityGroup_Status = 2
	// updating is a long operation because we must update all instances in SG
	SecurityGroup_UPDATING SecurityGroup_Status = 3
	SecurityGroup_DELETING SecurityGroup_Status = 4
)

// Enum value maps for SecurityGroup_Status.
var (
	SecurityGroup_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "UPDATING",
		4: "DELETING",
	}
	SecurityGroup_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"UPDATING":           3,
		"DELETING":           4,
	}
)

func (x SecurityGroup_Status) Enum() *SecurityGroup_Status {
	p := new(SecurityGroup_Status)
	*p = x
	return p
}

func (x SecurityGroup_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityGroup_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_vpc_v1_security_group_proto_enumTypes[0].Descriptor()
}

func (SecurityGroup_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_vpc_v1_security_group_proto_enumTypes[0]
}

func (x SecurityGroup_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityGroup_Status.Descriptor instead.
func (SecurityGroup_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP(), []int{0, 0}
}

type SecurityGroupRule_Direction int32

const (
	SecurityGroupRule_DIRECTION_UNSPECIFIED SecurityGroupRule_Direction = 0
	SecurityGroupRule_INGRESS               SecurityGroupRule_Direction = 1
	SecurityGroupRule_EGRESS                SecurityGroupRule_Direction = 2
)

// Enum value maps for SecurityGroupRule_Direction.
var (
	SecurityGroupRule_Direction_name = map[int32]string{
		0: "DIRECTION_UNSPECIFIED",
		1: "INGRESS",
		2: "EGRESS",
	}
	SecurityGroupRule_Direction_value = map[string]int32{
		"DIRECTION_UNSPECIFIED": 0,
		"INGRESS":               1,
		"EGRESS":                2,
	}
)

func (x SecurityGroupRule_Direction) Enum() *SecurityGroupRule_Direction {
	p := new(SecurityGroupRule_Direction)
	*p = x
	return p
}

func (x SecurityGroupRule_Direction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SecurityGroupRule_Direction) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_vpc_v1_security_group_proto_enumTypes[1].Descriptor()
}

func (SecurityGroupRule_Direction) Type() protoreflect.EnumType {
	return &file_yandex_cloud_vpc_v1_security_group_proto_enumTypes[1]
}

func (x SecurityGroupRule_Direction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SecurityGroupRule_Direction.Descriptor instead.
func (SecurityGroupRule_Direction) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP(), []int{1, 0}
}

type SecurityGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FolderId          string                 `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt         *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Name              string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description       string                 `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Labels            map[string]string      `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	NetworkId         string                 `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	Status            SecurityGroup_Status   `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.vpc.v1.SecurityGroup_Status" json:"status,omitempty"`
	Rules             []*SecurityGroupRule   `protobuf:"bytes,9,rep,name=rules,proto3" json:"rules,omitempty"`
	DefaultForNetwork bool                   `protobuf:"varint,10,opt,name=default_for_network,json=defaultForNetwork,proto3" json:"default_for_network,omitempty"`
}

func (x *SecurityGroup) Reset() {
	*x = SecurityGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroup) ProtoMessage() {}

func (x *SecurityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroup.ProtoReflect.Descriptor instead.
func (*SecurityGroup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP(), []int{0}
}

func (x *SecurityGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SecurityGroup) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *SecurityGroup) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *SecurityGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SecurityGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SecurityGroup) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SecurityGroup) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *SecurityGroup) GetStatus() SecurityGroup_Status {
	if x != nil {
		return x.Status
	}
	return SecurityGroup_STATUS_UNSPECIFIED
}

func (x *SecurityGroup) GetRules() []*SecurityGroupRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *SecurityGroup) GetDefaultForNetwork() bool {
	if x != nil {
		return x.DefaultForNetwork
	}
	return false
}

type SecurityGroupRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` //generated by api server after rule creation
	Description string                      `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string           `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Direction   SecurityGroupRule_Direction `protobuf:"varint,4,opt,name=direction,proto3,enum=yandex.cloud.vpc.v1.SecurityGroupRule_Direction" json:"direction,omitempty"`
	Ports       *PortRange                  `protobuf:"bytes,5,opt,name=ports,proto3" json:"ports,omitempty"` // null value means any
	// null value means any protocol
	// values from https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
	ProtocolName   string `protobuf:"bytes,6,opt,name=protocol_name,json=protocolName,proto3" json:"protocol_name,omitempty"`
	ProtocolNumber int64  `protobuf:"varint,7,opt,name=protocol_number,json=protocolNumber,proto3" json:"protocol_number,omitempty"`
	// Types that are assignable to Target:
	//	*SecurityGroupRule_CidrBlocks
	//	*SecurityGroupRule_SecurityGroupId
	//	*SecurityGroupRule_PredefinedTarget
	Target isSecurityGroupRule_Target `protobuf_oneof:"target"`
}

func (x *SecurityGroupRule) Reset() {
	*x = SecurityGroupRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SecurityGroupRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SecurityGroupRule) ProtoMessage() {}

func (x *SecurityGroupRule) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SecurityGroupRule.ProtoReflect.Descriptor instead.
func (*SecurityGroupRule) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP(), []int{1}
}

func (x *SecurityGroupRule) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SecurityGroupRule) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *SecurityGroupRule) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *SecurityGroupRule) GetDirection() SecurityGroupRule_Direction {
	if x != nil {
		return x.Direction
	}
	return SecurityGroupRule_DIRECTION_UNSPECIFIED
}

func (x *SecurityGroupRule) GetPorts() *PortRange {
	if x != nil {
		return x.Ports
	}
	return nil
}

func (x *SecurityGroupRule) GetProtocolName() string {
	if x != nil {
		return x.ProtocolName
	}
	return ""
}

func (x *SecurityGroupRule) GetProtocolNumber() int64 {
	if x != nil {
		return x.ProtocolNumber
	}
	return 0
}

func (m *SecurityGroupRule) GetTarget() isSecurityGroupRule_Target {
	if m != nil {
		return m.Target
	}
	return nil
}

func (x *SecurityGroupRule) GetCidrBlocks() *CidrBlocks {
	if x, ok := x.GetTarget().(*SecurityGroupRule_CidrBlocks); ok {
		return x.CidrBlocks
	}
	return nil
}

func (x *SecurityGroupRule) GetSecurityGroupId() string {
	if x, ok := x.GetTarget().(*SecurityGroupRule_SecurityGroupId); ok {
		return x.SecurityGroupId
	}
	return ""
}

func (x *SecurityGroupRule) GetPredefinedTarget() string {
	if x, ok := x.GetTarget().(*SecurityGroupRule_PredefinedTarget); ok {
		return x.PredefinedTarget
	}
	return ""
}

type isSecurityGroupRule_Target interface {
	isSecurityGroupRule_Target()
}

type SecurityGroupRule_CidrBlocks struct {
	CidrBlocks *CidrBlocks `protobuf:"bytes,8,opt,name=cidr_blocks,json=cidrBlocks,proto3,oneof"`
}

type SecurityGroupRule_SecurityGroupId struct {
	SecurityGroupId string `protobuf:"bytes,9,opt,name=security_group_id,json=securityGroupId,proto3,oneof"`
}

type SecurityGroupRule_PredefinedTarget struct {
	PredefinedTarget string `protobuf:"bytes,10,opt,name=predefined_target,json=predefinedTarget,proto3,oneof"`
}

func (*SecurityGroupRule_CidrBlocks) isSecurityGroupRule_Target() {}

func (*SecurityGroupRule_SecurityGroupId) isSecurityGroupRule_Target() {}

func (*SecurityGroupRule_PredefinedTarget) isSecurityGroupRule_Target() {}

type PortRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromPort int64 `protobuf:"varint,1,opt,name=from_port,json=fromPort,proto3" json:"from_port,omitempty"`
	ToPort   int64 `protobuf:"varint,2,opt,name=to_port,json=toPort,proto3" json:"to_port,omitempty"`
}

func (x *PortRange) Reset() {
	*x = PortRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PortRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PortRange) ProtoMessage() {}

func (x *PortRange) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PortRange.ProtoReflect.Descriptor instead.
func (*PortRange) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP(), []int{2}
}

func (x *PortRange) GetFromPort() int64 {
	if x != nil {
		return x.FromPort
	}
	return 0
}

func (x *PortRange) GetToPort() int64 {
	if x != nil {
		return x.ToPort
	}
	return 0
}

type CidrBlocks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V4CidrBlocks []string `protobuf:"bytes,1,rep,name=v4_cidr_blocks,json=v4CidrBlocks,proto3" json:"v4_cidr_blocks,omitempty"`
	V6CidrBlocks []string `protobuf:"bytes,2,rep,name=v6_cidr_blocks,json=v6CidrBlocks,proto3" json:"v6_cidr_blocks,omitempty"`
}

func (x *CidrBlocks) Reset() {
	*x = CidrBlocks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CidrBlocks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CidrBlocks) ProtoMessage() {}

func (x *CidrBlocks) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CidrBlocks.ProtoReflect.Descriptor instead.
func (*CidrBlocks) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP(), []int{3}
}

func (x *CidrBlocks) GetV4CidrBlocks() []string {
	if x != nil {
		return x.V4CidrBlocks
	}
	return nil
}

func (x *CidrBlocks) GetV6CidrBlocks() []string {
	if x != nil {
		return x.V6CidrBlocks
	}
	return nil
}

var File_yandex_cloud_vpc_v1_security_group_proto protoreflect.FileDescriptor

var file_yandex_cloud_vpc_v1_security_group_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd8, 0x04, 0x0a, 0x0d, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x46, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3c, 0x0a, 0x05, 0x72, 0x75, 0x6c,
	0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65,
	0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6f, 0x72,
	0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x56, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x22, 0x98, 0x05, 0x0a, 0x11, 0x53,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x54,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x27, 0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x0b, 0x63, 0x69, 0x64, 0x72,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x48, 0x00,
	0x52, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x2c, 0x0a, 0x11,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0f, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x11, 0x70, 0x72,
	0x65, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x70, 0x72, 0x65, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x64, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a, 0x09, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x15, 0x44, 0x49, 0x52, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x49, 0x4e, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x47, 0x52,
	0x45, 0x53, 0x53, 0x10, 0x02, 0x42, 0x0e, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x5b, 0x0a, 0x09, 0x50, 0x6f, 0x72, 0x74, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x28, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xfa, 0xc7, 0x31, 0x07, 0x30, 0x2d, 0x36, 0x35, 0x35,
	0x33, 0x35, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x50, 0x6f, 0x72, 0x74, 0x12, 0x24, 0x0a, 0x07,
	0x74, 0x6f, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0b, 0xfa,
	0xc7, 0x31, 0x07, 0x30, 0x2d, 0x36, 0x35, 0x35, 0x33, 0x35, 0x52, 0x06, 0x74, 0x6f, 0x50, 0x6f,
	0x72, 0x74, 0x22, 0x58, 0x0a, 0x0a, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x24, 0x0a, 0x0e, 0x76, 0x34, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x34, 0x43, 0x69, 0x64, 0x72,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x36, 0x5f, 0x63, 0x69, 0x64,
	0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x76, 0x36, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x42, 0x56, 0x0a, 0x17,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_vpc_v1_security_group_proto_rawDescOnce sync.Once
	file_yandex_cloud_vpc_v1_security_group_proto_rawDescData = file_yandex_cloud_vpc_v1_security_group_proto_rawDesc
)

func file_yandex_cloud_vpc_v1_security_group_proto_rawDescGZIP() []byte {
	file_yandex_cloud_vpc_v1_security_group_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_vpc_v1_security_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_vpc_v1_security_group_proto_rawDescData)
	})
	return file_yandex_cloud_vpc_v1_security_group_proto_rawDescData
}

var file_yandex_cloud_vpc_v1_security_group_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_vpc_v1_security_group_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_vpc_v1_security_group_proto_goTypes = []interface{}{
	(SecurityGroup_Status)(0),        // 0: yandex.cloud.vpc.v1.SecurityGroup.Status
	(SecurityGroupRule_Direction)(0), // 1: yandex.cloud.vpc.v1.SecurityGroupRule.Direction
	(*SecurityGroup)(nil),            // 2: yandex.cloud.vpc.v1.SecurityGroup
	(*SecurityGroupRule)(nil),        // 3: yandex.cloud.vpc.v1.SecurityGroupRule
	(*PortRange)(nil),                // 4: yandex.cloud.vpc.v1.PortRange
	(*CidrBlocks)(nil),               // 5: yandex.cloud.vpc.v1.CidrBlocks
	nil,                              // 6: yandex.cloud.vpc.v1.SecurityGroup.LabelsEntry
	nil,                              // 7: yandex.cloud.vpc.v1.SecurityGroupRule.LabelsEntry
	(*timestamppb.Timestamp)(nil),    // 8: google.protobuf.Timestamp
}
var file_yandex_cloud_vpc_v1_security_group_proto_depIdxs = []int32{
	8, // 0: yandex.cloud.vpc.v1.SecurityGroup.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: yandex.cloud.vpc.v1.SecurityGroup.labels:type_name -> yandex.cloud.vpc.v1.SecurityGroup.LabelsEntry
	0, // 2: yandex.cloud.vpc.v1.SecurityGroup.status:type_name -> yandex.cloud.vpc.v1.SecurityGroup.Status
	3, // 3: yandex.cloud.vpc.v1.SecurityGroup.rules:type_name -> yandex.cloud.vpc.v1.SecurityGroupRule
	7, // 4: yandex.cloud.vpc.v1.SecurityGroupRule.labels:type_name -> yandex.cloud.vpc.v1.SecurityGroupRule.LabelsEntry
	1, // 5: yandex.cloud.vpc.v1.SecurityGroupRule.direction:type_name -> yandex.cloud.vpc.v1.SecurityGroupRule.Direction
	4, // 6: yandex.cloud.vpc.v1.SecurityGroupRule.ports:type_name -> yandex.cloud.vpc.v1.PortRange
	5, // 7: yandex.cloud.vpc.v1.SecurityGroupRule.cidr_blocks:type_name -> yandex.cloud.vpc.v1.CidrBlocks
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_yandex_cloud_vpc_v1_security_group_proto_init() }
func file_yandex_cloud_vpc_v1_security_group_proto_init() {
	if File_yandex_cloud_vpc_v1_security_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityGroup); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SecurityGroupRule); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PortRange); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CidrBlocks); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yandex_cloud_vpc_v1_security_group_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SecurityGroupRule_CidrBlocks)(nil),
		(*SecurityGroupRule_SecurityGroupId)(nil),
		(*SecurityGroupRule_PredefinedTarget)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_vpc_v1_security_group_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_vpc_v1_security_group_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_vpc_v1_security_group_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_vpc_v1_security_group_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_vpc_v1_security_group_proto_msgTypes,
	}.Build()
	File_yandex_cloud_vpc_v1_security_group_proto = out.File
	file_yandex_cloud_vpc_v1_security_group_proto_rawDesc = nil
	file_yandex_cloud_vpc_v1_security_group_proto_goTypes = nil
	file_yandex_cloud_vpc_v1_security_group_proto_depIdxs = nil
}
