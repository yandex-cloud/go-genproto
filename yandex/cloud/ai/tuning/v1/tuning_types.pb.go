// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/tuning/v1/tuning_types.proto

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

type TuningTypeLora struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rank           int64   `protobuf:"varint,1,opt,name=rank,proto3" json:"rank,omitempty"`
	Alpha          float64 `protobuf:"fixed64,2,opt,name=alpha,proto3" json:"alpha,omitempty"`
	Initialization string  `protobuf:"bytes,3,opt,name=initialization,proto3" json:"initialization,omitempty"`
	Type           string  `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *TuningTypeLora) Reset() {
	*x = TuningTypeLora{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningTypeLora) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningTypeLora) ProtoMessage() {}

func (x *TuningTypeLora) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningTypeLora.ProtoReflect.Descriptor instead.
func (*TuningTypeLora) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescGZIP(), []int{0}
}

func (x *TuningTypeLora) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *TuningTypeLora) GetAlpha() float64 {
	if x != nil {
		return x.Alpha
	}
	return 0
}

func (x *TuningTypeLora) GetInitialization() string {
	if x != nil {
		return x.Initialization
	}
	return ""
}

func (x *TuningTypeLora) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type TuningTypePromptTune struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VirtualTokens int64 `protobuf:"varint,1,opt,name=virtual_tokens,json=virtualTokens,proto3" json:"virtual_tokens,omitempty"`
}

func (x *TuningTypePromptTune) Reset() {
	*x = TuningTypePromptTune{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningTypePromptTune) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningTypePromptTune) ProtoMessage() {}

func (x *TuningTypePromptTune) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningTypePromptTune.ProtoReflect.Descriptor instead.
func (*TuningTypePromptTune) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescGZIP(), []int{1}
}

func (x *TuningTypePromptTune) GetVirtualTokens() int64 {
	if x != nil {
		return x.VirtualTokens
	}
	return 0
}

var File_yandex_cloud_ai_tuning_v1_tuning_types_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e,
	0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x22, 0x76, 0x0a, 0x0e, 0x54, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x4c, 0x6f, 0x72, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x61, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x22, 0x3d, 0x0a, 0x14, 0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x75, 0x6e, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x76, 0x69, 0x72,
	0x74, 0x75, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x42, 0x63, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x3b, 0x66, 0x6f, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescData = file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDesc
)

func file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescData)
	})
	return file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDescData
}

var file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_tuning_v1_tuning_types_proto_goTypes = []any{
	(*TuningTypeLora)(nil),       // 0: yandex.cloud.ai.tuning.v1.TuningTypeLora
	(*TuningTypePromptTune)(nil), // 1: yandex.cloud.ai.tuning.v1.TuningTypePromptTune
}
var file_yandex_cloud_ai_tuning_v1_tuning_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_tuning_v1_tuning_types_proto_init() }
func file_yandex_cloud_ai_tuning_v1_tuning_types_proto_init() {
	if File_yandex_cloud_ai_tuning_v1_tuning_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TuningTypeLora); i {
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
		file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TuningTypePromptTune); i {
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
			RawDescriptor: file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_tuning_v1_tuning_types_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_tuning_v1_tuning_types_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_tuning_v1_tuning_types_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_tuning_v1_tuning_types_proto = out.File
	file_yandex_cloud_ai_tuning_v1_tuning_types_proto_rawDesc = nil
	file_yandex_cloud_ai_tuning_v1_tuning_types_proto_goTypes = nil
	file_yandex_cloud_ai_tuning_v1_tuning_types_proto_depIdxs = nil
}