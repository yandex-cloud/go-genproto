// Code generated by protoc-gen-goext. DO NOT EDIT.

package cic

func (m *Peering) SetPeeringSubnet(v string) {
	m.PeeringSubnet = v
}

func (m *Peering) SetPeerIp(v string) {
	m.PeerIp = v
}

func (m *Peering) SetCloudIp(v string) {
	m.CloudIp = v
}

func (m *Peering) SetPeerBgpAsn(v int64) {
	m.PeerBgpAsn = v
}

func (m *Peering) SetCloudBgpAsn(v int64) {
	m.CloudBgpAsn = v
}

func (m *Peering) SetPeerBgpMd5Key(v string) {
	m.PeerBgpMd5Key = v
}
