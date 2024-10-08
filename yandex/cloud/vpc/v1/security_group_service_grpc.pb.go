// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/vpc/v1/security_group_service.proto

package vpc

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
	SecurityGroupService_Get_FullMethodName            = "/yandex.cloud.vpc.v1.SecurityGroupService/Get"
	SecurityGroupService_List_FullMethodName           = "/yandex.cloud.vpc.v1.SecurityGroupService/List"
	SecurityGroupService_Create_FullMethodName         = "/yandex.cloud.vpc.v1.SecurityGroupService/Create"
	SecurityGroupService_Update_FullMethodName         = "/yandex.cloud.vpc.v1.SecurityGroupService/Update"
	SecurityGroupService_UpdateRules_FullMethodName    = "/yandex.cloud.vpc.v1.SecurityGroupService/UpdateRules"
	SecurityGroupService_UpdateRule_FullMethodName     = "/yandex.cloud.vpc.v1.SecurityGroupService/UpdateRule"
	SecurityGroupService_Delete_FullMethodName         = "/yandex.cloud.vpc.v1.SecurityGroupService/Delete"
	SecurityGroupService_Move_FullMethodName           = "/yandex.cloud.vpc.v1.SecurityGroupService/Move"
	SecurityGroupService_ListOperations_FullMethodName = "/yandex.cloud.vpc.v1.SecurityGroupService/ListOperations"
)

