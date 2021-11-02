// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/ai/vision/v1/face_detection.proto

package vision

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

type FaceAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// An array of detected faces for the specified image.
	Faces []*Face `protobuf:"bytes,1,rep,name=faces,proto3" json:"faces,omitempty"`
}

func (x *FaceAnnotation) Reset() {
	*x = FaceAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FaceAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FaceAnnotation) ProtoMessage() {}

func (x *FaceAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FaceAnnotation.ProtoReflect.Descriptor instead.
func (*FaceAnnotation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescGZIP(), []int{0}
}

func (x *FaceAnnotation) GetFaces() []*Face {
	if x != nil {
		return x.Faces
	}
	return nil
}

type Face struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Area on the image where the face is located.
	BoundingBox *Polygon `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
}

func (x *Face) Reset() {
	*x = Face{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Face) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Face) ProtoMessage() {}

func (x *Face) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Face.ProtoReflect.Descriptor instead.
func (*Face) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescGZIP(), []int{1}
}

func (x *Face) GetBoundingBox() *Polygon {
	if x != nil {
		return x.BoundingBox
	}
	return nil
}

var File_yandex_cloud_ai_vision_v1_face_detection_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x63, 0x65,
	0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x2a, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x0e, 0x46, 0x61, 0x63, 0x65, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a, 0x05, 0x66, 0x61, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x63, 0x65, 0x52, 0x05, 0x66, 0x61, 0x63, 0x65, 0x73,
	0x22, 0x4d, 0x0a, 0x04, 0x46, 0x61, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x62, 0x6f, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x62, 0x6f, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6c, 0x79, 0x67,
	0x6f, 0x6e, 0x52, 0x0b, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x6f, 0x78, 0x42,
	0x65, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescData = file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDesc
)

func file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescData)
	})
	return file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDescData
}

var file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_vision_v1_face_detection_proto_goTypes = []interface{}{
	(*FaceAnnotation)(nil), // 0: yandex.cloud.ai.vision.v1.FaceAnnotation
	(*Face)(nil),           // 1: yandex.cloud.ai.vision.v1.Face
	(*Polygon)(nil),        // 2: yandex.cloud.ai.vision.v1.Polygon
}
var file_yandex_cloud_ai_vision_v1_face_detection_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.ai.vision.v1.FaceAnnotation.faces:type_name -> yandex.cloud.ai.vision.v1.Face
	2, // 1: yandex.cloud.ai.vision.v1.Face.bounding_box:type_name -> yandex.cloud.ai.vision.v1.Polygon
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_vision_v1_face_detection_proto_init() }
func file_yandex_cloud_ai_vision_v1_face_detection_proto_init() {
	if File_yandex_cloud_ai_vision_v1_face_detection_proto != nil {
		return
	}
	file_yandex_cloud_ai_vision_v1_primitives_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FaceAnnotation); i {
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
		file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Face); i {
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
			RawDescriptor: file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_vision_v1_face_detection_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_vision_v1_face_detection_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_vision_v1_face_detection_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_vision_v1_face_detection_proto = out.File
	file_yandex_cloud_ai_vision_v1_face_detection_proto_rawDesc = nil
	file_yandex_cloud_ai_vision_v1_face_detection_proto_goTypes = nil
	file_yandex_cloud_ai_vision_v1_face_detection_proto_depIdxs = nil
}
