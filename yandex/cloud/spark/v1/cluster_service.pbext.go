// Code generated by protoc-gen-goext. DO NOT EDIT.

package spark

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

func (m *CreateClusterRequest) SetConfig(v *ClusterConfig) {
	m.Config = v
}

func (m *CreateClusterRequest) SetNetwork(v *NetworkConfig) {
	m.Network = v
}

func (m *CreateClusterRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateClusterRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateClusterRequest) SetLogging(v *LoggingConfig) {
	m.Logging = v
}

func (m *CreateClusterRequest) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
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
