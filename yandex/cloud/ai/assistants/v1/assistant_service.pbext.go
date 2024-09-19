// Code generated by protoc-gen-goext. DO NOT EDIT.

package assistants

import (
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/common"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *CreateAssistantRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateAssistantRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateAssistantRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateAssistantRequest) SetExpirationConfig(v *common.ExpirationConfig) {
	m.ExpirationConfig = v
}

func (m *CreateAssistantRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateAssistantRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *CreateAssistantRequest) SetInstruction(v string) {
	m.Instruction = v
}

func (m *CreateAssistantRequest) SetPromptTruncationOptions(v *PromptTruncationOptions) {
	m.PromptTruncationOptions = v
}

func (m *CreateAssistantRequest) SetCompletionOptions(v *CompletionOptions) {
	m.CompletionOptions = v
}

func (m *CreateAssistantRequest) SetTools(v []*Tool) {
	m.Tools = v
}

func (m *GetAssistantRequest) SetAssistantId(v string) {
	m.AssistantId = v
}

func (m *UpdateAssistantRequest) SetAssistantId(v string) {
	m.AssistantId = v
}

func (m *UpdateAssistantRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateAssistantRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateAssistantRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateAssistantRequest) SetExpirationConfig(v *common.ExpirationConfig) {
	m.ExpirationConfig = v
}

func (m *UpdateAssistantRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateAssistantRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *UpdateAssistantRequest) SetInstruction(v string) {
	m.Instruction = v
}

func (m *UpdateAssistantRequest) SetPromptTruncationOptions(v *PromptTruncationOptions) {
	m.PromptTruncationOptions = v
}

func (m *UpdateAssistantRequest) SetCompletionOptions(v *CompletionOptions) {
	m.CompletionOptions = v
}

func (m *UpdateAssistantRequest) SetTools(v []*Tool) {
	m.Tools = v
}

func (m *DeleteAssistantRequest) SetAssistantId(v string) {
	m.AssistantId = v
}

func (m *ListAssistantsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListAssistantsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAssistantsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListAssistantsResponse) SetAssistants(v []*Assistant) {
	m.Assistants = v
}

func (m *ListAssistantsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListAssistantVersionsRequest) SetAssistantId(v string) {
	m.AssistantId = v
}

func (m *ListAssistantVersionsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListAssistantVersionsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *AssistantVersion) SetId(v string) {
	m.Id = v
}

func (m *AssistantVersion) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *AssistantVersion) SetAssistant(v *Assistant) {
	m.Assistant = v
}

func (m *ListAssistantVersionsResponse) SetVersions(v []*AssistantVersion) {
	m.Versions = v
}

func (m *ListAssistantVersionsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}
