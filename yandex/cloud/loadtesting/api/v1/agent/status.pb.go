// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/api/v1/agent/status.proto

package agent

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

// Agent status.
type Status int32

const (
	// Status is not specified.
	Status_STATUS_UNSPECIFIED Status = 0
	// Agent is preparing a test to be executed.
	Status_PREPARING_TEST Status = 1
	// Agent is ready to take a test.
	Status_READY_FOR_TEST Status = 2
	// Agent is executing a test.
	Status_TESTING Status = 3
	// Agent application encountered an error and cannot operate normally.
	Status_TANK_FAILED Status = 4
	// Agent is waiting for resources to be allocated.
	Status_PROVISIONING Status = 5
	// Agent is being stopped.
	Status_STOPPING Status = 6
	// Agent is stopped.
	Status_STOPPED Status = 7
	// Agent is being started.
	Status_STARTING Status = 8
	// Agent is being restarted.
	Status_RESTARTING Status = 9
	// Agent is being updated.
	Status_UPDATING Status = 10
	// Agent encountered an error and cannot operate.
	Status_ERROR Status = 11
	// Agent is crashed and will be restarted automatically.
	Status_CRASHED Status = 12
	// Agent is being deleted.
	Status_DELETING Status = 13
	// Service is waiting for connection with agent to be established.
	Status_INITIALIZING_CONNECTION Status = 15
	// Service has lost connection with agent.
	Status_LOST_CONNECTION_WITH_AGENT Status = 16
	// Agent is uploading test artifacts.
	Status_UPLOADING_ARTIFACTS Status = 17
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "PREPARING_TEST",
		2:  "READY_FOR_TEST",
		3:  "TESTING",
		4:  "TANK_FAILED",
		5:  "PROVISIONING",
		6:  "STOPPING",
		7:  "STOPPED",
		8:  "STARTING",
		9:  "RESTARTING",
		10: "UPDATING",
		11: "ERROR",
		12: "CRASHED",
		13: "DELETING",
		15: "INITIALIZING_CONNECTION",
		16: "LOST_CONNECTION_WITH_AGENT",
		17: "UPLOADING_ARTIFACTS",
	}
	Status_value = map[string]int32{
		"STATUS_UNSPECIFIED":         0,
		"PREPARING_TEST":             1,
		"READY_FOR_TEST":             2,
		"TESTING":                    3,
		"TANK_FAILED":                4,
		"PROVISIONING":               5,
		"STOPPING":                   6,
		"STOPPED":                    7,
		"STARTING":                   8,
		"RESTARTING":                 9,
		"UPDATING":                   10,
		"ERROR":                      11,
		"CRASHED":                    12,
		"DELETING":                   13,
		"INITIALIZING_CONNECTION":    15,
		"LOST_CONNECTION_WITH_AGENT": 16,
		"UPLOADING_ARTIFACTS":        17,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_loadtesting_api_v1_agent_status_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_loadtesting_api_v1_agent_status_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescGZIP(), []int{0}
}

var File_yandex_cloud_loadtesting_api_v1_agent_status_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDesc = []byte{
	0x0a, 0x32, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2a, 0xbb, 0x02, 0x0a, 0x06,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12,
	0x0a, 0x0e, 0x50, 0x52, 0x45, 0x50, 0x41, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x45, 0x53, 0x54,
	0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x52, 0x45, 0x41, 0x44, 0x59, 0x5f, 0x46, 0x4f, 0x52, 0x5f,
	0x54, 0x45, 0x53, 0x54, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x45, 0x53, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x41, 0x4e, 0x4b, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x10, 0x04, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49, 0x4f,
	0x4e, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x49,
	0x4e, 0x47, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10,
	0x07, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12,
	0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12,
	0x0c, 0x0a, 0x08, 0x55, 0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0b, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x52, 0x41, 0x53,
	0x48, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e,
	0x47, 0x10, 0x0d, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x4c, 0x49, 0x5a,
	0x49, 0x4e, 0x47, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0f,
	0x12, 0x1e, 0x0a, 0x1a, 0x4c, 0x4f, 0x53, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x10, 0x10,
	0x12, 0x17, 0x0a, 0x13, 0x55, 0x50, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x52,
	0x54, 0x49, 0x46, 0x41, 0x43, 0x54, 0x53, 0x10, 0x11, 0x42, 0x7c, 0x0a, 0x29, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f,
	0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5a, 0x4f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescData = file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDesc
)

func file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_agent_status_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_loadtesting_api_v1_agent_status_proto_goTypes = []any{
	(Status)(0), // 0: yandex.cloud.loadtesting.api.v1.agent.Status
}
var file_yandex_cloud_loadtesting_api_v1_agent_status_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_agent_status_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_agent_status_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_agent_status_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_agent_status_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_agent_status_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_loadtesting_api_v1_agent_status_proto_enumTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_agent_status_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_agent_status_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_api_v1_agent_status_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_agent_status_proto_depIdxs = nil
}
