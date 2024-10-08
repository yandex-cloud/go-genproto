// Code generated by protoc-gen-goext. DO NOT EDIT.

package audittrails

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetTrailRequest) SetTrailId(v string) {
	m.TrailId = v
}

func (m *ListTrailsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListTrailsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListTrailsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListTrailsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListTrailsRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListTrailsResponse) SetTrails(v []*Trail) {
	m.Trails = v
}

func (m *ListTrailsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateTrailRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateTrailRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateTrailRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateTrailRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateTrailRequest) SetDestination(v *Trail_Destination) {
	m.Destination = v
}

func (m *CreateTrailRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateTrailRequest) SetFilter(v *Trail_Filter) {
	m.Filter = v
}

func (m *CreateTrailRequest) SetFilteringPolicy(v *Trail_FilteringPolicy) {
	m.FilteringPolicy = v
}

func (m *UpdateTrailRequest) SetTrailId(v string) {
	m.TrailId = v
}

func (m *UpdateTrailRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateTrailRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateTrailRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateTrailRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateTrailRequest) SetDestination(v *Trail_Destination) {
	m.Destination = v
}

func (m *UpdateTrailRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *UpdateTrailRequest) SetFilter(v *Trail_Filter) {
	m.Filter = v
}

func (m *UpdateTrailRequest) SetFilteringPolicy(v *Trail_FilteringPolicy) {
	m.FilteringPolicy = v
}

func (m *DeleteTrailRequest) SetTrailId(v string) {
	m.TrailId = v
}

func (m *CreateTrailMetadata) SetTrailId(v string) {
	m.TrailId = v
}

func (m *UpdateTrailMetadata) SetTrailId(v string) {
	m.TrailId = v
}

func (m *DeleteTrailMetadata) SetTrailId(v string) {
	m.TrailId = v
}

func (m *ListTrailOperationsRequest) SetTrailId(v string) {
	m.TrailId = v
}

func (m *ListTrailOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListTrailOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListTrailOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListTrailOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
