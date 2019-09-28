// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/redis/v1alpha/config/redis5_0.proto

package redis

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

type RedisConfig5_0_MaxmemoryPolicy int32

const (
	RedisConfig5_0_MAXMEMORY_POLICY_UNSPECIFIED RedisConfig5_0_MaxmemoryPolicy = 0
	// Try to remove less recently used (LRU) keys with `expire set`.
	RedisConfig5_0_VOLATILE_LRU RedisConfig5_0_MaxmemoryPolicy = 1
	// Remove less recently used (LRU) keys.
	RedisConfig5_0_ALLKEYS_LRU RedisConfig5_0_MaxmemoryPolicy = 2
	// Try to remove least frequently used (LFU) keys with `expire set`.
	RedisConfig5_0_VOLATILE_LFU RedisConfig5_0_MaxmemoryPolicy = 3
	// Remove least frequently used (LFU) keys.
	RedisConfig5_0_ALLKEYS_LFU RedisConfig5_0_MaxmemoryPolicy = 4
	// Try to remove keys with `expire set` randomly.
	RedisConfig5_0_VOLATILE_RANDOM RedisConfig5_0_MaxmemoryPolicy = 5
	// Remove keys randomly.
	RedisConfig5_0_ALLKEYS_RANDOM RedisConfig5_0_MaxmemoryPolicy = 6
	// Try to remove less recently used (LRU) keys with `expire set`
	// and shorter TTL first.
	RedisConfig5_0_VOLATILE_TTL RedisConfig5_0_MaxmemoryPolicy = 7
	// Return errors when memory limit was reached and commands could require
	// more memory to be used.
	RedisConfig5_0_NOEVICTION RedisConfig5_0_MaxmemoryPolicy = 8
)

var RedisConfig5_0_MaxmemoryPolicy_name = map[int32]string{
	0: "MAXMEMORY_POLICY_UNSPECIFIED",
	1: "VOLATILE_LRU",
	2: "ALLKEYS_LRU",
	3: "VOLATILE_LFU",
	4: "ALLKEYS_LFU",
	5: "VOLATILE_RANDOM",
	6: "ALLKEYS_RANDOM",
	7: "VOLATILE_TTL",
	8: "NOEVICTION",
}

var RedisConfig5_0_MaxmemoryPolicy_value = map[string]int32{
	"MAXMEMORY_POLICY_UNSPECIFIED": 0,
	"VOLATILE_LRU":                 1,
	"ALLKEYS_LRU":                  2,
	"VOLATILE_LFU":                 3,
	"ALLKEYS_LFU":                  4,
	"VOLATILE_RANDOM":              5,
	"ALLKEYS_RANDOM":               6,
	"VOLATILE_TTL":                 7,
	"NOEVICTION":                   8,
}

func (x RedisConfig5_0_MaxmemoryPolicy) String() string {
	return proto.EnumName(RedisConfig5_0_MaxmemoryPolicy_name, int32(x))
}

func (RedisConfig5_0_MaxmemoryPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e245dd92ac6ab7ee, []int{0, 0}
}

