// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/iam/v1/user_account.proto

package iam

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// Currently represents only [Yandex account](/docs/iam/concepts/users/accounts#passport).
type UserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the user account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to UserAccount:
	//
	//	*UserAccount_YandexPassportUserAccount
	//	*UserAccount_SamlUserAccount
	UserAccount isUserAccount_UserAccount `protobuf_oneof:"user_account"`
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAccount.ProtoReflect.Descriptor instead.
func (*UserAccount) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_user_account_proto_rawDescGZIP(), []int{0}
}

func (x *UserAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *UserAccount) GetUserAccount() isUserAccount_UserAccount {
	if m != nil {
		return m.UserAccount
	}
	return nil
}

func (x *UserAccount) GetYandexPassportUserAccount() *YandexPassportUserAccount {
	if x, ok := x.GetUserAccount().(*UserAccount_YandexPassportUserAccount); ok {
		return x.YandexPassportUserAccount
	}
	return nil
}

func (x *UserAccount) GetSamlUserAccount() *SamlUserAccount {
	if x, ok := x.GetUserAccount().(*UserAccount_SamlUserAccount); ok {
		return x.SamlUserAccount
	}
	return nil
}

type isUserAccount_UserAccount interface {
	isUserAccount_UserAccount()
}

type UserAccount_YandexPassportUserAccount struct {
	// A YandexPassportUserAccount resource.
	YandexPassportUserAccount *YandexPassportUserAccount `protobuf:"bytes,2,opt,name=yandex_passport_user_account,json=yandexPassportUserAccount,proto3,oneof"`
}

type UserAccount_SamlUserAccount struct {
	// A SAML federated user.
	SamlUserAccount *SamlUserAccount `protobuf:"bytes,3,opt,name=saml_user_account,json=samlUserAccount,proto3,oneof"`
}

func (*UserAccount_YandexPassportUserAccount) isUserAccount_UserAccount() {}

func (*UserAccount_SamlUserAccount) isUserAccount_UserAccount() {}

// A YandexPassportUserAccount resource.
// For more information, see [Yandex account](/docs/iam/concepts/users/accounts#passport).
type YandexPassportUserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Login of the Yandex user account.
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	// Default email of the Yandex user account.
	DefaultEmail string `protobuf:"bytes,2,opt,name=default_email,json=defaultEmail,proto3" json:"default_email,omitempty"`
}

func (x *YandexPassportUserAccount) Reset() {
	*x = YandexPassportUserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *YandexPassportUserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YandexPassportUserAccount) ProtoMessage() {}

func (x *YandexPassportUserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YandexPassportUserAccount.ProtoReflect.Descriptor instead.
func (*YandexPassportUserAccount) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_user_account_proto_rawDescGZIP(), []int{1}
}

func (x *YandexPassportUserAccount) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *YandexPassportUserAccount) GetDefaultEmail() string {
	if x != nil {
		return x.DefaultEmail
	}
	return ""
}

// A SAML federated user.
// For more information, see [federations](/docs/iam/concepts/federations).
type SamlUserAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the federation that the federation belongs to.
	FederationId string `protobuf:"bytes,1,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Name Id of the SAML federated user.
	// The name is unique within the federation. 1-256 characters long.
	NameId string `protobuf:"bytes,2,opt,name=name_id,json=nameId,proto3" json:"name_id,omitempty"`
	// Additional attributes of the SAML federated user.
	Attributes map[string]*SamlUserAccount_Attribute `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SamlUserAccount) Reset() {
	*x = SamlUserAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamlUserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamlUserAccount) ProtoMessage() {}

func (x *SamlUserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamlUserAccount.ProtoReflect.Descriptor instead.
func (*SamlUserAccount) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_user_account_proto_rawDescGZIP(), []int{2}
}

func (x *SamlUserAccount) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

func (x *SamlUserAccount) GetNameId() string {
	if x != nil {
		return x.NameId
	}
	return ""
}

func (x *SamlUserAccount) GetAttributes() map[string]*SamlUserAccount_Attribute {
	if x != nil {
		return x.Attributes
	}
	return nil
}

type SamlUserAccount_Attribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *SamlUserAccount_Attribute) Reset() {
	*x = SamlUserAccount_Attribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SamlUserAccount_Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamlUserAccount_Attribute) ProtoMessage() {}

