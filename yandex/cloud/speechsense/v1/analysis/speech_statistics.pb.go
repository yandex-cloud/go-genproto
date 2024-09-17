// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/speechsense/v1/analysis/speech_statistics.proto

package speechsense

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

type SpeechStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Total simultaneous speech duration in seconds
	TotalSimultaneousSpeechDurationSeconds int64 `protobuf:"varint,1,opt,name=total_simultaneous_speech_duration_seconds,json=totalSimultaneousSpeechDurationSeconds,proto3" json:"total_simultaneous_speech_duration_seconds,omitempty"`
	// Total simultaneous speech duration in ms
	TotalSimultaneousSpeechDurationMs int64 `protobuf:"varint,2,opt,name=total_simultaneous_speech_duration_ms,json=totalSimultaneousSpeechDurationMs,proto3" json:"total_simultaneous_speech_duration_ms,omitempty"`
	// Simultaneous speech ratio within audio segment
	TotalSimultaneousSpeechRatio float64 `protobuf:"fixed64,3,opt,name=total_simultaneous_speech_ratio,json=totalSimultaneousSpeechRatio,proto3" json:"total_simultaneous_speech_ratio,omitempty"`
	// Descriptive statistics for simultaneous speech duration distribution
	SimultaneousSpeechDurationEstimation *DescriptiveStatistics `protobuf:"bytes,4,opt,name=simultaneous_speech_duration_estimation,json=simultaneousSpeechDurationEstimation,proto3" json:"simultaneous_speech_duration_estimation,omitempty"`
}

func (x *SpeechStatistics) Reset() {
	*x = SpeechStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeechStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeechStatistics) ProtoMessage() {}

func (x *SpeechStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeechStatistics.ProtoReflect.Descriptor instead.
func (*SpeechStatistics) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *SpeechStatistics) GetTotalSimultaneousSpeechDurationSeconds() int64 {
	if x != nil {
		return x.TotalSimultaneousSpeechDurationSeconds
	}
	return 0
}

func (x *SpeechStatistics) GetTotalSimultaneousSpeechDurationMs() int64 {
	if x != nil {
		return x.TotalSimultaneousSpeechDurationMs
	}
	return 0
}

func (x *SpeechStatistics) GetTotalSimultaneousSpeechRatio() float64 {
	if x != nil {
		return x.TotalSimultaneousSpeechRatio
	}
	return 0
}

func (x *SpeechStatistics) GetSimultaneousSpeechDurationEstimation() *DescriptiveStatistics {
	if x != nil {
		return x.SimultaneousSpeechDurationEstimation
	}
	return nil
}

var File_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto protoreflect.FileDescriptor

var file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65,
	0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x73, 0x69, 0x73, 0x1a, 0x3c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9c, 0x03, 0x0a, 0x10, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x5a, 0x0a, 0x2a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x26, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x53, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x73, 0x12, 0x50, 0x0a, 0x25, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68,
	0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x21, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61,
	0x6e, 0x65, 0x6f, 0x75, 0x73, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x45, 0x0a, 0x1f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73,
	0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x5f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75,
	0x73, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x92, 0x01, 0x0a,
	0x27, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x73,
	0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x24, 0x73, 0x69, 0x6d,
	0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x44,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x97, 0x01, 0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e,
	0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x15,
	0x53, 0x70, 0x65, 0x65, 0x63, 0x68, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65,
	0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x3b,
	0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescOnce sync.Once
	file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescData = file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDesc
)

func file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescGZIP() []byte {
	file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescData)
	})
	return file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDescData
}

var file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_goTypes = []any{
	(*SpeechStatistics)(nil),      // 0: yandex.cloud.speechsense.v1.analysis.SpeechStatistics
	(*DescriptiveStatistics)(nil), // 1: yandex.cloud.speechsense.v1.analysis.DescriptiveStatistics
}
var file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.speechsense.v1.analysis.SpeechStatistics.simultaneous_speech_duration_estimation:type_name -> yandex.cloud.speechsense.v1.analysis.DescriptiveStatistics
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_init() }
func file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_init() {
	if File_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto != nil {
		return
	}
	file_yandex_cloud_speechsense_v1_analysis_statistics_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SpeechStatistics); i {
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
			RawDescriptor: file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_msgTypes,
	}.Build()
	File_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto = out.File
	file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_rawDesc = nil
	file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_goTypes = nil
	file_yandex_cloud_speechsense_v1_analysis_speech_statistics_proto_depIdxs = nil
}
