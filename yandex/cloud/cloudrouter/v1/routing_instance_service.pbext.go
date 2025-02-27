// Code generated by protoc-gen-goext. DO NOT EDIT.

package cloudrouter

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetRoutingInstanceRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *GetRoutingInstanceByCicPrivateConnectionIdRequest) SetCicPrivateConnectionId(v string) {
	m.CicPrivateConnectionId = v
}

func (m *GetRoutingInstanceByVpcNetworkIdRequest) SetVpcNetworkId(v string) {
	m.VpcNetworkId = v
}

func (m *ListRoutingInstancesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListRoutingInstancesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRoutingInstancesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRoutingInstancesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListRoutingInstancesResponse) SetRoutingInstances(v []*RoutingInstance) {
	m.RoutingInstances = v
}

func (m *ListRoutingInstancesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateRoutingInstanceRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateRoutingInstanceRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateRoutingInstanceRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateRoutingInstanceRequest) SetRegionId(v string) {
	m.RegionId = v
}

func (m *CreateRoutingInstanceRequest) SetVpcInfo(v []*RoutingInstance_VpcInfo) {
	m.VpcInfo = v
}

func (m *CreateRoutingInstanceRequest) SetCicPrivateConnectionInfo(v []*RoutingInstance_CicPrivateConnectionInfo) {
	m.CicPrivateConnectionInfo = v
}

func (m *CreateRoutingInstanceRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateRoutingInstanceMetadata) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *UpdateRoutingInstanceRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *UpdateRoutingInstanceRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateRoutingInstanceRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateRoutingInstanceRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateRoutingInstanceRequest) SetRegionId(v string) {
	m.RegionId = v
}

func (m *UpdateRoutingInstanceRequest) SetVpcInfo(v []*RoutingInstance_VpcInfo) {
	m.VpcInfo = v
}

func (m *UpdateRoutingInstanceRequest) SetCicPrivateConnectionInfo(v []*RoutingInstance_CicPrivateConnectionInfo) {
	m.CicPrivateConnectionInfo = v
}

func (m *UpdateRoutingInstanceRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateRoutingInstanceMetadata) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *UpsertPrefixesRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *UpsertPrefixesRequest) SetVpcNetworkId(v string) {
	m.VpcNetworkId = v
}

func (m *UpsertPrefixesRequest) SetVpcAzInfoPrefixes(v []*VpcAzInfoPrefixes) {
	m.VpcAzInfoPrefixes = v
}

func (m *RemovePrefixesRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *RemovePrefixesRequest) SetVpcNetworkId(v string) {
	m.VpcNetworkId = v
}

func (m *RemovePrefixesRequest) SetVpcAzInfoPrefixes(v []*VpcAzInfoPrefixes) {
	m.VpcAzInfoPrefixes = v
}

func (m *VpcAzInfoPrefixes) SetAzId(v string) {
	m.AzId = v
}

func (m *VpcAzInfoPrefixes) SetPrefixes(v []string) {
	m.Prefixes = v
}

func (m *AddPrivateConnectionRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *AddPrivateConnectionRequest) SetCicPrivateConnectionId(v string) {
	m.CicPrivateConnectionId = v
}

func (m *RemovePrivateConnectionRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *RemovePrivateConnectionRequest) SetCicPrivateConnectionId(v string) {
	m.CicPrivateConnectionId = v
}

func (m *DeleteRoutingInstanceRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *DeleteRoutingInstanceMetadata) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *ListRoutingInstanceOperationsRequest) SetRoutingInstanceId(v string) {
	m.RoutingInstanceId = v
}

func (m *ListRoutingInstanceOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListRoutingInstanceOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListRoutingInstanceOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListRoutingInstanceOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
