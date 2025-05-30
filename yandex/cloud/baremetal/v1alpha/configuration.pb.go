// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/baremetal/v1alpha/configuration.proto

package baremetal

import (
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

// CPU configuration.
type CPU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the CPU.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Vendor of the CPU.
	Vendor string `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	// Number of cores.
	Cores int64 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
}

func (x *CPU) Reset() {
	*x = CPU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CPU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CPU) ProtoMessage() {}

func (x *CPU) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CPU.ProtoReflect.Descriptor instead.
func (*CPU) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescGZIP(), []int{0}
}

func (x *CPU) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CPU) GetVendor() string {
	if x != nil {
		return x.Vendor
	}
	return ""
}

func (x *CPU) GetCores() int64 {
	if x != nil {
		return x.Cores
	}
	return 0
}

type DiskDriveConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Type of the disk drive.
	Type DiskDriveType `protobuf:"varint,1,opt,name=type,proto3,enum=yandex.cloud.baremetal.v1alpha.DiskDriveType" json:"type,omitempty"`
	// Number of disk drives.
	DiskCount int64 `protobuf:"varint,2,opt,name=disk_count,json=diskCount,proto3" json:"disk_count,omitempty"`
	// Size of a single disk drive in gibibytes (2^30 bytes).
	DiskSizeGib int64 `protobuf:"varint,3,opt,name=disk_size_gib,json=diskSizeGib,proto3" json:"disk_size_gib,omitempty"`
}

func (x *DiskDriveConfiguration) Reset() {
	*x = DiskDriveConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DiskDriveConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DiskDriveConfiguration) ProtoMessage() {}

func (x *DiskDriveConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DiskDriveConfiguration.ProtoReflect.Descriptor instead.
func (*DiskDriveConfiguration) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescGZIP(), []int{1}
}

func (x *DiskDriveConfiguration) GetType() DiskDriveType {
	if x != nil {
		return x.Type
	}
	return DiskDriveType_DISK_DRIVE_TYPE_UNSPECIFIED
}

func (x *DiskDriveConfiguration) GetDiskCount() int64 {
	if x != nil {
		return x.DiskCount
	}
	return 0
}

func (x *DiskDriveConfiguration) GetDiskSizeGib() int64 {
	if x != nil {
		return x.DiskSizeGib
	}
	return 0
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the configuration.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the configuration.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Random-access memory (RAM) size in gibibytes (2^30 bytes).
	MemoryGib int64 `protobuf:"varint,3,opt,name=memory_gib,json=memoryGib,proto3" json:"memory_gib,omitempty"`
	// CPU configuration.
	Cpu *CPU `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	// Array of disk drive configurations.
	DiskDrives []*DiskDriveConfiguration `protobuf:"bytes,5,rep,name=disk_drives,json=diskDrives,proto3" json:"disk_drives,omitempty"`
	// Network capacity or bandwidth in gigabits per second.
	NetworkCapacityGbps int64 `protobuf:"varint,6,opt,name=network_capacity_gbps,json=networkCapacityGbps,proto3" json:"network_capacity_gbps,omitempty"`
	// Number of cpu.
	CpuNum int64 `protobuf:"varint,8,opt,name=cpu_num,json=cpuNum,proto3" json:"cpu_num,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescGZIP(), []int{2}
}

func (x *Configuration) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Configuration) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Configuration) GetMemoryGib() int64 {
	if x != nil {
		return x.MemoryGib
	}
	return 0
}

func (x *Configuration) GetCpu() *CPU {
	if x != nil {
		return x.Cpu
	}
	return nil
}

func (x *Configuration) GetDiskDrives() []*DiskDriveConfiguration {
	if x != nil {
		return x.DiskDrives
	}
	return nil
}

func (x *Configuration) GetNetworkCapacityGbps() int64 {
	if x != nil {
		return x.NetworkCapacityGbps
	}
	return 0
}

func (x *Configuration) GetCpuNum() int64 {
	if x != nil {
		return x.CpuNum
	}
	return 0
}

var File_yandex_cloud_baremetal_v1alpha_configuration_proto protoreflect.FileDescriptor

var file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x1a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x03, 0x43, 0x50, 0x55, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x65,
	0x6e, 0x64, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x65, 0x6e, 0x64,
	0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x9e, 0x01, 0x0a, 0x16, 0x44, 0x69, 0x73,
	0x6b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x69, 0x73, 0x6b,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x5f, 0x67, 0x69, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x69,
	0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x47, 0x69, 0x62, 0x22, 0xb5, 0x02, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f, 0x67, 0x69, 0x62, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x47, 0x69, 0x62, 0x12, 0x35,
	0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x43, 0x50, 0x55,
	0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x57, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x44, 0x69, 0x73, 0x6b,
	0x44, 0x72, 0x69, 0x76, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x73, 0x12, 0x32,
	0x0a, 0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69,
	0x74, 0x79, 0x5f, 0x67, 0x62, 0x70, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x47, 0x62,
	0x70, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x70, 0x75, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x70, 0x75, 0x4e, 0x75, 0x6d, 0x4a, 0x04, 0x08, 0x07, 0x10,
	0x08, 0x42, 0x72, 0x0a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65,
	0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescOnce sync.Once
	file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescData = file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDesc
)

func file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescGZIP() []byte {
	file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescData)
	})
	return file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDescData
}

var file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_baremetal_v1alpha_configuration_proto_goTypes = []any{
	(*CPU)(nil),                    // 0: yandex.cloud.baremetal.v1alpha.CPU
	(*DiskDriveConfiguration)(nil), // 1: yandex.cloud.baremetal.v1alpha.DiskDriveConfiguration
	(*Configuration)(nil),          // 2: yandex.cloud.baremetal.v1alpha.Configuration
	(DiskDriveType)(0),             // 3: yandex.cloud.baremetal.v1alpha.DiskDriveType
}
var file_yandex_cloud_baremetal_v1alpha_configuration_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.baremetal.v1alpha.DiskDriveConfiguration.type:type_name -> yandex.cloud.baremetal.v1alpha.DiskDriveType
	0, // 1: yandex.cloud.baremetal.v1alpha.Configuration.cpu:type_name -> yandex.cloud.baremetal.v1alpha.CPU
	1, // 2: yandex.cloud.baremetal.v1alpha.Configuration.disk_drives:type_name -> yandex.cloud.baremetal.v1alpha.DiskDriveConfiguration
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_baremetal_v1alpha_configuration_proto_init() }
func file_yandex_cloud_baremetal_v1alpha_configuration_proto_init() {
	if File_yandex_cloud_baremetal_v1alpha_configuration_proto != nil {
		return
	}
	file_yandex_cloud_baremetal_v1alpha_disk_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CPU); i {
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
		file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*DiskDriveConfiguration); i {
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
		file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Configuration); i {
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
			RawDescriptor: file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_baremetal_v1alpha_configuration_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_baremetal_v1alpha_configuration_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_baremetal_v1alpha_configuration_proto_msgTypes,
	}.Build()
	File_yandex_cloud_baremetal_v1alpha_configuration_proto = out.File
	file_yandex_cloud_baremetal_v1alpha_configuration_proto_rawDesc = nil
	file_yandex_cloud_baremetal_v1alpha_configuration_proto_goTypes = nil
	file_yandex_cloud_baremetal_v1alpha_configuration_proto_depIdxs = nil
}
