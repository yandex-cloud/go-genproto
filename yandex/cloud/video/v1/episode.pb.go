// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/video/v1/episode.proto

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

type Episode_VisibilityStatus int32

const (
	Episode_VISIBILITY_STATUS_UNSPECIFIED Episode_VisibilityStatus = 0
	Episode_PUBLISHED                     Episode_VisibilityStatus = 1
	Episode_UNPUBLISHED                   Episode_VisibilityStatus = 2
)

// Enum value maps for Episode_VisibilityStatus.
var (
	Episode_VisibilityStatus_name = map[int32]string{
		0: "VISIBILITY_STATUS_UNSPECIFIED",
		1: "PUBLISHED",
		2: "UNPUBLISHED",
	}
	Episode_VisibilityStatus_value = map[string]int32{
		"VISIBILITY_STATUS_UNSPECIFIED": 0,
		"PUBLISHED":                     1,
		"UNPUBLISHED":                   2,
	}
)

func (x Episode_VisibilityStatus) Enum() *Episode_VisibilityStatus {
	p := new(Episode_VisibilityStatus)
	*p = x
	return p
}

func (x Episode_VisibilityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Episode_VisibilityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_video_v1_episode_proto_enumTypes[0].Descriptor()
}

func (Episode_VisibilityStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_video_v1_episode_proto_enumTypes[0]
}

func (x Episode_VisibilityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Episode_VisibilityStatus.Descriptor instead.
func (Episode_VisibilityStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_episode_proto_rawDescGZIP(), []int{0, 0}
}

type Episode struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the episode.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the stream. Optional, empty if the episode is linked to the line
	StreamId string `protobuf:"bytes,2,opt,name=stream_id,json=streamId,proto3" json:"stream_id,omitempty"`
	// ID of the line. Optional, empty if the episode is linked to the stream
	LineId string `protobuf:"bytes,3,opt,name=line_id,json=lineId,proto3" json:"line_id,omitempty"`
	// Channel title.
	Title string `protobuf:"bytes,4,opt,name=title,proto3" json:"title,omitempty"`
	// Channel description.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the thumbnail.
	ThumbnailId string `protobuf:"bytes,6,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
	// Episode start time.
	StartTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Episode finish time.
	FinishTime *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=finish_time,json=finishTime,proto3" json:"finish_time,omitempty"`
	// Enables episode DVR mode. DVR seconds determines how many last seconds of the stream are available.
	//
	// possible values:
	//   - `0`: infinite dvr size, the full length of the stream allowed to display
	//   - `>0`: size of dvr window in seconds, the minimum value is 30s
	DvrSeconds       int64                    `protobuf:"varint,9,opt,name=dvr_seconds,json=dvrSeconds,proto3" json:"dvr_seconds,omitempty"`
	VisibilityStatus Episode_VisibilityStatus `protobuf:"varint,10,opt,name=visibility_status,json=visibilityStatus,proto3,enum=yandex.cloud.video.v1.Episode_VisibilityStatus" json:"visibility_status,omitempty"`
	// Episode access rights.
	//
	// Types that are assignable to AccessRights:
	//
	//	*Episode_PublicAccess
	//	*Episode_AuthSystemAccess
	AccessRights isEpisode_AccessRights `protobuf_oneof:"access_rights"`
	// Time when episode was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time of last episode update.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Episode) Reset() {
	*x = Episode{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_episode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Episode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Episode) ProtoMessage() {}

func (x *Episode) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_episode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Episode.ProtoReflect.Descriptor instead.
func (*Episode) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_episode_proto_rawDescGZIP(), []int{0}
}

func (x *Episode) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Episode) GetStreamId() string {
	if x != nil {
		return x.StreamId
	}
	return ""
}

func (x *Episode) GetLineId() string {
	if x != nil {
		return x.LineId
	}
	return ""
}

func (x *Episode) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Episode) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Episode) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

func (x *Episode) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Episode) GetFinishTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishTime
	}
	return nil
}

func (x *Episode) GetDvrSeconds() int64 {
	if x != nil {
		return x.DvrSeconds
	}
	return 0
}

func (x *Episode) GetVisibilityStatus() Episode_VisibilityStatus {
	if x != nil {
		return x.VisibilityStatus
	}
	return Episode_VISIBILITY_STATUS_UNSPECIFIED
}

func (m *Episode) GetAccessRights() isEpisode_AccessRights {
	if m != nil {
		return m.AccessRights
	}
	return nil
}

func (x *Episode) GetPublicAccess() *EpisodePublicAccessRights {
	if x, ok := x.GetAccessRights().(*Episode_PublicAccess); ok {
		return x.PublicAccess
	}
	return nil
}

func (x *Episode) GetAuthSystemAccess() *EpisodeAuthSystemAccessRights {
	if x, ok := x.GetAccessRights().(*Episode_AuthSystemAccess); ok {
		return x.AuthSystemAccess
	}
	return nil
}

func (x *Episode) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Episode) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type isEpisode_AccessRights interface {
	isEpisode_AccessRights()
}

type Episode_PublicAccess struct {
	// Episode is available to everyone.
	PublicAccess *EpisodePublicAccessRights `protobuf:"bytes,1000,opt,name=public_access,json=publicAccess,proto3,oneof"`
}

