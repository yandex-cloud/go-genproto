// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/spqr/v1/cluster.proto

package spqr

import (
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Deployment environment.
type Cluster_Environment int32

const (
	Cluster_ENVIRONMENT_UNSPECIFIED Cluster_Environment = 0
	// Stable environment with a conservative update policy: only hotfixes
	// are applied during regular maintenance.
	Cluster_PRODUCTION Cluster_Environment = 1
	// Environment with more aggressive update policy: new versions
	// are rolled out irrespective of backward compatibility.
	Cluster_PRESTABLE Cluster_Environment = 2
)

// Enum value maps for Cluster_Environment.
var (
	Cluster_Environment_name = map[int32]string{
		0: "ENVIRONMENT_UNSPECIFIED",
		1: "PRODUCTION",
		2: "PRESTABLE",
	}
	Cluster_Environment_value = map[string]int32{
		"ENVIRONMENT_UNSPECIFIED": 0,
		"PRODUCTION":              1,
		"PRESTABLE":               2,
	}
)

func (x Cluster_Environment) Enum() *Cluster_Environment {
	p := new(Cluster_Environment)
	*p = x
	return p
}

func (x Cluster_Environment) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_Environment) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes[0].Descriptor()
}

func (Cluster_Environment) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes[0]
}

func (x Cluster_Environment) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Environment.Descriptor instead.
func (Cluster_Environment) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{0, 0}
}

type Cluster_Health int32

const (
	// State of the cluster is unknown ([Host.health] for every host in the cluster is UNKNOWN).
	Cluster_HEALTH_UNKNOWN Cluster_Health = 0
	// Cluster is alive and well ([Host.health] for every host in the cluster is ALIVE).
	Cluster_ALIVE Cluster_Health = 1
	// Cluster is inoperable ([Host.health] for every host in the cluster is DEAD).
	Cluster_DEAD Cluster_Health = 2
	// Cluster is working below capacity ([Host.health] for at least one host in the cluster is not ALIVE).
	Cluster_DEGRADED Cluster_Health = 3
)

// Enum value maps for Cluster_Health.
var (
	Cluster_Health_name = map[int32]string{
		0: "HEALTH_UNKNOWN",
		1: "ALIVE",
		2: "DEAD",
		3: "DEGRADED",
	}
	Cluster_Health_value = map[string]int32{
		"HEALTH_UNKNOWN": 0,
		"ALIVE":          1,
		"DEAD":           2,
		"DEGRADED":       3,
	}
)

func (x Cluster_Health) Enum() *Cluster_Health {
	p := new(Cluster_Health)
	*p = x
	return p
}

func (x Cluster_Health) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_Health) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes[1].Descriptor()
}

func (Cluster_Health) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes[1]
}

func (x Cluster_Health) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Health.Descriptor instead.
func (Cluster_Health) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{0, 1}
}

type Cluster_Status int32

const (
	// Cluster state is unknown.
	Cluster_STATUS_UNKNOWN Cluster_Status = 0
	// Cluster is being created.
	Cluster_CREATING Cluster_Status = 1
	// Cluster is running normally.
	Cluster_RUNNING Cluster_Status = 2
	// Cluster encountered a problem and cannot operate.
	Cluster_ERROR Cluster_Status = 3
	// Cluster is being updated.
	Cluster_UPDATING Cluster_Status = 4
	// Cluster is stopping.
	Cluster_STOPPING Cluster_Status = 5
	// Cluster stopped.
	Cluster_STOPPED Cluster_Status = 6
	// Cluster is starting.
	Cluster_STARTING Cluster_Status = 7
)

// Enum value maps for Cluster_Status.
var (
	Cluster_Status_name = map[int32]string{
		0: "STATUS_UNKNOWN",
		1: "CREATING",
		2: "RUNNING",
		3: "ERROR",
		4: "UPDATING",
		5: "STOPPING",
		6: "STOPPED",
		7: "STARTING",
	}
	Cluster_Status_value = map[string]int32{
		"STATUS_UNKNOWN": 0,
		"CREATING":       1,
		"RUNNING":        2,
		"ERROR":          3,
		"UPDATING":       4,
		"STOPPING":       5,
		"STOPPED":        6,
		"STARTING":       7,
	}
)

func (x Cluster_Status) Enum() *Cluster_Status {
	p := new(Cluster_Status)
	*p = x
	return p
}

