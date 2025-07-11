// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/apploadbalancer/v1/payload.proto

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

// A health check payload resource.
type Payload struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Payload.
	//
	// Types that are valid to be assigned to Payload:
	//
	//	*Payload_Text
	Payload       isPayload_Payload `protobuf_oneof:"payload"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Payload) Reset() {
	*x = Payload{}
	mi := &file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetPayload() isPayload_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Payload) GetText() string {
	if x != nil {
		if x, ok := x.Payload.(*Payload_Text); ok {
			return x.Text
		}
	}
	return ""
}

type isPayload_Payload interface {
	isPayload_Payload()
}

type Payload_Text struct {
	// Payload text.
	Text string `protobuf:"bytes,1,opt,name=text,proto3,oneof"`
}

func (*Payload_Text) isPayload_Payload() {}

var File_yandex_cloud_apploadbalancer_v1_payload_proto protoreflect.FileDescriptor

const file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc = "" +
	"\n" +
	"-yandex/cloud/apploadbalancer/v1/payload.proto\x12\x1fyandex.cloud.apploadbalancer.v1\x1a\x1dyandex/cloud/validation.proto\"8\n" +
	"\aPayload\x12\x1c\n" +
	"\x04text\x18\x01 \x01(\tB\x06\x8a\xc81\x02>0H\x00R\x04textB\x0f\n" +
	"\apayload\x12\x04\xc0\xc11\x01Bz\n" +
	"#yandex.cloud.api.apploadbalancer.v1ZSgithub.com/yandex-cloud/go-genproto/yandex/cloud/apploadbalancer/v1;apploadbalancerb\x06proto3"

var (
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescData []byte
)

func file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc)))
	})
	return file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescData
}

var file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_apploadbalancer_v1_payload_proto_goTypes = []any{
	(*Payload)(nil), // 0: yandex.cloud.apploadbalancer.v1.Payload
}
var file_yandex_cloud_apploadbalancer_v1_payload_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_apploadbalancer_v1_payload_proto_init() }
func file_yandex_cloud_apploadbalancer_v1_payload_proto_init() {
	if File_yandex_cloud_apploadbalancer_v1_payload_proto != nil {
		return
	}
	file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0].OneofWrappers = []any{
		(*Payload_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc), len(file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_apploadbalancer_v1_payload_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_apploadbalancer_v1_payload_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes,
	}.Build()
	File_yandex_cloud_apploadbalancer_v1_payload_proto = out.File
	file_yandex_cloud_apploadbalancer_v1_payload_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_payload_proto_depIdxs = nil
}
