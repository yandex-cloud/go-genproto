// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/greenplum/v1/resource_groups.proto

package greenplum

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	IsUserDefined *wrapperspb.BoolValue `protobuf:"bytes,2,opt,name=is_user_defined,json=isUserDefined,proto3" json:"is_user_defined,omitempty"`
	// References to CONCURRENCY from gp resource group parameter
	Concurrency *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	// References to CPU_RATE_LIMIT from gp resource group parameter
	CpuRateLimit *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=cpu_rate_limit,json=cpuRateLimit,proto3" json:"cpu_rate_limit,omitempty"`
	// References to MEMORY_LIMIT from gp resource group parameter
	MemoryLimit *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=memory_limit,json=memoryLimit,proto3" json:"memory_limit,omitempty"`
	// References to MEMORY_SHARED_QUOTA from gp resource group parameter
	MemorySharedQuota *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=memory_shared_quota,json=memorySharedQuota,proto3" json:"memory_shared_quota,omitempty"`
	// References to MEMORY_SPILL_RATIO from gp resource group parameter
	MemorySpillRatio *wrapperspb.Int64Value `protobuf:"bytes,7,opt,name=memory_spill_ratio,json=memorySpillRatio,proto3" json:"memory_spill_ratio,omitempty"`
}

func (x *ResourceGroup) Reset() {
	*x = ResourceGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceGroup) ProtoMessage() {}

func (x *ResourceGroup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceGroup.ProtoReflect.Descriptor instead.
func (*ResourceGroup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceGroup) GetIsUserDefined() *wrapperspb.BoolValue {
	if x != nil {
		return x.IsUserDefined
	}
	return nil
}

func (x *ResourceGroup) GetConcurrency() *wrapperspb.Int64Value {
	if x != nil {
		return x.Concurrency
	}
	return nil
}

func (x *ResourceGroup) GetCpuRateLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.CpuRateLimit
	}
	return nil
}

func (x *ResourceGroup) GetMemoryLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.MemoryLimit
	}
	return nil
}

func (x *ResourceGroup) GetMemorySharedQuota() *wrapperspb.Int64Value {
	if x != nil {
		return x.MemorySharedQuota
	}
	return nil
}

func (x *ResourceGroup) GetMemorySpillRatio() *wrapperspb.Int64Value {
	if x != nil {
		return x.MemorySpillRatio
	}
	return nil
}

var File_yandex_cloud_mdb_greenplum_v1_resource_groups_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDesc = []byte{
	0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70, 0x6c, 0x75,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x9b, 0x04, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x37, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x23, 0xe8, 0xc7, 0x31, 0x01, 0xf2, 0xc7, 0x31, 0x12, 0x5e, 0x5b, 0x5e,
	0x5c, 0x7c, 0x2f, 0x2a, 0x3f, 0x2e, 0x2c, 0x3b, 0x22, 0x27, 0x3c, 0x3e, 0x5d, 0x2b, 0x24, 0x8a,
	0xc8, 0x31, 0x05, 0x33, 0x2d, 0x32, 0x30, 0x30, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x42,
	0x0a, 0x0f, 0x69, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0d, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x66, 0x69, 0x6e,
	0x65, 0x64, 0x12, 0x46, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0xfa, 0xc7, 0x31, 0x03, 0x3e, 0x3d, 0x30, 0x52, 0x0b, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x4c, 0x0a, 0x0e, 0x63, 0x70,
	0x75, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x09, 0xfa, 0xc7, 0x31, 0x05, 0x31, 0x2d, 0x31, 0x30, 0x30, 0x52, 0x0c, 0x63, 0x70, 0x75, 0x52,
	0x61, 0x74, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x49, 0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa, 0xc7, 0x31,
	0x05, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x12, 0x56, 0x0a, 0x13, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa,
	0xc7, 0x31, 0x05, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x52, 0x11, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x53, 0x68, 0x61, 0x72, 0x65, 0x64, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x54, 0x0a, 0x12, 0x6d,
	0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x73, 0x70, 0x69, 0x6c, 0x6c, 0x5f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x09, 0xfa, 0xc7, 0x31, 0x05, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x52,
	0x10, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x53, 0x70, 0x69, 0x6c, 0x6c, 0x52, 0x61, 0x74, 0x69,
	0x6f, 0x42, 0x70, 0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70,
	0x6c, 0x75, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x67, 0x72, 0x65,
	0x65, 0x6e, 0x70, 0x6c, 0x75, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x70,
	0x6c, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescData = file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDesc
)

func file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDescData
}

var file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_goTypes = []any{
	(*ResourceGroup)(nil),         // 0: yandex.cloud.mdb.greenplum.v1.ResourceGroup
	(*wrapperspb.BoolValue)(nil),  // 1: google.protobuf.BoolValue
	(*wrapperspb.Int64Value)(nil), // 2: google.protobuf.Int64Value
}
var file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.mdb.greenplum.v1.ResourceGroup.is_user_defined:type_name -> google.protobuf.BoolValue
	2, // 1: yandex.cloud.mdb.greenplum.v1.ResourceGroup.concurrency:type_name -> google.protobuf.Int64Value
	2, // 2: yandex.cloud.mdb.greenplum.v1.ResourceGroup.cpu_rate_limit:type_name -> google.protobuf.Int64Value
	2, // 3: yandex.cloud.mdb.greenplum.v1.ResourceGroup.memory_limit:type_name -> google.protobuf.Int64Value
	2, // 4: yandex.cloud.mdb.greenplum.v1.ResourceGroup.memory_shared_quota:type_name -> google.protobuf.Int64Value
	2, // 5: yandex.cloud.mdb.greenplum.v1.ResourceGroup.memory_spill_ratio:type_name -> google.protobuf.Int64Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_init() }
func file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_init() {
	if File_yandex_cloud_mdb_greenplum_v1_resource_groups_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ResourceGroup); i {
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
			RawDescriptor: file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_greenplum_v1_resource_groups_proto = out.File
	file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_rawDesc = nil
	file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_goTypes = nil
	file_yandex_cloud_mdb_greenplum_v1_resource_groups_proto_depIdxs = nil
}
