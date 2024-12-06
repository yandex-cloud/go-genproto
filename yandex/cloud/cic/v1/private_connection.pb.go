// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/cic/v1/private_connection.proto

package cic

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A PrivateConnection resource.
type PrivateConnection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the privateConnection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the privateConnection.
	// The name must be unique within the folder.
	// Value must match the regular expression “\|[a-zA-Z]([-_a-zA-Z0-9]{0,61}[a-zA-Z0-9])?“.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description of the privateConnection. 0-256 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the folder that the privateConnection belongs to.
	FolderId string `protobuf:"bytes,5,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// ID of the region that the privateConnection belongs to.
	RegionId string `protobuf:"bytes,6,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// ID of the trunk_connection that the privateConnection belongs to.
	TrunkConnectionId string `protobuf:"bytes,7,opt,name=trunk_connection_id,json=trunkConnectionId,proto3" json:"trunk_connection_id,omitempty"`
	// VLAN_ID that the privateConnection uses in multiplexing.
	// Not used in connections over partners-II
	// Value range: [1, 4095]
	VlanId *wrapperspb.Int64Value `protobuf:"bytes,8,opt,name=vlan_id,json=vlanId,proto3" json:"vlan_id,omitempty"`
	// IPv4 peering config of connection
	Ipv4Peering *Peering `protobuf:"bytes,9,opt,name=ipv4_peering,json=ipv4Peering,proto3" json:"ipv4_peering,omitempty"`
	// IPv4 StaticRoute config of connection
	Ipv4StaticRoutes []*PrivateConnection_StaticRoute `protobuf:"bytes,18,rep,name=ipv4_static_routes,json=ipv4StaticRoutes,proto3" json:"ipv4_static_routes,omitempty"`
	// Resource labels, `key:value` pairs.
	// No more than 64 per resource.
	// The maximum string length in characters for each value is 63.
	// Each value must match the regular expression `[-_0-9a-z]*`.
	// The string length in characters for each key must be 1-63.
	// Each key must match the regular expression `[a-z][-_0-9a-z]*`.
	Labels map[string]string `protobuf:"bytes,24,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PrivateConnection) Reset() {
	*x = PrivateConnection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_private_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateConnection) ProtoMessage() {}

func (x *PrivateConnection) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_private_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateConnection.ProtoReflect.Descriptor instead.
func (*PrivateConnection) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_private_connection_proto_rawDescGZIP(), []int{0}
}

func (x *PrivateConnection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PrivateConnection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PrivateConnection) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PrivateConnection) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *PrivateConnection) GetRegionId() string {
	if x != nil {
		return x.RegionId
	}
	return ""
}

func (x *PrivateConnection) GetTrunkConnectionId() string {
	if x != nil {
		return x.TrunkConnectionId
	}
	return ""
}

func (x *PrivateConnection) GetVlanId() *wrapperspb.Int64Value {
	if x != nil {
		return x.VlanId
	}
	return nil
}

func (x *PrivateConnection) GetIpv4Peering() *Peering {
	if x != nil {
		return x.Ipv4Peering
	}
	return nil
}

func (x *PrivateConnection) GetIpv4StaticRoutes() []*PrivateConnection_StaticRoute {
	if x != nil {
		return x.Ipv4StaticRoutes
	}
	return nil
}

func (x *PrivateConnection) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type PrivateConnection_StaticRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Prefix.
	// It's an ip with format ipPrefix/length where address part of ipPrefix is 0.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// PeerIp.
	// It's an ip with just an ipAddress format without mask.
	// Will be removed in some next release
	//
	// Deprecated: Marked as deprecated in yandex/cloud/cic/v1/private_connection.proto.
	NextHop []string `protobuf:"bytes,2,rep,name=next_hop,json=nextHop,proto3" json:"next_hop,omitempty"`
}

func (x *PrivateConnection_StaticRoute) Reset() {
	*x = PrivateConnection_StaticRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_private_connection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateConnection_StaticRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateConnection_StaticRoute) ProtoMessage() {}

func (x *PrivateConnection_StaticRoute) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_private_connection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateConnection_StaticRoute.ProtoReflect.Descriptor instead.
func (*PrivateConnection_StaticRoute) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_private_connection_proto_rawDescGZIP(), []int{0, 1}
}

func (x *PrivateConnection_StaticRoute) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

// Deprecated: Marked as deprecated in yandex/cloud/cic/v1/private_connection.proto.
func (x *PrivateConnection_StaticRoute) GetNextHop() []string {
	if x != nil {
		return x.NextHop
	}
	return nil
}

var File_yandex_cloud_cic_v1_private_connection_proto protoreflect.FileDescriptor

var file_yandex_cloud_cic_v1_private_connection_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x04, 0x0a, 0x11, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x13,
	0x74, 0x72, 0x75, 0x6e, 0x6b, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x74, 0x72, 0x75, 0x6e, 0x6b,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x07,
	0x76, 0x6c, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x6c, 0x61, 0x6e,
	0x49, 0x64, 0x12, 0x3f, 0x0a, 0x0c, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x70, 0x65, 0x65, 0x72, 0x69,
	0x6e, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x65, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x52, 0x0b, 0x69, 0x70, 0x76, 0x34, 0x50, 0x65, 0x65, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x60, 0x0a, 0x12, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x63, 0x5f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x52, 0x10, 0x69, 0x70, 0x76, 0x34, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x73, 0x12, 0x4a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x18, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x0b,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x1d, 0x0a, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x68, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x07, 0x6e, 0x65, 0x78, 0x74, 0x48,
	0x6f, 0x70, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x12, 0x4a, 0x04,
	0x08, 0x13, 0x10, 0x18, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cic_v1_private_connection_proto_rawDescOnce sync.Once
	file_yandex_cloud_cic_v1_private_connection_proto_rawDescData = file_yandex_cloud_cic_v1_private_connection_proto_rawDesc
)

func file_yandex_cloud_cic_v1_private_connection_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cic_v1_private_connection_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cic_v1_private_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cic_v1_private_connection_proto_rawDescData)
	})
	return file_yandex_cloud_cic_v1_private_connection_proto_rawDescData
}

var file_yandex_cloud_cic_v1_private_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_cic_v1_private_connection_proto_goTypes = []any{
	(*PrivateConnection)(nil),             // 0: yandex.cloud.cic.v1.PrivateConnection
	nil,                                   // 1: yandex.cloud.cic.v1.PrivateConnection.LabelsEntry
	(*PrivateConnection_StaticRoute)(nil), // 2: yandex.cloud.cic.v1.PrivateConnection.StaticRoute
	(*wrapperspb.Int64Value)(nil),         // 3: google.protobuf.Int64Value
	(*Peering)(nil),                       // 4: yandex.cloud.cic.v1.Peering
}
var file_yandex_cloud_cic_v1_private_connection_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.cic.v1.PrivateConnection.vlan_id:type_name -> google.protobuf.Int64Value
	4, // 1: yandex.cloud.cic.v1.PrivateConnection.ipv4_peering:type_name -> yandex.cloud.cic.v1.Peering
	2, // 2: yandex.cloud.cic.v1.PrivateConnection.ipv4_static_routes:type_name -> yandex.cloud.cic.v1.PrivateConnection.StaticRoute
	1, // 3: yandex.cloud.cic.v1.PrivateConnection.labels:type_name -> yandex.cloud.cic.v1.PrivateConnection.LabelsEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cic_v1_private_connection_proto_init() }
func file_yandex_cloud_cic_v1_private_connection_proto_init() {
	if File_yandex_cloud_cic_v1_private_connection_proto != nil {
		return
	}
	file_yandex_cloud_cic_v1_peering_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cic_v1_private_connection_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateConnection); i {
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
		file_yandex_cloud_cic_v1_private_connection_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateConnection_StaticRoute); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_cic_v1_private_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_cic_v1_private_connection_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cic_v1_private_connection_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cic_v1_private_connection_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cic_v1_private_connection_proto = out.File
	file_yandex_cloud_cic_v1_private_connection_proto_rawDesc = nil
	file_yandex_cloud_cic_v1_private_connection_proto_goTypes = nil
	file_yandex_cloud_cic_v1_private_connection_proto_depIdxs = nil
}