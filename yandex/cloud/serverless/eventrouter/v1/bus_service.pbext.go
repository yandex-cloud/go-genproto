// Code generated by protoc-gen-goext. DO NOT EDIT.

package eventrouter

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetBusRequest) SetBusId(v string) {
	m.BusId = v
}

func (m *ListBusesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListBusesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBusesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBusesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListBusesResponse) SetBuses(v []*Bus) {
	m.Buses = v
}

func (m *ListBusesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateBusRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateBusRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateBusRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateBusRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateBusRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateBusMetadata) SetBusId(v string) {
	m.BusId = v
}

func (m *CreateBusMetadata) SetFolderId(v string) {
	m.FolderId = v
}

func (m *UpdateBusRequest) SetBusId(v string) {
	m.BusId = v
}

func (m *UpdateBusRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateBusRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateBusRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateBusRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateBusRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateBusMetadata) SetBusId(v string) {
	m.BusId = v
}

func (m *DeleteBusRequest) SetBusId(v string) {
	m.BusId = v
}

func (m *DeleteBusMetadata) SetBusId(v string) {
	m.BusId = v
}

func (m *ListBusOperationsRequest) SetBusId(v string) {
	m.BusId = v
}

func (m *ListBusOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListBusOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListBusOperationsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListBusOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListBusOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}