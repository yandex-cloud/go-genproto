// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/baremetal/v1alpha/hardware_pool_service.proto

package baremetal

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GetHardwarePoolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the HardwarePool resource to return.
	//
	// To get the hardware pool ID, use a [HardwarePoolService.List] request.
	HardwarePoolId string `protobuf:"bytes,1,opt,name=hardware_pool_id,json=hardwarePoolId,proto3" json:"hardware_pool_id,omitempty"`
}

func (x *GetHardwarePoolRequest) Reset() {
	*x = GetHardwarePoolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHardwarePoolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHardwarePoolRequest) ProtoMessage() {}

func (x *GetHardwarePoolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHardwarePoolRequest.ProtoReflect.Descriptor instead.
func (*GetHardwarePoolRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHardwarePoolRequest) GetHardwarePoolId() string {
	if x != nil {
		return x.HardwarePoolId
	}
	return ""
}

type ListHardwarePoolsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page to return. If the number of available
	// results is greater than `page_size`,
	// the service returns a [ListHardwarePoolsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value is 20.
	PageSize int64 `protobuf:"varint,100,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListHardwarePoolsResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,101,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListHardwarePoolsRequest) Reset() {
	*x = ListHardwarePoolsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHardwarePoolsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHardwarePoolsRequest) ProtoMessage() {}

func (x *ListHardwarePoolsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHardwarePoolsRequest.ProtoReflect.Descriptor instead.
func (*ListHardwarePoolsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListHardwarePoolsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListHardwarePoolsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListHardwarePoolsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of HardwarePool resources.
	HardwarePools []*HardwarePool `protobuf:"bytes,1,rep,name=hardware_pools,json=hardwarePools,proto3" json:"hardware_pools,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// [ListHardwarePoolsResponse.page_size], use `next_page_token` as the value
	// for the [ListHardwarePoolsResponse.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListHardwarePoolsResponse) Reset() {
	*x = ListHardwarePoolsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHardwarePoolsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHardwarePoolsResponse) ProtoMessage() {}

func (x *ListHardwarePoolsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHardwarePoolsResponse.ProtoReflect.Descriptor instead.
func (*ListHardwarePoolsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListHardwarePoolsResponse) GetHardwarePools() []*HardwarePool {
	if x != nil {
		return x.HardwarePools
	}
	return nil
}

func (x *ListHardwarePoolsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x32, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x68, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x4c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x10, 0x68, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x32, 0x30, 0x52, 0x0e, 0x68,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x22, 0x67, 0x0a,
	0x18, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x42, 0x09, 0xfa, 0xc7,
	0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x4a, 0x04, 0x08, 0x01, 0x10, 0x64, 0x22, 0x9e, 0x01, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x48,
	0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65,
	0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x48, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x0d, 0x68, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x64, 0x32, 0x83, 0x02, 0x0a, 0x13, 0x48, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x61, 0x72, 0x64, 0x77,
	0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61,
	0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x22, 0x00, 0x12, 0x7d,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x72, 0x64,
	0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x39, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f,
	0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x72, 0x0a,
	0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescData = file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDesc
)

func file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescData)
	})
	return file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDescData
}

var file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_goTypes = []any{
	(*GetHardwarePoolRequest)(nil),    // 0: yandex.cloud.baremetal.v1alpha.GetHardwarePoolRequest
	(*ListHardwarePoolsRequest)(nil),  // 1: yandex.cloud.baremetal.v1alpha.ListHardwarePoolsRequest
	(*ListHardwarePoolsResponse)(nil), // 2: yandex.cloud.baremetal.v1alpha.ListHardwarePoolsResponse
	(*HardwarePool)(nil),              // 3: yandex.cloud.baremetal.v1alpha.HardwarePool
}
var file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.baremetal.v1alpha.ListHardwarePoolsResponse.hardware_pools:type_name -> yandex.cloud.baremetal.v1alpha.HardwarePool
	0, // 1: yandex.cloud.baremetal.v1alpha.HardwarePoolService.Get:input_type -> yandex.cloud.baremetal.v1alpha.GetHardwarePoolRequest
	1, // 2: yandex.cloud.baremetal.v1alpha.HardwarePoolService.List:input_type -> yandex.cloud.baremetal.v1alpha.ListHardwarePoolsRequest
	3, // 3: yandex.cloud.baremetal.v1alpha.HardwarePoolService.Get:output_type -> yandex.cloud.baremetal.v1alpha.HardwarePool
	2, // 4: yandex.cloud.baremetal.v1alpha.HardwarePoolService.List:output_type -> yandex.cloud.baremetal.v1alpha.ListHardwarePoolsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_init() }
func file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_init() {
	if File_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto != nil {
		return
	}
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetHardwarePoolRequest); i {
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
		file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListHardwarePoolsRequest); i {
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
		file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListHardwarePoolsResponse); i {
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
			RawDescriptor: file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto = out.File
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_rawDesc = nil
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_goTypes = nil
	file_yandex_cloud_baremetal_v1alpha_hardware_pool_service_proto_depIdxs = nil
}
