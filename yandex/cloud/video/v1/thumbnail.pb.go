// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/video/v1/thumbnail.proto

package video

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

type Thumbnail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thumbnail.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the channel where the thumbnail was created.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Time when thumbnail was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Thumbnail) Reset() {
	*x = Thumbnail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thumbnail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thumbnail) ProtoMessage() {}

func (x *Thumbnail) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thumbnail.ProtoReflect.Descriptor instead.
func (*Thumbnail) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_proto_rawDescGZIP(), []int{0}
}

func (x *Thumbnail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Thumbnail) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Thumbnail) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_yandex_cloud_video_v1_thumbnail_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_thumbnail_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x7b, 0x0a, 0x09, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x64, 0x42, 0x5c, 0x0a, 0x19,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_thumbnail_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_thumbnail_proto_rawDescData = file_yandex_cloud_video_v1_thumbnail_proto_rawDesc
)

func file_yandex_cloud_video_v1_thumbnail_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_thumbnail_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_thumbnail_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_thumbnail_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_thumbnail_proto_rawDescData
}

var file_yandex_cloud_video_v1_thumbnail_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_video_v1_thumbnail_proto_goTypes = []any{
	(*Thumbnail)(nil),             // 0: yandex.cloud.video.v1.Thumbnail
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_video_v1_thumbnail_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.video.v1.Thumbnail.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_thumbnail_proto_init() }
func file_yandex_cloud_video_v1_thumbnail_proto_init() {
	if File_yandex_cloud_video_v1_thumbnail_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_thumbnail_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Thumbnail); i {
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
			RawDescriptor: file_yandex_cloud_video_v1_thumbnail_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_video_v1_thumbnail_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_thumbnail_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_video_v1_thumbnail_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_thumbnail_proto = out.File
	file_yandex_cloud_video_v1_thumbnail_proto_rawDesc = nil
	file_yandex_cloud_video_v1_thumbnail_proto_goTypes = nil
	file_yandex_cloud_video_v1_thumbnail_proto_depIdxs = nil
}
