// Code generated by protoc-gen-goext. DO NOT EDIT.

package cic

func (m *GetTrunkConnectionRequest) SetTrunkConnectionId(v string) {
	m.TrunkConnectionId = v
}

func (m *ListTrunkConnectionsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListTrunkConnectionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListTrunkConnectionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListTrunkConnectionsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListTrunkConnectionsResponse) SetTrunkConnections(v []*TrunkConnection) {
	m.TrunkConnections = v
}

func (m *ListTrunkConnectionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
