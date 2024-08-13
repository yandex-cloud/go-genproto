// Code generated by protoc-gen-goext. DO NOT EDIT.

package workload

func (m *GetFederatedCredentialRequest) SetFederatedCredentialId(v string) {
	m.FederatedCredentialId = v
}

func (m *ListFederatedCredentialsRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *ListFederatedCredentialsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListFederatedCredentialsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListFederatedCredentialsResponse) SetFederatedCredentials(v []*FederatedCredential) {
	m.FederatedCredentials = v
}

func (m *ListFederatedCredentialsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateFederatedCredentialRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateFederatedCredentialRequest) SetFederationId(v string) {
	m.FederationId = v
}

func (m *CreateFederatedCredentialRequest) SetExternalSubjectId(v string) {
	m.ExternalSubjectId = v
}

func (m *CreateFederatedCredentialMetadata) SetFederatedCredentialId(v string) {
	m.FederatedCredentialId = v
}

func (m *DeleteFederatedCredentialRequest) SetFederatedCredentialId(v string) {
	m.FederatedCredentialId = v
}

func (m *DeleteFederatedCredentialMetadata) SetFederatedCredentialId(v string) {
	m.FederatedCredentialId = v
}
