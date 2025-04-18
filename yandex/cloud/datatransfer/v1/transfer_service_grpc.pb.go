// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/datatransfer/v1/transfer_service.proto

package datatransfer

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
	TransferService_Create_FullMethodName     = "/yandex.cloud.datatransfer.v1.TransferService/Create"
	TransferService_Update_FullMethodName     = "/yandex.cloud.datatransfer.v1.TransferService/Update"
	TransferService_Delete_FullMethodName     = "/yandex.cloud.datatransfer.v1.TransferService/Delete"
	TransferService_List_FullMethodName       = "/yandex.cloud.datatransfer.v1.TransferService/List"
	TransferService_Get_FullMethodName        = "/yandex.cloud.datatransfer.v1.TransferService/Get"
	TransferService_Deactivate_FullMethodName = "/yandex.cloud.datatransfer.v1.TransferService/Deactivate"
	TransferService_Activate_FullMethodName   = "/yandex.cloud.datatransfer.v1.TransferService/Activate"
)

// TransferServiceClient is the client API for TransferService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransferServiceClient interface {
	// Creates a transfer in the specified folder.
	Create(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified transfer.
	Update(ctx context.Context, in *UpdateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified transfer.
	Delete(ctx context.Context, in *DeleteTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists transfers in the specified folder.
	List(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error)
	// Returns the specified transfer.
	//
	// To get the list of all available transfers, make a [List] request.
	Get(ctx context.Context, in *GetTransferRequest, opts ...grpc.CallOption) (*Transfer, error)
	// Deactivates the specified transfer.
	//
	// To get the list of all available transfers, make a [List] request.
	Deactivate(ctx context.Context, in *DeactivateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Activates the specified transfer.
	//
	// To get the list of all available transfers, make a [List] request.
	Activate(ctx context.Context, in *ActivateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type transferServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTransferServiceClient(cc grpc.ClientConnInterface) TransferServiceClient {
	return &transferServiceClient{cc}
}

func (c *transferServiceClient) Create(ctx context.Context, in *CreateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TransferService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) Update(ctx context.Context, in *UpdateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TransferService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) Delete(ctx context.Context, in *DeleteTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TransferService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) List(ctx context.Context, in *ListTransfersRequest, opts ...grpc.CallOption) (*ListTransfersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTransfersResponse)
	err := c.cc.Invoke(ctx, TransferService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) Get(ctx context.Context, in *GetTransferRequest, opts ...grpc.CallOption) (*Transfer, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Transfer)
	err := c.cc.Invoke(ctx, TransferService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) Deactivate(ctx context.Context, in *DeactivateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TransferService_Deactivate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *transferServiceClient) Activate(ctx context.Context, in *ActivateTransferRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TransferService_Activate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransferServiceServer is the server API for TransferService service.
// All implementations should embed UnimplementedTransferServiceServer
// for forward compatibility.
type TransferServiceServer interface {
	// Creates a transfer in the specified folder.
	Create(context.Context, *CreateTransferRequest) (*operation.Operation, error)
	// Updates the specified transfer.
	Update(context.Context, *UpdateTransferRequest) (*operation.Operation, error)
	// Deletes the specified transfer.
	Delete(context.Context, *DeleteTransferRequest) (*operation.Operation, error)
	// Lists transfers in the specified folder.
	List(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error)
	// Returns the specified transfer.
	//
	// To get the list of all available transfers, make a [List] request.
	Get(context.Context, *GetTransferRequest) (*Transfer, error)
	// Deactivates the specified transfer.
	//
	// To get the list of all available transfers, make a [List] request.
	Deactivate(context.Context, *DeactivateTransferRequest) (*operation.Operation, error)
	// Activates the specified transfer.
	//
	// To get the list of all available transfers, make a [List] request.
	Activate(context.Context, *ActivateTransferRequest) (*operation.Operation, error)
}

// UnimplementedTransferServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTransferServiceServer struct{}

func (UnimplementedTransferServiceServer) Create(context.Context, *CreateTransferRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTransferServiceServer) Update(context.Context, *UpdateTransferRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTransferServiceServer) Delete(context.Context, *DeleteTransferRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTransferServiceServer) List(context.Context, *ListTransfersRequest) (*ListTransfersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedTransferServiceServer) Get(context.Context, *GetTransferRequest) (*Transfer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTransferServiceServer) Deactivate(context.Context, *DeactivateTransferRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deactivate not implemented")
}
func (UnimplementedTransferServiceServer) Activate(context.Context, *ActivateTransferRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}
func (UnimplementedTransferServiceServer) testEmbeddedByValue() {}

// UnsafeTransferServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransferServiceServer will
// result in compilation errors.
type UnsafeTransferServiceServer interface {
	mustEmbedUnimplementedTransferServiceServer()
}

func RegisterTransferServiceServer(s grpc.ServiceRegistrar, srv TransferServiceServer) {
	// If the following call pancis, it indicates UnimplementedTransferServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TransferService_ServiceDesc, srv)
}

func _TransferService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Create(ctx, req.(*CreateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Update(ctx, req.(*UpdateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Delete(ctx, req.(*DeleteTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransfersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).List(ctx, req.(*ListTransfersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Get(ctx, req.(*GetTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_Deactivate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Deactivate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Deactivate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Deactivate(ctx, req.(*DeactivateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TransferService_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateTransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransferServiceServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TransferService_Activate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransferServiceServer).Activate(ctx, req.(*ActivateTransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransferService_ServiceDesc is the grpc.ServiceDesc for TransferService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransferService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.datatransfer.v1.TransferService",
	HandlerType: (*TransferServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TransferService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TransferService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TransferService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TransferService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _TransferService_Get_Handler,
		},
		{
			MethodName: "Deactivate",
			Handler:    _TransferService_Deactivate_Handler,
		},
		{
			MethodName: "Activate",
			Handler:    _TransferService_Activate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/datatransfer/v1/transfer_service.proto",
}
