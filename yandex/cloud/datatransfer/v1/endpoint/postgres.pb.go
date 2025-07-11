// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/datatransfer/v1/endpoint/postgres.proto

package endpoint

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type PostgresObjectTransferSettings struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Sequences
	//
	// CREATE SEQUENCE ...
	Sequence ObjectTransferStage `protobuf:"varint,1,opt,name=sequence,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"sequence,omitempty"`
	// Owned sequences
	//
	// CREATE SEQUENCE ... OWNED BY ...
	SequenceOwnedBy ObjectTransferStage `protobuf:"varint,2,opt,name=sequence_owned_by,json=sequenceOwnedBy,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"sequence_owned_by,omitempty"`
	// Tables
	//
	// CREATE TABLE ...
	Table ObjectTransferStage `protobuf:"varint,3,opt,name=table,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"table,omitempty"`
	// Primary keys
	//
	// ALTER TABLE ... ADD PRIMARY KEY ...
	PrimaryKey ObjectTransferStage `protobuf:"varint,4,opt,name=primary_key,json=primaryKey,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"primary_key,omitempty"`
	// Foreign keys
	//
	// ALTER TABLE ... ADD FOREIGN KEY ...
	FkConstraint ObjectTransferStage `protobuf:"varint,5,opt,name=fk_constraint,json=fkConstraint,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"fk_constraint,omitempty"`
	// Default values
	//
	// ALTER TABLE ... ALTER COLUMN ... SET DEFAULT ...
	DefaultValues ObjectTransferStage `protobuf:"varint,6,opt,name=default_values,json=defaultValues,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"default_values,omitempty"`
	// Constraints
	//
	// ALTER TABLE ... ADD CONSTRAINT ...
	Constraint ObjectTransferStage `protobuf:"varint,7,opt,name=constraint,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"constraint,omitempty"`
	// Indexes
	//
	// CREATE INDEX ...
	Index ObjectTransferStage `protobuf:"varint,8,opt,name=index,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"index,omitempty"`
	// Views
	//
	// CREATE VIEW ...
	View ObjectTransferStage `protobuf:"varint,9,opt,name=view,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"view,omitempty"`
	// Functions
	//
	// CREATE FUNCTION ...
	Function ObjectTransferStage `protobuf:"varint,10,opt,name=function,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"function,omitempty"`
	// Triggers
	//
	// CREATE TRIGGER ...
	Trigger ObjectTransferStage `protobuf:"varint,11,opt,name=trigger,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"trigger,omitempty"`
	// Types
	//
	// CREATE TYPE ...
	Type ObjectTransferStage `protobuf:"varint,12,opt,name=type,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"type,omitempty"`
	// Rules
	//
	// CREATE RULE ...
	Rule ObjectTransferStage `protobuf:"varint,13,opt,name=rule,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"rule,omitempty"`
	// Collations
	//
	// CREATE COLLATION ...
	Collation ObjectTransferStage `protobuf:"varint,14,opt,name=collation,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"collation,omitempty"`
	// Policies
	//
	// CREATE POLICY ...
	Policy ObjectTransferStage `protobuf:"varint,15,opt,name=policy,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"policy,omitempty"`
	// Casts
	//
	// CREATE CAST ...
	Cast ObjectTransferStage `protobuf:"varint,16,opt,name=cast,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"cast,omitempty"`
	// Materialized views
	//
	// CREATE MATERIALIZED VIEW ...
	MaterializedView ObjectTransferStage `protobuf:"varint,17,opt,name=materialized_view,json=materializedView,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"materialized_view,omitempty"`
	SequenceSet   ObjectTransferStage `protobuf:"varint,18,opt,name=sequence_set,json=sequenceSet,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage" json:"sequence_set,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostgresObjectTransferSettings) Reset() {
	*x = PostgresObjectTransferSettings{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostgresObjectTransferSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresObjectTransferSettings) ProtoMessage() {}

func (x *PostgresObjectTransferSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresObjectTransferSettings.ProtoReflect.Descriptor instead.
func (*PostgresObjectTransferSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescGZIP(), []int{0}
}

func (x *PostgresObjectTransferSettings) GetSequence() ObjectTransferStage {
	if x != nil {
		return x.Sequence
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetSequenceOwnedBy() ObjectTransferStage {
	if x != nil {
		return x.SequenceOwnedBy
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetTable() ObjectTransferStage {
	if x != nil {
		return x.Table
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetPrimaryKey() ObjectTransferStage {
	if x != nil {
		return x.PrimaryKey
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetFkConstraint() ObjectTransferStage {
	if x != nil {
		return x.FkConstraint
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetDefaultValues() ObjectTransferStage {
	if x != nil {
		return x.DefaultValues
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetConstraint() ObjectTransferStage {
	if x != nil {
		return x.Constraint
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetIndex() ObjectTransferStage {
	if x != nil {
		return x.Index
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetView() ObjectTransferStage {
	if x != nil {
		return x.View
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetFunction() ObjectTransferStage {
	if x != nil {
		return x.Function
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetTrigger() ObjectTransferStage {
	if x != nil {
		return x.Trigger
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetType() ObjectTransferStage {
	if x != nil {
		return x.Type
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetRule() ObjectTransferStage {
	if x != nil {
		return x.Rule
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetCollation() ObjectTransferStage {
	if x != nil {
		return x.Collation
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetPolicy() ObjectTransferStage {
	if x != nil {
		return x.Policy
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetCast() ObjectTransferStage {
	if x != nil {
		return x.Cast
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetMaterializedView() ObjectTransferStage {
	if x != nil {
		return x.MaterializedView
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

func (x *PostgresObjectTransferSettings) GetSequenceSet() ObjectTransferStage {
	if x != nil {
		return x.SequenceSet
	}
	return ObjectTransferStage_OBJECT_TRANSFER_STAGE_UNSPECIFIED
}

type OnPremisePostgres struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Will be used if the cluster ID is not specified.
	Port int64 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// Network interface for endpoint. If none will assume public ipv4
	SubnetId string   `protobuf:"bytes,4,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	Hosts    []string `protobuf:"bytes,5,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// TLS settings for server connection. Disabled by default.
	TlsMode       *TLSMode `protobuf:"bytes,6,opt,name=tls_mode,json=tlsMode,proto3" json:"tls_mode,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OnPremisePostgres) Reset() {
	*x = OnPremisePostgres{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OnPremisePostgres) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OnPremisePostgres) ProtoMessage() {}

func (x *OnPremisePostgres) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OnPremisePostgres.ProtoReflect.Descriptor instead.
func (*OnPremisePostgres) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescGZIP(), []int{1}
}

func (x *OnPremisePostgres) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *OnPremisePostgres) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *OnPremisePostgres) GetHosts() []string {
	if x != nil {
		return x.Hosts
	}
	return nil
}

func (x *OnPremisePostgres) GetTlsMode() *TLSMode {
	if x != nil {
		return x.TlsMode
	}
	return nil
}

type PostgresConnection struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Connection:
	//
	//	*PostgresConnection_MdbClusterId
	//	*PostgresConnection_OnPremise
	//	*PostgresConnection_ConnectionManagerConnection
	Connection    isPostgresConnection_Connection `protobuf_oneof:"connection"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PostgresConnection) Reset() {
	*x = PostgresConnection{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostgresConnection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresConnection) ProtoMessage() {}

func (x *PostgresConnection) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresConnection.ProtoReflect.Descriptor instead.
func (*PostgresConnection) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescGZIP(), []int{2}
}

