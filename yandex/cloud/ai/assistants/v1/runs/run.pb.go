// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/runs/run.proto

package runs

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1"
	threads "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1/threads"
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

// Enum representing the status of a run.
type RunState_RunStatus int32

const (
	// Default unspecified status.
	RunState_RUN_STATUS_UNSPECIFIED RunState_RunStatus = 0
	// Run has been created but has not started yet.
	RunState_PENDING RunState_RunStatus = 1
	// Run is currently in progress.
	RunState_IN_PROGRESS RunState_RunStatus = 2
	// Run has failed due to an error.
	RunState_FAILED RunState_RunStatus = 3
	// Run has completed successfully.
	RunState_COMPLETED RunState_RunStatus = 4
	// The run is waiting for tool calls to be executed and their results to be submitted.
	RunState_TOOL_CALLS RunState_RunStatus = 5
)

// Enum value maps for RunState_RunStatus.
var (
	RunState_RunStatus_name = map[int32]string{
		0: "RUN_STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "FAILED",
		4: "COMPLETED",
		5: "TOOL_CALLS",
	}
	RunState_RunStatus_value = map[string]int32{
		"RUN_STATUS_UNSPECIFIED": 0,
		"PENDING":                1,
		"IN_PROGRESS":            2,
		"FAILED":                 3,
		"COMPLETED":              4,
		"TOOL_CALLS":             5,
	}
)

func (x RunState_RunStatus) Enum() *RunState_RunStatus {
	p := new(RunState_RunStatus)
	*p = x
	return p
}

func (x RunState_RunStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RunState_RunStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_assistants_v1_runs_run_proto_enumTypes[0].Descriptor()
}

func (RunState_RunStatus) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_assistants_v1_runs_run_proto_enumTypes[0]
}

func (x RunState_RunStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RunState_RunStatus.Descriptor instead.
func (RunState_RunStatus) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescGZIP(), []int{1, 0}
}

