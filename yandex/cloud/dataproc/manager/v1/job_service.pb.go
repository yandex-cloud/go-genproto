// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/dataproc/manager/v1/job_service.proto

package dataproc_manager

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type ListJobsRequest struct {
	// Required. ID of the cluster to list Data Proc jobs of.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListJobs requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListJobs
	// request to get the next page of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// String that describes a display filter.
	Filter               string   `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsRequest) Reset()         { *m = ListJobsRequest{} }
func (m *ListJobsRequest) String() string { return proto.CompactTextString(m) }
func (*ListJobsRequest) ProtoMessage()    {}
func (*ListJobsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe06b71d256c11f, []int{0}
}

func (m *ListJobsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsRequest.Unmarshal(m, b)
}
func (m *ListJobsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsRequest.Marshal(b, m, deterministic)
}
func (m *ListJobsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsRequest.Merge(m, src)
}
func (m *ListJobsRequest) XXX_Size() int {
	return xxx_messageInfo_ListJobsRequest.Size(m)
}
func (m *ListJobsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsRequest proto.InternalMessageInfo

func (m *ListJobsRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *ListJobsRequest) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListJobsRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

func (m *ListJobsRequest) GetFilter() string {
	if m != nil {
		return m.Filter
	}
	return ""
}

type ListJobsResponse struct {
	// Requested list of Data Proc jobs.
	Jobs []*Job `protobuf:"bytes,1,rep,name=jobs,proto3" json:"jobs,omitempty"`
	// This token allows you to get the next page of results for ListJobs requests,
	// if the number of results is larger than `page_size` specified in the request.
	// To get the next page, specify the value of `next_page_token` as a value for
	// the `page_token` parameter in the next ListClusters request. Subsequent ListClusters
	// requests will have their own `next_page_token` to continue paging through the results.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListJobsResponse) Reset()         { *m = ListJobsResponse{} }
func (m *ListJobsResponse) String() string { return proto.CompactTextString(m) }
func (*ListJobsResponse) ProtoMessage()    {}
func (*ListJobsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe06b71d256c11f, []int{1}
}

func (m *ListJobsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListJobsResponse.Unmarshal(m, b)
}
func (m *ListJobsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListJobsResponse.Marshal(b, m, deterministic)
}
func (m *ListJobsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListJobsResponse.Merge(m, src)
}
func (m *ListJobsResponse) XXX_Size() int {
	return xxx_messageInfo_ListJobsResponse.Size(m)
}
func (m *ListJobsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListJobsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListJobsResponse proto.InternalMessageInfo

func (m *ListJobsResponse) GetJobs() []*Job {
	if m != nil {
		return m.Jobs
	}
	return nil
}

func (m *ListJobsResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

type UpdateJobStatusRequest struct {
	// Required. ID of the Data Proc cluster.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. ID of the Data Proc job to update.
	JobId string `protobuf:"bytes,2,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Required. New status of the job.
	Status               Job_Status `protobuf:"varint,3,opt,name=status,proto3,enum=yandex.cloud.dataproc.manager.v1.Job_Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UpdateJobStatusRequest) Reset()         { *m = UpdateJobStatusRequest{} }
func (m *UpdateJobStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateJobStatusRequest) ProtoMessage()    {}
func (*UpdateJobStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe06b71d256c11f, []int{2}
}

func (m *UpdateJobStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJobStatusRequest.Unmarshal(m, b)
}
func (m *UpdateJobStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJobStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateJobStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJobStatusRequest.Merge(m, src)
}
func (m *UpdateJobStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateJobStatusRequest.Size(m)
}
func (m *UpdateJobStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJobStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJobStatusRequest proto.InternalMessageInfo

func (m *UpdateJobStatusRequest) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *UpdateJobStatusRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *UpdateJobStatusRequest) GetStatus() Job_Status {
	if m != nil {
		return m.Status
	}
	return Job_STATUS_UNSPECIFIED
}

type UpdateJobStatusResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateJobStatusResponse) Reset()         { *m = UpdateJobStatusResponse{} }
func (m *UpdateJobStatusResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateJobStatusResponse) ProtoMessage()    {}
func (*UpdateJobStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fbe06b71d256c11f, []int{3}
}

func (m *UpdateJobStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateJobStatusResponse.Unmarshal(m, b)
}
func (m *UpdateJobStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateJobStatusResponse.Marshal(b, m, deterministic)
}
func (m *UpdateJobStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateJobStatusResponse.Merge(m, src)
}
func (m *UpdateJobStatusResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateJobStatusResponse.Size(m)
}
func (m *UpdateJobStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateJobStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateJobStatusResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ListJobsRequest)(nil), "yandex.cloud.dataproc.manager.v1.ListJobsRequest")
	proto.RegisterType((*ListJobsResponse)(nil), "yandex.cloud.dataproc.manager.v1.ListJobsResponse")
	proto.RegisterType((*UpdateJobStatusRequest)(nil), "yandex.cloud.dataproc.manager.v1.UpdateJobStatusRequest")
	proto.RegisterType((*UpdateJobStatusResponse)(nil), "yandex.cloud.dataproc.manager.v1.UpdateJobStatusResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/dataproc/manager/v1/job_service.proto", fileDescriptor_fbe06b71d256c11f)
}

var fileDescriptor_fbe06b71d256c11f = []byte{
	// 480 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x71, 0x93, 0x5a, 0xf5, 0xf0, 0xa7, 0x68, 0x0f, 0x60, 0x22, 0x21, 0x22, 0x0b, 0x68,
	0x84, 0xe8, 0x3a, 0x36, 0x42, 0xa2, 0xa2, 0x3d, 0x10, 0x71, 0x69, 0xc4, 0x01, 0xb9, 0xf4, 0xc2,
	0xc5, 0x5a, 0x7b, 0x17, 0xb3, 0x21, 0xf5, 0x1a, 0xef, 0xda, 0x2a, 0x95, 0x38, 0x72, 0x09, 0x2f,
	0xc2, 0x95, 0xa7, 0x20, 0xcf, 0xc4, 0x09, 0x79, 0xd7, 0x81, 0x16, 0x2a, 0xb9, 0xe5, 0x3a, 0x33,
	0xbf, 0x6f, 0x3e, 0x7f, 0xde, 0x81, 0xf0, 0x13, 0xc9, 0x29, 0x3b, 0xf6, 0xd3, 0xb9, 0xa8, 0xa8,
	0x4f, 0x89, 0x22, 0x45, 0x29, 0x52, 0xff, 0x88, 0xe4, 0x24, 0x63, 0xa5, 0x5f, 0x07, 0xfe, 0x4c,
	0x24, 0xb1, 0x64, 0x65, 0xcd, 0x53, 0x86, 0x8b, 0x52, 0x28, 0x81, 0x86, 0x86, 0xc1, 0x9a, 0xc1,
	0x2b, 0x06, 0xb7, 0x0c, 0xae, 0x83, 0xc1, 0xdd, 0x33, 0xaa, 0x35, 0x99, 0x73, 0x4a, 0x14, 0x17,
	0xb9, 0x11, 0x18, 0x3c, 0xba, 0xc8, 0x52, 0x33, 0xeb, 0x7d, 0xb7, 0x60, 0xf3, 0x15, 0x97, 0x6a,
	0x2a, 0x12, 0x19, 0xb1, 0x8f, 0x15, 0x93, 0x0a, 0x6d, 0x01, 0xa4, 0xf3, 0x4a, 0x2a, 0x56, 0xc6,
	0x9c, 0xba, 0xd6, 0xd0, 0x1a, 0x39, 0x93, 0x8d, 0xc5, 0x32, 0xe8, 0xef, 0xee, 0x3d, 0x1d, 0x47,
	0x4e, 0xdb, 0xdb, 0xa7, 0x68, 0x0b, 0x9c, 0x82, 0x64, 0x2c, 0x96, 0xfc, 0x84, 0xb9, 0x6b, 0x43,
	0x6b, 0xd4, 0x9b, 0xc0, 0xcf, 0x1f, 0x81, 0xbd, 0xbb, 0x17, 0x8c, 0xc7, 0xe3, 0x68, 0xa3, 0x69,
	0x1e, 0xf0, 0x13, 0x86, 0x46, 0x00, 0x7a, 0x50, 0x89, 0x0f, 0x2c, 0x77, 0x7b, 0x5a, 0xd1, 0x59,
	0x2c, 0x83, 0x75, 0x3d, 0x19, 0x69, 0x95, 0x37, 0x4d, 0x0f, 0x79, 0x60, 0xbf, 0xe3, 0x73, 0xc5,
	0x4a, 0xb7, 0xaf, 0xa7, 0x60, 0xb1, 0xfc, 0xad, 0xd7, 0x76, 0xbc, 0x0a, 0x6e, 0xfe, 0xb1, 0x2c,
	0x0b, 0x91, 0x4b, 0x86, 0x76, 0xa0, 0x3f, 0x13, 0x89, 0x74, 0xad, 0x61, 0x6f, 0x74, 0x35, 0x7c,
	0x80, 0xbb, 0x32, 0xc4, 0x53, 0x91, 0x44, 0x1a, 0x41, 0x0f, 0x61, 0x33, 0x67, 0xc7, 0x2a, 0x3e,
	0xe5, 0xb0, 0xf9, 0x16, 0x27, 0xba, 0xde, 0x94, 0x5f, 0xaf, 0xac, 0x79, 0xdf, 0x2c, 0xb8, 0x75,
	0x58, 0x50, 0xa2, 0xd8, 0x54, 0x24, 0x07, 0x8a, 0xa8, 0xea, 0xf2, 0x89, 0xdd, 0x03, 0xbb, 0xf9,
	0xe1, 0x9c, 0x9a, 0x15, 0xa7, 0x86, 0xd6, 0x67, 0x22, 0xd9, 0xa7, 0xe8, 0x25, 0xd8, 0x52, 0x4b,
	0xeb, 0x94, 0x6e, 0x84, 0x8f, 0x2f, 0xf4, 0x25, 0xb8, 0xb5, 0xd3, 0xb2, 0xde, 0x1d, 0xb8, 0xfd,
	0x8f, 0x53, 0x13, 0x54, 0xf8, 0x75, 0x0d, 0xa0, 0xa9, 0x9a, 0x27, 0x87, 0x2a, 0x80, 0x26, 0xcb,
	0x17, 0xa9, 0xe2, 0x35, 0x43, 0x41, 0xf7, 0xb6, 0xbf, 0x1e, 0xcb, 0x20, 0xbc, 0x0c, 0x62, 0x3c,
	0x78, 0x57, 0xd0, 0x17, 0x0b, 0xae, 0x19, 0x87, 0xc6, 0x1e, 0x7a, 0xd6, 0x2d, 0x73, 0x7e, 0xf6,
	0x83, 0x9d, 0xff, 0x20, 0x57, 0x3e, 0x26, 0x9f, 0xe1, 0xfe, 0x19, 0x9a, 0x14, 0xfc, 0x3c, 0x85,
	0xb7, 0x87, 0x19, 0x57, 0xef, 0xab, 0x04, 0xa7, 0xe2, 0xc8, 0x37, 0xc0, 0xb6, 0xb9, 0xae, 0x4c,
	0x6c, 0x67, 0x2c, 0xd7, 0xb7, 0xe4, 0x77, 0x9d, 0xdd, 0xf3, 0x55, 0x2d, 0x6e, 0x6b, 0x89, 0xad,
	0xc1, 0x27, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x26, 0xee, 0x4e, 0x43, 0x27, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// JobServiceClient is the client API for JobService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobServiceClient interface {
	// Retrieves a list of jobs for Data Proc cluster.
	ListActive(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error)
	// Currently used to update Job status.
	UpdateStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error)
}

type jobServiceClient struct {
	cc *grpc.ClientConn
}

func NewJobServiceClient(cc *grpc.ClientConn) JobServiceClient {
	return &jobServiceClient{cc}
}

func (c *jobServiceClient) ListActive(ctx context.Context, in *ListJobsRequest, opts ...grpc.CallOption) (*ListJobsResponse, error) {
	out := new(ListJobsResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.manager.v1.JobService/ListActive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobServiceClient) UpdateStatus(ctx context.Context, in *UpdateJobStatusRequest, opts ...grpc.CallOption) (*UpdateJobStatusResponse, error) {
	out := new(UpdateJobStatusResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.manager.v1.JobService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobServiceServer is the server API for JobService service.
type JobServiceServer interface {
	// Retrieves a list of jobs for Data Proc cluster.
	ListActive(context.Context, *ListJobsRequest) (*ListJobsResponse, error)
	// Currently used to update Job status.
	UpdateStatus(context.Context, *UpdateJobStatusRequest) (*UpdateJobStatusResponse, error)
}

// UnimplementedJobServiceServer can be embedded to have forward compatible implementations.
type UnimplementedJobServiceServer struct {
}

func (*UnimplementedJobServiceServer) ListActive(ctx context.Context, req *ListJobsRequest) (*ListJobsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActive not implemented")
}
func (*UnimplementedJobServiceServer) UpdateStatus(ctx context.Context, req *UpdateJobStatusRequest) (*UpdateJobStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}

func RegisterJobServiceServer(s *grpc.Server, srv JobServiceServer) {
	s.RegisterService(&_JobService_serviceDesc, srv)
}

func _JobService_ListActive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).ListActive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.manager.v1.JobService/ListActive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).ListActive(ctx, req.(*ListJobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateJobStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.manager.v1.JobService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobServiceServer).UpdateStatus(ctx, req.(*UpdateJobStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _JobService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dataproc.manager.v1.JobService",
	HandlerType: (*JobServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListActive",
			Handler:    _JobService_ListActive_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _JobService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dataproc/manager/v1/job_service.proto",
}
