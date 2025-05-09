// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/cdn/v1/shielding_service.proto

package cdn

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
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

// Request to retrieve shielding details for a specific resource.
type GetShieldingDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which to get shielding details.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *GetShieldingDetailsRequest) Reset() {
	*x = GetShieldingDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShieldingDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShieldingDetailsRequest) ProtoMessage() {}

func (x *GetShieldingDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShieldingDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetShieldingDetailsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetShieldingDetailsRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Request to activate shielding for a specific resource.
type ActivateShieldingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource to activate shielding for.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of the location to activate shielding in, allowing selection of a suitable geographical location.
	LocationId int64 `protobuf:"varint,2,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *ActivateShieldingRequest) Reset() {
	*x = ActivateShieldingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateShieldingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateShieldingRequest) ProtoMessage() {}

func (x *ActivateShieldingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateShieldingRequest.ProtoReflect.Descriptor instead.
func (*ActivateShieldingRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{1}
}

func (x *ActivateShieldingRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *ActivateShieldingRequest) GetLocationId() int64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

// Metadata for shielding activation, detailing the operations performed.
type ActivateShieldingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which shielding is being activated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *ActivateShieldingMetadata) Reset() {
	*x = ActivateShieldingMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateShieldingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateShieldingMetadata) ProtoMessage() {}

func (x *ActivateShieldingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateShieldingMetadata.ProtoReflect.Descriptor instead.
func (*ActivateShieldingMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{2}
}

func (x *ActivateShieldingMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Request to deactivate shielding for a specific resource.
type DeactivateShieldingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource to deactivate shielding for.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeactivateShieldingRequest) Reset() {
	*x = DeactivateShieldingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateShieldingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateShieldingRequest) ProtoMessage() {}

func (x *DeactivateShieldingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateShieldingRequest.ProtoReflect.Descriptor instead.
func (*DeactivateShieldingRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeactivateShieldingRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Metadata for shielding deactivation, detailing the operations performed.
type DeactivateShieldingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which shielding is being deactivated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeactivateShieldingMetadata) Reset() {
	*x = DeactivateShieldingMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeactivateShieldingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeactivateShieldingMetadata) ProtoMessage() {}

func (x *DeactivateShieldingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeactivateShieldingMetadata.ProtoReflect.Descriptor instead.
func (*DeactivateShieldingMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeactivateShieldingMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Request to update shielding parameters, including location adjustments.
type UpdateShieldingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which shielding parameters are being updated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of the location for updating shielding parameters, allowing for geographical adjustments.
	LocationId int64 `protobuf:"varint,2,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *UpdateShieldingRequest) Reset() {
	*x = UpdateShieldingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShieldingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShieldingRequest) ProtoMessage() {}

func (x *UpdateShieldingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShieldingRequest.ProtoReflect.Descriptor instead.
func (*UpdateShieldingRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateShieldingRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *UpdateShieldingRequest) GetLocationId() int64 {
	if x != nil {
		return x.LocationId
	}
	return 0
}

// Metadata for shielding updates, detailing the operations performed.
type UpdateShieldingMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource for which shielding parameters are being updated.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *UpdateShieldingMetadata) Reset() {
	*x = UpdateShieldingMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateShieldingMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateShieldingMetadata) ProtoMessage() {}

func (x *UpdateShieldingMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateShieldingMetadata.ProtoReflect.Descriptor instead.
func (*UpdateShieldingMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateShieldingMetadata) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

// Request to list available geographical locations for shielding.
type ListShieldingLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder for which to request a list of locations where shielding can be activated.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Maximum number of results per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results,
	// set [page_token] to the [ListShieldingLocationsResponse.next_page_token]
	// returned by a previous list response.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListShieldingLocationsRequest) Reset() {
	*x = ListShieldingLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShieldingLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShieldingLocationsRequest) ProtoMessage() {}

func (x *ListShieldingLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShieldingLocationsRequest.ProtoReflect.Descriptor instead.
func (*ListShieldingLocationsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListShieldingLocationsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListShieldingLocationsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListShieldingLocationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response for the list of available shielding locations.
type ListShieldingLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of available shielding locations, each representing a potential geographical location for the shielding server.
	ShieldingLocations []*ShieldingDetails `protobuf:"bytes,1,rep,name=shielding_locations,json=shieldingLocations,proto3" json:"shielding_locations,omitempty"`
	// Token for getting the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListShieldingLocationsResponse) Reset() {
	*x = ListShieldingLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListShieldingLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListShieldingLocationsResponse) ProtoMessage() {}

func (x *ListShieldingLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListShieldingLocationsResponse.ProtoReflect.Descriptor instead.
func (*ListShieldingLocationsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListShieldingLocationsResponse) GetShieldingLocations() []*ShieldingDetails {
	if x != nil {
		return x.ShieldingLocations
	}
	return nil
}

func (x *ListShieldingLocationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_cdn_v1_shielding_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_cdn_v1_shielding_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e,
	0x76, 0x31, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x4b, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x72, 0x0a,
	0x18, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa,
	0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x22, 0x4a, 0x0a, 0x19, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d,
	0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4b, 0x0a,
	0x1a, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x4c, 0x0a, 0x1b, 0x44, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x70, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x06, 0xfa, 0xc7, 0x31, 0x02, 0x3e, 0x30, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x17, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01,
	0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x64, 0x22, 0x9d, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a,
	0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09,
	0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xa0, 0x01, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x56, 0x0a, 0x13, 0x73, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x12, 0x73, 0x68, 0x69,
	0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xb0, 0x05, 0x0a, 0x10, 0x53, 0x68, 0x69, 0x65,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8f, 0x01, 0x0a,
	0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x31, 0xb2, 0xd2, 0x2a,
	0x2d, 0x0a, 0x19, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x53, 0x68,
	0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x9a,
	0x01, 0x0a, 0x0a, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x68,
	0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x38, 0xb2, 0xd2, 0x2a, 0x34, 0x0a, 0x1b, 0x44, 0x65, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5d, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x69, 0x65,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x89, 0x01, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x2f, 0xb2, 0xd2, 0x2a, 0x2b, 0x0a, 0x17, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x10, 0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x81, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x68, 0x69, 0x65,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x64, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x68, 0x69, 0x65, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x64, 0x6e, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x64, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x64, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescData = file_yandex_cloud_cdn_v1_shielding_service_proto_rawDesc
)

func file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescData)
	})
	return file_yandex_cloud_cdn_v1_shielding_service_proto_rawDescData
}

