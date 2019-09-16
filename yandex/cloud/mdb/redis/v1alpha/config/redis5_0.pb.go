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
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x6d, 0x6f, 0xd2, 0x5e,
	0x18, 0xc6, 0xff, 0x85, 0xbf, 0xdb, 0x3c, 0x28, 0xd4, 0xe3, 0x1b, 0x32, 0x1f, 0x82, 0xa8, 0x0b,
	0xd1, 0xf5, 0x14, 0x70, 0x98, 0x19, 0x1f, 0x81, 0x95, 0xa4, 0xb1, 0x50, 0x2c, 0x0f, 0x91, 0xf9,
	0x50, 0x0f, 0xf4, 0xd0, 0x35, 0x69, 0x39, 0x4d, 0x69, 0xd9, 0xd0, 0xf8, 0x89, 0xfc, 0x1a, 0x4b,
	0xfc, 0x3e, 0x7e, 0x02, 0xc3, 0x39, 0x94, 0x58, 0x5f, 0x2d, 0xee, 0xed, 0xd5, 0xeb, 0xfa, 0xf5,
	0xbe, 0x7b, 0xf5, 0x06, 0x07, 0x4b, 0x3c, 0xb3, 0xc8, 0x99, 0x3c, 0x71, 0x69, 0x64, 0xc9, 0x9e,
	0x35, 0x96, 0x03, 0x62, 0x39, 0x73, 0x79, 0x51, 0xc1, 0xae, 0x7f, 0x82, 0xe5, 0x09, 0x9d, 0x4d,
	0x1d, 0x9b, 0x8b, 0x35, 0xb3, 0x8c, 0xfc, 0x80, 0x86, 0x14, 0x3e, 0xe4, 0x29, 0xc4, 0x52, 0xc8,
	0xb3, 0xc6, 0x88, 0x19, 0xd0, 0x3a, 0x85, 0x78, 0x6a, 0xf7, 0xae, 0x4d, 0xa9, 0xed, 0x12, 0x99,
	0x85, 0xc6, 0xd1, 0x54, 0x3e, 0x0d, 0xb0, 0xef, 0x93, 0x60, 0xce, 0x31, 0xbb, 0x77, 0x12, 0x2f,
	0x5f, 0x60, 0xd7, 0xb1, 0x70, 0xe8, 0xd0, 0x19, 0x7f, 0x5c, 0x3c, 0x4f, 0x83, 0xac, 0xb1, 0xe2,
	0x36, 0x19, 0xae, 0x66, 0x96, 0xa1, 0x0f, 0x44, 0x0f, 0x9f, 0x79, 0xc4, 0xa3, 0xc1, 0xd2, 0xf4,
	0xa9, 0xeb, 0x4c, 0x96, 0x79, 0xa1, 0x20, 0x94, 0xb2, 0x55, 0x05, 0x5d, 0x68, 0x26, 0x94, 0x04,
	0xa2, 0x76, 0x4c, 0xeb, 0x32, 0x98, 0x91, 0xf3, 0x92, 0x02, 0xac, 0x81, 0xed, 0xd0, 0xf1, 0x08,
	0x8d, 0xc2, 0x7c, 0xaa, 0x20, 0x94, 0x32, 0xd5, 0x5b, 0x88, 0x6f, 0x85, 0xe2, 0xad, 0x90, 0x3a,
	0x0b, 0x9f, 0x1e, 0x0c, 0xb1, 0x1b, 0x11, 0x23, 0xf6, 0xc2, 0x06, 0xd8, 0xf1, 0xf1, 0x7c, 0x7e,
	0x4a, 0x03, 0x2b, 0x9f, 0x2e, 0x08, 0xa5, 0xab, 0x8d, 0xbd, 0x5f, 0x3f, 0x2b, 0xc5, 0x0f, 0x58,
	0xfa, 0x5a, 0x97, 0x8e, 0xcb, 0xd2, 0xb3, 0x37, 0x2f, 0x1f, 0xbf, 0x7e, 0x84, 0xf6, 0xef, 0xed,
	0xdd, 0x7f, 0xf0, 0xf9, 0xc5, 0x2b, 0x53, 0xfa, 0xf4, 0xed, 0x70, 0xbf, 0x52, 0x3d, 0xfc, 0x6e,
	0x6c, 0x72, 0xc5, 0x73, 0x01, 0xe4, 0xfe, 0x9a, 0x0f, 0x16, 0xc0, 0xed, 0x76, 0xfd, 0x7d, 0x5b,
	0x69, 0xeb, 0xc6, 0xc8, 0xec, 0xea, 0x9a, 0xda, 0x1c, 0x99, 0x83, 0x4e, 0xaf, 0xab, 0x34, 0xd5,
	0x96, 0xaa, 0x1c, 0x89, 0xff, 0x41, 0x11, 0x5c, 0x1b, 0xea, 0x5a, 0xbd, 0xaf, 0x6a, 0x8a, 0xa9,
	0x19, 0x03, 0x51, 0x80, 0x39, 0x90, 0xa9, 0x6b, 0xda, 0x5b, 0x65, 0xd4, 0x63, 0x42, 0x2a, 0x69,
	0x69, 0x0d, 0xc4, 0x74, 0xc2, 0xd2, 0x1a, 0x88, 0xff, 0xc3, 0x9b, 0x20, 0xb7, 0xb1, 0x18, 0xf5,
	0xce, 0x91, 0xde, 0x16, 0xaf, 0x40, 0x08, 0xb2, 0xb1, 0x6b, 0xad, 0x6d, 0x25, 0x58, 0xfd, 0xbe,
	0x26, 0x6e, 0xc3, 0x2c, 0x00, 0x1d, 0x5d, 0x19, 0xaa, 0xcd, 0xbe, 0xaa, 0x77, 0xc4, 0x9d, 0xe2,
	0x8f, 0x14, 0xb8, 0xf1, 0xc7, 0x57, 0xef, 0x91, 0x70, 0xd5, 0xe4, 0x17, 0x20, 0x92, 0xe9, 0x94,
	0x4c, 0x42, 0x67, 0x41, 0x4c, 0xde, 0x0d, 0x6b, 0x32, 0x53, 0xad, 0xfd, 0x53, 0x93, 0x46, 0x6e,
	0x83, 0xe3, 0x1a, 0x1c, 0x82, 0x4c, 0x34, 0x27, 0x41, 0x0c, 0x4f, 0x5d, 0x06, 0x0e, 0x56, 0xa4,
	0x35, 0xf7, 0x23, 0xc8, 0x5a, 0x64, 0x8a, 0x23, 0x37, 0x8c, 0xd1, 0xe9, 0xcb, 0xa0, 0xaf, 0xaf,
	0x61, 0x5c, 0x69, 0xbc, 0x3b, 0xd6, 0x6d, 0x27, 0x3c, 0x89, 0xc6, 0x68, 0x42, 0x3d, 0x99, 0x13,
	0x25, 0x7e, 0x20, 0x36, 0x95, 0x6c, 0x32, 0x63, 0xbf, 0x9d, 0x7c, 0xa1, 0xb3, 0x7d, 0xce, 0xc4,
	0xf1, 0x16, 0x8b, 0x3c, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x76, 0x80, 0xed, 0x6a, 0xec, 0x03,
	0x00, 0x00,
}
