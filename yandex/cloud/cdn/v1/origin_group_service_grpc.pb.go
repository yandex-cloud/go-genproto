// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/cdn/v1/origin_group_service.proto

package cdn

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
	OriginGroupService_Get_FullMethodName    = "/yandex.cloud.cdn.v1.OriginGroupService/Get"
	OriginGroupService_List_FullMethodName   = "/yandex.cloud.cdn.v1.OriginGroupService/List"
	OriginGroupService_Create_FullMethodName = "/yandex.cloud.cdn.v1.OriginGroupService/Create"
	OriginGroupService_Update_FullMethodName = "/yandex.cloud.cdn.v1.OriginGroupService/Update"
	OriginGroupService_Delete_FullMethodName = "/yandex.cloud.cdn.v1.OriginGroupService/Delete"
)

// OriginGroupServiceClient is the client API for OriginGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Origin Groups management service.
type OriginGroupServiceClient interface {
	// Gets origin group with specified origin group id.
	Get(ctx context.Context, in *GetOriginGroupRequest, opts ...grpc.CallOption) (*OriginGroup, error)
	// Lists origins of origin group.
	List(ctx context.Context, in *ListOriginGroupsRequest, opts ...grpc.CallOption) (*ListOriginGroupsResponse, error)
	// Creates origin group.
	Create(ctx context.Context, in *CreateOriginGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified origin group.
	//
	// Changes may take up to 15 minutes to apply. Afterwards, it is recommended to purge cache of the resources that
	// use the origin group via a [CacheService.Purge] request.
	Update(ctx context.Context, in *UpdateOriginGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes origin group with specified origin group id.
	Delete(ctx context.Context, in *DeleteOriginGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type originGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOriginGroupServiceClient(cc grpc.ClientConnInterface) OriginGroupServiceClient {
	return &originGroupServiceClient{cc}
}

func (c *originGroupServiceClient) Get(ctx context.Context, in *GetOriginGroupRequest, opts ...grpc.CallOption) (*OriginGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OriginGroup)
	err := c.cc.Invoke(ctx, OriginGroupService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originGroupServiceClient) List(ctx context.Context, in *ListOriginGroupsRequest, opts ...grpc.CallOption) (*ListOriginGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOriginGroupsResponse)
	err := c.cc.Invoke(ctx, OriginGroupService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originGroupServiceClient) Create(ctx context.Context, in *CreateOriginGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, OriginGroupService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originGroupServiceClient) Update(ctx context.Context, in *UpdateOriginGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, OriginGroupService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *originGroupServiceClient) Delete(ctx context.Context, in *DeleteOriginGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, OriginGroupService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OriginGroupServiceServer is the server API for OriginGroupService service.
// All implementations should embed UnimplementedOriginGroupServiceServer
// for forward compatibility.
//
// Origin Groups management service.
type OriginGroupServiceServer interface {
	// Gets origin group with specified origin group id.
	Get(context.Context, *GetOriginGroupRequest) (*OriginGroup, error)
	// Lists origins of origin group.
	List(context.Context, *ListOriginGroupsRequest) (*ListOriginGroupsResponse, error)
	// Creates origin group.
	Create(context.Context, *CreateOriginGroupRequest) (*operation.Operation, error)
	// Updates the specified origin group.
	//
	// Changes may take up to 15 minutes to apply. Afterwards, it is recommended to purge cache of the resources that
	// use the origin group via a [CacheService.Purge] request.
	Update(context.Context, *UpdateOriginGroupRequest) (*operation.Operation, error)
	// Deletes origin group with specified origin group id.
	Delete(context.Context, *DeleteOriginGroupRequest) (*operation.Operation, error)
}

// UnimplementedOriginGroupServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOriginGroupServiceServer struct{}

func (UnimplementedOriginGroupServiceServer) Get(context.Context, *GetOriginGroupRequest) (*OriginGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedOriginGroupServiceServer) List(context.Context, *ListOriginGroupsRequest) (*ListOriginGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedOriginGroupServiceServer) Create(context.Context, *CreateOriginGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOriginGroupServiceServer) Update(context.Context, *UpdateOriginGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedOriginGroupServiceServer) Delete(context.Context, *DeleteOriginGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedOriginGroupServiceServer) testEmbeddedByValue() {}

// UnsafeOriginGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OriginGroupServiceServer will
// result in compilation errors.
type UnsafeOriginGroupServiceServer interface {
	mustEmbedUnimplementedOriginGroupServiceServer()
}

func RegisterOriginGroupServiceServer(s grpc.ServiceRegistrar, srv OriginGroupServiceServer) {
	// If the following call pancis, it indicates UnimplementedOriginGroupServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OriginGroupService_ServiceDesc, srv)
}

func _OriginGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOriginGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginGroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginGroupServiceServer).Get(ctx, req.(*GetOriginGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOriginGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginGroupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginGroupServiceServer).List(ctx, req.(*ListOriginGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOriginGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginGroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginGroupServiceServer).Create(ctx, req.(*CreateOriginGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOriginGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginGroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginGroupServiceServer).Update(ctx, req.(*UpdateOriginGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OriginGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOriginGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OriginGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OriginGroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OriginGroupServiceServer).Delete(ctx, req.(*DeleteOriginGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OriginGroupService_ServiceDesc is the grpc.ServiceDesc for OriginGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OriginGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.cdn.v1.OriginGroupService",
	HandlerType: (*OriginGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OriginGroupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _OriginGroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _OriginGroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _OriginGroupService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _OriginGroupService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/cdn/v1/origin_group_service.proto",
}
