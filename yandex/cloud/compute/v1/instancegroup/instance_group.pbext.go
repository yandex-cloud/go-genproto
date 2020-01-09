// Code generated by protoc-gen-goext. DO NOT EDIT.

package instancegroup

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *InstanceGroup) SetId(v string) {
	m.Id = v
}

func (m *InstanceGroup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *InstanceGroup) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *InstanceGroup) SetName(v string) {
	m.Name = v
}

func (m *InstanceGroup) SetDescription(v string) {
	m.Description = v
}

func (m *InstanceGroup) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *InstanceGroup) SetInstanceTemplate(v *InstanceTemplate) {
	m.InstanceTemplate = v
}

func (m *InstanceGroup) SetScalePolicy(v *ScalePolicy) {
	m.ScalePolicy = v
}

func (m *InstanceGroup) SetDeployPolicy(v *DeployPolicy) {
	m.DeployPolicy = v
}

func (m *InstanceGroup) SetAllocationPolicy(v *AllocationPolicy) {
	m.AllocationPolicy = v
}

func (m *InstanceGroup) SetLoadBalancerState(v *LoadBalancerState) {
	m.LoadBalancerState = v
}

func (m *InstanceGroup) SetManagedInstancesState(v *ManagedInstancesState) {
	m.ManagedInstancesState = v
}

func (m *InstanceGroup) SetLoadBalancerSpec(v *LoadBalancerSpec) {
	m.LoadBalancerSpec = v
}

func (m *InstanceGroup) SetHealthChecksSpec(v *HealthChecksSpec) {
	m.HealthChecksSpec = v
}

func (m *InstanceGroup) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *InstanceGroup) SetStatus(v InstanceGroup_Status) {
	m.Status = v
}

func (m *LoadBalancerState) SetTargetGroupId(v string) {
	m.TargetGroupId = v
}

func (m *LoadBalancerState) SetStatusMessage(v string) {
	m.StatusMessage = v
}

func (m *ManagedInstancesState) SetTargetSize(v int64) {
	m.TargetSize = v
}

func (m *ManagedInstancesState) SetRunningActualCount(v int64) {
	m.RunningActualCount = v
}

func (m *ManagedInstancesState) SetRunningOutdatedCount(v int64) {
	m.RunningOutdatedCount = v
}

func (m *ManagedInstancesState) SetProcessingCount(v int64) {
	m.ProcessingCount = v
}

func (m *ManagedInstancesState_Statuses) SetCreating(v int64) {
	m.Creating = v
}

func (m *ManagedInstancesState_Statuses) SetStarting(v int64) {
	m.Starting = v
}

func (m *ManagedInstancesState_Statuses) SetOpening(v int64) {
	m.Opening = v
}

func (m *ManagedInstancesState_Statuses) SetWarming(v int64) {
	m.Warming = v
}

func (m *ManagedInstancesState_Statuses) SetRunning(v int64) {
	m.Running = v
}

func (m *ManagedInstancesState_Statuses) SetClosing(v int64) {
	m.Closing = v
}

func (m *ManagedInstancesState_Statuses) SetStopping(v int64) {
	m.Stopping = v
}

func (m *ManagedInstancesState_Statuses) SetUpdating(v int64) {
	m.Updating = v
}

func (m *ManagedInstancesState_Statuses) SetDeleting(v int64) {
	m.Deleting = v
}

