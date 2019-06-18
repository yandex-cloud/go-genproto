// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/user.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

var Permission_Privilege_name = map[int32]string{
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

var Permission_Privilege_value = map[string]int32{
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

func (x Permission_Privilege) String() string {
	return proto.EnumName(Permission_Privilege_name, int32(x))
}

func (Permission_Privilege) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a381b11fcae3a32d, []int{1, 0}
}

// A MySQL user. For more information, see
// the [documentation](/docs/managed-mysql/concepts).
type User struct {
	// Name of the MySQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the MySQL cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions          []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_a381b11fcae3a32d, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetClusterId() string {
	if m != nil {
		return m.ClusterId
	}
	return ""
}

func (m *User) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Permission struct {
	// Name of the database that the permission grants access to.
	DatabaseName string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Roles granted to the user within the database.
	Roles                []Permission_Privilege `protobuf:"varint,2,rep,packed,name=roles,proto3,enum=yandex.cloud.mdb.mysql.v1alpha.Permission_Privilege" json:"roles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *Permission) Reset()         { *m = Permission{} }
func (m *Permission) String() string { return proto.CompactTextString(m) }
func (*Permission) ProtoMessage()    {}
func (*Permission) Descriptor() ([]byte, []int) {
	return fileDescriptor_a381b11fcae3a32d, []int{1}
}

func (m *Permission) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permission.Unmarshal(m, b)
}
func (m *Permission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permission.Marshal(b, m, deterministic)
}
func (m *Permission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permission.Merge(m, src)
}
func (m *Permission) XXX_Size() int {
	return xxx_messageInfo_Permission.Size(m)
}
func (m *Permission) XXX_DiscardUnknown() {
	xxx_messageInfo_Permission.DiscardUnknown(m)
}

var xxx_messageInfo_Permission proto.InternalMessageInfo

func (m *Permission) GetDatabaseName() string {
	if m != nil {
		return m.DatabaseName
	}
	return ""
}

func (m *Permission) GetRoles() []Permission_Privilege {
	if m != nil {
		return m.Roles
	}
	return nil
}

type UserSpec struct {
	// Name of the MySQL user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the MySQL user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions          []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UserSpec) Reset()         { *m = UserSpec{} }
func (m *UserSpec) String() string { return proto.CompactTextString(m) }
func (*UserSpec) ProtoMessage()    {}
func (*UserSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a381b11fcae3a32d, []int{2}
}

func (m *UserSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserSpec.Unmarshal(m, b)
}
func (m *UserSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserSpec.Marshal(b, m, deterministic)
}
func (m *UserSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserSpec.Merge(m, src)
}
func (m *UserSpec) XXX_Size() int {
	return xxx_messageInfo_UserSpec.Size(m)
}
func (m *UserSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UserSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UserSpec proto.InternalMessageInfo

func (m *UserSpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserSpec) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserSpec) GetPermissions() []*Permission {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.mysql.v1alpha.Permission_Privilege", Permission_Privilege_name, Permission_Privilege_value)
	proto.RegisterType((*User)(nil), "yandex.cloud.mdb.mysql.v1alpha.User")
	proto.RegisterType((*Permission)(nil), "yandex.cloud.mdb.mysql.v1alpha.Permission")
	proto.RegisterType((*UserSpec)(nil), "yandex.cloud.mdb.mysql.v1alpha.UserSpec")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/user.proto", fileDescriptor_a381b11fcae3a32d)
}

var fileDescriptor_a381b11fcae3a32d = []byte{
	// 602 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcd, 0x6e, 0xd3, 0x4e,
	0x10, 0xff, 0xe7, 0xab, 0x4d, 0x26, 0x4d, 0xbb, 0x5d, 0xe9, 0x2f, 0x42, 0x51, 0xab, 0x2a, 0x5c,
	0xda, 0x4a, 0xb1, 0x71, 0xca, 0xa1, 0x08, 0x8a, 0x94, 0x8f, 0xa1, 0x35, 0x98, 0xc4, 0x5a, 0x3b,
	0x6d, 0x29, 0x42, 0x96, 0x13, 0x2f, 0xa9, 0x25, 0x27, 0x36, 0x76, 0xd2, 0x52, 0x8e, 0x5c, 0x90,
	0xfa, 0x34, 0x9c, 0x90, 0x78, 0x82, 0xf6, 0x51, 0x38, 0xf3, 0x04, 0x68, 0xed, 0x34, 0x84, 0x4b,
	0xc5, 0x81, 0xdb, 0xec, 0xfc, 0x3e, 0x66, 0x56, 0xfa, 0x0d, 0x6c, 0x5f, 0xda, 0x23, 0x87, 0x7f,
	0x94, 0xfb, 0x9e, 0x3f, 0x71, 0xe4, 0xa1, 0xd3, 0x93, 0x87, 0x97, 0xd1, 0x07, 0x4f, 0x3e, 0x57,
	0x6c, 0x2f, 0x38, 0xb3, 0xe5, 0x49, 0xc4, 0x43, 0x29, 0x08, 0xfd, 0xb1, 0x4f, 0x37, 0x12, 0xaa,
	0x14, 0x53, 0xa5, 0xa1, 0xd3, 0x93, 0x62, 0xaa, 0x34, 0xa5, 0xae, 0x6d, 0x0c, 0x7c, 0x7f, 0xe0,
	0x71, 0x39, 0x66, 0xf7, 0x26, 0xef, 0xe5, 0x8b, 0xd0, 0x0e, 0x02, 0x1e, 0x46, 0x89, 0x7e, 0x6d,
	0xfd, 0x8f, 0x51, 0xe7, 0xb6, 0xe7, 0x3a, 0xf6, 0xd8, 0xf5, 0x47, 0x09, 0x5c, 0xf9, 0x92, 0x82,
	0x6c, 0x37, 0xe2, 0x21, 0xa5, 0x90, 0x1d, 0xd9, 0x43, 0x5e, 0x4e, 0x6d, 0xa6, 0xb6, 0x0a, 0x2c,
	0xae, 0xe9, 0x3a, 0x40, 0xdf, 0x9b, 0x44, 0x63, 0x1e, 0x5a, 0xae, 0x53, 0x4e, 0xc7, 0x48, 0x61,
	0xda, 0x51, 0x1d, 0xaa, 0x41, 0x31, 0xe0, 0xe1, 0xd0, 0x8d, 0x22, 0xd7, 0x1f, 0x45, 0xe5, 0xcc,
	0x66, 0x66, 0xab, 0x58, 0xdb, 0x91, 0xee, 0x5e, 0x58, 0xd2, 0x67, 0x12, 0x36, 0x2f, 0xaf, 0x7c,
	0xcf, 0x00, 0xfc, 0xc6, 0xe8, 0x43, 0x28, 0x39, 0xf6, 0xd8, 0xee, 0xd9, 0x11, 0xb7, 0xe6, 0x16,
	0x5b, 0xba, 0x6d, 0xb6, 0xc5, 0x82, 0x06, 0xe4, 0x42, 0xdf, 0xe3, 0x51, 0x39, 0xbd, 0x99, 0xd9,
	0x5a, 0xae, 0x3d, 0xfe, 0xfb, 0xd9, 0x92, 0x1e, 0xba, 0xe7, 0xae, 0xc7, 0x07, 0xbc, 0xb1, 0xf8,
	0xf9, 0x46, 0xc9, 0x3c, 0xdf, 0x57, 0x58, 0xe2, 0x55, 0xf9, 0x9a, 0x86, 0xc2, 0x0c, 0xa5, 0xf7,
	0xe1, 0x7f, 0x9d, 0xa9, 0x47, 0xaa, 0x86, 0x07, 0x68, 0x75, 0xdb, 0x86, 0x8e, 0x4d, 0xf5, 0x85,
	0x8a, 0x2d, 0xf2, 0x1f, 0xa5, 0xb0, 0x5c, 0xd7, 0x34, 0x6b, 0x06, 0x1b, 0x24, 0x45, 0x0b, 0x90,
	0xab, 0x6b, 0x26, 0x32, 0x92, 0xa6, 0xab, 0x50, 0x8a, 0x4b, 0x8b, 0x75, 0xba, 0xa6, 0xda, 0x46,
	0x92, 0xa1, 0x00, 0x0b, 0x4d, 0x86, 0x75, 0x13, 0x49, 0x56, 0xa8, 0x93, 0x7a, 0x86, 0xe7, 0xe8,
	0x03, 0xb8, 0x37, 0xed, 0x99, 0xf8, 0x5a, 0xef, 0xb0, 0x3a, 0x7b, 0x63, 0x99, 0xf5, 0x86, 0x86,
	0x06, 0x59, 0xa0, 0x2b, 0x50, 0x9c, 0x82, 0x47, 0x2a, 0x1e, 0x93, 0x45, 0xe1, 0xd6, 0x42, 0x0d,
	0x4d, 0x24, 0x79, 0x9a, 0x87, 0x6c, 0x8b, 0x75, 0x74, 0x52, 0x10, 0x1b, 0xe0, 0x11, 0xb6, 0x4d,
	0x02, 0xb4, 0x08, 0x8b, 0x78, 0x82, 0xcd, 0xae, 0x89, 0xa4, 0x28, 0xfa, 0x6a, 0xbb, 0x85, 0x27,
	0x64, 0x49, 0x08, 0xd5, 0xb6, 0x81, 0xcc, 0x24, 0x25, 0xe1, 0xaa, 0x75, 0x9a, 0xaf, 0x6e, 0xc7,
	0x2c, 0x0b, 0xd0, 0x40, 0x0d, 0x9b, 0x26, 0x59, 0xa1, 0x25, 0x28, 0x18, 0x87, 0x9d, 0xe3, 0x64,
	0x20, 0x11, 0x7e, 0x26, 0x53, 0x0f, 0x0e, 0x90, 0x91, 0x55, 0xc1, 0xeb, 0xea, 0x2d, 0xf1, 0x17,
	0x5a, 0xf9, 0x96, 0x82, 0xbc, 0x48, 0x91, 0x11, 0xf0, 0x3e, 0x55, 0xe6, 0x93, 0xd4, 0x58, 0xff,
	0x71, 0xad, 0xa4, 0x7e, 0x5e, 0x2b, 0xa5, 0xb7, 0x76, 0xf5, 0x53, 0xbd, 0x7a, 0xfa, 0xa8, 0xfa,
	0xc4, 0x7a, 0xb7, 0x73, 0x75, 0xa3, 0x64, 0x9f, 0xed, 0xef, 0xd6, 0xa6, 0x41, 0xdb, 0x86, 0x7c,
	0x60, 0x47, 0xd1, 0x85, 0x1f, 0x4e, 0x63, 0xd6, 0x28, 0x09, 0xd9, 0xd5, 0x8d, 0x92, 0xdb, 0xab,
	0x2a, 0xb5, 0x3d, 0x36, 0x83, 0xff, 0x6d, 0xe8, 0x1a, 0x2f, 0x4f, 0x0f, 0x07, 0xee, 0xf8, 0x6c,
	0xd2, 0x93, 0xfa, 0xfe, 0x50, 0x4e, 0x4c, 0xaa, 0xc9, 0xa9, 0x0c, 0xfc, 0xea, 0x80, 0x8f, 0xe2,
	0x2b, 0x91, 0xef, 0x3e, 0xd7, 0xa7, 0xf1, 0xab, 0xb7, 0x10, 0x73, 0x77, 0x7f, 0x05, 0x00, 0x00,
	0xff, 0xff, 0x6e, 0x31, 0x78, 0x75, 0xdd, 0x03, 0x00, 0x00,
}
