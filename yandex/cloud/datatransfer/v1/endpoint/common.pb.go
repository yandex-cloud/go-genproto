// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datatransfer/v1/endpoint/common.proto

package endpoint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type ObjectTransferStage int32

const (
	ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED ObjectTransferStage = 0
	// Before data transfer
	ObjectTransferStage_BEFORE_DATA ObjectTransferStage = 1
	// After data transfer
	ObjectTransferStage_AFTER_DATA ObjectTransferStage = 2
	// Don't copy
	ObjectTransferStage_NEVER ObjectTransferStage = 3
)

// Enum value maps for ObjectTransferStage.
var (
	ObjectTransferStage_name = map[int32]string{
		0: "OBJECT_TRANSFER_STAGE_UNSPECIFIED",
		1: "BEFORE_DATA",
		2: "AFTER_DATA",
		3: "NEVER",
	}
	ObjectTransferStage_value = map[string]int32{
		"OBJECT_TRANSFER_STAGE_UNSPECIFIED": 0,
		"BEFORE_DATA":                       1,
		"AFTER_DATA":                        2,
		"NEVER":                             3,
	}
)

func (x ObjectTransferStage) Enum() *ObjectTransferStage {
	p := new(ObjectTransferStage)
	*p = x
	return p
}

func (x ObjectTransferStage) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ObjectTransferStage) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes[0].Descriptor()
}

func (ObjectTransferStage) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes[0]
}

func (x ObjectTransferStage) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ObjectTransferStage.Descriptor instead.
func (ObjectTransferStage) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{0}
}

type CleanupPolicy int32

const (
	CleanupPolicy_CLEANUP_POLICY_UNSPECIFIED CleanupPolicy = 0
	// Don't cleanup
	CleanupPolicy_DISABLED CleanupPolicy = 1
	// Drop
	CleanupPolicy_DROP CleanupPolicy = 2
	// Truncate
	CleanupPolicy_TRUNCATE CleanupPolicy = 3
)

// Enum value maps for CleanupPolicy.
var (
	CleanupPolicy_name = map[int32]string{
		0: "CLEANUP_POLICY_UNSPECIFIED",
		1: "DISABLED",
		2: "DROP",
		3: "TRUNCATE",
	}
	CleanupPolicy_value = map[string]int32{
		"CLEANUP_POLICY_UNSPECIFIED": 0,
		"DISABLED":                   1,
		"DROP":                       2,
		"TRUNCATE":                   3,
	}
)

func (x CleanupPolicy) Enum() *CleanupPolicy {
	p := new(CleanupPolicy)
	*p = x
	return p
}

func (x CleanupPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CleanupPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes[1].Descriptor()
}

func (CleanupPolicy) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes[1]
}

func (x CleanupPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CleanupPolicy.Descriptor instead.
func (CleanupPolicy) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{1}
}

type ColumnType int32

const (
	ColumnType_COLUMN_TYPE_UNSPECIFIED ColumnType = 0
	ColumnType_INT32                   ColumnType = 1
	ColumnType_INT16                   ColumnType = 2
	ColumnType_INT8                    ColumnType = 3
	ColumnType_UINT64                  ColumnType = 4
	ColumnType_UINT32                  ColumnType = 5
	ColumnType_UINT16                  ColumnType = 6
	ColumnType_UINT8                   ColumnType = 7
	ColumnType_DOUBLE                  ColumnType = 8
	ColumnType_BOOLEAN                 ColumnType = 9
	ColumnType_STRING                  ColumnType = 10
	ColumnType_UTF8                    ColumnType = 11
	ColumnType_ANY                     ColumnType = 12
	ColumnType_DATETIME                ColumnType = 13
	ColumnType_INT64                   ColumnType = 14
)

