// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/config/mysql5_7.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// Options and structure of `MysqlConfig5_7` reflects MySQL 5.7 configuration file
type MysqlConfig5_7 struct {
	// Size of the InnoDB buffer pool used for caching table and index data.
	//
	// For details, see [MySQL documentation for the parameter](https://dev.mysql.com/doc/refman/5.7/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).
	InnodbBufferPoolSize *wrappers.Int64Value `protobuf:"bytes,1,opt,name=innodb_buffer_pool_size,json=innodbBufferPoolSize,proto3" json:"innodb_buffer_pool_size,omitempty"`
	// The maximum permitted number of simultaneous client connections.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_max_connections).
	MaxConnections *wrappers.Int64Value `protobuf:"bytes,2,opt,name=max_connections,json=maxConnections,proto3" json:"max_connections,omitempty"`
	// Time that it takes to process a query before it is considered slow.
	//
	// For details, see [MySQL documentation for the variable](https://dev.mysql.com/doc/refman/5.7/en/server-system-variables.html#sysvar_long_query_time).
	LongQueryTime        *wrappers.DoubleValue `protobuf:"bytes,3,opt,name=long_query_time,json=longQueryTime,proto3" json:"long_query_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MysqlConfig5_7) Reset()         { *m = MysqlConfig5_7{} }
func (m *MysqlConfig5_7) String() string { return proto.CompactTextString(m) }
func (*MysqlConfig5_7) ProtoMessage()    {}
func (*MysqlConfig5_7) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ad7cd5b2ad0622, []int{0}
}

func (m *MysqlConfig5_7) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfig5_7.Unmarshal(m, b)
}
func (m *MysqlConfig5_7) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfig5_7.Marshal(b, m, deterministic)
}
func (m *MysqlConfig5_7) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfig5_7.Merge(m, src)
}
func (m *MysqlConfig5_7) XXX_Size() int {
	return xxx_messageInfo_MysqlConfig5_7.Size(m)
}
func (m *MysqlConfig5_7) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfig5_7.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfig5_7 proto.InternalMessageInfo

func (m *MysqlConfig5_7) GetInnodbBufferPoolSize() *wrappers.Int64Value {
	if m != nil {
		return m.InnodbBufferPoolSize
	}
	return nil
}

func (m *MysqlConfig5_7) GetMaxConnections() *wrappers.Int64Value {
	if m != nil {
		return m.MaxConnections
	}
	return nil
}

func (m *MysqlConfig5_7) GetLongQueryTime() *wrappers.DoubleValue {
	if m != nil {
		return m.LongQueryTime
	}
	return nil
}

type MysqlConfigSet5_7 struct {
	// Effective settings for a MySQL 5.7 cluster (a combination of settings defined
	// in [user_config] and [default_config]).
	EffectiveConfig *MysqlConfig5_7 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a MySQL 5.7 cluster.
	UserConfig *MysqlConfig5_7 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a MySQL 5.7 cluster.
	DefaultConfig        *MysqlConfig5_7 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MysqlConfigSet5_7) Reset()         { *m = MysqlConfigSet5_7{} }
func (m *MysqlConfigSet5_7) String() string { return proto.CompactTextString(m) }
func (*MysqlConfigSet5_7) ProtoMessage()    {}
func (*MysqlConfigSet5_7) Descriptor() ([]byte, []int) {
	return fileDescriptor_e2ad7cd5b2ad0622, []int{1}
}

func (m *MysqlConfigSet5_7) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MysqlConfigSet5_7.Unmarshal(m, b)
}
func (m *MysqlConfigSet5_7) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MysqlConfigSet5_7.Marshal(b, m, deterministic)
}
func (m *MysqlConfigSet5_7) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MysqlConfigSet5_7.Merge(m, src)
}
func (m *MysqlConfigSet5_7) XXX_Size() int {
	return xxx_messageInfo_MysqlConfigSet5_7.Size(m)
}
func (m *MysqlConfigSet5_7) XXX_DiscardUnknown() {
	xxx_messageInfo_MysqlConfigSet5_7.DiscardUnknown(m)
}

var xxx_messageInfo_MysqlConfigSet5_7 proto.InternalMessageInfo

func (m *MysqlConfigSet5_7) GetEffectiveConfig() *MysqlConfig5_7 {
	if m != nil {
		return m.EffectiveConfig
	}
	return nil
}

func (m *MysqlConfigSet5_7) GetUserConfig() *MysqlConfig5_7 {
	if m != nil {
		return m.UserConfig
	}
	return nil
}

func (m *MysqlConfigSet5_7) GetDefaultConfig() *MysqlConfig5_7 {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*MysqlConfig5_7)(nil), "yandex.cloud.mdb.mysql.v1alpha.config.MysqlConfig5_7")
	proto.RegisterType((*MysqlConfigSet5_7)(nil), "yandex.cloud.mdb.mysql.v1alpha.config.MysqlConfigSet5_7")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/config/mysql5_7.proto", fileDescriptor_e2ad7cd5b2ad0622)
}

var fileDescriptor_e2ad7cd5b2ad0622 = []byte{
	// 432 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0xd3, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0x06, 0x70, 0xba, 0x05, 0xd1, 0xa9, 0xbb, 0xab, 0x41, 0x70, 0xa9, 0x7f, 0x90, 0x82, 0xa0,
	0x17, 0x3b, 0x93, 0xac, 0xbb, 0xb6, 0x20, 0x7a, 0xb1, 0xed, 0x8d, 0x17, 0xa2, 0xa6, 0xd2, 0x0b,
	0x11, 0xc6, 0x49, 0x72, 0x92, 0x0e, 0xcc, 0xcc, 0x49, 0x93, 0xcc, 0xba, 0x5b, 0x7c, 0x05, 0x9f,
	0xc4, 0x07, 0xe9, 0x3b, 0xf5, 0x4a, 0x32, 0xb3, 0xab, 0x2e, 0x22, 0x2d, 0xec, 0xed, 0x49, 0xbe,
	0xdf, 0x09, 0xdf, 0x64, 0xc8, 0x78, 0x21, 0x4c, 0x06, 0x73, 0x96, 0x2a, 0xb4, 0x19, 0xd3, 0x59,
	0xc2, 0xf4, 0xa2, 0x3e, 0x53, 0x6c, 0x16, 0x09, 0x55, 0x9e, 0x0a, 0x96, 0xa2, 0xc9, 0x65, 0xe1,
	0x87, 0x13, 0xbe, 0x4f, 0xcb, 0x0a, 0x1b, 0x0c, 0x9e, 0xfa, 0x14, 0x75, 0x29, 0xaa, 0xb3, 0x84,
	0xba, 0x17, 0xe8, 0x32, 0x45, 0x7d, 0x6a, 0xf7, 0x71, 0x81, 0x58, 0x28, 0x60, 0x2e, 0x94, 0xd8,
	0x9c, 0x7d, 0xab, 0x44, 0x59, 0x42, 0x55, 0x7b, 0x66, 0xf7, 0xd1, 0xda, 0xf2, 0x99, 0x50, 0x32,
	0x13, 0x8d, 0x44, 0xe3, 0x1f, 0xef, 0xfd, 0xe8, 0x90, 0xde, 0xbb, 0xd6, 0x3d, 0x74, 0xdc, 0x84,
	0xef, 0x07, 0x82, 0xdc, 0x97, 0xc6, 0x60, 0x96, 0xf0, 0xc4, 0xe6, 0x39, 0x54, 0xbc, 0x44, 0x54,
	0xbc, 0x96, 0xe7, 0x30, 0xd8, 0x7a, 0xb2, 0xf5, 0x6c, 0x67, 0xf4, 0x80, 0xfa, 0x9d, 0x74, 0xb5,
	0x93, 0xbe, 0x35, 0xcd, 0xcb, 0xf1, 0x89, 0x50, 0x16, 0xa6, 0xdd, 0xcb, 0x8b, 0xe8, 0xd6, 0x9b,
	0xd7, 0x93, 0xd1, 0x78, 0x74, 0x70, 0x10, 0xc6, 0xf7, 0x3c, 0x35, 0x75, 0xd2, 0x07, 0x44, 0x75,
	0x2c, 0xcf, 0x21, 0x88, 0x49, 0x5f, 0x8b, 0x39, 0x4f, 0xd1, 0x18, 0x48, 0xdb, 0xaf, 0xa9, 0x07,
	0x9d, 0xab, 0xe9, 0xdb, 0x97, 0x17, 0xd1, 0xcd, 0x28, 0x1c, 0x46, 0x61, 0x18, 0x86, 0x71, 0x4f,
	0x8b, 0xf9, 0xe1, 0x1f, 0x20, 0x38, 0x22, 0x7d, 0x85, 0xa6, 0xe0, 0x67, 0x16, 0xaa, 0x05, 0x6f,
	0xa4, 0x86, 0xc1, 0xb6, 0x33, 0x1f, 0xfe, 0x63, 0x1e, 0xa1, 0x4d, 0x14, 0x38, 0x34, 0xee, 0xb6,
	0xa1, 0x8f, 0x6d, 0xe6, 0x93, 0xd4, 0xb0, 0xf7, 0xb3, 0x43, 0xee, 0xfe, 0xd5, 0xc7, 0x31, 0x34,
	0x6d, 0x25, 0x5f, 0xc9, 0x1d, 0xc8, 0xf3, 0x76, 0xd3, 0x0c, 0xb8, 0x2f, 0x7e, 0xd9, 0xc5, 0x84,
	0x5e, 0xeb, 0x98, 0xe8, 0x7a, 0xc7, 0x71, 0xff, 0x37, 0xe7, 0x67, 0xc1, 0x09, 0xd9, 0xb1, 0x35,
	0x54, 0x2b, 0xbc, 0xb3, 0x09, 0x4e, 0x5a, 0x69, 0xe9, 0x7e, 0x21, 0xbd, 0x0c, 0x72, 0x61, 0x55,
	0xb3, 0xa2, 0xb7, 0x37, 0xa1, 0xbb, 0x4b, 0xcc, 0x4f, 0xa6, 0xdf, 0xc9, 0xf3, 0x35, 0x46, 0x94,
	0xf2, 0xbf, 0xd4, 0xe7, 0xf7, 0x85, 0x6c, 0x4e, 0x6d, 0x42, 0x53, 0xd4, 0xcc, 0xa7, 0x86, 0xfe,
	0xa7, 0x2c, 0x70, 0x58, 0x80, 0x71, 0xa7, 0xc3, 0xae, 0x75, 0x55, 0x5e, 0xb9, 0x61, 0x72, 0xc3,
	0x45, 0x5e, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xed, 0xa2, 0x71, 0xc8, 0x60, 0x03, 0x00, 0x00,
}
