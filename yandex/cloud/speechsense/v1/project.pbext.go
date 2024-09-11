// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Project) SetId(v string) {
	m.Id = v
}

func (m *Project) SetName(v string) {
	m.Name = v
}

func (m *Project) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *Project) SetDescription(v string) {
	m.Description = v
}

func (m *Project) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *Project) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Project) SetModifiedBy(v string) {
	m.ModifiedBy = v
}

func (m *Project) SetModifiedAt(v *timestamppb.Timestamp) {
	m.ModifiedAt = v
}

func (m *Project) SetFilters(v []*FieldFilter) {
	m.Filters = v
}

func (m *FieldFilter) SetKey(v string) {
	m.Key = v
}

func (m *FieldFilter) SetFieldValue(v string) {
	m.FieldValue = v
}
