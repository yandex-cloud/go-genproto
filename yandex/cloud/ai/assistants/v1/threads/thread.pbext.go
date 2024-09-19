// Code generated by protoc-gen-goext. DO NOT EDIT.

package threads

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/common"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Thread) SetId(v string) {
	m.Id = v
}

func (m *Thread) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Thread) SetName(v string) {
	m.Name = v
}

func (m *Thread) SetDescription(v string) {
	m.Description = v
}

func (m *Thread) SetDefaultMessageAuthorId(v string) {
	m.DefaultMessageAuthorId = v
}

func (m *Thread) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *Thread) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Thread) SetUpdatedBy(v string) {
	m.UpdatedBy = v
}

func (m *Thread) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Thread) SetExpirationConfig(v *common.ExpirationConfig) {
	m.ExpirationConfig = v
}

func (m *Thread) SetExpiresAt(v *timestamppb.Timestamp) {
	m.ExpiresAt = v
}

func (m *Thread) SetLabels(v map[string]string) {
	m.Labels = v
}
