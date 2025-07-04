// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datasphere/v1/app_token_service.proto

package datasphere

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type AppTokenValidateRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// App token to validate.
	Token         string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AppTokenValidateRequest) Reset() {
	*x = AppTokenValidateRequest{}
	mi := &file_yandex_cloud_datasphere_v1_app_token_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AppTokenValidateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppTokenValidateRequest) ProtoMessage() {}

func (x *AppTokenValidateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_app_token_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppTokenValidateRequest.ProtoReflect.Descriptor instead.
func (*AppTokenValidateRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescGZIP(), []int{0}
}

func (x *AppTokenValidateRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_yandex_cloud_datasphere_v1_app_token_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDesc = "" +
	"\n" +
	"2yandex/cloud/datasphere/v1/app_token_service.proto\x12\x1ayandex.cloud.datasphere.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1bgoogle/protobuf/empty.proto\"/\n" +
	"\x17AppTokenValidateRequest\x12\x14\n" +
	"\x05token\x18\x01 \x01(\tR\x05token2\x9e\x01\n" +
	"\x0fAppTokenService\x12\x8a\x01\n" +
	"\bValidate\x123.yandex.cloud.datasphere.v1.AppTokenValidateRequest\x1a\x16.google.protobuf.Empty\"1\x82\xd3\xe4\x93\x02+\")/datasphere/v1/appTokens/{token}:validateBk\n" +
	"\x1eyandex.cloud.api.datasphere.v1ZIgithub.com/yandex-cloud/go-genproto/yandex/cloud/datasphere/v1;datasphereb\x06proto3"

var (
	file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescData []byte
)

func file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDesc), len(file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDesc)))
	})
	return file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDescData
}

var file_yandex_cloud_datasphere_v1_app_token_service_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_datasphere_v1_app_token_service_proto_goTypes = []any{
	(*AppTokenValidateRequest)(nil), // 0: yandex.cloud.datasphere.v1.AppTokenValidateRequest
	(*emptypb.Empty)(nil),           // 1: google.protobuf.Empty
}
var file_yandex_cloud_datasphere_v1_app_token_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.datasphere.v1.AppTokenService.Validate:input_type -> yandex.cloud.datasphere.v1.AppTokenValidateRequest
	1, // 1: yandex.cloud.datasphere.v1.AppTokenService.Validate:output_type -> google.protobuf.Empty
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v1_app_token_service_proto_init() }
func file_yandex_cloud_datasphere_v1_app_token_service_proto_init() {
	if File_yandex_cloud_datasphere_v1_app_token_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDesc), len(file_yandex_cloud_datasphere_v1_app_token_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_datasphere_v1_app_token_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v1_app_token_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datasphere_v1_app_token_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v1_app_token_service_proto = out.File
	file_yandex_cloud_datasphere_v1_app_token_service_proto_goTypes = nil
	file_yandex_cloud_datasphere_v1_app_token_service_proto_depIdxs = nil
}
