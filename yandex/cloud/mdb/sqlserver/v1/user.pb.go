// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/mdb/sqlserver/v1/user.proto

package sqlserver

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

type Permission_Role int32

const (
	Permission_ROLE_UNSPECIFIED Permission_Role = 0
	// Members of this fixed database role can perform all configuration and maintenance activities on the database, and can also drop the database in SQL Server.
	Permission_DB_OWNER Permission_Role = 1
	// Members of this fixed database role can modify role membership for custom roles only and manage permissions. They can potentially elevate their privileges and their actions should be monitored.
	Permission_DB_SECURITYADMIN Permission_Role = 2
	// Members of this fixed database role can add or remove access to the database for Windows logins, Windows groups, and SQL Server logins.
	Permission_DB_ACCESSADMIN Permission_Role = 3
	// Members of this fixed database role can back up the database.
	Permission_DB_BACKUPOPERATOR Permission_Role = 4
	// Members of this fixed database role can run any Data Definition Language (DDL) command in a database.
	Permission_DB_DDLADMIN Permission_Role = 5
	// Members of this fixed database role can add, delete, or change data in all user tables.
	Permission_DB_DATAWRITER Permission_Role = 6
	// Members of this fixed database role can read all data from all user tables.
	Permission_DB_DATAREADER Permission_Role = 7
	// Members of this fixed database role cannot add, modify, or delete any data in the user tables within a database. Denial has a higher priority than a grant, so you can use this role to quickly restrict one's privileges without explicitly revoking permissions or roles.
	Permission_DB_DENYDATAWRITER Permission_Role = 8
	// Members of this fixed database role cannot read any data in the user tables within a database. Denial has a higher priority than a grant, so you can use this role to quickly restrict one's privileges without explicitly revoking permissions or roles.
	Permission_DB_DENYDATAREADER Permission_Role = 9
)

// Enum value maps for Permission_Role.
var (
	Permission_Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "DB_OWNER",
		2: "DB_SECURITYADMIN",
		3: "DB_ACCESSADMIN",
		4: "DB_BACKUPOPERATOR",
		5: "DB_DDLADMIN",
		6: "DB_DATAWRITER",
		7: "DB_DATAREADER",
		8: "DB_DENYDATAWRITER",
		9: "DB_DENYDATAREADER",
	}
	Permission_Role_value = map[string]int32{
		"ROLE_UNSPECIFIED":  0,
		"DB_OWNER":          1,
		"DB_SECURITYADMIN":  2,
		"DB_ACCESSADMIN":    3,
		"DB_BACKUPOPERATOR": 4,
		"DB_DDLADMIN":       5,
		"DB_DATAWRITER":     6,
		"DB_DATAREADER":     7,
		"DB_DENYDATAWRITER": 8,
		"DB_DENYDATAREADER": 9,
	}
)

func (x Permission_Role) Enum() *Permission_Role {
	p := new(Permission_Role)
	*p = x
	return p
}

func (x Permission_Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Permission_Role) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_sqlserver_v1_user_proto_enumTypes[0].Descriptor()
}

func (Permission_Role) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_sqlserver_v1_user_proto_enumTypes[0]
}

func (x Permission_Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Permission_Role.Descriptor instead.
func (Permission_Role) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescGZIP(), []int{1, 0}
}

// An SQL Server user.
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the SQL Server user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the SQL Server cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[0]
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
	return file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescGZIP(), []int{0}
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

	// Name of the database the permission grants access to.
	DatabaseName string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Roles granted to the user within the database.
	Roles []Permission_Role `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=yandex.cloud.mdb.sqlserver.v1.Permission_Role" json:"roles,omitempty"`
}

func (x *Permission) Reset() {
	*x = Permission{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[1]
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
	return file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *Permission) GetRoles() []Permission_Role {
	if x != nil {
		return x.Roles
	}
	return nil
}

type UserSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the SQL Server user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the SQL Server user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserSpec) Reset() {
	*x = UserSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpec) ProtoMessage() {}

