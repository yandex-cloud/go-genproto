// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/serverless/workflows/v1/execution.proto

package workflows

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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

type Execution_Status int32

const (
	Execution_STATUS_UNSPECIFIED Execution_Status = 0
	// Workflow execution is being queued.
	Execution_QUEUED Execution_Status = 1
	// Workflow execution is running.
	Execution_RUNNING Execution_Status = 2
	// Workflow execution is being paused.
	Execution_PAUSED Execution_Status = 3
	// Workflow execution is stopped.
	Execution_STOPPED Execution_Status = 4
	// Workflow execution is failed.
	Execution_FAILED Execution_Status = 5
	// Workflow execution is finished.
	Execution_FINISHED Execution_Status = 6
)

// Enum value maps for Execution_Status.
var (
	Execution_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "QUEUED",
		2: "RUNNING",
		3: "PAUSED",
		4: "STOPPED",
		5: "FAILED",
		6: "FINISHED",
	}
	Execution_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"QUEUED":             1,
		"RUNNING":            2,
		"PAUSED":             3,
		"STOPPED":            4,
		"FAILED":             5,
		"FINISHED":           6,
	}
)

func (x Execution_Status) Enum() *Execution_Status {
	p := new(Execution_Status)
	*p = x
	return p
}

func (x Execution_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Execution_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_enumTypes[0].Descriptor()
}

func (Execution_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_serverless_workflows_v1_execution_proto_enumTypes[0]
}

func (x Execution_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Execution_Status.Descriptor instead.
func (Execution_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP(), []int{0, 0}
}

type Execution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Workflow execution. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the Workflow.
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Input data for the Workflow execution.
	Input *ExecutionInput `protobuf:"bytes,4,opt,name=input,proto3" json:"input,omitempty"`
	// Result of the Workflow execution.
	Result *ExecutionResult `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	// Error details, in case Workflow execution failed.
	Error *ExecutionError `protobuf:"bytes,6,opt,name=error,proto3" json:"error,omitempty"`
	// Status of the Workflow execution
	Status Execution_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.serverless.workflows.v1.Execution_Status" json:"status,omitempty"`
	// Start timestamp for the Workflow execution.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Duration of the Workflow execution.
	Duration *durationpb.Duration `protobuf:"bytes,9,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *Execution) Reset() {
	*x = Execution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Execution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Execution) ProtoMessage() {}

func (x *Execution) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Execution.ProtoReflect.Descriptor instead.
func (*Execution) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP(), []int{0}
}

func (x *Execution) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Execution) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *Execution) GetInput() *ExecutionInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *Execution) GetResult() *ExecutionResult {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *Execution) GetError() *ExecutionError {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *Execution) GetStatus() Execution_Status {
	if x != nil {
		return x.Status
	}
	return Execution_STATUS_UNSPECIFIED
}

func (x *Execution) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Execution) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type ExecutionPreview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Workflow execution. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the Workflow.
	WorkflowId string `protobuf:"bytes,2,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	// Status of the Workflow execution
	Status Execution_Status `protobuf:"varint,4,opt,name=status,proto3,enum=yandex.cloud.serverless.workflows.v1.Execution_Status" json:"status,omitempty"`
	// Start timestamp for the Workflow execution.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Duration of the Workflow execution.
	Duration *durationpb.Duration `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
}

func (x *ExecutionPreview) Reset() {
	*x = ExecutionPreview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionPreview) ProtoMessage() {}

func (x *ExecutionPreview) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionPreview.ProtoReflect.Descriptor instead.
func (*ExecutionPreview) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP(), []int{1}
}

func (x *ExecutionPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ExecutionPreview) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *ExecutionPreview) GetStatus() Execution_Status {
	if x != nil {
		return x.Status
	}
	return Execution_STATUS_UNSPECIFIED
}

func (x *ExecutionPreview) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *ExecutionPreview) GetDuration() *durationpb.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

type ExecutionInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Input:
	//
	//	*ExecutionInput_InputJson
	Input isExecutionInput_Input `protobuf_oneof:"input"`
}

func (x *ExecutionInput) Reset() {
	*x = ExecutionInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionInput) ProtoMessage() {}

func (x *ExecutionInput) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionInput.ProtoReflect.Descriptor instead.
func (*ExecutionInput) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP(), []int{2}
}

func (m *ExecutionInput) GetInput() isExecutionInput_Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (x *ExecutionInput) GetInputJson() string {
	if x, ok := x.GetInput().(*ExecutionInput_InputJson); ok {
		return x.InputJson
	}
	return ""
}

type isExecutionInput_Input interface {
	isExecutionInput_Input()
}

type ExecutionInput_InputJson struct {
	// JSON input data for the Workflow execution.
	InputJson string `protobuf:"bytes,1,opt,name=input_json,json=inputJson,proto3,oneof"`
}

func (*ExecutionInput_InputJson) isExecutionInput_Input() {}

type ExecutionResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Result:
	//
	//	*ExecutionResult_ResultJson
	Result isExecutionResult_Result `protobuf_oneof:"result"`
}

func (x *ExecutionResult) Reset() {
	*x = ExecutionResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionResult) ProtoMessage() {}

func (x *ExecutionResult) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionResult.ProtoReflect.Descriptor instead.
func (*ExecutionResult) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP(), []int{3}
}

func (m *ExecutionResult) GetResult() isExecutionResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (x *ExecutionResult) GetResultJson() string {
	if x, ok := x.GetResult().(*ExecutionResult_ResultJson); ok {
		return x.ResultJson
	}
	return ""
}

type isExecutionResult_Result interface {
	isExecutionResult_Result()
}

type ExecutionResult_ResultJson struct {
	// JSON result of the Workflow execution.
	ResultJson string `protobuf:"bytes,1,opt,name=result_json,json=resultJson,proto3,oneof"`
}

func (*ExecutionResult_ResultJson) isExecutionResult_Result() {}

type ExecutionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Error message of the Workflow execution.
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	// Error code of the Workflow execution.
	ErrorCode string `protobuf:"bytes,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
}

