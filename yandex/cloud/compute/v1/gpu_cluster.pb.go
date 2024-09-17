// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/compute/v1/gpu_cluster.proto

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

type GpuInterconnectType int32

const (
	GpuInterconnectType_GPU_INTERCONNECT_TYPE_UNSPECIFIED GpuInterconnectType = 0
	// InfiniBand interconnect.
	GpuInterconnectType_INFINIBAND GpuInterconnectType = 1
)

// Enum value maps for GpuInterconnectType.
var (
	GpuInterconnectType_name = map[int32]string{
		0: "GPU_INTERCONNECT_TYPE_UNSPECIFIED",
		1: "INFINIBAND",
	}
	GpuInterconnectType_value = map[string]int32{
		"GPU_INTERCONNECT_TYPE_UNSPECIFIED": 0,
		"INFINIBAND":                        1,
	}
)

func (x GpuInterconnectType) Enum() *GpuInterconnectType {
	p := new(GpuInterconnectType)
	*p = x
	return p
}

func (x GpuInterconnectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GpuInterconnectType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_compute_v1_gpu_cluster_proto_enumTypes[0].Descriptor()
}

func (GpuInterconnectType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_compute_v1_gpu_cluster_proto_enumTypes[0]
}

func (x GpuInterconnectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GpuInterconnectType.Descriptor instead.
func (GpuInterconnectType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{0}
}

type GpuCluster_Status int32

const (
	GpuCluster_STATUS_UNSPECIFIED GpuCluster_Status = 0
	// GPU cluster is being created.
	GpuCluster_CREATING GpuCluster_Status = 1
	// GPU cluster is ready to use.
	GpuCluster_READY GpuCluster_Status = 2
	// GPU cluster encountered a problem and cannot operate.
	GpuCluster_ERROR GpuCluster_Status = 3
	// GPU cluster is being deleted.
	GpuCluster_DELETING GpuCluster_Status = 4
)

// Enum value maps for GpuCluster_Status.
var (
	GpuCluster_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "READY",
		3: "ERROR",
		4: "DELETING",
	}
	GpuCluster_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"READY":              2,
		"ERROR":              3,
		"DELETING":           4,
	}
)

func (x GpuCluster_Status) Enum() *GpuCluster_Status {
	p := new(GpuCluster_Status)
	*p = x
	return p
}

func (x GpuCluster_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GpuCluster_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_compute_v1_gpu_cluster_proto_enumTypes[1].Descriptor()
}

func (GpuCluster_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_compute_v1_gpu_cluster_proto_enumTypes[1]
}

func (x GpuCluster_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GpuCluster_Status.Descriptor instead.
func (GpuCluster_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{0, 0}
}

// A GPU cluster. For details about the concept, see [documentation](/docs/compute/concepts/gpu-cluster).
type GpuCluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of GPU cluster.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the GPU cluster belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the GPU cluster.
	//
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the GPU cluster.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// GPU cluster labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the GPU cluster.
	Status GpuCluster_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.compute.v1.GpuCluster_Status" json:"status,omitempty"`
	// ID of the availability zone where the GPU cluster resides.
	ZoneId string `protobuf:"bytes,8,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Type of interconnect used for this GPU cluster.
	InterconnectType GpuInterconnectType `protobuf:"varint,9,opt,name=interconnect_type,json=interconnectType,proto3,enum=yandex.cloud.compute.v1.GpuInterconnectType" json:"interconnect_type,omitempty"`
}

func (x *GpuCluster) Reset() {
	*x = GpuCluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_gpu_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GpuCluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GpuCluster) ProtoMessage() {}

func (x *GpuCluster) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_gpu_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GpuCluster.ProtoReflect.Descriptor instead.
func (*GpuCluster) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *GpuCluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GpuCluster) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *GpuCluster) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GpuCluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GpuCluster) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *GpuCluster) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *GpuCluster) GetStatus() GpuCluster_Status {
	if x != nil {
		return x.Status
	}
	return GpuCluster_STATUS_UNSPECIFIED
}

func (x *GpuCluster) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *GpuCluster) GetInterconnectType() GpuInterconnectType {
	if x != nil {
		return x.InterconnectType
	}
	return GpuInterconnectType_GPU_INTERCONNECT_TYPE_UNSPECIFIED
}

var File_yandex_cloud_compute_v1_gpu_cluster_proto protoreflect.FileDescriptor

var file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDesc = []byte{
	0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x70, 0x75, 0x5f, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x04, 0x0a, 0x0a, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70, 0x75, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x59, 0x0a, 0x11, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x70,
	0x75, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x10, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x52,
	0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x09,
	0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52,
	0x4f, 0x52, 0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x04, 0x2a, 0x4c, 0x0a, 0x13, 0x47, 0x70, 0x75, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x21, 0x47, 0x50, 0x55,
	0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x46, 0x49, 0x4e, 0x49, 0x42, 0x41, 0x4e, 0x44, 0x10, 0x01,
	0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescData = file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDesc
)

func file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescData)
	})
	return file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDescData
}

var file_yandex_cloud_compute_v1_gpu_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_compute_v1_gpu_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_compute_v1_gpu_cluster_proto_goTypes = []any{
	(GpuInterconnectType)(0),      // 0: yandex.cloud.compute.v1.GpuInterconnectType
	(GpuCluster_Status)(0),        // 1: yandex.cloud.compute.v1.GpuCluster.Status
	(*GpuCluster)(nil),            // 2: yandex.cloud.compute.v1.GpuCluster
	nil,                           // 3: yandex.cloud.compute.v1.GpuCluster.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_compute_v1_gpu_cluster_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.compute.v1.GpuCluster.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: yandex.cloud.compute.v1.GpuCluster.labels:type_name -> yandex.cloud.compute.v1.GpuCluster.LabelsEntry
	1, // 2: yandex.cloud.compute.v1.GpuCluster.status:type_name -> yandex.cloud.compute.v1.GpuCluster.Status
	0, // 3: yandex.cloud.compute.v1.GpuCluster.interconnect_type:type_name -> yandex.cloud.compute.v1.GpuInterconnectType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_gpu_cluster_proto_init() }
func file_yandex_cloud_compute_v1_gpu_cluster_proto_init() {
	if File_yandex_cloud_compute_v1_gpu_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_compute_v1_gpu_cluster_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GpuCluster); i {
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
			RawDescriptor: file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_gpu_cluster_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_gpu_cluster_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_compute_v1_gpu_cluster_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_compute_v1_gpu_cluster_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_gpu_cluster_proto = out.File
	file_yandex_cloud_compute_v1_gpu_cluster_proto_rawDesc = nil
	file_yandex_cloud_compute_v1_gpu_cluster_proto_goTypes = nil
	file_yandex_cloud_compute_v1_gpu_cluster_proto_depIdxs = nil
}
