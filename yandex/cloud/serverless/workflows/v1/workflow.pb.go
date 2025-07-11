// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/serverless/workflows/v1/workflow.proto

package workflows

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
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

type Workflow_Status int32

const (
	Workflow_STATUS_UNSPECIFIED Workflow_Status = 0
	// Workflow is being created.
	Workflow_CREATING Workflow_Status = 1
	// Workflow is ready for use.
	Workflow_ACTIVE Workflow_Status = 2
	// Workflow is being updated.
	Workflow_UPDATING Workflow_Status = 3
	// Workflow is being deleted.
	Workflow_DELETING Workflow_Status = 4
	// Workflow failed. The only allowed action is delete.
	Workflow_ERROR Workflow_Status = 5
)

// Enum value maps for Workflow_Status.
var (
	Workflow_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "UPDATING",
		4: "DELETING",
		5: "ERROR",
	}
	Workflow_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"UPDATING":           3,
		"DELETING":           4,
		"ERROR":              5,
	}
)

func (x Workflow_Status) Enum() *Workflow_Status {
	p := new(Workflow_Status)
	*p = x
	return p
}

func (x Workflow_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Workflow_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_enumTypes[0].Descriptor()
}

func (Workflow_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_serverless_workflows_v1_workflow_proto_enumTypes[0]
}

func (x Workflow_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Workflow_Status.Descriptor instead.
func (Workflow_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescGZIP(), []int{0, 0}
}

type Workflow struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Workflow. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the Workflow belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Specification of the Workflow
	Specification *WorkflowSpecification `protobuf:"bytes,3,opt,name=specification,proto3" json:"specification,omitempty"`
	// Creation timestamp for the Workflow.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the Workflow. The name is unique within the folder.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Workflow.
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	// Workflow labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,7,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Status of the Workflow.
	Status Workflow_Status `protobuf:"varint,8,opt,name=status,proto3,enum=yandex.cloud.serverless.workflows.v1.Workflow_Status" json:"status,omitempty"`
	// Options for logging from the Workflow.
	LogOptions *LogOptions `protobuf:"bytes,9,opt,name=log_options,json=logOptions,proto3" json:"log_options,omitempty"`
	// ID of the VPC network Workflow will be executed in, in order to access private resources.
	NetworkId string `protobuf:"bytes,10,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// ID of the Service Account which will be used for resource access in Workflow execution.
	ServiceAccountId string `protobuf:"bytes,11,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Express execution mode.
	Express       bool `protobuf:"varint,12,opt,name=express,proto3" json:"express,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Workflow) Reset() {
	*x = Workflow{}
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Workflow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Workflow) ProtoMessage() {}

func (x *Workflow) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Workflow.ProtoReflect.Descriptor instead.
func (*Workflow) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescGZIP(), []int{0}
}

func (x *Workflow) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Workflow) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Workflow) GetSpecification() *WorkflowSpecification {
	if x != nil {
		return x.Specification
	}
	return nil
}

func (x *Workflow) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Workflow) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Workflow) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Workflow) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Workflow) GetStatus() Workflow_Status {
	if x != nil {
		return x.Status
	}
	return Workflow_STATUS_UNSPECIFIED
}

func (x *Workflow) GetLogOptions() *LogOptions {
	if x != nil {
		return x.LogOptions
	}
	return nil
}

func (x *Workflow) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Workflow) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Workflow) GetExpress() bool {
	if x != nil {
		return x.Express
	}
	return false
}

type WorkflowPreview struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Workflow. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the Workflow belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp for the Workflow.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the Workflow. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Workflow.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Workflow labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Status of the Workflow.
	Status Workflow_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.serverless.workflows.v1.Workflow_Status" json:"status,omitempty"`
	// Options for logging from the Workflow.
	LogOptions *LogOptions `protobuf:"bytes,8,opt,name=log_options,json=logOptions,proto3" json:"log_options,omitempty"`
	// ID of the VPC network Workflow will be executed in, in order to access private resources.
	NetworkId string `protobuf:"bytes,9,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// ID of the Service Account which will be used for resources access in Workflow execution.
	ServiceAccountId string `protobuf:"bytes,10,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Express execution mode.
	Express       bool `protobuf:"varint,11,opt,name=express,proto3" json:"express,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowPreview) Reset() {
	*x = WorkflowPreview{}
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowPreview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowPreview) ProtoMessage() {}

func (x *WorkflowPreview) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowPreview.ProtoReflect.Descriptor instead.
func (*WorkflowPreview) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowPreview) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *WorkflowPreview) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *WorkflowPreview) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *WorkflowPreview) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkflowPreview) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *WorkflowPreview) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *WorkflowPreview) GetStatus() Workflow_Status {
	if x != nil {
		return x.Status
	}
	return Workflow_STATUS_UNSPECIFIED
}

func (x *WorkflowPreview) GetLogOptions() *LogOptions {
	if x != nil {
		return x.LogOptions
	}
	return nil
}

func (x *WorkflowPreview) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *WorkflowPreview) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *WorkflowPreview) GetExpress() bool {
	if x != nil {
		return x.Express
	}
	return false
}

type WorkflowSpecification struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Spec:
	//
	//	*WorkflowSpecification_SpecYaml
	Spec          isWorkflowSpecification_Spec `protobuf_oneof:"spec"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowSpecification) Reset() {
	*x = WorkflowSpecification{}
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowSpecification) ProtoMessage() {}

