// Code generated by protoc-gen-goext. DO NOT EDIT.

package cloudregistry

func (m *IpPermission) SetAction(v IpPermission_Action) {
	m.Action = v
}

func (m *IpPermission) SetIp(v string) {
	m.Ip = v
}

func (m *IpPermissionDelta) SetAction(v IpPermissionDeltaAction) {
	m.Action = v
}

func (m *IpPermissionDelta) SetIpPermission(v *IpPermission) {
	m.IpPermission = v
}
