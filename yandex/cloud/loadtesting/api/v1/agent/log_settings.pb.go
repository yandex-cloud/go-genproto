// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/api/v1/agent/log_settings.proto

package agent

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

type LogSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Id of Yandex Cloud log group to upload agent logs to
	CloudLogGroupId string `protobuf:"bytes,1,opt,name=cloud_log_group_id,json=cloudLogGroupId,proto3" json:"cloud_log_group_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *LogSettings) Reset() {
	*x = LogSettings{}
	mi := &file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogSettings) ProtoMessage() {}

func (x *LogSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogSettings.ProtoReflect.Descriptor instead.
func (*LogSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescGZIP(), []int{0}
}

func (x *LogSettings) GetCloudLogGroupId() string {
	if x != nil {
		return x.CloudLogGroupId
	}
	return ""
}

var File_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto protoreflect.FileDescriptor

const file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDesc = "" +
	"\n" +
	"8yandex/cloud/loadtesting/api/v1/agent/log_settings.proto\x12%yandex.cloud.loadtesting.api.v1.agent\x1a\x1dyandex/cloud/validation.proto\"D\n" +
	"\vLogSettings\x125\n" +
	"\x12cloud_log_group_id\x18\x01 \x01(\tB\b\x8a\xc81\x04<=50R\x0fcloudLogGroupIdB|\n" +
	")yandex.cloud.api.loadtesting.api.v1.agentZOgithub.com/yandex-cloud/go-genproto/yandex/cloud/loadtesting/api/v1/agent;agentb\x06proto3"

var (
	file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescData []byte
)

func file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDesc), len(file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDesc)))
	})
	return file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_goTypes = []any{
	(*LogSettings)(nil), // 0: yandex.cloud.loadtesting.api.v1.agent.LogSettings
}
var file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDesc), len(file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_agent_log_settings_proto_depIdxs = nil
}
