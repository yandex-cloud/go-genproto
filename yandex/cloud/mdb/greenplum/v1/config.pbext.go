// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Resources) SetResourcePresetId(v string) {
	m.ResourcePresetId = v
}

func (m *Resources) SetDiskSize(v int64) {
	m.DiskSize = v
}

func (m *Resources) SetDiskTypeId(v string) {
	m.DiskTypeId = v
}

func (m *ConnectionPoolerConfig) SetMode(v ConnectionPoolerConfig_PoolMode) {
	m.Mode = v
}

func (m *ConnectionPoolerConfig) SetSize(v *wrapperspb.Int64Value) {
	m.Size = v
}

func (m *ConnectionPoolerConfig) SetClientIdleTimeout(v *wrapperspb.Int64Value) {
	m.ClientIdleTimeout = v
}

func (m *MasterSubclusterConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MasterSubclusterConfig) SetConfig(v *GreenplumMasterConfigSet) {
	m.Config = v
}

func (m *SegmentSubclusterConfig) SetResources(v *Resources) {
	m.Resources = v
}

func (m *SegmentSubclusterConfig) SetConfig(v *GreenplumSegmentConfigSet) {
	m.Config = v
}

func (m *GreenplumMasterConfig) SetLogLevel(v GreenplumMasterConfig_LogLevel) {
	m.LogLevel = v
}

func (m *GreenplumMasterConfig) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *GreenplumMasterConfig) SetTimezone(v *wrapperspb.StringValue) {
	m.Timezone = v
}

func (m *GreenplumMasterConfig) SetPool(v *ConnectionPoolerConfig) {
	m.Pool = v
}

func (m *GreenplumMasterConfig) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *GreenplumMasterConfig) SetRunawayDetectorActivationPercent(v *wrapperspb.Int64Value) {
	m.RunawayDetectorActivationPercent = v
}

func (m *GreenplumMasterConfig) SetTcpKeepalivesCount(v *wrapperspb.Int64Value) {
	m.TcpKeepalivesCount = v
}

func (m *GreenplumMasterConfig) SetTcpKeepalivesInterval(v *wrapperspb.Int64Value) {
	m.TcpKeepalivesInterval = v
}

func (m *GreenplumMasterConfig) SetReadableExternalTableTimeout(v *wrapperspb.Int64Value) {
	m.ReadableExternalTableTimeout = v
}

func (m *GreenplumMasterConfig) SetGpInterconnectSndQueueDepth(v *wrapperspb.Int64Value) {
	m.GpInterconnectSndQueueDepth = v
}

func (m *GreenplumMasterConfig) SetGpInterconnectQueueDepth(v *wrapperspb.Int64Value) {
	m.GpInterconnectQueueDepth = v
}

func (m *GreenplumMasterConfig) SetLogStatement(v GreenplumMasterConfig_LogStatement) {
	m.LogStatement = v
}

func (m *GreenplumMasterConfig) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *GreenplumMasterConfig) SetOptimizerAnalyzeRootPartition(v *wrapperspb.BoolValue) {
	m.OptimizerAnalyzeRootPartition = v
}

func (m *GreenplumMasterConfig) SetGpExternalMaxSegs(v *wrapperspb.Int64Value) {
	m.GpExternalMaxSegs = v
}

func (m *GreenplumMasterConfig) SetGpFtsProbeTimeout(v *wrapperspb.Int64Value) {
	m.GpFtsProbeTimeout = v
}

func (m *GreenplumMasterConfig) SetGpWorkfileCompression(v *wrapperspb.BoolValue) {
	m.GpWorkfileCompression = v
}

func (m *GreenplumMasterConfig) SetGpAutostatsModeInFunctions(v GreenplumMasterConfig_AutostatsModeInFunctions) {
	m.GpAutostatsModeInFunctions = v
}

func (m *GreenplumSegmentConfig) SetLogLevel(v GreenplumSegmentConfig_LogLevel) {
	m.LogLevel = v
}

func (m *GreenplumSegmentConfig) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *GreenplumSegmentConfig) SetMaxSlotWalKeepSize(v *wrapperspb.Int64Value) {
	m.MaxSlotWalKeepSize = v
}

func (m *GreenplumSegmentConfig) SetGpWorkfileLimitPerSegment(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitPerSegment = v
}

func (m *GreenplumSegmentConfig) SetGpWorkfileLimitPerQuery(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitPerQuery = v
}

func (m *GreenplumSegmentConfig) SetGpWorkfileLimitFilesPerQuery(v *wrapperspb.Int64Value) {
	m.GpWorkfileLimitFilesPerQuery = v
}

func (m *GreenplumSegmentConfig) SetGpResourceManager(v GreenplumSegmentConfig_GPResourceManager) {
	m.GpResourceManager = v
}

func (m *GreenplumSegmentConfig) SetGpResourceGroupCpuLimit(v *wrapperspb.FloatValue) {
	m.GpResourceGroupCpuLimit = v
}

func (m *GreenplumSegmentConfig) SetGpResourceGroupMemoryLimit(v *wrapperspb.FloatValue) {
	m.GpResourceGroupMemoryLimit = v
}

func (m *GreenplumMasterConfigSet) SetEffectiveConfig(v *GreenplumMasterConfig) {
	m.EffectiveConfig = v
}

func (m *GreenplumMasterConfigSet) SetUserConfig(v *GreenplumMasterConfig) {
	m.UserConfig = v
}

func (m *GreenplumMasterConfigSet) SetDefaultConfig(v *GreenplumMasterConfig) {
	m.DefaultConfig = v
}

func (m *GreenplumSegmentConfigSet) SetEffectiveConfig(v *GreenplumSegmentConfig) {
	m.EffectiveConfig = v
}

func (m *GreenplumSegmentConfigSet) SetUserConfig(v *GreenplumSegmentConfig) {
	m.UserConfig = v
}

func (m *GreenplumSegmentConfigSet) SetDefaultConfig(v *GreenplumSegmentConfig) {
	m.DefaultConfig = v
}