func (x *UserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[2]
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
	return file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescGZIP(), []int{2}
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

var File_yandex_cloud_mdb_sqlserver_v1_user_proto protoreflect.FileDescriptor

var file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDesc = []byte{
	0x0a, 0x28, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x64, 0x62, 0x2f, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0xd3, 0x02, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x23, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x07, 0x82, 0xc8, 0x31, 0x03, 0x3e, 0x3d, 0x31, 0x52, 0x05, 0x72,
	0x6f, 0x6c, 0x65, 0x73, 0x22, 0xd0, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x42, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x10,
	0x01, 0x12, 0x14, 0x0a, 0x10, 0x44, 0x42, 0x5f, 0x53, 0x45, 0x43, 0x55, 0x52, 0x49, 0x54, 0x59,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x44, 0x42, 0x5f, 0x41, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44,
	0x42, 0x5f, 0x42, 0x41, 0x43, 0x4b, 0x55, 0x50, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x4f, 0x52,
	0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x42, 0x5f, 0x44, 0x44, 0x4c, 0x41, 0x44, 0x4d, 0x49,
	0x4e, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x42, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x57, 0x52,
	0x49, 0x54, 0x45, 0x52, 0x10, 0x06, 0x12, 0x11, 0x0a, 0x0d, 0x44, 0x42, 0x5f, 0x44, 0x41, 0x54,
	0x41, 0x52, 0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x42, 0x5f,
	0x44, 0x45, 0x4e, 0x59, 0x44, 0x41, 0x54, 0x41, 0x57, 0x52, 0x49, 0x54, 0x45, 0x52, 0x10, 0x08,
	0x12, 0x15, 0x0a, 0x11, 0x44, 0x42, 0x5f, 0x44, 0x45, 0x4e, 0x59, 0x44, 0x41, 0x54, 0x41, 0x52,
	0x45, 0x41, 0x44, 0x45, 0x52, 0x10, 0x09, 0x22, 0xb5, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x1d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x33, 0x32,
	0xf2, 0xc7, 0x31, 0x0d, 0x5b, 0x61, 0x2d, 0x7a, 0x41, 0x2d, 0x5a, 0x30, 0x2d, 0x39, 0x5f, 0x5d,
	0x2a, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a,
	0xc8, 0x31, 0x05, 0x38, 0x2d, 0x31, 0x32, 0x38, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x4b, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x42,
	0x75, 0x0a, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6d, 0x64, 0x62, 0x2e, 0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x42, 0x03, 0x50, 0x53, 0x55, 0x5a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x64, 0x62, 0x2f,
	0x73, 0x71, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x71, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescData = file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDesc
)

func file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescData)
	})
	return file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDescData
}

var file_yandex_cloud_mdb_sqlserver_v1_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_sqlserver_v1_user_proto_goTypes = []interface{}{
	(Permission_Role)(0), // 0: yandex.cloud.mdb.sqlserver.v1.Permission.Role
	(*User)(nil),         // 1: yandex.cloud.mdb.sqlserver.v1.User
	(*Permission)(nil),   // 2: yandex.cloud.mdb.sqlserver.v1.Permission
	(*UserSpec)(nil),     // 3: yandex.cloud.mdb.sqlserver.v1.UserSpec
}
var file_yandex_cloud_mdb_sqlserver_v1_user_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.mdb.sqlserver.v1.User.permissions:type_name -> yandex.cloud.mdb.sqlserver.v1.Permission
	0, // 1: yandex.cloud.mdb.sqlserver.v1.Permission.roles:type_name -> yandex.cloud.mdb.sqlserver.v1.Permission.Role
	2, // 2: yandex.cloud.mdb.sqlserver.v1.UserSpec.permissions:type_name -> yandex.cloud.mdb.sqlserver.v1.Permission
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_sqlserver_v1_user_proto_init() }
func file_yandex_cloud_mdb_sqlserver_v1_user_proto_init() {
	if File_yandex_cloud_mdb_sqlserver_v1_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_sqlserver_v1_user_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_sqlserver_v1_user_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_sqlserver_v1_user_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_sqlserver_v1_user_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_sqlserver_v1_user_proto = out.File
	file_yandex_cloud_mdb_sqlserver_v1_user_proto_rawDesc = nil
	file_yandex_cloud_mdb_sqlserver_v1_user_proto_goTypes = nil
	file_yandex_cloud_mdb_sqlserver_v1_user_proto_depIdxs = nil
}
