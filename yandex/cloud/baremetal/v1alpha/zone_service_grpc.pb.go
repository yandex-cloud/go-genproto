// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/baremetal/v1alpha/zone_service.proto

package baremetal

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
	ZoneService_Get_FullMethodName  = "/yandex.cloud.baremetal.v1alpha.ZoneService/Get"
	ZoneService_List_FullMethodName = "/yandex.cloud.baremetal.v1alpha.ZoneService/List"
)

// ZoneServiceClient is the client API for ZoneService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods to retrieve information about availability zones.
type ZoneServiceClient interface {
	// Returns the specific Zone resource.
	//
	// To get the list of Zone resources, make a [List] request.
	Get(ctx context.Context, in *GetZoneRequest, opts ...grpc.CallOption) (*Zone, error)
	// Retrieves the list of Zone resources.
	List(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error)
}

type zoneServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewZoneServiceClient(cc grpc.ClientConnInterface) ZoneServiceClient {
	return &zoneServiceClient{cc}
}

func (c *zoneServiceClient) Get(ctx context.Context, in *GetZoneRequest, opts ...grpc.CallOption) (*Zone, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Zone)
	err := c.cc.Invoke(ctx, ZoneService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zoneServiceClient) List(ctx context.Context, in *ListZonesRequest, opts ...grpc.CallOption) (*ListZonesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListZonesResponse)
	err := c.cc.Invoke(ctx, ZoneService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZoneServiceServer is the server API for ZoneService service.
// All implementations should embed UnimplementedZoneServiceServer
// for forward compatibility.
//
// A set of methods to retrieve information about availability zones.
type ZoneServiceServer interface {
	// Returns the specific Zone resource.
	//
	// To get the list of Zone resources, make a [List] request.
	Get(context.Context, *GetZoneRequest) (*Zone, error)
	// Retrieves the list of Zone resources.
	List(context.Context, *ListZonesRequest) (*ListZonesResponse, error)
}

// UnimplementedZoneServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedZoneServiceServer struct{}

func (UnimplementedZoneServiceServer) Get(context.Context, *GetZoneRequest) (*Zone, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedZoneServiceServer) List(context.Context, *ListZonesRequest) (*ListZonesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedZoneServiceServer) testEmbeddedByValue() {}

// UnsafeZoneServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ZoneServiceServer will
// result in compilation errors.
type UnsafeZoneServiceServer interface {
	mustEmbedUnimplementedZoneServiceServer()
}

func RegisterZoneServiceServer(s grpc.ServiceRegistrar, srv ZoneServiceServer) {
	// If the following call pancis, it indicates UnimplementedZoneServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ZoneService_ServiceDesc, srv)
}

func _ZoneService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetZoneRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).Get(ctx, req.(*GetZoneRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ZoneService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListZonesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZoneServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ZoneService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZoneServiceServer).List(ctx, req.(*ListZonesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ZoneService_ServiceDesc is the grpc.ServiceDesc for ZoneService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ZoneService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.baremetal.v1alpha.ZoneService",
	HandlerType: (*ZoneServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ZoneService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ZoneService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/baremetal/v1alpha/zone_service.proto",
}
