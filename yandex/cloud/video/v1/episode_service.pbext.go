// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *GetEpisodeRequest) SetEpisodeId(v string) {
	m.EpisodeId = v
}

type ListEpisodesRequest_ParentId = isListEpisodesRequest_ParentId

func (m *ListEpisodesRequest) SetParentId(v ListEpisodesRequest_ParentId) {
	m.ParentId = v
}

func (m *ListEpisodesRequest) SetStreamId(v string) {
	m.ParentId = &ListEpisodesRequest_StreamId{
		StreamId: v,
	}
}

func (m *ListEpisodesRequest) SetLineId(v string) {
	m.ParentId = &ListEpisodesRequest_LineId{
		LineId: v,
	}
}

func (m *ListEpisodesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListEpisodesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListEpisodesRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListEpisodesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListEpisodesResponse) SetEpisodes(v []*Episode) {
	m.Episodes = v
}

func (m *ListEpisodesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *BatchGetEpisodesRequest) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *BatchGetEpisodesRequest) SetEpisodeIds(v []string) {
	m.EpisodeIds = v
}

func (m *BatchGetEpisodesResponse) SetEpisodes(v []*Episode) {
	m.Episodes = v
}

type CreateEpisodeRequest_ParentId = isCreateEpisodeRequest_ParentId

func (m *CreateEpisodeRequest) SetParentId(v CreateEpisodeRequest_ParentId) {
	m.ParentId = v
}

type CreateEpisodeRequest_AccessRights = isCreateEpisodeRequest_AccessRights

func (m *CreateEpisodeRequest) SetAccessRights(v CreateEpisodeRequest_AccessRights) {
	m.AccessRights = v
}

func (m *CreateEpisodeRequest) SetStreamId(v string) {
	m.ParentId = &CreateEpisodeRequest_StreamId{
		StreamId: v,
	}
}

func (m *CreateEpisodeRequest) SetLineId(v string) {
	m.ParentId = &CreateEpisodeRequest_LineId{
		LineId: v,
	}
}

func (m *CreateEpisodeRequest) SetTitle(v string) {
	m.Title = v
}

func (m *CreateEpisodeRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateEpisodeRequest) SetThumbnailId(v string) {
	m.ThumbnailId = v
}

func (m *CreateEpisodeRequest) SetStartTime(v *timestamppb.Timestamp) {
	m.StartTime = v
}

func (m *CreateEpisodeRequest) SetFinishTime(v *timestamppb.Timestamp) {
	m.FinishTime = v
}

func (m *CreateEpisodeRequest) SetDvrSeconds(v int64) {
	m.DvrSeconds = v
}

func (m *CreateEpisodeRequest) SetPublicAccess(v *EpisodePublicAccessParams) {
	m.AccessRights = &CreateEpisodeRequest_PublicAccess{
		PublicAccess: v,
	}
}

func (m *CreateEpisodeRequest) SetAuthSystemAccess(v *EpisodeAuthSystemAccessParams) {
	m.AccessRights = &CreateEpisodeRequest_AuthSystemAccess{
		AuthSystemAccess: v,
	}
}

func (m *CreateEpisodeRequest) SetSignUrlAccess(v *EpisodeSignURLAccessParams) {
	m.AccessRights = &CreateEpisodeRequest_SignUrlAccess{
		SignUrlAccess: v,
	}
}

func (m *CreateEpisodeMetadata) SetEpisodeId(v string) {
	m.EpisodeId = v
}

type UpdateEpisodeRequest_AccessRights = isUpdateEpisodeRequest_AccessRights

func (m *UpdateEpisodeRequest) SetAccessRights(v UpdateEpisodeRequest_AccessRights) {
	m.AccessRights = v
}

func (m *UpdateEpisodeRequest) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *UpdateEpisodeRequest) SetFieldMask(v *fieldmaskpb.FieldMask) {
	m.FieldMask = v
}

func (m *UpdateEpisodeRequest) SetTitle(v string) {
	m.Title = v
}

func (m *UpdateEpisodeRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateEpisodeRequest) SetThumbnailId(v string) {
	m.ThumbnailId = v
}

func (m *UpdateEpisodeRequest) SetStartTime(v *timestamppb.Timestamp) {
	m.StartTime = v
}

func (m *UpdateEpisodeRequest) SetFinishTime(v *timestamppb.Timestamp) {
	m.FinishTime = v
}

func (m *UpdateEpisodeRequest) SetDvrSeconds(v int64) {
	m.DvrSeconds = v
}

func (m *UpdateEpisodeRequest) SetPublicAccess(v *EpisodePublicAccessParams) {
	m.AccessRights = &UpdateEpisodeRequest_PublicAccess{
		PublicAccess: v,
	}
}

func (m *UpdateEpisodeRequest) SetAuthSystemAccess(v *EpisodeAuthSystemAccessParams) {
	m.AccessRights = &UpdateEpisodeRequest_AuthSystemAccess{
		AuthSystemAccess: v,
	}
}

func (m *UpdateEpisodeRequest) SetSignUrlAccess(v *EpisodeSignURLAccessParams) {
	m.AccessRights = &UpdateEpisodeRequest_SignUrlAccess{
		SignUrlAccess: v,
	}
}

func (m *UpdateEpisodeMetadata) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *DeleteEpisodeRequest) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *DeleteEpisodeMetadata) SetEpisodeId(v string) {
	m.EpisodeId = v
}

type PerformEpisodeActionRequest_Action = isPerformEpisodeActionRequest_Action

func (m *PerformEpisodeActionRequest) SetAction(v PerformEpisodeActionRequest_Action) {
	m.Action = v
}

func (m *PerformEpisodeActionRequest) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *PerformEpisodeActionRequest) SetPublish(v *PublishEpisodeAction) {
	m.Action = &PerformEpisodeActionRequest_Publish{
		Publish: v,
	}
}

func (m *PerformEpisodeActionRequest) SetUnpublish(v *UnpublishEpisodeAction) {
	m.Action = &PerformEpisodeActionRequest_Unpublish{
		Unpublish: v,
	}
}

func (m *PerformEpisodeActionMetadata) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *GetEpisodePlayerURLRequest) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *GetEpisodePlayerURLRequest) SetParams(v *EpisodePlayerParams) {
	m.Params = v
}

func (m *EpisodePlayerParams) SetMute(v bool) {
	m.Mute = v
}

func (m *EpisodePlayerParams) SetAutoplay(v bool) {
	m.Autoplay = v
}

func (m *EpisodePlayerParams) SetHidden(v bool) {
	m.Hidden = v
}

func (m *GetEpisodePlayerURLResponse) SetPlayerUrl(v string) {
	m.PlayerUrl = v
}

func (m *GetEpisodePlayerURLResponse) SetHtml(v string) {
	m.Html = v
}

func (m *GetEpisodeManifestsRequest) SetEpisodeId(v string) {
	m.EpisodeId = v
}

func (m *GetEpisodeManifestsResponse) SetManifests(v []*Manifest) {
	m.Manifests = v
}
