// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/iot/devices/v1/device_service.proto

package devices

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
	DeviceService_Get_FullMethodName               = "/yandex.cloud.iot.devices.v1.DeviceService/Get"
	DeviceService_GetByName_FullMethodName         = "/yandex.cloud.iot.devices.v1.DeviceService/GetByName"
	DeviceService_List_FullMethodName              = "/yandex.cloud.iot.devices.v1.DeviceService/List"
	DeviceService_Create_FullMethodName            = "/yandex.cloud.iot.devices.v1.DeviceService/Create"
	DeviceService_Update_FullMethodName            = "/yandex.cloud.iot.devices.v1.DeviceService/Update"
	DeviceService_Delete_FullMethodName            = "/yandex.cloud.iot.devices.v1.DeviceService/Delete"
	DeviceService_ListCertificates_FullMethodName  = "/yandex.cloud.iot.devices.v1.DeviceService/ListCertificates"
	DeviceService_AddCertificate_FullMethodName    = "/yandex.cloud.iot.devices.v1.DeviceService/AddCertificate"
	DeviceService_DeleteCertificate_FullMethodName = "/yandex.cloud.iot.devices.v1.DeviceService/DeleteCertificate"
	DeviceService_ListPasswords_FullMethodName     = "/yandex.cloud.iot.devices.v1.DeviceService/ListPasswords"
	DeviceService_AddPassword_FullMethodName       = "/yandex.cloud.iot.devices.v1.DeviceService/AddPassword"
	DeviceService_DeletePassword_FullMethodName    = "/yandex.cloud.iot.devices.v1.DeviceService/DeletePassword"
	DeviceService_ListOperations_FullMethodName    = "/yandex.cloud.iot.devices.v1.DeviceService/ListOperations"
)

