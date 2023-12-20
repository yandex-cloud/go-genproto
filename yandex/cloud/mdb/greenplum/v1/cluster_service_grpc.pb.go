// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/mdb/greenplum/v1/cluster_service.proto

package greenplum

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
	ClusterService_Get_FullMethodName              = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Get"
	ClusterService_List_FullMethodName             = "/yandex.cloud.mdb.greenplum.v1.ClusterService/List"
	ClusterService_Create_FullMethodName           = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Create"
	ClusterService_Update_FullMethodName           = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Update"
	ClusterService_Expand_FullMethodName           = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Expand"
	ClusterService_Delete_FullMethodName           = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Delete"
	ClusterService_Start_FullMethodName            = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Start"
	ClusterService_Stop_FullMethodName             = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Stop"
	ClusterService_ListOperations_FullMethodName   = "/yandex.cloud.mdb.greenplum.v1.ClusterService/ListOperations"
	ClusterService_ListMasterHosts_FullMethodName  = "/yandex.cloud.mdb.greenplum.v1.ClusterService/ListMasterHosts"
	ClusterService_ListSegmentHosts_FullMethodName = "/yandex.cloud.mdb.greenplum.v1.ClusterService/ListSegmentHosts"
	ClusterService_ListLogs_FullMethodName         = "/yandex.cloud.mdb.greenplum.v1.ClusterService/ListLogs"
	ClusterService_StreamLogs_FullMethodName       = "/yandex.cloud.mdb.greenplum.v1.ClusterService/StreamLogs"
	ClusterService_ListBackups_FullMethodName      = "/yandex.cloud.mdb.greenplum.v1.ClusterService/ListBackups"
	ClusterService_Backup_FullMethodName           = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Backup"
	ClusterService_Restore_FullMethodName          = "/yandex.cloud.mdb.greenplum.v1.ClusterService/Restore"
)

