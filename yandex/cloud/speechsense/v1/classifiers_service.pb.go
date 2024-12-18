// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/speechsense/v1/classifiers_service.proto

package speechsense

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ListClassifiersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Project id
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
}

func (x *ListClassifiersRequest) Reset() {
	*x = ListClassifiersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClassifiersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassifiersRequest) ProtoMessage() {}

func (x *ListClassifiersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassifiersRequest.ProtoReflect.Descriptor instead.
func (*ListClassifiersRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListClassifiersRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

type ListClassifiersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Classifiers belonging to the given project
	Classifiers []*Classifier `protobuf:"bytes,1,rep,name=classifiers,proto3" json:"classifiers,omitempty"`
}

func (x *ListClassifiersResponse) Reset() {
	*x = ListClassifiersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListClassifiersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListClassifiersResponse) ProtoMessage() {}

func (x *ListClassifiersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListClassifiersResponse.ProtoReflect.Descriptor instead.
func (*ListClassifiersResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListClassifiersResponse) GetClassifiers() []*Classifier {
	if x != nil {
		return x.Classifiers
	}
	return nil
}

var File_yandex_cloud_speechsense_v1_classifiers_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x37, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x22, 0x64, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73,
	0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x52, 0x0b, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x32,
	0xb5, 0x01, 0x0a, 0x12, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9e, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x25, 0x3a, 0x01, 0x2a, 0x22, 0x20, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65,
	0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x87, 0x01, 0x0a, 0x1f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x17, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescData = file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDesc
)

func file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescData)
	})
	return file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDescData
}

var file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_speechsense_v1_classifiers_service_proto_goTypes = []any{
	(*ListClassifiersRequest)(nil),  // 0: yandex.cloud.speechsense.v1.ListClassifiersRequest
	(*ListClassifiersResponse)(nil), // 1: yandex.cloud.speechsense.v1.ListClassifiersResponse
	(*Classifier)(nil),              // 2: yandex.cloud.speechsense.v1.Classifier
}
var file_yandex_cloud_speechsense_v1_classifiers_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.speechsense.v1.ListClassifiersResponse.classifiers:type_name -> yandex.cloud.speechsense.v1.Classifier
	0, // 1: yandex.cloud.speechsense.v1.ClassifiersService.List:input_type -> yandex.cloud.speechsense.v1.ListClassifiersRequest
	1, // 2: yandex.cloud.speechsense.v1.ClassifiersService.List:output_type -> yandex.cloud.speechsense.v1.ListClassifiersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_speechsense_v1_classifiers_service_proto_init() }
func file_yandex_cloud_speechsense_v1_classifiers_service_proto_init() {
	if File_yandex_cloud_speechsense_v1_classifiers_service_proto != nil {
		return
	}
	file_yandex_cloud_speechsense_v1_classifier_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListClassifiersRequest); i {
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
		file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListClassifiersResponse); i {
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
			RawDescriptor: file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_speechsense_v1_classifiers_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_speechsense_v1_classifiers_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_speechsense_v1_classifiers_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_speechsense_v1_classifiers_service_proto = out.File
	file_yandex_cloud_speechsense_v1_classifiers_service_proto_rawDesc = nil
	file_yandex_cloud_speechsense_v1_classifiers_service_proto_goTypes = nil
	file_yandex_cloud_speechsense_v1_classifiers_service_proto_depIdxs = nil
}
