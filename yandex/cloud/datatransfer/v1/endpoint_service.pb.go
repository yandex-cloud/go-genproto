// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datatransfer/v1/endpoint_service.proto

package datatransfer

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetEndpointRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of the endpoint to return.
	//
	// To get the endpoint ID, make an [EndpointService.List] request.
	EndpointId    string `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetEndpointRequest) Reset() {
	*x = GetEndpointRequest{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEndpointRequest) ProtoMessage() {}

func (x *GetEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEndpointRequest.ProtoReflect.Descriptor instead.
func (*GetEndpointRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetEndpointRequest) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

type ListEndpointsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of the folder containing the endpoints to be listed.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of endpoints to be sent in the response message. If the
	// folder contains more endpoints than `page_size`, `next_page_token` will be
	// included
	// in the response message. Include it into the subsequent `ListEndpointRequest` to
	// fetch the next page. Defaults to `100` if not specified. The maximum allowed
	// value
	// for this field is `1000`.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Opaque value identifying the endpoints page to be fetched. Should be empty in
	// the first `ListEndpointsRequest`. Subsequent requests should have this field
	// filled
	// with the `next_page_token` from the previous `ListEndpointsResponse`.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEndpointsRequest) Reset() {
	*x = ListEndpointsRequest{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEndpointsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointsRequest) ProtoMessage() {}

func (x *ListEndpointsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointsRequest.ProtoReflect.Descriptor instead.
func (*ListEndpointsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListEndpointsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListEndpointsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListEndpointsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListEndpointsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The list of endpoints. If there are more endpoints in the folder, then
	// `next_page_token` is a non-empty string to be included into the subsequent
	// `ListEndpointsRequest` to fetch the next endpoints page.
	Endpoints []*Endpoint `protobuf:"bytes,1,rep,name=endpoints,proto3" json:"endpoints,omitempty"`
	// Opaque value identifying the next endpoints page. This field is empty if there
	// are no more endpoints in the folder. Otherwise, it is non-empty and should be
	// included in the subsequent `ListEndpointsRequest` to fetch the next endpoints
	// page.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListEndpointsResponse) Reset() {
	*x = ListEndpointsResponse{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListEndpointsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEndpointsResponse) ProtoMessage() {}

func (x *ListEndpointsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEndpointsResponse.ProtoReflect.Descriptor instead.
func (*ListEndpointsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListEndpointsResponse) GetEndpoints() []*Endpoint {
	if x != nil {
		return x.Endpoints
	}
	return nil
}

func (x *ListEndpointsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateEndpointRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the folder to create the endpoint in.
	//
	// To get the folder ID, make a
	// [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the endpoint.
	//
	// The name must be unique within the folder.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the endpoint.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Endpoint labels as `key:value` pairs.
	//
	// For details about the concept, see [documentation]({{ api-url-prefix
	// }}/resource-manager/concepts/labels).
	Labels        map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Settings      *EndpointSettings `protobuf:"bytes,52,opt,name=settings,proto3" json:"settings,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEndpointRequest) Reset() {
	*x = CreateEndpointRequest{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEndpointRequest) ProtoMessage() {}

