// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/logging/v1/log_ingestion_service.proto

package logging

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
	LogIngestionService_Write_FullMethodName = "/yandex.cloud.logging.v1.LogIngestionService/Write"
)

// LogIngestionServiceClient is the client API for LogIngestionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for writing to log groups.
type LogIngestionServiceClient interface {
	// Write log entries to specified destination.
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
}

type logIngestionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogIngestionServiceClient(cc grpc.ClientConnInterface) LogIngestionServiceClient {
	return &logIngestionServiceClient{cc}
}

func (c *logIngestionServiceClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, LogIngestionService_Write_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogIngestionServiceServer is the server API for LogIngestionService service.
// All implementations should embed UnimplementedLogIngestionServiceServer
// for forward compatibility.
//
// A set of methods for writing to log groups.
type LogIngestionServiceServer interface {
	// Write log entries to specified destination.
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
}

// UnimplementedLogIngestionServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogIngestionServiceServer struct{}

func (UnimplementedLogIngestionServiceServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedLogIngestionServiceServer) testEmbeddedByValue() {}

// UnsafeLogIngestionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogIngestionServiceServer will
// result in compilation errors.
type UnsafeLogIngestionServiceServer interface {
	mustEmbedUnimplementedLogIngestionServiceServer()
}

func RegisterLogIngestionServiceServer(s grpc.ServiceRegistrar, srv LogIngestionServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogIngestionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogIngestionService_ServiceDesc, srv)
}

func _LogIngestionService_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogIngestionServiceServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogIngestionService_Write_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogIngestionServiceServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogIngestionService_ServiceDesc is the grpc.ServiceDesc for LogIngestionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogIngestionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.logging.v1.LogIngestionService",
	HandlerType: (*LogIngestionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _LogIngestionService_Write_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/logging/v1/log_ingestion_service.proto",
}
