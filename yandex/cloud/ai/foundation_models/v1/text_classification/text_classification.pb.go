// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/ai/foundation_models/v1/text_classification/text_classification.proto

package text_classification

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

// A pair of text label and corresponding confidence used in classification problems.
type ClassificationLabel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A label with a class name.
	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// Confidence of item's belonging to a class.
	Confidence float64 `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *ClassificationLabel) Reset() {
	*x = ClassificationLabel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationLabel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationLabel) ProtoMessage() {}

func (x *ClassificationLabel) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationLabel.ProtoReflect.Descriptor instead.
func (*ClassificationLabel) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescGZIP(), []int{0}
}

func (x *ClassificationLabel) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *ClassificationLabel) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

// Description of a sample for the classification task.
type ClassificationSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text sample.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Expected label for a given text.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *ClassificationSample) Reset() {
	*x = ClassificationSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationSample) ProtoMessage() {}

func (x *ClassificationSample) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationSample.ProtoReflect.Descriptor instead.
func (*ClassificationSample) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescGZIP(), []int{1}
}

func (x *ClassificationSample) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *ClassificationSample) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDesc = []byte{
	0x0a, 0x52, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4b,
	0x0a, 0x13, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x14, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x42, 0xb0, 0x01,
	0x0a, 0x3c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x78, 0x74, 0x5f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5a, 0x70,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x69, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3b, 0x74, 0x65, 0x78,
	0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescData = file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDesc
)

func file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescData)
	})
	return file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDescData
}

var file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_goTypes = []interface{}{
	(*ClassificationLabel)(nil),  // 0: yandex.cloud.ai.foundation_models.v1.text_classification.ClassificationLabel
	(*ClassificationSample)(nil), // 1: yandex.cloud.ai.foundation_models.v1.text_classification.ClassificationSample
}
var file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_init()
}
func file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_init() {
	if File_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationLabel); i {
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
		file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationSample); i {
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
			RawDescriptor: file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto = out.File
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_rawDesc = nil
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_goTypes = nil
	file_yandex_cloud_ai_foundation_models_v1_text_classification_text_classification_proto_depIdxs = nil
}