// Enum value maps for ColumnType.
var (
	ColumnType_name = map[int32]string{
		0:  "COLUMN_TYPE_UNSPECIFIED",
		1:  "INT32",
		2:  "INT16",
		3:  "INT8",
		4:  "UINT64",
		5:  "UINT32",
		6:  "UINT16",
		7:  "UINT8",
		8:  "DOUBLE",
		9:  "BOOLEAN",
		10: "STRING",
		11: "UTF8",
		12: "ANY",
		13: "DATETIME",
		14: "INT64",
	}
	ColumnType_value = map[string]int32{
		"COLUMN_TYPE_UNSPECIFIED": 0,
		"INT32":                   1,
		"INT16":                   2,
		"INT8":                    3,
		"UINT64":                  4,
		"UINT32":                  5,
		"UINT16":                  6,
		"UINT8":                   7,
		"DOUBLE":                  8,
		"BOOLEAN":                 9,
		"STRING":                  10,
		"UTF8":                    11,
		"ANY":                     12,
		"DATETIME":                13,
		"INT64":                   14,
	}
)

func (x ColumnType) Enum() *ColumnType {
	p := new(ColumnType)
	*p = x
	return p
}

func (x ColumnType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ColumnType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes[2].Descriptor()
}

func (ColumnType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes[2]
}

func (x ColumnType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ColumnType.Descriptor instead.
func (ColumnType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{2}
}

type AltName struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Source table name
	FromName string `protobuf:"bytes,1,opt,name=from_name,json=fromName,proto3" json:"from_name,omitempty"`
	// Target table name
	ToName        string `protobuf:"bytes,2,opt,name=to_name,json=toName,proto3" json:"to_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AltName) Reset() {
	*x = AltName{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AltName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AltName) ProtoMessage() {}

func (x *AltName) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AltName.ProtoReflect.Descriptor instead.
func (*AltName) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{0}
}

func (x *AltName) GetFromName() string {
	if x != nil {
		return x.FromName
	}
	return ""
}

func (x *AltName) GetToName() string {
	if x != nil {
		return x.ToName
	}
	return ""
}

type Secret struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*Secret_Raw
	Value         isSecret_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Secret) Reset() {
	*x = Secret{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Secret) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Secret) ProtoMessage() {}

func (x *Secret) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Secret.ProtoReflect.Descriptor instead.
func (*Secret) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{1}
}

func (x *Secret) GetValue() isSecret_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *Secret) GetRaw() string {
	if x != nil {
		if x, ok := x.Value.(*Secret_Raw); ok {
			return x.Raw
		}
	}
	return ""
}

type isSecret_Value interface {
	isSecret_Value()
}

type Secret_Raw struct {
	// Raw secret value
	Raw string `protobuf:"bytes,1,opt,name=raw,proto3,oneof"`
}

func (*Secret_Raw) isSecret_Value() {}

type ColSchema struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          ColumnType             `protobuf:"varint,2,opt,name=type,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ColumnType" json:"type,omitempty"`
	Key           bool                   `protobuf:"varint,3,opt,name=key,proto3" json:"key,omitempty"`
	Required      bool                   `protobuf:"varint,4,opt,name=required,proto3" json:"required,omitempty"`
	Path          string                 `protobuf:"bytes,5,opt,name=path,proto3" json:"path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColSchema) Reset() {
	*x = ColSchema{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColSchema) ProtoMessage() {}

func (x *ColSchema) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColSchema.ProtoReflect.Descriptor instead.
func (*ColSchema) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{2}
}

func (x *ColSchema) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ColSchema) GetType() ColumnType {
	if x != nil {
		return x.Type
	}
	return ColumnType_COLUMN_TYPE_UNSPECIFIED
}

func (x *ColSchema) GetKey() bool {
	if x != nil {
		return x.Key
	}
	return false
}

func (x *ColSchema) GetRequired() bool {
	if x != nil {
		return x.Required
	}
	return false
}

