// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/monitoring/v3/widget.proto

package monitoring

import (
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

// Widget.
type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. Widget layout position.
	Position *Widget_LayoutPosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// Required. Widget data.
	//
	// Types that are assignable to Widget:
	//
	//	*Widget_Text
	//	*Widget_Title
	//	*Widget_Chart
	//	*Widget_MultiSourceChart
	Widget isWidget_Widget `protobuf_oneof:"widget"`
	Links  []*LinkItem     `protobuf:"bytes,12,rep,name=links,proto3" json:"links,omitempty"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget.ProtoReflect.Descriptor instead.
func (*Widget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_monitoring_v3_widget_proto_rawDescGZIP(), []int{0}
}

func (x *Widget) GetPosition() *Widget_LayoutPosition {
	if x != nil {
		return x.Position
	}
	return nil
}

func (m *Widget) GetWidget() isWidget_Widget {
	if m != nil {
		return m.Widget
	}
	return nil
}

func (x *Widget) GetText() *TextWidget {
	if x, ok := x.GetWidget().(*Widget_Text); ok {
		return x.Text
	}
	return nil
}

func (x *Widget) GetTitle() *TitleWidget {
	if x, ok := x.GetWidget().(*Widget_Title); ok {
		return x.Title
	}
	return nil
}

func (x *Widget) GetChart() *ChartWidget {
	if x, ok := x.GetWidget().(*Widget_Chart); ok {
		return x.Chart
	}
	return nil
}

func (x *Widget) GetMultiSourceChart() *MultiSourceChartWidget {
	if x, ok := x.GetWidget().(*Widget_MultiSourceChart); ok {
		return x.MultiSourceChart
	}
	return nil
}

func (x *Widget) GetLinks() []*LinkItem {
	if x != nil {
		return x.Links
	}
	return nil
}

type isWidget_Widget interface {
	isWidget_Widget()
}

type Widget_Text struct {
	// Text widget.
	Text *TextWidget `protobuf:"bytes,2,opt,name=text,proto3,oneof"`
}

type Widget_Title struct {
	// Title widget.
	Title *TitleWidget `protobuf:"bytes,3,opt,name=title,proto3,oneof"`
}

type Widget_Chart struct {
	// Chart widget.
	Chart *ChartWidget `protobuf:"bytes,5,opt,name=chart,proto3,oneof"`
}

type Widget_MultiSourceChart struct {
	// Multi-source chart widget.
	MultiSourceChart *MultiSourceChartWidget `protobuf:"bytes,10,opt,name=multi_source_chart,json=multiSourceChart,proto3,oneof"`
}

func (*Widget_Text) isWidget_Widget() {}

func (*Widget_Title) isWidget_Widget() {}

func (*Widget_Chart) isWidget_Widget() {}

func (*Widget_MultiSourceChart) isWidget_Widget() {}

// Layout item for widget item positioning.
type Widget_LayoutPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Required. X-axis top-left corner coordinate.
	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Required. Y-axis top-left corner coordinate.
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	// Required. Weight.
	W int64 `protobuf:"varint,3,opt,name=w,proto3" json:"w,omitempty"`
	// Required. Height.
	H int64 `protobuf:"varint,4,opt,name=h,proto3" json:"h,omitempty"`
}

func (x *Widget_LayoutPosition) Reset() {
	*x = Widget_LayoutPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget_LayoutPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget_LayoutPosition) ProtoMessage() {}

func (x *Widget_LayoutPosition) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Widget_LayoutPosition.ProtoReflect.Descriptor instead.
func (*Widget_LayoutPosition) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_monitoring_v3_widget_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Widget_LayoutPosition) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Widget_LayoutPosition) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Widget_LayoutPosition) GetW() int64 {
	if x != nil {
		return x.W
	}
	return 0
}

func (x *Widget_LayoutPosition) GetH() int64 {
	if x != nil {
		return x.H
	}
	return 0
}

var File_yandex_cloud_monitoring_v3_widget_proto protoreflect.FileDescriptor

var file_yandex_cloud_monitoring_v3_widget_proto_rawDesc = []byte{
	0x0a, 0x27, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69,
	0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76,
	0x33, 0x2f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33,
	0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x3a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74,
	0x5f, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69,
	0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d, 0x04, 0x0a, 0x06, 0x57, 0x69,
	0x64, 0x67, 0x65, 0x74, 0x12, 0x4d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75,
	0x74, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x65, 0x78, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x12, 0x3f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x43,
	0x68, 0x61, 0x72, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68,
	0x61, 0x72, 0x74, 0x12, 0x62, 0x0a, 0x12, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x5f, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x4d, 0x75, 0x6c,
	0x74, 0x69, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x57, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x43, 0x68, 0x61, 0x72, 0x74, 0x12, 0x3a, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73,
	0x18, 0x0c, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x6f, 0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67,
	0x2e, 0x76, 0x33, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x1a, 0x48, 0x0a, 0x0e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x79, 0x12, 0x0c, 0x0a, 0x01, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x77, 0x12,
	0x0c, 0x0a, 0x01, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x68, 0x42, 0x08, 0x0a,
	0x06, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x4a, 0x04, 0x08,
	0x06, 0x10, 0x0a, 0x4a, 0x04, 0x08, 0x0b, 0x10, 0x0c, 0x42, 0x6b, 0x0a, 0x1e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x5a, 0x49, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x6f,
	0x6e, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x3b, 0x6d, 0x6f, 0x6e, 0x69,
	0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_monitoring_v3_widget_proto_rawDescOnce sync.Once
	file_yandex_cloud_monitoring_v3_widget_proto_rawDescData = file_yandex_cloud_monitoring_v3_widget_proto_rawDesc
)

func file_yandex_cloud_monitoring_v3_widget_proto_rawDescGZIP() []byte {
	file_yandex_cloud_monitoring_v3_widget_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_monitoring_v3_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_monitoring_v3_widget_proto_rawDescData)
	})
	return file_yandex_cloud_monitoring_v3_widget_proto_rawDescData
}

var file_yandex_cloud_monitoring_v3_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_monitoring_v3_widget_proto_goTypes = []any{
	(*Widget)(nil),                 // 0: yandex.cloud.monitoring.v3.Widget
	(*Widget_LayoutPosition)(nil),  // 1: yandex.cloud.monitoring.v3.Widget.LayoutPosition
	(*TextWidget)(nil),             // 2: yandex.cloud.monitoring.v3.TextWidget
	(*TitleWidget)(nil),            // 3: yandex.cloud.monitoring.v3.TitleWidget
	(*ChartWidget)(nil),            // 4: yandex.cloud.monitoring.v3.ChartWidget
	(*MultiSourceChartWidget)(nil), // 5: yandex.cloud.monitoring.v3.MultiSourceChartWidget
	(*LinkItem)(nil),               // 6: yandex.cloud.monitoring.v3.LinkItem
}
var file_yandex_cloud_monitoring_v3_widget_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.monitoring.v3.Widget.position:type_name -> yandex.cloud.monitoring.v3.Widget.LayoutPosition
	2, // 1: yandex.cloud.monitoring.v3.Widget.text:type_name -> yandex.cloud.monitoring.v3.TextWidget
	3, // 2: yandex.cloud.monitoring.v3.Widget.title:type_name -> yandex.cloud.monitoring.v3.TitleWidget
	4, // 3: yandex.cloud.monitoring.v3.Widget.chart:type_name -> yandex.cloud.monitoring.v3.ChartWidget
	5, // 4: yandex.cloud.monitoring.v3.Widget.multi_source_chart:type_name -> yandex.cloud.monitoring.v3.MultiSourceChartWidget
	6, // 5: yandex.cloud.monitoring.v3.Widget.links:type_name -> yandex.cloud.monitoring.v3.LinkItem
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_monitoring_v3_widget_proto_init() }
func file_yandex_cloud_monitoring_v3_widget_proto_init() {
	if File_yandex_cloud_monitoring_v3_widget_proto != nil {
		return
	}
	file_yandex_cloud_monitoring_v3_chart_widget_proto_init()
	file_yandex_cloud_monitoring_v3_text_widget_proto_init()
	file_yandex_cloud_monitoring_v3_title_widget_proto_init()
	file_yandex_cloud_monitoring_v3_multi_source_chart_widget_proto_init()
	file_yandex_cloud_monitoring_v3_link_item_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Widget); i {
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
		file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Widget_LayoutPosition); i {
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
	file_yandex_cloud_monitoring_v3_widget_proto_msgTypes[0].OneofWrappers = []any{
		(*Widget_Text)(nil),
		(*Widget_Title)(nil),
		(*Widget_Chart)(nil),
		(*Widget_MultiSourceChart)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_monitoring_v3_widget_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_monitoring_v3_widget_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_monitoring_v3_widget_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_monitoring_v3_widget_proto_msgTypes,
	}.Build()
	File_yandex_cloud_monitoring_v3_widget_proto = out.File
	file_yandex_cloud_monitoring_v3_widget_proto_rawDesc = nil
	file_yandex_cloud_monitoring_v3_widget_proto_goTypes = nil
	file_yandex_cloud_monitoring_v3_widget_proto_depIdxs = nil
}
