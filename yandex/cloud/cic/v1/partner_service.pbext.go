// Code generated by protoc-gen-goext. DO NOT EDIT.

package cic

func (m *GetPartnerRequest) SetPartnerId(v string) {
	m.PartnerId = v
}

func (m *ListPartnersRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPartnersRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPartnersRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListPartnersResponse) SetPartners(v []*Partner) {
	m.Partners = v
}

func (m *ListPartnersResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}