// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/logging/v1/sink.proto

package logging

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

type Sink struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Sink ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Sink folder ID.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Sink cloud ID.
	CloudId string `protobuf:"bytes,3,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// Sink creation time.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Sink name.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Sink description.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Sink labels.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Logs will be written to the sink on behalf of this service account
	ServiceAccountId string `protobuf:"bytes,8,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Logs destination
	//
	// Types that are valid to be assigned to Sink:
	//
	//	*Sink_Yds_
	//	*Sink_S3_
	Sink          isSink_Sink `protobuf_oneof:"sink"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sink) Reset() {
	*x = Sink{}
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink.ProtoReflect.Descriptor instead.
func (*Sink) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_sink_proto_rawDescGZIP(), []int{0}
}

func (x *Sink) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Sink) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Sink) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *Sink) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Sink) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Sink) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Sink) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Sink) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Sink) GetSink() isSink_Sink {
	if x != nil {
		return x.Sink
	}
	return nil
}

func (x *Sink) GetYds() *Sink_Yds {
	if x != nil {
		if x, ok := x.Sink.(*Sink_Yds_); ok {
			return x.Yds
		}
	}
	return nil
}

func (x *Sink) GetS3() *Sink_S3 {
	if x != nil {
		if x, ok := x.Sink.(*Sink_S3_); ok {
			return x.S3
		}
	}
	return nil
}

type isSink_Sink interface {
	isSink_Sink()
}

type Sink_Yds_ struct {
	// Yandex data stream
	Yds *Sink_Yds `protobuf:"bytes,9,opt,name=yds,proto3,oneof"`
}

type Sink_S3_ struct {
	// Object storage
	S3 *Sink_S3 `protobuf:"bytes,10,opt,name=s3,proto3,oneof"`
}

func (*Sink_Yds_) isSink_Sink() {}

func (*Sink_S3_) isSink_Sink() {}

type Sink_Yds struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Fully qualified name of data stream
	StreamName    string `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sink_Yds) Reset() {
	*x = Sink_Yds{}
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sink_Yds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink_Yds) ProtoMessage() {}

func (x *Sink_Yds) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink_Yds.ProtoReflect.Descriptor instead.
func (*Sink_Yds) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_sink_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Sink_Yds) GetStreamName() string {
	if x != nil {
		return x.StreamName
	}
	return ""
}

type Sink_S3 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Object storage bucket
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Prefix to use for saved log object names
	Prefix        string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sink_S3) Reset() {
	*x = Sink_S3{}
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sink_S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink_S3) ProtoMessage() {}

func (x *Sink_S3) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sink_S3.ProtoReflect.Descriptor instead.
func (*Sink_S3) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_sink_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Sink_S3) GetBucket() string {
	if x != nil {
		return x.Bucket
	}
	return ""
}

func (x *Sink_S3) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

var File_yandex_cloud_logging_v1_sink_proto protoreflect.FileDescriptor

const file_yandex_cloud_logging_v1_sink_proto_rawDesc = "" +
	"\n" +
	"\"yandex/cloud/logging/v1/sink.proto\x12\x17yandex.cloud.logging.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dyandex/cloud/validation.proto\"\xfd\x04\n" +
	"\x04Sink\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x12\x19\n" +
	"\bcloud_id\x18\x03 \x01(\tR\acloudId\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12A\n" +
	"\x06labels\x18\a \x03(\v2).yandex.cloud.logging.v1.Sink.LabelsEntryR\x06labels\x12,\n" +
	"\x12service_account_id\x18\b \x01(\tR\x10serviceAccountId\x125\n" +
	"\x03yds\x18\t \x01(\v2!.yandex.cloud.logging.v1.Sink.YdsH\x00R\x03yds\x122\n" +
	"\x02s3\x18\n" +
	" \x01(\v2 .yandex.cloud.logging.v1.Sink.S3H\x00R\x02s3\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\x1a1\n" +
	"\x03Yds\x12*\n" +
	"\vstream_name\x18\x01 \x01(\tB\t\x8a\xc81\x05<=512R\n" +
	"streamName\x1ad\n" +
	"\x02S3\x12:\n" +
	"\x06bucket\x18\x01 \x01(\tB\"\xf2\xc71\x1e[a-zA-Z0-9][-a-zA-Z0-9.]{2,62}R\x06bucket\x12\"\n" +
	"\x06prefix\x18\x02 \x01(\tB\n" +
	"\x8a\xc81\x06<=1024R\x06prefixB\f\n" +
	"\x04sink\x12\x04\xc0\xc11\x01Bb\n" +
	"\x1byandex.cloud.api.logging.v1ZCgithub.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1;loggingb\x06proto3"

var (
	file_yandex_cloud_logging_v1_sink_proto_rawDescOnce sync.Once
	file_yandex_cloud_logging_v1_sink_proto_rawDescData []byte
)

func file_yandex_cloud_logging_v1_sink_proto_rawDescGZIP() []byte {
	file_yandex_cloud_logging_v1_sink_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_logging_v1_sink_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_logging_v1_sink_proto_rawDesc), len(file_yandex_cloud_logging_v1_sink_proto_rawDesc)))
	})
	return file_yandex_cloud_logging_v1_sink_proto_rawDescData
}

var file_yandex_cloud_logging_v1_sink_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_logging_v1_sink_proto_goTypes = []any{
	(*Sink)(nil),                  // 0: yandex.cloud.logging.v1.Sink
	nil,                           // 1: yandex.cloud.logging.v1.Sink.LabelsEntry
	(*Sink_Yds)(nil),              // 2: yandex.cloud.logging.v1.Sink.Yds
	(*Sink_S3)(nil),               // 3: yandex.cloud.logging.v1.Sink.S3
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_logging_v1_sink_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.logging.v1.Sink.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.logging.v1.Sink.labels:type_name -> yandex.cloud.logging.v1.Sink.LabelsEntry
	2, // 2: yandex.cloud.logging.v1.Sink.yds:type_name -> yandex.cloud.logging.v1.Sink.Yds
	3, // 3: yandex.cloud.logging.v1.Sink.s3:type_name -> yandex.cloud.logging.v1.Sink.S3
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_logging_v1_sink_proto_init() }
func file_yandex_cloud_logging_v1_sink_proto_init() {
	if File_yandex_cloud_logging_v1_sink_proto != nil {
		return
	}
	file_yandex_cloud_logging_v1_sink_proto_msgTypes[0].OneofWrappers = []any{
		(*Sink_Yds_)(nil),
		(*Sink_S3_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_logging_v1_sink_proto_rawDesc), len(file_yandex_cloud_logging_v1_sink_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_logging_v1_sink_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_logging_v1_sink_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_logging_v1_sink_proto_msgTypes,
	}.Build()
	File_yandex_cloud_logging_v1_sink_proto = out.File
	file_yandex_cloud_logging_v1_sink_proto_goTypes = nil
	file_yandex_cloud_logging_v1_sink_proto_depIdxs = nil
}
