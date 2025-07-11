// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/opensearch/v1/config/opensearch.proto

package opensearch

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type OpenSearchConfig2 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// the maximum number of allowed boolean clauses in a query
	MaxClauseCount *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=max_clause_count,json=maxClauseCount,proto3" json:"max_clause_count,omitempty"`
	// the percentage or absolute value (10%, 512mb) of heap space that is allocated to fielddata
	FielddataCacheSize     string `protobuf:"bytes,4,opt,name=fielddata_cache_size,json=fielddataCacheSize,proto3" json:"fielddata_cache_size,omitempty"`
	ReindexRemoteWhitelist string `protobuf:"bytes,6,opt,name=reindex_remote_whitelist,json=reindexRemoteWhitelist,proto3" json:"reindex_remote_whitelist,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *OpenSearchConfig2) Reset() {
	*x = OpenSearchConfig2{}
	mi := &file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenSearchConfig2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSearchConfig2) ProtoMessage() {}

func (x *OpenSearchConfig2) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSearchConfig2.ProtoReflect.Descriptor instead.
func (*OpenSearchConfig2) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescGZIP(), []int{0}
}

func (x *OpenSearchConfig2) GetMaxClauseCount() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxClauseCount
	}
	return nil
}

func (x *OpenSearchConfig2) GetFielddataCacheSize() string {
	if x != nil {
		return x.FielddataCacheSize
	}
	return ""
}

func (x *OpenSearchConfig2) GetReindexRemoteWhitelist() string {
	if x != nil {
		return x.ReindexRemoteWhitelist
	}
	return ""
}

type OpenSearchConfigSet2 struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	EffectiveConfig *OpenSearchConfig2     `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	UserConfig      *OpenSearchConfig2     `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	DefaultConfig   *OpenSearchConfig2     `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *OpenSearchConfigSet2) Reset() {
	*x = OpenSearchConfigSet2{}
	mi := &file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OpenSearchConfigSet2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OpenSearchConfigSet2) ProtoMessage() {}

func (x *OpenSearchConfigSet2) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OpenSearchConfigSet2.ProtoReflect.Descriptor instead.
func (*OpenSearchConfigSet2) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescGZIP(), []int{1}
}

func (x *OpenSearchConfigSet2) GetEffectiveConfig() *OpenSearchConfig2 {
	if x != nil {
		return x.EffectiveConfig
	}
	return nil
}

func (x *OpenSearchConfigSet2) GetUserConfig() *OpenSearchConfig2 {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

func (x *OpenSearchConfigSet2) GetDefaultConfig() *OpenSearchConfig2 {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

var File_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDesc = "" +
	"\n" +
	"6yandex/cloud/mdb/opensearch/v1/config/opensearch.proto\x12%yandex.cloud.mdb.opensearch.v1.config\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1dyandex/cloud/validation.proto\"\xcc\x01\n" +
	"\x11OpenSearchConfig2\x12E\n" +
	"\x10max_clause_count\x18\x03 \x01(\v2\x1b.google.protobuf.Int64ValueR\x0emaxClauseCount\x120\n" +
	"\x14fielddata_cache_size\x18\x04 \x01(\tR\x12fielddataCacheSize\x128\n" +
	"\x18reindex_remote_whitelist\x18\x06 \x01(\tR\x16reindexRemoteWhitelistJ\x04\b\x05\x10\x06\"\xbd\x02\n" +
	"\x14OpenSearchConfigSet2\x12i\n" +
	"\x10effective_config\x18\x01 \x01(\v28.yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2B\x04\xe8\xc71\x01R\x0feffectiveConfig\x12Y\n" +
	"\vuser_config\x18\x02 \x01(\v28.yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2R\n" +
	"userConfig\x12_\n" +
	"\x0edefault_config\x18\x03 \x01(\v28.yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2R\rdefaultConfigBz\n" +
	"\"yandex.cloud.api.mdb.opensearch.v1ZTgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1/config;opensearchb\x06proto3"

var (
	file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDesc), len(file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDescData
}

var file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_goTypes = []any{
	(*OpenSearchConfig2)(nil),     // 0: yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2
	(*OpenSearchConfigSet2)(nil),  // 1: yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfigSet2
	(*wrapperspb.Int64Value)(nil), // 2: google.protobuf.Int64Value
}
var file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2.max_clause_count:type_name -> google.protobuf.Int64Value
	0, // 1: yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfigSet2.effective_config:type_name -> yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2
	0, // 2: yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfigSet2.user_config:type_name -> yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2
	0, // 3: yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfigSet2.default_config:type_name -> yandex.cloud.mdb.opensearch.v1.config.OpenSearchConfig2
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_init() }
func file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_init() {
	if File_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDesc), len(file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto = out.File
	file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_goTypes = nil
	file_yandex_cloud_mdb_opensearch_v1_config_opensearch_proto_depIdxs = nil
}
