// Code generated by protoc-gen-goext. DO NOT EDIT.

package airflow

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
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

func (m *Cluster) SetMonitoring(v []*Monitoring) {
	m.Monitoring = v
}

func (m *Cluster) SetConfig(v *ClusterConfig) {
	m.Config = v
}

func (m *Cluster) SetHealth(v Health) {
	m.Health = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetNetwork(v *NetworkConfig) {
	m.Network = v
}

func (m *Cluster) SetCodeSync(v *CodeSyncConfig) {
	m.CodeSync = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Cluster) SetWebserverUrl(v string) {
	m.WebserverUrl = v
}

func (m *Cluster) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Cluster) SetLogging(v *LoggingConfig) {
	m.Logging = v
}

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
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

func (m *ClusterConfig) SetVersionId(v string) {
	m.VersionId = v
}

func (m *ClusterConfig) SetAirflow(v *AirflowConfig) {
	m.Airflow = v
}

func (m *ClusterConfig) SetWebserver(v *WebserverConfig) {
	m.Webserver = v
}

func (m *ClusterConfig) SetScheduler(v *SchedulerConfig) {
	m.Scheduler = v
}

func (m *ClusterConfig) SetTriggerer(v *TriggererConfig) {
	m.Triggerer = v
}

func (m *ClusterConfig) SetWorker(v *WorkerConfig) {
	m.Worker = v
}

func (m *ClusterConfig) SetDependencies(v *Dependencies) {
	m.Dependencies = v
}

func (m *ClusterConfig) SetLockbox(v *LockboxConfig) {
	m.Lockbox = v
}

func (m *ClusterConfig) SetAirflowVersion(v string) {
	m.AirflowVersion = v
}

func (m *ClusterConfig) SetPythonVersion(v string) {
	m.PythonVersion = v
}

func (m *AirflowConfig) SetConfig(v map[string]string) {
	m.Config = v
}

func (m *WebserverConfig) SetCount(v int64) {
	m.Count = v
}

func (m *WebserverConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *SchedulerConfig) SetCount(v int64) {
	m.Count = v
}

func (m *SchedulerConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *TriggererConfig) SetCount(v int64) {
	m.Count = v
}

func (m *TriggererConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *WorkerConfig) SetMinCount(v int64) {
	m.MinCount = v
}

func (m *WorkerConfig) SetMaxCount(v int64) {
	m.MaxCount = v
}

func (m *WorkerConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Dependencies) SetPipPackages(v []string) {
	m.PipPackages = v
}

func (m *Dependencies) SetDebPackages(v []string) {
	m.DebPackages = v
}

func (m *NetworkConfig) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *NetworkConfig) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *S3Config) SetBucket(v string) {
	m.Bucket = v
}

func (m *GitSyncConfig) SetRepo(v string) {
	m.Repo = v
}

func (m *GitSyncConfig) SetBranch(v string) {
	m.Branch = v
}

func (m *GitSyncConfig) SetSubPath(v string) {
	m.SubPath = v
}

func (m *GitSyncConfig) SetSshKey(v string) {
	m.SshKey = v
}

type CodeSyncConfig_Source = isCodeSyncConfig_Source

func (m *CodeSyncConfig) SetSource(v CodeSyncConfig_Source) {
	m.Source = v
}

func (m *CodeSyncConfig) SetS3(v *S3Config) {
	m.Source = &CodeSyncConfig_S3{
		S3: v,
	}
}

func (m *CodeSyncConfig) SetGitSync(v *GitSyncConfig) {
	m.Source = &CodeSyncConfig_GitSync{
		GitSync: v,
	}
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

func (m *LoggingConfig) SetMinLevel(v v1.LogLevel_Level) {
	m.MinLevel = v
}

func (m *LockboxConfig) SetEnabled(v bool) {
	m.Enabled = v
}