func (x *WorkflowSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowSpecification.ProtoReflect.Descriptor instead.
func (*WorkflowSpecification) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescGZIP(), []int{2}
}

func (x *WorkflowSpecification) GetSpec() isWorkflowSpecification_Spec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *WorkflowSpecification) GetSpecYaml() string {
	if x != nil {
		if x, ok := x.Spec.(*WorkflowSpecification_SpecYaml); ok {
			return x.SpecYaml
		}
	}
	return ""
}

type isWorkflowSpecification_Spec interface {
	isWorkflowSpecification_Spec()
}

type WorkflowSpecification_SpecYaml struct {
	// Workflow specification in YAML format.
	SpecYaml string `protobuf:"bytes,1,opt,name=spec_yaml,json=specYaml,proto3,oneof"`
}

func (*WorkflowSpecification_SpecYaml) isWorkflowSpecification_Spec() {}

type LogOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Is logging from Workflow disabled.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Types that are valid to be assigned to Destination:
	//
	//	*LogOptions_LogGroupId
	//	*LogOptions_FolderId
	Destination isLogOptions_Destination `protobuf_oneof:"destination"`
	// Minimum logs level.
	//
	// See [LogLevel.Level] for details.
	MinLevel      v1.LogLevel_Level `protobuf:"varint,4,opt,name=min_level,json=minLevel,proto3,enum=yandex.cloud.logging.v1.LogLevel_Level" json:"min_level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogOptions) Reset() {
	*x = LogOptions{}
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogOptions) ProtoMessage() {}

func (x *LogOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogOptions.ProtoReflect.Descriptor instead.
func (*LogOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescGZIP(), []int{3}
}

func (x *LogOptions) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *LogOptions) GetDestination() isLogOptions_Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *LogOptions) GetLogGroupId() string {
	if x != nil {
		if x, ok := x.Destination.(*LogOptions_LogGroupId); ok {
			return x.LogGroupId
		}
	}
	return ""
}

func (x *LogOptions) GetFolderId() string {
	if x != nil {
		if x, ok := x.Destination.(*LogOptions_FolderId); ok {
			return x.FolderId
		}
	}
	return ""
}

func (x *LogOptions) GetMinLevel() v1.LogLevel_Level {
	if x != nil {
		return x.MinLevel
	}
	return v1.LogLevel_Level(0)
}

type isLogOptions_Destination interface {
	isLogOptions_Destination()
}

type LogOptions_LogGroupId struct {
	// ID of the logging group which should be used for Workflows logs.
	LogGroupId string `protobuf:"bytes,2,opt,name=log_group_id,json=logGroupId,proto3,oneof"`
}

type LogOptions_FolderId struct {
	// ID of the folder which default logging group should be used for Workflows.
	FolderId string `protobuf:"bytes,3,opt,name=folder_id,json=folderId,proto3,oneof"`
}

func (*LogOptions_LogGroupId) isLogOptions_Destination() {}

func (*LogOptions_FolderId) isLogOptions_Destination() {}

var File_yandex_cloud_serverless_workflows_v1_workflow_proto protoreflect.FileDescriptor

const file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDesc = "" +
	"\n" +
	"3yandex/cloud/serverless/workflows/v1/workflow.proto\x12$yandex.cloud.serverless.workflows.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a'yandex/cloud/logging/v1/log_entry.proto\x1a\x1dyandex/cloud/validation.proto\"\x86\x06\n" +
	"\bWorkflow\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x12a\n" +
	"\rspecification\x18\x03 \x01(\v2;.yandex.cloud.serverless.workflows.v1.WorkflowSpecificationR\rspecification\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x05 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x06 \x01(\tR\vdescription\x12R\n" +
	"\x06labels\x18\a \x03(\v2:.yandex.cloud.serverless.workflows.v1.Workflow.LabelsEntryR\x06labels\x12M\n" +
	"\x06status\x18\b \x01(\x0e25.yandex.cloud.serverless.workflows.v1.Workflow.StatusR\x06status\x12Q\n" +
	"\vlog_options\x18\t \x01(\v20.yandex.cloud.serverless.workflows.v1.LogOptionsR\n" +
	"logOptions\x12\x1d\n" +
	"\n" +
	"network_id\x18\n" +
	" \x01(\tR\tnetworkId\x12,\n" +
	"\x12service_account_id\x18\v \x01(\tR\x10serviceAccountId\x12\x18\n" +
	"\aexpress\x18\f \x01(\bR\aexpress\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"a\n" +
	"\x06Status\x12\x16\n" +
	"\x12STATUS_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bCREATING\x10\x01\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\f\n" +
	"\bUPDATING\x10\x03\x12\f\n" +
	"\bDELETING\x10\x04\x12\t\n" +
	"\x05ERROR\x10\x05\"\xce\x04\n" +
	"\x0fWorkflowPreview\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12Y\n" +
	"\x06labels\x18\x06 \x03(\v2A.yandex.cloud.serverless.workflows.v1.WorkflowPreview.LabelsEntryR\x06labels\x12M\n" +
	"\x06status\x18\a \x01(\x0e25.yandex.cloud.serverless.workflows.v1.Workflow.StatusR\x06status\x12Q\n" +
	"\vlog_options\x18\b \x01(\v20.yandex.cloud.serverless.workflows.v1.LogOptionsR\n" +
	"logOptions\x12\x1d\n" +
	"\n" +
	"network_id\x18\t \x01(\tR\tnetworkId\x12,\n" +
	"\x12service_account_id\x18\n" +
	" \x01(\tR\x10serviceAccountId\x12\x18\n" +
	"\aexpress\x18\v \x01(\bR\aexpress\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"D\n" +
	"\x15WorkflowSpecification\x12\x1d\n" +
	"\tspec_yaml\x18\x01 \x01(\tH\x00R\bspecYamlB\f\n" +
	"\x04spec\x12\x04\xc0\xc11\x01\"\xc0\x01\n" +
	"\n" +
	"LogOptions\x12\x1a\n" +
	"\bdisabled\x18\x01 \x01(\bR\bdisabled\x12\"\n" +
	"\flog_group_id\x18\x02 \x01(\tH\x00R\n" +
	"logGroupId\x12\x1d\n" +
	"\tfolder_id\x18\x03 \x01(\tH\x00R\bfolderId\x12D\n" +
	"\tmin_level\x18\x04 \x01(\x0e2'.yandex.cloud.logging.v1.LogLevel.LevelR\bminLevelB\r\n" +
	"\vdestinationB~\n" +
	"(yandex.cloud.api.serverless.workflows.v1ZRgithub.com/yandex-cloud/go-genproto/yandex/cloud/serverless/workflows/v1;workflowsb\x06proto3"

var (
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescData []byte
)

func file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDesc), len(file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDesc)))
	})
	return file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDescData
}

var file_yandex_cloud_serverless_workflows_v1_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_serverless_workflows_v1_workflow_proto_goTypes = []any{
	(Workflow_Status)(0),          // 0: yandex.cloud.serverless.workflows.v1.Workflow.Status
	(*Workflow)(nil),              // 1: yandex.cloud.serverless.workflows.v1.Workflow
	(*WorkflowPreview)(nil),       // 2: yandex.cloud.serverless.workflows.v1.WorkflowPreview
	(*WorkflowSpecification)(nil), // 3: yandex.cloud.serverless.workflows.v1.WorkflowSpecification
	(*LogOptions)(nil),            // 4: yandex.cloud.serverless.workflows.v1.LogOptions
	nil,                           // 5: yandex.cloud.serverless.workflows.v1.Workflow.LabelsEntry
	nil,                           // 6: yandex.cloud.serverless.workflows.v1.WorkflowPreview.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(v1.LogLevel_Level)(0),        // 8: yandex.cloud.logging.v1.LogLevel.Level
}
var file_yandex_cloud_serverless_workflows_v1_workflow_proto_depIdxs = []int32{
	3,  // 0: yandex.cloud.serverless.workflows.v1.Workflow.specification:type_name -> yandex.cloud.serverless.workflows.v1.WorkflowSpecification
	7,  // 1: yandex.cloud.serverless.workflows.v1.Workflow.created_at:type_name -> google.protobuf.Timestamp
	5,  // 2: yandex.cloud.serverless.workflows.v1.Workflow.labels:type_name -> yandex.cloud.serverless.workflows.v1.Workflow.LabelsEntry
	0,  // 3: yandex.cloud.serverless.workflows.v1.Workflow.status:type_name -> yandex.cloud.serverless.workflows.v1.Workflow.Status
	4,  // 4: yandex.cloud.serverless.workflows.v1.Workflow.log_options:type_name -> yandex.cloud.serverless.workflows.v1.LogOptions
	7,  // 5: yandex.cloud.serverless.workflows.v1.WorkflowPreview.created_at:type_name -> google.protobuf.Timestamp
	6,  // 6: yandex.cloud.serverless.workflows.v1.WorkflowPreview.labels:type_name -> yandex.cloud.serverless.workflows.v1.WorkflowPreview.LabelsEntry
	0,  // 7: yandex.cloud.serverless.workflows.v1.WorkflowPreview.status:type_name -> yandex.cloud.serverless.workflows.v1.Workflow.Status
	4,  // 8: yandex.cloud.serverless.workflows.v1.WorkflowPreview.log_options:type_name -> yandex.cloud.serverless.workflows.v1.LogOptions
	8,  // 9: yandex.cloud.serverless.workflows.v1.LogOptions.min_level:type_name -> yandex.cloud.logging.v1.LogLevel.Level
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_workflows_v1_workflow_proto_init() }
func file_yandex_cloud_serverless_workflows_v1_workflow_proto_init() {
	if File_yandex_cloud_serverless_workflows_v1_workflow_proto != nil {
		return
	}
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[2].OneofWrappers = []any{
		(*WorkflowSpecification_SpecYaml)(nil),
	}
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes[3].OneofWrappers = []any{
		(*LogOptions_LogGroupId)(nil),
		(*LogOptions_FolderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDesc), len(file_yandex_cloud_serverless_workflows_v1_workflow_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_workflows_v1_workflow_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_workflows_v1_workflow_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_serverless_workflows_v1_workflow_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_serverless_workflows_v1_workflow_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_workflows_v1_workflow_proto = out.File
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_goTypes = nil
	file_yandex_cloud_serverless_workflows_v1_workflow_proto_depIdxs = nil
}
