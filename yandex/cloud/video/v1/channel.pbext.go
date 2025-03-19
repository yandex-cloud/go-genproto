// Code generated by protoc-gen-goext. DO NOT EDIT.

package video

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Channel) SetId(v string) {
	m.Id = v
}

func (m *Channel) SetOrganizationId(v string) {
	m.OrganizationId = v
}

func (m *Channel) SetTitle(v string) {
	m.Title = v
}

func (m *Channel) SetDescription(v string) {
	m.Description = v
}

func (m *Channel) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Channel) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Channel) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Channel) SetSettings(v *ChannelSettings) {
	m.Settings = v
}

func (m *ChannelSettings) SetAdvertisement(v *AdvertisementSettings) {
	m.Advertisement = v
}

func (m *ChannelSettings) SetRefererVerification(v *RefererVerificationSettings) {
	m.RefererVerification = v
}

type AdvertisementSettings_Provider = isAdvertisementSettings_Provider

func (m *AdvertisementSettings) SetProvider(v AdvertisementSettings_Provider) {
	m.Provider = v
}

func (m *AdvertisementSettings) SetYandexDirect(v *AdvertisementSettings_YandexDirect) {
	m.Provider = &AdvertisementSettings_YandexDirect_{
		YandexDirect: v,
	}
}

func (m *AdvertisementSettings_YandexDirect) SetEnable(v bool) {
	m.Enable = v
}

func (m *AdvertisementSettings_YandexDirect) SetPageId(v int64) {
	m.PageId = v
}

func (m *AdvertisementSettings_YandexDirect) SetCategory(v int64) {
	m.Category = v
}

func (m *RefererVerificationSettings) SetEnable(v bool) {
	m.Enable = v
}

func (m *RefererVerificationSettings) SetAllowedDomains(v []string) {
	m.AllowedDomains = v
}
