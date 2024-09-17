// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/datasphere/v1/project_service.proto

package datasphere

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ProjectService_Create_FullMethodName              = "/yandex.cloud.datasphere.v1.ProjectService/Create"
	ProjectService_Update_FullMethodName              = "/yandex.cloud.datasphere.v1.ProjectService/Update"
	ProjectService_Delete_FullMethodName              = "/yandex.cloud.datasphere.v1.ProjectService/Delete"
	ProjectService_Open_FullMethodName                = "/yandex.cloud.datasphere.v1.ProjectService/Open"
	ProjectService_Get_FullMethodName                 = "/yandex.cloud.datasphere.v1.ProjectService/Get"
	ProjectService_List_FullMethodName                = "/yandex.cloud.datasphere.v1.ProjectService/List"
	ProjectService_GetUnitBalance_FullMethodName      = "/yandex.cloud.datasphere.v1.ProjectService/GetUnitBalance"
	ProjectService_SetUnitBalance_FullMethodName      = "/yandex.cloud.datasphere.v1.ProjectService/SetUnitBalance"
	ProjectService_Execute_FullMethodName             = "/yandex.cloud.datasphere.v1.ProjectService/Execute"
	ProjectService_GetCellOutputs_FullMethodName      = "/yandex.cloud.datasphere.v1.ProjectService/GetCellOutputs"
	ProjectService_GetStateVariables_FullMethodName   = "/yandex.cloud.datasphere.v1.ProjectService/GetStateVariables"
	ProjectService_GetNotebookMetadata_FullMethodName = "/yandex.cloud.datasphere.v1.ProjectService/GetNotebookMetadata"
)

