// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/containerregistry/v1/ip_permission.proto

package containerregistry

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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action IpPermission_Action `protobuf:"varint,1,opt,name=action,proto3,enum=yandex.cloud.containerregistry.v1.IpPermission_Action" json:"action,omitempty"`
	Ip     string              `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *IpPermission) Reset() {
	*x = IpPermission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpPermission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPermission) ProtoMessage() {}

func (x *IpPermission) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The action that is being performed on an ip permission.
	Action IpPermissionAction `protobuf:"varint,1,opt,name=action,proto3,enum=yandex.cloud.containerregistry.v1.IpPermissionAction" json:"action,omitempty"`
	// Ip permission.
	IpPermission *IpPermission `protobuf:"bytes,2,opt,name=ip_permission,json=ipPermission,proto3" json:"ip_permission,omitempty"`
}

func (x *IpPermissionDelta) Reset() {
	*x = IpPermissionDelta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IpPermissionDelta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IpPermissionDelta) ProtoMessage() {}

func (x *IpPermissionDelta) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x69, 0x70, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa4, 0x01, 0x0a, 0x0c, 0x49, 0x70,
	0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x06, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x34, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04,
	0x50, 0x55, 0x4c, 0x4c, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x55, 0x53, 0x48, 0x10, 0x02,
	0x22, 0xc4, 0x01, 0x0a, 0x11, 0x49, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x6c, 0x74, 0x61, 0x12, 0x53, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x70, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x04, 0xe8,
	0xc7, 0x31, 0x01, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x0d, 0x69,
	0x70, 0x5f, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x70, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0c, 0x69, 0x70, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0x4f, 0x0a, 0x12, 0x49, 0x70, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a,
	0x20, 0x49, 0x50, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x44, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x52, 0x45, 0x4d, 0x4f, 0x56, 0x45, 0x10, 0x02, 0x42, 0x80, 0x01, 0x0a, 0x25, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescData = file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc
)

func file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*IpPermission); i {
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
		file_yandex_cloud_containerregistry_v1_ip_permission_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*IpPermissionDelta); i {
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
			RawDescriptor: file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc,
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
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_rawDesc = nil
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_ip_permission_proto_depIdxs = nil
}
