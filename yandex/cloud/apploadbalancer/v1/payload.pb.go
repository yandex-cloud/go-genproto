// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/apploadbalancer/v1/payload.proto

package apploadbalancer

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

// A health check payload resource.
type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Payload.
	//
	// Types that are assignable to Payload:
	//
	//	*Payload_Text
	Payload isPayload_Payload `protobuf_oneof:"payload"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
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

func (m *Payload) GetPayload() isPayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *Payload) GetText() string {
	if x, ok := x.GetPayload().(*Payload_Text); ok {
		return x.Text
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

var file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x38, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1c, 0x0a, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0x8a, 0xc8, 0x31, 0x02, 0x3e, 0x30,
	0x48, 0x00, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x42, 0x0f, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01, 0x42, 0x7a, 0x0a, 0x23, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x70,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x70, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescOnce sync.Once
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescData = file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc
)

func file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescGZIP() []byte {
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDescData)
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
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Payload); i {
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
	file_yandex_cloud_apploadbalancer_v1_payload_proto_msgTypes[0].OneofWrappers = []any{
		(*Payload_Text)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc,
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
	file_yandex_cloud_apploadbalancer_v1_payload_proto_rawDesc = nil
	file_yandex_cloud_apploadbalancer_v1_payload_proto_goTypes = nil
	file_yandex_cloud_apploadbalancer_v1_payload_proto_depIdxs = nil
}
