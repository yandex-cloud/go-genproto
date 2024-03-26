// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/mongodb/v1/config"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *GetClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClustersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListClustersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClustersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClustersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClustersResponse) SetClusters(v []*Cluster) {
	m.Clusters = v
}

func (m *ListClustersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *CreateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *CreateClusterRequest) SetDatabaseSpecs(v []*DatabaseSpec) {
	m.DatabaseSpecs = v
}

func (m *CreateClusterRequest) SetUserSpecs(v []*UserSpec) {
	m.UserSpecs = v
}

func (m *CreateClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *UpdateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *UpdateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterRequest) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *MoveClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *MoveClusterMetadata) SetSourceFolderId(v string) {
	m.SourceFolderId = v
}

func (m *MoveClusterMetadata) SetDestinationFolderId(v string) {
	m.DestinationFolderId = v
}

func (m *BackupClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *BackupClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterRequest) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RestoreClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *RestoreClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *RestoreClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *RestoreClusterRequest) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *RestoreClusterRequest) SetConfigSpec(v *ConfigSpec) {
	m.ConfigSpec = v
}

func (m *RestoreClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *RestoreClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *RestoreClusterRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *RestoreClusterRequest) SetRecoveryTargetSpec(v *RestoreClusterRequest_RecoveryTargetSpec) {
	m.RecoveryTargetSpec = v
}

func (m *RestoreClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *RestoreClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *RestoreClusterRequest_RecoveryTargetSpec) SetTimestamp(v int64) {
	m.Timestamp = v
}

func (m *RestoreClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterMetadata) SetBackupId(v string) {
	m.BackupId = v
}

func (m *RescheduleMaintenanceRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceRequest) SetRescheduleType(v RescheduleMaintenanceRequest_RescheduleType) {
	m.RescheduleType = v
}

func (m *RescheduleMaintenanceRequest) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}

func (m *RescheduleMaintenanceMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RescheduleMaintenanceMetadata) SetDelayedUntil(v *timestamppb.Timestamp) {
	m.DelayedUntil = v
}

func (m *LogRecord) SetTimestamp(v *timestamppb.Timestamp) {
	m.Timestamp = v
}

func (m *LogRecord) SetMessage(v map[string]string) {
	m.Message = v
}

func (m *ListClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *ListClusterLogsRequest) SetServiceType(v ListClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *ListClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *ListClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *ListClusterLogsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterLogsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterLogsResponse) SetLogs(v []*LogRecord) {
	m.Logs = v
}

func (m *ListClusterLogsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *StreamLogRecord) SetRecord(v *LogRecord) {
	m.Record = v
}

func (m *StreamLogRecord) SetNextRecordToken(v string) {
	m.NextRecordToken = v
}

func (m *StreamClusterLogsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StreamClusterLogsRequest) SetColumnFilter(v []string) {
	m.ColumnFilter = v
}

func (m *StreamClusterLogsRequest) SetServiceType(v StreamClusterLogsRequest_ServiceType) {
	m.ServiceType = v
}

func (m *StreamClusterLogsRequest) SetFromTime(v *timestamppb.Timestamp) {
	m.FromTime = v
}

func (m *StreamClusterLogsRequest) SetToTime(v *timestamppb.Timestamp) {
	m.ToTime = v
}

func (m *StreamClusterLogsRequest) SetRecordToken(v string) {
	m.RecordToken = v
}

func (m *StreamClusterLogsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterOperationsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListClusterOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterBackupsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterBackupsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterBackupsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterBackupsResponse) SetBackups(v []*Backup) {
	m.Backups = v
}

func (m *ListClusterBackupsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *AddClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterHostsRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *AddClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *DeleteClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterHostsRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *DeleteClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *UpdateClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterHostsRequest) SetUpdateHostSpecs(v []*UpdateHostSpec) {
	m.UpdateHostSpecs = v
}

func (m *UpdateClusterHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *UpdateHostSpec) SetHostName(v string) {
	m.HostName = v
}

func (m *UpdateHostSpec) SetHidden(v *wrapperspb.BoolValue) {
	m.Hidden = v
}

func (m *UpdateHostSpec) SetSecondaryDelaySecs(v *wrapperspb.Int64Value) {
	m.SecondaryDelaySecs = v
}

func (m *UpdateHostSpec) SetPriority(v *wrapperspb.DoubleValue) {
	m.Priority = v
}

func (m *UpdateHostSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *UpdateHostSpec) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateHostSpec) SetTags(v map[string]string) {
	m.Tags = v
}

