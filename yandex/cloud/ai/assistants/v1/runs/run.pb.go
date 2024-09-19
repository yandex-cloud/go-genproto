// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
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
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RunState_RunStatus int32

const (
	RunState_RUN_STATUS_UNSPECIFIED RunState_RunStatus = 0
	RunState_PENDING                RunState_RunStatus = 1
	RunState_IN_PROGRESS            RunState_RunStatus = 2
	RunState_FAILED                 RunState_RunStatus = 3
	RunState_COMPLETED              RunState_RunStatus = 4
)

// Enum value maps for RunState_RunStatus.
var (
	RunState_RunStatus_name = map[int32]string{
		0: "RUN_STATUS_UNSPECIFIED",
		1: "PENDING",
		2: "IN_PROGRESS",
		3: "FAILED",
		4: "COMPLETED",
	}
	RunState_RunStatus_value = map[string]int32{
		"RUN_STATUS_UNSPECIFIED": 0,
		"PENDING":                1,
		"IN_PROGRESS":            2,
		"FAILED":                 3,
		"COMPLETED":              4,
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

type Run struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                            string                      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AssistantId                   string                      `protobuf:"bytes,2,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	ThreadId                      string                      `protobuf:"bytes,3,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	CreatedBy                     string                      `protobuf:"bytes,4,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt                     *timestamppb.Timestamp      `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Labels                        map[string]string           `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State                         *RunState                   `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty"`
	Usage                         *ContentUsage               `protobuf:"bytes,8,opt,name=usage,proto3" json:"usage,omitempty"`
	CustomPromptTruncationOptions *v1.PromptTruncationOptions `protobuf:"bytes,9,opt,name=custom_prompt_truncation_options,json=customPromptTruncationOptions,proto3" json:"custom_prompt_truncation_options,omitempty"`
	CustomCompletionOptions       *v1.CompletionOptions       `protobuf:"bytes,10,opt,name=custom_completion_options,json=customCompletionOptions,proto3" json:"custom_completion_options,omitempty"`
}

func (x *Run) Reset() {
	*x = Run{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Run) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Run) ProtoMessage() {}

func (x *Run) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

type RunState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status RunState_RunStatus `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.ai.assistants.v1.runs.RunState_RunStatus" json:"status,omitempty"`
	// Types that are assignable to StateData:
	//
	//	*RunState_Error
	//	*RunState_CompletedMessage
	StateData isRunState_StateData `protobuf_oneof:"StateData"`
}

func (x *RunState) Reset() {
	*x = RunState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunState) ProtoMessage() {}

func (x *RunState) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *RunState) GetStateData() isRunState_StateData {
	if m != nil {
		return m.StateData
	}
	return nil
}

func (x *RunState) GetError() *common.Error {
	if x, ok := x.GetStateData().(*RunState_Error); ok {
		return x.Error
	}
	return nil
}

func (x *RunState) GetCompletedMessage() *threads.Message {
	if x, ok := x.GetStateData().(*RunState_CompletedMessage); ok {
		return x.CompletedMessage
	}
	return nil
}

type isRunState_StateData interface {
	isRunState_StateData()
}

