// Code generated by protoc-gen-goext. DO NOT EDIT.

package files

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/common"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *File) SetId(v string) {
	m.Id = v
}

func (m *File) SetFolderId(v string) {
	m.FolderId = v
}

func (m *File) SetName(v string) {
	m.Name = v
}

func (m *File) SetDescription(v string) {
	m.Description = v
}

func (m *File) SetMimeType(v string) {
	m.MimeType = v
}

func (m *File) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *File) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *File) SetUpdatedBy(v string) {
	m.UpdatedBy = v
}

func (m *File) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *File) SetExpirationConfig(v *common.ExpirationConfig) {
	m.ExpirationConfig = v
}

func (m *File) SetExpiresAt(v *timestamppb.Timestamp) {
	m.ExpiresAt = v
}

func (m *File) SetLabels(v map[string]string) {
	m.Labels = v
}
