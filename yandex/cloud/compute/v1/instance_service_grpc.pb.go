// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package compute

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

// InstanceServiceClient is the client API for InstanceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceServiceClient interface {
	// Returns the specified Instance resource.
	//
	// To get the list of available Instance resources, make a [List] request.
	Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error)
	// Retrieves the list of Instance resources in the specified folder.
	List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error)
	// Creates an instance in the specified folder.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified instance.
	Update(ctx context.Context, in *UpdateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified instance.
	Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the metadata of the specified instance.
	UpdateMetadata(ctx context.Context, in *UpdateInstanceMetadataRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the serial port output of the specified Instance resource.
	GetSerialPortOutput(ctx context.Context, in *GetInstanceSerialPortOutputRequest, opts ...grpc.CallOption) (*GetInstanceSerialPortOutputResponse, error)
	// Stops the running instance.
	//
	// You can start the instance later using the [InstanceService.Start] method.
	Stop(ctx context.Context, in *StopInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the stopped instance.
	Start(ctx context.Context, in *StartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Restarts the running instance.
	Restart(ctx context.Context, in *RestartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Attaches the disk to the instance.
	AttachDisk(ctx context.Context, in *AttachInstanceDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Detaches the disk from the instance.
	DetachDisk(ctx context.Context, in *DetachInstanceDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Attaches the filesystem to the instance.
	//
	// The instance and the filesystem must reside in the same availability zone.
	//
	// To attach a filesystem, the instance must have a `STOPPED` status ([Instance.status]).
	// To check the instance status, make a [InstanceService.Get] request.
	// To stop the running instance, make a [InstanceService.Stop] request.
	//
	// To use the instance with an attached filesystem, the latter must be mounted.
	// For details, see [documentation](/docs/compute/operations/filesystem/attach-to-vm).
	AttachFilesystem(ctx context.Context, in *AttachInstanceFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Detaches the filesystem from the instance.
	//
	// To detach a filesystem, the instance must have a `STOPPED` status ([Instance.status]).
	// To check the instance status, make a [InstanceService.Get] request.
	// To stop the running instance, make a [InstanceService.Stop] request.
	DetachFilesystem(ctx context.Context, in *DetachInstanceFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Enables One-to-one NAT on the network interface.
	AddOneToOneNat(ctx context.Context, in *AddInstanceOneToOneNatRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Removes One-to-one NAT from the network interface.
	RemoveOneToOneNat(ctx context.Context, in *RemoveInstanceOneToOneNatRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified instance network interface.
	UpdateNetworkInterface(ctx context.Context, in *UpdateInstanceNetworkInterfaceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified instance.
	ListOperations(ctx context.Context, in *ListInstanceOperationsRequest, opts ...grpc.CallOption) (*ListInstanceOperationsResponse, error)
}

type instanceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceServiceClient(cc grpc.ClientConnInterface) InstanceServiceClient {
	return &instanceServiceClient{cc}
}

func (c *instanceServiceClient) Get(ctx context.Context, in *GetInstanceRequest, opts ...grpc.CallOption) (*Instance, error) {
	out := new(Instance)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) List(ctx context.Context, in *ListInstancesRequest, opts ...grpc.CallOption) (*ListInstancesResponse, error) {
	out := new(ListInstancesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) Create(ctx context.Context, in *CreateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) Update(ctx context.Context, in *UpdateInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) Delete(ctx context.Context, in *DeleteInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) UpdateMetadata(ctx context.Context, in *UpdateInstanceMetadataRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/UpdateMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) GetSerialPortOutput(ctx context.Context, in *GetInstanceSerialPortOutputRequest, opts ...grpc.CallOption) (*GetInstanceSerialPortOutputResponse, error) {
	out := new(GetInstanceSerialPortOutputResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/GetSerialPortOutput", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) Stop(ctx context.Context, in *StopInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) Start(ctx context.Context, in *StartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) Restart(ctx context.Context, in *RestartInstanceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/Restart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) AttachDisk(ctx context.Context, in *AttachInstanceDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/AttachDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) DetachDisk(ctx context.Context, in *DetachInstanceDiskRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/DetachDisk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) AttachFilesystem(ctx context.Context, in *AttachInstanceFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/AttachFilesystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) DetachFilesystem(ctx context.Context, in *DetachInstanceFilesystemRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/DetachFilesystem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) AddOneToOneNat(ctx context.Context, in *AddInstanceOneToOneNatRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/AddOneToOneNat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) RemoveOneToOneNat(ctx context.Context, in *RemoveInstanceOneToOneNatRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/RemoveOneToOneNat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) UpdateNetworkInterface(ctx context.Context, in *UpdateInstanceNetworkInterfaceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/UpdateNetworkInterface", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceServiceClient) ListOperations(ctx context.Context, in *ListInstanceOperationsRequest, opts ...grpc.CallOption) (*ListInstanceOperationsResponse, error) {
	out := new(ListInstanceOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.InstanceService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceServiceServer is the server API for InstanceService service.
// All implementations should embed UnimplementedInstanceServiceServer
// for forward compatibility
type InstanceServiceServer interface {
	// Returns the specified Instance resource.
	//
	// To get the list of available Instance resources, make a [List] request.
	Get(context.Context, *GetInstanceRequest) (*Instance, error)
	// Retrieves the list of Instance resources in the specified folder.
	List(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error)
	// Creates an instance in the specified folder.
	// Method starts an asynchronous operation that can be cancelled while it is in progress.
	Create(context.Context, *CreateInstanceRequest) (*operation.Operation, error)
	// Updates the specified instance.
	Update(context.Context, *UpdateInstanceRequest) (*operation.Operation, error)
	// Deletes the specified instance.
	Delete(context.Context, *DeleteInstanceRequest) (*operation.Operation, error)
	// Updates the metadata of the specified instance.
	UpdateMetadata(context.Context, *UpdateInstanceMetadataRequest) (*operation.Operation, error)
	// Returns the serial port output of the specified Instance resource.
	GetSerialPortOutput(context.Context, *GetInstanceSerialPortOutputRequest) (*GetInstanceSerialPortOutputResponse, error)
	// Stops the running instance.
	//
	// You can start the instance later using the [InstanceService.Start] method.
	Stop(context.Context, *StopInstanceRequest) (*operation.Operation, error)
	// Starts the stopped instance.
	Start(context.Context, *StartInstanceRequest) (*operation.Operation, error)
	// Restarts the running instance.
	Restart(context.Context, *RestartInstanceRequest) (*operation.Operation, error)
	// Attaches the disk to the instance.
	AttachDisk(context.Context, *AttachInstanceDiskRequest) (*operation.Operation, error)
	// Detaches the disk from the instance.
	DetachDisk(context.Context, *DetachInstanceDiskRequest) (*operation.Operation, error)
	// Attaches the filesystem to the instance.
	//
	// The instance and the filesystem must reside in the same availability zone.
	//
	// To attach a filesystem, the instance must have a `STOPPED` status ([Instance.status]).
	// To check the instance status, make a [InstanceService.Get] request.
	// To stop the running instance, make a [InstanceService.Stop] request.
	//
	// To use the instance with an attached filesystem, the latter must be mounted.
	// For details, see [documentation](/docs/compute/operations/filesystem/attach-to-vm).
	AttachFilesystem(context.Context, *AttachInstanceFilesystemRequest) (*operation.Operation, error)
	// Detaches the filesystem from the instance.
	//
	// To detach a filesystem, the instance must have a `STOPPED` status ([Instance.status]).
	// To check the instance status, make a [InstanceService.Get] request.
	// To stop the running instance, make a [InstanceService.Stop] request.
	DetachFilesystem(context.Context, *DetachInstanceFilesystemRequest) (*operation.Operation, error)
	// Enables One-to-one NAT on the network interface.
	AddOneToOneNat(context.Context, *AddInstanceOneToOneNatRequest) (*operation.Operation, error)
	// Removes One-to-one NAT from the network interface.
	RemoveOneToOneNat(context.Context, *RemoveInstanceOneToOneNatRequest) (*operation.Operation, error)
	// Updates the specified instance network interface.
	UpdateNetworkInterface(context.Context, *UpdateInstanceNetworkInterfaceRequest) (*operation.Operation, error)
	// Lists operations for the specified instance.
	ListOperations(context.Context, *ListInstanceOperationsRequest) (*ListInstanceOperationsResponse, error)
}

// UnimplementedInstanceServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInstanceServiceServer struct {
}

func (UnimplementedInstanceServiceServer) Get(context.Context, *GetInstanceRequest) (*Instance, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInstanceServiceServer) List(context.Context, *ListInstancesRequest) (*ListInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedInstanceServiceServer) Create(context.Context, *CreateInstanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInstanceServiceServer) Update(context.Context, *UpdateInstanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedInstanceServiceServer) Delete(context.Context, *DeleteInstanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedInstanceServiceServer) UpdateMetadata(context.Context, *UpdateInstanceMetadataRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (UnimplementedInstanceServiceServer) GetSerialPortOutput(context.Context, *GetInstanceSerialPortOutputRequest) (*GetInstanceSerialPortOutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSerialPortOutput not implemented")
}
func (UnimplementedInstanceServiceServer) Stop(context.Context, *StopInstanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedInstanceServiceServer) Start(context.Context, *StartInstanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedInstanceServiceServer) Restart(context.Context, *RestartInstanceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Restart not implemented")
}
func (UnimplementedInstanceServiceServer) AttachDisk(context.Context, *AttachInstanceDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachDisk not implemented")
}
func (UnimplementedInstanceServiceServer) DetachDisk(context.Context, *DetachInstanceDiskRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachDisk not implemented")
}
func (UnimplementedInstanceServiceServer) AttachFilesystem(context.Context, *AttachInstanceFilesystemRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachFilesystem not implemented")
}
func (UnimplementedInstanceServiceServer) DetachFilesystem(context.Context, *DetachInstanceFilesystemRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetachFilesystem not implemented")
}
func (UnimplementedInstanceServiceServer) AddOneToOneNat(context.Context, *AddInstanceOneToOneNatRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOneToOneNat not implemented")
}
func (UnimplementedInstanceServiceServer) RemoveOneToOneNat(context.Context, *RemoveInstanceOneToOneNatRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveOneToOneNat not implemented")
}
func (UnimplementedInstanceServiceServer) UpdateNetworkInterface(context.Context, *UpdateInstanceNetworkInterfaceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNetworkInterface not implemented")
}
func (UnimplementedInstanceServiceServer) ListOperations(context.Context, *ListInstanceOperationsRequest) (*ListInstanceOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}

// UnsafeInstanceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceServiceServer will
// result in compilation errors.
type UnsafeInstanceServiceServer interface {
	mustEmbedUnimplementedInstanceServiceServer()
}

func RegisterInstanceServiceServer(s grpc.ServiceRegistrar, srv InstanceServiceServer) {
	s.RegisterService(&InstanceService_ServiceDesc, srv)
}

func _InstanceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Get(ctx, req.(*GetInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).List(ctx, req.(*ListInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Create(ctx, req.(*CreateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Update(ctx, req.(*UpdateInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Delete(ctx, req.(*DeleteInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/UpdateMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).UpdateMetadata(ctx, req.(*UpdateInstanceMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_GetSerialPortOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceSerialPortOutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).GetSerialPortOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/GetSerialPortOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).GetSerialPortOutput(ctx, req.(*GetInstanceSerialPortOutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Stop(ctx, req.(*StopInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Start(ctx, req.(*StartInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_Restart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartInstanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).Restart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/Restart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).Restart(ctx, req.(*RestartInstanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_AttachDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachInstanceDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).AttachDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/AttachDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).AttachDisk(ctx, req.(*AttachInstanceDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_DetachDisk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachInstanceDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).DetachDisk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/DetachDisk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).DetachDisk(ctx, req.(*DetachInstanceDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_AttachFilesystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachInstanceFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).AttachFilesystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/AttachFilesystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).AttachFilesystem(ctx, req.(*AttachInstanceFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_DetachFilesystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetachInstanceFilesystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).DetachFilesystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/DetachFilesystem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).DetachFilesystem(ctx, req.(*DetachInstanceFilesystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_AddOneToOneNat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddInstanceOneToOneNatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).AddOneToOneNat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/AddOneToOneNat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).AddOneToOneNat(ctx, req.(*AddInstanceOneToOneNatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_RemoveOneToOneNat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveInstanceOneToOneNatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).RemoveOneToOneNat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/RemoveOneToOneNat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).RemoveOneToOneNat(ctx, req.(*RemoveInstanceOneToOneNatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_UpdateNetworkInterface_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceNetworkInterfaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).UpdateNetworkInterface(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/UpdateNetworkInterface",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).UpdateNetworkInterface(ctx, req.(*UpdateInstanceNetworkInterfaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.InstanceService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceServiceServer).ListOperations(ctx, req.(*ListInstanceOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstanceService_ServiceDesc is the grpc.ServiceDesc for InstanceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstanceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.InstanceService",
	HandlerType: (*InstanceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _InstanceService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _InstanceService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _InstanceService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _InstanceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _InstanceService_Delete_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _InstanceService_UpdateMetadata_Handler,
		},
		{
			MethodName: "GetSerialPortOutput",
			Handler:    _InstanceService_GetSerialPortOutput_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _InstanceService_Stop_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _InstanceService_Start_Handler,
		},
		{
			MethodName: "Restart",
			Handler:    _InstanceService_Restart_Handler,
		},
		{
			MethodName: "AttachDisk",
			Handler:    _InstanceService_AttachDisk_Handler,
		},
		{
			MethodName: "DetachDisk",
			Handler:    _InstanceService_DetachDisk_Handler,
		},
		{
			MethodName: "AttachFilesystem",
			Handler:    _InstanceService_AttachFilesystem_Handler,
		},
		{
			MethodName: "DetachFilesystem",
			Handler:    _InstanceService_DetachFilesystem_Handler,
		},
		{
			MethodName: "AddOneToOneNat",
			Handler:    _InstanceService_AddOneToOneNat_Handler,
		},
		{
			MethodName: "RemoveOneToOneNat",
			Handler:    _InstanceService_RemoveOneToOneNat_Handler,
		},
		{
			MethodName: "UpdateNetworkInterface",
			Handler:    _InstanceService_UpdateNetworkInterface_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _InstanceService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/instance_service.proto",
}
