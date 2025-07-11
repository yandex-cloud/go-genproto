// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/iam/v1/resource.proto

package iam

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// A Resource. For more information, see [Resource](/docs/iam/concepts/access-control/resources-with-access-control).
type Resource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The type of the resource, e.g. resource-manager.folder, billing.account, compute.snapshot, etc.
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Resource) Reset() {
	*x = Resource{}
	mi := &file_yandex_cloud_iam_v1_resource_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_resource_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_resource_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Resource) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_yandex_cloud_iam_v1_resource_proto protoreflect.FileDescriptor

const file_yandex_cloud_iam_v1_resource_proto_rawDesc = "" +
	"\n" +
	"\"yandex/cloud/iam/v1/resource.proto\x12\x13yandex.cloud.iam.v1\x1a\x1dyandex/cloud/validation.proto\"J\n" +
	"\bResource\x12\x1c\n" +
	"\x02id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\x02id\x12 \n" +
	"\x04type\x18\x02 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=64R\x04typeBV\n" +
	"\x17yandex.cloud.api.iam.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1;iamb\x06proto3"

var (
	file_yandex_cloud_iam_v1_resource_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_resource_proto_rawDescData []byte
)

func file_yandex_cloud_iam_v1_resource_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_resource_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_resource_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_iam_v1_resource_proto_rawDesc), len(file_yandex_cloud_iam_v1_resource_proto_rawDesc)))
	})
	return file_yandex_cloud_iam_v1_resource_proto_rawDescData
}

var file_yandex_cloud_iam_v1_resource_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_iam_v1_resource_proto_goTypes = []any{
	(*Resource)(nil), // 0: yandex.cloud.iam.v1.Resource
}
var file_yandex_cloud_iam_v1_resource_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_resource_proto_init() }
func file_yandex_cloud_iam_v1_resource_proto_init() {
	if File_yandex_cloud_iam_v1_resource_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_iam_v1_resource_proto_rawDesc), len(file_yandex_cloud_iam_v1_resource_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iam_v1_resource_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_resource_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_iam_v1_resource_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_resource_proto = out.File
	file_yandex_cloud_iam_v1_resource_proto_goTypes = nil
	file_yandex_cloud_iam_v1_resource_proto_depIdxs = nil
}