func (x *PostgresConnection) GetConnection() isPostgresConnection_Connection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *PostgresConnection) GetMdbClusterId() string {
	if x != nil {
		if x, ok := x.Connection.(*PostgresConnection_MdbClusterId); ok {
			return x.MdbClusterId
		}
	}
	return ""
}

func (x *PostgresConnection) GetOnPremise() *OnPremisePostgres {
	if x != nil {
		if x, ok := x.Connection.(*PostgresConnection_OnPremise); ok {
			return x.OnPremise
		}
	}
	return nil
}

func (x *PostgresConnection) GetConnectionManagerConnection() *ConnectionManagerConnection {
	if x != nil {
		if x, ok := x.Connection.(*PostgresConnection_ConnectionManagerConnection); ok {
			return x.ConnectionManagerConnection
		}
	}
	return nil
}

type isPostgresConnection_Connection interface {
	isPostgresConnection_Connection()
}

type PostgresConnection_MdbClusterId struct {
	// Managed Service for PostgreSQL cluster ID
	MdbClusterId string `protobuf:"bytes,1,opt,name=mdb_cluster_id,json=mdbClusterId,proto3,oneof"`
}

type PostgresConnection_OnPremise struct {
	// Connection options for on-premise PostgreSQL
	OnPremise *OnPremisePostgres `protobuf:"bytes,2,opt,name=on_premise,json=onPremise,proto3,oneof"`
}

