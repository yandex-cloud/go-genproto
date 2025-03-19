// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/video/v1/stream_line.proto

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

// Enum representing the status of an automatic stream line.
// Indicates whether the automatic line is active or deactivated.
type AutoLine_AutoLineStatus int32

const (
	// Auto line status unspecified.
	AutoLine_AUTO_LINE_STATUS_UNSPECIFIED AutoLine_AutoLineStatus = 0
	// The automatic line is deactivated and not currently active.
	AutoLine_DEACTIVATED AutoLine_AutoLineStatus = 1
	// The automatic line is active and operational.
	AutoLine_ACTIVE AutoLine_AutoLineStatus = 2
)

// Enum value maps for AutoLine_AutoLineStatus.
var (
	AutoLine_AutoLineStatus_name = map[int32]string{
		0: "AUTO_LINE_STATUS_UNSPECIFIED",
		1: "DEACTIVATED",
		2: "ACTIVE",
	}
	AutoLine_AutoLineStatus_value = map[string]int32{
		"AUTO_LINE_STATUS_UNSPECIFIED": 0,
		"DEACTIVATED":                  1,
		"ACTIVE":                       2,
	}
)

func (x AutoLine_AutoLineStatus) Enum() *AutoLine_AutoLineStatus {
	p := new(AutoLine_AutoLineStatus)
	*p = x
	return p
}

func (x AutoLine_AutoLineStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AutoLine_AutoLineStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_video_v1_stream_line_proto_enumTypes[0].Descriptor()
}

func (AutoLine_AutoLineStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_video_v1_stream_line_proto_enumTypes[0]
}

func (x AutoLine_AutoLineStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AutoLine_AutoLineStatus.Descriptor instead.
func (AutoLine_AutoLineStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{5, 0}
}

// Entity that is responsible for the incoming video signal settings.
type StreamLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the line.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the channel to which this stream line belongs.
	ChannelId string `protobuf:"bytes,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// Title of the stream line.
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	// ID of the thumbnail image associated with the stream line..
	ThumbnailId string `protobuf:"bytes,4,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
	// Specifies the input type and settings for the video signal source.
	//
	// Types that are assignable to InputType:
	//
	//	*StreamLine_RtmpPush
	//	*StreamLine_RtmpPull
	InputType isStreamLine_InputType `protobuf_oneof:"input_type"`
	// Specifies the control type of the stream line.
	//
	// Types that are assignable to LineType:
	//
	//	*StreamLine_ManualLine
	//	*StreamLine_AutoLine
	LineType isStreamLine_LineType `protobuf_oneof:"line_type"`
	// Time when the stream line was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time when the stream line was last updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,101,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Custom labels as “ key:value “ pairs. Maximum 64 per resource.
	Labels map[string]string `protobuf:"bytes,200,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *StreamLine) Reset() {
	*x = StreamLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamLine) ProtoMessage() {}

func (x *StreamLine) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamLine.ProtoReflect.Descriptor instead.
func (*StreamLine) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{0}
}

func (x *StreamLine) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StreamLine) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *StreamLine) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *StreamLine) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

func (m *StreamLine) GetInputType() isStreamLine_InputType {
	if m != nil {
		return m.InputType
	}
	return nil
}

func (x *StreamLine) GetRtmpPush() *RTMPPushInput {
	if x, ok := x.GetInputType().(*StreamLine_RtmpPush); ok {
		return x.RtmpPush
	}
	return nil
}

func (x *StreamLine) GetRtmpPull() *RTMPPullInput {
	if x, ok := x.GetInputType().(*StreamLine_RtmpPull); ok {
		return x.RtmpPull
	}
	return nil
}

func (m *StreamLine) GetLineType() isStreamLine_LineType {
	if m != nil {
		return m.LineType
	}
	return nil
}

func (x *StreamLine) GetManualLine() *ManualLine {
	if x, ok := x.GetLineType().(*StreamLine_ManualLine); ok {
		return x.ManualLine
	}
	return nil
}

func (x *StreamLine) GetAutoLine() *AutoLine {
	if x, ok := x.GetLineType().(*StreamLine_AutoLine); ok {
		return x.AutoLine
	}
	return nil
}

func (x *StreamLine) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *StreamLine) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *StreamLine) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type isStreamLine_InputType interface {
	isStreamLine_InputType()
}

type StreamLine_RtmpPush struct {
	// Real-Time Messaging Protocol (RTMP) push input settings.
	RtmpPush *RTMPPushInput `protobuf:"bytes,1000,opt,name=rtmp_push,json=rtmpPush,proto3,oneof"`
}

type StreamLine_RtmpPull struct {
	// Real-Time Messaging Protocol (RTMP) pull input type.
	RtmpPull *RTMPPullInput `protobuf:"bytes,1002,opt,name=rtmp_pull,json=rtmpPull,proto3,oneof"`
}

func (*StreamLine_RtmpPush) isStreamLine_InputType() {}

func (*StreamLine_RtmpPull) isStreamLine_InputType() {}

type isStreamLine_LineType interface {
	isStreamLine_LineType()
}

type StreamLine_ManualLine struct {
	// Manual control of stream.
	ManualLine *ManualLine `protobuf:"bytes,2000,opt,name=manual_line,json=manualLine,proto3,oneof"`
}

type StreamLine_AutoLine struct {
	// Automatic control of stream.
	AutoLine *AutoLine `protobuf:"bytes,2001,opt,name=auto_line,json=autoLine,proto3,oneof"`
}

func (*StreamLine_ManualLine) isStreamLine_LineType() {}

func (*StreamLine_AutoLine) isStreamLine_LineType() {}

// Represents the stream key used for pushing video streams.
type PushStreamKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique stream key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PushStreamKey) Reset() {
	*x = PushStreamKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushStreamKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushStreamKey) ProtoMessage() {}

func (x *PushStreamKey) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushStreamKey.ProtoReflect.Descriptor instead.
func (*PushStreamKey) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{1}
}

func (x *PushStreamKey) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

// Settings for an RTMP (Real-Time Messaging Protocol) push input.
// Used when the video stream is pushed to an RTMP server.
// @see https://en.wikipedia.org/wiki/Real-Time_Messaging_Protocol
type RTMPPushInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RTMP server url.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RTMPPushInput) Reset() {
	*x = RTMPPushInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RTMPPushInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RTMPPushInput) ProtoMessage() {}

func (x *RTMPPushInput) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RTMPPushInput.ProtoReflect.Descriptor instead.
func (*RTMPPushInput) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{2}
}

func (x *RTMPPushInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Settings for an RTMP pull input.
// Used when the service pulls the video stream from an RTMP source.
// @see https://en.wikipedia.org/wiki/Real-Time_Messaging_Protocol
type RTMPPullInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// RTMP url for receiving video signal.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *RTMPPullInput) Reset() {
	*x = RTMPPullInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RTMPPullInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RTMPPullInput) ProtoMessage() {}

func (x *RTMPPullInput) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RTMPPullInput.ProtoReflect.Descriptor instead.
func (*RTMPPullInput) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{3}
}

func (x *RTMPPullInput) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// Represents a manual line type where the stream control is handled manually.
// This means that stream start/stop actions are performed by the user.
type ManualLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ManualLine) Reset() {
	*x = ManualLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManualLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManualLine) ProtoMessage() {}

func (x *ManualLine) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManualLine.ProtoReflect.Descriptor instead.
func (*ManualLine) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{4}
}

// Represents an automatic line type where the stream control is handled automatically.
type AutoLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The status of the automatic line.
	Status AutoLine_AutoLineStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.video.v1.AutoLine_AutoLineStatus" json:"status,omitempty"`
}

