// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/ydb/v1/storage_type_service.proto

package ydb

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
	StorageTypeService_Get_FullMethodName  = "/yandex.cloud.ydb.v1.StorageTypeService/Get"
	StorageTypeService_List_FullMethodName = "/yandex.cloud.ydb.v1.StorageTypeService/List"
)

// StorageTypeServiceClient is the client API for StorageTypeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StorageTypeServiceClient interface {
	// Returns the specified storage types.
	Get(ctx context.Context, in *GetStorageTypeRequest, opts ...grpc.CallOption) (*StorageType, error)
	// Returns the list of available storage types.
	List(ctx context.Context, in *ListStorageTypesRequest, opts ...grpc.CallOption) (*ListStorageTypesResponse, error)
}

type storageTypeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStorageTypeServiceClient(cc grpc.ClientConnInterface) StorageTypeServiceClient {
	return &storageTypeServiceClient{cc}
}

func (c *storageTypeServiceClient) Get(ctx context.Context, in *GetStorageTypeRequest, opts ...grpc.CallOption) (*StorageType, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StorageType)
	err := c.cc.Invoke(ctx, StorageTypeService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storageTypeServiceClient) List(ctx context.Context, in *ListStorageTypesRequest, opts ...grpc.CallOption) (*ListStorageTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListStorageTypesResponse)
	err := c.cc.Invoke(ctx, StorageTypeService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StorageTypeServiceServer is the server API for StorageTypeService service.
// All implementations should embed UnimplementedStorageTypeServiceServer
// for forward compatibility.
type StorageTypeServiceServer interface {
	// Returns the specified storage types.
	Get(context.Context, *GetStorageTypeRequest) (*StorageType, error)
	// Returns the list of available storage types.
	List(context.Context, *ListStorageTypesRequest) (*ListStorageTypesResponse, error)
}

// UnimplementedStorageTypeServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStorageTypeServiceServer struct{}

func (UnimplementedStorageTypeServiceServer) Get(context.Context, *GetStorageTypeRequest) (*StorageType, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStorageTypeServiceServer) List(context.Context, *ListStorageTypesRequest) (*ListStorageTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStorageTypeServiceServer) testEmbeddedByValue() {}

// UnsafeStorageTypeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StorageTypeServiceServer will
// result in compilation errors.
type UnsafeStorageTypeServiceServer interface {
	mustEmbedUnimplementedStorageTypeServiceServer()
}

func RegisterStorageTypeServiceServer(s grpc.ServiceRegistrar, srv StorageTypeServiceServer) {
	// If the following call pancis, it indicates UnimplementedStorageTypeServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&StorageTypeService_ServiceDesc, srv)
}

func _StorageTypeService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStorageTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageTypeServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageTypeService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageTypeServiceServer).Get(ctx, req.(*GetStorageTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StorageTypeService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListStorageTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StorageTypeServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StorageTypeService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StorageTypeServiceServer).List(ctx, req.(*ListStorageTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StorageTypeService_ServiceDesc is the grpc.ServiceDesc for StorageTypeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StorageTypeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ydb.v1.StorageTypeService",
	HandlerType: (*StorageTypeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _StorageTypeService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _StorageTypeService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ydb/v1/storage_type_service.proto",
}
