// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/iam/v1/key.proto

package iam

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Key_Algorithm int32

const (
	Key_ALGORITHM_UNSPECIFIED Key_Algorithm = 0
	// RSA with a 2048-bit key size. Default value.
	Key_RSA_2048 Key_Algorithm = 1
	// RSA with a 4096-bit key size.
	Key_RSA_4096 Key_Algorithm = 2
)

// Enum value maps for Key_Algorithm.
var (
	Key_Algorithm_name = map[int32]string{
		0: "ALGORITHM_UNSPECIFIED",
		1: "RSA_2048",
		2: "RSA_4096",
	}
	Key_Algorithm_value = map[string]int32{
		"ALGORITHM_UNSPECIFIED": 0,
		"RSA_2048":              1,
		"RSA_4096":              2,
	}
)

func (x Key_Algorithm) Enum() *Key_Algorithm {
	p := new(Key_Algorithm)
	*p = x
	return p
}

func (x Key_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Key_Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_iam_v1_key_proto_enumTypes[0].Descriptor()
}

func (Key_Algorithm) Type() protoreflect.EnumType {
	return &file_yandex_cloud_iam_v1_key_proto_enumTypes[0]
}

func (x Key_Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Key_Algorithm.Descriptor instead.
func (Key_Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_key_proto_rawDescGZIP(), []int{0, 0}
}

// A Key resource. For more information, see [Authorized keys](/docs/iam/concepts/authorization/key).
type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Key resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are assignable to Subject:
	//
	//	*Key_UserAccountId
	//	*Key_ServiceAccountId
	Subject isKey_Subject `protobuf_oneof:"subject"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Description of the Key resource. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// An algorithm used to generate a key pair of the Key resource.
	KeyAlgorithm Key_Algorithm `protobuf:"varint,6,opt,name=key_algorithm,json=keyAlgorithm,proto3,enum=yandex.cloud.iam.v1.Key_Algorithm" json:"key_algorithm,omitempty"`
	// A public key of the Key resource.
	PublicKey string `protobuf:"bytes,7,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	// Timestamp for the last use of this key.
	LastUsedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_key_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_key_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_key_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (m *Key) GetSubject() isKey_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *Key) GetUserAccountId() string {
	if x, ok := x.GetSubject().(*Key_UserAccountId); ok {
		return x.UserAccountId
	}
	return ""
}

func (x *Key) GetServiceAccountId() string {
	if x, ok := x.GetSubject().(*Key_ServiceAccountId); ok {
		return x.ServiceAccountId
	}
	return ""
}

func (x *Key) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Key) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Key) GetKeyAlgorithm() Key_Algorithm {
	if x != nil {
		return x.KeyAlgorithm
	}
	return Key_ALGORITHM_UNSPECIFIED
}

func (x *Key) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *Key) GetLastUsedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUsedAt
	}
	return nil
}

type isKey_Subject interface {
	isKey_Subject()
}

type Key_UserAccountId struct {
	// ID of the user account that the Key resource belongs to.
	UserAccountId string `protobuf:"bytes,2,opt,name=user_account_id,json=userAccountId,proto3,oneof"`
}

type Key_ServiceAccountId struct {
	// ID of the service account that the Key resource belongs to.
	ServiceAccountId string `protobuf:"bytes,3,opt,name=service_account_id,json=serviceAccountId,proto3,oneof"`
}

func (*Key_UserAccountId) isKey_Subject() {}

func (*Key_ServiceAccountId) isKey_Subject() {}

var File_yandex_cloud_iam_v1_key_proto protoreflect.FileDescriptor

var file_yandex_cloud_iam_v1_key_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x6b, 0x65, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc7, 0x03, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x47, 0x0a, 0x0d, 0x6b, 0x65, 0x79, 0x5f, 0x61, 0x6c, 0x67, 0x6f,
	0x72, 0x69, 0x74, 0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x4b, 0x65, 0x79, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52,
	0x0c, 0x6b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x1d, 0x0a,
	0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x0c,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64, 0x41, 0x74, 0x22, 0x42, 0x0a, 0x09, 0x41, 0x6c,
	0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x19, 0x0a, 0x15, 0x41, 0x4c, 0x47, 0x4f, 0x52,
	0x49, 0x54, 0x48, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x53, 0x41, 0x5f, 0x32, 0x30, 0x34, 0x38, 0x10, 0x01,
	0x12, 0x0c, 0x0a, 0x08, 0x52, 0x53, 0x41, 0x5f, 0x34, 0x30, 0x39, 0x36, 0x10, 0x02, 0x42, 0x09,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x42,
	0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61, 0x6d,
	0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_iam_v1_key_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_key_proto_rawDescData = file_yandex_cloud_iam_v1_key_proto_rawDesc
)

func file_yandex_cloud_iam_v1_key_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_key_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_key_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_iam_v1_key_proto_rawDescData)
	})
	return file_yandex_cloud_iam_v1_key_proto_rawDescData
}

var file_yandex_cloud_iam_v1_key_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_iam_v1_key_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_iam_v1_key_proto_goTypes = []any{
	(Key_Algorithm)(0),            // 0: yandex.cloud.iam.v1.Key.Algorithm
	(*Key)(nil),                   // 1: yandex.cloud.iam.v1.Key
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_iam_v1_key_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.iam.v1.Key.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: yandex.cloud.iam.v1.Key.key_algorithm:type_name -> yandex.cloud.iam.v1.Key.Algorithm
	2, // 2: yandex.cloud.iam.v1.Key.last_used_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_key_proto_init() }
func file_yandex_cloud_iam_v1_key_proto_init() {
	if File_yandex_cloud_iam_v1_key_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_iam_v1_key_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Key); i {
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
	file_yandex_cloud_iam_v1_key_proto_msgTypes[0].OneofWrappers = []any{
		(*Key_UserAccountId)(nil),
		(*Key_ServiceAccountId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_iam_v1_key_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iam_v1_key_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_key_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_iam_v1_key_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_iam_v1_key_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_key_proto = out.File
	file_yandex_cloud_iam_v1_key_proto_rawDesc = nil
	file_yandex_cloud_iam_v1_key_proto_goTypes = nil
	file_yandex_cloud_iam_v1_key_proto_depIdxs = nil
}
