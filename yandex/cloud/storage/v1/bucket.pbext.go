// Code generated by protoc-gen-goext. DO NOT EDIT.

package storage

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Bucket) SetId(v string) {
	m.Id = v
}

func (m *Bucket) SetName(v string) {
	m.Name = v
}

func (m *Bucket) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Bucket) SetAnonymousAccessFlags(v *AnonymousAccessFlags) {
	m.AnonymousAccessFlags = v
}

func (m *Bucket) SetDefaultStorageClass(v string) {
	m.DefaultStorageClass = v
}

func (m *Bucket) SetVersioning(v Versioning) {
	m.Versioning = v
}

func (m *Bucket) SetMaxSize(v int64) {
	m.MaxSize = v
}

func (m *Bucket) SetPolicy(v *structpb.Struct) {
	m.Policy = v
}

func (m *Bucket) SetAcl(v *ACL) {
	m.Acl = v
}

func (m *Bucket) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Bucket) SetCors(v []*CorsRule) {
	m.Cors = v
}

func (m *Bucket) SetWebsiteSettings(v *WebsiteSettings) {
	m.WebsiteSettings = v
}

func (m *Bucket) SetLifecycleRules(v []*LifecycleRule) {
	m.LifecycleRules = v
}

func (m *Bucket) SetTags(v []*Tag) {
	m.Tags = v
}

func (m *Bucket) SetObjectLock(v *ObjectLock) {
	m.ObjectLock = v
}

func (m *Bucket) SetEncryption(v *Encryption) {
	m.Encryption = v
}

func (m *Bucket) SetAllowedPrivateEndpoints(v *BucketAllowedPrivateEndpoints) {
	m.AllowedPrivateEndpoints = v
}

func (m *Tag) SetKey(v string) {
	m.Key = v
}

func (m *Tag) SetValue(v string) {
	m.Value = v
}

func (m *ACL) SetGrants(v []*ACL_Grant) {
	m.Grants = v
}

func (m *ACL_Grant) SetPermission(v ACL_Grant_Permission) {
	m.Permission = v
}

func (m *ACL_Grant) SetGrantType(v ACL_Grant_GrantType) {
	m.GrantType = v
}

func (m *ACL_Grant) SetGranteeId(v string) {
	m.GranteeId = v
}

func (m *AnonymousAccessFlags) SetRead(v *wrapperspb.BoolValue) {
	m.Read = v
}

func (m *AnonymousAccessFlags) SetList(v *wrapperspb.BoolValue) {
	m.List = v
}

func (m *AnonymousAccessFlags) SetConfigRead(v *wrapperspb.BoolValue) {
	m.ConfigRead = v
}

func (m *CorsRule) SetId(v string) {
	m.Id = v
}

func (m *CorsRule) SetAllowedMethods(v []CorsRule_Method) {
	m.AllowedMethods = v
}

func (m *CorsRule) SetAllowedHeaders(v []string) {
	m.AllowedHeaders = v
}

func (m *CorsRule) SetAllowedOrigins(v []string) {
	m.AllowedOrigins = v
}

func (m *CorsRule) SetExposeHeaders(v []string) {
	m.ExposeHeaders = v
}

func (m *CorsRule) SetMaxAgeSeconds(v *wrapperspb.Int64Value) {
	m.MaxAgeSeconds = v
}

func (m *WebsiteSettings) SetIndex(v string) {
	m.Index = v
}

func (m *WebsiteSettings) SetError(v string) {
	m.Error = v
}

func (m *WebsiteSettings) SetRedirectAllRequests(v *WebsiteSettings_Scheme) {
	m.RedirectAllRequests = v
}

func (m *WebsiteSettings) SetRoutingRules(v []*WebsiteSettings_RoutingRule) {
	m.RoutingRules = v
}

func (m *WebsiteSettings_Scheme) SetProtocol(v WebsiteSettings_Protocol) {
	m.Protocol = v
}

func (m *WebsiteSettings_Scheme) SetHostname(v string) {
	m.Hostname = v
}

func (m *WebsiteSettings_Condition) SetHttpErrorCodeReturnedEquals(v string) {
	m.HttpErrorCodeReturnedEquals = v
}

func (m *WebsiteSettings_Condition) SetKeyPrefixEquals(v string) {
	m.KeyPrefixEquals = v
}

func (m *WebsiteSettings_Redirect) SetHostname(v string) {
	m.Hostname = v
}

func (m *WebsiteSettings_Redirect) SetHttpRedirectCode(v string) {
	m.HttpRedirectCode = v
}

func (m *WebsiteSettings_Redirect) SetProtocol(v WebsiteSettings_Protocol) {
	m.Protocol = v
}

func (m *WebsiteSettings_Redirect) SetReplaceKeyPrefixWith(v string) {
	m.ReplaceKeyPrefixWith = v
}

