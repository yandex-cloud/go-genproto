// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/mysql/v1alpha/user.proto

package mysql

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

type Permission_Privilege int32

const (
	Permission_PRIVILEGE_UNSPECIFIED Permission_Privilege = 0
	// All privileges that can be made available to the user.
	Permission_ALL_PRIVILEGES Permission_Privilege = 1
	// Altering tables.
	Permission_ALTER Permission_Privilege = 2
	// Altering stored routines (stored procedures and functions).
	Permission_ALTER_ROUTINE Permission_Privilege = 3
	// Creating tables or indexes.
	Permission_CREATE Permission_Privilege = 4
	// Creating stored routines.
	Permission_CREATE_ROUTINE Permission_Privilege = 5
	// Creating temporary tables.
	Permission_CREATE_TEMPORARY_TABLES Permission_Privilege = 6
	// Creating views.
	Permission_CREATE_VIEW Permission_Privilege = 7
	// Deleting tables.
	Permission_DELETE Permission_Privilege = 8
	// Removing tables or views.
	Permission_DROP Permission_Privilege = 9
	// Creating, altering, dropping, or displaying events for the Event Scheduler.
	Permission_EVENT Permission_Privilege = 10
	// Executing stored routines.
	Permission_EXECUTE Permission_Privilege = 11
	// Creating and removing indexes.
	Permission_INDEX Permission_Privilege = 12
	// Inserting rows into the database.
	Permission_INSERT Permission_Privilege = 13
	// Using LOCK TABLES statement for tables available with SELECT privilege.
	Permission_LOCK_TABLES Permission_Privilege = 14
	// Selecting rows from tables.
	//
	// Some SELECT statements can be allowed without the SELECT privilege. All
	// statements that read column values require the SELECT privilege. See
	// details in [MySQL documentation](https://dev.mysql.com/doc/refman/5.7/en/privileges-provided.html#priv_select).
	Permission_SELECT Permission_Privilege = 15
	// Using the SHOW CREATE VIEW statement. Also needed for views used with EXPLAIN.
	Permission_SHOW_VIEW Permission_Privilege = 16
	// Creating, removing, executing, or displaying triggers for a table.
	Permission_TRIGGER Permission_Privilege = 17
	// Updating rows in the database.
	Permission_UPDATE Permission_Privilege = 18
)

// Enum value maps for Permission_Privilege.
var (
	Permission_Privilege_name = map[int32]string{
		0:  "PRIVILEGE_UNSPECIFIED",
		1:  "ALL_PRIVILEGES",
		2:  "ALTER",
		3:  "ALTER_ROUTINE",
		4:  "CREATE",
		5:  "CREATE_ROUTINE",
		6:  "CREATE_TEMPORARY_TABLES",
		7:  "CREATE_VIEW",
		8:  "DELETE",
		9:  "DROP",
		10: "EVENT",
		11: "EXECUTE",
		12: "INDEX",
		13: "INSERT",
		14: "LOCK_TABLES",
		15: "SELECT",
		16: "SHOW_VIEW",
		17: "TRIGGER",
		18: "UPDATE",
	}
	Permission_Privilege_value = map[string]int32{
		"PRIVILEGE_UNSPECIFIED":   0,
		"ALL_PRIVILEGES":          1,
		"ALTER":                   2,
		"ALTER_ROUTINE":           3,
		"CREATE":                  4,
		"CREATE_ROUTINE":          5,
		"CREATE_TEMPORARY_TABLES": 6,
		"CREATE_VIEW":             7,
		"DELETE":                  8,
		"DROP":                    9,
		"EVENT":                   10,
		"EXECUTE":                 11,
		"INDEX":                   12,
		"INSERT":                  13,
		"LOCK_TABLES":             14,
		"SELECT":                  15,
		"SHOW_VIEW":               16,
		"TRIGGER":                 17,
		"UPDATE":                  18,
	}
)

func (x Permission_Privilege) Enum() *Permission_Privilege {
	p := new(Permission_Privilege)
	*p = x
	return p
}

func (x Permission_Privilege) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission_Privilege) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_mysql_v1alpha_user_proto_enumTypes[0].Descriptor()
}

func (Permission_Privilege) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_enumTypes[0]
}

func (x Permission_Privilege) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission_Privilege.Descriptor instead.
func (Permission_Privilege) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescGZIP(), []int{1, 0}
}

