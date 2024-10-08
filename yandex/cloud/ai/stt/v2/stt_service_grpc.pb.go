// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/ai/stt/v2/stt_service.proto

package stt

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
	SttService_LongRunningRecognize_FullMethodName = "/yandex.cloud.ai.stt.v2.SttService/LongRunningRecognize"
	SttService_StreamingRecognize_FullMethodName   = "/yandex.cloud.ai.stt.v2.SttService/StreamingRecognize"
)

// SttServiceClient is the client API for SttService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SttServiceClient interface {
	LongRunningRecognize(ctx context.Context, in *LongRunningRecognitionRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	StreamingRecognize(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingRecognitionRequest, StreamingRecognitionResponse], error)
}

type sttServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSttServiceClient(cc grpc.ClientConnInterface) SttServiceClient {
	return &sttServiceClient{cc}
}

func (c *sttServiceClient) LongRunningRecognize(ctx context.Context, in *LongRunningRecognitionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SttService_LongRunningRecognize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sttServiceClient) StreamingRecognize(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[StreamingRecognitionRequest, StreamingRecognitionResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &SttService_ServiceDesc.Streams[0], SttService_StreamingRecognize_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[StreamingRecognitionRequest, StreamingRecognitionResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SttService_StreamingRecognizeClient = grpc.BidiStreamingClient[StreamingRecognitionRequest, StreamingRecognitionResponse]

// SttServiceServer is the server API for SttService service.
// All implementations should embed UnimplementedSttServiceServer
// for forward compatibility.
type SttServiceServer interface {
	LongRunningRecognize(context.Context, *LongRunningRecognitionRequest) (*operation.Operation, error)
	StreamingRecognize(grpc.BidiStreamingServer[StreamingRecognitionRequest, StreamingRecognitionResponse]) error
}

// UnimplementedSttServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSttServiceServer struct{}

func (UnimplementedSttServiceServer) LongRunningRecognize(context.Context, *LongRunningRecognitionRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LongRunningRecognize not implemented")
}
func (UnimplementedSttServiceServer) StreamingRecognize(grpc.BidiStreamingServer[StreamingRecognitionRequest, StreamingRecognitionResponse]) error {
	return status.Errorf(codes.Unimplemented, "method StreamingRecognize not implemented")
}
func (UnimplementedSttServiceServer) testEmbeddedByValue() {}

// UnsafeSttServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SttServiceServer will
// result in compilation errors.
type UnsafeSttServiceServer interface {
	mustEmbedUnimplementedSttServiceServer()
}

func RegisterSttServiceServer(s grpc.ServiceRegistrar, srv SttServiceServer) {
	// If the following call pancis, it indicates UnimplementedSttServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SttService_ServiceDesc, srv)
}

func _SttService_LongRunningRecognize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LongRunningRecognitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SttServiceServer).LongRunningRecognize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SttService_LongRunningRecognize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SttServiceServer).LongRunningRecognize(ctx, req.(*LongRunningRecognitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SttService_StreamingRecognize_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SttServiceServer).StreamingRecognize(&grpc.GenericServerStream[StreamingRecognitionRequest, StreamingRecognitionResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type SttService_StreamingRecognizeServer = grpc.BidiStreamingServer[StreamingRecognitionRequest, StreamingRecognitionResponse]

// SttService_ServiceDesc is the grpc.ServiceDesc for SttService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SttService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.stt.v2.SttService",
	HandlerType: (*SttServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LongRunningRecognize",
			Handler:    _SttService_LongRunningRecognize_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingRecognize",
			Handler:       _SttService_StreamingRecognize_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "yandex/cloud/ai/stt/v2/stt_service.proto",
}
