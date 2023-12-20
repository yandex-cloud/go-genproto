// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetVideoRequest) SetVideoId(v string) {
	m.VideoId = v
}

func (m *ListVideoRequest) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *ListVideoRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListVideoRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListVideoRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListVideoRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListVideoResponse) SetVideos(v []*Video) {
	m.Videos = v
}

func (m *ListVideoResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

type CreateVideoRequest_Source = isCreateVideoRequest_Source

func (m *CreateVideoRequest) SetSource(v CreateVideoRequest_Source) {
	m.Source = v
}

type CreateVideoRequest_AccessRights = isCreateVideoRequest_AccessRights

func (m *CreateVideoRequest) SetAccessRights(v CreateVideoRequest_AccessRights) {
	m.AccessRights = v
}

func (m *CreateVideoRequest) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *CreateVideoRequest) SetTitle(v string) {
	m.Title = v
}

func (m *CreateVideoRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateVideoRequest) SetThumbnailId(v string) {
	m.ThumbnailId = v
}

func (m *CreateVideoRequest) SetTusd(v *VideoTUSDParams) {
	m.Source = &CreateVideoRequest_Tusd{
		Tusd: v,
	}
}

func (m *CreateVideoRequest) SetPublicAccess(v *VideoPublicAccessParams) {
	m.AccessRights = &CreateVideoRequest_PublicAccess{
		PublicAccess: v,
	}
}

func (m *CreateVideoRequest) SetAuthSystemAccess(v *VideoAuthSystemAccessParams) {
	m.AccessRights = &CreateVideoRequest_AuthSystemAccess{
		AuthSystemAccess: v,
	}
}

func (m *VideoTUSDParams) SetFileSize(v int64) {
	m.FileSize = v
}

func (m *CreateVideoMetadata) SetVideoId(v string) {
	m.VideoId = v
}

type UpdateVideoRequest_AccessRights = isUpdateVideoRequest_AccessRights

func (m *UpdateVideoRequest) SetAccessRights(v UpdateVideoRequest_AccessRights) {
	m.AccessRights = v
}

func (m *UpdateVideoRequest) SetVideoId(v string) {
	m.VideoId = v
}

func (m *UpdateVideoRequest) SetFieldMask(v *fieldmaskpb.FieldMask) {
	m.FieldMask = v
}

func (m *UpdateVideoRequest) SetTitle(v string) {
	m.Title = v
}

func (m *UpdateVideoRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateVideoRequest) SetThumbnailId(v string) {
	m.ThumbnailId = v
}

func (m *UpdateVideoRequest) SetPublicAccess(v *VideoPublicAccessParams) {
	m.AccessRights = &UpdateVideoRequest_PublicAccess{
		PublicAccess: v,
	}
}

func (m *UpdateVideoRequest) SetAuthSystemAccess(v *VideoAuthSystemAccessParams) {
	m.AccessRights = &UpdateVideoRequest_AuthSystemAccess{
		AuthSystemAccess: v,
	}
}

func (m *UpdateVideoMetadata) SetVideoId(v string) {
	m.VideoId = v
}

func (m *DeleteVideoRequest) SetVideoId(v string) {
	m.VideoId = v
}

func (m *DeleteVideoMetadata) SetVideoId(v string) {
	m.VideoId = v
}

type PerformVideoActionRequest_Action = isPerformVideoActionRequest_Action

func (m *PerformVideoActionRequest) SetAction(v PerformVideoActionRequest_Action) {
	m.Action = v
}

func (m *PerformVideoActionRequest) SetVideoId(v string) {
	m.VideoId = v
}

func (m *PerformVideoActionRequest) SetPublish(v *PublishVideoAction) {
	m.Action = &PerformVideoActionRequest_Publish{
		Publish: v,
	}
}

func (m *PerformVideoActionRequest) SetUnpublish(v *UnpublishVideoAction) {
	m.Action = &PerformVideoActionRequest_Unpublish{
		Unpublish: v,
	}
}

func (m *PerformVideoActionMetadata) SetVideoId(v string) {
	m.VideoId = v
}

func (m *GetVideoPlayerURLRequest) SetVideoId(v string) {
	m.VideoId = v
}

func (m *GetVideoPlayerURLRequest) SetParams(v *VideoPlayerParams) {
	m.Params = v
}

func (m *VideoPlayerParams) SetMute(v bool) {
	m.Mute = v
}

func (m *VideoPlayerParams) SetAutoplay(v bool) {
	m.Autoplay = v
}

func (m *VideoPlayerParams) SetHidden(v bool) {
	m.Hidden = v
}

func (m *GetVideoPlayerURLResponse) SetPlayerUrl(v string) {
	m.PlayerUrl = v
}

func (m *GetVideoPlayerURLResponse) SetHtml(v string) {
	m.Html = v
}
