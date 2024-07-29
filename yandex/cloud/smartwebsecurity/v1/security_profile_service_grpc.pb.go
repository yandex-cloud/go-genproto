// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/smartwebsecurity/v1/security_profile_service.proto

package smartwebsecurity

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
	SecurityProfileService_Get_FullMethodName    = "/yandex.cloud.smartwebsecurity.v1.SecurityProfileService/Get"
	SecurityProfileService_List_FullMethodName   = "/yandex.cloud.smartwebsecurity.v1.SecurityProfileService/List"
	SecurityProfileService_Create_FullMethodName = "/yandex.cloud.smartwebsecurity.v1.SecurityProfileService/Create"
	SecurityProfileService_Update_FullMethodName = "/yandex.cloud.smartwebsecurity.v1.SecurityProfileService/Update"
	SecurityProfileService_Delete_FullMethodName = "/yandex.cloud.smartwebsecurity.v1.SecurityProfileService/Delete"
)

// SecurityProfileServiceClient is the client API for SecurityProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SecurityProfileServiceClient interface {
	// Returns the specified SecurityProfile resource.
	Get(ctx context.Context, in *GetSecurityProfileRequest, opts ...grpc.CallOption) (*SecurityProfile, error)
	// Retrieves the list of SecurityProfile resources in the specified folder.
	List(ctx context.Context, in *ListSecurityProfilesRequest, opts ...grpc.CallOption) (*ListSecurityProfilesResponse, error)
	// Creates a security profile in the specified folder using the data specified in the request.
	Create(ctx context.Context, in *CreateSecurityProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified security profile.
	Update(ctx context.Context, in *UpdateSecurityProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified security profile.
	Delete(ctx context.Context, in *DeleteSecurityProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type securityProfileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSecurityProfileServiceClient(cc grpc.ClientConnInterface) SecurityProfileServiceClient {
	return &securityProfileServiceClient{cc}
}

func (c *securityProfileServiceClient) Get(ctx context.Context, in *GetSecurityProfileRequest, opts ...grpc.CallOption) (*SecurityProfile, error) {
	out := new(SecurityProfile)
	err := c.cc.Invoke(ctx, SecurityProfileService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityProfileServiceClient) List(ctx context.Context, in *ListSecurityProfilesRequest, opts ...grpc.CallOption) (*ListSecurityProfilesResponse, error) {
	out := new(ListSecurityProfilesResponse)
	err := c.cc.Invoke(ctx, SecurityProfileService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityProfileServiceClient) Create(ctx context.Context, in *CreateSecurityProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityProfileService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityProfileServiceClient) Update(ctx context.Context, in *UpdateSecurityProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityProfileService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *securityProfileServiceClient) Delete(ctx context.Context, in *DeleteSecurityProfileRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, SecurityProfileService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecurityProfileServiceServer is the server API for SecurityProfileService service.
// All implementations should embed UnimplementedSecurityProfileServiceServer
// for forward compatibility
type SecurityProfileServiceServer interface {
	// Returns the specified SecurityProfile resource.
	Get(context.Context, *GetSecurityProfileRequest) (*SecurityProfile, error)
	// Retrieves the list of SecurityProfile resources in the specified folder.
	List(context.Context, *ListSecurityProfilesRequest) (*ListSecurityProfilesResponse, error)
	// Creates a security profile in the specified folder using the data specified in the request.
	Create(context.Context, *CreateSecurityProfileRequest) (*operation.Operation, error)
	// Updates the specified security profile.
	Update(context.Context, *UpdateSecurityProfileRequest) (*operation.Operation, error)
	// Deletes the specified security profile.
	Delete(context.Context, *DeleteSecurityProfileRequest) (*operation.Operation, error)
}

// UnimplementedSecurityProfileServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSecurityProfileServiceServer struct {
}

func (UnimplementedSecurityProfileServiceServer) Get(context.Context, *GetSecurityProfileRequest) (*SecurityProfile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedSecurityProfileServiceServer) List(context.Context, *ListSecurityProfilesRequest) (*ListSecurityProfilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSecurityProfileServiceServer) Create(context.Context, *CreateSecurityProfileRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSecurityProfileServiceServer) Update(context.Context, *UpdateSecurityProfileRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSecurityProfileServiceServer) Delete(context.Context, *DeleteSecurityProfileRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeSecurityProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SecurityProfileServiceServer will
// result in compilation errors.
type UnsafeSecurityProfileServiceServer interface {
	mustEmbedUnimplementedSecurityProfileServiceServer()
}

func RegisterSecurityProfileServiceServer(s grpc.ServiceRegistrar, srv SecurityProfileServiceServer) {
	s.RegisterService(&SecurityProfileService_ServiceDesc, srv)
}

func _SecurityProfileService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSecurityProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityProfileServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityProfileService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityProfileServiceServer).Get(ctx, req.(*GetSecurityProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityProfileService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecurityProfilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityProfileServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityProfileService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityProfileServiceServer).List(ctx, req.(*ListSecurityProfilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityProfileService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSecurityProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityProfileServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityProfileService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityProfileServiceServer).Create(ctx, req.(*CreateSecurityProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityProfileService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSecurityProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityProfileServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityProfileService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityProfileServiceServer).Update(ctx, req.(*UpdateSecurityProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecurityProfileService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSecurityProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecurityProfileServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SecurityProfileService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecurityProfileServiceServer).Delete(ctx, req.(*DeleteSecurityProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SecurityProfileService_ServiceDesc is the grpc.ServiceDesc for SecurityProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SecurityProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.smartwebsecurity.v1.SecurityProfileService",
	HandlerType: (*SecurityProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _SecurityProfileService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _SecurityProfileService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _SecurityProfileService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SecurityProfileService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SecurityProfileService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/smartwebsecurity/v1/security_profile_service.proto",
}