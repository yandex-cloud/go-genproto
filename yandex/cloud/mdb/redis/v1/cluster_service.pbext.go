// Code generated by protoc-gen-goext. DO NOT EDIT.

package redis

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1/config"
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

func (m *CreateClusterRequest) SetHostSpecs(v []*HostSpec) {
	m.HostSpecs = v
}

func (m *CreateClusterRequest) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *CreateClusterRequest) SetSharded(v bool) {
	m.Sharded = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterRequest) SetTlsEnabled(v *wrapperspb.BoolValue) {
	m.TlsEnabled = v
}

func (m *CreateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateClusterRequest) SetPersistenceMode(v Cluster_PersistenceMode) {
	m.PersistenceMode = v
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

func (m *UpdateClusterRequest) SetPersistenceMode(v Cluster_PersistenceMode) {
	m.PersistenceMode = v
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

func (m *RestoreClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *RestoreClusterRequest) SetTlsEnabled(v *wrapperspb.BoolValue) {
	m.TlsEnabled = v
}

func (m *RestoreClusterRequest) SetPersistenceMode(v Cluster_PersistenceMode) {
	m.PersistenceMode = v
}

func (m *RestoreClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RestoreClusterMetadata) SetBackupId(v string) {
	m.BackupId = v
}

func (m *StartClusterFailoverRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterFailoverRequest) SetHostNames(v []string) {
	m.HostNames = v
}

func (m *StartClusterFailoverMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *StartClusterFailoverMetadata) SetHostNames(v []string) {
	m.HostNames = v
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

func (m *RebalanceClusterRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *RebalanceClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *HostSpec) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *HostSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *HostSpec) SetShardName(v string) {
	m.ShardName = v
}

type ConfigSpec_RedisSpec = isConfigSpec_RedisSpec

func (m *ConfigSpec) SetRedisSpec(v ConfigSpec_RedisSpec) {
	m.RedisSpec = v
}

func (m *ConfigSpec) SetVersion(v string) {
	m.Version = v
}

func (m *ConfigSpec) SetRedisConfig_5_0(v *config.RedisConfig5_0) {
	m.RedisSpec = &ConfigSpec_RedisConfig_5_0{
		RedisConfig_5_0: v,
	}
}

func (m *ConfigSpec) SetRedisConfig_6_0(v *config.RedisConfig6_0) {
	m.RedisSpec = &ConfigSpec_RedisConfig_6_0{
		RedisConfig_6_0: v,
	}
}

func (m *ConfigSpec) SetRedisConfig_6_2(v *config.RedisConfig6_2) {
	m.RedisSpec = &ConfigSpec_RedisConfig_6_2{
		RedisConfig_6_2: v,
	}
}

func (m *ConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ConfigSpec) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ConfigSpec) SetAccess(v *Access) {
	m.Access = v
}
