// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/video/v1/subtitle_service.proto

package video

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetSubtitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subtitle.
	SubtitleId string `protobuf:"bytes,1,opt,name=subtitle_id,json=subtitleId,proto3" json:"subtitle_id,omitempty"`
}

func (x *GetSubtitleRequest) Reset() {
	*x = GetSubtitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSubtitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSubtitleRequest) ProtoMessage() {}

func (x *GetSubtitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSubtitleRequest.ProtoReflect.Descriptor instead.
func (*GetSubtitleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetSubtitleRequest) GetSubtitleId() string {
	if x != nil {
		return x.SubtitleId
	}
	return ""
}

type ListSubtitlesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of the results per page to return.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,100,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token for getting the next page of the result.
	PageToken string `protobuf:"bytes,101,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// Types that are assignable to ParentId:
	//
	//	*ListSubtitlesRequest_VideoId
	ParentId isListSubtitlesRequest_ParentId `protobuf_oneof:"parent_id"`
}

func (x *ListSubtitlesRequest) Reset() {
	*x = ListSubtitlesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubtitlesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubtitlesRequest) ProtoMessage() {}

func (x *ListSubtitlesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubtitlesRequest.ProtoReflect.Descriptor instead.
func (*ListSubtitlesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListSubtitlesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSubtitlesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (m *ListSubtitlesRequest) GetParentId() isListSubtitlesRequest_ParentId {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (x *ListSubtitlesRequest) GetVideoId() string {
	if x, ok := x.GetParentId().(*ListSubtitlesRequest_VideoId); ok {
		return x.VideoId
	}
	return ""
}

type isListSubtitlesRequest_ParentId interface {
	isListSubtitlesRequest_ParentId()
}

type ListSubtitlesRequest_VideoId struct {
	// ID of the video.
	VideoId string `protobuf:"bytes,1000,opt,name=video_id,json=videoId,proto3,oneof"`
}

func (*ListSubtitlesRequest_VideoId) isListSubtitlesRequest_ParentId() {}

type ListSubtitlesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subtitles []*Subtitle `protobuf:"bytes,1,rep,name=subtitles,proto3" json:"subtitles,omitempty"`
	// Token for getting the next page.
	NextPageToken string `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSubtitlesResponse) Reset() {
	*x = ListSubtitlesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSubtitlesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSubtitlesResponse) ProtoMessage() {}

