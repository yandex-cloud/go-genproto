// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PlacementGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
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
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Placement strategy. To specify a placement strategy, send the corresponding
	// field containing approriate structure.
	//
	// Types that are valid to be assigned to PlacementStrategy:
	//
	//	*PlacementGroup_SpreadPlacementStrategy
	//	*PlacementGroup_PartitionPlacementStrategy
	PlacementStrategy isPlacementGroup_PlacementStrategy `protobuf_oneof:"placement_strategy"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *PlacementGroup) Reset() {
	*x = PlacementGroup{}
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PlacementGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlacementGroup) ProtoMessage() {}

func (x *PlacementGroup) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0]
	if x != nil {
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

func (x *PlacementGroup) GetPlacementStrategy() isPlacementGroup_PlacementStrategy {
	if x != nil {
		return x.PlacementStrategy
	}
	return nil
}

func (x *PlacementGroup) GetSpreadPlacementStrategy() *SpreadPlacementStrategy {
	if x != nil {
		if x, ok := x.PlacementStrategy.(*PlacementGroup_SpreadPlacementStrategy); ok {
			return x.SpreadPlacementStrategy
		}
	}
	return nil
}

func (x *PlacementGroup) GetPartitionPlacementStrategy() *PartitionPlacementStrategy {
	if x != nil {
		if x, ok := x.PlacementStrategy.(*PlacementGroup_PartitionPlacementStrategy); ok {
			return x.PartitionPlacementStrategy
		}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SpreadPlacementStrategy) Reset() {
	*x = SpreadPlacementStrategy{}
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SpreadPlacementStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpreadPlacementStrategy) ProtoMessage() {}

func (x *SpreadPlacementStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[1]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Partitions    int64                  `protobuf:"varint,1,opt,name=partitions,proto3" json:"partitions,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartitionPlacementStrategy) Reset() {
	*x = PartitionPlacementStrategy{}
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartitionPlacementStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartitionPlacementStrategy) ProtoMessage() {}

func (x *PartitionPlacementStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[2]
	if x != nil {
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

const file_yandex_cloud_compute_v1_placement_group_proto_rawDesc = "" +
	"\n" +
	"-yandex/cloud/compute/v1/placement_group.proto\x12\x17yandex.cloud.compute.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dyandex/cloud/validation.proto\"\xbb\x04\n" +
	"\x0ePlacementGroup\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12K\n" +
	"\x06labels\x18\x06 \x03(\v23.yandex.cloud.compute.v1.PlacementGroup.LabelsEntryR\x06labels\x12n\n" +
	"\x19spread_placement_strategy\x18\a \x01(\v20.yandex.cloud.compute.v1.SpreadPlacementStrategyH\x00R\x17spreadPlacementStrategy\x12w\n" +
	"\x1cpartition_placement_strategy\x18\b \x01(\v23.yandex.cloud.compute.v1.PartitionPlacementStrategyH\x00R\x1apartitionPlacementStrategy\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\x1a\n" +
	"\x12placement_strategy\x12\x04\xc0\xc11\x01\"\x19\n" +
	"\x17SpreadPlacementStrategy\"E\n" +
	"\x1aPartitionPlacementStrategy\x12'\n" +
	"\n" +
	"partitions\x18\x01 \x01(\x03B\a\xfa\xc71\x032-5R\n" +
	"partitionsBb\n" +
	"\x1byandex.cloud.api.compute.v1ZCgithub.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1;computeb\x06proto3"

var (
	file_yandex_cloud_compute_v1_placement_group_proto_rawDescOnce sync.Once
	file_yandex_cloud_compute_v1_placement_group_proto_rawDescData []byte
)

func file_yandex_cloud_compute_v1_placement_group_proto_rawDescGZIP() []byte {
	file_yandex_cloud_compute_v1_placement_group_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_compute_v1_placement_group_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_compute_v1_placement_group_proto_rawDesc), len(file_yandex_cloud_compute_v1_placement_group_proto_rawDesc)))
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
	file_yandex_cloud_compute_v1_placement_group_proto_msgTypes[0].OneofWrappers = []any{
		(*PlacementGroup_SpreadPlacementStrategy)(nil),
		(*PlacementGroup_PartitionPlacementStrategy)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_compute_v1_placement_group_proto_rawDesc), len(file_yandex_cloud_compute_v1_placement_group_proto_rawDesc)),
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
	file_yandex_cloud_compute_v1_placement_group_proto_goTypes = nil
	file_yandex_cloud_compute_v1_placement_group_proto_depIdxs = nil
}
