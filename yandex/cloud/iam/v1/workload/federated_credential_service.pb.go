// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/iam/v1/workload/federated_credential_service.proto

package workload

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFederatedCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federated credential to return.
	// To get the federated credential ID, make a [FederatedCredentialService.List] request.
	FederatedCredentialId string `protobuf:"bytes,1,opt,name=federated_credential_id,json=federatedCredentialId,proto3" json:"federated_credential_id,omitempty"`
}

func (x *GetFederatedCredentialRequest) Reset() {
	*x = GetFederatedCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFederatedCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFederatedCredentialRequest) ProtoMessage() {}

func (x *GetFederatedCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFederatedCredentialRequest.ProtoReflect.Descriptor instead.
func (*GetFederatedCredentialRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetFederatedCredentialRequest) GetFederatedCredentialId() string {
	if x != nil {
		return x.FederatedCredentialId
	}
	return ""
}

type ListFederatedCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the service account to list federated credentials for.
	// To get the the service account ID make a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListFederatedCredentialsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	// Default value: 100.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token]
	// to the [ListFederatedCredentialsResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListFederatedCredentialsRequest) Reset() {
	*x = ListFederatedCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederatedCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederatedCredentialsRequest) ProtoMessage() {}

func (x *ListFederatedCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederatedCredentialsRequest.ProtoReflect.Descriptor instead.
func (*ListFederatedCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListFederatedCredentialsRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *ListFederatedCredentialsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListFederatedCredentialsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListFederatedCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of federated credentials.
	FederatedCredentials []*FederatedCredential `protobuf:"bytes,1,rep,name=federated_credentials,json=federatedCredentials,proto3" json:"federated_credentials,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListFederatedCredentialsRequest.page_size], use
	// the [next_page_token] as the value
	// for the [ListFederatedCredentialsRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListFederatedCredentialsResponse) Reset() {
	*x = ListFederatedCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFederatedCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFederatedCredentialsResponse) ProtoMessage() {}

func (x *ListFederatedCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFederatedCredentialsResponse.ProtoReflect.Descriptor instead.
func (*ListFederatedCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListFederatedCredentialsResponse) GetFederatedCredentials() []*FederatedCredential {
	if x != nil {
		return x.FederatedCredentials
	}
	return nil
}

func (x *ListFederatedCredentialsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateFederatedCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the service account to create a federated credential for.
	// To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	ServiceAccountId string `protobuf:"bytes,1,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// ID of the workload identity federation that is used for authentication.
	FederationId string `protobuf:"bytes,2,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Id of the external subject.
	ExternalSubjectId string `protobuf:"bytes,3,opt,name=external_subject_id,json=externalSubjectId,proto3" json:"external_subject_id,omitempty"`
}

func (x *CreateFederatedCredentialRequest) Reset() {
	*x = CreateFederatedCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFederatedCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFederatedCredentialRequest) ProtoMessage() {}

func (x *CreateFederatedCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFederatedCredentialRequest.ProtoReflect.Descriptor instead.
func (*CreateFederatedCredentialRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateFederatedCredentialRequest) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CreateFederatedCredentialRequest) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

func (x *CreateFederatedCredentialRequest) GetExternalSubjectId() string {
	if x != nil {
		return x.ExternalSubjectId
	}
	return ""
}

type CreateFederatedCredentialMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federated credential that is being created.
	FederatedCredentialId string `protobuf:"bytes,1,opt,name=federated_credential_id,json=federatedCredentialId,proto3" json:"federated_credential_id,omitempty"`
}

func (x *CreateFederatedCredentialMetadata) Reset() {
	*x = CreateFederatedCredentialMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFederatedCredentialMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFederatedCredentialMetadata) ProtoMessage() {}

func (x *CreateFederatedCredentialMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFederatedCredentialMetadata.ProtoReflect.Descriptor instead.
func (*CreateFederatedCredentialMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateFederatedCredentialMetadata) GetFederatedCredentialId() string {
	if x != nil {
		return x.FederatedCredentialId
	}
	return ""
}

type DeleteFederatedCredentialRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federated credential key to delete.
	// To get the federated credential ID, use a [FederatedCredentialService.List] request.
	FederatedCredentialId string `protobuf:"bytes,1,opt,name=federated_credential_id,json=federatedCredentialId,proto3" json:"federated_credential_id,omitempty"`
}

func (x *DeleteFederatedCredentialRequest) Reset() {
	*x = DeleteFederatedCredentialRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFederatedCredentialRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFederatedCredentialRequest) ProtoMessage() {}

func (x *DeleteFederatedCredentialRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFederatedCredentialRequest.ProtoReflect.Descriptor instead.
func (*DeleteFederatedCredentialRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFederatedCredentialRequest) GetFederatedCredentialId() string {
	if x != nil {
		return x.FederatedCredentialId
	}
	return ""
}

type DeleteFederatedCredentialMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federated credential that is being deleted.
	FederatedCredentialId string `protobuf:"bytes,1,opt,name=federated_credential_id,json=federatedCredentialId,proto3" json:"federated_credential_id,omitempty"`
}

func (x *DeleteFederatedCredentialMetadata) Reset() {
	*x = DeleteFederatedCredentialMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFederatedCredentialMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFederatedCredentialMetadata) ProtoMessage() {}

func (x *DeleteFederatedCredentialMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFederatedCredentialMetadata.ProtoReflect.Descriptor instead.
func (*DeleteFederatedCredentialMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteFederatedCredentialMetadata) GetFederatedCredentialId() string {
	if x != nil {
		return x.FederatedCredentialId
	}
	return ""
}

var File_yandex_cloud_iam_v1_workload_federated_credential_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1c, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x37, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x65, 0x0a, 0x1d, 0x47, 0x65, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x44, 0x0a, 0x17, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x15, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xb1, 0x01, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x30,
	0x2d, 0x31, 0x30, 0x30, 0x30, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x29, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0x8a, 0xc8, 0x31, 0x06, 0x3c, 0x3d, 0x32, 0x30, 0x30, 0x30, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xb2, 0x01, 0x0a, 0x20, 0x4c,
	0x69, 0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x66, 0x0a, 0x15, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x52, 0x14, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0xcf, 0x01, 0x0a, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x3a, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x10,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64,
	0x12, 0x31, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x11,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x22, 0x5b, 0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a, 0x17, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x68,
	0x0a, 0x20, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x44, 0x0a, 0x17, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x52, 0x15, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x5b, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x36, 0x0a,
	0x17, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x49, 0x64, 0x32, 0xe6, 0x06, 0x0a, 0x1a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0xbe, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x3b, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x22, 0x47, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x41, 0x12, 0x3f, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xb4, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3d,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x12, 0x25, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77,
	0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0xd9, 0x01, 0x0a,
	0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x3e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f,
	0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6c, 0xb2, 0xd2, 0x2a, 0x38,
	0x0a, 0x21, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x13, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x3a, 0x01,
	0x2a, 0x22, 0x25, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c,
	0x6f, 0x61, 0x64, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0xf3, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x3e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x85, 0x01, 0xb2, 0xd2, 0x2a, 0x3a, 0x0a, 0x21, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x41, 0x2a, 0x3f, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x61, 0x6c, 0x73, 0x2f, 0x7b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x6d,
	0x0a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b,
	0x6c, 0x6f, 0x61, 0x64, 0x3b, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescData = file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDesc
)

func file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescData)
	})
	return file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDescData
}

var file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_goTypes = []any{
	(*GetFederatedCredentialRequest)(nil),     // 0: yandex.cloud.iam.v1.workload.GetFederatedCredentialRequest
	(*ListFederatedCredentialsRequest)(nil),   // 1: yandex.cloud.iam.v1.workload.ListFederatedCredentialsRequest
	(*ListFederatedCredentialsResponse)(nil),  // 2: yandex.cloud.iam.v1.workload.ListFederatedCredentialsResponse
	(*CreateFederatedCredentialRequest)(nil),  // 3: yandex.cloud.iam.v1.workload.CreateFederatedCredentialRequest
	(*CreateFederatedCredentialMetadata)(nil), // 4: yandex.cloud.iam.v1.workload.CreateFederatedCredentialMetadata
	(*DeleteFederatedCredentialRequest)(nil),  // 5: yandex.cloud.iam.v1.workload.DeleteFederatedCredentialRequest
	(*DeleteFederatedCredentialMetadata)(nil), // 6: yandex.cloud.iam.v1.workload.DeleteFederatedCredentialMetadata
	(*FederatedCredential)(nil),               // 7: yandex.cloud.iam.v1.workload.FederatedCredential
	(*operation.Operation)(nil),               // 8: yandex.cloud.operation.Operation
}
var file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_depIdxs = []int32{
	7, // 0: yandex.cloud.iam.v1.workload.ListFederatedCredentialsResponse.federated_credentials:type_name -> yandex.cloud.iam.v1.workload.FederatedCredential
	0, // 1: yandex.cloud.iam.v1.workload.FederatedCredentialService.Get:input_type -> yandex.cloud.iam.v1.workload.GetFederatedCredentialRequest
	1, // 2: yandex.cloud.iam.v1.workload.FederatedCredentialService.List:input_type -> yandex.cloud.iam.v1.workload.ListFederatedCredentialsRequest
	3, // 3: yandex.cloud.iam.v1.workload.FederatedCredentialService.Create:input_type -> yandex.cloud.iam.v1.workload.CreateFederatedCredentialRequest
	5, // 4: yandex.cloud.iam.v1.workload.FederatedCredentialService.Delete:input_type -> yandex.cloud.iam.v1.workload.DeleteFederatedCredentialRequest
	7, // 5: yandex.cloud.iam.v1.workload.FederatedCredentialService.Get:output_type -> yandex.cloud.iam.v1.workload.FederatedCredential
	2, // 6: yandex.cloud.iam.v1.workload.FederatedCredentialService.List:output_type -> yandex.cloud.iam.v1.workload.ListFederatedCredentialsResponse
	8, // 7: yandex.cloud.iam.v1.workload.FederatedCredentialService.Create:output_type -> yandex.cloud.operation.Operation
	8, // 8: yandex.cloud.iam.v1.workload.FederatedCredentialService.Delete:output_type -> yandex.cloud.operation.Operation
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_init() }
func file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_init() {
	if File_yandex_cloud_iam_v1_workload_federated_credential_service_proto != nil {
		return
	}
	file_yandex_cloud_iam_v1_workload_federated_credential_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetFederatedCredentialRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListFederatedCredentialsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListFederatedCredentialsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFederatedCredentialRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CreateFederatedCredentialMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFederatedCredentialRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteFederatedCredentialMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_workload_federated_credential_service_proto = out.File
	file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_rawDesc = nil
	file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_goTypes = nil
	file_yandex_cloud_iam_v1_workload_federated_credential_service_proto_depIdxs = nil
}