type PostgresConnection_ConnectionManagerConnection struct {
	ConnectionManagerConnection *ConnectionManagerConnection `protobuf:"bytes,3,opt,name=connection_manager_connection,json=connectionManagerConnection,proto3,oneof"`
}

func (*PostgresConnection_MdbClusterId) isPostgresConnection_Connection() {}

func (*PostgresConnection_OnPremise) isPostgresConnection_Connection() {}

func (*PostgresConnection_ConnectionManagerConnection) isPostgresConnection_Connection() {}

type PostgresSource struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database connection settings
	Connection *PostgresConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Database name
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// User for database access. not required as may be in connection
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Password for database access.
	Password *Secret `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Included tables
	//
	// If none or empty list is presented, all tables are replicated. Full table name
	// with schema. Can contain schema_name.* patterns.
	IncludeTables []string `protobuf:"bytes,5,rep,name=include_tables,json=includeTables,proto3" json:"include_tables,omitempty"`
	// Excluded tables
	//
	// If none or empty list is presented, all tables are replicated. Full table name
	// with schema. Can contain schema_name.* patterns.
	ExcludeTables []string `protobuf:"bytes,6,rep,name=exclude_tables,json=excludeTables,proto3" json:"exclude_tables,omitempty"`
	// Maximum lag of replication slot (in bytes); after exceeding this limit
	// replication will be aborted.
	SlotByteLagLimit int64 `protobuf:"varint,8,opt,name=slot_byte_lag_limit,json=slotByteLagLimit,proto3" json:"slot_byte_lag_limit,omitempty"`
	// Database schema for service tables (__consumer_keeper,
	// __data_transfer_mole_finder). Default is public
	ServiceSchema string `protobuf:"bytes,9,opt,name=service_schema,json=serviceSchema,proto3" json:"service_schema,omitempty"`
	// Select database objects to be transferred during activation or deactivation.
	ObjectTransferSettings *PostgresObjectTransferSettings `protobuf:"bytes,13,opt,name=object_transfer_settings,json=objectTransferSettings,proto3" json:"object_transfer_settings,omitempty"`
	// Security groups
	SecurityGroups []string `protobuf:"bytes,14,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PostgresSource) Reset() {
	*x = PostgresSource{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostgresSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresSource) ProtoMessage() {}

func (x *PostgresSource) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresSource.ProtoReflect.Descriptor instead.
func (*PostgresSource) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescGZIP(), []int{3}
}

func (x *PostgresSource) GetConnection() *PostgresConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *PostgresSource) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *PostgresSource) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PostgresSource) GetPassword() *Secret {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *PostgresSource) GetIncludeTables() []string {
	if x != nil {
		return x.IncludeTables
	}
	return nil
}

func (x *PostgresSource) GetExcludeTables() []string {
	if x != nil {
		return x.ExcludeTables
	}
	return nil
}

func (x *PostgresSource) GetSlotByteLagLimit() int64 {
	if x != nil {
		return x.SlotByteLagLimit
	}
	return 0
}

func (x *PostgresSource) GetServiceSchema() string {
	if x != nil {
		return x.ServiceSchema
	}
	return ""
}

func (x *PostgresSource) GetObjectTransferSettings() *PostgresObjectTransferSettings {
	if x != nil {
		return x.ObjectTransferSettings
	}
	return nil
}

func (x *PostgresSource) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

