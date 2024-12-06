// Code generated by protoc-gen-goext. DO NOT EDIT.

package cic

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PrivateConnection) SetId(v string) {
	m.Id = v
}

func (m *PrivateConnection) SetName(v string) {
	m.Name = v
}

func (m *PrivateConnection) SetDescription(v string) {
	m.Description = v
}

func (m *PrivateConnection) SetFolderId(v string) {
	m.FolderId = v
}

func (m *PrivateConnection) SetRegionId(v string) {
	m.RegionId = v
}

func (m *PrivateConnection) SetTrunkConnectionId(v string) {
	m.TrunkConnectionId = v
}

func (m *PrivateConnection) SetVlanId(v *wrapperspb.Int64Value) {
	m.VlanId = v
}

func (m *PrivateConnection) SetIpv4Peering(v *Peering) {
	m.Ipv4Peering = v
}

func (m *PrivateConnection) SetIpv4StaticRoutes(v []*PrivateConnection_StaticRoute) {
	m.Ipv4StaticRoutes = v
}

func (m *PrivateConnection) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *PrivateConnection_StaticRoute) SetPrefix(v string) {
	m.Prefix = v
}

func (m *PrivateConnection_StaticRoute) SetNextHop(v []string) {
	m.NextHop = v
}