// DeviceServiceClient is the client API for DeviceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing devices.
type DeviceServiceClient interface {
	// Returns the specified device.
	//
	// To get the list of available devices, make a [List] request.
	Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*Device, error)
	GetByName(ctx context.Context, in *GetByNameDeviceRequest, opts ...grpc.CallOption) (*Device, error)
	// Retrieves the list of devices in the specified registry.
	List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error)
	// Creates a device in the specified registry.
	Create(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified device.
	Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified device.
	Delete(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of device certificates for the specified device.
	ListCertificates(ctx context.Context, in *ListDeviceCertificatesRequest, opts ...grpc.CallOption) (*ListDeviceCertificatesResponse, error)
	// Adds a certificate.
	AddCertificate(ctx context.Context, in *AddDeviceCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified device certificate.
	DeleteCertificate(ctx context.Context, in *DeleteDeviceCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Retrieves the list of passwords for the specified device.
	ListPasswords(ctx context.Context, in *ListDevicePasswordsRequest, opts ...grpc.CallOption) (*ListDevicePasswordsResponse, error)
	// Adds password for the specified device.
	AddPassword(ctx context.Context, in *AddDevicePasswordRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified password.
	DeletePassword(ctx context.Context, in *DeleteDevicePasswordRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified device.
	ListOperations(ctx context.Context, in *ListDeviceOperationsRequest, opts ...grpc.CallOption) (*ListDeviceOperationsResponse, error)
}

type deviceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDeviceServiceClient(cc grpc.ClientConnInterface) DeviceServiceClient {
	return &deviceServiceClient{cc}
}

func (c *deviceServiceClient) Get(ctx context.Context, in *GetDeviceRequest, opts ...grpc.CallOption) (*Device, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Device)
	err := c.cc.Invoke(ctx, DeviceService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) GetByName(ctx context.Context, in *GetByNameDeviceRequest, opts ...grpc.CallOption) (*Device, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Device)
	err := c.cc.Invoke(ctx, DeviceService_GetByName_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) List(ctx context.Context, in *ListDevicesRequest, opts ...grpc.CallOption) (*ListDevicesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDevicesResponse)
	err := c.cc.Invoke(ctx, DeviceService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Create(ctx context.Context, in *CreateDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Update(ctx context.Context, in *UpdateDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) Delete(ctx context.Context, in *DeleteDeviceRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) ListCertificates(ctx context.Context, in *ListDeviceCertificatesRequest, opts ...grpc.CallOption) (*ListDeviceCertificatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDeviceCertificatesResponse)
	err := c.cc.Invoke(ctx, DeviceService_ListCertificates_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) AddCertificate(ctx context.Context, in *AddDeviceCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_AddCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) DeleteCertificate(ctx context.Context, in *DeleteDeviceCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_DeleteCertificate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) ListPasswords(ctx context.Context, in *ListDevicePasswordsRequest, opts ...grpc.CallOption) (*ListDevicePasswordsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDevicePasswordsResponse)
	err := c.cc.Invoke(ctx, DeviceService_ListPasswords_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) AddPassword(ctx context.Context, in *AddDevicePasswordRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_AddPassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) DeletePassword(ctx context.Context, in *DeleteDevicePasswordRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DeviceService_DeletePassword_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceServiceClient) ListOperations(ctx context.Context, in *ListDeviceOperationsRequest, opts ...grpc.CallOption) (*ListDeviceOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDeviceOperationsResponse)
	err := c.cc.Invoke(ctx, DeviceService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DeviceServiceServer is the server API for DeviceService service.
// All implementations should embed UnimplementedDeviceServiceServer
// for forward compatibility.
//
// A set of methods for managing devices.
type DeviceServiceServer interface {
	// Returns the specified device.
	//
	// To get the list of available devices, make a [List] request.
	Get(context.Context, *GetDeviceRequest) (*Device, error)
	GetByName(context.Context, *GetByNameDeviceRequest) (*Device, error)
	// Retrieves the list of devices in the specified registry.
	List(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error)
	// Creates a device in the specified registry.
	Create(context.Context, *CreateDeviceRequest) (*operation.Operation, error)
	// Updates the specified device.
	Update(context.Context, *UpdateDeviceRequest) (*operation.Operation, error)
	// Deletes the specified device.
	Delete(context.Context, *DeleteDeviceRequest) (*operation.Operation, error)
	// Retrieves the list of device certificates for the specified device.
	ListCertificates(context.Context, *ListDeviceCertificatesRequest) (*ListDeviceCertificatesResponse, error)
	// Adds a certificate.
	AddCertificate(context.Context, *AddDeviceCertificateRequest) (*operation.Operation, error)
	// Deletes the specified device certificate.
	DeleteCertificate(context.Context, *DeleteDeviceCertificateRequest) (*operation.Operation, error)
	// Retrieves the list of passwords for the specified device.
	ListPasswords(context.Context, *ListDevicePasswordsRequest) (*ListDevicePasswordsResponse, error)
	// Adds password for the specified device.
	AddPassword(context.Context, *AddDevicePasswordRequest) (*operation.Operation, error)
	// Deletes the specified password.
	DeletePassword(context.Context, *DeleteDevicePasswordRequest) (*operation.Operation, error)
	// Lists operations for the specified device.
	ListOperations(context.Context, *ListDeviceOperationsRequest) (*ListDeviceOperationsResponse, error)
}

// UnimplementedDeviceServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDeviceServiceServer struct{}

func (UnimplementedDeviceServiceServer) Get(context.Context, *GetDeviceRequest) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDeviceServiceServer) GetByName(context.Context, *GetByNameDeviceRequest) (*Device, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByName not implemented")
}
func (UnimplementedDeviceServiceServer) List(context.Context, *ListDevicesRequest) (*ListDevicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDeviceServiceServer) Create(context.Context, *CreateDeviceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDeviceServiceServer) Update(context.Context, *UpdateDeviceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDeviceServiceServer) Delete(context.Context, *DeleteDeviceRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDeviceServiceServer) ListCertificates(context.Context, *ListDeviceCertificatesRequest) (*ListDeviceCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCertificates not implemented")
}
func (UnimplementedDeviceServiceServer) AddCertificate(context.Context, *AddDeviceCertificateRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCertificate not implemented")
}
func (UnimplementedDeviceServiceServer) DeleteCertificate(context.Context, *DeleteDeviceCertificateRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCertificate not implemented")
}
func (UnimplementedDeviceServiceServer) ListPasswords(context.Context, *ListDevicePasswordsRequest) (*ListDevicePasswordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPasswords not implemented")
}
func (UnimplementedDeviceServiceServer) AddPassword(context.Context, *AddDevicePasswordRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPassword not implemented")
}
func (UnimplementedDeviceServiceServer) DeletePassword(context.Context, *DeleteDevicePasswordRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePassword not implemented")
}
func (UnimplementedDeviceServiceServer) ListOperations(context.Context, *ListDeviceOperationsRequest) (*ListDeviceOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedDeviceServiceServer) testEmbeddedByValue() {}

// UnsafeDeviceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DeviceServiceServer will
// result in compilation errors.
type UnsafeDeviceServiceServer interface {
	mustEmbedUnimplementedDeviceServiceServer()
}

func RegisterDeviceServiceServer(s grpc.ServiceRegistrar, srv DeviceServiceServer) {
	// If the following call pancis, it indicates UnimplementedDeviceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DeviceService_ServiceDesc, srv)
}

func _DeviceService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Get(ctx, req.(*GetDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_GetByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByNameDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).GetByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_GetByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).GetByName(ctx, req.(*GetByNameDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).List(ctx, req.(*ListDevicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Create(ctx, req.(*CreateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Update(ctx, req.(*UpdateDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).Delete(ctx, req.(*DeleteDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_ListCertificates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).ListCertificates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_ListCertificates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).ListCertificates(ctx, req.(*ListDeviceCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_AddCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).AddCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_AddCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).AddCertificate(ctx, req.(*AddDeviceCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_DeleteCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDeviceCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).DeleteCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_DeleteCertificate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).DeleteCertificate(ctx, req.(*DeleteDeviceCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_ListPasswords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDevicePasswordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).ListPasswords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_ListPasswords_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).ListPasswords(ctx, req.(*ListDevicePasswordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_AddPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDevicePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).AddPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_AddPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).AddPassword(ctx, req.(*AddDevicePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_DeletePassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDevicePasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).DeletePassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_DeletePassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).DeletePassword(ctx, req.(*DeleteDevicePasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DeviceService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDeviceOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DeviceServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DeviceService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DeviceServiceServer).ListOperations(ctx, req.(*ListDeviceOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DeviceService_ServiceDesc is the grpc.ServiceDesc for DeviceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DeviceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.iot.devices.v1.DeviceService",
	HandlerType: (*DeviceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _DeviceService_Get_Handler,
		},
		{
			MethodName: "GetByName",
			Handler:    _DeviceService_GetByName_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DeviceService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DeviceService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DeviceService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DeviceService_Delete_Handler,
		},
		{
			MethodName: "ListCertificates",
			Handler:    _DeviceService_ListCertificates_Handler,
		},
		{
			MethodName: "AddCertificate",
			Handler:    _DeviceService_AddCertificate_Handler,
		},
		{
			MethodName: "DeleteCertificate",
			Handler:    _DeviceService_DeleteCertificate_Handler,
		},
		{
			MethodName: "ListPasswords",
			Handler:    _DeviceService_ListPasswords_Handler,
		},
		{
			MethodName: "AddPassword",
			Handler:    _DeviceService_AddPassword_Handler,
		},
		{
			MethodName: "DeletePassword",
			Handler:    _DeviceService_DeletePassword_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _DeviceService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/iot/devices/v1/device_service.proto",
}
