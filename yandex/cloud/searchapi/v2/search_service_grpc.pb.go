// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/searchapi/v2/search_service.proto

package searchapi

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
	WebSearchAsyncService_Search_FullMethodName = "/yandex.cloud.searchapi.v2.WebSearchAsyncService/Search"
)

// WebSearchAsyncServiceClient is the client API for WebSearchAsyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for async search the Yandex search database.
type WebSearchAsyncServiceClient interface {
	Search(ctx context.Context, in *WebSearchRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type webSearchAsyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWebSearchAsyncServiceClient(cc grpc.ClientConnInterface) WebSearchAsyncServiceClient {
	return &webSearchAsyncServiceClient{cc}
}

func (c *webSearchAsyncServiceClient) Search(ctx context.Context, in *WebSearchRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, WebSearchAsyncService_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WebSearchAsyncServiceServer is the server API for WebSearchAsyncService service.
// All implementations should embed UnimplementedWebSearchAsyncServiceServer
// for forward compatibility.
//
// A set of methods for async search the Yandex search database.
type WebSearchAsyncServiceServer interface {
	Search(context.Context, *WebSearchRequest) (*operation.Operation, error)
}

// UnimplementedWebSearchAsyncServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedWebSearchAsyncServiceServer struct{}

func (UnimplementedWebSearchAsyncServiceServer) Search(context.Context, *WebSearchRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedWebSearchAsyncServiceServer) testEmbeddedByValue() {}

// UnsafeWebSearchAsyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WebSearchAsyncServiceServer will
// result in compilation errors.
type UnsafeWebSearchAsyncServiceServer interface {
	mustEmbedUnimplementedWebSearchAsyncServiceServer()
}

func RegisterWebSearchAsyncServiceServer(s grpc.ServiceRegistrar, srv WebSearchAsyncServiceServer) {
	// If the following call pancis, it indicates UnimplementedWebSearchAsyncServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&WebSearchAsyncService_ServiceDesc, srv)
}

func _WebSearchAsyncService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebSearchAsyncServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: WebSearchAsyncService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebSearchAsyncServiceServer).Search(ctx, req.(*WebSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WebSearchAsyncService_ServiceDesc is the grpc.ServiceDesc for WebSearchAsyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WebSearchAsyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.searchapi.v2.WebSearchAsyncService",
	HandlerType: (*WebSearchAsyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _WebSearchAsyncService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/searchapi/v2/search_service.proto",
}
