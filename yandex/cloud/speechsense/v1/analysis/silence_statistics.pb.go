// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/speechsense/v1/analysis/silence_statistics.proto

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

type SilenceStatistics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSimultaneousSilenceDurationMs int64 `protobuf:"varint,1,opt,name=total_simultaneous_silence_duration_ms,json=totalSimultaneousSilenceDurationMs,proto3" json:"total_simultaneous_silence_duration_ms,omitempty"`
	// Simultaneous silence ratio within audio segment
	TotalSimultaneousSilenceRatio float64 `protobuf:"fixed64,2,opt,name=total_simultaneous_silence_ratio,json=totalSimultaneousSilenceRatio,proto3" json:"total_simultaneous_silence_ratio,omitempty"`
	// Descriptive statistics for simultaneous silence duration distribution
	SimultaneousSilenceDurationEstimation   *DescriptiveStatistics `protobuf:"bytes,3,opt,name=simultaneous_silence_duration_estimation,json=simultaneousSilenceDurationEstimation,proto3" json:"simultaneous_silence_duration_estimation,omitempty"`
	TotalSimultaneousSilenceDurationSeconds int64                  `protobuf:"varint,4,opt,name=total_simultaneous_silence_duration_seconds,json=totalSimultaneousSilenceDurationSeconds,proto3" json:"total_simultaneous_silence_duration_seconds,omitempty"`
}

func (x *SilenceStatistics) Reset() {
	*x = SilenceStatistics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SilenceStatistics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SilenceStatistics) ProtoMessage() {}

func (x *SilenceStatistics) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SilenceStatistics.ProtoReflect.Descriptor instead.
func (*SilenceStatistics) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescGZIP(), []int{0}
}

func (x *SilenceStatistics) GetTotalSimultaneousSilenceDurationMs() int64 {
	if x != nil {
		return x.TotalSimultaneousSilenceDurationMs
	}
	return 0
}

func (x *SilenceStatistics) GetTotalSimultaneousSilenceRatio() float64 {
	if x != nil {
		return x.TotalSimultaneousSilenceRatio
	}
	return 0
}

func (x *SilenceStatistics) GetSimultaneousSilenceDurationEstimation() *DescriptiveStatistics {
	if x != nil {
		return x.SimultaneousSilenceDurationEstimation
	}
	return nil
}

func (x *SilenceStatistics) GetTotalSimultaneousSilenceDurationSeconds() int64 {
	if x != nil {
		return x.TotalSimultaneousSilenceDurationSeconds
	}
	return 0
}

var File_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto protoreflect.FileDescriptor

var file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70,
	0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61,
	0x6c, 0x79, 0x73, 0x69, 0x73, 0x1a, 0x3c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f,
	0x76, 0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa5, 0x03, 0x0a, 0x11, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x52, 0x0a, 0x26, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f,
	0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x53, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x6c, 0x65,
	0x6e, 0x63, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x73, 0x12, 0x47, 0x0a,
	0x20, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65,
	0x6f, 0x75, 0x73, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x1d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69,
	0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x52, 0x61, 0x74, 0x69, 0x6f, 0x12, 0x94, 0x01, 0x0a, 0x28, 0x73, 0x69, 0x6d, 0x75, 0x6c,
	0x74, 0x61, 0x6e, 0x65, 0x6f, 0x75, 0x73, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f,
	0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73,
	0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73,
	0x2e, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x76, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x25, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e,
	0x65, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5c, 0x0a,
	0x2b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61, 0x6e, 0x65,
	0x6f, 0x75, 0x73, 0x5f, 0x73, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x27, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x6d, 0x75, 0x6c, 0x74, 0x61,
	0x6e, 0x65, 0x6f, 0x75, 0x73, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63, 0x65, 0x44, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0x98, 0x01, 0x0a, 0x28,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x42, 0x16, 0x53, 0x69, 0x6c, 0x65, 0x6e, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x73, 0x69, 0x73, 0x3b, 0x73, 0x70, 0x65, 0x65, 0x63,
	0x68, 0x73, 0x65, 0x6e, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescOnce sync.Once
	file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescData = file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDesc
)

func file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescGZIP() []byte {
	file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescData)
	})
	return file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDescData
}

var file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_goTypes = []interface{}{
	(*SilenceStatistics)(nil),     // 0: yandex.cloud.speechsense.v1.analysis.SilenceStatistics
	(*DescriptiveStatistics)(nil), // 1: yandex.cloud.speechsense.v1.analysis.DescriptiveStatistics
}
var file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.speechsense.v1.analysis.SilenceStatistics.simultaneous_silence_duration_estimation:type_name -> yandex.cloud.speechsense.v1.analysis.DescriptiveStatistics
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_init() }
func file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_init() {
	if File_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto != nil {
		return
	}
	file_yandex_cloud_speechsense_v1_analysis_statistics_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SilenceStatistics); i {
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
			RawDescriptor: file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_msgTypes,
	}.Build()
	File_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto = out.File
	file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_rawDesc = nil
	file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_goTypes = nil
	file_yandex_cloud_speechsense_v1_analysis_silence_statistics_proto_depIdxs = nil
}