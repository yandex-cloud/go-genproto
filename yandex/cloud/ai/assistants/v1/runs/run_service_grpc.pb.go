// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/ai/assistants/v1/runs/run_service.proto

package runs

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
	RunService_Create_FullMethodName          = "/yandex.cloud.ai.assistants.v1.runs.RunService/Create"
	RunService_Listen_FullMethodName          = "/yandex.cloud.ai.assistants.v1.runs.RunService/Listen"
	RunService_Attach_FullMethodName          = "/yandex.cloud.ai.assistants.v1.runs.RunService/Attach"
	RunService_Get_FullMethodName             = "/yandex.cloud.ai.assistants.v1.runs.RunService/Get"
	RunService_GetLastByThread_FullMethodName = "/yandex.cloud.ai.assistants.v1.runs.RunService/GetLastByThread"
	RunService_List_FullMethodName            = "/yandex.cloud.ai.assistants.v1.runs.RunService/List"
)

// RunServiceClient is the client API for RunService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// RunService provides operations for managing runs.
type RunServiceClient interface {
	// Create a new run for a given assistant and thread.
	Create(ctx context.Context, in *CreateRunRequest, opts ...grpc.CallOption) (*Run, error)
	// Listen to events from a specific run.
	// If the run was created with `stream = false`, Listen will only respond with the final status of the run
	// and will not stream partial messages or intermediate events.
	Listen(ctx context.Context, in *ListenRunRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamEvent], error)
	// Bi-directional streaming method to interact with a specific run.
	// Like `Listen`, `Attach` streams events from the run, but also allows clients to send events back.
	// For example, it can be used to submit function call results when the run is waiting for user input.
	Attach(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AttachRunRequest, StreamEvent], error)
	// Retrieve details of a specific run by its ID.
	Get(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*Run, error)
	// Retrieves the most recent run for a specific thread.
	GetLastByThread(ctx context.Context, in *GetLastRunByThreadRequest, opts ...grpc.CallOption) (*Run, error)
	// List runs in a specific folder.
	List(ctx context.Context, in *ListRunsRequest, opts ...grpc.CallOption) (*ListRunsResponse, error)
}

type runServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRunServiceClient(cc grpc.ClientConnInterface) RunServiceClient {
	return &runServiceClient{cc}
}

func (c *runServiceClient) Create(ctx context.Context, in *CreateRunRequest, opts ...grpc.CallOption) (*Run, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Run)
	err := c.cc.Invoke(ctx, RunService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runServiceClient) Listen(ctx context.Context, in *ListenRunRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[StreamEvent], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RunService_ServiceDesc.Streams[0], RunService_Listen_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListenRunRequest, StreamEvent]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RunService_ListenClient = grpc.ServerStreamingClient[StreamEvent]

func (c *runServiceClient) Attach(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[AttachRunRequest, StreamEvent], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &RunService_ServiceDesc.Streams[1], RunService_Attach_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[AttachRunRequest, StreamEvent]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RunService_AttachClient = grpc.BidiStreamingClient[AttachRunRequest, StreamEvent]

func (c *runServiceClient) Get(ctx context.Context, in *GetRunRequest, opts ...grpc.CallOption) (*Run, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Run)
	err := c.cc.Invoke(ctx, RunService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runServiceClient) GetLastByThread(ctx context.Context, in *GetLastRunByThreadRequest, opts ...grpc.CallOption) (*Run, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Run)
	err := c.cc.Invoke(ctx, RunService_GetLastByThread_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runServiceClient) List(ctx context.Context, in *ListRunsRequest, opts ...grpc.CallOption) (*ListRunsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListRunsResponse)
	err := c.cc.Invoke(ctx, RunService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RunServiceServer is the server API for RunService service.
// All implementations should embed UnimplementedRunServiceServer
// for forward compatibility.
//
// RunService provides operations for managing runs.
type RunServiceServer interface {
	// Create a new run for a given assistant and thread.
	Create(context.Context, *CreateRunRequest) (*Run, error)
	// Listen to events from a specific run.
	// If the run was created with `stream = false`, Listen will only respond with the final status of the run
	// and will not stream partial messages or intermediate events.
	Listen(*ListenRunRequest, grpc.ServerStreamingServer[StreamEvent]) error
	// Bi-directional streaming method to interact with a specific run.
	// Like `Listen`, `Attach` streams events from the run, but also allows clients to send events back.
	// For example, it can be used to submit function call results when the run is waiting for user input.
	Attach(grpc.BidiStreamingServer[AttachRunRequest, StreamEvent]) error
	// Retrieve details of a specific run by its ID.
	Get(context.Context, *GetRunRequest) (*Run, error)
	// Retrieves the most recent run for a specific thread.
	GetLastByThread(context.Context, *GetLastRunByThreadRequest) (*Run, error)
	// List runs in a specific folder.
	List(context.Context, *ListRunsRequest) (*ListRunsResponse, error)
}

// UnimplementedRunServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRunServiceServer struct{}

func (UnimplementedRunServiceServer) Create(context.Context, *CreateRunRequest) (*Run, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRunServiceServer) Listen(*ListenRunRequest, grpc.ServerStreamingServer[StreamEvent]) error {
	return status.Errorf(codes.Unimplemented, "method Listen not implemented")
}
func (UnimplementedRunServiceServer) Attach(grpc.BidiStreamingServer[AttachRunRequest, StreamEvent]) error {
	return status.Errorf(codes.Unimplemented, "method Attach not implemented")
}
func (UnimplementedRunServiceServer) Get(context.Context, *GetRunRequest) (*Run, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRunServiceServer) GetLastByThread(context.Context, *GetLastRunByThreadRequest) (*Run, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLastByThread not implemented")
}
func (UnimplementedRunServiceServer) List(context.Context, *ListRunsRequest) (*ListRunsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRunServiceServer) testEmbeddedByValue() {}

// UnsafeRunServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RunServiceServer will
// result in compilation errors.
type UnsafeRunServiceServer interface {
	mustEmbedUnimplementedRunServiceServer()
}

func RegisterRunServiceServer(s grpc.ServiceRegistrar, srv RunServiceServer) {
	// If the following call pancis, it indicates UnimplementedRunServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RunService_ServiceDesc, srv)
}

func _RunService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).Create(ctx, req.(*CreateRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunService_Listen_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListenRunRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RunServiceServer).Listen(m, &grpc.GenericServerStream[ListenRunRequest, StreamEvent]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RunService_ListenServer = grpc.ServerStreamingServer[StreamEvent]

func _RunService_Attach_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RunServiceServer).Attach(&grpc.GenericServerStream[AttachRunRequest, StreamEvent]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type RunService_AttachServer = grpc.BidiStreamingServer[AttachRunRequest, StreamEvent]

func _RunService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).Get(ctx, req.(*GetRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunService_GetLastByThread_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLastRunByThreadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).GetLastByThread(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunService_GetLastByThread_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).GetLastByThread(ctx, req.(*GetLastRunByThreadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RunService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RunServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RunService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RunServiceServer).List(ctx, req.(*ListRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RunService_ServiceDesc is the grpc.ServiceDesc for RunService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RunService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.assistants.v1.runs.RunService",
	HandlerType: (*RunServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _RunService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _RunService_Get_Handler,
		},
		{
			MethodName: "GetLastByThread",
			Handler:    _RunService_GetLastByThread_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RunService_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Listen",
			Handler:       _RunService_Listen_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Attach",
			Handler:       _RunService_Attach_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "yandex/cloud/ai/assistants/v1/runs/run_service.proto",
}
