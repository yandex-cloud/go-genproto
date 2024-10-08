// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/api/v1/regression/widget.proto

package regression

import (
	report "github.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/report"
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

// Title size.
type TitleWidget_TitleSize int32

const (
	// Unspecified.
	TitleWidget_TITLE_SIZE_UNSPECIFIED TitleWidget_TitleSize = 0
	// Extra small.
	TitleWidget_TITLE_SIZE_XS TitleWidget_TitleSize = 1
	// Small.
	TitleWidget_TITLE_SIZE_S TitleWidget_TitleSize = 2
	// Medium.
	TitleWidget_TITLE_SIZE_M TitleWidget_TitleSize = 3
	// Large.
	TitleWidget_TITLE_SIZE_L TitleWidget_TitleSize = 4
)

// Enum value maps for TitleWidget_TitleSize.
var (
	TitleWidget_TitleSize_name = map[int32]string{
		0: "TITLE_SIZE_UNSPECIFIED",
		1: "TITLE_SIZE_XS",
		2: "TITLE_SIZE_S",
		3: "TITLE_SIZE_M",
		4: "TITLE_SIZE_L",
	}
	TitleWidget_TitleSize_value = map[string]int32{
		"TITLE_SIZE_UNSPECIFIED": 0,
		"TITLE_SIZE_XS":          1,
		"TITLE_SIZE_S":           2,
		"TITLE_SIZE_M":           3,
		"TITLE_SIZE_L":           4,
	}
)

func (x TitleWidget_TitleSize) Enum() *TitleWidget_TitleSize {
	p := new(TitleWidget_TitleSize)
	*p = x
	return p
}

func (x TitleWidget_TitleSize) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TitleWidget_TitleSize) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_enumTypes[0].Descriptor()
}

func (TitleWidget_TitleSize) Type() protoreflect.EnumType {
	return &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_enumTypes[0]
}

func (x TitleWidget_TitleSize) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TitleWidget_TitleSize.Descriptor instead.
func (TitleWidget_TitleSize) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP(), []int{3, 0}
}

// Regression dashboard widget.
type Widget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Widget position.
	Position *Widget_LayoutPosition `protobuf:"bytes,1,opt,name=position,proto3" json:"position,omitempty"`
	// Types that are assignable to Widget:
	//
	//	*Widget_Chart
	//	*Widget_Text
	//	*Widget_Title
	Widget isWidget_Widget `protobuf_oneof:"widget"`
}

func (x *Widget) Reset() {
	*x = Widget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget) ProtoMessage() {}

func (x *Widget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[0]
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
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP(), []int{0}
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

func (x *Widget) GetChart() *ChartWidget {
	if x, ok := x.GetWidget().(*Widget_Chart); ok {
		return x.Chart
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

type isWidget_Widget interface {
	isWidget_Widget()
}

type Widget_Chart struct {
	// Chart widget.
	Chart *ChartWidget `protobuf:"bytes,2,opt,name=chart,proto3,oneof"`
}

type Widget_Text struct {
	// Text widget.
	Text *TextWidget `protobuf:"bytes,3,opt,name=text,proto3,oneof"`
}

type Widget_Title struct {
	// Title widget.
	Title *TitleWidget `protobuf:"bytes,4,opt,name=title,proto3,oneof"`
}

func (*Widget_Chart) isWidget_Widget() {}

func (*Widget_Text) isWidget_Widget() {}

func (*Widget_Title) isWidget_Widget() {}

// Regression chart.
type ChartWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the chart.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the chart.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the chart.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Test filter selector to show KPI values for.
	FilterStr string `protobuf:"bytes,4,opt,name=filter_str,json=filterStr,proto3" json:"filter_str,omitempty"`
	// Test case to show KPI values for.
	TestCase string `protobuf:"bytes,5,opt,name=test_case,json=testCase,proto3" json:"test_case,omitempty"`
	// KPIs to show.
	Kpis []*report.Kpi `protobuf:"bytes,6,rep,name=kpis,proto3" json:"kpis,omitempty"`
}

func (x *ChartWidget) Reset() {
	*x = ChartWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChartWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChartWidget) ProtoMessage() {}

func (x *ChartWidget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChartWidget.ProtoReflect.Descriptor instead.
func (*ChartWidget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP(), []int{1}
}

func (x *ChartWidget) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChartWidget) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChartWidget) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ChartWidget) GetFilterStr() string {
	if x != nil {
		return x.FilterStr
	}
	return ""
}

func (x *ChartWidget) GetTestCase() string {
	if x != nil {
		return x.TestCase
	}
	return ""
}

func (x *ChartWidget) GetKpis() []*report.Kpi {
	if x != nil {
		return x.Kpis
	}
	return nil
}

// Text widget.
type TextWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Text string.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextWidget) Reset() {
	*x = TextWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextWidget) ProtoMessage() {}

func (x *TextWidget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextWidget.ProtoReflect.Descriptor instead.
func (*TextWidget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP(), []int{2}
}

func (x *TextWidget) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

// Title widget.
type TitleWidget struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Title string.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// Title size.
	Size TitleWidget_TitleSize `protobuf:"varint,2,opt,name=size,proto3,enum=yandex.cloud.loadtesting.api.v1.regression.TitleWidget_TitleSize" json:"size,omitempty"`
}

