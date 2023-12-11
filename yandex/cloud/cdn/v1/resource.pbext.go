// Code generated by protoc-gen-goext. DO NOT EDIT.

package cdn

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *SecondaryHostnames) SetValues(v []string) {
	m.Values = v
}

func (m *Resource) SetId(v string) {
	m.Id = v
}

func (m *Resource) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Resource) SetCname(v string) {
	m.Cname = v
}

func (m *Resource) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Resource) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *Resource) SetActive(v bool) {
	m.Active = v
}

func (m *Resource) SetOptions(v *ResourceOptions) {
	m.Options = v
}

func (m *Resource) SetSecondaryHostnames(v []string) {
	m.SecondaryHostnames = v
}

func (m *Resource) SetOriginGroupId(v int64) {
	m.OriginGroupId = v
}

func (m *Resource) SetOriginGroupName(v string) {
	m.OriginGroupName = v
}

func (m *Resource) SetOriginProtocol(v OriginProtocol) {
	m.OriginProtocol = v
}

func (m *Resource) SetSslCertificate(v *SSLCertificate) {
	m.SslCertificate = v
}

func (m *Resource) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *ResourceOptions) SetDisableCache(v *ResourceOptions_BoolOption) {
	m.DisableCache = v
}

func (m *ResourceOptions) SetEdgeCacheSettings(v *ResourceOptions_EdgeCacheSettings) {
	m.EdgeCacheSettings = v
}

func (m *ResourceOptions) SetBrowserCacheSettings(v *ResourceOptions_Int64Option) {
	m.BrowserCacheSettings = v
}

func (m *ResourceOptions) SetCacheHttpHeaders(v *ResourceOptions_StringsListOption) {
	m.CacheHttpHeaders = v
}

func (m *ResourceOptions) SetQueryParamsOptions(v *ResourceOptions_QueryParamsOptions) {
	m.QueryParamsOptions = v
}

func (m *ResourceOptions) SetSlice(v *ResourceOptions_BoolOption) {
	m.Slice = v
}

func (m *ResourceOptions) SetCompressionOptions(v *ResourceOptions_CompressionOptions) {
	m.CompressionOptions = v
}

func (m *ResourceOptions) SetRedirectOptions(v *ResourceOptions_RedirectOptions) {
	m.RedirectOptions = v
}

func (m *ResourceOptions) SetHostOptions(v *ResourceOptions_HostOptions) {
	m.HostOptions = v
}

func (m *ResourceOptions) SetStaticHeaders(v *ResourceOptions_StringsMapOption) {
	m.StaticHeaders = v
}

func (m *ResourceOptions) SetCors(v *ResourceOptions_StringsListOption) {
	m.Cors = v
}

func (m *ResourceOptions) SetStale(v *ResourceOptions_StringsListOption) {
	m.Stale = v
}

func (m *ResourceOptions) SetAllowedHttpMethods(v *ResourceOptions_StringsListOption) {
	m.AllowedHttpMethods = v
}

func (m *ResourceOptions) SetProxyCacheMethodsSet(v *ResourceOptions_BoolOption) {
	m.ProxyCacheMethodsSet = v
}

func (m *ResourceOptions) SetDisableProxyForceRanges(v *ResourceOptions_BoolOption) {
	m.DisableProxyForceRanges = v
}

func (m *ResourceOptions) SetStaticRequestHeaders(v *ResourceOptions_StringsMapOption) {
	m.StaticRequestHeaders = v
}

func (m *ResourceOptions) SetCustomServerName(v *ResourceOptions_StringOption) {
	m.CustomServerName = v
}

func (m *ResourceOptions) SetIgnoreCookie(v *ResourceOptions_BoolOption) {
	m.IgnoreCookie = v
}

func (m *ResourceOptions) SetRewrite(v *ResourceOptions_RewriteOption) {
	m.Rewrite = v
}

func (m *ResourceOptions) SetSecureKey(v *ResourceOptions_SecureKeyOption) {
	m.SecureKey = v
}

func (m *ResourceOptions) SetIpAddressAcl(v *ResourceOptions_IPAddresACL) {
	m.IpAddressAcl = v
}

func (m *ResourceOptions_BoolOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_BoolOption) SetValue(v bool) {
	m.Value = v
}

