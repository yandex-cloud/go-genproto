// Code generated by protoc-gen-goext. DO NOT EDIT.

package cloudregistry

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetRegistryRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListRegistriesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListRegistriesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRegistriesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRegistriesResponse) SetRegistries(v []*Registry) {
	m.Registries = v
}

func (m *ListRegistriesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateRegistryRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateRegistryRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateRegistryRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateRegistryRequest) SetKind(v Registry_Kind) {
	m.Kind = v
}

func (m *CreateRegistryRequest) SetType(v Registry_Type) {
	m.Type = v
}

func (m *CreateRegistryRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateRegistryRequest) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *CreateRegistryMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *UpdateRegistryRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *UpdateRegistryRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateRegistryRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateRegistryRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateRegistryRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateRegistryRequest) SetProperties(v map[string]string) {
	m.Properties = v
}

func (m *UpdateRegistryMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *DeleteRegistryMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *SetIpPermissionsRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *SetIpPermissionsRequest) SetIpPermissions(v []*IpPermission) {
	m.IpPermissions = v
}

func (m *UpdateIpPermissionsRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *UpdateIpPermissionsRequest) SetIpPermissionDeltas(v []*IpPermissionDelta) {
	m.IpPermissionDeltas = v
}

func (m *ListIpPermissionsRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListIpPermissionsResponse) SetPermissions(v []*IpPermission) {
	m.Permissions = v
}

func (m *SetIpPermissionsMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *UpdateIpPermissionsMetadata) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListArtifactsRequest) SetRegistryId(v string) {
	m.RegistryId = v
}

func (m *ListArtifactsRequest) SetPath(v string) {
	m.Path = v
}

func (m *ListArtifactsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListArtifactsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListArtifactsResponse) SetArtifacts(v []*Artifact) {
	m.Artifacts = v
}

func (m *ListArtifactsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
