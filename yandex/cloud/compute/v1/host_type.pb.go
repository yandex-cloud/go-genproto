// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/compute/v1/host_type.proto

package compute

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

// Represents host resources.
// Note: Platform can use hosts with different number of memory and cores.
// TODO: Do we need sockets here?
type HostType struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique type identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Total number of cores available for instances.
	Cores int64 `protobuf:"varint,2,opt,name=cores,proto3" json:"cores,omitempty"`
	// Ammount of memory available for instances.
	Memory int64 `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	// Number of local disks available for instances
	Disks int64 `protobuf:"varint,4,opt,name=disks,proto3" json:"disks,omitempty"`
	// Size of each local disk
	DiskSize      int64 `protobuf:"varint,5,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HostType) Reset() {
	*x = HostType{}
	mi := &file_yandex_cloud_compute_v1_host_type_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HostType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostType) ProtoMessage() {}

func (x *HostType) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_host_type_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostType.ProtoReflect.Descriptor instead.
func (*HostType) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_host_type_proto_rawDescGZIP(), []int{0}
}

func (x *HostType) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HostType) GetCores() int64 {
	if x != nil {
		return x.Cores
	}
	return 0
}

func (x *HostType) GetMemory() int64 {
	if x != nil {
		return x.Memory
	}
	return 0
}

func (x *HostType) GetDisks() int64 {
	if x != nil {
		return x.Disks
	}
	return 0
}

func (x *HostType) GetDiskSize() int64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

var File_yandex_cloud_compute_v1_host_type_proto protoreflect.FileDescriptor

const file_yandex_cloud_compute_v1_host_type_proto_rawDesc = "" +
	"\n" +
	"'yandex/cloud/compute/v1/host_type.proto\x12\x17yandex.cloud.compute.v1\"{\n" +
	"\bHostType\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x14\n" +
	"\x05cores\x18\x02 \x01(\x03R\x05cores\x12\x16\n" +
	"\x06memory\x18\x03 \x01(\x03R\x06memory\x12\x14\n" +
	"\x05disks\x18\x04 \x01(\x03R\x05disks\x12\x1b\n" +
	"\tdisk_size\x18\x05 \x01(\x03R\bdiskSizeBb\n" +
	"\x1byandex.cloud.api.compute.v1ZCgithub.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1;computeb\x06proto3"

var (
	file_yandex_cloud_compute_v1_host_type_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_host_type_proto_rawDescData []byte
)

func file_yandex_cloud_compute_v1_host_type_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_host_type_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_host_type_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_compute_v1_host_type_proto_rawDesc), len(file_yandex_cloud_compute_v1_host_type_proto_rawDesc)))
	})
	return file_yandex_cloud_compute_v1_host_type_proto_rawDescData
}

var file_yandex_cloud_compute_v1_host_type_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_compute_v1_host_type_proto_goTypes = []any{
	(*HostType)(nil), // 0: yandex.cloud.compute.v1.HostType
}
var file_yandex_cloud_compute_v1_host_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_host_type_proto_init() }
func file_yandex_cloud_compute_v1_host_type_proto_init() {
	if File_yandex_cloud_compute_v1_host_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_compute_v1_host_type_proto_rawDesc), len(file_yandex_cloud_compute_v1_host_type_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_host_type_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_host_type_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_compute_v1_host_type_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_host_type_proto = out.File
	file_yandex_cloud_compute_v1_host_type_proto_goTypes = nil
	file_yandex_cloud_compute_v1_host_type_proto_depIdxs = nil
}
