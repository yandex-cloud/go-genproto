// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/resourcemanager/v1/cloud_service.proto

package resourcemanager

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetCloudRequest struct {
	// ID of the Cloud resource to return.
	// To get the cloud ID, use a [CloudService.List] request.
	CloudId              string   `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCloudRequest) Reset()         { *m = GetCloudRequest{} }
func (m *GetCloudRequest) String() string { return proto.CompactTextString(m) }
func (*GetCloudRequest) ProtoMessage()    {}
func (*GetCloudRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{0}
}

func (m *GetCloudRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCloudRequest.Unmarshal(m, b)
}
func (m *GetCloudRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCloudRequest.Marshal(b, m, deterministic)
}
func (m *GetCloudRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCloudRequest.Merge(m, src)
}
func (m *GetCloudRequest) XXX_Size() int {
	return xxx_messageInfo_GetCloudRequest.Size(m)
}
func (m *GetCloudRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCloudRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCloudRequest proto.InternalMessageInfo

func (m *GetCloudRequest) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

type ListCloudsRequest struct {
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListCloudsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set [page_token]
	// to the [ListCloudsResponse.next_page_token]
	// returned by a previous list request to get the next page of results.
	PageToken string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on the [Cloud.name] field.
	// 2. An operator. Can be either `=` or `!=` for single values, `IN` or `NOT IN` for lists of values.
	// 3. The value. Must be 3-63 characters long and match the regular expression `^[a-z][-a-z0-9]{1,61}[a-z0-9]$`.
	Filter               string   `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCloudsRequest) Reset()         { *m = ListCloudsRequest{} }
func (m *ListCloudsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCloudsRequest) ProtoMessage()    {}
func (*ListCloudsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{1}
}

