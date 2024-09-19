// Code generated by protoc-gen-goext. DO NOT EDIT.

package searchindex

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/common"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type SearchIndex_IndexType = isSearchIndex_IndexType

func (m *SearchIndex) SetIndexType(v SearchIndex_IndexType) {
	m.IndexType = v
}

func (m *SearchIndex) SetId(v string) {
	m.Id = v
}

func (m *SearchIndex) SetFolderId(v string) {
	m.FolderId = v
}

func (m *SearchIndex) SetName(v string) {
	m.Name = v
}

func (m *SearchIndex) SetDescription(v string) {
	m.Description = v
}

func (m *SearchIndex) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *SearchIndex) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *SearchIndex) SetUpdatedBy(v string) {
	m.UpdatedBy = v
}

func (m *SearchIndex) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *SearchIndex) SetExpirationConfig(v *common.ExpirationConfig) {
	m.ExpirationConfig = v
}

func (m *SearchIndex) SetExpiresAt(v *timestamppb.Timestamp) {
	m.ExpiresAt = v
}

func (m *SearchIndex) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *SearchIndex) SetTextSearchIndex(v *TextSearchIndex) {
	m.IndexType = &SearchIndex_TextSearchIndex{
		TextSearchIndex: v,
	}
}

func (m *SearchIndex) SetVectorSearchIndex(v *VectorSearchIndex) {
	m.IndexType = &SearchIndex_VectorSearchIndex{
		VectorSearchIndex: v,
	}
}

func (m *TextSearchIndex) SetChunkingStrategy(v *ChunkingStrategy) {
	m.ChunkingStrategy = v
}

func (m *VectorSearchIndex) SetDocEmbedderUri(v string) {
	m.DocEmbedderUri = v
}

func (m *VectorSearchIndex) SetQueryEmbedderUri(v string) {
	m.QueryEmbedderUri = v
}

func (m *VectorSearchIndex) SetChunkingStrategy(v *ChunkingStrategy) {
	m.ChunkingStrategy = v
}
