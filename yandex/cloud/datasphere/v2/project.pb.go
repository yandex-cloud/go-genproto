// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/datasphere/v2/project.proto

package datasphere

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
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
	return file_yandex_cloud_datasphere_v2_project_proto_enumTypes[0].Descriptor()
}

func (Project_Settings_CommitMode) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datasphere_v2_project_proto_enumTypes[0]
}

func (x Project_Settings_CommitMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_Settings_CommitMode.Descriptor instead.
func (Project_Settings_CommitMode) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Project_Settings_Ide int32

const (
	Project_Settings_IDE_UNSPECIFIED Project_Settings_Ide = 0
	// Project running on JupyterLab IDE.
	Project_Settings_JUPYTER_LAB Project_Settings_Ide = 1
)

// Enum value maps for Project_Settings_Ide.
var (
	Project_Settings_Ide_name = map[int32]string{
		0: "IDE_UNSPECIFIED",
		1: "JUPYTER_LAB",
	}
	Project_Settings_Ide_value = map[string]int32{
		"IDE_UNSPECIFIED": 0,
		"JUPYTER_LAB":     1,
	}
)

func (x Project_Settings_Ide) Enum() *Project_Settings_Ide {
	p := new(Project_Settings_Ide)
	*p = x
	return p
}

func (x Project_Settings_Ide) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_Settings_Ide) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datasphere_v2_project_proto_enumTypes[1].Descriptor()
}

func (Project_Settings_Ide) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datasphere_v2_project_proto_enumTypes[1]
}

func (x Project_Settings_Ide) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_Settings_Ide.Descriptor instead.
func (Project_Settings_Ide) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0, 0, 1}
}

type Project_Settings_StaleExecutionTimeoutMode int32

const (
	Project_Settings_STALE_EXECUTION_TIMEOUT_MODE_UNSPECIFIED Project_Settings_StaleExecutionTimeoutMode = 0
	// Setting to automatically stop stale execution after one hour with low consumption.
	Project_Settings_ONE_HOUR Project_Settings_StaleExecutionTimeoutMode = 1
	// Setting to automatically stop stale execution after three hours with low consumption.
	Project_Settings_THREE_HOURS Project_Settings_StaleExecutionTimeoutMode = 2
	// Setting to never automatically stop stale executions.
	Project_Settings_NO_TIMEOUT Project_Settings_StaleExecutionTimeoutMode = 3
)

// Enum value maps for Project_Settings_StaleExecutionTimeoutMode.
var (
	Project_Settings_StaleExecutionTimeoutMode_name = map[int32]string{
		0: "STALE_EXECUTION_TIMEOUT_MODE_UNSPECIFIED",
		1: "ONE_HOUR",
		2: "THREE_HOURS",
		3: "NO_TIMEOUT",
	}
	Project_Settings_StaleExecutionTimeoutMode_value = map[string]int32{
		"STALE_EXECUTION_TIMEOUT_MODE_UNSPECIFIED": 0,
		"ONE_HOUR":    1,
		"THREE_HOURS": 2,
		"NO_TIMEOUT":  3,
	}
)

func (x Project_Settings_StaleExecutionTimeoutMode) Enum() *Project_Settings_StaleExecutionTimeoutMode {
	p := new(Project_Settings_StaleExecutionTimeoutMode)
	*p = x
	return p
}

func (x Project_Settings_StaleExecutionTimeoutMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_Settings_StaleExecutionTimeoutMode) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datasphere_v2_project_proto_enumTypes[2].Descriptor()
}

func (Project_Settings_StaleExecutionTimeoutMode) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datasphere_v2_project_proto_enumTypes[2]
}

func (x Project_Settings_StaleExecutionTimeoutMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_Settings_StaleExecutionTimeoutMode.Descriptor instead.
func (Project_Settings_StaleExecutionTimeoutMode) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0, 0, 2}
}

type Project_Settings_IdeExecutionMode int32

