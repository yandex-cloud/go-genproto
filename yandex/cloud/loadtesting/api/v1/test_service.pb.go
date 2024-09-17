// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/api/v1/test_service.proto

package loadtesting

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	test "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/test"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder to create a test in.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Test configuration associated with agents on which they will be executed.
	// In case of multiple configurations, a multitest will be created.
	Configurations []*test.SingleAgentConfiguration `protobuf:"bytes,2,rep,name=configurations,proto3" json:"configurations,omitempty"`
	// Test details. Name, tags etc.
	TestDetails *test.Details `protobuf:"bytes,3,opt,name=test_details,json=testDetails,proto3" json:"test_details,omitempty"`
}

func (x *CreateTestRequest) Reset() {
	*x = CreateTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestRequest) ProtoMessage() {}

func (x *CreateTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestRequest.ProtoReflect.Descriptor instead.
func (*CreateTestRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTestRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateTestRequest) GetConfigurations() []*test.SingleAgentConfiguration {
	if x != nil {
		return x.Configurations
	}
	return nil
}

func (x *CreateTestRequest) GetTestDetails() *test.Details {
	if x != nil {
		return x.TestDetails
	}
	return nil
}

type CreateTestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the test that is being created.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *CreateTestMetadata) Reset() {
	*x = CreateTestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTestMetadata) ProtoMessage() {}

func (x *CreateTestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTestMetadata.ProtoReflect.Descriptor instead.
func (*CreateTestMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTestMetadata) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type GetTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the test to return.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *GetTestRequest) Reset() {
	*x = GetTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTestRequest) ProtoMessage() {}

