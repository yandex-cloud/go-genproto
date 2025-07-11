// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/monitoring/v3/link.proto

package monitoring

import (
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

type Link struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Link title.
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// Open link in new tab.
	OpenInNewTab bool `protobuf:"varint,2,opt,name=open_in_new_tab,json=openInNewTab,proto3" json:"open_in_new_tab,omitempty"`
	// Types that are valid to be assigned to Type:
	//
	//	*Link_Url
	//	*Link_Dashboard_
	Type          isLink_Type `protobuf_oneof:"type"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Link) Reset() {
	*x = Link{}
	mi := &file_yandex_cloud_monitoring_v3_link_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_monitoring_v3_link_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_monitoring_v3_link_proto_rawDescGZIP(), []int{0}
}

func (x *Link) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Link) GetOpenInNewTab() bool {
	if x != nil {
		return x.OpenInNewTab
	}
	return false
}

func (x *Link) GetType() isLink_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *Link) GetUrl() string {
	if x != nil {
		if x, ok := x.Type.(*Link_Url); ok {
			return x.Url
		}
	}
	return ""
}

func (x *Link) GetDashboard() *Link_Dashboard {
	if x != nil {
		if x, ok := x.Type.(*Link_Dashboard_); ok {
			return x.Dashboard
		}
	}
	return nil
}

type isLink_Type interface {
	isLink_Type()
}

type Link_Url struct {
	// Url that can be a template with mustache expressions ({{expression}})
	// Variables available for template:
	//   - Dashboard time interval in diverse formats. Variables: __from:<format>, __to:<format>,
	//     where format is one of: iso, seconds or milliseconds.
	//   - Dashboard parameters
	Url string `protobuf:"bytes,3,opt,name=url,proto3,oneof"`
}

type Link_Dashboard_ struct {
	Dashboard *Link_Dashboard `protobuf:"bytes,4,opt,name=dashboard,proto3,oneof"`
}

func (*Link_Url) isLink_Type() {}

func (*Link_Dashboard_) isLink_Type() {}

type Link_Dashboard struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Container:
	//
	//	*Link_Dashboard_ProjectId
	//	*Link_Dashboard_FolderId
	Container      isLink_Dashboard_Container `protobuf_oneof:"container"`
	DashboardName  string                     `protobuf:"bytes,3,opt,name=dashboard_name,json=dashboardName,proto3" json:"dashboard_name,omitempty"`
	ApplyTimeRange bool                       `protobuf:"varint,4,opt,name=apply_time_range,json=applyTimeRange,proto3" json:"apply_time_range,omitempty"`
	// What parameter values to pass to dashboard when opening link
	// See Parametrization field in dashboard.proto
	Parameters    map[string]string `protobuf:"bytes,5,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Link_Dashboard) Reset() {
	*x = Link_Dashboard{}
	mi := &file_yandex_cloud_monitoring_v3_link_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Link_Dashboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link_Dashboard) ProtoMessage() {}

func (x *Link_Dashboard) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_monitoring_v3_link_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link_Dashboard.ProtoReflect.Descriptor instead.
func (*Link_Dashboard) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_monitoring_v3_link_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Link_Dashboard) GetContainer() isLink_Dashboard_Container {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *Link_Dashboard) GetProjectId() string {
	if x != nil {
		if x, ok := x.Container.(*Link_Dashboard_ProjectId); ok {
			return x.ProjectId
		}
	}
	return ""
}

func (x *Link_Dashboard) GetFolderId() string {
	if x != nil {
		if x, ok := x.Container.(*Link_Dashboard_FolderId); ok {
			return x.FolderId
		}
	}
	return ""
}

func (x *Link_Dashboard) GetDashboardName() string {
	if x != nil {
		return x.DashboardName
	}
	return ""
}

func (x *Link_Dashboard) GetApplyTimeRange() bool {
	if x != nil {
		return x.ApplyTimeRange
	}
	return false
}

func (x *Link_Dashboard) GetParameters() map[string]string {
	if x != nil {
		return x.Parameters
	}
	return nil
}

type isLink_Dashboard_Container interface {
	isLink_Dashboard_Container()
}

type Link_Dashboard_ProjectId struct {
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3,oneof"`
}

type Link_Dashboard_FolderId struct {
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3,oneof"`
}

func (*Link_Dashboard_ProjectId) isLink_Dashboard_Container() {}

