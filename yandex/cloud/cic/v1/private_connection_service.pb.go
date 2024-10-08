// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/cic/v1/private_connection_service.proto

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

type GetPrivateConnectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the PrivateConnection resource to return.
	// To get the privateConnection ID use a [PrivateConnectionService.List] request.
	PrivateConnectionId string `protobuf:"bytes,1,opt,name=private_connection_id,json=privateConnectionId,proto3" json:"private_connection_id,omitempty"`
}

func (x *GetPrivateConnectionRequest) Reset() {
	*x = GetPrivateConnectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPrivateConnectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPrivateConnectionRequest) ProtoMessage() {}

func (x *GetPrivateConnectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPrivateConnectionRequest.ProtoReflect.Descriptor instead.
func (*GetPrivateConnectionRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetPrivateConnectionRequest) GetPrivateConnectionId() string {
	if x != nil {
		return x.PrivateConnectionId
	}
	return ""
}

type ListPrivateConnectionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder to list privateConnections in.
	// To get the folder ID use a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListPrivatesConnectionResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListPrivatesConnectionResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [Subnet.name] field.
	// 2. An `=` operator.
	// 3. The value in double quotes (`"`). Must be 3-63 characters long and match the regular expression `[a-z][-a-z0-9]{1,61}[a-z0-9]`.
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListPrivateConnectionsRequest) Reset() {
	*x = ListPrivateConnectionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPrivateConnectionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPrivateConnectionsRequest) ProtoMessage() {}

func (x *ListPrivateConnectionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPrivateConnectionsRequest.ProtoReflect.Descriptor instead.
func (*ListPrivateConnectionsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListPrivateConnectionsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListPrivateConnectionsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListPrivateConnectionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListPrivateConnectionsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListPrivateConnectionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of PrivateConnection resources.
	PrivateConnections []*PrivateConnection `protobuf:"bytes,1,rep,name=private_connections,json=privateConnections,proto3" json:"private_connections,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListPrivateConnectionsRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListPrivateConnectionsRequest.page_token] query parameter
	// in the next list request. Subsequent list requests will have their own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListPrivateConnectionsResponse) Reset() {
	*x = ListPrivateConnectionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPrivateConnectionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPrivateConnectionsResponse) ProtoMessage() {}

func (x *ListPrivateConnectionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPrivateConnectionsResponse.ProtoReflect.Descriptor instead.
func (*ListPrivateConnectionsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListPrivateConnectionsResponse) GetPrivateConnections() []*PrivateConnection {
	if x != nil {
		return x.PrivateConnections
	}
	return nil
}

func (x *ListPrivateConnectionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_cic_v1_private_connection_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_cic_v1_private_connection_service_proto_rawDesc = []byte{
	0x0a, 0x34, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x15, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x13, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0xc1, 0x01, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74,
	0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x09, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31,
	0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a,
	0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x31,
	0x30, 0x30, 0x30, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x1e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57,
	0x0a, 0x13, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x12, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32,
	0xce, 0x02, 0x0a, 0x18, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9b, 0x01, 0x0a,
	0x03, 0x47, 0x65, 0x74, 0x12, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x93, 0x01, 0x0a, 0x04, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x63, 0x69, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x69, 0x63, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x69,
	0x63, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescData = file_yandex_cloud_cic_v1_private_connection_service_proto_rawDesc
)

func file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescData)
	})
	return file_yandex_cloud_cic_v1_private_connection_service_proto_rawDescData
}

var file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_cic_v1_private_connection_service_proto_goTypes = []any{
	(*GetPrivateConnectionRequest)(nil),    // 0: yandex.cloud.cic.v1.GetPrivateConnectionRequest
	(*ListPrivateConnectionsRequest)(nil),  // 1: yandex.cloud.cic.v1.ListPrivateConnectionsRequest
	(*ListPrivateConnectionsResponse)(nil), // 2: yandex.cloud.cic.v1.ListPrivateConnectionsResponse
	(*PrivateConnection)(nil),              // 3: yandex.cloud.cic.v1.PrivateConnection
}
var file_yandex_cloud_cic_v1_private_connection_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.cic.v1.ListPrivateConnectionsResponse.private_connections:type_name -> yandex.cloud.cic.v1.PrivateConnection
	0, // 1: yandex.cloud.cic.v1.PrivateConnectionService.Get:input_type -> yandex.cloud.cic.v1.GetPrivateConnectionRequest
	1, // 2: yandex.cloud.cic.v1.PrivateConnectionService.List:input_type -> yandex.cloud.cic.v1.ListPrivateConnectionsRequest
	3, // 3: yandex.cloud.cic.v1.PrivateConnectionService.Get:output_type -> yandex.cloud.cic.v1.PrivateConnection
	2, // 4: yandex.cloud.cic.v1.PrivateConnectionService.List:output_type -> yandex.cloud.cic.v1.ListPrivateConnectionsResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_cic_v1_private_connection_service_proto_init() }
func file_yandex_cloud_cic_v1_private_connection_service_proto_init() {
	if File_yandex_cloud_cic_v1_private_connection_service_proto != nil {
		return
	}
	file_yandex_cloud_cic_v1_private_connection_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetPrivateConnectionRequest); i {
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
		file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListPrivateConnectionsRequest); i {
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
		file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListPrivateConnectionsResponse); i {
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
			RawDescriptor: file_yandex_cloud_cic_v1_private_connection_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_cic_v1_private_connection_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_cic_v1_private_connection_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_cic_v1_private_connection_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_cic_v1_private_connection_service_proto = out.File
	file_yandex_cloud_cic_v1_private_connection_service_proto_rawDesc = nil
	file_yandex_cloud_cic_v1_private_connection_service_proto_goTypes = nil
	file_yandex_cloud_cic_v1_private_connection_service_proto_depIdxs = nil
}
