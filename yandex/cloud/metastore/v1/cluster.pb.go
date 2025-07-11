// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/metastore/v1/cluster.proto

package metastore

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
	return file_yandex_cloud_metastore_v1_cluster_proto_enumTypes[0].Descriptor()
}

func (Cluster_Health) Type() protoreflect.EnumType {
	return &file_yandex_cloud_metastore_v1_cluster_proto_enumTypes[0]
}

func (x Cluster_Health) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Health.Descriptor instead.
func (Cluster_Health) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{0, 0}
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
	return file_yandex_cloud_metastore_v1_cluster_proto_enumTypes[1].Descriptor()
}

func (Cluster_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_metastore_v1_cluster_proto_enumTypes[1]
}

func (x Cluster_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cluster_Status.Descriptor instead.
func (Cluster_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{0, 1}
}

// Hive Metastore Cluster.
type Cluster struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Metastore cluster.
	// This ID is assigned by MDB at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the Metastore cluster belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the Metastore cluster.
	// The name is unique within the folder. 1-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the Metastore cluster. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Custom labels for the Metastore cluster as “ key:value “ pairs.
	// Maximum 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Aggregated cluster health.
	Health Cluster_Health `protobuf:"varint,8,opt,name=health,proto3,enum=yandex.cloud.metastore.v1.Cluster_Health" json:"health,omitempty"`
	// Current state of the cluster.
	Status Cluster_Status `protobuf:"varint,9,opt,name=status,proto3,enum=yandex.cloud.metastore.v1.Cluster_Status" json:"status,omitempty"`
	// Deletion Protection inhibits deletion of the cluster
	DeletionProtection bool `protobuf:"varint,16,opt,name=deletion_protection,json=deletionProtection,proto3" json:"deletion_protection,omitempty"`
	// Metastore server version
	Version string `protobuf:"bytes,17,opt,name=version,proto3" json:"version,omitempty"`
	// Metastore network
	NetworkId string `protobuf:"bytes,18,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	// IP address of metastore server balancer endpoint
	EndpointIp string `protobuf:"bytes,19,opt,name=endpoint_ip,json=endpointIp,proto3" json:"endpoint_ip,omitempty"`
	// Metastore cluster configuration
	ClusterConfig *ClusterConfig `protobuf:"bytes,20,opt,name=cluster_config,json=clusterConfig,proto3" json:"cluster_config,omitempty"`
	// Service account that will be used to access a YC resources
	ServiceAccountId string `protobuf:"bytes,21,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Cloud logging configuration
	Logging *LoggingConfig `protobuf:"bytes,22,opt,name=logging,proto3" json:"logging,omitempty"`
	// Network related configuration options.
	Network *NetworkConfig `protobuf:"bytes,23,opt,name=network,proto3" json:"network,omitempty"`
	// Window of maintenance operations.
	MaintenanceWindow *MaintenanceWindow `protobuf:"bytes,24,opt,name=maintenance_window,json=maintenanceWindow,proto3" json:"maintenance_window,omitempty"`
	// Maintenance operation planned at nearest maintenance_window.
	PlannedOperation *MaintenanceOperation `protobuf:"bytes,25,opt,name=planned_operation,json=plannedOperation,proto3" json:"planned_operation,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[0]
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
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{0}
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

func (x *Cluster) GetDeletionProtection() bool {
	if x != nil {
		return x.DeletionProtection
	}
	return false
}

func (x *Cluster) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *Cluster) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Cluster) GetEndpointIp() string {
	if x != nil {
		return x.EndpointIp
	}
	return ""
}

func (x *Cluster) GetClusterConfig() *ClusterConfig {
	if x != nil {
		return x.ClusterConfig
	}
	return nil
}

func (x *Cluster) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Cluster) GetLogging() *LoggingConfig {
	if x != nil {
		return x.Logging
	}
	return nil
}

func (x *Cluster) GetNetwork() *NetworkConfig {
	if x != nil {
		return x.Network
	}
	return nil
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

type ClusterConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resources     *Resources             `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterConfig) Reset() {
	*x = ClusterConfig{}
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterConfig) ProtoMessage() {}

