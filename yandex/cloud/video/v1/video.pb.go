// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/video/v1/video.proto

package video

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Video_VideoStatus int32

const (
	// Video status unspecified.
	Video_VIDEO_STATUS_UNSPECIFIED Video_VideoStatus = 0
	// Waiting for the whole number of bytes to be loaded.
	Video_WAIT_UPLOADING Video_VideoStatus = 1
	// Video processing.
	Video_PROCESSING Video_VideoStatus = 4
	// Video is ready, processing is completed.
	Video_READY Video_VideoStatus = 5
	// An error occurred during video processing.
	Video_ERROR Video_VideoStatus = 7
)

// Enum value maps for Video_VideoStatus.
var (
	Video_VideoStatus_name = map[int32]string{
		0: "VIDEO_STATUS_UNSPECIFIED",
		1: "WAIT_UPLOADING",
		4: "PROCESSING",
		5: "READY",
		7: "ERROR",
	}
	Video_VideoStatus_value = map[string]int32{
		"VIDEO_STATUS_UNSPECIFIED": 0,
		"WAIT_UPLOADING":           1,
		"PROCESSING":               4,
		"READY":                    5,
		"ERROR":                    7,
	}
)

func (x Video_VideoStatus) Enum() *Video_VideoStatus {
	p := new(Video_VideoStatus)
	*p = x
	return p
}

func (x Video_VideoStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Video_VideoStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_video_v1_video_proto_enumTypes[0].Descriptor()
}

func (Video_VideoStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_video_v1_video_proto_enumTypes[0]
}

func (x Video_VideoStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Video_VideoStatus.Descriptor instead.
func (Video_VideoStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_video_proto_rawDescGZIP(), []int{0, 0}
}

type Video_VisibilityStatus int32

const (
	// Visibility status unspecified.
	Video_VISIBILITY_STATUS_UNSPECIFIED Video_VisibilityStatus = 0
	// Video is published and available for viewing.
	Video_PUBLISHED Video_VisibilityStatus = 1
	// Video is unpublished, only admin can watch.
	Video_UNPUBLISHED Video_VisibilityStatus = 2
)

// Enum value maps for Video_VisibilityStatus.
var (
	Video_VisibilityStatus_name = map[int32]string{
		0: "VISIBILITY_STATUS_UNSPECIFIED",
		1: "PUBLISHED",
		2: "UNPUBLISHED",
	}
	Video_VisibilityStatus_value = map[string]int32{
		"VISIBILITY_STATUS_UNSPECIFIED": 0,
		"PUBLISHED":                     1,
		"UNPUBLISHED":                   2,
	}
)

func (x Video_VisibilityStatus) Enum() *Video_VisibilityStatus {
	p := new(Video_VisibilityStatus)
	*p = x
	return p
}

func (x Video_VisibilityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Video_VisibilityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_video_v1_video_proto_enumTypes[1].Descriptor()
}

func (Video_VisibilityStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_video_v1_video_proto_enumTypes[1]
}

func (x Video_VisibilityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Video_VisibilityStatus.Descriptor instead.
func (Video_VisibilityStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_video_proto_rawDescGZIP(), []int{0, 1}
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the video.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the channel where the video was created.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Video title.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// Video description.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the thumbnail.
	ThumbnailId string `protobuf:"bytes,5,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
	// Video status.
	Status Video_VideoStatus `protobuf:"varint,6,opt,name=status,proto3,enum=yandex.cloud.video.v1.Video_VideoStatus" json:"status,omitempty"`
	// Video duration. Optional, may be empty until the transcoding result is ready.
	Duration *durationpb.Duration `protobuf:"bytes,8,opt,name=duration,proto3" json:"duration,omitempty"`
	// Video visibility status.
	VisibilityStatus Video_VisibilityStatus `protobuf:"varint,9,opt,name=visibility_status,json=visibilityStatus,proto3,enum=yandex.cloud.video.v1.Video_VisibilityStatus" json:"visibility_status,omitempty"`
	// Source type.
	//
	// Types that are assignable to Source:
	//
	//	*Video_Tusd
	Source isVideo_Source `protobuf_oneof:"source"`
	// Video access rights.
	//
	// Types that are assignable to AccessRights:
	//
	//	*Video_PublicAccess
	//	*Video_AuthSystemAccess
	AccessRights isVideo_AccessRights `protobuf_oneof:"access_rights"`
	// Time when video was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time of last video update.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Custom labels as “ key:value “ pairs. Maximum 64 per resource.
	Labels map[string]string `protobuf:"bytes,200,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_video_proto_rawDescGZIP(), []int{0}
}

func (x *Video) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Video) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Video) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Video) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

