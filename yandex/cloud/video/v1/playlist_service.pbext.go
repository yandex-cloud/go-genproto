// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetPlaylistRequest) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *ListPlaylistsRequest) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *ListPlaylistsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListPlaylistsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListPlaylistsRequest) SetOrderBy(v string) {
	m.OrderBy = v
}

func (m *ListPlaylistsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListPlaylistsResponse) SetPlaylists(v []*Playlist) {
	m.Playlists = v
}

func (m *ListPlaylistsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *CreatePlaylistRequest) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *CreatePlaylistRequest) SetTitle(v string) {
	m.Title = v
}

func (m *CreatePlaylistRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreatePlaylistRequest) SetItems(v []*PlaylistItem) {
	m.Items = v
}

func (m *CreatePlaylistMetadata) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *UpdatePlaylistRequest) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *UpdatePlaylistRequest) SetFieldMask(v *fieldmaskpb.FieldMask) {
	m.FieldMask = v
}

func (m *UpdatePlaylistRequest) SetTitle(v string) {
	m.Title = v
}

func (m *UpdatePlaylistRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdatePlaylistRequest) SetItems(v []*PlaylistItem) {
	m.Items = v
}

func (m *UpdatePlaylistMetadata) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *DeletePlaylistRequest) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *DeletePlaylistMetadata) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *BatchDeletePlaylistsRequest) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *BatchDeletePlaylistsRequest) SetPlaylistIds(v []string) {
	m.PlaylistIds = v
}

func (m *BatchDeletePlaylistsMetadata) SetPlaylistIds(v []string) {
	m.PlaylistIds = v
}

func (m *GetPlaylistPlayerURLRequest) SetPlaylistId(v string) {
	m.PlaylistId = v
}

func (m *GetPlaylistPlayerURLRequest) SetParams(v *PlaylistPlayerParams) {
	m.Params = v
}

func (m *PlaylistPlayerParams) SetMute(v bool) {
	m.Mute = v
}

func (m *PlaylistPlayerParams) SetAutoplay(v bool) {
	m.Autoplay = v
}

func (m *PlaylistPlayerParams) SetHidden(v bool) {
	m.Hidden = v
}

func (m *GetPlaylistPlayerURLResponse) SetPlayerUrl(v string) {
	m.PlayerUrl = v
}

func (m *GetPlaylistPlayerURLResponse) SetHtml(v string) {
	m.Html = v
}