func (x Cluster_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cluster_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes[2].Descriptor()
}

func (Cluster_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes[2]
}

func (x Cluster_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Status.Descriptor instead.
func (Cluster_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{0, 2}
}

// A managed SPQR cluster. For more information, see the [documentation](/docs/managed-spqr/concepts).
type Cluster struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the SPQR cluster.
	// This ID is assigned by MDB at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the SPQR cluster belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the SPQR cluster.
	// The name is unique within the folder. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the SPQR cluster. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Custom labels for the SPQR cluster as “ key:value “ pairs. Maximum 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Deployment environment of the SPQR cluster.
	Environment Cluster_Environment `protobuf:"varint,7,opt,name=environment,proto3,enum=yandex.cloud.mdb.spqr.v1.Cluster_Environment" json:"environment,omitempty"`
	// Description of monitoring systems relevant to the SPQR cluster.
	Monitoring []*Monitoring `protobuf:"bytes,8,rep,name=monitoring,proto3" json:"monitoring,omitempty"`
	// Configuration of the SPQR cluster.
	Config *ClusterConfig `protobuf:"bytes,9,opt,name=config,proto3" json:"config,omitempty"`
	// ID of the network that the cluster belongs to.
	NetworkId string `protobuf:"bytes,10,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// Aggregated cluster health.
	Health Cluster_Health `protobuf:"varint,11,opt,name=health,proto3,enum=yandex.cloud.mdb.spqr.v1.Cluster_Health" json:"health,omitempty"`
	// Current state of the cluster.
	Status Cluster_Status `protobuf:"varint,12,opt,name=status,proto3,enum=yandex.cloud.mdb.spqr.v1.Cluster_Status" json:"status,omitempty"`
	// Maintenance window for the cluster.
	MaintenanceWindow *MaintenanceWindow `protobuf:"bytes,13,opt,name=maintenance_window,json=maintenanceWindow,proto3" json:"maintenance_window,omitempty"`
	// Planned maintenance operation to be started for the cluster within the nearest [maintenance_window].
	PlannedOperation *MaintenanceOperation `protobuf:"bytes,14,opt,name=planned_operation,json=plannedOperation,proto3" json:"planned_operation,omitempty"`
	// User security groups
	SecurityGroupIds []string `protobuf:"bytes,15,rep,name=security_group_ids,json=securityGroupIds,proto3" json:"security_group_ids,omitempty"`
	// Deletion Protection inhibits deletion of the cluster
	DeletionProtection bool `protobuf:"varint,16,opt,name=deletion_protection,json=deletionProtection,proto3" json:"deletion_protection,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cluster) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Cluster) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Cluster) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Cluster) GetEnvironment() Cluster_Environment {
	if x != nil {
		return x.Environment
	}
	return Cluster_ENVIRONMENT_UNSPECIFIED
}

func (x *Cluster) GetMonitoring() []*Monitoring {
	if x != nil {
		return x.Monitoring
	}
	return nil
}

func (x *Cluster) GetConfig() *ClusterConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Cluster) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Cluster) GetHealth() Cluster_Health {
	if x != nil {
		return x.Health
	}
	return Cluster_HEALTH_UNKNOWN
}

func (x *Cluster) GetStatus() Cluster_Status {
	if x != nil {
		return x.Status
	}
	return Cluster_STATUS_UNKNOWN
}

func (x *Cluster) GetMaintenanceWindow() *MaintenanceWindow {
	if x != nil {
		return x.MaintenanceWindow
	}
	return nil
}

func (x *Cluster) GetPlannedOperation() *MaintenanceOperation {
	if x != nil {
		return x.PlannedOperation
	}
	return nil
}

func (x *Cluster) GetSecurityGroupIds() []string {
	if x != nil {
		return x.SecurityGroupIds
	}
	return nil
}

func (x *Cluster) GetDeletionProtection() bool {
	if x != nil {
		return x.DeletionProtection
	}
	return false
}