type PostgresTarget struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Database connection settings
	Connection *PostgresConnection `protobuf:"bytes,1,opt,name=connection,proto3" json:"connection,omitempty"`
	// Database name
	Database string `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	// User for database access. not required as may be in connection
	User string `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	// Password for database access.
	Password *Secret `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Cleanup policy for activate, reactivate and reupload processes. Default is
	// truncate.
	CleanupPolicy CleanupPolicy `protobuf:"varint,5,opt,name=cleanup_policy,json=cleanupPolicy,proto3,enum=yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy" json:"cleanup_policy,omitempty"`
	// Security groups
	SecurityGroups            []string `protobuf:"bytes,7,rep,name=security_groups,json=securityGroups,proto3" json:"security_groups,omitempty"`
	IsSchemaMigrationDisabled bool     `protobuf:"varint,8,opt,name=is_schema_migration_disabled,json=isSchemaMigrationDisabled,proto3" json:"is_schema_migration_disabled,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *PostgresTarget) Reset() {
	*x = PostgresTarget{}
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PostgresTarget) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresTarget) ProtoMessage() {}

func (x *PostgresTarget) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresTarget.ProtoReflect.Descriptor instead.
func (*PostgresTarget) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescGZIP(), []int{4}
}

func (x *PostgresTarget) GetConnection() *PostgresConnection {
	if x != nil {
		return x.Connection
	}
	return nil
}

func (x *PostgresTarget) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *PostgresTarget) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *PostgresTarget) GetPassword() *Secret {
	if x != nil {
		return x.Password
	}
	return nil
}

func (x *PostgresTarget) GetCleanupPolicy() CleanupPolicy {
	if x != nil {
		return x.CleanupPolicy
	}
	return CleanupPolicy_CLEANUP_POLICY_UNSPECIFIED
}

func (x *PostgresTarget) GetSecurityGroups() []string {
	if x != nil {
		return x.SecurityGroups
	}
	return nil
}

func (x *PostgresTarget) GetIsSchemaMigrationDisabled() bool {
	if x != nil {
		return x.IsSchemaMigrationDisabled
	}
	return false
}

var File_yandex_cloud_datatransfer_v1_endpoint_postgres_proto protoreflect.FileDescriptor

