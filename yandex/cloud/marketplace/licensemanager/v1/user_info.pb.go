// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/marketplace/licensemanager/v1/user_info.proto

package licensemanager

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

type UserInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID, reserved field
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of legal person
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// INN of legal person
	Inn string `protobuf:"bytes,3,opt,name=inn,proto3" json:"inn,omitempty"`
	// KPP of legal person
	Kpp           string `protobuf:"bytes,4,opt,name=kpp,proto3" json:"kpp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	mi := &file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetInn() string {
	if x != nil {
		return x.Inn
	}
	return ""
}

func (x *UserInfo) GetKpp() string {
	if x != nil {
		return x.Kpp
	}
	return ""
}

var File_yandex_cloud_marketplace_licensemanager_v1_user_info_proto protoreflect.FileDescriptor

const file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDesc = "" +
	"\n" +
	":yandex/cloud/marketplace/licensemanager/v1/user_info.proto\x12*yandex.cloud.marketplace.licensemanager.v1\"R\n" +
	"\bUserInfo\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x10\n" +
	"\x03inn\x18\x03 \x01(\tR\x03inn\x12\x10\n" +
	"\x03kpp\x18\x04 \x01(\tR\x03kppB\x8f\x01\n" +
	".yandex.cloud.api.marketplace.licensemanager.v1Z]github.com/yandex-cloud/go-genproto/yandex/cloud/marketplace/licensemanager/v1;licensemanagerb\x06proto3"

var (
	file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescOnce sync.Once
	file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescData []byte
)

func file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescGZIP() []byte {
	file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDesc), len(file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDesc)))
	})
	return file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDescData
}

var file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_goTypes = []any{
	(*UserInfo)(nil), // 0: yandex.cloud.marketplace.licensemanager.v1.UserInfo
}
var file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_init() }
func file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_init() {
	if File_yandex_cloud_marketplace_licensemanager_v1_user_info_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDesc), len(file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_msgTypes,
	}.Build()
	File_yandex_cloud_marketplace_licensemanager_v1_user_info_proto = out.File
	file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_goTypes = nil
	file_yandex_cloud_marketplace_licensemanager_v1_user_info_proto_depIdxs = nil
}
