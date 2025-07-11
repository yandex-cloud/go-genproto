// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/iam/v1/api_key.proto

package iam

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// An ApiKey resource. For more information, see [Api-Key](/docs/iam/concepts/authorization/api-key).
type ApiKey struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the API Key.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the service account that the API key belongs to.
	ServiceAccountId string `protobuf:"bytes,2,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Description of the API key. 0-256 characters long.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Timestamp for the last authentication using this API key.
	LastUsedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
	// Draft
	// Scope of the API key. 0-256 characters long.
	//
	// Deprecated: Marked as deprecated in yandex/cloud/iam/v1/api_key.proto.
	Scope string `protobuf:"bytes,6,opt,name=scope,proto3" json:"scope,omitempty"`
	// Scopes of the API key. 0-256 characters long.
	Scopes []string `protobuf:"bytes,8,rep,name=scopes,proto3" json:"scopes,omitempty"`
	// API key expiration timestamp.
	ExpiresAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApiKey) Reset() {
	*x = ApiKey{}
	mi := &file_yandex_cloud_iam_v1_api_key_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApiKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApiKey) ProtoMessage() {}

func (x *ApiKey) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_api_key_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApiKey.ProtoReflect.Descriptor instead.
func (*ApiKey) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_api_key_proto_rawDescGZIP(), []int{0}
}

func (x *ApiKey) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ApiKey) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *ApiKey) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *ApiKey) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ApiKey) GetLastUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsedAt
	}
	return nil
}

// Deprecated: Marked as deprecated in yandex/cloud/iam/v1/api_key.proto.
func (x *ApiKey) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *ApiKey) GetScopes() []string {
	if x != nil {
		return x.Scopes
	}
	return nil
}

func (x *ApiKey) GetExpiresAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ExpiresAt
	}
	return nil
}

var File_yandex_cloud_iam_v1_api_key_proto protoreflect.FileDescriptor

const file_yandex_cloud_iam_v1_api_key_proto_rawDesc = "" +
	"\n" +
	"!yandex/cloud/iam/v1/api_key.proto\x12\x13yandex.cloud.iam.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xce\x02\n" +
	"\x06ApiKey\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12,\n" +
	"\x12service_account_id\x18\x02 \x01(\tR\x10serviceAccountId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12 \n" +
	"\vdescription\x18\x04 \x01(\tR\vdescription\x12<\n" +
	"\flast_used_at\x18\x05 \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"lastUsedAt\x12\x18\n" +
	"\x05scope\x18\x06 \x01(\tB\x02\x18\x01R\x05scope\x12\x16\n" +
	"\x06scopes\x18\b \x03(\tR\x06scopes\x129\n" +
	"\n" +
	"expires_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\texpiresAtBV\n" +
	"\x17yandex.cloud.api.iam.v1Z;github.com/yandex-cloud/go-genproto/yandex/cloud/iam/v1;iamb\x06proto3"

var (
	file_yandex_cloud_iam_v1_api_key_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_api_key_proto_rawDescData []byte
)

func file_yandex_cloud_iam_v1_api_key_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_api_key_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_api_key_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_iam_v1_api_key_proto_rawDesc), len(file_yandex_cloud_iam_v1_api_key_proto_rawDesc)))
	})
	return file_yandex_cloud_iam_v1_api_key_proto_rawDescData
}

var file_yandex_cloud_iam_v1_api_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_iam_v1_api_key_proto_goTypes = []any{
	(*ApiKey)(nil),                // 0: yandex.cloud.iam.v1.ApiKey
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_iam_v1_api_key_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.iam.v1.ApiKey.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.iam.v1.ApiKey.last_used_at:type_name -> google.protobuf.Timestamp
	1, // 2: yandex.cloud.iam.v1.ApiKey.expires_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_api_key_proto_init() }
func file_yandex_cloud_iam_v1_api_key_proto_init() {
	if File_yandex_cloud_iam_v1_api_key_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_iam_v1_api_key_proto_rawDesc), len(file_yandex_cloud_iam_v1_api_key_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iam_v1_api_key_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_api_key_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_iam_v1_api_key_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_api_key_proto = out.File
	file_yandex_cloud_iam_v1_api_key_proto_goTypes = nil
	file_yandex_cloud_iam_v1_api_key_proto_depIdxs = nil
}