// Represents a run of an assistant over a specific thread of messages.
type Run struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Unique identifier of the run.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Identifier for the assistant that is being run.
	AssistantId string `protobuf:"bytes,2,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	// Identifier for the thread of messages that this run is associated with.
	ThreadId string `protobuf:"bytes,3,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	// Identifier of the subject who created this run.
	CreatedBy string `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// Timestamp representing when the run was created.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Set of key-value pairs that can be used to organize and categorize the run.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Current state of the run, including its status and any associated data.
	State *RunState `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	// Information about the content usage during the run, such as the number of [tokens](/docs/foundation-models/concepts/yandexgpt/tokens) used by the completion model.
	Usage *ContentUsage `protobuf:"bytes,8,opt,name=usage,proto3" json:"usage,omitempty"`
	// Configuration options for truncating the prompt when the token count exceeds a specified limit.
	// If specified, these options will override the assistant's prompt truncation settings for this run.
	CustomPromptTruncationOptions *v1.PromptTruncationOptions `protobuf:"bytes,9,opt,name=custom_prompt_truncation_options,json=customPromptTruncationOptions,proto3" json:"custom_prompt_truncation_options,omitempty"`
	// Configuration options for completion generation.
	// If specified, these options will override the assistant's completion settings for this run.
	CustomCompletionOptions *v1.CompletionOptions `protobuf:"bytes,10,opt,name=custom_completion_options,json=customCompletionOptions,proto3" json:"custom_completion_options,omitempty"`
	// List of tools that are available for the assistant to use in this run.
	Tools []*v1.Tool `protobuf:"bytes,11,rep,name=tools,proto3" json:"tools,omitempty"`
	// Specifies the format of the model's response.
	CustomResponseFormat *v1.ResponseFormat `protobuf:"bytes,12,opt,name=custom_response_format,json=customResponseFormat,proto3" json:"custom_response_format,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *Run) Reset() {
	*x = Run{}
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Run.ProtoReflect.Descriptor instead.
func (*Run) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescGZIP(), []int{0}
}

func (x *Run) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Run) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

func (x *Run) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *Run) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Run) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Run) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Run) GetState() *RunState {
	if x != nil {
		return x.State
	}
	return nil
}

func (x *Run) GetUsage() *ContentUsage {
	if x != nil {
		return x.Usage
	}
	return nil
}

func (x *Run) GetCustomPromptTruncationOptions() *v1.PromptTruncationOptions {
	if x != nil {
		return x.CustomPromptTruncationOptions
	}
	return nil
}

func (x *Run) GetCustomCompletionOptions() *v1.CompletionOptions {
	if x != nil {
		return x.CustomCompletionOptions
	}
	return nil
}

func (x *Run) GetTools() []*v1.Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *Run) GetCustomResponseFormat() *v1.ResponseFormat {
	if x != nil {
		return x.CustomResponseFormat
	}
	return nil
}

// Represents the current state of a run.
type RunState struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Current status of a run.
	Status RunState_RunStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.ai.assistants.v1.runs.RunState_RunStatus" json:"status,omitempty"`
	// Oneof field to capture additional data depending on the state of a run.
	//
	// Types that are valid to be assigned to StateData:
	//
	//	*RunState_Error
	//	*RunState_CompletedMessage
	//	*RunState_ToolCallList
	StateData     isRunState_StateData `protobuf_oneof:"StateData"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RunState) Reset() {
	*x = RunState{}
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RunState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunState) ProtoMessage() {}

func (x *RunState) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunState.ProtoReflect.Descriptor instead.
func (*RunState) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescGZIP(), []int{1}
}

func (x *RunState) GetStatus() RunState_RunStatus {
	if x != nil {
		return x.Status
	}
	return RunState_RUN_STATUS_UNSPECIFIED
}

func (x *RunState) GetStateData() isRunState_StateData {
	if x != nil {
		return x.StateData
	}
	return nil
}

func (x *RunState) GetError() *common.Error {
	if x != nil {
		if x, ok := x.StateData.(*RunState_Error); ok {
			return x.Error
		}
	}
	return nil
}

func (x *RunState) GetCompletedMessage() *threads.Message {
	if x != nil {
		if x, ok := x.StateData.(*RunState_CompletedMessage); ok {
			return x.CompletedMessage
		}
	}
	return nil
}

func (x *RunState) GetToolCallList() *v1.ToolCallList {
	if x != nil {
		if x, ok := x.StateData.(*RunState_ToolCallList); ok {
			return x.ToolCallList
		}
	}
	return nil
}

type isRunState_StateData interface {
	isRunState_StateData()
}

type RunState_Error struct {
	// Error information if a run has failed.
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type RunState_CompletedMessage struct {
	// Final message generated by an assistant if a run has completed successfully.
	CompletedMessage *threads.Message `protobuf:"bytes,3,opt,name=completed_message,json=completedMessage,proto3,oneof"`
}

type RunState_ToolCallList struct {
	// A list of tool calls requested by the assistant.
	ToolCallList *v1.ToolCallList `protobuf:"bytes,4,opt,name=tool_call_list,json=toolCallList,proto3,oneof"`
}

func (*RunState_Error) isRunState_StateData() {}

func (*RunState_CompletedMessage) isRunState_StateData() {}

func (*RunState_ToolCallList) isRunState_StateData() {}

// Represents the content usage during a run, such as the number of [tokens](/docs/foundation-models/concepts/yandexgpt/tokens) used by the completion model.
type ContentUsage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of tokens used in the prompt.
	PromptTokens int64 `protobuf:"varint,1,opt,name=prompt_tokens,json=promptTokens,proto3" json:"prompt_tokens,omitempty"`
	// The number of tokens used in the completion response.
	CompletionTokens int64 `protobuf:"varint,2,opt,name=completion_tokens,json=completionTokens,proto3" json:"completion_tokens,omitempty"`
	// The total number of tokens used (prompt + completion).
	TotalTokens   int64 `protobuf:"varint,3,opt,name=total_tokens,json=totalTokens,proto3" json:"total_tokens,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContentUsage) Reset() {
	*x = ContentUsage{}
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContentUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentUsage) ProtoMessage() {}

func (x *ContentUsage) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContentUsage.ProtoReflect.Descriptor instead.
func (*ContentUsage) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescGZIP(), []int{2}
}

