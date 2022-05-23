// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/compute/v1/snapshot.proto

package compute

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Snapshot_Status int32

const (
	Snapshot_STATUS_UNSPECIFIED Snapshot_Status = 0
	// Snapshot is being created.
	Snapshot_CREATING Snapshot_Status = 1
	// Snapshot is ready to use.
	Snapshot_READY Snapshot_Status = 2
	// Snapshot encountered a problem and cannot operate.
	Snapshot_ERROR Snapshot_Status = 3
	// Snapshot is being deleted.
	Snapshot_DELETING Snapshot_Status = 4
)

// Enum value maps for Snapshot_Status.
var (
	Snapshot_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "ERROR",
		4: "DELETING",
	}
	Snapshot_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"READY":              2,
		"ERROR":              3,
		"DELETING":           4,
	}
)

func (x Snapshot_Status) Enum() *Snapshot_Status {
	p := new(Snapshot_Status)
	*p = x
	return p
}

func (x Snapshot_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Snapshot_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_compute_v1_snapshot_proto_enumTypes[0].Descriptor()
}

func (Snapshot_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_compute_v1_snapshot_proto_enumTypes[0]
}

func (x Snapshot_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Snapshot_Status.Descriptor instead.
func (Snapshot_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_snapshot_proto_rawDescGZIP(), []int{0, 0}
}

// A Snapshot resource. For more information, see [Snapshots](/docs/compute/concepts/snapshot).
type Snapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the snapshot.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the snapshot belongs to.
	FolderId  string                 `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the snapshot. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the snapshot. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Size of the snapshot, specified in bytes.
	StorageSize int64 `protobuf:"varint,7,opt,name=storage_size,json=storageSize,proto3" json:"storage_size,omitempty"` // delta from prev snapshot from same disk
	// Size of the disk when the snapshot was created, specified in bytes.
	DiskSize int64 `protobuf:"varint,8,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"` // at snapshot moment
	// License IDs that indicate which licenses are attached to this resource.
	// License IDs are used to calculate additional charges for the use of the virtual machine.
	//
	// The correct license ID is generated by Yandex Cloud. IDs are inherited by new resources created from this resource.
	//
	// If you know the license IDs, specify them when you create the image.
	// For example, if you create a disk image using a third-party utility and load it into Yandex Object Storage, the license IDs will be lost.
	// You can specify them in the [yandex.cloud.compute.v1.ImageService.Create] request.
	ProductIds []string `protobuf:"bytes,9,rep,name=product_ids,json=productIds,proto3" json:"product_ids,omitempty"`
	// Current status of the snapshot.
	Status Snapshot_Status `protobuf:"varint,10,opt,name=status,proto3,enum=yandex.cloud.compute.v1.Snapshot_Status" json:"status,omitempty"`
	// ID of the source disk used to create this snapshot.
	SourceDiskId string `protobuf:"bytes,11,opt,name=source_disk_id,json=sourceDiskId,proto3" json:"source_disk_id,omitempty"`
}

func (x *Snapshot) Reset() {
	*x = Snapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Snapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Snapshot) ProtoMessage() {}

func (x *Snapshot) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Snapshot.ProtoReflect.Descriptor instead.
func (*Snapshot) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *Snapshot) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Snapshot) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Snapshot) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Snapshot) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Snapshot) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Snapshot) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Snapshot) GetStorageSize() int64 {
	if x != nil {
		return x.StorageSize
	}
	return 0
}

func (x *Snapshot) GetDiskSize() int64 {
	if x != nil {
		return x.DiskSize
	}
	return 0
}

func (x *Snapshot) GetProductIds() []string {
	if x != nil {
		return x.ProductIds
	}
	return nil
}

func (x *Snapshot) GetStatus() Snapshot_Status {
	if x != nil {
		return x.Status
	}
	return Snapshot_STATUS_UNSPECIFIED
}

func (x *Snapshot) GetSourceDiskId() string {
	if x != nil {
		return x.SourceDiskId
	}
	return ""
}

var File_yandex_cloud_compute_v1_snapshot_proto protoreflect.FileDescriptor

var file_yandex_cloud_compute_v1_snapshot_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xc7, 0x04, 0x0a, 0x08, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x64, 0x73, 0x12, 0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x69, 0x73, 0x6b, 0x49, 0x64, 0x1a, 0x39, 0x0a,
	0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44,
	0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x42, 0x62, 0x0a, 0x1b,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_compute_v1_snapshot_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_snapshot_proto_rawDescData = file_yandex_cloud_compute_v1_snapshot_proto_rawDesc
)

func file_yandex_cloud_compute_v1_snapshot_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_snapshot_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_compute_v1_snapshot_proto_rawDescData)
	})
	return file_yandex_cloud_compute_v1_snapshot_proto_rawDescData
}

var file_yandex_cloud_compute_v1_snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_compute_v1_snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_compute_v1_snapshot_proto_goTypes = []interface{}{
	(Snapshot_Status)(0),          // 0: yandex.cloud.compute.v1.Snapshot.Status
	(*Snapshot)(nil),              // 1: yandex.cloud.compute.v1.Snapshot
	nil,                           // 2: yandex.cloud.compute.v1.Snapshot.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_compute_v1_snapshot_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.compute.v1.Snapshot.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.compute.v1.Snapshot.labels:type_name -> yandex.cloud.compute.v1.Snapshot.LabelsEntry
	0, // 2: yandex.cloud.compute.v1.Snapshot.status:type_name -> yandex.cloud.compute.v1.Snapshot.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_snapshot_proto_init() }
func file_yandex_cloud_compute_v1_snapshot_proto_init() {
	if File_yandex_cloud_compute_v1_snapshot_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_compute_v1_snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Snapshot); i {
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
			RawDescriptor: file_yandex_cloud_compute_v1_snapshot_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_snapshot_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_snapshot_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_compute_v1_snapshot_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_compute_v1_snapshot_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_snapshot_proto = out.File
	file_yandex_cloud_compute_v1_snapshot_proto_rawDesc = nil
	file_yandex_cloud_compute_v1_snapshot_proto_goTypes = nil
	file_yandex_cloud_compute_v1_snapshot_proto_depIdxs = nil
}
