// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.17.3
// source: yandex/cloud/ai/foundation_models/v1/foundation_models_service.proto

package foundation_models

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

const (
	TextGenerationService_Completion_FullMethodName = "/yandex.cloud.ai.foundation_models.v1.TextGenerationService/Completion"
)

// TextGenerationServiceClient is the client API for TextGenerationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextGenerationServiceClient interface {
	// RPC method for generating text completions.
	Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (TextGenerationService_CompletionClient, error)
}

type textGenerationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextGenerationServiceClient(cc grpc.ClientConnInterface) TextGenerationServiceClient {
	return &textGenerationServiceClient{cc}
}

func (c *textGenerationServiceClient) Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (TextGenerationService_CompletionClient, error) {
	stream, err := c.cc.NewStream(ctx, &TextGenerationService_ServiceDesc.Streams[0], TextGenerationService_Completion_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &textGenerationServiceCompletionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TextGenerationService_CompletionClient interface {
	Recv() (*CompletionResponse, error)
	grpc.ClientStream
}

type textGenerationServiceCompletionClient struct {
	grpc.ClientStream
}

func (x *textGenerationServiceCompletionClient) Recv() (*CompletionResponse, error) {
	m := new(CompletionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TextGenerationServiceServer is the server API for TextGenerationService service.
// All implementations should embed UnimplementedTextGenerationServiceServer
// for forward compatibility
type TextGenerationServiceServer interface {
	// RPC method for generating text completions.
	Completion(*CompletionRequest, TextGenerationService_CompletionServer) error
}

// UnimplementedTextGenerationServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTextGenerationServiceServer struct {
}

func (UnimplementedTextGenerationServiceServer) Completion(*CompletionRequest, TextGenerationService_CompletionServer) error {
	return status.Errorf(codes.Unimplemented, "method Completion not implemented")
}

// UnsafeTextGenerationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextGenerationServiceServer will
// result in compilation errors.
type UnsafeTextGenerationServiceServer interface {
	mustEmbedUnimplementedTextGenerationServiceServer()
}

func RegisterTextGenerationServiceServer(s grpc.ServiceRegistrar, srv TextGenerationServiceServer) {
	s.RegisterService(&TextGenerationService_ServiceDesc, srv)
}

func _TextGenerationService_Completion_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CompletionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TextGenerationServiceServer).Completion(m, &textGenerationServiceCompletionServer{stream})
}

type TextGenerationService_CompletionServer interface {
	Send(*CompletionResponse) error
	grpc.ServerStream
}

type textGenerationServiceCompletionServer struct {
	grpc.ServerStream
}

func (x *textGenerationServiceCompletionServer) Send(m *CompletionResponse) error {
	return x.ServerStream.SendMsg(m)
}

// TextGenerationService_ServiceDesc is the grpc.ServiceDesc for TextGenerationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextGenerationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.foundation_models.v1.TextGenerationService",
	HandlerType: (*TextGenerationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Completion",
			Handler:       _TextGenerationService_Completion_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "yandex/cloud/ai/foundation_models/v1/foundation_models_service.proto",
}

const (
	TextGenerationAsyncService_Completion_FullMethodName = "/yandex.cloud.ai.foundation_models.v1.TextGenerationAsyncService/Completion"
)

// TextGenerationAsyncServiceClient is the client API for TextGenerationAsyncService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextGenerationAsyncServiceClient interface {
	// RPC method for generating text completions in asynchronous mode.
	Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*operation.Operation, error)
}

type textGenerationAsyncServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextGenerationAsyncServiceClient(cc grpc.ClientConnInterface) TextGenerationAsyncServiceClient {
	return &textGenerationAsyncServiceClient{cc}
}

func (c *textGenerationAsyncServiceClient) Completion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, TextGenerationAsyncService_Completion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextGenerationAsyncServiceServer is the server API for TextGenerationAsyncService service.
// All implementations should embed UnimplementedTextGenerationAsyncServiceServer
// for forward compatibility
type TextGenerationAsyncServiceServer interface {
	// RPC method for generating text completions in asynchronous mode.
	Completion(context.Context, *CompletionRequest) (*operation.Operation, error)
}

// UnimplementedTextGenerationAsyncServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTextGenerationAsyncServiceServer struct {
}

func (UnimplementedTextGenerationAsyncServiceServer) Completion(context.Context, *CompletionRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Completion not implemented")
}

// UnsafeTextGenerationAsyncServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextGenerationAsyncServiceServer will
// result in compilation errors.
type UnsafeTextGenerationAsyncServiceServer interface {
	mustEmbedUnimplementedTextGenerationAsyncServiceServer()
}