func (x *ColSchema) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type TLSMode struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to TlsMode:
	//
	//	*TLSMode_Disabled
	//	*TLSMode_Enabled
	TlsMode       isTLSMode_TlsMode `protobuf_oneof:"tls_mode"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLSMode) Reset() {
	*x = TLSMode{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLSMode) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSMode) ProtoMessage() {}

func (x *TLSMode) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSMode.ProtoReflect.Descriptor instead.
func (*TLSMode) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{3}
}

func (x *TLSMode) GetTlsMode() isTLSMode_TlsMode {
	if x != nil {
		return x.TlsMode
	}
	return nil
}

func (x *TLSMode) GetDisabled() *emptypb.Empty {
	if x != nil {
		if x, ok := x.TlsMode.(*TLSMode_Disabled); ok {
			return x.Disabled
		}
	}
	return nil
}

func (x *TLSMode) GetEnabled() *TLSConfig {
	if x != nil {
		if x, ok := x.TlsMode.(*TLSMode_Enabled); ok {
			return x.Enabled
		}
	}
	return nil
}

type isTLSMode_TlsMode interface {
	isTLSMode_TlsMode()
}

type TLSMode_Disabled struct {
	Disabled *emptypb.Empty `protobuf:"bytes,1,opt,name=disabled,proto3,oneof"`
}

type TLSMode_Enabled struct {
	Enabled *TLSConfig `protobuf:"bytes,2,opt,name=enabled,proto3,oneof"`
}

func (*TLSMode_Disabled) isTLSMode_TlsMode() {}

func (*TLSMode_Enabled) isTLSMode_TlsMode() {}

type TLSConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// CA certificate
	//
	// X.509 certificate of the certificate authority which issued the server's
	// certificate, in PEM format. When CA certificate is specified TLS is used to
	// connect to the server.
	CaCertificate string `protobuf:"bytes,1,opt,name=ca_certificate,json=caCertificate,proto3" json:"ca_certificate,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TLSConfig) Reset() {
	*x = TLSConfig{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TLSConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TLSConfig) ProtoMessage() {}

func (x *TLSConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TLSConfig.ProtoReflect.Descriptor instead.
func (*TLSConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{4}
}

func (x *TLSConfig) GetCaCertificate() string {
	if x != nil {
		return x.CaCertificate
	}
	return ""
}

type ColumnValue struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Value:
	//
	//	*ColumnValue_StringValue
	Value         isColumnValue_Value `protobuf_oneof:"value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ColumnValue) Reset() {
	*x = ColumnValue{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ColumnValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ColumnValue) ProtoMessage() {}

func (x *ColumnValue) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ColumnValue.ProtoReflect.Descriptor instead.
func (*ColumnValue) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{5}
}

func (x *ColumnValue) GetValue() isColumnValue_Value {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *ColumnValue) GetStringValue() string {
	if x != nil {
		if x, ok := x.Value.(*ColumnValue_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

type isColumnValue_Value interface {
	isColumnValue_Value()
}

type ColumnValue_StringValue struct {
	StringValue string `protobuf:"bytes,1,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*ColumnValue_StringValue) isColumnValue_Value() {}

type DataTransformationOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Cloud function
	CloudFunction string `protobuf:"bytes,1,opt,name=cloud_function,json=cloudFunction,proto3" json:"cloud_function,omitempty"`
	// Number of retries
	NumberOfRetries int64 `protobuf:"varint,2,opt,name=number_of_retries,json=numberOfRetries,proto3" json:"number_of_retries,omitempty"`
	// Buffer size for function
	BufferSize string `protobuf:"bytes,3,opt,name=buffer_size,json=bufferSize,proto3" json:"buffer_size,omitempty"`
	// Flush interval
	BufferFlushInterval string `protobuf:"bytes,4,opt,name=buffer_flush_interval,json=bufferFlushInterval,proto3" json:"buffer_flush_interval,omitempty"`
	// Invocation timeout
	InvocationTimeout string `protobuf:"bytes,5,opt,name=invocation_timeout,json=invocationTimeout,proto3" json:"invocation_timeout,omitempty"`
	// Service account
	ServiceAccountId string `protobuf:"bytes,8,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *DataTransformationOptions) Reset() {
	*x = DataTransformationOptions{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataTransformationOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataTransformationOptions) ProtoMessage() {}

func (x *DataTransformationOptions) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataTransformationOptions.ProtoReflect.Descriptor instead.
func (*DataTransformationOptions) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{6}
}

func (x *DataTransformationOptions) GetCloudFunction() string {
	if x != nil {
		return x.CloudFunction
	}
	return ""
}

func (x *DataTransformationOptions) GetNumberOfRetries() int64 {
	if x != nil {
		return x.NumberOfRetries
	}
	return 0
}

func (x *DataTransformationOptions) GetBufferSize() string {
	if x != nil {
		return x.BufferSize
	}
	return ""
}

func (x *DataTransformationOptions) GetBufferFlushInterval() string {
	if x != nil {
		return x.BufferFlushInterval
	}
	return ""
}

func (x *DataTransformationOptions) GetInvocationTimeout() string {
	if x != nil {
		return x.InvocationTimeout
	}
	return ""
}

func (x *DataTransformationOptions) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

type FieldList struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Column schema
	Fields        []*ColSchema `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *FieldList) Reset() {
	*x = FieldList{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FieldList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FieldList) ProtoMessage() {}

func (x *FieldList) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FieldList.ProtoReflect.Descriptor instead.
func (*FieldList) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{7}
}

