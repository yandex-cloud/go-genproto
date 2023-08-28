// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/ai/llm/v1alpha/llm.proto

package llm

import (
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

// Sets the generation options.
type GenerationOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enables streaming of the partially generated text.
	PartialResults bool `protobuf:"varint,1,opt,name=partial_results,json=partialResults,proto3" json:"partial_results,omitempty"`
	// Affects creativity and randomness of the responses. It is a double number between 0 and infinity. A low temperature causes the responses to be straightforward, a high temperature causes high-level creativity and randomness.
	Temperature *wrapperspb.DoubleValue `protobuf:"bytes,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	// Sets response limit in tokens. It is a int number between 1 and 2000.
	MaxTokens *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
}

func (x *GenerationOptions) Reset() {
	*x = GenerationOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerationOptions) ProtoMessage() {}

func (x *GenerationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerationOptions.ProtoReflect.Descriptor instead.
func (*GenerationOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescGZIP(), []int{0}
}

func (x *GenerationOptions) GetPartialResults() bool {
	if x != nil {
		return x.PartialResults
	}
	return false
}

func (x *GenerationOptions) GetTemperature() *wrapperspb.DoubleValue {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *GenerationOptions) GetMaxTokens() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxTokens
	}
	return nil
}

// Contains the response, its score and length in tokens.
type Alternative struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text of the response.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Text log likelihood.
	Score float64 `protobuf:"fixed64,2,opt,name=score,proto3" json:"score,omitempty"`
	// Number of tokens in the response.
	NumTokens int64 `protobuf:"varint,3,opt,name=num_tokens,json=numTokens,proto3" json:"num_tokens,omitempty"`
}

func (x *Alternative) Reset() {
	*x = Alternative{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Alternative) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Alternative) ProtoMessage() {}

func (x *Alternative) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Alternative.ProtoReflect.Descriptor instead.
func (*Alternative) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescGZIP(), []int{1}
}

func (x *Alternative) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Alternative) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

func (x *Alternative) GetNumTokens() int64 {
	if x != nil {
		return x.NumTokens
	}
	return 0
}

// Contains description of the message for Chat.
type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message sender.
	Role string `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	// Text of the message.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Contains description of the LLM token (base coding unit).
type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Internal token id.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Token text representation.
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	// Defines if token is special or not. Special tokens define the behaviour of the model and are not visible for users.
	Special bool `protobuf:"varint,3,opt,name=special,proto3" json:"special,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescGZIP(), []int{3}
}

func (x *Token) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Token) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Token) GetSpecial() bool {
	if x != nil {
		return x.Special
	}
	return false
}

var File_yandex_cloud_ai_llm_v1alpha_llm_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x6c, 0x6c,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x70, 0x61, 0x72, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22,
	0x56, 0x0a, 0x0b, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75,
	0x6d, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0x31, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x45, 0x0a, 0x05, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x70, 0x65, 0x63, 0x69, 0x61,
	0x6c, 0x42, 0x66, 0x0a, 0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x6c, 0x6c, 0x6d, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x6c, 0x6c, 0x6d, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x6c, 0x6c, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescData = file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDesc
)

func file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescData)
	})
	return file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDescData
}

var file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_ai_llm_v1alpha_llm_proto_goTypes = []interface{}{
	(*GenerationOptions)(nil),      // 0: yandex.cloud.ai.llm.v1alpha.GenerationOptions
	(*Alternative)(nil),            // 1: yandex.cloud.ai.llm.v1alpha.Alternative
	(*Message)(nil),                // 2: yandex.cloud.ai.llm.v1alpha.Message
	(*Token)(nil),                  // 3: yandex.cloud.ai.llm.v1alpha.Token
	(*wrapperspb.DoubleValue)(nil), // 4: google.protobuf.DoubleValue
	(*wrapperspb.Int64Value)(nil),  // 5: google.protobuf.Int64Value
}
var file_yandex_cloud_ai_llm_v1alpha_llm_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.ai.llm.v1alpha.GenerationOptions.temperature:type_name -> google.protobuf.DoubleValue
	5, // 1: yandex.cloud.ai.llm.v1alpha.GenerationOptions.max_tokens:type_name -> google.protobuf.Int64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_llm_v1alpha_llm_proto_init() }
func file_yandex_cloud_ai_llm_v1alpha_llm_proto_init() {
	if File_yandex_cloud_ai_llm_v1alpha_llm_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerationOptions); i {
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
		file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Alternative); i {
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
		file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
			RawDescriptor: file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_llm_v1alpha_llm_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_llm_v1alpha_llm_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_llm_v1alpha_llm_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_llm_v1alpha_llm_proto = out.File
	file_yandex_cloud_ai_llm_v1alpha_llm_proto_rawDesc = nil
	file_yandex_cloud_ai_llm_v1alpha_llm_proto_goTypes = nil
	file_yandex_cloud_ai_llm_v1alpha_llm_proto_depIdxs = nil
}
