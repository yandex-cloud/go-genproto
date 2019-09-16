// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/operation/operation_service.proto

package operation

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetOperationRequest struct {
	// ID of the Operation resource to return.
	OperationId          string   `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationRequest) Reset()         { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()    {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65df5639ee3bf772, []int{0}
}

func (m *GetOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationRequest.Unmarshal(m, b)
}
func (m *GetOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationRequest.Merge(m, src)
}
func (m *GetOperationRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationRequest.Size(m)
}
func (m *GetOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationRequest proto.InternalMessageInfo

func (m *GetOperationRequest) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

type CancelOperationRequest struct {
	// ID of the operation to cancel.
	OperationId          string   `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CancelOperationRequest) Reset()         { *m = CancelOperationRequest{} }
func (m *CancelOperationRequest) String() string { return proto.CompactTextString(m) }
func (*CancelOperationRequest) ProtoMessage()    {}
func (*CancelOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65df5639ee3bf772, []int{1}
}

func (m *CancelOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CancelOperationRequest.Unmarshal(m, b)
}
func (m *CancelOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CancelOperationRequest.Marshal(b, m, deterministic)
}
func (m *CancelOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CancelOperationRequest.Merge(m, src)
}
func (m *CancelOperationRequest) XXX_Size() int {
	return xxx_messageInfo_CancelOperationRequest.Size(m)
}
func (m *CancelOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CancelOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CancelOperationRequest proto.InternalMessageInfo

func (m *CancelOperationRequest) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetOperationRequest)(nil), "yandex.cloud.operation.GetOperationRequest")
	proto.RegisterType((*CancelOperationRequest)(nil), "yandex.cloud.operation.CancelOperationRequest")
}

func init() {
	proto.RegisterFile("yandex/cloud/operation/operation_service.proto", fileDescriptor_65df5639ee3bf772)
}

var fileDescriptor_65df5639ee3bf772 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xab, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0xcf, 0x2f, 0x48, 0x2d, 0x4a, 0x2c, 0xc9,
	0xcc, 0xcf, 0x43, 0xb0, 0xe2, 0x8b, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0xf5, 0x0a, 0x8a, 0xf2,
	0x4b, 0xf2, 0x85, 0xc4, 0x20, 0xea, 0xf5, 0xc0, 0xea, 0xf5, 0xe0, 0xaa, 0xa4, 0x64, 0xd2, 0xf3,
	0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0xc0, 0xc2,
	0xc5, 0x10, 0x5d, 0x52, 0x6a, 0x84, 0x6c, 0x81, 0xaa, 0x93, 0x45, 0x51, 0x57, 0x96, 0x98, 0x93,
	0x99, 0x82, 0x24, 0xad, 0x64, 0xc7, 0x25, 0xec, 0x9e, 0x5a, 0xe2, 0x0f, 0xd3, 0x14, 0x94, 0x5a,
	0x58, 0x9a, 0x5a, 0x5c, 0x22, 0xa4, 0xce, 0xc5, 0x83, 0x70, 0x6e, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0xa7, 0x13, 0xcb, 0x8b, 0xe3, 0x86, 0x8c, 0x41, 0xdc, 0x70, 0x19, 0xcf, 0x14, 0x25,
	0x47, 0x2e, 0x31, 0xe7, 0xc4, 0xbc, 0xe4, 0xd4, 0x1c, 0xb2, 0x8d, 0x30, 0x9a, 0xc6, 0xc4, 0x25,
	0x00, 0xd7, 0x1d, 0x0c, 0x09, 0x1a, 0xa1, 0x4a, 0x2e, 0x66, 0xf7, 0xd4, 0x12, 0x21, 0x6d, 0x3d,
	0xec, 0x81, 0xa3, 0x87, 0xc5, 0xd1, 0x52, 0x8a, 0xb8, 0x14, 0xc3, 0x55, 0x2a, 0x29, 0x35, 0x5d,
	0x7e, 0x32, 0x99, 0x49, 0x46, 0x48, 0x0a, 0x11, 0x4e, 0xc5, 0xfa, 0xd5, 0xc8, 0xee, 0xac, 0x15,
	0x6a, 0x63, 0xe4, 0x62, 0x83, 0xf8, 0x49, 0x48, 0x0f, 0x97, 0x89, 0xd8, 0xfd, 0x4c, 0x8c, 0x0b,
	0x34, 0xc1, 0x2e, 0x50, 0x16, 0x52, 0xc4, 0xed, 0x02, 0xab, 0x64, 0xb0, 0xe9, 0x4e, 0x6e, 0x51,
	0x2e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0x93, 0x75, 0x21,
	0xf1, 0x98, 0x9e, 0xaf, 0x9b, 0x9e, 0x9a, 0x07, 0x8e, 0x42, 0x7d, 0xec, 0x09, 0xc1, 0x1a, 0xce,
	0x4a, 0x62, 0x03, 0xab, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xa7, 0x20, 0x3f, 0x99,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OperationServiceClient is the client API for OperationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OperationServiceClient interface {
	// Returns the specified Operation resource.
	Get(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error)
	// Cancels the specified operation.
	Cancel(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*Operation, error)
}

type operationServiceClient struct {
	cc *grpc.ClientConn
}

func NewOperationServiceClient(cc *grpc.ClientConn) OperationServiceClient {
	return &operationServiceClient{cc}
}

func (c *operationServiceClient) Get(ctx context.Context, in *GetOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.operation.OperationService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationServiceClient) Cancel(ctx context.Context, in *CancelOperationRequest, opts ...grpc.CallOption) (*Operation, error) {
	out := new(Operation)
	err := c.cc.Invoke(ctx, "/yandex.cloud.operation.OperationService/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationServiceServer is the server API for OperationService service.
type OperationServiceServer interface {
	// Returns the specified Operation resource.
	Get(context.Context, *GetOperationRequest) (*Operation, error)
	// Cancels the specified operation.
	Cancel(context.Context, *CancelOperationRequest) (*Operation, error)
}

// UnimplementedOperationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOperationServiceServer struct {
}

func (*UnimplementedOperationServiceServer) Get(ctx context.Context, req *GetOperationRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedOperationServiceServer) Cancel(ctx context.Context, req *CancelOperationRequest) (*Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}

func RegisterOperationServiceServer(s *grpc.Server, srv OperationServiceServer) {
	s.RegisterService(&_OperationService_serviceDesc, srv)
}

func _OperationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.operation.OperationService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Get(ctx, req.(*GetOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationService_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelOperationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationServiceServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.operation.OperationService/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationServiceServer).Cancel(ctx, req.(*CancelOperationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OperationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.operation.OperationService",
	HandlerType: (*OperationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _OperationService_Get_Handler,
		},
		{
			MethodName: "Cancel",
			Handler:    _OperationService_Cancel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/operation/operation_service.proto",
}