const (
	Project_Settings_IDE_EXECUTION_MODE_UNSPECIFIED Project_Settings_IdeExecutionMode = 0
	// VM is allocated for specific execution and deallocated after the execution ends.
	Project_Settings_SERVERLESS Project_Settings_IdeExecutionMode = 1
	// VM is allocated at the first execution and stays allocated until manually deallocated.
	// Or until timeout exceeded.
	Project_Settings_DEDICATED Project_Settings_IdeExecutionMode = 2
)

// Enum value maps for Project_Settings_IdeExecutionMode.
var (
	Project_Settings_IdeExecutionMode_name = map[int32]string{
		0: "IDE_EXECUTION_MODE_UNSPECIFIED",
		1: "SERVERLESS",
		2: "DEDICATED",
	}
	Project_Settings_IdeExecutionMode_value = map[string]int32{
		"IDE_EXECUTION_MODE_UNSPECIFIED": 0,
		"SERVERLESS":                     1,
		"DEDICATED":                      2,
	}
)

func (x Project_Settings_IdeExecutionMode) Enum() *Project_Settings_IdeExecutionMode {
	p := new(Project_Settings_IdeExecutionMode)
	*p = x
	return p
}

func (x Project_Settings_IdeExecutionMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Project_Settings_IdeExecutionMode) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datasphere_v2_project_proto_enumTypes[3].Descriptor()
}

func (Project_Settings_IdeExecutionMode) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datasphere_v2_project_proto_enumTypes[3]
}

func (x Project_Settings_IdeExecutionMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Project_Settings_IdeExecutionMode.Descriptor instead.
func (Project_Settings_IdeExecutionMode) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0, 0, 3}
}