func (x *ClusterConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[1]
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
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterConfig) GetResources() *Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

type NetworkConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// IDs of VPC network subnets where instances of the cluster are attached.
	SubnetIds []string `protobuf:"bytes,1,rep,name=subnet_ids,json=subnetIds,proto3" json:"subnet_ids,omitempty"`
	// User security groups.
	SecurityGroupIds []string `protobuf:"bytes,2,rep,name=security_group_ids,json=securityGroupIds,proto3" json:"security_group_ids,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *NetworkConfig) Reset() {
	*x = NetworkConfig{}
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NetworkConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkConfig) ProtoMessage() {}

func (x *NetworkConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkConfig.ProtoReflect.Descriptor instead.
func (*NetworkConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkConfig) GetSubnetIds() []string {
	if x != nil {
		return x.SubnetIds
	}
	return nil
}

func (x *NetworkConfig) GetSecurityGroupIds() []string {
	if x != nil {
		return x.SecurityGroupIds
	}
	return nil
}

type Resources struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the preset for computational resources available to a pod (CPU, memory etc.).
	ResourcePresetId string `protobuf:"bytes,1,opt,name=resource_preset_id,json=resourcePresetId,proto3" json:"resource_preset_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Resources) Reset() {
	*x = Resources{}
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Resources) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resources) ProtoMessage() {}

func (x *Resources) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resources.ProtoReflect.Descriptor instead.
func (*Resources) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{3}
}

func (x *Resources) GetResourcePresetId() string {
	if x != nil {
		return x.ResourcePresetId
	}
	return ""
}

type LoggingConfig struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Enabled bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Types that are valid to be assigned to Destination:
	//
	//	*LoggingConfig_FolderId
	//	*LoggingConfig_LogGroupId
	Destination   isLoggingConfig_Destination `protobuf_oneof:"destination"`
	MinLevel      v1.LogLevel_Level           `protobuf:"varint,4,opt,name=min_level,json=minLevel,proto3,enum=yandex.cloud.logging.v1.LogLevel_Level" json:"min_level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LoggingConfig) Reset() {
	*x = LoggingConfig{}
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoggingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoggingConfig) ProtoMessage() {}

func (x *LoggingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoggingConfig.ProtoReflect.Descriptor instead.
func (*LoggingConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP(), []int{4}
}

func (x *LoggingConfig) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *LoggingConfig) GetDestination() isLoggingConfig_Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *LoggingConfig) GetFolderId() string {
	if x != nil {
		if x, ok := x.Destination.(*LoggingConfig_FolderId); ok {
			return x.FolderId
		}
	}
	return ""
}

func (x *LoggingConfig) GetLogGroupId() string {
	if x != nil {
		if x, ok := x.Destination.(*LoggingConfig_LogGroupId); ok {
			return x.LogGroupId
		}
	}
	return ""
}

func (x *LoggingConfig) GetMinLevel() v1.LogLevel_Level {
	if x != nil {
		return x.MinLevel
	}
	return v1.LogLevel_Level(0)
}

type isLoggingConfig_Destination interface {
	isLoggingConfig_Destination()
}

type LoggingConfig_FolderId struct {
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3,oneof"`
}

type LoggingConfig_LogGroupId struct {
	LogGroupId string `protobuf:"bytes,3,opt,name=log_group_id,json=logGroupId,proto3,oneof"`
}

func (*LoggingConfig_FolderId) isLoggingConfig_Destination() {}

func (*LoggingConfig_LogGroupId) isLoggingConfig_Destination() {}

var File_yandex_cloud_metastore_v1_cluster_proto protoreflect.FileDescriptor

