// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dataproc_manager

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataprocManagerServiceClient is the client API for DataprocManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataprocManagerServiceClient interface {
	// Sends a status report from a host
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type dataprocManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataprocManagerServiceClient(cc grpc.ClientConnInterface) DataprocManagerServiceClient {
	return &dataprocManagerServiceClient{cc}
}

func (c *dataprocManagerServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.manager.v1.DataprocManagerService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataprocManagerServiceServer is the server API for DataprocManagerService service.
// All implementations should embed UnimplementedDataprocManagerServiceServer
// for forward compatibility
type DataprocManagerServiceServer interface {
	// Sends a status report from a host
	Report(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedDataprocManagerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedDataprocManagerServiceServer struct {
}

func (UnimplementedDataprocManagerServiceServer) Report(context.Context, *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

// UnsafeDataprocManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataprocManagerServiceServer will
// result in compilation errors.
type UnsafeDataprocManagerServiceServer interface {
	mustEmbedUnimplementedDataprocManagerServiceServer()
}

func RegisterDataprocManagerServiceServer(s grpc.ServiceRegistrar, srv DataprocManagerServiceServer) {
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
		FullMethod: "/yandex.cloud.dataproc.manager.v1.DataprocManagerService/Report",
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
