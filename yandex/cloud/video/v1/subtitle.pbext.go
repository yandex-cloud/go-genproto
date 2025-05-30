// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Subtitle_ParentId = isSubtitle_ParentId

func (m *Subtitle) SetParentId(v Subtitle_ParentId) {
	m.ParentId = v
}

func (m *Subtitle) SetId(v string) {
	m.Id = v
}

func (m *Subtitle) SetLanguage(v string) {
	m.Language = v
}

func (m *Subtitle) SetLabel(v string) {
	m.Label = v
}

func (m *Subtitle) SetStatus(v Subtitle_SubtitleStatus) {
	m.Status = v
}

func (m *Subtitle) SetSourceType(v Subtitle_SubtitleSourceType) {
	m.SourceType = v
}

func (m *Subtitle) SetFilename(v string) {
	m.Filename = v
}

func (m *Subtitle) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Subtitle) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Subtitle) SetVideoId(v string) {
	m.ParentId = &Subtitle_VideoId{
		VideoId: v,
	}
}
