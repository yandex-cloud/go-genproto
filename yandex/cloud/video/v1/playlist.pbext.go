// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Playlist) SetId(v string) {
	m.Id = v
}

func (m *Playlist) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *Playlist) SetTitle(v string) {
	m.Title = v
}

func (m *Playlist) SetDescription(v string) {
	m.Description = v
}

func (m *Playlist) SetItems(v []*PlaylistItem) {
	m.Items = v
}

func (m *Playlist) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Playlist) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

type PlaylistItem_Id = isPlaylistItem_Id

func (m *PlaylistItem) SetId(v PlaylistItem_Id) {
	m.Id = v
}

func (m *PlaylistItem) SetVideoId(v string) {
	m.Id = &PlaylistItem_VideoId{
		VideoId: v,
	}
}

func (m *PlaylistItem) SetEpisodeId(v string) {
	m.Id = &PlaylistItem_EpisodeId{
		EpisodeId: v,
	}
}

func (m *PlaylistItem) SetPosition(v int64) {
	m.Position = v
}
