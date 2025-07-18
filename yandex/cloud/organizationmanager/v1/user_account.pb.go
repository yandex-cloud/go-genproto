// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/organizationmanager/v1/user_account.proto

package organizationmanager

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Currently represents only [Yandex account](/docs/iam/concepts/users/accounts#passport).
type UserAccount struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the user account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to UserAccount:
	//
	//	*UserAccount_YandexPassportUserAccount
	//	*UserAccount_SamlUserAccount
	UserAccount   isUserAccount_UserAccount `protobuf_oneof:"user_account"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserAccount) Reset() {
	*x = UserAccount{}
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAccount) ProtoMessage() {}

func (x *UserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[0]
	if x != nil {
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
	return file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescGZIP(), []int{0}
}

func (x *UserAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserAccount) GetUserAccount() isUserAccount_UserAccount {
	if x != nil {
		return x.UserAccount
	}
	return nil
}

func (x *UserAccount) GetYandexPassportUserAccount() *YandexPassportUserAccount {
	if x != nil {
		if x, ok := x.UserAccount.(*UserAccount_YandexPassportUserAccount); ok {
			return x.YandexPassportUserAccount
		}
	}
	return nil
}

func (x *UserAccount) GetSamlUserAccount() *SamlUserAccount {
	if x != nil {
		if x, ok := x.UserAccount.(*UserAccount_SamlUserAccount); ok {
			return x.SamlUserAccount
		}
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// Login of the Yandex user account.
	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	// Default email of the Yandex user account.
	DefaultEmail  string `protobuf:"bytes,2,opt,name=default_email,json=defaultEmail,proto3" json:"default_email,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YandexPassportUserAccount) Reset() {
	*x = YandexPassportUserAccount{}
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YandexPassportUserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YandexPassportUserAccount) ProtoMessage() {}

