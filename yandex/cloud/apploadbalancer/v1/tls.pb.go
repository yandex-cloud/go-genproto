// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/apploadbalancer/v1/tls.proto

package apploadbalancer

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

// A TLS validation context resource.
type ValidationContext struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// TLS certificate issued by a trusted certificate authority (CA).
	//
	// Types that are valid to be assigned to TrustedCa:
	//
	//	*ValidationContext_TrustedCaId
	//	*ValidationContext_TrustedCaBytes
	TrustedCa     isValidationContext_TrustedCa `protobuf_oneof:"trusted_ca"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ValidationContext) Reset() {
	*x = ValidationContext{}
	mi := &file_yandex_cloud_apploadbalancer_v1_tls_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidationContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidationContext) ProtoMessage() {}

func (x *ValidationContext) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_tls_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidationContext.ProtoReflect.Descriptor instead.
func (*ValidationContext) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescGZIP(), []int{0}
}

func (x *ValidationContext) GetTrustedCa() isValidationContext_TrustedCa {
	if x != nil {
		return x.TrustedCa
	}
	return nil
}

func (x *ValidationContext) GetTrustedCaId() string {
	if x != nil {
		if x, ok := x.TrustedCa.(*ValidationContext_TrustedCaId); ok {
			return x.TrustedCaId
		}
	}
	return ""
}

func (x *ValidationContext) GetTrustedCaBytes() string {
	if x != nil {
		if x, ok := x.TrustedCa.(*ValidationContext_TrustedCaBytes); ok {
			return x.TrustedCaBytes
		}
	}
	return ""
}

type isValidationContext_TrustedCa interface {
	isValidationContext_TrustedCa()
}

type ValidationContext_TrustedCaId struct {
	TrustedCaId string `protobuf:"bytes,1,opt,name=trusted_ca_id,json=trustedCaId,proto3,oneof"`
}

type ValidationContext_TrustedCaBytes struct {
	// X.509 certificate contents in PEM format.
	TrustedCaBytes string `protobuf:"bytes,2,opt,name=trusted_ca_bytes,json=trustedCaBytes,proto3,oneof"`
}

func (*ValidationContext_TrustedCaId) isValidationContext_TrustedCa() {}

func (*ValidationContext_TrustedCaBytes) isValidationContext_TrustedCa() {}

var File_yandex_cloud_apploadbalancer_v1_tls_proto protoreflect.FileDescriptor

const file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDesc = "" +
	"\n" +
	")yandex/cloud/apploadbalancer/v1/tls.proto\x12\x1fyandex.cloud.apploadbalancer.v1\x1a\x1dyandex/cloud/validation.proto\"y\n" +
	"\x11ValidationContext\x12$\n" +
	"\rtrusted_ca_id\x18\x01 \x01(\tH\x00R\vtrustedCaId\x12*\n" +
	"\x10trusted_ca_bytes\x18\x02 \x01(\tH\x00R\x0etrustedCaBytesB\x12\n" +
	"\n" +
	"trusted_ca\x12\x04\xc0\xc11\x01Bz\n" +
	"#yandex.cloud.api.apploadbalancer.v1ZSgithub.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1;apploadbalancerb\x06proto3"

var (
	file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescData []byte
)

func file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDesc)))
	})
	return file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDescData
}

var file_yandex_cloud_apploadbalancer_v1_tls_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_apploadbalancer_v1_tls_proto_goTypes = []any{
	(*ValidationContext)(nil), // 0: yandex.cloud.apploadbalancer.v1.ValidationContext
}
var file_yandex_cloud_apploadbalancer_v1_tls_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_apploadbalancer_v1_tls_proto_init() }
func file_yandex_cloud_apploadbalancer_v1_tls_proto_init() {
	if File_yandex_cloud_apploadbalancer_v1_tls_proto != nil {
		return
	}
	file_yandex_cloud_apploadbalancer_v1_tls_proto_msgTypes[0].OneofWrappers = []any{
		(*ValidationContext_TrustedCaId)(nil),
		(*ValidationContext_TrustedCaBytes)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_tls_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_apploadbalancer_v1_tls_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_apploadbalancer_v1_tls_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_apploadbalancer_v1_tls_proto_msgTypes,
	}.Build()
	File_yandex_cloud_apploadbalancer_v1_tls_proto = out.File
	file_yandex_cloud_apploadbalancer_v1_tls_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_tls_proto_depIdxs = nil
}
