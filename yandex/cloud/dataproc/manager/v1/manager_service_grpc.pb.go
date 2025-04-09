// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/dataproc/manager/v1/manager_service.proto

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
	DataprocManagerService_Report_FullMethodName = "/yandex.cloud.dataproc.manager.v1.DataprocManagerService/Report"
)

// DataprocManagerServiceClient is the client API for DataprocManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Yandex Data Processing manager service definition.
type DataprocManagerServiceClient interface {
	// Sends a status report from a host.
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type dataprocManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataprocManagerServiceClient(cc grpc.ClientConnInterface) DataprocManagerServiceClient {
	return &dataprocManagerServiceClient{cc}
}

func (c *dataprocManagerServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, DataprocManagerService_Report_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataprocManagerServiceServer is the server API for DataprocManagerService service.
// All implementations should embed UnimplementedDataprocManagerServiceServer
// for forward compatibility.
//
// Yandex Data Processing manager service definition.
type DataprocManagerServiceServer interface {
	// Sends a status report from a host.
	Report(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedDataprocManagerServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDataprocManagerServiceServer struct{}

func (UnimplementedDataprocManagerServiceServer) Report(context.Context, *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}
func (UnimplementedDataprocManagerServiceServer) testEmbeddedByValue() {}

// UnsafeDataprocManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataprocManagerServiceServer will
// result in compilation errors.
type UnsafeDataprocManagerServiceServer interface {
	mustEmbedUnimplementedDataprocManagerServiceServer()
}

func RegisterDataprocManagerServiceServer(s grpc.ServiceRegistrar, srv DataprocManagerServiceServer) {
	// If the following call pancis, it indicates UnimplementedDataprocManagerServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DataprocManagerService_ServiceDesc, srv)
}

func _DataprocManagerService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataprocManagerServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DataprocManagerService_Report_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataprocManagerServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DataprocManagerService_ServiceDesc is the grpc.ServiceDesc for DataprocManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataprocManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dataproc.manager.v1.DataprocManagerService",
	HandlerType: (*DataprocManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _DataprocManagerService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dataproc/manager/v1/manager_service.proto",
}