// A Project resource.
type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the project.
	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the project. 1-63 characters long.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the project. 0-256 characters long.
	Description string            `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Labels      map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	CreatedById string            `protobuf:"bytes,6,opt,name=created_by_id,json=createdById,proto3" json:"created_by_id,omitempty"`
	// Settings of the project.
	Settings *Project_Settings `protobuf:"bytes,7,opt,name=settings,proto3" json:"settings,omitempty"`
	// Limits of the project.
	Limits *Project_Limits `protobuf:"bytes,8,opt,name=limits,proto3" json:"limits,omitempty"`
	// ID of the community that the project belongs to.
	CommunityId string `protobuf:"bytes,11,opt,name=community_id,json=communityId,proto3" json:"community_id,omitempty"`
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v2_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v2_project_proto_msgTypes[0]
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
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetId() string {
	if x != nil {
		return x.Id
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

func (x *Project) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Project) GetCreatedById() string {
	if x != nil {
		return x.CreatedById
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

func (x *Project) GetCommunityId() string {
	if x != nil {
		return x.CommunityId
	}
	return ""
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
	CommitMode Project_Settings_CommitMode `protobuf:"varint,4,opt,name=commit_mode,json=commitMode,proto3,enum=yandex.cloud.datasphere.v2.Project_Settings_CommitMode" json:"commit_mode,omitempty"`
	// Network interfaces security groups.
	SecurityGroupIds []string `protobuf:"bytes,5,rep,name=security_group_ids,json=securityGroupIds,proto3" json:"security_group_ids,omitempty"`
	// Is early access preview enabled for the project.
	EarlyAccess bool `protobuf:"varint,6,opt,name=early_access,json=earlyAccess,proto3" json:"early_access,omitempty"`
	// Project IDE.
	Ide Project_Settings_Ide `protobuf:"varint,7,opt,name=ide,proto3,enum=yandex.cloud.datasphere.v2.Project_Settings_Ide" json:"ide,omitempty"`
	// Default project folder ID.
	DefaultFolderId string `protobuf:"bytes,8,opt,name=default_folder_id,json=defaultFolderId,proto3" json:"default_folder_id,omitempty"`
	// Timeout to automatically stop stale executions.
	StaleExecTimeoutMode Project_Settings_StaleExecutionTimeoutMode `protobuf:"varint,9,opt,name=stale_exec_timeout_mode,json=staleExecTimeoutMode,proto3,enum=yandex.cloud.datasphere.v2.Project_Settings_StaleExecutionTimeoutMode" json:"stale_exec_timeout_mode,omitempty"`
	// VM allocation mode.
	IdeExecutionMode Project_Settings_IdeExecutionMode `protobuf:"varint,10,opt,name=ide_execution_mode,json=ideExecutionMode,proto3,enum=yandex.cloud.datasphere.v2.Project_Settings_IdeExecutionMode" json:"ide_execution_mode,omitempty"`
	// Timeout for VM deallocation.
	VmInactivityTimeout *durationpb.Duration `protobuf:"bytes,11,opt,name=vm_inactivity_timeout,json=vmInactivityTimeout,proto3" json:"vm_inactivity_timeout,omitempty"`
	// Default VM configuration for DEDICATED mode.
	DefaultDedicatedSpec string `protobuf:"bytes,12,opt,name=default_dedicated_spec,json=defaultDedicatedSpec,proto3" json:"default_dedicated_spec,omitempty"`
}

func (x *Project_Settings) Reset() {
	*x = Project_Settings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v2_project_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_Settings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_Settings) ProtoMessage() {}

func (x *Project_Settings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v2_project_proto_msgTypes[1]
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
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0, 0}
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

func (x *Project_Settings) GetEarlyAccess() bool {
	if x != nil {
		return x.EarlyAccess
	}
	return false
}

func (x *Project_Settings) GetIde() Project_Settings_Ide {
	if x != nil {
		return x.Ide
	}
	return Project_Settings_IDE_UNSPECIFIED
}

func (x *Project_Settings) GetDefaultFolderId() string {
	if x != nil {
		return x.DefaultFolderId
	}
	return ""
}

func (x *Project_Settings) GetStaleExecTimeoutMode() Project_Settings_StaleExecutionTimeoutMode {
	if x != nil {
		return x.StaleExecTimeoutMode
	}
	return Project_Settings_STALE_EXECUTION_TIMEOUT_MODE_UNSPECIFIED
}

func (x *Project_Settings) GetIdeExecutionMode() Project_Settings_IdeExecutionMode {
	if x != nil {
		return x.IdeExecutionMode
	}
	return Project_Settings_IDE_EXECUTION_MODE_UNSPECIFIED
}

func (x *Project_Settings) GetVmInactivityTimeout() *durationpb.Duration {
	if x != nil {
		return x.VmInactivityTimeout
	}
	return nil
}

func (x *Project_Settings) GetDefaultDedicatedSpec() string {
	if x != nil {
		return x.DefaultDedicatedSpec
	}
	return ""
}

type Project_Limits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of units that can be spent per hour.
	MaxUnitsPerHour *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=max_units_per_hour,json=maxUnitsPerHour,proto3" json:"max_units_per_hour,omitempty"`
	// The number of units that can be spent on the one execution.
	MaxUnitsPerExecution *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=max_units_per_execution,json=maxUnitsPerExecution,proto3" json:"max_units_per_execution,omitempty"`
}

func (x *Project_Limits) Reset() {
	*x = Project_Limits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v2_project_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project_Limits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project_Limits) ProtoMessage() {}

func (x *Project_Limits) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v2_project_proto_msgTypes[2]
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
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP(), []int{0, 1}
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

var File_yandex_cloud_datasphere_v2_project_proto protoreflect.FileDescriptor

