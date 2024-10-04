// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/cic/v1/point_of_presence_service.proto

package cic

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GetPointOfPresenceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the PointOfPresence resource to return.
	// To get the pointOfPresence ID use a [PointOfPresenceService.List] request.
	PointOfPresenceId string `protobuf:"bytes,1,opt,name=point_of_presence_id,json=pointOfPresenceId,proto3" json:"point_of_presence_id,omitempty"`
}

func (x *GetPointOfPresenceRequest) Reset() {
	*x = GetPointOfPresenceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPointOfPresenceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPointOfPresenceRequest) ProtoMessage() {}

func (x *GetPointOfPresenceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPointOfPresenceRequest.ProtoReflect.Descriptor instead.
func (*GetPointOfPresenceRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPointOfPresenceRequest) GetPointOfPresenceId() string {
	if x != nil {
		return x.PointOfPresenceId
	}
	return ""
}

type ListPointOfPresencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListPointOfPresencesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListPointOfPresencesResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Subnet.name] field.
	// 2. An `=` operator.
	// 3. The value in double quotes (`"`). Must be 3-63 characters long and match the regular expression `[a-z][-a-z0-9]{1,61}[a-z0-9]`.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListPointOfPresencesRequest) Reset() {
	*x = ListPointOfPresencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPointOfPresencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPointOfPresencesRequest) ProtoMessage() {}

func (x *ListPointOfPresencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPointOfPresencesRequest.ProtoReflect.Descriptor instead.
func (*ListPointOfPresencesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPointOfPresencesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPointOfPresencesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListPointOfPresencesRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListPointOfPresencesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of PointOfPresence resources.
	PointOfPresences []*PointOfPresence `protobuf:"bytes,1,rep,name=point_of_presences,json=pointOfPresences,proto3" json:"point_of_presences,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListPointOfPresencesRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListPointOfPresencesRequest.page_token] query parameter
	// in the next list request. Subsequent list requests will have their own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPointOfPresencesResponse) Reset() {
	*x = ListPointOfPresencesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPointOfPresencesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPointOfPresencesResponse) ProtoMessage() {}

func (x *ListPointOfPresencesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPointOfPresencesResponse.ProtoReflect.Descriptor instead.
func (*ListPointOfPresencesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPointOfPresencesResponse) GetPointOfPresences() []*PointOfPresence {
	if x != nil {
		return x.PointOfPresences
	}
	return nil
}

func (x *ListPointOfPresencesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_cic_v1_point_of_presence_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDesc = []byte{
	0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x70,
	0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74,
	0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3d, 0x0a, 0x14, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x11, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x22, 0x9a, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66,
	0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a,
	0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x22, 0x9a, 0x01,
	0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65,
	0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x12, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65,
	0x52, 0x10, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xbf, 0x02, 0x0a, 0x16, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x94, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72,
	0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x37, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x31, 0x12, 0x2f, 0x2f, 0x63, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x7b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x6f, 0x66, 0x5f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x8d, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1a, 0x12, 0x18, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x4f, 0x66, 0x50, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x42, 0x56, 0x0a, 0x17,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31,
	0x3b, 0x63, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescData = file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDesc
)

func file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescData)
	})
	return file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDescData
}

var file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_cic_v1_point_of_presence_service_proto_goTypes = []any{
	(*GetPointOfPresenceRequest)(nil),    // 0: yandex.cloud.cic.v1.GetPointOfPresenceRequest
	(*ListPointOfPresencesRequest)(nil),  // 1: yandex.cloud.cic.v1.ListPointOfPresencesRequest
	(*ListPointOfPresencesResponse)(nil), // 2: yandex.cloud.cic.v1.ListPointOfPresencesResponse
	(*PointOfPresence)(nil),              // 3: yandex.cloud.cic.v1.PointOfPresence
}
var file_yandex_cloud_cic_v1_point_of_presence_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.cic.v1.ListPointOfPresencesResponse.point_of_presences:type_name -> yandex.cloud.cic.v1.PointOfPresence
	0, // 1: yandex.cloud.cic.v1.PointOfPresenceService.Get:input_type -> yandex.cloud.cic.v1.GetPointOfPresenceRequest
	1, // 2: yandex.cloud.cic.v1.PointOfPresenceService.List:input_type -> yandex.cloud.cic.v1.ListPointOfPresencesRequest
	3, // 3: yandex.cloud.cic.v1.PointOfPresenceService.Get:output_type -> yandex.cloud.cic.v1.PointOfPresence
	2, // 4: yandex.cloud.cic.v1.PointOfPresenceService.List:output_type -> yandex.cloud.cic.v1.ListPointOfPresencesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cic_v1_point_of_presence_service_proto_init() }
func file_yandex_cloud_cic_v1_point_of_presence_service_proto_init() {
	if File_yandex_cloud_cic_v1_point_of_presence_service_proto != nil {
		return
	}
	file_yandex_cloud_cic_v1_point_of_presence_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPointOfPresenceRequest); i {
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
		file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListPointOfPresencesRequest); i {
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
		file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListPointOfPresencesResponse); i {
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
			RawDescriptor: file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_cic_v1_point_of_presence_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cic_v1_point_of_presence_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cic_v1_point_of_presence_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cic_v1_point_of_presence_service_proto = out.File
	file_yandex_cloud_cic_v1_point_of_presence_service_proto_rawDesc = nil
	file_yandex_cloud_cic_v1_point_of_presence_service_proto_goTypes = nil
	file_yandex_cloud_cic_v1_point_of_presence_service_proto_depIdxs = nil
}
