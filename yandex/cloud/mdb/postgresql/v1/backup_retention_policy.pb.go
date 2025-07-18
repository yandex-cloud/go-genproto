// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/postgresql/v1/backup_retention_policy.proto

package postgresql

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// Message to describe a crontab schedule.
type CronTab struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	DayOfMonth    string                 `protobuf:"bytes,3,opt,name=day_of_month,json=dayOfMonth,proto3" json:"day_of_month,omitempty"`
	Month         string                 `protobuf:"bytes,4,opt,name=month,proto3" json:"month,omitempty"`
	DayOfWeek     string                 `protobuf:"bytes,5,opt,name=day_of_week,json=dayOfWeek,proto3" json:"day_of_week,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CronTab) Reset() {
	*x = CronTab{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CronTab) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CronTab) ProtoMessage() {}

func (x *CronTab) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CronTab.ProtoReflect.Descriptor instead.
func (*CronTab) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescGZIP(), []int{0}
}

func (x *CronTab) GetDayOfMonth() string {
	if x != nil {
		return x.DayOfMonth
	}
	return ""
}

func (x *CronTab) GetMonth() string {
	if x != nil {
		return x.Month
	}
	return ""
}

func (x *CronTab) GetDayOfWeek() string {
	if x != nil {
		return x.DayOfWeek
	}
	return ""
}

// Message to describe a retention policy for cluster backups.
type BackupRetentionPolicy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Required. Policy ID.
	PolicyId string `protobuf:"bytes,1,opt,name=policy_id,json=policyId,proto3" json:"policy_id,omitempty"`
	// PostgreSQL cluster ID.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Required. Policy name.
	PolicyName string `protobuf:"bytes,3,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// CronTab schedule.
	Cron *CronTab `protobuf:"bytes,5,opt,name=cron,proto3" json:"cron,omitempty"`
	// Retention duration.
	RetainForDays int64 `protobuf:"varint,6,opt,name=retain_for_days,json=retainForDays,proto3" json:"retain_for_days,omitempty"`
	// Human-readable description.
	Description   string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BackupRetentionPolicy) Reset() {
	*x = BackupRetentionPolicy{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BackupRetentionPolicy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BackupRetentionPolicy) ProtoMessage() {}

func (x *BackupRetentionPolicy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BackupRetentionPolicy.ProtoReflect.Descriptor instead.
func (*BackupRetentionPolicy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescGZIP(), []int{1}
}

func (x *BackupRetentionPolicy) GetPolicyId() string {
	if x != nil {
		return x.PolicyId
	}
	return ""
}

func (x *BackupRetentionPolicy) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *BackupRetentionPolicy) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *BackupRetentionPolicy) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BackupRetentionPolicy) GetCron() *CronTab {
	if x != nil {
		return x.Cron
	}
	return nil
}

func (x *BackupRetentionPolicy) GetRetainForDays() int64 {
	if x != nil {
		return x.RetainForDays
	}
	return 0
}

func (x *BackupRetentionPolicy) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDesc = "" +
	"\n" +
	"<yandex/cloud/mdb/postgresql/v1/backup_retention_policy.proto\x12\x1eyandex.cloud.mdb.postgresql.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dyandex/cloud/validation.proto\"{\n" +
	"\aCronTab\x12 \n" +
	"\fday_of_month\x18\x03 \x01(\tR\n" +
	"dayOfMonth\x12\x14\n" +
	"\x05month\x18\x04 \x01(\tR\x05month\x12\x1e\n" +
	"\vday_of_week\x18\x05 \x01(\tR\tdayOfWeekJ\x04\b\x01\x10\x02J\x04\b\x02\x10\x03R\x06minuteR\x04hour\"\xd0\x02\n" +
	"\x15BackupRetentionPolicy\x12!\n" +
	"\tpolicy_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\bpolicyId\x12+\n" +
	"\n" +
	"cluster_id\x18\x02 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12%\n" +
	"\vpolicy_name\x18\x03 \x01(\tB\x04\xe8\xc71\x01R\n" +
	"policyName\x129\n" +
	"\n" +
	"created_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12;\n" +
	"\x04cron\x18\x05 \x01(\v2'.yandex.cloud.mdb.postgresql.v1.CronTabR\x04cron\x12&\n" +
	"\x0fretain_for_days\x18\x06 \x01(\x03R\rretainForDays\x12 \n" +
	"\vdescription\x18\a \x01(\tR\vdescriptionBs\n" +
	"\"yandex.cloud.api.mdb.postgresql.v1ZMgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1;postgresqlb\x06proto3"

var (
	file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDesc), len(file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDescData
}

var file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_goTypes = []any{
	(*CronTab)(nil),               // 0: yandex.cloud.mdb.postgresql.v1.CronTab
	(*BackupRetentionPolicy)(nil), // 1: yandex.cloud.mdb.postgresql.v1.BackupRetentionPolicy
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.mdb.postgresql.v1.BackupRetentionPolicy.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: yandex.cloud.mdb.postgresql.v1.BackupRetentionPolicy.cron:type_name -> yandex.cloud.mdb.postgresql.v1.CronTab
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_init() }
func file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_init() {
	if File_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDesc), len(file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto = out.File
	file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_goTypes = nil
	file_yandex_cloud_mdb_postgresql_v1_backup_retention_policy_proto_depIdxs = nil
}
