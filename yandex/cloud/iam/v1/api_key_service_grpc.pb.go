// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/iam/v1/api_key_service.proto

package iam

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
	ApiKeyService_List_FullMethodName           = "/yandex.cloud.iam.v1.ApiKeyService/List"
	ApiKeyService_Get_FullMethodName            = "/yandex.cloud.iam.v1.ApiKeyService/Get"
	ApiKeyService_Create_FullMethodName         = "/yandex.cloud.iam.v1.ApiKeyService/Create"
	ApiKeyService_Update_FullMethodName         = "/yandex.cloud.iam.v1.ApiKeyService/Update"
	ApiKeyService_Delete_FullMethodName         = "/yandex.cloud.iam.v1.ApiKeyService/Delete"
	ApiKeyService_ListOperations_FullMethodName = "/yandex.cloud.iam.v1.ApiKeyService/ListOperations"
	ApiKeyService_ListScopes_FullMethodName     = "/yandex.cloud.iam.v1.ApiKeyService/ListScopes"
)

// ApiKeyServiceClient is the client API for ApiKeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing API keys.
type ApiKeyServiceClient interface {
	// Retrieves the list of API keys for the specified service account.
	List(ctx context.Context, in *ListApiKeysRequest, opts ...grpc.CallOption) (*ListApiKeysResponse, error)
	// Returns the specified API key.
	//
	// To get the list of available API keys, make a [List] request.
	Get(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*ApiKey, error)
	// Creates an API key for the specified service account.
	Create(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error)
	// Updates the specified API key.
	Update(ctx context.Context, in *UpdateApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified API key.
	Delete(ctx context.Context, in *DeleteApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of operations for the specified API key.
	ListOperations(ctx context.Context, in *ListApiKeyOperationsRequest, opts ...grpc.CallOption) (*ListApiKeyOperationsResponse, error)
	// Retrieves the list of scopes.
	ListScopes(ctx context.Context, in *ListApiKeyScopesRequest, opts ...grpc.CallOption) (*ListApiKeyScopesResponse, error)
}

type apiKeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewApiKeyServiceClient(cc grpc.ClientConnInterface) ApiKeyServiceClient {
	return &apiKeyServiceClient{cc}
}

func (c *apiKeyServiceClient) List(ctx context.Context, in *ListApiKeysRequest, opts ...grpc.CallOption) (*ListApiKeysResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListApiKeysResponse)
	err := c.cc.Invoke(ctx, ApiKeyService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Get(ctx context.Context, in *GetApiKeyRequest, opts ...grpc.CallOption) (*ApiKey, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ApiKey)
	err := c.cc.Invoke(ctx, ApiKeyService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Create(ctx context.Context, in *CreateApiKeyRequest, opts ...grpc.CallOption) (*CreateApiKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateApiKeyResponse)
	err := c.cc.Invoke(ctx, ApiKeyService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Update(ctx context.Context, in *UpdateApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ApiKeyService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) Delete(ctx context.Context, in *DeleteApiKeyRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ApiKeyService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) ListOperations(ctx context.Context, in *ListApiKeyOperationsRequest, opts ...grpc.CallOption) (*ListApiKeyOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListApiKeyOperationsResponse)
	err := c.cc.Invoke(ctx, ApiKeyService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiKeyServiceClient) ListScopes(ctx context.Context, in *ListApiKeyScopesRequest, opts ...grpc.CallOption) (*ListApiKeyScopesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListApiKeyScopesResponse)
	err := c.cc.Invoke(ctx, ApiKeyService_ListScopes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApiKeyServiceServer is the server API for ApiKeyService service.
// All implementations should embed UnimplementedApiKeyServiceServer
// for forward compatibility.
//
// A set of methods for managing API keys.
type ApiKeyServiceServer interface {
	// Retrieves the list of API keys for the specified service account.
	List(context.Context, *ListApiKeysRequest) (*ListApiKeysResponse, error)
	// Returns the specified API key.
	//
	// To get the list of available API keys, make a [List] request.
	Get(context.Context, *GetApiKeyRequest) (*ApiKey, error)
	// Creates an API key for the specified service account.
	Create(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error)
	// Updates the specified API key.
	Update(context.Context, *UpdateApiKeyRequest) (*operation.Operation, error)
	// Deletes the specified API key.
	Delete(context.Context, *DeleteApiKeyRequest) (*operation.Operation, error)
	// Retrieves the list of operations for the specified API key.
	ListOperations(context.Context, *ListApiKeyOperationsRequest) (*ListApiKeyOperationsResponse, error)
	// Retrieves the list of scopes.
	ListScopes(context.Context, *ListApiKeyScopesRequest) (*ListApiKeyScopesResponse, error)
}

// UnimplementedApiKeyServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedApiKeyServiceServer struct{}

func (UnimplementedApiKeyServiceServer) List(context.Context, *ListApiKeysRequest) (*ListApiKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedApiKeyServiceServer) Get(context.Context, *GetApiKeyRequest) (*ApiKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedApiKeyServiceServer) Create(context.Context, *CreateApiKeyRequest) (*CreateApiKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedApiKeyServiceServer) Update(context.Context, *UpdateApiKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedApiKeyServiceServer) Delete(context.Context, *DeleteApiKeyRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedApiKeyServiceServer) ListOperations(context.Context, *ListApiKeyOperationsRequest) (*ListApiKeyOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedApiKeyServiceServer) ListScopes(context.Context, *ListApiKeyScopesRequest) (*ListApiKeyScopesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListScopes not implemented")
}
func (UnimplementedApiKeyServiceServer) testEmbeddedByValue() {}

// UnsafeApiKeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiKeyServiceServer will
// result in compilation errors.
type UnsafeApiKeyServiceServer interface {
	mustEmbedUnimplementedApiKeyServiceServer()
}

func RegisterApiKeyServiceServer(s grpc.ServiceRegistrar, srv ApiKeyServiceServer) {
	// If the following call pancis, it indicates UnimplementedApiKeyServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ApiKeyService_ServiceDesc, srv)
}

func _ApiKeyService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).List(ctx, req.(*ListApiKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Get(ctx, req.(*GetApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Create(ctx, req.(*CreateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Update(ctx, req.(*UpdateApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteApiKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).Delete(ctx, req.(*DeleteApiKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiKeyOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).ListOperations(ctx, req.(*ListApiKeyOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiKeyService_ListScopes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListApiKeyScopesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiKeyServiceServer).ListScopes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ApiKeyService_ListScopes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiKeyServiceServer).ListScopes(ctx, req.(*ListApiKeyScopesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ApiKeyService_ServiceDesc is the grpc.ServiceDesc for ApiKeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiKeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.ApiKeyService",
	HandlerType: (*ApiKeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _ApiKeyService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ApiKeyService_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ApiKeyService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ApiKeyService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApiKeyService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ApiKeyService_ListOperations_Handler,
		},
		{
			MethodName: "ListScopes",
			Handler:    _ApiKeyService_ListScopes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/api_key_service.proto",
}