// Monitoring system.
type Monitoring struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the monitoring system.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the monitoring system.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Link to the monitoring system charts for the SPQR cluster.
	Link          string `protobuf:"bytes,3,opt,name=link,proto3" json:"link,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Monitoring) Reset() {
	*x = Monitoring{}
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Monitoring) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Monitoring) ProtoMessage() {}

func (x *Monitoring) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Monitoring.ProtoReflect.Descriptor instead.
func (*Monitoring) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *Monitoring) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Monitoring) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Monitoring) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

type ClusterConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Configuration for SPQR servers in the cluster.
	SpqrConfig *SPQRConfig `protobuf:"bytes,1,opt,name=spqr_config,json=spqrConfig,proto3" json:"spqr_config,omitempty"`
	// Time to start the daily backup, in the UTC timezone.
	BackupWindowStart *timeofday.TimeOfDay `protobuf:"bytes,2,opt,name=backup_window_start,json=backupWindowStart,proto3" json:"backup_window_start,omitempty"`
	// Retain period of automatically created backup in days
	BackupRetainPeriodDays *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=backup_retain_period_days,json=backupRetainPeriodDays,proto3" json:"backup_retain_period_days,omitempty"`
	// Access policy to DB
	Access        *Access               `protobuf:"bytes,4,opt,name=access,proto3" json:"access,omitempty"`
	SoxAudit      *wrapperspb.BoolValue `protobuf:"bytes,5,opt,name=sox_audit,json=soxAudit,proto3" json:"sox_audit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterConfig.ProtoReflect.Descriptor instead.
func (*ClusterConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterConfig) GetSpqrConfig() *SPQRConfig {
	if x != nil {
		return x.SpqrConfig
	}
	return nil
}

func (x *ClusterConfig) GetBackupWindowStart() *timeofday.TimeOfDay {
	if x != nil {
		return x.BackupWindowStart
	}
	return nil
}

func (x *ClusterConfig) GetBackupRetainPeriodDays() *wrapperspb.Int64Value {
	if x != nil {
		return x.BackupRetainPeriodDays
	}
	return nil
}

func (x *ClusterConfig) GetAccess() *Access {
	if x != nil {
		return x.Access
	}
	return nil
}

func (x *ClusterConfig) GetSoxAudit() *wrapperspb.BoolValue {
	if x != nil {
		return x.SoxAudit
	}
	return nil
}

type Access struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Allow access for DataLens.
	DataLens bool `protobuf:"varint,1,opt,name=data_lens,json=dataLens,proto3" json:"data_lens,omitempty"`
	// Allow access for Web SQL.
	WebSql bool `protobuf:"varint,2,opt,name=web_sql,json=webSql,proto3" json:"web_sql,omitempty"`
	// Allow access for DataTransfer.
	DataTransfer bool `protobuf:"varint,3,opt,name=data_transfer,json=dataTransfer,proto3" json:"data_transfer,omitempty"`
	// Allow access for Serverless.
	// NOTE: Do not propagate to public API until Serverless integration is required.
	Serverless    bool `protobuf:"varint,4,opt,name=serverless,proto3" json:"serverless,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Access) Reset() {
	*x = Access{}
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Access) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Access) ProtoMessage() {}

func (x *Access) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Access.ProtoReflect.Descriptor instead.
func (*Access) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Access) GetDataLens() bool {
	if x != nil {
		return x.DataLens
	}
	return false
}

func (x *Access) GetWebSql() bool {
	if x != nil {
		return x.WebSql
	}
	return false
}

func (x *Access) GetDataTransfer() bool {
	if x != nil {
		return x.DataTransfer
	}
	return false
}

func (x *Access) GetServerless() bool {
	if x != nil {
		return x.Serverless
	}
	return false
}

