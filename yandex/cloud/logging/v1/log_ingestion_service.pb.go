// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/logging/v1/log_ingestion_service.proto

package logging

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	status "google.golang.org/genproto/googleapis/rpc/status"
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

type WriteRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Log entries destination.
	//
	// See [Destination] for details.
	Destination *Destination `protobuf:"bytes,1,opt,name=destination,proto3" json:"destination,omitempty"`
	// Common resource (type, ID) specification for log entries.
	Resource *LogEntryResource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// List of log entries.
	Entries []*IncomingLogEntry `protobuf:"bytes,3,rep,name=entries,proto3" json:"entries,omitempty"`
	// Log entries defaults.
	//
	// See [LogEntryDefaults] for details.
	Defaults      *LogEntryDefaults `protobuf:"bytes,4,opt,name=defaults,proto3" json:"defaults,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRequest) GetDestination() *Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *WriteRequest) GetResource() *LogEntryResource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *WriteRequest) GetEntries() []*IncomingLogEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

func (x *WriteRequest) GetDefaults() *LogEntryDefaults {
	if x != nil {
		return x.Defaults
	}
	return nil
}

type WriteResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Map<idx, status> of ingest failures.
	//
	// If entry with idx N is absent, it was ingested successfully.
	Errors        map[int64]*status.Status `protobuf:"bytes,1,rep,name=errors,proto3" json:"errors,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescGZIP(), []int{1}
}

func (x *WriteResponse) GetErrors() map[int64]*status.Status {
	if x != nil {
		return x.Errors
	}
	return nil
}

var File_yandex_cloud_logging_v1_log_ingestion_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc = "" +
	"\n" +
	"3yandex/cloud/logging/v1/log_ingestion_service.proto\x12\x17yandex.cloud.logging.v1\x1a\x17google/rpc/status.proto\x1a'yandex/cloud/logging/v1/log_entry.proto\x1a*yandex/cloud/logging/v1/log_resource.proto\x1a\x1dyandex/cloud/validation.proto\"\xba\x02\n" +
	"\fWriteRequest\x12L\n" +
	"\vdestination\x18\x01 \x01(\v2$.yandex.cloud.logging.v1.DestinationB\x04\xe8\xc71\x01R\vdestination\x12E\n" +
	"\bresource\x18\x02 \x01(\v2).yandex.cloud.logging.v1.LogEntryResourceR\bresource\x12N\n" +
	"\aentries\x18\x03 \x03(\v2).yandex.cloud.logging.v1.IncomingLogEntryB\t\x82\xc81\x051-100R\aentries\x12E\n" +
	"\bdefaults\x18\x04 \x01(\v2).yandex.cloud.logging.v1.LogEntryDefaultsR\bdefaults\"\xaa\x01\n" +
	"\rWriteResponse\x12J\n" +
	"\x06errors\x18\x01 \x03(\v22.yandex.cloud.logging.v1.WriteResponse.ErrorsEntryR\x06errors\x1aM\n" +
	"\vErrorsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\x03R\x03key\x12(\n" +
	"\x05value\x18\x02 \x01(\v2\x12.google.rpc.StatusR\x05value:\x028\x012m\n" +
	"\x13LogIngestionService\x12V\n" +
	"\x05Write\x12%.yandex.cloud.logging.v1.WriteRequest\x1a&.yandex.cloud.logging.v1.WriteResponseBb\n" +
	"\x1byandex.cloud.api.logging.v1ZCgithub.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1;loggingb\x06proto3"

var (
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData []byte
)

func file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc), len(file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc)))
	})
	return file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDescData
}

var file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_logging_v1_log_ingestion_service_proto_goTypes = []any{
	(*WriteRequest)(nil),     // 0: yandex.cloud.logging.v1.WriteRequest
	(*WriteResponse)(nil),    // 1: yandex.cloud.logging.v1.WriteResponse
	nil,                      // 2: yandex.cloud.logging.v1.WriteResponse.ErrorsEntry
	(*Destination)(nil),      // 3: yandex.cloud.logging.v1.Destination
	(*LogEntryResource)(nil), // 4: yandex.cloud.logging.v1.LogEntryResource
	(*IncomingLogEntry)(nil), // 5: yandex.cloud.logging.v1.IncomingLogEntry
	(*LogEntryDefaults)(nil), // 6: yandex.cloud.logging.v1.LogEntryDefaults
	(*status.Status)(nil),    // 7: google.rpc.Status
}
var file_yandex_cloud_logging_v1_log_ingestion_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.logging.v1.WriteRequest.destination:type_name -> yandex.cloud.logging.v1.Destination
	4, // 1: yandex.cloud.logging.v1.WriteRequest.resource:type_name -> yandex.cloud.logging.v1.LogEntryResource
	5, // 2: yandex.cloud.logging.v1.WriteRequest.entries:type_name -> yandex.cloud.logging.v1.IncomingLogEntry
	6, // 3: yandex.cloud.logging.v1.WriteRequest.defaults:type_name -> yandex.cloud.logging.v1.LogEntryDefaults
	2, // 4: yandex.cloud.logging.v1.WriteResponse.errors:type_name -> yandex.cloud.logging.v1.WriteResponse.ErrorsEntry
	7, // 5: yandex.cloud.logging.v1.WriteResponse.ErrorsEntry.value:type_name -> google.rpc.Status
	0, // 6: yandex.cloud.logging.v1.LogIngestionService.Write:input_type -> yandex.cloud.logging.v1.WriteRequest
	1, // 7: yandex.cloud.logging.v1.LogIngestionService.Write:output_type -> yandex.cloud.logging.v1.WriteResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_logging_v1_log_ingestion_service_proto_init() }
func file_yandex_cloud_logging_v1_log_ingestion_service_proto_init() {
	if File_yandex_cloud_logging_v1_log_ingestion_service_proto != nil {
		return
	}
	file_yandex_cloud_logging_v1_log_entry_proto_init()
	file_yandex_cloud_logging_v1_log_resource_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc), len(file_yandex_cloud_logging_v1_log_ingestion_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_logging_v1_log_ingestion_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_logging_v1_log_ingestion_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_logging_v1_log_ingestion_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_logging_v1_log_ingestion_service_proto = out.File
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_goTypes = nil
	file_yandex_cloud_logging_v1_log_ingestion_service_proto_depIdxs = nil
}
