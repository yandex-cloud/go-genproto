// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/marketplace/metering/v1/product_usage_service.proto

package metering

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
	ProductUsageService_Write_FullMethodName = "/yandex.cloud.marketplace.metering.v1.ProductUsageService/Write"
)

// ProductUsageServiceClient is the client API for ProductUsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing product's usage with product instances.
type ProductUsageServiceClient interface {
	// Writes product's usage (authenticated by publisher's service account)
	Write(ctx context.Context, in *WriteUsageRequest, opts ...grpc.CallOption) (*WriteUsageResponse, error)
}

type productUsageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductUsageServiceClient(cc grpc.ClientConnInterface) ProductUsageServiceClient {
	return &productUsageServiceClient{cc}
}

func (c *productUsageServiceClient) Write(ctx context.Context, in *WriteUsageRequest, opts ...grpc.CallOption) (*WriteUsageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteUsageResponse)
	err := c.cc.Invoke(ctx, ProductUsageService_Write_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductUsageServiceServer is the server API for ProductUsageService service.
// All implementations should embed UnimplementedProductUsageServiceServer
// for forward compatibility.
//
// A set of methods for managing product's usage with product instances.
type ProductUsageServiceServer interface {
	// Writes product's usage (authenticated by publisher's service account)
	Write(context.Context, *WriteUsageRequest) (*WriteUsageResponse, error)
}

// UnimplementedProductUsageServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProductUsageServiceServer struct{}

func (UnimplementedProductUsageServiceServer) Write(context.Context, *WriteUsageRequest) (*WriteUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedProductUsageServiceServer) testEmbeddedByValue() {}

// UnsafeProductUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductUsageServiceServer will
// result in compilation errors.
type UnsafeProductUsageServiceServer interface {
	mustEmbedUnimplementedProductUsageServiceServer()
}

func RegisterProductUsageServiceServer(s grpc.ServiceRegistrar, srv ProductUsageServiceServer) {
	// If the following call pancis, it indicates UnimplementedProductUsageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProductUsageService_ServiceDesc, srv)
}

func _ProductUsageService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductUsageServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProductUsageService_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductUsageServiceServer).Write(ctx, req.(*WriteUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductUsageService_ServiceDesc is the grpc.ServiceDesc for ProductUsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductUsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.marketplace.metering.v1.ProductUsageService",
	HandlerType: (*ProductUsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _ProductUsageService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/marketplace/metering/v1/product_usage_service.proto",
}
