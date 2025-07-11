// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ydb/v1/storage_type_service.proto

package ydb

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type GetStorageTypeRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. ID of the storage type to return.
	StorageTypeId string `protobuf:"bytes,1,opt,name=storage_type_id,json=storageTypeId,proto3" json:"storage_type_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStorageTypeRequest) Reset() {
	*x = GetStorageTypeRequest{}
	mi := &file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStorageTypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStorageTypeRequest) ProtoMessage() {}

func (x *GetStorageTypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStorageTypeRequest.ProtoReflect.Descriptor instead.
func (*GetStorageTypeRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetStorageTypeRequest) GetStorageTypeId() string {
	if x != nil {
		return x.StorageTypeId
	}
	return ""
}

type ListStorageTypesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of results per page that should be returned. If the number of available
	// results is larger than `page_size`, the service returns a `next_page_token` that can be used
	// to get the next page of results in subsequent ListStorageTypes requests.
	// Acceptable values are 0 to 1000, inclusive. Default value: 100.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. Set `page_token` to the `next_page_token` returned by a previous ListStorageTypes
	// request to get the next page of results.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStorageTypesRequest) Reset() {
	*x = ListStorageTypesRequest{}
	mi := &file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStorageTypesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageTypesRequest) ProtoMessage() {}

func (x *ListStorageTypesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageTypesRequest.ProtoReflect.Descriptor instead.
func (*ListStorageTypesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListStorageTypesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListStorageTypesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListStorageTypesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Requested list of storage types.
	StorageTypes []*StorageType `protobuf:"bytes,1,rep,name=storage_types,json=storageTypes,proto3" json:"storage_types,omitempty"`
	// This token allows you to get the next page of results for ListStorageTypes requests,
	// if the number of results is larger than `page_size` specified in the request.
	// To get the next page, specify the value of `next_page_token` as a value for
	// the `page_token` parameter in the next ListStorageTypes request. Subsequent ListStorageTypes
	// requests will have their own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStorageTypesResponse) Reset() {
	*x = ListStorageTypesResponse{}
	mi := &file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStorageTypesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStorageTypesResponse) ProtoMessage() {}

func (x *ListStorageTypesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStorageTypesResponse.ProtoReflect.Descriptor instead.
func (*ListStorageTypesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListStorageTypesResponse) GetStorageTypes() []*StorageType {
	if x != nil {
		return x.StorageTypes
	}
	return nil
}

func (x *ListStorageTypesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_ydb_v1_storage_type_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDesc = "" +
	"\n" +
	".yandex/cloud/ydb/v1/storage_type_service.proto\x12\x13yandex.cloud.ydb.v1\x1a\x1cgoogle/api/annotations.proto\x1a&yandex/cloud/ydb/v1/storage_type.proto\x1a\x1dyandex/cloud/validation.proto\"E\n" +
	"\x15GetStorageTypeRequest\x12,\n" +
	"\x0fstorage_type_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\rstorageTypeId\"l\n" +
	"\x17ListStorageTypesRequest\x12'\n" +
	"\tpage_size\x18\x01 \x01(\x03B\n" +
	"\xfa\xc71\x060-1000R\bpageSize\x12(\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tB\t\x8a\xc81\x05<=100R\tpageToken\"\x89\x01\n" +
	"\x18ListStorageTypesResponse\x12E\n" +
	"\rstorage_types\x18\x01 \x03(\v2 .yandex.cloud.ydb.v1.StorageTypeR\fstorageTypes\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken2\x9e\x02\n" +
	"\x12StorageTypeService\x12\x83\x01\n" +
	"\x03Get\x12*.yandex.cloud.ydb.v1.GetStorageTypeRequest\x1a .yandex.cloud.ydb.v1.StorageType\".\x82\xd3\xe4\x93\x02(\x12&/ydb/v1/storageTypes/{storage_type_id}\x12\x81\x01\n" +
	"\x04List\x12,.yandex.cloud.ydb.v1.ListStorageTypesRequest\x1a-.yandex.cloud.ydb.v1.ListStorageTypesResponse\"\x1c\x82\xd3\xe4\x93\x02\x16\x12\x14/ydb/v1/storageTypesBV\n" +
	"\x17yandex.cloud.api.ydb.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/ydb/v1;ydbb\x06proto3"

var (
	file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescData []byte
)

func file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDesc), len(file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDesc)))
	})
	return file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDescData
}

var file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_ydb_v1_storage_type_service_proto_goTypes = []any{
	(*GetStorageTypeRequest)(nil),    // 0: yandex.cloud.ydb.v1.GetStorageTypeRequest
	(*ListStorageTypesRequest)(nil),  // 1: yandex.cloud.ydb.v1.ListStorageTypesRequest
	(*ListStorageTypesResponse)(nil), // 2: yandex.cloud.ydb.v1.ListStorageTypesResponse
	(*StorageType)(nil),              // 3: yandex.cloud.ydb.v1.StorageType
}
var file_yandex_cloud_ydb_v1_storage_type_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.ydb.v1.ListStorageTypesResponse.storage_types:type_name -> yandex.cloud.ydb.v1.StorageType
	0, // 1: yandex.cloud.ydb.v1.StorageTypeService.Get:input_type -> yandex.cloud.ydb.v1.GetStorageTypeRequest
	1, // 2: yandex.cloud.ydb.v1.StorageTypeService.List:input_type -> yandex.cloud.ydb.v1.ListStorageTypesRequest
	3, // 3: yandex.cloud.ydb.v1.StorageTypeService.Get:output_type -> yandex.cloud.ydb.v1.StorageType
	2, // 4: yandex.cloud.ydb.v1.StorageTypeService.List:output_type -> yandex.cloud.ydb.v1.ListStorageTypesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ydb_v1_storage_type_service_proto_init() }
func file_yandex_cloud_ydb_v1_storage_type_service_proto_init() {
	if File_yandex_cloud_ydb_v1_storage_type_service_proto != nil {
		return
	}
	file_yandex_cloud_ydb_v1_storage_type_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDesc), len(file_yandex_cloud_ydb_v1_storage_type_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ydb_v1_storage_type_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ydb_v1_storage_type_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ydb_v1_storage_type_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ydb_v1_storage_type_service_proto = out.File
	file_yandex_cloud_ydb_v1_storage_type_service_proto_goTypes = nil
	file_yandex_cloud_ydb_v1_storage_type_service_proto_depIdxs = nil
}
