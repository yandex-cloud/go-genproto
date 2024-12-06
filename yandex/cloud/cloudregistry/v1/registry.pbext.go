// Code generated by protoc-gen-goext. DO NOT EDIT.

package cloudregistry

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Registry) SetId(v string) {
	m.Id = v
}

func (m *Registry) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Registry) SetName(v string) {
	m.Name = v
}

func (m *Registry) SetKind(v Registry_Kind) {
	m.Kind = v
}

func (m *Registry) SetType(v Registry_Type) {
	m.Type = v
}

func (m *Registry) SetStatus(v Registry_Status) {
	m.Status = v
}

func (m *Registry) SetDescription(v string) {
	m.Description = v
}

func (m *Registry) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Registry) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *Registry) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Registry) SetModifiedAt(v *timestamppb.Timestamp) {
	m.ModifiedAt = v
}