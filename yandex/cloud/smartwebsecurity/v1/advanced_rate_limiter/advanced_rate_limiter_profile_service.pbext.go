// Code generated by protoc-gen-goext. DO NOT EDIT.

package smartwebsecurity

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetAdvancedRateLimiterProfileRequest) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}

func (m *ListAdvancedRateLimiterProfilesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListAdvancedRateLimiterProfilesResponse) SetAdvancedRateLimiterProfiles(v []*AdvancedRateLimiterProfile) {
	m.AdvancedRateLimiterProfiles = v
}

func (m *CreateAdvancedRateLimiterProfileRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateAdvancedRateLimiterProfileRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateAdvancedRateLimiterProfileRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateAdvancedRateLimiterProfileRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateAdvancedRateLimiterProfileRequest) SetAdvancedRateLimiterRules(v []*AdvancedRateLimiterRule) {
	m.AdvancedRateLimiterRules = v
}

func (m *CreateAdvancedRateLimiterProfileMetadata) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}

func (m *UpdateAdvancedRateLimiterProfileRequest) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}

func (m *UpdateAdvancedRateLimiterProfileRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateAdvancedRateLimiterProfileRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateAdvancedRateLimiterProfileRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateAdvancedRateLimiterProfileRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateAdvancedRateLimiterProfileRequest) SetAdvancedRateLimiterRules(v []*AdvancedRateLimiterRule) {
	m.AdvancedRateLimiterRules = v
}

func (m *UpdateAdvancedRateLimiterProfileMetadata) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}

func (m *DeleteAdvancedRateLimiterProfileRequest) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}

func (m *DeleteAdvancedRateLimiterProfileMetadata) SetAdvancedRateLimiterProfileId(v string) {
	m.AdvancedRateLimiterProfileId = v
}
