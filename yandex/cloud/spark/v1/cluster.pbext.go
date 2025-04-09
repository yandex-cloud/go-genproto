// Code generated by protoc-gen-goext. DO NOT EDIT.

package spark

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

func (m *Cluster) SetConfig(v *ClusterConfig) {
	m.Config = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetNetwork(v *NetworkConfig) {
	m.Network = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Cluster) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Cluster) SetLogging(v *LoggingConfig) {
	m.Logging = v
}

func (m *Cluster) SetHealth(v Health) {
	m.Health = v
}

func (m *Cluster) SetLinks(v []*UILink) {
	m.Links = v
}

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
}

func (m *ClusterConfig) SetResourcePools(v *ResourcePools) {
	m.ResourcePools = v
}

func (m *ClusterConfig) SetHistoryServer(v *HistoryServerConfig) {
	m.HistoryServer = v
}

func (m *ClusterConfig) SetDependencies(v *Dependencies) {
	m.Dependencies = v
}

func (m *ClusterConfig) SetMetastore(v *Metastore) {
	m.Metastore = v
}

func (m *HistoryServerConfig) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *NetworkConfig) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *NetworkConfig) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *ResourcePools) SetDriver(v *ResourcePool) {
	m.Driver = v
}

func (m *ResourcePools) SetExecutor(v *ResourcePool) {
	m.Executor = v
}

func (m *ResourcePool) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *ResourcePool) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

type ScalePolicy_ScaleType = isScalePolicy_ScaleType

func (m *ScalePolicy) SetScaleType(v ScalePolicy_ScaleType) {
	m.ScaleType = v
}

func (m *ScalePolicy) SetFixedScale(v *ScalePolicy_FixedScale) {
	m.ScaleType = &ScalePolicy_FixedScale_{
		FixedScale: v,
	}
}

func (m *ScalePolicy) SetAutoScale(v *ScalePolicy_AutoScale) {
	m.ScaleType = &ScalePolicy_AutoScale_{
		AutoScale: v,
	}
}

func (m *ScalePolicy_FixedScale) SetSize(v int64) {
	m.Size = v
}

func (m *ScalePolicy_AutoScale) SetMinSize(v int64) {
	m.MinSize = v
}

func (m *ScalePolicy_AutoScale) SetMaxSize(v int64) {
	m.MaxSize = v
}

func (m *Dependencies) SetPipPackages(v []string) {
	m.PipPackages = v
}

func (m *Dependencies) SetDebPackages(v []string) {
	m.DebPackages = v
}

func (m *Metastore) SetClusterId(v string) {
	m.ClusterId = v
}

type LoggingConfig_Destination = isLoggingConfig_Destination

func (m *LoggingConfig) SetDestination(v LoggingConfig_Destination) {
	m.Destination = v
}

func (m *LoggingConfig) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *LoggingConfig) SetFolderId(v string) {
	m.Destination = &LoggingConfig_FolderId{
		FolderId: v,
	}
}

func (m *LoggingConfig) SetLogGroupId(v string) {
	m.Destination = &LoggingConfig_LogGroupId{
		LogGroupId: v,
	}
}

func (m *UILink) SetName(v string) {
	m.Name = v
}

func (m *UILink) SetUrl(v string) {
	m.Url = v
}
