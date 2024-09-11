// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/iam/v1/workload/federated_credential_service.proto

package workload

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
	FederatedCredentialService_Get_FullMethodName    = "/yandex.cloud.iam.v1.workload.FederatedCredentialService/Get"
	FederatedCredentialService_List_FullMethodName   = "/yandex.cloud.iam.v1.workload.FederatedCredentialService/List"
	FederatedCredentialService_Create_FullMethodName = "/yandex.cloud.iam.v1.workload.FederatedCredentialService/Create"
	FederatedCredentialService_Delete_FullMethodName = "/yandex.cloud.iam.v1.workload.FederatedCredentialService/Delete"
)

// FederatedCredentialServiceClient is the client API for FederatedCredentialService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FederatedCredentialServiceClient interface {
	// Returns the specified federated credential.
	//
	// To get the list of available federated credentials, make a [List] request.
	Get(ctx context.Context, in *GetFederatedCredentialRequest, opts ...grpc.CallOption) (*FederatedCredential, error)
	// Retrieves the list of federated credentials for the specified service account.
	List(ctx context.Context, in *ListFederatedCredentialsRequest, opts ...grpc.CallOption) (*ListFederatedCredentialsResponse, error)
	// Creates a federated credential for the specified service account.
	Create(ctx context.Context, in *CreateFederatedCredentialRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified federated credential.
	Delete(ctx context.Context, in *DeleteFederatedCredentialRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type federatedCredentialServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFederatedCredentialServiceClient(cc grpc.ClientConnInterface) FederatedCredentialServiceClient {
	return &federatedCredentialServiceClient{cc}
}

func (c *federatedCredentialServiceClient) Get(ctx context.Context, in *GetFederatedCredentialRequest, opts ...grpc.CallOption) (*FederatedCredential, error) {
	out := new(FederatedCredential)
	err := c.cc.Invoke(ctx, FederatedCredentialService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialServiceClient) List(ctx context.Context, in *ListFederatedCredentialsRequest, opts ...grpc.CallOption) (*ListFederatedCredentialsResponse, error) {
	out := new(ListFederatedCredentialsResponse)
	err := c.cc.Invoke(ctx, FederatedCredentialService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialServiceClient) Create(ctx context.Context, in *CreateFederatedCredentialRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FederatedCredentialService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *federatedCredentialServiceClient) Delete(ctx context.Context, in *DeleteFederatedCredentialRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FederatedCredentialService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FederatedCredentialServiceServer is the server API for FederatedCredentialService service.
// All implementations should embed UnimplementedFederatedCredentialServiceServer
// for forward compatibility
type FederatedCredentialServiceServer interface {
	// Returns the specified federated credential.
	//
	// To get the list of available federated credentials, make a [List] request.
	Get(context.Context, *GetFederatedCredentialRequest) (*FederatedCredential, error)
	// Retrieves the list of federated credentials for the specified service account.
	List(context.Context, *ListFederatedCredentialsRequest) (*ListFederatedCredentialsResponse, error)
	// Creates a federated credential for the specified service account.
	Create(context.Context, *CreateFederatedCredentialRequest) (*operation.Operation, error)
	// Deletes the specified federated credential.
	Delete(context.Context, *DeleteFederatedCredentialRequest) (*operation.Operation, error)
}

// UnimplementedFederatedCredentialServiceServer should be embedded to have forward compatible implementations.
type UnimplementedFederatedCredentialServiceServer struct {
}

func (UnimplementedFederatedCredentialServiceServer) Get(context.Context, *GetFederatedCredentialRequest) (*FederatedCredential, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFederatedCredentialServiceServer) List(context.Context, *ListFederatedCredentialsRequest) (*ListFederatedCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFederatedCredentialServiceServer) Create(context.Context, *CreateFederatedCredentialRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFederatedCredentialServiceServer) Delete(context.Context, *DeleteFederatedCredentialRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeFederatedCredentialServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FederatedCredentialServiceServer will
// result in compilation errors.
type UnsafeFederatedCredentialServiceServer interface {
	mustEmbedUnimplementedFederatedCredentialServiceServer()
}

func RegisterFederatedCredentialServiceServer(s grpc.ServiceRegistrar, srv FederatedCredentialServiceServer) {
	s.RegisterService(&FederatedCredentialService_ServiceDesc, srv)
}

func _FederatedCredentialService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFederatedCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialServiceServer).Get(ctx, req.(*GetFederatedCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFederatedCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialServiceServer).List(ctx, req.(*ListFederatedCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFederatedCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialServiceServer).Create(ctx, req.(*CreateFederatedCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FederatedCredentialService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFederatedCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FederatedCredentialServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FederatedCredentialService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FederatedCredentialServiceServer).Delete(ctx, req.(*DeleteFederatedCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FederatedCredentialService_ServiceDesc is the grpc.ServiceDesc for FederatedCredentialService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FederatedCredentialService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iam.v1.workload.FederatedCredentialService",
	HandlerType: (*FederatedCredentialServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FederatedCredentialService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FederatedCredentialService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FederatedCredentialService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FederatedCredentialService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iam/v1/workload/federated_credential_service.proto",
}