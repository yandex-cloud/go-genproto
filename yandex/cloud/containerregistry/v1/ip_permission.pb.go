// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/containerregistry/v1/ip_permission.proto

package containerregistry

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

type IpPermissionAction int32

const (
	IpPermissionAction_IP_PERMISSION_ACTION_UNSPECIFIED IpPermissionAction = 0
	// Addition of an ip permission.
	IpPermissionAction_ADD IpPermissionAction = 1
	// Removal of an ip permission.
	IpPermissionAction_REMOVE IpPermissionAction = 2
)

// Enum value maps for IpPermissionAction.
var (
	IpPermissionAction_name = map[int32]string{
		0: "IP_PERMISSION_ACTION_UNSPECIFIED",
		1: "ADD",
		2: "REMOVE",
	}
	IpPermissionAction_value = map[string]int32{
		"IP_PERMISSION_ACTION_UNSPECIFIED": 0,
		"ADD":                              1,
		"REMOVE":                           2,
	}
)

func (x IpPermissionAction) Enum() *IpPermissionAction {
	p := new(IpPermissionAction)
	*p = x
	return p
}

func (x IpPermissionAction) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IpPermissionAction) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_enumTypes[0].Descriptor()
}

func (IpPermissionAction) Type() protoreflect.EnumType {
	return &file_yandex_cloud_containerregistry_v1_ip_permission_proto_enumTypes[0]
}

func (x IpPermissionAction) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IpPermissionAction.Descriptor instead.
func (IpPermissionAction) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescGZIP(), []int{0}
}

type IpPermission_Action int32

const (
	IpPermission_ACTION_UNSPECIFIED IpPermission_Action = 0
	IpPermission_PULL               IpPermission_Action = 1
	IpPermission_PUSH               IpPermission_Action = 2
)

// Enum value maps for IpPermission_Action.
var (
	IpPermission_Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "PULL",
		2: "PUSH",
	}
	IpPermission_Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"PULL":               1,
		"PUSH":               2,
	}
)

func (x IpPermission_Action) Enum() *IpPermission_Action {
	p := new(IpPermission_Action)
	*p = x
	return p
}

func (x IpPermission_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IpPermission_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_enumTypes[1].Descriptor()
}

func (IpPermission_Action) Type() protoreflect.EnumType {
	return &file_yandex_cloud_containerregistry_v1_ip_permission_proto_enumTypes[1]
}

func (x IpPermission_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IpPermission_Action.Descriptor instead.
func (IpPermission_Action) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescGZIP(), []int{0, 0}
}

type IpPermission struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Action        IpPermission_Action    `protobuf:"varint,1,opt,name=action,proto3,enum=yandex.cloud.containerregistry.v1.IpPermission_Action" json:"action,omitempty"`
	Ip            string                 `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IpPermission) Reset() {
	*x = IpPermission{}
	mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPermission) ProtoMessage() {}

func (x *IpPermission) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPermission.ProtoReflect.Descriptor instead.
func (*IpPermission) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescGZIP(), []int{0}
}

func (x *IpPermission) GetAction() IpPermission_Action {
	if x != nil {
		return x.Action
	}
	return IpPermission_ACTION_UNSPECIFIED
}

func (x *IpPermission) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type IpPermissionDelta struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The action that is being performed on an ip permission.
	Action IpPermissionAction `protobuf:"varint,1,opt,name=action,proto3,enum=yandex.cloud.containerregistry.v1.IpPermissionAction" json:"action,omitempty"`
	// Ip permission.
	IpPermission  *IpPermission `protobuf:"bytes,2,opt,name=ip_permission,json=ipPermission,proto3" json:"ip_permission,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IpPermissionDelta) Reset() {
	*x = IpPermissionDelta{}
	mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IpPermissionDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPermissionDelta) ProtoMessage() {}

func (x *IpPermissionDelta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IpPermissionDelta.ProtoReflect.Descriptor instead.
func (*IpPermissionDelta) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescGZIP(), []int{1}
}

