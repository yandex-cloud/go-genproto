// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/dataset/v1/dataset.proto

package fomo

import (
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

// Status of the dataset.
type DatasetInfo_Status int32

const (
	DatasetInfo_STATUS_UNSPECIFIED DatasetInfo_Status = 0
	DatasetInfo_DRAFT              DatasetInfo_Status = 1
	DatasetInfo_VALIDATING         DatasetInfo_Status = 2
	DatasetInfo_READY              DatasetInfo_Status = 3
	DatasetInfo_INVALID            DatasetInfo_Status = 4
	DatasetInfo_DELETING           DatasetInfo_Status = 5
)

// Enum value maps for DatasetInfo_Status.
var (
	DatasetInfo_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "DRAFT",
		2: "VALIDATING",
		3: "READY",
		4: "INVALID",
		5: "DELETING",
	}
	DatasetInfo_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"DRAFT":              1,
		"VALIDATING":         2,
		"READY":              3,
		"INVALID":            4,
		"DELETING":           5,
	}
)

func (x DatasetInfo_Status) Enum() *DatasetInfo_Status {
	p := new(DatasetInfo_Status)
	*p = x
	return p
}

func (x DatasetInfo_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DatasetInfo_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_dataset_v1_dataset_proto_enumTypes[0].Descriptor()
}

func (DatasetInfo_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_dataset_v1_dataset_proto_enumTypes[0]
}

func (x DatasetInfo_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DatasetInfo_Status.Descriptor instead.
func (DatasetInfo_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescGZIP(), []int{0, 0}
}

// Information about the dataset.
type DatasetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the dataset.
	DatasetId string `protobuf:"bytes,1,opt,name=dataset_id,json=datasetId,proto3" json:"dataset_id,omitempty"`
	// Folder ID of the dataset.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the dataset.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the dataset.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Metadata of the dataset.
	Metadata string `protobuf:"bytes,5,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Status of the dataset.
	Status DatasetInfo_Status `protobuf:"varint,6,opt,name=status,proto3,enum=yandex.cloud.ai.dataset.v1.DatasetInfo_Status" json:"status,omitempty"`
	// Task type of the dataset.
	TaskType string `protobuf:"bytes,7,opt,name=task_type,json=taskType,proto3" json:"task_type,omitempty"`
	// Create dataset timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Update dataset timestamp.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Number of rows in the dataset.
	Rows int64 `protobuf:"varint,10,opt,name=rows,proto3" json:"rows,omitempty"`
	// Size of the dataset.
	SizeBytes int64 `protobuf:"varint,11,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	// Deprecated. Use created_by instead
	//
	// Deprecated: Marked as deprecated in yandex/cloud/ai/dataset/v1/dataset.proto.
	CreatedById string `protobuf:"bytes,12,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	// Labels of the dataset
	Labels map[string]string `protobuf:"bytes,13,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// User ID of the dataset's creator.
	CreatedBy string `protobuf:"bytes,14,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// User ID of the dataset's last updater.
	UpdatedBy       string             `protobuf:"bytes,15,opt,name=updated_by,json=updatedBy,proto3" json:"updated_by,omitempty"`
	ValidationError []*ValidationError `protobuf:"bytes,21,rep,name=validation_error,json=validationError,proto3" json:"validation_error,omitempty"`
}

func (x *DatasetInfo) Reset() {
	*x = DatasetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DatasetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DatasetInfo) ProtoMessage() {}

func (x *DatasetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DatasetInfo.ProtoReflect.Descriptor instead.
func (*DatasetInfo) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescGZIP(), []int{0}
}

func (x *DatasetInfo) GetDatasetId() string {
	if x != nil {
		return x.DatasetId
	}
	return ""
}

func (x *DatasetInfo) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *DatasetInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DatasetInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *DatasetInfo) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *DatasetInfo) GetStatus() DatasetInfo_Status {
	if x != nil {
		return x.Status
	}
	return DatasetInfo_STATUS_UNSPECIFIED
}

func (x *DatasetInfo) GetTaskType() string {
	if x != nil {
		return x.TaskType
	}
	return ""
}

func (x *DatasetInfo) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *DatasetInfo) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *DatasetInfo) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

func (x *DatasetInfo) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

// Deprecated: Marked as deprecated in yandex/cloud/ai/dataset/v1/dataset.proto.
func (x *DatasetInfo) GetCreatedById() string {
	if x != nil {
		return x.CreatedById
	}
	return ""
}

func (x *DatasetInfo) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *DatasetInfo) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *DatasetInfo) GetUpdatedBy() string {
	if x != nil {
		return x.UpdatedBy
	}
	return ""
}

func (x *DatasetInfo) GetValidationError() []*ValidationError {
	if x != nil {
		return x.ValidationError
	}
	return nil
}

// Information about dataset validation error.
type ValidationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the validation error.
	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// Description of the validation error.
	ErrorDescription string `protobuf:"bytes,2,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	// Row numbers in which the error occurred.
	RowNumbers []int64 `protobuf:"varint,3,rep,packed,name=row_numbers,json=rowNumbers,proto3" json:"row_numbers,omitempty"`
}