func (m *EnableClusterShardingRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *EnableClusterShardingRequest) SetMongocfg(v *EnableClusterShardingRequest_MongoCfg) {
	m.Mongocfg = v
}

func (m *EnableClusterShardingRequest) SetMongos(v *EnableClusterShardingRequest_Mongos) {
	m.Mongos = v
}

func (m *EnableClusterShardingRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *EnableClusterShardingRequest) SetMongoinfra(v *EnableClusterShardingRequest_MongoInfra) {
	m.Mongoinfra = v
}

func (m *EnableClusterShardingRequest_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *EnableClusterShardingRequest_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *EnableClusterShardingRequest_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *EnableClusterShardingMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *ListClusterShardsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterShardsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterShardsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterShardsResponse) SetShards(v []*Shard) {
	m.Shards = v
}

func (m *ListClusterShardsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *AddClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *AddClusterShardRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *AddClusterShardMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *AddClusterShardMetadata) SetShardName(v string) {
	m.ShardName = v
}

func (m *DeleteClusterShardRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterShardRequest) SetShardName(v string) {
	m.ShardName = v
}

func (m *DeleteClusterShardMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterShardMetadata) SetShardName(v string) {
	m.ShardName = v
}

func (m *ResetupHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ResetupHostsRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *ResetupHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ResetupHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *RestartHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestartHostsRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *RestartHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestartHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *StepdownHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StepdownHostsRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *StepdownHostsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StepdownHostsMetadata) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *HostSpec) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *HostSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *HostSpec) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *HostSpec) SetType(v Host_Type) {
	m.Type = v
}

func (m *HostSpec) SetShardName(v string) {
	m.ShardName = v
}

func (m *HostSpec) SetHidden(v *wrapperspb.BoolValue) {
	m.Hidden = v
}

func (m *HostSpec) SetSecondaryDelaySecs(v *wrapperspb.Int64Value) {
	m.SecondaryDelaySecs = v
}

func (m *HostSpec) SetPriority(v *wrapperspb.DoubleValue) {
	m.Priority = v
}

func (m *HostSpec) SetTags(v map[string]string) {
	m.Tags = v
}

func (m *MongodbSpec3_6) SetMongod(v *MongodbSpec3_6_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec3_6) SetMongocfg(v *MongodbSpec3_6_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec3_6) SetMongos(v *MongodbSpec3_6_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec3_6) SetMongoinfra(v *MongodbSpec3_6_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec3_6_Mongod) SetConfig(v *config.MongodConfig3_6) {
	m.Config = v
}

func (m *MongodbSpec3_6_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec3_6_MongoCfg) SetConfig(v *config.MongoCfgConfig3_6) {
	m.Config = v
}

func (m *MongodbSpec3_6_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec3_6_Mongos) SetConfig(v *config.MongosConfig3_6) {
	m.Config = v
}

func (m *MongodbSpec3_6_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec3_6_MongoInfra) SetConfigMongos(v *config.MongosConfig3_6) {
	m.ConfigMongos = v
}

func (m *MongodbSpec3_6_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig3_6) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec3_6_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_0) SetMongod(v *MongodbSpec4_0_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec4_0) SetMongocfg(v *MongodbSpec4_0_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec4_0) SetMongos(v *MongodbSpec4_0_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec4_0) SetMongoinfra(v *MongodbSpec4_0_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec4_0_Mongod) SetConfig(v *config.MongodConfig4_0) {
	m.Config = v
}

func (m *MongodbSpec4_0_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_0_MongoCfg) SetConfig(v *config.MongoCfgConfig4_0) {
	m.Config = v
}

func (m *MongodbSpec4_0_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_0_Mongos) SetConfig(v *config.MongosConfig4_0) {
	m.Config = v
}

func (m *MongodbSpec4_0_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_0_MongoInfra) SetConfigMongos(v *config.MongosConfig4_0) {
	m.ConfigMongos = v
}

func (m *MongodbSpec4_0_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig4_0) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec4_0_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_2) SetMongod(v *MongodbSpec4_2_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec4_2) SetMongocfg(v *MongodbSpec4_2_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec4_2) SetMongos(v *MongodbSpec4_2_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec4_2) SetMongoinfra(v *MongodbSpec4_2_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec4_2_Mongod) SetConfig(v *config.MongodConfig4_2) {
	m.Config = v
}

func (m *MongodbSpec4_2_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_2_MongoCfg) SetConfig(v *config.MongoCfgConfig4_2) {
	m.Config = v
}

func (m *MongodbSpec4_2_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_2_Mongos) SetConfig(v *config.MongosConfig4_2) {
	m.Config = v
}

