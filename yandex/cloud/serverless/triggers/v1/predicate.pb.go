// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/serverless/triggers/v1/predicate.proto

package triggers

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

type Predicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Predicate:
	//
	//	*Predicate_AndPredicate
	//	*Predicate_FieldValuePredicate
	Predicate isPredicate_Predicate `protobuf_oneof:"predicate"`
}

func (x *Predicate) Reset() {
	*x = Predicate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Predicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Predicate) ProtoMessage() {}

func (x *Predicate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Predicate.ProtoReflect.Descriptor instead.
func (*Predicate) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescGZIP(), []int{0}
}

func (m *Predicate) GetPredicate() isPredicate_Predicate {
	if m != nil {
		return m.Predicate
	}
	return nil
}

func (x *Predicate) GetAndPredicate() *AndPredicate {
	if x, ok := x.GetPredicate().(*Predicate_AndPredicate); ok {
		return x.AndPredicate
	}
	return nil
}

func (x *Predicate) GetFieldValuePredicate() *FieldValuePredicate {
	if x, ok := x.GetPredicate().(*Predicate_FieldValuePredicate); ok {
		return x.FieldValuePredicate
	}
	return nil
}

type isPredicate_Predicate interface {
	isPredicate_Predicate()
}

type Predicate_AndPredicate struct {
	AndPredicate *AndPredicate `protobuf:"bytes,2,opt,name=and_predicate,json=andPredicate,proto3,oneof"`
}

type Predicate_FieldValuePredicate struct {
	FieldValuePredicate *FieldValuePredicate `protobuf:"bytes,4,opt,name=field_value_predicate,json=fieldValuePredicate,proto3,oneof"`
}

func (*Predicate_AndPredicate) isPredicate_Predicate() {}

func (*Predicate_FieldValuePredicate) isPredicate_Predicate() {}

type AndPredicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Predicate []*Predicate `protobuf:"bytes,1,rep,name=predicate,proto3" json:"predicate,omitempty"`
}

func (x *AndPredicate) Reset() {
	*x = AndPredicate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AndPredicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AndPredicate) ProtoMessage() {}

func (x *AndPredicate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AndPredicate.ProtoReflect.Descriptor instead.
func (*AndPredicate) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescGZIP(), []int{1}
}

func (x *AndPredicate) GetPredicate() []*Predicate {
	if x != nil {
		return x.Predicate
	}
	return nil
}

type FieldValuePredicate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FieldPath string `protobuf:"bytes,1,opt,name=field_path,json=fieldPath,proto3" json:"field_path,omitempty"`
	// Types that are assignable to Value:
	//
	//	*FieldValuePredicate_Exact
	//	*FieldValuePredicate_Prefix
	//	*FieldValuePredicate_Suffix
	Value isFieldValuePredicate_Value `protobuf_oneof:"value"`
}

func (x *FieldValuePredicate) Reset() {
	*x = FieldValuePredicate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FieldValuePredicate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldValuePredicate) ProtoMessage() {}

func (x *FieldValuePredicate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldValuePredicate.ProtoReflect.Descriptor instead.
func (*FieldValuePredicate) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescGZIP(), []int{2}
}

func (x *FieldValuePredicate) GetFieldPath() string {
	if x != nil {
		return x.FieldPath
	}
	return ""
}

func (m *FieldValuePredicate) GetValue() isFieldValuePredicate_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *FieldValuePredicate) GetExact() string {
	if x, ok := x.GetValue().(*FieldValuePredicate_Exact); ok {
		return x.Exact
	}
	return ""
}

