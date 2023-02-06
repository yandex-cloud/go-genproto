// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/vpc/v1/subnet.proto

package vpc

import (
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

type IpVersion int32

const (
	IpVersion_IP_VERSION_UNSPECIFIED IpVersion = 0
	IpVersion_IPV4                   IpVersion = 1
	IpVersion_IPV6                   IpVersion = 2
)

// Enum value maps for IpVersion.
var (
	IpVersion_name = map[int32]string{
		0: "IP_VERSION_UNSPECIFIED",
		1: "IPV4",
		2: "IPV6",
	}
	IpVersion_value = map[string]int32{
		"IP_VERSION_UNSPECIFIED": 0,
		"IPV4":                   1,
		"IPV6":                   2,
	}
)

func (x IpVersion) Enum() *IpVersion {
	p := new(IpVersion)
	*p = x
	return p
}

func (x IpVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IpVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_vpc_v1_subnet_proto_enumTypes[0].Descriptor()
}

func (IpVersion) Type() protoreflect.EnumType {
	return &file_yandex_cloud_vpc_v1_subnet_proto_enumTypes[0]
}

func (x IpVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IpVersion.Descriptor instead.
func (IpVersion) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_subnet_proto_rawDescGZIP(), []int{0}
}

// A Subnet resource. For more information, see [Subnets](/docs/vpc/concepts/subnets).
type Subnet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subnet.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the subnet belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the subnet. The name is unique within the project. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Optional description of the subnet. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as “ key:value “ pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the network the subnet belongs to.
	NetworkId string `protobuf:"bytes,7,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// ID of the availability zone where the subnet resides.
	ZoneId string `protobuf:"bytes,8,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"` // if subnet will be zonal
	// CIDR block.
	// The range of internal addresses that are defined for this subnet.
	// This field can be set only at Subnet resource creation time and cannot be changed.
	// For example, 10.0.0.0/22 or 192.168.0.0/24.
	// Minimum subnet size is /28, maximum subnet size is /16.
	V4CidrBlocks []string `protobuf:"bytes,10,rep,name=v4_cidr_blocks,json=v4CidrBlocks,proto3" json:"v4_cidr_blocks,omitempty"`
	// IPv6 not available yet.
	V6CidrBlocks []string `protobuf:"bytes,11,rep,name=v6_cidr_blocks,json=v6CidrBlocks,proto3" json:"v6_cidr_blocks,omitempty"`
	// ID of route table the subnet is linked to.
	RouteTableId string       `protobuf:"bytes,12,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
	DhcpOptions  *DhcpOptions `protobuf:"bytes,13,opt,name=dhcp_options,json=dhcpOptions,proto3" json:"dhcp_options,omitempty"`
}

func (x *Subnet) Reset() {
	*x = Subnet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_subnet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subnet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subnet) ProtoMessage() {}

func (x *Subnet) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_subnet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subnet.ProtoReflect.Descriptor instead.
func (*Subnet) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_subnet_proto_rawDescGZIP(), []int{0}
}

func (x *Subnet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subnet) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Subnet) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Subnet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subnet) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Subnet) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Subnet) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Subnet) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *Subnet) GetV4CidrBlocks() []string {
	if x != nil {
		return x.V4CidrBlocks
	}
	return nil
}

func (x *Subnet) GetV6CidrBlocks() []string {
	if x != nil {
		return x.V6CidrBlocks
	}
	return nil
}

func (x *Subnet) GetRouteTableId() string {
	if x != nil {
		return x.RouteTableId
	}
	return ""
}

func (x *Subnet) GetDhcpOptions() *DhcpOptions {
	if x != nil {
		return x.DhcpOptions
	}
	return nil
}

type DhcpOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DomainNameServers []string `protobuf:"bytes,1,rep,name=domain_name_servers,json=domainNameServers,proto3" json:"domain_name_servers,omitempty"`
	DomainName        string   `protobuf:"bytes,2,opt,name=domain_name,json=domainName,proto3" json:"domain_name,omitempty"`
	NtpServers        []string `protobuf:"bytes,3,rep,name=ntp_servers,json=ntpServers,proto3" json:"ntp_servers,omitempty"`
}