const file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDesc = "" +
	"\n" +
	"4yandex/cloud/datatransfer/v1/endpoint/postgres.proto\x12%yandex.cloud.datatransfer.v1.endpoint\x1a2yandex/cloud/datatransfer/v1/endpoint/common.proto\"\xe5\f\n" +
	"\x1ePostgresObjectTransferSettings\x12V\n" +
	"\bsequence\x18\x01 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\bsequence\x12f\n" +
	"\x11sequence_owned_by\x18\x02 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x0fsequenceOwnedBy\x12P\n" +
	"\x05table\x18\x03 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x05table\x12[\n" +
	"\vprimary_key\x18\x04 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\n" +
	"primaryKey\x12_\n" +
	"\rfk_constraint\x18\x05 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\ffkConstraint\x12a\n" +
	"\x0edefault_values\x18\x06 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\rdefaultValues\x12Z\n" +
	"\n" +
	"constraint\x18\a \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\n" +
	"constraint\x12P\n" +
	"\x05index\x18\b \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x05index\x12N\n" +
	"\x04view\x18\t \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x04view\x12V\n" +
	"\bfunction\x18\n" +
	" \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\bfunction\x12T\n" +
	"\atrigger\x18\v \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\atrigger\x12N\n" +
	"\x04type\x18\f \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x04type\x12N\n" +
	"\x04rule\x18\r \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x04rule\x12X\n" +
	"\tcollation\x18\x0e \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\tcollation\x12R\n" +
	"\x06policy\x18\x0f \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x06policy\x12N\n" +
	"\x04cast\x18\x10 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x04cast\x12g\n" +
	"\x11materialized_view\x18\x11 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\x10materializedView\x12]\n" +
	"\fsequence_set\x18\x12 \x01(\x0e2:.yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStageR\vsequenceSet\"\xb1\x01\n" +
	"\x11OnPremisePostgres\x12\x12\n" +
	"\x04port\x18\x02 \x01(\x03R\x04port\x12\x1b\n" +
	"\tsubnet_id\x18\x04 \x01(\tR\bsubnetId\x12\x14\n" +
	"\x05hosts\x18\x05 \x03(\tR\x05hosts\x12I\n" +
	"\btls_mode\x18\x06 \x01(\v2..yandex.cloud.datatransfer.v1.endpoint.TLSModeR\atlsModeJ\x04\b\x01\x10\x02J\x04\b\x03\x10\x04\"\xb0\x02\n" +
	"\x12PostgresConnection\x12&\n" +
	"\x0emdb_cluster_id\x18\x01 \x01(\tH\x00R\fmdbClusterId\x12Y\n" +
	"\n" +
	"on_premise\x18\x02 \x01(\v28.yandex.cloud.datatransfer.v1.endpoint.OnPremisePostgresH\x00R\tonPremise\x12\x88\x01\n" +
	"\x1dconnection_manager_connection\x18\x03 \x01(\v2B.yandex.cloud.datatransfer.v1.endpoint.ConnectionManagerConnectionH\x00R\x1bconnectionManagerConnectionB\f\n" +
	"\n" +
	"connection\"\xc0\x04\n" +
	"\x0ePostgresSource\x12Y\n" +
	"\n" +
	"connection\x18\x01 \x01(\v29.yandex.cloud.datatransfer.v1.endpoint.PostgresConnectionR\n" +
	"connection\x12\x1a\n" +
	"\bdatabase\x18\x02 \x01(\tR\bdatabase\x12\x12\n" +
	"\x04user\x18\x03 \x01(\tR\x04user\x12I\n" +
	"\bpassword\x18\x04 \x01(\v2-.yandex.cloud.datatransfer.v1.endpoint.SecretR\bpassword\x12%\n" +
	"\x0einclude_tables\x18\x05 \x03(\tR\rincludeTables\x12%\n" +
	"\x0eexclude_tables\x18\x06 \x03(\tR\rexcludeTables\x12-\n" +
	"\x13slot_byte_lag_limit\x18\b \x01(\x03R\x10slotByteLagLimit\x12%\n" +
	"\x0eservice_schema\x18\t \x01(\tR\rserviceSchema\x12\x7f\n" +
	"\x18object_transfer_settings\x18\r \x01(\v2E.yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettingsR\x16objectTransferSettings\x12'\n" +
	"\x0fsecurity_groups\x18\x0e \x03(\tR\x0esecurityGroupsJ\x04\b\a\x10\bJ\x04\b\n" +
	"\x10\r\"\xb3\x03\n" +
	"\x0ePostgresTarget\x12Y\n" +
	"\n" +
	"connection\x18\x01 \x01(\v29.yandex.cloud.datatransfer.v1.endpoint.PostgresConnectionR\n" +
	"connection\x12\x1a\n" +
	"\bdatabase\x18\x02 \x01(\tR\bdatabase\x12\x12\n" +
	"\x04user\x18\x03 \x01(\tR\x04user\x12I\n" +
	"\bpassword\x18\x04 \x01(\v2-.yandex.cloud.datatransfer.v1.endpoint.SecretR\bpassword\x12[\n" +
	"\x0ecleanup_policy\x18\x05 \x01(\x0e24.yandex.cloud.datatransfer.v1.endpoint.CleanupPolicyR\rcleanupPolicy\x12'\n" +
	"\x0fsecurity_groups\x18\a \x03(\tR\x0esecurityGroups\x12?\n" +
	"\x1cis_schema_migration_disabled\x18\b \x01(\bR\x19isSchemaMigrationDisabledJ\x04\b\x06\x10\aB\xa7\x01\n" +
	")yandex.cloud.api.datatransfer.v1.endpointZRgithub.com/yandex-cloud/go-genproto/yandex/cloud/datatransfer/v1/endpoint;endpoint\xaa\x02%Yandex.Cloud.Datatransfer.V1.EndPointb\x06proto3"

var (
	file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescOnce sync.Once
	file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescData []byte
)

func file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDesc)))
	})
	return file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDescData
}

