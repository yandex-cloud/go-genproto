// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/vpc/v1/address.proto

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

type Address_Type int32

const (
	Address_TYPE_UNSPECIFIED Address_Type = 0
	// Internal IP address.
	Address_INTERNAL Address_Type = 1
	// Public IP address.
	Address_EXTERNAL Address_Type = 2
)

// Enum value maps for Address_Type.
var (
	Address_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "INTERNAL",
		2: "EXTERNAL",
	}
	Address_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"INTERNAL":         1,
		"EXTERNAL":         2,
	}
)

func (x Address_Type) Enum() *Address_Type {
	p := new(Address_Type)
	*p = x
	return p
}

func (x Address_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Address_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_vpc_v1_address_proto_enumTypes[0].Descriptor()
}

func (Address_Type) Type() protoreflect.EnumType {
	return &file_yandex_cloud_vpc_v1_address_proto_enumTypes[0]
}

func (x Address_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Address_Type.Descriptor instead.
func (Address_Type) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{0, 0}
}

type Address_IpVersion int32

const (
	Address_IP_VERSION_UNSPECIFIED Address_IpVersion = 0
	// IPv4 address.
	Address_IPV4 Address_IpVersion = 1
	// IPv6 address.
	Address_IPV6 Address_IpVersion = 2
)

// Enum value maps for Address_IpVersion.
var (
	Address_IpVersion_name = map[int32]string{
		0: "IP_VERSION_UNSPECIFIED",
		1: "IPV4",
		2: "IPV6",
	}
	Address_IpVersion_value = map[string]int32{
		"IP_VERSION_UNSPECIFIED": 0,
		"IPV4":                   1,
		"IPV6":                   2,
	}
)

func (x Address_IpVersion) Enum() *Address_IpVersion {
	p := new(Address_IpVersion)
	*p = x
	return p
}

func (x Address_IpVersion) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Address_IpVersion) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_vpc_v1_address_proto_enumTypes[1].Descriptor()
}

func (Address_IpVersion) Type() protoreflect.EnumType {
	return &file_yandex_cloud_vpc_v1_address_proto_enumTypes[1]
}

func (x Address_IpVersion) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Address_IpVersion.Descriptor instead.
func (Address_IpVersion) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{0, 1}
}

// An Address resource. For more information, see [Address](/docs/vpc/concepts/address).
type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the address. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the address belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the address.
	// The name is unique within the folder.
	// Value must match the regular expression “\|[a-zA-Z]([-_a-zA-Z0-9]{0,61}[a-zA-Z0-9])?“.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the address. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Address labels as `key:value` pairs.
	// No more than 64 per resource.
	// The maximum string length in characters for each value is 63.
	// Each value must match the regular expression `[-_0-9a-z]*`.
	// The string length in characters for each key must be 1-63.
	// Each key must match the regular expression `[a-z][-_0-9a-z]*`.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// External ipv4 address specification.
	//
	// Types that are assignable to Address:
	//
	//	*Address_ExternalIpv4Address
	Address isAddress_Address `protobuf_oneof:"address"`
	// Specifies if address is reserved or not.
	Reserved bool `protobuf:"varint,15,opt,name=reserved,proto3" json:"reserved,omitempty"`
	// Specifies if address is used or not.
	Used bool `protobuf:"varint,16,opt,name=used,proto3" json:"used,omitempty"`
	// Type of the IP address.
	Type Address_Type `protobuf:"varint,17,opt,name=type,proto3,enum=yandex.cloud.vpc.v1.Address_Type" json:"type,omitempty"`
	// Version of the IP address.
	IpVersion Address_IpVersion `protobuf:"varint,18,opt,name=ip_version,json=ipVersion,proto3,enum=yandex.cloud.vpc.v1.Address_IpVersion" json:"ip_version,omitempty"`
	// Specifies if address protected from deletion.
	DeletionProtection bool `protobuf:"varint,19,opt,name=deletion_protection,json=deletionProtection,proto3" json:"deletion_protection,omitempty"`
	// Optional DNS record specifications
	DnsRecords []*DnsRecord `protobuf:"bytes,20,rep,name=dns_records,json=dnsRecords,proto3" json:"dns_records,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Address) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Address) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Address) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Address) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Address) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (m *Address) GetAddress() isAddress_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (x *Address) GetExternalIpv4Address() *ExternalIpv4Address {
	if x, ok := x.GetAddress().(*Address_ExternalIpv4Address); ok {
		return x.ExternalIpv4Address
	}
	return nil
}

func (x *Address) GetReserved() bool {
	if x != nil {
		return x.Reserved
	}
	return false
}

func (x *Address) GetUsed() bool {
	if x != nil {
		return x.Used
	}
	return false
}

func (x *Address) GetType() Address_Type {
	if x != nil {
		return x.Type
	}
	return Address_TYPE_UNSPECIFIED
}

func (x *Address) GetIpVersion() Address_IpVersion {
	if x != nil {
		return x.IpVersion
	}
	return Address_IP_VERSION_UNSPECIFIED
}

func (x *Address) GetDeletionProtection() bool {
	if x != nil {
		return x.DeletionProtection
	}
	return false
}

func (x *Address) GetDnsRecords() []*DnsRecord {
	if x != nil {
		return x.DnsRecords
	}
	return nil
}

type isAddress_Address interface {
	isAddress_Address()
}

type Address_ExternalIpv4Address struct {
	ExternalIpv4Address *ExternalIpv4Address `protobuf:"bytes,7,opt,name=external_ipv4_address,json=externalIpv4Address,proto3,oneof"`
}

func (*Address_ExternalIpv4Address) isAddress_Address() {}

type ExternalIpv4Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Value of address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Availability zone from which the address will be allocated.
	ZoneId string `protobuf:"bytes,2,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Parameters of the allocated address, for example DDoS Protection.
	Requirements *AddressRequirements `protobuf:"bytes,3,opt,name=requirements,proto3" json:"requirements,omitempty"`
}

