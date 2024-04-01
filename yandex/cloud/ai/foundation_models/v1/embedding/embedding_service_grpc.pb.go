// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/ai/foundation_models/v1/embedding/embedding_service.proto

package foundation_models

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

const (
	EmbeddingsService_TextEmbedding_FullMethodName = "/yandex.cloud.ai.foundation_models.v1.EmbeddingsService/TextEmbedding"
)

// EmbeddingsServiceClient is the client API for EmbeddingsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbeddingsServiceClient interface {
	// RPC method for obtaining embeddings from text data.
	TextEmbedding(ctx context.Context, in *TextEmbeddingRequest, opts ...grpc.CallOption) (*TextEmbeddingResponse, error)
}

type embeddingsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbeddingsServiceClient(cc grpc.ClientConnInterface) EmbeddingsServiceClient {
	return &embeddingsServiceClient{cc}
}

func (c *embeddingsServiceClient) TextEmbedding(ctx context.Context, in *TextEmbeddingRequest, opts ...grpc.CallOption) (*TextEmbeddingResponse, error) {
	out := new(TextEmbeddingResponse)
	err := c.cc.Invoke(ctx, EmbeddingsService_TextEmbedding_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmbeddingsServiceServer is the server API for EmbeddingsService service.
// All implementations should embed UnimplementedEmbeddingsServiceServer
// for forward compatibility
type EmbeddingsServiceServer interface {
	// RPC method for obtaining embeddings from text data.
	TextEmbedding(context.Context, *TextEmbeddingRequest) (*TextEmbeddingResponse, error)
}

// UnimplementedEmbeddingsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedEmbeddingsServiceServer struct {
}

func (UnimplementedEmbeddingsServiceServer) TextEmbedding(context.Context, *TextEmbeddingRequest) (*TextEmbeddingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextEmbedding not implemented")
}

// UnsafeEmbeddingsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbeddingsServiceServer will
// result in compilation errors.
type UnsafeEmbeddingsServiceServer interface {
	mustEmbedUnimplementedEmbeddingsServiceServer()
}

func RegisterEmbeddingsServiceServer(s grpc.ServiceRegistrar, srv EmbeddingsServiceServer) {
	s.RegisterService(&EmbeddingsService_ServiceDesc, srv)
}

func _EmbeddingsService_TextEmbedding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextEmbeddingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbeddingsServiceServer).TextEmbedding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmbeddingsService_TextEmbedding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbeddingsServiceServer).TextEmbedding(ctx, req.(*TextEmbeddingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmbeddingsService_ServiceDesc is the grpc.ServiceDesc for EmbeddingsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmbeddingsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.foundation_models.v1.EmbeddingsService",
	HandlerType: (*EmbeddingsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TextEmbedding",
			Handler:    _EmbeddingsService_TextEmbedding_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/foundation_models/v1/embedding/embedding_service.proto",
}