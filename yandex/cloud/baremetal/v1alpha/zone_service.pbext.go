// Code generated by protoc-gen-goext. DO NOT EDIT.

package baremetal

func (m *GetZoneRequest) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *ListZonesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListZonesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListZonesResponse) SetZones(v []*Zone) {
	m.Zones = v
}

func (m *ListZonesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