func (x *ExternalIpv4Address) Reset() {
	*x = ExternalIpv4Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalIpv4Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalIpv4Address) ProtoMessage() {}

func (x *ExternalIpv4Address) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalIpv4Address.ProtoReflect.Descriptor instead.
func (*ExternalIpv4Address) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{1}
}

func (x *ExternalIpv4Address) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ExternalIpv4Address) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *ExternalIpv4Address) GetRequirements() *AddressRequirements {
	if x != nil {
		return x.Requirements
	}
	return nil
}

type AddressRequirements struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DDoS protection provider ID.
	DdosProtectionProvider string `protobuf:"bytes,1,opt,name=ddos_protection_provider,json=ddosProtectionProvider,proto3" json:"ddos_protection_provider,omitempty"`
	// Capability to send SMTP traffic.
	OutgoingSmtpCapability string `protobuf:"bytes,2,opt,name=outgoing_smtp_capability,json=outgoingSmtpCapability,proto3" json:"outgoing_smtp_capability,omitempty"`
}

func (x *AddressRequirements) Reset() {
	*x = AddressRequirements{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressRequirements) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressRequirements) ProtoMessage() {}

func (x *AddressRequirements) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressRequirements.ProtoReflect.Descriptor instead.
func (*AddressRequirements) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{2}
}

func (x *AddressRequirements) GetDdosProtectionProvider() string {
	if x != nil {
		return x.DdosProtectionProvider
	}
	return ""
}

func (x *AddressRequirements) GetOutgoingSmtpCapability() string {
	if x != nil {
		return x.OutgoingSmtpCapability
	}
	return ""
}

type DnsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// DNS record name (absolute or relative to the DNS zone in use).
	Fqdn string `protobuf:"bytes,1,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	// ID of the public DNS zone.
	DnsZoneId string `protobuf:"bytes,2,opt,name=dns_zone_id,json=dnsZoneId,proto3" json:"dns_zone_id,omitempty"`
	// TTL of record.
	Ttl int64 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// If the PTR record is required, this parameter must be set to "true".
	Ptr bool `protobuf:"varint,4,opt,name=ptr,proto3" json:"ptr,omitempty"`
}

func (x *DnsRecord) Reset() {
	*x = DnsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DnsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsRecord) ProtoMessage() {}

func (x *DnsRecord) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_vpc_v1_address_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsRecord.ProtoReflect.Descriptor instead.
func (*DnsRecord) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP(), []int{3}
}

func (x *DnsRecord) GetFqdn() string {
	if x != nil {
		return x.Fqdn
	}
	return ""
}

func (x *DnsRecord) GetDnsZoneId() string {
	if x != nil {
		return x.DnsZoneId
	}
	return ""
}

func (x *DnsRecord) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *DnsRecord) GetPtr() bool {
	if x != nil {
		return x.Ptr
	}
	return false
}

var File_yandex_cloud_vpc_v1_address_proto protoreflect.FileDescriptor

var file_yandex_cloud_vpc_v1_address_proto_rawDesc = []byte{
	0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x06, 0x0a, 0x07, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x5e, 0x0a, 0x15, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x5f, 0x69, 0x70, 0x76, 0x34, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x76, 0x34, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x69,
	0x70, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x49, 0x70,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x69, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x12, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3f, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0a, 0x64, 0x6e, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x38, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45,
	0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x10, 0x02, 0x22, 0x3b, 0x0a, 0x09, 0x49, 0x70, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x50, 0x5f, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x50, 0x56, 0x34, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x49, 0x50, 0x56, 0x36, 0x10, 0x02, 0x42, 0x0f, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x0f, 0x22, 0x96, 0x01,
	0x0a, 0x13, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x70, 0x76, 0x34, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x70,
	0x63, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x89, 0x01, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x38,
	0x0a, 0x18, 0x64, 0x64, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x16, 0x64, 0x64, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x18, 0x6f, 0x75, 0x74, 0x67,
	0x6f, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x6d, 0x74, 0x70, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6f, 0x75, 0x74, 0x67,
	0x6f, 0x69, 0x6e, 0x67, 0x53, 0x6d, 0x74, 0x70, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x63, 0x0a, 0x09, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x71, 0x64, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x71, 0x64, 0x6e, 0x12, 0x1e, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x6e, 0x73, 0x5a, 0x6f, 0x6e,
	0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x74, 0x74, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x70, 0x74, 0x72, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x70, 0x63, 0x2e,
	0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x70, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_vpc_v1_address_proto_rawDescOnce sync.Once
	file_yandex_cloud_vpc_v1_address_proto_rawDescData = file_yandex_cloud_vpc_v1_address_proto_rawDesc
)

func file_yandex_cloud_vpc_v1_address_proto_rawDescGZIP() []byte {
	file_yandex_cloud_vpc_v1_address_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_vpc_v1_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_vpc_v1_address_proto_rawDescData)
	})
	return file_yandex_cloud_vpc_v1_address_proto_rawDescData
}

var file_yandex_cloud_vpc_v1_address_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_vpc_v1_address_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_vpc_v1_address_proto_goTypes = []any{
	(Address_Type)(0),             // 0: yandex.cloud.vpc.v1.Address.Type
	(Address_IpVersion)(0),        // 1: yandex.cloud.vpc.v1.Address.IpVersion
	(*Address)(nil),               // 2: yandex.cloud.vpc.v1.Address
	(*ExternalIpv4Address)(nil),   // 3: yandex.cloud.vpc.v1.ExternalIpv4Address
	(*AddressRequirements)(nil),   // 4: yandex.cloud.vpc.v1.AddressRequirements
	(*DnsRecord)(nil),             // 5: yandex.cloud.vpc.v1.DnsRecord
	nil,                           // 6: yandex.cloud.vpc.v1.Address.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_yandex_cloud_vpc_v1_address_proto_depIdxs = []int32{
	7, // 0: yandex.cloud.vpc.v1.Address.created_at:type_name -> google.protobuf.Timestamp
	6, // 1: yandex.cloud.vpc.v1.Address.labels:type_name -> yandex.cloud.vpc.v1.Address.LabelsEntry
	3, // 2: yandex.cloud.vpc.v1.Address.external_ipv4_address:type_name -> yandex.cloud.vpc.v1.ExternalIpv4Address
	0, // 3: yandex.cloud.vpc.v1.Address.type:type_name -> yandex.cloud.vpc.v1.Address.Type
	1, // 4: yandex.cloud.vpc.v1.Address.ip_version:type_name -> yandex.cloud.vpc.v1.Address.IpVersion
	5, // 5: yandex.cloud.vpc.v1.Address.dns_records:type_name -> yandex.cloud.vpc.v1.DnsRecord
	4, // 6: yandex.cloud.vpc.v1.ExternalIpv4Address.requirements:type_name -> yandex.cloud.vpc.v1.AddressRequirements
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_vpc_v1_address_proto_init() }
func file_yandex_cloud_vpc_v1_address_proto_init() {
	if File_yandex_cloud_vpc_v1_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Address); i {
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
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExternalIpv4Address); i {
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
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*AddressRequirements); i {
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
		file_yandex_cloud_vpc_v1_address_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DnsRecord); i {
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
	file_yandex_cloud_vpc_v1_address_proto_msgTypes[0].OneofWrappers = []any{
		(*Address_ExternalIpv4Address)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_vpc_v1_address_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_vpc_v1_address_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_vpc_v1_address_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_vpc_v1_address_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_vpc_v1_address_proto_msgTypes,
	}.Build()
	File_yandex_cloud_vpc_v1_address_proto = out.File
	file_yandex_cloud_vpc_v1_address_proto_rawDesc = nil
	file_yandex_cloud_vpc_v1_address_proto_goTypes = nil
	file_yandex_cloud_vpc_v1_address_proto_depIdxs = nil
}