var File_yandex_cloud_mdb_spqr_v1_cluster_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDesc = "" +
	"\n" +
	"&yandex/cloud/mdb/spqr/v1/cluster.proto\x12\x18yandex.cloud.mdb.spqr.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1bgoogle/type/timeofday.proto\x1a%yandex/cloud/mdb/spqr/v1/config.proto\x1a*yandex/cloud/mdb/spqr/v1/maintenance.proto\"\xc3\t\n" +
	"\aCluster\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12E\n" +
	"\x06labels\x18\x06 \x03(\v2-.yandex.cloud.mdb.spqr.v1.Cluster.LabelsEntryR\x06labels\x12O\n" +
	"\venvironment\x18\a \x01(\x0e2-.yandex.cloud.mdb.spqr.v1.Cluster.EnvironmentR\venvironment\x12D\n" +
	"\n" +
	"monitoring\x18\b \x03(\v2$.yandex.cloud.mdb.spqr.v1.MonitoringR\n" +
	"monitoring\x12?\n" +
	"\x06config\x18\t \x01(\v2'.yandex.cloud.mdb.spqr.v1.ClusterConfigR\x06config\x12\x1d\n" +
	"\n" +
	"network_id\x18\n" +
	" \x01(\tR\tnetworkId\x12@\n" +
	"\x06health\x18\v \x01(\x0e2(.yandex.cloud.mdb.spqr.v1.Cluster.HealthR\x06health\x12@\n" +
	"\x06status\x18\f \x01(\x0e2(.yandex.cloud.mdb.spqr.v1.Cluster.StatusR\x06status\x12Z\n" +
	"\x12maintenance_window\x18\r \x01(\v2+.yandex.cloud.mdb.spqr.v1.MaintenanceWindowR\x11maintenanceWindow\x12[\n" +
	"\x11planned_operation\x18\x0e \x01(\v2..yandex.cloud.mdb.spqr.v1.MaintenanceOperationR\x10plannedOperation\x12,\n" +
	"\x12security_group_ids\x18\x0f \x03(\tR\x10securityGroupIds\x12/\n" +
	"\x13deletion_protection\x18\x10 \x01(\bR\x12deletionProtection\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"I\n" +
	"\vEnvironment\x12\x1b\n" +
	"\x17ENVIRONMENT_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"PRODUCTION\x10\x01\x12\r\n" +
	"\tPRESTABLE\x10\x02\"?\n" +
	"\x06Health\x12\x12\n" +
	"\x0eHEALTH_UNKNOWN\x10\x00\x12\t\n" +
	"\x05ALIVE\x10\x01\x12\b\n" +
	"\x04DEAD\x10\x02\x12\f\n" +
	"\bDEGRADED\x10\x03\"y\n" +
	"\x06Status\x12\x12\n" +
	"\x0eSTATUS_UNKNOWN\x10\x00\x12\f\n" +
	"\bCREATING\x10\x01\x12\v\n" +
	"\aRUNNING\x10\x02\x12\t\n" +
	"\x05ERROR\x10\x03\x12\f\n" +
	"\bUPDATING\x10\x04\x12\f\n" +
	"\bSTOPPING\x10\x05\x12\v\n" +
	"\aSTOPPED\x10\x06\x12\f\n" +
	"\bSTARTING\x10\a\"V\n" +
	"\n" +
	"Monitoring\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x02 \x01(\tR\vdescription\x12\x12\n" +
	"\x04link\x18\x03 \x01(\tR\x04link\"\xe9\x02\n" +
	"\rClusterConfig\x12E\n" +
	"\vspqr_config\x18\x01 \x01(\v2$.yandex.cloud.mdb.spqr.v1.SPQRConfigR\n" +
	"spqrConfig\x12F\n" +
	"\x13backup_window_start\x18\x02 \x01(\v2\x16.google.type.TimeOfDayR\x11backupWindowStart\x12V\n" +
	"\x19backup_retain_period_days\x18\x03 \x01(\v2\x1b.google.protobuf.Int64ValueR\x16backupRetainPeriodDays\x128\n" +
	"\x06access\x18\x04 \x01(\v2 .yandex.cloud.mdb.spqr.v1.AccessR\x06access\x127\n" +
	"\tsox_audit\x18\x05 \x01(\v2\x1a.google.protobuf.BoolValueR\bsoxAudit\"\x83\x01\n" +
	"\x06Access\x12\x1b\n" +
	"\tdata_lens\x18\x01 \x01(\bR\bdataLens\x12\x17\n" +
	"\aweb_sql\x18\x02 \x01(\bR\x06webSql\x12#\n" +
	"\rdata_transfer\x18\x03 \x01(\bR\fdataTransfer\x12\x1e\n" +
	"\n" +
	"serverless\x18\x04 \x01(\bR\n" +
	"serverlessBa\n" +
	"\x1cyandex.cloud.api.mdb.spqr.v1ZAgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/spqr/v1;spqrb\x06proto3"

var (
	file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDesc), len(file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDescData
}

