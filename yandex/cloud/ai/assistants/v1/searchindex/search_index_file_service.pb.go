// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/searchindex/search_index_file_service.proto

package searchindex

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
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

// Request message for creating multiple files within a search index.
type BatchCreateSearchIndexFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIds       []string `protobuf:"bytes,1,rep,name=file_ids,json=fileIds,proto3" json:"file_ids,omitempty"`
	SearchIndexId string   `protobuf:"bytes,2,opt,name=search_index_id,json=searchIndexId,proto3" json:"search_index_id,omitempty"`
}

func (x *BatchCreateSearchIndexFileRequest) Reset() {
	*x = BatchCreateSearchIndexFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateSearchIndexFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateSearchIndexFileRequest) ProtoMessage() {}

func (x *BatchCreateSearchIndexFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateSearchIndexFileRequest.ProtoReflect.Descriptor instead.
func (*BatchCreateSearchIndexFileRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP(), []int{0}
}

func (x *BatchCreateSearchIndexFileRequest) GetFileIds() []string {
	if x != nil {
		return x.FileIds
	}
	return nil
}

func (x *BatchCreateSearchIndexFileRequest) GetSearchIndexId() string {
	if x != nil {
		return x.SearchIndexId
	}
	return ""
}

// Response message for the BatchCreate operation.
type BatchCreateSearchIndexFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files []*SearchIndexFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *BatchCreateSearchIndexFileResponse) Reset() {
	*x = BatchCreateSearchIndexFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchCreateSearchIndexFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateSearchIndexFileResponse) ProtoMessage() {}

func (x *BatchCreateSearchIndexFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchCreateSearchIndexFileResponse.ProtoReflect.Descriptor instead.
func (*BatchCreateSearchIndexFileResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP(), []int{1}
}

func (x *BatchCreateSearchIndexFileResponse) GetFiles() []*SearchIndexFile {
	if x != nil {
		return x.Files
	}
	return nil
}

// Request message for retrieving a file from a search index.
type GetSearchIndexFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the file to retrieve.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// ID of the search index that contains the file.
	SearchIndexId string `protobuf:"bytes,2,opt,name=search_index_id,json=searchIndexId,proto3" json:"search_index_id,omitempty"`
}

func (x *GetSearchIndexFileRequest) Reset() {
	*x = GetSearchIndexFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSearchIndexFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSearchIndexFileRequest) ProtoMessage() {}

func (x *GetSearchIndexFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSearchIndexFileRequest.ProtoReflect.Descriptor instead.
func (*GetSearchIndexFileRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetSearchIndexFileRequest) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *GetSearchIndexFileRequest) GetSearchIndexId() string {
	if x != nil {
		return x.SearchIndexId
	}
	return ""
}

// Request message for listing files in a specific search index.
type ListSearchIndexFilesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the search index whose files will be listed.
	SearchIndexId string `protobuf:"bytes,1,opt,name=search_index_id,json=searchIndexId,proto3" json:"search_index_id,omitempty"`
	// Maximum number of files to return per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to retrieve the next page of results.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListSearchIndexFilesRequest) Reset() {
	*x = ListSearchIndexFilesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSearchIndexFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSearchIndexFilesRequest) ProtoMessage() {}

func (x *ListSearchIndexFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSearchIndexFilesRequest.ProtoReflect.Descriptor instead.
func (*ListSearchIndexFilesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListSearchIndexFilesRequest) GetSearchIndexId() string {
	if x != nil {
		return x.SearchIndexId
	}
	return ""
}

func (x *ListSearchIndexFilesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListSearchIndexFilesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for the list operation.
type ListSearchIndexFilesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of files in the specified search index.
	Files []*SearchIndexFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// Token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListSearchIndexFilesResponse) Reset() {
	*x = ListSearchIndexFilesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSearchIndexFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSearchIndexFilesResponse) ProtoMessage() {}

func (x *ListSearchIndexFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSearchIndexFilesResponse.ProtoReflect.Descriptor instead.
func (*ListSearchIndexFilesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP(), []int{4}
}

func (x *ListSearchIndexFilesResponse) GetFiles() []*SearchIndexFile {
	if x != nil {
		return x.Files
	}
	return nil
}

func (x *ListSearchIndexFilesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x29, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x1a, 0x41, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x21, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42, 0x06, 0x82, 0xc8, 0x31, 0x02, 0x3e, 0x30, 0x52,
	0x07, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x22, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x68,
	0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7,
	0x31, 0x01, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x22, 0x87, 0x01, 0x0a, 0x1b, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x98, 0x01, 0x0a, 0x1c, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x9a, 0x05,
	0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdd, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0xb2, 0xd2, 0x2a, 0x24, 0x12,
	0x22, 0x42, 0x61, 0x74, 0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x3a, 0x01, 0x2a, 0x22, 0x2a, 0x2f, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x3a, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0xcb, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x44, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69,
	0x6c, 0x65, 0x22, 0x42, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3c, 0x12, 0x3a, 0x2f, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x7b, 0x66, 0x69,
	0x6c, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xd1, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x46, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x47, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x38, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x32, 0x12, 0x30, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x46, 0x69, 0x6c, 0x65, 0x2f, 0x7b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x8a, 0x01, 0x0a, 0x2d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5a, 0x59, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x3b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescData = file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc
)

func file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescData)
	})
	return file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_goTypes = []any{
	(*BatchCreateSearchIndexFileRequest)(nil),  // 0: yandex.cloud.ai.assistants.v1.searchindex.BatchCreateSearchIndexFileRequest
	(*BatchCreateSearchIndexFileResponse)(nil), // 1: yandex.cloud.ai.assistants.v1.searchindex.BatchCreateSearchIndexFileResponse
	(*GetSearchIndexFileRequest)(nil),          // 2: yandex.cloud.ai.assistants.v1.searchindex.GetSearchIndexFileRequest
	(*ListSearchIndexFilesRequest)(nil),        // 3: yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesRequest
	(*ListSearchIndexFilesResponse)(nil),       // 4: yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesResponse
	(*SearchIndexFile)(nil),                    // 5: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFile
	(*operation.Operation)(nil),                // 6: yandex.cloud.operation.Operation
}
var file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.ai.assistants.v1.searchindex.BatchCreateSearchIndexFileResponse.files:type_name -> yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFile
	5, // 1: yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesResponse.files:type_name -> yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFile
	0, // 2: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileService.BatchCreate:input_type -> yandex.cloud.ai.assistants.v1.searchindex.BatchCreateSearchIndexFileRequest
	2, // 3: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileService.Get:input_type -> yandex.cloud.ai.assistants.v1.searchindex.GetSearchIndexFileRequest
	3, // 4: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileService.List:input_type -> yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesRequest
	6, // 5: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileService.BatchCreate:output_type -> yandex.cloud.operation.Operation
	5, // 6: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileService.Get:output_type -> yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFile
	4, // 7: yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileService.List:output_type -> yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_init() }
func file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*BatchCreateSearchIndexFileRequest); i {
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
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*BatchCreateSearchIndexFileResponse); i {
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
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetSearchIndexFileRequest); i {
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
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListSearchIndexFilesRequest); i {
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
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListSearchIndexFilesResponse); i {
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
			RawDescriptor: file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto = out.File
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc = nil
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_depIdxs = nil
}
