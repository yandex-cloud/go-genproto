// Code generated by protoc-gen-goext. DO NOT EDIT.

package opensearch

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Backup) SetId(v string) {
	m.Id = v
}

func (m *Backup) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Backup) SetSourceClusterId(v string) {
	m.SourceClusterId = v
}

func (m *Backup) SetStartedAt(v *timestamppb.Timestamp) {
	m.StartedAt = v
}

func (m *Backup) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Backup) SetIndices(v []string) {
	m.Indices = v
}

func (m *Backup) SetOpensearchVersion(v string) {
	m.OpensearchVersion = v
}

func (m *Backup) SetSizeBytes(v int64) {
	m.SizeBytes = v
}

func (m *Backup) SetIndicesTotal(v int64) {
	m.IndicesTotal = v
}

func (m *SnapshotManagement) SetSnapshotSchedule(v *SnapshotSchedule) {
	m.SnapshotSchedule = v
}

func (m *SnapshotManagement) SetSnapshotMaxAgeDays(v *wrapperspb.Int64Value) {
	m.SnapshotMaxAgeDays = v
}

type SnapshotSchedule_Schedule = isSnapshotSchedule_Schedule

func (m *SnapshotSchedule) SetSchedule(v SnapshotSchedule_Schedule) {
	m.Schedule = v
}

func (m *SnapshotSchedule) SetHourlySnapshotSchedule(v *HourlySnapshotSchedule) {
	m.Schedule = &SnapshotSchedule_HourlySnapshotSchedule{
		HourlySnapshotSchedule: v,
	}
}

func (m *SnapshotSchedule) SetDailySnapshotSchedule(v *DailySnapshotSchedule) {
	m.Schedule = &SnapshotSchedule_DailySnapshotSchedule{
		DailySnapshotSchedule: v,
	}
}

func (m *SnapshotSchedule) SetWeeklySnapshotSchedule(v *WeeklySnapshotSchedule) {
	m.Schedule = &SnapshotSchedule_WeeklySnapshotSchedule{
		WeeklySnapshotSchedule: v,
	}
}

func (m *HourlySnapshotSchedule) SetMinute(v int64) {
	m.Minute = v
}

func (m *DailySnapshotSchedule) SetHour(v int64) {
	m.Hour = v
}

func (m *DailySnapshotSchedule) SetMinute(v int64) {
	m.Minute = v
}

func (m *WeeklySnapshotSchedule) SetDay(v WeeklySnapshotSchedule_WeekDay) {
	m.Day = v
}

func (m *WeeklySnapshotSchedule) SetHour(v int64) {
	m.Hour = v
}

func (m *WeeklySnapshotSchedule) SetMinute(v int64) {
	m.Minute = v
}
