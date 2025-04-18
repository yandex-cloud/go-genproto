// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/baremetal/v1alpha/public_subnet_service.proto

package baremetal

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
	PublicSubnetService_Get_FullMethodName            = "/yandex.cloud.baremetal.v1alpha.PublicSubnetService/Get"
	PublicSubnetService_List_FullMethodName           = "/yandex.cloud.baremetal.v1alpha.PublicSubnetService/List"
	PublicSubnetService_Create_FullMethodName         = "/yandex.cloud.baremetal.v1alpha.PublicSubnetService/Create"
	PublicSubnetService_Update_FullMethodName         = "/yandex.cloud.baremetal.v1alpha.PublicSubnetService/Update"
	PublicSubnetService_Delete_FullMethodName         = "/yandex.cloud.baremetal.v1alpha.PublicSubnetService/Delete"
	PublicSubnetService_ListOperations_FullMethodName = "/yandex.cloud.baremetal.v1alpha.PublicSubnetService/ListOperations"
)

// PublicSubnetServiceClient is the client API for PublicSubnetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing PublicSubnet resources.
type PublicSubnetServiceClient interface {
	// Returns the specific PublicSubnet resource.
	//
	// To get the list of available PublicSubnet resources, make a [List] request.
	Get(ctx context.Context, in *GetPublicSubnetRequest, opts ...grpc.CallOption) (*PublicSubnet, error)
	// Retrieves the list of PublicSubnet resources in the specified folder.
	List(ctx context.Context, in *ListPublicSubnetRequest, opts ...grpc.CallOption) (*ListPublicSubnetResponse, error)
	// Creates a public subnet in the specified folder.
	Create(ctx context.Context, in *CreatePublicSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified public subnet.
	Update(ctx context.Context, in *UpdatePublicSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified public subnet.
	//
	// Deleting a public subnet removes its data permanently and is irreversible.
	Delete(ctx context.Context, in *DeletePublicSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified public subnet.
	ListOperations(ctx context.Context, in *ListPublicSubnetOperationsRequest, opts ...grpc.CallOption) (*ListPublicSubnetOperationsResponse, error)
}

type publicSubnetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPublicSubnetServiceClient(cc grpc.ClientConnInterface) PublicSubnetServiceClient {
	return &publicSubnetServiceClient{cc}
}

func (c *publicSubnetServiceClient) Get(ctx context.Context, in *GetPublicSubnetRequest, opts ...grpc.CallOption) (*PublicSubnet, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublicSubnet)
	err := c.cc.Invoke(ctx, PublicSubnetService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicSubnetServiceClient) List(ctx context.Context, in *ListPublicSubnetRequest, opts ...grpc.CallOption) (*ListPublicSubnetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPublicSubnetResponse)
	err := c.cc.Invoke(ctx, PublicSubnetService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicSubnetServiceClient) Create(ctx context.Context, in *CreatePublicSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, PublicSubnetService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicSubnetServiceClient) Update(ctx context.Context, in *UpdatePublicSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, PublicSubnetService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicSubnetServiceClient) Delete(ctx context.Context, in *DeletePublicSubnetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, PublicSubnetService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publicSubnetServiceClient) ListOperations(ctx context.Context, in *ListPublicSubnetOperationsRequest, opts ...grpc.CallOption) (*ListPublicSubnetOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListPublicSubnetOperationsResponse)
	err := c.cc.Invoke(ctx, PublicSubnetService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PublicSubnetServiceServer is the server API for PublicSubnetService service.
// All implementations should embed UnimplementedPublicSubnetServiceServer
// for forward compatibility.
//
// A set of methods for managing PublicSubnet resources.
type PublicSubnetServiceServer interface {
	// Returns the specific PublicSubnet resource.
	//
	// To get the list of available PublicSubnet resources, make a [List] request.
	Get(context.Context, *GetPublicSubnetRequest) (*PublicSubnet, error)
	// Retrieves the list of PublicSubnet resources in the specified folder.
	List(context.Context, *ListPublicSubnetRequest) (*ListPublicSubnetResponse, error)
	// Creates a public subnet in the specified folder.
	Create(context.Context, *CreatePublicSubnetRequest) (*operation.Operation, error)
	// Updates the specified public subnet.
	Update(context.Context, *UpdatePublicSubnetRequest) (*operation.Operation, error)
	// Deletes the specified public subnet.
	//
	// Deleting a public subnet removes its data permanently and is irreversible.
	Delete(context.Context, *DeletePublicSubnetRequest) (*operation.Operation, error)
	// Lists operations for the specified public subnet.
	ListOperations(context.Context, *ListPublicSubnetOperationsRequest) (*ListPublicSubnetOperationsResponse, error)
}

// UnimplementedPublicSubnetServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPublicSubnetServiceServer struct{}

func (UnimplementedPublicSubnetServiceServer) Get(context.Context, *GetPublicSubnetRequest) (*PublicSubnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPublicSubnetServiceServer) List(context.Context, *ListPublicSubnetRequest) (*ListPublicSubnetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedPublicSubnetServiceServer) Create(context.Context, *CreatePublicSubnetRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedPublicSubnetServiceServer) Update(context.Context, *UpdatePublicSubnetRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedPublicSubnetServiceServer) Delete(context.Context, *DeletePublicSubnetRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPublicSubnetServiceServer) ListOperations(context.Context, *ListPublicSubnetOperationsRequest) (*ListPublicSubnetOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedPublicSubnetServiceServer) testEmbeddedByValue() {}

// UnsafePublicSubnetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PublicSubnetServiceServer will
// result in compilation errors.
type UnsafePublicSubnetServiceServer interface {
	mustEmbedUnimplementedPublicSubnetServiceServer()
}

func RegisterPublicSubnetServiceServer(s grpc.ServiceRegistrar, srv PublicSubnetServiceServer) {
	// If the following call pancis, it indicates UnimplementedPublicSubnetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PublicSubnetService_ServiceDesc, srv)
}

func _PublicSubnetService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicSubnetServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicSubnetService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicSubnetServiceServer).Get(ctx, req.(*GetPublicSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicSubnetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicSubnetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicSubnetService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicSubnetServiceServer).List(ctx, req.(*ListPublicSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicSubnetService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicSubnetServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicSubnetService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicSubnetServiceServer).Create(ctx, req.(*CreatePublicSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicSubnetService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublicSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicSubnetServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicSubnetService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicSubnetServiceServer).Update(ctx, req.(*UpdatePublicSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicSubnetService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePublicSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicSubnetServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicSubnetService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicSubnetServiceServer).Delete(ctx, req.(*DeletePublicSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PublicSubnetService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPublicSubnetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublicSubnetServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PublicSubnetService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublicSubnetServiceServer).ListOperations(ctx, req.(*ListPublicSubnetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PublicSubnetService_ServiceDesc is the grpc.ServiceDesc for PublicSubnetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PublicSubnetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.baremetal.v1alpha.PublicSubnetService",
	HandlerType: (*PublicSubnetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _PublicSubnetService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _PublicSubnetService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _PublicSubnetService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PublicSubnetService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PublicSubnetService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _PublicSubnetService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/baremetal/v1alpha/public_subnet_service.proto",
}
