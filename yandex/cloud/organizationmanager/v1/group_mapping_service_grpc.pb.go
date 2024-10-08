// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/organizationmanager/v1/group_mapping_service.proto

package organizationmanager

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GroupMappingService_Get_FullMethodName         = "/yandex.cloud.organizationmanager.v1.GroupMappingService/Get"
	GroupMappingService_Create_FullMethodName      = "/yandex.cloud.organizationmanager.v1.GroupMappingService/Create"
	GroupMappingService_Update_FullMethodName      = "/yandex.cloud.organizationmanager.v1.GroupMappingService/Update"
	GroupMappingService_Delete_FullMethodName      = "/yandex.cloud.organizationmanager.v1.GroupMappingService/Delete"
	GroupMappingService_ListItems_FullMethodName   = "/yandex.cloud.organizationmanager.v1.GroupMappingService/ListItems"
	GroupMappingService_UpdateItems_FullMethodName = "/yandex.cloud.organizationmanager.v1.GroupMappingService/UpdateItems"
)

// GroupMappingServiceClient is the client API for GroupMappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RPC service dedicated for federation group mapping.
type GroupMappingServiceClient interface {
	// Returns a group mapping configured for the specific federation
	// If a federation does not exist this call will return an error
	//
	//	NOT_FOUND will be returned
	//
	// If a federation exist, but has not ever been configured for group mapping
	//
	//	the call FAILED_PRECONDITION will be returned.
	Get(ctx context.Context, in *GetGroupMappingRequest, opts ...grpc.CallOption) (*GetGroupMappingResponse, error)
	// Adds a group mapping for a federation
	// If mapping already exist, ALREADY_EXISTS will be returned
	Create(ctx context.Context, in *CreateGroupMappingRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates an existing group mapping for a federation
	// Errors:
	// - if federation is not found
	// In case of any error, no changes are applied to existing group mapping
	//
	// This call is idempotent. The following actions do nothing:
	// - enabling when already enabled
	// - disabling when disabled
	// Such parts of request will be ignored. Others will be applied.
	Update(ctx context.Context, in *UpdateGroupMappingRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes a group mapping. This will remove all the mapping items
	// cascade.
	Delete(ctx context.Context, in *DeleteGroupMappingRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns all the group mappings items
	//
	// Filtering is only supported by external_group_id or internal_group_id
	ListItems(ctx context.Context, in *ListGroupMappingItemsRequest, opts ...grpc.CallOption) (*ListGroupMappingItemsResponse, error)
	// Updates group mapping items for a specified federation
	// Errors:
	// - if federation is not found
	// - if internal group in the mapping added does not exist
	// In case of any error, no changes are applied to existing group mapping
	//
	// This call is idempotent. The following actions do nothing:
	// - adding group mapping items that are already present
	// - removing group mapping items that are not present
	// Such parts of request will be ignored. Others will be applied.
	UpdateItems(ctx context.Context, in *UpdateGroupMappingItemsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type groupMappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupMappingServiceClient(cc grpc.ClientConnInterface) GroupMappingServiceClient {
	return &groupMappingServiceClient{cc}
}

func (c *groupMappingServiceClient) Get(ctx context.Context, in *GetGroupMappingRequest, opts ...grpc.CallOption) (*GetGroupMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetGroupMappingResponse)
	err := c.cc.Invoke(ctx, GroupMappingService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMappingServiceClient) Create(ctx context.Context, in *CreateGroupMappingRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GroupMappingService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMappingServiceClient) Update(ctx context.Context, in *UpdateGroupMappingRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GroupMappingService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMappingServiceClient) Delete(ctx context.Context, in *DeleteGroupMappingRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GroupMappingService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMappingServiceClient) ListItems(ctx context.Context, in *ListGroupMappingItemsRequest, opts ...grpc.CallOption) (*ListGroupMappingItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListGroupMappingItemsResponse)
	err := c.cc.Invoke(ctx, GroupMappingService_ListItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupMappingServiceClient) UpdateItems(ctx context.Context, in *UpdateGroupMappingItemsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GroupMappingService_UpdateItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupMappingServiceServer is the server API for GroupMappingService service.
// All implementations should embed UnimplementedGroupMappingServiceServer
// for forward compatibility.
//
// RPC service dedicated for federation group mapping.
type GroupMappingServiceServer interface {
	// Returns a group mapping configured for the specific federation
	// If a federation does not exist this call will return an error
	//
	//	NOT_FOUND will be returned
	//
	// If a federation exist, but has not ever been configured for group mapping
	//
	//	the call FAILED_PRECONDITION will be returned.
	Get(context.Context, *GetGroupMappingRequest) (*GetGroupMappingResponse, error)
	// Adds a group mapping for a federation
	// If mapping already exist, ALREADY_EXISTS will be returned
	Create(context.Context, *CreateGroupMappingRequest) (*operation.Operation, error)
	// Updates an existing group mapping for a federation
	// Errors:
	// - if federation is not found
	// In case of any error, no changes are applied to existing group mapping
	//
	// This call is idempotent. The following actions do nothing:
	// - enabling when already enabled
	// - disabling when disabled
	// Such parts of request will be ignored. Others will be applied.
	Update(context.Context, *UpdateGroupMappingRequest) (*operation.Operation, error)
	// Deletes a group mapping. This will remove all the mapping items
	// cascade.
	Delete(context.Context, *DeleteGroupMappingRequest) (*operation.Operation, error)
	// Returns all the group mappings items
	//
	// Filtering is only supported by external_group_id or internal_group_id
	ListItems(context.Context, *ListGroupMappingItemsRequest) (*ListGroupMappingItemsResponse, error)
	// Updates group mapping items for a specified federation
	// Errors:
	// - if federation is not found
	// - if internal group in the mapping added does not exist
	// In case of any error, no changes are applied to existing group mapping
	//
	// This call is idempotent. The following actions do nothing:
	// - adding group mapping items that are already present
	// - removing group mapping items that are not present
	// Such parts of request will be ignored. Others will be applied.
	UpdateItems(context.Context, *UpdateGroupMappingItemsRequest) (*operation.Operation, error)
}

// UnimplementedGroupMappingServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGroupMappingServiceServer struct{}

func (UnimplementedGroupMappingServiceServer) Get(context.Context, *GetGroupMappingRequest) (*GetGroupMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGroupMappingServiceServer) Create(context.Context, *CreateGroupMappingRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGroupMappingServiceServer) Update(context.Context, *UpdateGroupMappingRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGroupMappingServiceServer) Delete(context.Context, *DeleteGroupMappingRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGroupMappingServiceServer) ListItems(context.Context, *ListGroupMappingItemsRequest) (*ListGroupMappingItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListItems not implemented")
}
func (UnimplementedGroupMappingServiceServer) UpdateItems(context.Context, *UpdateGroupMappingItemsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItems not implemented")
}
func (UnimplementedGroupMappingServiceServer) testEmbeddedByValue() {}

// UnsafeGroupMappingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupMappingServiceServer will
// result in compilation errors.
type UnsafeGroupMappingServiceServer interface {
	mustEmbedUnimplementedGroupMappingServiceServer()
}

func RegisterGroupMappingServiceServer(s grpc.ServiceRegistrar, srv GroupMappingServiceServer) {
	// If the following call pancis, it indicates UnimplementedGroupMappingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GroupMappingService_ServiceDesc, srv)
}

func _GroupMappingService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMappingServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMappingService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMappingServiceServer).Get(ctx, req.(*GetGroupMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMappingService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMappingServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMappingService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMappingServiceServer).Create(ctx, req.(*CreateGroupMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMappingService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMappingServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMappingService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMappingServiceServer).Update(ctx, req.(*UpdateGroupMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMappingService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMappingServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMappingService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMappingServiceServer).Delete(ctx, req.(*DeleteGroupMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMappingService_ListItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGroupMappingItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMappingServiceServer).ListItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMappingService_ListItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMappingServiceServer).ListItems(ctx, req.(*ListGroupMappingItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GroupMappingService_UpdateItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMappingItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupMappingServiceServer).UpdateItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GroupMappingService_UpdateItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupMappingServiceServer).UpdateItems(ctx, req.(*UpdateGroupMappingItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GroupMappingService_ServiceDesc is the grpc.ServiceDesc for GroupMappingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GroupMappingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.organizationmanager.v1.GroupMappingService",
	HandlerType: (*GroupMappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GroupMappingService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GroupMappingService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GroupMappingService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GroupMappingService_Delete_Handler,
		},
		{
			MethodName: "ListItems",
			Handler:    _GroupMappingService_ListItems_Handler,
		},
		{
			MethodName: "UpdateItems",
			Handler:    _GroupMappingService_UpdateItems_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/organizationmanager/v1/group_mapping_service.proto",
}