func (x *FieldList) GetFields() []*ColSchema {
	if x != nil {
		return x.Fields
	}
	return nil
}

type DataSchema struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Schema:
	//
	//	*DataSchema_JsonFields
	//	*DataSchema_Fields
	Schema        isDataSchema_Schema `protobuf_oneof:"schema"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DataSchema) Reset() {
	*x = DataSchema{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DataSchema) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSchema) ProtoMessage() {}

func (x *DataSchema) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSchema.ProtoReflect.Descriptor instead.
func (*DataSchema) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{8}
}

func (x *DataSchema) GetSchema() isDataSchema_Schema {
	if x != nil {
		return x.Schema
	}
	return nil
}

func (x *DataSchema) GetJsonFields() string {
	if x != nil {
		if x, ok := x.Schema.(*DataSchema_JsonFields); ok {
			return x.JsonFields
		}
	}
	return ""
}

func (x *DataSchema) GetFields() *FieldList {
	if x != nil {
		if x, ok := x.Schema.(*DataSchema_Fields); ok {
			return x.Fields
		}
	}
	return nil
}

type isDataSchema_Schema interface {
	isDataSchema_Schema()
}

type DataSchema_JsonFields struct {
	JsonFields string `protobuf:"bytes,1,opt,name=json_fields,json=jsonFields,proto3,oneof"`
}

type DataSchema_Fields struct {
	Fields *FieldList `protobuf:"bytes,2,opt,name=fields,proto3,oneof"`
}

func (*DataSchema_JsonFields) isDataSchema_Schema() {}

func (*DataSchema_Fields) isDataSchema_Schema() {}

// No authentication
type NoAuth struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NoAuth) Reset() {
	*x = NoAuth{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NoAuth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoAuth) ProtoMessage() {}

func (x *NoAuth) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoAuth.ProtoReflect.Descriptor instead.
func (*NoAuth) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{9}
}

type ConnectionManagerConnection struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ConnectionId  string                 `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectionManagerConnection) Reset() {
	*x = ConnectionManagerConnection{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectionManagerConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionManagerConnection) ProtoMessage() {}

func (x *ConnectionManagerConnection) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionManagerConnection.ProtoReflect.Descriptor instead.
func (*ConnectionManagerConnection) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP(), []int{10}
}

func (x *ConnectionManagerConnection) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

var File_yandex_cloud_datatransfer_v1_endpoint_common_proto protoreflect.FileDescriptor

