// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/cloudregistry/v1/artifact_service.proto

package cloudregistry

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
	ArtifactService_Get_FullMethodName    = "/yandex.cloud.cloudregistry.v1.ArtifactService/Get"
	ArtifactService_Delete_FullMethodName = "/yandex.cloud.cloudregistry.v1.ArtifactService/Delete"
)

// ArtifactServiceClient is the client API for ArtifactService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing Artifacts.
type ArtifactServiceClient interface {
	// Returns the specified artifact resource.
	//
	// To get the list of available artifact resources, make [RegistryService.ListArtifacts] method call.
	Get(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*Artifact, error)
	// Deletes the specified artifact.
	Delete(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type artifactServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewArtifactServiceClient(cc grpc.ClientConnInterface) ArtifactServiceClient {
	return &artifactServiceClient{cc}
}

func (c *artifactServiceClient) Get(ctx context.Context, in *GetArtifactRequest, opts ...grpc.CallOption) (*Artifact, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Artifact)
	err := c.cc.Invoke(ctx, ArtifactService_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artifactServiceClient) Delete(ctx context.Context, in *DeleteArtifactRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, ArtifactService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArtifactServiceServer is the server API for ArtifactService service.
// All implementations should embed UnimplementedArtifactServiceServer
// for forward compatibility.
//
// A set of methods for managing Artifacts.
type ArtifactServiceServer interface {
	// Returns the specified artifact resource.
	//
	// To get the list of available artifact resources, make [RegistryService.ListArtifacts] method call.
	Get(context.Context, *GetArtifactRequest) (*Artifact, error)
	// Deletes the specified artifact.
	Delete(context.Context, *DeleteArtifactRequest) (*operation.Operation, error)
}

// UnimplementedArtifactServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedArtifactServiceServer struct{}

func (UnimplementedArtifactServiceServer) Get(context.Context, *GetArtifactRequest) (*Artifact, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedArtifactServiceServer) Delete(context.Context, *DeleteArtifactRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedArtifactServiceServer) testEmbeddedByValue() {}

// UnsafeArtifactServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArtifactServiceServer will
// result in compilation errors.
type UnsafeArtifactServiceServer interface {
	mustEmbedUnimplementedArtifactServiceServer()
}

func RegisterArtifactServiceServer(s grpc.ServiceRegistrar, srv ArtifactServiceServer) {
	// If the following call pancis, it indicates UnimplementedArtifactServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ArtifactService_ServiceDesc, srv)
}

func _ArtifactService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactServiceServer).Get(ctx, req.(*GetArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtifactService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteArtifactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtifactServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ArtifactService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtifactServiceServer).Delete(ctx, req.(*DeleteArtifactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArtifactService_ServiceDesc is the grpc.ServiceDesc for ArtifactService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArtifactService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.cloudregistry.v1.ArtifactService",
	HandlerType: (*ArtifactServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ArtifactService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ArtifactService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/cloudregistry/v1/artifact_service.proto",
}