func (x *DhcpOptions) Reset() {
	*x = DhcpOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_subnet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DhcpOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DhcpOptions) ProtoMessage() {}

func (x *DhcpOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_subnet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DhcpOptions.ProtoReflect.Descriptor instead.
func (*DhcpOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_subnet_proto_rawDescGZIP(), []int{1}
}

func (x *DhcpOptions) GetDomainNameServers() []string {
	if x != nil {
		return x.DomainNameServers
	}
	return nil
}

func (x *DhcpOptions) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *DhcpOptions) GetNtpServers() []string {
	if x != nil {
		return x.NtpServers
	}
	return nil
}

var File_yandex_cloud_vpc_v1_subnet_proto protoreflect.FileDescriptor

var file_yandex_cloud_vpc_v1_subnet_proto_rawDesc = []byte{
	0x0a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x04, 0x0a, 0x06, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x76, 0x34,
	0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0c, 0x76, 0x34, 0x43, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x12, 0x24, 0x0a, 0x0e, 0x76, 0x36, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x76, 0x36, 0x43, 0x69, 0x64, 0x72,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x5f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x43, 0x0a, 0x0c,
	0x64, 0x68, 0x63, 0x70, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x68, 0x63, 0x70, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x64, 0x68, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7f, 0x0a, 0x0b,
	0x44, 0x68, 0x63, 0x70, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x4e, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x6e, 0x74, 0x70, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x6e, 0x74, 0x70, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x2a, 0x3b, 0x0a,
	0x09, 0x49, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x50,
	0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x36, 0x10, 0x02, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_vpc_v1_subnet_proto_rawDescOnce sync.Once
	file_yandex_cloud_vpc_v1_subnet_proto_rawDescData = file_yandex_cloud_vpc_v1_subnet_proto_rawDesc
)

func file_yandex_cloud_vpc_v1_subnet_proto_rawDescGZIP() []byte {
	file_yandex_cloud_vpc_v1_subnet_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_vpc_v1_subnet_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_vpc_v1_subnet_proto_rawDescData)
	})
	return file_yandex_cloud_vpc_v1_subnet_proto_rawDescData
}

var file_yandex_cloud_vpc_v1_subnet_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_vpc_v1_subnet_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_vpc_v1_subnet_proto_goTypes = []interface{}{
	(IpVersion)(0),                // 0: yandex.cloud.vpc.v1.IpVersion
	(*Subnet)(nil),                // 1: yandex.cloud.vpc.v1.Subnet
	(*DhcpOptions)(nil),           // 2: yandex.cloud.vpc.v1.DhcpOptions
	nil,                           // 3: yandex.cloud.vpc.v1.Subnet.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_vpc_v1_subnet_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.vpc.v1.Subnet.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: yandex.cloud.vpc.v1.Subnet.labels:type_name -> yandex.cloud.vpc.v1.Subnet.LabelsEntry
	2, // 2: yandex.cloud.vpc.v1.Subnet.dhcp_options:type_name -> yandex.cloud.vpc.v1.DhcpOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_vpc_v1_subnet_proto_init() }
func file_yandex_cloud_vpc_v1_subnet_proto_init() {
	if File_yandex_cloud_vpc_v1_subnet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_vpc_v1_subnet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subnet); i {
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
		file_yandex_cloud_vpc_v1_subnet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DhcpOptions); i {
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
			RawDescriptor: file_yandex_cloud_vpc_v1_subnet_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_vpc_v1_subnet_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_vpc_v1_subnet_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_vpc_v1_subnet_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_vpc_v1_subnet_proto_msgTypes,
	}.Build()
	File_yandex_cloud_vpc_v1_subnet_proto = out.File
	file_yandex_cloud_vpc_v1_subnet_proto_rawDesc = nil
	file_yandex_cloud_vpc_v1_subnet_proto_goTypes = nil
	file_yandex_cloud_vpc_v1_subnet_proto_depIdxs = nil
}
