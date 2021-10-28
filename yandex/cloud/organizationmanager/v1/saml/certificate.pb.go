// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: yandex/cloud/organizationmanager/v1/saml/certificate.proto

package saml

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// A certificate.
type Certificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the certificate.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the federation that the certificate belongs to.
	FederationId string `protobuf:"bytes,2,opt,name=federation_id,json=federationId,proto3" json:"federation_id,omitempty"`
	// Name of the certificate.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the certificate.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Certificate data in PEM format.
	Data string `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Certificate) Reset() {
	*x = Certificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Certificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Certificate) ProtoMessage() {}

func (x *Certificate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Certificate.ProtoReflect.Descriptor instead.
func (*Certificate) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescGZIP(), []int{0}
}

func (x *Certificate) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Certificate) GetFederationId() string {
	if x != nil {
		return x.FederationId
	}
	return ""
}

func (x *Certificate) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Certificate) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Certificate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Certificate) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_yandex_cloud_organizationmanager_v1_saml_certificate_proto protoreflect.FileDescriptor

var file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x28, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x02, 0x0a, 0x0b, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0d, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0c, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xf2, 0xc7, 0x31, 0x1d, 0x7c, 0x5b, 0x61, 0x2d, 0x7a,
	0x5d, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x31, 0x7d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x36, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x07, 0x3c, 0x3d,
	0x33, 0x32, 0x30, 0x30, 0x30, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x81, 0x01, 0x0a, 0x2c,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x73, 0x61, 0x6d, 0x6c, 0x5a, 0x51, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6d, 0x6c, 0x3b, 0x73, 0x61, 0x6d, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescOnce sync.Once
	file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescData = file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDesc
)

func file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescGZIP() []byte {
	file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescData)
	})
	return file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDescData
}

var file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_goTypes = []interface{}{
	(*Certificate)(nil),           // 0: yandex.cloud.organizationmanager.v1.saml.Certificate
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.organizationmanager.v1.saml.Certificate.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_init() }
func file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_init() {
	if File_yandex_cloud_organizationmanager_v1_saml_certificate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Certificate); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_msgTypes,
	}.Build()
	File_yandex_cloud_organizationmanager_v1_saml_certificate_proto = out.File
	file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_rawDesc = nil
	file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_goTypes = nil
	file_yandex_cloud_organizationmanager_v1_saml_certificate_proto_depIdxs = nil
}