func (x *AutoLine) Reset() {
	*x = AutoLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoLine) ProtoMessage() {}

func (x *AutoLine) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_stream_line_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoLine.ProtoReflect.Descriptor instead.
func (*AutoLine) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP(), []int{5}
}

func (x *AutoLine) GetStatus() AutoLine_AutoLineStatus {
	if x != nil {
		return x.Status
	}
	return AutoLine_AUTO_LINE_STATUS_UNSPECIFIED
}

var File_yandex_cloud_video_v1_stream_line_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_stream_line_proto_rawDesc = []byte{
	0x0a, 0x27, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc1, 0x05, 0x0a, 0x0a, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x72, 0x74, 0x6d, 0x70,
	0x5f, 0x70, 0x75, 0x73, 0x68, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x54, 0x4d, 0x50, 0x50, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x74, 0x6d, 0x70, 0x50, 0x75, 0x73, 0x68, 0x12, 0x44,
	0x0a, 0x09, 0x72, 0x74, 0x6d, 0x70, 0x5f, 0x70, 0x75, 0x6c, 0x6c, 0x18, 0xea, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x54, 0x4d, 0x50, 0x50,
	0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x48, 0x00, 0x52, 0x08, 0x72, 0x74, 0x6d, 0x70,
	0x50, 0x75, 0x6c, 0x6c, 0x12, 0x45, 0x0a, 0x0b, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x5f, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0xd0, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x48, 0x01, 0x52,
	0x0a, 0x6d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x61,
	0x75, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0xd1, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x69, 0x6e, 0x65,
	0x48, 0x01, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x39, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x46, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0xc8, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4c, 0x69, 0x6e, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x4a, 0x04, 0x08, 0x05, 0x10, 0x64, 0x4a, 0x05, 0x08, 0x66, 0x10, 0xc8, 0x01, 0x4a, 0x06, 0x08,
	0xc9, 0x01, 0x10, 0xe8, 0x07, 0x4a, 0x06, 0x08, 0xe9, 0x07, 0x10, 0xea, 0x07, 0x4a, 0x06, 0x08,
	0xeb, 0x07, 0x10, 0xd0, 0x0f, 0x22, 0x21, 0x0a, 0x0d, 0x50, 0x75, 0x73, 0x68, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x21, 0x0a, 0x0d, 0x52, 0x54, 0x4d, 0x50,
	0x50, 0x75, 0x73, 0x68, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x21, 0x0a, 0x0d, 0x52,
	0x54, 0x4d, 0x50, 0x50, 0x75, 0x6c, 0x6c, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x0c,
	0x0a, 0x0a, 0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x4c, 0x69, 0x6e, 0x65, 0x22, 0xa3, 0x01, 0x0a,
	0x08, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x4c,
	0x69, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x4f, 0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x4c, 0x69, 0x6e, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x1c, 0x41, 0x55, 0x54, 0x4f, 0x5f, 0x4c, 0x49, 0x4e, 0x45,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x45, 0x41, 0x43, 0x54, 0x49, 0x56,
	0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45,
	0x10, 0x02, 0x42, 0x5c, 0x0a, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x5a,
	0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_stream_line_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_stream_line_proto_rawDescData = file_yandex_cloud_video_v1_stream_line_proto_rawDesc
)

