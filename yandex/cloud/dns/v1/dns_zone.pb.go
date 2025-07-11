// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/dns/v1/dns_zone.proto

package dns

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A DNS zone. For details about the concept, see [DNS zones](/docs/dns/concepts/dns-zone).
type DnsZone struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the DNS zone. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the DNS zone belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the DNS zone.
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the DNS zone.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// DNS zone labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// DNS zone suffix.
	Zone string `protobuf:"bytes,7,opt,name=zone,proto3" json:"zone,omitempty"`
	// Privately visible zone settings.
	// Specifies whether records within the zone are visible from a VPC networks only.
	PrivateVisibility *PrivateVisibility `protobuf:"bytes,8,opt,name=private_visibility,json=privateVisibility,proto3" json:"private_visibility,omitempty"`
	// Publicly visible zone settings.
	// Indicates whether records within the zone are publicly visible.
	PublicVisibility *PublicVisibility `protobuf:"bytes,9,opt,name=public_visibility,json=publicVisibility,proto3" json:"public_visibility,omitempty"`
	// Prevents accidental zone removal.
	DeletionProtection bool `protobuf:"varint,10,opt,name=deletion_protection,json=deletionProtection,proto3" json:"deletion_protection,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *DnsZone) Reset() {
	*x = DnsZone{}
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DnsZone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DnsZone) ProtoMessage() {}

func (x *DnsZone) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DnsZone.ProtoReflect.Descriptor instead.
func (*DnsZone) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{0}
}

func (x *DnsZone) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DnsZone) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *DnsZone) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DnsZone) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DnsZone) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DnsZone) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DnsZone) GetZone() string {
	if x != nil {
		return x.Zone
	}
	return ""
}

func (x *DnsZone) GetPrivateVisibility() *PrivateVisibility {
	if x != nil {
		return x.PrivateVisibility
	}
	return nil
}

func (x *DnsZone) GetPublicVisibility() *PublicVisibility {
	if x != nil {
		return x.PublicVisibility
	}
	return nil
}

func (x *DnsZone) GetDeletionProtection() bool {
	if x != nil {
		return x.DeletionProtection
	}
	return false
}

// A record set. For details about the concept, see [Resource record](/docs/dns/concepts/resource-record).
type RecordSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Domain name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Record type.
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// Time to live in seconds.
	Ttl int64 `protobuf:"varint,3,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// Data of the record set.
	Data          []string `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RecordSet) Reset() {
	*x = RecordSet{}
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RecordSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordSet) ProtoMessage() {}

func (x *RecordSet) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordSet.ProtoReflect.Descriptor instead.
func (*RecordSet) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{1}
}

func (x *RecordSet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecordSet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RecordSet) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

func (x *RecordSet) GetData() []string {
	if x != nil {
		return x.Data
	}
	return nil
}

// Configuration for privately visible zones.
type PrivateVisibility struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Network IDs.
	NetworkIds    []string `protobuf:"bytes,1,rep,name=network_ids,json=networkIds,proto3" json:"network_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PrivateVisibility) Reset() {
	*x = PrivateVisibility{}
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PrivateVisibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateVisibility) ProtoMessage() {}

func (x *PrivateVisibility) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateVisibility.ProtoReflect.Descriptor instead.
func (*PrivateVisibility) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateVisibility) GetNetworkIds() []string {
	if x != nil {
		return x.NetworkIds
	}
	return nil
}

// Configuration for publicly visible zones.
type PublicVisibility struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PublicVisibility) Reset() {
	*x = PublicVisibility{}
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicVisibility) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicVisibility) ProtoMessage() {}

func (x *PublicVisibility) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicVisibility.ProtoReflect.Descriptor instead.
func (*PublicVisibility) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP(), []int{3}
}

var File_yandex_cloud_dns_v1_dns_zone_proto protoreflect.FileDescriptor

