// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: yandex/cloud/operation/operation.proto

package operation

import (
	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
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

// An Operation resource. For more information, see [Operation](/docs/api-design-guide/concepts/operation).
type Operation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the operation.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Description of the operation. 0-256 characters long.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"` // ex: Create VM, Stop VM, Delete Disk, Snapshot Disk, etc
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the user or service account who initiated the operation.
	CreatedBy string `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// The time when the Operation resource was last modified.
	ModifiedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	// If the value is `false`, it means the operation is still in progress.
	// If `true`, the operation is completed, and either `error` or `response` is available.
	Done bool `protobuf:"varint,6,opt,name=done,proto3" json:"done,omitempty"`
	// Service-specific metadata associated with the operation.
	// It typically contains the ID of the target resource that the operation is performed on.
	// Any method that returns a long-running operation should document the metadata type, if any.
	Metadata *anypb.Any `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// The operation result.
	// If `done == false` and there was no failure detected, neither `error` nor `response` is set.
	// If `done == false` and there was a failure detected, `error` is set.
	// If `done == true`, exactly one of `error` or `response` is set.
	//
	// Types that are assignable to Result:
	//	*Operation_Error
	//	*Operation_Response
	Result isOperation_Result `protobuf_oneof:"result"`
}

func (x *Operation) Reset() {
	*x = Operation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_operation_operation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Operation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Operation) ProtoMessage() {}

func (x *Operation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_operation_operation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Operation.ProtoReflect.Descriptor instead.
func (*Operation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_operation_operation_proto_rawDescGZIP(), []int{0}
}

func (x *Operation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Operation) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Operation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Operation) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Operation) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

func (x *Operation) GetDone() bool {
	if x != nil {
		return x.Done
	}
	return false
}

func (x *Operation) GetMetadata() *anypb.Any {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (m *Operation) GetResult() isOperation_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *Operation) GetError() *status.Status {
	if x, ok := x.GetResult().(*Operation_Error); ok {
		return x.Error
	}
	return nil
}

func (x *Operation) GetResponse() *anypb.Any {
	if x, ok := x.GetResult().(*Operation_Response); ok {
		return x.Response
	}
	return nil
}

type isOperation_Result interface {
	isOperation_Result()
}

type Operation_Error struct {
	// The error result of the operation in case of failure or cancellation.
	Error *status.Status `protobuf:"bytes,8,opt,name=error,proto3,oneof"`
}

type Operation_Response struct {
	// The normal response of the operation in case of success.
	// If the original method returns no data on success, such as Delete,
	// the response is [google.protobuf.Empty].
	// If the original method is the standard Create/Update,
	// the response should be the target resource of the operation.
	// Any method that returns a long-running operation should document the response type, if any.
	Response *anypb.Any `protobuf:"bytes,9,opt,name=response,proto3,oneof"`
}

func (*Operation_Error) isOperation_Result() {}

func (*Operation_Response) isOperation_Result() {}

var File_yandex_cloud_operation_operation_proto protoreflect.FileDescriptor

var file_yandex_cloud_operation_operation_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x3b, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x6f, 0x6e, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x64, 0x6f, 0x6e, 0x65,
	0x12, 0x30, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x32,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x48, 0x00, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x62, 0x0a, 0x1a,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_operation_operation_proto_rawDescOnce sync.Once
	file_yandex_cloud_operation_operation_proto_rawDescData = file_yandex_cloud_operation_operation_proto_rawDesc
)

func file_yandex_cloud_operation_operation_proto_rawDescGZIP() []byte {
	file_yandex_cloud_operation_operation_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_operation_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_operation_operation_proto_rawDescData)
	})
	return file_yandex_cloud_operation_operation_proto_rawDescData
}

var file_yandex_cloud_operation_operation_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_operation_operation_proto_goTypes = []interface{}{
	(*Operation)(nil),             // 0: yandex.cloud.operation.Operation
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 2: google.protobuf.Any
	(*status.Status)(nil),         // 3: google.rpc.Status
}
var file_yandex_cloud_operation_operation_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.operation.Operation.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.operation.Operation.modified_at:type_name -> google.protobuf.Timestamp
	2, // 2: yandex.cloud.operation.Operation.metadata:type_name -> google.protobuf.Any
	3, // 3: yandex.cloud.operation.Operation.error:type_name -> google.rpc.Status
	2, // 4: yandex.cloud.operation.Operation.response:type_name -> google.protobuf.Any
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_operation_operation_proto_init() }
func file_yandex_cloud_operation_operation_proto_init() {
	if File_yandex_cloud_operation_operation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_operation_operation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Operation); i {
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
	file_yandex_cloud_operation_operation_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Operation_Error)(nil),
		(*Operation_Response)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_operation_operation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_operation_operation_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_operation_operation_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_operation_operation_proto_msgTypes,
	}.Build()
	File_yandex_cloud_operation_operation_proto = out.File
	file_yandex_cloud_operation_operation_proto_rawDesc = nil
	file_yandex_cloud_operation_operation_proto_goTypes = nil
	file_yandex_cloud_operation_operation_proto_depIdxs = nil
}