func file_yandex_cloud_video_v1_stream_line_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_stream_line_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_stream_line_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_stream_line_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_stream_line_proto_rawDescData
}

var file_yandex_cloud_video_v1_stream_line_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_video_v1_stream_line_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_video_v1_stream_line_proto_goTypes = []any{
	(AutoLine_AutoLineStatus)(0),  // 0: yandex.cloud.video.v1.AutoLine.AutoLineStatus
	(*StreamLine)(nil),            // 1: yandex.cloud.video.v1.StreamLine
	(*PushStreamKey)(nil),         // 2: yandex.cloud.video.v1.PushStreamKey
	(*RTMPPushInput)(nil),         // 3: yandex.cloud.video.v1.RTMPPushInput
	(*RTMPPullInput)(nil),         // 4: yandex.cloud.video.v1.RTMPPullInput
	(*ManualLine)(nil),            // 5: yandex.cloud.video.v1.ManualLine
	(*AutoLine)(nil),              // 6: yandex.cloud.video.v1.AutoLine
	nil,                           // 7: yandex.cloud.video.v1.StreamLine.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
}
var file_yandex_cloud_video_v1_stream_line_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.video.v1.StreamLine.rtmp_push:type_name -> yandex.cloud.video.v1.RTMPPushInput
	4, // 1: yandex.cloud.video.v1.StreamLine.rtmp_pull:type_name -> yandex.cloud.video.v1.RTMPPullInput
	5, // 2: yandex.cloud.video.v1.StreamLine.manual_line:type_name -> yandex.cloud.video.v1.ManualLine
	6, // 3: yandex.cloud.video.v1.StreamLine.auto_line:type_name -> yandex.cloud.video.v1.AutoLine
	8, // 4: yandex.cloud.video.v1.StreamLine.created_at:type_name -> google.protobuf.Timestamp
	8, // 5: yandex.cloud.video.v1.StreamLine.updated_at:type_name -> google.protobuf.Timestamp
	7, // 6: yandex.cloud.video.v1.StreamLine.labels:type_name -> yandex.cloud.video.v1.StreamLine.LabelsEntry
	0, // 7: yandex.cloud.video.v1.AutoLine.status:type_name -> yandex.cloud.video.v1.AutoLine.AutoLineStatus
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_stream_line_proto_init() }
func file_yandex_cloud_video_v1_stream_line_proto_init() {
	if File_yandex_cloud_video_v1_stream_line_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_stream_line_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StreamLine); i {
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
		file_yandex_cloud_video_v1_stream_line_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PushStreamKey); i {
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
		file_yandex_cloud_video_v1_stream_line_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*RTMPPushInput); i {
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
		file_yandex_cloud_video_v1_stream_line_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*RTMPPullInput); i {
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
		file_yandex_cloud_video_v1_stream_line_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ManualLine); i {
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
		file_yandex_cloud_video_v1_stream_line_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AutoLine); i {
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
	file_yandex_cloud_video_v1_stream_line_proto_msgTypes[0].OneofWrappers = []any{
		(*StreamLine_RtmpPush)(nil),
		(*StreamLine_RtmpPull)(nil),
		(*StreamLine_ManualLine)(nil),
		(*StreamLine_AutoLine)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_video_v1_stream_line_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_video_v1_stream_line_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_stream_line_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_video_v1_stream_line_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_video_v1_stream_line_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_stream_line_proto = out.File
	file_yandex_cloud_video_v1_stream_line_proto_rawDesc = nil
	file_yandex_cloud_video_v1_stream_line_proto_goTypes = nil
	file_yandex_cloud_video_v1_stream_line_proto_depIdxs = nil
}