func (x *ExecutionError) Reset() {
	*x = ExecutionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionError) ProtoMessage() {}

func (x *ExecutionError) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionError.ProtoReflect.Descriptor instead.
func (*ExecutionError) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP(), []int{4}
}

func (x *ExecutionError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExecutionError) GetErrorCode() string {
	if x != nil {
		return x.ErrorCode
	}
	return ""
}

var File_yandex_cloud_serverless_workflows_v1_execution_proto protoreflect.FileDescriptor

var file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDesc = []byte{
	0x0a, 0x34, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x04, 0x0a,
	0x09, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x4a, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x6c, 0x65, 0x73, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73,
	0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x4a, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x4e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x77, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a,
	0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x46, 0x49, 0x4e, 0x49, 0x53, 0x48, 0x45, 0x44,
	0x10, 0x06, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x8b, 0x02, 0x0a, 0x10, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x12, 0x4e,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x40, 0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x0a, 0x0a, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x42, 0x0d, 0x0a, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x44, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x0b, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x42, 0x0e,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x49,
	0x0a, 0x0e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x7e, 0x0a, 0x28, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x52, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescData = file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDesc
)

func file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescData)
	})
	return file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDescData
}

var file_yandex_cloud_serverless_workflows_v1_execution_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_serverless_workflows_v1_execution_proto_goTypes = []interface{}{
	(Execution_Status)(0),         // 0: yandex.cloud.serverless.workflows.v1.Execution.Status
	(*Execution)(nil),             // 1: yandex.cloud.serverless.workflows.v1.Execution
	(*ExecutionPreview)(nil),      // 2: yandex.cloud.serverless.workflows.v1.ExecutionPreview
	(*ExecutionInput)(nil),        // 3: yandex.cloud.serverless.workflows.v1.ExecutionInput
	(*ExecutionResult)(nil),       // 4: yandex.cloud.serverless.workflows.v1.ExecutionResult
	(*ExecutionError)(nil),        // 5: yandex.cloud.serverless.workflows.v1.ExecutionError
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 7: google.protobuf.Duration
}
var file_yandex_cloud_serverless_workflows_v1_execution_proto_depIdxs = []int32{
	3, // 0: yandex.cloud.serverless.workflows.v1.Execution.input:type_name -> yandex.cloud.serverless.workflows.v1.ExecutionInput
	4, // 1: yandex.cloud.serverless.workflows.v1.Execution.result:type_name -> yandex.cloud.serverless.workflows.v1.ExecutionResult
	5, // 2: yandex.cloud.serverless.workflows.v1.Execution.error:type_name -> yandex.cloud.serverless.workflows.v1.ExecutionError
	0, // 3: yandex.cloud.serverless.workflows.v1.Execution.status:type_name -> yandex.cloud.serverless.workflows.v1.Execution.Status
	6, // 4: yandex.cloud.serverless.workflows.v1.Execution.started_at:type_name -> google.protobuf.Timestamp
	7, // 5: yandex.cloud.serverless.workflows.v1.Execution.duration:type_name -> google.protobuf.Duration
	0, // 6: yandex.cloud.serverless.workflows.v1.ExecutionPreview.status:type_name -> yandex.cloud.serverless.workflows.v1.Execution.Status
	6, // 7: yandex.cloud.serverless.workflows.v1.ExecutionPreview.started_at:type_name -> google.protobuf.Timestamp
	7, // 8: yandex.cloud.serverless.workflows.v1.ExecutionPreview.duration:type_name -> google.protobuf.Duration
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_workflows_v1_execution_proto_init() }
func file_yandex_cloud_serverless_workflows_v1_execution_proto_init() {
	if File_yandex_cloud_serverless_workflows_v1_execution_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Execution); i {
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
		file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionPreview); i {
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
		file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionInput); i {
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
		file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionResult); i {
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
		file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionError); i {
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
	file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ExecutionInput_InputJson)(nil),
	}
	file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*ExecutionResult_ResultJson)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_workflows_v1_execution_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_workflows_v1_execution_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_serverless_workflows_v1_execution_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_serverless_workflows_v1_execution_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_workflows_v1_execution_proto = out.File
	file_yandex_cloud_serverless_workflows_v1_execution_proto_rawDesc = nil
	file_yandex_cloud_serverless_workflows_v1_execution_proto_goTypes = nil
	file_yandex_cloud_serverless_workflows_v1_execution_proto_depIdxs = nil
}