const file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc = "" +
	"\n" +
	"\"yandex/cloud/dns/v1/dns_zone.proto\x12\x13yandex.cloud.dns.v1\x1a\x1dyandex/cloud/validation.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\x94\x04\n" +
	"\aDnsZone\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12@\n" +
	"\x06labels\x18\x06 \x03(\v2(.yandex.cloud.dns.v1.DnsZone.LabelsEntryR\x06labels\x12\x12\n" +
	"\x04zone\x18\a \x01(\tR\x04zone\x12U\n" +
	"\x12private_visibility\x18\b \x01(\v2&.yandex.cloud.dns.v1.PrivateVisibilityR\x11privateVisibility\x12R\n" +
	"\x11public_visibility\x18\t \x01(\v2%.yandex.cloud.dns.v1.PublicVisibilityR\x10publicVisibility\x12/\n" +
	"\x13deletion_protection\x18\n" +
	" \x01(\bR\x12deletionProtection\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\x98\x01\n" +
	"\tRecordSet\x12\x1d\n" +
	"\x04name\x18\x01 \x01(\tB\t\x8a\xc81\x051-254R\x04name\x12\x1c\n" +
	"\x04type\x18\x02 \x01(\tB\b\x8a\xc81\x041-20R\x04type\x12\"\n" +
	"\x03ttl\x18\x03 \x01(\x03B\x10\xfa\xc71\f0-2147483647R\x03ttl\x12*\n" +
	"\x04data\x18\x04 \x03(\tB\x16\x82\xc81\x051-100\x8a\xc81\x051-255\x90\xc81\x01R\x04data\"I\n" +
	"\x11PrivateVisibility\x124\n" +
	"\vnetwork_ids\x18\x01 \x03(\tB\x13\x82\xc81\x050-100\x8a\xc81\x0220\x90\xc81\x01R\n" +
	"networkIds\"\x12\n" +
	"\x10PublicVisibilityBV\n" +
	"\x17yandex.cloud.api.dns.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/dns/v1;dnsb\x06proto3"

var (
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDescOnce sync.Once
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData []byte
)

func file_yandex_cloud_dns_v1_dns_zone_proto_rawDescGZIP() []byte {
	file_yandex_cloud_dns_v1_dns_zone_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc), len(file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc)))
	})
	return file_yandex_cloud_dns_v1_dns_zone_proto_rawDescData
}

var file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_dns_v1_dns_zone_proto_goTypes = []any{
	(*DnsZone)(nil),               // 0: yandex.cloud.dns.v1.DnsZone
	(*RecordSet)(nil),             // 1: yandex.cloud.dns.v1.RecordSet
	(*PrivateVisibility)(nil),     // 2: yandex.cloud.dns.v1.PrivateVisibility
	(*PublicVisibility)(nil),      // 3: yandex.cloud.dns.v1.PublicVisibility
	nil,                           // 4: yandex.cloud.dns.v1.DnsZone.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_yandex_cloud_dns_v1_dns_zone_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.dns.v1.DnsZone.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.dns.v1.DnsZone.labels:type_name -> yandex.cloud.dns.v1.DnsZone.LabelsEntry
	2, // 2: yandex.cloud.dns.v1.DnsZone.private_visibility:type_name -> yandex.cloud.dns.v1.PrivateVisibility
	3, // 3: yandex.cloud.dns.v1.DnsZone.public_visibility:type_name -> yandex.cloud.dns.v1.PublicVisibility
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_dns_v1_dns_zone_proto_init() }
func file_yandex_cloud_dns_v1_dns_zone_proto_init() {
	if File_yandex_cloud_dns_v1_dns_zone_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc), len(file_yandex_cloud_dns_v1_dns_zone_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_dns_v1_dns_zone_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_dns_v1_dns_zone_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_dns_v1_dns_zone_proto_msgTypes,
	}.Build()
	File_yandex_cloud_dns_v1_dns_zone_proto = out.File
	file_yandex_cloud_dns_v1_dns_zone_proto_goTypes = nil
	file_yandex_cloud_dns_v1_dns_zone_proto_depIdxs = nil
}
