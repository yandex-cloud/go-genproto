// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

func (m *CreateProjectRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateProjectRequest) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *CreateProjectRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateProjectRequest) SetFilters(v []*FieldFilter) {
	m.Filters = v
}

func (m *CreateProjectMetadata) SetId(v string) {
	m.Id = v
}
