// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/apploadbalancer/v1/http_router_service.proto

package apploadbalancer

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GetHttpRouterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router to return.
	//
	// To get the HTTP router ID, make a [HttpRouterService.List] request.
	HttpRouterId  string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetHttpRouterRequest) Reset() {
	*x = GetHttpRouterRequest{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetHttpRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHttpRouterRequest) ProtoMessage() {}

func (x *GetHttpRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHttpRouterRequest.ProtoReflect.Descriptor instead.
func (*GetHttpRouterRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetHttpRouterRequest) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

type ListHttpRoutersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the folder to list HTTP routers in.
	//
	// To get the folder ID, make a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than `page_size`, the service returns a [ListHttpRoutersResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListHttpRoutersResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters HTTP routers listed in the response.
	//
	// The expression must specify:
	// 1. The field name. Currently you can use filtering only on [HttpRouter.name] field.
	// 2. An `=` operator.
	// 3. The value in double quotes (`"`). Must be 3-63 characters long and match the regular expression `[a-z][-a-z0-9]{1,61}[a-z0-9]`.
	// Example of a filter: `name=my-http-router`.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHttpRoutersRequest) Reset() {
	*x = ListHttpRoutersRequest{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHttpRoutersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHttpRoutersRequest) ProtoMessage() {}

func (x *ListHttpRoutersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHttpRoutersRequest.ProtoReflect.Descriptor instead.
func (*ListHttpRoutersRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListHttpRoutersRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListHttpRoutersRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListHttpRoutersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListHttpRoutersRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListHttpRoutersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of HTTP routers in the specified folder.
	HttpRouters []*HttpRouter `protobuf:"bytes,1,rep,name=http_routers,json=httpRouters,proto3" json:"http_routers,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// the specified [ListHttpRoutersRequest.page_size], use `next_page_token` as the value
	// for the [ListHttpRoutersRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHttpRoutersResponse) Reset() {
	*x = ListHttpRoutersResponse{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHttpRoutersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHttpRoutersResponse) ProtoMessage() {}

func (x *ListHttpRoutersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHttpRoutersResponse.ProtoReflect.Descriptor instead.
func (*ListHttpRoutersResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListHttpRoutersResponse) GetHttpRouters() []*HttpRouter {
	if x != nil {
		return x.HttpRouters
	}
	return nil
}

func (x *ListHttpRoutersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type DeleteHttpRouterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router to delete.
	//
	// To get the HTTP router ID, make a [HttpRouterService.List] request.
	HttpRouterId  string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteHttpRouterRequest) Reset() {
	*x = DeleteHttpRouterRequest{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteHttpRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHttpRouterRequest) ProtoMessage() {}

func (x *DeleteHttpRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHttpRouterRequest.ProtoReflect.Descriptor instead.
func (*DeleteHttpRouterRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteHttpRouterRequest) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

type DeleteHttpRouterMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router that is being deleted.
	HttpRouterId  string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteHttpRouterMetadata) Reset() {
	*x = DeleteHttpRouterMetadata{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteHttpRouterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHttpRouterMetadata) ProtoMessage() {}

func (x *DeleteHttpRouterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHttpRouterMetadata.ProtoReflect.Descriptor instead.
func (*DeleteHttpRouterMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteHttpRouterMetadata) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

type UpdateHttpRouterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router to update.
	//
	// To get the HTTP router ID, make a [HttpRouterService.List] request.
	HttpRouterId string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	// Field mask that specifies which attributes of the HTTP router should be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// New name for the HTTP router.
	// The name must be unique within the folder.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// New description of the HTTP router.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// HTTP router labels as `key:value` pairs.
	// For details about the concept, see [documentation](/docs/overview/concepts/services#labels).
	//
	// Existing set of labels is completely replaced by the provided set, so if you just want
	// to add or remove a label:
	// 1. Get the current set of labels with a [HttpRouterService.Get] request.
	// 2. Add or remove a label in this set.
	// 3. Send the new set in this field.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// New virtual hosts that combine routes inside the router.
	// For details about the concept, see [documentation](/docs/application-load-balancer/concepts/http-router#virtual-host).
	//
	// Only one virtual host with no authority (default match) can be specified.
	//
	// Existing list of virtual hosts is completely replaced by the specified list, so if you just want to add or remove
	// a virtual host, make a [VirtualHostService.Create] request or a [VirtualHostService.Delete] request.
	VirtualHosts []*VirtualHost `protobuf:"bytes,6,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// New route options for the HTTP router.
	RouteOptions  *RouteOptions `protobuf:"bytes,8,opt,name=route_options,json=routeOptions,proto3" json:"route_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateHttpRouterRequest) Reset() {
	*x = UpdateHttpRouterRequest{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateHttpRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHttpRouterRequest) ProtoMessage() {}

func (x *UpdateHttpRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHttpRouterRequest.ProtoReflect.Descriptor instead.
func (*UpdateHttpRouterRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateHttpRouterRequest) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

func (x *UpdateHttpRouterRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateHttpRouterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateHttpRouterRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateHttpRouterRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *UpdateHttpRouterRequest) GetVirtualHosts() []*VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

func (x *UpdateHttpRouterRequest) GetRouteOptions() *RouteOptions {
	if x != nil {
		return x.RouteOptions
	}
	return nil
}

type UpdateHttpRouterMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router that is being updated.
	HttpRouterId  string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateHttpRouterMetadata) Reset() {
	*x = UpdateHttpRouterMetadata{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateHttpRouterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHttpRouterMetadata) ProtoMessage() {}

func (x *UpdateHttpRouterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHttpRouterMetadata.ProtoReflect.Descriptor instead.
func (*UpdateHttpRouterMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateHttpRouterMetadata) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

type CreateHttpRouterRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the folder to create an HTTP router in.
	//
	// To get the folder ID, make a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the HTTP router.
	// The name must be unique within the folder.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the HTTP router.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// HTTP router labels as `key:value` pairs.
	// For details about the concept, see [documentation](/docs/overview/concepts/services#labels).
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Virtual hosts that combine routes inside the router.
	// For details about the concept, see [documentation](/docs/application-load-balancer/concepts/http-router#virtual-host).
	//
	// Only one virtual host with no authority (default match) can be specified.
	VirtualHosts []*VirtualHost `protobuf:"bytes,5,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// Route options for the HTTP router.
	RouteOptions  *RouteOptions `protobuf:"bytes,7,opt,name=route_options,json=routeOptions,proto3" json:"route_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateHttpRouterRequest) Reset() {
	*x = CreateHttpRouterRequest{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHttpRouterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHttpRouterRequest) ProtoMessage() {}

func (x *CreateHttpRouterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHttpRouterRequest.ProtoReflect.Descriptor instead.
func (*CreateHttpRouterRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateHttpRouterRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateHttpRouterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateHttpRouterRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateHttpRouterRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateHttpRouterRequest) GetVirtualHosts() []*VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

func (x *CreateHttpRouterRequest) GetRouteOptions() *RouteOptions {
	if x != nil {
		return x.RouteOptions
	}
	return nil
}

type CreateHttpRouterMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router that is being created.
	HttpRouterId  string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateHttpRouterMetadata) Reset() {
	*x = CreateHttpRouterMetadata{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateHttpRouterMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHttpRouterMetadata) ProtoMessage() {}

func (x *CreateHttpRouterMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHttpRouterMetadata.ProtoReflect.Descriptor instead.
func (*CreateHttpRouterMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateHttpRouterMetadata) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

type ListHttpRouterOperationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the HTTP router to get operations for.
	//
	// To get the HTTP router ID, use a [HttpRouterService.List] request.
	HttpRouterId string `protobuf:"bytes,1,opt,name=http_router_id,json=httpRouterId,proto3" json:"http_router_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than [page_size], the service returns a [ListHttpRouterOperationsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListHttpRouterOperationsResponse.next_page_token] returned by a previous list request.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHttpRouterOperationsRequest) Reset() {
	*x = ListHttpRouterOperationsRequest{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHttpRouterOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHttpRouterOperationsRequest) ProtoMessage() {}

func (x *ListHttpRouterOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHttpRouterOperationsRequest.ProtoReflect.Descriptor instead.
func (*ListHttpRouterOperationsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListHttpRouterOperationsRequest) GetHttpRouterId() string {
	if x != nil {
		return x.HttpRouterId
	}
	return ""
}

func (x *ListHttpRouterOperationsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListHttpRouterOperationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListHttpRouterOperationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of operations for the specified HTTP router.
	Operations []*operation.Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// the specified [ListHttpRouterOperationsRequest.page_size], use `next_page_token` as the value
	// for the [ListHttpRouterOperationsRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListHttpRouterOperationsResponse) Reset() {
	*x = ListHttpRouterOperationsResponse{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListHttpRouterOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHttpRouterOperationsResponse) ProtoMessage() {}

func (x *ListHttpRouterOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHttpRouterOperationsResponse.ProtoReflect.Descriptor instead.
func (*ListHttpRouterOperationsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP(), []int{10}
}

func (x *ListHttpRouterOperationsResponse) GetOperations() []*operation.Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *ListHttpRouterOperationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_apploadbalancer_v1_http_router_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDesc = "" +
	"\n" +
	"9yandex/cloud/apploadbalancer/v1/http_router_service.proto\x12\x1fyandex.cloud.apploadbalancer.v1\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\x1a yandex/cloud/api/operation.proto\x1a&yandex/cloud/operation/operation.proto\x1a1yandex/cloud/apploadbalancer/v1/http_router.proto\x1a2yandex/cloud/apploadbalancer/v1/virtual_host.proto\x1a\x1dyandex/cloud/validation.proto\"B\n" +
	"\x14GetHttpRouterRequest\x12*\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\fhttpRouterId\"\xb2\x01\n" +
	"\x16ListHttpRoutersRequest\x12!\n" +
	"\tfolder_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bfolderId\x12'\n" +
	"\tpage_size\x18\x02 \x01(\x03B\n" +
	"\xfa\xc71\x060-1000R\bpageSize\x12(\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\t\x8a\xc81\x05<=100R\tpageToken\x12\"\n" +
	"\x06filter\x18\x04 \x01(\tB\n" +
	"\x8a\xc81\x06<=1000R\x06filter\"\x91\x01\n" +
	"\x17ListHttpRoutersResponse\x12N\n" +
	"\fhttp_routers\x18\x01 \x03(\v2+.yandex.cloud.apploadbalancer.v1.HttpRouterR\vhttpRouters\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"E\n" +
	"\x17DeleteHttpRouterRequest\x12*\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\fhttpRouterId\"@\n" +
	"\x18DeleteHttpRouterMetadata\x12$\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tR\fhttpRouterId\"\xf7\x04\n" +
	"\x17UpdateHttpRouterRequest\x12*\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\fhttpRouterId\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12:\n" +
	"\x04name\x18\x03 \x01(\tB&\xf2\xc71\"([a-z]([-a-z0-9]{0,61}[a-z0-9])?)?R\x04name\x12+\n" +
	"\vdescription\x18\x04 \x01(\tB\t\x8a\xc81\x05<=256R\vdescription\x12\xa1\x01\n" +
	"\x06labels\x18\x05 \x03(\v2D.yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.LabelsEntryBC\xf2\xc71\x0f[-_./\\@0-9a-z]*\x82\xc81\x04<=64\x8a\xc81\x04<=63\xb2\xc81\x1c\x12\x14[a-z][-_./\\@0-9a-z]*\x1a\x041-63R\x06labels\x12Q\n" +
	"\rvirtual_hosts\x18\x06 \x03(\v2,.yandex.cloud.apploadbalancer.v1.VirtualHostR\fvirtualHosts\x12R\n" +
	"\rroute_options\x18\b \x01(\v2-.yandex.cloud.apploadbalancer.v1.RouteOptionsR\frouteOptions\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01J\x04\b\a\x10\b\"@\n" +
	"\x18UpdateHttpRouterMetadata\x12$\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tR\fhttpRouterId\"\xb1\x04\n" +
	"\x17CreateHttpRouterRequest\x12!\n" +
	"\tfolder_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bfolderId\x12:\n" +
	"\x04name\x18\x02 \x01(\tB&\xf2\xc71\"([a-z]([-a-z0-9]{0,61}[a-z0-9])?)?R\x04name\x12+\n" +
	"\vdescription\x18\x03 \x01(\tB\t\x8a\xc81\x05<=256R\vdescription\x12\xa1\x01\n" +
	"\x06labels\x18\x04 \x03(\v2D.yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest.LabelsEntryBC\xf2\xc71\x0f[-_./\\@0-9a-z]*\x82\xc81\x04<=64\x8a\xc81\x04<=63\xb2\xc81\x1c\x12\x14[a-z][-_./\\@0-9a-z]*\x1a\x041-63R\x06labels\x12Q\n" +
	"\rvirtual_hosts\x18\x05 \x03(\v2,.yandex.cloud.apploadbalancer.v1.VirtualHostR\fvirtualHosts\x12R\n" +
	"\rroute_options\x18\a \x01(\v2-.yandex.cloud.apploadbalancer.v1.RouteOptionsR\frouteOptions\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01J\x04\b\x06\x10\a\"@\n" +
	"\x18CreateHttpRouterMetadata\x12$\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tR\fhttpRouterId\"\xa8\x01\n" +
	"\x1fListHttpRouterOperationsRequest\x122\n" +
	"\x0ehttp_router_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\fhttpRouterId\x12'\n" +
	"\tpage_size\x18\x02 \x01(\x03B\n" +
	"\xfa\xc71\x06<=1000R\bpageSize\x12(\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\t\x8a\xc81\x05<=100R\tpageToken\"\x8d\x01\n" +
	" ListHttpRouterOperationsResponse\x12A\n" +
	"\n" +
	"operations\x18\x01 \x03(\v2!.yandex.cloud.operation.OperationR\n" +
	"operations\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken2\x9f\t\n" +
	"\x11HttpRouterService\x12\xa3\x01\n" +
	"\x03Get\x125.yandex.cloud.apploadbalancer.v1.GetHttpRouterRequest\x1a+.yandex.cloud.apploadbalancer.v1.HttpRouter\"8\x82\xd3\xe4\x93\x022\x120/apploadbalancer/v1/httpRouters/{http_router_id}\x12\xa2\x01\n" +
	"\x04List\x127.yandex.cloud.apploadbalancer.v1.ListHttpRoutersRequest\x1a8.yandex.cloud.apploadbalancer.v1.ListHttpRoutersResponse\"'\x82\xd3\xe4\x93\x02!\x12\x1f/apploadbalancer/v1/httpRouters\x12\xbb\x01\n" +
	"\x06Create\x128.yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest\x1a!.yandex.cloud.operation.Operation\"T\xb2\xd2*&\n" +
	"\x18CreateHttpRouterMetadata\x12\n" +
	"HttpRouter\x82\xd3\xe4\x93\x02$:\x01*\"\x1f/apploadbalancer/v1/httpRouters\x12\xcc\x01\n" +
	"\x06Update\x128.yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest\x1a!.yandex.cloud.operation.Operation\"e\xb2\xd2*&\n" +
	"\x18UpdateHttpRouterMetadata\x12\n" +
	"HttpRouter\x82\xd3\xe4\x93\x025:\x01*20/apploadbalancer/v1/httpRouters/{http_router_id}\x12\xd4\x01\n" +
	"\x06Delete\x128.yandex.cloud.apploadbalancer.v1.DeleteHttpRouterRequest\x1a!.yandex.cloud.operation.Operation\"m\xb2\xd2*1\n" +
	"\x18DeleteHttpRouterMetadata\x12\x15google.protobuf.Empty\x82\xd3\xe4\x93\x022*0/apploadbalancer/v1/httpRouters/{http_router_id}\x12\xda\x01\n" +
	"\x0eListOperations\x12@.yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsRequest\x1aA.yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsResponse\"C\x82\xd3\xe4\x93\x02=\x12;/apploadbalancer/v1/httpRouters/{http_router_id}/operationsBz\n" +
	"#yandex.cloud.api.apploadbalancer.v1ZSgithub.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1;apploadbalancerb\x06proto3"

var (
	file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescData []byte
)

func file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDesc)))
	})
	return file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDescData
}

var file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_goTypes = []any{
	(*GetHttpRouterRequest)(nil),             // 0: yandex.cloud.apploadbalancer.v1.GetHttpRouterRequest
	(*ListHttpRoutersRequest)(nil),           // 1: yandex.cloud.apploadbalancer.v1.ListHttpRoutersRequest
	(*ListHttpRoutersResponse)(nil),          // 2: yandex.cloud.apploadbalancer.v1.ListHttpRoutersResponse
	(*DeleteHttpRouterRequest)(nil),          // 3: yandex.cloud.apploadbalancer.v1.DeleteHttpRouterRequest
	(*DeleteHttpRouterMetadata)(nil),         // 4: yandex.cloud.apploadbalancer.v1.DeleteHttpRouterMetadata
	(*UpdateHttpRouterRequest)(nil),          // 5: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest
	(*UpdateHttpRouterMetadata)(nil),         // 6: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterMetadata
	(*CreateHttpRouterRequest)(nil),          // 7: yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest
	(*CreateHttpRouterMetadata)(nil),         // 8: yandex.cloud.apploadbalancer.v1.CreateHttpRouterMetadata
	(*ListHttpRouterOperationsRequest)(nil),  // 9: yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsRequest
	(*ListHttpRouterOperationsResponse)(nil), // 10: yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsResponse
	nil,                                      // 11: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.LabelsEntry
	nil,                                      // 12: yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest.LabelsEntry
	(*HttpRouter)(nil),                       // 13: yandex.cloud.apploadbalancer.v1.HttpRouter
	(*fieldmaskpb.FieldMask)(nil),            // 14: google.protobuf.FieldMask
	(*VirtualHost)(nil),                      // 15: yandex.cloud.apploadbalancer.v1.VirtualHost
	(*RouteOptions)(nil),                     // 16: yandex.cloud.apploadbalancer.v1.RouteOptions
	(*operation.Operation)(nil),              // 17: yandex.cloud.operation.Operation
}
var file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_depIdxs = []int32{
	13, // 0: yandex.cloud.apploadbalancer.v1.ListHttpRoutersResponse.http_routers:type_name -> yandex.cloud.apploadbalancer.v1.HttpRouter
	14, // 1: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.update_mask:type_name -> google.protobuf.FieldMask
	11, // 2: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.labels:type_name -> yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.LabelsEntry
	15, // 3: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.virtual_hosts:type_name -> yandex.cloud.apploadbalancer.v1.VirtualHost
	16, // 4: yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest.route_options:type_name -> yandex.cloud.apploadbalancer.v1.RouteOptions
	12, // 5: yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest.labels:type_name -> yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest.LabelsEntry
	15, // 6: yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest.virtual_hosts:type_name -> yandex.cloud.apploadbalancer.v1.VirtualHost
	16, // 7: yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest.route_options:type_name -> yandex.cloud.apploadbalancer.v1.RouteOptions
	17, // 8: yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsResponse.operations:type_name -> yandex.cloud.operation.Operation
	0,  // 9: yandex.cloud.apploadbalancer.v1.HttpRouterService.Get:input_type -> yandex.cloud.apploadbalancer.v1.GetHttpRouterRequest
	1,  // 10: yandex.cloud.apploadbalancer.v1.HttpRouterService.List:input_type -> yandex.cloud.apploadbalancer.v1.ListHttpRoutersRequest
	7,  // 11: yandex.cloud.apploadbalancer.v1.HttpRouterService.Create:input_type -> yandex.cloud.apploadbalancer.v1.CreateHttpRouterRequest
	5,  // 12: yandex.cloud.apploadbalancer.v1.HttpRouterService.Update:input_type -> yandex.cloud.apploadbalancer.v1.UpdateHttpRouterRequest
	3,  // 13: yandex.cloud.apploadbalancer.v1.HttpRouterService.Delete:input_type -> yandex.cloud.apploadbalancer.v1.DeleteHttpRouterRequest
	9,  // 14: yandex.cloud.apploadbalancer.v1.HttpRouterService.ListOperations:input_type -> yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsRequest
	13, // 15: yandex.cloud.apploadbalancer.v1.HttpRouterService.Get:output_type -> yandex.cloud.apploadbalancer.v1.HttpRouter
	2,  // 16: yandex.cloud.apploadbalancer.v1.HttpRouterService.List:output_type -> yandex.cloud.apploadbalancer.v1.ListHttpRoutersResponse
	17, // 17: yandex.cloud.apploadbalancer.v1.HttpRouterService.Create:output_type -> yandex.cloud.operation.Operation
	17, // 18: yandex.cloud.apploadbalancer.v1.HttpRouterService.Update:output_type -> yandex.cloud.operation.Operation
	17, // 19: yandex.cloud.apploadbalancer.v1.HttpRouterService.Delete:output_type -> yandex.cloud.operation.Operation
	10, // 20: yandex.cloud.apploadbalancer.v1.HttpRouterService.ListOperations:output_type -> yandex.cloud.apploadbalancer.v1.ListHttpRouterOperationsResponse
	15, // [15:21] is the sub-list for method output_type
	9,  // [9:15] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_init() }
func file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_init() {
	if File_yandex_cloud_apploadbalancer_v1_http_router_service_proto != nil {
		return
	}
	file_yandex_cloud_apploadbalancer_v1_http_router_proto_init()
	file_yandex_cloud_apploadbalancer_v1_virtual_host_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_apploadbalancer_v1_http_router_service_proto = out.File
	file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_http_router_service_proto_depIdxs = nil
}