// Fields and structure of `RedisConfig` reflects Redis configuration file
// parameters.
type RedisConfig5_0 struct {
	// Redis key eviction policy for a dataset that reaches maximum memory,
	// available to the host. Redis maxmemory setting depends on Managed
	// Service for Redis [host class](/docs/managed-redis/concepts/instance-types).
	//
	// All policies are described in detail in [Redis documentation](https://redis.io/topics/lru-cache).
	MaxmemoryPolicy RedisConfig5_0_MaxmemoryPolicy `protobuf:"varint,1,opt,name=maxmemory_policy,json=maxmemoryPolicy,proto3,enum=yandex.cloud.mdb.redis.v1alpha.config.RedisConfig5_0_MaxmemoryPolicy" json:"maxmemory_policy,omitempty"`
	// Time that Redis keeps the connection open while the client is idle.
	// If no new command is sent during that time, the connection is closed.
	Timeout *wrappers.Int64Value `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Authentication password.
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisConfig5_0) Reset()         { *m = RedisConfig5_0{} }
func (m *RedisConfig5_0) String() string { return proto.CompactTextString(m) }
func (*RedisConfig5_0) ProtoMessage()    {}
func (*RedisConfig5_0) Descriptor() ([]byte, []int) {
	return fileDescriptor_e245dd92ac6ab7ee, []int{0}
}

func (m *RedisConfig5_0) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisConfig5_0.Unmarshal(m, b)
}
func (m *RedisConfig5_0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisConfig5_0.Marshal(b, m, deterministic)
}
func (m *RedisConfig5_0) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisConfig5_0.Merge(m, src)
}
func (m *RedisConfig5_0) XXX_Size() int {
	return xxx_messageInfo_RedisConfig5_0.Size(m)
}
func (m *RedisConfig5_0) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisConfig5_0.DiscardUnknown(m)
}

var xxx_messageInfo_RedisConfig5_0 proto.InternalMessageInfo

func (m *RedisConfig5_0) GetMaxmemoryPolicy() RedisConfig5_0_MaxmemoryPolicy {
	if m != nil {
		return m.MaxmemoryPolicy
	}
	return RedisConfig5_0_MAXMEMORY_POLICY_UNSPECIFIED
}

func (m *RedisConfig5_0) GetTimeout() *wrappers.Int64Value {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RedisConfig5_0) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RedisConfigSet5_0 struct {
	// Effective settings for a Redis 5.0 cluster (a combination of settings
	// defined in [user_config] and [default_config]).
	EffectiveConfig *RedisConfig5_0 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a Redis 5.0 cluster.
	UserConfig *RedisConfig5_0 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a Redis 5.0 cluster.
	DefaultConfig        *RedisConfig5_0 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RedisConfigSet5_0) Reset()         { *m = RedisConfigSet5_0{} }
func (m *RedisConfigSet5_0) String() string { return proto.CompactTextString(m) }
func (*RedisConfigSet5_0) ProtoMessage()    {}
func (*RedisConfigSet5_0) Descriptor() ([]byte, []int) {
	return fileDescriptor_e245dd92ac6ab7ee, []int{1}
}

func (m *RedisConfigSet5_0) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisConfigSet5_0.Unmarshal(m, b)
}
func (m *RedisConfigSet5_0) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisConfigSet5_0.Marshal(b, m, deterministic)
}
func (m *RedisConfigSet5_0) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisConfigSet5_0.Merge(m, src)
}
func (m *RedisConfigSet5_0) XXX_Size() int {
	return xxx_messageInfo_RedisConfigSet5_0.Size(m)
}
func (m *RedisConfigSet5_0) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisConfigSet5_0.DiscardUnknown(m)
}

var xxx_messageInfo_RedisConfigSet5_0 proto.InternalMessageInfo

func (m *RedisConfigSet5_0) GetEffectiveConfig() *RedisConfig5_0 {
	if m != nil {
		return m.EffectiveConfig
	}
	return nil
}

func (m *RedisConfigSet5_0) GetUserConfig() *RedisConfig5_0 {
	if m != nil {
		return m.UserConfig
	}
	return nil
}

func (m *RedisConfigSet5_0) GetDefaultConfig() *RedisConfig5_0 {
	if m != nil {
		return m.DefaultConfig
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.mdb.redis.v1alpha.config.RedisConfig5_0_MaxmemoryPolicy", RedisConfig5_0_MaxmemoryPolicy_name, RedisConfig5_0_MaxmemoryPolicy_value)
	proto.RegisterType((*RedisConfig5_0)(nil), "yandex.cloud.mdb.redis.v1alpha.config.RedisConfig5_0")
	proto.RegisterType((*RedisConfigSet5_0)(nil), "yandex.cloud.mdb.redis.v1alpha.config.RedisConfigSet5_0")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/redis/v1alpha/config/redis5_0.proto", fileDescriptor_e245dd92ac6ab7ee)
}

var fileDescriptor_e245dd92ac6ab7ee = []byte{
	// 537 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xfb, 0x6e, 0xd2, 0x50,
	0x1c, 0xc7, 0x2d, 0xe8, 0x36, 0x0f, 0x0a, 0xf5, 0xf8, 0x0f, 0x99, 0x97, 0x20, 0xea, 0x82, 0xba,
	0x9e, 0x02, 0x0e, 0x33, 0xe3, 0x15, 0x58, 0x49, 0x1a, 0x0b, 0x25, 0xe5, 0x12, 0x99, 0x97, 0x7a,
	0xa0, 0x87, 0xae, 0x49, 0xcb, 0x69, 0x7a, 0x61, 0xc3, 0xcb, 0x13, 0xf9, 0x1a, 0x4b, 0x7c, 0x1f,
	0x9f, 0xc0, 0xd0, 0x43, 0x89, 0x35, 0x31, 0x59, 0xb6, 0x7f, 0xbf, 0xfd, 0x7e, 0x3e, 0x3d, 0xa7,
	0xbf, 0xfe, 0xc0, 0xde, 0x02, 0xcf, 0x0c, 0x72, 0x22, 0x4e, 0x6c, 0x1a, 0x1a, 0xa2, 0x63, 0x8c,
	0x45, 0x8f, 0x18, 0x96, 0x2f, 0xce, 0x2b, 0xd8, 0x76, 0x8f, 0xb0, 0x38, 0xa1, 0xb3, 0xa9, 0x65,
	0xb2, 0xb0, 0xa6, 0x97, 0x91, 0xeb, 0xd1, 0x80, 0xc2, 0x87, 0x8c, 0x42, 0x11, 0x85, 0x1c, 0x63,
	0x8c, 0xa2, 0x02, 0x5a, 0x51, 0x88, 0x51, 0xdb, 0x77, 0x4d, 0x4a, 0x4d, 0x9b, 0x88, 0x11, 0x34,
	0x0e, 0xa7, 0xe2, 0xb1, 0x87, 0x5d, 0x97, 0x78, 0x3e, 0xd3, 0x6c, 0xdf, 0x49, 0xbc, 0x7c, 0x8e,
	0x6d, 0xcb, 0xc0, 0x81, 0x45, 0x67, 0xec, 0x71, 0xf1, 0x34, 0x0d, 0xb2, 0xda, 0xd2, 0xdb, 0x8c,
	0x74, 0x35, 0xbd, 0x0c, 0x5d, 0xc0, 0x3b, 0xf8, 0xc4, 0x21, 0x0e, 0xf5, 0x16, 0xba, 0x4b, 0x6d,
	0x6b, 0xb2, 0xc8, 0x73, 0x05, 0xae, 0x94, 0xad, 0x4a, 0xe8, 0x4c, 0x67, 0x42, 0x49, 0x21, 0x6a,
	0xc7, 0xb6, 0x6e, 0x24, 0xd3, 0x72, 0x4e, 0x32, 0x80, 0x35, 0xb0, 0x19, 0x58, 0x0e, 0xa1, 0x61,
	0x90, 0x4f, 0x15, 0xb8, 0x52, 0xa6, 0x7a, 0x0b, 0xb1, 0x5b, 0xa1, 0xf8, 0x56, 0x48, 0x9e, 0x05,
	0xcf, 0xf6, 0x86, 0xd8, 0x0e, 0x89, 0x16, 0x77, 0x61, 0x03, 0x6c, 0xb9, 0xd8, 0xf7, 0x8f, 0xa9,
	0x67, 0xe4, 0xd3, 0x05, 0xae, 0x74, 0xb5, 0xb1, 0xf3, 0xfb, 0x57, 0xa5, 0xf8, 0x01, 0x0b, 0x5f,
	0xeb, 0xc2, 0x61, 0x59, 0x78, 0xfe, 0xf6, 0xd5, 0x93, 0x37, 0x8f, 0xd1, 0xee, 0xbd, 0x9d, 0xfb,
	0x0f, 0x3e, 0xbf, 0x7c, 0xad, 0x0b, 0x9f, 0xbe, 0xed, 0xef, 0x56, 0xaa, 0xfb, 0x3f, 0xb4, 0x35,
	0x57, 0x3c, 0xe5, 0x40, 0xee, 0x9f, 0xf3, 0xc1, 0x02, 0xb8, 0xdd, 0xae, 0xbf, 0x6f, 0x4b, 0x6d,
	0x55, 0x1b, 0xe9, 0x5d, 0x55, 0x91, 0x9b, 0x23, 0x7d, 0xd0, 0xe9, 0x75, 0xa5, 0xa6, 0xdc, 0x92,
	0xa5, 0x03, 0xfe, 0x12, 0xe4, 0xc1, 0xb5, 0xa1, 0xaa, 0xd4, 0xfb, 0xb2, 0x22, 0xe9, 0x8a, 0x36,
	0xe0, 0x39, 0x98, 0x03, 0x99, 0xba, 0xa2, 0xbc, 0x93, 0x46, 0xbd, 0x28, 0x48, 0x25, 0x2b, 0xad,
	0x01, 0x9f, 0x4e, 0x54, 0x5a, 0x03, 0xfe, 0x32, 0xbc, 0x09, 0x72, 0xeb, 0x8a, 0x56, 0xef, 0x1c,
	0xa8, 0x6d, 0xfe, 0x0a, 0x84, 0x20, 0x1b, 0xb7, 0x56, 0xd9, 0x46, 0xc2, 0xd5, 0xef, 0x2b, 0xfc,
	0x26, 0xcc, 0x02, 0xd0, 0x51, 0xa5, 0xa1, 0xdc, 0xec, 0xcb, 0x6a, 0x87, 0xdf, 0x2a, 0xfe, 0x4c,
	0x81, 0x1b, 0x7f, 0x7d, 0xf5, 0x1e, 0x09, 0x96, 0x93, 0xfc, 0x02, 0x78, 0x32, 0x9d, 0x92, 0x49,
	0x60, 0xcd, 0x89, 0xce, 0x66, 0x13, 0x4d, 0x32, 0x53, 0xad, 0x9d, 0x6b, 0x92, 0x5a, 0x6e, 0xad,
	0x63, 0x19, 0x1c, 0x82, 0x4c, 0xe8, 0x13, 0x2f, 0x96, 0xa7, 0x2e, 0x22, 0x07, 0x4b, 0xd3, 0xca,
	0xfb, 0x11, 0x64, 0x0d, 0x32, 0xc5, 0xa1, 0x1d, 0xc4, 0xea, 0xf4, 0x45, 0xd4, 0xd7, 0x57, 0x32,
	0x96, 0x34, 0xbe, 0x83, 0x47, 0x09, 0x0d, 0x76, 0xad, 0xff, 0xaa, 0x0e, 0x55, 0xd3, 0x0a, 0x8e,
	0xc2, 0x31, 0x9a, 0x50, 0x47, 0x64, 0x94, 0xc0, 0x76, 0xc9, 0xa4, 0x82, 0x49, 0x66, 0xd1, 0x1f,
	0x2a, 0x9e, 0x69, 0xc3, 0x5f, 0x44, 0xe1, 0x78, 0x23, 0x42, 0x9e, 0xfe, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xcc, 0xe8, 0xca, 0x05, 0x17, 0x04, 0x00, 0x00,
}
