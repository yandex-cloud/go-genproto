// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/k8s/v1/version_service.proto

package k8s

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

type ListVersionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListVersionsRequest) Reset() {
	*x = ListVersionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsRequest) ProtoMessage() {}

func (x *ListVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsRequest.ProtoReflect.Descriptor instead.
func (*ListVersionsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_version_service_proto_rawDescGZIP(), []int{0}
}

type ListVersionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Versions available in the specified release channel.
	AvailableVersions []*AvailableVersions `protobuf:"bytes,1,rep,name=available_versions,json=availableVersions,proto3" json:"available_versions,omitempty"`
}

func (x *ListVersionsResponse) Reset() {
	*x = ListVersionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListVersionsResponse) ProtoMessage() {}

func (x *ListVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListVersionsResponse.ProtoReflect.Descriptor instead.
func (*ListVersionsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_version_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListVersionsResponse) GetAvailableVersions() []*AvailableVersions {
	if x != nil {
		return x.AvailableVersions
	}
	return nil
}

type AvailableVersions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Release channel: `RAPID`, `REGULAR` or `STABLE`. For more details see [documentation](https://cloud.yandex.ru/docs/managed-kubernetes/concepts/release-channels-and-updates).
	ReleaseChannel ReleaseChannel `protobuf:"varint,1,opt,name=release_channel,json=releaseChannel,proto3,enum=yandex.cloud.k8s.v1.ReleaseChannel" json:"release_channel,omitempty"`
	// Version of Kubernetes components.
	Versions []string `protobuf:"bytes,2,rep,name=versions,proto3" json:"versions,omitempty"`
}

func (x *AvailableVersions) Reset() {
	*x = AvailableVersions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AvailableVersions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AvailableVersions) ProtoMessage() {}

func (x *AvailableVersions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AvailableVersions.ProtoReflect.Descriptor instead.
func (*AvailableVersions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_k8s_v1_version_service_proto_rawDescGZIP(), []int{2}
}

func (x *AvailableVersions) GetReleaseChannel() ReleaseChannel {
	if x != nil {
		return x.ReleaseChannel
	}
	return ReleaseChannel_RELEASE_CHANNEL_UNSPECIFIED
}

func (x *AvailableVersions) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

var File_yandex_cloud_k8s_v1_version_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_k8s_v1_version_service_proto_rawDesc = []byte{
	0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b,
	0x38, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x38, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6d, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x55, 0x0a, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x7d, 0x0a, 0x11, 0x41, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4c, 0x0a, 0x0f,
	0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x65,
	0x61, 0x73, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x32, 0x97, 0x01, 0x0a, 0x0e, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6b, 0x38, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x12,
	0x1f, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6b, 0x38, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6b, 0x38,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6b, 0x38, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_k8s_v1_version_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_k8s_v1_version_service_proto_rawDescData = file_yandex_cloud_k8s_v1_version_service_proto_rawDesc
)

func file_yandex_cloud_k8s_v1_version_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_k8s_v1_version_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_k8s_v1_version_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_k8s_v1_version_service_proto_rawDescData)
	})
	return file_yandex_cloud_k8s_v1_version_service_proto_rawDescData
}

var file_yandex_cloud_k8s_v1_version_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_k8s_v1_version_service_proto_goTypes = []interface{}{
	(*ListVersionsRequest)(nil),  // 0: yandex.cloud.k8s.v1.ListVersionsRequest
	(*ListVersionsResponse)(nil), // 1: yandex.cloud.k8s.v1.ListVersionsResponse
	(*AvailableVersions)(nil),    // 2: yandex.cloud.k8s.v1.AvailableVersions
	(ReleaseChannel)(0),          // 3: yandex.cloud.k8s.v1.ReleaseChannel
}
var file_yandex_cloud_k8s_v1_version_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.k8s.v1.ListVersionsResponse.available_versions:type_name -> yandex.cloud.k8s.v1.AvailableVersions
	3, // 1: yandex.cloud.k8s.v1.AvailableVersions.release_channel:type_name -> yandex.cloud.k8s.v1.ReleaseChannel
	0, // 2: yandex.cloud.k8s.v1.VersionService.List:input_type -> yandex.cloud.k8s.v1.ListVersionsRequest
	1, // 3: yandex.cloud.k8s.v1.VersionService.List:output_type -> yandex.cloud.k8s.v1.ListVersionsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_k8s_v1_version_service_proto_init() }
func file_yandex_cloud_k8s_v1_version_service_proto_init() {
	if File_yandex_cloud_k8s_v1_version_service_proto != nil {
		return
	}
	file_yandex_cloud_k8s_v1_cluster_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVersionsRequest); i {
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
		file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListVersionsResponse); i {
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
		file_yandex_cloud_k8s_v1_version_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AvailableVersions); i {
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
			RawDescriptor: file_yandex_cloud_k8s_v1_version_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_k8s_v1_version_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_k8s_v1_version_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_k8s_v1_version_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_k8s_v1_version_service_proto = out.File
	file_yandex_cloud_k8s_v1_version_service_proto_rawDesc = nil
	file_yandex_cloud_k8s_v1_version_service_proto_goTypes = nil
	file_yandex_cloud_k8s_v1_version_service_proto_depIdxs = nil
}
