// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/searchapi/v2/search_query.proto

package searchapi

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type SearchQuery_SearchType int32

const (
	SearchQuery_SEARCH_TYPE_UNSPECIFIED SearchQuery_SearchType = 0
	// Russian search type (default), yandex.ru search domain name will be used.
	SearchQuery_SEARCH_TYPE_RU SearchQuery_SearchType = 1
	// Turkish search type, yandex.tr search domain name will be used.
	SearchQuery_SEARCH_TYPE_TR SearchQuery_SearchType = 2
	// International search type, yandex.com search domain name will be used.
	SearchQuery_SEARCH_TYPE_COM SearchQuery_SearchType = 3
)

// Enum value maps for SearchQuery_SearchType.
var (
	SearchQuery_SearchType_name = map[int32]string{
		0: "SEARCH_TYPE_UNSPECIFIED",
		1: "SEARCH_TYPE_RU",
		2: "SEARCH_TYPE_TR",
		3: "SEARCH_TYPE_COM",
	}
	SearchQuery_SearchType_value = map[string]int32{
		"SEARCH_TYPE_UNSPECIFIED": 0,
		"SEARCH_TYPE_RU":          1,
		"SEARCH_TYPE_TR":          2,
		"SEARCH_TYPE_COM":         3,
	}
)

func (x SearchQuery_SearchType) Enum() *SearchQuery_SearchType {
	p := new(SearchQuery_SearchType)
	*p = x
	return p
}

func (x SearchQuery_SearchType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchQuery_SearchType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes[0].Descriptor()
}

func (SearchQuery_SearchType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes[0]
}

func (x SearchQuery_SearchType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchQuery_SearchType.Descriptor instead.
func (SearchQuery_SearchType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_searchapi_v2_search_query_proto_rawDescGZIP(), []int{0, 0}
}

type SearchQuery_FamilyMode int32

const (
	SearchQuery_FAMILY_MODE_UNSPECIFIED SearchQuery_FamilyMode = 0
	// Filtering is disabled. Search results include any documents regardless of their contents.
	SearchQuery_FAMILY_MODE_NONE SearchQuery_FamilyMode = 1
	// Moderate filter (default value). Documents of the Adult category are excluded from search results
	// unless a query is explicitly made for searching resources of this category.
	SearchQuery_FAMILY_MODE_MODERATE SearchQuery_FamilyMode = 2
	// Regardless of a search query, documents of the Adult category
	// and those with profanity are excluded from search results.
	SearchQuery_FAMILY_MODE_STRICT SearchQuery_FamilyMode = 3
)

// Enum value maps for SearchQuery_FamilyMode.
var (
	SearchQuery_FamilyMode_name = map[int32]string{
		0: "FAMILY_MODE_UNSPECIFIED",
		1: "FAMILY_MODE_NONE",
		2: "FAMILY_MODE_MODERATE",
		3: "FAMILY_MODE_STRICT",
	}
	SearchQuery_FamilyMode_value = map[string]int32{
		"FAMILY_MODE_UNSPECIFIED": 0,
		"FAMILY_MODE_NONE":        1,
		"FAMILY_MODE_MODERATE":    2,
		"FAMILY_MODE_STRICT":      3,
	}
)

func (x SearchQuery_FamilyMode) Enum() *SearchQuery_FamilyMode {
	p := new(SearchQuery_FamilyMode)
	*p = x
	return p
}

func (x SearchQuery_FamilyMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchQuery_FamilyMode) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes[1].Descriptor()
}

func (SearchQuery_FamilyMode) Type() protoreflect.EnumType {
	return &file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes[1]
}

func (x SearchQuery_FamilyMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchQuery_FamilyMode.Descriptor instead.
func (SearchQuery_FamilyMode) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_searchapi_v2_search_query_proto_rawDescGZIP(), []int{0, 1}
}

type SearchQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Search type that determines the domain name that will be used for the search queries.
	SearchType SearchQuery_SearchType `protobuf:"varint,1,opt,name=search_type,json=searchType,proto3,enum=yandex.cloud.searchapi.v2.SearchQuery_SearchType" json:"search_type,omitempty"`
	// Search query text
	QueryText string `protobuf:"bytes,2,opt,name=query_text,json=queryText,proto3" json:"query_text,omitempty"`
	// Rule for filtering search results and determines whether any documents should be excluded.
	FamilyMode SearchQuery_FamilyMode `protobuf:"varint,3,opt,name=family_mode,json=familyMode,proto3,enum=yandex.cloud.searchapi.v2.SearchQuery_FamilyMode" json:"family_mode,omitempty"`
	// The number of a requested page with search results
	Page int64 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *SearchQuery) Reset() {
	*x = SearchQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuery) ProtoMessage() {}

func (x *SearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchQuery.ProtoReflect.Descriptor instead.
func (*SearchQuery) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_searchapi_v2_search_query_proto_rawDescGZIP(), []int{0}
}

func (x *SearchQuery) GetSearchType() SearchQuery_SearchType {
	if x != nil {
		return x.SearchType
	}
	return SearchQuery_SEARCH_TYPE_UNSPECIFIED
}

func (x *SearchQuery) GetQueryText() string {
	if x != nil {
		return x.QueryText
	}
	return ""
}

func (x *SearchQuery) GetFamilyMode() SearchQuery_FamilyMode {
	if x != nil {
		return x.FamilyMode
	}
	return SearchQuery_FAMILY_MODE_UNSPECIFIED
}

func (x *SearchQuery) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

var File_yandex_cloud_searchapi_v2_search_query_proto protoreflect.FileDescriptor

var file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x2f, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x03, 0x0a, 0x0b, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x58, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2c, 0x0a, 0x0a, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05,
	0x3c, 0x3d, 0x34, 0x30, 0x30, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x54, 0x65, 0x78, 0x74,
	0x12, 0x52, 0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x32, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x46, 0x61,
	0x6d, 0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x07, 0xfa, 0xc7, 0x31, 0x03, 0x3e, 0x3d, 0x30, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x22, 0x66, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1b, 0x0a, 0x17, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e,
	0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x55, 0x10, 0x01,
	0x12, 0x12, 0x0a, 0x0e, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x54, 0x52, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x10, 0x03, 0x22, 0x71, 0x0a, 0x0a, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x46, 0x41, 0x4d, 0x49, 0x4c,
	0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x46, 0x41,
	0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x52, 0x41,
	0x54, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x5f, 0x4d,
	0x4f, 0x44, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x43, 0x54, 0x10, 0x03, 0x42, 0x65, 0x0a, 0x1a,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x2e, 0x76, 0x32, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x32, 0x3b, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDescOnce sync.Once
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData = file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc
)

func file_yandex_cloud_searchapi_v2_search_query_proto_rawDescGZIP() []byte {
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData)
	})
	return file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData
}

var file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_searchapi_v2_search_query_proto_goTypes = []any{
	(SearchQuery_SearchType)(0), // 0: yandex.cloud.searchapi.v2.SearchQuery.SearchType
	(SearchQuery_FamilyMode)(0), // 1: yandex.cloud.searchapi.v2.SearchQuery.FamilyMode
	(*SearchQuery)(nil),         // 2: yandex.cloud.searchapi.v2.SearchQuery
}
var file_yandex_cloud_searchapi_v2_search_query_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.searchapi.v2.SearchQuery.search_type:type_name -> yandex.cloud.searchapi.v2.SearchQuery.SearchType
	1, // 1: yandex.cloud.searchapi.v2.SearchQuery.family_mode:type_name -> yandex.cloud.searchapi.v2.SearchQuery.FamilyMode
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_searchapi_v2_search_query_proto_init() }
func file_yandex_cloud_searchapi_v2_search_query_proto_init() {
	if File_yandex_cloud_searchapi_v2_search_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SearchQuery); i {
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
			RawDescriptor: file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_searchapi_v2_search_query_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_searchapi_v2_search_query_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes,
	}.Build()
	File_yandex_cloud_searchapi_v2_search_query_proto = out.File
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc = nil
	file_yandex_cloud_searchapi_v2_search_query_proto_goTypes = nil
	file_yandex_cloud_searchapi_v2_search_query_proto_depIdxs = nil
}
