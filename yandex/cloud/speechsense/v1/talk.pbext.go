// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

import (
	analysis "github.com/yandex-cloud/go-genproto/yandex/cloud/speechsense/v1/analysis"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Talk) SetId(v string) {
	m.Id = v
}

func (m *Talk) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *Talk) SetSpaceId(v string) {
	m.SpaceId = v
}

func (m *Talk) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *Talk) SetProjectIds(v []string) {
	m.ProjectIds = v
}

func (m *Talk) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *Talk) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Talk) SetModifiedBy(v string) {
	m.ModifiedBy = v
}

func (m *Talk) SetModifiedAt(v *timestamppb.Timestamp) {
	m.ModifiedAt = v
}

func (m *Talk) SetTalkFields(v []*Field) {
	m.TalkFields = v
}

func (m *Talk) SetTranscription(v *analysis.Transcription) {
	m.Transcription = v
}

func (m *Talk) SetSpeechStatistics(v *analysis.SpeechStatistics) {
	m.SpeechStatistics = v
}

func (m *Talk) SetSilenceStatistics(v *analysis.SilenceStatistics) {
	m.SilenceStatistics = v
}

func (m *Talk) SetInterruptsStatistics(v *analysis.InterruptsStatistics) {
	m.InterruptsStatistics = v
}

func (m *Talk) SetConversationStatistics(v *analysis.ConversationStatistics) {
	m.ConversationStatistics = v
}

func (m *Talk) SetPoints(v *analysis.Points) {
	m.Points = v
}

func (m *Talk) SetTextClassifiers(v *analysis.TextClassifiers) {
	m.TextClassifiers = v
}

func (m *Field) SetName(v string) {
	m.Name = v
}

func (m *Field) SetValue(v string) {
	m.Value = v
}

func (m *Field) SetType(v FieldType) {
	m.Type = v
}
