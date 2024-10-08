// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/kms/v1/symmetric_crypto_service.proto

package kms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SymmetricCryptoService_Encrypt_FullMethodName         = "/yandex.cloud.kms.v1.SymmetricCryptoService/Encrypt"
	SymmetricCryptoService_Decrypt_FullMethodName         = "/yandex.cloud.kms.v1.SymmetricCryptoService/Decrypt"
	SymmetricCryptoService_ReEncrypt_FullMethodName       = "/yandex.cloud.kms.v1.SymmetricCryptoService/ReEncrypt"
	SymmetricCryptoService_GenerateDataKey_FullMethodName = "/yandex.cloud.kms.v1.SymmetricCryptoService/GenerateDataKey"
)

// SymmetricCryptoServiceClient is the client API for SymmetricCryptoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Set of methods that perform symmetric encryption and decryption.
type SymmetricCryptoServiceClient interface {
	// Encrypts given plaintext with the specified key.
	Encrypt(ctx context.Context, in *SymmetricEncryptRequest, opts ...grpc.CallOption) (*SymmetricEncryptResponse, error)
	// Decrypts the given ciphertext with the specified key.
	Decrypt(ctx context.Context, in *SymmetricDecryptRequest, opts ...grpc.CallOption) (*SymmetricDecryptResponse, error)
	// Re-encrypts a ciphertext with the specified KMS key.
	ReEncrypt(ctx context.Context, in *SymmetricReEncryptRequest, opts ...grpc.CallOption) (*SymmetricReEncryptResponse, error)
	// Generates a new symmetric data encryption key (not a KMS key) and returns
	// the generated key as plaintext and as ciphertext encrypted with the specified symmetric KMS key.
	GenerateDataKey(ctx context.Context, in *GenerateDataKeyRequest, opts ...grpc.CallOption) (*GenerateDataKeyResponse, error)
}

type symmetricCryptoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSymmetricCryptoServiceClient(cc grpc.ClientConnInterface) SymmetricCryptoServiceClient {
	return &symmetricCryptoServiceClient{cc}
}

func (c *symmetricCryptoServiceClient) Encrypt(ctx context.Context, in *SymmetricEncryptRequest, opts ...grpc.CallOption) (*SymmetricEncryptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymmetricEncryptResponse)
	err := c.cc.Invoke(ctx, SymmetricCryptoService_Encrypt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *symmetricCryptoServiceClient) Decrypt(ctx context.Context, in *SymmetricDecryptRequest, opts ...grpc.CallOption) (*SymmetricDecryptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymmetricDecryptResponse)
	err := c.cc.Invoke(ctx, SymmetricCryptoService_Decrypt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *symmetricCryptoServiceClient) ReEncrypt(ctx context.Context, in *SymmetricReEncryptRequest, opts ...grpc.CallOption) (*SymmetricReEncryptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SymmetricReEncryptResponse)
	err := c.cc.Invoke(ctx, SymmetricCryptoService_ReEncrypt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *symmetricCryptoServiceClient) GenerateDataKey(ctx context.Context, in *GenerateDataKeyRequest, opts ...grpc.CallOption) (*GenerateDataKeyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenerateDataKeyResponse)
	err := c.cc.Invoke(ctx, SymmetricCryptoService_GenerateDataKey_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SymmetricCryptoServiceServer is the server API for SymmetricCryptoService service.
// All implementations should embed UnimplementedSymmetricCryptoServiceServer
// for forward compatibility.
//
// Set of methods that perform symmetric encryption and decryption.
type SymmetricCryptoServiceServer interface {
	// Encrypts given plaintext with the specified key.
	Encrypt(context.Context, *SymmetricEncryptRequest) (*SymmetricEncryptResponse, error)
	// Decrypts the given ciphertext with the specified key.
	Decrypt(context.Context, *SymmetricDecryptRequest) (*SymmetricDecryptResponse, error)
	// Re-encrypts a ciphertext with the specified KMS key.
	ReEncrypt(context.Context, *SymmetricReEncryptRequest) (*SymmetricReEncryptResponse, error)
	// Generates a new symmetric data encryption key (not a KMS key) and returns
	// the generated key as plaintext and as ciphertext encrypted with the specified symmetric KMS key.
	GenerateDataKey(context.Context, *GenerateDataKeyRequest) (*GenerateDataKeyResponse, error)
}

// UnimplementedSymmetricCryptoServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSymmetricCryptoServiceServer struct{}

func (UnimplementedSymmetricCryptoServiceServer) Encrypt(context.Context, *SymmetricEncryptRequest) (*SymmetricEncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Encrypt not implemented")
}
func (UnimplementedSymmetricCryptoServiceServer) Decrypt(context.Context, *SymmetricDecryptRequest) (*SymmetricDecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedSymmetricCryptoServiceServer) ReEncrypt(context.Context, *SymmetricReEncryptRequest) (*SymmetricReEncryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReEncrypt not implemented")
}
func (UnimplementedSymmetricCryptoServiceServer) GenerateDataKey(context.Context, *GenerateDataKeyRequest) (*GenerateDataKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateDataKey not implemented")
}
func (UnimplementedSymmetricCryptoServiceServer) testEmbeddedByValue() {}

// UnsafeSymmetricCryptoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SymmetricCryptoServiceServer will
// result in compilation errors.
type UnsafeSymmetricCryptoServiceServer interface {
	mustEmbedUnimplementedSymmetricCryptoServiceServer()
}

func RegisterSymmetricCryptoServiceServer(s grpc.ServiceRegistrar, srv SymmetricCryptoServiceServer) {
	// If the following call pancis, it indicates UnimplementedSymmetricCryptoServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SymmetricCryptoService_ServiceDesc, srv)
}

func _SymmetricCryptoService_Encrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymmetricEncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).Encrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SymmetricCryptoService_Encrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).Encrypt(ctx, req.(*SymmetricEncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SymmetricCryptoService_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymmetricDecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SymmetricCryptoService_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).Decrypt(ctx, req.(*SymmetricDecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SymmetricCryptoService_ReEncrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SymmetricReEncryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).ReEncrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SymmetricCryptoService_ReEncrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).ReEncrypt(ctx, req.(*SymmetricReEncryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SymmetricCryptoService_GenerateDataKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateDataKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SymmetricCryptoServiceServer).GenerateDataKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SymmetricCryptoService_GenerateDataKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SymmetricCryptoServiceServer).GenerateDataKey(ctx, req.(*GenerateDataKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SymmetricCryptoService_ServiceDesc is the grpc.ServiceDesc for SymmetricCryptoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SymmetricCryptoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.kms.v1.SymmetricCryptoService",
	HandlerType: (*SymmetricCryptoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Encrypt",
			Handler:    _SymmetricCryptoService_Encrypt_Handler,
		},
		{
			MethodName: "Decrypt",
			Handler:    _SymmetricCryptoService_Decrypt_Handler,
		},
		{
			MethodName: "ReEncrypt",
			Handler:    _SymmetricCryptoService_ReEncrypt_Handler,
		},
		{
			MethodName: "GenerateDataKey",
			Handler:    _SymmetricCryptoService_GenerateDataKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/kms/v1/symmetric_crypto_service.proto",
}
