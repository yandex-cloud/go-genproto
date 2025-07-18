// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/iot/broker/v1/broker.proto

package broker

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

type Broker_Status int32

const (
	Broker_STATUS_UNSPECIFIED Broker_Status = 0
	// Broker is being created.
	Broker_CREATING Broker_Status = 1
	// Broker is ready to use.
	Broker_ACTIVE Broker_Status = 2
	// Broker is being deleted.
	Broker_DELETING Broker_Status = 3
)

// Enum value maps for Broker_Status.
var (
	Broker_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "CREATING",
		2: "ACTIVE",
		3: "DELETING",
	}
	Broker_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"CREATING":           1,
		"ACTIVE":             2,
		"DELETING":           3,
	}
)

func (x Broker_Status) Enum() *Broker_Status {
	p := new(Broker_Status)
	*p = x
	return p
}

func (x Broker_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Broker_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_iot_broker_v1_broker_proto_enumTypes[0].Descriptor()
}

func (Broker_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_iot_broker_v1_broker_proto_enumTypes[0]
}

func (x Broker_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Broker_Status.Descriptor instead.
func (Broker_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_broker_v1_broker_proto_rawDescGZIP(), []int{0, 0}
}

// A broker.
type Broker struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the broker.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the broker belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the broker. The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the broker. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `key:value` pairs. Maximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Status of the broker.
	Status Broker_Status `protobuf:"varint,7,opt,name=status,proto3,enum=yandex.cloud.iot.broker.v1.Broker_Status" json:"status,omitempty"`
	// Options for logging broker events
	LogOptions    *LogOptions `protobuf:"bytes,8,opt,name=log_options,json=logOptions,proto3" json:"log_options,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Broker) Reset() {
	*x = Broker{}
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Broker) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Broker) ProtoMessage() {}

func (x *Broker) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Broker.ProtoReflect.Descriptor instead.
func (*Broker) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_broker_v1_broker_proto_rawDescGZIP(), []int{0}
}

func (x *Broker) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Broker) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Broker) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Broker) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Broker) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Broker) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Broker) GetStatus() Broker_Status {
	if x != nil {
		return x.Status
	}
	return Broker_STATUS_UNSPECIFIED
}

func (x *Broker) GetLogOptions() *LogOptions {
	if x != nil {
		return x.LogOptions
	}
	return nil
}

// A broker certificate.
type BrokerCertificate struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the broker that the certificate belongs to.
	BrokerId string `protobuf:"bytes,1,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	// SHA256 hash of the certificates.
	Fingerprint string `protobuf:"bytes,2,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Public part of the certificate.
	CertificateData string `protobuf:"bytes,3,opt,name=certificate_data,json=certificateData,proto3" json:"certificate_data,omitempty"`
	// Creation timestamp.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrokerCertificate) Reset() {
	*x = BrokerCertificate{}
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrokerCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerCertificate) ProtoMessage() {}

func (x *BrokerCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerCertificate.ProtoReflect.Descriptor instead.
func (*BrokerCertificate) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_broker_v1_broker_proto_rawDescGZIP(), []int{1}
}

func (x *BrokerCertificate) GetBrokerId() string {
	if x != nil {
		return x.BrokerId
	}
	return ""
}

func (x *BrokerCertificate) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *BrokerCertificate) GetCertificateData() string {
	if x != nil {
		return x.CertificateData
	}
	return ""
}

func (x *BrokerCertificate) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// A broker password.
type BrokerPassword struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the broker that the password belongs to.
	BrokerId string `protobuf:"bytes,1,opt,name=broker_id,json=brokerId,proto3" json:"broker_id,omitempty"`
	// ID of the password.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Creation timestamp.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BrokerPassword) Reset() {
	*x = BrokerPassword{}
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BrokerPassword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BrokerPassword) ProtoMessage() {}

func (x *BrokerPassword) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BrokerPassword.ProtoReflect.Descriptor instead.
func (*BrokerPassword) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_iot_broker_v1_broker_proto_rawDescGZIP(), []int{2}
}

func (x *BrokerPassword) GetBrokerId() string {
	if x != nil {
		return x.BrokerId
	}
	return ""
}

func (x *BrokerPassword) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BrokerPassword) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type LogOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Is logging from broker disabled.
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
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogOptions) ProtoMessage() {}

func (x *LogOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[3]
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
	return file_yandex_cloud_iot_broker_v1_broker_proto_rawDescGZIP(), []int{3}
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

var File_yandex_cloud_iot_broker_v1_broker_proto protoreflect.FileDescriptor

const file_yandex_cloud_iot_broker_v1_broker_proto_rawDesc = "" +
	"\n" +
	"'yandex/cloud/iot/broker/v1/broker.proto\x12\x1ayandex.cloud.iot.broker.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a'yandex/cloud/logging/v1/log_entry.proto\x1a\x1dyandex/cloud/validation.proto\"\xff\x03\n" +
	"\x06Broker\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1b\n" +
	"\tfolder_id\x18\x02 \x01(\tR\bfolderId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x12\n" +
	"\x04name\x18\x04 \x01(\tR\x04name\x12 \n" +
	"\vdescription\x18\x05 \x01(\tR\vdescription\x12F\n" +
	"\x06labels\x18\x06 \x03(\v2..yandex.cloud.iot.broker.v1.Broker.LabelsEntryR\x06labels\x12A\n" +
	"\x06status\x18\a \x01(\x0e2).yandex.cloud.iot.broker.v1.Broker.StatusR\x06status\x12G\n" +
	"\vlog_options\x18\b \x01(\v2&.yandex.cloud.iot.broker.v1.LogOptionsR\n" +
	"logOptions\x1a9\n" +
	"\vLabelsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value:\x028\x01\"H\n" +
	"\x06Status\x12\x16\n" +
	"\x12STATUS_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bCREATING\x10\x01\x12\n" +
	"\n" +
	"\x06ACTIVE\x10\x02\x12\f\n" +
	"\bDELETING\x10\x03\"\xb8\x01\n" +
	"\x11BrokerCertificate\x12\x1b\n" +
	"\tbroker_id\x18\x01 \x01(\tR\bbrokerId\x12 \n" +
	"\vfingerprint\x18\x02 \x01(\tR\vfingerprint\x12)\n" +
	"\x10certificate_data\x18\x03 \x01(\tR\x0fcertificateData\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"x\n" +
	"\x0eBrokerPassword\x12\x1b\n" +
	"\tbroker_id\x18\x01 \x01(\tR\bbrokerId\x12\x0e\n" +
	"\x02id\x18\x02 \x01(\tR\x02id\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"\x8a\x02\n" +
	"\n" +
	"LogOptions\x12\x1a\n" +
	"\bdisabled\x18\x01 \x01(\bR\bdisabled\x12G\n" +
	"\flog_group_id\x18\x02 \x01(\tB#\xf2\xc71\x1f([a-zA-Z][-a-zA-Z0-9_.]{0,63})?H\x00R\n" +
	"logGroupId\x12B\n" +
	"\tfolder_id\x18\x03 \x01(\tB#\xf2\xc71\x1f([a-zA-Z][-a-zA-Z0-9_.]{0,63})?H\x00R\bfolderId\x12D\n" +
	"\tmin_level\x18\x04 \x01(\x0e2'.yandex.cloud.logging.v1.LogLevel.LevelR\bminLevelB\r\n" +
	"\vdestinationBg\n" +
	"\x1eyandex.cloud.api.iot.broker.v1ZEgithub.com/yandex-cloud/go-genproto/yandex/cloud/iot/broker/v1;brokerb\x06proto3"

var (
	file_yandex_cloud_iot_broker_v1_broker_proto_rawDescOnce sync.Once
	file_yandex_cloud_iot_broker_v1_broker_proto_rawDescData []byte
)

func file_yandex_cloud_iot_broker_v1_broker_proto_rawDescGZIP() []byte {
	file_yandex_cloud_iot_broker_v1_broker_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_iot_broker_v1_broker_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_iot_broker_v1_broker_proto_rawDesc), len(file_yandex_cloud_iot_broker_v1_broker_proto_rawDesc)))
	})
	return file_yandex_cloud_iot_broker_v1_broker_proto_rawDescData
}

var file_yandex_cloud_iot_broker_v1_broker_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_iot_broker_v1_broker_proto_goTypes = []any{
	(Broker_Status)(0),            // 0: yandex.cloud.iot.broker.v1.Broker.Status
	(*Broker)(nil),                // 1: yandex.cloud.iot.broker.v1.Broker
	(*BrokerCertificate)(nil),     // 2: yandex.cloud.iot.broker.v1.BrokerCertificate
	(*BrokerPassword)(nil),        // 3: yandex.cloud.iot.broker.v1.BrokerPassword
	(*LogOptions)(nil),            // 4: yandex.cloud.iot.broker.v1.LogOptions
	nil,                           // 5: yandex.cloud.iot.broker.v1.Broker.LabelsEntry
	(*timestamppb.Timestamp)(nil), // 6: google.protobuf.Timestamp
	(v1.LogLevel_Level)(0),        // 7: yandex.cloud.logging.v1.LogLevel.Level
}
var file_yandex_cloud_iot_broker_v1_broker_proto_depIdxs = []int32{
	6, // 0: yandex.cloud.iot.broker.v1.Broker.created_at:type_name -> google.protobuf.Timestamp
	5, // 1: yandex.cloud.iot.broker.v1.Broker.labels:type_name -> yandex.cloud.iot.broker.v1.Broker.LabelsEntry
	0, // 2: yandex.cloud.iot.broker.v1.Broker.status:type_name -> yandex.cloud.iot.broker.v1.Broker.Status
	4, // 3: yandex.cloud.iot.broker.v1.Broker.log_options:type_name -> yandex.cloud.iot.broker.v1.LogOptions
	6, // 4: yandex.cloud.iot.broker.v1.BrokerCertificate.created_at:type_name -> google.protobuf.Timestamp
	6, // 5: yandex.cloud.iot.broker.v1.BrokerPassword.created_at:type_name -> google.protobuf.Timestamp
	7, // 6: yandex.cloud.iot.broker.v1.LogOptions.min_level:type_name -> yandex.cloud.logging.v1.LogLevel.Level
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_iot_broker_v1_broker_proto_init() }
func file_yandex_cloud_iot_broker_v1_broker_proto_init() {
	if File_yandex_cloud_iot_broker_v1_broker_proto != nil {
		return
	}
	file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes[3].OneofWrappers = []any{
		(*LogOptions_LogGroupId)(nil),
		(*LogOptions_FolderId)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_iot_broker_v1_broker_proto_rawDesc), len(file_yandex_cloud_iot_broker_v1_broker_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_iot_broker_v1_broker_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_iot_broker_v1_broker_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_iot_broker_v1_broker_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_iot_broker_v1_broker_proto_msgTypes,
	}.Build()
	File_yandex_cloud_iot_broker_v1_broker_proto = out.File
	file_yandex_cloud_iot_broker_v1_broker_proto_goTypes = nil
	file_yandex_cloud_iot_broker_v1_broker_proto_depIdxs = nil
}