func (*Link_Dashboard_FolderId) isLink_Dashboard_Container() {}

var File_yandex_cloud_monitoring_v3_link_proto protoreflect.FileDescriptor

const file_yandex_cloud_monitoring_v3_link_proto_rawDesc = "" +
	"\n" +
	"%yandex/cloud/monitoring/v3/link.proto\x12\x1ayandex.cloud.monitoring.v3\"\xf2\x03\n" +
	"\x04Link\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12%\n" +
	"\x0fopen_in_new_tab\x18\x02 \x01(\bR\fopenInNewTab\x12\x12\n" +
	"\x03url\x18\x03 \x01(\tH\x00R\x03url\x12J\n" +
	"\tdashboard\x18\x04 \x01(\v2*.yandex.cloud.monitoring.v3.Link.DashboardH\x00R\tdashboard\x1a\xc4\x02\n" +
	"\tDashboard\x12\x1f\n" +
	"\n" +
	"project_id\x18\x01 \x01(\tH\x00R\tprojectId\x12\x1d\n" +
	"\tfolder_id\x18\x02 \x01(\tH\x00R\bfolderId\x12%\n" +
	"\x0edashboard_name\x18\x03 \x01(\tR\rdashboardName\x12(\n" +
	"\x10apply_time_range\x18\x04 \x01(\bR\x0eapplyTimeRange\x12Z\n" +
	"\n" +
	"parameters\x18\x05 \x03(\v2:.yandex.cloud.monitoring.v3.Link.Dashboard.ParametersEntryR\n" +
	"parameters\x1a=\n" +
	"\x0fParametersEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01B\v\n" +
	"\tcontainerB\x06\n" +
	"\x04typeBk\n" +
	"\x1eyandex.cloud.api.monitoring.v3ZIgithub.com/yandex-cloud/go-genproto/yandex/cloud/monitoring/v3;monitoringb\x06proto3"

var (
	file_yandex_cloud_monitoring_v3_link_proto_rawDescOnce sync.Once
	file_yandex_cloud_monitoring_v3_link_proto_rawDescData []byte
)

func file_yandex_cloud_monitoring_v3_link_proto_rawDescGZIP() []byte {
	file_yandex_cloud_monitoring_v3_link_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_monitoring_v3_link_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_monitoring_v3_link_proto_rawDesc), len(file_yandex_cloud_monitoring_v3_link_proto_rawDesc)))
	})
	return file_yandex_cloud_monitoring_v3_link_proto_rawDescData
}

var file_yandex_cloud_monitoring_v3_link_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_monitoring_v3_link_proto_goTypes = []any{
	(*Link)(nil),           // 0: yandex.cloud.monitoring.v3.Link
	(*Link_Dashboard)(nil), // 1: yandex.cloud.monitoring.v3.Link.Dashboard
	nil,                    // 2: yandex.cloud.monitoring.v3.Link.Dashboard.ParametersEntry
}
var file_yandex_cloud_monitoring_v3_link_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.monitoring.v3.Link.dashboard:type_name -> yandex.cloud.monitoring.v3.Link.Dashboard
	2, // 1: yandex.cloud.monitoring.v3.Link.Dashboard.parameters:type_name -> yandex.cloud.monitoring.v3.Link.Dashboard.ParametersEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_monitoring_v3_link_proto_init() }
func file_yandex_cloud_monitoring_v3_link_proto_init() {
	if File_yandex_cloud_monitoring_v3_link_proto != nil {
		return
	}
	file_yandex_cloud_monitoring_v3_link_proto_msgTypes[0].OneofWrappers = []any{
		(*Link_Url)(nil),
		(*Link_Dashboard_)(nil),
	}
	file_yandex_cloud_monitoring_v3_link_proto_msgTypes[1].OneofWrappers = []any{
		(*Link_Dashboard_ProjectId)(nil),
		(*Link_Dashboard_FolderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_monitoring_v3_link_proto_rawDesc), len(file_yandex_cloud_monitoring_v3_link_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_monitoring_v3_link_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_monitoring_v3_link_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_monitoring_v3_link_proto_msgTypes,
	}.Build()
	File_yandex_cloud_monitoring_v3_link_proto = out.File
	file_yandex_cloud_monitoring_v3_link_proto_goTypes = nil
	file_yandex_cloud_monitoring_v3_link_proto_depIdxs = nil
}