func (m *WebsiteSettings_Redirect) SetReplaceKeyWith(v string) {
	m.ReplaceKeyWith = v
}

func (m *WebsiteSettings_RoutingRule) SetCondition(v *WebsiteSettings_Condition) {
	m.Condition = v
}

func (m *WebsiteSettings_RoutingRule) SetRedirect(v *WebsiteSettings_Redirect) {
	m.Redirect = v
}

func (m *LifecycleRule) SetId(v *wrapperspb.StringValue) {
	m.Id = v
}

func (m *LifecycleRule) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *LifecycleRule) SetFilter(v *LifecycleRule_RuleFilter) {
	m.Filter = v
}

func (m *LifecycleRule) SetExpiration(v *LifecycleRule_Expiration) {
	m.Expiration = v
}

func (m *LifecycleRule) SetTransitions(v []*LifecycleRule_Transition) {
	m.Transitions = v
}

func (m *LifecycleRule) SetAbortIncompleteMultipartUpload(v *LifecycleRule_AfterDays) {
	m.AbortIncompleteMultipartUpload = v
}

func (m *LifecycleRule) SetNoncurrentExpiration(v *LifecycleRule_NoncurrentExpiration) {
	m.NoncurrentExpiration = v
}

func (m *LifecycleRule) SetNoncurrentTransitions(v []*LifecycleRule_NoncurrentTransition) {
	m.NoncurrentTransitions = v
}

func (m *LifecycleRule) SetNoncurrentDeleteMarkers(v *LifecycleRule_NoncurrentDeleteMarkers) {
	m.NoncurrentDeleteMarkers = v
}

func (m *LifecycleRule_AfterDays) SetDaysAfterExpiration(v *wrapperspb.Int64Value) {
	m.DaysAfterExpiration = v
}

func (m *LifecycleRule_NoncurrentDeleteMarkers) SetNoncurrentDays(v *wrapperspb.Int64Value) {
	m.NoncurrentDays = v
}

func (m *LifecycleRule_NoncurrentExpiration) SetNoncurrentDays(v *wrapperspb.Int64Value) {
	m.NoncurrentDays = v
}

func (m *LifecycleRule_NoncurrentTransition) SetNoncurrentDays(v *wrapperspb.Int64Value) {
	m.NoncurrentDays = v
}

func (m *LifecycleRule_NoncurrentTransition) SetStorageClass(v string) {
	m.StorageClass = v
}

func (m *LifecycleRule_Transition) SetDate(v *timestamppb.Timestamp) {
	m.Date = v
}

func (m *LifecycleRule_Transition) SetDays(v *wrapperspb.Int64Value) {
	m.Days = v
}

func (m *LifecycleRule_Transition) SetStorageClass(v string) {
	m.StorageClass = v
}

func (m *LifecycleRule_Expiration) SetDate(v *timestamppb.Timestamp) {
	m.Date = v
}

func (m *LifecycleRule_Expiration) SetDays(v *wrapperspb.Int64Value) {
	m.Days = v
}

func (m *LifecycleRule_Expiration) SetExpiredObjectDeleteMarker(v *wrapperspb.BoolValue) {
	m.ExpiredObjectDeleteMarker = v
}

func (m *LifecycleRule_RuleFilter) SetPrefix(v string) {
	m.Prefix = v
}

func (m *LifecycleRule_RuleFilter) SetObjectSizeGreaterThan(v *wrapperspb.Int64Value) {
	m.ObjectSizeGreaterThan = v
}

func (m *LifecycleRule_RuleFilter) SetObjectSizeLessThan(v *wrapperspb.Int64Value) {
	m.ObjectSizeLessThan = v
}

func (m *LifecycleRule_RuleFilter) SetTag(v *Tag) {
	m.Tag = v
}

func (m *LifecycleRule_RuleFilter) SetAndOperator(v *LifecycleRule_RuleFilter_And) {
	m.AndOperator = v
}

func (m *LifecycleRule_RuleFilter_And) SetPrefix(v string) {
	m.Prefix = v
}

func (m *LifecycleRule_RuleFilter_And) SetObjectSizeGreaterThan(v *wrapperspb.Int64Value) {
	m.ObjectSizeGreaterThan = v
}

func (m *LifecycleRule_RuleFilter_And) SetObjectSizeLessThan(v *wrapperspb.Int64Value) {
	m.ObjectSizeLessThan = v
}

func (m *LifecycleRule_RuleFilter_And) SetTag(v []*Tag) {
	m.Tag = v
}

func (m *Counters) SetSimpleObjectSize(v int64) {
	m.SimpleObjectSize = v
}

func (m *Counters) SetSimpleObjectCount(v int64) {
	m.SimpleObjectCount = v
}

