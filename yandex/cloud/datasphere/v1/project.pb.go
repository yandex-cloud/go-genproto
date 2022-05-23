// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: yandex/cloud/datasphere/v1/project.proto

package datasphere

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Project_Settings_CommitMode int32

const (
	Project_Settings_COMMIT_MODE_UNSPECIFIED Project_Settings_CommitMode = 0
	// Commit happens after the execution of a cell or group of cells or after completion with an error.
	Project_Settings_STANDARD Project_Settings_CommitMode = 1
	// Commit happens periodically.
	// Also, automatic saving of state occurs when switching to another type of computing resource.
	Project_Settings_AUTO Project_Settings_CommitMode = 2
)

// Enum value maps for Project_Settings_CommitMode.
var (
	Project_Settings_CommitMode_name = map[int32]string{
		0: "COMMIT_MODE_UNSPECIFIED",
		1: "STANDARD",
		2: "AUTO",
	}
	Project_Settings_CommitMode_value = map[string]int32{
		"COMMIT_MODE_UNSPECIFIED": 0,
		"STANDARD":                1,
		"AUTO":                    2,
	}
)

func (x Project_Settings_CommitMode) Enum() *Project_Settings_CommitMode {
	p := new(Project_Settings_CommitMode)
	*p = x
	return p
}

func (x Project_Settings_CommitMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_Settings_CommitMode) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datasphere_v1_project_proto_enumTypes[0].Descriptor()
}

func (Project_Settings_CommitMode) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datasphere_v1_project_proto_enumTypes[0]
}

func (x Project_Settings_CommitMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_Settings_CommitMode.Descriptor instead.
func (Project_Settings_CommitMode) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_proto_rawDescGZIP(), []int{0, 0, 0}
}

// A Project resource.
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the project.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the project belongs to.
	FolderId  string                 `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the project. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the project. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Settings of the project.
	Settings *Project_Settings `protobuf:"bytes,6,opt,name=settings,proto3" json:"settings,omitempty"`
	// Limits of the project.
	Limits *Project_Limits `protobuf:"bytes,7,opt,name=limits,proto3" json:"limits,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Project) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Project) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Project) GetSettings() *Project_Settings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *Project) GetLimits() *Project_Limits {
	if x != nil {
		return x.Limits
	}
	return nil
}

type Project_Settings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the service account, on whose behalf all operations with clusters will be performed.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// ID of the subnet where the DataProc cluster resides.
	// Currently only subnets created in the availability zone ru-central1-a are supported.
	SubnetId string `protobuf:"bytes,2,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// ID of the DataProc cluster.
	DataProcClusterId string `protobuf:"bytes,3,opt,name=data_proc_cluster_id,json=dataProcClusterId,proto3" json:"data_proc_cluster_id,omitempty"`
	// Commit mode that is assigned to the project.
	CommitMode Project_Settings_CommitMode `protobuf:"varint,4,opt,name=commit_mode,json=commitMode,proto3,enum=yandex.cloud.datasphere.v1.Project_Settings_CommitMode" json:"commit_mode,omitempty"`
	// Network interfaces security groups.
	SecurityGroupIds []string `protobuf:"bytes,5,rep,name=security_group_ids,json=securityGroupIds,proto3" json:"security_group_ids,omitempty"`
}

func (x *Project_Settings) Reset() {
	*x = Project_Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_Settings) ProtoMessage() {}

func (x *Project_Settings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project_Settings.ProtoReflect.Descriptor instead.
func (*Project_Settings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Project_Settings) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Project_Settings) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *Project_Settings) GetDataProcClusterId() string {
	if x != nil {
		return x.DataProcClusterId
	}
	return ""
}

func (x *Project_Settings) GetCommitMode() Project_Settings_CommitMode {
	if x != nil {
		return x.CommitMode
	}
	return Project_Settings_COMMIT_MODE_UNSPECIFIED
}

func (x *Project_Settings) GetSecurityGroupIds() []string {
	if x != nil {
		return x.SecurityGroupIds
	}
	return nil
}

type Project_Limits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of units that can be spent per hour.
	MaxUnitsPerHour *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=max_units_per_hour,json=maxUnitsPerHour,proto3" json:"max_units_per_hour,omitempty"`
	// The number of units that can be spent on the one execution.
	MaxUnitsPerExecution *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=max_units_per_execution,json=maxUnitsPerExecution,proto3" json:"max_units_per_execution,omitempty"`
}