type RunState_Error struct {
	Error *common.Error `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type RunState_CompletedMessage struct {
	CompletedMessage *threads.Message `protobuf:"bytes,3,opt,name=completed_message,json=completedMessage,proto3,oneof"`
}

func (*RunState_Error) isRunState_StateData() {}

func (*RunState_CompletedMessage) isRunState_StateData() {}

type ContentUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PromptTokens     int64 `protobuf:"varint,1,opt,name=prompt_tokens,json=promptTokens,proto3" json:"prompt_tokens,omitempty"`
	CompletionTokens int64 `protobuf:"varint,2,opt,name=completion_tokens,json=completionTokens,proto3" json:"completion_tokens,omitempty"`
	TotalTokens      int64 `protobuf:"varint,3,opt,name=total_tokens,json=totalTokens,proto3" json:"total_tokens,omitempty"`
}

func (x *ContentUsage) Reset() {
	*x = ContentUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContentUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContentUsage) ProtoMessage() {}

func (x *ContentUsage) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
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

var file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x75, 0x6e, 0x73, 0x2f, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e,
	0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x75,
	0x6e, 0x73, 0x1a, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x61, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x05, 0x0a, 0x03, 0x52, 0x75,
	0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x75, 0x6e, 0x73,
	0x2e, 0x52, 0x75, 0x6e, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x42, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x75, 0x6e, 0x73, 0x2e, 0x52, 0x75, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x46, 0x0a, 0x05,
	0x75, 0x73, 0x61, 0x67, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x75, 0x6e, 0x73,
	0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x75,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x7f, 0x0a, 0x20, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x70,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x6d, 0x70, 0x74, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x1d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x50, 0x72,
	0x6f, 0x6d, 0x70, 0x74, 0x54, 0x72, 0x75, 0x6e, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x6c, 0x0a, 0x19, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x17, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xdf,
	0x02, 0x0a, 0x08, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x75, 0x6e, 0x73,
	0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x35, 0x0a, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x5d, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e,
	0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52,
	0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x60, 0x0a, 0x09, 0x52, 0x75, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a,
	0x0a, 0x16, 0x52, 0x55, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45,
	0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52,
	0x4f, 0x47, 0x52, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x04, 0x42, 0x0b, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x22, 0x83, 0x01, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x42, 0x75, 0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x75, 0x6e, 0x73,
	0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x75, 0x6e, 0x73, 0x3b, 0x72, 0x75, 0x6e, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescData = file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc
)

func file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDescData)
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
	(*common.Error)(nil),               // 8: yandex.cloud.ai.common.Error
	(*threads.Message)(nil),            // 9: yandex.cloud.ai.assistants.v1.threads.Message
}
var file_yandex_cloud_ai_assistants_v1_runs_run_proto_depIdxs = []int32{
	5, // 0: yandex.cloud.ai.assistants.v1.runs.Run.created_at:type_name -> google.protobuf.Timestamp
	4, // 1: yandex.cloud.ai.assistants.v1.runs.Run.labels:type_name -> yandex.cloud.ai.assistants.v1.runs.Run.LabelsEntry
	2, // 2: yandex.cloud.ai.assistants.v1.runs.Run.state:type_name -> yandex.cloud.ai.assistants.v1.runs.RunState
	3, // 3: yandex.cloud.ai.assistants.v1.runs.Run.usage:type_name -> yandex.cloud.ai.assistants.v1.runs.ContentUsage
	6, // 4: yandex.cloud.ai.assistants.v1.runs.Run.custom_prompt_truncation_options:type_name -> yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	7, // 5: yandex.cloud.ai.assistants.v1.runs.Run.custom_completion_options:type_name -> yandex.cloud.ai.assistants.v1.CompletionOptions
	0, // 6: yandex.cloud.ai.assistants.v1.runs.RunState.status:type_name -> yandex.cloud.ai.assistants.v1.runs.RunState.RunStatus
	8, // 7: yandex.cloud.ai.assistants.v1.runs.RunState.error:type_name -> yandex.cloud.ai.common.Error
	9, // 8: yandex.cloud.ai.assistants.v1.runs.RunState.completed_message:type_name -> yandex.cloud.ai.assistants.v1.threads.Message
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_runs_run_proto_init() }
func file_yandex_cloud_ai_assistants_v1_runs_run_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_runs_run_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Run); i {
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
		file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RunState); i {
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
		file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ContentUsage); i {
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
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_msgTypes[1].OneofWrappers = []any{
		(*RunState_Error)(nil),
		(*RunState_CompletedMessage)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc,
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
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_rawDesc = nil
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_runs_run_proto_depIdxs = nil
}