const file_yandex_cloud_metastore_v1_cluster_proto_rawDesc = "" +
	"\n" +
	"'yandex/cloud/metastore/v1/cluster.proto\x12\x19yandex.cloud.metastore.v1\x1a+yandex/cloud/metastore/v1/maintenance.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a'yandex/cloud/logging/v1/log_entry.proto\x1a\x1dyandex/cloud/validation.proto\"\xcf\t\n" +
	"\aCluster\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12F\n" +
	"\x06labels\x18\x06 \x03(\v2..yandex.cloud.metastore.v1.Cluster.LabelsEntryR\x06labels\x12A\n" +
	"\x06health\x18\b \x01(\x0e2).yandex.cloud.metastore.v1.Cluster.HealthR\x06health\x12A\n" +
	"\x06status\x18\t \x01(\x0e2).yandex.cloud.metastore.v1.Cluster.StatusR\x06status\x12/\n" +
	"\x13deletion_protection\x18\x10 \x01(\bR\x12deletionProtection\x12\x18\n" +
	"\aversion\x18\x11 \x01(\tR\aversion\x12\x1d\n" +
	"\n" +
	"network_id\x18\x12 \x01(\tR\tnetworkId\x12\x1f\n" +
	"\vendpoint_ip\x18\x13 \x01(\tR\n" +
	"endpointIp\x12O\n" +
	"\x0ecluster_config\x18\x14 \x01(\v2(.yandex.cloud.metastore.v1.ClusterConfigR\rclusterConfig\x126\n" +
	"\x12service_account_id\x18\x15 \x01(\tB\b\x8a\xc81\x04<=50R\x10serviceAccountId\x12B\n" +
	"\alogging\x18\x16 \x01(\v2(.yandex.cloud.metastore.v1.LoggingConfigR\alogging\x12B\n" +
	"\anetwork\x18\x17 \x01(\v2(.yandex.cloud.metastore.v1.NetworkConfigR\anetwork\x12[\n" +
	"\x12maintenance_window\x18\x18 \x01(\v2,.yandex.cloud.metastore.v1.MaintenanceWindowR\x11maintenanceWindow\x12\\\n" +
	"\x11planned_operation\x18\x19 \x01(\v2/.yandex.cloud.metastore.v1.MaintenanceOperationR\x10plannedOperation\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"?\n" +
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
	"\bSTARTING\x10\aJ\x04\b\a\x10\bJ\x04\b\n" +
	"\x10\x10\"Y\n" +
	"\rClusterConfig\x12B\n" +
	"\tresources\x18\x02 \x01(\v2$.yandex.cloud.metastore.v1.ResourcesR\tresourcesJ\x04\b\x01\x10\x02\"\\\n" +
	"\rNetworkConfig\x12\x1d\n" +
	"\n" +
	"subnet_ids\x18\x01 \x03(\tR\tsubnetIds\x12,\n" +
	"\x12security_group_ids\x18\x02 \x03(\tR\x10securityGroupIds\"9\n" +
	"\tResources\x12,\n" +
	"\x12resource_preset_id\x18\x01 \x01(\tR\x10resourcePresetId\"\x8b\x02\n" +
	"\rLoggingConfig\x12\x18\n" +
	"\aenabled\x18\x01 \x01(\bR\aenabled\x12B\n" +
	"\tfolder_id\x18\x02 \x01(\tB#\xf2\xc71\x1f([a-zA-Z][-a-zA-Z0-9_.]{0,63})?H\x00R\bfolderId\x12G\n" +
	"\flog_group_id\x18\x03 \x01(\tB#\xf2\xc71\x1f([a-zA-Z][-a-zA-Z0-9_.]{0,63})?H\x00R\n" +
	"logGroupId\x12D\n" +
	"\tmin_level\x18\x04 \x01(\x0e2'.yandex.cloud.logging.v1.LogLevel.LevelR\bminLevelB\r\n" +
	"\vdestinationBh\n" +
	"\x1dyandex.cloud.api.metastore.v1ZGgithub.com/yandex-cloud/go-genproto/yandex/cloud/metastore/v1;metastoreb\x06proto3"

var (
	file_yandex_cloud_metastore_v1_cluster_proto_rawDescOnce sync.Once
	file_yandex_cloud_metastore_v1_cluster_proto_rawDescData []byte
)

func file_yandex_cloud_metastore_v1_cluster_proto_rawDescGZIP() []byte {
	file_yandex_cloud_metastore_v1_cluster_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_metastore_v1_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_metastore_v1_cluster_proto_rawDesc), len(file_yandex_cloud_metastore_v1_cluster_proto_rawDesc)))
	})
	return file_yandex_cloud_metastore_v1_cluster_proto_rawDescData
}

