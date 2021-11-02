// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package triggers

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TriggerServiceClient is the client API for TriggerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TriggerServiceClient interface {
	// Returns the specified trigger.
	//
	// To get the list of all available triggers, make a [List] request.
	Get(ctx context.Context, in *GetTriggerRequest, opts ...grpc.CallOption) (*Trigger, error)
	// Retrieves the list of triggers in the specified folder.
	List(ctx context.Context, in *ListTriggersRequest, opts ...grpc.CallOption) (*ListTriggersResponse, error)
	// Creates a trigger in the specified folder.
	Create(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified trigger.
	Update(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified trigger.
	Delete(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Pauses the specified trigger.
	Pause(ctx context.Context, in *PauseTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Restarts the specified trigger.
	Resume(ctx context.Context, in *ResumeTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified trigger.
	ListOperations(ctx context.Context, in *ListTriggerOperationsRequest, opts ...grpc.CallOption) (*ListTriggerOperationsResponse, error)
}

type triggerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTriggerServiceClient(cc grpc.ClientConnInterface) TriggerServiceClient {
	return &triggerServiceClient{cc}
}

func (c *triggerServiceClient) Get(ctx context.Context, in *GetTriggerRequest, opts ...grpc.CallOption) (*Trigger, error) {
	out := new(Trigger)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) List(ctx context.Context, in *ListTriggersRequest, opts ...grpc.CallOption) (*ListTriggersResponse, error) {
	out := new(ListTriggersResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Create(ctx context.Context, in *CreateTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Update(ctx context.Context, in *UpdateTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Delete(ctx context.Context, in *DeleteTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Pause(ctx context.Context, in *PauseTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) Resume(ctx context.Context, in *ResumeTriggerRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *triggerServiceClient) ListOperations(ctx context.Context, in *ListTriggerOperationsRequest, opts ...grpc.CallOption) (*ListTriggerOperationsResponse, error) {
	out := new(ListTriggerOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.serverless.triggers.v1.TriggerService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TriggerServiceServer is the server API for TriggerService service.
// All implementations should embed UnimplementedTriggerServiceServer
// for forward compatibility
type TriggerServiceServer interface {
	// Returns the specified trigger.
	//
	// To get the list of all available triggers, make a [List] request.
	Get(context.Context, *GetTriggerRequest) (*Trigger, error)
	// Retrieves the list of triggers in the specified folder.
	List(context.Context, *ListTriggersRequest) (*ListTriggersResponse, error)
	// Creates a trigger in the specified folder.
	Create(context.Context, *CreateTriggerRequest) (*operation.Operation, error)
	// Updates the specified trigger.
	Update(context.Context, *UpdateTriggerRequest) (*operation.Operation, error)
	// Deletes the specified trigger.
	Delete(context.Context, *DeleteTriggerRequest) (*operation.Operation, error)
	// Pauses the specified trigger.
	Pause(context.Context, *PauseTriggerRequest) (*operation.Operation, error)
	// Restarts the specified trigger.
	Resume(context.Context, *ResumeTriggerRequest) (*operation.Operation, error)
	// Lists operations for the specified trigger.
	ListOperations(context.Context, *ListTriggerOperationsRequest) (*ListTriggerOperationsResponse, error)
}

// UnimplementedTriggerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTriggerServiceServer struct {
}

func (UnimplementedTriggerServiceServer) Get(context.Context, *GetTriggerRequest) (*Trigger, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTriggerServiceServer) List(context.Context, *ListTriggersRequest) (*ListTriggersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTriggerServiceServer) Create(context.Context, *CreateTriggerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTriggerServiceServer) Update(context.Context, *UpdateTriggerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTriggerServiceServer) Delete(context.Context, *DeleteTriggerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTriggerServiceServer) Pause(context.Context, *PauseTriggerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedTriggerServiceServer) Resume(context.Context, *ResumeTriggerRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedTriggerServiceServer) ListOperations(context.Context, *ListTriggerOperationsRequest) (*ListTriggerOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeTriggerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TriggerServiceServer will
// result in compilation errors.
type UnsafeTriggerServiceServer interface {
	mustEmbedUnimplementedTriggerServiceServer()
}

func RegisterTriggerServiceServer(s grpc.ServiceRegistrar, srv TriggerServiceServer) {
	s.RegisterService(&TriggerService_ServiceDesc, srv)
}

func _TriggerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Get(ctx, req.(*GetTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTriggersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).List(ctx, req.(*ListTriggersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Create(ctx, req.(*CreateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Update(ctx, req.(*UpdateTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Delete(ctx, req.(*DeleteTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Pause(ctx, req.(*PauseTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeTriggerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).Resume(ctx, req.(*ResumeTriggerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TriggerService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTriggerOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TriggerServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.serverless.triggers.v1.TriggerService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TriggerServiceServer).ListOperations(ctx, req.(*ListTriggerOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TriggerService_ServiceDesc is the grpc.ServiceDesc for TriggerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TriggerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.serverless.triggers.v1.TriggerService",
	HandlerType: (*TriggerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TriggerService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TriggerService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TriggerService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TriggerService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TriggerService_Delete_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _TriggerService_Pause_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _TriggerService_Resume_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _TriggerService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/serverless/triggers/v1/trigger_service.proto",
}
