// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/speechsense/v1/analysis/text_classifiers.proto

package speechsense

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TextClassifiers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassificationResult []*ClassificationResult `protobuf:"bytes,1,rep,name=classification_result,json=classificationResult,proto3" json:"classification_result,omitempty"`
}

func (x *TextClassifiers) Reset() {
	*x = TextClassifiers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextClassifiers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextClassifiers) ProtoMessage() {}

func (x *TextClassifiers) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextClassifiers.ProtoReflect.Descriptor instead.
func (*TextClassifiers) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescGZIP(), []int{0}
}

func (x *TextClassifiers) GetClassificationResult() []*ClassificationResult {
	if x != nil {
		return x.ClassificationResult
	}
	return nil
}

type ClassificationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Classifier name
	Classifier string `protobuf:"bytes,1,opt,name=classifier,proto3" json:"classifier,omitempty"`
	// Classifier statistics
	ClassifierStatistics []*ClassifierStatistics `protobuf:"bytes,2,rep,name=classifier_statistics,json=classifierStatistics,proto3" json:"classifier_statistics,omitempty"`
}

func (x *ClassificationResult) Reset() {
	*x = ClassificationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassificationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassificationResult) ProtoMessage() {}

func (x *ClassificationResult) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassificationResult.ProtoReflect.Descriptor instead.
func (*ClassificationResult) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescGZIP(), []int{1}
}

func (x *ClassificationResult) GetClassifier() string {
	if x != nil {
		return x.Classifier
	}
	return ""
}

func (x *ClassificationResult) GetClassifierStatistics() []*ClassifierStatistics {
	if x != nil {
		return x.ClassifierStatistics
	}
	return nil
}

type ClassifierStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Channel number, null for whole talk
	ChannelNumber *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=channel_number,json=channelNumber,proto3" json:"channel_number,omitempty"`
	// classifier total count
	TotalCount int64 `protobuf:"varint,2,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Represents various histograms build on top of classifiers
	Histograms []*Histogram `protobuf:"bytes,3,rep,name=histograms,proto3" json:"histograms,omitempty"`
}

func (x *ClassifierStatistics) Reset() {
	*x = ClassifierStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifierStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifierStatistics) ProtoMessage() {}

func (x *ClassifierStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifierStatistics.ProtoReflect.Descriptor instead.
func (*ClassifierStatistics) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescGZIP(), []int{2}
}

func (x *ClassifierStatistics) GetChannelNumber() *wrapperspb.Int64Value {
	if x != nil {
		return x.ChannelNumber
	}
	return nil
}

func (x *ClassifierStatistics) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *ClassifierStatistics) GetHistograms() []*Histogram {
	if x != nil {
		return x.Histograms
	}
	return nil
}

type Histogram struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// histogram count values. For example:
	// if len(count_values) = 2, it means that histogram is 50/50,
	// if len(count_values) = 3 - [0] value represents first third, [1] - second third, [2] - last third, etc.
	CountValues []int64 `protobuf:"varint,1,rep,packed,name=count_values,json=countValues,proto3" json:"count_values,omitempty"`
}

func (x *Histogram) Reset() {
	*x = Histogram{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Histogram) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Histogram) ProtoMessage() {}

func (x *Histogram) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Histogram.ProtoReflect.Descriptor instead.
func (*Histogram) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescGZIP(), []int{3}
}

func (x *Histogram) GetCountValues() []int64 {
	if x != nil {
		return x.CountValues
	}
	return nil
}

var File_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto protoreflect.FileDescriptor

var file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x73, 0x69, 0x73, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x0f, 0x54, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x6f, 0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x14, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xa7, 0x01, 0x0a, 0x14, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x12, 0x6f, 0x0a, 0x15, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61,
	0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x14, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x14, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x42, 0x0a, 0x0e, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0d, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x4f, 0x0a, 0x0a, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x48, 0x69, 0x73, 0x74,
	0x6f, 0x67, 0x72, 0x61, 0x6d, 0x52, 0x0a, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d,
	0x73, 0x22, 0x2e, 0x0a, 0x09, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x67, 0x72, 0x61, 0x6d, 0x12, 0x21,
	0x0a, 0x0c, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x73, 0x42, 0x96, 0x01, 0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x14,
	0x54, 0x65, 0x78, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x3b, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescOnce sync.Once
	file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescData = file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDesc
)

func file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescGZIP() []byte {
	file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescData)
	})
	return file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDescData
}

var file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_goTypes = []interface{}{
	(*TextClassifiers)(nil),       // 0: yandex.cloud.speechsense.v1.analysis.TextClassifiers
	(*ClassificationResult)(nil),  // 1: yandex.cloud.speechsense.v1.analysis.ClassificationResult
	(*ClassifierStatistics)(nil),  // 2: yandex.cloud.speechsense.v1.analysis.ClassifierStatistics
	(*Histogram)(nil),             // 3: yandex.cloud.speechsense.v1.analysis.Histogram
	(*wrapperspb.Int64Value)(nil), // 4: google.protobuf.Int64Value
}
var file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.speechsense.v1.analysis.TextClassifiers.classification_result:type_name -> yandex.cloud.speechsense.v1.analysis.ClassificationResult
	2, // 1: yandex.cloud.speechsense.v1.analysis.ClassificationResult.classifier_statistics:type_name -> yandex.cloud.speechsense.v1.analysis.ClassifierStatistics
	4, // 2: yandex.cloud.speechsense.v1.analysis.ClassifierStatistics.channel_number:type_name -> google.protobuf.Int64Value
	3, // 3: yandex.cloud.speechsense.v1.analysis.ClassifierStatistics.histograms:type_name -> yandex.cloud.speechsense.v1.analysis.Histogram
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_init() }
func file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_init() {
	if File_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextClassifiers); i {
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
		file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassificationResult); i {
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
		file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifierStatistics); i {
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
		file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Histogram); i {
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
			RawDescriptor: file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_msgTypes,
	}.Build()
	File_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto = out.File
	file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_rawDesc = nil
	file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_goTypes = nil
	file_yandex_cloud_speechsense_v1_analysis_text_classifiers_proto_depIdxs = nil
}