var file_yandex_cloud_metastore_v1_cluster_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_metastore_v1_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_metastore_v1_cluster_proto_goTypes = []any{
	(Cluster_Health)(0),           // 0: yandex.cloud.metastore.v1.Cluster.Health
	(Cluster_Status)(0),           // 1: yandex.cloud.metastore.v1.Cluster.Status
	(*Cluster)(nil),               // 2: yandex.cloud.metastore.v1.Cluster
	(*ClusterConfig)(nil),         // 3: yandex.cloud.metastore.v1.ClusterConfig
	(*NetworkConfig)(nil),         // 4: yandex.cloud.metastore.v1.NetworkConfig
	(*Resources)(nil),             // 5: yandex.cloud.metastore.v1.Resources
	(*LoggingConfig)(nil),         // 6: yandex.cloud.metastore.v1.LoggingConfig
	nil,                           // 7: yandex.cloud.metastore.v1.Cluster.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*MaintenanceWindow)(nil),     // 9: yandex.cloud.metastore.v1.MaintenanceWindow
	(*MaintenanceOperation)(nil),  // 10: yandex.cloud.metastore.v1.MaintenanceOperation
	(v1.LogLevel_Level)(0),        // 11: yandex.cloud.logging.v1.LogLevel.Level
}
var file_yandex_cloud_metastore_v1_cluster_proto_depIdxs = []int32{
	8,  // 0: yandex.cloud.metastore.v1.Cluster.created_at:type_name -> google.protobuf.Timestamp
	7,  // 1: yandex.cloud.metastore.v1.Cluster.labels:type_name -> yandex.cloud.metastore.v1.Cluster.LabelsEntry
	0,  // 2: yandex.cloud.metastore.v1.Cluster.health:type_name -> yandex.cloud.metastore.v1.Cluster.Health
	1,  // 3: yandex.cloud.metastore.v1.Cluster.status:type_name -> yandex.cloud.metastore.v1.Cluster.Status
	3,  // 4: yandex.cloud.metastore.v1.Cluster.cluster_config:type_name -> yandex.cloud.metastore.v1.ClusterConfig
	6,  // 5: yandex.cloud.metastore.v1.Cluster.logging:type_name -> yandex.cloud.metastore.v1.LoggingConfig
	4,  // 6: yandex.cloud.metastore.v1.Cluster.network:type_name -> yandex.cloud.metastore.v1.NetworkConfig
	9,  // 7: yandex.cloud.metastore.v1.Cluster.maintenance_window:type_name -> yandex.cloud.metastore.v1.MaintenanceWindow
	10, // 8: yandex.cloud.metastore.v1.Cluster.planned_operation:type_name -> yandex.cloud.metastore.v1.MaintenanceOperation
	5,  // 9: yandex.cloud.metastore.v1.ClusterConfig.resources:type_name -> yandex.cloud.metastore.v1.Resources
	11, // 10: yandex.cloud.metastore.v1.LoggingConfig.min_level:type_name -> yandex.cloud.logging.v1.LogLevel.Level
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_yandex_cloud_metastore_v1_cluster_proto_init() }
func file_yandex_cloud_metastore_v1_cluster_proto_init() {
	if File_yandex_cloud_metastore_v1_cluster_proto != nil {
		return
	}
	file_yandex_cloud_metastore_v1_maintenance_proto_init()
	file_yandex_cloud_metastore_v1_cluster_proto_msgTypes[4].OneofWrappers = []any{
		(*LoggingConfig_FolderId)(nil),
		(*LoggingConfig_LogGroupId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_metastore_v1_cluster_proto_rawDesc), len(file_yandex_cloud_metastore_v1_cluster_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_metastore_v1_cluster_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_metastore_v1_cluster_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_metastore_v1_cluster_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_metastore_v1_cluster_proto_msgTypes,
	}.Build()
	File_yandex_cloud_metastore_v1_cluster_proto = out.File
	file_yandex_cloud_metastore_v1_cluster_proto_goTypes = nil
	file_yandex_cloud_metastore_v1_cluster_proto_depIdxs = nil
}
