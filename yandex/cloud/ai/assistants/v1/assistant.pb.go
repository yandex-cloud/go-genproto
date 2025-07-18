// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/assistant.proto

package assistants

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/common"
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

// Assistant represents an AI assistant configuration with various settings and metadata.
type Assistant struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique identifier of the assistant.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the assistant belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the assistant.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the assistant.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Identifier of the subject who created this assistant.
	CreatedBy string `protobuf:"bytes,5,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// Timestamp representing when the assistant was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Identifier of the subject who last updated this assistant.
	UpdatedBy string `protobuf:"bytes,7,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	// Timestamp representing the last time this assistant was updated.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Configuration for the expiration of the assistant, defining when and how the assistant will expire.
	ExpirationConfig *common.ExpirationConfig `protobuf:"bytes,9,opt,name=expiration_config,json=expirationConfig,proto3" json:"expiration_config,omitempty"`
	// Timestamp representing when the assistant will expire.
	ExpiresAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	// Set of key-value pairs that can be used to organize and categorize the assistant.
	Labels map[string]string `protobuf:"bytes,11,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The [ID of the model](/docs/foundation-models/concepts/yandexgpt/models) to be used for completion generation.
	ModelUri string `protobuf:"bytes,12,opt,name=model_uri,json=modelUri,proto3" json:"model_uri,omitempty"`
	// Instructions or guidelines that the assistant should follow when generating responses or performing tasks.
	// These instructions can help guide the assistant's behavior and responses.
	Instruction string `protobuf:"bytes,13,opt,name=instruction,proto3" json:"instruction,omitempty"`
	// Configuration options for truncating the prompt when the token count exceeds a specified limit.
	PromptTruncationOptions *PromptTruncationOptions `protobuf:"bytes,14,opt,name=prompt_truncation_options,json=promptTruncationOptions,proto3" json:"prompt_truncation_options,omitempty"`
	// Configuration options for completion generation.
	CompletionOptions *CompletionOptions `protobuf:"bytes,15,opt,name=completion_options,json=completionOptions,proto3" json:"completion_options,omitempty"`
	// List of tools that the assistant can use to perform additional tasks.
	// One example is the SearchIndexTool, which is used for Retrieval-Augmented Generation (RAG).
	Tools []*Tool `protobuf:"bytes,16,rep,name=tools,proto3" json:"tools,omitempty"`
	// Specifies the format of the model's response.
	ResponseFormat *ResponseFormat `protobuf:"bytes,17,opt,name=response_format,json=responseFormat,proto3" json:"response_format,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *Assistant) Reset() {
	*x = Assistant{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Assistant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Assistant) ProtoMessage() {}

func (x *Assistant) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Assistant.ProtoReflect.Descriptor instead.
func (*Assistant) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescGZIP(), []int{0}
}

func (x *Assistant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Assistant) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Assistant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Assistant) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Assistant) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Assistant) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Assistant) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *Assistant) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Assistant) GetExpirationConfig() *common.ExpirationConfig {
	if x != nil {
		return x.ExpirationConfig
	}
	return nil
}

func (x *Assistant) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

func (x *Assistant) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Assistant) GetModelUri() string {
	if x != nil {
		return x.ModelUri
	}
	return ""
}

func (x *Assistant) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *Assistant) GetPromptTruncationOptions() *PromptTruncationOptions {
	if x != nil {
		return x.PromptTruncationOptions
	}
	return nil
}

func (x *Assistant) GetCompletionOptions() *CompletionOptions {
	if x != nil {
		return x.CompletionOptions
	}
	return nil
}

func (x *Assistant) GetTools() []*Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *Assistant) GetResponseFormat() *ResponseFormat {
	if x != nil {
		return x.ResponseFormat
	}
	return nil
}