func (x *SamlUserAccount_Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_user_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SamlUserAccount_Attribute.ProtoReflect.Descriptor instead.
func (*SamlUserAccount_Attribute) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_user_account_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SamlUserAccount_Attribute) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_yandex_cloud_iam_v1_user_account_proto protoreflect.FileDescriptor

var file_yandex_cloud_iam_v1_user_account_proto_rawDesc = []byte{
	0x0a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x71, 0x0a, 0x1c,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x59, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x50,
	0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x48, 0x00, 0x52, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x50, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x52, 0x0a, 0x11, 0x73, 0x61, 0x6d, 0x6c, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x61, 0x6d, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x48, 0x00, 0x52, 0x0f, 0x73, 0x61, 0x6d, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x42, 0x14, 0x0a, 0x0c, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x22, 0x56, 0x0a, 0x19, 0x59, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x50, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x23, 0x0a, 0x0d,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0xd4, 0x02, 0x0a, 0x0f, 0x53, 0x61, 0x6d, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0c, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x07, 0x6e, 0x61, 0x6d, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a,
	0xc8, 0x31, 0x05, 0x31, 0x2d, 0x32, 0x35, 0x36, 0x52, 0x06, 0x6e, 0x61, 0x6d, 0x65, 0x49, 0x64,
	0x12, 0x54, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x1a, 0x21, 0x0a, 0x09, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x6d, 0x0a, 0x0f, 0x41, 0x74, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x44,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x6d, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_iam_v1_user_account_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_user_account_proto_rawDescData = file_yandex_cloud_iam_v1_user_account_proto_rawDesc
)

func file_yandex_cloud_iam_v1_user_account_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_user_account_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_iam_v1_user_account_proto_rawDescData)
	})
	return file_yandex_cloud_iam_v1_user_account_proto_rawDescData
}

var file_yandex_cloud_iam_v1_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_iam_v1_user_account_proto_goTypes = []any{
	(*UserAccount)(nil),               // 0: yandex.cloud.iam.v1.UserAccount
	(*YandexPassportUserAccount)(nil), // 1: yandex.cloud.iam.v1.YandexPassportUserAccount
	(*SamlUserAccount)(nil),           // 2: yandex.cloud.iam.v1.SamlUserAccount
	(*SamlUserAccount_Attribute)(nil), // 3: yandex.cloud.iam.v1.SamlUserAccount.Attribute
	nil,                               // 4: yandex.cloud.iam.v1.SamlUserAccount.AttributesEntry
}
var file_yandex_cloud_iam_v1_user_account_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.iam.v1.UserAccount.yandex_passport_user_account:type_name -> yandex.cloud.iam.v1.YandexPassportUserAccount
	2, // 1: yandex.cloud.iam.v1.UserAccount.saml_user_account:type_name -> yandex.cloud.iam.v1.SamlUserAccount
	4, // 2: yandex.cloud.iam.v1.SamlUserAccount.attributes:type_name -> yandex.cloud.iam.v1.SamlUserAccount.AttributesEntry
	3, // 3: yandex.cloud.iam.v1.SamlUserAccount.AttributesEntry.value:type_name -> yandex.cloud.iam.v1.SamlUserAccount.Attribute
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_user_account_proto_init() }
func file_yandex_cloud_iam_v1_user_account_proto_init() {
	if File_yandex_cloud_iam_v1_user_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_iam_v1_user_account_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*UserAccount); i {
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
		file_yandex_cloud_iam_v1_user_account_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*YandexPassportUserAccount); i {
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
		file_yandex_cloud_iam_v1_user_account_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SamlUserAccount); i {
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
		file_yandex_cloud_iam_v1_user_account_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SamlUserAccount_Attribute); i {
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
	file_yandex_cloud_iam_v1_user_account_proto_msgTypes[0].OneofWrappers = []any{
		(*UserAccount_YandexPassportUserAccount)(nil),
		(*UserAccount_SamlUserAccount)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_iam_v1_user_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iam_v1_user_account_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_user_account_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_iam_v1_user_account_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_user_account_proto = out.File
	file_yandex_cloud_iam_v1_user_account_proto_rawDesc = nil
	file_yandex_cloud_iam_v1_user_account_proto_goTypes = nil
	file_yandex_cloud_iam_v1_user_account_proto_depIdxs = nil
}