func (x *ValidationError) Reset() {
	*x = ValidationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationError) ProtoMessage() {}

func (x *ValidationError) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationError.ProtoReflect.Descriptor instead.
func (*ValidationError) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescGZIP(), []int{1}
}

func (x *ValidationError) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *ValidationError) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

func (x *ValidationError) GetRowNumbers() []int64 {
	if x != nil {
		return x.RowNumbers
	}
	return nil
}

var File_yandex_cloud_ai_dataset_v1_dataset_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x06, 0x0a, 0x0b, 0x44, 0x61, 0x74, 0x61,
	0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x65, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x72, 0x6f, 0x77, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x56, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x15, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x0f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x39,
	0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x61, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x44,
	0x52, 0x41, 0x46, 0x54, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41,
	0x54, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10,
	0x03, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x4a, 0x04, 0x08, 0x10,
	0x10, 0x15, 0x22, 0x75, 0x0a, 0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x6f, 0x77, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x72,
	0x6f, 0x77, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x42, 0x65, 0x0a, 0x1e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x6f, 0x6d, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescData = file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDesc
)

func file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescData)
	})
	return file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDescData
}

var file_yandex_cloud_ai_dataset_v1_dataset_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_ai_dataset_v1_dataset_proto_goTypes = []any{
	(DatasetInfo_Status)(0),       // 0: yandex.cloud.ai.dataset.v1.DatasetInfo.Status
	(*DatasetInfo)(nil),           // 1: yandex.cloud.ai.dataset.v1.DatasetInfo
	(*ValidationError)(nil),       // 2: yandex.cloud.ai.dataset.v1.ValidationError
	nil,                           // 3: yandex.cloud.ai.dataset.v1.DatasetInfo.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_yandex_cloud_ai_dataset_v1_dataset_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.ai.dataset.v1.DatasetInfo.status:type_name -> yandex.cloud.ai.dataset.v1.DatasetInfo.Status
	4, // 1: yandex.cloud.ai.dataset.v1.DatasetInfo.created_at:type_name -> google.protobuf.Timestamp
	4, // 2: yandex.cloud.ai.dataset.v1.DatasetInfo.updated_at:type_name -> google.protobuf.Timestamp
	3, // 3: yandex.cloud.ai.dataset.v1.DatasetInfo.labels:type_name -> yandex.cloud.ai.dataset.v1.DatasetInfo.LabelsEntry
	2, // 4: yandex.cloud.ai.dataset.v1.DatasetInfo.validation_error:type_name -> yandex.cloud.ai.dataset.v1.ValidationError
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_dataset_v1_dataset_proto_init() }
func file_yandex_cloud_ai_dataset_v1_dataset_proto_init() {
	if File_yandex_cloud_ai_dataset_v1_dataset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*DatasetInfo); i {
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
		file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ValidationError); i {
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
			RawDescriptor: file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_dataset_v1_dataset_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_dataset_v1_dataset_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_dataset_v1_dataset_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_dataset_v1_dataset_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_dataset_v1_dataset_proto = out.File
	file_yandex_cloud_ai_dataset_v1_dataset_proto_rawDesc = nil
	file_yandex_cloud_ai_dataset_v1_dataset_proto_goTypes = nil
	file_yandex_cloud_ai_dataset_v1_dataset_proto_depIdxs = nil
}
