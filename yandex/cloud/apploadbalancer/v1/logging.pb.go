// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/apploadbalancer/v1/logging.proto

package apploadbalancer

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	code "google.golang.org/genproto/googleapis/rpc/code"
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

type HttpCodeInterval int32

const (
	HttpCodeInterval_HTTP_CODE_INTERVAL_UNSPECIFIED HttpCodeInterval = 0
	HttpCodeInterval_HTTP_1XX                       HttpCodeInterval = 1
	HttpCodeInterval_HTTP_2XX                       HttpCodeInterval = 2
	HttpCodeInterval_HTTP_3XX                       HttpCodeInterval = 3
	HttpCodeInterval_HTTP_4XX                       HttpCodeInterval = 4
	HttpCodeInterval_HTTP_5XX                       HttpCodeInterval = 5
	HttpCodeInterval_HTTP_ALL                       HttpCodeInterval = 6
)

// Enum value maps for HttpCodeInterval.
var (
	HttpCodeInterval_name = map[int32]string{
		0: "HTTP_CODE_INTERVAL_UNSPECIFIED",
		1: "HTTP_1XX",
		2: "HTTP_2XX",
		3: "HTTP_3XX",
		4: "HTTP_4XX",
		5: "HTTP_5XX",
		6: "HTTP_ALL",
	}
	HttpCodeInterval_value = map[string]int32{
		"HTTP_CODE_INTERVAL_UNSPECIFIED": 0,
		"HTTP_1XX":                       1,
		"HTTP_2XX":                       2,
		"HTTP_3XX":                       3,
		"HTTP_4XX":                       4,
		"HTTP_5XX":                       5,
		"HTTP_ALL":                       6,
	}
)

func (x HttpCodeInterval) Enum() *HttpCodeInterval {
	p := new(HttpCodeInterval)
	*p = x
	return p
}

func (x HttpCodeInterval) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpCodeInterval) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_apploadbalancer_v1_logging_proto_enumTypes[0].Descriptor()
}

func (HttpCodeInterval) Type() protoreflect.EnumType {
	return &file_yandex_cloud_apploadbalancer_v1_logging_proto_enumTypes[0]
}

func (x HttpCodeInterval) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpCodeInterval.Descriptor instead.
func (HttpCodeInterval) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescGZIP(), []int{0}
}

// LogDiscardRule discards a fraction of logs with certain codes.
// If neither codes or intervals are provided, rule applies to all logs.
type LogDiscardRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// HTTP codes that should be discarded.
	HttpCodes []int64 `protobuf:"varint,1,rep,packed,name=http_codes,json=httpCodes,proto3" json:"http_codes,omitempty"`
	// Groups of HTTP codes like 4xx that should be discarded.
	HttpCodeIntervals []HttpCodeInterval `protobuf:"varint,2,rep,packed,name=http_code_intervals,json=httpCodeIntervals,proto3,enum=yandex.cloud.apploadbalancer.v1.HttpCodeInterval" json:"http_code_intervals,omitempty"`
	// GRPC codes that should be discarded
	GrpcCodes []code.Code `protobuf:"varint,3,rep,packed,name=grpc_codes,json=grpcCodes,proto3,enum=google.rpc.Code" json:"grpc_codes,omitempty"`
	// Percent of logs to be discarded: 0 - keep all, 100 or unset - discard all
	DiscardPercent *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=discard_percent,json=discardPercent,proto3" json:"discard_percent,omitempty"`
}

func (x *LogDiscardRule) Reset() {
	*x = LogDiscardRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogDiscardRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogDiscardRule) ProtoMessage() {}

func (x *LogDiscardRule) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogDiscardRule.ProtoReflect.Descriptor instead.
func (*LogDiscardRule) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescGZIP(), []int{0}
}

func (x *LogDiscardRule) GetHttpCodes() []int64 {
	if x != nil {
		return x.HttpCodes
	}
	return nil
}

func (x *LogDiscardRule) GetHttpCodeIntervals() []HttpCodeInterval {
	if x != nil {
		return x.HttpCodeIntervals
	}
	return nil
}

func (x *LogDiscardRule) GetGrpcCodes() []code.Code {
	if x != nil {
		return x.GrpcCodes
	}
	return nil
}

func (x *LogDiscardRule) GetDiscardPercent() *wrapperspb.Int64Value {
	if x != nil {
		return x.DiscardPercent
	}
	return nil
}

type LogOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Cloud Logging log group ID to store access logs.
	// If not set then logs will be stored in default log group for the folder
	// where load balancer located.
	LogGroupId string `protobuf:"bytes,1,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// ordered list of rules, first matching rule applies
	DiscardRules []*LogDiscardRule `protobuf:"bytes,2,rep,name=discard_rules,json=discardRules,proto3" json:"discard_rules,omitempty"`
	// Do not send logs to Cloud Logging log group.
	Disable bool `protobuf:"varint,3,opt,name=disable,proto3" json:"disable,omitempty"`
}

func (x *LogOptions) Reset() {
	*x = LogOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogOptions) ProtoMessage() {}

func (x *LogOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogOptions.ProtoReflect.Descriptor instead.
func (*LogOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescGZIP(), []int{1}
}

func (x *LogOptions) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

func (x *LogOptions) GetDiscardRules() []*LogDiscardRule {
	if x != nil {
		return x.DiscardRules
	}
	return nil
}

func (x *LogOptions) GetDisable() bool {
	if x != nil {
		return x.Disable
	}
	return false
}

var File_yandex_cloud_apploadbalancer_v1_logging_proto protoreflect.FileDescriptor

var file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x44, 0x69,
	0x73, 0x63, 0x61, 0x72, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x03, 0x42, 0x0b, 0xfa,
	0xc7, 0x31, 0x07, 0x31, 0x30, 0x30, 0x2d, 0x35, 0x39, 0x39, 0x52, 0x09, 0x68, 0x74, 0x74, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x61, 0x0a, 0x13, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0e, 0x32, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x52, 0x11, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x09,
	0x67, 0x72, 0x70, 0x63, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x4f, 0x0a, 0x0f, 0x64, 0x69, 0x73,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x09, 0xfa, 0xc7, 0x31, 0x05, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63,
	0x61, 0x72, 0x64, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x22, 0x9e, 0x01, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x67,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6c, 0x6f, 0x67, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x0d, 0x64,
	0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x44, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x52,
	0x75, 0x6c, 0x65, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x61, 0x72, 0x64, 0x52, 0x75, 0x6c, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x64, 0x69, 0x73, 0x61, 0x62, 0x6c, 0x65, 0x2a, 0x8a, 0x01, 0x0a, 0x10,
	0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x12, 0x22, 0x0a, 0x1e, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x43, 0x4f, 0x44, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x45, 0x52, 0x56, 0x41, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x31, 0x58, 0x58,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x32, 0x58, 0x58, 0x10, 0x02,
	0x12, 0x0c, 0x0a, 0x08, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x33, 0x58, 0x58, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x48, 0x54, 0x54, 0x50, 0x5f, 0x34, 0x58, 0x58, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08,
	0x48, 0x54, 0x54, 0x50, 0x5f, 0x35, 0x58, 0x58, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x48, 0x54,
	0x54, 0x50, 0x5f, 0x41, 0x4c, 0x4c, 0x10, 0x06, 0x42, 0x7a, 0x0a, 0x23, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x5a,
	0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescData = file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDesc
)

func file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescData)
	})
	return file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDescData
}

var file_yandex_cloud_apploadbalancer_v1_logging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_apploadbalancer_v1_logging_proto_goTypes = []interface{}{
	(HttpCodeInterval)(0),         // 0: yandex.cloud.apploadbalancer.v1.HttpCodeInterval
	(*LogDiscardRule)(nil),        // 1: yandex.cloud.apploadbalancer.v1.LogDiscardRule
	(*LogOptions)(nil),            // 2: yandex.cloud.apploadbalancer.v1.LogOptions
	(code.Code)(0),                // 3: google.rpc.Code
	(*wrapperspb.Int64Value)(nil), // 4: google.protobuf.Int64Value
}
var file_yandex_cloud_apploadbalancer_v1_logging_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.apploadbalancer.v1.LogDiscardRule.http_code_intervals:type_name -> yandex.cloud.apploadbalancer.v1.HttpCodeInterval
	3, // 1: yandex.cloud.apploadbalancer.v1.LogDiscardRule.grpc_codes:type_name -> google.rpc.Code
	4, // 2: yandex.cloud.apploadbalancer.v1.LogDiscardRule.discard_percent:type_name -> google.protobuf.Int64Value
	1, // 3: yandex.cloud.apploadbalancer.v1.LogOptions.discard_rules:type_name -> yandex.cloud.apploadbalancer.v1.LogDiscardRule
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_apploadbalancer_v1_logging_proto_init() }
func file_yandex_cloud_apploadbalancer_v1_logging_proto_init() {
	if File_yandex_cloud_apploadbalancer_v1_logging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogDiscardRule); i {
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
		file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogOptions); i {
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
			RawDescriptor: file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_apploadbalancer_v1_logging_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_apploadbalancer_v1_logging_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_apploadbalancer_v1_logging_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_apploadbalancer_v1_logging_proto_msgTypes,
	}.Build()
	File_yandex_cloud_apploadbalancer_v1_logging_proto = out.File
	file_yandex_cloud_apploadbalancer_v1_logging_proto_rawDesc = nil
	file_yandex_cloud_apploadbalancer_v1_logging_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_logging_proto_depIdxs = nil
}
