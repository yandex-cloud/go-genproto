// Code generated by protoc-gen-goext. DO NOT EDIT.

package dataproc

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
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

func (m *CreateSubclusterConfigSpec) SetName(v string) {
	m.Name = v
}

func (m *CreateSubclusterConfigSpec) SetRole(v Role) {
	m.Role = v
}

func (m *CreateSubclusterConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *CreateSubclusterConfigSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *CreateSubclusterConfigSpec) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *CreateSubclusterConfigSpec) SetAutoscalingConfig(v *AutoscalingConfig) {
	m.AutoscalingConfig = v
}

func (m *UpdateSubclusterConfigSpec) SetId(v string) {
	m.Id = v
}

func (m *UpdateSubclusterConfigSpec) SetName(v string) {
	m.Name = v
}

func (m *UpdateSubclusterConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *UpdateSubclusterConfigSpec) SetHostsCount(v int64) {
	m.HostsCount = v
}

func (m *UpdateSubclusterConfigSpec) SetAutoscalingConfig(v *AutoscalingConfig) {
	m.AutoscalingConfig = v
}

func (m *CreateClusterConfigSpec) SetVersionId(v string) {
	m.VersionId = v
}

func (m *CreateClusterConfigSpec) SetHadoop(v *HadoopConfig) {
	m.Hadoop = v
}

func (m *CreateClusterConfigSpec) SetSubclustersSpec(v []*CreateSubclusterConfigSpec) {
	m.SubclustersSpec = v
}

func (m *UpdateClusterConfigSpec) SetSubclustersSpec(v []*UpdateSubclusterConfigSpec) {
	m.SubclustersSpec = v
}

func (m *UpdateClusterConfigSpec) SetHadoop(v *HadoopConfig) {
	m.Hadoop = v
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

func (m *CreateClusterRequest) SetConfigSpec(v *CreateClusterConfigSpec) {
	m.ConfigSpec = v
}

func (m *CreateClusterRequest) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *CreateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateClusterRequest) SetBucket(v string) {
	m.Bucket = v
}

func (m *CreateClusterRequest) SetUiProxy(v bool) {
	m.UiProxy = v
}

func (m *CreateClusterRequest) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *CreateClusterRequest) SetHostGroupIds(v []string) {
	m.HostGroupIds = v
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

func (m *UpdateClusterRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateClusterRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateClusterRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateClusterRequest) SetConfigSpec(v *UpdateClusterConfigSpec) {
	m.ConfigSpec = v
}

func (m *UpdateClusterRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateClusterRequest) SetBucket(v string) {
	m.Bucket = v
}

func (m *UpdateClusterRequest) SetDecommissionTimeout(v int64) {
	m.DecommissionTimeout = v
}

func (m *UpdateClusterRequest) SetUiProxy(v bool) {
	m.UiProxy = v
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

func (m *DeleteClusterRequest) SetDecommissionTimeout(v int64) {
	m.DecommissionTimeout = v
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

func (m *StopClusterRequest) SetDecommissionTimeout(v int64) {
	m.DecommissionTimeout = v
}

func (m *StopClusterMetadata) SetClusterId(v string) {
	m.ClusterId = v
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

func (m *ListClusterHostsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterHostsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterHostsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterHostsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListUILinksRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UILink) SetName(v string) {
	m.Name = v
}

func (m *UILink) SetUrl(v string) {
	m.Url = v
}

func (m *ListUILinksResponse) SetLinks(v []*UILink) {
	m.Links = v
}
