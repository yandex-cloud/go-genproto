// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	config "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1/config"
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Cluster) SetId(v string) {
	m.Id = v
}

func (m *Cluster) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Cluster) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Cluster) SetName(v string) {
	m.Name = v
}

func (m *Cluster) SetDescription(v string) {
	m.Description = v
}

func (m *Cluster) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Cluster) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *Cluster) SetMonitoring(v []*Monitoring) {
	m.Monitoring = v
}

func (m *Cluster) SetConfig(v *ClusterConfig) {
	m.Config = v
}

func (m *Cluster) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Cluster) SetHealth(v Cluster_Health) {
	m.Health = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
}

func (m *Cluster) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Cluster) SetHostGroupIds(v []string) {
	m.HostGroupIds = v
}

func (m *Monitoring) SetName(v string) {
	m.Name = v
}

func (m *Monitoring) SetDescription(v string) {
	m.Description = v
}

func (m *Monitoring) SetLink(v string) {
	m.Link = v
}

type ClusterConfig_PostgresqlConfig = isClusterConfig_PostgresqlConfig

func (m *ClusterConfig) SetPostgresqlConfig(v ClusterConfig_PostgresqlConfig) {
	m.PostgresqlConfig = v
}

func (m *ClusterConfig) SetVersion(v string) {
	m.Version = v
}