var file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_mdb_spqr_v1_cluster_proto_goTypes = []any{
	(Cluster_Environment)(0),      // 0: yandex.cloud.mdb.spqr.v1.Cluster.Environment
	(Cluster_Health)(0),           // 1: yandex.cloud.mdb.spqr.v1.Cluster.Health
	(Cluster_Status)(0),           // 2: yandex.cloud.mdb.spqr.v1.Cluster.Status
	(*Cluster)(nil),               // 3: yandex.cloud.mdb.spqr.v1.Cluster
	(*Monitoring)(nil),            // 4: yandex.cloud.mdb.spqr.v1.Monitoring
	(*ClusterConfig)(nil),         // 5: yandex.cloud.mdb.spqr.v1.ClusterConfig
	(*Access)(nil),                // 6: yandex.cloud.mdb.spqr.v1.Access
	nil,                           // 7: yandex.cloud.mdb.spqr.v1.Cluster.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*MaintenanceWindow)(nil),     // 9: yandex.cloud.mdb.spqr.v1.MaintenanceWindow
	(*MaintenanceOperation)(nil),  // 10: yandex.cloud.mdb.spqr.v1.MaintenanceOperation
	(*SPQRConfig)(nil),            // 11: yandex.cloud.mdb.spqr.v1.SPQRConfig
	(*timeofday.TimeOfDay)(nil),   // 12: google.type.TimeOfDay
	(*wrapperspb.Int64Value)(nil), // 13: google.protobuf.Int64Value
	(*wrapperspb.BoolValue)(nil),  // 14: google.protobuf.BoolValue
}
var file_yandex_cloud_mdb_spqr_v1_cluster_proto_depIdxs = []int32{
	8,  // 0: yandex.cloud.mdb.spqr.v1.Cluster.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: yandex.cloud.mdb.spqr.v1.Cluster.labels:type_name -> yandex.cloud.mdb.spqr.v1.Cluster.LabelsEntry
	0,  // 2: yandex.cloud.mdb.spqr.v1.Cluster.environment:type_name -> yandex.cloud.mdb.spqr.v1.Cluster.Environment
	4,  // 3: yandex.cloud.mdb.spqr.v1.Cluster.monitoring:type_name -> yandex.cloud.mdb.spqr.v1.Monitoring
	5,  // 4: yandex.cloud.mdb.spqr.v1.Cluster.config:type_name -> yandex.cloud.mdb.spqr.v1.ClusterConfig
	1,  // 5: yandex.cloud.mdb.spqr.v1.Cluster.health:type_name -> yandex.cloud.mdb.spqr.v1.Cluster.Health
	2,  // 6: yandex.cloud.mdb.spqr.v1.Cluster.status:type_name -> yandex.cloud.mdb.spqr.v1.Cluster.Status
	9,  // 7: yandex.cloud.mdb.spqr.v1.Cluster.maintenance_window:type_name -> yandex.cloud.mdb.spqr.v1.MaintenanceWindow
	10, // 8: yandex.cloud.mdb.spqr.v1.Cluster.planned_operation:type_name -> yandex.cloud.mdb.spqr.v1.MaintenanceOperation
	11, // 9: yandex.cloud.mdb.spqr.v1.ClusterConfig.spqr_config:type_name -> yandex.cloud.mdb.spqr.v1.SPQRConfig
	12, // 10: yandex.cloud.mdb.spqr.v1.ClusterConfig.backup_window_start:type_name -> google.type.TimeOfDay
	13, // 11: yandex.cloud.mdb.spqr.v1.ClusterConfig.backup_retain_period_days:type_name -> google.protobuf.Int64Value
	6,  // 12: yandex.cloud.mdb.spqr.v1.ClusterConfig.access:type_name -> yandex.cloud.mdb.spqr.v1.Access
	14, // 13: yandex.cloud.mdb.spqr.v1.ClusterConfig.sox_audit:type_name -> google.protobuf.BoolValue
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_spqr_v1_cluster_proto_init() }
func file_yandex_cloud_mdb_spqr_v1_cluster_proto_init() {
	if File_yandex_cloud_mdb_spqr_v1_cluster_proto != nil {
		return
	}
	file_yandex_cloud_mdb_spqr_v1_config_proto_init()
	file_yandex_cloud_mdb_spqr_v1_maintenance_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDesc), len(file_yandex_cloud_mdb_spqr_v1_cluster_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_spqr_v1_cluster_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_spqr_v1_cluster_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_spqr_v1_cluster_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_spqr_v1_cluster_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_spqr_v1_cluster_proto = out.File
	file_yandex_cloud_mdb_spqr_v1_cluster_proto_goTypes = nil
	file_yandex_cloud_mdb_spqr_v1_cluster_proto_depIdxs = nil
}
