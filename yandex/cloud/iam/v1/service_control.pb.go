// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/iam/v1/service_control.proto

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

type Service_Status int32

const (
	Service_STATUS_UNSPECIFIED Service_Status = 0
	// The service is enabled.
	Service_ENABLED Service_Status = 1
	// The service is paused.
	Service_PAUSED Service_Status = 2
	// The service is disabled.
	Service_DISABLED Service_Status = 3
	// The service is being enabled.
	Service_ENABLING Service_Status = 4
	// The service is being resumed.
	Service_RESUMING Service_Status = 5
	// The service is being paused.
	Service_PAUSING Service_Status = 6
	// The service is being disabled.
	Service_DISABLING Service_Status = 7
	// The service is in error state.
	Service_ERROR Service_Status = 8
)

// Enum value maps for Service_Status.
var (
	Service_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "ENABLED",
		2: "PAUSED",
		3: "DISABLED",
		4: "ENABLING",
		5: "RESUMING",
		6: "PAUSING",
		7: "DISABLING",
		8: "ERROR",
	}
	Service_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"ENABLED":            1,
		"PAUSED":             2,
		"DISABLED":           3,
		"ENABLING":           4,
		"RESUMING":           5,
		"PAUSING":            6,
		"DISABLING":          7,
		"ERROR":              8,
	}
)

func (x Service_Status) Enum() *Service_Status {
	p := new(Service_Status)
	*p = x
	return p
}

func (x Service_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_iam_v1_service_control_proto_enumTypes[0].Descriptor()
}

func (Service_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_iam_v1_service_control_proto_enumTypes[0]
}

func (x Service_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service_Status.Descriptor instead.
func (Service_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_service_control_proto_rawDescGZIP(), []int{0, 0}
}

// A Service.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the service.
	ServiceId string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	// Resource that the service belongs to.
	Resource *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	// Time of the last status update of the service.
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// Current status of the service.
	Status Service_Status `protobuf:"varint,4,opt,name=status,proto3,enum=yandex.cloud.iam.v1.Service_Status" json:"status,omitempty"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iam_v1_service_control_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iam_v1_service_control_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iam_v1_service_control_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Service) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Service) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Service) GetStatus() Service_Status {
	if x != nil {
		return x.Status
	}
	return Service_STATUS_UNSPECIFIED
}

var File_yandex_cloud_iam_v1_service_control_proto protoreflect.FileDescriptor

var file_yandex_cloud_iam_v1_service_control_proto_rawDesc = []byte{
	0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x22, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x69, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe8, 0x02, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64,
	0x12, 0x39, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x8a, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x50, 0x41, 0x55, 0x53, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0c, 0x0a,
	0x08, 0x45, 0x4e, 0x41, 0x42, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x45, 0x53, 0x55, 0x4d, 0x49, 0x4e, 0x47, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x55,
	0x53, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c,
	0x49, 0x4e, 0x47, 0x10, 0x07, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x08,
	0x42, 0x56, 0x0a, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x5a, 0x3b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x61,
	0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x69, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_iam_v1_service_control_proto_rawDescOnce sync.Once
	file_yandex_cloud_iam_v1_service_control_proto_rawDescData = file_yandex_cloud_iam_v1_service_control_proto_rawDesc
)

func file_yandex_cloud_iam_v1_service_control_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iam_v1_service_control_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iam_v1_service_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_iam_v1_service_control_proto_rawDescData)
	})
	return file_yandex_cloud_iam_v1_service_control_proto_rawDescData
}

var file_yandex_cloud_iam_v1_service_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_iam_v1_service_control_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_iam_v1_service_control_proto_goTypes = []interface{}{
	(Service_Status)(0),           // 0: yandex.cloud.iam.v1.Service.Status
	(*Service)(nil),               // 1: yandex.cloud.iam.v1.Service
	(*Resource)(nil),              // 2: yandex.cloud.iam.v1.Resource
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_yandex_cloud_iam_v1_service_control_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.iam.v1.Service.resource:type_name -> yandex.cloud.iam.v1.Resource
	3, // 1: yandex.cloud.iam.v1.Service.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: yandex.cloud.iam.v1.Service.status:type_name -> yandex.cloud.iam.v1.Service.Status
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iam_v1_service_control_proto_init() }
func file_yandex_cloud_iam_v1_service_control_proto_init() {
	if File_yandex_cloud_iam_v1_service_control_proto != nil {
		return
	}
	file_yandex_cloud_iam_v1_resource_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_iam_v1_service_control_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
			RawDescriptor: file_yandex_cloud_iam_v1_service_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iam_v1_service_control_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iam_v1_service_control_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_iam_v1_service_control_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_iam_v1_service_control_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iam_v1_service_control_proto = out.File
	file_yandex_cloud_iam_v1_service_control_proto_rawDesc = nil
	file_yandex_cloud_iam_v1_service_control_proto_goTypes = nil
	file_yandex_cloud_iam_v1_service_control_proto_depIdxs = nil
}
