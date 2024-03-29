// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/marketplace/metering/v1/image_product_usage_service.proto

package metering

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	ImageProductUsageService_Write_FullMethodName = "/yandex.cloud.marketplace.metering.v1.ImageProductUsageService/Write"
)

// ImageProductUsageServiceClient is the client API for ImageProductUsageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageProductUsageServiceClient interface {
	// Writes image product's usage. Authentication is by user's service account.
	Write(ctx context.Context, in *WriteImageProductUsageRequest, opts ...grpc.CallOption) (*WriteImageProductUsageResponse, error)
}

type imageProductUsageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageProductUsageServiceClient(cc grpc.ClientConnInterface) ImageProductUsageServiceClient {
	return &imageProductUsageServiceClient{cc}
}

func (c *imageProductUsageServiceClient) Write(ctx context.Context, in *WriteImageProductUsageRequest, opts ...grpc.CallOption) (*WriteImageProductUsageResponse, error) {
	out := new(WriteImageProductUsageResponse)
	err := c.cc.Invoke(ctx, ImageProductUsageService_Write_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageProductUsageServiceServer is the server API for ImageProductUsageService service.
// All implementations should embed UnimplementedImageProductUsageServiceServer
// for forward compatibility
type ImageProductUsageServiceServer interface {
	// Writes image product's usage. Authentication is by user's service account.
	Write(context.Context, *WriteImageProductUsageRequest) (*WriteImageProductUsageResponse, error)
}

// UnimplementedImageProductUsageServiceServer should be embedded to have forward compatible implementations.
type UnimplementedImageProductUsageServiceServer struct {
}

func (UnimplementedImageProductUsageServiceServer) Write(context.Context, *WriteImageProductUsageRequest) (*WriteImageProductUsageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}

// UnsafeImageProductUsageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageProductUsageServiceServer will
// result in compilation errors.
type UnsafeImageProductUsageServiceServer interface {
	mustEmbedUnimplementedImageProductUsageServiceServer()
}

func RegisterImageProductUsageServiceServer(s grpc.ServiceRegistrar, srv ImageProductUsageServiceServer) {
	s.RegisterService(&ImageProductUsageService_ServiceDesc, srv)
}

func _ImageProductUsageService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteImageProductUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageProductUsageServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageProductUsageService_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageProductUsageServiceServer).Write(ctx, req.(*WriteImageProductUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageProductUsageService_ServiceDesc is the grpc.ServiceDesc for ImageProductUsageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageProductUsageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.marketplace.metering.v1.ImageProductUsageService",
	HandlerType: (*ImageProductUsageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _ImageProductUsageService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/marketplace/metering/v1/image_product_usage_service.proto",
}
