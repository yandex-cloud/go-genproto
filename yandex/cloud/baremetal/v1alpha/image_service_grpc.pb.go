// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/baremetal/v1alpha/image_service.proto

package baremetal

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
	ImageService_Get_FullMethodName            = "/yandex.cloud.baremetal.v1alpha.ImageService/Get"
	ImageService_List_FullMethodName           = "/yandex.cloud.baremetal.v1alpha.ImageService/List"
	ImageService_Create_FullMethodName         = "/yandex.cloud.baremetal.v1alpha.ImageService/Create"
	ImageService_Update_FullMethodName         = "/yandex.cloud.baremetal.v1alpha.ImageService/Update"
	ImageService_Delete_FullMethodName         = "/yandex.cloud.baremetal.v1alpha.ImageService/Delete"
	ImageService_ListOperations_FullMethodName = "/yandex.cloud.baremetal.v1alpha.ImageService/ListOperations"
)

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Image resources.
type ImageServiceClient interface {
	// Returns the specific Image resource.
	//
	// To get the list of available Image resources, make a [List] request.
	Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error)
	// Retrieves the list of Image resources in the specified folder.
	List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error)
	// Creates an image in the specified folder.
	Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Updates the specified image.
	Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Deletes the specified image.
	//
	// Deleting an image removes its data permanently and is irreversible.
	Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Lists operations for the specified image.
	ListOperations(ctx context.Context, in *ListImageOperationsRequest, opts ...grpc.CallOption) (*ListImageOperationsResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) Get(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (*Image, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Image)
	err := c.cc.Invoke(ctx, ImageService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) List(ctx context.Context, in *ListImagesRequest, opts ...grpc.CallOption) (*ListImagesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListImagesResponse)
	err := c.cc.Invoke(ctx, ImageService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) Create(ctx context.Context, in *CreateImageRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ImageService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) Update(ctx context.Context, in *UpdateImageRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ImageService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) Delete(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ImageService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ListOperations(ctx context.Context, in *ListImageOperationsRequest, opts ...grpc.CallOption) (*ListImageOperationsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListImageOperationsResponse)
	err := c.cc.Invoke(ctx, ImageService_ListOperations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations should embed UnimplementedImageServiceServer
// for forward compatibility.
//
// A set of methods for managing Image resources.
type ImageServiceServer interface {
	// Returns the specific Image resource.
	//
	// To get the list of available Image resources, make a [List] request.
	Get(context.Context, *GetImageRequest) (*Image, error)
	// Retrieves the list of Image resources in the specified folder.
	List(context.Context, *ListImagesRequest) (*ListImagesResponse, error)
	// Creates an image in the specified folder.
	Create(context.Context, *CreateImageRequest) (*operation.Operation, error)
	// Updates the specified image.
	Update(context.Context, *UpdateImageRequest) (*operation.Operation, error)
	// Deletes the specified image.
	//
	// Deleting an image removes its data permanently and is irreversible.
	Delete(context.Context, *DeleteImageRequest) (*operation.Operation, error)
	// Lists operations for the specified image.
	ListOperations(context.Context, *ListImageOperationsRequest) (*ListImageOperationsResponse, error)
}

// UnimplementedImageServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageServiceServer struct{}

func (UnimplementedImageServiceServer) Get(context.Context, *GetImageRequest) (*Image, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedImageServiceServer) List(context.Context, *ListImagesRequest) (*ListImagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedImageServiceServer) Create(context.Context, *CreateImageRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedImageServiceServer) Update(context.Context, *UpdateImageRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedImageServiceServer) Delete(context.Context, *DeleteImageRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedImageServiceServer) ListOperations(context.Context, *ListImageOperationsRequest) (*ListImageOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperations not implemented")
}
func (UnimplementedImageServiceServer) testEmbeddedByValue() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	// If the following call pancis, it indicates UnimplementedImageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Get(ctx, req.(*GetImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).List(ctx, req.(*ListImagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Create(ctx, req.(*CreateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Update(ctx, req.(*UpdateImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).Delete(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ListOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListImageOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ListOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ImageService_ListOperations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ListOperations(ctx, req.(*ListImageOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.baremetal.v1alpha.ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ImageService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ImageService_List_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ImageService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ImageService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ImageService_Delete_Handler,
		},
		{
			MethodName: "ListOperations",
			Handler:    _ImageService_ListOperations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/baremetal/v1alpha/image_service.proto",
}