func (m *ResourceOptions_StringOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_StringOption) SetValue(v string) {
	m.Value = v
}

func (m *ResourceOptions_Int64Option) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_Int64Option) SetValue(v int64) {
	m.Value = v
}

func (m *ResourceOptions_StringsListOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_StringsListOption) SetValue(v []string) {
	m.Value = v
}

func (m *ResourceOptions_StringsMapOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_StringsMapOption) SetValue(v map[string]string) {
	m.Value = v
}

func (m *ResourceOptions_CachingTimes) SetSimpleValue(v int64) {
	m.SimpleValue = v
}

func (m *ResourceOptions_CachingTimes) SetCustomValues(v map[string]int64) {
	m.CustomValues = v
}

type ResourceOptions_EdgeCacheSettings_ValuesVariant = isResourceOptions_EdgeCacheSettings_ValuesVariant

func (m *ResourceOptions_EdgeCacheSettings) SetValuesVariant(v ResourceOptions_EdgeCacheSettings_ValuesVariant) {
	m.ValuesVariant = v
}

func (m *ResourceOptions_EdgeCacheSettings) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_EdgeCacheSettings) SetValue(v *ResourceOptions_CachingTimes) {
	m.ValuesVariant = &ResourceOptions_EdgeCacheSettings_Value{
		Value: v,
	}
}

func (m *ResourceOptions_EdgeCacheSettings) SetDefaultValue(v int64) {
	m.ValuesVariant = &ResourceOptions_EdgeCacheSettings_DefaultValue{
		DefaultValue: v,
	}
}

func (m *ResourceOptions_StringVariableMapOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_StringVariableMapOption) SetValue(v map[string]*ResourceOptions_StringVariableMapOption_OneofString) {
	m.Value = v
}

type ResourceOptions_StringVariableMapOption_OneofString_StringOption = isResourceOptions_StringVariableMapOption_OneofString_StringOption

func (m *ResourceOptions_StringVariableMapOption_OneofString) SetStringOption(v ResourceOptions_StringVariableMapOption_OneofString_StringOption) {
	m.StringOption = v
}

func (m *ResourceOptions_StringVariableMapOption_OneofString) SetValue(v *ResourceOptions_StringOption) {
	m.StringOption = &ResourceOptions_StringVariableMapOption_OneofString_Value{
		Value: v,
	}
}

func (m *ResourceOptions_StringVariableMapOption_OneofString) SetValues(v *ResourceOptions_StringsListOption) {
	m.StringOption = &ResourceOptions_StringVariableMapOption_OneofString_Values{
		Values: v,
	}
}

type ResourceOptions_QueryParamsOptions_QueryParamsVariant = isResourceOptions_QueryParamsOptions_QueryParamsVariant

func (m *ResourceOptions_QueryParamsOptions) SetQueryParamsVariant(v ResourceOptions_QueryParamsOptions_QueryParamsVariant) {
	m.QueryParamsVariant = v
}

func (m *ResourceOptions_QueryParamsOptions) SetIgnoreQueryString(v *ResourceOptions_BoolOption) {
	m.QueryParamsVariant = &ResourceOptions_QueryParamsOptions_IgnoreQueryString{
		IgnoreQueryString: v,
	}
}

func (m *ResourceOptions_QueryParamsOptions) SetQueryParamsWhitelist(v *ResourceOptions_StringsListOption) {
	m.QueryParamsVariant = &ResourceOptions_QueryParamsOptions_QueryParamsWhitelist{
		QueryParamsWhitelist: v,
	}
}

func (m *ResourceOptions_QueryParamsOptions) SetQueryParamsBlacklist(v *ResourceOptions_StringsListOption) {
	m.QueryParamsVariant = &ResourceOptions_QueryParamsOptions_QueryParamsBlacklist{
		QueryParamsBlacklist: v,
	}
}

type ResourceOptions_RedirectOptions_RedirectVariant = isResourceOptions_RedirectOptions_RedirectVariant

func (m *ResourceOptions_RedirectOptions) SetRedirectVariant(v ResourceOptions_RedirectOptions_RedirectVariant) {
	m.RedirectVariant = v
}

