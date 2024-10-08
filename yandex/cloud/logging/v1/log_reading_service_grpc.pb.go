// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/logging/v1/log_reading_service.proto

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
	LogReadingService_Read_FullMethodName = "/yandex.cloud.logging.v1.LogReadingService/Read"
)

// LogReadingServiceClient is the client API for LogReadingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for reading from log groups.
type LogReadingServiceClient interface {
	// Read log entries from the specified log group.
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
}

type logReadingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLogReadingServiceClient(cc grpc.ClientConnInterface) LogReadingServiceClient {
	return &logReadingServiceClient{cc}
}

func (c *logReadingServiceClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, LogReadingService_Read_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogReadingServiceServer is the server API for LogReadingService service.
// All implementations should embed UnimplementedLogReadingServiceServer
// for forward compatibility.
//
// A set of methods for reading from log groups.
type LogReadingServiceServer interface {
	// Read log entries from the specified log group.
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
}

// UnimplementedLogReadingServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLogReadingServiceServer struct{}

func (UnimplementedLogReadingServiceServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedLogReadingServiceServer) testEmbeddedByValue() {}

// UnsafeLogReadingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogReadingServiceServer will
// result in compilation errors.
type UnsafeLogReadingServiceServer interface {
	mustEmbedUnimplementedLogReadingServiceServer()
}

func RegisterLogReadingServiceServer(s grpc.ServiceRegistrar, srv LogReadingServiceServer) {
	// If the following call pancis, it indicates UnimplementedLogReadingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LogReadingService_ServiceDesc, srv)
}

func _LogReadingService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogReadingServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LogReadingService_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogReadingServiceServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LogReadingService_ServiceDesc is the grpc.ServiceDesc for LogReadingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogReadingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.logging.v1.LogReadingService",
	HandlerType: (*LogReadingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _LogReadingService_Read_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/logging/v1/log_reading_service.proto",
}
