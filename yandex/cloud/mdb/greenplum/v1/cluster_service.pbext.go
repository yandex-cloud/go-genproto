// Code generated by protoc-gen-goext. DO NOT EDIT.

package greenplum

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
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

func (m *CreateClusterRequest) SetConfig(v *GreenplumConfig) {
	m.Config = v
}

func (m *CreateClusterRequest) SetMasterConfig(v *MasterSubclusterConfigSpec) {
	m.MasterConfig = v
}

func (m *CreateClusterRequest) SetSegmentConfig(v *SegmentSubclusterConfigSpec) {
	m.SegmentConfig = v
}

func (m *CreateClusterRequest) SetMasterHostCount(v int64) {
	m.MasterHostCount = v
}

func (m *CreateClusterRequest) SetSegmentInHost(v int64) {
	m.SegmentInHost = v
}

func (m *CreateClusterRequest) SetSegmentHostCount(v int64) {
	m.SegmentHostCount = v
}

func (m *CreateClusterRequest) SetUserName(v string) {
	m.UserName = v
}

func (m *CreateClusterRequest) SetUserPassword(v string) {
	m.UserPassword = v
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

func (m *CreateClusterRequest) SetHostGroupIds(v []string) {
	m.HostGroupIds = v
}

func (m *CreateClusterMetadata) SetClusterId(v string) {
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

func (m *ListClusterHostsResponse) SetHosts(v []*Host) {
	m.Hosts = v
}

func (m *ListClusterHostsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *MasterSubclusterConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *MasterSubclusterConfigSpec) SetConfig(v *GreenplumMasterConfig) {
	m.Config = v
}

func (m *SegmentSubclusterConfigSpec) SetResources(v *Resources) {
	m.Resources = v
}

func (m *SegmentSubclusterConfigSpec) SetConfig(v *GreenplumSegmentConfig) {
	m.Config = v
}
