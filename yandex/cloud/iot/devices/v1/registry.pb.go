// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/iot/devices/v1/registry.proto

package devices

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/logging/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Registry_Status int32

const (
	Registry_STATUS_UNSPECIFIED Registry_Status = 0
	// Registry is being created.
	Registry_CREATING Registry_Status = 1
	// Registry is ready to use.
	Registry_ACTIVE Registry_Status = 2
	// Registry is being deleted.
	Registry_DELETING Registry_Status = 3
	// Registry is disabled.
	Registry_DISABLED Registry_Status = 4
)

// Enum value maps for Registry_Status.
var (
	Registry_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
		4: "DISABLED",
	}
	Registry_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"DELETING":           3,
		"DISABLED":           4,
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
	state protoimpl.MessageState `protogen:"open.v1"`
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
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Status of the registry.
	Status Registry_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.iot.devices.v1.Registry_Status" json:"status,omitempty"`
	// ID of the logs group for the specified registry.
	LogGroupId string `protobuf:"bytes,8,opt,name=log_group_id,json=logGroupId,proto3" json:"log_group_id,omitempty"`
	// Options for logging registry events
	LogOptions    *LogOptions `protobuf:"bytes,9,opt,name=log_options,json=logOptions,proto3" json:"log_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Registry) Reset() {
	*x = Registry{}
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Registry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Registry) ProtoMessage() {}

func (x *Registry) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[0]
	if x != nil {
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

func (x *Registry) GetLogOptions() *LogOptions {
	if x != nil {
		return x.LogOptions
	}
	return nil
}

// A registry certificate. For more information, see [Managing registry certificates](/docs/iot-core/operations/certificates/registry-certificates).
type RegistryCertificate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the registry that the certificate belongs to.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// SHA256 hash of the certificates.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Public part of the certificate.
	CertificateData string `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	// Creation timestamp.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryCertificate) Reset() {
	*x = RegistryCertificate{}
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryCertificate) ProtoMessage() {}

func (x *RegistryCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[1]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the device that the alias belongs to.
	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	// Prefix of a canonical topic name to be aliased, e.g. `$devices/abcdef`.
	TopicPrefix string `protobuf:"bytes,2,opt,name=topic_prefix,json=topicPrefix,proto3" json:"topic_prefix,omitempty"`
	// Alias of a device topic.
	Alias         string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeviceAlias) Reset() {
	*x = DeviceAlias{}
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeviceAlias) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeviceAlias) ProtoMessage() {}

func (x *DeviceAlias) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[2]
	if x != nil {
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
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the registry that the password belongs to.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// ID of the password.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Creation timestamp.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RegistryPassword) Reset() {
	*x = RegistryPassword{}
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegistryPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryPassword) ProtoMessage() {}

func (x *RegistryPassword) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[3]
	if x != nil {
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

// A Yandex Data Streams export.
type DataStreamExport struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the YDS export.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the YDS export.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the registry that the YDS export belongs to.
	RegistryId string `protobuf:"bytes,3,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// MQTT topic whose messages export to YDS.
	MqttTopicFilter string `protobuf:"bytes,4,opt,name=mqtt_topic_filter,json=mqttTopicFilter,proto3" json:"mqtt_topic_filter,omitempty"`
	// YDS database.
	Database string `protobuf:"bytes,5,opt,name=database,proto3" json:"database,omitempty"`
	// YDS stream name.
	Stream string `protobuf:"bytes,6,opt,name=stream,proto3" json:"stream,omitempty"`
	// ID of the service account which has permission to write to data stream.
	ServiceAccountId string `protobuf:"bytes,7,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// Creation timestamp.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataStreamExport) Reset() {
	*x = DataStreamExport{}
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataStreamExport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStreamExport) ProtoMessage() {}

func (x *DataStreamExport) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStreamExport.ProtoReflect.Descriptor instead.
func (*DataStreamExport) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{4}
}

func (x *DataStreamExport) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataStreamExport) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DataStreamExport) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *DataStreamExport) GetMqttTopicFilter() string {
	if x != nil {
		return x.MqttTopicFilter
	}
	return ""
}

func (x *DataStreamExport) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *DataStreamExport) GetStream() string {
	if x != nil {
		return x.Stream
	}
	return ""
}