func (x *Video) GetStatus() Video_VideoStatus {
	if x != nil {
		return x.Status
	}
	return Video_VIDEO_STATUS_UNSPECIFIED
}

func (x *Video) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Video) GetVisibilityStatus() Video_VisibilityStatus {
	if x != nil {
		return x.VisibilityStatus
	}
	return Video_VISIBILITY_STATUS_UNSPECIFIED
}

func (m *Video) GetSource() isVideo_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *Video) GetTusd() *VideoTUSDSource {
	if x, ok := x.GetSource().(*Video_Tusd); ok {
		return x.Tusd
	}
	return nil
}

func (m *Video) GetAccessRights() isVideo_AccessRights {
	if m != nil {
		return m.AccessRights
	}
	return nil
}

func (x *Video) GetPublicAccess() *VideoPublicAccessRights {
	if x, ok := x.GetAccessRights().(*Video_PublicAccess); ok {
		return x.PublicAccess
	}
	return nil
}

func (x *Video) GetAuthSystemAccess() *VideoAuthSystemAccessRights {
	if x, ok := x.GetAccessRights().(*Video_AuthSystemAccess); ok {
		return x.AuthSystemAccess
	}
	return nil
}

func (x *Video) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Video) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Video) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type isVideo_Source interface {
	isVideo_Source()
}

type Video_Tusd struct {
	// Upload video using the tus protocol.
	Tusd *VideoTUSDSource `protobuf:"bytes,1000,opt,name=tusd,proto3,oneof"`
}

func (*Video_Tusd) isVideo_Source() {}

type isVideo_AccessRights interface {
	isVideo_AccessRights()
}

type Video_PublicAccess struct {
	// Video is available to everyone.
	PublicAccess *VideoPublicAccessRights `protobuf:"bytes,2000,opt,name=public_access,json=publicAccess,proto3,oneof"`
}

type Video_AuthSystemAccess struct {
	// Checking access rights using the authorization system.
	AuthSystemAccess *VideoAuthSystemAccessRights `protobuf:"bytes,2002,opt,name=auth_system_access,json=authSystemAccess,proto3,oneof"`
}

func (*Video_PublicAccess) isVideo_AccessRights() {}

func (*Video_AuthSystemAccess) isVideo_AccessRights() {}

type VideoTUSDSource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// URL for uploading video via the tus protocol.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *VideoTUSDSource) Reset() {
	*x = VideoTUSDSource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoTUSDSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoTUSDSource) ProtoMessage() {}

func (x *VideoTUSDSource) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoTUSDSource.ProtoReflect.Descriptor instead.
func (*VideoTUSDSource) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_video_proto_rawDescGZIP(), []int{1}
}

func (x *VideoTUSDSource) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type VideoPublicAccessRights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoPublicAccessRights) Reset() {
	*x = VideoPublicAccessRights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoPublicAccessRights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoPublicAccessRights) ProtoMessage() {}

func (x *VideoPublicAccessRights) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoPublicAccessRights.ProtoReflect.Descriptor instead.
func (*VideoPublicAccessRights) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_video_proto_rawDescGZIP(), []int{2}
}

type VideoAuthSystemAccessRights struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VideoAuthSystemAccessRights) Reset() {
	*x = VideoAuthSystemAccessRights{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoAuthSystemAccessRights) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoAuthSystemAccessRights) ProtoMessage() {}

func (x *VideoAuthSystemAccessRights) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_video_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoAuthSystemAccessRights.ProtoReflect.Descriptor instead.
func (*VideoAuthSystemAccessRights) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_video_proto_rawDescGZIP(), []int{3}
}