var file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_goTypes = []any{
	(*PostgresObjectTransferSettings)(nil), // 0: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings
	(*OnPremisePostgres)(nil),              // 1: yandex.cloud.datatransfer.v1.endpoint.OnPremisePostgres
	(*PostgresConnection)(nil),             // 2: yandex.cloud.datatransfer.v1.endpoint.PostgresConnection
	(*PostgresSource)(nil),                 // 3: yandex.cloud.datatransfer.v1.endpoint.PostgresSource
	(*PostgresTarget)(nil),                 // 4: yandex.cloud.datatransfer.v1.endpoint.PostgresTarget
	(ObjectTransferStage)(0),               // 5: yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	(*TLSMode)(nil),                        // 6: yandex.cloud.datatransfer.v1.endpoint.TLSMode
	(*ConnectionManagerConnection)(nil),    // 7: yandex.cloud.datatransfer.v1.endpoint.ConnectionManagerConnection
	(*Secret)(nil),                         // 8: yandex.cloud.datatransfer.v1.endpoint.Secret
	(CleanupPolicy)(0),                     // 9: yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy
}
var file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_depIdxs = []int32{
	5,  // 0: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.sequence:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 1: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.sequence_owned_by:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 2: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.table:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 3: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.primary_key:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 4: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.fk_constraint:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 5: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.default_values:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 6: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.constraint:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 7: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.index:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 8: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.view:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 9: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.function:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 10: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.trigger:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 11: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.type:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 12: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.rule:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 13: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.collation:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 14: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.policy:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 15: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.cast:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 16: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.materialized_view:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	5,  // 17: yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings.sequence_set:type_name -> yandex.cloud.datatransfer.v1.endpoint.ObjectTransferStage
	6,  // 18: yandex.cloud.datatransfer.v1.endpoint.OnPremisePostgres.tls_mode:type_name -> yandex.cloud.datatransfer.v1.endpoint.TLSMode
	1,  // 19: yandex.cloud.datatransfer.v1.endpoint.PostgresConnection.on_premise:type_name -> yandex.cloud.datatransfer.v1.endpoint.OnPremisePostgres
	7,  // 20: yandex.cloud.datatransfer.v1.endpoint.PostgresConnection.connection_manager_connection:type_name -> yandex.cloud.datatransfer.v1.endpoint.ConnectionManagerConnection
	2,  // 21: yandex.cloud.datatransfer.v1.endpoint.PostgresSource.connection:type_name -> yandex.cloud.datatransfer.v1.endpoint.PostgresConnection
	8,  // 22: yandex.cloud.datatransfer.v1.endpoint.PostgresSource.password:type_name -> yandex.cloud.datatransfer.v1.endpoint.Secret
	0,  // 23: yandex.cloud.datatransfer.v1.endpoint.PostgresSource.object_transfer_settings:type_name -> yandex.cloud.datatransfer.v1.endpoint.PostgresObjectTransferSettings
	2,  // 24: yandex.cloud.datatransfer.v1.endpoint.PostgresTarget.connection:type_name -> yandex.cloud.datatransfer.v1.endpoint.PostgresConnection
	8,  // 25: yandex.cloud.datatransfer.v1.endpoint.PostgresTarget.password:type_name -> yandex.cloud.datatransfer.v1.endpoint.Secret
	9,  // 26: yandex.cloud.datatransfer.v1.endpoint.PostgresTarget.cleanup_policy:type_name -> yandex.cloud.datatransfer.v1.endpoint.CleanupPolicy
	27, // [27:27] is the sub-list for method output_type
	27, // [27:27] is the sub-list for method input_type
	27, // [27:27] is the sub-list for extension type_name
	27, // [27:27] is the sub-list for extension extendee
	0,  // [0:27] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_init() }
func file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_init() {
	if File_yandex_cloud_datatransfer_v1_endpoint_postgres_proto != nil {
		return
	}
	file_yandex_cloud_datatransfer_v1_endpoint_common_proto_init()
	file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes[2].OneofWrappers = []any{
		(*PostgresConnection_MdbClusterId)(nil),
		(*PostgresConnection_OnPremise)(nil),
		(*PostgresConnection_ConnectionManagerConnection)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDesc), len(file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datatransfer_v1_endpoint_postgres_proto = out.File
	file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_goTypes = nil
	file_yandex_cloud_datatransfer_v1_endpoint_postgres_proto_depIdxs = nil
}