var file_yandex_cloud_datasphere_v2_project_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68,
	0x65, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x0d, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x05,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76,
	0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x22, 0x0a,
	0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x48, 0x0a, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x08, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x06, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x52, 0x06, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x74, 0x79,
	0x49, 0x64, 0x1a, 0xd3, 0x08, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12,
	0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72,
	0x6f, 0x63, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x58, 0x0a, 0x0b, 0x63,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x10, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x64, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x65, 0x61, 0x72, 0x6c, 0x79, 0x5f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x61, 0x72, 0x6c, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x42, 0x0a, 0x03, 0x69, 0x64, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32,
	0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x2e, 0x49, 0x64, 0x65, 0x52, 0x03, 0x69, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x64, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x46, 0x6f,
	0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x7d, 0x0a, 0x17, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x5f,
	0x65, 0x78, 0x65, 0x63, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x6f, 0x64,
	0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x46, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72,
	0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x53, 0x74, 0x61, 0x6c, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x52,
	0x14, 0x73, 0x74, 0x61, 0x6c, 0x65, 0x45, 0x78, 0x65, 0x63, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x6b, 0x0a, 0x12, 0x69, 0x64, 0x65, 0x5f, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x3d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x2e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x49, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f, 0x64, 0x65,
	0x52, 0x10, 0x69, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x6f,
	0x64, 0x65, 0x12, 0x4d, 0x0a, 0x15, 0x76, 0x6d, 0x5f, 0x69, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x76, 0x6d,
	0x49, 0x6e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x12, 0x34, 0x0a, 0x16, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x64, 0x65, 0x64,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x14, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x44, 0x65, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x64, 0x53, 0x70, 0x65, 0x63, 0x22, 0x41, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x69,
	0x74, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x4f, 0x4d, 0x4d, 0x49, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x4e, 0x44, 0x41, 0x52, 0x44, 0x10, 0x01,
	0x12, 0x08, 0x0a, 0x04, 0x41, 0x55, 0x54, 0x4f, 0x10, 0x02, 0x22, 0x2b, 0x0a, 0x03, 0x49, 0x64,
	0x65, 0x12, 0x13, 0x0a, 0x0f, 0x49, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4a, 0x55, 0x50, 0x59, 0x54, 0x45,
	0x52, 0x5f, 0x4c, 0x41, 0x42, 0x10, 0x01, 0x22, 0x78, 0x0a, 0x19, 0x53, 0x74, 0x61, 0x6c, 0x65,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74,
	0x4d, 0x6f, 0x64, 0x65, 0x12, 0x2c, 0x0a, 0x28, 0x53, 0x54, 0x41, 0x4c, 0x45, 0x5f, 0x45, 0x58,
	0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f,
	0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x4e, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x10, 0x01,
	0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x48, 0x52, 0x45, 0x45, 0x5f, 0x48, 0x4f, 0x55, 0x52, 0x53, 0x10,
	0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x4e, 0x4f, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x10,
	0x03, 0x22, 0x55, 0x0a, 0x10, 0x49, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x44, 0x45, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x53, 0x45, 0x52,
	0x56, 0x45, 0x52, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x44,
	0x49, 0x43, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x1a, 0xa6, 0x01, 0x0a, 0x06, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x73, 0x12, 0x48, 0x0a, 0x12, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x75, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x6d, 0x61,
	0x78, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x48, 0x6f, 0x75, 0x72, 0x12, 0x52, 0x0a,
	0x17, 0x6d, 0x61, 0x78, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x14, 0x6d, 0x61, 0x78,
	0x55, 0x6e, 0x69, 0x74, 0x73, 0x50, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08, 0x0a,
	0x10, 0x0b, 0x4a, 0x04, 0x08, 0x09, 0x10, 0x0a, 0x42, 0x6b, 0x0a, 0x1e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x32, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x32, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datasphere_v2_project_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v2_project_proto_rawDescData = file_yandex_cloud_datasphere_v2_project_proto_rawDesc
)

func file_yandex_cloud_datasphere_v2_project_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v2_project_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v2_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datasphere_v2_project_proto_rawDescData)
	})
	return file_yandex_cloud_datasphere_v2_project_proto_rawDescData
}

