// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/greenplum/v1/user.proto

package greenplum

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// User name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// User password. Used only in create and update requests
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Resource group for user's queries
	ResourceGroup string `protobuf:"bytes,3,opt,name=resource_group,json=resourceGroup,proto3" json:"resource_group,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetResourceGroup() string {
	if x != nil {
		return x.ResourceGroup
	}
	return ""
}

var File_yandex_cloud_mdb_greenplum_v1_user_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65,
	0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xaf, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x41, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2d, 0xe8, 0xc7, 0x31, 0x01, 0xf2, 0xc7, 0x31, 0x1d, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d,
	0x5a, 0x5f, 0x5d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x7b,
	0x30, 0x2c, 0x36, 0x32, 0x7d, 0x24, 0x8a, 0xc8, 0x31, 0x04, 0x31, 0x2d, 0x36, 0x33, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x36, 0x2d, 0x32, 0x30,
	0x30, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x3d, 0x0a, 0x0e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x16, 0xf2, 0xc7, 0x31, 0x12, 0x5e, 0x5b, 0x5e, 0x5c, 0x7c, 0x2f, 0x2a,
	0x3f, 0x2e, 0x2c, 0x3b, 0x22, 0x27, 0x3c, 0x3e, 0x5d, 0x2b, 0x24, 0x52, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x42, 0x70, 0x0a, 0x21, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x5a,
	0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f,
	0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescData = file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDesc
)

func file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDescData
}

var file_yandex_cloud_mdb_greenplum_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_mdb_greenplum_v1_user_proto_goTypes = []any{
	(*User)(nil), // 0: yandex.cloud.mdb.greenplum.v1.User
}
var file_yandex_cloud_mdb_greenplum_v1_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_greenplum_v1_user_proto_init() }
func file_yandex_cloud_mdb_greenplum_v1_user_proto_init() {
	if File_yandex_cloud_mdb_greenplum_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_greenplum_v1_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
			RawDescriptor: file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_greenplum_v1_user_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_greenplum_v1_user_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_greenplum_v1_user_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_greenplum_v1_user_proto = out.File
	file_yandex_cloud_mdb_greenplum_v1_user_proto_rawDesc = nil
	file_yandex_cloud_mdb_greenplum_v1_user_proto_goTypes = nil
	file_yandex_cloud_mdb_greenplum_v1_user_proto_depIdxs = nil
}
