// Code generated by protoc-gen-goext. DO NOT EDIT.

package awscompatibility

import (
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

func (m *GetAccessKeyRequest) SetAccessKeyId(v string) {
	m.AccessKeyId = v
}

func (m *ListAccessKeysRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *ListAccessKeysRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAccessKeysRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAccessKeysResponse) SetAccessKeys(v []*AccessKey) {
	m.AccessKeys = v
}

func (m *ListAccessKeysResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateAccessKeyRequest) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *CreateAccessKeyRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateAccessKeyResponse) SetAccessKey(v *AccessKey) {
	m.AccessKey = v
}

func (m *CreateAccessKeyResponse) SetSecret(v string) {
	m.Secret = v
}

func (m *DeleteAccessKeyRequest) SetAccessKeyId(v string) {
	m.AccessKeyId = v
}

func (m *DeleteAccessKeyMetadata) SetAccessKeyId(v string) {
	m.AccessKeyId = v
}

func (m *ListAccessKeyOperationsRequest) SetAccessKeyId(v string) {
	m.AccessKeyId = v
}

func (m *ListAccessKeyOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAccessKeyOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAccessKeyOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListAccessKeyOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
