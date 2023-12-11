// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type StreamLine_InputType = isStreamLine_InputType

func (m *StreamLine) SetInputType(v StreamLine_InputType) {
	m.InputType = v
}

type StreamLine_LineType = isStreamLine_LineType

func (m *StreamLine) SetLineType(v StreamLine_LineType) {
	m.LineType = v
}

func (m *StreamLine) SetId(v string) {
	m.Id = v
}

func (m *StreamLine) SetChannelId(v string) {
	m.ChannelId = v
}

func (m *StreamLine) SetTitle(v string) {
	m.Title = v
}

func (m *StreamLine) SetThumbnailId(v string) {
	m.ThumbnailId = v
}

func (m *StreamLine) SetRtmpPush(v *RTMPPushInput) {
	m.InputType = &StreamLine_RtmpPush{
		RtmpPush: v,
	}
}

func (m *StreamLine) SetSrtPush(v *SRTPushInput) {
	m.InputType = &StreamLine_SrtPush{
		SrtPush: v,
	}
}

func (m *StreamLine) SetRtmpPull(v *RTMPPullInput) {
	m.InputType = &StreamLine_RtmpPull{
		RtmpPull: v,
	}
}

func (m *StreamLine) SetSrtPull(v *SRTPullInput) {
	m.InputType = &StreamLine_SrtPull{
		SrtPull: v,
	}
}

func (m *StreamLine) SetTcpPull(v *TCPPullInput) {
	m.InputType = &StreamLine_TcpPull{
		TcpPull: v,
	}
}

func (m *StreamLine) SetRtspPull(v *RTSPPullInput) {
	m.InputType = &StreamLine_RtspPull{
		RtspPull: v,
	}
}

func (m *StreamLine) SetManualLine(v *ManualLine) {
	m.LineType = &StreamLine_ManualLine{
		ManualLine: v,
	}
}

func (m *StreamLine) SetAutoLine(v *AutoLine) {
	m.LineType = &StreamLine_AutoLine{
		AutoLine: v,
	}
}

func (m *StreamLine) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *StreamLine) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *PushStreamKey) SetKey(v string) {
	m.Key = v
}

func (m *RTMPPushInput) SetUrl(v string) {
	m.Url = v
}

func (m *SRTPushInput) SetUrl(v string) {
	m.Url = v
}

func (m *RTMPPullInput) SetUrl(v string) {
	m.Url = v
}

func (m *SRTPullInput) SetUrl(v string) {
	m.Url = v
}

func (m *TCPPullInput) SetUrl(v string) {
	m.Url = v
}

func (m *RTSPPullInput) SetUrl(v string) {
	m.Url = v
}

func (m *AutoLine) SetStatus(v AutoLine_AutoLineStatus) {
	m.Status = v
}