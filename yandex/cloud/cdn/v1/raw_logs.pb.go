// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/cdn/v1/raw_logs.proto

package cdn

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

// Provider side statuses of Raw logs processing.
type RawLogsStatus int32

const (
	RawLogsStatus_RAW_LOGS_STATUS_UNSPECIFIED RawLogsStatus = 0
	// Raw logs wasn't activated.
	RawLogsStatus_RAW_LOGS_STATUS_NOT_ACTIVATED RawLogsStatus = 1
	// Raw logs was activated, and logs storing process works as expected.
	RawLogsStatus_RAW_LOGS_STATUS_OK RawLogsStatus = 2
	// Raw logs was activated, but CDN provider has been failed to store logs.
	RawLogsStatus_RAW_LOGS_STATUS_FAILED RawLogsStatus = 3
	// Raw logs was activated, but logs storing process is expected.
	RawLogsStatus_RAW_LOGS_STATUS_PENDING RawLogsStatus = 4
)

// Enum value maps for RawLogsStatus.
var (
	RawLogsStatus_name = map[int32]string{
		0: "RAW_LOGS_STATUS_UNSPECIFIED",
		1: "RAW_LOGS_STATUS_NOT_ACTIVATED",
		2: "RAW_LOGS_STATUS_OK",
		3: "RAW_LOGS_STATUS_FAILED",
		4: "RAW_LOGS_STATUS_PENDING",
	}
	RawLogsStatus_value = map[string]int32{
		"RAW_LOGS_STATUS_UNSPECIFIED":   0,
		"RAW_LOGS_STATUS_NOT_ACTIVATED": 1,
		"RAW_LOGS_STATUS_OK":            2,
		"RAW_LOGS_STATUS_FAILED":        3,
		"RAW_LOGS_STATUS_PENDING":       4,
	}
)

func (x RawLogsStatus) Enum() *RawLogsStatus {
	p := new(RawLogsStatus)
	*p = x
	return p
}

func (x RawLogsStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RawLogsStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_cdn_v1_raw_logs_proto_enumTypes[0].Descriptor()
}

func (RawLogsStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_cdn_v1_raw_logs_proto_enumTypes[0]
}

func (x RawLogsStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RawLogsStatus.Descriptor instead.
func (RawLogsStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescGZIP(), []int{0}
}

// User settings for Raw logs.
type RawLogsSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Destination S3 bucket name, note that the suer should be owner of the bucket.
	BucketName string `protobuf:"bytes,1,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	// Bucket region, unused for now, could be blank.
	BucketRegion string `protobuf:"bytes,2,opt,name=bucket_region,json=bucketRegion,proto3" json:"bucket_region,omitempty"`
	// file_prefix: prefix each log object name with specified prefix.
	//
	// The prefix makes it simpler for you to locate the log objects.
	// For example, if you specify the prefix value logs/, each log object that
	// S3 creates begins with the logs/ prefix in its key, so pseudo S3 folders
	// could be setup.
	FilePrefix    string `protobuf:"bytes,3,opt,name=file_prefix,json=filePrefix,proto3" json:"file_prefix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RawLogsSettings) Reset() {
	*x = RawLogsSettings{}
	mi := &file_yandex_cloud_cdn_v1_raw_logs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RawLogsSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawLogsSettings) ProtoMessage() {}

func (x *RawLogsSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_raw_logs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawLogsSettings.ProtoReflect.Descriptor instead.
func (*RawLogsSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescGZIP(), []int{0}
}

func (x *RawLogsSettings) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *RawLogsSettings) GetBucketRegion() string {
	if x != nil {
		return x.BucketRegion
	}
	return ""
}

func (x *RawLogsSettings) GetFilePrefix() string {
	if x != nil {
		return x.FilePrefix
	}
	return ""
}

var File_yandex_cloud_cdn_v1_raw_logs_proto protoreflect.FileDescriptor

const file_yandex_cloud_cdn_v1_raw_logs_proto_rawDesc = "" +
	"\n" +
	"\"yandex/cloud/cdn/v1/raw_logs.proto\x12\x13yandex.cloud.cdn.v1\x1a\x1dyandex/cloud/validation.proto\"\x9c\x01\n" +
	"\x0fRawLogsSettings\x12/\n" +
	"\vbucket_name\x18\x01 \x01(\tB\x0e\xe8\xc71\x01\x8a\xc81\x06<=1024R\n" +
	"bucketName\x12-\n" +
	"\rbucket_region\x18\x02 \x01(\tB\b\x8a\xc81\x04<=50R\fbucketRegion\x12)\n" +
	"\vfile_prefix\x18\x03 \x01(\tB\b\x8a\xc81\x04<=50R\n" +
	"filePrefix*\xa4\x01\n" +
	"\rRawLogsStatus\x12\x1f\n" +
	"\x1bRAW_LOGS_STATUS_UNSPECIFIED\x10\x00\x12!\n" +
	"\x1dRAW_LOGS_STATUS_NOT_ACTIVATED\x10\x01\x12\x16\n" +
	"\x12RAW_LOGS_STATUS_OK\x10\x02\x12\x1a\n" +
	"\x16RAW_LOGS_STATUS_FAILED\x10\x03\x12\x1b\n" +
	"\x17RAW_LOGS_STATUS_PENDING\x10\x04BV\n" +
	"\x17yandex.cloud.api.cdn.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/cdn/v1;cdnb\x06proto3"

var (
	file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescData []byte
)

func file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_cdn_v1_raw_logs_proto_rawDesc), len(file_yandex_cloud_cdn_v1_raw_logs_proto_rawDesc)))
	})
	return file_yandex_cloud_cdn_v1_raw_logs_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_raw_logs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_cdn_v1_raw_logs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_cdn_v1_raw_logs_proto_goTypes = []any{
	(RawLogsStatus)(0),      // 0: yandex.cloud.cdn.v1.RawLogsStatus
	(*RawLogsSettings)(nil), // 1: yandex.cloud.cdn.v1.RawLogsSettings
}
var file_yandex_cloud_cdn_v1_raw_logs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_raw_logs_proto_init() }
func file_yandex_cloud_cdn_v1_raw_logs_proto_init() {
	if File_yandex_cloud_cdn_v1_raw_logs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_cdn_v1_raw_logs_proto_rawDesc), len(file_yandex_cloud_cdn_v1_raw_logs_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_raw_logs_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_raw_logs_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_cdn_v1_raw_logs_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_cdn_v1_raw_logs_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_raw_logs_proto = out.File
	file_yandex_cloud_cdn_v1_raw_logs_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_raw_logs_proto_depIdxs = nil
}
