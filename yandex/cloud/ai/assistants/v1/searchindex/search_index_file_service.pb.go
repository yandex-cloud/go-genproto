// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Request message for creating multiple files within a search index.
type BatchCreateSearchIndexFileRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	FileIds       []string               `protobuf:"bytes,1,rep,name=file_ids,json=fileIds,proto3" json:"file_ids,omitempty"`
	SearchIndexId string                 `protobuf:"bytes,2,opt,name=search_index_id,json=searchIndexId,proto3" json:"search_index_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchCreateSearchIndexFileRequest) Reset() {
	*x = BatchCreateSearchIndexFileRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchCreateSearchIndexFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateSearchIndexFileRequest) ProtoMessage() {}

func (x *BatchCreateSearchIndexFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[0]
	if x != nil {
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Files         []*SearchIndexFile     `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchCreateSearchIndexFileResponse) Reset() {
	*x = BatchCreateSearchIndexFileResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchCreateSearchIndexFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchCreateSearchIndexFileResponse) ProtoMessage() {}

func (x *BatchCreateSearchIndexFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the file to retrieve.
	FileId string `protobuf:"bytes,1,opt,name=file_id,json=fileId,proto3" json:"file_id,omitempty"`
	// ID of the search index that contains the file.
	SearchIndexId string `protobuf:"bytes,2,opt,name=search_index_id,json=searchIndexId,proto3" json:"search_index_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetSearchIndexFileRequest) Reset() {
	*x = GetSearchIndexFileRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetSearchIndexFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSearchIndexFileRequest) ProtoMessage() {}

func (x *GetSearchIndexFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[2]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the search index whose files will be listed.
	SearchIndexId string `protobuf:"bytes,1,opt,name=search_index_id,json=searchIndexId,proto3" json:"search_index_id,omitempty"`
	// Maximum number of files to return per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to retrieve the next page of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSearchIndexFilesRequest) Reset() {
	*x = ListSearchIndexFilesRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSearchIndexFilesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSearchIndexFilesRequest) ProtoMessage() {}

func (x *ListSearchIndexFilesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[3]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of files in the specified search index.
	Files []*SearchIndexFile `protobuf:"bytes,1,rep,name=files,proto3" json:"files,omitempty"`
	// Token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListSearchIndexFilesResponse) Reset() {
	*x = ListSearchIndexFilesResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListSearchIndexFilesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSearchIndexFilesResponse) ProtoMessage() {}

func (x *ListSearchIndexFilesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_msgTypes[4]
	if x != nil {
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

const file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc = "" +
	"\n" +
	"Iyandex/cloud/ai/assistants/v1/searchindex/search_index_file_service.proto\x12)yandex.cloud.ai.assistants.v1.searchindex\x1aAyandex/cloud/ai/assistants/v1/searchindex/search_index_file.proto\x1a yandex/cloud/api/operation.proto\x1a&yandex/cloud/operation/operation.proto\x1a\x1dyandex/cloud/validation.proto\x1a\x1cgoogle/api/annotations.proto\"t\n" +
	"!BatchCreateSearchIndexFileRequest\x12!\n" +
	"\bfile_ids\x18\x01 \x03(\tB\x06\x82\xc81\x02>0R\afileIds\x12,\n" +
	"\x0fsearch_index_id\x18\x02 \x01(\tB\x04\xe8\xc71\x01R\rsearchIndexId\"v\n" +
	"\"BatchCreateSearchIndexFileResponse\x12P\n" +
	"\x05files\x18\x01 \x03(\v2:.yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileR\x05files\"h\n" +
	"\x19GetSearchIndexFileRequest\x12\x1d\n" +
	"\afile_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x06fileId\x12,\n" +
	"\x0fsearch_index_id\x18\x02 \x01(\tB\x04\xe8\xc71\x01R\rsearchIndexId\"\x87\x01\n" +
	"\x1bListSearchIndexFilesRequest\x12,\n" +
	"\x0fsearch_index_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\rsearchIndexId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"\x98\x01\n" +
	"\x1cListSearchIndexFilesResponse\x12P\n" +
	"\x05files\x18\x01 \x03(\v2:.yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFileR\x05files\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken2\x9a\x05\n" +
	"\x16SearchIndexFileService\x12\xdd\x01\n" +
	"\vBatchCreate\x12L.yandex.cloud.ai.assistants.v1.searchindex.BatchCreateSearchIndexFileRequest\x1a!.yandex.cloud.operation.Operation\"]\xb2\xd2*$\x12\"BatchCreateSearchIndexFileResponse\x82\xd3\xe4\x93\x02/:\x01*\"*/assistants/v1/searchIndexFile:batchCreate\x12\xcb\x01\n" +
	"\x03Get\x12D.yandex.cloud.ai.assistants.v1.searchindex.GetSearchIndexFileRequest\x1a:.yandex.cloud.ai.assistants.v1.searchindex.SearchIndexFile\"B\x82\xd3\xe4\x93\x02<\x12:/assistants/v1/searchIndexFile/{search_index_id}/{file_id}\x12\xd1\x01\n" +
	"\x04List\x12F.yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesRequest\x1aG.yandex.cloud.ai.assistants.v1.searchindex.ListSearchIndexFilesResponse\"8\x82\xd3\xe4\x93\x022\x120/assistants/v1/searchIndexFile/{search_index_id}B\x8a\x01\n" +
	"-yandex.cloud.api.ai.assistants.v1.searchindexZYgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1/searchindex;searchindexb\x06proto3"

var (
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescData []byte
)

func file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc)))
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_rawDesc)),
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
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_searchindex_search_index_file_service_proto_depIdxs = nil
}