func (x *Project_Limits) Reset() {
	*x = Project_Limits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_Limits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_Limits) ProtoMessage() {}

func (x *Project_Limits) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project_Limits.ProtoReflect.Descriptor instead.
func (*Project_Limits) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Project_Limits) GetMaxUnitsPerHour() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxUnitsPerHour
	}
	return nil
}

func (x *Project_Limits) GetMaxUnitsPerExecution() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxUnitsPerExecution
	}
	return nil
}

var File_yandex_cloud_datasphere_v1_project_proto protoreflect.FileDescriptor

var file_yandex_cloud_datasphere_v1_project_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68,
	0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb2, 0x06, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x1a,
	0xd1, 0x02, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x2c, 0x0a, 0x12,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x70, 0x72, 0x6f, 0x63, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x63, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x37, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d,
	0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x10,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x73,
	0x22, 0x41, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53,
	0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54,
	0x4f, 0x10, 0x02, 0x1a, 0xa6, 0x01, 0x0a, 0x06, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x48,
	0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x68, 0x6f, 0x75, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x74,
	0x73, 0x50, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x52, 0x0a, 0x17, 0x6d, 0x61, 0x78, 0x5f,
	0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36,
	0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6d, 0x61, 0x78, 0x55, 0x6e, 0x69, 0x74, 0x73,
	0x50, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x6b, 0x0a, 0x1e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x49,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_datasphere_v1_project_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v1_project_proto_rawDescData = file_yandex_cloud_datasphere_v1_project_proto_rawDesc
)

func file_yandex_cloud_datasphere_v1_project_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v1_project_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datasphere_v1_project_proto_rawDescData)
	})
	return file_yandex_cloud_datasphere_v1_project_proto_rawDescData
}

var file_yandex_cloud_datasphere_v1_project_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_datasphere_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_datasphere_v1_project_proto_goTypes = []interface{}{
	(Project_Settings_CommitMode)(0), // 0: yandex.cloud.datasphere.v1.Project.Settings.CommitMode
	(*Project)(nil),                  // 1: yandex.cloud.datasphere.v1.Project
	(*Project_Settings)(nil),         // 2: yandex.cloud.datasphere.v1.Project.Settings
	(*Project_Limits)(nil),           // 3: yandex.cloud.datasphere.v1.Project.Limits
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
	(*wrapperspb.Int64Value)(nil),    // 5: google.protobuf.Int64Value
}
var file_yandex_cloud_datasphere_v1_project_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.datasphere.v1.Project.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: yandex.cloud.datasphere.v1.Project.settings:type_name -> yandex.cloud.datasphere.v1.Project.Settings
	3, // 2: yandex.cloud.datasphere.v1.Project.limits:type_name -> yandex.cloud.datasphere.v1.Project.Limits
	0, // 3: yandex.cloud.datasphere.v1.Project.Settings.commit_mode:type_name -> yandex.cloud.datasphere.v1.Project.Settings.CommitMode
	5, // 4: yandex.cloud.datasphere.v1.Project.Limits.max_units_per_hour:type_name -> google.protobuf.Int64Value
	5, // 5: yandex.cloud.datasphere.v1.Project.Limits.max_units_per_execution:type_name -> google.protobuf.Int64Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v1_project_proto_init() }
func file_yandex_cloud_datasphere_v1_project_proto_init() {
	if File_yandex_cloud_datasphere_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datasphere_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
		file_yandex_cloud_datasphere_v1_project_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project_Settings); i {
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
		file_yandex_cloud_datasphere_v1_project_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project_Limits); i {
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
			RawDescriptor: file_yandex_cloud_datasphere_v1_project_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datasphere_v1_project_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v1_project_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_datasphere_v1_project_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_datasphere_v1_project_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v1_project_proto = out.File
	file_yandex_cloud_datasphere_v1_project_proto_rawDesc = nil
	file_yandex_cloud_datasphere_v1_project_proto_goTypes = nil
	file_yandex_cloud_datasphere_v1_project_proto_depIdxs = nil
}
