// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/assistant_service.proto

package assistants

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	common "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
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

// Request to create a new assistant.
type CreateAssistantRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	FolderId string                 `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the assistant.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the assistant.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Expiration configuration for the assistant.
	ExpirationConfig *common.ExpirationConfig `protobuf:"bytes,4,opt,name=expiration_config,json=expirationConfig,proto3" json:"expiration_config,omitempty"`
	// Set of key-value pairs to label the user.
	Labels map[string]string `protobuf:"bytes,5,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The [ID of the model](/docs/foundation-models/concepts/yandexgpt/models) to be used for completion generation.
	ModelUri string `protobuf:"bytes,6,opt,name=model_uri,json=modelUri,proto3" json:"model_uri,omitempty"`
	// Instructions or guidelines that the assistant should follow when generating responses or performing tasks.
	// These instructions can help guide the assistant's behavior and responses.
	Instruction string `protobuf:"bytes,7,opt,name=instruction,proto3" json:"instruction,omitempty"`
	// Configuration options for truncating the prompt when the token count exceeds a specified limit.
	PromptTruncationOptions *PromptTruncationOptions `protobuf:"bytes,8,opt,name=prompt_truncation_options,json=promptTruncationOptions,proto3" json:"prompt_truncation_options,omitempty"`
	// Configuration options for completion generation.
	CompletionOptions *CompletionOptions `protobuf:"bytes,9,opt,name=completion_options,json=completionOptions,proto3" json:"completion_options,omitempty"`
	// List of tools that the assistant can use to perform additional tasks.
	// One example is the SearchIndexTool, which is used for Retrieval-Augmented Generation (RAG).
	Tools []*Tool `protobuf:"bytes,10,rep,name=tools,proto3" json:"tools,omitempty"`
	// Specifies the format of the model's response.
	ResponseFormat *ResponseFormat `protobuf:"bytes,11,opt,name=response_format,json=responseFormat,proto3" json:"response_format,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *CreateAssistantRequest) Reset() {
	*x = CreateAssistantRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateAssistantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAssistantRequest) ProtoMessage() {}

func (x *CreateAssistantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAssistantRequest.ProtoReflect.Descriptor instead.
func (*CreateAssistantRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateAssistantRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateAssistantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateAssistantRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateAssistantRequest) GetExpirationConfig() *common.ExpirationConfig {
	if x != nil {
		return x.ExpirationConfig
	}
	return nil
}

func (x *CreateAssistantRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateAssistantRequest) GetModelUri() string {
	if x != nil {
		return x.ModelUri
	}
	return ""
}

func (x *CreateAssistantRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *CreateAssistantRequest) GetPromptTruncationOptions() *PromptTruncationOptions {
	if x != nil {
		return x.PromptTruncationOptions
	}
	return nil
}

func (x *CreateAssistantRequest) GetCompletionOptions() *CompletionOptions {
	if x != nil {
		return x.CompletionOptions
	}
	return nil
}

func (x *CreateAssistantRequest) GetTools() []*Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *CreateAssistantRequest) GetResponseFormat() *ResponseFormat {
	if x != nil {
		return x.ResponseFormat
	}
	return nil
}

// Request message for retrieving an assistant by ID.
type GetAssistantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the assistant to retrieve.
	AssistantId   string `protobuf:"bytes,1,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAssistantRequest) Reset() {
	*x = GetAssistantRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAssistantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssistantRequest) ProtoMessage() {}

