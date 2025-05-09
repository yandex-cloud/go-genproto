// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/k8s/v1/cluster_service.proto

package k8s

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
	ClusterService_Get_FullMethodName                   = "/yandex.cloud.k8s.v1.ClusterService/Get"
	ClusterService_List_FullMethodName                  = "/yandex.cloud.k8s.v1.ClusterService/List"
	ClusterService_Create_FullMethodName                = "/yandex.cloud.k8s.v1.ClusterService/Create"
	ClusterService_Update_FullMethodName                = "/yandex.cloud.k8s.v1.ClusterService/Update"
	ClusterService_Delete_FullMethodName                = "/yandex.cloud.k8s.v1.ClusterService/Delete"
	ClusterService_Stop_FullMethodName                  = "/yandex.cloud.k8s.v1.ClusterService/Stop"
	ClusterService_Start_FullMethodName                 = "/yandex.cloud.k8s.v1.ClusterService/Start"
	ClusterService_RescheduleMaintenance_FullMethodName = "/yandex.cloud.k8s.v1.ClusterService/RescheduleMaintenance"
	ClusterService_ListNodeGroups_FullMethodName        = "/yandex.cloud.k8s.v1.ClusterService/ListNodeGroups"
	ClusterService_ListOperations_FullMethodName        = "/yandex.cloud.k8s.v1.ClusterService/ListOperations"
	ClusterService_ListNodes_FullMethodName             = "/yandex.cloud.k8s.v1.ClusterService/ListNodes"
	ClusterService_ListAccessBindings_FullMethodName    = "/yandex.cloud.k8s.v1.ClusterService/ListAccessBindings"
	ClusterService_SetAccessBindings_FullMethodName     = "/yandex.cloud.k8s.v1.ClusterService/SetAccessBindings"
	ClusterService_UpdateAccessBindings_FullMethodName  = "/yandex.cloud.k8s.v1.ClusterService/UpdateAccessBindings"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Kubernetes cluster.
