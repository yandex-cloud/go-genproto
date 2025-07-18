// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datatransfer/v1/endpoint/yds.proto

package endpoint

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

type YdsCompressionCodec int32

const (
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_UNSPECIFIED YdsCompressionCodec = 0
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_RAW         YdsCompressionCodec = 1
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_GZIP        YdsCompressionCodec = 2
	YdsCompressionCodec_YDS_COMPRESSION_CODEC_ZSTD        YdsCompressionCodec = 4
)

// Enum value maps for YdsCompressionCodec.
var (
	YdsCompressionCodec_name = map[int32]string{
		0: "YDS_COMPRESSION_CODEC_UNSPECIFIED",
		1: "YDS_COMPRESSION_CODEC_RAW",
		2: "YDS_COMPRESSION_CODEC_GZIP",
		4: "YDS_COMPRESSION_CODEC_ZSTD",
	}
	YdsCompressionCodec_value = map[string]int32{
		"YDS_COMPRESSION_CODEC_UNSPECIFIED": 0,
		"YDS_COMPRESSION_CODEC_RAW":         1,
		"YDS_COMPRESSION_CODEC_GZIP":        2,
		"YDS_COMPRESSION_CODEC_ZSTD":        4,
	}
)

func (x YdsCompressionCodec) Enum() *YdsCompressionCodec {
	p := new(YdsCompressionCodec)
	*p = x
	return p
}

func (x YdsCompressionCodec) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (YdsCompressionCodec) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes[0].Descriptor()
}

func (YdsCompressionCodec) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes[0]
}

func (x YdsCompressionCodec) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use YdsCompressionCodec.Descriptor instead.
func (YdsCompressionCodec) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP(), []int{0}
}

type YDSSource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Stream
	Stream string `protobuf:"bytes,2,opt,name=stream,proto3" json:"stream,omitempty"`
	// SA which has read access to the stream.
	ServiceAccountId string `protobuf:"bytes,8,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Compression codec
	SupportedCodecs []YdsCompressionCodec `protobuf:"varint,9,rep,packed,name=supported_codecs,json=supportedCodecs,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec" json:"supported_codecs,omitempty"`
	// Data parsing rules
	Parser *Parser `protobuf:"bytes,10,opt,name=parser,proto3" json:"parser,omitempty"`
	// Should continue working, if consumer read lag exceed TTL of topic
	// False: stop the transfer in error state, if detected lost data. True: continue
	// working with losing part of data
	AllowTtlRewind bool `protobuf:"varint,11,opt,name=allow_ttl_rewind,json=allowTtlRewind,proto3" json:"allow_ttl_rewind,omitempty"`
	// for dedicated db
	Endpoint string `protobuf:"bytes,20,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Network interface for endpoint. If none will assume public ipv4
	SubnetId string `protobuf:"bytes,30,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,34,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	// for important streams
	Consumer      string `protobuf:"bytes,35,opt,name=consumer,proto3" json:"consumer,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YDSSource) Reset() {
	*x = YDSSource{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YDSSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YDSSource) ProtoMessage() {}

func (x *YDSSource) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YDSSource.ProtoReflect.Descriptor instead.
func (*YDSSource) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP(), []int{0}
}

func (x *YDSSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *YDSSource) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *YDSSource) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *YDSSource) GetSupportedCodecs() []YdsCompressionCodec {
	if x != nil {
		return x.SupportedCodecs
	}
	return nil
}

func (x *YDSSource) GetParser() *Parser {
	if x != nil {
		return x.Parser
	}
	return nil
}

func (x *YDSSource) GetAllowTtlRewind() bool {
	if x != nil {
		return x.AllowTtlRewind
	}
	return false
}

func (x *YDSSource) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *YDSSource) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *YDSSource) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *YDSSource) GetConsumer() string {
	if x != nil {
		return x.Consumer
	}
	return ""
}

