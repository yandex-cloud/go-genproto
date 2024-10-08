// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/certificatemanager/v1/certificate_service.proto

package certificatemanager

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
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CertificateService_Get_FullMethodName                  = "/yandex.cloud.certificatemanager.v1.CertificateService/Get"
	CertificateService_List_FullMethodName                 = "/yandex.cloud.certificatemanager.v1.CertificateService/List"
	CertificateService_ListVersions_FullMethodName         = "/yandex.cloud.certificatemanager.v1.CertificateService/ListVersions"
	CertificateService_Create_FullMethodName               = "/yandex.cloud.certificatemanager.v1.CertificateService/Create"
	CertificateService_Update_FullMethodName               = "/yandex.cloud.certificatemanager.v1.CertificateService/Update"
	CertificateService_Delete_FullMethodName               = "/yandex.cloud.certificatemanager.v1.CertificateService/Delete"
	CertificateService_RequestNew_FullMethodName           = "/yandex.cloud.certificatemanager.v1.CertificateService/RequestNew"
	CertificateService_ListOperations_FullMethodName       = "/yandex.cloud.certificatemanager.v1.CertificateService/ListOperations"
	CertificateService_ListAccessBindings_FullMethodName   = "/yandex.cloud.certificatemanager.v1.CertificateService/ListAccessBindings"
	CertificateService_SetAccessBindings_FullMethodName    = "/yandex.cloud.certificatemanager.v1.CertificateService/SetAccessBindings"
	CertificateService_UpdateAccessBindings_FullMethodName = "/yandex.cloud.certificatemanager.v1.CertificateService/UpdateAccessBindings"
)

