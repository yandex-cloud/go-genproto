// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/organizationmanager/v1/ssh_certificate_service.proto

package organizationmanager

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GenerateSshCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Scope:
	//
	//	*GenerateSshCertificateRequest_CloudId
	//	*GenerateSshCertificateRequest_OrganizationId
	Scope isGenerateSshCertificateRequest_Scope `protobuf_oneof:"scope"`
	// Types that are assignable to Subject:
	//
	//	*GenerateSshCertificateRequest_SubjectId
	//	*GenerateSshCertificateRequest_OsLogin
	Subject   isGenerateSshCertificateRequest_Subject `protobuf_oneof:"subject"`
	PublicKey string                                  `protobuf:"bytes,5,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *GenerateSshCertificateRequest) Reset() {
	*x = GenerateSshCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSshCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSshCertificateRequest) ProtoMessage() {}

func (x *GenerateSshCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSshCertificateRequest.ProtoReflect.Descriptor instead.
func (*GenerateSshCertificateRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescGZIP(), []int{0}
}

func (m *GenerateSshCertificateRequest) GetScope() isGenerateSshCertificateRequest_Scope {
	if m != nil {
		return m.Scope
	}
	return nil
}

func (x *GenerateSshCertificateRequest) GetCloudId() string {
	if x, ok := x.GetScope().(*GenerateSshCertificateRequest_CloudId); ok {
		return x.CloudId
	}
	return ""
}

func (x *GenerateSshCertificateRequest) GetOrganizationId() string {
	if x, ok := x.GetScope().(*GenerateSshCertificateRequest_OrganizationId); ok {
		return x.OrganizationId
	}
	return ""
}

func (m *GenerateSshCertificateRequest) GetSubject() isGenerateSshCertificateRequest_Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (x *GenerateSshCertificateRequest) GetSubjectId() string {
	if x, ok := x.GetSubject().(*GenerateSshCertificateRequest_SubjectId); ok {
		return x.SubjectId
	}
	return ""
}

func (x *GenerateSshCertificateRequest) GetOsLogin() string {
	if x, ok := x.GetSubject().(*GenerateSshCertificateRequest_OsLogin); ok {
		return x.OsLogin
	}
	return ""
}

func (x *GenerateSshCertificateRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

type isGenerateSshCertificateRequest_Scope interface {
	isGenerateSshCertificateRequest_Scope()
}

type GenerateSshCertificateRequest_CloudId struct {
	CloudId string `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3,oneof"` // the cloud must be attached to an organization
}

type GenerateSshCertificateRequest_OrganizationId struct {
	OrganizationId string `protobuf:"bytes,2,opt,name=organization_id,json=organizationId,proto3,oneof"`
}

func (*GenerateSshCertificateRequest_CloudId) isGenerateSshCertificateRequest_Scope() {}

func (*GenerateSshCertificateRequest_OrganizationId) isGenerateSshCertificateRequest_Scope() {}

type isGenerateSshCertificateRequest_Subject interface {
	isGenerateSshCertificateRequest_Subject()
}

type GenerateSshCertificateRequest_SubjectId struct {
	SubjectId string `protobuf:"bytes,3,opt,name=subject_id,json=subjectId,proto3,oneof"` // specify subject to generate certificate for default login
}

type GenerateSshCertificateRequest_OsLogin struct {
	OsLogin string `protobuf:"bytes,4,opt,name=os_login,json=osLogin,proto3,oneof"` // specify os_login for a specific login
}

func (*GenerateSshCertificateRequest_SubjectId) isGenerateSshCertificateRequest_Subject() {}

func (*GenerateSshCertificateRequest_OsLogin) isGenerateSshCertificateRequest_Subject() {}

type GenerateSshCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// as per specification https://cvsweb.openbsd.org/src/usr.bin/ssh/PROTOCOL.certkeys?annotate=HEAD
	SignedCertificate string `protobuf:"bytes,1,opt,name=signed_certificate,json=signedCertificate,proto3" json:"signed_certificate,omitempty"`
}

func (x *GenerateSshCertificateResponse) Reset() {
	*x = GenerateSshCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateSshCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateSshCertificateResponse) ProtoMessage() {}

func (x *GenerateSshCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateSshCertificateResponse.ProtoReflect.Descriptor instead.
func (*GenerateSshCertificateResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateSshCertificateResponse) GetSignedCertificate() string {
	if x != nil {
		return x.SignedCertificate
	}
	return ""
}

var File_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDesc = []byte{
	0x0a, 0x41, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x68, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x23, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x97, 0x02, 0x0a, 0x1d, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x53, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c,
	0x3d, 0x35, 0x30, 0x48, 0x00, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x33,
	0x0a, 0x0f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35,
	0x30, 0x48, 0x01, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x25,
	0x0a, 0x08, 0x6f, 0x73, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x08, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x33, 0x32, 0x48, 0x01, 0x52, 0x07, 0x6f, 0x73,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2e, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f,
	0x6b, 0x65, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0xe8, 0xc7, 0x31, 0x01, 0x8a,
	0xc8, 0x31, 0x07, 0x3c, 0x3d, 0x31, 0x35, 0x30, 0x30, 0x30, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x42, 0x0d, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x04,
	0xc0, 0xc1, 0x31, 0x01, 0x42, 0x09, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22,
	0x4f, 0x0a, 0x1e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x63, 0x65, 0x72, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73,
	0x69, 0x67, 0x6e, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x32, 0xeb, 0x01, 0x0a, 0x15, 0x53, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xd1, 0x01, 0x0a, 0x08, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x12, 0x42, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x43, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e,
	0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x53, 0x73, 0x68, 0x43, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x3c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x3a, 0x01, 0x2a, 0x22, 0x31, 0x2f, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x73, 0x68, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x73, 0x3a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x42, 0x86,
	0x01, 0x0a, 0x27, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x5a, 0x5b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2f, 0x76, 0x31, 0x3b, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescData = file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDesc
)

func file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescData)
	})
	return file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDescData
}

var file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_goTypes = []any{
	(*GenerateSshCertificateRequest)(nil),  // 0: yandex.cloud.organizationmanager.v1.GenerateSshCertificateRequest
	(*GenerateSshCertificateResponse)(nil), // 1: yandex.cloud.organizationmanager.v1.GenerateSshCertificateResponse
}
var file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.organizationmanager.v1.SshCertificateService.Generate:input_type -> yandex.cloud.organizationmanager.v1.GenerateSshCertificateRequest
	1, // 1: yandex.cloud.organizationmanager.v1.SshCertificateService.Generate:output_type -> yandex.cloud.organizationmanager.v1.GenerateSshCertificateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_init() }
func file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_init() {
	if File_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateSshCertificateRequest); i {
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
		file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateSshCertificateResponse); i {
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
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes[0].OneofWrappers = []any{
		(*GenerateSshCertificateRequest_CloudId)(nil),
		(*GenerateSshCertificateRequest_OrganizationId)(nil),
		(*GenerateSshCertificateRequest_SubjectId)(nil),
		(*GenerateSshCertificateRequest_OsLogin)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto = out.File
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_rawDesc = nil
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_goTypes = nil
	file_yandex_cloud_organizationmanager_v1_ssh_certificate_service_proto_depIdxs = nil
}
