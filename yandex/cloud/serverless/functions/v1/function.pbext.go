// Code generated by protoc-gen-goext. DO NOT EDIT.

package functions

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Function) SetId(v string) {
	m.Id = v
}

func (m *Function) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Function) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Function) SetName(v string) {
	m.Name = v
}

func (m *Function) SetDescription(v string) {
	m.Description = v
}

func (m *Function) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Function) SetHttpInvokeUrl(v string) {
	m.HttpInvokeUrl = v
}

func (m *Function) SetStatus(v Function_Status) {
	m.Status = v
}

func (m *Version) SetId(v string) {
	m.Id = v
}

func (m *Version) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *Version) SetDescription(v string) {
	m.Description = v
}

func (m *Version) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Version) SetRuntime(v string) {
	m.Runtime = v
}

func (m *Version) SetEntrypoint(v string) {
	m.Entrypoint = v
}

func (m *Version) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Version) SetExecutionTimeout(v *durationpb.Duration) {
	m.ExecutionTimeout = v
}

func (m *Version) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Version) SetImageSize(v int64) {
	m.ImageSize = v
}

func (m *Version) SetStatus(v Version_Status) {
	m.Status = v
}

func (m *Version) SetTags(v []string) {
	m.Tags = v
}

func (m *Version) SetEnvironment(v map[string]string) {
	m.Environment = v
}

func (m *Version) SetConnectivity(v *Connectivity) {
	m.Connectivity = v
}

func (m *Version) SetNamedServiceAccounts(v map[string]string) {
	m.NamedServiceAccounts = v
}

func (m *Version) SetSecrets(v []*Secret) {
	m.Secrets = v
}

func (m *Version) SetLogOptions(v *LogOptions) {
	m.LogOptions = v
}

func (m *Version) SetStorageMounts(v []*StorageMount) {
	m.StorageMounts = v
}

func (m *Version) SetAsyncInvocationConfig(v *AsyncInvocationConfig) {
	m.AsyncInvocationConfig = v
}

func (m *Version) SetTmpfsSize(v int64) {
	m.TmpfsSize = v
}

func (m *Version) SetConcurrency(v int64) {
	m.Concurrency = v
}

func (m *Resources) SetMemory(v int64) {
	m.Memory = v
}

func (m *Package) SetBucketName(v string) {
	m.BucketName = v
}

func (m *Package) SetObjectName(v string) {
	m.ObjectName = v
}

func (m *Package) SetSha256(v string) {
	m.Sha256 = v
}

func (m *Connectivity) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Connectivity) SetSubnetId(v []string) {
	m.SubnetId = v
}

func (m *ScalingPolicy) SetFunctionId(v string) {
	m.FunctionId = v
}

func (m *ScalingPolicy) SetTag(v string) {
	m.Tag = v
}

func (m *ScalingPolicy) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *ScalingPolicy) SetModifiedAt(v *timestamppb.Timestamp) {
	m.ModifiedAt = v
}

func (m *ScalingPolicy) SetProvisionedInstancesCount(v int64) {
	m.ProvisionedInstancesCount = v
}

func (m *ScalingPolicy) SetZoneInstancesLimit(v int64) {
	m.ZoneInstancesLimit = v
}

func (m *ScalingPolicy) SetZoneRequestsLimit(v int64) {
	m.ZoneRequestsLimit = v
}

type Secret_Reference = isSecret_Reference

func (m *Secret) SetReference(v Secret_Reference) {
	m.Reference = v
}

func (m *Secret) SetId(v string) {
	m.Id = v
}

func (m *Secret) SetVersionId(v string) {
	m.VersionId = v
}

func (m *Secret) SetKey(v string) {
	m.Key = v
}

func (m *Secret) SetEnvironmentVariable(v string) {
	m.Reference = &Secret_EnvironmentVariable{
		EnvironmentVariable: v,
	}
}

type LogOptions_Destination = isLogOptions_Destination

func (m *LogOptions) SetDestination(v LogOptions_Destination) {
	m.Destination = v
}

func (m *LogOptions) SetDisabled(v bool) {
	m.Disabled = v
}

func (m *LogOptions) SetLogGroupId(v string) {
	m.Destination = &LogOptions_LogGroupId{
		LogGroupId: v,
	}
}

func (m *LogOptions) SetFolderId(v string) {
	m.Destination = &LogOptions_FolderId{
		FolderId: v,
	}
}

func (m *LogOptions) SetMinLevel(v v1.LogLevel_Level) {
	m.MinLevel = v
}

func (m *StorageMount) SetBucketId(v string) {
	m.BucketId = v
}

func (m *StorageMount) SetPrefix(v string) {
	m.Prefix = v
}

func (m *StorageMount) SetMountPointName(v string) {
	m.MountPointName = v
}

func (m *StorageMount) SetReadOnly(v bool) {
	m.ReadOnly = v
}

func (m *AsyncInvocationConfig) SetRetriesCount(v int64) {
	m.RetriesCount = v
}

func (m *AsyncInvocationConfig) SetSuccessTarget(v *AsyncInvocationConfig_ResponseTarget) {
	m.SuccessTarget = v
}

func (m *AsyncInvocationConfig) SetFailureTarget(v *AsyncInvocationConfig_ResponseTarget) {
	m.FailureTarget = v
}

func (m *AsyncInvocationConfig) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

type AsyncInvocationConfig_ResponseTarget_Target = isAsyncInvocationConfig_ResponseTarget_Target

func (m *AsyncInvocationConfig_ResponseTarget) SetTarget(v AsyncInvocationConfig_ResponseTarget_Target) {
	m.Target = v
}

func (m *AsyncInvocationConfig_ResponseTarget) SetEmptyTarget(v *EmptyTarget) {
	m.Target = &AsyncInvocationConfig_ResponseTarget_EmptyTarget{
		EmptyTarget: v,
	}
}

func (m *AsyncInvocationConfig_ResponseTarget) SetYmqTarget(v *YMQTarget) {
	m.Target = &AsyncInvocationConfig_ResponseTarget_YmqTarget{
		YmqTarget: v,
	}
}

func (m *YMQTarget) SetQueueArn(v string) {
	m.QueueArn = v
}

func (m *YMQTarget) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}
