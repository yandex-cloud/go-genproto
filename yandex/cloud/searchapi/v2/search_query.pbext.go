// Code generated by protoc-gen-goext. DO NOT EDIT.

package searchapi

func (m *SearchQuery) SetSearchType(v SearchQuery_SearchType) {
	m.SearchType = v
}

func (m *SearchQuery) SetQueryText(v string) {
	m.QueryText = v
}

func (m *SearchQuery) SetFamilyMode(v SearchQuery_FamilyMode) {
	m.FamilyMode = v
}

func (m *SearchQuery) SetPage(v int64) {
	m.Page = v
}

func (m *SearchQuery) SetFixTypoMode(v SearchQuery_FixTypoMode) {
	m.FixTypoMode = v
}
