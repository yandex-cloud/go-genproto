// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/cdn/v1/origin.proto

package cdn

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// An origin. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type Origin struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the origin.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the parent origin group.
	OriginGroupId int64 `protobuf:"varint,2,opt,name=origin_group_id,json=originGroupId,proto3" json:"origin_group_id,omitempty"`
	// IP address or Domain name of your origin and the port (if custom).
	// Used if [meta] variant is `common`.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// The setting allows to enable or disable an Origin source in the Origins group.
	//
	// It has two possible values:
	//
	// True - The origin is enabled and used as a source for the CDN. An origins
	// group must contain at least one enabled origin.
	// False - The origin is disabled and the CDN is not using it to pull content.
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Specifies whether the origin is used in its origin group as backup.
	// A backup origin is used when one of active origins becomes unavailable.
	Backup bool `protobuf:"varint,5,opt,name=backup,proto3" json:"backup,omitempty"`
	// Set up origin of the content.
	Meta          *OriginMeta `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Origin) Reset() {
	*x = Origin{}
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{0}
}

func (x *Origin) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Origin) GetOriginGroupId() int64 {
	if x != nil {
		return x.OriginGroupId
	}
	return 0
}

func (x *Origin) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Origin) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Origin) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *Origin) GetMeta() *OriginMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Origin parameters. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginParams struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Source: IP address or Domain name of your origin and the port (if custom).
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// The setting allows to enable or disable an Origin source in the Origins group.
	//
	// It has two possible values:
	//
	// True - The origin is enabled and used as a source for the CDN. An origins
	// group must contain at least one enabled origins. False - The origin is disabled
	// and the CDN is not using it to pull content.
	Enabled bool `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// backup option has two possible values:
	//
	//	True - The option is active. The origin will not be used until one of
	//	       active origins become unavailable.
	//	False - The option is disabled.
	Backup bool `protobuf:"varint,3,opt,name=backup,proto3" json:"backup,omitempty"`
	// Set up origin of the content.
	Meta          *OriginMeta `protobuf:"bytes,4,opt,name=meta,proto3" json:"meta,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OriginParams) Reset() {
	*x = OriginParams{}
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OriginParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginParams) ProtoMessage() {}

func (x *OriginParams) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginParams.ProtoReflect.Descriptor instead.
func (*OriginParams) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{1}
}

func (x *OriginParams) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *OriginParams) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *OriginParams) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *OriginParams) GetMeta() *OriginMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

// Origin type. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginMeta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Type of the origin.
	//
	// Types that are valid to be assigned to OriginMetaVariant:
	//
	//	*OriginMeta_Common
	//	*OriginMeta_Bucket
	//	*OriginMeta_Website
	//	*OriginMeta_Balancer
	OriginMetaVariant isOriginMeta_OriginMetaVariant `protobuf_oneof:"origin_meta_variant"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *OriginMeta) Reset() {
	*x = OriginMeta{}
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OriginMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginMeta) ProtoMessage() {}

