// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/cdn/v1/origin_service.proto

package cdn

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder that the origin belongs to.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// [origin_id] group ID to request origin from.
	OriginId int64 `protobuf:"varint,2,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
}

func (x *GetOriginRequest) Reset() {
	*x = GetOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOriginRequest) ProtoMessage() {}

func (x *GetOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOriginRequest.ProtoReflect.Descriptor instead.
func (*GetOriginRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetOriginRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *GetOriginRequest) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

type ListOriginsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder that the origin belongs to.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// ID of the group to request origins from.
	OriginGroupId int64 `protobuf:"varint,2,opt,name=origin_group_id,json=originGroupId,proto3" json:"origin_group_id,omitempty"`
}

func (x *ListOriginsRequest) Reset() {
	*x = ListOriginsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOriginsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOriginsRequest) ProtoMessage() {}

func (x *ListOriginsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOriginsRequest.ProtoReflect.Descriptor instead.
func (*ListOriginsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListOriginsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListOriginsRequest) GetOriginGroupId() int64 {
	if x != nil {
		return x.OriginGroupId
	}
	return 0
}

type ListOriginsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Origin from response.
	Origins []*Origin `protobuf:"bytes,1,rep,name=origins,proto3" json:"origins,omitempty"`
}

func (x *ListOriginsResponse) Reset() {
	*x = ListOriginsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOriginsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOriginsResponse) ProtoMessage() {}

func (x *ListOriginsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOriginsResponse.ProtoReflect.Descriptor instead.
func (*ListOriginsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListOriginsResponse) GetOrigins() []*Origin {
	if x != nil {
		return x.Origins
	}
	return nil
}

type CreateOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder that the origin belongs to.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// [origin_group_id] group ID to request origins from.
	OriginGroupId int64 `protobuf:"varint,2,opt,name=origin_group_id,json=originGroupId,proto3" json:"origin_group_id,omitempty"`
	// IP address or Domain name of your origin and the port (if custom).
	// Used if [meta] variant is `common`.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// The setting allows to enable or disable an Origin source in the Origins group.
	//
	// It has two possible values:
	//
	// True - The origin is enabled and used as a source for the CDN. An origins
	// group must contain at least one enabled origin. Default value.
	// False - The origin is disabled and the CDN is not using it to pull content.
	Enabled *wrapperspb.BoolValue `protobuf:"bytes,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Specifies whether the origin is used in its origin group as backup.
	// A backup origin is used when one of active origins becomes unavailable.
	//
	// Default value: False.
	Backup *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=backup,proto3" json:"backup,omitempty"`
	// Set up origin of the content.
	Meta *OriginMeta `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *CreateOriginRequest) Reset() {
	*x = CreateOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOriginRequest) ProtoMessage() {}

func (x *CreateOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOriginRequest.ProtoReflect.Descriptor instead.
func (*CreateOriginRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateOriginRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateOriginRequest) GetOriginGroupId() int64 {
	if x != nil {
		return x.OriginGroupId
	}
	return 0
}

func (x *CreateOriginRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreateOriginRequest) GetEnabled() *wrapperspb.BoolValue {
	if x != nil {
		return x.Enabled
	}
	return nil
}

func (x *CreateOriginRequest) GetBackup() *wrapperspb.BoolValue {
	if x != nil {
		return x.Backup
	}
	return nil
}

func (x *CreateOriginRequest) GetMeta() *OriginMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type CreateOriginMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the origin.
	OriginId int64 `protobuf:"varint,1,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
	// ID pf the parent origins group.
	OriginGroupId int64 `protobuf:"varint,2,opt,name=origin_group_id,json=originGroupId,proto3" json:"origin_group_id,omitempty"`
}

func (x *CreateOriginMetadata) Reset() {
	*x = CreateOriginMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOriginMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOriginMetadata) ProtoMessage() {}

func (x *CreateOriginMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOriginMetadata.ProtoReflect.Descriptor instead.
func (*CreateOriginMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateOriginMetadata) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

func (x *CreateOriginMetadata) GetOriginGroupId() int64 {
	if x != nil {
		return x.OriginGroupId
	}
	return 0
}

type UpdateOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder that the origin belongs to.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// ID of the origin.
	OriginId int64 `protobuf:"varint,2,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
	// IP address or Domain name of your origin and the port (if custom).
	// Used if [meta] variant is `common`.
	// Required.
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	// The setting allows to enable or disable an Origin source in the Origins group.
	//
	// It has two possible values:
	//
	// True - The origin is enabled and used as a source for the CDN. An origins
	// group must contain at least one enabled origin. Default value.
	// False - The origin is disabled and the CDN is not using it to pull content.
	//
	// Required.
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Specifies whether the origin is used in its origin group as backup.
	// A backup origin is used when one of active origins becomes unavailable.
	//
	// Required.
	Backup bool `protobuf:"varint,5,opt,name=backup,proto3" json:"backup,omitempty"`
	// Set up type of the origin.
	Meta *OriginMeta `protobuf:"bytes,6,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *UpdateOriginRequest) Reset() {
	*x = UpdateOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOriginRequest) ProtoMessage() {}

