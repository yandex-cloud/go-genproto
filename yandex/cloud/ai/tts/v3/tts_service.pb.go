// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/ai/tts/v3/tts_service.proto

package tts

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_yandex_cloud_ai_tts_v3_tts_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_tts_v3_tts_service_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x74, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x74, 0x73, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x70, 0x65, 0x65,
	0x63, 0x68, 0x6b, 0x69, 0x74, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x1a, 0x20, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x74, 0x74,
	0x73, 0x2f, 0x76, 0x33, 0x2f, 0x74, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x82,
	0x01, 0x0a, 0x0b, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x7a, 0x65, 0x72, 0x12, 0x73,
	0x0a, 0x12, 0x55, 0x74, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x79, 0x6e, 0x74, 0x68,
	0x65, 0x73, 0x69, 0x73, 0x12, 0x2b, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x6b, 0x69, 0x74,
	0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x74, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x63,
	0x65, 0x53, 0x79, 0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x73, 0x70, 0x65, 0x65, 0x63, 0x68, 0x6b, 0x69, 0x74, 0x2e, 0x74, 0x74,
	0x73, 0x2e, 0x76, 0x33, 0x2e, 0x55, 0x74, 0x74, 0x65, 0x72, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x79,
	0x6e, 0x74, 0x68, 0x65, 0x73, 0x69, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x30, 0x01, 0x42, 0x5c, 0x0a, 0x1a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x74, 0x73, 0x2e, 0x76,
	0x33, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x74, 0x74, 0x73, 0x2f, 0x76, 0x33, 0x3b, 0x74, 0x74,
	0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_yandex_cloud_ai_tts_v3_tts_service_proto_goTypes = []interface{}{
	(*UtteranceSynthesisRequest)(nil),  // 0: speechkit.tts.v3.UtteranceSynthesisRequest
	(*UtteranceSynthesisResponse)(nil), // 1: speechkit.tts.v3.UtteranceSynthesisResponse
}
var file_yandex_cloud_ai_tts_v3_tts_service_proto_depIdxs = []int32{
	0, // 0: speechkit.tts.v3.Synthesizer.UtteranceSynthesis:input_type -> speechkit.tts.v3.UtteranceSynthesisRequest
	1, // 1: speechkit.tts.v3.Synthesizer.UtteranceSynthesis:output_type -> speechkit.tts.v3.UtteranceSynthesisResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_tts_v3_tts_service_proto_init() }
func file_yandex_cloud_ai_tts_v3_tts_service_proto_init() {
	if File_yandex_cloud_ai_tts_v3_tts_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_tts_v3_tts_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_ai_tts_v3_tts_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ai_tts_v3_tts_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_tts_v3_tts_service_proto_depIdxs,
	}.Build()
	File_yandex_cloud_ai_tts_v3_tts_service_proto = out.File
	file_yandex_cloud_ai_tts_v3_tts_service_proto_rawDesc = nil
	file_yandex_cloud_ai_tts_v3_tts_service_proto_goTypes = nil
	file_yandex_cloud_ai_tts_v3_tts_service_proto_depIdxs = nil
}
