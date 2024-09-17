// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Sink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

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
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Logs will be written to the sink on behalf of this service account
	ServiceAccountId string `protobuf:"bytes,8,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Logs destination
	//
	// Types that are assignable to Sink:
	//
	//	*Sink_Yds_
	//	*Sink_S3_
	Sink isSink_Sink `protobuf_oneof:"sink"`
}

func (x *Sink) Reset() {
	*x = Sink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink) ProtoMessage() {}

func (x *Sink) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *Sink) GetSink() isSink_Sink {
	if m != nil {
		return m.Sink
	}
	return nil
}

func (x *Sink) GetYds() *Sink_Yds {
	if x, ok := x.GetSink().(*Sink_Yds_); ok {
		return x.Yds
	}
	return nil
}

func (x *Sink) GetS3() *Sink_S3 {
	if x, ok := x.GetSink().(*Sink_S3_); ok {
		return x.S3
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Fully qualified name of data stream
	StreamName string `protobuf:"bytes,1,opt,name=stream_name,json=streamName,proto3" json:"stream_name,omitempty"`
}

func (x *Sink_Yds) Reset() {
	*x = Sink_Yds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink_Yds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink_Yds) ProtoMessage() {}

func (x *Sink_Yds) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Object storage bucket
	Bucket string `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	// Prefix to use for saved log object names
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *Sink_S3) Reset() {
	*x = Sink_S3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sink_S3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sink_S3) ProtoMessage() {}

func (x *Sink_S3) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_sink_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_yandex_cloud_logging_v1_sink_proto_rawDesc = []byte{
	0x0a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x6e, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfd, 0x04,
	0x0a, 0x04, 0x53, 0x69, 0x6e, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x41, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x2e, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x35, 0x0a, 0x03, 0x79, 0x64, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67,
	0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x6e, 0x6b, 0x2e, 0x59, 0x64, 0x73,
	0x48, 0x00, 0x52, 0x03, 0x79, 0x64, 0x73, 0x12, 0x32, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69,
	0x6e, 0x6b, 0x2e, 0x53, 0x33, 0x48, 0x00, 0x52, 0x02, 0x73, 0x33, 0x1a, 0x39, 0x0a, 0x0b, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x31, 0x0a, 0x03, 0x59, 0x64, 0x73, 0x12, 0x2a, 0x0a,
	0x0b, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x35, 0x31, 0x32, 0x52, 0x0a, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x64, 0x0a, 0x02, 0x53, 0x33, 0x12,
	0x3a, 0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x22, 0xf2, 0xc7, 0x31, 0x1e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5d,
	0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x2e, 0x5d, 0x7b, 0x32, 0x2c,
	0x36, 0x32, 0x7d, 0x52, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x70,
	0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31,
	0x06, 0x3c, 0x3d, 0x31, 0x30, 0x32, 0x34, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42,
	0x0c, 0x0a, 0x04, 0x73, 0x69, 0x6e, 0x6b, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x42, 0x62, 0x0a,
	0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x67, 0x67, 0x69, 0x6e,
	0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_logging_v1_sink_proto_rawDescOnce sync.Once
	file_yandex_cloud_logging_v1_sink_proto_rawDescData = file_yandex_cloud_logging_v1_sink_proto_rawDesc
)

func file_yandex_cloud_logging_v1_sink_proto_rawDescGZIP() []byte {
	file_yandex_cloud_logging_v1_sink_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_logging_v1_sink_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_logging_v1_sink_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_logging_v1_sink_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Sink); i {
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
		file_yandex_cloud_logging_v1_sink_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Sink_Yds); i {
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
		file_yandex_cloud_logging_v1_sink_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Sink_S3); i {
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
	file_yandex_cloud_logging_v1_sink_proto_msgTypes[0].OneofWrappers = []any{
		(*Sink_Yds_)(nil),
		(*Sink_S3_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_logging_v1_sink_proto_rawDesc,
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
	file_yandex_cloud_logging_v1_sink_proto_rawDesc = nil
	file_yandex_cloud_logging_v1_sink_proto_goTypes = nil
	file_yandex_cloud_logging_v1_sink_proto_depIdxs = nil
}
