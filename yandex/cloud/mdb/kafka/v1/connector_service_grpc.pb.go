// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package kafka

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

// ConnectorServiceClient is the client API for ConnectorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ConnectorServiceClient interface {
	// Returns the specified Apache Kafka Connector resource.
	//
	// To get the list of available Apache Kafka Connector resources, make a [List] request.
	Get(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*Connector, error)
	// Retrieves the list of Apache Kafka Connector resources in the specified cluster.
	List(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error)
	// Creates a new Apache Kafka connector in the specified cluster.
	Create(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified Apache Kafka connector.
	Delete(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Resume the specified Apache Kafka connector.
	Resume(ctx context.Context, in *ResumeConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Pause the specified Apache Kafka connector.
	Pause(ctx context.Context, in *PauseConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type connectorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewConnectorServiceClient(cc grpc.ClientConnInterface) ConnectorServiceClient {
	return &connectorServiceClient{cc}
}

func (c *connectorServiceClient) Get(ctx context.Context, in *GetConnectorRequest, opts ...grpc.CallOption) (*Connector, error) {
	out := new(Connector)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ConnectorService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorServiceClient) List(ctx context.Context, in *ListConnectorsRequest, opts ...grpc.CallOption) (*ListConnectorsResponse, error) {
	out := new(ListConnectorsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ConnectorService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorServiceClient) Create(ctx context.Context, in *CreateConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ConnectorService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorServiceClient) Delete(ctx context.Context, in *DeleteConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ConnectorService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorServiceClient) Resume(ctx context.Context, in *ResumeConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ConnectorService/Resume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *connectorServiceClient) Pause(ctx context.Context, in *PauseConnectorRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.kafka.v1.ConnectorService/Pause", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConnectorServiceServer is the server API for ConnectorService service.
// All implementations should embed UnimplementedConnectorServiceServer
// for forward compatibility
type ConnectorServiceServer interface {
	// Returns the specified Apache Kafka Connector resource.
	//
	// To get the list of available Apache Kafka Connector resources, make a [List] request.
	Get(context.Context, *GetConnectorRequest) (*Connector, error)
	// Retrieves the list of Apache Kafka Connector resources in the specified cluster.
	List(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error)
	// Creates a new Apache Kafka connector in the specified cluster.
	Create(context.Context, *CreateConnectorRequest) (*operation.Operation, error)
	// Deletes the specified Apache Kafka connector.
	Delete(context.Context, *DeleteConnectorRequest) (*operation.Operation, error)
	// Resume the specified Apache Kafka connector.
	Resume(context.Context, *ResumeConnectorRequest) (*operation.Operation, error)
	// Pause the specified Apache Kafka connector.
	Pause(context.Context, *PauseConnectorRequest) (*operation.Operation, error)
}

// UnimplementedConnectorServiceServer should be embedded to have forward compatible implementations.
type UnimplementedConnectorServiceServer struct {
}

func (UnimplementedConnectorServiceServer) Get(context.Context, *GetConnectorRequest) (*Connector, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedConnectorServiceServer) List(context.Context, *ListConnectorsRequest) (*ListConnectorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedConnectorServiceServer) Create(context.Context, *CreateConnectorRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedConnectorServiceServer) Delete(context.Context, *DeleteConnectorRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedConnectorServiceServer) Resume(context.Context, *ResumeConnectorRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resume not implemented")
}
func (UnimplementedConnectorServiceServer) Pause(context.Context, *PauseConnectorRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}

// UnsafeConnectorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ConnectorServiceServer will
// result in compilation errors.
type UnsafeConnectorServiceServer interface {
	mustEmbedUnimplementedConnectorServiceServer()
}

func RegisterConnectorServiceServer(s grpc.ServiceRegistrar, srv ConnectorServiceServer) {
	s.RegisterService(&ConnectorService_ServiceDesc, srv)
}

func _ConnectorService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ConnectorService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).Get(ctx, req.(*GetConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListConnectorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ConnectorService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).List(ctx, req.(*ListConnectorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ConnectorService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).Create(ctx, req.(*CreateConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ConnectorService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).Delete(ctx, req.(*DeleteConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorService_Resume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResumeConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).Resume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ConnectorService/Resume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).Resume(ctx, req.(*ResumeConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConnectorService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PauseConnectorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConnectorServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.kafka.v1.ConnectorService/Pause",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConnectorServiceServer).Pause(ctx, req.(*PauseConnectorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ConnectorService_ServiceDesc is the grpc.ServiceDesc for ConnectorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ConnectorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.kafka.v1.ConnectorService",
	HandlerType: (*ConnectorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ConnectorService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ConnectorService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ConnectorService_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ConnectorService_Delete_Handler,
		},
		{
			MethodName: "Resume",
			Handler:    _ConnectorService_Resume_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _ConnectorService_Pause_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/kafka/v1/connector_service.proto",
}