func (m *ListCloudsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudsRequest.Unmarshal(m, b)
}
func (m *ListCloudsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudsRequest.Marshal(b, m, deterministic)
}
func (m *ListCloudsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudsRequest.Merge(m, src)
}
func (m *ListCloudsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCloudsRequest.Size(m)
}
func (m *ListCloudsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudsRequest proto.InternalMessageInfo

func (m *ListCloudsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCloudsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListCloudsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type ListCloudsResponse struct {
	// List of Cloud resources.
	Clouds []*Cloud `protobuf:"bytes,1,rep,name=clouds,proto3" json:"clouds,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListCloudsRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListCloudsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCloudsResponse) Reset()         { *m = ListCloudsResponse{} }
func (m *ListCloudsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCloudsResponse) ProtoMessage()    {}
func (*ListCloudsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{2}
}

func (m *ListCloudsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudsResponse.Unmarshal(m, b)
}
func (m *ListCloudsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudsResponse.Marshal(b, m, deterministic)
}
func (m *ListCloudsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudsResponse.Merge(m, src)
}
func (m *ListCloudsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCloudsResponse.Size(m)
}
func (m *ListCloudsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudsResponse proto.InternalMessageInfo

func (m *ListCloudsResponse) GetClouds() []*Cloud {
	if m != nil {
		return m.Clouds
	}
	return nil
}

func (m *ListCloudsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type ListCloudOperationsRequest struct {
	// ID of the Cloud resource to list operations for.
	CloudId string `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListCloudOperationsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set [page_token]
	// to the [ListCloudOperationsResponse.next_page_token]
	// returned by a previous list request to get the next page of results.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCloudOperationsRequest) Reset()         { *m = ListCloudOperationsRequest{} }
func (m *ListCloudOperationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListCloudOperationsRequest) ProtoMessage()    {}
func (*ListCloudOperationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{3}
}

func (m *ListCloudOperationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudOperationsRequest.Unmarshal(m, b)
}
func (m *ListCloudOperationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudOperationsRequest.Marshal(b, m, deterministic)
}
func (m *ListCloudOperationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudOperationsRequest.Merge(m, src)
}
func (m *ListCloudOperationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListCloudOperationsRequest.Size(m)
}
func (m *ListCloudOperationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudOperationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudOperationsRequest proto.InternalMessageInfo

func (m *ListCloudOperationsRequest) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func (m *ListCloudOperationsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListCloudOperationsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListCloudOperationsResponse struct {
	// List of operations for the specified cloud.
	Operations []*operation.Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListCloudOperationsRequest.page_size], use the [next_page_token] as the value
	// for the [ListCloudOperationsRequest.page_token] query parameter in the next list request.
	// Each subsequent list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListCloudOperationsResponse) Reset()         { *m = ListCloudOperationsResponse{} }
func (m *ListCloudOperationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListCloudOperationsResponse) ProtoMessage()    {}
func (*ListCloudOperationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{4}
}

func (m *ListCloudOperationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListCloudOperationsResponse.Unmarshal(m, b)
}
func (m *ListCloudOperationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListCloudOperationsResponse.Marshal(b, m, deterministic)
}
func (m *ListCloudOperationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListCloudOperationsResponse.Merge(m, src)
}
func (m *ListCloudOperationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListCloudOperationsResponse.Size(m)
}
func (m *ListCloudOperationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListCloudOperationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListCloudOperationsResponse proto.InternalMessageInfo

func (m *ListCloudOperationsResponse) GetOperations() []*operation.Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *ListCloudOperationsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateCloudRequest struct {
	// ID of the cloud to update.
	// To get the cloud ID, use a [CloudService.List] request.
	CloudId string `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// Field mask that specifies which fields of the cloud are going to be updated.
	UpdateMask *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Name of the cloud.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the cloud.
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCloudRequest) Reset()         { *m = UpdateCloudRequest{} }
func (m *UpdateCloudRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateCloudRequest) ProtoMessage()    {}
func (*UpdateCloudRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{5}
}

func (m *UpdateCloudRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCloudRequest.Unmarshal(m, b)
}
func (m *UpdateCloudRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCloudRequest.Marshal(b, m, deterministic)
}
func (m *UpdateCloudRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCloudRequest.Merge(m, src)
}
func (m *UpdateCloudRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateCloudRequest.Size(m)
}
func (m *UpdateCloudRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCloudRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCloudRequest proto.InternalMessageInfo

func (m *UpdateCloudRequest) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func (m *UpdateCloudRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func (m *UpdateCloudRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateCloudRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateCloudMetadata struct {
	// ID of the cloud that is being updated.
	CloudId              string   `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateCloudMetadata) Reset()         { *m = UpdateCloudMetadata{} }
func (m *UpdateCloudMetadata) String() string { return proto.CompactTextString(m) }
func (*UpdateCloudMetadata) ProtoMessage()    {}
func (*UpdateCloudMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0ca0d4f81d9e76e, []int{6}
}

func (m *UpdateCloudMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCloudMetadata.Unmarshal(m, b)
}
func (m *UpdateCloudMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCloudMetadata.Marshal(b, m, deterministic)
}
func (m *UpdateCloudMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCloudMetadata.Merge(m, src)
}
func (m *UpdateCloudMetadata) XXX_Size() int {
	return xxx_messageInfo_UpdateCloudMetadata.Size(m)
}
func (m *UpdateCloudMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCloudMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCloudMetadata proto.InternalMessageInfo

func (m *UpdateCloudMetadata) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCloudRequest)(nil), "yandex.cloud.resourcemanager.v1.GetCloudRequest")
	proto.RegisterType((*ListCloudsRequest)(nil), "yandex.cloud.resourcemanager.v1.ListCloudsRequest")
	proto.RegisterType((*ListCloudsResponse)(nil), "yandex.cloud.resourcemanager.v1.ListCloudsResponse")
	proto.RegisterType((*ListCloudOperationsRequest)(nil), "yandex.cloud.resourcemanager.v1.ListCloudOperationsRequest")
	proto.RegisterType((*ListCloudOperationsResponse)(nil), "yandex.cloud.resourcemanager.v1.ListCloudOperationsResponse")
	proto.RegisterType((*UpdateCloudRequest)(nil), "yandex.cloud.resourcemanager.v1.UpdateCloudRequest")
	proto.RegisterType((*UpdateCloudMetadata)(nil), "yandex.cloud.resourcemanager.v1.UpdateCloudMetadata")
}

func init() {
	proto.RegisterFile("yandex/cloud/resourcemanager/v1/cloud_service.proto", fileDescriptor_c0ca0d4f81d9e76e)
}

var fileDescriptor_c0ca0d4f81d9e76e = []byte{
	// 886 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0xd6, 0x6c, 0xb6, 0xa1, 0xfb, 0x52, 0xa8, 0x3a, 0x80, 0x14, 0xdc, 0x56, 0x0d, 0x8e, 0x14,
	0x42, 0x16, 0xff, 0x4a, 0xd4, 0xa2, 0xdd, 0xa4, 0x0b, 0x4d, 0x81, 0x0a, 0x89, 0x0a, 0x70, 0xe0,
	0xc2, 0xaa, 0x8a, 0x66, 0xe3, 0x59, 0x63, 0x6d, 0x62, 0x1b, 0x8f, 0x13, 0xb5, 0x29, 0xbd, 0xc0,
	0x05, 0x72, 0xe1, 0x00, 0x07, 0xfe, 0x8e, 0x72, 0xe4, 0xce, 0xee, 0xb9, 0x5c, 0x39, 0x70, 0xe0,
	0x80, 0xc4, 0x8d, 0x0b, 0x12, 0xa7, 0xca, 0xe3, 0x1f, 0x9b, 0x38, 0xee, 0xc6, 0xe9, 0x69, 0xd7,
	0x7e, 0xdf, 0x37, 0xef, 0xfb, 0xbe, 0xbc, 0x37, 0x09, 0xb4, 0x1e, 0x10, 0xdb, 0xa0, 0xf7, 0x95,
	0xc1, 0xd0, 0x19, 0x1b, 0x8a, 0x47, 0x99, 0x33, 0xf6, 0x06, 0x74, 0x44, 0x6c, 0x62, 0x52, 0x4f,
	0x99, 0x68, 0x61, 0xa1, 0xcf, 0xa8, 0x37, 0xb1, 0x06, 0x54, 0x76, 0x3d, 0xc7, 0x77, 0xf0, 0xb5,
	0x90, 0x24, 0xf3, 0x9a, 0x9c, 0x22, 0xc9, 0x13, 0x4d, 0xb8, 0x62, 0x3a, 0x8e, 0x39, 0xa4, 0x0a,
	0x71, 0x2d, 0x85, 0xd8, 0xb6, 0xe3, 0x13, 0xdf, 0x72, 0x6c, 0x16, 0xd2, 0x85, 0x4a, 0x54, 0xe5,
	0x4f, 0x07, 0xe3, 0x43, 0xe5, 0xd0, 0xa2, 0x43, 0xa3, 0x3f, 0x22, 0xec, 0x28, 0x42, 0x6c, 0xe7,
	0x52, 0x15, 0x1f, 0xb7, 0x00, 0x0e, 0x5a, 0x3a, 0x2e, 0xf5, 0x78, 0xc7, 0x6c, 0xc4, 0x60, 0x40,
	0x19, 0x8b, 0xfe, 0x44, 0x88, 0xda, 0x02, 0x22, 0xe1, 0x2f, 0x9d, 0x74, 0x75, 0x01, 0x37, 0x21,
	0x43, 0xcb, 0x98, 0x2b, 0x8b, 0xbb, 0x70, 0xf1, 0x0e, 0xf5, 0x6f, 0x07, 0x45, 0x9d, 0x7e, 0x35,
	0xa6, 0xcc, 0xc7, 0x6f, 0xc0, 0xf9, 0x30, 0x42, 0xcb, 0x28, 0xa3, 0x0a, 0xaa, 0x6f, 0x75, 0x2f,
	0xfc, 0x7d, 0xac, 0xa1, 0xd9, 0x89, 0xb6, 0xd9, 0xb9, 0x79, 0x5d, 0xd5, 0x5f, 0xe0, 0xd5, 0x0f,
	0x0d, 0x71, 0x86, 0xe0, 0xd2, 0x47, 0x16, 0x0b, 0xd9, 0xec, 0x94, 0xbe, 0xe5, 0x12, 0x93, 0xf6,
	0x99, 0x35, 0xa5, 0x9c, 0x5f, 0xe8, 0xc2, 0xff, 0xc7, 0x5a, 0xb1, 0x73, 0x53, 0x53, 0x55, 0x55,
	0x3f, 0x1f, 0x14, 0x7b, 0xd6, 0x94, 0xe2, 0x3a, 0x00, 0x07, 0xfa, 0xce, 0x11, 0xb5, 0xcb, 0x1b,
	0xbc, 0xd3, 0xd6, 0xec, 0x44, 0x3b, 0xc7, 0x91, 0x3a, 0x3f, 0xe5, 0xb3, 0xa0, 0x86, 0x45, 0x28,
	0x1e, 0x5a, 0x43, 0x9f, 0x7a, 0xe5, 0x02, 0x47, 0xc1, 0xec, 0x24, 0x39, 0x2f, 0xaa, 0x88, 0x5f,
	0x03, 0x9e, 0xd7, 0xc2, 0x5c, 0xc7, 0x66, 0x14, 0xef, 0x41, 0x91, 0xab, 0x65, 0x65, 0x54, 0x29,
	0xd4, 0x4b, 0xcd, 0x9a, 0xbc, 0x62, 0x10, 0xe4, 0x30, 0x8a, 0x88, 0x85, 0x6b, 0x70, 0xd1, 0xa6,
	0xf7, 0xfd, 0x7e, 0x5a, 0xa8, 0xfe, 0x62, 0xf0, 0xfa, 0x93, 0x58, 0xa1, 0xf8, 0x33, 0x02, 0x21,
	0x69, 0xff, 0x71, 0xfc, 0x11, 0xb0, 0x75, 0x23, 0x5d, 0x0c, 0x6f, 0x23, 0x77, 0x78, 0x85, 0x67,
	0x87, 0x27, 0x7e, 0x87, 0xe0, 0x72, 0xa6, 0xb4, 0x28, 0xa2, 0x5b, 0x00, 0xc9, 0xcc, 0xc4, 0x31,
	0xbd, 0xbe, 0x18, 0xd3, 0xe9, 0x4c, 0x25, 0x7c, 0x7d, 0x8e, 0x94, 0x3b, 0xa5, 0x3f, 0x11, 0xe0,
	0xcf, 0x5d, 0x83, 0xf8, 0xf4, 0xb9, 0x06, 0x0e, 0xb7, 0xa1, 0x34, 0xe6, 0x74, 0xbe, 0x79, 0xbc,
	0x47, 0xa9, 0x29, 0xc8, 0xe1, 0x72, 0xca, 0xf1, 0x72, 0xca, 0x1f, 0x04, 0xcb, 0x79, 0x97, 0xb0,
	0x23, 0x1d, 0x42, 0x78, 0xf0, 0x3f, 0x7e, 0x1b, 0x36, 0x6d, 0x32, 0xa2, 0x51, 0x56, 0xd5, 0x7f,
	0x8f, 0xb5, 0x6b, 0xfb, 0x44, 0x9a, 0xde, 0xab, 0xef, 0x4b, 0x44, 0x9a, 0xaa, 0xd2, 0xce, 0xbd,
	0x87, 0xea, 0x5b, 0x37, 0xb4, 0x47, 0xfb, 0xd1, 0xd3, 0x9b, 0xef, 0xe8, 0x9c, 0x80, 0xb7, 0xa1,
	0x64, 0x50, 0x36, 0xf0, 0x2c, 0x37, 0x70, 0x5b, 0xde, 0x9c, 0xcf, 0xba, 0x79, 0xfd, 0x86, 0x3e,
	0x5f, 0x15, 0x55, 0x78, 0x79, 0xce, 0xe1, 0x5d, 0xea, 0x13, 0x83, 0xf8, 0x04, 0xbf, 0x96, 0xb6,
	0x98, 0x98, 0x6a, 0xfe, 0x01, 0x70, 0x81, 0x83, 0x7b, 0xe1, 0x8d, 0x85, 0x7f, 0x40, 0x50, 0xb8,
	0x43, 0x7d, 0xac, 0xae, 0x9c, 0xd5, 0xd4, 0xe6, 0x0a, 0x39, 0xa7, 0x5b, 0x94, 0xbf, 0xf9, 0xfd,
	0xaf, 0x1f, 0x37, 0xea, 0xb8, 0x96, 0x5c, 0x54, 0x52, 0xfa, 0xa6, 0x62, 0xca, 0xc3, 0x58, 0xf0,
	0x23, 0xfc, 0x13, 0x82, 0xcd, 0x60, 0x84, 0x70, 0x73, 0x65, 0x83, 0xa5, 0xfb, 0x40, 0x68, 0xad,
	0xc5, 0x09, 0x87, 0x52, 0xac, 0x72, 0x85, 0x57, 0xf1, 0xe5, 0x33, 0x14, 0xe2, 0x5f, 0x10, 0x14,
	0xc3, 0xb0, 0xf1, 0xea, 0x26, 0xcb, 0x73, 0x27, 0xac, 0x9e, 0x72, 0xf1, 0xd3, 0xc7, 0x4f, 0x1a,
	0x57, 0xb2, 0x3f, 0xd2, 0x73, 0xfc, 0x91, 0xeb, 0xdc, 0x6e, 0xe6, 0x4c, 0x72, 0x17, 0x35, 0xf0,
	0x6f, 0x08, 0x5e, 0x0a, 0x1c, 0x9f, 0xae, 0x22, 0x6e, 0xe7, 0x8f, 0x68, 0xe9, 0x6e, 0x11, 0x3a,
	0xcf, 0x47, 0x8e, 0x82, 0xde, 0xe1, 0x06, 0x5a, 0x58, 0xcb, 0x67, 0x40, 0x99, 0xdb, 0xfa, 0x5f,
	0x51, 0x78, 0xe5, 0xde, 0xe2, 0x5f, 0x4b, 0x5d, 0xcb, 0x36, 0x2c, 0xdb, 0x64, 0x58, 0x5e, 0xd4,
	0x13, 0x7d, 0x69, 0x2d, 0x03, 0x63, 0xfd, 0x4a, 0x6e, 0x7c, 0x24, 0xf9, 0x3d, 0x2e, 0x79, 0x0f,
	0x77, 0xce, 0x94, 0x1c, 0xd7, 0x78, 0xec, 0xc3, 0x65, 0x99, 0xff, 0x20, 0xb8, 0xd4, 0xa3, 0xe9,
	0xb7, 0x52, 0xa6, 0x98, 0x25, 0xdc, 0x1a, 0x13, 0xf4, 0x2d, 0x7a, 0xfc, 0xa4, 0xb1, 0x03, 0x95,
	0x67, 0x1d, 0x95, 0xcc, 0xd3, 0xab, 0xe9, 0x6b, 0xec, 0xfd, 0x91, 0xeb, 0x3f, 0xe0, 0x5e, 0xdf,
	0x15, 0xdb, 0xf9, 0xbd, 0xb2, 0x74, 0x87, 0x60, 0xe8, 0xfe, 0x43, 0xf0, 0x4a, 0x38, 0xc4, 0x29,
	0xc3, 0x6a, 0xa6, 0xe1, 0x2c, 0xe8, 0x1a, 0x9e, 0xbf, 0x0f, 0x3c, 0x77, 0xa0, 0x7a, 0xc6, 0x69,
	0x79, 0x6c, 0xdf, 0x16, 0xf7, 0xf2, 0xdb, 0x1e, 0x67, 0x34, 0xd9, 0x45, 0x8d, 0xee, 0x14, 0xaa,
	0x8b, 0x0e, 0x5d, 0x2b, 0x63, 0x47, 0xbe, 0xe8, 0x99, 0x96, 0xff, 0xe5, 0xf8, 0x40, 0x1e, 0x38,
	0x23, 0x25, 0xc4, 0x4b, 0xe1, 0x2f, 0x26, 0xd3, 0x91, 0x4c, 0x6a, 0x73, 0x75, 0xca, 0x8a, 0xdf,
	0x78, 0xed, 0xd4, 0xab, 0x83, 0x22, 0xa7, 0xb5, 0x9e, 0x06, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x80,
	0x26, 0x44, 0xb3, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CloudServiceClient is the client API for CloudService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudServiceClient interface {
	// Returns the specified Cloud resource.
	//
	// To get the list of available Cloud resources, make a [List] request.
	Get(ctx context.Context, in *GetCloudRequest, opts ...grpc.CallOption) (*Cloud, error)
	// Retrieves the list of Cloud resources.
	List(ctx context.Context, in *ListCloudsRequest, opts ...grpc.CallOption) (*ListCloudsResponse, error)
	// Updates the specified cloud.
	Update(ctx context.Context, in *UpdateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified cloud.
	ListOperations(ctx context.Context, in *ListCloudOperationsRequest, opts ...grpc.CallOption) (*ListCloudOperationsResponse, error)
	// Lists access bindings for the specified cloud.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified cloud.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified cloud.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type cloudServiceClient struct {
	cc *grpc.ClientConn
}

func NewCloudServiceClient(cc *grpc.ClientConn) CloudServiceClient {
	return &cloudServiceClient{cc}
}

func (c *cloudServiceClient) Get(ctx context.Context, in *GetCloudRequest, opts ...grpc.CallOption) (*Cloud, error) {
	out := new(Cloud)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) List(ctx context.Context, in *ListCloudsRequest, opts ...grpc.CallOption) (*ListCloudsResponse, error) {
	out := new(ListCloudsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) Update(ctx context.Context, in *UpdateCloudRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) ListOperations(ctx context.Context, in *ListCloudOperationsRequest, opts ...grpc.CallOption) (*ListCloudOperationsResponse, error) {
	out := new(ListCloudOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/ListAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/SetAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.resourcemanager.v1.CloudService/UpdateAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudServiceServer is the server API for CloudService service.
type CloudServiceServer interface {
	// Returns the specified Cloud resource.
	//
	// To get the list of available Cloud resources, make a [List] request.
	Get(context.Context, *GetCloudRequest) (*Cloud, error)
	// Retrieves the list of Cloud resources.
	List(context.Context, *ListCloudsRequest) (*ListCloudsResponse, error)
	// Updates the specified cloud.
	Update(context.Context, *UpdateCloudRequest) (*operation.Operation, error)
	// Lists operations for the specified cloud.
	ListOperations(context.Context, *ListCloudOperationsRequest) (*ListCloudOperationsResponse, error)
	// Lists access bindings for the specified cloud.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified cloud.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified cloud.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedCloudServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCloudServiceServer struct {
}

func (*UnimplementedCloudServiceServer) Get(ctx context.Context, req *GetCloudRequest) (*Cloud, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCloudServiceServer) List(ctx context.Context, req *ListCloudsRequest) (*ListCloudsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedCloudServiceServer) Update(ctx context.Context, req *UpdateCloudRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedCloudServiceServer) ListOperations(ctx context.Context, req *ListCloudOperationsRequest) (*ListCloudOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (*UnimplementedCloudServiceServer) ListAccessBindings(ctx context.Context, req *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (*UnimplementedCloudServiceServer) SetAccessBindings(ctx context.Context, req *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (*UnimplementedCloudServiceServer) UpdateAccessBindings(ctx context.Context, req *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

func RegisterCloudServiceServer(s *grpc.Server, srv CloudServiceServer) {
	s.RegisterService(&_CloudService_serviceDesc, srv)
}

func _CloudService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Get(ctx, req.(*GetCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).List(ctx, req.(*ListCloudsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCloudRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).Update(ctx, req.(*UpdateCloudRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCloudOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).ListOperations(ctx, req.(*ListCloudOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/ListAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/SetAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.resourcemanager.v1.CloudService/UpdateAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.resourcemanager.v1.CloudService",
	HandlerType: (*CloudServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CloudService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CloudService_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CloudService_Update_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _CloudService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _CloudService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _CloudService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _CloudService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/resourcemanager/v1/cloud_service.proto",
}