const file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDesc = "" +
	"\n" +
	"2yandex/cloud/datatransfer/v1/endpoint/common.proto\x12%yandex.cloud.datatransfer.v1.endpoint\x1a\x1bgoogle/protobuf/empty.proto\"?\n" +
	"\aAltName\x12\x1b\n" +
	"\tfrom_name\x18\x01 \x01(\tR\bfromName\x12\x17\n" +
	"\ato_name\x18\x02 \x01(\tR\x06toName\"%\n" +
	"\x06Secret\x12\x12\n" +
	"\x03raw\x18\x01 \x01(\tH\x00R\x03rawB\a\n" +
	"\x05value\"\xa8\x01\n" +
	"\tColSchema\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12E\n" +
	"\x04type\x18\x02 \x01(\x0e21.yandex.cloud.datatransfer.v1.endpoint.ColumnTypeR\x04type\x12\x10\n" +
	"\x03key\x18\x03 \x01(\bR\x03key\x12\x1a\n" +
	"\brequired\x18\x04 \x01(\bR\brequired\x12\x12\n" +
	"\x04path\x18\x05 \x01(\tR\x04path\"\x99\x01\n" +
	"\aTLSMode\x124\n" +
	"\bdisabled\x18\x01 \x01(\v2\x16.google.protobuf.EmptyH\x00R\bdisabled\x12L\n" +
	"\aenabled\x18\x02 \x01(\v20.yandex.cloud.datatransfer.v1.endpoint.TLSConfigH\x00R\aenabledB\n" +
	"\n" +
	"\btls_mode\"2\n" +
	"\tTLSConfig\x12%\n" +
	"\x0eca_certificate\x18\x01 \x01(\tR\rcaCertificate\";\n" +
	"\vColumnValue\x12#\n" +
	"\fstring_value\x18\x01 \x01(\tH\x00R\vstringValueB\a\n" +
	"\x05value\"\xa6\x02\n" +
	"\x19DataTransformationOptions\x12%\n" +
	"\x0ecloud_function\x18\x01 \x01(\tR\rcloudFunction\x12*\n" +
	"\x11number_of_retries\x18\x02 \x01(\x03R\x0fnumberOfRetries\x12\x1f\n" +
	"\vbuffer_size\x18\x03 \x01(\tR\n" +
	"bufferSize\x122\n" +
	"\x15buffer_flush_interval\x18\x04 \x01(\tR\x13bufferFlushInterval\x12-\n" +
	"\x12invocation_timeout\x18\x05 \x01(\tR\x11invocationTimeout\x12,\n" +
	"\x12service_account_id\x18\b \x01(\tR\x10serviceAccountIdJ\x04\b\x06\x10\b\"[\n" +
	"\tFieldList\x12H\n" +
	"\x06fields\x18\x02 \x03(\v20.yandex.cloud.datatransfer.v1.endpoint.ColSchemaR\x06fieldsJ\x04\b\x01\x10\x02\"\x85\x01\n" +
	"\n" +
	"DataSchema\x12!\n" +
	"\vjson_fields\x18\x01 \x01(\tH\x00R\n" +
	"jsonFields\x12J\n" +
	"\x06fields\x18\x02 \x01(\v20.yandex.cloud.datatransfer.v1.endpoint.FieldListH\x00R\x06fieldsB\b\n" +
	"\x06schema\"\b\n" +
	"\x06NoAuth\"B\n" +
	"\x1bConnectionManagerConnection\x12#\n" +
	"\rconnection_id\x18\x01 \x01(\tR\fconnectionId*h\n" +
	"\x13ObjectTransferStage\x12%\n" +
	"!OBJECT_TRANSFER_STAGE_UNSPECIFIED\x10\x00\x12\x0f\n" +
	"\vBEFORE_DATA\x10\x01\x12\x0e\n" +
	"\n" +
	"AFTER_DATA\x10\x02\x12\t\n" +
	"\x05NEVER\x10\x03*U\n" +
	"\rCleanupPolicy\x12\x1e\n" +
	"\x1aCLEANUP_POLICY_UNSPECIFIED\x10\x00\x12\f\n" +
	"\bDISABLED\x10\x01\x12\b\n" +
	"\x04DROP\x10\x02\x12\f\n" +
	"\bTRUNCATE\x10\x03*\xc9\x01\n" +
	"\n" +
	"ColumnType\x12\x1b\n" +
	"\x17COLUMN_TYPE_UNSPECIFIED\x10\x00\x12\t\n" +
	"\x05INT32\x10\x01\x12\t\n" +
	"\x05INT16\x10\x02\x12\b\n" +
	"\x04INT8\x10\x03\x12\n" +
	"\n" +
	"\x06UINT64\x10\x04\x12\n" +
	"\n" +
	"\x06UINT32\x10\x05\x12\n" +
	"\n" +
	"\x06UINT16\x10\x06\x12\t\n" +
	"\x05UINT8\x10\a\x12\n" +
	"\n" +
	"\x06DOUBLE\x10\b\x12\v\n" +
	"\aBOOLEAN\x10\t\x12\n" +
	"\n" +
	"\x06STRING\x10\n" +
	"\x12\b\n" +
	"\x04UTF8\x10\v\x12\a\n" +
	"\x03ANY\x10\f\x12\f\n" +
	"\bDATETIME\x10\r\x12\t\n" +
	"\x05INT64\x10\x0eB\xa7\x01\n" +
	")yandex.cloud.api.datatransfer.v1.endpointZRgithub.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1/endpoint;endpoint\xaa\x02%Yandex.Cloud.Datatransfer.V1.EndPointb\x06proto3"