func (x *ListSubtitlesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSubtitlesResponse.ProtoReflect.Descriptor instead.
func (*ListSubtitlesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListSubtitlesResponse) GetSubtitles() []*Subtitle {
	if x != nil {
		return x.Subtitles
	}
	return nil
}

func (x *ListSubtitlesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateSubtitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Subtitle language in any of the following formats:
	// * three-letter code according to ISO 639-2/T, ISO 639-2/B, or ISO 639-3
	// * two-letter code according to ISO 639-1
	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
	// Contains the subtitle label (or title) that will be displayed on screen during video playback.
	// Should provide a concise and accurate representation of the spoken content.
	// If not provided, it will be auto-generated based on the specified language.
	Label string `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	// Types that are assignable to ParentId:
	//
	//	*CreateSubtitleRequest_VideoId
	ParentId isCreateSubtitleRequest_ParentId `protobuf_oneof:"parent_id"`
	// Source type.
	//
	// Types that are assignable to Source:
	//
	//	*CreateSubtitleRequest_Upload
	Source isCreateSubtitleRequest_Source `protobuf_oneof:"source"`
}

func (x *CreateSubtitleRequest) Reset() {
	*x = CreateSubtitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubtitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubtitleRequest) ProtoMessage() {}

func (x *CreateSubtitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubtitleRequest.ProtoReflect.Descriptor instead.
func (*CreateSubtitleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateSubtitleRequest) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

func (x *CreateSubtitleRequest) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (m *CreateSubtitleRequest) GetParentId() isCreateSubtitleRequest_ParentId {
	if m != nil {
		return m.ParentId
	}
	return nil
}

func (x *CreateSubtitleRequest) GetVideoId() string {
	if x, ok := x.GetParentId().(*CreateSubtitleRequest_VideoId); ok {
		return x.VideoId
	}
	return ""
}

func (m *CreateSubtitleRequest) GetSource() isCreateSubtitleRequest_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (x *CreateSubtitleRequest) GetUpload() *SubtitleUploadParams {
	if x, ok := x.GetSource().(*CreateSubtitleRequest_Upload); ok {
		return x.Upload
	}
	return nil
}

type isCreateSubtitleRequest_ParentId interface {
	isCreateSubtitleRequest_ParentId()
}

type CreateSubtitleRequest_VideoId struct {
	// ID of the video.
	VideoId string `protobuf:"bytes,1000,opt,name=video_id,json=videoId,proto3,oneof"`
}

func (*CreateSubtitleRequest_VideoId) isCreateSubtitleRequest_ParentId() {}

type isCreateSubtitleRequest_Source interface {
	isCreateSubtitleRequest_Source()
}

type CreateSubtitleRequest_Upload struct {
	// Upload subtitle.
	Upload *SubtitleUploadParams `protobuf:"bytes,1100,opt,name=upload,proto3,oneof"`
}

func (*CreateSubtitleRequest_Upload) isCreateSubtitleRequest_Source() {}

type SubtitleUploadParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *SubtitleUploadParams) Reset() {
	*x = SubtitleUploadParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtitleUploadParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtitleUploadParams) ProtoMessage() {}

func (x *SubtitleUploadParams) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtitleUploadParams.ProtoReflect.Descriptor instead.
func (*SubtitleUploadParams) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{4}
}

func (x *SubtitleUploadParams) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type CreateSubtitleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subtitle.
	SubtitleId string `protobuf:"bytes,1,opt,name=subtitle_id,json=subtitleId,proto3" json:"subtitle_id,omitempty"`
}

func (x *CreateSubtitleMetadata) Reset() {
	*x = CreateSubtitleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubtitleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubtitleMetadata) ProtoMessage() {}

func (x *CreateSubtitleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubtitleMetadata.ProtoReflect.Descriptor instead.
func (*CreateSubtitleMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateSubtitleMetadata) GetSubtitleId() string {
	if x != nil {
		return x.SubtitleId
	}
	return ""
}

type GenerateSubtitleUploadURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subtitle.
	SubtitleId string `protobuf:"bytes,1,opt,name=subtitle_id,json=subtitleId,proto3" json:"subtitle_id,omitempty"`
}

func (x *GenerateSubtitleUploadURLRequest) Reset() {
	*x = GenerateSubtitleUploadURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSubtitleUploadURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSubtitleUploadURLRequest) ProtoMessage() {}

func (x *GenerateSubtitleUploadURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSubtitleUploadURLRequest.ProtoReflect.Descriptor instead.
func (*GenerateSubtitleUploadURLRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{6}
}

func (x *GenerateSubtitleUploadURLRequest) GetSubtitleId() string {
	if x != nil {
		return x.SubtitleId
	}
	return ""
}

type GenerateSubtitleUploadURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upload url.
	UploadUrl string `protobuf:"bytes,1,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
}

func (x *GenerateSubtitleUploadURLResponse) Reset() {
	*x = GenerateSubtitleUploadURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSubtitleUploadURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSubtitleUploadURLResponse) ProtoMessage() {}

func (x *GenerateSubtitleUploadURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSubtitleUploadURLResponse.ProtoReflect.Descriptor instead.
func (*GenerateSubtitleUploadURLResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateSubtitleUploadURLResponse) GetUploadUrl() string {
	if x != nil {
		return x.UploadUrl
	}
	return ""
}

type DeleteSubtitleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subtitle.
	SubtitleId string `protobuf:"bytes,1,opt,name=subtitle_id,json=subtitleId,proto3" json:"subtitle_id,omitempty"`
}

func (x *DeleteSubtitleRequest) Reset() {
	*x = DeleteSubtitleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubtitleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubtitleRequest) ProtoMessage() {}

func (x *DeleteSubtitleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubtitleRequest.ProtoReflect.Descriptor instead.
func (*DeleteSubtitleRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteSubtitleRequest) GetSubtitleId() string {
	if x != nil {
		return x.SubtitleId
	}
	return ""
}

type DeleteSubtitleMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the subtitle.
	SubtitleId string `protobuf:"bytes,1,opt,name=subtitle_id,json=subtitleId,proto3" json:"subtitle_id,omitempty"`
}

func (x *DeleteSubtitleMetadata) Reset() {
	*x = DeleteSubtitleMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSubtitleMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSubtitleMetadata) ProtoMessage() {}

func (x *DeleteSubtitleMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSubtitleMetadata.ProtoReflect.Descriptor instead.
func (*DeleteSubtitleMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteSubtitleMetadata) GetSubtitleId() string {
	if x != nil {
		return x.SubtitleId
	}
	return ""
}

var File_yandex_cloud_video_v1_subtitle_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_subtitle_service_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x43, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8,
	0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x73, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x09, 0xfa, 0xc7, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x2a, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x8a, 0xc8,
	0x31, 0x07, 0x3c, 0x3d, 0x31, 0x35, 0x30, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x48, 0x00, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x42, 0x11, 0x0a, 0x09,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x64, 0x4a, 0x05, 0x08, 0x66, 0x10, 0xe8, 0x07, 0x22, 0x84, 0x01, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x09, 0x73, 0x75, 0x62, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08,
	0x02, 0x10, 0x64, 0x22, 0xf4, 0x01, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a,
	0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x07, 0x8a, 0xc8, 0x31, 0x03, 0x32, 0x2d, 0x33, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x18, 0xe8, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x48, 0x00, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x12, 0x46, 0x0a, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0xcc, 0x08, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x48, 0x01,
	0x52, 0x06, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x11, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x42, 0x0e, 0x0a, 0x06, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x4a, 0x05, 0x08, 0x03, 0x10,
	0xe8, 0x07, 0x4a, 0x06, 0x08, 0xe9, 0x07, 0x10, 0xcc, 0x08, 0x22, 0x38, 0x0a, 0x14, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x20, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f,
	0x0a, 0x0b, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x22,
	0x51, 0x0a, 0x20, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8,
	0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x49, 0x64, 0x22, 0x42, 0x0a, 0x21, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x46, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x22, 0x39,
	0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x62, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x49, 0x64, 0x32, 0xb1, 0x06, 0x0a, 0x0f, 0x53, 0x75,
	0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7c, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73,
	0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x7e, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x44, 0xb2, 0xd2, 0x2a, 0x22, 0x0a, 0x16, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x08, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x12, 0xc3, 0x01,
	0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x55, 0x52, 0x4c, 0x12, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x35, 0x12, 0x33,
	0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x3a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x55, 0x52, 0x4c, 0x12, 0xb7, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2c,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x5c, 0xb2, 0xd2, 0x2a, 0x2f, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x75, 0x62,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x2a, 0x21, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x73, 0x2f,
	0x7b, 0x73, 0x75, 0x62, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x5c, 0x0a,
	0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x5a, 0x3f, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_subtitle_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_subtitle_service_proto_rawDescData = file_yandex_cloud_video_v1_subtitle_service_proto_rawDesc
)

func file_yandex_cloud_video_v1_subtitle_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_subtitle_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_subtitle_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_subtitle_service_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_subtitle_service_proto_rawDescData
}