// ClusterServiceClient is the client API for ClusterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClusterServiceClient interface {
	// Returns the specified Greenplum® cluster.
	//
	// To get the list of all available Greenplum® clusters, make a [List] request.
	Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error)
	// Retrieves a list of Greenplum® clusters that belong to the specified folder.
	List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Creates a Greenplum® cluster in the specified folder.
	Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified Greenplum® cluster.
	Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Expands the specified Greenplum® cluster.
	Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified Greenplum® cluster.
	Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the specified Greenplum® cluster.
	Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops the specified Greenplum® cluster.
	Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of Operation resources for the specified cluster.
	ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error)
	// Retrieves a list of master hosts for the specified cluster.
	ListMasterHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error)
	// Retrieves a list of segment hosts for the specified cluster.
	ListSegmentHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error)
	// Retrieves logs for the specified Greenplum® cluster.
	ListLogs(ctx context.Context, in *ListClusterLogsRequest, opts ...grpc.CallOption) (*ListClusterLogsResponse, error)
	// Same as [ListLogs] but using server-side streaming. Also allows for `tail -f` semantics.
	StreamLogs(ctx context.Context, in *StreamClusterLogsRequest, opts ...grpc.CallOption) (ClusterService_StreamLogsClient, error)
	// Retrieves a list of available backups for the specified Greenplum® cluster.
	ListBackups(ctx context.Context, in *ListClusterBackupsRequest, opts ...grpc.CallOption) (*ListClusterBackupsResponse, error)
	// Creates a backup for the specified Greenplum cluster.
	Backup(ctx context.Context, in *BackupClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Creates a new Greenplum® cluster using the specified backup.
	Restore(ctx context.Context, in *RestoreClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type clusterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClusterServiceClient(cc grpc.ClientConnInterface) ClusterServiceClient {
	return &clusterServiceClient{cc}
}

func (c *clusterServiceClient) Get(ctx context.Context, in *GetClusterRequest, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, ClusterService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) List(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, ClusterService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Create(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Update(ctx context.Context, in *UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Expand(ctx context.Context, in *ExpandRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Expand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Delete(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Start(ctx context.Context, in *StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Stop(ctx context.Context, in *StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListOperations(ctx context.Context, in *ListClusterOperationsRequest, opts ...grpc.CallOption) (*ListClusterOperationsResponse, error) {
	out := new(ListClusterOperationsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListOperations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListMasterHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error) {
	out := new(ListClusterHostsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListMasterHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListSegmentHosts(ctx context.Context, in *ListClusterHostsRequest, opts ...grpc.CallOption) (*ListClusterHostsResponse, error) {
	out := new(ListClusterHostsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListSegmentHosts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) ListLogs(ctx context.Context, in *ListClusterLogsRequest, opts ...grpc.CallOption) (*ListClusterLogsResponse, error) {
	out := new(ListClusterLogsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListLogs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) StreamLogs(ctx context.Context, in *StreamClusterLogsRequest, opts ...grpc.CallOption) (ClusterService_StreamLogsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClusterService_ServiceDesc.Streams[0], ClusterService_StreamLogs_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &clusterServiceStreamLogsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ClusterService_StreamLogsClient interface {
	Recv() (*StreamLogRecord, error)
	grpc.ClientStream
}

type clusterServiceStreamLogsClient struct {
	grpc.ClientStream
}

func (x *clusterServiceStreamLogsClient) Recv() (*StreamLogRecord, error) {
	m := new(StreamLogRecord)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clusterServiceClient) ListBackups(ctx context.Context, in *ListClusterBackupsRequest, opts ...grpc.CallOption) (*ListClusterBackupsResponse, error) {
	out := new(ListClusterBackupsResponse)
	err := c.cc.Invoke(ctx, ClusterService_ListBackups_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Backup(ctx context.Context, in *BackupClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Backup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterServiceClient) Restore(ctx context.Context, in *RestoreClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ClusterService_Restore_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClusterServiceServer is the server API for ClusterService service.
// All implementations should embed UnimplementedClusterServiceServer
// for forward compatibility
type ClusterServiceServer interface {
	// Returns the specified Greenplum® cluster.
	//
	// To get the list of all available Greenplum® clusters, make a [List] request.
	Get(context.Context, *GetClusterRequest) (*Cluster, error)
	// Retrieves a list of Greenplum® clusters that belong to the specified folder.
	List(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Creates a Greenplum® cluster in the specified folder.
	Create(context.Context, *CreateClusterRequest) (*operation.Operation, error)
	// Updates the specified Greenplum® cluster.
	Update(context.Context, *UpdateClusterRequest) (*operation.Operation, error)
	// Expands the specified Greenplum® cluster.
	Expand(context.Context, *ExpandRequest) (*operation.Operation, error)
	// Deletes the specified Greenplum® cluster.
	Delete(context.Context, *DeleteClusterRequest) (*operation.Operation, error)
	// Starts the specified Greenplum® cluster.
	Start(context.Context, *StartClusterRequest) (*operation.Operation, error)
	// Stops the specified Greenplum® cluster.
	Stop(context.Context, *StopClusterRequest) (*operation.Operation, error)
	// Retrieves the list of Operation resources for the specified cluster.
	ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error)
	// Retrieves a list of master hosts for the specified cluster.
	ListMasterHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error)
	// Retrieves a list of segment hosts for the specified cluster.
	ListSegmentHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error)
	// Retrieves logs for the specified Greenplum® cluster.
	ListLogs(context.Context, *ListClusterLogsRequest) (*ListClusterLogsResponse, error)
	// Same as [ListLogs] but using server-side streaming. Also allows for `tail -f` semantics.
	StreamLogs(*StreamClusterLogsRequest, ClusterService_StreamLogsServer) error
	// Retrieves a list of available backups for the specified Greenplum® cluster.
	ListBackups(context.Context, *ListClusterBackupsRequest) (*ListClusterBackupsResponse, error)
	// Creates a backup for the specified Greenplum cluster.
	Backup(context.Context, *BackupClusterRequest) (*operation.Operation, error)
	// Creates a new Greenplum® cluster using the specified backup.
	Restore(context.Context, *RestoreClusterRequest) (*operation.Operation, error)
}

// UnimplementedClusterServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClusterServiceServer struct {
}

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
func (UnimplementedClusterServiceServer) Expand(context.Context, *ExpandRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expand not implemented")
}
func (UnimplementedClusterServiceServer) Delete(context.Context, *DeleteClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedClusterServiceServer) Start(context.Context, *StartClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedClusterServiceServer) Stop(context.Context, *StopClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedClusterServiceServer) ListOperations(context.Context, *ListClusterOperationsRequest) (*ListClusterOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedClusterServiceServer) ListMasterHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMasterHosts not implemented")
}
func (UnimplementedClusterServiceServer) ListSegmentHosts(context.Context, *ListClusterHostsRequest) (*ListClusterHostsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSegmentHosts not implemented")
}
func (UnimplementedClusterServiceServer) ListLogs(context.Context, *ListClusterLogsRequest) (*ListClusterLogsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogs not implemented")
}
func (UnimplementedClusterServiceServer) StreamLogs(*StreamClusterLogsRequest, ClusterService_StreamLogsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLogs not implemented")
}
func (UnimplementedClusterServiceServer) ListBackups(context.Context, *ListClusterBackupsRequest) (*ListClusterBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBackups not implemented")
}
func (UnimplementedClusterServiceServer) Backup(context.Context, *BackupClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Backup not implemented")
}
func (UnimplementedClusterServiceServer) Restore(context.Context, *RestoreClusterRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restore not implemented")
}

// UnsafeClusterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClusterServiceServer will
// result in compilation errors.
type UnsafeClusterServiceServer interface {
	mustEmbedUnimplementedClusterServiceServer()
}

func RegisterClusterServiceServer(s grpc.ServiceRegistrar, srv ClusterServiceServer) {
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

func _ClusterService_Expand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExpandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Expand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Expand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Expand(ctx, req.(*ExpandRequest))
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

func _ClusterService_ListMasterHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListMasterHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListMasterHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListMasterHosts(ctx, req.(*ListClusterHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListSegmentHosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterHostsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListSegmentHosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListSegmentHosts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListSegmentHosts(ctx, req.(*ListClusterHostsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_ListLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterLogsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListLogs(ctx, req.(*ListClusterLogsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_StreamLogs_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamClusterLogsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ClusterServiceServer).StreamLogs(m, &clusterServiceStreamLogsServer{stream})
}

type ClusterService_StreamLogsServer interface {
	Send(*StreamLogRecord) error
	grpc.ServerStream
}

type clusterServiceStreamLogsServer struct {
	grpc.ServerStream
}

func (x *clusterServiceStreamLogsServer) Send(m *StreamLogRecord) error {
	return x.ServerStream.SendMsg(m)
}

func _ClusterService_ListBackups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClusterBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).ListBackups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_ListBackups_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).ListBackups(ctx, req.(*ListClusterBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Backup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BackupClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Backup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Backup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Backup(ctx, req.(*BackupClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClusterService_Restore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestoreClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServiceServer).Restore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClusterService_Restore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServiceServer).Restore(ctx, req.(*RestoreClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClusterService_ServiceDesc is the grpc.ServiceDesc for ClusterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClusterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.greenplum.v1.ClusterService",
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
			MethodName: "Expand",
			Handler:    _ClusterService_Expand_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ClusterService_Delete_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ClusterService_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ClusterService_Stop_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ClusterService_ListOperations_Handler,
		},
		{
			MethodName: "ListMasterHosts",
			Handler:    _ClusterService_ListMasterHosts_Handler,
		},
		{
			MethodName: "ListSegmentHosts",
			Handler:    _ClusterService_ListSegmentHosts_Handler,
		},
		{
			MethodName: "ListLogs",
			Handler:    _ClusterService_ListLogs_Handler,
		},
		{
			MethodName: "ListBackups",
			Handler:    _ClusterService_ListBackups_Handler,
		},
		{
			MethodName: "Backup",
			Handler:    _ClusterService_Backup_Handler,
		},
		{
			MethodName: "Restore",
			Handler:    _ClusterService_Restore_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLogs",
			Handler:       _ClusterService_StreamLogs_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "yandex/cloud/mdb/greenplum/v1/cluster_service.proto",
}