func (m *MongodbSpec4_2_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_2_MongoInfra) SetConfigMongos(v *config.MongosConfig4_2) {
	m.ConfigMongos = v
}

func (m *MongodbSpec4_2_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig4_2) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec4_2_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4) SetMongod(v *MongodbSpec4_4_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec4_4) SetMongocfg(v *MongodbSpec4_4_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec4_4) SetMongos(v *MongodbSpec4_4_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec4_4) SetMongoinfra(v *MongodbSpec4_4_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec4_4_Mongod) SetConfig(v *config.MongodConfig4_4) {
	m.Config = v
}

func (m *MongodbSpec4_4_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4_MongoCfg) SetConfig(v *config.MongoCfgConfig4_4) {
	m.Config = v
}

func (m *MongodbSpec4_4_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4_Mongos) SetConfig(v *config.MongosConfig4_4) {
	m.Config = v
}

func (m *MongodbSpec4_4_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4_MongoInfra) SetConfigMongos(v *config.MongosConfig4_4) {
	m.ConfigMongos = v
}

func (m *MongodbSpec4_4_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig4_4) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec4_4_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4Enterprise) SetMongod(v *MongodbSpec4_4Enterprise_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec4_4Enterprise) SetMongocfg(v *MongodbSpec4_4Enterprise_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec4_4Enterprise) SetMongos(v *MongodbSpec4_4Enterprise_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec4_4Enterprise) SetMongoinfra(v *MongodbSpec4_4Enterprise_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec4_4Enterprise_Mongod) SetConfig(v *config.MongodConfig4_4Enterprise) {
	m.Config = v
}

func (m *MongodbSpec4_4Enterprise_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4Enterprise_MongoCfg) SetConfig(v *config.MongoCfgConfig4_4Enterprise) {
	m.Config = v
}

func (m *MongodbSpec4_4Enterprise_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4Enterprise_Mongos) SetConfig(v *config.MongosConfig4_4Enterprise) {
	m.Config = v
}

func (m *MongodbSpec4_4Enterprise_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec4_4Enterprise_MongoInfra) SetConfigMongos(v *config.MongosConfig4_4Enterprise) {
	m.ConfigMongos = v
}

func (m *MongodbSpec4_4Enterprise_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig4_4Enterprise) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec4_4Enterprise_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0) SetMongod(v *MongodbSpec5_0_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec5_0) SetMongocfg(v *MongodbSpec5_0_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec5_0) SetMongos(v *MongodbSpec5_0_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec5_0) SetMongoinfra(v *MongodbSpec5_0_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec5_0_Mongod) SetConfig(v *config.MongodConfig5_0) {
	m.Config = v
}

func (m *MongodbSpec5_0_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0_MongoCfg) SetConfig(v *config.MongoCfgConfig5_0) {
	m.Config = v
}

func (m *MongodbSpec5_0_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0_Mongos) SetConfig(v *config.MongosConfig5_0) {
	m.Config = v
}

func (m *MongodbSpec5_0_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0_MongoInfra) SetConfigMongos(v *config.MongosConfig5_0) {
	m.ConfigMongos = v
}

func (m *MongodbSpec5_0_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig5_0) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec5_0_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0Enterprise) SetMongod(v *MongodbSpec5_0Enterprise_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec5_0Enterprise) SetMongocfg(v *MongodbSpec5_0Enterprise_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec5_0Enterprise) SetMongos(v *MongodbSpec5_0Enterprise_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec5_0Enterprise) SetMongoinfra(v *MongodbSpec5_0Enterprise_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec5_0Enterprise_Mongod) SetConfig(v *config.MongodConfig5_0Enterprise) {
	m.Config = v
}

func (m *MongodbSpec5_0Enterprise_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0Enterprise_MongoCfg) SetConfig(v *config.MongoCfgConfig5_0Enterprise) {
	m.Config = v
}

func (m *MongodbSpec5_0Enterprise_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0Enterprise_Mongos) SetConfig(v *config.MongosConfig5_0Enterprise) {
	m.Config = v
}

func (m *MongodbSpec5_0Enterprise_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec5_0Enterprise_MongoInfra) SetConfigMongos(v *config.MongosConfig5_0Enterprise) {
	m.ConfigMongos = v
}