func (x *GetAssistantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssistantRequest.ProtoReflect.Descriptor instead.
func (*GetAssistantRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssistantRequest) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

// Request message for updating an existing assistant.
type UpdateAssistantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the assistant to update.
	AssistantId string `protobuf:"bytes,1,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	// Field mask specifying which fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// New name for the assistant.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// New description for the assistant.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// New expiration configuration for the assistant.
	ExpirationConfig *common.ExpirationConfig `protobuf:"bytes,5,opt,name=expiration_config,json=expirationConfig,proto3" json:"expiration_config,omitempty"`
	// New set of labels for the assistant.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// The new [ID of the model](/docs/foundation-models/concepts/yandexgpt/models) to be used for completion generation.
	ModelUri string `protobuf:"bytes,7,opt,name=model_uri,json=modelUri,proto3" json:"model_uri,omitempty"`
	// New instructions or guidelines for the assistant to follow.
	Instruction string `protobuf:"bytes,8,opt,name=instruction,proto3" json:"instruction,omitempty"`
	// New configuration for truncating the prompt.
	PromptTruncationOptions *PromptTruncationOptions `protobuf:"bytes,9,opt,name=prompt_truncation_options,json=promptTruncationOptions,proto3" json:"prompt_truncation_options,omitempty"`
	// New configuration for completion generation.
	CompletionOptions *CompletionOptions `protobuf:"bytes,10,opt,name=completion_options,json=completionOptions,proto3" json:"completion_options,omitempty"`
	// New list of tools the assistant can use.
	Tools          []*Tool         `protobuf:"bytes,11,rep,name=tools,proto3" json:"tools,omitempty"`
	ResponseFormat *ResponseFormat `protobuf:"bytes,12,opt,name=response_format,json=responseFormat,proto3" json:"response_format,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *UpdateAssistantRequest) Reset() {
	*x = UpdateAssistantRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateAssistantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAssistantRequest) ProtoMessage() {}

func (x *UpdateAssistantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAssistantRequest.ProtoReflect.Descriptor instead.
func (*UpdateAssistantRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAssistantRequest) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

func (x *UpdateAssistantRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateAssistantRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAssistantRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateAssistantRequest) GetExpirationConfig() *common.ExpirationConfig {
	if x != nil {
		return x.ExpirationConfig
	}
	return nil
}

func (x *UpdateAssistantRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *UpdateAssistantRequest) GetModelUri() string {
	if x != nil {
		return x.ModelUri
	}
	return ""
}

func (x *UpdateAssistantRequest) GetInstruction() string {
	if x != nil {
		return x.Instruction
	}
	return ""
}

func (x *UpdateAssistantRequest) GetPromptTruncationOptions() *PromptTruncationOptions {
	if x != nil {
		return x.PromptTruncationOptions
	}
	return nil
}

func (x *UpdateAssistantRequest) GetCompletionOptions() *CompletionOptions {
	if x != nil {
		return x.CompletionOptions
	}
	return nil
}

func (x *UpdateAssistantRequest) GetTools() []*Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

func (x *UpdateAssistantRequest) GetResponseFormat() *ResponseFormat {
	if x != nil {
		return x.ResponseFormat
	}
	return nil
}

// Request message for deleting an assistant by ID.
type DeleteAssistantRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the assistant to delete.
	AssistantId   string `protobuf:"bytes,1,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAssistantRequest) Reset() {
	*x = DeleteAssistantRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAssistantRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssistantRequest) ProtoMessage() {}

func (x *DeleteAssistantRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssistantRequest.ProtoReflect.Descriptor instead.
func (*DeleteAssistantRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteAssistantRequest) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

// Response message for the delete operation.
type DeleteAssistantResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteAssistantResponse) Reset() {
	*x = DeleteAssistantResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteAssistantResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAssistantResponse) ProtoMessage() {}

func (x *DeleteAssistantResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAssistantResponse.ProtoReflect.Descriptor instead.
func (*DeleteAssistantResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{4}
}

// Request message for listing assistants in a specific folder.
type ListAssistantsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Folder ID from which to list assistants.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Maximum number of assistants to return per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to retrieve the next page of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssistantsRequest) Reset() {
	*x = ListAssistantsRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssistantsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssistantsRequest) ProtoMessage() {}

func (x *ListAssistantsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssistantsRequest.ProtoReflect.Descriptor instead.
func (*ListAssistantsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListAssistantsRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListAssistantsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAssistantsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for the list operation.
type ListAssistantsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of assistants in the specified folder.
	Assistants []*Assistant `protobuf:"bytes,1,rep,name=assistants,proto3" json:"assistants,omitempty"`
	// Token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssistantsResponse) Reset() {
	*x = ListAssistantsResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssistantsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssistantsResponse) ProtoMessage() {}

func (x *ListAssistantsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssistantsResponse.ProtoReflect.Descriptor instead.
func (*ListAssistantsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListAssistantsResponse) GetAssistants() []*Assistant {
	if x != nil {
		return x.Assistants
	}
	return nil
}

func (x *ListAssistantsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

// Request to list all versions of a specific assistant.
type ListAssistantVersionsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the assistant whose versions are to be listed.
	AssistantId string `protobuf:"bytes,1,opt,name=assistant_id,json=assistantId,proto3" json:"assistant_id,omitempty"`
	// Maximum number of versions to return per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to retrieve the next page of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssistantVersionsRequest) Reset() {
	*x = ListAssistantVersionsRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssistantVersionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssistantVersionsRequest) ProtoMessage() {}

func (x *ListAssistantVersionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssistantVersionsRequest.ProtoReflect.Descriptor instead.
func (*ListAssistantVersionsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{7}
}

func (x *ListAssistantVersionsRequest) GetAssistantId() string {
	if x != nil {
		return x.AssistantId
	}
	return ""
}

func (x *ListAssistantVersionsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAssistantVersionsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Represents a specific version of an assistant.
type AssistantVersion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the assistant version.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Mask specifying which fields were updated in this version.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Assistant configuration for this version.
	Assistant     *Assistant `protobuf:"bytes,3,opt,name=assistant,proto3" json:"assistant,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AssistantVersion) Reset() {
	*x = AssistantVersion{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AssistantVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AssistantVersion) ProtoMessage() {}

func (x *AssistantVersion) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AssistantVersion.ProtoReflect.Descriptor instead.
func (*AssistantVersion) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{8}
}