type Episode_AuthSystemAccess struct {
	// Checking access rights using the authorization system.
	AuthSystemAccess *EpisodeAuthSystemAccessRights `protobuf:"bytes,1002,opt,name=auth_system_access,json=authSystemAccess,proto3,oneof"`
}

func (*Episode_PublicAccess) isEpisode_AccessRights() {}

func (*Episode_AuthSystemAccess) isEpisode_AccessRights() {}

type EpisodePublicAccessRights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EpisodePublicAccessRights) Reset() {
	*x = EpisodePublicAccessRights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_episode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpisodePublicAccessRights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpisodePublicAccessRights) ProtoMessage() {}

func (x *EpisodePublicAccessRights) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_episode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpisodePublicAccessRights.ProtoReflect.Descriptor instead.
func (*EpisodePublicAccessRights) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_episode_proto_rawDescGZIP(), []int{1}
}

type EpisodeAuthSystemAccessRights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EpisodeAuthSystemAccessRights) Reset() {
	*x = EpisodeAuthSystemAccessRights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_episode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpisodeAuthSystemAccessRights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpisodeAuthSystemAccessRights) ProtoMessage() {}

func (x *EpisodeAuthSystemAccessRights) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_episode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpisodeAuthSystemAccessRights.ProtoReflect.Descriptor instead.
func (*EpisodeAuthSystemAccessRights) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_episode_proto_rawDescGZIP(), []int{2}
}

var File_yandex_cloud_video_v1_episode_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_episode_proto_rawDesc = []byte{
	0x0a, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd5, 0x06,
	0x0a, 0x07, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6c, 0x69, 0x6e, 0x65, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x76, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x76, 0x72, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x5c, 0x0a, 0x11, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x2e, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x10, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x58, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x65, 0x0a, 0x12, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x41, 0x75, 0x74, 0x68, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x48, 0x00,
	0x52, 0x10, 0x61, 0x75, 0x74, 0x68, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x55, 0x0a, 0x10, 0x56, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a, 0x1d,
	0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x55, 0x4e, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10, 0x02, 0x42,
	0x0f, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x4a, 0x05, 0x08, 0x66, 0x10, 0xe8, 0x07, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x64, 0x4a, 0x06, 0x08,
	0xe9, 0x07, 0x10, 0xea, 0x07, 0x22, 0x1b, 0x0a, 0x19, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68,
	0x74, 0x73, 0x22, 0x1f, 0x0a, 0x1d, 0x45, 0x70, 0x69, 0x73, 0x6f, 0x64, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67,
	0x68, 0x74, 0x73, 0x42, 0x5c, 0x0a, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_episode_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_episode_proto_rawDescData = file_yandex_cloud_video_v1_episode_proto_rawDesc
)

func file_yandex_cloud_video_v1_episode_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_episode_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_episode_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_episode_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_episode_proto_rawDescData
}

var file_yandex_cloud_video_v1_episode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_video_v1_episode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_video_v1_episode_proto_goTypes = []any{
	(Episode_VisibilityStatus)(0),         // 0: yandex.cloud.video.v1.Episode.VisibilityStatus
	(*Episode)(nil),                       // 1: yandex.cloud.video.v1.Episode
	(*EpisodePublicAccessRights)(nil),     // 2: yandex.cloud.video.v1.EpisodePublicAccessRights
	(*EpisodeAuthSystemAccessRights)(nil), // 3: yandex.cloud.video.v1.EpisodeAuthSystemAccessRights
	(*timestamppb.Timestamp)(nil),         // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_video_v1_episode_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.video.v1.Episode.start_time:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.video.v1.Episode.finish_time:type_name -> google.protobuf.Timestamp
	0, // 2: yandex.cloud.video.v1.Episode.visibility_status:type_name -> yandex.cloud.video.v1.Episode.VisibilityStatus
	2, // 3: yandex.cloud.video.v1.Episode.public_access:type_name -> yandex.cloud.video.v1.EpisodePublicAccessRights
	3, // 4: yandex.cloud.video.v1.Episode.auth_system_access:type_name -> yandex.cloud.video.v1.EpisodeAuthSystemAccessRights
	4, // 5: yandex.cloud.video.v1.Episode.created_at:type_name -> google.protobuf.Timestamp
	4, // 6: yandex.cloud.video.v1.Episode.updated_at:type_name -> google.protobuf.Timestamp
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_episode_proto_init() }
func file_yandex_cloud_video_v1_episode_proto_init() {
	if File_yandex_cloud_video_v1_episode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_episode_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Episode); i {
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
		file_yandex_cloud_video_v1_episode_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EpisodePublicAccessRights); i {
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
		file_yandex_cloud_video_v1_episode_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EpisodeAuthSystemAccessRights); i {
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
	file_yandex_cloud_video_v1_episode_proto_msgTypes[0].OneofWrappers = []any{
		(*Episode_PublicAccess)(nil),
		(*Episode_AuthSystemAccess)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_video_v1_episode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_video_v1_episode_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_episode_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_video_v1_episode_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_video_v1_episode_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_episode_proto = out.File
	file_yandex_cloud_video_v1_episode_proto_rawDesc = nil
	file_yandex_cloud_video_v1_episode_proto_goTypes = nil
	file_yandex_cloud_video_v1_episode_proto_depIdxs = nil
}