func (m *MongodbSpec5_0Enterprise_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig5_0Enterprise) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec5_0Enterprise_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0) SetMongod(v *MongodbSpec6_0_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec6_0) SetMongocfg(v *MongodbSpec6_0_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec6_0) SetMongos(v *MongodbSpec6_0_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec6_0) SetMongoinfra(v *MongodbSpec6_0_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec6_0_Mongod) SetConfig(v *config.MongodConfig6_0) {
	m.Config = v
}

func (m *MongodbSpec6_0_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0_MongoCfg) SetConfig(v *config.MongoCfgConfig6_0) {
	m.Config = v
}

func (m *MongodbSpec6_0_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0_Mongos) SetConfig(v *config.MongosConfig6_0) {
	m.Config = v
}

func (m *MongodbSpec6_0_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0_MongoInfra) SetConfigMongos(v *config.MongosConfig6_0) {
	m.ConfigMongos = v
}

func (m *MongodbSpec6_0_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig6_0) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec6_0_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0Enterprise) SetMongod(v *MongodbSpec6_0Enterprise_Mongod) {
	m.Mongod = v
}

func (m *MongodbSpec6_0Enterprise) SetMongocfg(v *MongodbSpec6_0Enterprise_MongoCfg) {
	m.Mongocfg = v
}

func (m *MongodbSpec6_0Enterprise) SetMongos(v *MongodbSpec6_0Enterprise_Mongos) {
	m.Mongos = v
}

func (m *MongodbSpec6_0Enterprise) SetMongoinfra(v *MongodbSpec6_0Enterprise_MongoInfra) {
	m.Mongoinfra = v
}

func (m *MongodbSpec6_0Enterprise_Mongod) SetConfig(v *config.MongodConfig6_0Enterprise) {
	m.Config = v
}

func (m *MongodbSpec6_0Enterprise_Mongod) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0Enterprise_MongoCfg) SetConfig(v *config.MongoCfgConfig6_0Enterprise) {
	m.Config = v
}

func (m *MongodbSpec6_0Enterprise_MongoCfg) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0Enterprise_Mongos) SetConfig(v *config.MongosConfig6_0Enterprise) {
	m.Config = v
}

func (m *MongodbSpec6_0Enterprise_Mongos) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MongodbSpec6_0Enterprise_MongoInfra) SetConfigMongos(v *config.MongosConfig6_0Enterprise) {
	m.ConfigMongos = v
}

func (m *MongodbSpec6_0Enterprise_MongoInfra) SetConfigMongocfg(v *config.MongoCfgConfig6_0Enterprise) {
	m.ConfigMongocfg = v
}

func (m *MongodbSpec6_0Enterprise_MongoInfra) SetResources(v *Resources) {
	m.Resources = v
}

type ConfigSpec_MongodbSpec = isConfigSpec_MongodbSpec

func (m *ConfigSpec) SetMongodbSpec(v ConfigSpec_MongodbSpec) {
	m.MongodbSpec = v
}

func (m *ConfigSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigSpec) SetFeatureCompatibilityVersion(v string) {
	m.FeatureCompatibilityVersion = v
}

func (m *ConfigSpec) SetMongodbSpec_3_6(v *MongodbSpec3_6) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_3_6{
		MongodbSpec_3_6: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_4_0(v *MongodbSpec4_0) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_4_0{
		MongodbSpec_4_0: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_4_2(v *MongodbSpec4_2) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_4_2{
		MongodbSpec_4_2: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_4_4(v *MongodbSpec4_4) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_4_4{
		MongodbSpec_4_4: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_5_0(v *MongodbSpec5_0) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_5_0{
		MongodbSpec_5_0: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_6_0(v *MongodbSpec6_0) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_6_0{
		MongodbSpec_6_0: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_4_4Enterprise(v *MongodbSpec4_4Enterprise) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_4_4Enterprise{
		MongodbSpec_4_4Enterprise: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_5_0Enterprise(v *MongodbSpec5_0Enterprise) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_5_0Enterprise{
		MongodbSpec_5_0Enterprise: v,
	}
}

func (m *ConfigSpec) SetMongodbSpec_6_0Enterprise(v *MongodbSpec6_0Enterprise) {
	m.MongodbSpec = &ConfigSpec_MongodbSpec_6_0Enterprise{
		MongodbSpec_6_0Enterprise: v,
	}
}

func (m *ConfigSpec) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ConfigSpec) SetBackupRetainPeriodDays(v *wrapperspb.Int64Value) {
	m.BackupRetainPeriodDays = v
}

func (m *ConfigSpec) SetPerformanceDiagnostics(v *PerformanceDiagnosticsConfig) {
	m.PerformanceDiagnostics = v
}

func (m *ConfigSpec) SetAccess(v *Access) {
	m.Access = v
}
