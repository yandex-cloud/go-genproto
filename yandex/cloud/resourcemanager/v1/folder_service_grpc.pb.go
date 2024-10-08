// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/resourcemanager/v1/folder_service.proto

package resourcemanager

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	FolderService_Get_FullMethodName                  = "/yandex.cloud.resourcemanager.v1.FolderService/Get"
	FolderService_List_FullMethodName                 = "/yandex.cloud.resourcemanager.v1.FolderService/List"
	FolderService_Create_FullMethodName               = "/yandex.cloud.resourcemanager.v1.FolderService/Create"
	FolderService_Update_FullMethodName               = "/yandex.cloud.resourcemanager.v1.FolderService/Update"
	FolderService_Delete_FullMethodName               = "/yandex.cloud.resourcemanager.v1.FolderService/Delete"
	FolderService_ListOperations_FullMethodName       = "/yandex.cloud.resourcemanager.v1.FolderService/ListOperations"
	FolderService_ListAccessBindings_FullMethodName   = "/yandex.cloud.resourcemanager.v1.FolderService/ListAccessBindings"
	FolderService_SetAccessBindings_FullMethodName    = "/yandex.cloud.resourcemanager.v1.FolderService/SetAccessBindings"
	FolderService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.resourcemanager.v1.FolderService/UpdateAccessBindings"
)

// FolderServiceClient is the client API for FolderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Folder resources.
type FolderServiceClient interface {
	// Returns the specified Folder resource.
	//
	// To get the list of available Folder resources, make a [List] request.
	Get(ctx context.Context, in *GetFolderRequest, opts ...grpc.CallOption) (*Folder, error)
	// Retrieves the list of Folder resources in the specified cloud.
	List(ctx context.Context, in *ListFoldersRequest, opts ...grpc.CallOption) (*ListFoldersResponse, error)
	// Creates a folder in the specified cloud.
	Create(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified folder.
	Update(ctx context.Context, in *UpdateFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified folder.
	Delete(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified folder.
	ListOperations(ctx context.Context, in *ListFolderOperationsRequest, opts ...grpc.CallOption) (*ListFolderOperationsResponse, error)
	// Lists access bindings for the specified folder.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified folder.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified folder.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type folderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFolderServiceClient(cc grpc.ClientConnInterface) FolderServiceClient {
	return &folderServiceClient{cc}
}

func (c *folderServiceClient) Get(ctx context.Context, in *GetFolderRequest, opts ...grpc.CallOption) (*Folder, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Folder)
	err := c.cc.Invoke(ctx, FolderService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) List(ctx context.Context, in *ListFoldersRequest, opts ...grpc.CallOption) (*ListFoldersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFoldersResponse)
	err := c.cc.Invoke(ctx, FolderService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) Create(ctx context.Context, in *CreateFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FolderService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) Update(ctx context.Context, in *UpdateFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FolderService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) Delete(ctx context.Context, in *DeleteFolderRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FolderService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) ListOperations(ctx context.Context, in *ListFolderOperationsRequest, opts ...grpc.CallOption) (*ListFolderOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListFolderOperationsResponse)
	err := c.cc.Invoke(ctx, FolderService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, FolderService_ListAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FolderService_SetAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *folderServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, FolderService_UpdateAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FolderServiceServer is the server API for FolderService service.
// All implementations should embed UnimplementedFolderServiceServer
// for forward compatibility.
//
// A set of methods for managing Folder resources.
type FolderServiceServer interface {
	// Returns the specified Folder resource.
	//
	// To get the list of available Folder resources, make a [List] request.
	Get(context.Context, *GetFolderRequest) (*Folder, error)
	// Retrieves the list of Folder resources in the specified cloud.
	List(context.Context, *ListFoldersRequest) (*ListFoldersResponse, error)
	// Creates a folder in the specified cloud.
	Create(context.Context, *CreateFolderRequest) (*operation.Operation, error)
	// Updates the specified folder.
	Update(context.Context, *UpdateFolderRequest) (*operation.Operation, error)
	// Deletes the specified folder.
	Delete(context.Context, *DeleteFolderRequest) (*operation.Operation, error)
	// Lists operations for the specified folder.
	ListOperations(context.Context, *ListFolderOperationsRequest) (*ListFolderOperationsResponse, error)
	// Lists access bindings for the specified folder.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified folder.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified folder.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedFolderServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedFolderServiceServer struct{}

func (UnimplementedFolderServiceServer) Get(context.Context, *GetFolderRequest) (*Folder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedFolderServiceServer) List(context.Context, *ListFoldersRequest) (*ListFoldersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedFolderServiceServer) Create(context.Context, *CreateFolderRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedFolderServiceServer) Update(context.Context, *UpdateFolderRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFolderServiceServer) Delete(context.Context, *DeleteFolderRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFolderServiceServer) ListOperations(context.Context, *ListFolderOperationsRequest) (*ListFolderOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedFolderServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedFolderServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedFolderServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedFolderServiceServer) testEmbeddedByValue() {}

// UnsafeFolderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FolderServiceServer will
// result in compilation errors.
type UnsafeFolderServiceServer interface {
	mustEmbedUnimplementedFolderServiceServer()
}

func RegisterFolderServiceServer(s grpc.ServiceRegistrar, srv FolderServiceServer) {
	// If the following call pancis, it indicates UnimplementedFolderServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&FolderService_ServiceDesc, srv)
}

func _FolderService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).Get(ctx, req.(*GetFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFoldersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).List(ctx, req.(*ListFoldersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).Create(ctx, req.(*CreateFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).Update(ctx, req.(*UpdateFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFolderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).Delete(ctx, req.(*DeleteFolderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFolderOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).ListOperations(ctx, req.(*ListFolderOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FolderService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FolderServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: FolderService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FolderServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FolderService_ServiceDesc is the grpc.ServiceDesc for FolderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FolderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.resourcemanager.v1.FolderService",
	HandlerType: (*FolderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _FolderService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _FolderService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _FolderService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _FolderService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _FolderService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _FolderService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _FolderService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _FolderService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _FolderService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/resourcemanager/v1/folder_service.proto",
}
