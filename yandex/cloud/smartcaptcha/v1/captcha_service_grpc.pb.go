// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/smartcaptcha/v1/captcha_service.proto

package smartcaptcha

import (
	context "context"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CaptchaService_Get_FullMethodName          = "/yandex.cloud.smartcaptcha.v1.CaptchaService/Get"
	CaptchaService_GetSecretKey_FullMethodName = "/yandex.cloud.smartcaptcha.v1.CaptchaService/GetSecretKey"
	CaptchaService_List_FullMethodName         = "/yandex.cloud.smartcaptcha.v1.CaptchaService/List"
	CaptchaService_Create_FullMethodName       = "/yandex.cloud.smartcaptcha.v1.CaptchaService/Create"
	CaptchaService_Update_FullMethodName       = "/yandex.cloud.smartcaptcha.v1.CaptchaService/Update"
	CaptchaService_Delete_FullMethodName       = "/yandex.cloud.smartcaptcha.v1.CaptchaService/Delete"
)

// CaptchaServiceClient is the client API for CaptchaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Captcha resources.
type CaptchaServiceClient interface {
	// Returns the specified Captcha resource.
	Get(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*Captcha, error)
	// Returns the secret data of specified Captcha resource.
	GetSecretKey(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*CaptchaSecretKey, error)
	// Retrieves the list of Captcha resources in the specified folder.
	List(ctx context.Context, in *ListCaptchasRequest, opts ...grpc.CallOption) (*ListCaptchasResponse, error)
	// Creates a captcha in the specified folder using the data specified in the request.
	Create(ctx context.Context, in *CreateCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified captcha.
	Update(ctx context.Context, in *UpdateCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified captcha.
	Delete(ctx context.Context, in *DeleteCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type captchaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCaptchaServiceClient(cc grpc.ClientConnInterface) CaptchaServiceClient {
	return &captchaServiceClient{cc}
}

func (c *captchaServiceClient) Get(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*Captcha, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Captcha)
	err := c.cc.Invoke(ctx, CaptchaService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) GetSecretKey(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*CaptchaSecretKey, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CaptchaSecretKey)
	err := c.cc.Invoke(ctx, CaptchaService_GetSecretKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) List(ctx context.Context, in *ListCaptchasRequest, opts ...grpc.CallOption) (*ListCaptchasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCaptchasResponse)
	err := c.cc.Invoke(ctx, CaptchaService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) Create(ctx context.Context, in *CreateCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CaptchaService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) Update(ctx context.Context, in *UpdateCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CaptchaService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *captchaServiceClient) Delete(ctx context.Context, in *DeleteCaptchaRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CaptchaService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CaptchaServiceServer is the server API for CaptchaService service.
// All implementations should embed UnimplementedCaptchaServiceServer
// for forward compatibility.
//
// A set of methods for managing Captcha resources.
type CaptchaServiceServer interface {
	// Returns the specified Captcha resource.
	Get(context.Context, *GetCaptchaRequest) (*Captcha, error)
	// Returns the secret data of specified Captcha resource.
	GetSecretKey(context.Context, *GetCaptchaRequest) (*CaptchaSecretKey, error)
	// Retrieves the list of Captcha resources in the specified folder.
	List(context.Context, *ListCaptchasRequest) (*ListCaptchasResponse, error)
	// Creates a captcha in the specified folder using the data specified in the request.
	Create(context.Context, *CreateCaptchaRequest) (*operation.Operation, error)
	// Updates the specified captcha.
	Update(context.Context, *UpdateCaptchaRequest) (*operation.Operation, error)
	// Deletes the specified captcha.
	Delete(context.Context, *DeleteCaptchaRequest) (*operation.Operation, error)
}

// UnimplementedCaptchaServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCaptchaServiceServer struct{}

func (UnimplementedCaptchaServiceServer) Get(context.Context, *GetCaptchaRequest) (*Captcha, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCaptchaServiceServer) GetSecretKey(context.Context, *GetCaptchaRequest) (*CaptchaSecretKey, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSecretKey not implemented")
}
func (UnimplementedCaptchaServiceServer) List(context.Context, *ListCaptchasRequest) (*ListCaptchasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCaptchaServiceServer) Create(context.Context, *CreateCaptchaRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCaptchaServiceServer) Update(context.Context, *UpdateCaptchaRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCaptchaServiceServer) Delete(context.Context, *DeleteCaptchaRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCaptchaServiceServer) testEmbeddedByValue() {}

// UnsafeCaptchaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CaptchaServiceServer will
// result in compilation errors.
type UnsafeCaptchaServiceServer interface {
	mustEmbedUnimplementedCaptchaServiceServer()
}

func RegisterCaptchaServiceServer(s grpc.ServiceRegistrar, srv CaptchaServiceServer) {
	// If the following call pancis, it indicates UnimplementedCaptchaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CaptchaService_ServiceDesc, srv)
}

func _CaptchaService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).Get(ctx, req.(*GetCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_GetSecretKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).GetSecretKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_GetSecretKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).GetSecretKey(ctx, req.(*GetCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCaptchasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).List(ctx, req.(*ListCaptchasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).Create(ctx, req.(*CreateCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).Update(ctx, req.(*UpdateCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CaptchaService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CaptchaServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CaptchaService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CaptchaServiceServer).Delete(ctx, req.(*DeleteCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CaptchaService_ServiceDesc is the grpc.ServiceDesc for CaptchaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CaptchaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.smartcaptcha.v1.CaptchaService",
	HandlerType: (*CaptchaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CaptchaService_Get_Handler,
		},
		{
			MethodName: "GetSecretKey",
			Handler:    _CaptchaService_GetSecretKey_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CaptchaService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CaptchaService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CaptchaService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CaptchaService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/smartcaptcha/v1/captcha_service.proto",
}