func (x *UpdateOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOriginRequest.ProtoReflect.Descriptor instead.
func (*UpdateOriginRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateOriginRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *UpdateOriginRequest) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

func (x *UpdateOriginRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *UpdateOriginRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *UpdateOriginRequest) GetBackup() bool {
	if x != nil {
		return x.Backup
	}
	return false
}

func (x *UpdateOriginRequest) GetMeta() *OriginMeta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type UpdateOriginMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the origin.
	OriginId int64 `protobuf:"varint,1,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
	// Parent origins group ID.
	OriginGroupId int64 `protobuf:"varint,2,opt,name=origin_group_id,json=originGroupId,proto3" json:"origin_group_id,omitempty"`
}

func (x *UpdateOriginMetadata) Reset() {
	*x = UpdateOriginMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOriginMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOriginMetadata) ProtoMessage() {}

func (x *UpdateOriginMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOriginMetadata.ProtoReflect.Descriptor instead.
func (*UpdateOriginMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateOriginMetadata) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

func (x *UpdateOriginMetadata) GetOriginGroupId() int64 {
	if x != nil {
		return x.OriginGroupId
	}
	return 0
}

type DeleteOriginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder that the origin belongs to.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// ID of the origin.
	OriginId int64 `protobuf:"varint,2,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
}

func (x *DeleteOriginRequest) Reset() {
	*x = DeleteOriginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOriginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOriginRequest) ProtoMessage() {}

func (x *DeleteOriginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOriginRequest.ProtoReflect.Descriptor instead.
func (*DeleteOriginRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteOriginRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *DeleteOriginRequest) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

type DeleteOriginMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the origin.
	OriginId int64 `protobuf:"varint,1,opt,name=origin_id,json=originId,proto3" json:"origin_id,omitempty"`
}

func (x *DeleteOriginMetadata) Reset() {
	*x = DeleteOriginMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOriginMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOriginMetadata) ProtoMessage() {}

func (x *DeleteOriginMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOriginMetadata.ProtoReflect.Descriptor instead.
func (*DeleteOriginMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteOriginMetadata) GetOriginId() int64 {
	if x != nil {
		return x.OriginId
	}
	return 0
}

var File_yandex_cloud_cdn_v1_origin_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_cdn_v1_origin_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64,
	0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09,
	0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02,
	0x3e, 0x30, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x6f, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x0d,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x22, 0x4c, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x52, 0x07, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x22, 0xb5, 0x02, 0x0a, 0x13,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2e,
	0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52,
	0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x24,
	0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x06, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x32, 0x0a, 0x06, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f,
	0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x12, 0x33,
	0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d,
	0x65, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x09, 0x6f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06,
	0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x2e, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e,
	0x30, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64,
	0x22, 0xe4, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x08,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74,
	0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x23, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x0f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa,
	0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x0d, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x64, 0x22, 0x65, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e,
	0x30, 0x52, 0x08, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x14, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x09, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x08,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x49, 0x64, 0x32, 0xd9, 0x05, 0x0a, 0x0d, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x69, 0x67, 0x69, 0x6e, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f,
	0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x7b,
	0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x72, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12, 0x0f, 0x2f,
	0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x12, 0x93,
	0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0xb2, 0xd2, 0x2a, 0x1e, 0x0a, 0x14, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14,
	0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x48, 0xb2, 0xd2,
	0x2a, 0x1e, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x3a, 0x01, 0x2a, 0x32, 0x1b, 0x2f, 0x63, 0x64, 0x6e, 0x2f,
	0x76, 0x31, 0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xab, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54,
	0xb2, 0xd2, 0x2a, 0x2d, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x2a, 0x1b, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31,
	0x2f, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x73, 0x2f, 0x7b, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x7d, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x5a,
	0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x64, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cdn_v1_origin_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_origin_service_proto_rawDescData = file_yandex_cloud_cdn_v1_origin_service_proto_rawDesc
)

