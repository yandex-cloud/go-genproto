// Code generated by protoc-gen-goext. DO NOT EDIT.

package apigateway

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *ApiGateway) SetId(v string) {
	m.Id = v
}

func (m *ApiGateway) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ApiGateway) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *ApiGateway) SetName(v string) {
	m.Name = v
}

func (m *ApiGateway) SetDescription(v string) {
	m.Description = v
}

func (m *ApiGateway) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *ApiGateway) SetStatus(v ApiGateway_Status) {
	m.Status = v
}

func (m *ApiGateway) SetDomain(v string) {
	m.Domain = v
}

func (m *ApiGateway) SetLogGroupId(v string) {
	m.LogGroupId = v
}

func (m *ApiGateway) SetAttachedDomains(v []*AttachedDomain) {
	m.AttachedDomains = v
}

func (m *AttachedDomain) SetDomainId(v string) {
	m.DomainId = v
}

func (m *AttachedDomain) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *AttachedDomain) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *AttachedDomain) SetDomain(v string) {
	m.Domain = v
}
