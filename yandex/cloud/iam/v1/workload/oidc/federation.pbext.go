// Code generated by protoc-gen-goext. DO NOT EDIT.

package oidc

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Federation) SetId(v string) {
	m.Id = v
}

func (m *Federation) SetName(v string) {
	m.Name = v
}

func (m *Federation) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Federation) SetDescription(v string) {
	m.Description = v
}

func (m *Federation) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *Federation) SetAudiences(v []string) {
	m.Audiences = v
}

func (m *Federation) SetIssuer(v string) {
	m.Issuer = v
}

func (m *Federation) SetJwksUrl(v string) {
	m.JwksUrl = v
}

func (m *Federation) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Federation) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}
