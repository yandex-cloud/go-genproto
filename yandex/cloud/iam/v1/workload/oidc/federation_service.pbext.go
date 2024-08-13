// Code generated by protoc-gen-goext. DO NOT EDIT.

package oidc

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetFederationRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *ListFederationsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListFederationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFederationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFederationsResponse) SetFederations(v []*Federation) {
	m.Federations = v
}

func (m *ListFederationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateFederationRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateFederationRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateFederationRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateFederationRequest) SetDisabled(v bool) {
	m.Disabled = v
}

func (m *CreateFederationRequest) SetAudiences(v []string) {
	m.Audiences = v
}

func (m *CreateFederationRequest) SetIssuer(v string) {
	m.Issuer = v
}

func (m *CreateFederationRequest) SetJwksUrl(v string) {
	m.JwksUrl = v
}

func (m *CreateFederationRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateFederationMetadata) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateFederationRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateFederationRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateFederationRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateFederationRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateFederationRequest) SetDisabled(v bool) {
	m.Disabled = v
}

func (m *UpdateFederationRequest) SetAudiences(v []string) {
	m.Audiences = v
}

func (m *UpdateFederationRequest) SetJwksUrl(v string) {
	m.JwksUrl = v
}

func (m *UpdateFederationRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateFederationMetadata) SetFederationId(v string) {
	m.FederationId = v
}

func (m *DeleteFederationRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *DeleteFederationMetadata) SetFederationId(v string) {
	m.FederationId = v
}