// CertificateServiceClient is the client API for CertificateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing certificates.
type CertificateServiceClient interface {
	// Returns the specified certificate.
	//
	// To get the list of available certificates, make a [List] request.
	Get(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*Certificate, error)
	// Returns the list of certificates in the specified folder.
	List(ctx context.Context, in *ListCertificatesRequest, opts ...grpc.CallOption) (*ListCertificatesResponse, error)
	ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error)
	// Creates a certificate in the specified folder.
	Create(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified certificate.
	Update(ctx context.Context, in *UpdateCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified certificate.
	Delete(ctx context.Context, in *DeleteCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Request a certificate in the specified folder.
	RequestNew(ctx context.Context, in *RequestNewCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified certificate.
	ListOperations(ctx context.Context, in *ListCertificateOperationsRequest, opts ...grpc.CallOption) (*ListCertificateOperationsResponse, error)
	// Lists existing access bindings for the specified certificate.
	ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the certificate.
	SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates access bindings for the specified certificate.
	UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type certificateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateServiceClient(cc grpc.ClientConnInterface) CertificateServiceClient {
	return &certificateServiceClient{cc}
}

func (c *certificateServiceClient) Get(ctx context.Context, in *GetCertificateRequest, opts ...grpc.CallOption) (*Certificate, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Certificate)
	err := c.cc.Invoke(ctx, CertificateService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) List(ctx context.Context, in *ListCertificatesRequest, opts ...grpc.CallOption) (*ListCertificatesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCertificatesResponse)
	err := c.cc.Invoke(ctx, CertificateService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) ListVersions(ctx context.Context, in *ListVersionsRequest, opts ...grpc.CallOption) (*ListVersionsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListVersionsResponse)
	err := c.cc.Invoke(ctx, CertificateService_ListVersions_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) Create(ctx context.Context, in *CreateCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CertificateService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) Update(ctx context.Context, in *UpdateCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CertificateService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) Delete(ctx context.Context, in *DeleteCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CertificateService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) RequestNew(ctx context.Context, in *RequestNewCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CertificateService_RequestNew_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) ListOperations(ctx context.Context, in *ListCertificateOperationsRequest, opts ...grpc.CallOption) (*ListCertificateOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCertificateOperationsResponse)
	err := c.cc.Invoke(ctx, CertificateService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) ListAccessBindings(ctx context.Context, in *access.ListAccessBindingsRequest, opts ...grpc.CallOption) (*access.ListAccessBindingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(access.ListAccessBindingsResponse)
	err := c.cc.Invoke(ctx, CertificateService_ListAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) SetAccessBindings(ctx context.Context, in *access.SetAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CertificateService_SetAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateServiceClient) UpdateAccessBindings(ctx context.Context, in *access.UpdateAccessBindingsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, CertificateService_UpdateAccessBindings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateServiceServer is the server API for CertificateService service.
// All implementations should embed UnimplementedCertificateServiceServer
// for forward compatibility.
//
// A set of methods for managing certificates.
type CertificateServiceServer interface {
	// Returns the specified certificate.
	//
	// To get the list of available certificates, make a [List] request.
	Get(context.Context, *GetCertificateRequest) (*Certificate, error)
	// Returns the list of certificates in the specified folder.
	List(context.Context, *ListCertificatesRequest) (*ListCertificatesResponse, error)
	ListVersions(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error)
	// Creates a certificate in the specified folder.
	Create(context.Context, *CreateCertificateRequest) (*operation.Operation, error)
	// Updates the specified certificate.
	Update(context.Context, *UpdateCertificateRequest) (*operation.Operation, error)
	// Deletes the specified certificate.
	Delete(context.Context, *DeleteCertificateRequest) (*operation.Operation, error)
	// Request a certificate in the specified folder.
	RequestNew(context.Context, *RequestNewCertificateRequest) (*operation.Operation, error)
	// Lists operations for the specified certificate.
	ListOperations(context.Context, *ListCertificateOperationsRequest) (*ListCertificateOperationsResponse, error)
	// Lists existing access bindings for the specified certificate.
	ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error)
	// Sets access bindings for the certificate.
	SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error)
	// Updates access bindings for the specified certificate.
	UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error)
}

// UnimplementedCertificateServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCertificateServiceServer struct{}

func (UnimplementedCertificateServiceServer) Get(context.Context, *GetCertificateRequest) (*Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCertificateServiceServer) List(context.Context, *ListCertificatesRequest) (*ListCertificatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedCertificateServiceServer) ListVersions(context.Context, *ListVersionsRequest) (*ListVersionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVersions not implemented")
}
func (UnimplementedCertificateServiceServer) Create(context.Context, *CreateCertificateRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCertificateServiceServer) Update(context.Context, *UpdateCertificateRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCertificateServiceServer) Delete(context.Context, *DeleteCertificateRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCertificateServiceServer) RequestNew(context.Context, *RequestNewCertificateRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestNew not implemented")
}
func (UnimplementedCertificateServiceServer) ListOperations(context.Context, *ListCertificateOperationsRequest) (*ListCertificateOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedCertificateServiceServer) ListAccessBindings(context.Context, *access.ListAccessBindingsRequest) (*access.ListAccessBindingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAccessBindings not implemented")
}
func (UnimplementedCertificateServiceServer) SetAccessBindings(context.Context, *access.SetAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetAccessBindings not implemented")
}
func (UnimplementedCertificateServiceServer) UpdateAccessBindings(context.Context, *access.UpdateAccessBindingsRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAccessBindings not implemented")
}
func (UnimplementedCertificateServiceServer) testEmbeddedByValue() {}

// UnsafeCertificateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateServiceServer will
// result in compilation errors.
type UnsafeCertificateServiceServer interface {
	mustEmbedUnimplementedCertificateServiceServer()
}

func RegisterCertificateServiceServer(s grpc.ServiceRegistrar, srv CertificateServiceServer) {
	// If the following call pancis, it indicates UnimplementedCertificateServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CertificateService_ServiceDesc, srv)
}

func _CertificateService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Get(ctx, req.(*GetCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).List(ctx, req.(*ListCertificatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_ListVersions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVersionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ListVersions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_ListVersions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ListVersions(ctx, req.(*ListVersionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Create(ctx, req.(*CreateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Update(ctx, req.(*UpdateCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).Delete(ctx, req.(*DeleteCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_RequestNew_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestNewCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).RequestNew(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_RequestNew_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).RequestNew(ctx, req.(*RequestNewCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCertificateOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ListOperations(ctx, req.(*ListCertificateOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_ListAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.ListAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).ListAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_ListAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).ListAccessBindings(ctx, req.(*access.ListAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_SetAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.SetAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).SetAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_SetAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).SetAccessBindings(ctx, req.(*access.SetAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateService_UpdateAccessBindings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(access.UpdateAccessBindingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateServiceServer).UpdateAccessBindings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CertificateService_UpdateAccessBindings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateServiceServer).UpdateAccessBindings(ctx, req.(*access.UpdateAccessBindingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CertificateService_ServiceDesc is the grpc.ServiceDesc for CertificateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.certificatemanager.v1.CertificateService",
	HandlerType: (*CertificateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _CertificateService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _CertificateService_List_Handler,
		},
		{
			MethodName: "ListVersions",
			Handler:    _CertificateService_ListVersions_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _CertificateService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CertificateService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateService_Delete_Handler,
		},
		{
			MethodName: "RequestNew",
			Handler:    _CertificateService_RequestNew_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _CertificateService_ListOperations_Handler,
		},
		{
			MethodName: "ListAccessBindings",
			Handler:    _CertificateService_ListAccessBindings_Handler,
		},
		{
			MethodName: "SetAccessBindings",
			Handler:    _CertificateService_SetAccessBindings_Handler,
		},
		{
			MethodName: "UpdateAccessBindings",
			Handler:    _CertificateService_UpdateAccessBindings_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/certificatemanager/v1/certificate_service.proto",
}
