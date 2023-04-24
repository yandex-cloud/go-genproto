// Code generated by protoc-gen-goext. DO NOT EDIT.

package organizationmanager

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetGroupMappingRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *GetGroupMappingResponse) SetGroupMapping(v *GroupMapping) {
	m.GroupMapping = v
}

func (m *CreateGroupMappingRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *CreateGroupMappingRequest) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *CreateGroupMappingMetadata) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateGroupMappingRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateGroupMappingRequest) SetUpdatedFields(v *fieldmaskpb.FieldMask) {
	m.UpdatedFields = v
}

func (m *UpdateGroupMappingRequest) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *UpdateGroupMappingMetadata) SetFederationId(v string) {
	m.FederationId = v
}

func (m *DeleteGroupMappingRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *DeleteGroupMappingMetadata) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateGroupMappingItemsRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateGroupMappingItemsRequest) SetGroupMappingItemDeltas(v []*GroupMappingItemDelta) {
	m.GroupMappingItemDeltas = v
}

func (m *GroupMappingItemDelta) SetItem(v *GroupMappingItem) {
	m.Item = v
}

func (m *GroupMappingItemDelta) SetAction(v GroupMappingItemDelta_Action) {
	m.Action = v
}

func (m *UpdateGroupMappingItemsMetadata) SetFederationId(v string) {
	m.FederationId = v
}

func (m *UpdateGroupMappingItemsResponse) SetGroupMappingItemDeltas(v []*GroupMappingItemDelta) {
	m.GroupMappingItemDeltas = v
}

func (m *ListGroupMappingItemsRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *ListGroupMappingItemsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListGroupMappingItemsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListGroupMappingItemsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListGroupMappingItemsResponse) SetGroupMappingItems(v []*GroupMappingItem) {
	m.GroupMappingItems = v
}

func (m *ListGroupMappingItemsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}