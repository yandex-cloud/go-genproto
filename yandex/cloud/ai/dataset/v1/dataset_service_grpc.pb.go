// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: yandex/cloud/ai/dataset/v1/dataset_service.proto

package fomo

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
	DatasetService_Describe_FullMethodName                   = "/yandex.cloud.ai.dataset.v1.DatasetService/Describe"
	DatasetService_Validate_FullMethodName                   = "/yandex.cloud.ai.dataset.v1.DatasetService/Validate"
	DatasetService_Create_FullMethodName                     = "/yandex.cloud.ai.dataset.v1.DatasetService/Create"
	DatasetService_Update_FullMethodName                     = "/yandex.cloud.ai.dataset.v1.DatasetService/Update"
	DatasetService_Delete_FullMethodName                     = "/yandex.cloud.ai.dataset.v1.DatasetService/Delete"
	DatasetService_List_FullMethodName                       = "/yandex.cloud.ai.dataset.v1.DatasetService/List"
	DatasetService_ListUploadFormats_FullMethodName          = "/yandex.cloud.ai.dataset.v1.DatasetService/ListUploadFormats"
	DatasetService_ListUploadSchemas_FullMethodName          = "/yandex.cloud.ai.dataset.v1.DatasetService/ListUploadSchemas"
	DatasetService_GetUploadDraftUrl_FullMethodName          = "/yandex.cloud.ai.dataset.v1.DatasetService/GetUploadDraftUrl"
	DatasetService_GetDownloadUrls_FullMethodName            = "/yandex.cloud.ai.dataset.v1.DatasetService/GetDownloadUrls"
	DatasetService_StartMultipartUploadDraft_FullMethodName  = "/yandex.cloud.ai.dataset.v1.DatasetService/StartMultipartUploadDraft"
	DatasetService_FinishMultipartUploadDraft_FullMethodName = "/yandex.cloud.ai.dataset.v1.DatasetService/FinishMultipartUploadDraft"
	DatasetService_ListTypes_FullMethodName                  = "/yandex.cloud.ai.dataset.v1.DatasetService/ListTypes"
	DatasetService_GetPreview_FullMethodName                 = "/yandex.cloud.ai.dataset.v1.DatasetService/GetPreview"
	DatasetService_ListOperationsIds_FullMethodName          = "/yandex.cloud.ai.dataset.v1.DatasetService/ListOperationsIds"
)

// DatasetServiceClient is the client API for DatasetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// A set of methods for managing datasets.
type DatasetServiceClient interface {
	// Returns dataset information by dataset id.
	Describe(ctx context.Context, in *DescribeDatasetRequest, opts ...grpc.CallOption) (*DescribeDatasetResponse, error)
	// Starts dataset validation process.
	Validate(ctx context.Context, in *ValidateDatasetRequest, opts ...grpc.CallOption) (*operation.Operation, error)
	// Creates dataset.
	Create(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error)
	// Updates dataset.
	Update(ctx context.Context, in *UpdateDatasetRequest, opts ...grpc.CallOption) (*UpdateDatasetResponse, error)
	// Deletes dataset.
	Delete(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*DeleteDatasetResponse, error)
	// Lists datasets in specified folder.
	List(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error)
	// Deprecated: Do not use.
	// Deprecated. Use ListUploadSchemas.
	ListUploadFormats(ctx context.Context, in *ListUploadFormatsRequest, opts ...grpc.CallOption) (*ListUploadFormatsResponse, error)
	// Lists supported dataset upload formats types and schemas for the specified dataset task type.
	ListUploadSchemas(ctx context.Context, in *ListUploadSchemasRequest, opts ...grpc.CallOption) (*ListUploadSchemasResponse, error)
	// Returns an S3 presigned URL for dataset upload.
	// This method only applicable if the dataset size does not exceed 5GB.
	GetUploadDraftUrl(ctx context.Context, in *GetUploadDraftUrlRequest, opts ...grpc.CallOption) (*GetUploadDraftUrlResponse, error)
	// Get urls to download dataset
	GetDownloadUrls(ctx context.Context, in *GetDownloadUrlsRequest, opts ...grpc.CallOption) (*GetDownloadUrlsResponse, error)
	// Returns a list of S3 presigned URLs for multipart upload of dataset.
	StartMultipartUploadDraft(ctx context.Context, in *StartMultipartUploadDraftRequest, opts ...grpc.CallOption) (*StartMultipartUploadDraftResponse, error)
	// Finishes multipart upload of the dataset.
	FinishMultipartUploadDraft(ctx context.Context, in *FinishMultipartUploadDraftRequest, opts ...grpc.CallOption) (*FinishMultipartUploadDraftResponse, error)
	// Returns a list of dataset types
	ListTypes(ctx context.Context, in *ListTypesRequest, opts ...grpc.CallOption) (*ListTypesResponse, error)
	// Returns a preview of dataset types
	GetPreview(ctx context.Context, in *GetDatasetPreviewRequest, opts ...grpc.CallOption) (*GetDatasetPreviewResponse, error)
	ListOperationsIds(ctx context.Context, in *ListOperationsIdsRequest, opts ...grpc.CallOption) (*ListOperationsIdsResponse, error)
}

type datasetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatasetServiceClient(cc grpc.ClientConnInterface) DatasetServiceClient {
	return &datasetServiceClient{cc}
}

func (c *datasetServiceClient) Describe(ctx context.Context, in *DescribeDatasetRequest, opts ...grpc.CallOption) (*DescribeDatasetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeDatasetResponse)
	err := c.cc.Invoke(ctx, DatasetService_Describe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) Validate(ctx context.Context, in *ValidateDatasetRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(operation.Operation)
	err := c.cc.Invoke(ctx, DatasetService_Validate_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) Create(ctx context.Context, in *CreateDatasetRequest, opts ...grpc.CallOption) (*CreateDatasetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDatasetResponse)
	err := c.cc.Invoke(ctx, DatasetService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) Update(ctx context.Context, in *UpdateDatasetRequest, opts ...grpc.CallOption) (*UpdateDatasetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDatasetResponse)
	err := c.cc.Invoke(ctx, DatasetService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) Delete(ctx context.Context, in *DeleteDatasetRequest, opts ...grpc.CallOption) (*DeleteDatasetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDatasetResponse)
	err := c.cc.Invoke(ctx, DatasetService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) List(ctx context.Context, in *ListDatasetsRequest, opts ...grpc.CallOption) (*ListDatasetsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDatasetsResponse)
	err := c.cc.Invoke(ctx, DatasetService_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Deprecated: Do not use.
func (c *datasetServiceClient) ListUploadFormats(ctx context.Context, in *ListUploadFormatsRequest, opts ...grpc.CallOption) (*ListUploadFormatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUploadFormatsResponse)
	err := c.cc.Invoke(ctx, DatasetService_ListUploadFormats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) ListUploadSchemas(ctx context.Context, in *ListUploadSchemasRequest, opts ...grpc.CallOption) (*ListUploadSchemasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListUploadSchemasResponse)
	err := c.cc.Invoke(ctx, DatasetService_ListUploadSchemas_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetUploadDraftUrl(ctx context.Context, in *GetUploadDraftUrlRequest, opts ...grpc.CallOption) (*GetUploadDraftUrlResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetUploadDraftUrlResponse)
	err := c.cc.Invoke(ctx, DatasetService_GetUploadDraftUrl_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetDownloadUrls(ctx context.Context, in *GetDownloadUrlsRequest, opts ...grpc.CallOption) (*GetDownloadUrlsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDownloadUrlsResponse)
	err := c.cc.Invoke(ctx, DatasetService_GetDownloadUrls_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) StartMultipartUploadDraft(ctx context.Context, in *StartMultipartUploadDraftRequest, opts ...grpc.CallOption) (*StartMultipartUploadDraftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartMultipartUploadDraftResponse)
	err := c.cc.Invoke(ctx, DatasetService_StartMultipartUploadDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) FinishMultipartUploadDraft(ctx context.Context, in *FinishMultipartUploadDraftRequest, opts ...grpc.CallOption) (*FinishMultipartUploadDraftResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FinishMultipartUploadDraftResponse)
	err := c.cc.Invoke(ctx, DatasetService_FinishMultipartUploadDraft_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) ListTypes(ctx context.Context, in *ListTypesRequest, opts ...grpc.CallOption) (*ListTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListTypesResponse)
	err := c.cc.Invoke(ctx, DatasetService_ListTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) GetPreview(ctx context.Context, in *GetDatasetPreviewRequest, opts ...grpc.CallOption) (*GetDatasetPreviewResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatasetPreviewResponse)
	err := c.cc.Invoke(ctx, DatasetService_GetPreview_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *datasetServiceClient) ListOperationsIds(ctx context.Context, in *ListOperationsIdsRequest, opts ...grpc.CallOption) (*ListOperationsIdsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListOperationsIdsResponse)
	err := c.cc.Invoke(ctx, DatasetService_ListOperationsIds_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatasetServiceServer is the server API for DatasetService service.
// All implementations should embed UnimplementedDatasetServiceServer
// for forward compatibility.
//
// A set of methods for managing datasets.
type DatasetServiceServer interface {
	// Returns dataset information by dataset id.
	Describe(context.Context, *DescribeDatasetRequest) (*DescribeDatasetResponse, error)
	// Starts dataset validation process.
	Validate(context.Context, *ValidateDatasetRequest) (*operation.Operation, error)
	// Creates dataset.
	Create(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error)
	// Updates dataset.
	Update(context.Context, *UpdateDatasetRequest) (*UpdateDatasetResponse, error)
	// Deletes dataset.
	Delete(context.Context, *DeleteDatasetRequest) (*DeleteDatasetResponse, error)
	// Lists datasets in specified folder.
	List(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error)
	// Deprecated: Do not use.
	// Deprecated. Use ListUploadSchemas.
	ListUploadFormats(context.Context, *ListUploadFormatsRequest) (*ListUploadFormatsResponse, error)
	// Lists supported dataset upload formats types and schemas for the specified dataset task type.
	ListUploadSchemas(context.Context, *ListUploadSchemasRequest) (*ListUploadSchemasResponse, error)
	// Returns an S3 presigned URL for dataset upload.
	// This method only applicable if the dataset size does not exceed 5GB.
	GetUploadDraftUrl(context.Context, *GetUploadDraftUrlRequest) (*GetUploadDraftUrlResponse, error)
	// Get urls to download dataset
	GetDownloadUrls(context.Context, *GetDownloadUrlsRequest) (*GetDownloadUrlsResponse, error)
	// Returns a list of S3 presigned URLs for multipart upload of dataset.
	StartMultipartUploadDraft(context.Context, *StartMultipartUploadDraftRequest) (*StartMultipartUploadDraftResponse, error)
	// Finishes multipart upload of the dataset.
	FinishMultipartUploadDraft(context.Context, *FinishMultipartUploadDraftRequest) (*FinishMultipartUploadDraftResponse, error)
	// Returns a list of dataset types
	ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error)
	// Returns a preview of dataset types
	GetPreview(context.Context, *GetDatasetPreviewRequest) (*GetDatasetPreviewResponse, error)
	ListOperationsIds(context.Context, *ListOperationsIdsRequest) (*ListOperationsIdsResponse, error)
}

// UnimplementedDatasetServiceServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatasetServiceServer struct{}

func (UnimplementedDatasetServiceServer) Describe(context.Context, *DescribeDatasetRequest) (*DescribeDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Describe not implemented")
}
func (UnimplementedDatasetServiceServer) Validate(context.Context, *ValidateDatasetRequest) (*operation.Operation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}
func (UnimplementedDatasetServiceServer) Create(context.Context, *CreateDatasetRequest) (*CreateDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDatasetServiceServer) Update(context.Context, *UpdateDatasetRequest) (*UpdateDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDatasetServiceServer) Delete(context.Context, *DeleteDatasetRequest) (*DeleteDatasetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDatasetServiceServer) List(context.Context, *ListDatasetsRequest) (*ListDatasetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDatasetServiceServer) ListUploadFormats(context.Context, *ListUploadFormatsRequest) (*ListUploadFormatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUploadFormats not implemented")
}
func (UnimplementedDatasetServiceServer) ListUploadSchemas(context.Context, *ListUploadSchemasRequest) (*ListUploadSchemasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUploadSchemas not implemented")
}
func (UnimplementedDatasetServiceServer) GetUploadDraftUrl(context.Context, *GetUploadDraftUrlRequest) (*GetUploadDraftUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUploadDraftUrl not implemented")
}
func (UnimplementedDatasetServiceServer) GetDownloadUrls(context.Context, *GetDownloadUrlsRequest) (*GetDownloadUrlsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDownloadUrls not implemented")
}
func (UnimplementedDatasetServiceServer) StartMultipartUploadDraft(context.Context, *StartMultipartUploadDraftRequest) (*StartMultipartUploadDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartMultipartUploadDraft not implemented")
}
func (UnimplementedDatasetServiceServer) FinishMultipartUploadDraft(context.Context, *FinishMultipartUploadDraftRequest) (*FinishMultipartUploadDraftResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinishMultipartUploadDraft not implemented")
}
func (UnimplementedDatasetServiceServer) ListTypes(context.Context, *ListTypesRequest) (*ListTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTypes not implemented")
}
func (UnimplementedDatasetServiceServer) GetPreview(context.Context, *GetDatasetPreviewRequest) (*GetDatasetPreviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPreview not implemented")
}
func (UnimplementedDatasetServiceServer) ListOperationsIds(context.Context, *ListOperationsIdsRequest) (*ListOperationsIdsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperationsIds not implemented")
}
func (UnimplementedDatasetServiceServer) testEmbeddedByValue() {}

// UnsafeDatasetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatasetServiceServer will
// result in compilation errors.
type UnsafeDatasetServiceServer interface {
	mustEmbedUnimplementedDatasetServiceServer()
}

func RegisterDatasetServiceServer(s grpc.ServiceRegistrar, srv DatasetServiceServer) {
	// If the following call pancis, it indicates UnimplementedDatasetServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DatasetService_ServiceDesc, srv)
}

func _DatasetService_Describe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).Describe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_Describe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).Describe(ctx, req.(*DescribeDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_Validate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).Validate(ctx, req.(*ValidateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).Create(ctx, req.(*CreateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).Update(ctx, req.(*UpdateDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).Delete(ctx, req.(*DeleteDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatasetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).List(ctx, req.(*ListDatasetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_ListUploadFormats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUploadFormatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).ListUploadFormats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_ListUploadFormats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).ListUploadFormats(ctx, req.(*ListUploadFormatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_ListUploadSchemas_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUploadSchemasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).ListUploadSchemas(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_ListUploadSchemas_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).ListUploadSchemas(ctx, req.(*ListUploadSchemasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetUploadDraftUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUploadDraftUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetUploadDraftUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_GetUploadDraftUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetUploadDraftUrl(ctx, req.(*GetUploadDraftUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetDownloadUrls_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDownloadUrlsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetDownloadUrls(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_GetDownloadUrls_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetDownloadUrls(ctx, req.(*GetDownloadUrlsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_StartMultipartUploadDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartMultipartUploadDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).StartMultipartUploadDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_StartMultipartUploadDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).StartMultipartUploadDraft(ctx, req.(*StartMultipartUploadDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_FinishMultipartUploadDraft_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishMultipartUploadDraftRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).FinishMultipartUploadDraft(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_FinishMultipartUploadDraft_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).FinishMultipartUploadDraft(ctx, req.(*FinishMultipartUploadDraftRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_ListTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).ListTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_ListTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).ListTypes(ctx, req.(*ListTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_GetPreview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatasetPreviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).GetPreview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_GetPreview_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).GetPreview(ctx, req.(*GetDatasetPreviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatasetService_ListOperationsIds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationsIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatasetServiceServer).ListOperationsIds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatasetService_ListOperationsIds_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatasetServiceServer).ListOperationsIds(ctx, req.(*ListOperationsIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatasetService_ServiceDesc is the grpc.ServiceDesc for DatasetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatasetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.dataset.v1.DatasetService",
	HandlerType: (*DatasetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Describe",
			Handler:    _DatasetService_Describe_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _DatasetService_Validate_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _DatasetService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DatasetService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DatasetService_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DatasetService_List_Handler,
		},
		{
			MethodName: "ListUploadFormats",
			Handler:    _DatasetService_ListUploadFormats_Handler,
		},
		{
			MethodName: "ListUploadSchemas",
			Handler:    _DatasetService_ListUploadSchemas_Handler,
		},
		{
			MethodName: "GetUploadDraftUrl",
			Handler:    _DatasetService_GetUploadDraftUrl_Handler,
		},
		{
			MethodName: "GetDownloadUrls",
			Handler:    _DatasetService_GetDownloadUrls_Handler,
		},
		{
			MethodName: "StartMultipartUploadDraft",
			Handler:    _DatasetService_StartMultipartUploadDraft_Handler,
		},
		{
			MethodName: "FinishMultipartUploadDraft",
			Handler:    _DatasetService_FinishMultipartUploadDraft_Handler,
		},
		{
			MethodName: "ListTypes",
			Handler:    _DatasetService_ListTypes_Handler,
		},
		{
			MethodName: "GetPreview",
			Handler:    _DatasetService_GetPreview_Handler,
		},
		{
			MethodName: "ListOperationsIds",
			Handler:    _DatasetService_ListOperationsIds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/dataset/v1/dataset_service.proto",
}