func (x *GetTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTestRequest.ProtoReflect.Descriptor instead.
func (*GetTestRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTestRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type StopTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the test to stop.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *StopTestRequest) Reset() {
	*x = StopTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTestRequest) ProtoMessage() {}

func (x *StopTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTestRequest.ProtoReflect.Descriptor instead.
func (*StopTestRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{3}
}

func (x *StopTestRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type StopTestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the test that is being stopped.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *StopTestMetadata) Reset() {
	*x = StopTestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopTestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopTestMetadata) ProtoMessage() {}

func (x *StopTestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopTestMetadata.ProtoReflect.Descriptor instead.
func (*StopTestMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{4}
}

func (x *StopTestMetadata) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type DeleteTestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the test to delete.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *DeleteTestRequest) Reset() {
	*x = DeleteTestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestRequest) ProtoMessage() {}

func (x *DeleteTestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestRequest.ProtoReflect.Descriptor instead.
func (*DeleteTestRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteTestRequest) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type DeleteTestMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the test that is being deleted.
	TestId string `protobuf:"bytes,1,opt,name=test_id,json=testId,proto3" json:"test_id,omitempty"`
}

func (x *DeleteTestMetadata) Reset() {
	*x = DeleteTestMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTestMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTestMetadata) ProtoMessage() {}

func (x *DeleteTestMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTestMetadata.ProtoReflect.Descriptor instead.
func (*DeleteTestMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTestMetadata) GetTestId() string {
	if x != nil {
		return x.TestId
	}
	return ""
}

type ListTestsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder to list tests in.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than `page_size`, the service returns a [ListTestsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListTestsResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters tests listed in the response.
	//
	// The filter expression may contain multiple field expressions joined by `AND`.
	// The field expression must specify:
	// 1. The field name.
	// 2. An operator:
	//   - `=`, `!=`, `<`, `<=`, `>`, `>=`, `CONTAINS`, `:` for single values.
	//   - `IN` or `NOT IN` for lists of values.
	//
	// 3. The value. String values must be encosed in `"`, boolean values are {`true`, `false`}, timestamp values in ISO-8601.
	//
	// Currently supported fields:
	// - `id` [yandex.cloud.loadtesting.api.v1.test.Test.id]
	//   - operators: `=`, `!=`, `IN`, `NOT IN`
	//
	// - `details.name` [yandex.cloud.loadtesting.api.v1.test.Details.name]
	//   - operators: `=`, `!=`, `IN`, `NOT IN`, `CONTAINS`
	//
	// - `details.tags.<TAG_NAME>` [yandex.cloud.loadtesting.api.v1.test.Details.tags]
	//   - operators: `:`
	//
	// - `summary.status` [yandex.cloud.loadtesting.api.v1.test.Summary.status]
	//   - operators: `=`, `!=`, `IN`, `NOT IN`
	//
	// - `summary.is_finished` [yandex.cloud.loadtesting.api.v1.test.Summary.is_finished]
	//   - operators: `=`
	//
	// - `summary.created_at` [yandex.cloud.loadtesting.api.v1.test.Summary.created_at]
	//   - operators: `<`, `<=`, `>`, `>=`
	//
	// - `summary.created_by` [yandex.cloud.loadtesting.api.v1.test.Summary.created_by]
	//   - operators: `=`, `!=`, `IN`, `NOT IN`
	//
	// Examples:
	// - `summary.status IN ("DONE", "ERROR") AND details.tags.author:"yandex"`
	// - `summary.is_finished = true AND details.name CONTAINS "nightly-test"`
	Filter string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
}

func (x *ListTestsRequest) Reset() {
	*x = ListTestsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestsRequest) ProtoMessage() {}

func (x *ListTestsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestsRequest.ProtoReflect.Descriptor instead.
func (*ListTestsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListTestsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListTestsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListTestsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListTestsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListTestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of tests in the specified folder.
	Tests []*test.Test `protobuf:"bytes,1,rep,name=tests,proto3" json:"tests,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// the specified [ListTestsRequest.page_size], use `next_page_token` as the value
	// for the [ListTestsRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListTestsResponse) Reset() {
	*x = ListTestsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTestsResponse) ProtoMessage() {}

func (x *ListTestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTestsResponse.ProtoReflect.Descriptor instead.
func (*ListTestsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP(), []int{8}
}

func (x *ListTestsResponse) GetTests() []*test.Test {
	if x != nil {
		return x.Tests
	}
	return nil
}

func (x *ListTestsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_loadtesting_api_v1_test_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x45, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x66, 0x0a, 0x0e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x41,
	0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x50, 0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x2f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x06, 0x74, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22,
	0x2b, 0x0a, 0x10, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2c, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x17, 0x0a, 0x07, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x74, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x83, 0x01, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x22,
	0x7d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xcf,
	0x06, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa8,
	0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x47, 0xb2, 0xd2, 0x2a, 0x1f, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x3a, 0x01, 0x2a, 0x22, 0x19,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x8f, 0x01, 0x0a, 0x03, 0x47, 0x65,
	0x74, 0x12, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x12, 0x23, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x73, 0x2f, 0x7b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb1, 0x01, 0x0a, 0x04,
	0x53, 0x74, 0x6f, 0x70, 0x12, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0xb2, 0xd2, 0x2a, 0x1d, 0x0a,
	0x10, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x09, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x2d, 0x3a, 0x01, 0x2a, 0x22, 0x28, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x2f, 0x7b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x73, 0x74, 0x6f, 0x70, 0x12,
	0xbb, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x32, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x5a, 0xb2, 0xd2, 0x2a, 0x2b, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x25, 0x2a, 0x23, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2f, 0x7b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x90, 0x01,
	0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x54, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x12, 0x19, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x73,
	0x42, 0x76, 0x0a, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x61,
	0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescData = file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDesc
)

func file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_loadtesting_api_v1_test_service_proto_goTypes = []any{
	(*CreateTestRequest)(nil),             // 0: yandex.cloud.loadtesting.api.v1.CreateTestRequest
	(*CreateTestMetadata)(nil),            // 1: yandex.cloud.loadtesting.api.v1.CreateTestMetadata
	(*GetTestRequest)(nil),                // 2: yandex.cloud.loadtesting.api.v1.GetTestRequest
	(*StopTestRequest)(nil),               // 3: yandex.cloud.loadtesting.api.v1.StopTestRequest
	(*StopTestMetadata)(nil),              // 4: yandex.cloud.loadtesting.api.v1.StopTestMetadata
	(*DeleteTestRequest)(nil),             // 5: yandex.cloud.loadtesting.api.v1.DeleteTestRequest
	(*DeleteTestMetadata)(nil),            // 6: yandex.cloud.loadtesting.api.v1.DeleteTestMetadata
	(*ListTestsRequest)(nil),              // 7: yandex.cloud.loadtesting.api.v1.ListTestsRequest
	(*ListTestsResponse)(nil),             // 8: yandex.cloud.loadtesting.api.v1.ListTestsResponse
	(*test.SingleAgentConfiguration)(nil), // 9: yandex.cloud.loadtesting.api.v1.test.SingleAgentConfiguration
	(*test.Details)(nil),                  // 10: yandex.cloud.loadtesting.api.v1.test.Details
	(*test.Test)(nil),                     // 11: yandex.cloud.loadtesting.api.v1.test.Test
	(*operation.Operation)(nil),           // 12: yandex.cloud.operation.Operation
}
var file_yandex_cloud_loadtesting_api_v1_test_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.loadtesting.api.v1.CreateTestRequest.configurations:type_name -> yandex.cloud.loadtesting.api.v1.test.SingleAgentConfiguration
	10, // 1: yandex.cloud.loadtesting.api.v1.CreateTestRequest.test_details:type_name -> yandex.cloud.loadtesting.api.v1.test.Details
	11, // 2: yandex.cloud.loadtesting.api.v1.ListTestsResponse.tests:type_name -> yandex.cloud.loadtesting.api.v1.test.Test
	0,  // 3: yandex.cloud.loadtesting.api.v1.TestService.Create:input_type -> yandex.cloud.loadtesting.api.v1.CreateTestRequest
	2,  // 4: yandex.cloud.loadtesting.api.v1.TestService.Get:input_type -> yandex.cloud.loadtesting.api.v1.GetTestRequest
	3,  // 5: yandex.cloud.loadtesting.api.v1.TestService.Stop:input_type -> yandex.cloud.loadtesting.api.v1.StopTestRequest
	5,  // 6: yandex.cloud.loadtesting.api.v1.TestService.Delete:input_type -> yandex.cloud.loadtesting.api.v1.DeleteTestRequest
	7,  // 7: yandex.cloud.loadtesting.api.v1.TestService.List:input_type -> yandex.cloud.loadtesting.api.v1.ListTestsRequest
	12, // 8: yandex.cloud.loadtesting.api.v1.TestService.Create:output_type -> yandex.cloud.operation.Operation
	11, // 9: yandex.cloud.loadtesting.api.v1.TestService.Get:output_type -> yandex.cloud.loadtesting.api.v1.test.Test
	12, // 10: yandex.cloud.loadtesting.api.v1.TestService.Stop:output_type -> yandex.cloud.operation.Operation
	12, // 11: yandex.cloud.loadtesting.api.v1.TestService.Delete:output_type -> yandex.cloud.operation.Operation
	8,  // 12: yandex.cloud.loadtesting.api.v1.TestService.List:output_type -> yandex.cloud.loadtesting.api.v1.ListTestsResponse
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_test_service_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_test_service_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_test_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*StopTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*StopTestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTestRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTestMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListTestsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*ListTestsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_test_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_test_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_api_v1_test_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_test_service_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_test_service_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_api_v1_test_service_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_test_service_proto_depIdxs = nil
}
