// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/loadtesting/api/v1/agent_service.proto

package loadtesting

import (
	context "context"
	agent "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/agent"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AgentService_Create_FullMethodName = "/yandex.cloud.loadtesting.api.v1.AgentService/Create"
	AgentService_Get_FullMethodName    = "/yandex.cloud.loadtesting.api.v1.AgentService/Get"
	AgentService_List_FullMethodName   = "/yandex.cloud.loadtesting.api.v1.AgentService/List"
	AgentService_Delete_FullMethodName = "/yandex.cloud.loadtesting.api.v1.AgentService/Delete"
	AgentService_Update_FullMethodName = "/yandex.cloud.loadtesting.api.v1.AgentService/Update"
)

// AgentServiceClient is the client API for AgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgentServiceClient interface {
	// Creates an agent in the specified folder.
	//
	// Also creates a corresponding compute instance.
	Create(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified agent.
	//
	// To get the list of all available agents, make a [List] request.
	Get(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*agent.Agent, error)
	// Retrieves the list of agents in the specified folder.
	List(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error)
	// Deletes the specified agent.
	//
	// Also deletes a corresponding compute instance.
	Delete(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified agent.
	Update(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type agentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentServiceClient(cc grpc.ClientConnInterface) AgentServiceClient {
	return &agentServiceClient{cc}
}

func (c *agentServiceClient) Create(ctx context.Context, in *CreateAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AgentService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Get(ctx context.Context, in *GetAgentRequest, opts ...grpc.CallOption) (*agent.Agent, error) {
	out := new(agent.Agent)
	err := c.cc.Invoke(ctx, AgentService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) List(ctx context.Context, in *ListAgentsRequest, opts ...grpc.CallOption) (*ListAgentsResponse, error) {
	out := new(ListAgentsResponse)
	err := c.cc.Invoke(ctx, AgentService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Delete(ctx context.Context, in *DeleteAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AgentService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentServiceClient) Update(ctx context.Context, in *UpdateAgentRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, AgentService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentServiceServer is the server API for AgentService service.
// All implementations should embed UnimplementedAgentServiceServer
// for forward compatibility
type AgentServiceServer interface {
	// Creates an agent in the specified folder.
	//
	// Also creates a corresponding compute instance.
	Create(context.Context, *CreateAgentRequest) (*operation.Operation, error)
	// Returns the specified agent.
	//
	// To get the list of all available agents, make a [List] request.
	Get(context.Context, *GetAgentRequest) (*agent.Agent, error)
	// Retrieves the list of agents in the specified folder.
	List(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error)
	// Deletes the specified agent.
	//
	// Also deletes a corresponding compute instance.
	Delete(context.Context, *DeleteAgentRequest) (*operation.Operation, error)
	// Updates the specified agent.
	Update(context.Context, *UpdateAgentRequest) (*operation.Operation, error)
}

// UnimplementedAgentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAgentServiceServer struct {
}

func (UnimplementedAgentServiceServer) Create(context.Context, *CreateAgentRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAgentServiceServer) Get(context.Context, *GetAgentRequest) (*agent.Agent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAgentServiceServer) List(context.Context, *ListAgentsRequest) (*ListAgentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedAgentServiceServer) Delete(context.Context, *DeleteAgentRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAgentServiceServer) Update(context.Context, *UpdateAgentRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

// UnsafeAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgentServiceServer will
// result in compilation errors.
type UnsafeAgentServiceServer interface {
	mustEmbedUnimplementedAgentServiceServer()
}

func RegisterAgentServiceServer(s grpc.ServiceRegistrar, srv AgentServiceServer) {
	s.RegisterService(&AgentService_ServiceDesc, srv)
}

func _AgentService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Create(ctx, req.(*CreateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Get(ctx, req.(*GetAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).List(ctx, req.(*ListAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Delete(ctx, req.(*DeleteAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgentService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AgentService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServiceServer).Update(ctx, req.(*UpdateAgentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgentService_ServiceDesc is the grpc.ServiceDesc for AgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.loadtesting.api.v1.AgentService",
	HandlerType: (*AgentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AgentService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _AgentService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _AgentService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AgentService_Delete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AgentService_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/loadtesting/api/v1/agent_service.proto",
}