func (x *ContentUsage) GetPromptTokens() int64 {
	if x != nil {
		return x.PromptTokens
	}
	return 0
}

func (x *ContentUsage) GetCompletionTokens() int64 {
	if x != nil {
		return x.CompletionTokens
	}
	return 0
}

func (x *ContentUsage) GetTotalTokens() int64 {
	if x != nil {
		return x.TotalTokens
	}
	return 0
}

var File_yandex_cloud_ai_assistants_v1_runs_run_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc = "" +
	"\n" +
	",yandex/cloud/ai/assistants/v1/runs/run.proto\x12\"yandex.cloud.ai.assistants.v1.runs\x1a#yandex/cloud/ai/common/common.proto\x1a*yandex/cloud/ai/assistants/v1/common.proto\x1a3yandex/cloud/ai/assistants/v1/threads/message.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd2\x06\n" +
	"\x03Run\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12!\n" +
	"\fassistant_id\x18\x02 \x01(\tR\vassistantId\x12\x1b\n" +
	"\tthread_id\x18\x03 \x01(\tR\bthreadId\x12\x1d\n" +
	"\n" +
	"created_by\x18\x04 \x01(\tR\tcreatedBy\x129\n" +
	"\n" +
	"created_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12K\n" +
	"\x06labels\x18\x06 \x03(\v23.yandex.cloud.ai.assistants.v1.runs.Run.LabelsEntryR\x06labels\x12B\n" +
	"\x05state\x18\a \x01(\v2,.yandex.cloud.ai.assistants.v1.runs.RunStateR\x05state\x12F\n" +
	"\x05usage\x18\b \x01(\v20.yandex.cloud.ai.assistants.v1.runs.ContentUsageR\x05usage\x12\x7f\n" +
	" custom_prompt_truncation_options\x18\t \x01(\v26.yandex.cloud.ai.assistants.v1.PromptTruncationOptionsR\x1dcustomPromptTruncationOptions\x12l\n" +
	"\x19custom_completion_options\x18\n" +
	" \x01(\v20.yandex.cloud.ai.assistants.v1.CompletionOptionsR\x17customCompletionOptions\x129\n" +
	"\x05tools\x18\v \x03(\v2#.yandex.cloud.ai.assistants.v1.ToolR\x05tools\x12c\n" +
	"\x16custom_response_format\x18\f \x01(\v2-.yandex.cloud.ai.assistants.v1.ResponseFormatR\x14customResponseFormat\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"\xc4\x03\n" +
	"\bRunState\x12N\n" +
	"\x06status\x18\x01 \x01(\x0e26.yandex.cloud.ai.assistants.v1.runs.RunState.RunStatusR\x06status\x125\n" +
	"\x05error\x18\x02 \x01(\v2\x1d.yandex.cloud.ai.common.ErrorH\x00R\x05error\x12]\n" +
	"\x11completed_message\x18\x03 \x01(\v2..yandex.cloud.ai.assistants.v1.threads.MessageH\x00R\x10completedMessage\x12S\n" +
	"\x0etool_call_list\x18\x04 \x01(\v2+.yandex.cloud.ai.assistants.v1.ToolCallListH\x00R\ftoolCallList\"p\n" +
	"\tRunStatus\x12\x1a\n" +
	"\x16RUN_STATUS_UNSPECIFIED\x10\x00\x12\v\n" +
	"\aPENDING\x10\x01\x12\x0f\n" +
	"\vIN_PROGRESS\x10\x02\x12\n" +
	"\n" +
	"\x06FAILED\x10\x03\x12\r\n" +
	"\tCOMPLETED\x10\x04\x12\x0e\n" +
	"\n" +
	"TOOL_CALLS\x10\x05B\v\n" +
	"\tStateData\"\x83\x01\n" +
	"\fContentUsage\x12#\n" +
	"\rprompt_tokens\x18\x01 \x01(\x03R\fpromptTokens\x12+\n" +
	"\x11completion_tokens\x18\x02 \x01(\x03R\x10completionTokens\x12!\n" +
	"\ftotal_tokens\x18\x03 \x01(\x03R\vtotalTokensBu\n" +
	"&yandex.cloud.api.ai.assistants.v1.runsZKgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1/runs;runsb\x06proto3"

