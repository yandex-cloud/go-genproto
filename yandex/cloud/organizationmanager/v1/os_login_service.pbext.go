// Code generated by protoc-gen-goext. DO NOT EDIT.

package organizationmanager

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *GetOsLoginSettingsRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *OsLoginSettings) SetUserSshKeySettings(v *UserSshKeySettings) {
	m.UserSshKeySettings = v
}

func (m *OsLoginSettings) SetSshCertificateSettings(v *SshCertificateSettings) {
	m.SshCertificateSettings = v
}

func (m *UserSshKeySettings) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *UserSshKeySettings) SetAllowManageOwnKeys(v bool) {
	m.AllowManageOwnKeys = v
}

func (m *SshCertificateSettings) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *UpdateOsLoginSettingsRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *UpdateOsLoginSettingsRequest) SetUserSshKeySettings(v *UpdateOsLoginSettingsRequest_UserSshKeySettings) {
	m.UserSshKeySettings = v
}

func (m *UpdateOsLoginSettingsRequest) SetSshCertificateSettings(v *UpdateOsLoginSettingsRequest_SshCertificateSettings) {
	m.SshCertificateSettings = v
}

func (m *UpdateOsLoginSettingsRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateOsLoginSettingsRequest_UserSshKeySettings) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *UpdateOsLoginSettingsRequest_UserSshKeySettings) SetAllowManageOwnKeys(v bool) {
	m.AllowManageOwnKeys = v
}

func (m *UpdateOsLoginSettingsRequest_SshCertificateSettings) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *SetDefaultOsLoginProfileRequest) SetOsLoginProfileId(v string) {
	m.OsLoginProfileId = v
}

func (m *GetOsLoginProfileRequest) SetOsLoginProfileId(v string) {
	m.OsLoginProfileId = v
}

func (m *ListOsLoginProfilesRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *ListOsLoginProfilesRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListOsLoginProfilesRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListOsLoginProfilesRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListOsLoginProfilesResponse) SetProfiles(v []*OsLoginProfile) {
	m.Profiles = v
}

func (m *ListOsLoginProfilesResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *OsLoginProfile) SetId(v string) {
	m.Id = v
}

func (m *OsLoginProfile) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *OsLoginProfile) SetSubjectId(v string) {
	m.SubjectId = v
}

func (m *OsLoginProfile) SetLogin(v string) {
	m.Login = v
}

func (m *OsLoginProfile) SetUid(v int64) {
	m.Uid = v
}

func (m *OsLoginProfile) SetIsDefault(v bool) {
	m.IsDefault = v
}

func (m *OsLoginProfile) SetHomeDirectory(v string) {
	m.HomeDirectory = v
}

func (m *OsLoginProfile) SetShell(v string) {
	m.Shell = v
}

func (m *UpdateOsLoginProfileRequest) SetOsLoginProfileId(v string) {
	m.OsLoginProfileId = v
}

func (m *UpdateOsLoginProfileRequest) SetLogin(v string) {
	m.Login = v
}

func (m *UpdateOsLoginProfileRequest) SetUid(v int64) {
	m.Uid = v
}

func (m *UpdateOsLoginProfileRequest) SetHomeDirectory(v string) {
	m.HomeDirectory = v
}

func (m *UpdateOsLoginProfileRequest) SetShell(v string) {
	m.Shell = v
}

func (m *UpdateOsLoginProfileRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *DeleteOsLoginProfileRequest) SetId(v string) {
	m.Id = v
}

func (m *CreateOsLoginProfileRequest) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *CreateOsLoginProfileRequest) SetSubjectId(v string) {
	m.SubjectId = v
}

func (m *CreateOsLoginProfileRequest) SetLogin(v string) {
	m.Login = v
}

func (m *CreateOsLoginProfileRequest) SetUid(v int64) {
	m.Uid = v
}

func (m *CreateOsLoginProfileRequest) SetHomeDirectory(v string) {
	m.HomeDirectory = v
}

func (m *CreateOsLoginProfileRequest) SetShell(v string) {
	m.Shell = v
}

func (m *UpdateOsLoginProfileMetadata) SetOsLoginProfileId(v string) {
	m.OsLoginProfileId = v
}

func (m *DeleteOsLoginProfileMetadata) SetOsLoginProfileId(v string) {
	m.OsLoginProfileId = v
}

func (m *CreateOsLoginProfileMetadata) SetOsLoginProfileId(v string) {
	m.OsLoginProfileId = v
}

func (m *CreateOsLoginProfileMetadata) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *CreateOsLoginProfileMetadata) SetSubjectId(v string) {
	m.SubjectId = v
}

func (m *UpdateOsLoginSettingsMetadata) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *SetDefaultOsLoginProfileMetadata) SetPreviousDefaultProfileId(v string) {
	m.PreviousDefaultProfileId = v
}

func (m *SetDefaultOsLoginProfileMetadata) SetCurrentDefaultProfileId(v string) {
	m.CurrentDefaultProfileId = v
}
