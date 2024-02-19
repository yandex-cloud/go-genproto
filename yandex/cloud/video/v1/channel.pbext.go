// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Channel) SetId(v string) {
	m.Id = v
}

func (m *Channel) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *Channel) SetTitle(v string) {
	m.Title = v
}

func (m *Channel) SetDescription(v string) {
	m.Description = v
}

func (m *Channel) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Channel) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Channel) SetLabels(v map[string]string) {
	m.Labels = v
}
