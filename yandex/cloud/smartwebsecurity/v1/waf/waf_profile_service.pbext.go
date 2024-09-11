// Code generated by protoc-gen-goext. DO NOT EDIT.

package smartwebsecurity

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetWafProfileRequest) SetWafProfileId(v string) {
	m.WafProfileId = v
}

func (m *ListWafProfilesRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListWafProfilesResponse) SetWafProfiles(v []*WafProfile) {
	m.WafProfiles = v
}

type CreateWafProfileRequest_RuleSet = isCreateWafProfileRequest_RuleSet

func (m *CreateWafProfileRequest) SetRuleSet(v CreateWafProfileRequest_RuleSet) {
	m.RuleSet = v
}

func (m *CreateWafProfileRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateWafProfileRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateWafProfileRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateWafProfileRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateWafProfileRequest) SetRules(v []*WafProfileRule) {
	m.Rules = v
}

func (m *CreateWafProfileRequest) SetExclusionRules(v []*WafProfileExclusionRule) {
	m.ExclusionRules = v
}

func (m *CreateWafProfileRequest) SetCoreRuleSet(v *WafProfile_CoreRuleSet) {
	m.RuleSet = &CreateWafProfileRequest_CoreRuleSet{
		CoreRuleSet: v,
	}
}

func (m *CreateWafProfileRequest) SetAnalyzeRequestBody(v *WafProfile_AnalyzeRequestBody) {
	m.AnalyzeRequestBody = v
}

func (m *CreateWafProfileMetadata) SetWafProfileId(v string) {
	m.WafProfileId = v
}

type UpdateWafProfileRequest_RuleSet = isUpdateWafProfileRequest_RuleSet

func (m *UpdateWafProfileRequest) SetRuleSet(v UpdateWafProfileRequest_RuleSet) {
	m.RuleSet = v
}

func (m *UpdateWafProfileRequest) SetWafProfileId(v string) {
	m.WafProfileId = v
}

func (m *UpdateWafProfileRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateWafProfileRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateWafProfileRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateWafProfileRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateWafProfileRequest) SetRules(v []*WafProfileRule) {
	m.Rules = v
}

func (m *UpdateWafProfileRequest) SetExclusionRules(v []*WafProfileExclusionRule) {
	m.ExclusionRules = v
}

func (m *UpdateWafProfileRequest) SetCoreRuleSet(v *WafProfile_CoreRuleSet) {
	m.RuleSet = &UpdateWafProfileRequest_CoreRuleSet{
		CoreRuleSet: v,
	}
}

func (m *UpdateWafProfileRequest) SetAnalyzeRequestBody(v *WafProfile_AnalyzeRequestBody) {
	m.AnalyzeRequestBody = v
}

func (m *UpdateWafProfileMetadata) SetWafProfileId(v string) {
	m.WafProfileId = v
}

func (m *DeleteWafProfileRequest) SetWafProfileId(v string) {
	m.WafProfileId = v
}

func (m *DeleteWafProfileMetadata) SetWafProfileId(v string) {
	m.WafProfileId = v
}