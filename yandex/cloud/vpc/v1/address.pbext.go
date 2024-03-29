// Code generated by protoc-gen-goext. DO NOT EDIT.

package vpc

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Address_Address = isAddress_Address

func (m *Address) SetAddress(v Address_Address) {
	m.Address = v
}

func (m *Address) SetId(v string) {
	m.Id = v
}

func (m *Address) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Address) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Address) SetName(v string) {
	m.Name = v
}

func (m *Address) SetDescription(v string) {
	m.Description = v
}

func (m *Address) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Address) SetExternalIpv4Address(v *ExternalIpv4Address) {
	m.Address = &Address_ExternalIpv4Address{
		ExternalIpv4Address: v,
	}
}

func (m *Address) SetReserved(v bool) {
	m.Reserved = v
}

func (m *Address) SetUsed(v bool) {
	m.Used = v
}

func (m *Address) SetType(v Address_Type) {
	m.Type = v
}

func (m *Address) SetIpVersion(v Address_IpVersion) {
	m.IpVersion = v
}

func (m *Address) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Address) SetDnsRecords(v []*DnsRecord) {
	m.DnsRecords = v
}

func (m *ExternalIpv4Address) SetAddress(v string) {
	m.Address = v
}

func (m *ExternalIpv4Address) SetZoneId(v string) {
	m.ZoneId = v
}

func (m *ExternalIpv4Address) SetRequirements(v *AddressRequirements) {
	m.Requirements = v
}

func (m *AddressRequirements) SetDdosProtectionProvider(v string) {
	m.DdosProtectionProvider = v
}

func (m *AddressRequirements) SetOutgoingSmtpCapability(v string) {
	m.OutgoingSmtpCapability = v
}

func (m *DnsRecord) SetFqdn(v string) {
	m.Fqdn = v
}

func (m *DnsRecord) SetDnsZoneId(v string) {
	m.DnsZoneId = v
}

func (m *DnsRecord) SetTtl(v int64) {
	m.Ttl = v
}

func (m *DnsRecord) SetPtr(v bool) {
	m.Ptr = v
}
