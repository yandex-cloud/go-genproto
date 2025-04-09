// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/tuning/v1/tuning_task.proto

package fomo

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type TuningTask_Status int32

const (
	TuningTask_STATUS_UNSPECIFIED TuningTask_Status = 0
	TuningTask_CREATED            TuningTask_Status = 1
	TuningTask_PENDING            TuningTask_Status = 2
	TuningTask_IN_PROGRESS        TuningTask_Status = 3
	TuningTask_COMPLETED          TuningTask_Status = 4
	TuningTask_FAILED             TuningTask_Status = 5
	TuningTask_CANCELED           TuningTask_Status = 6
	TuningTask_DRAFT              TuningTask_Status = 7
)

// Enum value maps for TuningTask_Status.
var (
	TuningTask_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATED",
		2: "PENDING",
		3: "IN_PROGRESS",
		4: "COMPLETED",
		5: "FAILED",
		6: "CANCELED",
		7: "DRAFT",
	}
	TuningTask_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATED":            1,
		"PENDING":            2,
		"IN_PROGRESS":        3,
		"COMPLETED":          4,
		"FAILED":             5,
		"CANCELED":           6,
		"DRAFT":              7,
	}
)

func (x TuningTask_Status) Enum() *TuningTask_Status {
	p := new(TuningTask_Status)
	*p = x
	return p
}

func (x TuningTask_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TuningTask_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_tuning_v1_tuning_task_proto_enumTypes[0].Descriptor()
}

func (TuningTask_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_tuning_v1_tuning_task_proto_enumTypes[0]
}

func (x TuningTask_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TuningTask_Status.Descriptor instead.
func (TuningTask_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescGZIP(), []int{0, 0}
}

type TuningTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId         string                 `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	OperationId    string                 `protobuf:"bytes,3,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	Status         TuningTask_Status      `protobuf:"varint,4,opt,name=status,proto3,enum=yandex.cloud.ai.tuning.v1.TuningTask_Status" json:"status,omitempty"`
	FolderId       string                 `protobuf:"bytes,5,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedBy      string                 `protobuf:"bytes,6,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	CreatedAt      *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StartedAt      *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	FinishedAt     *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	SourceModelUri string                 `protobuf:"bytes,10,opt,name=source_model_uri,json=sourceModelUri,proto3" json:"source_model_uri,omitempty"`
	TargetModelUri string                 `protobuf:"bytes,11,opt,name=target_model_uri,json=targetModelUri,proto3" json:"target_model_uri,omitempty"`
	Name           string                 `protobuf:"bytes,12,opt,name=name,proto3" json:"name,omitempty"`
	Description    string                 `protobuf:"bytes,13,opt,name=description,proto3" json:"description,omitempty"`
	Labels         map[string]string      `protobuf:"bytes,14,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TuningTask) Reset() {
	*x = TuningTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_tuning_v1_tuning_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TuningTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TuningTask) ProtoMessage() {}

func (x *TuningTask) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_tuning_v1_tuning_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TuningTask.ProtoReflect.Descriptor instead.
func (*TuningTask) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescGZIP(), []int{0}
}

func (x *TuningTask) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TuningTask) GetOperationId() string {
	if x != nil {
		return x.OperationId
	}
	return ""
}

func (x *TuningTask) GetStatus() TuningTask_Status {
	if x != nil {
		return x.Status
	}
	return TuningTask_STATUS_UNSPECIFIED
}

func (x *TuningTask) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *TuningTask) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *TuningTask) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *TuningTask) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *TuningTask) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *TuningTask) GetSourceModelUri() string {
	if x != nil {
		return x.SourceModelUri
	}
	return ""
}

func (x *TuningTask) GetTargetModelUri() string {
	if x != nil {
		return x.TargetModelUri
	}
	return ""
}

func (x *TuningTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TuningTask) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *TuningTask) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

var File_yandex_cloud_ai_tuning_v1_tuning_task_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x74,
	0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x06, 0x0a, 0x0a, 0x54, 0x75, 0x6e,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x62, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55,
	0x72, 0x69, 0x12, 0x28, 0x0a, 0x10, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x55, 0x72, 0x69, 0x12, 0x2f, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1b, 0xf2, 0xc7, 0x31, 0x17,
	0x5e, 0x5b, 0x41, 0x2d, 0x5a, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x20, 0x5f, 0x2d, 0x5d, 0x7b,
	0x33, 0x2c, 0x31, 0x30, 0x30, 0x7d, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x69, 0x2e, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x75, 0x6e, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x7f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x45, 0x41, 0x54,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53, 0x53,
	0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10,
	0x04, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0c, 0x0a,
	0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x44,
	0x52, 0x41, 0x46, 0x54, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x63, 0x0a, 0x1d,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x61, 0x69, 0x2e, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x5a, 0x42, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x61, 0x69, 0x2f, 0x74, 0x75, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x6f, 0x6d,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescData = file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDesc
)

func file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescData)
	})
	return file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDescData
}

var file_yandex_cloud_ai_tuning_v1_tuning_task_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_ai_tuning_v1_tuning_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_ai_tuning_v1_tuning_task_proto_goTypes = []any{
	(TuningTask_Status)(0),        // 0: yandex.cloud.ai.tuning.v1.TuningTask.Status
	(*TuningTask)(nil),            // 1: yandex.cloud.ai.tuning.v1.TuningTask
	nil,                           // 2: yandex.cloud.ai.tuning.v1.TuningTask.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_ai_tuning_v1_tuning_task_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.ai.tuning.v1.TuningTask.status:type_name -> yandex.cloud.ai.tuning.v1.TuningTask.Status
	3, // 1: yandex.cloud.ai.tuning.v1.TuningTask.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: yandex.cloud.ai.tuning.v1.TuningTask.started_at:type_name -> google.protobuf.Timestamp
	3, // 3: yandex.cloud.ai.tuning.v1.TuningTask.finished_at:type_name -> google.protobuf.Timestamp
	2, // 4: yandex.cloud.ai.tuning.v1.TuningTask.labels:type_name -> yandex.cloud.ai.tuning.v1.TuningTask.LabelsEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_tuning_v1_tuning_task_proto_init() }
func file_yandex_cloud_ai_tuning_v1_tuning_task_proto_init() {
	if File_yandex_cloud_ai_tuning_v1_tuning_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_tuning_v1_tuning_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TuningTask); i {
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
			RawDescriptor: file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_tuning_v1_tuning_task_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_tuning_v1_tuning_task_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_tuning_v1_tuning_task_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_tuning_v1_tuning_task_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_tuning_v1_tuning_task_proto = out.File
	file_yandex_cloud_ai_tuning_v1_tuning_task_proto_rawDesc = nil
	file_yandex_cloud_ai_tuning_v1_tuning_task_proto_goTypes = nil
	file_yandex_cloud_ai_tuning_v1_tuning_task_proto_depIdxs = nil
}