func (x *OriginMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginMeta.ProtoReflect.Descriptor instead.
func (*OriginMeta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{2}
}

func (x *OriginMeta) GetOriginMetaVariant() isOriginMeta_OriginMetaVariant {
	if x != nil {
		return x.OriginMetaVariant
	}
	return nil
}

func (x *OriginMeta) GetCommon() *OriginNamedMeta {
	if x != nil {
		if x, ok := x.OriginMetaVariant.(*OriginMeta_Common); ok {
			return x.Common
		}
	}
	return nil
}

func (x *OriginMeta) GetBucket() *OriginNamedMeta {
	if x != nil {
		if x, ok := x.OriginMetaVariant.(*OriginMeta_Bucket); ok {
			return x.Bucket
		}
	}
	return nil
}

func (x *OriginMeta) GetWebsite() *OriginNamedMeta {
	if x != nil {
		if x, ok := x.OriginMetaVariant.(*OriginMeta_Website); ok {
			return x.Website
		}
	}
	return nil
}

func (x *OriginMeta) GetBalancer() *OriginBalancerMeta {
	if x != nil {
		if x, ok := x.OriginMetaVariant.(*OriginMeta_Balancer); ok {
			return x.Balancer
		}
	}
	return nil
}

type isOriginMeta_OriginMetaVariant interface {
	isOriginMeta_OriginMetaVariant()
}

type OriginMeta_Common struct {
	// A server with a domain name linked to it
	Common *OriginNamedMeta `protobuf:"bytes,1,opt,name=common,proto3,oneof"`
}

type OriginMeta_Bucket struct {
	// An Object Storage bucket not configured as a static site hosting.
	Bucket *OriginNamedMeta `protobuf:"bytes,2,opt,name=bucket,proto3,oneof"`
}

type OriginMeta_Website struct {
	// An Object Storage bucket configured as a static site hosting.
	Website *OriginNamedMeta `protobuf:"bytes,3,opt,name=website,proto3,oneof"`
}

type OriginMeta_Balancer struct {
	// An L7 load balancer from Application Load Balancer.
	// CDN servers will access the load balancer at one of its IP addresses that must be selected in the origin settings.
	Balancer *OriginBalancerMeta `protobuf:"bytes,4,opt,name=balancer,proto3,oneof"`
}

func (*OriginMeta_Common) isOriginMeta_OriginMetaVariant() {}

func (*OriginMeta_Bucket) isOriginMeta_OriginMetaVariant() {}

func (*OriginMeta_Website) isOriginMeta_OriginMetaVariant() {}

func (*OriginMeta_Balancer) isOriginMeta_OriginMetaVariant() {}

// Origin info. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginNamedMeta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the origin.
	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OriginNamedMeta) Reset() {
	*x = OriginNamedMeta{}
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OriginNamedMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginNamedMeta) ProtoMessage() {}

func (x *OriginNamedMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginNamedMeta.ProtoReflect.Descriptor instead.
func (*OriginNamedMeta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{3}
}

func (x *OriginNamedMeta) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Application Load Balancer origin info. For details about the concept, see [documentation](/docs/cdn/concepts/origins).
type OriginBalancerMeta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the origin.
	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OriginBalancerMeta) Reset() {
	*x = OriginBalancerMeta{}
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OriginBalancerMeta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OriginBalancerMeta) ProtoMessage() {}

