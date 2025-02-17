// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/tuning/v1/tuning_error.proto

package fomo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TuningError_Type int32

const (
	TuningError_TYPE_UNSPECIFIED TuningError_Type = 0
	TuningError_PUBLIC           TuningError_Type = 1
	TuningError_INTERNAL         TuningError_Type = 2
)

// Enum value maps for TuningError_Type.
var (
	TuningError_Type_name = map[int32]string{
		0: "TYPE_UNSPECIFIED",
		1: "PUBLIC",
		2: "INTERNAL",
	}
	TuningError_Type_value = map[string]int32{
		"TYPE_UNSPECIFIED": 0,
		"PUBLIC":           1,
		"INTERNAL":         2,
	}
)

func (x TuningError_Type) Enum() *TuningError_Type {
	p := new(TuningError_Type)
	*p = x
	return p
}

func (x TuningError_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TuningError_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_tuning_v1_tuning_error_proto_enumTypes[0].Descriptor()
}

func (TuningError_Type) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_tuning_v1_tuning_error_proto_enumTypes[0]
}

func (x TuningError_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TuningError_Type.Descriptor instead.
func (TuningError_Type) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescGZIP(), []int{0, 0}
}

type TuningError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TuningTaskId string           `protobuf:"bytes,1,opt,name=tuning_task_id,json=tuningTaskId,proto3" json:"tuning_task_id,omitempty"`
	Message      string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type         TuningError_Type `protobuf:"varint,3,opt,name=type,proto3,enum=yandex.cloud.ai.tuning.v1.TuningError_Type" json:"type,omitempty"`
}

func (x *TuningError) Reset() {
	*x = TuningError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_tuning_v1_tuning_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningError) ProtoMessage() {}

func (x *TuningError) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_tuning_v1_tuning_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningError.ProtoReflect.Descriptor instead.
func (*TuningError) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescGZIP(), []int{0}
}

func (x *TuningError) GetTuningTaskId() string {
	if x != nil {
		return x.TuningTaskId
	}
	return ""
}

func (x *TuningError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *TuningError) GetType() TuningError_Type {
	if x != nil {
		return x.Type
	}
	return TuningError_TYPE_UNSPECIFIED
}

var File_yandex_cloud_ai_tuning_v1_tuning_error_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e,
	0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0xc6, 0x01, 0x0a, 0x0b, 0x54, 0x75,
	0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x04, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x14, 0x0a, 0x10, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x55, 0x42, 0x4c,
	0x49, 0x43, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x10, 0x02, 0x42, 0x63, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x31, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x3b, 0x66, 0x6f, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescData = file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDesc
)

func file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescData)
	})
	return file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDescData
}

var file_yandex_cloud_ai_tuning_v1_tuning_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_ai_tuning_v1_tuning_error_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_ai_tuning_v1_tuning_error_proto_goTypes = []any{
	(TuningError_Type)(0), // 0: yandex.cloud.ai.tuning.v1.TuningError.Type
	(*TuningError)(nil),   // 1: yandex.cloud.ai.tuning.v1.TuningError
}
var file_yandex_cloud_ai_tuning_v1_tuning_error_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.ai.tuning.v1.TuningError.type:type_name -> yandex.cloud.ai.tuning.v1.TuningError.Type
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_tuning_v1_tuning_error_proto_init() }
func file_yandex_cloud_ai_tuning_v1_tuning_error_proto_init() {
	if File_yandex_cloud_ai_tuning_v1_tuning_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_tuning_v1_tuning_error_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TuningError); i {
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
			RawDescriptor: file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_tuning_v1_tuning_error_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_tuning_v1_tuning_error_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_tuning_v1_tuning_error_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_tuning_v1_tuning_error_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_tuning_v1_tuning_error_proto = out.File
	file_yandex_cloud_ai_tuning_v1_tuning_error_proto_rawDesc = nil
	file_yandex_cloud_ai_tuning_v1_tuning_error_proto_goTypes = nil
	file_yandex_cloud_ai_tuning_v1_tuning_error_proto_depIdxs = nil
}
