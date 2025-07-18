// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datasphere/v2/community.proto

package datasphere

import (
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

type Community struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the community.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Time when community was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the community.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the comminuty.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Labels of the community.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// ID of the user who created the community.
	CreatedById string `protobuf:"bytes,6,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	// ID of the organization to which community belongs.
	OrganizationId string `protobuf:"bytes,10,opt,name=organization_id,json=organizationId,proto3" json:"organization_id,omitempty"`
	// ID of the zone where this community was created
	ZoneId        string `protobuf:"bytes,11,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Community) Reset() {
	*x = Community{}
	mi := &file_yandex_cloud_datasphere_v2_community_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Community) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Community) ProtoMessage() {}

func (x *Community) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v2_community_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Community.ProtoReflect.Descriptor instead.
func (*Community) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_community_proto_rawDescGZIP(), []int{0}
}

func (x *Community) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Community) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Community) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Community) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Community) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Community) GetCreatedById() string {
	if x != nil {
		return x.CreatedById
	}
	return ""
}

func (x *Community) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *Community) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

var File_yandex_cloud_datasphere_v2_community_proto protoreflect.FileDescriptor

const file_yandex_cloud_datasphere_v2_community_proto_rawDesc = "" +
	"\n" +
	"*yandex/cloud/datasphere/v2/community.proto\x12\x1ayandex.cloud.datasphere.v2\x1a\x1fgoogle/protobuf/timestamp.proto\"\xfe\x02\n" +
	"\tCommunity\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12I\n" +
	"\x06labels\x18\x05 \x03(\v21.yandex.cloud.datasphere.v2.Community.LabelsEntryR\x06labels\x12\"\n" +
	"\rcreated_by_id\x18\x06 \x01(\tR\vcreatedById\x12'\n" +
	"\x0forganization_id\x18\n" +
	" \x01(\tR\x0eorganizationId\x12\x17\n" +
	"\azone_id\x18\v \x01(\tR\x06zoneId\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01J\x04\b\a\x10\n" +
	"Bk\n" +
	"\x1eyandex.cloud.api.datasphere.v2ZIgithub.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v2;datasphereb\x06proto3"

var (
	file_yandex_cloud_datasphere_v2_community_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v2_community_proto_rawDescData []byte
)

func file_yandex_cloud_datasphere_v2_community_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v2_community_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v2_community_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datasphere_v2_community_proto_rawDesc), len(file_yandex_cloud_datasphere_v2_community_proto_rawDesc)))
	})
	return file_yandex_cloud_datasphere_v2_community_proto_rawDescData
}

var file_yandex_cloud_datasphere_v2_community_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_datasphere_v2_community_proto_goTypes = []any{
	(*Community)(nil),             // 0: yandex.cloud.datasphere.v2.Community
	nil,                           // 1: yandex.cloud.datasphere.v2.Community.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_datasphere_v2_community_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.datasphere.v2.Community.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.datasphere.v2.Community.labels:type_name -> yandex.cloud.datasphere.v2.Community.LabelsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v2_community_proto_init() }
func file_yandex_cloud_datasphere_v2_community_proto_init() {
	if File_yandex_cloud_datasphere_v2_community_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datasphere_v2_community_proto_rawDesc), len(file_yandex_cloud_datasphere_v2_community_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datasphere_v2_community_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v2_community_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datasphere_v2_community_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v2_community_proto = out.File
	file_yandex_cloud_datasphere_v2_community_proto_goTypes = nil
	file_yandex_cloud_datasphere_v2_community_proto_depIdxs = nil
}
