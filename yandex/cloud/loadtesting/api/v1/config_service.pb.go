// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/loadtesting/api/v1/config_service.proto

package loadtesting

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/config"
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

type CreateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder to create a config in.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Config content.
	//
	// Types that are assignable to Config:
	//
	//	*CreateConfigRequest_YamlString
	Config isCreateConfigRequest_Config `protobuf_oneof:"config"`
	// Name of the config.
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateConfigRequest) Reset() {
	*x = CreateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConfigRequest) ProtoMessage() {}

func (x *CreateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateConfigRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateConfigRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (m *CreateConfigRequest) GetConfig() isCreateConfigRequest_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *CreateConfigRequest) GetYamlString() string {
	if x, ok := x.GetConfig().(*CreateConfigRequest_YamlString); ok {
		return x.YamlString
	}
	return ""
}

func (x *CreateConfigRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type isCreateConfigRequest_Config interface {
	isCreateConfigRequest_Config()
}

type CreateConfigRequest_YamlString struct {
	// Config content provided as a string in YAML format.
	YamlString string `protobuf:"bytes,2,opt,name=yaml_string,json=yamlString,proto3,oneof"`
}

func (*CreateConfigRequest_YamlString) isCreateConfigRequest_Config() {}

type CreateConfigMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the config that is being created.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *CreateConfigMetadata) Reset() {
	*x = CreateConfigMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateConfigMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateConfigMetadata) ProtoMessage() {}

func (x *CreateConfigMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateConfigMetadata.ProtoReflect.Descriptor instead.
func (*CreateConfigMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateConfigMetadata) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the config to return.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetConfigRequest) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

type ListConfigsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder to list configs in.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than `page_size`, the service returns a [ListConfigsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListConfigsResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters tests listed in the response.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListConfigsRequest) Reset() {
	*x = ListConfigsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigsRequest) ProtoMessage() {}

func (x *ListConfigsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigsRequest.ProtoReflect.Descriptor instead.
func (*ListConfigsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListConfigsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListConfigsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListConfigsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListConfigsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListConfigsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of configs in the specified folder.
	Configs []*config.Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// the specified [ListConfigsRequest.page_size], use `next_page_token` as the value
	// for the [ListConfigsRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListConfigsResponse) Reset() {
	*x = ListConfigsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListConfigsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListConfigsResponse) ProtoMessage() {}

func (x *ListConfigsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListConfigsResponse.ProtoReflect.Descriptor instead.
func (*ListConfigsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListConfigsResponse) GetConfigs() []*config.Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

func (x *ListConfigsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DeleteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the config to deleted.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *DeleteConfigRequest) Reset() {
	*x = DeleteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigRequest) ProtoMessage() {}

func (x *DeleteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteConfigRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteConfigRequest) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

type DeleteConfigMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the config that is being deleted.
	ConfigId string `protobuf:"bytes,1,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *DeleteConfigMetadata) Reset() {
	*x = DeleteConfigMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteConfigMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteConfigMetadata) ProtoMessage() {}

func (x *DeleteConfigMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteConfigMetadata.ProtoReflect.Descriptor instead.
func (*DeleteConfigMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteConfigMetadata) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

var File_yandex_cloud_loadtesting_api_v1_config_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x33,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61,
	0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0b, 0x79, 0x61, 0x6d, 0x6c, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x79,
	0x61, 0x6d, 0x6c, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x08, 0x0a,
	0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x0a, 0x22, 0x33, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x49, 0x64, 0x22, 0x91, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0x87, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x32, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x64, 0x22, 0x33, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x32, 0xbf, 0x05, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xb2, 0x01, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x4f, 0xb2, 0xd2, 0x2a, 0x25, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0d, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x20, 0x3a, 0x01, 0x2a, 0x22, 0x1b, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x12, 0x99, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x2f, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x29, 0x12, 0x27, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x96, 0x01,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x12, 0xc3, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x60, 0xb2, 0xd2, 0x2a, 0x2d,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x2a, 0x27, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x2f, 0x7b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x76, 0x0a, 0x23,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescData = file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDesc
)

func file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_loadtesting_api_v1_config_service_proto_goTypes = []interface{}{
	(*CreateConfigRequest)(nil),  // 0: yandex.cloud.loadtesting.api.v1.CreateConfigRequest
	(*CreateConfigMetadata)(nil), // 1: yandex.cloud.loadtesting.api.v1.CreateConfigMetadata
	(*GetConfigRequest)(nil),     // 2: yandex.cloud.loadtesting.api.v1.GetConfigRequest
	(*ListConfigsRequest)(nil),   // 3: yandex.cloud.loadtesting.api.v1.ListConfigsRequest
	(*ListConfigsResponse)(nil),  // 4: yandex.cloud.loadtesting.api.v1.ListConfigsResponse
	(*DeleteConfigRequest)(nil),  // 5: yandex.cloud.loadtesting.api.v1.DeleteConfigRequest
	(*DeleteConfigMetadata)(nil), // 6: yandex.cloud.loadtesting.api.v1.DeleteConfigMetadata
	(*config.Config)(nil),        // 7: yandex.cloud.loadtesting.api.v1.config.Config
	(*operation.Operation)(nil),  // 8: yandex.cloud.operation.Operation
}
var file_yandex_cloud_loadtesting_api_v1_config_service_proto_depIdxs = []int32{
	7, // 0: yandex.cloud.loadtesting.api.v1.ListConfigsResponse.configs:type_name -> yandex.cloud.loadtesting.api.v1.config.Config
	0, // 1: yandex.cloud.loadtesting.api.v1.ConfigService.Create:input_type -> yandex.cloud.loadtesting.api.v1.CreateConfigRequest
	2, // 2: yandex.cloud.loadtesting.api.v1.ConfigService.Get:input_type -> yandex.cloud.loadtesting.api.v1.GetConfigRequest
	3, // 3: yandex.cloud.loadtesting.api.v1.ConfigService.List:input_type -> yandex.cloud.loadtesting.api.v1.ListConfigsRequest
	5, // 4: yandex.cloud.loadtesting.api.v1.ConfigService.Delete:input_type -> yandex.cloud.loadtesting.api.v1.DeleteConfigRequest
	8, // 5: yandex.cloud.loadtesting.api.v1.ConfigService.Create:output_type -> yandex.cloud.operation.Operation
	7, // 6: yandex.cloud.loadtesting.api.v1.ConfigService.Get:output_type -> yandex.cloud.loadtesting.api.v1.config.Config
	4, // 7: yandex.cloud.loadtesting.api.v1.ConfigService.List:output_type -> yandex.cloud.loadtesting.api.v1.ListConfigsResponse
	8, // 8: yandex.cloud.loadtesting.api.v1.ConfigService.Delete:output_type -> yandex.cloud.operation.Operation
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_config_service_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_config_service_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_config_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConfigRequest); i {
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
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateConfigMetadata); i {
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
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigsRequest); i {
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
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListConfigsResponse); i {
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
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigRequest); i {
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
		file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteConfigMetadata); i {
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
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CreateConfigRequest_YamlString)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_config_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_config_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_api_v1_config_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_config_service_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_config_service_proto_depIdxs = nil
}
