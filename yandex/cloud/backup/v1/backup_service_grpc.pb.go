// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/backup/v1/backup_service.proto

package backup

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
	BackupService_List_FullMethodName               = "/yandex.cloud.backup.v1.BackupService/List"
	BackupService_ListArchives_FullMethodName       = "/yandex.cloud.backup.v1.BackupService/ListArchives"
	BackupService_ListFiles_FullMethodName          = "/yandex.cloud.backup.v1.BackupService/ListFiles"
	BackupService_Get_FullMethodName                = "/yandex.cloud.backup.v1.BackupService/Get"
	BackupService_StartRecovery_FullMethodName      = "/yandex.cloud.backup.v1.BackupService/StartRecovery"
	BackupService_StartFilesRecovery_FullMethodName = "/yandex.cloud.backup.v1.BackupService/StartFilesRecovery"
	BackupService_Delete_FullMethodName             = "/yandex.cloud.backup.v1.BackupService/Delete"
)

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BackupServiceClient interface {
	// List backups using filters.
	List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error)
	// List archives that holds backups for specified folder or
	// specified [Compute Cloud instance](/docs/backup/concepts/vm-connection#os).
	ListArchives(ctx context.Context, in *ListArchivesRequest, opts ...grpc.CallOption) (*ListArchivesResponse, error)
	// ListFiles of the backup.
	ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error)
	// Get backup by its id.
	Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error)
	// Start recovery process of specified backup to specific Compute Cloud instance.
	//
	// For details, see [Restoring a VM from a backup](/docs/backup/operations/backup-vm/recover).
	StartRecovery(ctx context.Context, in *StartRecoveryRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// StartFilesRecovery runs recovery process of selected files to specific Compute Cloud instance.
	StartFilesRecovery(ctx context.Context, in *StartFilesRecoveryRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Delete specific backup.
	Delete(ctx context.Context, in *DeleteBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type backupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBackupServiceClient(cc grpc.ClientConnInterface) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error) {
	out := new(ListBackupsResponse)
	err := c.cc.Invoke(ctx, BackupService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListArchives(ctx context.Context, in *ListArchivesRequest, opts ...grpc.CallOption) (*ListArchivesResponse, error) {
	out := new(ListArchivesResponse)
	err := c.cc.Invoke(ctx, BackupService_ListArchives_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) ListFiles(ctx context.Context, in *ListFilesRequest, opts ...grpc.CallOption) (*ListFilesResponse, error) {
	out := new(ListFilesResponse)
	err := c.cc.Invoke(ctx, BackupService_ListFiles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error) {
	out := new(Backup)
	err := c.cc.Invoke(ctx, BackupService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) StartRecovery(ctx context.Context, in *StartRecoveryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BackupService_StartRecovery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) StartFilesRecovery(ctx context.Context, in *StartFilesRecoveryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BackupService_StartFilesRecovery_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) Delete(ctx context.Context, in *DeleteBackupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, BackupService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
// All implementations should embed UnimplementedBackupServiceServer
// for forward compatibility
type BackupServiceServer interface {
	// List backups using filters.
	List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error)
	// List archives that holds backups for specified folder or
	// specified [Compute Cloud instance](/docs/backup/concepts/vm-connection#os).
	ListArchives(context.Context, *ListArchivesRequest) (*ListArchivesResponse, error)
	// ListFiles of the backup.
	ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error)
	// Get backup by its id.
	Get(context.Context, *GetBackupRequest) (*Backup, error)
	// Start recovery process of specified backup to specific Compute Cloud instance.
	//
	// For details, see [Restoring a VM from a backup](/docs/backup/operations/backup-vm/recover).
	StartRecovery(context.Context, *StartRecoveryRequest) (*operation.Operation, error)
	// StartFilesRecovery runs recovery process of selected files to specific Compute Cloud instance.
	StartFilesRecovery(context.Context, *StartFilesRecoveryRequest) (*operation.Operation, error)
	// Delete specific backup.
	Delete(context.Context, *DeleteBackupRequest) (*operation.Operation, error)
}

// UnimplementedBackupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (UnimplementedBackupServiceServer) List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedBackupServiceServer) ListArchives(context.Context, *ListArchivesRequest) (*ListArchivesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListArchives not implemented")
}
func (UnimplementedBackupServiceServer) ListFiles(context.Context, *ListFilesRequest) (*ListFilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFiles not implemented")
}
func (UnimplementedBackupServiceServer) Get(context.Context, *GetBackupRequest) (*Backup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBackupServiceServer) StartRecovery(context.Context, *StartRecoveryRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartRecovery not implemented")
}
func (UnimplementedBackupServiceServer) StartFilesRecovery(context.Context, *StartFilesRecoveryRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartFilesRecovery not implemented")
}
func (UnimplementedBackupServiceServer) Delete(context.Context, *DeleteBackupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeBackupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BackupServiceServer will
// result in compilation errors.
type UnsafeBackupServiceServer interface {
	mustEmbedUnimplementedBackupServiceServer()
}

func RegisterBackupServiceServer(s grpc.ServiceRegistrar, srv BackupServiceServer) {
	s.RegisterService(&BackupService_ServiceDesc, srv)
}

func _BackupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).List(ctx, req.(*ListBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListArchives_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArchivesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListArchives(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ListArchives_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListArchives(ctx, req.(*ListArchivesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_ListFiles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).ListFiles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_ListFiles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).ListFiles(ctx, req.(*ListFilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Get(ctx, req.(*GetBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_StartRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).StartRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_StartRecovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).StartRecovery(ctx, req.(*StartRecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_StartFilesRecovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartFilesRecoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).StartFilesRecovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_StartFilesRecovery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).StartFilesRecovery(ctx, req.(*StartFilesRecoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BackupService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Delete(ctx, req.(*DeleteBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BackupService_ServiceDesc is the grpc.ServiceDesc for BackupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BackupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.backup.v1.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _BackupService_List_Handler,
		},
		{
			MethodName: "ListArchives",
			Handler:    _BackupService_ListArchives_Handler,
		},
		{
			MethodName: "ListFiles",
			Handler:    _BackupService_ListFiles_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _BackupService_Get_Handler,
		},
		{
			MethodName: "StartRecovery",
			Handler:    _BackupService_StartRecovery_Handler,
		},
		{
			MethodName: "StartFilesRecovery",
			Handler:    _BackupService_StartFilesRecovery_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BackupService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/backup/v1/backup_service.proto",
}