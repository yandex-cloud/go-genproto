// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/vpc/v1/gateway_service.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GatewayService_Get_FullMethodName            = "/yandex.cloud.vpc.v1.GatewayService/Get"
	GatewayService_List_FullMethodName           = "/yandex.cloud.vpc.v1.GatewayService/List"
	GatewayService_Create_FullMethodName         = "/yandex.cloud.vpc.v1.GatewayService/Create"
	GatewayService_Update_FullMethodName         = "/yandex.cloud.vpc.v1.GatewayService/Update"
	GatewayService_Delete_FullMethodName         = "/yandex.cloud.vpc.v1.GatewayService/Delete"
	GatewayService_ListOperations_FullMethodName = "/yandex.cloud.vpc.v1.GatewayService/ListOperations"
	GatewayService_Move_FullMethodName           = "/yandex.cloud.vpc.v1.GatewayService/Move"
)

// GatewayServiceClient is the client API for GatewayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayServiceClient interface {
	// Returns the specified Gateway resource.
	//
	// To get the list of all available Gateway resources, make a [List] request.
	Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error)
	// Retrieves the list of Gateway resources in the specified folder.
	List(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*ListGatewaysResponse, error)
	// Creates a gateway in the specified folder.
	Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified gateway.
	Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified gateway.
	Delete(ctx context.Context, in *DeleteGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// List operations for the specified gateway.
	ListOperations(ctx context.Context, in *ListGatewayOperationsRequest, opts ...grpc.CallOption) (*ListGatewayOperationsResponse, error)
	// Move a gateway to another folder
	Move(ctx context.Context, in *MoveGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type gatewayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayServiceClient(cc grpc.ClientConnInterface) GatewayServiceClient {
	return &gatewayServiceClient{cc}
}

func (c *gatewayServiceClient) Get(ctx context.Context, in *GetGatewayRequest, opts ...grpc.CallOption) (*Gateway, error) {
	out := new(Gateway)
	err := c.cc.Invoke(ctx, GatewayService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) List(ctx context.Context, in *ListGatewaysRequest, opts ...grpc.CallOption) (*ListGatewaysResponse, error) {
	out := new(ListGatewaysResponse)
	err := c.cc.Invoke(ctx, GatewayService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Create(ctx context.Context, in *CreateGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GatewayService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Update(ctx context.Context, in *UpdateGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GatewayService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Delete(ctx context.Context, in *DeleteGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GatewayService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) ListOperations(ctx context.Context, in *ListGatewayOperationsRequest, opts ...grpc.CallOption) (*ListGatewayOperationsResponse, error) {
	out := new(ListGatewayOperationsResponse)
	err := c.cc.Invoke(ctx, GatewayService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayServiceClient) Move(ctx context.Context, in *MoveGatewayRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, GatewayService_Move_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayServiceServer is the server API for GatewayService service.
// All implementations should embed UnimplementedGatewayServiceServer
// for forward compatibility
type GatewayServiceServer interface {
	// Returns the specified Gateway resource.
	//
	// To get the list of all available Gateway resources, make a [List] request.
	Get(context.Context, *GetGatewayRequest) (*Gateway, error)
	// Retrieves the list of Gateway resources in the specified folder.
	List(context.Context, *ListGatewaysRequest) (*ListGatewaysResponse, error)
	// Creates a gateway in the specified folder.
	Create(context.Context, *CreateGatewayRequest) (*operation.Operation, error)
	// Updates the specified gateway.
	Update(context.Context, *UpdateGatewayRequest) (*operation.Operation, error)
	// Deletes the specified gateway.
	Delete(context.Context, *DeleteGatewayRequest) (*operation.Operation, error)
	// List operations for the specified gateway.
	ListOperations(context.Context, *ListGatewayOperationsRequest) (*ListGatewayOperationsResponse, error)
	// Move a gateway to another folder
	Move(context.Context, *MoveGatewayRequest) (*operation.Operation, error)
}

// UnimplementedGatewayServiceServer should be embedded to have forward compatible implementations.
type UnimplementedGatewayServiceServer struct {
}

func (UnimplementedGatewayServiceServer) Get(context.Context, *GetGatewayRequest) (*Gateway, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedGatewayServiceServer) List(context.Context, *ListGatewaysRequest) (*ListGatewaysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedGatewayServiceServer) Create(context.Context, *CreateGatewayRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedGatewayServiceServer) Update(context.Context, *UpdateGatewayRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedGatewayServiceServer) Delete(context.Context, *DeleteGatewayRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedGatewayServiceServer) ListOperations(context.Context, *ListGatewayOperationsRequest) (*ListGatewayOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedGatewayServiceServer) Move(context.Context, *MoveGatewayRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}

// UnsafeGatewayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayServiceServer will
// result in compilation errors.
type UnsafeGatewayServiceServer interface {
	mustEmbedUnimplementedGatewayServiceServer()
}

func RegisterGatewayServiceServer(s grpc.ServiceRegistrar, srv GatewayServiceServer) {
	s.RegisterService(&GatewayService_ServiceDesc, srv)
}

func _GatewayService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Get(ctx, req.(*GetGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewaysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).List(ctx, req.(*ListGatewaysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Create(ctx, req.(*CreateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Update(ctx, req.(*UpdateGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Delete(ctx, req.(*DeleteGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListGatewayOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).ListOperations(ctx, req.(*ListGatewayOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveGatewayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayService_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayServiceServer).Move(ctx, req.(*MoveGatewayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayService_ServiceDesc is the grpc.ServiceDesc for GatewayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.vpc.v1.GatewayService",
	HandlerType: (*GatewayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _GatewayService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _GatewayService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _GatewayService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _GatewayService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _GatewayService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _GatewayService_ListOperations_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _GatewayService_Move_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/vpc/v1/gateway_service.proto",
}
