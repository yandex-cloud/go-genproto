// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/compute/v1/placement_group.proto

package compute

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type PlacementGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the placement group. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the placement group belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the placement group.
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the placement group. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Placement group labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Placement strategy. To specify a placement strategy, send the corresponding
	// field containing approriate structure.
	//
	// Types that are assignable to PlacementStrategy:
	//
	//	*PlacementGroup_SpreadPlacementStrategy
	//	*PlacementGroup_PartitionPlacementStrategy
	PlacementStrategy isPlacementGroup_PlacementStrategy `protobuf_oneof:"placement_strategy"`
}

func (x *PlacementGroup) Reset() {
	*x = PlacementGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlacementGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementGroup) ProtoMessage() {}

func (x *PlacementGroup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlacementGroup.ProtoReflect.Descriptor instead.
func (*PlacementGroup) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_placement_group_proto_rawDescGZIP(), []int{0}
}

func (x *PlacementGroup) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlacementGroup) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *PlacementGroup) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PlacementGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PlacementGroup) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *PlacementGroup) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (m *PlacementGroup) GetPlacementStrategy() isPlacementGroup_PlacementStrategy {
	if m != nil {
		return m.PlacementStrategy
	}
	return nil
}

func (x *PlacementGroup) GetSpreadPlacementStrategy() *SpreadPlacementStrategy {
	if x, ok := x.GetPlacementStrategy().(*PlacementGroup_SpreadPlacementStrategy); ok {
		return x.SpreadPlacementStrategy
	}
	return nil
}

func (x *PlacementGroup) GetPartitionPlacementStrategy() *PartitionPlacementStrategy {
	if x, ok := x.GetPlacementStrategy().(*PlacementGroup_PartitionPlacementStrategy); ok {
		return x.PartitionPlacementStrategy
	}
	return nil
}

type isPlacementGroup_PlacementStrategy interface {
	isPlacementGroup_PlacementStrategy()
}

type PlacementGroup_SpreadPlacementStrategy struct {
	// Anti-affinity placement strategy (`spread`). Instances are distributed
	// over distinct failure domains.
	SpreadPlacementStrategy *SpreadPlacementStrategy `protobuf:"bytes,7,opt,name=spread_placement_strategy,json=spreadPlacementStrategy,proto3,oneof"`
}

type PlacementGroup_PartitionPlacementStrategy struct {
	PartitionPlacementStrategy *PartitionPlacementStrategy `protobuf:"bytes,8,opt,name=partition_placement_strategy,json=partitionPlacementStrategy,proto3,oneof"`
}

func (*PlacementGroup_SpreadPlacementStrategy) isPlacementGroup_PlacementStrategy() {}

func (*PlacementGroup_PartitionPlacementStrategy) isPlacementGroup_PlacementStrategy() {}

// This is an empty structure that must be passed to explicitly
// specify the required placement strategy.
type SpreadPlacementStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SpreadPlacementStrategy) Reset() {
	*x = SpreadPlacementStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpreadPlacementStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpreadPlacementStrategy) ProtoMessage() {}

func (x *SpreadPlacementStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpreadPlacementStrategy.ProtoReflect.Descriptor instead.
func (*SpreadPlacementStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_placement_group_proto_rawDescGZIP(), []int{1}
}

type PartitionPlacementStrategy struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Partitions int64 `protobuf:"varint,1,opt,name=partitions,proto3" json:"partitions,omitempty"`
}

func (x *PartitionPlacementStrategy) Reset() {
	*x = PartitionPlacementStrategy{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PartitionPlacementStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionPlacementStrategy) ProtoMessage() {}

func (x *PartitionPlacementStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartitionPlacementStrategy.ProtoReflect.Descriptor instead.
func (*PartitionPlacementStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_compute_v1_placement_group_proto_rawDescGZIP(), []int{2}
}

func (x *PartitionPlacementStrategy) GetPartitions() int64 {
	if x != nil {
		return x.Partitions
	}
	return 0
}

var File_yandex_cloud_compute_v1_placement_group_proto protoreflect.FileDescriptor

var file_yandex_cloud_compute_v1_placement_group_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x04, 0x0a, 0x0e, 0x50, 0x6c, 0x61,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x6e, 0x0a, 0x19, 0x73, 0x70, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x67, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x48, 0x00, 0x52, 0x17, 0x73,
	0x70, 0x72, 0x65, 0x61, 0x64, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12, 0x77, 0x0a, 0x1c, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x48, 0x00, 0x52, 0x1a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x1a,
	0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x1a, 0x0a, 0x12, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79,
	0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x19, 0x0a, 0x17, 0x53, 0x70, 0x72, 0x65, 0x61, 0x64,
	0x50, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67,
	0x79, 0x22, 0x45, 0x0a, 0x1a, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c,
	0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x72, 0x61, 0x74, 0x65, 0x67, 0x79, 0x12,
	0x27, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0xc7, 0x31, 0x03, 0x32, 0x2d, 0x35, 0x52, 0x0a, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_compute_v1_placement_group_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_placement_group_proto_rawDescData = file_yandex_cloud_compute_v1_placement_group_proto_rawDesc
)

func file_yandex_cloud_compute_v1_placement_group_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_placement_group_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_placement_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_compute_v1_placement_group_proto_rawDescData)
	})
	return file_yandex_cloud_compute_v1_placement_group_proto_rawDescData
}

var file_yandex_cloud_compute_v1_placement_group_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_compute_v1_placement_group_proto_goTypes = []any{
	(*PlacementGroup)(nil),             // 0: yandex.cloud.compute.v1.PlacementGroup
	(*SpreadPlacementStrategy)(nil),    // 1: yandex.cloud.compute.v1.SpreadPlacementStrategy
	(*PartitionPlacementStrategy)(nil), // 2: yandex.cloud.compute.v1.PartitionPlacementStrategy
	nil,                                // 3: yandex.cloud.compute.v1.PlacementGroup.LabelsEntry
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_compute_v1_placement_group_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.compute.v1.PlacementGroup.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: yandex.cloud.compute.v1.PlacementGroup.labels:type_name -> yandex.cloud.compute.v1.PlacementGroup.LabelsEntry
	1, // 2: yandex.cloud.compute.v1.PlacementGroup.spread_placement_strategy:type_name -> yandex.cloud.compute.v1.SpreadPlacementStrategy
	2, // 3: yandex.cloud.compute.v1.PlacementGroup.partition_placement_strategy:type_name -> yandex.cloud.compute.v1.PartitionPlacementStrategy
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_compute_v1_placement_group_proto_init() }
func file_yandex_cloud_compute_v1_placement_group_proto_init() {
	if File_yandex_cloud_compute_v1_placement_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PlacementGroup); i {
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
		file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SpreadPlacementStrategy); i {
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
		file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PartitionPlacementStrategy); i {
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
	file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0].OneofWrappers = []any{
		(*PlacementGroup_SpreadPlacementStrategy)(nil),
		(*PlacementGroup_PartitionPlacementStrategy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_compute_v1_placement_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_compute_v1_placement_group_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_compute_v1_placement_group_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_compute_v1_placement_group_proto_msgTypes,
	}.Build()
	File_yandex_cloud_compute_v1_placement_group_proto = out.File
	file_yandex_cloud_compute_v1_placement_group_proto_rawDesc = nil
	file_yandex_cloud_compute_v1_placement_group_proto_goTypes = nil
	file_yandex_cloud_compute_v1_placement_group_proto_depIdxs = nil
}