var (
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescData []byte
)

func file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_runs_run_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_ai_assistants_v1_runs_run_proto_goTypes = []any{
	(RunState_RunStatus)(0),            // 0: yandex.cloud.ai.assistants.v1.runs.RunState.RunStatus
	(*Run)(nil),                        // 1: yandex.cloud.ai.assistants.v1.runs.Run
	(*RunState)(nil),                   // 2: yandex.cloud.ai.assistants.v1.runs.RunState
	(*ContentUsage)(nil),               // 3: yandex.cloud.ai.assistants.v1.runs.ContentUsage
	nil,                                // 4: yandex.cloud.ai.assistants.v1.runs.Run.LabelsEntry
	(*timestamppb.Timestamp)(nil),      // 5: google.protobuf.Timestamp
	(*v1.PromptTruncationOptions)(nil), // 6: yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	(*v1.CompletionOptions)(nil),       // 7: yandex.cloud.ai.assistants.v1.CompletionOptions
	(*v1.Tool)(nil),                    // 8: yandex.cloud.ai.assistants.v1.Tool
	(*v1.ResponseFormat)(nil),          // 9: yandex.cloud.ai.assistants.v1.ResponseFormat
	(*common.Error)(nil),               // 10: yandex.cloud.ai.common.Error
	(*threads.Message)(nil),            // 11: yandex.cloud.ai.assistants.v1.threads.Message
	(*v1.ToolCallList)(nil),            // 12: yandex.cloud.ai.assistants.v1.ToolCallList
}
var file_yandex_cloud_ai_assistants_v1_runs_run_proto_depIdxs = []int32{
	5,  // 0: yandex.cloud.ai.assistants.v1.runs.Run.created_at:type_name -> google.protobuf.Timestamp
	4,  // 1: yandex.cloud.ai.assistants.v1.runs.Run.labels:type_name -> yandex.cloud.ai.assistants.v1.runs.Run.LabelsEntry
	2,  // 2: yandex.cloud.ai.assistants.v1.runs.Run.state:type_name -> yandex.cloud.ai.assistants.v1.runs.RunState
	3,  // 3: yandex.cloud.ai.assistants.v1.runs.Run.usage:type_name -> yandex.cloud.ai.assistants.v1.runs.ContentUsage
	6,  // 4: yandex.cloud.ai.assistants.v1.runs.Run.custom_prompt_truncation_options:type_name -> yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	7,  // 5: yandex.cloud.ai.assistants.v1.runs.Run.custom_completion_options:type_name -> yandex.cloud.ai.assistants.v1.CompletionOptions
	8,  // 6: yandex.cloud.ai.assistants.v1.runs.Run.tools:type_name -> yandex.cloud.ai.assistants.v1.Tool
	9,  // 7: yandex.cloud.ai.assistants.v1.runs.Run.custom_response_format:type_name -> yandex.cloud.ai.assistants.v1.ResponseFormat
	0,  // 8: yandex.cloud.ai.assistants.v1.runs.RunState.status:type_name -> yandex.cloud.ai.assistants.v1.runs.RunState.RunStatus
	10, // 9: yandex.cloud.ai.assistants.v1.runs.RunState.error:type_name -> yandex.cloud.ai.common.Error
	11, // 10: yandex.cloud.ai.assistants.v1.runs.RunState.completed_message:type_name -> yandex.cloud.ai.assistants.v1.threads.Message
	12, // 11: yandex.cloud.ai.assistants.v1.runs.RunState.tool_call_list:type_name -> yandex.cloud.ai.assistants.v1.ToolCallList
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_runs_run_proto_init() }
func file_yandex_cloud_ai_assistants_v1_runs_run_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_runs_run_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1].OneofWrappers = []any{
		(*RunState_Error)(nil),
		(*RunState_CompletedMessage)(nil),
		(*RunState_ToolCallList)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_runs_run_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_runs_run_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_assistants_v1_runs_run_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_runs_run_proto = out.File
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_depIdxs = nil
}
