// Code generated by protoc-gen-goext. DO NOT EDIT.

package regression

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Dashboard) SetId(v string) {
	m.Id = v
}

func (m *Dashboard) SetName(v string) {
	m.Name = v
}

func (m *Dashboard) SetDescription(v string) {
	m.Description = v
}

func (m *Dashboard) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Dashboard) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Dashboard) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *Dashboard) SetUpdatedBy(v string) {
	m.UpdatedBy = v
}

func (m *Dashboard) SetEtag(v string) {
	m.Etag = v
}

func (m *Dashboard) SetContent(v *Dashboard_Content) {
	m.Content = v
}

func (m *Dashboard_Content) SetWidgets(v []*Widget) {
	m.Widgets = v
}
