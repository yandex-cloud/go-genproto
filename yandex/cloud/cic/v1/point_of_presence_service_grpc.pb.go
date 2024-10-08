// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/cic/v1/point_of_presence_service.proto

package cic

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
	PointOfPresenceService_Get_FullMethodName  = "/yandex.cloud.cic.v1.PointOfPresenceService/Get"
	PointOfPresenceService_List_FullMethodName = "/yandex.cloud.cic.v1.PointOfPresenceService/List"
)

// PointOfPresenceServiceClient is the client API for PointOfPresenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing PointOfPresence resources.
type PointOfPresenceServiceClient interface {
	// Returns the specified PointOfPresence resource.
	//
	// To get the list of available PointOfPresence resources, make a [List] request.
	Get(ctx context.Context, in *GetPointOfPresenceRequest, opts ...grpc.CallOption) (*PointOfPresence, error)
	// Retrieves the list of PointOfPresence resources in the specified folder.
	List(ctx context.Context, in *ListPointOfPresencesRequest, opts ...grpc.CallOption) (*ListPointOfPresencesResponse, error)
}

type pointOfPresenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPointOfPresenceServiceClient(cc grpc.ClientConnInterface) PointOfPresenceServiceClient {
	return &pointOfPresenceServiceClient{cc}
}

func (c *pointOfPresenceServiceClient) Get(ctx context.Context, in *GetPointOfPresenceRequest, opts ...grpc.CallOption) (*PointOfPresence, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PointOfPresence)
	err := c.cc.Invoke(ctx, PointOfPresenceService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pointOfPresenceServiceClient) List(ctx context.Context, in *ListPointOfPresencesRequest, opts ...grpc.CallOption) (*ListPointOfPresencesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPointOfPresencesResponse)
	err := c.cc.Invoke(ctx, PointOfPresenceService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PointOfPresenceServiceServer is the server API for PointOfPresenceService service.
// All implementations should embed UnimplementedPointOfPresenceServiceServer
// for forward compatibility.
//
// A set of methods for managing PointOfPresence resources.
type PointOfPresenceServiceServer interface {
	// Returns the specified PointOfPresence resource.
	//
	// To get the list of available PointOfPresence resources, make a [List] request.
	Get(context.Context, *GetPointOfPresenceRequest) (*PointOfPresence, error)
	// Retrieves the list of PointOfPresence resources in the specified folder.
	List(context.Context, *ListPointOfPresencesRequest) (*ListPointOfPresencesResponse, error)
}

// UnimplementedPointOfPresenceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPointOfPresenceServiceServer struct{}

func (UnimplementedPointOfPresenceServiceServer) Get(context.Context, *GetPointOfPresenceRequest) (*PointOfPresence, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPointOfPresenceServiceServer) List(context.Context, *ListPointOfPresencesRequest) (*ListPointOfPresencesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPointOfPresenceServiceServer) testEmbeddedByValue() {}

// UnsafePointOfPresenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PointOfPresenceServiceServer will
// result in compilation errors.
type UnsafePointOfPresenceServiceServer interface {
	mustEmbedUnimplementedPointOfPresenceServiceServer()
}

func RegisterPointOfPresenceServiceServer(s grpc.ServiceRegistrar, srv PointOfPresenceServiceServer) {
	// If the following call pancis, it indicates UnimplementedPointOfPresenceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PointOfPresenceService_ServiceDesc, srv)
}

func _PointOfPresenceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPointOfPresenceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointOfPresenceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PointOfPresenceService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointOfPresenceServiceServer).Get(ctx, req.(*GetPointOfPresenceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PointOfPresenceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPointOfPresencesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PointOfPresenceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PointOfPresenceService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PointOfPresenceServiceServer).List(ctx, req.(*ListPointOfPresencesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PointOfPresenceService_ServiceDesc is the grpc.ServiceDesc for PointOfPresenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PointOfPresenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.cic.v1.PointOfPresenceService",
	HandlerType: (*PointOfPresenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PointOfPresenceService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PointOfPresenceService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/cic/v1/point_of_presence_service.proto",
}
