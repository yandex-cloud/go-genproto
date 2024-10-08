// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/datasphere/v1/app_token_service.proto

package datasphere

import (
	context "context"
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
	AppTokenService_Validate_FullMethodName = "/yandex.cloud.datasphere.v1.AppTokenService/Validate"
)

// AppTokenServiceClient is the client API for AppTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing app tokens.
type AppTokenServiceClient interface {
	// Validates app token.
	Validate(ctx context.Context, in *AppTokenValidateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type appTokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppTokenServiceClient(cc grpc.ClientConnInterface) AppTokenServiceClient {
	return &appTokenServiceClient{cc}
}

func (c *appTokenServiceClient) Validate(ctx context.Context, in *AppTokenValidateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, AppTokenService_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppTokenServiceServer is the server API for AppTokenService service.
// All implementations should embed UnimplementedAppTokenServiceServer
// for forward compatibility.
//
// A set of methods for managing app tokens.
type AppTokenServiceServer interface {
	// Validates app token.
	Validate(context.Context, *AppTokenValidateRequest) (*emptypb.Empty, error)
}

// UnimplementedAppTokenServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAppTokenServiceServer struct{}

func (UnimplementedAppTokenServiceServer) Validate(context.Context, *AppTokenValidateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedAppTokenServiceServer) testEmbeddedByValue() {}

// UnsafeAppTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppTokenServiceServer will
// result in compilation errors.
type UnsafeAppTokenServiceServer interface {
	mustEmbedUnimplementedAppTokenServiceServer()
}

func RegisterAppTokenServiceServer(s grpc.ServiceRegistrar, srv AppTokenServiceServer) {
	// If the following call pancis, it indicates UnimplementedAppTokenServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AppTokenService_ServiceDesc, srv)
}

func _AppTokenService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppTokenValidateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppTokenServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppTokenService_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppTokenServiceServer).Validate(ctx, req.(*AppTokenValidateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppTokenService_ServiceDesc is the grpc.ServiceDesc for AppTokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppTokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.datasphere.v1.AppTokenService",
	HandlerType: (*AppTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Validate",
			Handler:    _AppTokenService_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/datasphere/v1/app_token_service.proto",
}
