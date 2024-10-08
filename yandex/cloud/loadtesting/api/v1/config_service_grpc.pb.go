// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/loadtesting/api/v1/config_service.proto

package loadtesting

import (
	context "context"
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/config"
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
	ConfigService_Create_FullMethodName = "/yandex.cloud.loadtesting.api.v1.ConfigService/Create"
	ConfigService_Get_FullMethodName    = "/yandex.cloud.loadtesting.api.v1.ConfigService/Get"
	ConfigService_List_FullMethodName   = "/yandex.cloud.loadtesting.api.v1.ConfigService/List"
	ConfigService_Delete_FullMethodName = "/yandex.cloud.loadtesting.api.v1.ConfigService/Delete"
)

// ConfigServiceClient is the client API for ConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing test configurations.
type ConfigServiceClient interface {
	// Creates a test config in the specified folder.
	Create(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified config.
	//
	// To get the list of all available configs, make a [List] request.
	Get(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*config.Config, error)
	// Retrieves the list of configs in the specified folder.
	List(ctx context.Context, in *ListConfigsRequest, opts ...grpc.CallOption) (*ListConfigsResponse, error)
	// Deletes the specified config.
	Delete(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type configServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConfigServiceClient(cc grpc.ClientConnInterface) ConfigServiceClient {
	return &configServiceClient{cc}
}

func (c *configServiceClient) Create(ctx context.Context, in *CreateConfigRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ConfigService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) Get(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*config.Config, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(config.Config)
	err := c.cc.Invoke(ctx, ConfigService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) List(ctx context.Context, in *ListConfigsRequest, opts ...grpc.CallOption) (*ListConfigsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListConfigsResponse)
	err := c.cc.Invoke(ctx, ConfigService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configServiceClient) Delete(ctx context.Context, in *DeleteConfigRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ConfigService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigServiceServer is the server API for ConfigService service.
// All implementations should embed UnimplementedConfigServiceServer
// for forward compatibility.
//
// A set of methods for managing test configurations.
type ConfigServiceServer interface {
	// Creates a test config in the specified folder.
	Create(context.Context, *CreateConfigRequest) (*operation.Operation, error)
	// Returns the specified config.
	//
	// To get the list of all available configs, make a [List] request.
	Get(context.Context, *GetConfigRequest) (*config.Config, error)
	// Retrieves the list of configs in the specified folder.
	List(context.Context, *ListConfigsRequest) (*ListConfigsResponse, error)
	// Deletes the specified config.
	Delete(context.Context, *DeleteConfigRequest) (*operation.Operation, error)
}

// UnimplementedConfigServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedConfigServiceServer struct{}

func (UnimplementedConfigServiceServer) Create(context.Context, *CreateConfigRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedConfigServiceServer) Get(context.Context, *GetConfigRequest) (*config.Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConfigServiceServer) List(context.Context, *ListConfigsRequest) (*ListConfigsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedConfigServiceServer) Delete(context.Context, *DeleteConfigRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedConfigServiceServer) testEmbeddedByValue() {}

// UnsafeConfigServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConfigServiceServer will
// result in compilation errors.
type UnsafeConfigServiceServer interface {
	mustEmbedUnimplementedConfigServiceServer()
}

func RegisterConfigServiceServer(s grpc.ServiceRegistrar, srv ConfigServiceServer) {
	// If the following call pancis, it indicates UnimplementedConfigServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ConfigService_ServiceDesc, srv)
}

func _ConfigService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Create(ctx, req.(*CreateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Get(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConfigsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).List(ctx, req.(*ListConfigsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ConfigService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigServiceServer).Delete(ctx, req.(*DeleteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConfigService_ServiceDesc is the grpc.ServiceDesc for ConfigService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConfigService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.loadtesting.api.v1.ConfigService",
	HandlerType: (*ConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ConfigService_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ConfigService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ConfigService_List_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ConfigService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/loadtesting/api/v1/config_service.proto",
}
