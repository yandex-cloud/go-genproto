// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/video/v1/stream_service.proto

package video

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
	StreamService_Get_FullMethodName           = "/yandex.cloud.video.v1.StreamService/Get"
	StreamService_List_FullMethodName          = "/yandex.cloud.video.v1.StreamService/List"
	StreamService_BatchGet_FullMethodName      = "/yandex.cloud.video.v1.StreamService/BatchGet"
	StreamService_Create_FullMethodName        = "/yandex.cloud.video.v1.StreamService/Create"
	StreamService_Update_FullMethodName        = "/yandex.cloud.video.v1.StreamService/Update"
	StreamService_Delete_FullMethodName        = "/yandex.cloud.video.v1.StreamService/Delete"
	StreamService_BatchDelete_FullMethodName   = "/yandex.cloud.video.v1.StreamService/BatchDelete"
	StreamService_PerformAction_FullMethodName = "/yandex.cloud.video.v1.StreamService/PerformAction"
)

// StreamServiceClient is the client API for StreamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Stream management service.
type StreamServiceClient interface {
	// Get the specific stream.
	Get(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (*Stream, error)
	// List streams for channel.
	List(ctx context.Context, in *ListStreamsRequest, opts ...grpc.CallOption) (*ListStreamsResponse, error)
	// Batch get streams for channel.
	BatchGet(ctx context.Context, in *BatchGetStreamsRequest, opts ...grpc.CallOption) (*BatchGetStreamsResponse, error)
	// Create stream.
	Create(ctx context.Context, in *CreateStreamRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Update stream.
	Update(ctx context.Context, in *UpdateStreamRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Delete stream.
	Delete(ctx context.Context, in *DeleteStreamRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Batch delete streams.
	BatchDelete(ctx context.Context, in *BatchDeleteStreamsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Perform an action on the stream.
	PerformAction(ctx context.Context, in *PerformStreamActionRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type streamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamServiceClient(cc grpc.ClientConnInterface) StreamServiceClient {
	return &streamServiceClient{cc}
}

func (c *streamServiceClient) Get(ctx context.Context, in *GetStreamRequest, opts ...grpc.CallOption) (*Stream, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Stream)
	err := c.cc.Invoke(ctx, StreamService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) List(ctx context.Context, in *ListStreamsRequest, opts ...grpc.CallOption) (*ListStreamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStreamsResponse)
	err := c.cc.Invoke(ctx, StreamService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) BatchGet(ctx context.Context, in *BatchGetStreamsRequest, opts ...grpc.CallOption) (*BatchGetStreamsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BatchGetStreamsResponse)
	err := c.cc.Invoke(ctx, StreamService_BatchGet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Create(ctx context.Context, in *CreateStreamRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, StreamService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Update(ctx context.Context, in *UpdateStreamRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, StreamService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) Delete(ctx context.Context, in *DeleteStreamRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, StreamService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) BatchDelete(ctx context.Context, in *BatchDeleteStreamsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, StreamService_BatchDelete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamServiceClient) PerformAction(ctx context.Context, in *PerformStreamActionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, StreamService_PerformAction_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServiceServer is the server API for StreamService service.
// All implementations should embed UnimplementedStreamServiceServer
// for forward compatibility.
//
// Stream management service.
type StreamServiceServer interface {
	// Get the specific stream.
	Get(context.Context, *GetStreamRequest) (*Stream, error)
	// List streams for channel.
	List(context.Context, *ListStreamsRequest) (*ListStreamsResponse, error)
	// Batch get streams for channel.
	BatchGet(context.Context, *BatchGetStreamsRequest) (*BatchGetStreamsResponse, error)
	// Create stream.
	Create(context.Context, *CreateStreamRequest) (*operation.Operation, error)
	// Update stream.
	Update(context.Context, *UpdateStreamRequest) (*operation.Operation, error)
	// Delete stream.
	Delete(context.Context, *DeleteStreamRequest) (*operation.Operation, error)
	// Batch delete streams.
	BatchDelete(context.Context, *BatchDeleteStreamsRequest) (*operation.Operation, error)
	// Perform an action on the stream.
	PerformAction(context.Context, *PerformStreamActionRequest) (*operation.Operation, error)
}

// UnimplementedStreamServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStreamServiceServer struct{}

func (UnimplementedStreamServiceServer) Get(context.Context, *GetStreamRequest) (*Stream, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStreamServiceServer) List(context.Context, *ListStreamsRequest) (*ListStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStreamServiceServer) BatchGet(context.Context, *BatchGetStreamsRequest) (*BatchGetStreamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (UnimplementedStreamServiceServer) Create(context.Context, *CreateStreamRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStreamServiceServer) Update(context.Context, *UpdateStreamRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStreamServiceServer) Delete(context.Context, *DeleteStreamRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStreamServiceServer) BatchDelete(context.Context, *BatchDeleteStreamsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelete not implemented")
}
func (UnimplementedStreamServiceServer) PerformAction(context.Context, *PerformStreamActionRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformAction not implemented")
}
func (UnimplementedStreamServiceServer) testEmbeddedByValue() {}

// UnsafeStreamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServiceServer will
// result in compilation errors.
type UnsafeStreamServiceServer interface {
	mustEmbedUnimplementedStreamServiceServer()
}

func RegisterStreamServiceServer(s grpc.ServiceRegistrar, srv StreamServiceServer) {
	// If the following call pancis, it indicates UnimplementedStreamServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StreamService_ServiceDesc, srv)
}

func _StreamService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Get(ctx, req.(*GetStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).List(ctx, req.(*ListStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_BatchGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).BatchGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_BatchGet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).BatchGet(ctx, req.(*BatchGetStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Create(ctx, req.(*CreateStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Update(ctx, req.(*UpdateStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).Delete(ctx, req.(*DeleteStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_BatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchDeleteStreamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).BatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_BatchDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).BatchDelete(ctx, req.(*BatchDeleteStreamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamService_PerformAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformStreamActionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServiceServer).PerformAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamService_PerformAction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServiceServer).PerformAction(ctx, req.(*PerformStreamActionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamService_ServiceDesc is the grpc.ServiceDesc for StreamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.video.v1.StreamService",
	HandlerType: (*StreamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _StreamService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StreamService_List_Handler,
		},
		{
			MethodName: "BatchGet",
			Handler:    _StreamService_BatchGet_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _StreamService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _StreamService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StreamService_Delete_Handler,
		},
		{
			MethodName: "BatchDelete",
			Handler:    _StreamService_BatchDelete_Handler,
		},
		{
			MethodName: "PerformAction",
			Handler:    _StreamService_PerformAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/video/v1/stream_service.proto",
}
