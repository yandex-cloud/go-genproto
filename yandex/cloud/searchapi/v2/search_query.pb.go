// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/searchapi/v2/search_query.proto

package searchapi

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type SearchQuery_SearchType int32

const (
	SearchQuery_SEARCH_TYPE_UNSPECIFIED SearchQuery_SearchType = 0
	// Russian search type (default), yandex.ru search domain name will be used.
	SearchQuery_SEARCH_TYPE_RU SearchQuery_SearchType = 1
	// Turkish search type, yandex.tr search domain name will be used.
	SearchQuery_SEARCH_TYPE_TR SearchQuery_SearchType = 2
	// International search type, yandex.com search domain name will be used.
	SearchQuery_SEARCH_TYPE_COM SearchQuery_SearchType = 3
	// Kazakh search type, yandex.kz search domain name will be used.
	SearchQuery_SEARCH_TYPE_KK SearchQuery_SearchType = 4
	// Belarusian search type, yandex.by search domain name will be used.
	SearchQuery_SEARCH_TYPE_BE SearchQuery_SearchType = 5
	// Uzbek search type, yandex.uz search domain name will be used.
	SearchQuery_SEARCH_TYPE_UZ SearchQuery_SearchType = 6
)

// Enum value maps for SearchQuery_SearchType.
var (
	SearchQuery_SearchType_name = map[int32]string{
		0: "SEARCH_TYPE_UNSPECIFIED",
		1: "SEARCH_TYPE_RU",
		2: "SEARCH_TYPE_TR",
		3: "SEARCH_TYPE_COM",
		4: "SEARCH_TYPE_KK",
		5: "SEARCH_TYPE_BE",
		6: "SEARCH_TYPE_UZ",
	}
	SearchQuery_SearchType_value = map[string]int32{
		"SEARCH_TYPE_UNSPECIFIED": 0,
		"SEARCH_TYPE_RU":          1,
		"SEARCH_TYPE_TR":          2,
		"SEARCH_TYPE_COM":         3,
		"SEARCH_TYPE_KK":          4,
		"SEARCH_TYPE_BE":          5,
		"SEARCH_TYPE_UZ":          6,
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

type SearchQuery_FixTypoMode int32

const (
	SearchQuery_FIX_TYPO_MODE_UNSPECIFIED SearchQuery_FixTypoMode = 0
	// Automatically correct typos (default value).
	SearchQuery_FIX_TYPO_MODE_ON SearchQuery_FixTypoMode = 1
	// Autocorrection is off.
	SearchQuery_FIX_TYPO_MODE_OFF SearchQuery_FixTypoMode = 2
)

// Enum value maps for SearchQuery_FixTypoMode.
var (
	SearchQuery_FixTypoMode_name = map[int32]string{
		0: "FIX_TYPO_MODE_UNSPECIFIED",
		1: "FIX_TYPO_MODE_ON",
		2: "FIX_TYPO_MODE_OFF",
	}
	SearchQuery_FixTypoMode_value = map[string]int32{
		"FIX_TYPO_MODE_UNSPECIFIED": 0,
		"FIX_TYPO_MODE_ON":          1,
		"FIX_TYPO_MODE_OFF":         2,
	}
)

func (x SearchQuery_FixTypoMode) Enum() *SearchQuery_FixTypoMode {
	p := new(SearchQuery_FixTypoMode)
	*p = x
	return p
}

func (x SearchQuery_FixTypoMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SearchQuery_FixTypoMode) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes[2].Descriptor()
}

func (SearchQuery_FixTypoMode) Type() protoreflect.EnumType {
	return &file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes[2]
}

func (x SearchQuery_FixTypoMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SearchQuery_FixTypoMode.Descriptor instead.
func (SearchQuery_FixTypoMode) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_searchapi_v2_search_query_proto_rawDescGZIP(), []int{0, 2}
}

type SearchQuery struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Search type that determines the domain name that will be used for the search queries.
	SearchType SearchQuery_SearchType `protobuf:"varint,1,opt,name=search_type,json=searchType,proto3,enum=yandex.cloud.searchapi.v2.SearchQuery_SearchType" json:"search_type,omitempty"`
	// Search query text
	QueryText string `protobuf:"bytes,2,opt,name=query_text,json=queryText,proto3" json:"query_text,omitempty"`
	// Rule for filtering search results and determines whether any documents should be excluded.
	FamilyMode SearchQuery_FamilyMode `protobuf:"varint,3,opt,name=family_mode,json=familyMode,proto3,enum=yandex.cloud.searchapi.v2.SearchQuery_FamilyMode" json:"family_mode,omitempty"`
	// The number of a requested page with search results
	Page int64 `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	// Typos autocorrections mode
	FixTypoMode   SearchQuery_FixTypoMode `protobuf:"varint,5,opt,name=fix_typo_mode,json=fixTypoMode,proto3,enum=yandex.cloud.searchapi.v2.SearchQuery_FixTypoMode" json:"fix_typo_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SearchQuery) Reset() {
	*x = SearchQuery{}
	mi := &file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SearchQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchQuery) ProtoMessage() {}

