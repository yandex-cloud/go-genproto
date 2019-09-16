// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/backup_service.proto

package mysql

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

type GetBackupRequest struct {
	// ID of the backup to return information about.
	// To get the backup ID, use a [ClusterService.ListBackups] request.
	BackupId             string   `protobuf:"bytes,1,opt,name=backup_id,json=backupId,proto3" json:"backup_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetBackupRequest) Reset()         { *m = GetBackupRequest{} }
func (m *GetBackupRequest) String() string { return proto.CompactTextString(m) }
func (*GetBackupRequest) ProtoMessage()    {}
func (*GetBackupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_221a1eb58843722c, []int{0}
}

func (m *GetBackupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetBackupRequest.Unmarshal(m, b)
}
func (m *GetBackupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetBackupRequest.Marshal(b, m, deterministic)
}
func (m *GetBackupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBackupRequest.Merge(m, src)
}
func (m *GetBackupRequest) XXX_Size() int {
	return xxx_messageInfo_GetBackupRequest.Size(m)
}
func (m *GetBackupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBackupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBackupRequest proto.InternalMessageInfo

func (m *GetBackupRequest) GetBackupId() string {
	if m != nil {
		return m.BackupId
	}
	return ""
}

type ListBackupsRequest struct {
	// ID of the folder to list backups in.
	// To get the folder ID, use a [yandex.cloud.resourcemanager.v1.FolderService.List] request.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListBackupsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token.  To get the next page of results, Set [page_token] to the [ListBackupsResponse.next_page_token]
	// returned by a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackupsRequest) Reset()         { *m = ListBackupsRequest{} }
func (m *ListBackupsRequest) String() string { return proto.CompactTextString(m) }
func (*ListBackupsRequest) ProtoMessage()    {}
func (*ListBackupsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_221a1eb58843722c, []int{1}
}

func (m *ListBackupsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackupsRequest.Unmarshal(m, b)
}
func (m *ListBackupsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackupsRequest.Marshal(b, m, deterministic)
}
func (m *ListBackupsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackupsRequest.Merge(m, src)
}
func (m *ListBackupsRequest) XXX_Size() int {
	return xxx_messageInfo_ListBackupsRequest.Size(m)
}
func (m *ListBackupsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackupsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackupsRequest proto.InternalMessageInfo

func (m *ListBackupsRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *ListBackupsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListBackupsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

type ListBackupsResponse struct {
	// List of MySQL backups.
	Backups []*Backup `protobuf:"bytes,1,rep,name=backups,proto3" json:"backups,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListBackupsRequest.page_size], use the [next_page_token] as the value
	// for the [ListBackupsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBackupsResponse) Reset()         { *m = ListBackupsResponse{} }
func (m *ListBackupsResponse) String() string { return proto.CompactTextString(m) }
func (*ListBackupsResponse) ProtoMessage()    {}
func (*ListBackupsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_221a1eb58843722c, []int{2}
}

func (m *ListBackupsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBackupsResponse.Unmarshal(m, b)
}
func (m *ListBackupsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBackupsResponse.Marshal(b, m, deterministic)
}
func (m *ListBackupsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBackupsResponse.Merge(m, src)
}
func (m *ListBackupsResponse) XXX_Size() int {
	return xxx_messageInfo_ListBackupsResponse.Size(m)
}
func (m *ListBackupsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBackupsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBackupsResponse proto.InternalMessageInfo

func (m *ListBackupsResponse) GetBackups() []*Backup {
	if m != nil {
		return m.Backups
	}
	return nil
}

func (m *ListBackupsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

func init() {
	proto.RegisterType((*GetBackupRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.GetBackupRequest")
	proto.RegisterType((*ListBackupsRequest)(nil), "yandex.cloud.mdb.mysql.v1alpha.ListBackupsRequest")
	proto.RegisterType((*ListBackupsResponse)(nil), "yandex.cloud.mdb.mysql.v1alpha.ListBackupsResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/backup_service.proto", fileDescriptor_221a1eb58843722c)
}

var fileDescriptor_221a1eb58843722c = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbf, 0x6f, 0x13, 0x31,
	0x14, 0xc7, 0xe5, 0x24, 0x94, 0x9c, 0xa1, 0x02, 0x99, 0x25, 0x8a, 0xa0, 0x0a, 0x37, 0x94, 0xf0,
	0x23, 0xe7, 0xbb, 0x44, 0x9d, 0x68, 0x25, 0x94, 0xa5, 0x14, 0x31, 0xa0, 0x2b, 0x13, 0x4b, 0xe4,
	0x8b, 0x1f, 0x57, 0xab, 0x77, 0xf6, 0x35, 0x76, 0xa2, 0xb6, 0x88, 0x85, 0x31, 0x23, 0x1d, 0xf8,
	0x73, 0x18, 0xdb, 0x9d, 0x7f, 0x81, 0x81, 0xbf, 0x81, 0x09, 0x9d, 0x7d, 0x01, 0x5a, 0x50, 0xda,
	0x8e, 0xf6, 0xfb, 0x7e, 0xde, 0xfb, 0xea, 0x7d, 0x1f, 0x1e, 0x1c, 0x31, 0xc9, 0xe1, 0x90, 0x8e,
	0x33, 0x35, 0xe5, 0x34, 0xe7, 0x09, 0xcd, 0x8f, 0xf4, 0x41, 0x46, 0x67, 0x11, 0xcb, 0x8a, 0x3d,
	0x46, 0x13, 0x36, 0xde, 0x9f, 0x16, 0x23, 0x0d, 0x93, 0x99, 0x18, 0x43, 0x50, 0x4c, 0x94, 0x51,
	0x64, 0xcd, 0x41, 0x81, 0x85, 0x82, 0x9c, 0x27, 0x81, 0x85, 0x82, 0x0a, 0x6a, 0xdf, 0x4f, 0x95,
	0x4a, 0x33, 0xa0, 0xac, 0x10, 0x94, 0x49, 0xa9, 0x0c, 0x33, 0x42, 0x49, 0xed, 0xe8, 0xf6, 0x83,
	0x73, 0x23, 0x67, 0x2c, 0x13, 0xdc, 0xd6, 0xab, 0xf2, 0xd3, 0x2b, 0x39, 0x72, 0x62, 0x7f, 0x03,
	0xdf, 0xdd, 0x06, 0x33, 0xb4, 0x5f, 0x31, 0x1c, 0x4c, 0x41, 0x1b, 0xf2, 0x10, 0x7b, 0x95, 0x6b,
	0xc1, 0x5b, 0xa8, 0x83, 0xba, 0xde, 0xb0, 0xf1, 0xe3, 0x34, 0x42, 0x71, 0xd3, 0x7d, 0xef, 0x70,
	0xff, 0x33, 0xc2, 0xe4, 0xb5, 0xd0, 0x15, 0xa8, 0x17, 0xe4, 0x63, 0xec, 0xbd, 0x57, 0x19, 0x87,
	0xc9, 0x1f, 0xf2, 0x76, 0x49, 0xce, 0xcf, 0xa2, 0xc6, 0xe6, 0xd6, 0x46, 0x18, 0x37, 0x5d, 0x79,
	0x87, 0x93, 0x47, 0xd8, 0x2b, 0x58, 0x0a, 0x23, 0x2d, 0x8e, 0xa1, 0x55, 0xeb, 0xa0, 0x6e, 0x7d,
	0x88, 0x7f, 0x9e, 0x46, 0x2b, 0x9b, 0x5b, 0x51, 0x18, 0x86, 0x71, 0xb3, 0x2c, 0xee, 0x8a, 0x63,
	0x20, 0x5d, 0x8c, 0xad, 0xd0, 0xa8, 0x7d, 0x90, 0xad, 0xba, 0x6d, 0xea, 0xcd, 0xcf, 0xa2, 0x1b,
	0x56, 0x19, 0xdb, 0x2e, 0x6f, 0xcb, 0x9a, 0x3f, 0x47, 0xf8, 0xde, 0x39, 0x53, 0xba, 0x50, 0x52,
	0x03, 0x79, 0x81, 0x6f, 0x3a, 0xe3, 0xba, 0x85, 0x3a, 0xf5, 0xee, 0xad, 0xfe, 0x7a, 0xb0, 0x7c,
	0xff, 0x41, 0xb5, 0x8f, 0x05, 0x46, 0x22, 0x7c, 0x47, 0xc2, 0xa1, 0x19, 0xfd, 0x65, 0xa4, 0x76,
	0xd1, 0xc8, 0x6a, 0xa9, 0x78, 0xb3, 0x30, 0xd3, 0xff, 0x5a, 0xc3, 0xab, 0xae, 0xcd, 0xae, 0x8b,
	0x9e, 0x9c, 0x20, 0x5c, 0xdf, 0x06, 0x43, 0xc2, 0xcb, 0xa6, 0x5f, 0x0c, 0xa4, 0x7d, 0x45, 0xbf,
	0x7e, 0xff, 0xd3, 0xb7, 0xef, 0x27, 0xb5, 0x67, 0xe4, 0x09, 0xcd, 0x99, 0x64, 0x29, 0xf0, 0xde,
	0xff, 0x92, 0xd7, 0xf4, 0xc3, 0xef, 0x78, 0x3f, 0x92, 0x2f, 0x08, 0x37, 0xca, 0xa5, 0x91, 0xfe,
	0x65, 0x43, 0xfe, 0xcd, 0xbb, 0x3d, 0xb8, 0x16, 0xe3, 0xe2, 0xf0, 0xd7, 0xad, 0xcb, 0x0e, 0x59,
	0x5b, 0xee, 0x72, 0xf8, 0xea, 0xdd, 0xcb, 0x54, 0x98, 0xbd, 0x69, 0x12, 0x8c, 0x55, 0x4e, 0xdd,
	0xa0, 0x9e, 0x3b, 0xea, 0x54, 0xf5, 0x52, 0x90, 0xf6, 0x82, 0xe9, 0xf2, 0x6b, 0x7f, 0x6e, 0x5f,
	0xc9, 0x8a, 0xd5, 0x0e, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x69, 0xee, 0xc3, 0x1c, 0xae, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BackupServiceClient is the client API for BackupService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BackupServiceClient interface {
	// Returns the specified MySQL backup.
	//
	// To get the list of available MySQL backups, make a [List] request.
	Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error)
	// Retrieves the list of MySQL backups available for the specified folder.
	List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error)
}

type backupServiceClient struct {
	cc *grpc.ClientConn
}

func NewBackupServiceClient(cc *grpc.ClientConn) BackupServiceClient {
	return &backupServiceClient{cc}
}

func (c *backupServiceClient) Get(ctx context.Context, in *GetBackupRequest, opts ...grpc.CallOption) (*Backup, error) {
	out := new(Backup)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.BackupService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *backupServiceClient) List(ctx context.Context, in *ListBackupsRequest, opts ...grpc.CallOption) (*ListBackupsResponse, error) {
	out := new(ListBackupsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.mdb.mysql.v1alpha.BackupService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BackupServiceServer is the server API for BackupService service.
type BackupServiceServer interface {
	// Returns the specified MySQL backup.
	//
	// To get the list of available MySQL backups, make a [List] request.
	Get(context.Context, *GetBackupRequest) (*Backup, error)
	// Retrieves the list of MySQL backups available for the specified folder.
	List(context.Context, *ListBackupsRequest) (*ListBackupsResponse, error)
}

// UnimplementedBackupServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBackupServiceServer struct {
}

func (*UnimplementedBackupServiceServer) Get(ctx context.Context, req *GetBackupRequest) (*Backup, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedBackupServiceServer) List(ctx context.Context, req *ListBackupsRequest) (*ListBackupsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterBackupServiceServer(s *grpc.Server, srv BackupServiceServer) {
	s.RegisterService(&_BackupService_serviceDesc, srv)
}

func _BackupService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.BackupService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).Get(ctx, req.(*GetBackupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BackupService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBackupsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BackupServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.mdb.mysql.v1alpha.BackupService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BackupServiceServer).List(ctx, req.(*ListBackupsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BackupService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.mdb.mysql.v1alpha.BackupService",
	HandlerType: (*BackupServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BackupService_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _BackupService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/mdb/mysql/v1alpha/backup_service.proto",
}