var (
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescData []byte
)

func file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDesc)))
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_yandex_cloud_datatransfer_v1_endpoint_common_proto_goTypes = []any{
	(ObjectTransferStage)(0),            // 0: yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	(CleanupPolicy)(0),                  // 1: yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy
	(ColumnType)(0),                     // 2: yandex.cloud.datatransfer.v1.endpoint.ColumnType
	(*AltName)(nil),                     // 3: yandex.cloud.datatransfer.v1.endpoint.AltName
	(*Secret)(nil),                      // 4: yandex.cloud.datatransfer.v1.endpoint.Secret
	(*ColSchema)(nil),                   // 5: yandex.cloud.datatransfer.v1.endpoint.ColSchema
	(*TLSMode)(nil),                     // 6: yandex.cloud.datatransfer.v1.endpoint.TLSMode
	(*TLSConfig)(nil),                   // 7: yandex.cloud.datatransfer.v1.endpoint.TLSConfig
	(*ColumnValue)(nil),                 // 8: yandex.cloud.datatransfer.v1.endpoint.ColumnValue
	(*DataTransformationOptions)(nil),   // 9: yandex.cloud.datatransfer.v1.endpoint.DataTransformationOptions
	(*FieldList)(nil),                   // 10: yandex.cloud.datatransfer.v1.endpoint.FieldList
	(*DataSchema)(nil),                  // 11: yandex.cloud.datatransfer.v1.endpoint.DataSchema
	(*NoAuth)(nil),                      // 12: yandex.cloud.datatransfer.v1.endpoint.NoAuth
	(*ConnectionManagerConnection)(nil), // 13: yandex.cloud.datatransfer.v1.endpoint.ConnectionManagerConnection
	(*emptypb.Empty)(nil),               // 14: google.protobuf.Empty
}
var file_yandex_cloud_datatransfer_v1_endpoint_common_proto_depIdxs = []int32{
	2,  // 0: yandex.cloud.datatransfer.v1.endpoint.ColSchema.type:type_name -> yandex.cloud.datatransfer.v1.endpoint.ColumnType
	14, // 1: yandex.cloud.datatransfer.v1.endpoint.TLSMode.disabled:type_name -> google.protobuf.Empty
	7,  // 2: yandex.cloud.datatransfer.v1.endpoint.TLSMode.enabled:type_name -> yandex.cloud.datatransfer.v1.endpoint.TLSConfig
	5,  // 3: yandex.cloud.datatransfer.v1.endpoint.FieldList.fields:type_name -> yandex.cloud.datatransfer.v1.endpoint.ColSchema
	10, // 4: yandex.cloud.datatransfer.v1.endpoint.DataSchema.fields:type_name -> yandex.cloud.datatransfer.v1.endpoint.FieldList
	5,  // [5:5] is the sub-list for method output_type
	5,  // [5:5] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_common_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_common_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_common_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[1].OneofWrappers = []any{
		(*Secret_Raw)(nil),
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[3].OneofWrappers = []any{
		(*TLSMode_Disabled)(nil),
		(*TLSMode_Enabled)(nil),
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[5].OneofWrappers = []any{
		(*ColumnValue_StringValue)(nil),
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes[8].OneofWrappers = []any{
		(*DataSchema_JsonFields)(nil),
		(*DataSchema_Fields)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_common_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_common_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_common_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_datatransfer_v1_endpoint_common_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_common_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_common_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_depIdxs = nil
}