func (m *ManagedInstancesState_Statuses) SetFailed(v int64) {
	m.Failed = v
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

func (m *ScalePolicy_AutoScale) SetMinZoneSize(v int64) {
	m.MinZoneSize = v
}

func (m *ScalePolicy_AutoScale) SetMaxSize(v int64) {
	m.MaxSize = v
}

func (m *ScalePolicy_AutoScale) SetMeasurementDuration(v *duration.Duration) {
	m.MeasurementDuration = v
}

func (m *ScalePolicy_AutoScale) SetWarmupDuration(v *duration.Duration) {
	m.WarmupDuration = v
}

func (m *ScalePolicy_AutoScale) SetStabilizationDuration(v *duration.Duration) {
	m.StabilizationDuration = v
}

func (m *ScalePolicy_AutoScale) SetInitialSize(v int64) {
	m.InitialSize = v
}

func (m *ScalePolicy_AutoScale) SetCpuUtilizationRule(v *ScalePolicy_CpuUtilizationRule) {
	m.CpuUtilizationRule = v
}

func (m *ScalePolicy_AutoScale) SetCustomRules(v []*ScalePolicy_CustomRule) {
	m.CustomRules = v
}

func (m *ScalePolicy_CpuUtilizationRule) SetUtilizationTarget(v float64) {
	m.UtilizationTarget = v
}

func (m *ScalePolicy_CustomRule) SetRuleType(v ScalePolicy_CustomRule_RuleType) {
	m.RuleType = v
}

func (m *ScalePolicy_CustomRule) SetMetricType(v ScalePolicy_CustomRule_MetricType) {
	m.MetricType = v
}

func (m *ScalePolicy_CustomRule) SetMetricName(v string) {
	m.MetricName = v
}

func (m *ScalePolicy_CustomRule) SetTarget(v float64) {
	m.Target = v
}

func (m *ScalePolicy_FixedScale) SetSize(v int64) {
	m.Size = v
}

func (m *DeployPolicy) SetMaxUnavailable(v int64) {
	m.MaxUnavailable = v
}

func (m *DeployPolicy) SetMaxDeleting(v int64) {
	m.MaxDeleting = v
}

func (m *DeployPolicy) SetMaxCreating(v int64) {
	m.MaxCreating = v
}

func (m *DeployPolicy) SetMaxExpansion(v int64) {
	m.MaxExpansion = v
}

func (m *DeployPolicy) SetStartupDuration(v *duration.Duration) {
	m.StartupDuration = v
}

func (m *AllocationPolicy) SetZones(v []*AllocationPolicy_Zone) {
	m.Zones = v
}

func (m *AllocationPolicy_Zone) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *InstanceTemplate) SetDescription(v string) {
	m.Description = v
}

func (m *InstanceTemplate) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *InstanceTemplate) SetPlatformId(v string) {
	m.PlatformId = v
}

func (m *InstanceTemplate) SetResourcesSpec(v *ResourcesSpec) {
	m.ResourcesSpec = v
}

func (m *InstanceTemplate) SetMetadata(v map[string]string) {
	m.Metadata = v
}

func (m *InstanceTemplate) SetBootDiskSpec(v *AttachedDiskSpec) {
	m.BootDiskSpec = v
}

func (m *InstanceTemplate) SetSecondaryDiskSpecs(v []*AttachedDiskSpec) {
	m.SecondaryDiskSpecs = v
}

func (m *InstanceTemplate) SetNetworkInterfaceSpecs(v []*NetworkInterfaceSpec) {
	m.NetworkInterfaceSpecs = v
}

func (m *InstanceTemplate) SetSchedulingPolicy(v *SchedulingPolicy) {
	m.SchedulingPolicy = v
}

func (m *InstanceTemplate) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *ResourcesSpec) SetMemory(v int64) {
	m.Memory = v
}

func (m *ResourcesSpec) SetCores(v int64) {
	m.Cores = v
}

func (m *ResourcesSpec) SetCoreFraction(v int64) {
	m.CoreFraction = v
}

func (m *ResourcesSpec) SetGpus(v int64) {
	m.Gpus = v
}

func (m *AttachedDiskSpec) SetMode(v AttachedDiskSpec_Mode) {
	m.Mode = v
}

func (m *AttachedDiskSpec) SetDeviceName(v string) {
	m.DeviceName = v
}

func (m *AttachedDiskSpec) SetDiskSpec(v *AttachedDiskSpec_DiskSpec) {
	m.DiskSpec = v
}

type AttachedDiskSpec_DiskSpec_SourceOneof = isAttachedDiskSpec_DiskSpec_SourceOneof

func (m *AttachedDiskSpec_DiskSpec) SetSourceOneof(v AttachedDiskSpec_DiskSpec_SourceOneof) {
	m.SourceOneof = v
}

func (m *AttachedDiskSpec_DiskSpec) SetDescription(v string) {
	m.Description = v
}

func (m *AttachedDiskSpec_DiskSpec) SetTypeId(v string) {
	m.TypeId = v
}

func (m *AttachedDiskSpec_DiskSpec) SetSize(v int64) {
	m.Size = v
}

func (m *AttachedDiskSpec_DiskSpec) SetImageId(v string) {
	m.SourceOneof = &AttachedDiskSpec_DiskSpec_ImageId{
		ImageId: v,
	}
}

func (m *AttachedDiskSpec_DiskSpec) SetSnapshotId(v string) {
	m.SourceOneof = &AttachedDiskSpec_DiskSpec_SnapshotId{
		SnapshotId: v,
	}
}

func (m *NetworkInterfaceSpec) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *NetworkInterfaceSpec) SetSubnetIds(v []string) {
	m.SubnetIds = v
}