func file_yandex_cloud_cdn_v1_origin_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_origin_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_origin_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cdn_v1_origin_service_proto_rawDescData)
	})
	return file_yandex_cloud_cdn_v1_origin_service_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_cdn_v1_origin_service_proto_goTypes = []any{
	(*GetOriginRequest)(nil),     // 0: yandex.cloud.cdn.v1.GetOriginRequest
	(*ListOriginsRequest)(nil),   // 1: yandex.cloud.cdn.v1.ListOriginsRequest
	(*ListOriginsResponse)(nil),  // 2: yandex.cloud.cdn.v1.ListOriginsResponse
	(*CreateOriginRequest)(nil),  // 3: yandex.cloud.cdn.v1.CreateOriginRequest
	(*CreateOriginMetadata)(nil), // 4: yandex.cloud.cdn.v1.CreateOriginMetadata
	(*UpdateOriginRequest)(nil),  // 5: yandex.cloud.cdn.v1.UpdateOriginRequest
	(*UpdateOriginMetadata)(nil), // 6: yandex.cloud.cdn.v1.UpdateOriginMetadata
	(*DeleteOriginRequest)(nil),  // 7: yandex.cloud.cdn.v1.DeleteOriginRequest
	(*DeleteOriginMetadata)(nil), // 8: yandex.cloud.cdn.v1.DeleteOriginMetadata
	(*Origin)(nil),               // 9: yandex.cloud.cdn.v1.Origin
	(*wrapperspb.BoolValue)(nil), // 10: google.protobuf.BoolValue
	(*OriginMeta)(nil),           // 11: yandex.cloud.cdn.v1.OriginMeta
	(*operation.Operation)(nil),  // 12: yandex.cloud.operation.Operation
}
var file_yandex_cloud_cdn_v1_origin_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.cdn.v1.ListOriginsResponse.origins:type_name -> yandex.cloud.cdn.v1.Origin
	10, // 1: yandex.cloud.cdn.v1.CreateOriginRequest.enabled:type_name -> google.protobuf.BoolValue
	10, // 2: yandex.cloud.cdn.v1.CreateOriginRequest.backup:type_name -> google.protobuf.BoolValue
	11, // 3: yandex.cloud.cdn.v1.CreateOriginRequest.meta:type_name -> yandex.cloud.cdn.v1.OriginMeta
	11, // 4: yandex.cloud.cdn.v1.UpdateOriginRequest.meta:type_name -> yandex.cloud.cdn.v1.OriginMeta
	0,  // 5: yandex.cloud.cdn.v1.OriginService.Get:input_type -> yandex.cloud.cdn.v1.GetOriginRequest
	1,  // 6: yandex.cloud.cdn.v1.OriginService.List:input_type -> yandex.cloud.cdn.v1.ListOriginsRequest
	3,  // 7: yandex.cloud.cdn.v1.OriginService.Create:input_type -> yandex.cloud.cdn.v1.CreateOriginRequest
	5,  // 8: yandex.cloud.cdn.v1.OriginService.Update:input_type -> yandex.cloud.cdn.v1.UpdateOriginRequest
	7,  // 9: yandex.cloud.cdn.v1.OriginService.Delete:input_type -> yandex.cloud.cdn.v1.DeleteOriginRequest
	9,  // 10: yandex.cloud.cdn.v1.OriginService.Get:output_type -> yandex.cloud.cdn.v1.Origin
	2,  // 11: yandex.cloud.cdn.v1.OriginService.List:output_type -> yandex.cloud.cdn.v1.ListOriginsResponse
	12, // 12: yandex.cloud.cdn.v1.OriginService.Create:output_type -> yandex.cloud.operation.Operation
	12, // 13: yandex.cloud.cdn.v1.OriginService.Update:output_type -> yandex.cloud.operation.Operation
	12, // 14: yandex.cloud.cdn.v1.OriginService.Delete:output_type -> yandex.cloud.operation.Operation
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_origin_service_proto_init() }
func file_yandex_cloud_cdn_v1_origin_service_proto_init() {
	if File_yandex_cloud_cdn_v1_origin_service_proto != nil {
		return
	}
	file_yandex_cloud_cdn_v1_origin_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetOriginRequest); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListOriginsRequest); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListOriginsResponse); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOriginRequest); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOriginMetadata); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateOriginRequest); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateOriginMetadata); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteOriginRequest); i {
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
		file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteOriginMetadata); i {
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
			RawDescriptor: file_yandex_cloud_cdn_v1_origin_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_origin_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_origin_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_origin_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_origin_service_proto = out.File
	file_yandex_cloud_cdn_v1_origin_service_proto_rawDesc = nil
	file_yandex_cloud_cdn_v1_origin_service_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_origin_service_proto_depIdxs = nil
}