var file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_cdn_v1_shielding_service_proto_goTypes = []any{
	(*GetShieldingDetailsRequest)(nil),     // 0: yandex.cloud.cdn.v1.GetShieldingDetailsRequest
	(*ActivateShieldingRequest)(nil),       // 1: yandex.cloud.cdn.v1.ActivateShieldingRequest
	(*ActivateShieldingMetadata)(nil),      // 2: yandex.cloud.cdn.v1.ActivateShieldingMetadata
	(*DeactivateShieldingRequest)(nil),     // 3: yandex.cloud.cdn.v1.DeactivateShieldingRequest
	(*DeactivateShieldingMetadata)(nil),    // 4: yandex.cloud.cdn.v1.DeactivateShieldingMetadata
	(*UpdateShieldingRequest)(nil),         // 5: yandex.cloud.cdn.v1.UpdateShieldingRequest
	(*UpdateShieldingMetadata)(nil),        // 6: yandex.cloud.cdn.v1.UpdateShieldingMetadata
	(*ListShieldingLocationsRequest)(nil),  // 7: yandex.cloud.cdn.v1.ListShieldingLocationsRequest
	(*ListShieldingLocationsResponse)(nil), // 8: yandex.cloud.cdn.v1.ListShieldingLocationsResponse
	(*ShieldingDetails)(nil),               // 9: yandex.cloud.cdn.v1.ShieldingDetails
	(*operation.Operation)(nil),            // 10: yandex.cloud.operation.Operation
}
var file_yandex_cloud_cdn_v1_shielding_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.cdn.v1.ListShieldingLocationsResponse.shielding_locations:type_name -> yandex.cloud.cdn.v1.ShieldingDetails
	1,  // 1: yandex.cloud.cdn.v1.ShieldingService.Activate:input_type -> yandex.cloud.cdn.v1.ActivateShieldingRequest
	3,  // 2: yandex.cloud.cdn.v1.ShieldingService.Deactivate:input_type -> yandex.cloud.cdn.v1.DeactivateShieldingRequest
	0,  // 3: yandex.cloud.cdn.v1.ShieldingService.Get:input_type -> yandex.cloud.cdn.v1.GetShieldingDetailsRequest
	5,  // 4: yandex.cloud.cdn.v1.ShieldingService.Update:input_type -> yandex.cloud.cdn.v1.UpdateShieldingRequest
	7,  // 5: yandex.cloud.cdn.v1.ShieldingService.ListAvailableLocations:input_type -> yandex.cloud.cdn.v1.ListShieldingLocationsRequest
	10, // 6: yandex.cloud.cdn.v1.ShieldingService.Activate:output_type -> yandex.cloud.operation.Operation
	10, // 7: yandex.cloud.cdn.v1.ShieldingService.Deactivate:output_type -> yandex.cloud.operation.Operation
	9,  // 8: yandex.cloud.cdn.v1.ShieldingService.Get:output_type -> yandex.cloud.cdn.v1.ShieldingDetails
	10, // 9: yandex.cloud.cdn.v1.ShieldingService.Update:output_type -> yandex.cloud.operation.Operation
	8,  // 10: yandex.cloud.cdn.v1.ShieldingService.ListAvailableLocations:output_type -> yandex.cloud.cdn.v1.ListShieldingLocationsResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cdn_v1_shielding_service_proto_init() }
func file_yandex_cloud_cdn_v1_shielding_service_proto_init() {
	if File_yandex_cloud_cdn_v1_shielding_service_proto != nil {
		return
	}
	file_yandex_cloud_cdn_v1_shielding_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetShieldingDetailsRequest); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ActivateShieldingRequest); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ActivateShieldingMetadata); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*DeactivateShieldingRequest); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeactivateShieldingMetadata); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateShieldingRequest); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateShieldingMetadata); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListShieldingLocationsRequest); i {
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
		file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListShieldingLocationsResponse); i {
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
			RawDescriptor: file_yandex_cloud_cdn_v1_shielding_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_cdn_v1_shielding_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cdn_v1_shielding_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cdn_v1_shielding_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cdn_v1_shielding_service_proto = out.File
	file_yandex_cloud_cdn_v1_shielding_service_proto_rawDesc = nil
	file_yandex_cloud_cdn_v1_shielding_service_proto_goTypes = nil
	file_yandex_cloud_cdn_v1_shielding_service_proto_depIdxs = nil
}
