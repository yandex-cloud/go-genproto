// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/baremetal/v1alpha/zone_service.proto

package baremetal

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

type GetZoneRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Zone resource to return.
	//
	// To get the zone ID, use a [ZoneService.List] request.
	ZoneId        string `protobuf:"bytes,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetZoneRequest) Reset() {
	*x = GetZoneRequest{}
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneRequest) ProtoMessage() {}

func (x *GetZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetZoneRequest.ProtoReflect.Descriptor instead.
func (*GetZoneRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetZoneRequest) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

type ListZonesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of results per page to return. If the number of available
	// results is greater than `page_size`,
	// the service returns a [ListZonesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value is 20.
	PageSize int64 `protobuf:"varint,100,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListZonesResponse.next_page_token] returned by a previous list request.
	PageToken     string `protobuf:"bytes,101,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListZonesRequest) Reset() {
	*x = ListZonesRequest{}
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListZonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZonesRequest) ProtoMessage() {}

func (x *ListZonesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZonesRequest.ProtoReflect.Descriptor instead.
func (*ListZonesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListZonesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListZonesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListZonesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of Zone resources.
	Zones []*Zone `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// [ListZonesRequest.page_size], use `next_page_token` as the value
	// for the [ListZonesRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListZonesResponse) Reset() {
	*x = ListZonesResponse{}
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListZonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZonesResponse) ProtoMessage() {}

func (x *ListZonesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListZonesResponse.ProtoReflect.Descriptor instead.
func (*ListZonesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListZonesResponse) GetZones() []*Zone {
	if x != nil {
		return x.Zones
	}
	return nil
}

func (x *ListZonesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_baremetal_v1alpha_zone_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc = "" +
	"\n" +
	"1yandex/cloud/baremetal/v1alpha/zone_service.proto\x12\x1eyandex.cloud.baremetal.v1alpha\x1a)yandex/cloud/baremetal/v1alpha/zone.proto\x1a\x1dyandex/cloud/validation.proto\"3\n" +
	"\x0eGetZoneRequest\x12!\n" +
	"\azone_id\x18\x01 \x01(\tB\b\x8a\xc81\x04<=20R\x06zoneId\"_\n" +
	"\x10ListZonesRequest\x12&\n" +
	"\tpage_size\x18d \x01(\x03B\t\xfa\xc71\x05<=100R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18e \x01(\tR\tpageTokenJ\x04\b\x01\x10d\"}\n" +
	"\x11ListZonesResponse\x12:\n" +
	"\x05zones\x18\x01 \x03(\v2$.yandex.cloud.baremetal.v1alpha.ZoneR\x05zones\x12&\n" +
	"\x0fnext_page_token\x18d \x01(\tR\rnextPageTokenJ\x04\b\x02\x10d2\xdb\x01\n" +
	"\vZoneService\x12]\n" +
	"\x03Get\x12..yandex.cloud.baremetal.v1alpha.GetZoneRequest\x1a$.yandex.cloud.baremetal.v1alpha.Zone\"\x00\x12m\n" +
	"\x04List\x120.yandex.cloud.baremetal.v1alpha.ListZonesRequest\x1a1.yandex.cloud.baremetal.v1alpha.ListZonesResponse\"\x00Br\n" +
	"\"yandex.cloud.api.baremetal.v1alphaZLgithub.com/yandex-cloud/go-genproto/yandex/cloud/baremetal/v1alpha;baremetalb\x06proto3"

var (
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescData []byte
)

func file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc), len(file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc)))
	})
	return file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescData
}

var file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_baremetal_v1alpha_zone_service_proto_goTypes = []any{
	(*GetZoneRequest)(nil),    // 0: yandex.cloud.baremetal.v1alpha.GetZoneRequest
	(*ListZonesRequest)(nil),  // 1: yandex.cloud.baremetal.v1alpha.ListZonesRequest
	(*ListZonesResponse)(nil), // 2: yandex.cloud.baremetal.v1alpha.ListZonesResponse
	(*Zone)(nil),              // 3: yandex.cloud.baremetal.v1alpha.Zone
}
var file_yandex_cloud_baremetal_v1alpha_zone_service_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.baremetal.v1alpha.ListZonesResponse.zones:type_name -> yandex.cloud.baremetal.v1alpha.Zone
	0, // 1: yandex.cloud.baremetal.v1alpha.ZoneService.Get:input_type -> yandex.cloud.baremetal.v1alpha.GetZoneRequest
	1, // 2: yandex.cloud.baremetal.v1alpha.ZoneService.List:input_type -> yandex.cloud.baremetal.v1alpha.ListZonesRequest
	3, // 3: yandex.cloud.baremetal.v1alpha.ZoneService.Get:output_type -> yandex.cloud.baremetal.v1alpha.Zone
	2, // 4: yandex.cloud.baremetal.v1alpha.ZoneService.List:output_type -> yandex.cloud.baremetal.v1alpha.ListZonesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_baremetal_v1alpha_zone_service_proto_init() }
func file_yandex_cloud_baremetal_v1alpha_zone_service_proto_init() {
	if File_yandex_cloud_baremetal_v1alpha_zone_service_proto != nil {
		return
	}
	file_yandex_cloud_baremetal_v1alpha_zone_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc), len(file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_baremetal_v1alpha_zone_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_baremetal_v1alpha_zone_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_baremetal_v1alpha_zone_service_proto = out.File
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_goTypes = nil
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_depIdxs = nil
}