func (m *ResourceOptions_RedirectOptions) SetRedirectHttpToHttps(v *ResourceOptions_BoolOption) {
	m.RedirectVariant = &ResourceOptions_RedirectOptions_RedirectHttpToHttps{
		RedirectHttpToHttps: v,
	}
}

func (m *ResourceOptions_RedirectOptions) SetRedirectHttpsToHttp(v *ResourceOptions_BoolOption) {
	m.RedirectVariant = &ResourceOptions_RedirectOptions_RedirectHttpsToHttp{
		RedirectHttpsToHttp: v,
	}
}

type ResourceOptions_HostOptions_HostVariant = isResourceOptions_HostOptions_HostVariant

func (m *ResourceOptions_HostOptions) SetHostVariant(v ResourceOptions_HostOptions_HostVariant) {
	m.HostVariant = v
}

func (m *ResourceOptions_HostOptions) SetHost(v *ResourceOptions_StringOption) {
	m.HostVariant = &ResourceOptions_HostOptions_Host{
		Host: v,
	}
}

func (m *ResourceOptions_HostOptions) SetForwardHostHeader(v *ResourceOptions_BoolOption) {
	m.HostVariant = &ResourceOptions_HostOptions_ForwardHostHeader{
		ForwardHostHeader: v,
	}
}

type ResourceOptions_CompressionOptions_CompressionVariant = isResourceOptions_CompressionOptions_CompressionVariant

func (m *ResourceOptions_CompressionOptions) SetCompressionVariant(v ResourceOptions_CompressionOptions_CompressionVariant) {
	m.CompressionVariant = v
}

func (m *ResourceOptions_CompressionOptions) SetFetchCompressed(v *ResourceOptions_BoolOption) {
	m.CompressionVariant = &ResourceOptions_CompressionOptions_FetchCompressed{
		FetchCompressed: v,
	}
}

func (m *ResourceOptions_CompressionOptions) SetGzipOn(v *ResourceOptions_BoolOption) {
	m.CompressionVariant = &ResourceOptions_CompressionOptions_GzipOn{
		GzipOn: v,
	}
}

func (m *ResourceOptions_CompressionOptions) SetBrotliCompression(v *ResourceOptions_StringsListOption) {
	m.CompressionVariant = &ResourceOptions_CompressionOptions_BrotliCompression{
		BrotliCompression: v,
	}
}

func (m *ResourceOptions_RewriteOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_RewriteOption) SetBody(v string) {
	m.Body = v
}

func (m *ResourceOptions_RewriteOption) SetFlag(v RewriteFlag) {
	m.Flag = v
}

func (m *ResourceOptions_SecureKeyOption) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_SecureKeyOption) SetKey(v string) {
	m.Key = v
}

func (m *ResourceOptions_SecureKeyOption) SetType(v SecureKeyURLType) {
	m.Type = v
}

func (m *ResourceOptions_IPAddresACL) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *ResourceOptions_IPAddresACL) SetPolicyType(v PolicyType) {
	m.PolicyType = v
}

func (m *ResourceOptions_IPAddresACL) SetExceptedValues(v *ResourceOptions_StringsListOption) {
	m.ExceptedValues = v
}

func (m *SSLTargetCertificate) SetType(v SSLCertificateType) {
	m.Type = v
}

func (m *SSLTargetCertificate) SetData(v *SSLCertificateData) {
	m.Data = v
}

func (m *SSLCertificate) SetType(v SSLCertificateType) {
	m.Type = v
}

func (m *SSLCertificate) SetStatus(v SSLCertificateStatus) {
	m.Status = v
}

func (m *SSLCertificate) SetData(v *SSLCertificateData) {
	m.Data = v
}

type SSLCertificateData_SslCertificateDataVariant = isSSLCertificateData_SslCertificateDataVariant

func (m *SSLCertificateData) SetSslCertificateDataVariant(v SSLCertificateData_SslCertificateDataVariant) {
	m.SslCertificateDataVariant = v
}

func (m *SSLCertificateData) SetCm(v *SSLCertificateCMData) {
	m.SslCertificateDataVariant = &SSLCertificateData_Cm{
		Cm: v,
	}
}

func (m *SSLCertificateCMData) SetId(v string) {
	m.Id = v
}
