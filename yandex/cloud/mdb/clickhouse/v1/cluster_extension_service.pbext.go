// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

func (m *GetClusterExtensionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *GetClusterExtensionRequest) SetExtensionName(v string) {
	m.ExtensionName = v
}

func (m *ListClusterExtensionsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListClusterExtensionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListClusterExtensionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListClusterExtensionsResponse) SetExtensions(v []*ClusterExtension) {
	m.Extensions = v
}

func (m *ListClusterExtensionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateClusterExtensionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateClusterExtensionRequest) SetExtensionSpec(v *ExtensionSpec) {
	m.ExtensionSpec = v
}

func (m *CreateClusterExtensionMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *CreateClusterExtensionMetadata) SetExtensionName(v string) {
	m.ExtensionName = v
}

func (m *DeleteClusterExtensionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterExtensionRequest) SetExtensionName(v string) {
	m.ExtensionName = v
}

func (m *DeleteClusterExtensionMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *DeleteClusterExtensionMetadata) SetExtensionName(v string) {
	m.ExtensionName = v
}

func (m *UpdateClusterExtensionRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterExtensionRequest) SetExtensionSpec(v *ExtensionSpec) {
	m.ExtensionSpec = v
}

func (m *UpdateClusterExtensionMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateClusterExtensionMetadata) SetExtensionName(v string) {
	m.ExtensionName = v
}

func (m *SetClusterExtensionsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *SetClusterExtensionsRequest) SetExtensionSpecs(v []*ExtensionSpec) {
	m.ExtensionSpecs = v
}

func (m *SetClusterExtensionsMetadata) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *SetClusterExtensionsMetadata) SetAddedExtensionNames(v []string) {
	m.AddedExtensionNames = v
}

func (m *SetClusterExtensionsMetadata) SetUpdatedExtensionNames(v []string) {
	m.UpdatedExtensionNames = v
}

func (m *SetClusterExtensionsMetadata) SetDeletedExtensionNames(v []string) {
	m.DeletedExtensionNames = v
}

func (m *ExtensionSpec) SetName(v string) {
	m.Name = v
}

func (m *ExtensionSpec) SetVersion(v string) {
	m.Version = v
}
