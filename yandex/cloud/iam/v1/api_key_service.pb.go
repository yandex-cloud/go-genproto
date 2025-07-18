// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/iam/v1/api_key_service.proto

package iam

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetApiKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the API key to return.
	// To get the API key ID, use a [ApiKeyService.List] request.
	ApiKeyId      string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetApiKeyRequest) Reset() {
	*x = GetApiKeyRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetApiKeyRequest) ProtoMessage() {}

func (x *GetApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetApiKeyRequest.ProtoReflect.Descriptor instead.
func (*GetApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetApiKeyRequest) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

type ListApiKeysRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the service account to list API keys for.
	// To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	// If not specified, it defaults to the subject that made the request.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"` // use current subject identity if this not set
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListApiKeysResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token]
	// to the [ListApiKeysResponse.next_page_token]
	// returned by a previous list request.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApiKeysRequest) Reset() {
	*x = ListApiKeysRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApiKeysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeysRequest) ProtoMessage() {}

func (x *ListApiKeysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeysRequest.ProtoReflect.Descriptor instead.
func (*ListApiKeysRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListApiKeysRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *ListApiKeysRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListApiKeysRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListApiKeysResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of API keys.
	ApiKeys []*ApiKey `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListApiKeysRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListApiKeysRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApiKeysResponse) Reset() {
	*x = ListApiKeysResponse{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApiKeysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeysResponse) ProtoMessage() {}

func (x *ListApiKeysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeysResponse.ProtoReflect.Descriptor instead.
func (*ListApiKeysResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListApiKeysResponse) GetApiKeys() []*ApiKey {
	if x != nil {
		return x.ApiKeys
	}
	return nil
}

func (x *ListApiKeysResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateApiKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the service account to create an API key for.
	// To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	// If not specified, it defaults to the subject that made the request.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"` // use current subject identity if this not set
	// Description of the API key.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// Draft
	// Scope of the API key.
	//
	// Deprecated: Marked as deprecated in yandex/cloud/iam/v1/api_key_service.proto.
	Scope string `protobuf:"bytes,3,opt,name=scope,proto3" json:"scope,omitempty"`
	// Scopes of the API key.
	Scopes []string `protobuf:"bytes,5,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// API key expiration timestamp, if not specified, then the API key doesn't expire
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateApiKeyRequest) Reset() {
	*x = CreateApiKeyRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiKeyRequest) ProtoMessage() {}

func (x *CreateApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiKeyRequest.ProtoReflect.Descriptor instead.
func (*CreateApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateApiKeyRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CreateApiKeyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Deprecated: Marked as deprecated in yandex/cloud/iam/v1/api_key_service.proto.
func (x *CreateApiKeyRequest) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *CreateApiKeyRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *CreateApiKeyRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type CreateApiKeyResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ApiKey resource.
	ApiKey *ApiKey `protobuf:"bytes,1,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// Secret part of the API key. This secret key you may use in the requests for authentication.
	Secret        string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateApiKeyResponse) Reset() {
	*x = CreateApiKeyResponse{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateApiKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateApiKeyResponse) ProtoMessage() {}

func (x *CreateApiKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateApiKeyResponse.ProtoReflect.Descriptor instead.
func (*CreateApiKeyResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateApiKeyResponse) GetApiKey() *ApiKey {
	if x != nil {
		return x.ApiKey
	}
	return nil
}

func (x *CreateApiKeyResponse) GetSecret() string {
	if x != nil {
		return x.Secret
	}
	return ""
}

type UpdateApiKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the ApiKey resource to update.
	// To get the API key ID, use a [ApiKeyService.List] request.
	ApiKeyId string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	// Field mask that specifies which fields of the ApiKey resource are going to be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Description of the API key.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Scopes of the API key.
	Scopes []string `protobuf:"bytes,4,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// API key expiration timestamp, if not specified, then the API key doesn't expire
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateApiKeyRequest) Reset() {
	*x = UpdateApiKeyRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiKeyRequest) ProtoMessage() {}

func (x *UpdateApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiKeyRequest.ProtoReflect.Descriptor instead.
func (*UpdateApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateApiKeyRequest) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *UpdateApiKeyRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateApiKeyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateApiKeyRequest) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *UpdateApiKeyRequest) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

type UpdateApiKeyMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the ApiKey resource that is being updated.
	ApiKeyId      string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateApiKeyMetadata) Reset() {
	*x = UpdateApiKeyMetadata{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateApiKeyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateApiKeyMetadata) ProtoMessage() {}

func (x *UpdateApiKeyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateApiKeyMetadata.ProtoReflect.Descriptor instead.
func (*UpdateApiKeyMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateApiKeyMetadata) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

type DeleteApiKeyRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the API key to delete.
	// To get the API key ID, use a [ApiKeyService.List] request.
	ApiKeyId      string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteApiKeyRequest) Reset() {
	*x = DeleteApiKeyRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteApiKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiKeyRequest) ProtoMessage() {}

func (x *DeleteApiKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiKeyRequest.ProtoReflect.Descriptor instead.
func (*DeleteApiKeyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteApiKeyRequest) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

type DeleteApiKeyMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the API key that is being deleted.
	ApiKeyId      string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteApiKeyMetadata) Reset() {
	*x = DeleteApiKeyMetadata{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteApiKeyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteApiKeyMetadata) ProtoMessage() {}

func (x *DeleteApiKeyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteApiKeyMetadata.ProtoReflect.Descriptor instead.
func (*DeleteApiKeyMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteApiKeyMetadata) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

type ListApiKeyOperationsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the key to list operations for.
	ApiKeyId string `protobuf:"bytes,1,opt,name=api_key_id,json=apiKeyId,proto3" json:"api_key_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListApiKeyOperationsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListApiKeyOperationsResponse.next_page_token] returned by a previous list request.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApiKeyOperationsRequest) Reset() {
	*x = ListApiKeyOperationsRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApiKeyOperationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeyOperationsRequest) ProtoMessage() {}

func (x *ListApiKeyOperationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeyOperationsRequest.ProtoReflect.Descriptor instead.
func (*ListApiKeyOperationsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListApiKeyOperationsRequest) GetApiKeyId() string {
	if x != nil {
		return x.ApiKeyId
	}
	return ""
}

func (x *ListApiKeyOperationsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListApiKeyOperationsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListApiKeyOperationsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of operations for the specified API key.
	Operations []*operation.Operation `protobuf:"bytes,1,rep,name=operations,proto3" json:"operations,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListApiKeyOperationsRequest.page_size], use the [next_page_token] as the value
	// for the [ListApiKeyOperationsRequest.page_token] query parameter in the next list request.
	// Each subsequent list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApiKeyOperationsResponse) Reset() {
	*x = ListApiKeyOperationsResponse{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApiKeyOperationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeyOperationsResponse) ProtoMessage() {}

func (x *ListApiKeyOperationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeyOperationsResponse.ProtoReflect.Descriptor instead.
func (*ListApiKeyOperationsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{10}
}

func (x *ListApiKeyOperationsResponse) GetOperations() []*operation.Operation {
	if x != nil {
		return x.Operations
	}
	return nil
}

func (x *ListApiKeyOperationsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type ListApiKeyScopesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListApiKeyScopesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the
	// [ListApiKeyScopesResponse.next_page_token] returned by a previous list request.
	PageToken     string `protobuf:"bytes,2,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApiKeyScopesRequest) Reset() {
	*x = ListApiKeyScopesRequest{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApiKeyScopesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeyScopesRequest) ProtoMessage() {}

func (x *ListApiKeyScopesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeyScopesRequest.ProtoReflect.Descriptor instead.
func (*ListApiKeyScopesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{11}
}

func (x *ListApiKeyScopesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListApiKeyScopesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListApiKeyScopesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of scopes
	Scopes []string `protobuf:"bytes,1,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListApiKeyScopesRequest.page_size], use the [next_page_token] as the value
	// for the [ListApiKeyScopesRequest.page_token] query parameter in the next list request.
	// Each subsequent list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListApiKeyScopesResponse) Reset() {
	*x = ListApiKeyScopesResponse{}
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListApiKeyScopesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListApiKeyScopesResponse) ProtoMessage() {}

func (x *ListApiKeyScopesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListApiKeyScopesResponse.ProtoReflect.Descriptor instead.
func (*ListApiKeyScopesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP(), []int{12}
}

func (x *ListApiKeyScopesResponse) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *ListApiKeyScopesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_iam_v1_api_key_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_iam_v1_api_key_service_proto_rawDesc = "" +
	"\n" +
	")yandex/cloud/iam/v1/api_key_service.proto\x12\x13yandex.cloud.iam.v1\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a yandex/cloud/api/operation.proto\x1a!yandex/cloud/iam/v1/api_key.proto\x1a&yandex/cloud/operation/operation.proto\x1a\x1dyandex/cloud/validation.proto\">\n" +
	"\x10GetApiKeyRequest\x12*\n" +
	"\n" +
	"api_key_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\bapiKeyId\"\xa0\x01\n" +
	"\x12ListApiKeysRequest\x126\n" +
	"\x12service_account_id\x18\x01 \x01(\tB\b\x8a\xc81\x04<=50R\x10serviceAccountId\x12'\n" +
	"\tpage_size\x18\x02 \x01(\x03B\n" +
	"\xfa\xc71\x060-1000R\bpageSize\x12)\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\n" +
	"\x8a\xc81\x06<=2000R\tpageToken\"u\n" +
	"\x13ListApiKeysResponse\x126\n" +
	"\bapi_keys\x18\x01 \x03(\v2\x1b.yandex.cloud.iam.v1.ApiKeyR\aapiKeys\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\x84\x02\n" +
	"\x13CreateApiKeyRequest\x126\n" +
	"\x12service_account_id\x18\x01 \x01(\tB\b\x8a\xc81\x04<=50R\x10serviceAccountId\x12+\n" +
	"\vdescription\x18\x02 \x01(\tB\t\x8a\xc81\x05<=256R\vdescription\x12!\n" +
	"\x05scope\x18\x03 \x01(\tB\v\x8a\xc81\x05<=256\x18\x01R\x05scope\x12*\n" +
	"\x06scopes\x18\x05 \x03(\tB\x12\x82\xc81\x05<=100\x8a\xc81\x05<=256R\x06scopes\x129\n" +
	"\n" +
	"expires_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\"d\n" +
	"\x14CreateApiKeyResponse\x124\n" +
	"\aapi_key\x18\x01 \x01(\v2\x1b.yandex.cloud.iam.v1.ApiKeyR\x06apiKey\x12\x16\n" +
	"\x06secret\x18\x02 \x01(\tR\x06secret\"\x92\x02\n" +
	"\x13UpdateApiKeyRequest\x12*\n" +
	"\n" +
	"api_key_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\bapiKeyId\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12+\n" +
	"\vdescription\x18\x03 \x01(\tB\t\x8a\xc81\x05<=256R\vdescription\x12*\n" +
	"\x06scopes\x18\x04 \x03(\tB\x12\x82\xc81\x051-100\x8a\xc81\x05<=256R\x06scopes\x129\n" +
	"\n" +
	"expires_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAt\"4\n" +
	"\x14UpdateApiKeyMetadata\x12\x1c\n" +
	"\n" +
	"api_key_id\x18\x01 \x01(\tR\bapiKeyId\"A\n" +
	"\x13DeleteApiKeyRequest\x12*\n" +
	"\n" +
	"api_key_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\bapiKeyId\"4\n" +
	"\x14DeleteApiKeyMetadata\x12\x1c\n" +
	"\n" +
	"api_key_id\x18\x01 \x01(\tR\bapiKeyId\"\x9d\x01\n" +
	"\x1bListApiKeyOperationsRequest\x12*\n" +
	"\n" +
	"api_key_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\bapiKeyId\x12'\n" +
	"\tpage_size\x18\x02 \x01(\x03B\n" +
	"\xfa\xc71\x060-1000R\bpageSize\x12)\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\n" +
	"\x8a\xc81\x06<=2000R\tpageToken\"\x89\x01\n" +
	"\x1cListApiKeyOperationsResponse\x12A\n" +
	"\n" +
	"operations\x18\x01 \x03(\v2!.yandex.cloud.operation.OperationR\n" +
	"operations\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"m\n" +
	"\x17ListApiKeyScopesRequest\x12'\n" +
	"\tpage_size\x18\x01 \x01(\x03B\n" +
	"\xfa\xc71\x060-1000R\bpageSize\x12)\n" +
	"\n" +
	"page_token\x18\x02 \x01(\tB\n" +
	"\x8a\xc81\x06<=2000R\tpageToken\"Z\n" +
	"\x18ListApiKeyScopesResponse\x12\x16\n" +
	"\x06scopes\x18\x01 \x03(\tR\x06scopes\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken2\xf6\a\n" +
	"\rApiKeyService\x12r\n" +
	"\x04List\x12'.yandex.cloud.iam.v1.ListApiKeysRequest\x1a(.yandex.cloud.iam.v1.ListApiKeysResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/iam/v1/apiKeys\x12o\n" +
	"\x03Get\x12%.yandex.cloud.iam.v1.GetApiKeyRequest\x1a\x1b.yandex.cloud.iam.v1.ApiKey\"$\x82\xd3\xe4\x93\x02\x1e\x12\x1c/iam/v1/apiKeys/{api_key_id}\x12y\n" +
	"\x06Create\x12(.yandex.cloud.iam.v1.CreateApiKeyRequest\x1a).yandex.cloud.iam.v1.CreateApiKeyResponse\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/iam/v1/apiKeys\x12\xa0\x01\n" +
	"\x06Update\x12(.yandex.cloud.iam.v1.UpdateApiKeyRequest\x1a!.yandex.cloud.operation.Operation\"I\xb2\xd2*\x1e\n" +
	"\x14UpdateApiKeyMetadata\x12\x06ApiKey\x82\xd3\xe4\x93\x02!:\x01*2\x1c/iam/v1/apiKeys/{api_key_id}\x12\xac\x01\n" +
	"\x06Delete\x12(.yandex.cloud.iam.v1.DeleteApiKeyRequest\x1a!.yandex.cloud.operation.Operation\"U\xb2\xd2*-\n" +
	"\x14DeleteApiKeyMetadata\x12\x15google.protobuf.Empty\x82\xd3\xe4\x93\x02\x1e*\x1c/iam/v1/apiKeys/{api_key_id}\x12\xa6\x01\n" +
	"\x0eListOperations\x120.yandex.cloud.iam.v1.ListApiKeyOperationsRequest\x1a1.yandex.cloud.iam.v1.ListApiKeyOperationsResponse\"/\x82\xd3\xe4\x93\x02)\x12'/iam/v1/apiKeys/{api_key_id}/operations\x12\x89\x01\n" +
	"\n" +
	"ListScopes\x12,.yandex.cloud.iam.v1.ListApiKeyScopesRequest\x1a-.yandex.cloud.iam.v1.ListApiKeyScopesResponse\"\x1e\x82\xd3\xe4\x93\x02\x18\x12\x16/iam/v1/apiKeys/scopesBV\n" +
	"\x17yandex.cloud.api.iam.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1;iamb\x06proto3"

var (
	file_yandex_cloud_iam_v1_api_key_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_api_key_service_proto_rawDescData []byte
)

func file_yandex_cloud_iam_v1_api_key_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_api_key_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_api_key_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_iam_v1_api_key_service_proto_rawDesc), len(file_yandex_cloud_iam_v1_api_key_service_proto_rawDesc)))
	})
	return file_yandex_cloud_iam_v1_api_key_service_proto_rawDescData
}

var file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_yandex_cloud_iam_v1_api_key_service_proto_goTypes = []any{
	(*GetApiKeyRequest)(nil),             // 0: yandex.cloud.iam.v1.GetApiKeyRequest
	(*ListApiKeysRequest)(nil),           // 1: yandex.cloud.iam.v1.ListApiKeysRequest
	(*ListApiKeysResponse)(nil),          // 2: yandex.cloud.iam.v1.ListApiKeysResponse
	(*CreateApiKeyRequest)(nil),          // 3: yandex.cloud.iam.v1.CreateApiKeyRequest
	(*CreateApiKeyResponse)(nil),         // 4: yandex.cloud.iam.v1.CreateApiKeyResponse
	(*UpdateApiKeyRequest)(nil),          // 5: yandex.cloud.iam.v1.UpdateApiKeyRequest
	(*UpdateApiKeyMetadata)(nil),         // 6: yandex.cloud.iam.v1.UpdateApiKeyMetadata
	(*DeleteApiKeyRequest)(nil),          // 7: yandex.cloud.iam.v1.DeleteApiKeyRequest
	(*DeleteApiKeyMetadata)(nil),         // 8: yandex.cloud.iam.v1.DeleteApiKeyMetadata
	(*ListApiKeyOperationsRequest)(nil),  // 9: yandex.cloud.iam.v1.ListApiKeyOperationsRequest
	(*ListApiKeyOperationsResponse)(nil), // 10: yandex.cloud.iam.v1.ListApiKeyOperationsResponse
	(*ListApiKeyScopesRequest)(nil),      // 11: yandex.cloud.iam.v1.ListApiKeyScopesRequest
	(*ListApiKeyScopesResponse)(nil),     // 12: yandex.cloud.iam.v1.ListApiKeyScopesResponse
	(*ApiKey)(nil),                       // 13: yandex.cloud.iam.v1.ApiKey
	(*timestamppb.Timestamp)(nil),        // 14: google.protobuf.Timestamp
	(*fieldmaskpb.FieldMask)(nil),        // 15: google.protobuf.FieldMask
	(*operation.Operation)(nil),          // 16: yandex.cloud.operation.Operation
}
var file_yandex_cloud_iam_v1_api_key_service_proto_depIdxs = []int32{
	13, // 0: yandex.cloud.iam.v1.ListApiKeysResponse.api_keys:type_name -> yandex.cloud.iam.v1.ApiKey
	14, // 1: yandex.cloud.iam.v1.CreateApiKeyRequest.expires_at:type_name -> google.protobuf.Timestamp
	13, // 2: yandex.cloud.iam.v1.CreateApiKeyResponse.api_key:type_name -> yandex.cloud.iam.v1.ApiKey
	15, // 3: yandex.cloud.iam.v1.UpdateApiKeyRequest.update_mask:type_name -> google.protobuf.FieldMask
	14, // 4: yandex.cloud.iam.v1.UpdateApiKeyRequest.expires_at:type_name -> google.protobuf.Timestamp
	16, // 5: yandex.cloud.iam.v1.ListApiKeyOperationsResponse.operations:type_name -> yandex.cloud.operation.Operation
	1,  // 6: yandex.cloud.iam.v1.ApiKeyService.List:input_type -> yandex.cloud.iam.v1.ListApiKeysRequest
	0,  // 7: yandex.cloud.iam.v1.ApiKeyService.Get:input_type -> yandex.cloud.iam.v1.GetApiKeyRequest
	3,  // 8: yandex.cloud.iam.v1.ApiKeyService.Create:input_type -> yandex.cloud.iam.v1.CreateApiKeyRequest
	5,  // 9: yandex.cloud.iam.v1.ApiKeyService.Update:input_type -> yandex.cloud.iam.v1.UpdateApiKeyRequest
	7,  // 10: yandex.cloud.iam.v1.ApiKeyService.Delete:input_type -> yandex.cloud.iam.v1.DeleteApiKeyRequest
	9,  // 11: yandex.cloud.iam.v1.ApiKeyService.ListOperations:input_type -> yandex.cloud.iam.v1.ListApiKeyOperationsRequest
	11, // 12: yandex.cloud.iam.v1.ApiKeyService.ListScopes:input_type -> yandex.cloud.iam.v1.ListApiKeyScopesRequest
	2,  // 13: yandex.cloud.iam.v1.ApiKeyService.List:output_type -> yandex.cloud.iam.v1.ListApiKeysResponse
	13, // 14: yandex.cloud.iam.v1.ApiKeyService.Get:output_type -> yandex.cloud.iam.v1.ApiKey
	4,  // 15: yandex.cloud.iam.v1.ApiKeyService.Create:output_type -> yandex.cloud.iam.v1.CreateApiKeyResponse
	16, // 16: yandex.cloud.iam.v1.ApiKeyService.Update:output_type -> yandex.cloud.operation.Operation
	16, // 17: yandex.cloud.iam.v1.ApiKeyService.Delete:output_type -> yandex.cloud.operation.Operation
	10, // 18: yandex.cloud.iam.v1.ApiKeyService.ListOperations:output_type -> yandex.cloud.iam.v1.ListApiKeyOperationsResponse
	12, // 19: yandex.cloud.iam.v1.ApiKeyService.ListScopes:output_type -> yandex.cloud.iam.v1.ListApiKeyScopesResponse
	13, // [13:20] is the sub-list for method output_type
	6,  // [6:13] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_api_key_service_proto_init() }
func file_yandex_cloud_iam_v1_api_key_service_proto_init() {
	if File_yandex_cloud_iam_v1_api_key_service_proto != nil {
		return
	}
	file_yandex_cloud_iam_v1_api_key_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_iam_v1_api_key_service_proto_rawDesc), len(file_yandex_cloud_iam_v1_api_key_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_iam_v1_api_key_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_api_key_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_iam_v1_api_key_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_api_key_service_proto = out.File
	file_yandex_cloud_iam_v1_api_key_service_proto_goTypes = nil
	file_yandex_cloud_iam_v1_api_key_service_proto_depIdxs = nil
}
