// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package instancegroup

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// InstanceGroupServiceClient is the client API for InstanceGroupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type InstanceGroupServiceClient interface {
	// Returns the specified InstanceGroup resource.
	//
	// To get the list of available InstanceGroup resources, make a [List] request.
	Get(ctx context.Context, in *GetInstanceGroupRequest, opts ...grpc.CallOption) (*InstanceGroup, error)
	// Retrieves the list of InstanceGroup resources in the specified folder.
	List(ctx context.Context, in *ListInstanceGroupsRequest, opts ...grpc.CallOption) (*ListInstanceGroupsResponse, error)
	// Creates an instance group in the specified folder.
	// This method starts an operation that can be cancelled by another operation.
	Create(ctx context.Context, in *CreateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Creates an instance group in the specified folder from a YAML file.
	// This method starts an operation that can be cancelled by another operation.
	CreateFromYaml(ctx context.Context, in *CreateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified instance group.
	// This method starts an operation that can be cancelled by another operation.
	Update(ctx context.Context, in *UpdateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified instance group from a YAML file.
	// This method starts an operation that can be cancelled by another operation.
	UpdateFromYaml(ctx context.Context, in *UpdateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stops the specified instance group.
	Stop(ctx context.Context, in *StopInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Starts the specified instance group.
	Start(ctx context.Context, in *StartInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified instance group.
	Delete(ctx context.Context, in *DeleteInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists instances for the specified instance group.
	ListInstances(ctx context.Context, in *ListInstanceGroupInstancesRequest, opts ...grpc.CallOption) (*ListInstanceGroupInstancesResponse, error)
	// Delete instances from the instance group.
	DeleteInstances(ctx context.Context, in *DeleteInstancesRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Stop instances from the instance group.
	StopInstances(ctx context.Context, in *StopInstancesRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified instance group.
	ListOperations(ctx context.Context, in *ListInstanceGroupOperationsRequest, opts ...grpc.CallOption) (*ListInstanceGroupOperationsResponse, error)
	// Lists logs for the specified instance group.
	ListLogRecords(ctx context.Context, in *ListInstanceGroupLogRecordsRequest, opts ...grpc.CallOption) (*ListInstanceGroupLogRecordsResponse, error)
	// Lists existing access bindings for the specified instance group.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified instance group.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified instance group.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Resumes all processes regarding management of the specified instance group,
	// i.e. scaling, checking instances' health, auto-healing and updating them.
	ResumeProcesses(ctx context.Context, in *ResumeInstanceGroupProcessesRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Pauses all processes regarding management of the specified instance group,
	// i.e. scaling, checking instances' health, auto-healing and updating them. Running instances are not stopped.
	PauseProcesses(ctx context.Context, in *PauseInstanceGroupProcessesRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type instanceGroupServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewInstanceGroupServiceClient(cc grpc.ClientConnInterface) InstanceGroupServiceClient {
	return &instanceGroupServiceClient{cc}
}

func (c *instanceGroupServiceClient) Get(ctx context.Context, in *GetInstanceGroupRequest, opts ...grpc.CallOption) (*InstanceGroup, error) {
	out := new(InstanceGroup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) List(ctx context.Context, in *ListInstanceGroupsRequest, opts ...grpc.CallOption) (*ListInstanceGroupsResponse, error) {
	out := new(ListInstanceGroupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) Create(ctx context.Context, in *CreateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) CreateFromYaml(ctx context.Context, in *CreateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/CreateFromYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) Update(ctx context.Context, in *UpdateInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) UpdateFromYaml(ctx context.Context, in *UpdateInstanceGroupFromYamlRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/UpdateFromYaml", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) Stop(ctx context.Context, in *StopInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) Start(ctx context.Context, in *StartInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) Delete(ctx context.Context, in *DeleteInstanceGroupRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) ListInstances(ctx context.Context, in *ListInstanceGroupInstancesRequest, opts ...grpc.CallOption) (*ListInstanceGroupInstancesResponse, error) {
	out := new(ListInstanceGroupInstancesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) DeleteInstances(ctx context.Context, in *DeleteInstancesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/DeleteInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) StopInstances(ctx context.Context, in *StopInstancesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/StopInstances", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) ListOperations(ctx context.Context, in *ListInstanceGroupOperationsRequest, opts ...grpc.CallOption) (*ListInstanceGroupOperationsResponse, error) {
	out := new(ListInstanceGroupOperationsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) ListLogRecords(ctx context.Context, in *ListInstanceGroupLogRecordsRequest, opts ...grpc.CallOption) (*ListInstanceGroupLogRecordsResponse, error) {
	out := new(ListInstanceGroupLogRecordsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListLogRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/SetAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/UpdateAccessBindings", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) ResumeProcesses(ctx context.Context, in *ResumeInstanceGroupProcessesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ResumeProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *instanceGroupServiceClient) PauseProcesses(ctx context.Context, in *PauseInstanceGroupProcessesRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/PauseProcesses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InstanceGroupServiceServer is the server API for InstanceGroupService service.
// All implementations should embed UnimplementedInstanceGroupServiceServer
// for forward compatibility
type InstanceGroupServiceServer interface {
	// Returns the specified InstanceGroup resource.
	//
	// To get the list of available InstanceGroup resources, make a [List] request.
	Get(context.Context, *GetInstanceGroupRequest) (*InstanceGroup, error)
	// Retrieves the list of InstanceGroup resources in the specified folder.
	List(context.Context, *ListInstanceGroupsRequest) (*ListInstanceGroupsResponse, error)
	// Creates an instance group in the specified folder.
	// This method starts an operation that can be cancelled by another operation.
	Create(context.Context, *CreateInstanceGroupRequest) (*operation.Operation, error)
	// Creates an instance group in the specified folder from a YAML file.
	// This method starts an operation that can be cancelled by another operation.
	CreateFromYaml(context.Context, *CreateInstanceGroupFromYamlRequest) (*operation.Operation, error)
	// Updates the specified instance group.
	// This method starts an operation that can be cancelled by another operation.
	Update(context.Context, *UpdateInstanceGroupRequest) (*operation.Operation, error)
	// Updates the specified instance group from a YAML file.
	// This method starts an operation that can be cancelled by another operation.
	UpdateFromYaml(context.Context, *UpdateInstanceGroupFromYamlRequest) (*operation.Operation, error)
	// Stops the specified instance group.
	Stop(context.Context, *StopInstanceGroupRequest) (*operation.Operation, error)
	// Starts the specified instance group.
	Start(context.Context, *StartInstanceGroupRequest) (*operation.Operation, error)
	// Deletes the specified instance group.
	Delete(context.Context, *DeleteInstanceGroupRequest) (*operation.Operation, error)
	// Lists instances for the specified instance group.
	ListInstances(context.Context, *ListInstanceGroupInstancesRequest) (*ListInstanceGroupInstancesResponse, error)
	// Delete instances from the instance group.
	DeleteInstances(context.Context, *DeleteInstancesRequest) (*operation.Operation, error)
	// Stop instances from the instance group.
	StopInstances(context.Context, *StopInstancesRequest) (*operation.Operation, error)
	// Lists operations for the specified instance group.
	ListOperations(context.Context, *ListInstanceGroupOperationsRequest) (*ListInstanceGroupOperationsResponse, error)
	// Lists logs for the specified instance group.
	ListLogRecords(context.Context, *ListInstanceGroupLogRecordsRequest) (*ListInstanceGroupLogRecordsResponse, error)
	// Lists existing access bindings for the specified instance group.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the specified instance group.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified instance group.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
	// Resumes all processes regarding management of the specified instance group,
	// i.e. scaling, checking instances' health, auto-healing and updating them.
	ResumeProcesses(context.Context, *ResumeInstanceGroupProcessesRequest) (*operation.Operation, error)
	// Pauses all processes regarding management of the specified instance group,
	// i.e. scaling, checking instances' health, auto-healing and updating them. Running instances are not stopped.
	PauseProcesses(context.Context, *PauseInstanceGroupProcessesRequest) (*operation.Operation, error)
}

// UnimplementedInstanceGroupServiceServer should be embedded to have forward compatible implementations.
type UnimplementedInstanceGroupServiceServer struct {
}

func (UnimplementedInstanceGroupServiceServer) Get(context.Context, *GetInstanceGroupRequest) (*InstanceGroup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedInstanceGroupServiceServer) List(context.Context, *ListInstanceGroupsRequest) (*ListInstanceGroupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedInstanceGroupServiceServer) Create(context.Context, *CreateInstanceGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedInstanceGroupServiceServer) CreateFromYaml(context.Context, *CreateInstanceGroupFromYamlRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFromYaml not implemented")
}
func (UnimplementedInstanceGroupServiceServer) Update(context.Context, *UpdateInstanceGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedInstanceGroupServiceServer) UpdateFromYaml(context.Context, *UpdateInstanceGroupFromYamlRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFromYaml not implemented")
}
func (UnimplementedInstanceGroupServiceServer) Stop(context.Context, *StopInstanceGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedInstanceGroupServiceServer) Start(context.Context, *StartInstanceGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedInstanceGroupServiceServer) Delete(context.Context, *DeleteInstanceGroupRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedInstanceGroupServiceServer) ListInstances(context.Context, *ListInstanceGroupInstancesRequest) (*ListInstanceGroupInstancesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListInstances not implemented")
}
func (UnimplementedInstanceGroupServiceServer) DeleteInstances(context.Context, *DeleteInstancesRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInstances not implemented")
}
func (UnimplementedInstanceGroupServiceServer) StopInstances(context.Context, *StopInstancesRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopInstances not implemented")
}
func (UnimplementedInstanceGroupServiceServer) ListOperations(context.Context, *ListInstanceGroupOperationsRequest) (*ListInstanceGroupOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedInstanceGroupServiceServer) ListLogRecords(context.Context, *ListInstanceGroupLogRecordsRequest) (*ListInstanceGroupLogRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLogRecords not implemented")
}
func (UnimplementedInstanceGroupServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedInstanceGroupServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedInstanceGroupServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedInstanceGroupServiceServer) ResumeProcesses(context.Context, *ResumeInstanceGroupProcessesRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResumeProcesses not implemented")
}
func (UnimplementedInstanceGroupServiceServer) PauseProcesses(context.Context, *PauseInstanceGroupProcessesRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PauseProcesses not implemented")
}

// UnsafeInstanceGroupServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to InstanceGroupServiceServer will
// result in compilation errors.
type UnsafeInstanceGroupServiceServer interface {
	mustEmbedUnimplementedInstanceGroupServiceServer()
}

func RegisterInstanceGroupServiceServer(s grpc.ServiceRegistrar, srv InstanceGroupServiceServer) {
	s.RegisterService(&InstanceGroupService_ServiceDesc, srv)
}

func _InstanceGroupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetInstanceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).Get(ctx, req.(*GetInstanceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceGroupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).List(ctx, req.(*ListInstanceGroupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).Create(ctx, req.(*CreateInstanceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_CreateFromYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateInstanceGroupFromYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).CreateFromYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/CreateFromYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).CreateFromYaml(ctx, req.(*CreateInstanceGroupFromYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).Update(ctx, req.(*UpdateInstanceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_UpdateFromYaml_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInstanceGroupFromYamlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).UpdateFromYaml(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/UpdateFromYaml",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).UpdateFromYaml(ctx, req.(*UpdateInstanceGroupFromYamlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopInstanceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).Stop(ctx, req.(*StopInstanceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartInstanceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).Start(ctx, req.(*StartInstanceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstanceGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).Delete(ctx, req.(*DeleteInstanceGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_ListInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceGroupInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).ListInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).ListInstances(ctx, req.(*ListInstanceGroupInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_DeleteInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).DeleteInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/DeleteInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).DeleteInstances(ctx, req.(*DeleteInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_StopInstances_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopInstancesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).StopInstances(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/StopInstances",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).StopInstances(ctx, req.(*StopInstancesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceGroupOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).ListOperations(ctx, req.(*ListInstanceGroupOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_ListLogRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListInstanceGroupLogRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).ListLogRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListLogRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).ListLogRecords(ctx, req.(*ListInstanceGroupLogRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ListAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/SetAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/UpdateAccessBindings",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_ResumeProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeInstanceGroupProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).ResumeProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/ResumeProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).ResumeProcesses(ctx, req.(*ResumeInstanceGroupProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InstanceGroupService_PauseProcesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseInstanceGroupProcessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InstanceGroupServiceServer).PauseProcesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.compute.v1.instancegroup.InstanceGroupService/PauseProcesses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InstanceGroupServiceServer).PauseProcesses(ctx, req.(*PauseInstanceGroupProcessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// InstanceGroupService_ServiceDesc is the grpc.ServiceDesc for InstanceGroupService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var InstanceGroupService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.compute.v1.instancegroup.InstanceGroupService",
	HandlerType: (*InstanceGroupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _InstanceGroupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _InstanceGroupService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _InstanceGroupService_Create_Handler,
		},
		{
			MethodName: "CreateFromYaml",
			Handler:    _InstanceGroupService_CreateFromYaml_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _InstanceGroupService_Update_Handler,
		},
		{
			MethodName: "UpdateFromYaml",
			Handler:    _InstanceGroupService_UpdateFromYaml_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _InstanceGroupService_Stop_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _InstanceGroupService_Start_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _InstanceGroupService_Delete_Handler,
		},
		{
			MethodName: "ListInstances",
			Handler:    _InstanceGroupService_ListInstances_Handler,
		},
		{
			MethodName: "DeleteInstances",
			Handler:    _InstanceGroupService_DeleteInstances_Handler,
		},
		{
			MethodName: "StopInstances",
			Handler:    _InstanceGroupService_StopInstances_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _InstanceGroupService_ListOperations_Handler,
		},
		{
			MethodName: "ListLogRecords",
			Handler:    _InstanceGroupService_ListLogRecords_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _InstanceGroupService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _InstanceGroupService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _InstanceGroupService_UpdateAccessBindings_Handler,
		},
		{
			MethodName: "ResumeProcesses",
			Handler:    _InstanceGroupService_ResumeProcesses_Handler,
		},
		{
			MethodName: "PauseProcesses",
			Handler:    _InstanceGroupService_PauseProcesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/compute/v1/instancegroup/instance_group_service.proto",
}
