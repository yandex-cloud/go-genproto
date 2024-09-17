// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/video/v1/thumbnail_service.proto

package video

import (
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

type ListThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the channel.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The maximum number of the results per page to return. Default value: 100.
	PageSize int64 `protobuf:"varint,100,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token for getting the next page of the result.
	PageToken string `protobuf:"bytes,101,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListThumbnailRequest) Reset() {
	*x = ListThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListThumbnailRequest) ProtoMessage() {}

func (x *ListThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListThumbnailRequest.ProtoReflect.Descriptor instead.
func (*ListThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListThumbnailRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *ListThumbnailRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListThumbnailRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListThumbnailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of thumbnails.
	Thumbnails []*Thumbnail `protobuf:"bytes,1,rep,name=thumbnails,proto3" json:"thumbnails,omitempty"`
	// Token for getting the next page.
	NextPageToken string `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListThumbnailResponse) Reset() {
	*x = ListThumbnailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListThumbnailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListThumbnailResponse) ProtoMessage() {}

func (x *ListThumbnailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListThumbnailResponse.ProtoReflect.Descriptor instead.
func (*ListThumbnailResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListThumbnailResponse) GetThumbnails() []*Thumbnail {
	if x != nil {
		return x.Thumbnails
	}
	return nil
}

func (x *ListThumbnailResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateThumbnailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the channel.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
}

func (x *CreateThumbnailRequest) Reset() {
	*x = CreateThumbnailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateThumbnailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateThumbnailRequest) ProtoMessage() {}

func (x *CreateThumbnailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateThumbnailRequest.ProtoReflect.Descriptor instead.
func (*CreateThumbnailRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateThumbnailRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

type CreateThumbnailMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thumbnail.
	ThumbnailId string `protobuf:"bytes,1,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
}

func (x *CreateThumbnailMetadata) Reset() {
	*x = CreateThumbnailMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateThumbnailMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateThumbnailMetadata) ProtoMessage() {}

func (x *CreateThumbnailMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateThumbnailMetadata.ProtoReflect.Descriptor instead.
func (*CreateThumbnailMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateThumbnailMetadata) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

type BatchGenerateDownloadURLsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the channel.
	ChannelId string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// List of thumbnails IDs.
	ThumbnailIds []string `protobuf:"bytes,2,rep,name=thumbnail_ids,json=thumbnailIds,proto3" json:"thumbnail_ids,omitempty"`
}

func (x *BatchGenerateDownloadURLsRequest) Reset() {
	*x = BatchGenerateDownloadURLsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGenerateDownloadURLsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGenerateDownloadURLsRequest) ProtoMessage() {}

func (x *BatchGenerateDownloadURLsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGenerateDownloadURLsRequest.ProtoReflect.Descriptor instead.
func (*BatchGenerateDownloadURLsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{4}
}

func (x *BatchGenerateDownloadURLsRequest) GetChannelId() string {
	if x != nil {
		return x.ChannelId
	}
	return ""
}

func (x *BatchGenerateDownloadURLsRequest) GetThumbnailIds() []string {
	if x != nil {
		return x.ThumbnailIds
	}
	return nil
}

type BatchGenerateDownloadURLsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of download urls.
	DownloadUrls []*ThumbnailDownloadURL `protobuf:"bytes,1,rep,name=download_urls,json=downloadUrls,proto3" json:"download_urls,omitempty"`
}

func (x *BatchGenerateDownloadURLsResponse) Reset() {
	*x = BatchGenerateDownloadURLsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGenerateDownloadURLsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGenerateDownloadURLsResponse) ProtoMessage() {}

func (x *BatchGenerateDownloadURLsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGenerateDownloadURLsResponse.ProtoReflect.Descriptor instead.
func (*BatchGenerateDownloadURLsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{5}
}

func (x *BatchGenerateDownloadURLsResponse) GetDownloadUrls() []*ThumbnailDownloadURL {
	if x != nil {
		return x.DownloadUrls
	}
	return nil
}

type ThumbnailDownloadURL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thumbnail.
	ThumbnailId string `protobuf:"bytes,1,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
	// Download url.
	DownloadUrl string `protobuf:"bytes,2,opt,name=download_url,json=downloadUrl,proto3" json:"download_url,omitempty"`
}

func (x *ThumbnailDownloadURL) Reset() {
	*x = ThumbnailDownloadURL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThumbnailDownloadURL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThumbnailDownloadURL) ProtoMessage() {}

func (x *ThumbnailDownloadURL) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThumbnailDownloadURL.ProtoReflect.Descriptor instead.
func (*ThumbnailDownloadURL) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{6}
}

func (x *ThumbnailDownloadURL) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

func (x *ThumbnailDownloadURL) GetDownloadUrl() string {
	if x != nil {
		return x.DownloadUrl
	}
	return ""
}

type GenerateThumbnailUploadURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thumbnail.
	ThumbnailId string `protobuf:"bytes,1,opt,name=thumbnail_id,json=thumbnailId,proto3" json:"thumbnail_id,omitempty"`
}

func (x *GenerateThumbnailUploadURLRequest) Reset() {
	*x = GenerateThumbnailUploadURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateThumbnailUploadURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateThumbnailUploadURLRequest) ProtoMessage() {}

func (x *GenerateThumbnailUploadURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateThumbnailUploadURLRequest.ProtoReflect.Descriptor instead.
func (*GenerateThumbnailUploadURLRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{7}
}

func (x *GenerateThumbnailUploadURLRequest) GetThumbnailId() string {
	if x != nil {
		return x.ThumbnailId
	}
	return ""
}

type GenerateThumbnailUploadURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Upload url.
	UploadUrl string `protobuf:"bytes,1,opt,name=upload_url,json=uploadUrl,proto3" json:"upload_url,omitempty"`
}

func (x *GenerateThumbnailUploadURLResponse) Reset() {
	*x = GenerateThumbnailUploadURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateThumbnailUploadURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateThumbnailUploadURLResponse) ProtoMessage() {}

func (x *GenerateThumbnailUploadURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateThumbnailUploadURLResponse.ProtoReflect.Descriptor instead.
func (*GenerateThumbnailUploadURLResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP(), []int{8}
}

func (x *GenerateThumbnailUploadURLResponse) GetUploadUrl() string {
	if x != nil {
		return x.UploadUrl
	}
	return ""
}

var File_yandex_cloud_video_v1_thumbnail_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_video_v1_thumbnail_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x15, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69, 0x64,
	0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x77, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x64, 0x22, 0x87,
	0x01, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x0a, 0x74, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x0a,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x64, 0x22, 0x37, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x22, 0x3c, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c,
	0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x22,
	0x66, 0x0a, 0x20, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x73, 0x22, 0x75, 0x0a, 0x21, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64,
	0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d,
	0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x75, 0x6d,
	0x62, 0x6e, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c,
	0x52, 0x0c, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x73, 0x22, 0x5c,
	0x0a, 0x14, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x22, 0x46, 0x0a, 0x21,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x49, 0x64, 0x22, 0x43, 0x0a, 0x22, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55,
	0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x72, 0x6c, 0x32, 0xd2, 0x05, 0x0a, 0x10, 0x54, 0x68,
	0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x7f,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0xa3, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x47, 0xb2, 0xd2,
	0x2a, 0x24, 0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x09, 0x54, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x3a, 0x01, 0x2a, 0x22,
	0x14, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x73, 0x12, 0xc9, 0x01, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55,
	0x52, 0x4c, 0x73, 0x12, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61,
	0x64, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x33, 0x3a, 0x01,
	0x2a, 0x22, 0x2e, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x75,
	0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x3a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c,
	0x73, 0x12, 0xca, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x38, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x39, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x3a, 0x3a, 0x01, 0x2a, 0x22, 0x35, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x73, 0x2f, 0x7b, 0x74,
	0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x67, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x42, 0x5c,
	0x0a, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x76, 0x31, 0x5a, 0x3f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescData = file_yandex_cloud_video_v1_thumbnail_service_proto_rawDesc
)

func file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescData)
	})
	return file_yandex_cloud_video_v1_thumbnail_service_proto_rawDescData
}

var file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_video_v1_thumbnail_service_proto_goTypes = []any{
	(*ListThumbnailRequest)(nil),               // 0: yandex.cloud.video.v1.ListThumbnailRequest
	(*ListThumbnailResponse)(nil),              // 1: yandex.cloud.video.v1.ListThumbnailResponse
	(*CreateThumbnailRequest)(nil),             // 2: yandex.cloud.video.v1.CreateThumbnailRequest
	(*CreateThumbnailMetadata)(nil),            // 3: yandex.cloud.video.v1.CreateThumbnailMetadata
	(*BatchGenerateDownloadURLsRequest)(nil),   // 4: yandex.cloud.video.v1.BatchGenerateDownloadURLsRequest
	(*BatchGenerateDownloadURLsResponse)(nil),  // 5: yandex.cloud.video.v1.BatchGenerateDownloadURLsResponse
	(*ThumbnailDownloadURL)(nil),               // 6: yandex.cloud.video.v1.ThumbnailDownloadURL
	(*GenerateThumbnailUploadURLRequest)(nil),  // 7: yandex.cloud.video.v1.GenerateThumbnailUploadURLRequest
	(*GenerateThumbnailUploadURLResponse)(nil), // 8: yandex.cloud.video.v1.GenerateThumbnailUploadURLResponse
	(*Thumbnail)(nil),                          // 9: yandex.cloud.video.v1.Thumbnail
	(*operation.Operation)(nil),                // 10: yandex.cloud.operation.Operation
}
var file_yandex_cloud_video_v1_thumbnail_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.video.v1.ListThumbnailResponse.thumbnails:type_name -> yandex.cloud.video.v1.Thumbnail
	6,  // 1: yandex.cloud.video.v1.BatchGenerateDownloadURLsResponse.download_urls:type_name -> yandex.cloud.video.v1.ThumbnailDownloadURL
	0,  // 2: yandex.cloud.video.v1.ThumbnailService.List:input_type -> yandex.cloud.video.v1.ListThumbnailRequest
	2,  // 3: yandex.cloud.video.v1.ThumbnailService.Create:input_type -> yandex.cloud.video.v1.CreateThumbnailRequest
	4,  // 4: yandex.cloud.video.v1.ThumbnailService.BatchGenerateDownloadURLs:input_type -> yandex.cloud.video.v1.BatchGenerateDownloadURLsRequest
	7,  // 5: yandex.cloud.video.v1.ThumbnailService.GenerateUploadURL:input_type -> yandex.cloud.video.v1.GenerateThumbnailUploadURLRequest
	1,  // 6: yandex.cloud.video.v1.ThumbnailService.List:output_type -> yandex.cloud.video.v1.ListThumbnailResponse
	10, // 7: yandex.cloud.video.v1.ThumbnailService.Create:output_type -> yandex.cloud.operation.Operation
	5,  // 8: yandex.cloud.video.v1.ThumbnailService.BatchGenerateDownloadURLs:output_type -> yandex.cloud.video.v1.BatchGenerateDownloadURLsResponse
	8,  // 9: yandex.cloud.video.v1.ThumbnailService.GenerateUploadURL:output_type -> yandex.cloud.video.v1.GenerateThumbnailUploadURLResponse
	6,  // [6:10] is the sub-list for method output_type
	2,  // [2:6] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_video_v1_thumbnail_service_proto_init() }
func file_yandex_cloud_video_v1_thumbnail_service_proto_init() {
	if File_yandex_cloud_video_v1_thumbnail_service_proto != nil {
		return
	}
	file_yandex_cloud_video_v1_thumbnail_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListThumbnailRequest); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListThumbnailResponse); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateThumbnailRequest); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateThumbnailMetadata); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*BatchGenerateDownloadURLsRequest); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*BatchGenerateDownloadURLsResponse); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ThumbnailDownloadURL); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateThumbnailUploadURLRequest); i {
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
		file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateThumbnailUploadURLResponse); i {
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
			RawDescriptor: file_yandex_cloud_video_v1_thumbnail_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_video_v1_thumbnail_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_video_v1_thumbnail_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_video_v1_thumbnail_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_video_v1_thumbnail_service_proto = out.File
	file_yandex_cloud_video_v1_thumbnail_service_proto_rawDesc = nil
	file_yandex_cloud_video_v1_thumbnail_service_proto_goTypes = nil
	file_yandex_cloud_video_v1_thumbnail_service_proto_depIdxs = nil
}
