// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/serverless/eventrouter/v1/rule_service.proto

package eventrouter

import (
	context "context"
	access "github.com/yandex-cloud/go-genproto/yandex/cloud/access"
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
	RuleService_Get_FullMethodName                  = "/yandex.cloud.serverless.eventrouter.v1.RuleService/Get"
	RuleService_List_FullMethodName                 = "/yandex.cloud.serverless.eventrouter.v1.RuleService/List"
	RuleService_Create_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.RuleService/Create"
	RuleService_Update_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.RuleService/Update"
	RuleService_Delete_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.RuleService/Delete"
	RuleService_Enable_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.RuleService/Enable"
	RuleService_Disable_FullMethodName              = "/yandex.cloud.serverless.eventrouter.v1.RuleService/Disable"
	RuleService_ListAccessBindings_FullMethodName   = "/yandex.cloud.serverless.eventrouter.v1.RuleService/ListAccessBindings"
	RuleService_SetAccessBindings_FullMethodName    = "/yandex.cloud.serverless.eventrouter.v1.RuleService/SetAccessBindings"
	RuleService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.serverless.eventrouter.v1.RuleService/UpdateAccessBindings"
	RuleService_ListOperations_FullMethodName       = "/yandex.cloud.serverless.eventrouter.v1.RuleService/ListOperations"
)

// RuleServiceClient is the client API for RuleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuleServiceClient interface {
	// Returns the specified rules.
	// To get the list of all available buses, make a [List] request.
	Get(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error)
	// Retrieves the list of rules in the specified folder.
	List(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error)
	// Creates a rule in the specified folder.
	Create(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified rule.
	Update(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified rule.
	Delete(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Enables the specified rule.
	Enable(ctx context.Context, in *EnableRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Disables the specified rule.
	Disable(ctx context.Context, in *DisableRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists existing access bindings for the specified rule.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the rule.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified rule.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified rule.
	ListOperations(ctx context.Context, in *ListRuleOperationsRequest, opts ...grpc.CallOption) (*ListRuleOperationsResponse, error)
}

type ruleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuleServiceClient(cc grpc.ClientConnInterface) RuleServiceClient {
	return &ruleServiceClient{cc}
}

func (c *ruleServiceClient) Get(ctx context.Context, in *GetRuleRequest, opts ...grpc.CallOption) (*Rule, error) {
	out := new(Rule)
	err := c.cc.Invoke(ctx, RuleService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) List(ctx context.Context, in *ListRulesRequest, opts ...grpc.CallOption) (*ListRulesResponse, error) {
	out := new(ListRulesResponse)
	err := c.cc.Invoke(ctx, RuleService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Create(ctx context.Context, in *CreateRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Update(ctx context.Context, in *UpdateRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Delete(ctx context.Context, in *DeleteRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Enable(ctx context.Context, in *EnableRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_Enable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) Disable(ctx context.Context, in *DisableRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_Disable_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, RuleService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, RuleService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ruleServiceClient) ListOperations(ctx context.Context, in *ListRuleOperationsRequest, opts ...grpc.CallOption) (*ListRuleOperationsResponse, error) {
	out := new(ListRuleOperationsResponse)
	err := c.cc.Invoke(ctx, RuleService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuleServiceServer is the server API for RuleService service.
// All implementations should embed UnimplementedRuleServiceServer
// for forward compatibility
type RuleServiceServer interface {
	// Returns the specified rules.
	// To get the list of all available buses, make a [List] request.
	Get(context.Context, *GetRuleRequest) (*Rule, error)
	// Retrieves the list of rules in the specified folder.
	List(context.Context, *ListRulesRequest) (*ListRulesResponse, error)
	// Creates a rule in the specified folder.
	Create(context.Context, *CreateRuleRequest) (*operation.Operation, error)
	// Updates the specified rule.
	Update(context.Context, *UpdateRuleRequest) (*operation.Operation, error)
	// Deletes the specified rule.
	Delete(context.Context, *DeleteRuleRequest) (*operation.Operation, error)
	// Enables the specified rule.
	Enable(context.Context, *EnableRuleRequest) (*operation.Operation, error)
	// Disables the specified rule.
	Disable(context.Context, *DisableRuleRequest) (*operation.Operation, error)
	// Lists existing access bindings for the specified rule.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the rule.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified rule.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
	// Lists operations for the specified rule.
	ListOperations(context.Context, *ListRuleOperationsRequest) (*ListRuleOperationsResponse, error)
}

// UnimplementedRuleServiceServer should be embedded to have forward compatible implementations.
type UnimplementedRuleServiceServer struct {
}

func (UnimplementedRuleServiceServer) Get(context.Context, *GetRuleRequest) (*Rule, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRuleServiceServer) List(context.Context, *ListRulesRequest) (*ListRulesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedRuleServiceServer) Create(context.Context, *CreateRuleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedRuleServiceServer) Update(context.Context, *UpdateRuleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRuleServiceServer) Delete(context.Context, *DeleteRuleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedRuleServiceServer) Enable(context.Context, *EnableRuleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Enable not implemented")
}
func (UnimplementedRuleServiceServer) Disable(context.Context, *DisableRuleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disable not implemented")
}
func (UnimplementedRuleServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedRuleServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedRuleServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedRuleServiceServer) ListOperations(context.Context, *ListRuleOperationsRequest) (*ListRuleOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeRuleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuleServiceServer will
// result in compilation errors.
type UnsafeRuleServiceServer interface {
	mustEmbedUnimplementedRuleServiceServer()
}

func RegisterRuleServiceServer(s grpc.ServiceRegistrar, srv RuleServiceServer) {
	s.RegisterService(&RuleService_ServiceDesc, srv)
}

func _RuleService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Get(ctx, req.(*GetRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).List(ctx, req.(*ListRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Create(ctx, req.(*CreateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Update(ctx, req.(*UpdateRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Delete(ctx, req.(*DeleteRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Enable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnableRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Enable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Enable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Enable(ctx, req.(*EnableRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_Disable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisableRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).Disable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_Disable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).Disable(ctx, req.(*DisableRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RuleService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRuleOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuleServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RuleService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuleServiceServer).ListOperations(ctx, req.(*ListRuleOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RuleService_ServiceDesc is the grpc.ServiceDesc for RuleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.serverless.eventrouter.v1.RuleService",
	HandlerType: (*RuleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _RuleService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _RuleService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _RuleService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _RuleService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _RuleService_Delete_Handler,
		},
		{
			MethodName: "Enable",
			Handler:    _RuleService_Enable_Handler,
		},
		{
			MethodName: "Disable",
			Handler:    _RuleService_Disable_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _RuleService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _RuleService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _RuleService_UpdateAccessBindings_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _RuleService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/serverless/eventrouter/v1/rule_service.proto",
}
