// Code generated by protoc-gen-goext. DO NOT EDIT.

package cic

func (m *GetPublicConnectionRequest) SetPublicConnectionId(v string) {
	m.PublicConnectionId = v
}

func (m *ListPublicConnectionsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListPublicConnectionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPublicConnectionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPublicConnectionsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListPublicConnectionsResponse) SetPublicConnections(v []*PublicConnection) {
	m.PublicConnections = v
}

func (m *ListPublicConnectionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}