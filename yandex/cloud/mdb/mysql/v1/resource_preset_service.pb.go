// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/mysql/v1/resource_preset_service.proto

package mysql

import (
	context "context"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type GetResourcePresetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the resource preset to return.
	// To get the resource preset ID, use a [ResourcePresetService.List] request.
	ResourcePresetId string `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
}

func (x *GetResourcePresetRequest) Reset() {
	*x = GetResourcePresetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResourcePresetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePresetRequest) ProtoMessage() {}

func (x *GetResourcePresetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePresetRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePresetRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetResourcePresetRequest) GetResourcePresetId() string {
	if x != nil {
		return x.ResourcePresetId
	}
	return ""
}

type ListResourcePresetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListResourcePresetsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the [ListResourcePresetsResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListResourcePresetsRequest) Reset() {
	*x = ListResourcePresetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcePresetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcePresetsRequest) ProtoMessage() {}

func (x *ListResourcePresetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcePresetsRequest.ProtoReflect.Descriptor instead.
func (*ListResourcePresetsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListResourcePresetsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListResourcePresetsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListResourcePresetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of resource presets.
	ResourcePresets []*ResourcePreset `protobuf:"bytes,1,rep,name=resource_presets,json=resourcePresets,proto3" json:"resource_presets,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListResourcePresetsRequest.page_size], use the [next_page_token] as the value
	// for the [ListResourcePresetsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListResourcePresetsResponse) Reset() {
	*x = ListResourcePresetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResourcePresetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResourcePresetsResponse) ProtoMessage() {}

func (x *ListResourcePresetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResourcePresetsResponse.ProtoReflect.Descriptor instead.
func (*ListResourcePresetsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListResourcePresetsResponse) GetResourcePresets() []*ResourcePreset {
	if x != nil {
		return x.ResourcePresets
	}
	return nil
}

func (x *ListResourcePresetsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDesc = []byte{
	0x0a, 0x37, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71,
	0x6c, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01,
	0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x22, 0x6f, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x30, 0x2d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a,
	0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0xa6, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62,
	0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x12, 0x31, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xe2, 0x02, 0x0a,
	0x15, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa5, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x33,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x22, 0x3e,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x38, 0x12, 0x36, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x2d, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xa0,
	0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x2d, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x72, 0x65, 0x73, 0x65, 0x74,
	0x73, 0x42, 0x64, 0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e,
	0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76,
	0x31, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescData = file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDesc
)

func file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDescData
}

var file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_goTypes = []interface{}{
	(*GetResourcePresetRequest)(nil),    // 0: yandex.cloud.mdb.mysql.v1.GetResourcePresetRequest
	(*ListResourcePresetsRequest)(nil),  // 1: yandex.cloud.mdb.mysql.v1.ListResourcePresetsRequest
	(*ListResourcePresetsResponse)(nil), // 2: yandex.cloud.mdb.mysql.v1.ListResourcePresetsResponse
	(*ResourcePreset)(nil),              // 3: yandex.cloud.mdb.mysql.v1.ResourcePreset
}
var file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.mdb.mysql.v1.ListResourcePresetsResponse.resource_presets:type_name -> yandex.cloud.mdb.mysql.v1.ResourcePreset
	0, // 1: yandex.cloud.mdb.mysql.v1.ResourcePresetService.Get:input_type -> yandex.cloud.mdb.mysql.v1.GetResourcePresetRequest
	1, // 2: yandex.cloud.mdb.mysql.v1.ResourcePresetService.List:input_type -> yandex.cloud.mdb.mysql.v1.ListResourcePresetsRequest
	3, // 3: yandex.cloud.mdb.mysql.v1.ResourcePresetService.Get:output_type -> yandex.cloud.mdb.mysql.v1.ResourcePreset
	2, // 4: yandex.cloud.mdb.mysql.v1.ResourcePresetService.List:output_type -> yandex.cloud.mdb.mysql.v1.ListResourcePresetsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_init() }
func file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_init() {
	if File_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto != nil {
		return
	}
	file_yandex_cloud_mdb_mysql_v1_resource_preset_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResourcePresetRequest); i {
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
		file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcePresetsRequest); i {
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
		file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResourcePresetsResponse); i {
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
			RawDescriptor: file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto = out.File
	file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_rawDesc = nil
	file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_goTypes = nil
	file_yandex_cloud_mdb_mysql_v1_resource_preset_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResourcePresetServiceClient is the client API for ResourcePresetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourcePresetServiceClient interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error)
}

type resourcePresetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourcePresetServiceClient(cc grpc.ClientConnInterface) ResourcePresetServiceClient {
	return &resourcePresetServiceClient{cc}
}

func (c *resourcePresetServiceClient) Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error) {
	out := new(ResourcePreset)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcePresetServiceClient) List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error) {
	out := new(ListResourcePresetsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcePresetServiceServer is the server API for ResourcePresetService service.
type ResourcePresetServiceServer interface {
	// Returns the specified resource preset.
	//
	// To get the list of available resource presets, make a [List] request.
	Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error)
	// Retrieves the list of available resource presets.
	List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error)
}

// UnimplementedResourcePresetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedResourcePresetServiceServer struct {
}

func (*UnimplementedResourcePresetServiceServer) Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedResourcePresetServiceServer) List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterResourcePresetServiceServer(s *grpc.Server, srv ResourcePresetServiceServer) {
	s.RegisterService(&_ResourcePresetService_serviceDesc, srv)
}

func _ResourcePresetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetResourcePresetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).Get(ctx, req.(*GetResourcePresetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourcePresetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListResourcePresetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourcePresetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1.ResourcePresetService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).List(ctx, req.(*ListResourcePresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourcePresetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.mysql.v1.ResourcePresetService",
	HandlerType: (*ResourcePresetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ResourcePresetService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResourcePresetService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/mysql/v1/resource_preset_service.proto",
}
