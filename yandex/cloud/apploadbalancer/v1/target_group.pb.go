// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.6.1
// source: yandex/cloud/apploadbalancer/v1/target_group.proto

package apploadbalancer

import (
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TargetGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output only. ID of the target group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name is unique within the folder. 3-63 characters long.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the target group. 0-256 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the folder that the target group belongs to.
	FolderId string `protobuf:"bytes,4,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// NOTE: all endpoints must use the same address_type - either ip or hostname.
	Targets []*Target `protobuf:"bytes,6,rep,name=targets,proto3" json:"targets,omitempty"`
	// Creation timestamp for the target group.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *TargetGroup) Reset() {
	*x = TargetGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetGroup) ProtoMessage() {}

func (x *TargetGroup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetGroup.ProtoReflect.Descriptor instead.
func (*TargetGroup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescGZIP(), []int{0}
}

func (x *TargetGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TargetGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TargetGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TargetGroup) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *TargetGroup) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TargetGroup) GetTargets() []*Target {
	if x != nil {
		return x.Targets
	}
	return nil
}

func (x *TargetGroup) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Target struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to AddressType:
	//	*Target_IpAddress
	AddressType isTarget_AddressType `protobuf_oneof:"address_type"`
	// ID of the subnet that target connected to.
	SubnetId string `protobuf:"bytes,3,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
}

func (x *Target) Reset() {
	*x = Target{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Target) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Target) ProtoMessage() {}

func (x *Target) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Target.ProtoReflect.Descriptor instead.
func (*Target) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescGZIP(), []int{1}
}

func (m *Target) GetAddressType() isTarget_AddressType {
	if m != nil {
		return m.AddressType
	}
	return nil
}

func (x *Target) GetIpAddress() string {
	if x, ok := x.GetAddressType().(*Target_IpAddress); ok {
		return x.IpAddress
	}
	return ""
}

func (x *Target) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

type isTarget_AddressType interface {
	isTarget_AddressType()
}

type Target_IpAddress struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,json=ipAddress,proto3,oneof"`
}

func (*Target_IpAddress) isTarget_AddressType() {}

var File_yandex_cloud_apploadbalancer_v1_target_group_proto protoreflect.FileDescriptor

var file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfb, 0x02, 0x0a, 0x0b, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x50, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x07, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x52, 0x07, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x5c, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1f, 0x0a,
	0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x42, 0x14, 0x0a, 0x0c, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x04, 0xc0, 0xc1, 0x31,
	0x01, 0x42, 0x7a, 0x0a, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescData = file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDesc
)

func file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescData)
	})
	return file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDescData
}

var file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_apploadbalancer_v1_target_group_proto_goTypes = []interface{}{
	(*TargetGroup)(nil),         // 0: yandex.cloud.apploadbalancer.v1.TargetGroup
	(*Target)(nil),              // 1: yandex.cloud.apploadbalancer.v1.Target
	nil,                         // 2: yandex.cloud.apploadbalancer.v1.TargetGroup.LabelsEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_apploadbalancer_v1_target_group_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.apploadbalancer.v1.TargetGroup.labels:type_name -> yandex.cloud.apploadbalancer.v1.TargetGroup.LabelsEntry
	1, // 1: yandex.cloud.apploadbalancer.v1.TargetGroup.targets:type_name -> yandex.cloud.apploadbalancer.v1.Target
	3, // 2: yandex.cloud.apploadbalancer.v1.TargetGroup.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_apploadbalancer_v1_target_group_proto_init() }
func file_yandex_cloud_apploadbalancer_v1_target_group_proto_init() {
	if File_yandex_cloud_apploadbalancer_v1_target_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetGroup); i {
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
		file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Target); i {
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
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Target_IpAddress)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_apploadbalancer_v1_target_group_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_apploadbalancer_v1_target_group_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_apploadbalancer_v1_target_group_proto_msgTypes,
	}.Build()
	File_yandex_cloud_apploadbalancer_v1_target_group_proto = out.File
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_rawDesc = nil
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_target_group_proto_depIdxs = nil
}
