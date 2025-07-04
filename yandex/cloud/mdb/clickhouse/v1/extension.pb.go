// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/clickhouse/v1/extension.proto

package clickhouse

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Extension struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Extension name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. All extension available versions.
	Versions      []*ExtensionVersion `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Extension) Reset() {
	*x = Extension{}
	mi := &file_yandex_cloud_mdb_clickhouse_v1_extension_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Extension) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Extension) ProtoMessage() {}

func (x *Extension) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_clickhouse_v1_extension_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Extension.ProtoReflect.Descriptor instead.
func (*Extension) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescGZIP(), []int{0}
}

func (x *Extension) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Extension) GetVersions() []*ExtensionVersion {
	if x != nil {
		return x.Versions
	}
	return nil
}

type ExtensionVersion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Version ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Is default version.
	Default bool `protobuf:"varint,2,opt,name=default,proto3" json:"default,omitempty"`
	// Is version deprecated.
	Deprecated    bool `protobuf:"varint,3,opt,name=deprecated,proto3" json:"deprecated,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExtensionVersion) Reset() {
	*x = ExtensionVersion{}
	mi := &file_yandex_cloud_mdb_clickhouse_v1_extension_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExtensionVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExtensionVersion) ProtoMessage() {}

func (x *ExtensionVersion) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_clickhouse_v1_extension_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExtensionVersion.ProtoReflect.Descriptor instead.
func (*ExtensionVersion) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescGZIP(), []int{1}
}

func (x *ExtensionVersion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExtensionVersion) GetDefault() bool {
	if x != nil {
		return x.Default
	}
	return false
}

func (x *ExtensionVersion) GetDeprecated() bool {
	if x != nil {
		return x.Deprecated
	}
	return false
}

var File_yandex_cloud_mdb_clickhouse_v1_extension_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDesc = "" +
	"\n" +
	".yandex/cloud/mdb/clickhouse/v1/extension.proto\x12\x1eyandex.cloud.mdb.clickhouse.v1\x1a\x1dyandex/cloud/validation.proto\"s\n" +
	"\tExtension\x12\x18\n" +
	"\x04name\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x04name\x12L\n" +
	"\bversions\x18\x02 \x03(\v20.yandex.cloud.mdb.clickhouse.v1.ExtensionVersionR\bversions\"b\n" +
	"\x10ExtensionVersion\x12\x14\n" +
	"\x02id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x02id\x12\x18\n" +
	"\adefault\x18\x02 \x01(\bR\adefault\x12\x1e\n" +
	"\n" +
	"deprecated\x18\x03 \x01(\bR\n" +
	"deprecatedBs\n" +
	"\"yandex.cloud.api.mdb.clickhouse.v1ZMgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/clickhouse/v1;clickhouseb\x06proto3"

var (
	file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDesc), len(file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDescData
}

var file_yandex_cloud_mdb_clickhouse_v1_extension_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_clickhouse_v1_extension_proto_goTypes = []any{
	(*Extension)(nil),        // 0: yandex.cloud.mdb.clickhouse.v1.Extension
	(*ExtensionVersion)(nil), // 1: yandex.cloud.mdb.clickhouse.v1.ExtensionVersion
}
var file_yandex_cloud_mdb_clickhouse_v1_extension_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.mdb.clickhouse.v1.Extension.versions:type_name -> yandex.cloud.mdb.clickhouse.v1.ExtensionVersion
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_clickhouse_v1_extension_proto_init() }
func file_yandex_cloud_mdb_clickhouse_v1_extension_proto_init() {
	if File_yandex_cloud_mdb_clickhouse_v1_extension_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDesc), len(file_yandex_cloud_mdb_clickhouse_v1_extension_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_clickhouse_v1_extension_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_clickhouse_v1_extension_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_clickhouse_v1_extension_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_clickhouse_v1_extension_proto = out.File
	file_yandex_cloud_mdb_clickhouse_v1_extension_proto_goTypes = nil
	file_yandex_cloud_mdb_clickhouse_v1_extension_proto_depIdxs = nil
}