func RegisterTextGenerationAsyncServiceServer(s grpc.ServiceRegistrar, srv TextGenerationAsyncServiceServer) {
	s.RegisterService(&TextGenerationAsyncService_ServiceDesc, srv)
}

func _TextGenerationAsyncService_Completion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationAsyncServiceServer).Completion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TextGenerationAsyncService_Completion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationAsyncServiceServer).Completion(ctx, req.(*CompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextGenerationAsyncService_ServiceDesc is the grpc.ServiceDesc for TextGenerationAsyncService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextGenerationAsyncService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.foundation_models.v1.TextGenerationAsyncService",
	HandlerType: (*TextGenerationAsyncServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Completion",
			Handler:    _TextGenerationAsyncService_Completion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/foundation_models/v1/foundation_models_service.proto",
}

const (
	TokenizerService_Tokenize_FullMethodName           = "/yandex.cloud.ai.foundation_models.v1.TokenizerService/Tokenize"
	TokenizerService_TokenizeCompletion_FullMethodName = "/yandex.cloud.ai.foundation_models.v1.TokenizerService/TokenizeCompletion"
)

// TokenizerServiceClient is the client API for TokenizerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenizerServiceClient interface {
	// RPC method for tokenizing text.
	Tokenize(ctx context.Context, in *TokenizeRequest, opts ...grpc.CallOption) (*TokenizeResponse, error)
	// RPC method for tokenizing content of CompletionRequest
	TokenizeCompletion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*TokenizeResponse, error)
}

type tokenizerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenizerServiceClient(cc grpc.ClientConnInterface) TokenizerServiceClient {
	return &tokenizerServiceClient{cc}
}

func (c *tokenizerServiceClient) Tokenize(ctx context.Context, in *TokenizeRequest, opts ...grpc.CallOption) (*TokenizeResponse, error) {
	out := new(TokenizeResponse)
	err := c.cc.Invoke(ctx, TokenizerService_Tokenize_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenizerServiceClient) TokenizeCompletion(ctx context.Context, in *CompletionRequest, opts ...grpc.CallOption) (*TokenizeResponse, error) {
	out := new(TokenizeResponse)
	err := c.cc.Invoke(ctx, TokenizerService_TokenizeCompletion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TokenizerServiceServer is the server API for TokenizerService service.
// All implementations should embed UnimplementedTokenizerServiceServer
// for forward compatibility
type TokenizerServiceServer interface {
	// RPC method for tokenizing text.
	Tokenize(context.Context, *TokenizeRequest) (*TokenizeResponse, error)
	// RPC method for tokenizing content of CompletionRequest
	TokenizeCompletion(context.Context, *CompletionRequest) (*TokenizeResponse, error)
}

// UnimplementedTokenizerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTokenizerServiceServer struct {
}

func (UnimplementedTokenizerServiceServer) Tokenize(context.Context, *TokenizeRequest) (*TokenizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Tokenize not implemented")
}
func (UnimplementedTokenizerServiceServer) TokenizeCompletion(context.Context, *CompletionRequest) (*TokenizeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TokenizeCompletion not implemented")
}

// UnsafeTokenizerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenizerServiceServer will
// result in compilation errors.
type UnsafeTokenizerServiceServer interface {
	mustEmbedUnimplementedTokenizerServiceServer()
}

func RegisterTokenizerServiceServer(s grpc.ServiceRegistrar, srv TokenizerServiceServer) {
	s.RegisterService(&TokenizerService_ServiceDesc, srv)
}

func _TokenizerService_Tokenize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenizeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenizerServiceServer).Tokenize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenizerService_Tokenize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenizerServiceServer).Tokenize(ctx, req.(*TokenizeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenizerService_TokenizeCompletion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompletionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenizerServiceServer).TokenizeCompletion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TokenizerService_TokenizeCompletion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenizerServiceServer).TokenizeCompletion(ctx, req.(*CompletionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TokenizerService_ServiceDesc is the grpc.ServiceDesc for TokenizerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenizerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.foundation_models.v1.TokenizerService",
	HandlerType: (*TokenizerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Tokenize",
			Handler:    _TokenizerService_Tokenize_Handler,
		},
		{
			MethodName: "TokenizeCompletion",
			Handler:    _TokenizerService_TokenizeCompletion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/foundation_models/v1/foundation_models_service.proto",
}

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
	Metadata: "yandex/cloud/ai/foundation_models/v1/foundation_models_service.proto",
}