var File_yandex_cloud_ai_assistants_v1_assistant_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDesc = "" +
	"\n" +
	"-yandex/cloud/ai/assistants/v1/assistant.proto\x12\x1dyandex.cloud.ai.assistants.v1\x1a#yandex/cloud/ai/common/common.proto\x1a*yandex/cloud/ai/assistants/v1/common.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe4\a\n" +
	"\tAssistant\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12\x1d\n" +
	"\n" +
	"created_by\x18\x05 \x01(\tR\tcreatedBy\x129\n" +
	"\n" +
	"created_at\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x1d\n" +
	"\n" +
	"updated_by\x18\a \x01(\tR\tupdatedBy\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\x12U\n" +
	"\x11expiration_config\x18\t \x01(\v2(.yandex.cloud.ai.common.ExpirationConfigR\x10expirationConfig\x129\n" +
	"\n" +
	"expires_at\x18\n" +
	" \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\x12L\n" +
	"\x06labels\x18\v \x03(\v24.yandex.cloud.ai.assistants.v1.Assistant.LabelsEntryR\x06labels\x12\x1b\n" +
	"\tmodel_uri\x18\f \x01(\tR\bmodelUri\x12 \n" +
	"\vinstruction\x18\r \x01(\tR\vinstruction\x12r\n" +
	"\x19prompt_truncation_options\x18\x0e \x01(\v26.yandex.cloud.ai.assistants.v1.PromptTruncationOptionsR\x17promptTruncationOptions\x12_\n" +
	"\x12completion_options\x18\x0f \x01(\v20.yandex.cloud.ai.assistants.v1.CompletionOptionsR\x11completionOptions\x129\n" +
	"\x05tools\x18\x10 \x03(\v2#.yandex.cloud.ai.assistants.v1.ToolR\x05tools\x12V\n" +
	"\x0fresponse_format\x18\x11 \x01(\v2-.yandex.cloud.ai.assistants.v1.ResponseFormatR\x0eresponseFormat\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01Bq\n" +
	"!yandex.cloud.api.ai.assistants.v1ZLgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1;assistantsb\x06proto3"

var (
	file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescData []byte
)

func file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_assistant_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_assistants_v1_assistant_proto_goTypes = []any{
	(*Assistant)(nil),               // 0: yandex.cloud.ai.assistants.v1.Assistant
	nil,                             // 1: yandex.cloud.ai.assistants.v1.Assistant.LabelsEntry
	(*timestamppb.Timestamp)(nil),   // 2: google.protobuf.Timestamp
	(*common.ExpirationConfig)(nil), // 3: yandex.cloud.ai.common.ExpirationConfig
	(*PromptTruncationOptions)(nil), // 4: yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	(*CompletionOptions)(nil),       // 5: yandex.cloud.ai.assistants.v1.CompletionOptions
	(*Tool)(nil),                    // 6: yandex.cloud.ai.assistants.v1.Tool
	(*ResponseFormat)(nil),          // 7: yandex.cloud.ai.assistants.v1.ResponseFormat
}
var file_yandex_cloud_ai_assistants_v1_assistant_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.ai.assistants.v1.Assistant.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.ai.assistants.v1.Assistant.updated_at:type_name -> google.protobuf.Timestamp
	3, // 2: yandex.cloud.ai.assistants.v1.Assistant.expiration_config:type_name -> yandex.cloud.ai.common.ExpirationConfig
	2, // 3: yandex.cloud.ai.assistants.v1.Assistant.expires_at:type_name -> google.protobuf.Timestamp
	1, // 4: yandex.cloud.ai.assistants.v1.Assistant.labels:type_name -> yandex.cloud.ai.assistants.v1.Assistant.LabelsEntry
	4, // 5: yandex.cloud.ai.assistants.v1.Assistant.prompt_truncation_options:type_name -> yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	5, // 6: yandex.cloud.ai.assistants.v1.Assistant.completion_options:type_name -> yandex.cloud.ai.assistants.v1.CompletionOptions
	6, // 7: yandex.cloud.ai.assistants.v1.Assistant.tools:type_name -> yandex.cloud.ai.assistants.v1.Tool
	7, // 8: yandex.cloud.ai.assistants.v1.Assistant.response_format:type_name -> yandex.cloud.ai.assistants.v1.ResponseFormat
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_assistant_proto_init() }
func file_yandex_cloud_ai_assistants_v1_assistant_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_assistant_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_assistant_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_assistant_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_assistant_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_assistant_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_assistant_proto = out.File
	file_yandex_cloud_ai_assistants_v1_assistant_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_assistant_proto_depIdxs = nil
}
