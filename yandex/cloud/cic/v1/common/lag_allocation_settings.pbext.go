// Code generated by protoc-gen-goext. DO NOT EDIT.

package cic

type LagAllocationSettingsRequest_Lag = isLagAllocationSettingsRequest_Lag

func (m *LagAllocationSettingsRequest) SetLag(v LagAllocationSettingsRequest_Lag) {
	m.Lag = v
}

func (m *LagAllocationSettingsRequest) SetLagSize(v int64) {
	m.Lag = &LagAllocationSettingsRequest_LagSize{
		LagSize: v,
	}
}

func (m *LagAllocationSettings) SetLagInfo(v *LagInfo) {
	m.LagInfo = v
}