func (x *AssistantVersion) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AssistantVersion) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *AssistantVersion) GetAssistant() *Assistant {
	if x != nil {
		return x.Assistant
	}
	return nil
}

// Response message containing the list versions operation.
type ListAssistantVersionsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of assistant versions.
	Versions []*AssistantVersion `protobuf:"bytes,1,rep,name=versions,proto3" json:"versions,omitempty"`
	// Token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAssistantVersionsResponse) Reset() {
	*x = ListAssistantVersionsResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAssistantVersionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssistantVersionsResponse) ProtoMessage() {}

func (x *ListAssistantVersionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssistantVersionsResponse.ProtoReflect.Descriptor instead.
func (*ListAssistantVersionsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP(), []int{9}
}

func (x *ListAssistantVersionsResponse) GetVersions() []*AssistantVersion {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *ListAssistantVersionsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_ai_assistants_v1_assistant_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDesc = "" +
	"\n" +
	"5yandex/cloud/ai/assistants/v1/assistant_service.proto\x12\x1dyandex.cloud.ai.assistants.v1\x1a#yandex/cloud/ai/common/common.proto\x1a-yandex/cloud/ai/assistants/v1/assistant.proto\x1a*yandex/cloud/ai/assistants/v1/common.proto\x1a\x1dyandex/cloud/validation.proto\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\"\x8b\x06\n" +
	"\x16CreateAssistantRequest\x12!\n" +
	"\tfolder_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bfolderId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12U\n" +
	"\x11expiration_config\x18\x04 \x01(\v2(.yandex.cloud.ai.common.ExpirationConfigR\x10expirationConfig\x12Y\n" +
	"\x06labels\x18\x05 \x03(\v2A.yandex.cloud.ai.assistants.v1.CreateAssistantRequest.LabelsEntryR\x06labels\x12!\n" +
	"\tmodel_uri\x18\x06 \x01(\tB\x04\xe8\xc71\x01R\bmodelUri\x12 \n" +
	"\vinstruction\x18\a \x01(\tR\vinstruction\x12r\n" +
	"\x19prompt_truncation_options\x18\b \x01(\v26.yandex.cloud.ai.assistants.v1.PromptTruncationOptionsR\x17promptTruncationOptions\x12_\n" +
	"\x12completion_options\x18\t \x01(\v20.yandex.cloud.ai.assistants.v1.CompletionOptionsR\x11completionOptions\x129\n" +
	"\x05tools\x18\n" +
	" \x03(\v2#.yandex.cloud.ai.assistants.v1.ToolR\x05tools\x12V\n" +
	"\x0fresponse_format\x18\v \x01(\v2-.yandex.cloud.ai.assistants.v1.ResponseFormatR\x0eresponseFormat\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\">\n" +
	"\x13GetAssistantRequest\x12'\n" +
	"\fassistant_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\vassistantId\"\xce\x06\n" +
	"\x16UpdateAssistantRequest\x12'\n" +
	"\fassistant_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\vassistantId\x12A\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskB\x04\xe8\xc71\x01R\n" +
	"updateMask\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12U\n" +
	"\x11expiration_config\x18\x05 \x01(\v2(.yandex.cloud.ai.common.ExpirationConfigR\x10expirationConfig\x12Y\n" +
	"\x06labels\x18\x06 \x03(\v2A.yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.LabelsEntryR\x06labels\x12\x1b\n" +
	"\tmodel_uri\x18\a \x01(\tR\bmodelUri\x12 \n" +
	"\vinstruction\x18\b \x01(\tR\vinstruction\x12r\n" +
	"\x19prompt_truncation_options\x18\t \x01(\v26.yandex.cloud.ai.assistants.v1.PromptTruncationOptionsR\x17promptTruncationOptions\x12_\n" +
	"\x12completion_options\x18\n" +
	" \x01(\v20.yandex.cloud.ai.assistants.v1.CompletionOptionsR\x11completionOptions\x129\n" +
	"\x05tools\x18\v \x03(\v2#.yandex.cloud.ai.assistants.v1.ToolR\x05tools\x12V\n" +
	"\x0fresponse_format\x18\f \x01(\v2-.yandex.cloud.ai.assistants.v1.ResponseFormatR\x0eresponseFormat\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"A\n" +
	"\x16DeleteAssistantRequest\x12'\n" +
	"\fassistant_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\vassistantId\"\x19\n" +
	"\x17DeleteAssistantResponse\"v\n" +
	"\x15ListAssistantsRequest\x12!\n" +
	"\tfolder_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bfolderId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"\x8a\x01\n" +
	"\x16ListAssistantsResponse\x12H\n" +
	"\n" +
	"assistants\x18\x01 \x03(\v2(.yandex.cloud.ai.assistants.v1.AssistantR\n" +
	"assistants\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\x83\x01\n" +
	"\x1cListAssistantVersionsRequest\x12'\n" +
	"\fassistant_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\vassistantId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"\xa7\x01\n" +
	"\x10AssistantVersion\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12;\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12F\n" +
	"\tassistant\x18\x03 \x01(\v2(.yandex.cloud.ai.assistants.v1.AssistantR\tassistant\"\x94\x01\n" +
	"\x1dListAssistantVersionsResponse\x12K\n" +
	"\bversions\x18\x01 \x03(\v2/.yandex.cloud.ai.assistants.v1.AssistantVersionR\bversions\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken2\xde\a\n" +
	"\x10AssistantService\x12\x8f\x01\n" +
	"\x06Create\x125.yandex.cloud.ai.assistants.v1.CreateAssistantRequest\x1a(.yandex.cloud.ai.assistants.v1.Assistant\"$\x82\xd3\xe4\x93\x02\x1e:\x01*\"\x19/assistants/v1/assistants\x12\x95\x01\n" +
	"\x03Get\x122.yandex.cloud.ai.assistants.v1.GetAssistantRequest\x1a(.yandex.cloud.ai.assistants.v1.Assistant\"0\x82\xd3\xe4\x93\x02*\x12(/assistants/v1/assistants/{assistant_id}\x12\x9e\x01\n" +
	"\x06Update\x125.yandex.cloud.ai.assistants.v1.UpdateAssistantRequest\x1a(.yandex.cloud.ai.assistants.v1.Assistant\"3\x82\xd3\xe4\x93\x02-:\x01*2(/assistants/v1/assistants/{assistant_id}\x12\xa9\x01\n" +
	"\x06Delete\x125.yandex.cloud.ai.assistants.v1.DeleteAssistantRequest\x1a6.yandex.cloud.ai.assistants.v1.DeleteAssistantResponse\"0\x82\xd3\xe4\x93\x02**(/assistants/v1/assistants/{assistant_id}\x12\x96\x01\n" +
	"\x04List\x124.yandex.cloud.ai.assistants.v1.ListAssistantsRequest\x1a5.yandex.cloud.ai.assistants.v1.ListAssistantsResponse\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/assistants/v1/assistants\x12\xb9\x01\n" +
	"\fListVersions\x12;.yandex.cloud.ai.assistants.v1.ListAssistantVersionsRequest\x1a<.yandex.cloud.ai.assistants.v1.ListAssistantVersionsResponse\".\x82\xd3\xe4\x93\x02(\x12&/assistants/v1/assistants:listVersionsBq\n" +
	"!yandex.cloud.api.ai.assistants.v1ZLgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1;assistantsb\x06proto3"

var (
	file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescData []byte
)

func file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_yandex_cloud_ai_assistants_v1_assistant_service_proto_goTypes = []any{
	(*CreateAssistantRequest)(nil),        // 0: yandex.cloud.ai.assistants.v1.CreateAssistantRequest
	(*GetAssistantRequest)(nil),           // 1: yandex.cloud.ai.assistants.v1.GetAssistantRequest
	(*UpdateAssistantRequest)(nil),        // 2: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest
	(*DeleteAssistantRequest)(nil),        // 3: yandex.cloud.ai.assistants.v1.DeleteAssistantRequest
	(*DeleteAssistantResponse)(nil),       // 4: yandex.cloud.ai.assistants.v1.DeleteAssistantResponse
	(*ListAssistantsRequest)(nil),         // 5: yandex.cloud.ai.assistants.v1.ListAssistantsRequest
	(*ListAssistantsResponse)(nil),        // 6: yandex.cloud.ai.assistants.v1.ListAssistantsResponse
	(*ListAssistantVersionsRequest)(nil),  // 7: yandex.cloud.ai.assistants.v1.ListAssistantVersionsRequest
	(*AssistantVersion)(nil),              // 8: yandex.cloud.ai.assistants.v1.AssistantVersion
	(*ListAssistantVersionsResponse)(nil), // 9: yandex.cloud.ai.assistants.v1.ListAssistantVersionsResponse
	nil,                                   // 10: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.LabelsEntry
	nil,                                   // 11: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.LabelsEntry
	(*common.ExpirationConfig)(nil),       // 12: yandex.cloud.ai.common.ExpirationConfig
	(*PromptTruncationOptions)(nil),       // 13: yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	(*CompletionOptions)(nil),             // 14: yandex.cloud.ai.assistants.v1.CompletionOptions
	(*Tool)(nil),                          // 15: yandex.cloud.ai.assistants.v1.Tool
	(*ResponseFormat)(nil),                // 16: yandex.cloud.ai.assistants.v1.ResponseFormat
	(*fieldmaskpb.FieldMask)(nil),         // 17: google.protobuf.FieldMask
	(*Assistant)(nil),                     // 18: yandex.cloud.ai.assistants.v1.Assistant
}
var file_yandex_cloud_ai_assistants_v1_assistant_service_proto_depIdxs = []int32{
	12, // 0: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.expiration_config:type_name -> yandex.cloud.ai.common.ExpirationConfig
	10, // 1: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.labels:type_name -> yandex.cloud.ai.assistants.v1.CreateAssistantRequest.LabelsEntry
	13, // 2: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.prompt_truncation_options:type_name -> yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	14, // 3: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.completion_options:type_name -> yandex.cloud.ai.assistants.v1.CompletionOptions
	15, // 4: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.tools:type_name -> yandex.cloud.ai.assistants.v1.Tool
	16, // 5: yandex.cloud.ai.assistants.v1.CreateAssistantRequest.response_format:type_name -> yandex.cloud.ai.assistants.v1.ResponseFormat
	17, // 6: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.update_mask:type_name -> google.protobuf.FieldMask
	12, // 7: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.expiration_config:type_name -> yandex.cloud.ai.common.ExpirationConfig
	11, // 8: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.labels:type_name -> yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.LabelsEntry
	13, // 9: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.prompt_truncation_options:type_name -> yandex.cloud.ai.assistants.v1.PromptTruncationOptions
	14, // 10: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.completion_options:type_name -> yandex.cloud.ai.assistants.v1.CompletionOptions
	15, // 11: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.tools:type_name -> yandex.cloud.ai.assistants.v1.Tool
	16, // 12: yandex.cloud.ai.assistants.v1.UpdateAssistantRequest.response_format:type_name -> yandex.cloud.ai.assistants.v1.ResponseFormat
	18, // 13: yandex.cloud.ai.assistants.v1.ListAssistantsResponse.assistants:type_name -> yandex.cloud.ai.assistants.v1.Assistant
	17, // 14: yandex.cloud.ai.assistants.v1.AssistantVersion.update_mask:type_name -> google.protobuf.FieldMask
	18, // 15: yandex.cloud.ai.assistants.v1.AssistantVersion.assistant:type_name -> yandex.cloud.ai.assistants.v1.Assistant
	8,  // 16: yandex.cloud.ai.assistants.v1.ListAssistantVersionsResponse.versions:type_name -> yandex.cloud.ai.assistants.v1.AssistantVersion
	0,  // 17: yandex.cloud.ai.assistants.v1.AssistantService.Create:input_type -> yandex.cloud.ai.assistants.v1.CreateAssistantRequest
	1,  // 18: yandex.cloud.ai.assistants.v1.AssistantService.Get:input_type -> yandex.cloud.ai.assistants.v1.GetAssistantRequest
	2,  // 19: yandex.cloud.ai.assistants.v1.AssistantService.Update:input_type -> yandex.cloud.ai.assistants.v1.UpdateAssistantRequest
	3,  // 20: yandex.cloud.ai.assistants.v1.AssistantService.Delete:input_type -> yandex.cloud.ai.assistants.v1.DeleteAssistantRequest
	5,  // 21: yandex.cloud.ai.assistants.v1.AssistantService.List:input_type -> yandex.cloud.ai.assistants.v1.ListAssistantsRequest
	7,  // 22: yandex.cloud.ai.assistants.v1.AssistantService.ListVersions:input_type -> yandex.cloud.ai.assistants.v1.ListAssistantVersionsRequest
	18, // 23: yandex.cloud.ai.assistants.v1.AssistantService.Create:output_type -> yandex.cloud.ai.assistants.v1.Assistant
	18, // 24: yandex.cloud.ai.assistants.v1.AssistantService.Get:output_type -> yandex.cloud.ai.assistants.v1.Assistant
	18, // 25: yandex.cloud.ai.assistants.v1.AssistantService.Update:output_type -> yandex.cloud.ai.assistants.v1.Assistant
	4,  // 26: yandex.cloud.ai.assistants.v1.AssistantService.Delete:output_type -> yandex.cloud.ai.assistants.v1.DeleteAssistantResponse
	6,  // 27: yandex.cloud.ai.assistants.v1.AssistantService.List:output_type -> yandex.cloud.ai.assistants.v1.ListAssistantsResponse
	9,  // 28: yandex.cloud.ai.assistants.v1.AssistantService.ListVersions:output_type -> yandex.cloud.ai.assistants.v1.ListAssistantVersionsResponse
	23, // [23:29] is the sub-list for method output_type
	17, // [17:23] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_assistant_service_proto_init() }
func file_yandex_cloud_ai_assistants_v1_assistant_service_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_assistant_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_assistant_proto_init()
	file_yandex_cloud_ai_assistants_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_assistant_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_assistant_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_assistant_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_assistant_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_assistant_service_proto = out.File
	file_yandex_cloud_ai_assistants_v1_assistant_service_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_assistant_service_proto_depIdxs = nil
}
