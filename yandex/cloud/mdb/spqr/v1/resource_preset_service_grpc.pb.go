// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/mdb/spqr/v1/resource_preset_service.proto

package spqr

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ResourcePresetService_Get_FullMethodName  = "/yandex.cloud.mdb.spqr.v1.ResourcePresetService/Get"
	ResourcePresetService_List_FullMethodName = "/yandex.cloud.mdb.spqr.v1.ResourcePresetService/List"
)

// ResourcePresetServiceClient is the client API for ResourcePresetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing ResourcePreset resources.
type ResourcePresetServiceClient interface {
	// Returns the specified ResourcePreset resource.
	//
	// To get the list of available ResourcePreset resources, make a [List] request.
	Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error)
	// Retrieves the list of available ResourcePreset resources.
	List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error)
}

type resourcePresetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourcePresetServiceClient(cc grpc.ClientConnInterface) ResourcePresetServiceClient {
	return &resourcePresetServiceClient{cc}
}

func (c *resourcePresetServiceClient) Get(ctx context.Context, in *GetResourcePresetRequest, opts ...grpc.CallOption) (*ResourcePreset, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourcePreset)
	err := c.cc.Invoke(ctx, ResourcePresetService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourcePresetServiceClient) List(ctx context.Context, in *ListResourcePresetsRequest, opts ...grpc.CallOption) (*ListResourcePresetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResourcePresetsResponse)
	err := c.cc.Invoke(ctx, ResourcePresetService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourcePresetServiceServer is the server API for ResourcePresetService service.
// All implementations should embed UnimplementedResourcePresetServiceServer
// for forward compatibility.
//
// A set of methods for managing ResourcePreset resources.
type ResourcePresetServiceServer interface {
	// Returns the specified ResourcePreset resource.
	//
	// To get the list of available ResourcePreset resources, make a [List] request.
	Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error)
	// Retrieves the list of available ResourcePreset resources.
	List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error)
}

// UnimplementedResourcePresetServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedResourcePresetServiceServer struct{}

func (UnimplementedResourcePresetServiceServer) Get(context.Context, *GetResourcePresetRequest) (*ResourcePreset, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedResourcePresetServiceServer) List(context.Context, *ListResourcePresetsRequest) (*ListResourcePresetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResourcePresetServiceServer) testEmbeddedByValue() {}

// UnsafeResourcePresetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourcePresetServiceServer will
// result in compilation errors.
type UnsafeResourcePresetServiceServer interface {
	mustEmbedUnimplementedResourcePresetServiceServer()
}

func RegisterResourcePresetServiceServer(s grpc.ServiceRegistrar, srv ResourcePresetServiceServer) {
	// If the following call pancis, it indicates UnimplementedResourcePresetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ResourcePresetService_ServiceDesc, srv)
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
		FullMethod: ResourcePresetService_Get_FullMethodName,
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
		FullMethod: ResourcePresetService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourcePresetServiceServer).List(ctx, req.(*ListResourcePresetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourcePresetService_ServiceDesc is the grpc.ServiceDesc for ResourcePresetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourcePresetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.spqr.v1.ResourcePresetService",
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
	Metadata: "yandex/cloud/mdb/spqr/v1/resource_preset_service.proto",
}
