// Code generated by protoc-gen-goext. DO NOT EDIT.

package backup

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

func (m *ListResourcesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListResourcesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListResourcesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListResourcesRequest) SetType(v ResourceType) {
	m.Type = v
}

func (m *ListResourcesResponse) SetResources(v []*Resource) {
	m.Resources = v
}

func (m *ListResourcesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *GetResourceRequest) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *GetResourceRequest) SetIncludeTenantInfo(v bool) {
	m.IncludeTenantInfo = v
}

func (m *GetResourceResponse) SetResource(v *Resource) {
	m.Resource = v
}

func (m *DeleteResourceRequest) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *DeleteResourceRequest) SetResourceId(v string) {
	m.ResourceId = v
}

func (m *DeleteResourceMetadata) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *ListTasksRequest) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *ListTasksRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListTasksRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListTasksResponse) SetTasks(v []*Task) {
	m.Tasks = v
}

func (m *ListTasksResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListDirectoryRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListDirectoryRequest) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *ListDirectoryRequest) SetPath(v string) {
	m.Path = v
}

func (m *ListDirectoryResponse) SetItems(v []*ListDirectoryResponse_FilesystemItem) {
	m.Items = v
}

func (m *ListDirectoryResponse_FilesystemItem) SetName(v string) {
	m.Name = v
}

func (m *ListDirectoryResponse_FilesystemItem) SetType(v ListDirectoryResponse_FilesystemItem_Type) {
	m.Type = v
}

func (m *ListDirectoryResponse_FilesystemItem) SetFileType(v ListDirectoryResponse_FilesystemItem_Type) {
	m.FileType = v
}

func (m *ListDirectoryResponse_FilesystemItem) SetSize(v int64) {
	m.Size = v
}

func (m *CreateDirectoryRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateDirectoryRequest) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *CreateDirectoryRequest) SetPath(v string) {
	m.Path = v
}

func (m *CreateDirectoryMetadata) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *CreateDirectoryMetadata) SetPath(v string) {
	m.Path = v
}

func (m *ListResourceOperationsRequest) SetComputeInstanceId(v string) {
	m.ComputeInstanceId = v
}

func (m *ListResourceOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListResourceOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListResourceOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListResourceOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
