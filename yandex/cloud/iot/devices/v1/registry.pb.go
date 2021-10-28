// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: yandex/cloud/iot/devices/v1/registry.proto

package devices

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

type Registry_Status int32

const (
	Registry_STATUS_UNSPECIFIED Registry_Status = 0
	// Registry is being created.
	Registry_CREATING Registry_Status = 1
	// Registry is ready to use.
	Registry_ACTIVE Registry_Status = 2
	// Registry is being deleted.
	Registry_DELETING Registry_Status = 3
)

// Enum value maps for Registry_Status.
var (
	Registry_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
	}
	Registry_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"DELETING":           3,
	}
)

func (x Registry_Status) Enum() *Registry_Status {
	p := new(Registry_Status)
	*p = x
	return p
}

func (x Registry_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Registry_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_iot_devices_v1_registry_proto_enumTypes[0].Descriptor()
}

func (Registry_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_iot_devices_v1_registry_proto_enumTypes[0]
}

func (x Registry_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Registry_Status.Descriptor instead.
func (Registry_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{0, 0}
}

// A registry. For more information, see [Registry](/docs/iot-core/concepts/index#registry).
type Registry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the registry.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the registry belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the registry. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the registry. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Status of the registry.
	Status Registry_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.iot.devices.v1.Registry_Status" json:"status,omitempty"`
	// ID of the logs group for the specified registry.
	LogGroupId string `protobuf:"bytes,8,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
}

func (x *Registry) Reset() {
	*x = Registry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Registry.ProtoReflect.Descriptor instead.
func (*Registry) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{0}
}

func (x *Registry) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Registry) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Registry) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Registry) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Registry) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Registry) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Registry) GetStatus() Registry_Status {
	if x != nil {
		return x.Status
	}
	return Registry_STATUS_UNSPECIFIED
}

func (x *Registry) GetLogGroupId() string {
	if x != nil {
		return x.LogGroupId
	}
	return ""
}

// A registry certificate. For more information, see [Managing registry certificates](/docs/iot-core/operations/certificates/registry-certificates).
type RegistryCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the registry that the certificate belongs to.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// SHA256 hash of the certificates.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Public part of the certificate.
	CertificateData string `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *RegistryCertificate) Reset() {
	*x = RegistryCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryCertificate) ProtoMessage() {}

func (x *RegistryCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryCertificate.ProtoReflect.Descriptor instead.
func (*RegistryCertificate) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{1}
}

func (x *RegistryCertificate) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *RegistryCertificate) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *RegistryCertificate) GetCertificateData() string {
	if x != nil {
		return x.CertificateData
	}
	return ""
}

func (x *RegistryCertificate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// A device topic alias.
//
// Alias is an alternate name of a device topic assigned by the user. Map alias to canonical topic name prefix, e.g. `my/custom/alias` match to `$device/abcdef/events`. For more information, see [Using topic aliases](/docs/iot-core/concepts/topic#aliases).
type DeviceAlias struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the device that the alias belongs to.
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Prefix of a canonical topic name to be aliased, e.g. `$devices/abcdef`.
	TopicPrefix string `protobuf:"bytes,2,opt,name=topic_prefix,json=topicPrefix,proto3" json:"topic_prefix,omitempty"`
	// Alias of a device topic.
	Alias string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
}

func (x *DeviceAlias) Reset() {
	*x = DeviceAlias{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeviceAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceAlias) ProtoMessage() {}

func (x *DeviceAlias) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeviceAlias.ProtoReflect.Descriptor instead.
func (*DeviceAlias) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{2}
}

func (x *DeviceAlias) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *DeviceAlias) GetTopicPrefix() string {
	if x != nil {
		return x.TopicPrefix
	}
	return ""
}

func (x *DeviceAlias) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

// A registry password.
type RegistryPassword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the registry that the password belongs to.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// ID of the password.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *RegistryPassword) Reset() {
	*x = RegistryPassword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryPassword) ProtoMessage() {}

func (x *RegistryPassword) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryPassword.ProtoReflect.Descriptor instead.
func (*RegistryPassword) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{3}
}

func (x *RegistryPassword) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *RegistryPassword) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegistryPassword) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_yandex_cloud_iot_devices_v1_registry_proto protoreflect.FileDescriptor

var file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x6f, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x03, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x49, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x44, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69,
	0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x5f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x67,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x48, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x12,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x45, 0x41, 0x54, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x02, 0x12, 0x0c,
	0x0a, 0x08, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x03, 0x22, 0xbe, 0x01, 0x0a,
	0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70,
	0x72, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67,
	0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x63, 0x0a,
	0x0b, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x70,
	0x69, 0x63, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x14, 0x0a, 0x05,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x22, 0x7e, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x42, 0x6a, 0x0a, 0x1f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x69, 0x6f, 0x74, 0x2e, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67,
	0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x6f, 0x74, 0x2f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDescOnce sync.Once
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData = file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc
)

func file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData)
	})
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData
}

var file_yandex_cloud_iot_devices_v1_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_iot_devices_v1_registry_proto_goTypes = []interface{}{
	(Registry_Status)(0),          // 0: yandex.cloud.iot.devices.v1.Registry.Status
	(*Registry)(nil),              // 1: yandex.cloud.iot.devices.v1.Registry
	(*RegistryCertificate)(nil),   // 2: yandex.cloud.iot.devices.v1.RegistryCertificate
	(*DeviceAlias)(nil),           // 3: yandex.cloud.iot.devices.v1.DeviceAlias
	(*RegistryPassword)(nil),      // 4: yandex.cloud.iot.devices.v1.RegistryPassword
	nil,                           // 5: yandex.cloud.iot.devices.v1.Registry.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
}
var file_yandex_cloud_iot_devices_v1_registry_proto_depIdxs = []int32{
	6, // 0: yandex.cloud.iot.devices.v1.Registry.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: yandex.cloud.iot.devices.v1.Registry.labels:type_name -> yandex.cloud.iot.devices.v1.Registry.LabelsEntry
	0, // 2: yandex.cloud.iot.devices.v1.Registry.status:type_name -> yandex.cloud.iot.devices.v1.Registry.Status
	6, // 3: yandex.cloud.iot.devices.v1.RegistryCertificate.created_at:type_name -> google.protobuf.Timestamp
	6, // 4: yandex.cloud.iot.devices.v1.RegistryPassword.created_at:type_name -> google.protobuf.Timestamp
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iot_devices_v1_registry_proto_init() }
func file_yandex_cloud_iot_devices_v1_registry_proto_init() {
	if File_yandex_cloud_iot_devices_v1_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Registry); i {
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
		file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryCertificate); i {
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
		file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeviceAlias); i {
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
		file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryPassword); i {
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
			RawDescriptor: file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iot_devices_v1_registry_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iot_devices_v1_registry_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_iot_devices_v1_registry_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iot_devices_v1_registry_proto = out.File
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc = nil
	file_yandex_cloud_iot_devices_v1_registry_proto_goTypes = nil
	file_yandex_cloud_iot_devices_v1_registry_proto_depIdxs = nil
}