func (x *IpPermissionDelta) GetAction() IpPermissionAction {
	if x != nil {
		return x.Action
	}
	return IpPermissionAction_IP_PERMISSION_ACTION_UNSPECIFIED
}

func (x *IpPermissionDelta) GetIpPermission() *IpPermission {
	if x != nil {
		return x.IpPermission
	}
	return nil
}

var File_yandex_cloud_containerregistry_v1_ip_permission_proto protoreflect.FileDescriptor

const file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc = "" +
	"\n" +
	"5yandex/cloud/containerregistry/v1/ip_permission.proto\x12!yandex.cloud.containerregistry.v1\x1a\x1dyandex/cloud/validation.proto\"\xa4\x01\n" +
	"\fIpPermission\x12N\n" +
	"\x06action\x18\x01 \x01(\x0e26.yandex.cloud.containerregistry.v1.IpPermission.ActionR\x06action\x12\x0e\n" +
	"\x02ip\x18\x02 \x01(\tR\x02ip\"4\n" +
	"\x06Action\x12\x16\n" +
	"\x12ACTION_UNSPECIFIED\x10\x00\x12\b\n" +
	"\x04PULL\x10\x01\x12\b\n" +
	"\x04PUSH\x10\x02\"\xc4\x01\n" +
	"\x11IpPermissionDelta\x12S\n" +
	"\x06action\x18\x01 \x01(\x0e25.yandex.cloud.containerregistry.v1.IpPermissionActionB\x04\xe8\xc71\x01R\x06action\x12Z\n" +
	"\rip_permission\x18\x02 \x01(\v2/.yandex.cloud.containerregistry.v1.IpPermissionB\x04\xe8\xc71\x01R\fipPermission*O\n" +
	"\x12IpPermissionAction\x12$\n" +
	" IP_PERMISSION_ACTION_UNSPECIFIED\x10\x00\x12\a\n" +
	"\x03ADD\x10\x01\x12\n" +
	"\n" +
	"\x06REMOVE\x10\x02B\x80\x01\n" +
	"%yandex.cloud.api.containerregistry.v1ZWgithub.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1;containerregistryb\x06proto3"

var (
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescData []byte
)

func file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc), len(file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc)))
	})
	return file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescData
}

var file_yandex_cloud_containerregistry_v1_ip_permission_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_containerregistry_v1_ip_permission_proto_goTypes = []any{
	(IpPermissionAction)(0),   // 0: yandex.cloud.containerregistry.v1.IpPermissionAction
	(IpPermission_Action)(0),  // 1: yandex.cloud.containerregistry.v1.IpPermission.Action
	(*IpPermission)(nil),      // 2: yandex.cloud.containerregistry.v1.IpPermission
	(*IpPermissionDelta)(nil), // 3: yandex.cloud.containerregistry.v1.IpPermissionDelta
}
var file_yandex_cloud_containerregistry_v1_ip_permission_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.containerregistry.v1.IpPermission.action:type_name -> yandex.cloud.containerregistry.v1.IpPermission.Action
	0, // 1: yandex.cloud.containerregistry.v1.IpPermissionDelta.action:type_name -> yandex.cloud.containerregistry.v1.IpPermissionAction
	2, // 2: yandex.cloud.containerregistry.v1.IpPermissionDelta.ip_permission:type_name -> yandex.cloud.containerregistry.v1.IpPermission
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_containerregistry_v1_ip_permission_proto_init() }
func file_yandex_cloud_containerregistry_v1_ip_permission_proto_init() {
	if File_yandex_cloud_containerregistry_v1_ip_permission_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc), len(file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_containerregistry_v1_ip_permission_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_containerregistry_v1_ip_permission_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_containerregistry_v1_ip_permission_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes,
	}.Build()
	File_yandex_cloud_containerregistry_v1_ip_permission_proto = out.File
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_depIdxs = nil
}
