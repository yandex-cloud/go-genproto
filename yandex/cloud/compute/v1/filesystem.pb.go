// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: yandex/cloud/compute/v1/filesystem.proto

package compute

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Filesystem_Status int32

const (
	Filesystem_STATUS_UNSPECIFIED Filesystem_Status = 0
	// Filesystem is being created.
	Filesystem_CREATING Filesystem_Status = 1
	// Filesystem is ready to use.
	Filesystem_READY Filesystem_Status = 2
	// Filesystem encountered a problem and cannot operate.
	Filesystem_ERROR Filesystem_Status = 3
	// Filesystem is being deleted.
	Filesystem_DELETING Filesystem_Status = 4
)

// Enum value maps for Filesystem_Status.
var (
	Filesystem_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "ERROR",
		4: "DELETING",
	}
	Filesystem_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"READY":              2,
		"ERROR":              3,
		"DELETING":           4,
	}
)

func (x Filesystem_Status) Enum() *Filesystem_Status {
	p := new(Filesystem_Status)
	*p = x
	return p
}

func (x Filesystem_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Filesystem_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_compute_v1_filesystem_proto_enumTypes[0].Descriptor()
}

func (Filesystem_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_compute_v1_filesystem_proto_enumTypes[0]
}

func (x Filesystem_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Filesystem_Status.Descriptor instead.
func (Filesystem_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_filesystem_proto_rawDescGZIP(), []int{0, 0}
}

// A Filesystem resource. For more information, see [Filesystems](/docs/compute/concepts/filesystem).
type Filesystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the filesystem.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the filesystem belongs to.
	FolderId  string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the filesystem. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the filesystem. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the filesystem type.
	TypeId string `protobuf:"bytes,7,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	// ID of the availability zone where the filesystem resides.
	ZoneId string `protobuf:"bytes,8,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Size of the filesystem, specified in bytes.
	Size int64 `protobuf:"varint,9,opt,name=size,proto3" json:"size,omitempty"`
	// Block size of the filesystem, specifiedin bytes.
	BlockSize int64 `protobuf:"varint,10,opt,name=block_size,json=blockSize,proto3" json:"block_size,omitempty"`
	// Current status of the filesystem.
	Status Filesystem_Status `protobuf:"varint,11,opt,name=status,proto3,enum=yandex.cloud.compute.v1.Filesystem_Status" json:"status,omitempty"`
}

func (x *Filesystem) Reset() {
	*x = Filesystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filesystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filesystem) ProtoMessage() {}

func (x *Filesystem) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filesystem.ProtoReflect.Descriptor instead.
func (*Filesystem) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_filesystem_proto_rawDescGZIP(), []int{0}
}

func (x *Filesystem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Filesystem) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Filesystem) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Filesystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Filesystem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Filesystem) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Filesystem) GetTypeId() string {
	if x != nil {
		return x.TypeId
	}
	return ""
}

func (x *Filesystem) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *Filesystem) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Filesystem) GetBlockSize() int64 {
	if x != nil {
		return x.BlockSize
	}
	return 0
}

func (x *Filesystem) GetStatus() Filesystem_Status {
	if x != nil {
		return x.Status
	}
	return Filesystem_STATUS_UNSPECIFIED
}

var File_yandex_cloud_compute_v1_filesystem_proto protoreflect.FileDescriptor

var file_yandex_cloud_compute_v1_filesystem_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab, 0x04, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x79,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x79, 0x70,
	0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x42, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_compute_v1_filesystem_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_filesystem_proto_rawDescData = file_yandex_cloud_compute_v1_filesystem_proto_rawDesc
)

func file_yandex_cloud_compute_v1_filesystem_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_filesystem_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_compute_v1_filesystem_proto_rawDescData)
	})
	return file_yandex_cloud_compute_v1_filesystem_proto_rawDescData
}

var file_yandex_cloud_compute_v1_filesystem_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_compute_v1_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_compute_v1_filesystem_proto_goTypes = []interface{}{
	(Filesystem_Status)(0),      // 0: yandex.cloud.compute.v1.Filesystem.Status
	(*Filesystem)(nil),          // 1: yandex.cloud.compute.v1.Filesystem
	nil,                         // 2: yandex.cloud.compute.v1.Filesystem.LabelsEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_compute_v1_filesystem_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.compute.v1.Filesystem.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.compute.v1.Filesystem.labels:type_name -> yandex.cloud.compute.v1.Filesystem.LabelsEntry
	0, // 2: yandex.cloud.compute.v1.Filesystem.status:type_name -> yandex.cloud.compute.v1.Filesystem.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_filesystem_proto_init() }
func file_yandex_cloud_compute_v1_filesystem_proto_init() {
	if File_yandex_cloud_compute_v1_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_compute_v1_filesystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filesystem); i {
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
			RawDescriptor: file_yandex_cloud_compute_v1_filesystem_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_filesystem_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_filesystem_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_compute_v1_filesystem_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_compute_v1_filesystem_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_filesystem_proto = out.File
	file_yandex_cloud_compute_v1_filesystem_proto_rawDesc = nil
	file_yandex_cloud_compute_v1_filesystem_proto_goTypes = nil
	file_yandex_cloud_compute_v1_filesystem_proto_depIdxs = nil
}