func (m *Counters) SetObjectsPartsSize(v int64) {
	m.ObjectsPartsSize = v
}

func (m *Counters) SetObjectsPartsCount(v int64) {
	m.ObjectsPartsCount = v
}

func (m *Counters) SetMultipartObjectsSize(v int64) {
	m.MultipartObjectsSize = v
}

func (m *Counters) SetMultipartObjectsCount(v int64) {
	m.MultipartObjectsCount = v
}

func (m *Counters) SetActiveMultipartCount(v int64) {
	m.ActiveMultipartCount = v
}

func (m *OptionalSizeByClass) SetStorageClass(v string) {
	m.StorageClass = v
}

func (m *OptionalSizeByClass) SetClassSize(v *wrapperspb.Int64Value) {
	m.ClassSize = v
}

func (m *SizeByClass) SetStorageClass(v string) {
	m.StorageClass = v
}

func (m *SizeByClass) SetClassSize(v int64) {
	m.ClassSize = v
}

func (m *CountersByClass) SetStorageClass(v string) {
	m.StorageClass = v
}

func (m *CountersByClass) SetCounters(v *Counters) {
	m.Counters = v
}

func (m *BucketStats) SetName(v string) {
	m.Name = v
}

func (m *BucketStats) SetMaxSize(v *wrapperspb.Int64Value) {
	m.MaxSize = v
}

func (m *BucketStats) SetUsedSize(v int64) {
	m.UsedSize = v
}

func (m *BucketStats) SetStorageClassMaxSizes(v []*OptionalSizeByClass) {
	m.StorageClassMaxSizes = v
}

func (m *BucketStats) SetStorageClassUsedSizes(v []*SizeByClass) {
	m.StorageClassUsedSizes = v
}

func (m *BucketStats) SetStorageClassCounters(v []*CountersByClass) {
	m.StorageClassCounters = v
}

func (m *BucketStats) SetDefaultStorageClass(v *wrapperspb.StringValue) {
	m.DefaultStorageClass = v
}

func (m *BucketStats) SetAnonymousAccessFlags(v *AnonymousAccessFlags) {
	m.AnonymousAccessFlags = v
}

func (m *BucketStats) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *BucketStats) SetUpdatedAt(v *timestamppb.Timestamp) {
	m.UpdatedAt = v
}

func (m *HTTPSConfig) SetName(v string) {
	m.Name = v
}

func (m *HTTPSConfig) SetSourceType(v HTTPSConfig_SourceType) {
	m.SourceType = v
}

func (m *HTTPSConfig) SetIssuer(v *wrapperspb.StringValue) {
	m.Issuer = v
}

func (m *HTTPSConfig) SetSubject(v *wrapperspb.StringValue) {
	m.Subject = v
}

func (m *HTTPSConfig) SetDnsNames(v []string) {
	m.DnsNames = v
}

func (m *HTTPSConfig) SetNotBefore(v *timestamppb.Timestamp) {
	m.NotBefore = v
}

func (m *HTTPSConfig) SetNotAfter(v *timestamppb.Timestamp) {
	m.NotAfter = v
}

func (m *HTTPSConfig) SetCertificateId(v string) {
	m.CertificateId = v
}

func (m *ObjectLock) SetStatus(v ObjectLock_ObjectLockStatus) {
	m.Status = v
}

func (m *ObjectLock) SetDefaultRetention(v *ObjectLock_DefaultRetention) {
	m.DefaultRetention = v
}

type ObjectLock_DefaultRetention_Period = isObjectLock_DefaultRetention_Period

func (m *ObjectLock_DefaultRetention) SetPeriod(v ObjectLock_DefaultRetention_Period) {
	m.Period = v
}

func (m *ObjectLock_DefaultRetention) SetMode(v ObjectLock_DefaultRetention_Mode) {
	m.Mode = v
}

func (m *ObjectLock_DefaultRetention) SetDays(v int64) {
	m.Period = &ObjectLock_DefaultRetention_Days{
		Days: v,
	}
}

func (m *ObjectLock_DefaultRetention) SetYears(v int64) {
	m.Period = &ObjectLock_DefaultRetention_Years{
		Years: v,
	}
}

func (m *Encryption) SetRules(v []*Encryption_EncryptionRule) {
	m.Rules = v
}

func (m *Encryption_EncryptionRule) SetKmsMasterKeyId(v string) {
	m.KmsMasterKeyId = v
}

func (m *Encryption_EncryptionRule) SetSseAlgorithm(v string) {
	m.SseAlgorithm = v
}

func (m *BucketAllowedPrivateEndpoints) SetEnabled(v bool) {
	m.Enabled = v
}

func (m *BucketAllowedPrivateEndpoints) SetPrivateEndpoints(v []string) {
	m.PrivateEndpoints = v
}