// SecurityGroupServiceClient is the client API for SecurityGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing SecurityGroup resources.
type SecurityGroupServiceClient interface {
	// Returns the specified SecurityGroup resource.
	//
	// To get the list of all available SecurityGroup resources, make a [List] request.
	Get(ctx context.Context, in *GetSecurityGroupRequest, opts ...grpc.CallOption) (*SecurityGroup, error)
	// Retrieves the list of SecurityGroup resources in the specified folder.
	List(ctx context.Context, in *ListSecurityGroupsRequest, opts ...grpc.CallOption) (*ListSecurityGroupsResponse, error)
	// Creates a security group in the specified folder and network.
	Create(ctx context.Context, in *CreateSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified security group.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Update(ctx context.Context, in *UpdateSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the rules of the specified security group.
	UpdateRules(ctx context.Context, in *UpdateSecurityGroupRulesRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified rule.
	UpdateRule(ctx context.Context, in *UpdateSecurityGroupRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified security group.
	Delete(ctx context.Context, in *DeleteSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Moves security groups to another folder.
	Move(ctx context.Context, in *MoveSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified security groups.
	ListOperations(ctx context.Context, in *ListSecurityGroupOperationsRequest, opts ...grpc.CallOption) (*ListSecurityGroupOperationsResponse, error)
}

type securityGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityGroupServiceClient(cc grpc.ClientConnInterface) SecurityGroupServiceClient {
	return &securityGroupServiceClient{cc}
}

func (c *securityGroupServiceClient) Get(ctx context.Context, in *GetSecurityGroupRequest, opts ...grpc.CallOption) (*SecurityGroup, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SecurityGroup)
	err := c.cc.Invoke(ctx, SecurityGroupService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) List(ctx context.Context, in *ListSecurityGroupsRequest, opts ...grpc.CallOption) (*ListSecurityGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSecurityGroupsResponse)
	err := c.cc.Invoke(ctx, SecurityGroupService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) Create(ctx context.Context, in *CreateSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityGroupService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) Update(ctx context.Context, in *UpdateSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityGroupService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) UpdateRules(ctx context.Context, in *UpdateSecurityGroupRulesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityGroupService_UpdateRules_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) UpdateRule(ctx context.Context, in *UpdateSecurityGroupRuleRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityGroupService_UpdateRule_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) Delete(ctx context.Context, in *DeleteSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityGroupService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) Move(ctx context.Context, in *MoveSecurityGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityGroupService_Move_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityGroupServiceClient) ListOperations(ctx context.Context, in *ListSecurityGroupOperationsRequest, opts ...grpc.CallOption) (*ListSecurityGroupOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSecurityGroupOperationsResponse)
	err := c.cc.Invoke(ctx, SecurityGroupService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityGroupServiceServer is the server API for SecurityGroupService service.
// All implementations should embed UnimplementedSecurityGroupServiceServer
// for forward compatibility.
//
// A set of methods for managing SecurityGroup resources.
type SecurityGroupServiceServer interface {
	// Returns the specified SecurityGroup resource.
	//
	// To get the list of all available SecurityGroup resources, make a [List] request.
	Get(context.Context, *GetSecurityGroupRequest) (*SecurityGroup, error)
	// Retrieves the list of SecurityGroup resources in the specified folder.
	List(context.Context, *ListSecurityGroupsRequest) (*ListSecurityGroupsResponse, error)
	// Creates a security group in the specified folder and network.
	Create(context.Context, *CreateSecurityGroupRequest) (*operation.Operation, error)
	// Updates the specified security group.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Update(context.Context, *UpdateSecurityGroupRequest) (*operation.Operation, error)
	// Updates the rules of the specified security group.
	UpdateRules(context.Context, *UpdateSecurityGroupRulesRequest) (*operation.Operation, error)
	// Updates the specified rule.
	UpdateRule(context.Context, *UpdateSecurityGroupRuleRequest) (*operation.Operation, error)
	// Deletes the specified security group.
	Delete(context.Context, *DeleteSecurityGroupRequest) (*operation.Operation, error)
	// Moves security groups to another folder.
	Move(context.Context, *MoveSecurityGroupRequest) (*operation.Operation, error)
	// Lists operations for the specified security groups.
	ListOperations(context.Context, *ListSecurityGroupOperationsRequest) (*ListSecurityGroupOperationsResponse, error)
}

// UnimplementedSecurityGroupServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSecurityGroupServiceServer struct{}

func (UnimplementedSecurityGroupServiceServer) Get(context.Context, *GetSecurityGroupRequest) (*SecurityGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSecurityGroupServiceServer) List(context.Context, *ListSecurityGroupsRequest) (*ListSecurityGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSecurityGroupServiceServer) Create(context.Context, *CreateSecurityGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSecurityGroupServiceServer) Update(context.Context, *UpdateSecurityGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSecurityGroupServiceServer) UpdateRules(context.Context, *UpdateSecurityGroupRulesRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRules not implemented")
}
func (UnimplementedSecurityGroupServiceServer) UpdateRule(context.Context, *UpdateSecurityGroupRuleRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRule not implemented")
}
func (UnimplementedSecurityGroupServiceServer) Delete(context.Context, *DeleteSecurityGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSecurityGroupServiceServer) Move(context.Context, *MoveSecurityGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedSecurityGroupServiceServer) ListOperations(context.Context, *ListSecurityGroupOperationsRequest) (*ListSecurityGroupOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedSecurityGroupServiceServer) testEmbeddedByValue() {}

// UnsafeSecurityGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityGroupServiceServer will
// result in compilation errors.
type UnsafeSecurityGroupServiceServer interface {
	mustEmbedUnimplementedSecurityGroupServiceServer()
}

func RegisterSecurityGroupServiceServer(s grpc.ServiceRegistrar, srv SecurityGroupServiceServer) {
	// If the following call pancis, it indicates UnimplementedSecurityGroupServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SecurityGroupService_ServiceDesc, srv)
}

func _SecurityGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).Get(ctx, req.(*GetSecurityGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).List(ctx, req.(*ListSecurityGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecurityGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).Create(ctx, req.(*CreateSecurityGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecurityGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).Update(ctx, req.(*UpdateSecurityGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_UpdateRules_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecurityGroupRulesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).UpdateRules(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_UpdateRules_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).UpdateRules(ctx, req.(*UpdateSecurityGroupRulesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_UpdateRule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecurityGroupRuleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).UpdateRule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_UpdateRule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).UpdateRule(ctx, req.(*UpdateSecurityGroupRuleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecurityGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).Delete(ctx, req.(*DeleteSecurityGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveSecurityGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).Move(ctx, req.(*MoveSecurityGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityGroupService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityGroupOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityGroupServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityGroupService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityGroupServiceServer).ListOperations(ctx, req.(*ListSecurityGroupOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityGroupService_ServiceDesc is the grpc.ServiceDesc for SecurityGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.vpc.v1.SecurityGroupService",
	HandlerType: (*SecurityGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SecurityGroupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SecurityGroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SecurityGroupService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SecurityGroupService_Update_Handler,
		},
		{
			MethodName: "UpdateRules",
			Handler:    _SecurityGroupService_UpdateRules_Handler,
		},
		{
			MethodName: "UpdateRule",
			Handler:    _SecurityGroupService_UpdateRule_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SecurityGroupService_Delete_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _SecurityGroupService_Move_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _SecurityGroupService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/vpc/v1/security_group_service.proto",
}
