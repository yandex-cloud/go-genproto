// Code generated by protoc-gen-goext. DO NOT EDIT.

package datasphere

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *CreateProjectRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateProjectRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateProjectRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateProjectRequest) SetSettings(v *Project_Settings) {
	m.Settings = v
}

func (m *CreateProjectRequest) SetLimits(v *Project_Limits) {
	m.Limits = v
}

func (m *CreateProjectMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *UpdateProjectRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *UpdateProjectRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateProjectRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateProjectRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateProjectRequest) SetSettings(v *Project_Settings) {
	m.Settings = v
}

func (m *UpdateProjectRequest) SetLimits(v *Project_Limits) {
	m.Limits = v
}

func (m *UpdateProjectMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *DeleteProjectRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *DeleteProjectMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *OpenProjectRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *OpenProjectMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *OpenProjectResponse) SetProjectUrl(v string) {
	m.ProjectUrl = v
}

func (m *OpenProjectResponse) SetSessionToken(v string) {
	m.SessionToken = v
}

func (m *GetProjectRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *ListProjectsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListProjectsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListProjectsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListProjectsResponse) SetProjects(v []*Project) {
	m.Projects = v
}

func (m *ListProjectsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *GetUnitBalanceRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *GetUnitBalanceResponse) SetUnitBalance(v *wrapperspb.Int64Value) {
	m.UnitBalance = v
}

func (m *SetUnitBalanceRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *SetUnitBalanceRequest) SetUnitBalance(v *wrapperspb.Int64Value) {
	m.UnitBalance = v
}

type ProjectExecutionRequest_Target = isProjectExecutionRequest_Target

func (m *ProjectExecutionRequest) SetTarget(v ProjectExecutionRequest_Target) {
	m.Target = v
}

func (m *ProjectExecutionRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *ProjectExecutionRequest) SetNotebookId(v string) {
	m.Target = &ProjectExecutionRequest_NotebookId{
		NotebookId: v,
	}
}

func (m *ProjectExecutionRequest) SetCellId(v string) {
	m.Target = &ProjectExecutionRequest_CellId{
		CellId: v,
	}
}

func (m *ProjectExecutionRequest) SetInputVariables(v *structpb.Struct) {
	m.InputVariables = v
}

func (m *ProjectExecutionRequest) SetOutputVariableNames(v []string) {
	m.OutputVariableNames = v
}

type ProjectExecutionMetadata_Target = isProjectExecutionMetadata_Target

func (m *ProjectExecutionMetadata) SetTarget(v ProjectExecutionMetadata_Target) {
	m.Target = v
}

func (m *ProjectExecutionMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *ProjectExecutionMetadata) SetNotebookId(v string) {
	m.Target = &ProjectExecutionMetadata_NotebookId{
		NotebookId: v,
	}
}

func (m *ProjectExecutionMetadata) SetCellId(v string) {
	m.Target = &ProjectExecutionMetadata_CellId{
		CellId: v,
	}
}
