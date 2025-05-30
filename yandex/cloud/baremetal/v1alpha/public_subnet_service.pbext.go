// Code generated by protoc-gen-goext. DO NOT EDIT.

package baremetal

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetPublicSubnetRequest) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *ListPublicSubnetRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListPublicSubnetRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPublicSubnetRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPublicSubnetRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListPublicSubnetRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListPublicSubnetResponse) SetPublicSubnets(v []*PublicSubnet) {
	m.PublicSubnets = v
}

func (m *ListPublicSubnetResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreatePublicSubnetRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreatePublicSubnetRequest) SetName(v string) {
	m.Name = v
}

func (m *CreatePublicSubnetRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreatePublicSubnetRequest) SetHardwarePoolIds(v []string) {
	m.HardwarePoolIds = v
}

func (m *CreatePublicSubnetRequest) SetPrefixLength(v int64) {
	m.PrefixLength = v
}

func (m *CreatePublicSubnetRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreatePublicSubnetMetadata) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *UpdatePublicSubnetRequest) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *UpdatePublicSubnetRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdatePublicSubnetRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdatePublicSubnetRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdatePublicSubnetRequest) SetHardwarePoolIds(v []string) {
	m.HardwarePoolIds = v
}

func (m *UpdatePublicSubnetRequest) SetType(v PublicSubnetType) {
	m.Type = v
}

func (m *UpdatePublicSubnetRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdatePublicSubnetMetadata) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *DeletePublicSubnetRequest) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *DeletePublicSubnetMetadata) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *ListPublicSubnetOperationsRequest) SetPublicSubnetId(v string) {
	m.PublicSubnetId = v
}

func (m *ListPublicSubnetOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPublicSubnetOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPublicSubnetOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListPublicSubnetOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