// A MySQL user. For more information, see
// the [documentation](/docs/managed-mysql/concepts).
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the MySQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the MySQL cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *User) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type Permission struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the database that the permission grants access to.
	DatabaseName string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Roles granted to the user within the database.
	Roles []Permission_Privilege `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=yandex.cloud.mdb.mysql.v1alpha.Permission_Privilege" json:"roles,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Permission.ProtoReflect.Descriptor instead.
func (*Permission) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *Permission) GetRoles() []Permission_Privilege {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UserSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the MySQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the MySQL user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserSpec) Reset() {
	*x = UserSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpec) ProtoMessage() {}

func (x *UserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSpec.ProtoReflect.Descriptor instead.
func (*UserSpec) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserSpec) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserSpec) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSpec) GetPermissions() []*Permission {
	if x != nil {
		return x.Permissions
	}
	return nil
}

var File_yandex_cloud_mdb_mysql_v1alpha_user_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDesc = []byte{
	0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79,
	0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1d, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d,
	0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x22, 0xb9, 0x03, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x53, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x34, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x42, 0x07, 0x82,
	0xc8, 0x31, 0x03, 0x3e, 0x3d, 0x31, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22, 0xb0, 0x02,
	0x0a, 0x09, 0x50, 0x72, 0x69, 0x76, 0x69, 0x6c, 0x65, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50,
	0x52, 0x49, 0x56, 0x49, 0x4c, 0x45, 0x47, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x41, 0x4c, 0x4c, 0x5f, 0x50, 0x52,
	0x49, 0x56, 0x49, 0x4c, 0x45, 0x47, 0x45, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x4c,
	0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x52,
	0x4f, 0x55, 0x54, 0x49, 0x4e, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x10, 0x04, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f, 0x52,
	0x4f, 0x55, 0x54, 0x49, 0x4e, 0x45, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17, 0x43, 0x52, 0x45, 0x41,
	0x54, 0x45, 0x5f, 0x54, 0x45, 0x4d, 0x50, 0x4f, 0x52, 0x41, 0x52, 0x59, 0x5f, 0x54, 0x41, 0x42,
	0x4c, 0x45, 0x53, 0x10, 0x06, 0x12, 0x0f, 0x0a, 0x0b, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x5f,
	0x56, 0x49, 0x45, 0x57, 0x10, 0x07, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45,
	0x10, 0x08, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x52, 0x4f, 0x50, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x45, 0x10, 0x0b, 0x12, 0x09, 0x0a, 0x05, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x10, 0x0c, 0x12,
	0x0a, 0x0a, 0x06, 0x49, 0x4e, 0x53, 0x45, 0x52, 0x54, 0x10, 0x0d, 0x12, 0x0f, 0x0a, 0x0b, 0x4c,
	0x4f, 0x43, 0x4b, 0x5f, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x53, 0x10, 0x0e, 0x12, 0x0a, 0x0a, 0x06,
	0x53, 0x45, 0x4c, 0x45, 0x43, 0x54, 0x10, 0x0f, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x48, 0x4f, 0x57,
	0x5f, 0x56, 0x49, 0x45, 0x57, 0x10, 0x10, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52, 0x49, 0x47, 0x47,
	0x45, 0x52, 0x10, 0x11, 0x12, 0x0a, 0x0a, 0x06, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x10, 0x12,
	0x22, 0xb6, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x33, 0x32, 0xf2, 0xc7, 0x31, 0x0d, 0x5b, 0x61, 0x2d,
	0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d, 0x2a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x29, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05, 0x38, 0x2d, 0x31, 0x32,
	0x38, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x4c, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6d, 0x64, 0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x6e, 0x0a, 0x22, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64,
	0x62, 0x2e, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x3b, 0x6d, 0x79, 0x73, 0x71, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescData = file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDesc
)

func file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDescData
}

var file_yandex_cloud_mdb_mysql_v1alpha_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_mysql_v1alpha_user_proto_goTypes = []interface{}{
	(Permission_Privilege)(0), // 0: yandex.cloud.mdb.mysql.v1alpha.Permission.Privilege
	(*User)(nil),              // 1: yandex.cloud.mdb.mysql.v1alpha.User
	(*Permission)(nil),        // 2: yandex.cloud.mdb.mysql.v1alpha.Permission
	(*UserSpec)(nil),          // 3: yandex.cloud.mdb.mysql.v1alpha.UserSpec
}
var file_yandex_cloud_mdb_mysql_v1alpha_user_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.mdb.mysql.v1alpha.User.permissions:type_name -> yandex.cloud.mdb.mysql.v1alpha.Permission
	0, // 1: yandex.cloud.mdb.mysql.v1alpha.Permission.roles:type_name -> yandex.cloud.mdb.mysql.v1alpha.Permission.Privilege
	2, // 2: yandex.cloud.mdb.mysql.v1alpha.UserSpec.permissions:type_name -> yandex.cloud.mdb.mysql.v1alpha.Permission
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_mysql_v1alpha_user_proto_init() }
func file_yandex_cloud_mdb_mysql_v1alpha_user_proto_init() {
	if File_yandex_cloud_mdb_mysql_v1alpha_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Permission); i {
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
		file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSpec); i {
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
			RawDescriptor: file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_mysql_v1alpha_user_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_mysql_v1alpha_user_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_mysql_v1alpha_user_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_mysql_v1alpha_user_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_mysql_v1alpha_user_proto = out.File
	file_yandex_cloud_mdb_mysql_v1alpha_user_proto_rawDesc = nil
	file_yandex_cloud_mdb_mysql_v1alpha_user_proto_goTypes = nil
	file_yandex_cloud_mdb_mysql_v1alpha_user_proto_depIdxs = nil
}