func (x *FieldValuePredicate) GetPrefix() string {
	if x, ok := x.GetValue().(*FieldValuePredicate_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (x *FieldValuePredicate) GetSuffix() string {
	if x, ok := x.GetValue().(*FieldValuePredicate_Suffix); ok {
		return x.Suffix
	}
	return ""
}

type isFieldValuePredicate_Value interface {
	isFieldValuePredicate_Value()
}

type FieldValuePredicate_Exact struct {
	Exact string `protobuf:"bytes,3,opt,name=exact,proto3,oneof"` // string representation of the value matches exactly to the given string
}

type FieldValuePredicate_Prefix struct {
	Prefix string `protobuf:"bytes,8,opt,name=prefix,proto3,oneof"` // value has given prefix
}

type FieldValuePredicate_Suffix struct {
	Suffix string `protobuf:"bytes,9,opt,name=suffix,proto3,oneof"` // value has given suffix
}

func (*FieldValuePredicate_Exact) isFieldValuePredicate_Value() {}

func (*FieldValuePredicate_Prefix) isFieldValuePredicate_Value() {}

func (*FieldValuePredicate_Suffix) isFieldValuePredicate_Value() {}

var File_yandex_cloud_serverless_triggers_v1_predicate_proto protoreflect.FileDescriptor

var file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDesc = []byte{
	0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65,
	0x72, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74,
	0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x09, 0x50, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x58, 0x0a, 0x0d, 0x61, 0x6e, 0x64, 0x5f, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x48, 0x00, 0x52, 0x0c, 0x61, 0x6e, 0x64, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x12, 0x6e, 0x0a, 0x15, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x38, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67,
	0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x13, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x42, 0x11, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x04,
	0xc0, 0xc1, 0x31, 0x01, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x22, 0x5c, 0x0a, 0x0c, 0x41, 0x6e,
	0x64, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x70, 0x72,
	0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x09, 0x70,
	0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x13, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x50, 0x72, 0x65, 0x64, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x12, 0x23, 0x0a, 0x0a, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x09, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x78, 0x61, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x18, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x42, 0x0d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x04, 0xc0, 0xc1, 0x31, 0x01,
	0x4a, 0x04, 0x08, 0x04, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x42, 0x7b, 0x0a, 0x27,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x74, 0x72, 0x69, 0x67,
	0x67, 0x65, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x6c, 0x65, 0x73, 0x73, 0x2f, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x74, 0x72, 0x69, 0x67, 0x67, 0x65, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescData = file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDesc
)

func file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescData)
	})
	return file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDescData
}

var file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_serverless_triggers_v1_predicate_proto_goTypes = []any{
	(*Predicate)(nil),           // 0: yandex.cloud.serverless.triggers.v1.Predicate
	(*AndPredicate)(nil),        // 1: yandex.cloud.serverless.triggers.v1.AndPredicate
	(*FieldValuePredicate)(nil), // 2: yandex.cloud.serverless.triggers.v1.FieldValuePredicate
}
var file_yandex_cloud_serverless_triggers_v1_predicate_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.serverless.triggers.v1.Predicate.and_predicate:type_name -> yandex.cloud.serverless.triggers.v1.AndPredicate
	2, // 1: yandex.cloud.serverless.triggers.v1.Predicate.field_value_predicate:type_name -> yandex.cloud.serverless.triggers.v1.FieldValuePredicate
	0, // 2: yandex.cloud.serverless.triggers.v1.AndPredicate.predicate:type_name -> yandex.cloud.serverless.triggers.v1.Predicate
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_triggers_v1_predicate_proto_init() }
func file_yandex_cloud_serverless_triggers_v1_predicate_proto_init() {
	if File_yandex_cloud_serverless_triggers_v1_predicate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Predicate); i {
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
		file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AndPredicate); i {
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
		file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*FieldValuePredicate); i {
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
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[0].OneofWrappers = []any{
		(*Predicate_AndPredicate)(nil),
		(*Predicate_FieldValuePredicate)(nil),
	}
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes[2].OneofWrappers = []any{
		(*FieldValuePredicate_Exact)(nil),
		(*FieldValuePredicate_Prefix)(nil),
		(*FieldValuePredicate_Suffix)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_triggers_v1_predicate_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_triggers_v1_predicate_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_serverless_triggers_v1_predicate_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_triggers_v1_predicate_proto = out.File
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_rawDesc = nil
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_goTypes = nil
	file_yandex_cloud_serverless_triggers_v1_predicate_proto_depIdxs = nil
}
