// Code generated by protoc-gen-goext. DO NOT EDIT.

package cloudrouter

func (m *RoutingInstance) SetId(v string) {
	m.Id = v
}

func (m *RoutingInstance) SetName(v string) {
	m.Name = v
}

func (m *RoutingInstance) SetDescription(v string) {
	m.Description = v
}

func (m *RoutingInstance) SetFolderId(v string) {
	m.FolderId = v
}

func (m *RoutingInstance) SetRegionId(v string) {
	m.RegionId = v
}

func (m *RoutingInstance) SetVpcInfo(v []*RoutingInstance_VpcInfo) {
	m.VpcInfo = v
}

func (m *RoutingInstance) SetCicPrivateConnectionInfo(v []*RoutingInstance_CicPrivateConnectionInfo) {
	m.CicPrivateConnectionInfo = v
}

func (m *RoutingInstance) SetStatus(v RoutingInstance_Status) {
	m.Status = v
}

func (m *RoutingInstance) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *RoutingInstance_CicPrivateConnectionInfo) SetCicPrivateConnectionId(v string) {
	m.CicPrivateConnectionId = v
}

func (m *RoutingInstance_VpcInfo) SetVpcNetworkId(v string) {
	m.VpcNetworkId = v
}

func (m *RoutingInstance_VpcInfo) SetAzInfos(v []*RoutingInstance_VpcAzInfo) {
	m.AzInfos = v
}

func (m *RoutingInstance_VpcAzInfo) SetManualInfo(v *RoutingInstance_VpcManualInfo) {
	m.ManualInfo = v
}

func (m *RoutingInstance_VpcManualInfo) SetAzId(v string) {
	m.AzId = v
}

func (m *RoutingInstance_VpcManualInfo) SetPrefixes(v []string) {
	m.Prefixes = v
}
