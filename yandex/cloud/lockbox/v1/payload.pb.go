// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/lockbox/v1/payload.proto

package lockbox

import (
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

// A payload.
type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the version that the payload belongs to.
	VersionId string `protobuf:"bytes,1,opt,name=version_id,json=versionId,proto3" json:"version_id,omitempty"`
	// Payload entries.
	Entries []*Payload_Entry `protobuf:"bytes,2,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[0]
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
	return file_yandex_cloud_lockbox_v1_payload_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetVersionId() string {
	if x != nil {
		return x.VersionId
	}
	return ""
}

func (x *Payload) GetEntries() []*Payload_Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type Payload_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Non-confidential key of the entry.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// Confidential value of the entry.
	//
	// Types that are assignable to Value:
	//
	//	*Payload_Entry_TextValue
	//	*Payload_Entry_BinaryValue
	Value isPayload_Entry_Value `protobuf_oneof:"value"`
}

func (x *Payload_Entry) Reset() {
	*x = Payload_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload_Entry) ProtoMessage() {}

func (x *Payload_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload_Entry.ProtoReflect.Descriptor instead.
func (*Payload_Entry) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_lockbox_v1_payload_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Payload_Entry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (m *Payload_Entry) GetValue() isPayload_Entry_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Payload_Entry) GetTextValue() string {
	if x, ok := x.GetValue().(*Payload_Entry_TextValue); ok {
		return x.TextValue
	}
	return ""
}

func (x *Payload_Entry) GetBinaryValue() []byte {
	if x, ok := x.GetValue().(*Payload_Entry_BinaryValue); ok {
		return x.BinaryValue
	}
	return nil
}

type isPayload_Entry_Value interface {
	isPayload_Entry_Value()
}

type Payload_Entry_TextValue struct {
	// Text value.
	TextValue string `protobuf:"bytes,2,opt,name=text_value,json=textValue,proto3,oneof"`
}

type Payload_Entry_BinaryValue struct {
	// Binary value.
	BinaryValue []byte `protobuf:"bytes,3,opt,name=binary_value,json=binaryValue,proto3,oneof"`
}

func (*Payload_Entry_TextValue) isPayload_Entry_Value() {}

func (*Payload_Entry_BinaryValue) isPayload_Entry_Value() {}

var File_yandex_cloud_lockbox_v1_payload_proto protoreflect.FileDescriptor

var file_yandex_cloud_lockbox_v1_payload_proto_rawDesc = []byte{
	0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31,
	0x22, 0xd4, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x68, 0x0a,
	0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1f, 0x0a, 0x0a, 0x74, 0x65, 0x78, 0x74,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09,
	0x74, 0x65, 0x78, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x62, 0x69, 0x6e,
	0x61, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x0b, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x63, 0x6b,
	0x62, 0x6f, 0x78, 0x2e, 0x76, 0x31, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78,
	0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63, 0x6b, 0x62, 0x6f, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_lockbox_v1_payload_proto_rawDescOnce sync.Once
	file_yandex_cloud_lockbox_v1_payload_proto_rawDescData = file_yandex_cloud_lockbox_v1_payload_proto_rawDesc
)

func file_yandex_cloud_lockbox_v1_payload_proto_rawDescGZIP() []byte {
	file_yandex_cloud_lockbox_v1_payload_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_lockbox_v1_payload_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_lockbox_v1_payload_proto_rawDescData)
	})
	return file_yandex_cloud_lockbox_v1_payload_proto_rawDescData
}

var file_yandex_cloud_lockbox_v1_payload_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_lockbox_v1_payload_proto_goTypes = []any{
	(*Payload)(nil),       // 0: yandex.cloud.lockbox.v1.Payload
	(*Payload_Entry)(nil), // 1: yandex.cloud.lockbox.v1.Payload.Entry
}
var file_yandex_cloud_lockbox_v1_payload_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.lockbox.v1.Payload.entries:type_name -> yandex.cloud.lockbox.v1.Payload.Entry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_lockbox_v1_payload_proto_init() }
func file_yandex_cloud_lockbox_v1_payload_proto_init() {
	if File_yandex_cloud_lockbox_v1_payload_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[0].Exporter = func(v any, i int) any {
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
		file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Payload_Entry); i {
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
	file_yandex_cloud_lockbox_v1_payload_proto_msgTypes[1].OneofWrappers = []any{
		(*Payload_Entry_TextValue)(nil),
		(*Payload_Entry_BinaryValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_lockbox_v1_payload_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_lockbox_v1_payload_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_lockbox_v1_payload_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_lockbox_v1_payload_proto_msgTypes,
	}.Build()
	File_yandex_cloud_lockbox_v1_payload_proto = out.File
	file_yandex_cloud_lockbox_v1_payload_proto_rawDesc = nil
	file_yandex_cloud_lockbox_v1_payload_proto_goTypes = nil
	file_yandex_cloud_lockbox_v1_payload_proto_depIdxs = nil
}