func (x *DataStreamExport) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *DataStreamExport) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type LogOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Is logging from registry disabled.
	Disabled bool `protobuf:"varint,1,opt,name=disabled,proto3" json:"disabled,omitempty"`
	// Log entries destination.
	//
	// Types that are valid to be assigned to Destination:
	//
	//	*LogOptions_LogGroupId
	//	*LogOptions_FolderId
	Destination isLogOptions_Destination `protobuf_oneof:"destination"`
	// Minimum log entry level.
	//
	// See [LogLevel.Level] for details.
	MinLevel      v1.LogLevel_Level `protobuf:"varint,4,opt,name=min_level,json=minLevel,proto3,enum=yandex.cloud.logging.v1.LogLevel_Level" json:"min_level,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LogOptions) Reset() {
	*x = LogOptions{}
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogOptions) ProtoMessage() {}

func (x *LogOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogOptions.ProtoReflect.Descriptor instead.
func (*LogOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP(), []int{5}
}

func (x *LogOptions) GetDisabled() bool {
	if x != nil {
		return x.Disabled
	}
	return false
}

func (x *LogOptions) GetDestination() isLogOptions_Destination {
	if x != nil {
		return x.Destination
	}
	return nil
}

func (x *LogOptions) GetLogGroupId() string {
	if x != nil {
		if x, ok := x.Destination.(*LogOptions_LogGroupId); ok {
			return x.LogGroupId
		}
	}
	return ""
}

func (x *LogOptions) GetFolderId() string {
	if x != nil {
		if x, ok := x.Destination.(*LogOptions_FolderId); ok {
			return x.FolderId
		}
	}
	return ""
}

func (x *LogOptions) GetMinLevel() v1.LogLevel_Level {
	if x != nil {
		return x.MinLevel
	}
	return v1.LogLevel_Level(0)
}

type isLogOptions_Destination interface {
	isLogOptions_Destination()
}

type LogOptions_LogGroupId struct {
	// Entry should be written to log group resolved by ID.
	LogGroupId string `protobuf:"bytes,2,opt,name=log_group_id,json=logGroupId,proto3,oneof"`
}

type LogOptions_FolderId struct {
	// Entry should be written to default log group for specified folder.
	FolderId string `protobuf:"bytes,3,opt,name=folder_id,json=folderId,proto3,oneof"`
}

func (*LogOptions_LogGroupId) isLogOptions_Destination() {}

func (*LogOptions_FolderId) isLogOptions_Destination() {}

var File_yandex_cloud_iot_devices_v1_registry_proto protoreflect.FileDescriptor

const file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc = "" +
	"\n" +
	"*yandex/cloud/iot/devices/v1/registry.proto\x12\x1byandex.cloud.iot.devices.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a'yandex/cloud/logging/v1/log_entry.proto\x1a\x1dyandex/cloud/validation.proto\"\xb8\x04\n" +
	"\bRegistry\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12I\n" +
	"\x06labels\x18\x06 \x03(\v21.yandex.cloud.iot.devices.v1.Registry.LabelsEntryR\x06labels\x12D\n" +
	"\x06status\x18\a \x01(\x0e2,.yandex.cloud.iot.devices.v1.Registry.StatusR\x06status\x12 \n" +
	"\flog_group_id\x18\b \x01(\tR\n" +
	"logGroupId\x12H\n" +
	"\vlog_options\x18\t \x01(\v2'.yandex.cloud.iot.devices.v1.LogOptionsR\n" +
	"logOptions\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"V\n" +
	"\x06Status\x12\x16\n" +
	"\x12STATUS_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bCREATING\x10\x01\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\f\n" +
	"\bDELETING\x10\x03\x12\f\n" +
	"\bDISABLED\x10\x04\"\xbe\x01\n" +
	"\x13RegistryCertificate\x12\x1f\n" +
	"\vregistry_id\x18\x01 \x01(\tR\n" +
	"registryId\x12 \n" +
	"\vfingerprint\x18\x02 \x01(\tR\vfingerprint\x12)\n" +
	"\x10certificate_data\x18\x03 \x01(\tR\x0fcertificateData\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"c\n" +
	"\vDeviceAlias\x12\x1b\n" +
	"\tdevice_id\x18\x01 \x01(\tR\bdeviceId\x12!\n" +
	"\ftopic_prefix\x18\x02 \x01(\tR\vtopicPrefix\x12\x14\n" +
	"\x05alias\x18\x03 \x01(\tR\x05alias\"~\n" +
	"\x10RegistryPassword\x12\x1f\n" +
	"\vregistry_id\x18\x01 \x01(\tR\n" +
	"registryId\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"\xa0\x02\n" +
	"\x10DataStreamExport\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x1f\n" +
	"\vregistry_id\x18\x03 \x01(\tR\n" +
	"registryId\x12*\n" +
	"\x11mqtt_topic_filter\x18\x04 \x01(\tR\x0fmqttTopicFilter\x12\x1a\n" +
	"\bdatabase\x18\x05 \x01(\tR\bdatabase\x12\x16\n" +
	"\x06stream\x18\x06 \x01(\tR\x06stream\x12,\n" +
	"\x12service_account_id\x18\a \x01(\tR\x10serviceAccountId\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"\x8a\x02\n" +
	"\n" +
	"LogOptions\x12\x1a\n" +
	"\bdisabled\x18\x01 \x01(\bR\bdisabled\x12G\n" +
	"\flog_group_id\x18\x02 \x01(\tB#\xf2\xc71\x1f([a-zA-Z][-a-zA-Z0-9_.]{0,63})?H\x00R\n" +
	"logGroupId\x12B\n" +
	"\tfolder_id\x18\x03 \x01(\tB#\xf2\xc71\x1f([a-zA-Z][-a-zA-Z0-9_.]{0,63})?H\x00R\bfolderId\x12D\n" +
	"\tmin_level\x18\x04 \x01(\x0e2'.yandex.cloud.logging.v1.LogLevel.LevelR\bminLevelB\r\n" +
	"\vdestinationBj\n" +
	"\x1fyandex.cloud.api.iot.devices.v1ZGgithub.com/yandex-cloud/go-genproto/yandex/cloud/iot/devices/v1;devicesb\x06proto3"

var (
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDescOnce sync.Once
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData []byte
)

func file_yandex_cloud_iot_devices_v1_registry_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iot_devices_v1_registry_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc), len(file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc)))
	})
	return file_yandex_cloud_iot_devices_v1_registry_proto_rawDescData
}

var file_yandex_cloud_iot_devices_v1_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_iot_devices_v1_registry_proto_goTypes = []any{
	(Registry_Status)(0),          // 0: yandex.cloud.iot.devices.v1.Registry.Status
	(*Registry)(nil),              // 1: yandex.cloud.iot.devices.v1.Registry
	(*RegistryCertificate)(nil),   // 2: yandex.cloud.iot.devices.v1.RegistryCertificate
	(*DeviceAlias)(nil),           // 3: yandex.cloud.iot.devices.v1.DeviceAlias
	(*RegistryPassword)(nil),      // 4: yandex.cloud.iot.devices.v1.RegistryPassword
	(*DataStreamExport)(nil),      // 5: yandex.cloud.iot.devices.v1.DataStreamExport
	(*LogOptions)(nil),            // 6: yandex.cloud.iot.devices.v1.LogOptions
	nil,                           // 7: yandex.cloud.iot.devices.v1.Registry.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(v1.LogLevel_Level)(0),        // 9: yandex.cloud.logging.v1.LogLevel.Level
}
var file_yandex_cloud_iot_devices_v1_registry_proto_depIdxs = []int32{
	8, // 0: yandex.cloud.iot.devices.v1.Registry.created_at:type_name -> google.protobuf.Timestamp
	7, // 1: yandex.cloud.iot.devices.v1.Registry.labels:type_name -> yandex.cloud.iot.devices.v1.Registry.LabelsEntry
	0, // 2: yandex.cloud.iot.devices.v1.Registry.status:type_name -> yandex.cloud.iot.devices.v1.Registry.Status
	6, // 3: yandex.cloud.iot.devices.v1.Registry.log_options:type_name -> yandex.cloud.iot.devices.v1.LogOptions
	8, // 4: yandex.cloud.iot.devices.v1.RegistryCertificate.created_at:type_name -> google.protobuf.Timestamp
	8, // 5: yandex.cloud.iot.devices.v1.RegistryPassword.created_at:type_name -> google.protobuf.Timestamp
	8, // 6: yandex.cloud.iot.devices.v1.DataStreamExport.created_at:type_name -> google.protobuf.Timestamp
	9, // 7: yandex.cloud.iot.devices.v1.LogOptions.min_level:type_name -> yandex.cloud.logging.v1.LogLevel.Level
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iot_devices_v1_registry_proto_init() }
func file_yandex_cloud_iot_devices_v1_registry_proto_init() {
	if File_yandex_cloud_iot_devices_v1_registry_proto != nil {
		return
	}
	file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes[5].OneofWrappers = []any{
		(*LogOptions_LogGroupId)(nil),
		(*LogOptions_FolderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc), len(file_yandex_cloud_iot_devices_v1_registry_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iot_devices_v1_registry_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iot_devices_v1_registry_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_iot_devices_v1_registry_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_iot_devices_v1_registry_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iot_devices_v1_registry_proto = out.File
	file_yandex_cloud_iot_devices_v1_registry_proto_goTypes = nil
	file_yandex_cloud_iot_devices_v1_registry_proto_depIdxs = nil
}