type YDSTarget struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database
	Database string `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	// Stream
	Stream string `protobuf:"bytes,2,opt,name=stream,proto3" json:"stream,omitempty"`
	// SA which has read access to the stream.
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Save transaction order
	// Not to split events queue into separate per-table queues.
	// Incompatible with setting Topic prefix, only with Topic full name.
	SaveTxOrder      bool                `protobuf:"varint,4,opt,name=save_tx_order,json=saveTxOrder,proto3" json:"save_tx_order,omitempty"`
	CompressionCodec YdsCompressionCodec `protobuf:"varint,5,opt,name=compression_codec,json=compressionCodec,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec" json:"compression_codec,omitempty"`
	// Data serialization format
	Serializer *Serializer `protobuf:"bytes,8,opt,name=serializer,proto3" json:"serializer,omitempty"`
	// for dedicated db
	Endpoint string `protobuf:"bytes,20,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Network interface for endpoint. If none will assume public ipv4
	SubnetId string `protobuf:"bytes,30,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,34,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *YDSTarget) Reset() {
	*x = YDSTarget{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YDSTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YDSTarget) ProtoMessage() {}

func (x *YDSTarget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YDSTarget.ProtoReflect.Descriptor instead.
func (*YDSTarget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP(), []int{1}
}

func (x *YDSTarget) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *YDSTarget) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *YDSTarget) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *YDSTarget) GetSaveTxOrder() bool {
	if x != nil {
		return x.SaveTxOrder
	}
	return false
}

func (x *YDSTarget) GetCompressionCodec() YdsCompressionCodec {
	if x != nil {
		return x.CompressionCodec
	}
	return YdsCompressionCodec_YDS_COMPRESSION_CODEC_UNSPECIFIED
}

func (x *YDSTarget) GetSerializer() *Serializer {
	if x != nil {
		return x.Serializer
	}
	return nil
}

func (x *YDSTarget) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *YDSTarget) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *YDSTarget) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

var File_yandex_cloud_datatransfer_v1_endpoint_yds_proto protoreflect.FileDescriptor

const file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc = "" +
	"\n" +
	"/yandex/cloud/datatransfer/v1/endpoint/yds.proto\x12%yandex.cloud.datatransfer.v1.endpoint\x1a3yandex/cloud/datatransfer/v1/endpoint/parsers.proto\x1a7yandex/cloud/datatransfer/v1/endpoint/serializers.proto\"\xdb\x03\n" +
	"\tYDSSource\x12\x1a\n" +
	"\bdatabase\x18\x01 \x01(\tR\bdatabase\x12\x16\n" +
	"\x06stream\x18\x02 \x01(\tR\x06stream\x12,\n" +
	"\x12service_account_id\x18\b \x01(\tR\x10serviceAccountId\x12e\n" +
	"\x10supported_codecs\x18\t \x03(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodecR\x0fsupportedCodecs\x12E\n" +
	"\x06parser\x18\n" +
	" \x01(\v2-.yandex.cloud.datatransfer.v1.endpoint.ParserR\x06parser\x12(\n" +
	"\x10allow_ttl_rewind\x18\v \x01(\bR\x0eallowTtlRewind\x12\x1a\n" +
	"\bendpoint\x18\x14 \x01(\tR\bendpoint\x12\x1b\n" +
	"\tsubnet_id\x18\x1e \x01(\tR\bsubnetId\x12'\n" +
	"\x0fsecurity_groups\x18\" \x03(\tR\x0esecurityGroups\x12\x1a\n" +
	"\bconsumer\x18# \x01(\tR\bconsumerJ\x04\b\x03\x10\bJ\x04\b\f\x10\x14J\x04\b\x15\x10\x1eJ\x04\b\x1f\x10\"\"\xc7\x03\n" +
	"\tYDSTarget\x12\x1a\n" +
	"\bdatabase\x18\x01 \x01(\tR\bdatabase\x12\x16\n" +
	"\x06stream\x18\x02 \x01(\tR\x06stream\x12,\n" +
	"\x12service_account_id\x18\x03 \x01(\tR\x10serviceAccountId\x12\"\n" +
	"\rsave_tx_order\x18\x04 \x01(\bR\vsaveTxOrder\x12g\n" +
	"\x11compression_codec\x18\x05 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodecR\x10compressionCodec\x12Q\n" +
	"\n" +
	"serializer\x18\b \x01(\v21.yandex.cloud.datatransfer.v1.endpoint.SerializerR\n" +
	"serializer\x12\x1a\n" +
	"\bendpoint\x18\x14 \x01(\tR\bendpoint\x12\x1b\n" +
	"\tsubnet_id\x18\x1e \x01(\tR\bsubnetId\x12'\n" +
	"\x0fsecurity_groups\x18\" \x03(\tR\x0esecurityGroupsJ\x04\b\x06\x10\bJ\x04\b\t\x10\x14J\x04\b\x15\x10\x1eJ\x04\b\x1f\x10\"*\x9b\x01\n" +
	"\x13YdsCompressionCodec\x12%\n" +
	"!YDS_COMPRESSION_CODEC_UNSPECIFIED\x10\x00\x12\x1d\n" +
	"\x19YDS_COMPRESSION_CODEC_RAW\x10\x01\x12\x1e\n" +
	"\x1aYDS_COMPRESSION_CODEC_GZIP\x10\x02\x12\x1e\n" +
	"\x1aYDS_COMPRESSION_CODEC_ZSTD\x10\x04B\xa7\x01\n" +
	")yandex.cloud.api.datatransfer.v1.endpointZRgithub.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1/endpoint;endpoint\xaa\x02%Yandex.Cloud.Datatransfer.V1.EndPointb\x06proto3"

var (
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData []byte
)

func file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc)))
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_goTypes = []any{
	(YdsCompressionCodec)(0), // 0: yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec
	(*YDSSource)(nil),        // 1: yandex.cloud.datatransfer.v1.endpoint.YDSSource
	(*YDSTarget)(nil),        // 2: yandex.cloud.datatransfer.v1.endpoint.YDSTarget
	(*Parser)(nil),           // 3: yandex.cloud.datatransfer.v1.endpoint.Parser
	(*Serializer)(nil),       // 4: yandex.cloud.datatransfer.v1.endpoint.Serializer
}
var file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.datatransfer.v1.endpoint.YDSSource.supported_codecs:type_name -> yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec
	3, // 1: yandex.cloud.datatransfer.v1.endpoint.YDSSource.parser:type_name -> yandex.cloud.datatransfer.v1.endpoint.Parser
	0, // 2: yandex.cloud.datatransfer.v1.endpoint.YDSTarget.compression_codec:type_name -> yandex.cloud.datatransfer.v1.endpoint.YdsCompressionCodec
	4, // 3: yandex.cloud.datatransfer.v1.endpoint.YDSTarget.serializer:type_name -> yandex.cloud.datatransfer.v1.endpoint.Serializer
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_yds_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_parsers_proto_init()
	file_yandex_cloud_datatransfer_v1_endpoint_serializers_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_yds_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_yds_proto_depIdxs = nil
}