func (m *NetworkInterfaceSpec) SetPrimaryV4AddressSpec(v *PrimaryAddressSpec) {
	m.PrimaryV4AddressSpec = v
}

func (m *NetworkInterfaceSpec) SetPrimaryV6AddressSpec(v *PrimaryAddressSpec) {
	m.PrimaryV6AddressSpec = v
}

func (m *PrimaryAddressSpec) SetOneToOneNatSpec(v *OneToOneNatSpec) {
	m.OneToOneNatSpec = v
}

func (m *OneToOneNatSpec) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

func (m *SchedulingPolicy) SetPreemptible(v bool) {
	m.Preemptible = v
}

func (m *LoadBalancerSpec) SetTargetGroupSpec(v *TargetGroupSpec) {
	m.TargetGroupSpec = v
}

func (m *TargetGroupSpec) SetName(v string) {
	m.Name = v
}

func (m *TargetGroupSpec) SetDescription(v string) {
	m.Description = v
}

func (m *TargetGroupSpec) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *HealthChecksSpec) SetHealthCheckSpecs(v []*HealthCheckSpec) {
	m.HealthCheckSpecs = v
}

type HealthCheckSpec_HealthCheckOptions = isHealthCheckSpec_HealthCheckOptions

func (m *HealthCheckSpec) SetHealthCheckOptions(v HealthCheckSpec_HealthCheckOptions) {
	m.HealthCheckOptions = v
}

func (m *HealthCheckSpec) SetInterval(v *duration.Duration) {
	m.Interval = v
}

func (m *HealthCheckSpec) SetTimeout(v *duration.Duration) {
	m.Timeout = v
}

func (m *HealthCheckSpec) SetUnhealthyThreshold(v int64) {
	m.UnhealthyThreshold = v
}

func (m *HealthCheckSpec) SetHealthyThreshold(v int64) {
	m.HealthyThreshold = v
}

func (m *HealthCheckSpec) SetTcpOptions(v *HealthCheckSpec_TcpOptions) {
	m.HealthCheckOptions = &HealthCheckSpec_TcpOptions_{
		TcpOptions: v,
	}
}

func (m *HealthCheckSpec) SetHttpOptions(v *HealthCheckSpec_HttpOptions) {
	m.HealthCheckOptions = &HealthCheckSpec_HttpOptions_{
		HttpOptions: v,
	}
}

func (m *HealthCheckSpec_TcpOptions) SetPort(v int64) {
	m.Port = v
}

func (m *HealthCheckSpec_HttpOptions) SetPort(v int64) {
	m.Port = v
}

func (m *HealthCheckSpec_HttpOptions) SetPath(v string) {
	m.Path = v
}

func (m *ManagedInstance) SetId(v string) {
	m.Id = v
}

func (m *ManagedInstance) SetStatus(v ManagedInstance_Status) {
	m.Status = v
}

func (m *ManagedInstance) SetInstanceId(v string) {
	m.InstanceId = v
}

func (m *ManagedInstance) SetFqdn(v string) {
	m.Fqdn = v
}

func (m *ManagedInstance) SetName(v string) {
	m.Name = v
}

func (m *ManagedInstance) SetStatusMessage(v string) {
	m.StatusMessage = v
}

func (m *ManagedInstance) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *ManagedInstance) SetNetworkInterfaces(v []*NetworkInterface) {
	m.NetworkInterfaces = v
}

func (m *ManagedInstance) SetStatusChangedAt(v *timestamp.Timestamp) {
	m.StatusChangedAt = v
}

func (m *NetworkInterface) SetIndex(v string) {
	m.Index = v
}

func (m *NetworkInterface) SetMacAddress(v string) {
	m.MacAddress = v
}

func (m *NetworkInterface) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *NetworkInterface) SetPrimaryV4Address(v *PrimaryAddress) {
	m.PrimaryV4Address = v
}

func (m *NetworkInterface) SetPrimaryV6Address(v *PrimaryAddress) {
	m.PrimaryV6Address = v
}

func (m *PrimaryAddress) SetAddress(v string) {
	m.Address = v
}

func (m *PrimaryAddress) SetOneToOneNat(v *OneToOneNat) {
	m.OneToOneNat = v
}

func (m *OneToOneNat) SetAddress(v string) {
	m.Address = v
}

func (m *OneToOneNat) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

func (m *LogRecord) SetTimestamp(v *timestamp.Timestamp) {
	m.Timestamp = v
}

func (m *LogRecord) SetMessage(v string) {
	m.Message = v
}
