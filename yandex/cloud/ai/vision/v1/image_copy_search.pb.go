// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/vision/v1/image_copy_search.proto

package vision

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

type ImageCopySearchAnnotation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Number of image copies
	CopyCount int64 `protobuf:"varint,1,opt,name=copy_count,json=copyCount,proto3" json:"copy_count,omitempty"`
	// Top relevance result of image copy search
	TopResults    []*CopyMatch `protobuf:"bytes,2,rep,name=top_results,json=topResults,proto3" json:"top_results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ImageCopySearchAnnotation) Reset() {
	*x = ImageCopySearchAnnotation{}
	mi := &file_yandex_cloud_ai_vision_v1_image_copy_search_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ImageCopySearchAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImageCopySearchAnnotation) ProtoMessage() {}

func (x *ImageCopySearchAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_image_copy_search_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImageCopySearchAnnotation.ProtoReflect.Descriptor instead.
func (*ImageCopySearchAnnotation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescGZIP(), []int{0}
}

func (x *ImageCopySearchAnnotation) GetCopyCount() int64 {
	if x != nil {
		return x.CopyCount
	}
	return 0
}

func (x *ImageCopySearchAnnotation) GetTopResults() []*CopyMatch {
	if x != nil {
		return x.TopResults
	}
	return nil
}

type CopyMatch struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// url of image
	ImageUrl string `protobuf:"bytes,1,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	// url of page that contains image
	PageUrl string `protobuf:"bytes,2,opt,name=page_url,json=pageUrl,proto3" json:"page_url,omitempty"`
	// page title that contains image
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// image description
	Description   string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CopyMatch) Reset() {
	*x = CopyMatch{}
	mi := &file_yandex_cloud_ai_vision_v1_image_copy_search_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CopyMatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CopyMatch) ProtoMessage() {}

func (x *CopyMatch) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_image_copy_search_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CopyMatch.ProtoReflect.Descriptor instead.
func (*CopyMatch) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescGZIP(), []int{1}
}

func (x *CopyMatch) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *CopyMatch) GetPageUrl() string {
	if x != nil {
		return x.PageUrl
	}
	return ""
}

func (x *CopyMatch) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CopyMatch) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_yandex_cloud_ai_vision_v1_image_copy_search_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDesc = "" +
	"\n" +
	"1yandex/cloud/ai/vision/v1/image_copy_search.proto\x12\x19yandex.cloud.ai.vision.v1\"\x81\x01\n" +
	"\x19ImageCopySearchAnnotation\x12\x1d\n" +
	"\n" +
	"copy_count\x18\x01 \x01(\x03R\tcopyCount\x12E\n" +
	"\vtop_results\x18\x02 \x03(\v2$.yandex.cloud.ai.vision.v1.CopyMatchR\n" +
	"topResults\"{\n" +
	"\tCopyMatch\x12\x1b\n" +
	"\timage_url\x18\x01 \x01(\tR\bimageUrl\x12\x19\n" +
	"\bpage_url\x18\x02 \x01(\tR\apageUrl\x12\x14\n" +
	"\x05title\x18\x03 \x01(\tR\x05title\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescriptionBe\n" +
	"\x1dyandex.cloud.api.ai.vision.v1ZDgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/vision/v1;visionb\x06proto3"

var (
	file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescData []byte
)

func file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDesc), len(file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDescData
}

var file_yandex_cloud_ai_vision_v1_image_copy_search_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_vision_v1_image_copy_search_proto_goTypes = []any{
	(*ImageCopySearchAnnotation)(nil), // 0: yandex.cloud.ai.vision.v1.ImageCopySearchAnnotation
	(*CopyMatch)(nil),                 // 1: yandex.cloud.ai.vision.v1.CopyMatch
}
var file_yandex_cloud_ai_vision_v1_image_copy_search_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.ai.vision.v1.ImageCopySearchAnnotation.top_results:type_name -> yandex.cloud.ai.vision.v1.CopyMatch
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_vision_v1_image_copy_search_proto_init() }
func file_yandex_cloud_ai_vision_v1_image_copy_search_proto_init() {
	if File_yandex_cloud_ai_vision_v1_image_copy_search_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDesc), len(file_yandex_cloud_ai_vision_v1_image_copy_search_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_vision_v1_image_copy_search_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_vision_v1_image_copy_search_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_vision_v1_image_copy_search_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_vision_v1_image_copy_search_proto = out.File
	file_yandex_cloud_ai_vision_v1_image_copy_search_proto_goTypes = nil
	file_yandex_cloud_ai_vision_v1_image_copy_search_proto_depIdxs = nil
}