var File_yandex_cloud_video_v1_video_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_video_proto_rawDesc = []byte{
	0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xda, 0x08, 0x0a, 0x05,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12,
	0x40, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x56, 0x69,
	0x64, 0x65, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5a, 0x0a, 0x11, 0x76, 0x69, 0x73, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x10, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x3d, 0x0a, 0x04, 0x74, 0x75, 0x73, 0x64, 0x18, 0xe8, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x54, 0x55, 0x53, 0x44, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x48, 0x00, 0x52, 0x04, 0x74,
	0x75, 0x73, 0x64, 0x12, 0x56, 0x0a, 0x0d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x18, 0xd0, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x48, 0x01, 0x52, 0x0c, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x63, 0x0a, 0x12, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0xd2, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x48, 0x01, 0x52, 0x10,
	0x61, 0x75, 0x74, 0x68, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x64,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x41, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x56, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x22, 0x65, 0x0a, 0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x18, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x12, 0x0a, 0x0e, 0x57, 0x41, 0x49, 0x54, 0x5f, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53,
	0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x05,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x22, 0x55, 0x0a, 0x10, 0x56,
	0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x21, 0x0a, 0x1d, 0x56, 0x49, 0x53, 0x49, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45, 0x44,
	0x10, 0x02, 0x42, 0x08, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x0f, 0x0a, 0x0d,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x72, 0x69, 0x67, 0x68, 0x74, 0x73, 0x4a, 0x05, 0x08,
	0x66, 0x10, 0xc8, 0x01, 0x4a, 0x06, 0x08, 0xc9, 0x01, 0x10, 0xe8, 0x07, 0x4a, 0x06, 0x08, 0xe9,
	0x07, 0x10, 0xd0, 0x0f, 0x4a, 0x04, 0x08, 0x0a, 0x10, 0x64, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08,
	0x4a, 0x06, 0x08, 0xd1, 0x0f, 0x10, 0xd2, 0x0f, 0x22, 0x23, 0x0a, 0x0f, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x54, 0x55, 0x53, 0x44, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x19, 0x0a,
	0x17, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x1d, 0x0a, 0x1b, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x41, 0x75, 0x74, 0x68, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x69, 0x67, 0x68, 0x74, 0x73, 0x42, 0x5c, 0x0a, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_video_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_video_proto_rawDescData = file_yandex_cloud_video_v1_video_proto_rawDesc
)

func file_yandex_cloud_video_v1_video_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_video_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_video_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_video_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_video_proto_rawDescData
}

var file_yandex_cloud_video_v1_video_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_video_v1_video_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_video_v1_video_proto_goTypes = []any{
	(Video_VideoStatus)(0),              // 0: yandex.cloud.video.v1.Video.VideoStatus
	(Video_VisibilityStatus)(0),         // 1: yandex.cloud.video.v1.Video.VisibilityStatus
	(*Video)(nil),                       // 2: yandex.cloud.video.v1.Video
	(*VideoTUSDSource)(nil),             // 3: yandex.cloud.video.v1.VideoTUSDSource
	(*VideoPublicAccessRights)(nil),     // 4: yandex.cloud.video.v1.VideoPublicAccessRights
	(*VideoAuthSystemAccessRights)(nil), // 5: yandex.cloud.video.v1.VideoAuthSystemAccessRights
	nil,                                 // 6: yandex.cloud.video.v1.Video.LabelsEntry
	(*durationpb.Duration)(nil),         // 7: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),       // 8: google.protobuf.Timestamp
}
var file_yandex_cloud_video_v1_video_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.video.v1.Video.status:type_name -> yandex.cloud.video.v1.Video.VideoStatus
	7, // 1: yandex.cloud.video.v1.Video.duration:type_name -> google.protobuf.Duration
	1, // 2: yandex.cloud.video.v1.Video.visibility_status:type_name -> yandex.cloud.video.v1.Video.VisibilityStatus
	3, // 3: yandex.cloud.video.v1.Video.tusd:type_name -> yandex.cloud.video.v1.VideoTUSDSource
	4, // 4: yandex.cloud.video.v1.Video.public_access:type_name -> yandex.cloud.video.v1.VideoPublicAccessRights
	5, // 5: yandex.cloud.video.v1.Video.auth_system_access:type_name -> yandex.cloud.video.v1.VideoAuthSystemAccessRights
	8, // 6: yandex.cloud.video.v1.Video.created_at:type_name -> google.protobuf.Timestamp
	8, // 7: yandex.cloud.video.v1.Video.updated_at:type_name -> google.protobuf.Timestamp
	6, // 8: yandex.cloud.video.v1.Video.labels:type_name -> yandex.cloud.video.v1.Video.LabelsEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_video_proto_init() }
func file_yandex_cloud_video_v1_video_proto_init() {
	if File_yandex_cloud_video_v1_video_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_video_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Video); i {
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
		file_yandex_cloud_video_v1_video_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*VideoTUSDSource); i {
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
		file_yandex_cloud_video_v1_video_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*VideoPublicAccessRights); i {
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
		file_yandex_cloud_video_v1_video_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*VideoAuthSystemAccessRights); i {
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
	file_yandex_cloud_video_v1_video_proto_msgTypes[0].OneofWrappers = []any{
		(*Video_Tusd)(nil),
		(*Video_PublicAccess)(nil),
		(*Video_AuthSystemAccess)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_video_v1_video_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_video_v1_video_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_video_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_video_v1_video_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_video_v1_video_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_video_proto = out.File
	file_yandex_cloud_video_v1_video_proto_rawDesc = nil
	file_yandex_cloud_video_v1_video_proto_goTypes = nil
	file_yandex_cloud_video_v1_video_proto_depIdxs = nil
}