func (x *CreateEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEndpointRequest.ProtoReflect.Descriptor instead.
func (*CreateEndpointRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateEndpointRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateEndpointRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateEndpointRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateEndpointRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateEndpointRequest) GetSettings() *EndpointSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

type CreateEndpointMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EndpointId    string                 `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateEndpointMetadata) Reset() {
	*x = CreateEndpointMetadata{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateEndpointMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEndpointMetadata) ProtoMessage() {}

func (x *CreateEndpointMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEndpointMetadata.ProtoReflect.Descriptor instead.
func (*CreateEndpointMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateEndpointMetadata) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

type UpdateEndpointRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of the endpoint to be updated.
	EndpointId string `protobuf:"bytes,10,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	// The new endpoint name. Must be unique within the folder.
	Name string `protobuf:"bytes,11,opt,name=name,proto3" json:"name,omitempty"`
	// The new description for the endpoint.
	Description string `protobuf:"bytes,12,opt,name=description,proto3" json:"description,omitempty"`
	// Endpoint labels as `key:value` pairs.
	//
	// For details about the concept, see [documentation]({{ api-url-prefix
	// }}/resource-manager/concepts/labels).
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The new endpoint settings.
	Settings *EndpointSettings `protobuf:"bytes,52,opt,name=settings,proto3" json:"settings,omitempty"`
	// Field mask specifying endpoint fields to be updated. Semantics for this field is
	// described here:
	// <https://pkg.go.dev/google.golang.org/protobuf/types/known/fieldmaskpb#FieldMask>
	// The only exception: if the repeated field is specified in the mask, then
	// the new value replaces the old one instead of being appended to the old one.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,60,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEndpointRequest) Reset() {
	*x = UpdateEndpointRequest{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEndpointRequest) ProtoMessage() {}

func (x *UpdateEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEndpointRequest.ProtoReflect.Descriptor instead.
func (*UpdateEndpointRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateEndpointRequest) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

func (x *UpdateEndpointRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateEndpointRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateEndpointRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *UpdateEndpointRequest) GetSettings() *EndpointSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *UpdateEndpointRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type UpdateEndpointMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EndpointId    string                 `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateEndpointMetadata) Reset() {
	*x = UpdateEndpointMetadata{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateEndpointMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEndpointMetadata) ProtoMessage() {}

func (x *UpdateEndpointMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEndpointMetadata.ProtoReflect.Descriptor instead.
func (*UpdateEndpointMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateEndpointMetadata) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

type DeleteEndpointRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifier of the endpoint to delete.
	//
	// To get the list of all available endpoints, make a [List] request.
	EndpointId    string `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEndpointRequest) Reset() {
	*x = DeleteEndpointRequest{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEndpointRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointRequest) ProtoMessage() {}

func (x *DeleteEndpointRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointRequest.ProtoReflect.Descriptor instead.
func (*DeleteEndpointRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteEndpointRequest) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

type DeleteEndpointMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EndpointId    string                 `protobuf:"bytes,1,opt,name=endpoint_id,json=endpointId,proto3" json:"endpoint_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteEndpointMetadata) Reset() {
	*x = DeleteEndpointMetadata{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteEndpointMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEndpointMetadata) ProtoMessage() {}

func (x *DeleteEndpointMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEndpointMetadata.ProtoReflect.Descriptor instead.
func (*DeleteEndpointMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteEndpointMetadata) GetEndpointId() string {
	if x != nil {
		return x.EndpointId
	}
	return ""
}

var File_yandex_cloud_datatransfer_v1_endpoint_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDesc = "" +
	"\n" +
	"3yandex/cloud/datatransfer/v1/endpoint_service.proto\x12\x1cyandex.cloud.datatransfer.v1\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\x1a+yandex/cloud/datatransfer/v1/endpoint.proto\x1a yandex/cloud/api/operation.proto\x1a&yandex/cloud/operation/operation.proto\"5\n" +
	"\x12GetEndpointRequest\x12\x1f\n" +
	"\vendpoint_id\x18\x01 \x01(\tR\n" +
	"endpointId\"o\n" +
	"\x14ListEndpointsRequest\x12\x1b\n" +
	"\tfolder_id\x18\x01 \x01(\tR\bfolderId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"\x85\x01\n" +
	"\x15ListEndpointsResponse\x12D\n" +
	"\tendpoints\x18\x01 \x03(\v2&.yandex.cloud.datatransfer.v1.EndpointR\tendpoints\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\xd0\x02\n" +
	"\x15CreateEndpointRequest\x12\x1b\n" +
	"\tfolder_id\x18\x01 \x01(\tR\bfolderId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12W\n" +
	"\x06labels\x18\x04 \x03(\v2?.yandex.cloud.datatransfer.v1.CreateEndpointRequest.LabelsEntryR\x06labels\x12J\n" +
	"\bsettings\x184 \x01(\v2..yandex.cloud.datatransfer.v1.EndpointSettingsR\bsettings\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01J\x04\b\x05\x104\"9\n" +
	"\x16CreateEndpointMetadata\x12\x1f\n" +
	"\vendpoint_id\x18\x01 \x01(\tR\n" +
	"endpointId\"\x9d\x03\n" +
	"\x15UpdateEndpointRequest\x12\x1f\n" +
	"\vendpoint_id\x18\n" +
	" \x01(\tR\n" +
	"endpointId\x12\x12\n" +
	"\x04name\x18\v \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\f \x01(\tR\vdescription\x12W\n" +
	"\x06labels\x18\r \x03(\v2?.yandex.cloud.datatransfer.v1.UpdateEndpointRequest.LabelsEntryR\x06labels\x12J\n" +
	"\bsettings\x184 \x01(\v2..yandex.cloud.datatransfer.v1.EndpointSettingsR\bsettings\x12;\n" +
	"\vupdate_mask\x18< \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01J\x04\b\x01\x10\n" +
	"J\x04\b\x0e\x104J\x04\b5\x10<\"9\n" +
	"\x16UpdateEndpointMetadata\x12\x1f\n" +
	"\vendpoint_id\x18\x01 \x01(\tR\n" +
	"endpointId\"8\n" +
	"\x15DeleteEndpointRequest\x12\x1f\n" +
	"\vendpoint_id\x18\x01 \x01(\tR\n" +
	"endpointId\"9\n" +
	"\x16DeleteEndpointMetadata\x12\x1f\n" +
	"\vendpoint_id\x18\x01 \x01(\tR\n" +
	"endpointId2\xbd\x06\n" +
	"\x0fEndpointService\x12\x83\x01\n" +
	"\x03Get\x120.yandex.cloud.datatransfer.v1.GetEndpointRequest\x1a&.yandex.cloud.datatransfer.v1.Endpoint\"\"\x82\xd3\xe4\x93\x02\x1c\x12\x1a/v1/endpoint/{endpoint_id}\x12\x97\x01\n" +
	"\x04List\x122.yandex.cloud.datatransfer.v1.ListEndpointsRequest\x1a3.yandex.cloud.datatransfer.v1.ListEndpointsResponse\"&\x82\xd3\xe4\x93\x02 \x12\x1e/v1/endpoints/list/{folder_id}\x12\x9f\x01\n" +
	"\x06Create\x123.yandex.cloud.datatransfer.v1.CreateEndpointRequest\x1a!.yandex.cloud.operation.Operation\"=\xb2\xd2*\"\n" +
	"\x16CreateEndpointMetadata\x12\bEndpoint\x82\xd3\xe4\x93\x02\x11:\x01*\"\f/v1/endpoint\x12\xad\x01\n" +
	"\x06Update\x123.yandex.cloud.datatransfer.v1.UpdateEndpointRequest\x1a!.yandex.cloud.operation.Operation\"K\xb2\xd2*\"\n" +
	"\x16UpdateEndpointMetadata\x12\bEndpoint\x82\xd3\xe4\x93\x02\x1f:\x01*2\x1a/v1/endpoint/{endpoint_id}\x12\xb7\x01\n" +
	"\x06Delete\x123.yandex.cloud.datatransfer.v1.DeleteEndpointRequest\x1a!.yandex.cloud.operation.Operation\"U\xb2\xd2*/\n" +
	"\x16DeleteEndpointMetadata\x12\x15google.protobuf.Empty\x82\xd3\xe4\x93\x02\x1c*\x1a/v1/endpoint/{endpoint_id}Bq\n" +
	" yandex.cloud.api.datatransfer.v1ZMgithub.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1;datatransferb\x06proto3"

var (
	file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescData []byte
)

func file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDesc)))
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_yandex_cloud_datatransfer_v1_endpoint_service_proto_goTypes = []any{
	(*GetEndpointRequest)(nil),     // 0: yandex.cloud.datatransfer.v1.GetEndpointRequest
	(*ListEndpointsRequest)(nil),   // 1: yandex.cloud.datatransfer.v1.ListEndpointsRequest
	(*ListEndpointsResponse)(nil),  // 2: yandex.cloud.datatransfer.v1.ListEndpointsResponse
	(*CreateEndpointRequest)(nil),  // 3: yandex.cloud.datatransfer.v1.CreateEndpointRequest
	(*CreateEndpointMetadata)(nil), // 4: yandex.cloud.datatransfer.v1.CreateEndpointMetadata
	(*UpdateEndpointRequest)(nil),  // 5: yandex.cloud.datatransfer.v1.UpdateEndpointRequest
	(*UpdateEndpointMetadata)(nil), // 6: yandex.cloud.datatransfer.v1.UpdateEndpointMetadata
	(*DeleteEndpointRequest)(nil),  // 7: yandex.cloud.datatransfer.v1.DeleteEndpointRequest
	(*DeleteEndpointMetadata)(nil), // 8: yandex.cloud.datatransfer.v1.DeleteEndpointMetadata
	nil,                            // 9: yandex.cloud.datatransfer.v1.CreateEndpointRequest.LabelsEntry
	nil,                            // 10: yandex.cloud.datatransfer.v1.UpdateEndpointRequest.LabelsEntry
	(*Endpoint)(nil),               // 11: yandex.cloud.datatransfer.v1.Endpoint
	(*EndpointSettings)(nil),       // 12: yandex.cloud.datatransfer.v1.EndpointSettings
	(*fieldmaskpb.FieldMask)(nil),  // 13: google.protobuf.FieldMask
	(*operation.Operation)(nil),    // 14: yandex.cloud.operation.Operation
}
var file_yandex_cloud_datatransfer_v1_endpoint_service_proto_depIdxs = []int32{
	11, // 0: yandex.cloud.datatransfer.v1.ListEndpointsResponse.endpoints:type_name -> yandex.cloud.datatransfer.v1.Endpoint
	9,  // 1: yandex.cloud.datatransfer.v1.CreateEndpointRequest.labels:type_name -> yandex.cloud.datatransfer.v1.CreateEndpointRequest.LabelsEntry
	12, // 2: yandex.cloud.datatransfer.v1.CreateEndpointRequest.settings:type_name -> yandex.cloud.datatransfer.v1.EndpointSettings
	10, // 3: yandex.cloud.datatransfer.v1.UpdateEndpointRequest.labels:type_name -> yandex.cloud.datatransfer.v1.UpdateEndpointRequest.LabelsEntry
	12, // 4: yandex.cloud.datatransfer.v1.UpdateEndpointRequest.settings:type_name -> yandex.cloud.datatransfer.v1.EndpointSettings
	13, // 5: yandex.cloud.datatransfer.v1.UpdateEndpointRequest.update_mask:type_name -> google.protobuf.FieldMask
	0,  // 6: yandex.cloud.datatransfer.v1.EndpointService.Get:input_type -> yandex.cloud.datatransfer.v1.GetEndpointRequest
	1,  // 7: yandex.cloud.datatransfer.v1.EndpointService.List:input_type -> yandex.cloud.datatransfer.v1.ListEndpointsRequest
	3,  // 8: yandex.cloud.datatransfer.v1.EndpointService.Create:input_type -> yandex.cloud.datatransfer.v1.CreateEndpointRequest
	5,  // 9: yandex.cloud.datatransfer.v1.EndpointService.Update:input_type -> yandex.cloud.datatransfer.v1.UpdateEndpointRequest
	7,  // 10: yandex.cloud.datatransfer.v1.EndpointService.Delete:input_type -> yandex.cloud.datatransfer.v1.DeleteEndpointRequest
	11, // 11: yandex.cloud.datatransfer.v1.EndpointService.Get:output_type -> yandex.cloud.datatransfer.v1.Endpoint
	2,  // 12: yandex.cloud.datatransfer.v1.EndpointService.List:output_type -> yandex.cloud.datatransfer.v1.ListEndpointsResponse
	14, // 13: yandex.cloud.datatransfer.v1.EndpointService.Create:output_type -> yandex.cloud.operation.Operation
	14, // 14: yandex.cloud.datatransfer.v1.EndpointService.Update:output_type -> yandex.cloud.operation.Operation
	14, // 15: yandex.cloud.datatransfer.v1.EndpointService.Delete:output_type -> yandex.cloud.operation.Operation
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_service_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_service_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_service_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_service_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_service_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_service_proto_depIdxs = nil
}
