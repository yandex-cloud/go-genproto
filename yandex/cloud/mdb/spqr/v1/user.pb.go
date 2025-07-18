// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/spqr/v1/user.proto

package spqr

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// A SPQR User resource. For more information, see the
// [Developer's Guide](/docs/managed-spqr/concepts).
type User struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the SPQR user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the SPQR cluster the user belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Set of permissions granted to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// SPQR Settings for this user
	Settings *UserSettings `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
	// User grants
	Grants        []string `protobuf:"bytes,5,rep,name=grants,proto3" json:"grants,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[0]
	if x != nil {
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
	return file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescGZIP(), []int{0}
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

func (x *User) GetSettings() *UserSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *User) GetGrants() []string {
	if x != nil {
		return x.Grants
	}
	return nil
}

type Permission struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the database that the permission grants access to.
	DatabaseName  string `protobuf:"bytes,1,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Permission) Reset() {
	*x = Permission{}
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Permission) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Permission) ProtoMessage() {}

func (x *Permission) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[1]
	if x != nil {
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
	return file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *Permission) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type UserSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the SPQR user.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Password of the SPQR user.
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// Set of permissions to grant to the user.
	Permissions []*Permission `protobuf:"bytes,3,rep,name=permissions,proto3" json:"permissions,omitempty"`
	// SPQR Settings for this user
	Settings *UserSettings `protobuf:"bytes,4,opt,name=settings,proto3" json:"settings,omitempty"`
	// User grants
	Grants        []string `protobuf:"bytes,5,rep,name=grants,proto3" json:"grants,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserSpec) Reset() {
	*x = UserSpec{}
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSpec) ProtoMessage() {}

func (x *UserSpec) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[2]
	if x != nil {
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
	return file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescGZIP(), []int{2}
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

func (x *UserSpec) GetSettings() *UserSettings {
	if x != nil {
		return x.Settings
	}
	return nil
}

func (x *UserSpec) GetGrants() []string {
	if x != nil {
		return x.Grants
	}
	return nil
}

type UserSettings struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ConnectionLimit   *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=connection_limit,json=connectionLimit,proto3" json:"connection_limit,omitempty"`
	ConnectionRetries *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=connection_retries,json=connectionRetries,proto3" json:"connection_retries,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *UserSettings) Reset() {
	*x = UserSettings{}
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSettings) ProtoMessage() {}

func (x *UserSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSettings.ProtoReflect.Descriptor instead.
func (*UserSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserSettings) GetConnectionLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.ConnectionLimit
	}
	return nil
}

func (x *UserSettings) GetConnectionRetries() *wrapperspb.Int64Value {
	if x != nil {
		return x.ConnectionRetries
	}
	return nil
}

var File_yandex_cloud_mdb_spqr_v1_user_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_spqr_v1_user_proto_rawDesc = "" +
	"\n" +
	"#yandex/cloud/mdb/spqr/v1/user.proto\x12\x18yandex.cloud.mdb.spqr.v1\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1dyandex/cloud/validation.proto\"\xf9\x01\n" +
	"\x04User\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x02 \x01(\tR\tclusterId\x12F\n" +
	"\vpermissions\x18\x03 \x03(\v2$.yandex.cloud.mdb.spqr.v1.PermissionR\vpermissions\x12B\n" +
	"\bsettings\x18\x04 \x01(\v2&.yandex.cloud.mdb.spqr.v1.UserSettingsR\bsettings\x122\n" +
	"\x06grants\x18\x05 \x03(\tB\x1a\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\x06grants\"1\n" +
	"\n" +
	"Permission\x12#\n" +
	"\rdatabase_name\x18\x01 \x01(\tR\fdatabaseName\"\xa9\x02\n" +
	"\bUserSpec\x122\n" +
	"\x04name\x18\x01 \x01(\tB\x1e\xe8\xc71\x01\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\x04name\x12)\n" +
	"\bpassword\x18\x02 \x01(\tB\r\xe8\xc71\x01\x8a\xc81\x058-128R\bpassword\x12F\n" +
	"\vpermissions\x18\x03 \x03(\v2$.yandex.cloud.mdb.spqr.v1.PermissionR\vpermissions\x12B\n" +
	"\bsettings\x18\x04 \x01(\v2&.yandex.cloud.mdb.spqr.v1.UserSettingsR\bsettings\x122\n" +
	"\x06grants\x18\x05 \x03(\tB\x1a\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\x06grants\"\xa2\x01\n" +
	"\fUserSettings\x12F\n" +
	"\x10connection_limit\x18\x01 \x01(\v2\x1b.google.protobuf.Int64ValueR\x0fconnectionLimit\x12J\n" +
	"\x12connection_retries\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\x11connectionRetriesBa\n" +
	"\x1cyandex.cloud.api.mdb.spqr.v1ZAgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/spqr/v1;spqrb\x06proto3"

var (
	file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_spqr_v1_user_proto_rawDesc), len(file_yandex_cloud_mdb_spqr_v1_user_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_spqr_v1_user_proto_rawDescData
}

var file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_mdb_spqr_v1_user_proto_goTypes = []any{
	(*User)(nil),                  // 0: yandex.cloud.mdb.spqr.v1.User
	(*Permission)(nil),            // 1: yandex.cloud.mdb.spqr.v1.Permission
	(*UserSpec)(nil),              // 2: yandex.cloud.mdb.spqr.v1.UserSpec
	(*UserSettings)(nil),          // 3: yandex.cloud.mdb.spqr.v1.UserSettings
	(*wrapperspb.Int64Value)(nil), // 4: google.protobuf.Int64Value
}
var file_yandex_cloud_mdb_spqr_v1_user_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.mdb.spqr.v1.User.permissions:type_name -> yandex.cloud.mdb.spqr.v1.Permission
	3, // 1: yandex.cloud.mdb.spqr.v1.User.settings:type_name -> yandex.cloud.mdb.spqr.v1.UserSettings
	1, // 2: yandex.cloud.mdb.spqr.v1.UserSpec.permissions:type_name -> yandex.cloud.mdb.spqr.v1.Permission
	3, // 3: yandex.cloud.mdb.spqr.v1.UserSpec.settings:type_name -> yandex.cloud.mdb.spqr.v1.UserSettings
	4, // 4: yandex.cloud.mdb.spqr.v1.UserSettings.connection_limit:type_name -> google.protobuf.Int64Value
	4, // 5: yandex.cloud.mdb.spqr.v1.UserSettings.connection_retries:type_name -> google.protobuf.Int64Value
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_spqr_v1_user_proto_init() }
func file_yandex_cloud_mdb_spqr_v1_user_proto_init() {
	if File_yandex_cloud_mdb_spqr_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_spqr_v1_user_proto_rawDesc), len(file_yandex_cloud_mdb_spqr_v1_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_spqr_v1_user_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_spqr_v1_user_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_spqr_v1_user_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_spqr_v1_user_proto = out.File
	file_yandex_cloud_mdb_spqr_v1_user_proto_goTypes = nil
	file_yandex_cloud_mdb_spqr_v1_user_proto_depIdxs = nil
}
