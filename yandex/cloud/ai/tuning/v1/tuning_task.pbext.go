// Code generated by protoc-gen-goext. DO NOT EDIT.

package fomo

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *TuningTask) SetTaskId(v string) {
	m.TaskId = v
}

func (m *TuningTask) SetOperationId(v string) {
	m.OperationId = v
}

func (m *TuningTask) SetStatus(v TuningTask_Status) {
	m.Status = v
}

func (m *TuningTask) SetFolderId(v string) {
	m.FolderId = v
}

func (m *TuningTask) SetCreatedBy(v string) {
	m.CreatedBy = v
}

func (m *TuningTask) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *TuningTask) SetStartedAt(v *timestamppb.Timestamp) {
	m.StartedAt = v
}

func (m *TuningTask) SetFinishedAt(v *timestamppb.Timestamp) {
	m.FinishedAt = v
}

func (m *TuningTask) SetSourceModelUri(v string) {
	m.SourceModelUri = v
}

func (m *TuningTask) SetTargetModelUri(v string) {
	m.TargetModelUri = v
}

func (m *TuningTask) SetName(v string) {
	m.Name = v
}

func (m *TuningTask) SetDescription(v string) {
	m.Description = v
}

func (m *TuningTask) SetLabels(v map[string]string) {
	m.Labels = v
}