var file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_yandex_cloud_video_v1_subtitle_service_proto_goTypes = []any{
	(*GetSubtitleRequest)(nil),                // 0: yandex.cloud.video.v1.GetSubtitleRequest
	(*ListSubtitlesRequest)(nil),              // 1: yandex.cloud.video.v1.ListSubtitlesRequest
	(*ListSubtitlesResponse)(nil),             // 2: yandex.cloud.video.v1.ListSubtitlesResponse
	(*CreateSubtitleRequest)(nil),             // 3: yandex.cloud.video.v1.CreateSubtitleRequest
	(*SubtitleUploadParams)(nil),              // 4: yandex.cloud.video.v1.SubtitleUploadParams
	(*CreateSubtitleMetadata)(nil),            // 5: yandex.cloud.video.v1.CreateSubtitleMetadata
	(*GenerateSubtitleUploadURLRequest)(nil),  // 6: yandex.cloud.video.v1.GenerateSubtitleUploadURLRequest
	(*GenerateSubtitleUploadURLResponse)(nil), // 7: yandex.cloud.video.v1.GenerateSubtitleUploadURLResponse
	(*DeleteSubtitleRequest)(nil),             // 8: yandex.cloud.video.v1.DeleteSubtitleRequest
	(*DeleteSubtitleMetadata)(nil),            // 9: yandex.cloud.video.v1.DeleteSubtitleMetadata
	(*Subtitle)(nil),                          // 10: yandex.cloud.video.v1.Subtitle
	(*operation.Operation)(nil),               // 11: yandex.cloud.operation.Operation
}
var file_yandex_cloud_video_v1_subtitle_service_proto_depIdxs = []int32{
	10, // 0: yandex.cloud.video.v1.ListSubtitlesResponse.subtitles:type_name -> yandex.cloud.video.v1.Subtitle
	4,  // 1: yandex.cloud.video.v1.CreateSubtitleRequest.upload:type_name -> yandex.cloud.video.v1.SubtitleUploadParams
	0,  // 2: yandex.cloud.video.v1.SubtitleService.Get:input_type -> yandex.cloud.video.v1.GetSubtitleRequest
	1,  // 3: yandex.cloud.video.v1.SubtitleService.List:input_type -> yandex.cloud.video.v1.ListSubtitlesRequest
	3,  // 4: yandex.cloud.video.v1.SubtitleService.Create:input_type -> yandex.cloud.video.v1.CreateSubtitleRequest
	6,  // 5: yandex.cloud.video.v1.SubtitleService.GenerateUploadURL:input_type -> yandex.cloud.video.v1.GenerateSubtitleUploadURLRequest
	8,  // 6: yandex.cloud.video.v1.SubtitleService.Delete:input_type -> yandex.cloud.video.v1.DeleteSubtitleRequest
	10, // 7: yandex.cloud.video.v1.SubtitleService.Get:output_type -> yandex.cloud.video.v1.Subtitle
	2,  // 8: yandex.cloud.video.v1.SubtitleService.List:output_type -> yandex.cloud.video.v1.ListSubtitlesResponse
	11, // 9: yandex.cloud.video.v1.SubtitleService.Create:output_type -> yandex.cloud.operation.Operation
	7,  // 10: yandex.cloud.video.v1.SubtitleService.GenerateUploadURL:output_type -> yandex.cloud.video.v1.GenerateSubtitleUploadURLResponse
	11, // 11: yandex.cloud.video.v1.SubtitleService.Delete:output_type -> yandex.cloud.operation.Operation
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_subtitle_service_proto_init() }
func file_yandex_cloud_video_v1_subtitle_service_proto_init() {
	if File_yandex_cloud_video_v1_subtitle_service_proto != nil {
		return
	}
	file_yandex_cloud_video_v1_subtitle_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetSubtitleRequest); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubtitlesRequest); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListSubtitlesResponse); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSubtitleRequest); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*SubtitleUploadParams); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateSubtitleMetadata); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateSubtitleUploadURLRequest); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateSubtitleUploadURLResponse); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSubtitleRequest); i {
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
		file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteSubtitleMetadata); i {
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
	file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[1].OneofWrappers = []any{
		(*ListSubtitlesRequest_VideoId)(nil),
	}
	file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes[3].OneofWrappers = []any{
		(*CreateSubtitleRequest_VideoId)(nil),
		(*CreateSubtitleRequest_Upload)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_video_v1_subtitle_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_video_v1_subtitle_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_subtitle_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_video_v1_subtitle_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_subtitle_service_proto = out.File
	file_yandex_cloud_video_v1_subtitle_service_proto_rawDesc = nil
	file_yandex_cloud_video_v1_subtitle_service_proto_goTypes = nil
	file_yandex_cloud_video_v1_subtitle_service_proto_depIdxs = nil
}
