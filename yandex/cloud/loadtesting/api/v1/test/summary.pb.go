// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/api/v1/test/summary.proto

package test

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

// Process of test and some results
type Summary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Status of the test.
	Status Status `protobuf:"varint,1,opt,name=status,proto3,enum=yandex.cloud.loadtesting.api.v1.test.Status" json:"status,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// UA or SA that created the test.
	CreatedBy string `protobuf:"bytes,3,opt,name=created_by,json=createdBy,proto3" json:"created_by,omitempty"`
	// Test start timestamp.
	//
	// Empty if the test has not been started yet.
	StartedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// Test finish timestamp.
	//
	// Empty if the test has not been finished yet.
	FinishedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=finished_at,json=finishedAt,proto3" json:"finished_at,omitempty"`
	// Indicates whether the test is finished.
	IsFinished bool `protobuf:"varint,6,opt,name=is_finished,json=isFinished,proto3" json:"is_finished,omitempty"`
	// Error message.
	Error string `protobuf:"bytes,7,opt,name=error,proto3" json:"error,omitempty"`
	// Detected imbalance point.
	//
	// Contains information about a state at the moment it has been
	// [auto-stopped](/docs/load-testing/concepts/auto-stop).
	//
	// Empty if no auto-stop occured.
	ImbalancePoint *ImbalancePoint `protobuf:"bytes,8,opt,name=imbalance_point,json=imbalancePoint,proto3" json:"imbalance_point,omitempty"`
	// ID of the agent that executed the test.
	AssignedAgentId string `protobuf:"bytes,9,opt,name=assigned_agent_id,json=assignedAgentId,proto3" json:"assigned_agent_id,omitempty"`
	// Test output artifacts.
	//
	// Link to the artifacts output target containing `.log` and other files collected
	// during test execution.
	Artifacts *FilePointer `protobuf:"bytes,10,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
}

func (x *Summary) Reset() {
	*x = Summary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_test_summary_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Summary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Summary) ProtoMessage() {}

func (x *Summary) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_test_summary_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Summary.ProtoReflect.Descriptor instead.
func (*Summary) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescGZIP(), []int{0}
}

func (x *Summary) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_UNSPECIFIED
}

func (x *Summary) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Summary) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *Summary) GetStartedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.StartedAt
	}
	return nil
}

func (x *Summary) GetFinishedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.FinishedAt
	}
	return nil
}

func (x *Summary) GetIsFinished() bool {
	if x != nil {
		return x.IsFinished
	}
	return false
}

func (x *Summary) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Summary) GetImbalancePoint() *ImbalancePoint {
	if x != nil {
		return x.ImbalancePoint
	}
	return nil
}

func (x *Summary) GetAssignedAgentId() string {
	if x != nil {
		return x.AssignedAgentId
	}
	return ""
}

func (x *Summary) GetArtifacts() *FilePointer {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

var File_yandex_cloud_loadtesting_api_v1_test_summary_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x3a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x69, 0x6d, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x31, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f,
	0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xb4, 0x04, 0x0a, 0x07, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12,
	0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x39, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x66, 0x69, 0x6e,
	0x69, 0x73, 0x68, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x69,
	0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73,
	0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x5d,
	0x0a, 0x0f, 0x69, 0x6d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x49,
	0x6d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x0e, 0x69,
	0x6d, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2a, 0x0a,
	0x11, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4f, 0x0a, 0x09, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64,
	0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x42, 0x79, 0x0a, 0x28, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x3b, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescData = file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDesc
)

func file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_test_summary_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_loadtesting_api_v1_test_summary_proto_goTypes = []any{
	(*Summary)(nil),               // 0: yandex.cloud.loadtesting.api.v1.test.Summary
	(Status)(0),                   // 1: yandex.cloud.loadtesting.api.v1.test.Status
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*ImbalancePoint)(nil),        // 3: yandex.cloud.loadtesting.api.v1.test.ImbalancePoint
	(*FilePointer)(nil),           // 4: yandex.cloud.loadtesting.api.v1.test.FilePointer
}
var file_yandex_cloud_loadtesting_api_v1_test_summary_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.loadtesting.api.v1.test.Summary.status:type_name -> yandex.cloud.loadtesting.api.v1.test.Status
	2, // 1: yandex.cloud.loadtesting.api.v1.test.Summary.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: yandex.cloud.loadtesting.api.v1.test.Summary.started_at:type_name -> google.protobuf.Timestamp
	2, // 3: yandex.cloud.loadtesting.api.v1.test.Summary.finished_at:type_name -> google.protobuf.Timestamp
	3, // 4: yandex.cloud.loadtesting.api.v1.test.Summary.imbalance_point:type_name -> yandex.cloud.loadtesting.api.v1.test.ImbalancePoint
	4, // 5: yandex.cloud.loadtesting.api.v1.test.Summary.artifacts:type_name -> yandex.cloud.loadtesting.api.v1.test.FilePointer
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_test_summary_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_test_summary_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_test_summary_proto != nil {
		return
	}
	file_yandex_cloud_loadtesting_api_v1_test_file_pointer_proto_init()
	file_yandex_cloud_loadtesting_api_v1_test_imbalance_point_proto_init()
	file_yandex_cloud_loadtesting_api_v1_test_status_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_api_v1_test_summary_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Summary); i {
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
			RawDescriptor: file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_test_summary_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_test_summary_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_api_v1_test_summary_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_test_summary_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_test_summary_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_api_v1_test_summary_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_test_summary_proto_depIdxs = nil
}