func (m *ClusterConfig) SetPostgresqlConfig_9_6(v *config.PostgresqlConfigSet9_6) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_9_6{
		PostgresqlConfig_9_6: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_10_1C(v *config.PostgresqlConfigSet10_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_10_1C{
		PostgresqlConfig_10_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_10(v *config.PostgresqlConfigSet10) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_10{
		PostgresqlConfig_10: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_11(v *config.PostgresqlConfigSet11) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_11{
		PostgresqlConfig_11: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_11_1C(v *config.PostgresqlConfigSet11_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_11_1C{
		PostgresqlConfig_11_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_12(v *config.PostgresqlConfigSet12) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_12{
		PostgresqlConfig_12: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_12_1C(v *config.PostgresqlConfigSet12_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_12_1C{
		PostgresqlConfig_12_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_13(v *config.PostgresqlConfigSet13) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_13{
		PostgresqlConfig_13: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_13_1C(v *config.PostgresqlConfigSet13_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_13_1C{
		PostgresqlConfig_13_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_14(v *config.PostgresqlConfigSet14) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_14{
		PostgresqlConfig_14: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_14_1C(v *config.PostgresqlConfigSet14_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_14_1C{
		PostgresqlConfig_14_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_15(v *config.PostgresqlConfigSet15) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_15{
		PostgresqlConfig_15: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_15_1C(v *config.PostgresqlConfigSet15_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_15_1C{
		PostgresqlConfig_15_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_16(v *config.PostgresqlConfigSet16) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_16{
		PostgresqlConfig_16: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_16_1C(v *config.PostgresqlConfigSet16_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_16_1C{
		PostgresqlConfig_16_1C: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_17(v *config.PostgresqlConfigSet17) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_17{
		PostgresqlConfig_17: v,
	}
}

func (m *ClusterConfig) SetPostgresqlConfig_17_1C(v *config.PostgresqlConfigSet17_1C) {
	m.PostgresqlConfig = &ClusterConfig_PostgresqlConfig_17_1C{
		PostgresqlConfig_17_1C: v,
	}
}

func (m *ClusterConfig) SetPoolerConfig(v *ConnectionPoolerConfig) {
	m.PoolerConfig = v
}

func (m *ClusterConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *ClusterConfig) SetAutofailover(v *wrapperspb.BoolValue) {
	m.Autofailover = v
}

func (m *ClusterConfig) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ClusterConfig) SetBackupRetainPeriodDays(v *wrapperspb.Int64Value) {
	m.BackupRetainPeriodDays = v
}

func (m *ClusterConfig) SetAccess(v *Access) {
	m.Access = v
}

func (m *ClusterConfig) SetPerformanceDiagnostics(v *PerformanceDiagnostics) {
	m.PerformanceDiagnostics = v
}

func (m *ClusterConfig) SetDiskSizeAutoscaling(v *DiskSizeAutoscaling) {
	m.DiskSizeAutoscaling = v
}

func (m *ConnectionPoolerConfig) SetPoolingMode(v ConnectionPoolerConfig_PoolingMode) {
	m.PoolingMode = v
}

func (m *ConnectionPoolerConfig) SetPoolDiscard(v *wrapperspb.BoolValue) {
	m.PoolDiscard = v
}

func (m *Host) SetName(v string) {
	m.Name = v
}

func (m *Host) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Host) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *Host) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Host) SetRole(v Host_Role) {
	m.Role = v
}

func (m *Host) SetHealth(v Host_Health) {
	m.Health = v
}

func (m *Host) SetServices(v []*Service) {
	m.Services = v
}

func (m *Host) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *Host) SetReplicationSource(v string) {
	m.ReplicationSource = v
}

func (m *Host) SetPriority(v *wrapperspb.Int64Value) {
	m.Priority = v
}

func (m *Host) SetConfig(v *HostConfig) {
	m.Config = v
}

func (m *Host) SetAssignPublicIp(v bool) {
	m.AssignPublicIp = v
}

func (m *Host) SetReplicaType(v Host_ReplicaType) {
	m.ReplicaType = v
}

type HostConfig_PostgresqlConfig = isHostConfig_PostgresqlConfig

func (m *HostConfig) SetPostgresqlConfig(v HostConfig_PostgresqlConfig) {
	m.PostgresqlConfig = v
}

func (m *HostConfig) SetPostgresqlConfig_9_6(v *config.PostgresqlHostConfig9_6) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_9_6{
		PostgresqlConfig_9_6: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_10_1C(v *config.PostgresqlHostConfig10_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_10_1C{
		PostgresqlConfig_10_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_10(v *config.PostgresqlHostConfig10) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_10{
		PostgresqlConfig_10: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_11(v *config.PostgresqlHostConfig11) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_11{
		PostgresqlConfig_11: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_11_1C(v *config.PostgresqlHostConfig11_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_11_1C{
		PostgresqlConfig_11_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_12(v *config.PostgresqlHostConfig12) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_12{
		PostgresqlConfig_12: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_12_1C(v *config.PostgresqlHostConfig12_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_12_1C{
		PostgresqlConfig_12_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_13(v *config.PostgresqlHostConfig13) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_13{
		PostgresqlConfig_13: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_13_1C(v *config.PostgresqlHostConfig13_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_13_1C{
		PostgresqlConfig_13_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_14(v *config.PostgresqlHostConfig14) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_14{
		PostgresqlConfig_14: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_14_1C(v *config.PostgresqlHostConfig14_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_14_1C{
		PostgresqlConfig_14_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_15(v *config.PostgresqlHostConfig15) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_15{
		PostgresqlConfig_15: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_15_1C(v *config.PostgresqlHostConfig15_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_15_1C{
		PostgresqlConfig_15_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_16(v *config.PostgresqlHostConfig16) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_16{
		PostgresqlConfig_16: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_16_1C(v *config.PostgresqlHostConfig16_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_16_1C{
		PostgresqlConfig_16_1C: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_17(v *config.PostgresqlHostConfig17) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_17{
		PostgresqlConfig_17: v,
	}
}

func (m *HostConfig) SetPostgresqlConfig_17_1C(v *config.PostgresqlHostConfig17_1C) {
	m.PostgresqlConfig = &HostConfig_PostgresqlConfig_17_1C{
		PostgresqlConfig_17_1C: v,
	}
}

func (m *Service) SetType(v Service_Type) {
	m.Type = v
}

func (m *Service) SetHealth(v Service_Health) {
	m.Health = v
}

func (m *Resources) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *Resources) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *Resources) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *Access) SetDataLens(v bool) {
	m.DataLens = v
}

func (m *Access) SetWebSql(v bool) {
	m.WebSql = v
}

func (m *Access) SetServerless(v bool) {
	m.Serverless = v
}

func (m *Access) SetDataTransfer(v bool) {
	m.DataTransfer = v
}

func (m *Access) SetYandexQuery(v bool) {
	m.YandexQuery = v
}

func (m *PerformanceDiagnostics) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *PerformanceDiagnostics) SetSessionsSamplingInterval(v int64) {
	m.SessionsSamplingInterval = v
}

func (m *PerformanceDiagnostics) SetStatementsSamplingInterval(v int64) {
	m.StatementsSamplingInterval = v
}

func (m *DiskSizeAutoscaling) SetPlannedUsageThreshold(v int64) {
	m.PlannedUsageThreshold = v
}

func (m *DiskSizeAutoscaling) SetEmergencyUsageThreshold(v int64) {
	m.EmergencyUsageThreshold = v
}

func (m *DiskSizeAutoscaling) SetDiskSizeLimit(v int64) {
	m.DiskSizeLimit = v
}
