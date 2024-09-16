// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/serverless/eventrouter/v1/bus_service.proto

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
	BusService_Get_FullMethodName                  = "/yandex.cloud.serverless.eventrouter.v1.BusService/Get"
	BusService_List_FullMethodName                 = "/yandex.cloud.serverless.eventrouter.v1.BusService/List"
	BusService_Create_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.BusService/Create"
	BusService_Update_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.BusService/Update"
	BusService_Delete_FullMethodName               = "/yandex.cloud.serverless.eventrouter.v1.BusService/Delete"
	BusService_ListOperations_FullMethodName       = "/yandex.cloud.serverless.eventrouter.v1.BusService/ListOperations"
	BusService_ListAccessBindings_FullMethodName   = "/yandex.cloud.serverless.eventrouter.v1.BusService/ListAccessBindings"
	BusService_SetAccessBindings_FullMethodName    = "/yandex.cloud.serverless.eventrouter.v1.BusService/SetAccessBindings"
	BusService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.serverless.eventrouter.v1.BusService/UpdateAccessBindings"
)

// BusServiceClient is the client API for BusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusServiceClient interface {
	// Returns the specified bus.
	// To get the list of all available buses, make a [List] request.
	Get(ctx context.Context, in *GetBusRequest, opts ...grpc.CallOption) (*Bus, error)
	// Retrieves the list of buses in the specified folder.
	List(ctx context.Context, in *ListBusesRequest, opts ...grpc.CallOption) (*ListBusesResponse, error)
	// Creates a bus in the specified folder.
	Create(ctx context.Context, in *CreateBusRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified bus.
	Update(ctx context.Context, in *UpdateBusRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified bus.
	Delete(ctx context.Context, in *DeleteBusRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified bus.
	ListOperations(ctx context.Context, in *ListBusOperationsRequest, opts ...grpc.CallOption) (*ListBusOperationsResponse, error)
	// Lists existing access bindings for the specified bus.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the bus.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified bus.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type busServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusServiceClient(cc grpc.ClientConnInterface) BusServiceClient {
	return &busServiceClient{cc}
}

func (c *busServiceClient) Get(ctx context.Context, in *GetBusRequest, opts ...grpc.CallOption) (*Bus, error) {
	out := new(Bus)
	err := c.cc.Invoke(ctx, BusService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) List(ctx context.Context, in *ListBusesRequest, opts ...grpc.CallOption) (*ListBusesResponse, error) {
	out := new(ListBusesResponse)
	err := c.cc.Invoke(ctx, BusService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) Create(ctx context.Context, in *CreateBusRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BusService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) Update(ctx context.Context, in *UpdateBusRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BusService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) Delete(ctx context.Context, in *DeleteBusRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BusService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) ListOperations(ctx context.Context, in *ListBusOperationsRequest, opts ...grpc.CallOption) (*ListBusOperationsResponse, error) {
	out := new(ListBusOperationsResponse)
	err := c.cc.Invoke(ctx, BusService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, BusService_ListAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BusService_SetAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BusService_UpdateAccessBindings_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusServiceServer is the server API for BusService service.
// All implementations should embed UnimplementedBusServiceServer
// for forward compatibility
type BusServiceServer interface {
	// Returns the specified bus.
	// To get the list of all available buses, make a [List] request.
	Get(context.Context, *GetBusRequest) (*Bus, error)
	// Retrieves the list of buses in the specified folder.
	List(context.Context, *ListBusesRequest) (*ListBusesResponse, error)
	// Creates a bus in the specified folder.
	Create(context.Context, *CreateBusRequest) (*operation.Operation, error)
	// Updates the specified bus.
	Update(context.Context, *UpdateBusRequest) (*operation.Operation, error)
	// Deletes the specified bus.
	Delete(context.Context, *DeleteBusRequest) (*operation.Operation, error)
	// Lists operations for the specified bus.
	ListOperations(context.Context, *ListBusOperationsRequest) (*ListBusOperationsResponse, error)
	// Lists existing access bindings for the specified bus.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the bus.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified bus.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedBusServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBusServiceServer struct {
}

func (UnimplementedBusServiceServer) Get(context.Context, *GetBusRequest) (*Bus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBusServiceServer) List(context.Context, *ListBusesRequest) (*ListBusesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBusServiceServer) Create(context.Context, *CreateBusRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBusServiceServer) Update(context.Context, *UpdateBusRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBusServiceServer) Delete(context.Context, *DeleteBusRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBusServiceServer) ListOperations(context.Context, *ListBusOperationsRequest) (*ListBusOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedBusServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedBusServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedBusServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}

// UnsafeBusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusServiceServer will
// result in compilation errors.
type UnsafeBusServiceServer interface {
	mustEmbedUnimplementedBusServiceServer()
}

func RegisterBusServiceServer(s grpc.ServiceRegistrar, srv BusServiceServer) {
	s.RegisterService(&BusService_ServiceDesc, srv)
}

func _BusService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).Get(ctx, req.(*GetBusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).List(ctx, req.(*ListBusesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).Create(ctx, req.(*CreateBusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).Update(ctx, req.(*UpdateBusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).Delete(ctx, req.(*DeleteBusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).ListOperations(ctx, req.(*ListBusOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BusService_ServiceDesc is the grpc.ServiceDesc for BusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.serverless.eventrouter.v1.BusService",
	HandlerType: (*BusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BusService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BusService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _BusService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BusService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BusService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _BusService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _BusService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _BusService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _BusService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/serverless/eventrouter/v1/bus_service.proto",
}
