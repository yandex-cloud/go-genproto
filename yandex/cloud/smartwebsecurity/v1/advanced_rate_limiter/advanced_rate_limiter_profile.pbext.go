// Code generated by protoc-gen-goext. DO NOT EDIT.

package smartwebsecurity

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/smartwebsecurity/v1"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *AdvancedRateLimiterProfile) SetId(v string) {
	m.Id = v
}

func (m *AdvancedRateLimiterProfile) SetFolderId(v string) {
	m.FolderId = v
}

func (m *AdvancedRateLimiterProfile) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *AdvancedRateLimiterProfile) SetName(v string) {
	m.Name = v
}

func (m *AdvancedRateLimiterProfile) SetDescription(v string) {
	m.Description = v
}

func (m *AdvancedRateLimiterProfile) SetAdvancedRateLimiterRules(v []*AdvancedRateLimiterRule) {
	m.AdvancedRateLimiterRules = v
}

func (m *AdvancedRateLimiterProfile) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *AdvancedRateLimiterProfile) SetCloudId(v string) {
	m.CloudId = v
}

type AdvancedRateLimiterRule_RuleSpecifier = isAdvancedRateLimiterRule_RuleSpecifier

func (m *AdvancedRateLimiterRule) SetRuleSpecifier(v AdvancedRateLimiterRule_RuleSpecifier) {
	m.RuleSpecifier = v
}

func (m *AdvancedRateLimiterRule) SetName(v string) {
	m.Name = v
}

func (m *AdvancedRateLimiterRule) SetPriority(v int64) {
	m.Priority = v
}

func (m *AdvancedRateLimiterRule) SetDescription(v string) {
	m.Description = v
}

func (m *AdvancedRateLimiterRule) SetDryRun(v bool) {
	m.DryRun = v
}

func (m *AdvancedRateLimiterRule) SetStaticQuota(v *AdvancedRateLimiterRule_StaticQuota) {
	m.RuleSpecifier = &AdvancedRateLimiterRule_StaticQuota_{
		StaticQuota: v,
	}
}

func (m *AdvancedRateLimiterRule) SetDynamicQuota(v *AdvancedRateLimiterRule_DynamicQuota) {
	m.RuleSpecifier = &AdvancedRateLimiterRule_DynamicQuota_{
		DynamicQuota: v,
	}
}

func (m *AdvancedRateLimiterRule_StaticQuota) SetAction(v AdvancedRateLimiterRule_Action) {
	m.Action = v
}

func (m *AdvancedRateLimiterRule_StaticQuota) SetCondition(v *v1.Condition) {
	m.Condition = v
}

func (m *AdvancedRateLimiterRule_StaticQuota) SetLimit(v int64) {
	m.Limit = v
}

func (m *AdvancedRateLimiterRule_StaticQuota) SetPeriod(v int64) {
	m.Period = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota) SetAction(v AdvancedRateLimiterRule_Action) {
	m.Action = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota) SetCondition(v *v1.Condition) {
	m.Condition = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota) SetLimit(v int64) {
	m.Limit = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota) SetPeriod(v int64) {
	m.Period = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota) SetCharacteristics(v []*AdvancedRateLimiterRule_DynamicQuota_Characteristic) {
	m.Characteristics = v
}

type AdvancedRateLimiterRule_DynamicQuota_Characteristic_CharacteristicSpecifier = isAdvancedRateLimiterRule_DynamicQuota_Characteristic_CharacteristicSpecifier

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic) SetCharacteristicSpecifier(v AdvancedRateLimiterRule_DynamicQuota_Characteristic_CharacteristicSpecifier) {
	m.CharacteristicSpecifier = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic) SetSimpleCharacteristic(v *AdvancedRateLimiterRule_DynamicQuota_Characteristic_SimpleCharacteristic) {
	m.CharacteristicSpecifier = &AdvancedRateLimiterRule_DynamicQuota_Characteristic_SimpleCharacteristic_{
		SimpleCharacteristic: v,
	}
}

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic) SetKeyCharacteristic(v *AdvancedRateLimiterRule_DynamicQuota_Characteristic_KeyCharacteristic) {
	m.CharacteristicSpecifier = &AdvancedRateLimiterRule_DynamicQuota_Characteristic_KeyCharacteristic_{
		KeyCharacteristic: v,
	}
}

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic) SetCaseInsensitive(v bool) {
	m.CaseInsensitive = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic_SimpleCharacteristic) SetType(v AdvancedRateLimiterRule_DynamicQuota_Characteristic_SimpleCharacteristic_Type) {
	m.Type = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic_KeyCharacteristic) SetType(v AdvancedRateLimiterRule_DynamicQuota_Characteristic_KeyCharacteristic_Type) {
	m.Type = v
}

func (m *AdvancedRateLimiterRule_DynamicQuota_Characteristic_KeyCharacteristic) SetValue(v string) {
	m.Value = v
}
