// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/baremetal/v1alpha/zone_service.proto

package baremetal

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

type GetZoneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Zone resource to return.
	//
	// To get the zone ID, use a [ZoneService.List] request.
	ZoneId string `protobuf:"bytes,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
}

func (x *GetZoneRequest) Reset() {
	*x = GetZoneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetZoneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetZoneRequest) ProtoMessage() {}

func (x *GetZoneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The maximum number of results per page to return. If the number of available
	// results is greater than `page_size`,
	// the service returns a [ListZonesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value is 20.
	PageSize int64 `protobuf:"varint,100,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set `page_token` to the
	// [ListZonesResponse.next_page_token] returned by a previous list request.
	PageToken string `protobuf:"bytes,101,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListZonesRequest) Reset() {
	*x = ListZonesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZonesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZonesRequest) ProtoMessage() {}

func (x *ListZonesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of Zone resources.
	Zones []*Zone `protobuf:"bytes,1,rep,name=zones,proto3" json:"zones,omitempty"`
	// Token for getting the next page of the list. If the number of results is greater than
	// [ListZonesRequest.page_size], use `next_page_token` as the value
	// for the [ListZonesRequest.page_token] parameter in the next list request.
	//
	// Each subsequent page will have its own `next_page_token` to continue paging through the results.
	NextPageToken string `protobuf:"bytes,100,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListZonesResponse) Reset() {
	*x = ListZonesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListZonesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListZonesResponse) ProtoMessage() {}

func (x *ListZonesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x1a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x32, 0x30, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65,
	0x49, 0x64, 0x22, 0x5f, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x64, 0x20, 0x01, 0x28, 0x03, 0x42, 0x09, 0xfa, 0xc7, 0x31, 0x05, 0x3c,
	0x3d, 0x31, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x65, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x64, 0x22, 0x7d, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x7a, 0x6f, 0x6e, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x52, 0x05, 0x7a,
	0x6f, 0x6e, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08, 0x02,
	0x10, 0x64, 0x32, 0xdb, 0x01, 0x0a, 0x0b, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5d, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x47, 0x65, 0x74, 0x5a, 0x6f,
	0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x5a, 0x6f, 0x6e, 0x65, 0x22,
	0x00, 0x12, 0x6d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x5a,
	0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x5a, 0x6f, 0x6e, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x72, 0x0a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescData = file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc
)

func file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetZoneRequest); i {
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
		file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListZonesRequest); i {
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
		file_yandex_cloud_baremetal_v1alpha_zone_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListZonesResponse); i {
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
			RawDescriptor: file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc,
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
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_rawDesc = nil
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_goTypes = nil
	file_yandex_cloud_baremetal_v1alpha_zone_service_proto_depIdxs = nil
}