// ProjectServiceClient is the client API for ProjectService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Project resources.
type ProjectServiceClient interface {
	// Creates a project in the specified folder.
	Create(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified project.
	Update(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified project.
	Delete(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Opens the specified project.
	Open(ctx context.Context, in *OpenProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns the specified project.
	Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error)
	// Lists projects for the specified folder.
	List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	// Returns the unit balance of the specified project.
	GetUnitBalance(ctx context.Context, in *GetUnitBalanceRequest, opts ...grpc.CallOption) (*GetUnitBalanceResponse, error)
	// Sets the unit balance of the specified project.
	SetUnitBalance(ctx context.Context, in *SetUnitBalanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Executes code in the specified cell or notebook.
	Execute(ctx context.Context, in *ProjectExecutionRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Returns outputs of the specified cell.
	GetCellOutputs(ctx context.Context, in *CellOutputsRequest, opts ...grpc.CallOption) (*CellOutputsResponse, error)
	// Returns state variables of the specified notebook.
	GetStateVariables(ctx context.Context, in *GetStateVariablesRequest, opts ...grpc.CallOption) (*GetStateVariablesResponse, error)
	// Returns metadata of the specified notebook.
	GetNotebookMetadata(ctx context.Context, in *GetNotebookMetadataRequest, opts ...grpc.CallOption) (*GetNotebookMetadataResponse, error)
}

type projectServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectServiceClient(cc grpc.ClientConnInterface) ProjectServiceClient {
	return &projectServiceClient{cc}
}

func (c *projectServiceClient) Create(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Update(ctx context.Context, in *UpdateProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Delete(ctx context.Context, in *DeleteProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Open(ctx context.Context, in *OpenProjectRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Open_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Get(ctx context.Context, in *GetProjectRequest, opts ...grpc.CallOption) (*Project, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Project)
	err := c.cc.Invoke(ctx, ProjectService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) List(ctx context.Context, in *ListProjectsRequest, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, ProjectService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetUnitBalance(ctx context.Context, in *GetUnitBalanceRequest, opts ...grpc.CallOption) (*GetUnitBalanceResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUnitBalanceResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetUnitBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) SetUnitBalance(ctx context.Context, in *SetUnitBalanceRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, ProjectService_SetUnitBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) Execute(ctx context.Context, in *ProjectExecutionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ProjectService_Execute_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetCellOutputs(ctx context.Context, in *CellOutputsRequest, opts ...grpc.CallOption) (*CellOutputsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CellOutputsResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetCellOutputs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetStateVariables(ctx context.Context, in *GetStateVariablesRequest, opts ...grpc.CallOption) (*GetStateVariablesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetStateVariablesResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetStateVariables_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectServiceClient) GetNotebookMetadata(ctx context.Context, in *GetNotebookMetadataRequest, opts ...grpc.CallOption) (*GetNotebookMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetNotebookMetadataResponse)
	err := c.cc.Invoke(ctx, ProjectService_GetNotebookMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectServiceServer is the server API for ProjectService service.
// All implementations should embed UnimplementedProjectServiceServer
// for forward compatibility.
//
// A set of methods for managing Project resources.
type ProjectServiceServer interface {
	// Creates a project in the specified folder.
	Create(context.Context, *CreateProjectRequest) (*operation.Operation, error)
	// Updates the specified project.
	Update(context.Context, *UpdateProjectRequest) (*operation.Operation, error)
	// Deletes the specified project.
	Delete(context.Context, *DeleteProjectRequest) (*operation.Operation, error)
	// Opens the specified project.
	Open(context.Context, *OpenProjectRequest) (*operation.Operation, error)
	// Returns the specified project.
	Get(context.Context, *GetProjectRequest) (*Project, error)
	// Lists projects for the specified folder.
	List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error)
	// Returns the unit balance of the specified project.
	GetUnitBalance(context.Context, *GetUnitBalanceRequest) (*GetUnitBalanceResponse, error)
	// Sets the unit balance of the specified project.
	SetUnitBalance(context.Context, *SetUnitBalanceRequest) (*emptypb.Empty, error)
	// Executes code in the specified cell or notebook.
	Execute(context.Context, *ProjectExecutionRequest) (*operation.Operation, error)
	// Returns outputs of the specified cell.
	GetCellOutputs(context.Context, *CellOutputsRequest) (*CellOutputsResponse, error)
	// Returns state variables of the specified notebook.
	GetStateVariables(context.Context, *GetStateVariablesRequest) (*GetStateVariablesResponse, error)
	// Returns metadata of the specified notebook.
	GetNotebookMetadata(context.Context, *GetNotebookMetadataRequest) (*GetNotebookMetadataResponse, error)
}

// UnimplementedProjectServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedProjectServiceServer struct{}

func (UnimplementedProjectServiceServer) Create(context.Context, *CreateProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedProjectServiceServer) Update(context.Context, *UpdateProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedProjectServiceServer) Delete(context.Context, *DeleteProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedProjectServiceServer) Open(context.Context, *OpenProjectRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Open not implemented")
}
func (UnimplementedProjectServiceServer) Get(context.Context, *GetProjectRequest) (*Project, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedProjectServiceServer) List(context.Context, *ListProjectsRequest) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedProjectServiceServer) GetUnitBalance(context.Context, *GetUnitBalanceRequest) (*GetUnitBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnitBalance not implemented")
}
func (UnimplementedProjectServiceServer) SetUnitBalance(context.Context, *SetUnitBalanceRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUnitBalance not implemented")
}
func (UnimplementedProjectServiceServer) Execute(context.Context, *ProjectExecutionRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedProjectServiceServer) GetCellOutputs(context.Context, *CellOutputsRequest) (*CellOutputsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCellOutputs not implemented")
}
func (UnimplementedProjectServiceServer) GetStateVariables(context.Context, *GetStateVariablesRequest) (*GetStateVariablesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStateVariables not implemented")
}
func (UnimplementedProjectServiceServer) GetNotebookMetadata(context.Context, *GetNotebookMetadataRequest) (*GetNotebookMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNotebookMetadata not implemented")
}
func (UnimplementedProjectServiceServer) testEmbeddedByValue() {}

// UnsafeProjectServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectServiceServer will
// result in compilation errors.
type UnsafeProjectServiceServer interface {
	mustEmbedUnimplementedProjectServiceServer()
}

func RegisterProjectServiceServer(s grpc.ServiceRegistrar, srv ProjectServiceServer) {
	// If the following call pancis, it indicates UnimplementedProjectServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ProjectService_ServiceDesc, srv)
}

func _ProjectService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Create(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Update(ctx, req.(*UpdateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Delete(ctx, req.(*DeleteProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Open_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OpenProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Open(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Open_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Open(ctx, req.(*OpenProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Get(ctx, req.(*GetProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).List(ctx, req.(*ListProjectsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetUnitBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUnitBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetUnitBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetUnitBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetUnitBalance(ctx, req.(*GetUnitBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_SetUnitBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUnitBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).SetUnitBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_SetUnitBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).SetUnitBalance(ctx, req.(*SetUnitBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_Execute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).Execute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_Execute_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).Execute(ctx, req.(*ProjectExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetCellOutputs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CellOutputsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetCellOutputs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetCellOutputs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetCellOutputs(ctx, req.(*CellOutputsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetStateVariables_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStateVariablesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetStateVariables(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetStateVariables_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetStateVariables(ctx, req.(*GetStateVariablesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProjectService_GetNotebookMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNotebookMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectServiceServer).GetNotebookMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProjectService_GetNotebookMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectServiceServer).GetNotebookMetadata(ctx, req.(*GetNotebookMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProjectService_ServiceDesc is the grpc.ServiceDesc for ProjectService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProjectService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.datasphere.v1.ProjectService",
	HandlerType: (*ProjectServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ProjectService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ProjectService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ProjectService_Delete_Handler,
		},
		{
			MethodName: "Open",
			Handler:    _ProjectService_Open_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _ProjectService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ProjectService_List_Handler,
		},
		{
			MethodName: "GetUnitBalance",
			Handler:    _ProjectService_GetUnitBalance_Handler,
		},
		{
			MethodName: "SetUnitBalance",
			Handler:    _ProjectService_SetUnitBalance_Handler,
		},
		{
			MethodName: "Execute",
			Handler:    _ProjectService_Execute_Handler,
		},
		{
			MethodName: "GetCellOutputs",
			Handler:    _ProjectService_GetCellOutputs_Handler,
		},
		{
			MethodName: "GetStateVariables",
			Handler:    _ProjectService_GetStateVariables_Handler,
		},
		{
			MethodName: "GetNotebookMetadata",
			Handler:    _ProjectService_GetNotebookMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/datasphere/v1/project_service.proto",
}
