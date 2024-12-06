// Code generated by protoc-gen-goext. DO NOT EDIT.

package searchapi

func (m *SortSpec) SetSortMode(v SortSpec_SortMode) {
	m.SortMode = v
}

func (m *SortSpec) SetSortOrder(v SortSpec_SortOrder) {
	m.SortOrder = v
}

func (m *GroupSpec) SetGroupMode(v GroupSpec_GroupMode) {
	m.GroupMode = v
}

func (m *GroupSpec) SetGroupsOnPage(v int64) {
	m.GroupsOnPage = v
}

func (m *GroupSpec) SetDocsInGroup(v int64) {
	m.DocsInGroup = v
}

func (m *WebSearchRequest) SetQuery(v *SearchQuery) {
	m.Query = v
}

func (m *WebSearchRequest) SetSortSpec(v *SortSpec) {
	m.SortSpec = v
}

func (m *WebSearchRequest) SetGroupSpec(v *GroupSpec) {
	m.GroupSpec = v
}

func (m *WebSearchRequest) SetMaxPassages(v int64) {
	m.MaxPassages = v
}

func (m *WebSearchRequest) SetRegion(v string) {
	m.Region = v
}

func (m *WebSearchRequest) SetL10N(v WebSearchRequest_Localization) {
	m.L10N = v
}

func (m *WebSearchRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *WebSearchRequest) SetResponseFormat(v WebSearchRequest_Format) {
	m.ResponseFormat = v
}

func (m *WebSearchRequest) SetUserAgent(v string) {
	m.UserAgent = v
}

func (m *WebSearchResponse) SetRawData(v []byte) {
	m.RawData = v
}