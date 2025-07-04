// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/apploadbalancer/v1/http_router.proto

package apploadbalancer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// An HTTP router resource.
// For details about the concept, see [documentation](/docs/application-load-balancer/concepts/http-router).
type HttpRouter struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the router. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the router. The name is unique within the folder.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the router.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the folder that the router belongs to.
	FolderId string `protobuf:"bytes,4,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Router labels as `key:value` pairs.
	// For details about the concept, see [documentation](/docs/overview/concepts/services#labels).
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Virtual hosts that combine routes inside the router.
	// For details about the concept, see [documentation](/docs/application-load-balancer/concepts/http-router#virtual-host).
	//
	// Only one virtual host with no authority (default match) can be specified.
	VirtualHosts []*VirtualHost `protobuf:"bytes,6,rep,name=virtual_hosts,json=virtualHosts,proto3" json:"virtual_hosts,omitempty"`
	// Creation timestamp.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	RouteOptions  *RouteOptions          `protobuf:"bytes,8,opt,name=route_options,json=routeOptions,proto3" json:"route_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *HttpRouter) Reset() {
	*x = HttpRouter{}
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *HttpRouter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpRouter) ProtoMessage() {}

func (x *HttpRouter) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_http_router_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpRouter.ProtoReflect.Descriptor instead.
func (*HttpRouter) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescGZIP(), []int{0}
}

func (x *HttpRouter) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HttpRouter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HttpRouter) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *HttpRouter) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *HttpRouter) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *HttpRouter) GetVirtualHosts() []*VirtualHost {
	if x != nil {
		return x.VirtualHosts
	}
	return nil
}

func (x *HttpRouter) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *HttpRouter) GetRouteOptions() *RouteOptions {
	if x != nil {
		return x.RouteOptions
	}
	return nil
}

var File_yandex_cloud_apploadbalancer_v1_http_router_proto protoreflect.FileDescriptor

const file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDesc = "" +
	"\n" +
	"1yandex/cloud/apploadbalancer/v1/http_router.proto\x12\x1fyandex.cloud.apploadbalancer.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a2yandex/cloud/apploadbalancer/v1/virtual_host.proto\"\xdd\x03\n" +
	"\n" +
	"HttpRouter\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x1b\n" +
	"\tfolder_id\x18\x04 \x01(\tR\bfolderId\x12O\n" +
	"\x06labels\x18\x05 \x03(\v27.yandex.cloud.apploadbalancer.v1.HttpRouter.LabelsEntryR\x06labels\x12Q\n" +
	"\rvirtual_hosts\x18\x06 \x03(\v2,.yandex.cloud.apploadbalancer.v1.VirtualHostR\fvirtualHosts\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12R\n" +
	"\rroute_options\x18\b \x01(\v2-.yandex.cloud.apploadbalancer.v1.RouteOptionsR\frouteOptions\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01Bz\n" +
	"#yandex.cloud.api.apploadbalancer.v1ZSgithub.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1;apploadbalancerb\x06proto3"

var (
	file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescData []byte
)

func file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDesc)))
	})
	return file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDescData
}

var file_yandex_cloud_apploadbalancer_v1_http_router_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_apploadbalancer_v1_http_router_proto_goTypes = []any{
	(*HttpRouter)(nil),            // 0: yandex.cloud.apploadbalancer.v1.HttpRouter
	nil,                           // 1: yandex.cloud.apploadbalancer.v1.HttpRouter.LabelsEntry
	(*VirtualHost)(nil),           // 2: yandex.cloud.apploadbalancer.v1.VirtualHost
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*RouteOptions)(nil),          // 4: yandex.cloud.apploadbalancer.v1.RouteOptions
}
var file_yandex_cloud_apploadbalancer_v1_http_router_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.apploadbalancer.v1.HttpRouter.labels:type_name -> yandex.cloud.apploadbalancer.v1.HttpRouter.LabelsEntry
	2, // 1: yandex.cloud.apploadbalancer.v1.HttpRouter.virtual_hosts:type_name -> yandex.cloud.apploadbalancer.v1.VirtualHost
	3, // 2: yandex.cloud.apploadbalancer.v1.HttpRouter.created_at:type_name -> google.protobuf.Timestamp
	4, // 3: yandex.cloud.apploadbalancer.v1.HttpRouter.route_options:type_name -> yandex.cloud.apploadbalancer.v1.RouteOptions
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_apploadbalancer_v1_http_router_proto_init() }
func file_yandex_cloud_apploadbalancer_v1_http_router_proto_init() {
	if File_yandex_cloud_apploadbalancer_v1_http_router_proto != nil {
		return
	}
	file_yandex_cloud_apploadbalancer_v1_virtual_host_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_http_router_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_apploadbalancer_v1_http_router_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_apploadbalancer_v1_http_router_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_apploadbalancer_v1_http_router_proto_msgTypes,
	}.Build()
	File_yandex_cloud_apploadbalancer_v1_http_router_proto = out.File
	file_yandex_cloud_apploadbalancer_v1_http_router_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_http_router_proto_depIdxs = nil
}
