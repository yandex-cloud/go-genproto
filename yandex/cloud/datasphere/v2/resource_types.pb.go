// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datasphere/v2/resource_types.proto

package datasphere

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

type ResourceType int32

const (
	ResourceType_RESOURCE_TYPE_UNSPECIFIED  ResourceType = 0
	ResourceType_RESOURCE_TYPE_SECRET       ResourceType = 2
	ResourceType_RESOURCE_TYPE_DOCKER_IMAGE ResourceType = 3
	ResourceType_RESOURCE_TYPE_DATASET      ResourceType = 4
	ResourceType_RESOURCE_TYPE_S3           ResourceType = 5
	ResourceType_RESOURCE_TYPE_NODE         ResourceType = 6
	ResourceType_RESOURCE_TYPE_ALIAS        ResourceType = 8
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RESOURCE_TYPE_UNSPECIFIED",
		2: "RESOURCE_TYPE_SECRET",
		3: "RESOURCE_TYPE_DOCKER_IMAGE",
		4: "RESOURCE_TYPE_DATASET",
		5: "RESOURCE_TYPE_S3",
		6: "RESOURCE_TYPE_NODE",
		8: "RESOURCE_TYPE_ALIAS",
	}
	ResourceType_value = map[string]int32{
		"RESOURCE_TYPE_UNSPECIFIED":  0,
		"RESOURCE_TYPE_SECRET":       2,
		"RESOURCE_TYPE_DOCKER_IMAGE": 3,
		"RESOURCE_TYPE_DATASET":      4,
		"RESOURCE_TYPE_S3":           5,
		"RESOURCE_TYPE_NODE":         6,
		"RESOURCE_TYPE_ALIAS":        8,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datasphere_v2_resource_types_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datasphere_v2_resource_types_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescGZIP(), []int{0}
}

var File_yandex_cloud_datasphere_v2_resource_types_proto protoreflect.FileDescriptor

const file_yandex_cloud_datasphere_v2_resource_types_proto_rawDesc = "" +
	"\n" +
	"/yandex/cloud/datasphere/v2/resource_types.proto\x12\x1ayandex.cloud.datasphere.v2*\xd5\x01\n" +
	"\fResourceType\x12\x1d\n" +
	"\x19RESOURCE_TYPE_UNSPECIFIED\x10\x00\x12\x18\n" +
	"\x14RESOURCE_TYPE_SECRET\x10\x02\x12\x1e\n" +
	"\x1aRESOURCE_TYPE_DOCKER_IMAGE\x10\x03\x12\x19\n" +
	"\x15RESOURCE_TYPE_DATASET\x10\x04\x12\x14\n" +
	"\x10RESOURCE_TYPE_S3\x10\x05\x12\x16\n" +
	"\x12RESOURCE_TYPE_NODE\x10\x06\x12\x17\n" +
	"\x13RESOURCE_TYPE_ALIAS\x10\b\"\x04\b\x01\x10\x01\"\x04\b\a\x10\aB|\n" +
	"\x1eyandex.cloud.api.datasphere.v2B\x0fDsResourceTypesZIgithub.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2;datasphereb\x06proto3"

var (
	file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescData []byte
)

func file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datasphere_v2_resource_types_proto_rawDesc), len(file_yandex_cloud_datasphere_v2_resource_types_proto_rawDesc)))
	})
	return file_yandex_cloud_datasphere_v2_resource_types_proto_rawDescData
}

var file_yandex_cloud_datasphere_v2_resource_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_datasphere_v2_resource_types_proto_goTypes = []any{
	(ResourceType)(0), // 0: yandex.cloud.datasphere.v2.ResourceType
}
var file_yandex_cloud_datasphere_v2_resource_types_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v2_resource_types_proto_init() }
func file_yandex_cloud_datasphere_v2_resource_types_proto_init() {
	if File_yandex_cloud_datasphere_v2_resource_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datasphere_v2_resource_types_proto_rawDesc), len(file_yandex_cloud_datasphere_v2_resource_types_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datasphere_v2_resource_types_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v2_resource_types_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_datasphere_v2_resource_types_proto_enumTypes,
	}.Build()
	File_yandex_cloud_datasphere_v2_resource_types_proto = out.File
	file_yandex_cloud_datasphere_v2_resource_types_proto_goTypes = nil
	file_yandex_cloud_datasphere_v2_resource_types_proto_depIdxs = nil
}
