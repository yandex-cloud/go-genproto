// Code generated by protoc-gen-goext. DO NOT EDIT.

package loadbalancer

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetNetworkLoadBalancerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *ListNetworkLoadBalancersRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListNetworkLoadBalancersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListNetworkLoadBalancersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListNetworkLoadBalancersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListNetworkLoadBalancersResponse) SetNetworkLoadBalancers(v []*NetworkLoadBalancer) {
	m.NetworkLoadBalancers = v
}

func (m *ListNetworkLoadBalancersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateNetworkLoadBalancerRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateNetworkLoadBalancerRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateNetworkLoadBalancerRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateNetworkLoadBalancerRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateNetworkLoadBalancerRequest) SetRegionId(v string) {
	m.RegionId = v
}

func (m *CreateNetworkLoadBalancerRequest) SetType(v NetworkLoadBalancer_Type) {
	m.Type = v
}

func (m *CreateNetworkLoadBalancerRequest) SetListenerSpecs(v []*ListenerSpec) {
	m.ListenerSpecs = v
}

func (m *CreateNetworkLoadBalancerRequest) SetAttachedTargetGroups(v []*AttachedTargetGroup) {
	m.AttachedTargetGroups = v
}

func (m *CreateNetworkLoadBalancerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetListenerSpecs(v []*ListenerSpec) {
	m.ListenerSpecs = v
}

func (m *UpdateNetworkLoadBalancerRequest) SetAttachedTargetGroups(v []*AttachedTargetGroup) {
	m.AttachedTargetGroups = v
}

func (m *UpdateNetworkLoadBalancerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *DeleteNetworkLoadBalancerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *DeleteNetworkLoadBalancerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *StartNetworkLoadBalancerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *StartNetworkLoadBalancerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *StopNetworkLoadBalancerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *StopNetworkLoadBalancerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *AttachNetworkLoadBalancerTargetGroupRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *AttachNetworkLoadBalancerTargetGroupRequest) SetAttachedTargetGroup(v *AttachedTargetGroup) {
	m.AttachedTargetGroup = v
}

func (m *AttachNetworkLoadBalancerTargetGroupMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *AttachNetworkLoadBalancerTargetGroupMetadata) SetTargetGroupId(v string) {
	m.TargetGroupId = v
}

func (m *DetachNetworkLoadBalancerTargetGroupRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *DetachNetworkLoadBalancerTargetGroupRequest) SetTargetGroupId(v string) {
	m.TargetGroupId = v
}

func (m *DetachNetworkLoadBalancerTargetGroupMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *DetachNetworkLoadBalancerTargetGroupMetadata) SetTargetGroupId(v string) {
	m.TargetGroupId = v
}

func (m *AddNetworkLoadBalancerListenerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *AddNetworkLoadBalancerListenerRequest) SetListenerSpec(v *ListenerSpec) {
	m.ListenerSpec = v
}

func (m *AddNetworkLoadBalancerListenerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *RemoveNetworkLoadBalancerListenerRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *RemoveNetworkLoadBalancerListenerRequest) SetListenerName(v string) {
	m.ListenerName = v
}

func (m *RemoveNetworkLoadBalancerListenerMetadata) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *ListNetworkLoadBalancerOperationsRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *ListNetworkLoadBalancerOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListNetworkLoadBalancerOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListNetworkLoadBalancerOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListNetworkLoadBalancerOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *GetTargetStatesRequest) SetNetworkLoadBalancerId(v string) {
	m.NetworkLoadBalancerId = v
}

func (m *GetTargetStatesRequest) SetTargetGroupId(v string) {
	m.TargetGroupId = v
}

func (m *GetTargetStatesResponse) SetTargetStates(v []*TargetState) {
	m.TargetStates = v
}

func (m *ExternalAddressSpec) SetAddress(v string) {
	m.Address = v
}

func (m *ExternalAddressSpec) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

func (m *InternalAddressSpec) SetAddress(v string) {
	m.Address = v
}

func (m *InternalAddressSpec) SetSubnetId(v string) {
	m.SubnetId = v
}

func (m *InternalAddressSpec) SetIpVersion(v IpVersion) {
	m.IpVersion = v
}

type ListenerSpec_Address = isListenerSpec_Address

func (m *ListenerSpec) SetAddress(v ListenerSpec_Address) {
	m.Address = v
}

func (m *ListenerSpec) SetName(v string) {
	m.Name = v
}

func (m *ListenerSpec) SetPort(v int64) {
	m.Port = v
}

func (m *ListenerSpec) SetProtocol(v Listener_Protocol) {
	m.Protocol = v
}

func (m *ListenerSpec) SetExternalAddressSpec(v *ExternalAddressSpec) {
	m.Address = &ListenerSpec_ExternalAddressSpec{
		ExternalAddressSpec: v,
	}
}

func (m *ListenerSpec) SetInternalAddressSpec(v *InternalAddressSpec) {
	m.Address = &ListenerSpec_InternalAddressSpec{
		InternalAddressSpec: v,
	}
}

func (m *ListenerSpec) SetTargetPort(v int64) {
	m.TargetPort = v
}