var file_yandex_cloud_datasphere_v2_project_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_yandex_cloud_datasphere_v2_project_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_datasphere_v2_project_proto_goTypes = []any{
	(Project_Settings_CommitMode)(0),                // 0: yandex.cloud.datasphere.v2.Project.Settings.CommitMode
	(Project_Settings_Ide)(0),                       // 1: yandex.cloud.datasphere.v2.Project.Settings.Ide
	(Project_Settings_StaleExecutionTimeoutMode)(0), // 2: yandex.cloud.datasphere.v2.Project.Settings.StaleExecutionTimeoutMode
	(Project_Settings_IdeExecutionMode)(0),          // 3: yandex.cloud.datasphere.v2.Project.Settings.IdeExecutionMode
	(*Project)(nil),                                 // 4: yandex.cloud.datasphere.v2.Project
	(*Project_Settings)(nil),                        // 5: yandex.cloud.datasphere.v2.Project.Settings
	(*Project_Limits)(nil),                          // 6: yandex.cloud.datasphere.v2.Project.Limits
	nil,                                             // 7: yandex.cloud.datasphere.v2.Project.LabelsEntry
	(*timestamppb.Timestamp)(nil),                   // 8: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),                     // 9: google.protobuf.Duration
	(*wrapperspb.Int64Value)(nil),                   // 10: google.protobuf.Int64Value
}
var file_yandex_cloud_datasphere_v2_project_proto_depIdxs = []int32{
	8,  // 0: yandex.cloud.datasphere.v2.Project.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: yandex.cloud.datasphere.v2.Project.labels:type_name -> yandex.cloud.datasphere.v2.Project.LabelsEntry
	5,  // 2: yandex.cloud.datasphere.v2.Project.settings:type_name -> yandex.cloud.datasphere.v2.Project.Settings
	6,  // 3: yandex.cloud.datasphere.v2.Project.limits:type_name -> yandex.cloud.datasphere.v2.Project.Limits
	0,  // 4: yandex.cloud.datasphere.v2.Project.Settings.commit_mode:type_name -> yandex.cloud.datasphere.v2.Project.Settings.CommitMode
	1,  // 5: yandex.cloud.datasphere.v2.Project.Settings.ide:type_name -> yandex.cloud.datasphere.v2.Project.Settings.Ide
	2,  // 6: yandex.cloud.datasphere.v2.Project.Settings.stale_exec_timeout_mode:type_name -> yandex.cloud.datasphere.v2.Project.Settings.StaleExecutionTimeoutMode
	3,  // 7: yandex.cloud.datasphere.v2.Project.Settings.ide_execution_mode:type_name -> yandex.cloud.datasphere.v2.Project.Settings.IdeExecutionMode
	9,  // 8: yandex.cloud.datasphere.v2.Project.Settings.vm_inactivity_timeout:type_name -> google.protobuf.Duration
	10, // 9: yandex.cloud.datasphere.v2.Project.Limits.max_units_per_hour:type_name -> google.protobuf.Int64Value
	10, // 10: yandex.cloud.datasphere.v2.Project.Limits.max_units_per_execution:type_name -> google.protobuf.Int64Value
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v2_project_proto_init() }
func file_yandex_cloud_datasphere_v2_project_proto_init() {
	if File_yandex_cloud_datasphere_v2_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datasphere_v2_project_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_yandex_cloud_datasphere_v2_project_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_yandex_cloud_datasphere_v2_project_proto_msgTypes[2].Exporter = func(v any, i int) any {
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
			RawDescriptor: file_yandex_cloud_datasphere_v2_project_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datasphere_v2_project_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v2_project_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_datasphere_v2_project_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_datasphere_v2_project_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v2_project_proto = out.File
	file_yandex_cloud_datasphere_v2_project_proto_rawDesc = nil
	file_yandex_cloud_datasphere_v2_project_proto_goTypes = nil
	file_yandex_cloud_datasphere_v2_project_proto_depIdxs = nil
}
