// Code generated by protoc-gen-goext. DO NOT EDIT.

package lockbox

import (
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
)

type PayloadEntryChange_Value = isPayloadEntryChange_Value

func (m *PayloadEntryChange) SetValue(v PayloadEntryChange_Value) {
	m.Value = v
}

func (m *PayloadEntryChange) SetKey(v string) {
	m.Key = v
}

func (m *PayloadEntryChange) SetTextValue(v string) {
	m.Value = &PayloadEntryChange_TextValue{
		TextValue: v,
	}
}

func (m *PayloadEntryChange) SetBinaryValue(v []byte) {
	m.Value = &PayloadEntryChange_BinaryValue{
		BinaryValue: v,
	}
}

func (m *GetSecretRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ListSecretsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListSecretsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListSecretsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListSecretsResponse) SetSecrets(v []*Secret) {
	m.Secrets = v
}

func (m *ListSecretsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreateSecretRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateSecretRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateSecretRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateSecretRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateSecretRequest) SetKmsKeyId(v string) {
	m.KmsKeyId = v
}

func (m *CreateSecretRequest) SetVersionDescription(v string) {
	m.VersionDescription = v
}

func (m *CreateSecretRequest) SetVersionPayloadEntries(v []*PayloadEntryChange) {
	m.VersionPayloadEntries = v
}

func (m *CreateSecretRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *CreateSecretMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *CreateSecretMetadata) SetVersionId(v string) {
	m.VersionId = v
}

func (m *UpdateSecretRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *UpdateSecretRequest) SetUpdateMask(v *field_mask.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateSecretRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateSecretRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateSecretRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateSecretRequest) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *UpdateSecretMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *DeleteSecretRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *DeleteSecretMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ActivateSecretRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ActivateSecretMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *DeactivateSecretRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *DeactivateSecretMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *AddVersionRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *AddVersionRequest) SetDescription(v string) {
	m.Description = v
}

func (m *AddVersionRequest) SetPayloadEntries(v []*PayloadEntryChange) {
	m.PayloadEntries = v
}

func (m *AddVersionRequest) SetBaseVersionId(v string) {
	m.BaseVersionId = v
}

func (m *AddVersionMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *AddVersionMetadata) SetVersionId(v string) {
	m.VersionId = v
}

func (m *ListVersionsRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ListVersionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListVersionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListVersionsResponse) SetVersions(v []*Version) {
	m.Versions = v
}

func (m *ListVersionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ScheduleVersionDestructionRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ScheduleVersionDestructionRequest) SetVersionId(v string) {
	m.VersionId = v
}

func (m *ScheduleVersionDestructionRequest) SetPendingPeriod(v *duration.Duration) {
	m.PendingPeriod = v
}

func (m *ScheduleVersionDestructionMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ScheduleVersionDestructionMetadata) SetVersionId(v string) {
	m.VersionId = v
}

func (m *ScheduleVersionDestructionMetadata) SetDestroyAt(v *timestamp.Timestamp) {
	m.DestroyAt = v
}

func (m *CancelVersionDestructionRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *CancelVersionDestructionRequest) SetVersionId(v string) {
	m.VersionId = v
}

func (m *CancelVersionDestructionMetadata) SetSecretId(v string) {
	m.SecretId = v
}

func (m *CancelVersionDestructionMetadata) SetVersionId(v string) {
	m.VersionId = v
}

func (m *ListSecretOperationsRequest) SetSecretId(v string) {
	m.SecretId = v
}

func (m *ListSecretOperationsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListSecretOperationsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListSecretOperationsResponse) SetOperations(v []*operation.Operation) {
	m.Operations = v
}

func (m *ListSecretOperationsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}