func (x *YandexPassportUserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[1]
	if x != nil {
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
	return file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescGZIP(), []int{1}
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
// For more information, see [federations](/docs/iam/concepts/users/accounts#saml-federation).
type SamlUserAccount struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the federation that the federation belongs to.
	FederationId string `protobuf:"bytes,1,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Name Id of the SAML federated user.
	// The name is unique within the federation. 1-256 characters long.
	NameId string `protobuf:"bytes,2,opt,name=name_id,json=nameId,proto3" json:"name_id,omitempty"`
	// Additional attributes of the SAML federated user.
	Attributes    map[string]*SamlUserAccount_Attribute `protobuf:"bytes,3,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SamlUserAccount) Reset() {
	*x = SamlUserAccount{}
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SamlUserAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamlUserAccount) ProtoMessage() {}

func (x *SamlUserAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[2]
	if x != nil {
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
	return file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescGZIP(), []int{2}
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
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         []string               `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SamlUserAccount_Attribute) Reset() {
	*x = SamlUserAccount_Attribute{}
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SamlUserAccount_Attribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SamlUserAccount_Attribute) ProtoMessage() {}

func (x *SamlUserAccount_Attribute) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[3]
	if x != nil {
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
	return file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SamlUserAccount_Attribute) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_yandex_cloud_organizationmanager_v1_user_account_proto protoreflect.FileDescriptor

const file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDesc = "" +
	"\n" +
	"6yandex/cloud/organizationmanager/v1/user_account.proto\x12#yandex.cloud.organizationmanager.v1\x1a\x1dyandex/cloud/validation.proto\"\x9b\x02\n" +
	"\vUserAccount\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x81\x01\n" +
	"\x1cyandex_passport_user_account\x18\x02 \x01(\v2>.yandex.cloud.organizationmanager.v1.YandexPassportUserAccountH\x00R\x19yandexPassportUserAccount\x12b\n" +
	"\x11saml_user_account\x18\x03 \x01(\v24.yandex.cloud.organizationmanager.v1.SamlUserAccountH\x00R\x0fsamlUserAccountB\x14\n" +
	"\fuser_account\x12\x04\xc0\xc11\x01\"V\n" +
	"\x19YandexPassportUserAccount\x12\x14\n" +
	"\x05login\x18\x01 \x01(\tR\x05login\x12#\n" +
	"\rdefault_email\x18\x02 \x01(\tR\fdefaultEmail\"\xf4\x02\n" +
	"\x0fSamlUserAccount\x121\n" +
	"\rfederation_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\ffederationId\x12&\n" +
	"\aname_id\x18\x02 \x01(\tB\r\xe8\xc71\x01\x8a\xc81\x051-256R\x06nameId\x12d\n" +
	"\n" +
	"attributes\x18\x03 \x03(\v2D.yandex.cloud.organizationmanager.v1.SamlUserAccount.AttributesEntryR\n" +
	"attributes\x1a!\n" +
	"\tAttribute\x12\x14\n" +
	"\x05value\x18\x01 \x03(\tR\x05value\x1a}\n" +
	"\x0fAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12T\n" +
	"\x05value\x18\x02 \x01(\v2>.yandex.cloud.organizationmanager.v1.SamlUserAccount.AttributeR\x05value:\x028\x01B\x86\x01\n" +
	"'yandex.cloud.api.organizationmanager.v1Z[github.com/yandex-cloud/go-genproto/yandex/cloud/organizationmanager/v1;organizationmanagerb\x06proto3"

var (
	file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescOnce sync.Once
	file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescData []byte
)

func file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescGZIP() []byte {
	file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDesc), len(file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDesc)))
	})
	return file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDescData
}

var file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_organizationmanager_v1_user_account_proto_goTypes = []any{
	(*UserAccount)(nil),               // 0: yandex.cloud.organizationmanager.v1.UserAccount
	(*YandexPassportUserAccount)(nil), // 1: yandex.cloud.organizationmanager.v1.YandexPassportUserAccount
	(*SamlUserAccount)(nil),           // 2: yandex.cloud.organizationmanager.v1.SamlUserAccount
	(*SamlUserAccount_Attribute)(nil), // 3: yandex.cloud.organizationmanager.v1.SamlUserAccount.Attribute
	nil,                               // 4: yandex.cloud.organizationmanager.v1.SamlUserAccount.AttributesEntry
}
var file_yandex_cloud_organizationmanager_v1_user_account_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.organizationmanager.v1.UserAccount.yandex_passport_user_account:type_name -> yandex.cloud.organizationmanager.v1.YandexPassportUserAccount
	2, // 1: yandex.cloud.organizationmanager.v1.UserAccount.saml_user_account:type_name -> yandex.cloud.organizationmanager.v1.SamlUserAccount
	4, // 2: yandex.cloud.organizationmanager.v1.SamlUserAccount.attributes:type_name -> yandex.cloud.organizationmanager.v1.SamlUserAccount.AttributesEntry
	3, // 3: yandex.cloud.organizationmanager.v1.SamlUserAccount.AttributesEntry.value:type_name -> yandex.cloud.organizationmanager.v1.SamlUserAccount.Attribute
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_organizationmanager_v1_user_account_proto_init() }
func file_yandex_cloud_organizationmanager_v1_user_account_proto_init() {
	if File_yandex_cloud_organizationmanager_v1_user_account_proto != nil {
		return
	}
	file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes[0].OneofWrappers = []any{
		(*UserAccount_YandexPassportUserAccount)(nil),
		(*UserAccount_SamlUserAccount)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDesc), len(file_yandex_cloud_organizationmanager_v1_user_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_organizationmanager_v1_user_account_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_organizationmanager_v1_user_account_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_organizationmanager_v1_user_account_proto_msgTypes,
	}.Build()
	File_yandex_cloud_organizationmanager_v1_user_account_proto = out.File
	file_yandex_cloud_organizationmanager_v1_user_account_proto_goTypes = nil
	file_yandex_cloud_organizationmanager_v1_user_account_proto_depIdxs = nil
}
