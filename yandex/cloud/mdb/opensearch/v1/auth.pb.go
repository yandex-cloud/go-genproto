// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/opensearch/v1/auth.proto

package opensearch

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type AuthSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// SAML settings
	Saml          *SAMLSettings `protobuf:"bytes,1,opt,name=saml,proto3" json:"saml,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AuthSettings) Reset() {
	*x = AuthSettings{}
	mi := &file_yandex_cloud_mdb_opensearch_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthSettings) ProtoMessage() {}

func (x *AuthSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_opensearch_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthSettings.ProtoReflect.Descriptor instead.
func (*AuthSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescGZIP(), []int{0}
}

func (x *AuthSettings) GetSaml() *SAMLSettings {
	if x != nil {
		return x.Saml
	}
	return nil
}

type SAMLSettings struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Enabled bool                   `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Required. The entity ID of your IdP.
	IdpEntityId string `protobuf:"bytes,2,opt,name=idp_entity_id,json=idpEntityId,proto3" json:"idp_entity_id,omitempty"`
	// Required. The SAML 2.0 metadata file of your IdP.
	IdpMetadataFile []byte `protobuf:"bytes,3,opt,name=idp_metadata_file,json=idpMetadataFile,proto3" json:"idp_metadata_file,omitempty"`
	// Required. The entity ID of the service provider.
	SpEntityId string `protobuf:"bytes,4,opt,name=sp_entity_id,json=spEntityId,proto3" json:"sp_entity_id,omitempty"`
	// Required. The OpenSearch Dashboards base URL.
	DashboardsUrl string `protobuf:"bytes,5,opt,name=dashboards_url,json=dashboardsUrl,proto3" json:"dashboards_url,omitempty"`
	// Optional. The attribute in the SAML response where the roles are stored. If not configured, no roles are used.
	RolesKey string `protobuf:"bytes,6,opt,name=roles_key,json=rolesKey,proto3" json:"roles_key,omitempty"`
	// Optional. The attribute in the SAML response where the subject is stored. If not configured, the NameID attribute is used.
	SubjectKey string `protobuf:"bytes,7,opt,name=subject_key,json=subjectKey,proto3" json:"subject_key,omitempty"`
	// default jwt expiration timeout.
	JwtDefaultExpirationTimeout *wrapperspb.Int64Value `protobuf:"bytes,8,opt,name=jwt_default_expiration_timeout,json=jwtDefaultExpirationTimeout,proto3" json:"jwt_default_expiration_timeout,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *SAMLSettings) Reset() {
	*x = SAMLSettings{}
	mi := &file_yandex_cloud_mdb_opensearch_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SAMLSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SAMLSettings) ProtoMessage() {}

func (x *SAMLSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_opensearch_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SAMLSettings.ProtoReflect.Descriptor instead.
func (*SAMLSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SAMLSettings) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *SAMLSettings) GetIdpEntityId() string {
	if x != nil {
		return x.IdpEntityId
	}
	return ""
}

func (x *SAMLSettings) GetIdpMetadataFile() []byte {
	if x != nil {
		return x.IdpMetadataFile
	}
	return nil
}

func (x *SAMLSettings) GetSpEntityId() string {
	if x != nil {
		return x.SpEntityId
	}
	return ""
}

func (x *SAMLSettings) GetDashboardsUrl() string {
	if x != nil {
		return x.DashboardsUrl
	}
	return ""
}

func (x *SAMLSettings) GetRolesKey() string {
	if x != nil {
		return x.RolesKey
	}
	return ""
}

func (x *SAMLSettings) GetSubjectKey() string {
	if x != nil {
		return x.SubjectKey
	}
	return ""
}

func (x *SAMLSettings) GetJwtDefaultExpirationTimeout() *wrapperspb.Int64Value {
	if x != nil {
		return x.JwtDefaultExpirationTimeout
	}
	return nil
}

var File_yandex_cloud_mdb_opensearch_v1_auth_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDesc = "" +
	"\n" +
	")yandex/cloud/mdb/opensearch/v1/auth.proto\x12\x1eyandex.cloud.mdb.opensearch.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1dyandex/cloud/validation.proto\"P\n" +
	"\fAuthSettings\x12@\n" +
	"\x04saml\x18\x01 \x01(\v2,.yandex.cloud.mdb.opensearch.v1.SAMLSettingsR\x04saml\"\xa5\x03\n" +
	"\fSAMLSettings\x12\x18\n" +
	"\aenabled\x18\x01 \x01(\bR\aenabled\x12-\n" +
	"\ridp_entity_id\x18\x02 \x01(\tB\t\x8a\xc81\x05<=250R\vidpEntityId\x127\n" +
	"\x11idp_metadata_file\x18\x03 \x01(\fB\v\x8a\xc81\a<=10000R\x0fidpMetadataFile\x12+\n" +
	"\fsp_entity_id\x18\x04 \x01(\tB\t\x8a\xc81\x05<=250R\n" +
	"spEntityId\x120\n" +
	"\x0edashboards_url\x18\x05 \x01(\tB\t\x8a\xc81\x05<=250R\rdashboardsUrl\x12&\n" +
	"\troles_key\x18\x06 \x01(\tB\t\x8a\xc81\x05<=250R\brolesKey\x12*\n" +
	"\vsubject_key\x18\a \x01(\tB\t\x8a\xc81\x05<=250R\n" +
	"subjectKey\x12`\n" +
	"\x1ejwt_default_expiration_timeout\x18\b \x01(\v2\x1b.google.protobuf.Int64ValueR\x1bjwtDefaultExpirationTimeoutBs\n" +
	"\"yandex.cloud.api.mdb.opensearch.v1ZMgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/opensearch/v1;opensearchb\x06proto3"

var (
	file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDesc), len(file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDescData
}

var file_yandex_cloud_mdb_opensearch_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_opensearch_v1_auth_proto_goTypes = []any{
	(*AuthSettings)(nil),          // 0: yandex.cloud.mdb.opensearch.v1.AuthSettings
	(*SAMLSettings)(nil),          // 1: yandex.cloud.mdb.opensearch.v1.SAMLSettings
	(*wrapperspb.Int64Value)(nil), // 2: google.protobuf.Int64Value
}
var file_yandex_cloud_mdb_opensearch_v1_auth_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.mdb.opensearch.v1.AuthSettings.saml:type_name -> yandex.cloud.mdb.opensearch.v1.SAMLSettings
	2, // 1: yandex.cloud.mdb.opensearch.v1.SAMLSettings.jwt_default_expiration_timeout:type_name -> google.protobuf.Int64Value
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_opensearch_v1_auth_proto_init() }
func file_yandex_cloud_mdb_opensearch_v1_auth_proto_init() {
	if File_yandex_cloud_mdb_opensearch_v1_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDesc), len(file_yandex_cloud_mdb_opensearch_v1_auth_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_opensearch_v1_auth_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_opensearch_v1_auth_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_opensearch_v1_auth_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_opensearch_v1_auth_proto = out.File
	file_yandex_cloud_mdb_opensearch_v1_auth_proto_goTypes = nil
	file_yandex_cloud_mdb_opensearch_v1_auth_proto_depIdxs = nil
}