type ClusterServiceClient interface {
	// Returns the specified Kubernetes cluster.
	//
	// To get the list of available Kubernetes cluster, make a [List] request.
	Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Retrieves the list of Kubernetes cluster in the specified folder.
	List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates a Kubernetes cluster in the specified folder.
	Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified Kubernetes cluster.
	Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified Kubernetes cluster.
	Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops the specified Kubernetes cluster.
	Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the specified Kubernetes cluster.
	Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Reschedules mandatory maintenance for the specified cluster.
	RescheduleMaintenance(ctx context.Context, in *RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists nodegroup for the specified Kubernetes cluster.
	ListNodeGroups(ctx context.Context, in *ListClusterNodeGroupsRequest, opts ...grpc.CallOption) (*ListClusterNodeGroupsResponse, error)
	// Lists operations for the specified Kubernetes cluster.
	ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error)
	// Lists cluster's nodes.
	ListNodes(ctx context.Context, in *ListClusterNodesRequest, opts ...grpc.CallOption) (*ListClusterNodesResponse, error)
	// Lists cluster's access bindings
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets cluster's access bindings
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates cluster's access bindings
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, ClusterService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Stop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Start_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) RescheduleMaintenance(ctx context.Context, in *RescheduleMaintenanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_RescheduleMaintenance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListNodeGroups(ctx context.Context, in *ListClusterNodeGroupsRequest, opts ...grpc.CallOption) (*ListClusterNodeGroupsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListClusterNodeGroupsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListNodeGroups_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListClusterOperationsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListNodes(ctx context.Context, in *ListClusterNodesRequest, opts ...grpc.CallOption) (*ListClusterNodesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListClusterNodesResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListNodes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_SetAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_UpdateAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility.
//
// A set of methods for managing Kubernetes cluster.
type ClusterServiceServer interface {
	// Returns the specified Kubernetes cluster.
	//
	// To get the list of available Kubernetes cluster, make a [List] request.
	Get(context.Context, *GetClusterRequest) (*Cluster, error)
	// Retrieves the list of Kubernetes cluster in the specified folder.
	List(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates a Kubernetes cluster in the specified folder.
	Create(context.Context, *CreateClusterRequest) (*operation.Operation, error)
	// Updates the specified Kubernetes cluster.
	Update(context.Context, *UpdateClusterRequest) (*operation.Operation, error)
	// Deletes the specified Kubernetes cluster.
	Delete(context.Context, *DeleteClusterRequest) (*operation.Operation, error)
	// Stops the specified Kubernetes cluster.
	Stop(context.Context, *StopClusterRequest) (*operation.Operation, error)
	// Starts the specified Kubernetes cluster.
	Start(context.Context, *StartClusterRequest) (*operation.Operation, error)
	// Reschedules mandatory maintenance for the specified cluster.
	RescheduleMaintenance(context.Context, *RescheduleMaintenanceRequest) (*operation.Operation, error)
	// Lists nodegroup for the specified Kubernetes cluster.
	ListNodeGroups(context.Context, *ListClusterNodeGroupsRequest) (*ListClusterNodeGroupsResponse, error)
	// Lists operations for the specified Kubernetes cluster.
	ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error)
	// Lists cluster's nodes.
	ListNodes(context.Context, *ListClusterNodesRequest) (*ListClusterNodesResponse, error)
	// Lists cluster's access bindings
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets cluster's access bindings
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates cluster's access bindings
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedClusterServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedClusterServiceServer struct{}

func (UnimplementedClusterServiceServer) Get(context.Context, *GetClusterRequest) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedClusterServiceServer) List(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedClusterServiceServer) Create(context.Context, *CreateClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedClusterServiceServer) Update(context.Context, *UpdateClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedClusterServiceServer) Delete(context.Context, *DeleteClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClusterServiceServer) Stop(context.Context, *StopClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedClusterServiceServer) Start(context.Context, *StartClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedClusterServiceServer) RescheduleMaintenance(context.Context, *RescheduleMaintenanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RescheduleMaintenance not implemented")
}
func (UnimplementedClusterServiceServer) ListNodeGroups(context.Context, *ListClusterNodeGroupsRequest) (*ListClusterNodeGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodeGroups not implemented")
}
func (UnimplementedClusterServiceServer) ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedClusterServiceServer) ListNodes(context.Context, *ListClusterNodesRequest) (*ListClusterNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedClusterServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedClusterServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedClusterServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedClusterServiceServer) testEmbeddedByValue() {}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
	// If the following call pancis, it indicates UnimplementedClusterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ClusterService_ServiceDesc, srv)
}

func _ClusterService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Get(ctx, req.(*GetClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).List(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Create(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Update(ctx, req.(*UpdateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Delete(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Stop(ctx, req.(*StopClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Start(ctx, req.(*StartClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_RescheduleMaintenance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RescheduleMaintenanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).RescheduleMaintenance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_RescheduleMaintenance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).RescheduleMaintenance(ctx, req.(*RescheduleMaintenanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListNodeGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterNodeGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListNodeGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListNodeGroups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListNodeGroups(ctx, req.(*ListClusterNodeGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListOperations(ctx, req.(*ListClusterOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListNodes(ctx, req.(*ListClusterNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.k8s.v1.ClusterService",
	HandlerType: (*ClusterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ClusterService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ClusterService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ClusterService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ClusterService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterService_Delete_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ClusterService_Stop_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ClusterService_Start_Handler,
		},
		{
			MethodName: "RescheduleMaintenance",
			Handler:    _ClusterService_RescheduleMaintenance_Handler,
		},
		{
			MethodName: "ListNodeGroups",
			Handler:    _ClusterService_ListNodeGroups_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ClusterService_ListOperations_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _ClusterService_ListNodes_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _ClusterService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _ClusterService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _ClusterService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/k8s/v1/cluster_service.proto",
}