func (x *TitleWidget) Reset() {
	*x = TitleWidget{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TitleWidget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TitleWidget) ProtoMessage() {}

func (x *TitleWidget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TitleWidget.ProtoReflect.Descriptor instead.
func (*TitleWidget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP(), []int{3}
}

func (x *TitleWidget) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *TitleWidget) GetSize() TitleWidget_TitleSize {
	if x != nil {
		return x.Size
	}
	return TitleWidget_TITLE_SIZE_UNSPECIFIED
}

// Widget position.
type Widget_LayoutPosition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// X.
	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y.
	Y int64 `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	// Width.
	Width int64 `protobuf:"varint,3,opt,name=width,proto3" json:"width,omitempty"`
	// Height.
	Height int64 `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *Widget_LayoutPosition) Reset() {
	*x = Widget_LayoutPosition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Widget_LayoutPosition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Widget_LayoutPosition) ProtoMessage() {}

func (x *Widget_LayoutPosition) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[4]
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
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP(), []int{0, 0}
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

func (x *Widget_LayoutPosition) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Widget_LayoutPosition) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

var File_yandex_cloud_loadtesting_api_v1_regression_widget_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDesc = []byte{
	0x0a, 0x37, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x77, 0x69, 0x64,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x30, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x6b, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x03, 0x0a, 0x06, 0x57, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x12, 0x5d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x4f, 0x0a, 0x05, 0x63, 0x68, 0x61, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x68,
	0x61, 0x72, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x61,
	0x72, 0x74, 0x12, 0x4c, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x65,
	0x78, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x4f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x74,
	0x6c, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x48, 0x00, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x1a, 0x5a, 0x0a, 0x0e, 0x4c, 0x61, 0x79, 0x6f, 0x75, 0x74, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01,
	0x78, 0x12, 0x0c, 0x0a, 0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x42, 0x08, 0x0a,
	0x06, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x72,
	0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x53, 0x74, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x74, 0x65, 0x73, 0x74, 0x5f, 0x63, 0x61, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x74, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x04, 0x6b, 0x70, 0x69,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x4b, 0x70, 0x69, 0x52, 0x04, 0x6b, 0x70, 0x69, 0x73, 0x22, 0x20, 0x0a, 0x0a, 0x54, 0x65,
	0x78, 0x74, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0xea, 0x01, 0x0a,
	0x0b, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x12, 0x55, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x41,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f,
	0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x70, 0x0a, 0x09, 0x54, 0x69, 0x74, 0x6c, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x49,
	0x5a, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x58,
	0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x53, 0x49, 0x5a,
	0x45, 0x5f, 0x53, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x54, 0x4c, 0x45, 0x5f, 0x53,
	0x49, 0x5a, 0x45, 0x5f, 0x4d, 0x10, 0x03, 0x12, 0x10, 0x0a, 0x0c, 0x54, 0x49, 0x54, 0x4c, 0x45,
	0x5f, 0x53, 0x49, 0x5a, 0x45, 0x5f, 0x4c, 0x10, 0x04, 0x42, 0x8b, 0x01, 0x0a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5a, 0x59, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x72, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x3b, 0x72, 0x65, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescData = file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDesc
)

func file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_goTypes = []any{
	(TitleWidget_TitleSize)(0),    // 0: yandex.cloud.loadtesting.api.v1.regression.TitleWidget.TitleSize
	(*Widget)(nil),                // 1: yandex.cloud.loadtesting.api.v1.regression.Widget
	(*ChartWidget)(nil),           // 2: yandex.cloud.loadtesting.api.v1.regression.ChartWidget
	(*TextWidget)(nil),            // 3: yandex.cloud.loadtesting.api.v1.regression.TextWidget
	(*TitleWidget)(nil),           // 4: yandex.cloud.loadtesting.api.v1.regression.TitleWidget
	(*Widget_LayoutPosition)(nil), // 5: yandex.cloud.loadtesting.api.v1.regression.Widget.LayoutPosition
	(*report.Kpi)(nil),            // 6: yandex.cloud.loadtesting.api.v1.report.Kpi
}
var file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.loadtesting.api.v1.regression.Widget.position:type_name -> yandex.cloud.loadtesting.api.v1.regression.Widget.LayoutPosition
	2, // 1: yandex.cloud.loadtesting.api.v1.regression.Widget.chart:type_name -> yandex.cloud.loadtesting.api.v1.regression.ChartWidget
	3, // 2: yandex.cloud.loadtesting.api.v1.regression.Widget.text:type_name -> yandex.cloud.loadtesting.api.v1.regression.TextWidget
	4, // 3: yandex.cloud.loadtesting.api.v1.regression.Widget.title:type_name -> yandex.cloud.loadtesting.api.v1.regression.TitleWidget
	6, // 4: yandex.cloud.loadtesting.api.v1.regression.ChartWidget.kpis:type_name -> yandex.cloud.loadtesting.api.v1.report.Kpi
	0, // 5: yandex.cloud.loadtesting.api.v1.regression.TitleWidget.size:type_name -> yandex.cloud.loadtesting.api.v1.regression.TitleWidget.TitleSize
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_regression_widget_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ChartWidget); i {
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
		file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*TextWidget); i {
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
		file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TitleWidget); i {
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
		file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[4].Exporter = func(v any, i int) any {
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
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes[0].OneofWrappers = []any{
		(*Widget_Chart)(nil),
		(*Widget_Text)(nil),
		(*Widget_Title)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_regression_widget_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_regression_widget_proto_depIdxs = nil
}
