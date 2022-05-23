// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: yandex/cloud/ai/vision/v2/image_classifier_service.proto

package vision

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

// ImageClassifierServiceClient is the client API for ImageClassifierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClassifierServiceClient interface {
	Annotate(ctx context.Context, in *AnnotationRequest, opts ...grpc.CallOption) (*AnnotationResponse, error)
}

type imageClassifierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClassifierServiceClient(cc grpc.ClientConnInterface) ImageClassifierServiceClient {
	return &imageClassifierServiceClient{cc}
}

func (c *imageClassifierServiceClient) Annotate(ctx context.Context, in *AnnotationRequest, opts ...grpc.CallOption) (*AnnotationResponse, error) {
	out := new(AnnotationResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ai.vision.v2.ImageClassifierService/Annotate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageClassifierServiceServer is the server API for ImageClassifierService service.
// All implementations should embed UnimplementedImageClassifierServiceServer
// for forward compatibility
type ImageClassifierServiceServer interface {
	Annotate(context.Context, *AnnotationRequest) (*AnnotationResponse, error)
}

// UnimplementedImageClassifierServiceServer should be embedded to have forward compatible implementations.
type UnimplementedImageClassifierServiceServer struct {
}

func (UnimplementedImageClassifierServiceServer) Annotate(context.Context, *AnnotationRequest) (*AnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Annotate not implemented")
}

// UnsafeImageClassifierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageClassifierServiceServer will
// result in compilation errors.
type UnsafeImageClassifierServiceServer interface {
	mustEmbedUnimplementedImageClassifierServiceServer()
}

func RegisterImageClassifierServiceServer(s grpc.ServiceRegistrar, srv ImageClassifierServiceServer) {
	s.RegisterService(&ImageClassifierService_ServiceDesc, srv)
}

func _ImageClassifierService_Annotate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageClassifierServiceServer).Annotate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ai.vision.v2.ImageClassifierService/Annotate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageClassifierServiceServer).Annotate(ctx, req.(*AnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageClassifierService_ServiceDesc is the grpc.ServiceDesc for ImageClassifierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageClassifierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.vision.v2.ImageClassifierService",
	HandlerType: (*ImageClassifierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Annotate",
			Handler:    _ImageClassifierService_Annotate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/vision/v2/image_classifier_service.proto",
}