func (x *OriginBalancerMeta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OriginBalancerMeta.ProtoReflect.Descriptor instead.
func (*OriginBalancerMeta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP(), []int{4}
}

func (x *OriginBalancerMeta) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_yandex_cloud_cdn_v1_origin_proto protoreflect.FileDescriptor

const file_yandex_cloud_cdn_v1_origin_proto_rawDesc = "" +
	"\n" +
	" yandex/cloud/cdn/v1/origin.proto\x12\x13yandex.cloud.cdn.v1\"\xbf\x01\n" +
	"\x06Origin\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12&\n" +
	"\x0forigin_group_id\x18\x02 \x01(\x03R\roriginGroupId\x12\x16\n" +
	"\x06source\x18\x03 \x01(\tR\x06source\x12\x18\n" +
	"\aenabled\x18\x04 \x01(\bR\aenabled\x12\x16\n" +
	"\x06backup\x18\x05 \x01(\bR\x06backup\x123\n" +
	"\x04meta\x18\x06 \x01(\v2\x1f.yandex.cloud.cdn.v1.OriginMetaR\x04meta\"\x8d\x01\n" +
	"\fOriginParams\x12\x16\n" +
	"\x06source\x18\x01 \x01(\tR\x06source\x12\x18\n" +
	"\aenabled\x18\x02 \x01(\bR\aenabled\x12\x16\n" +
	"\x06backup\x18\x03 \x01(\bR\x06backup\x123\n" +
	"\x04meta\x18\x04 \x01(\v2\x1f.yandex.cloud.cdn.v1.OriginMetaR\x04meta\"\xac\x02\n" +
	"\n" +
	"OriginMeta\x12>\n" +
	"\x06common\x18\x01 \x01(\v2$.yandex.cloud.cdn.v1.OriginNamedMetaH\x00R\x06common\x12>\n" +
	"\x06bucket\x18\x02 \x01(\v2$.yandex.cloud.cdn.v1.OriginNamedMetaH\x00R\x06bucket\x12@\n" +
	"\awebsite\x18\x03 \x01(\v2$.yandex.cloud.cdn.v1.OriginNamedMetaH\x00R\awebsite\x12E\n" +
	"\bbalancer\x18\x04 \x01(\v2'.yandex.cloud.cdn.v1.OriginBalancerMetaH\x00R\bbalancerB\x15\n" +
	"\x13origin_meta_variant\"%\n" +
	"\x0fOriginNamedMeta\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\"$\n" +
	"\x12OriginBalancerMeta\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02idBV\n" +
	"\x17yandex.cloud.api.cdn.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/cdn/v1;cdnb\x06proto3"

var (
	file_yandex_cloud_cdn_v1_origin_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_origin_proto_rawDescData []byte
)

func file_yandex_cloud_cdn_v1_origin_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_origin_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_origin_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_cdn_v1_origin_proto_rawDesc), len(file_yandex_cloud_cdn_v1_origin_proto_rawDesc)))
	})
	return file_yandex_cloud_cdn_v1_origin_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_origin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_cdn_v1_origin_proto_goTypes = []any{
	(*Origin)(nil),             // 0: yandex.cloud.cdn.v1.Origin
	(*OriginParams)(nil),       // 1: yandex.cloud.cdn.v1.OriginParams
	(*OriginMeta)(nil),         // 2: yandex.cloud.cdn.v1.OriginMeta
	(*OriginNamedMeta)(nil),    // 3: yandex.cloud.cdn.v1.OriginNamedMeta
	(*OriginBalancerMeta)(nil), // 4: yandex.cloud.cdn.v1.OriginBalancerMeta
}
var file_yandex_cloud_cdn_v1_origin_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.cdn.v1.Origin.meta:type_name -> yandex.cloud.cdn.v1.OriginMeta
	2, // 1: yandex.cloud.cdn.v1.OriginParams.meta:type_name -> yandex.cloud.cdn.v1.OriginMeta
	3, // 2: yandex.cloud.cdn.v1.OriginMeta.common:type_name -> yandex.cloud.cdn.v1.OriginNamedMeta
	3, // 3: yandex.cloud.cdn.v1.OriginMeta.bucket:type_name -> yandex.cloud.cdn.v1.OriginNamedMeta
	3, // 4: yandex.cloud.cdn.v1.OriginMeta.website:type_name -> yandex.cloud.cdn.v1.OriginNamedMeta
	4, // 5: yandex.cloud.cdn.v1.OriginMeta.balancer:type_name -> yandex.cloud.cdn.v1.OriginBalancerMeta
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_origin_proto_init() }
func file_yandex_cloud_cdn_v1_origin_proto_init() {
	if File_yandex_cloud_cdn_v1_origin_proto != nil {
		return
	}
	file_yandex_cloud_cdn_v1_origin_proto_msgTypes[2].OneofWrappers = []any{
		(*OriginMeta_Common)(nil),
		(*OriginMeta_Bucket)(nil),
		(*OriginMeta_Website)(nil),
		(*OriginMeta_Balancer)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_cdn_v1_origin_proto_rawDesc), len(file_yandex_cloud_cdn_v1_origin_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_origin_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_origin_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_origin_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_origin_proto = out.File
	file_yandex_cloud_cdn_v1_origin_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_origin_proto_depIdxs = nil
}
