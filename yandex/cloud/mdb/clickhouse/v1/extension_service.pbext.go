// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

func (m *GetExtensionRequest) SetExtensionName(v string) {
	m.ExtensionName = v
}

func (m *GetExtensionRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListExtensionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListExtensionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListExtensionsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListExtensionsResponse) SetExtensions(v []*Extension) {
	m.Extensions = v
}

func (m *ListExtensionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
