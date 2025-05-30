// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

type StreamTalkRequest_Event = isStreamTalkRequest_Event

func (m *StreamTalkRequest) SetEvent(v StreamTalkRequest_Event) {
	m.Event = v
}

func (m *StreamTalkRequest) SetMetadata(v *TalkMetadata) {
	m.Event = &StreamTalkRequest_Metadata{
		Metadata: v,
	}
}

func (m *StreamTalkRequest) SetAudio(v *AudioStreamingRequest) {
	m.Event = &StreamTalkRequest_Audio{
		Audio: v,
	}
}

func (m *UploadTalkRequest) SetTalkId(v string) {
	m.TalkId = v
}

func (m *UploadTalkRequest) SetMetadata(v *TalkMetadata) {
	m.Metadata = v
}

func (m *UploadTalkRequest) SetAudio(v *AudioRequest) {
	m.Audio = v
}

func (m *UploadTalkResponse) SetTalkId(v string) {
	m.TalkId = v
}

func (m *UploadTextRequest) SetTalkId(v string) {
	m.TalkId = v
}

func (m *UploadTextRequest) SetMetadata(v *TalkMetadata) {
	m.Metadata = v
}

func (m *UploadTextRequest) SetTextContent(v *TextContent) {
	m.TextContent = v
}

func (m *UploadTextResponse) SetTalkId(v string) {
	m.TalkId = v
}

func (m *TalkMetadata) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *TalkMetadata) SetFields(v map[string]string) {
	m.Fields = v
}

func (m *TalkMetadata) SetUsers(v []*UserMetadata) {
	m.Users = v
}

func (m *UserMetadata) SetId(v string) {
	m.Id = v
}

func (m *UserMetadata) SetRole(v UserRole) {
	m.Role = v
}

func (m *UserMetadata) SetFields(v map[string]string) {
	m.Fields = v
}

func (m *SearchTalkRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *SearchTalkRequest) SetSpaceId(v string) {
	m.SpaceId = v
}

func (m *SearchTalkRequest) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *SearchTalkRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *SearchTalkRequest) SetFilters(v []*Filter) {
	m.Filters = v
}

func (m *SearchTalkRequest) SetQuery(v *Query) {
	m.Query = v
}

func (m *SearchTalkRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *SearchTalkRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *SearchTalkRequest) SetSortData(v *SortData) {
	m.SortData = v
}

func (m *SearchTalkResponse) SetTalkIds(v []string) {
	m.TalkIds = v
}

func (m *SearchTalkResponse) SetTalksCount(v int64) {
	m.TalksCount = v
}

func (m *SearchTalkResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *GetTalkRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *GetTalkRequest) SetSpaceId(v string) {
	m.SpaceId = v
}

func (m *GetTalkRequest) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *GetTalkRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *GetTalkRequest) SetTalkIds(v []string) {
	m.TalkIds = v
}

func (m *GetTalkRequest) SetResultsMask(v *fieldmaskpb.FieldMask) {
	m.ResultsMask = v
}

func (m *GetTalkResponse) SetTalk(v []*Talk) {
	m.Talk = v
}

func (m *UploadBadgeMetadata) SetBadgeId(v string) {
	m.BadgeId = v
}

func (m *UploadBadgeResponse) SetBadgeId(v string) {
	m.BadgeId = v
}

func (m *UploadBadgeResponse) SetTalkIds(v []string) {
	m.TalkIds = v
}
