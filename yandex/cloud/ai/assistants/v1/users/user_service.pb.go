// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/users/user_service.proto

package users

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

// Request message for creating a new user.
type CreateUserRequest struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	FolderId string                 `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the user.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the user.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Source      string `protobuf:"bytes,4,opt,name=source,proto3" json:"source,omitempty"`
	// Expiration configuration for the user.
	ExpirationConfig *common.ExpirationConfig `protobuf:"bytes,5,opt,name=expiration_config,json=expirationConfig,proto3" json:"expiration_config,omitempty"`
	// Set of key-value pairs to label the user.
	Labels        map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateUserRequest) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *CreateUserRequest) GetExpirationConfig() *common.ExpirationConfig {
	if x != nil {
		return x.ExpirationConfig
	}
	return nil
}

func (x *CreateUserRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Request message for retrieving a user by ID.
type GetUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user to retrieve.
	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Request message for updating an existing user.
type UpdateUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user to update.
	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// A field mask specifying which fields to update.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// New name for the user.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// New description for the user.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// New expiration configuration for the user.
	ExpirationConfig *common.ExpirationConfig `protobuf:"bytes,5,opt,name=expiration_config,json=expirationConfig,proto3" json:"expiration_config,omitempty"`
	// New set of labels for the user.
	Labels        map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateUserRequest) GetExpirationConfig() *common.ExpirationConfig {
	if x != nil {
		return x.ExpirationConfig
	}
	return nil
}

func (x *UpdateUserRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Request message for deleting a user by ID.
type DeleteUserRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user to delete.
	UserId        string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response message for the delete operation.
type DeleteUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserResponse) Reset() {
	*x = DeleteUserResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserResponse) ProtoMessage() {}

func (x *DeleteUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{4}
}

// Request message for listing users in a specific folder.
type ListUsersRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Folder ID from which to list users.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Maximum number of users to return per page.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Token to retrieve the next page of results.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersRequest) Reset() {
	*x = ListUsersRequest{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersRequest) ProtoMessage() {}

func (x *ListUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersRequest.ProtoReflect.Descriptor instead.
func (*ListUsersRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{5}
}

func (x *ListUsersRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *ListUsersRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListUsersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// Response message for the list operation.
type ListUsersResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of users in the specified folder.
	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	// Token to retrieve the next page of results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP(), []int{6}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *ListUsersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_yandex_cloud_ai_assistants_v1_users_user_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDesc = "" +
	"\n" +
	"6yandex/cloud/ai/assistants/v1/users/user_service.proto\x12#yandex.cloud.ai.assistants.v1.users\x1a#yandex/cloud/ai/common/common.proto\x1a.yandex/cloud/ai/assistants/v1/users/user.proto\x1a\x1dyandex/cloud/validation.proto\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\"\xf2\x02\n" +
	"\x11CreateUserRequest\x12!\n" +
	"\tfolder_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bfolderId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x03 \x01(\tR\vdescription\x12\x16\n" +
	"\x06source\x18\x04 \x01(\tR\x06source\x12U\n" +
	"\x11expiration_config\x18\x05 \x01(\v2(.yandex.cloud.ai.common.ExpirationConfigR\x10expirationConfig\x12Z\n" +
	"\x06labels\x18\x06 \x03(\v2B.yandex.cloud.ai.assistants.v1.users.CreateUserRequest.LabelsEntryR\x06labels\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"/\n" +
	"\x0eGetUserRequest\x12\x1d\n" +
	"\auser_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x06userId\"\x99\x03\n" +
	"\x11UpdateUserRequest\x12\x1d\n" +
	"\auser_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x06userId\x12A\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskB\x04\xe8\xc71\x01R\n" +
	"updateMask\x12\x12\n" +
	"\x04name\x18\x03 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12U\n" +
	"\x11expiration_config\x18\x05 \x01(\v2(.yandex.cloud.ai.common.ExpirationConfigR\x10expirationConfig\x12Z\n" +
	"\x06labels\x18\x06 \x03(\v2B.yandex.cloud.ai.assistants.v1.users.UpdateUserRequest.LabelsEntryR\x06labels\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"2\n" +
	"\x11DeleteUserRequest\x12\x1d\n" +
	"\auser_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x06userId\"\x14\n" +
	"\x12DeleteUserResponse\"q\n" +
	"\x10ListUsersRequest\x12!\n" +
	"\tfolder_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bfolderId\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\x03R\bpageSize\x12\x1d\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tR\tpageToken\"|\n" +
	"\x11ListUsersResponse\x12?\n" +
	"\x05users\x18\x01 \x03(\v2).yandex.cloud.ai.assistants.v1.users.UserR\x05users\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken2\xe6\x05\n" +
	"\vUserService\x12\x87\x01\n" +
	"\x06Create\x126.yandex.cloud.ai.assistants.v1.users.CreateUserRequest\x1a).yandex.cloud.ai.assistants.v1.users.User\"\x1a\x82\xd3\xe4\x93\x02\x14:\x01*\"\x0f/users/v1/users\x12\x88\x01\n" +
	"\x03Get\x123.yandex.cloud.ai.assistants.v1.users.GetUserRequest\x1a).yandex.cloud.ai.assistants.v1.users.User\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/users/v1/users/{user_id}\x12\x91\x01\n" +
	"\x06Update\x126.yandex.cloud.ai.assistants.v1.users.UpdateUserRequest\x1a).yandex.cloud.ai.assistants.v1.users.User\"$\x82\xd3\xe4\x93\x02\x1e:\x01*2\x19/users/v1/users/{user_id}\x12\x9c\x01\n" +
	"\x06Delete\x126.yandex.cloud.ai.assistants.v1.users.DeleteUserRequest\x1a7.yandex.cloud.ai.assistants.v1.users.DeleteUserResponse\"!\x82\xd3\xe4\x93\x02\x1b*\x19/users/v1/users/{user_id}\x12\x8e\x01\n" +
	"\x04List\x125.yandex.cloud.ai.assistants.v1.users.ListUsersRequest\x1a6.yandex.cloud.ai.assistants.v1.users.ListUsersResponse\"\x17\x82\xd3\xe4\x93\x02\x11\x12\x0f/users/v1/usersBx\n" +
	"'yandex.cloud.api.ai.assistants.v1.usersZMgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1/users;usersb\x06proto3"

var (
	file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescData []byte
)

func file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_ai_assistants_v1_users_user_service_proto_goTypes = []any{
	(*CreateUserRequest)(nil),       // 0: yandex.cloud.ai.assistants.v1.users.CreateUserRequest
	(*GetUserRequest)(nil),          // 1: yandex.cloud.ai.assistants.v1.users.GetUserRequest
	(*UpdateUserRequest)(nil),       // 2: yandex.cloud.ai.assistants.v1.users.UpdateUserRequest
	(*DeleteUserRequest)(nil),       // 3: yandex.cloud.ai.assistants.v1.users.DeleteUserRequest
	(*DeleteUserResponse)(nil),      // 4: yandex.cloud.ai.assistants.v1.users.DeleteUserResponse
	(*ListUsersRequest)(nil),        // 5: yandex.cloud.ai.assistants.v1.users.ListUsersRequest
	(*ListUsersResponse)(nil),       // 6: yandex.cloud.ai.assistants.v1.users.ListUsersResponse
	nil,                             // 7: yandex.cloud.ai.assistants.v1.users.CreateUserRequest.LabelsEntry
	nil,                             // 8: yandex.cloud.ai.assistants.v1.users.UpdateUserRequest.LabelsEntry
	(*common.ExpirationConfig)(nil), // 9: yandex.cloud.ai.common.ExpirationConfig
	(*fieldmaskpb.FieldMask)(nil),   // 10: google.protobuf.FieldMask
	(*User)(nil),                    // 11: yandex.cloud.ai.assistants.v1.users.User
}
var file_yandex_cloud_ai_assistants_v1_users_user_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.ai.assistants.v1.users.CreateUserRequest.expiration_config:type_name -> yandex.cloud.ai.common.ExpirationConfig
	7,  // 1: yandex.cloud.ai.assistants.v1.users.CreateUserRequest.labels:type_name -> yandex.cloud.ai.assistants.v1.users.CreateUserRequest.LabelsEntry
	10, // 2: yandex.cloud.ai.assistants.v1.users.UpdateUserRequest.update_mask:type_name -> google.protobuf.FieldMask
	9,  // 3: yandex.cloud.ai.assistants.v1.users.UpdateUserRequest.expiration_config:type_name -> yandex.cloud.ai.common.ExpirationConfig
	8,  // 4: yandex.cloud.ai.assistants.v1.users.UpdateUserRequest.labels:type_name -> yandex.cloud.ai.assistants.v1.users.UpdateUserRequest.LabelsEntry
	11, // 5: yandex.cloud.ai.assistants.v1.users.ListUsersResponse.users:type_name -> yandex.cloud.ai.assistants.v1.users.User
	0,  // 6: yandex.cloud.ai.assistants.v1.users.UserService.Create:input_type -> yandex.cloud.ai.assistants.v1.users.CreateUserRequest
	1,  // 7: yandex.cloud.ai.assistants.v1.users.UserService.Get:input_type -> yandex.cloud.ai.assistants.v1.users.GetUserRequest
	2,  // 8: yandex.cloud.ai.assistants.v1.users.UserService.Update:input_type -> yandex.cloud.ai.assistants.v1.users.UpdateUserRequest
	3,  // 9: yandex.cloud.ai.assistants.v1.users.UserService.Delete:input_type -> yandex.cloud.ai.assistants.v1.users.DeleteUserRequest
	5,  // 10: yandex.cloud.ai.assistants.v1.users.UserService.List:input_type -> yandex.cloud.ai.assistants.v1.users.ListUsersRequest
	11, // 11: yandex.cloud.ai.assistants.v1.users.UserService.Create:output_type -> yandex.cloud.ai.assistants.v1.users.User
	11, // 12: yandex.cloud.ai.assistants.v1.users.UserService.Get:output_type -> yandex.cloud.ai.assistants.v1.users.User
	11, // 13: yandex.cloud.ai.assistants.v1.users.UserService.Update:output_type -> yandex.cloud.ai.assistants.v1.users.User
	4,  // 14: yandex.cloud.ai.assistants.v1.users.UserService.Delete:output_type -> yandex.cloud.ai.assistants.v1.users.DeleteUserResponse
	6,  // 15: yandex.cloud.ai.assistants.v1.users.UserService.List:output_type -> yandex.cloud.ai.assistants.v1.users.ListUsersResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_users_user_service_proto_init() }
func file_yandex_cloud_ai_assistants_v1_users_user_service_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_users_user_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_users_user_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_users_user_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_users_user_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_users_user_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_users_user_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_users_user_service_proto = out.File
	file_yandex_cloud_ai_assistants_v1_users_user_service_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_users_user_service_proto_depIdxs = nil
}
