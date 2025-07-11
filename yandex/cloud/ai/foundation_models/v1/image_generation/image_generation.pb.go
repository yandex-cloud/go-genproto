// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/foundation_models/v1/image_generation/image_generation.proto

package image_generation

import (
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

// The text descriptions and weights that the model uses to generate an image.
type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Text describing the image.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Message weight. Negative values indicate negative messages. Note: Currently not supported.
	Weight        float64 `protobuf:"fixed64,2,opt,name=weight,proto3" json:"weight,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes[0]
	if x != nil {
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
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Message) GetWeight() float64 {
	if x != nil {
		return x.Weight
	}
	return 0
}

type AspectRatio struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Weight of width in image.
	WidthRatio int64 `protobuf:"varint,1,opt,name=width_ratio,json=widthRatio,proto3" json:"width_ratio,omitempty"`
	// Weight of height in image.
	HeightRatio   int64 `protobuf:"varint,2,opt,name=height_ratio,json=heightRatio,proto3" json:"height_ratio,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AspectRatio) Reset() {
	*x = AspectRatio{}
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AspectRatio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AspectRatio) ProtoMessage() {}

func (x *AspectRatio) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AspectRatio.ProtoReflect.Descriptor instead.
func (*AspectRatio) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescGZIP(), []int{1}
}

func (x *AspectRatio) GetWidthRatio() int64 {
	if x != nil {
		return x.WidthRatio
	}
	return 0
}

func (x *AspectRatio) GetHeightRatio() int64 {
	if x != nil {
		return x.HeightRatio
	}
	return 0
}

type ImageGenerationOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The [MIME type](https://en.wikipedia.org/wiki/Media_type) of generated image format.
	// For possible specifications, see [documentation](/docs/foundation-models/concepts).
	MimeType string `protobuf:"bytes,1,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Seed for image generation. It serves as a starting point for image generation from noise. If set to 0 or not provided, a randomly generated value will be used.
	Seed int64 `protobuf:"varint,2,opt,name=seed,proto3" json:"seed,omitempty"`
	// Aspect ratio of generated image.
	AspectRatio   *AspectRatio `protobuf:"bytes,3,opt,name=aspect_ratio,json=aspectRatio,proto3" json:"aspect_ratio,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageGenerationOptions) Reset() {
	*x = ImageGenerationOptions{}
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageGenerationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageGenerationOptions) ProtoMessage() {}

func (x *ImageGenerationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageGenerationOptions.ProtoReflect.Descriptor instead.
func (*ImageGenerationOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescGZIP(), []int{2}
}

func (x *ImageGenerationOptions) GetMimeType() string {
	if x != nil {
		return x.MimeType
	}
	return ""
}

func (x *ImageGenerationOptions) GetSeed() int64 {
	if x != nil {
		return x.Seed
	}
	return 0
}

func (x *ImageGenerationOptions) GetAspectRatio() *AspectRatio {
	if x != nil {
		return x.AspectRatio
	}
	return nil
}

var File_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDesc = "" +
	"\n" +
	"Lyandex/cloud/ai/foundation_models/v1/image_generation/image_generation.proto\x125yandex.cloud.ai.foundation_models.v1.image_generation\"5\n" +
	"\aMessage\x12\x12\n" +
	"\x04text\x18\x01 \x01(\tR\x04text\x12\x16\n" +
	"\x06weight\x18\x02 \x01(\x01R\x06weight\"Q\n" +
	"\vAspectRatio\x12\x1f\n" +
	"\vwidth_ratio\x18\x01 \x01(\x03R\n" +
	"widthRatio\x12!\n" +
	"\fheight_ratio\x18\x02 \x01(\x03R\vheightRatio\"\xb0\x01\n" +
	"\x16ImageGenerationOptions\x12\x1b\n" +
	"\tmime_type\x18\x01 \x01(\tR\bmimeType\x12\x12\n" +
	"\x04seed\x18\x02 \x01(\x03R\x04seed\x12e\n" +
	"\faspect_ratio\x18\x03 \x01(\v2B.yandex.cloud.ai.foundation_models.v1.image_generation.AspectRatioR\vaspectRatioB\xa7\x01\n" +
	"9yandex.cloud.api.ai.foundation_models.v1.image_generationZjgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/foundation_models/v1/image_generation;image_generationb\x06proto3"

var (
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescData []byte
)

func file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDesc), len(file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDescData
}

var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_goTypes = []any{
	(*Message)(nil),                // 0: yandex.cloud.ai.foundation_models.v1.image_generation.Message
	(*AspectRatio)(nil),            // 1: yandex.cloud.ai.foundation_models.v1.image_generation.AspectRatio
	(*ImageGenerationOptions)(nil), // 2: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationOptions
}
var file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.ai.foundation_models.v1.image_generation.ImageGenerationOptions.aspect_ratio:type_name -> yandex.cloud.ai.foundation_models.v1.image_generation.AspectRatio
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_init() }
func file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_init() {
	if File_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDesc), len(file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto = out.File
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_goTypes = nil
	file_yandex_cloud_ai_foundation_models_v1_image_generation_image_generation_proto_depIdxs = nil
}