func (x *SearchQuery) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes[0]
	if x != nil {
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

func (x *SearchQuery) GetFixTypoMode() SearchQuery_FixTypoMode {
	if x != nil {
		return x.FixTypoMode
	}
	return SearchQuery_FIX_TYPO_MODE_UNSPECIFIED
}

var File_yandex_cloud_searchapi_v2_search_query_proto protoreflect.FileDescriptor

const file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc = "" +
	"\n" +
	",yandex/cloud/searchapi/v2/search_query.proto\x12\x19yandex.cloud.searchapi.v2\x1a\x1dyandex/cloud/validation.proto\"\xd1\x05\n" +
	"\vSearchQuery\x12X\n" +
	"\vsearch_type\x18\x01 \x01(\x0e21.yandex.cloud.searchapi.v2.SearchQuery.SearchTypeB\x04\xe8\xc71\x01R\n" +
	"searchType\x12,\n" +
	"\n" +
	"query_text\x18\x02 \x01(\tB\r\xe8\xc71\x01\x8a\xc81\x05<=400R\tqueryText\x12R\n" +
	"\vfamily_mode\x18\x03 \x01(\x0e21.yandex.cloud.searchapi.v2.SearchQuery.FamilyModeR\n" +
	"familyMode\x12\x1b\n" +
	"\x04page\x18\x04 \x01(\x03B\a\xfa\xc71\x03>=0R\x04page\x12V\n" +
	"\rfix_typo_mode\x18\x05 \x01(\x0e22.yandex.cloud.searchapi.v2.SearchQuery.FixTypoModeR\vfixTypoMode\"\xa2\x01\n" +
	"\n" +
	"SearchType\x12\x1b\n" +
	"\x17SEARCH_TYPE_UNSPECIFIED\x10\x00\x12\x12\n" +
	"\x0eSEARCH_TYPE_RU\x10\x01\x12\x12\n" +
	"\x0eSEARCH_TYPE_TR\x10\x02\x12\x13\n" +
	"\x0fSEARCH_TYPE_COM\x10\x03\x12\x12\n" +
	"\x0eSEARCH_TYPE_KK\x10\x04\x12\x12\n" +
	"\x0eSEARCH_TYPE_BE\x10\x05\x12\x12\n" +
	"\x0eSEARCH_TYPE_UZ\x10\x06\"q\n" +
	"\n" +
	"FamilyMode\x12\x1b\n" +
	"\x17FAMILY_MODE_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10FAMILY_MODE_NONE\x10\x01\x12\x18\n" +
	"\x14FAMILY_MODE_MODERATE\x10\x02\x12\x16\n" +
	"\x12FAMILY_MODE_STRICT\x10\x03\"Y\n" +
	"\vFixTypoMode\x12\x1d\n" +
	"\x19FIX_TYPO_MODE_UNSPECIFIED\x10\x00\x12\x14\n" +
	"\x10FIX_TYPO_MODE_ON\x10\x01\x12\x15\n" +
	"\x11FIX_TYPO_MODE_OFF\x10\x02Be\n" +
	"\x1ayandex.cloud.api.search.v2ZGgithub.com/yandex-cloud/go-genproto/yandex/cloud/searchapi/v2;searchapib\x06proto3"

var (
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDescOnce sync.Once
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData []byte
)

func file_yandex_cloud_searchapi_v2_search_query_proto_rawDescGZIP() []byte {
	file_yandex_cloud_searchapi_v2_search_query_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc), len(file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc)))
	})
	return file_yandex_cloud_searchapi_v2_search_query_proto_rawDescData
}

var file_yandex_cloud_searchapi_v2_search_query_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_yandex_cloud_searchapi_v2_search_query_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_searchapi_v2_search_query_proto_goTypes = []any{
	(SearchQuery_SearchType)(0),  // 0: yandex.cloud.searchapi.v2.SearchQuery.SearchType
	(SearchQuery_FamilyMode)(0),  // 1: yandex.cloud.searchapi.v2.SearchQuery.FamilyMode
	(SearchQuery_FixTypoMode)(0), // 2: yandex.cloud.searchapi.v2.SearchQuery.FixTypoMode
	(*SearchQuery)(nil),          // 3: yandex.cloud.searchapi.v2.SearchQuery
}
var file_yandex_cloud_searchapi_v2_search_query_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.searchapi.v2.SearchQuery.search_type:type_name -> yandex.cloud.searchapi.v2.SearchQuery.SearchType
	1, // 1: yandex.cloud.searchapi.v2.SearchQuery.family_mode:type_name -> yandex.cloud.searchapi.v2.SearchQuery.FamilyMode
	2, // 2: yandex.cloud.searchapi.v2.SearchQuery.fix_typo_mode:type_name -> yandex.cloud.searchapi.v2.SearchQuery.FixTypoMode
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_searchapi_v2_search_query_proto_init() }
func file_yandex_cloud_searchapi_v2_search_query_proto_init() {
	if File_yandex_cloud_searchapi_v2_search_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc), len(file_yandex_cloud_searchapi_v2_search_query_proto_rawDesc)),
			NumEnums:      3,
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
	file_yandex_cloud_searchapi_v2_search_query_proto_goTypes = nil
	file_yandex_cloud_searchapi_v2_search_query_proto_depIdxs = nil
}
