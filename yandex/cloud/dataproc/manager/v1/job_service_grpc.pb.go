// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/dataproc/manager/v1/job_service.proto

package dataproc_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	JobService_ListActive_FullMethodName          = "/yandex.cloud.dataproc.manager.v1.JobService/ListActive"
	JobService_UpdateStatus_FullMethodName        = "/yandex.cloud.dataproc.manager.v1.JobService/UpdateStatus"
	JobService_ListSupportActive_FullMethodName   = "/yandex.cloud.dataproc.manager.v1.JobService/ListSupportActive"
	JobService_UpdateSupportStatus_FullMethodName = "/yandex.cloud.dataproc.manager.v1.JobService/UpdateSupportStatus"
	JobService_SaveSupportLog_FullMethodName      = "/yandex.cloud.dataproc.manager.v1.JobService/SaveSupportLog"
)

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobServiceClient interface {
	// Retrieves a list of jobs for Yandex Data Processing cluster.
	ListActive(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	// Currently used to update Job status.
	UpdateStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error)
	// Retrieves a list of support jobs for Yandex Data Processing cluster.
	ListSupportActive(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListSupportJobsResponse, error)
	// Currently used to update support job status.
	UpdateSupportStatus(ctx context.Context, in *UpdateSupportJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error)
	// Save support job output.
	SaveSupportLog(ctx context.Context, in *SaveSupportJobLogRequest, opts ...grpc.CallOption) (*SaveSupportJobLogResponse, error)
}

type jobServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobServiceClient(cc grpc.ClientConnInterface) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) ListActive(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, JobService_ListActive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) UpdateStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateJobStatusResponse)
	err := c.cc.Invoke(ctx, JobService_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) ListSupportActive(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListSupportJobsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSupportJobsResponse)
	err := c.cc.Invoke(ctx, JobService_ListSupportActive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) UpdateSupportStatus(ctx context.Context, in *UpdateSupportJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateJobStatusResponse)
	err := c.cc.Invoke(ctx, JobService_UpdateSupportStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) SaveSupportLog(ctx context.Context, in *SaveSupportJobLogRequest, opts ...grpc.CallOption) (*SaveSupportJobLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SaveSupportJobLogResponse)
	err := c.cc.Invoke(ctx, JobService_SaveSupportLog_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
// All implementations should embed UnimplementedJobServiceServer
// for forward compatibility.
type JobServiceServer interface {
	// Retrieves a list of jobs for Yandex Data Processing cluster.
	ListActive(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	// Currently used to update Job status.
	UpdateStatus(context.Context, *UpdateJobStatusRequest) (*UpdateJobStatusResponse, error)
	// Retrieves a list of support jobs for Yandex Data Processing cluster.
	ListSupportActive(context.Context, *ListJobsRequest) (*ListSupportJobsResponse, error)
	// Currently used to update support job status.
	UpdateSupportStatus(context.Context, *UpdateSupportJobStatusRequest) (*UpdateJobStatusResponse, error)
	// Save support job output.
	SaveSupportLog(context.Context, *SaveSupportJobLogRequest) (*SaveSupportJobLogResponse, error)
}

// UnimplementedJobServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedJobServiceServer struct{}

func (UnimplementedJobServiceServer) ListActive(context.Context, *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActive not implemented")
}
func (UnimplementedJobServiceServer) UpdateStatus(context.Context, *UpdateJobStatusRequest) (*UpdateJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedJobServiceServer) ListSupportActive(context.Context, *ListJobsRequest) (*ListSupportJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSupportActive not implemented")
}
func (UnimplementedJobServiceServer) UpdateSupportStatus(context.Context, *UpdateSupportJobStatusRequest) (*UpdateJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSupportStatus not implemented")
}
func (UnimplementedJobServiceServer) SaveSupportLog(context.Context, *SaveSupportJobLogRequest) (*SaveSupportJobLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSupportLog not implemented")
}
func (UnimplementedJobServiceServer) testEmbeddedByValue() {}

// UnsafeJobServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobServiceServer will
// result in compilation errors.
type UnsafeJobServiceServer interface {
	mustEmbedUnimplementedJobServiceServer()
}

func RegisterJobServiceServer(s grpc.ServiceRegistrar, srv JobServiceServer) {
	// If the following call pancis, it indicates UnimplementedJobServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&JobService_ServiceDesc, srv)
}

func _JobService_ListActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ListActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_ListActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ListActive(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).UpdateStatus(ctx, req.(*UpdateJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_ListSupportActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ListSupportActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_ListSupportActive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ListSupportActive(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_UpdateSupportStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSupportJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).UpdateSupportStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_UpdateSupportStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).UpdateSupportStatus(ctx, req.(*UpdateSupportJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_SaveSupportLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSupportJobLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).SaveSupportLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobService_SaveSupportLog_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).SaveSupportLog(ctx, req.(*SaveSupportJobLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobService_ServiceDesc is the grpc.ServiceDesc for JobService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dataproc.manager.v1.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListActive",
			Handler:    _JobService_ListActive_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _JobService_UpdateStatus_Handler,
		},
		{
			MethodName: "ListSupportActive",
			Handler:    _JobService_ListSupportActive_Handler,
		},
		{
			MethodName: "UpdateSupportStatus",
			Handler:    _JobService_UpdateSupportStatus_Handler,
		},
		{
			MethodName: "SaveSupportLog",
			Handler:    _JobService_SaveSupportLog_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dataproc/manager/v1/job_service.proto",
}
