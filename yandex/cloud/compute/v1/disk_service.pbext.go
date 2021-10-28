// Code generated by protoc-gen-goext. DO NOT EDIT.

package compute

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetDiskRequest) SetDiskId(v string) {
	m.DiskId = v
}

func (m *ListDisksRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListDisksRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDisksRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDisksRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListDisksResponse) SetDisks(v []*Disk) {
	m.Disks = v
}

func (m *ListDisksResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type CreateDiskRequest_Source = isCreateDiskRequest_Source

func (m *CreateDiskRequest) SetSource(v CreateDiskRequest_Source) {
	m.Source = v
}

func (m *CreateDiskRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateDiskRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateDiskRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateDiskRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateDiskRequest) SetTypeId(v string) {
	m.TypeId = v
}

func (m *CreateDiskRequest) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *CreateDiskRequest) SetSize(v int64) {
	m.Size = v
}

func (m *CreateDiskRequest) SetImageId(v string) {
	m.Source = &CreateDiskRequest_ImageId{
		ImageId: v,
	}
}

func (m *CreateDiskRequest) SetSnapshotId(v string) {
	m.Source = &CreateDiskRequest_SnapshotId{
		SnapshotId: v,
	}
}

func (m *CreateDiskRequest) SetBlockSize(v int64) {
	m.BlockSize = v
}

func (m *CreateDiskRequest) SetDiskPlacementPolicy(v *DiskPlacementPolicy) {
	m.DiskPlacementPolicy = v
}

func (m *CreateDiskMetadata) SetDiskId(v string) {
	m.DiskId = v
}

func (m *UpdateDiskRequest) SetDiskId(v string) {
	m.DiskId = v
}

func (m *UpdateDiskRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateDiskRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateDiskRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateDiskRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateDiskRequest) SetSize(v int64) {
	m.Size = v
}

func (m *UpdateDiskRequest) SetDiskPlacementPolicy(v *DiskPlacementPolicy) {
	m.DiskPlacementPolicy = v
}

func (m *UpdateDiskMetadata) SetDiskId(v string) {
	m.DiskId = v
}

func (m *DeleteDiskRequest) SetDiskId(v string) {
	m.DiskId = v
}

func (m *DeleteDiskMetadata) SetDiskId(v string) {
	m.DiskId = v
}

func (m *ListDiskOperationsRequest) SetDiskId(v string) {
	m.DiskId = v
}

func (m *ListDiskOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDiskOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDiskOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListDiskOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
