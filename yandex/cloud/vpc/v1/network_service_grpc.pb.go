// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/vpc/v1/network_service.proto

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	NetworkService_Get_FullMethodName                = "/yandex.cloud.vpc.v1.NetworkService/Get"
	NetworkService_List_FullMethodName               = "/yandex.cloud.vpc.v1.NetworkService/List"
	NetworkService_Create_FullMethodName             = "/yandex.cloud.vpc.v1.NetworkService/Create"
	NetworkService_Update_FullMethodName             = "/yandex.cloud.vpc.v1.NetworkService/Update"
	NetworkService_Delete_FullMethodName             = "/yandex.cloud.vpc.v1.NetworkService/Delete"
	NetworkService_ListSubnets_FullMethodName        = "/yandex.cloud.vpc.v1.NetworkService/ListSubnets"
	NetworkService_ListSecurityGroups_FullMethodName = "/yandex.cloud.vpc.v1.NetworkService/ListSecurityGroups"
	NetworkService_ListRouteTables_FullMethodName    = "/yandex.cloud.vpc.v1.NetworkService/ListRouteTables"
	NetworkService_ListOperations_FullMethodName     = "/yandex.cloud.vpc.v1.NetworkService/ListOperations"
	NetworkService_Move_FullMethodName               = "/yandex.cloud.vpc.v1.NetworkService/Move"
)

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Network resources.
type NetworkServiceClient interface {
	// Returns the specified Network resource.
	//
	// Get the list of available Network resources by making a [List] request.
	Get(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error)
	// Retrieves the list of Network resources in the specified folder.
	List(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error)
	// Creates a network in the specified folder using the data specified in the request.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified network.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Update(ctx context.Context, in *UpdateNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified network.
	Delete(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists subnets from the specified network.
	ListSubnets(ctx context.Context, in *ListNetworkSubnetsRequest, opts ...grpc.CallOption) (*ListNetworkSubnetsResponse, error)
	// Lists security groups from the specified network.
	ListSecurityGroups(ctx context.Context, in *ListNetworkSecurityGroupsRequest, opts ...grpc.CallOption) (*ListNetworkSecurityGroupsResponse, error)
	// Lists route tables from the specified network.
	ListRouteTables(ctx context.Context, in *ListNetworkRouteTablesRequest, opts ...grpc.CallOption) (*ListNetworkRouteTablesResponse, error)
	// Lists operations for the specified network.
	ListOperations(ctx context.Context, in *ListNetworkOperationsRequest, opts ...grpc.CallOption) (*ListNetworkOperationsResponse, error)
	// Move network to another folder.
	Move(ctx context.Context, in *MoveNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type networkServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNetworkServiceClient(cc grpc.ClientConnInterface) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) Get(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*Network, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Network)
	err := c.cc.Invoke(ctx, NetworkService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) List(ctx context.Context, in *ListNetworksRequest, opts ...grpc.CallOption) (*ListNetworksResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNetworksResponse)
	err := c.cc.Invoke(ctx, NetworkService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Create(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, NetworkService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Update(ctx context.Context, in *UpdateNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, NetworkService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Delete(ctx context.Context, in *DeleteNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, NetworkService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ListSubnets(ctx context.Context, in *ListNetworkSubnetsRequest, opts ...grpc.CallOption) (*ListNetworkSubnetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNetworkSubnetsResponse)
	err := c.cc.Invoke(ctx, NetworkService_ListSubnets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ListSecurityGroups(ctx context.Context, in *ListNetworkSecurityGroupsRequest, opts ...grpc.CallOption) (*ListNetworkSecurityGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNetworkSecurityGroupsResponse)
	err := c.cc.Invoke(ctx, NetworkService_ListSecurityGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ListRouteTables(ctx context.Context, in *ListNetworkRouteTablesRequest, opts ...grpc.CallOption) (*ListNetworkRouteTablesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNetworkRouteTablesResponse)
	err := c.cc.Invoke(ctx, NetworkService_ListRouteTables_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ListOperations(ctx context.Context, in *ListNetworkOperationsRequest, opts ...grpc.CallOption) (*ListNetworkOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListNetworkOperationsResponse)
	err := c.cc.Invoke(ctx, NetworkService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) Move(ctx context.Context, in *MoveNetworkRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, NetworkService_Move_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
// All implementations should embed UnimplementedNetworkServiceServer
// for forward compatibility.
//
// A set of methods for managing Network resources.
type NetworkServiceServer interface {
	// Returns the specified Network resource.
	//
	// Get the list of available Network resources by making a [List] request.
	Get(context.Context, *GetNetworkRequest) (*Network, error)
	// Retrieves the list of Network resources in the specified folder.
	List(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error)
	// Creates a network in the specified folder using the data specified in the request.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(context.Context, *CreateNetworkRequest) (*operation.Operation, error)
	// Updates the specified network.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Update(context.Context, *UpdateNetworkRequest) (*operation.Operation, error)
	// Deletes the specified network.
	Delete(context.Context, *DeleteNetworkRequest) (*operation.Operation, error)
	// Lists subnets from the specified network.
	ListSubnets(context.Context, *ListNetworkSubnetsRequest) (*ListNetworkSubnetsResponse, error)
	// Lists security groups from the specified network.
	ListSecurityGroups(context.Context, *ListNetworkSecurityGroupsRequest) (*ListNetworkSecurityGroupsResponse, error)
	// Lists route tables from the specified network.
	ListRouteTables(context.Context, *ListNetworkRouteTablesRequest) (*ListNetworkRouteTablesResponse, error)
	// Lists operations for the specified network.
	ListOperations(context.Context, *ListNetworkOperationsRequest) (*ListNetworkOperationsResponse, error)
	// Move network to another folder.
	Move(context.Context, *MoveNetworkRequest) (*operation.Operation, error)
}

// UnimplementedNetworkServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNetworkServiceServer struct{}

func (UnimplementedNetworkServiceServer) Get(context.Context, *GetNetworkRequest) (*Network, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedNetworkServiceServer) List(context.Context, *ListNetworksRequest) (*ListNetworksResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedNetworkServiceServer) Create(context.Context, *CreateNetworkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedNetworkServiceServer) Update(context.Context, *UpdateNetworkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedNetworkServiceServer) Delete(context.Context, *DeleteNetworkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedNetworkServiceServer) ListSubnets(context.Context, *ListNetworkSubnetsRequest) (*ListNetworkSubnetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSubnets not implemented")
}
func (UnimplementedNetworkServiceServer) ListSecurityGroups(context.Context, *ListNetworkSecurityGroupsRequest) (*ListNetworkSecurityGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSecurityGroups not implemented")
}
func (UnimplementedNetworkServiceServer) ListRouteTables(context.Context, *ListNetworkRouteTablesRequest) (*ListNetworkRouteTablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRouteTables not implemented")
}
func (UnimplementedNetworkServiceServer) ListOperations(context.Context, *ListNetworkOperationsRequest) (*ListNetworkOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedNetworkServiceServer) Move(context.Context, *MoveNetworkRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Move not implemented")
}
func (UnimplementedNetworkServiceServer) testEmbeddedByValue() {}

// UnsafeNetworkServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NetworkServiceServer will
// result in compilation errors.
type UnsafeNetworkServiceServer interface {
	mustEmbedUnimplementedNetworkServiceServer()
}

func RegisterNetworkServiceServer(s grpc.ServiceRegistrar, srv NetworkServiceServer) {
	// If the following call pancis, it indicates UnimplementedNetworkServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&NetworkService_ServiceDesc, srv)
}

func _NetworkService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Get(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).List(ctx, req.(*ListNetworksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Create(ctx, req.(*CreateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Update(ctx, req.(*UpdateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Delete(ctx, req.(*DeleteNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ListSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkSubnetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_ListSubnets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListSubnets(ctx, req.(*ListNetworkSubnetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ListSecurityGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkSecurityGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListSecurityGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_ListSecurityGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListSecurityGroups(ctx, req.(*ListNetworkSecurityGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ListRouteTables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkRouteTablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListRouteTables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_ListRouteTables_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListRouteTables(ctx, req.(*ListNetworkRouteTablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNetworkOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ListOperations(ctx, req.(*ListNetworkOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_Move_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MoveNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).Move(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NetworkService_Move_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).Move(ctx, req.(*MoveNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NetworkService_ServiceDesc is the grpc.ServiceDesc for NetworkService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NetworkService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.vpc.v1.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _NetworkService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NetworkService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _NetworkService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _NetworkService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _NetworkService_Delete_Handler,
		},
		{
			MethodName: "ListSubnets",
			Handler:    _NetworkService_ListSubnets_Handler,
		},
		{
			MethodName: "ListSecurityGroups",
			Handler:    _NetworkService_ListSecurityGroups_Handler,
		},
		{
			MethodName: "ListRouteTables",
			Handler:    _NetworkService_ListRouteTables_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _NetworkService_ListOperations_Handler,
		},
		{
			MethodName: "Move",
			Handler:    _NetworkService_Move_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/vpc/v1/network_service.proto",
}
