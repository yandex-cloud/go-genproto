// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/redis/v1/config/redis6_0.proto

package redis

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

type RedisConfig6_0_MaxmemoryPolicy int32

const (
	RedisConfig6_0_MAXMEMORY_POLICY_UNSPECIFIED RedisConfig6_0_MaxmemoryPolicy = 0
	// Try to remove less recently used (LRU) keys with `expire set`.
	RedisConfig6_0_VOLATILE_LRU RedisConfig6_0_MaxmemoryPolicy = 1
	// Remove less recently used (LRU) keys.
	RedisConfig6_0_ALLKEYS_LRU RedisConfig6_0_MaxmemoryPolicy = 2
	// Try to remove least frequently used (LFU) keys with `expire set`.
	RedisConfig6_0_VOLATILE_LFU RedisConfig6_0_MaxmemoryPolicy = 3
	// Remove least frequently used (LFU) keys.
	RedisConfig6_0_ALLKEYS_LFU RedisConfig6_0_MaxmemoryPolicy = 4
	// Try to remove keys with `expire set` randomly.
	RedisConfig6_0_VOLATILE_RANDOM RedisConfig6_0_MaxmemoryPolicy = 5
	// Remove keys randomly.
	RedisConfig6_0_ALLKEYS_RANDOM RedisConfig6_0_MaxmemoryPolicy = 6
	// Try to remove less recently used (LRU) keys with `expire set`
	// and shorter TTL first.
	RedisConfig6_0_VOLATILE_TTL RedisConfig6_0_MaxmemoryPolicy = 7
	// Return errors when memory limit was reached and commands could require
	// more memory to be used.
	RedisConfig6_0_NOEVICTION RedisConfig6_0_MaxmemoryPolicy = 8
)

// Enum value maps for RedisConfig6_0_MaxmemoryPolicy.
var (
	RedisConfig6_0_MaxmemoryPolicy_name = map[int32]string{
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
	RedisConfig6_0_MaxmemoryPolicy_value = map[string]int32{
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
)

func (x RedisConfig6_0_MaxmemoryPolicy) Enum() *RedisConfig6_0_MaxmemoryPolicy {
	p := new(RedisConfig6_0_MaxmemoryPolicy)
	*p = x
	return p
}

func (x RedisConfig6_0_MaxmemoryPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RedisConfig6_0_MaxmemoryPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes[0].Descriptor()
}

func (RedisConfig6_0_MaxmemoryPolicy) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes[0]
}

func (x RedisConfig6_0_MaxmemoryPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RedisConfig6_0_MaxmemoryPolicy.Descriptor instead.
func (RedisConfig6_0_MaxmemoryPolicy) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{0, 0}
}

// Fields and structure of `RedisConfig` reflects Redis configuration file
// parameters.
type RedisConfig6_0 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Redis key eviction policy for a dataset that reaches maximum memory,
	// available to the host. Redis maxmemory setting depends on Managed
	// Service for Redis [host class](/docs/managed-redis/concepts/instance-types).
	//
	// All policies are described in detail in [Redis documentation](https://redis.io/topics/lru-cache).
	MaxmemoryPolicy RedisConfig6_0_MaxmemoryPolicy `protobuf:"varint,1,opt,name=maxmemory_policy,json=maxmemoryPolicy,proto3,enum=yandex.cloud.mdb.redis.v1.config.RedisConfig6_0_MaxmemoryPolicy" json:"maxmemory_policy,omitempty"`
	// Time that Redis keeps the connection open while the client is idle.
	// If no new command is sent during that time, the connection is closed.
	Timeout *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// Authentication password.
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	// Number of database buckets on a single redis-server process.
	Databases *wrapperspb.Int64Value `protobuf:"bytes,4,opt,name=databases,proto3" json:"databases,omitempty"`
	// Threshold for logging slow requests to server in microseconds (log only slower than it).
	SlowlogLogSlowerThan *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=slowlog_log_slower_than,json=slowlogLogSlowerThan,proto3" json:"slowlog_log_slower_than,omitempty"`
	// Max slow requests number to log.
	SlowlogMaxLen *wrapperspb.Int64Value `protobuf:"bytes,6,opt,name=slowlog_max_len,json=slowlogMaxLen,proto3" json:"slowlog_max_len,omitempty"`
	// String setting for pub\sub functionality.
	NotifyKeyspaceEvents string `protobuf:"bytes,7,opt,name=notify_keyspace_events,json=notifyKeyspaceEvents,proto3" json:"notify_keyspace_events,omitempty"`
	// Redis connection output buffers limits for pubsub operations.
	ClientOutputBufferLimitPubsub *RedisConfig6_0_ClientOutputBufferLimit `protobuf:"bytes,8,opt,name=client_output_buffer_limit_pubsub,json=clientOutputBufferLimitPubsub,proto3" json:"client_output_buffer_limit_pubsub,omitempty"`
	// Redis connection output buffers limits for clients.
	ClientOutputBufferLimitNormal *RedisConfig6_0_ClientOutputBufferLimit `protobuf:"bytes,9,opt,name=client_output_buffer_limit_normal,json=clientOutputBufferLimitNormal,proto3" json:"client_output_buffer_limit_normal,omitempty"`
	unknownFields                 protoimpl.UnknownFields
	sizeCache                     protoimpl.SizeCache
}

func (x *RedisConfig6_0) Reset() {
	*x = RedisConfig6_0{}
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedisConfig6_0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfig6_0) ProtoMessage() {}

func (x *RedisConfig6_0) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfig6_0.ProtoReflect.Descriptor instead.
func (*RedisConfig6_0) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{0}
}

func (x *RedisConfig6_0) GetMaxmemoryPolicy() RedisConfig6_0_MaxmemoryPolicy {
	if x != nil {
		return x.MaxmemoryPolicy
	}
	return RedisConfig6_0_MAXMEMORY_POLICY_UNSPECIFIED
}

func (x *RedisConfig6_0) GetTimeout() *wrapperspb.Int64Value {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RedisConfig6_0) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RedisConfig6_0) GetDatabases() *wrapperspb.Int64Value {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *RedisConfig6_0) GetSlowlogLogSlowerThan() *wrapperspb.Int64Value {
	if x != nil {
		return x.SlowlogLogSlowerThan
	}
	return nil
}

func (x *RedisConfig6_0) GetSlowlogMaxLen() *wrapperspb.Int64Value {
	if x != nil {
		return x.SlowlogMaxLen
	}
	return nil
}

func (x *RedisConfig6_0) GetNotifyKeyspaceEvents() string {
	if x != nil {
		return x.NotifyKeyspaceEvents
	}
	return ""
}

func (x *RedisConfig6_0) GetClientOutputBufferLimitPubsub() *RedisConfig6_0_ClientOutputBufferLimit {
	if x != nil {
		return x.ClientOutputBufferLimitPubsub
	}
	return nil
}

func (x *RedisConfig6_0) GetClientOutputBufferLimitNormal() *RedisConfig6_0_ClientOutputBufferLimit {
	if x != nil {
		return x.ClientOutputBufferLimitNormal
	}
	return nil
}

type RedisConfigSet6_0 struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Effective settings for a Redis 6.0 cluster (a combination of settings
	// defined in [user_config] and [default_config]).
	EffectiveConfig *RedisConfig6_0 `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a Redis 6.0 cluster.
	UserConfig *RedisConfig6_0 `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a Redis 6.0 cluster.
	DefaultConfig *RedisConfig6_0 `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RedisConfigSet6_0) Reset() {
	*x = RedisConfigSet6_0{}
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedisConfigSet6_0) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfigSet6_0) ProtoMessage() {}

func (x *RedisConfigSet6_0) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfigSet6_0.ProtoReflect.Descriptor instead.
func (*RedisConfigSet6_0) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{1}
}

func (x *RedisConfigSet6_0) GetEffectiveConfig() *RedisConfig6_0 {
	if x != nil {
		return x.EffectiveConfig
	}
	return nil
}

func (x *RedisConfigSet6_0) GetUserConfig() *RedisConfig6_0 {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

func (x *RedisConfigSet6_0) GetDefaultConfig() *RedisConfig6_0 {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

type RedisConfig6_0_ClientOutputBufferLimit struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Total limit in bytes.
	HardLimit *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=hard_limit,json=hardLimit,proto3" json:"hard_limit,omitempty"`
	// Limit in bytes during certain time period.
	SoftLimit *wrapperspb.Int64Value `protobuf:"bytes,3,opt,name=soft_limit,json=softLimit,proto3" json:"soft_limit,omitempty"`
	// Seconds for soft limit.
	SoftSeconds   *wrapperspb.Int64Value `protobuf:"bytes,5,opt,name=soft_seconds,json=softSeconds,proto3" json:"soft_seconds,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RedisConfig6_0_ClientOutputBufferLimit) Reset() {
	*x = RedisConfig6_0_ClientOutputBufferLimit{}
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedisConfig6_0_ClientOutputBufferLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfig6_0_ClientOutputBufferLimit) ProtoMessage() {}

func (x *RedisConfig6_0_ClientOutputBufferLimit) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfig6_0_ClientOutputBufferLimit.ProtoReflect.Descriptor instead.
func (*RedisConfig6_0_ClientOutputBufferLimit) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RedisConfig6_0_ClientOutputBufferLimit) GetHardLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.HardLimit
	}
	return nil
}

func (x *RedisConfig6_0_ClientOutputBufferLimit) GetSoftLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.SoftLimit
	}
	return nil
}

func (x *RedisConfig6_0_ClientOutputBufferLimit) GetSoftSeconds() *wrapperspb.Int64Value {
	if x != nil {
		return x.SoftSeconds
	}
	return nil
}

var File_yandex_cloud_mdb_redis_v1_config_redis6_0_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc = "" +
	"\n" +
	"/yandex/cloud/mdb/redis/v1/config/redis6_0.proto\x12 yandex.cloud.mdb.redis.v1.config\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1dyandex/cloud/validation.proto\"\xa3\n" +
	"\n" +
	"\x0eRedisConfig6_0\x12k\n" +
	"\x10maxmemory_policy\x18\x01 \x01(\x0e2@.yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.MaxmemoryPolicyR\x0fmaxmemoryPolicy\x125\n" +
	"\atimeout\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\atimeout\x12B\n" +
	"\bpassword\x18\x03 \x01(\tB&\xf2\xc71\"[a-zA-Z0-9@=+?*.,!&#$^<>_-]{8,128}R\bpassword\x12A\n" +
	"\tdatabases\x18\x04 \x01(\v2\x1b.google.protobuf.Int64ValueB\x06\xfa\xc71\x02>0R\tdatabases\x12[\n" +
	"\x17slowlog_log_slower_than\x18\x05 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\x14slowlogLogSlowerThan\x12L\n" +
	"\x0fslowlog_max_len\x18\x06 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\rslowlogMaxLen\x12O\n" +
	"\x16notify_keyspace_events\x18\a \x01(\tB\x19\xf2\xc71\x15[KEg$lshzxeAtm]{0,13}R\x14notifyKeyspaceEvents\x12\x92\x01\n" +
	"!client_output_buffer_limit_pubsub\x18\b \x01(\v2H.yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimitR\x1dclientOutputBufferLimitPubsub\x12\x92\x01\n" +
	"!client_output_buffer_limit_normal\x18\t \x01(\v2H.yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimitR\x1dclientOutputBufferLimitNormal\x1a\xf8\x01\n" +
	"\x17ClientOutputBufferLimit\x12C\n" +
	"\n" +
	"hard_limit\x18\x01 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\thardLimit\x12C\n" +
	"\n" +
	"soft_limit\x18\x03 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\tsoftLimit\x12G\n" +
	"\fsoft_seconds\x18\x05 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\vsoftSecondsJ\x04\b\x04\x10\x05J\x04\b\x02\x10\x03\"\xc4\x01\n" +
	"\x0fMaxmemoryPolicy\x12 \n" +
	"\x1cMAXMEMORY_POLICY_UNSPECIFIED\x10\x00\x12\x10\n" +
	"\fVOLATILE_LRU\x10\x01\x12\x0f\n" +
	"\vALLKEYS_LRU\x10\x02\x12\x10\n" +
	"\fVOLATILE_LFU\x10\x03\x12\x0f\n" +
	"\vALLKEYS_LFU\x10\x04\x12\x13\n" +
	"\x0fVOLATILE_RANDOM\x10\x05\x12\x12\n" +
	"\x0eALLKEYS_RANDOM\x10\x06\x12\x10\n" +
	"\fVOLATILE_TTL\x10\a\x12\x0e\n" +
	"\n" +
	"NOEVICTION\x10\b\"\x9c\x02\n" +
	"\x11RedisConfigSet6_0\x12[\n" +
	"\x10effective_config\x18\x01 \x01(\v20.yandex.cloud.mdb.redis.v1.config.RedisConfig6_0R\x0feffectiveConfig\x12Q\n" +
	"\vuser_config\x18\x02 \x01(\v20.yandex.cloud.mdb.redis.v1.config.RedisConfig6_0R\n" +
	"userConfig\x12W\n" +
	"\x0edefault_config\x18\x03 \x01(\v20.yandex.cloud.mdb.redis.v1.config.RedisConfig6_0R\rdefaultConfigBr\n" +
	"$yandex.cloud.api.mdb.redis.v1.configZJgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1/config;redisb\x06proto3"

var (
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc), len(file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDescData
}

var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_goTypes = []any{
	(RedisConfig6_0_MaxmemoryPolicy)(0),            // 0: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.MaxmemoryPolicy
	(*RedisConfig6_0)(nil),                         // 1: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	(*RedisConfigSet6_0)(nil),                      // 2: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0
	(*RedisConfig6_0_ClientOutputBufferLimit)(nil), // 3: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimit
	(*wrapperspb.Int64Value)(nil),                  // 4: google.protobuf.Int64Value
}
var file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_depIdxs = []int32{
	0,  // 0: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.maxmemory_policy:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.MaxmemoryPolicy
	4,  // 1: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.timeout:type_name -> google.protobuf.Int64Value
	4,  // 2: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.databases:type_name -> google.protobuf.Int64Value
	4,  // 3: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.slowlog_log_slower_than:type_name -> google.protobuf.Int64Value
	4,  // 4: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.slowlog_max_len:type_name -> google.protobuf.Int64Value
	3,  // 5: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.client_output_buffer_limit_pubsub:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimit
	3,  // 6: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.client_output_buffer_limit_normal:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimit
	1,  // 7: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0.effective_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	1,  // 8: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0.user_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	1,  // 9: yandex.cloud.mdb.redis.v1.config.RedisConfigSet6_0.default_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig6_0
	4,  // 10: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimit.hard_limit:type_name -> google.protobuf.Int64Value
	4,  // 11: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimit.soft_limit:type_name -> google.protobuf.Int64Value
	4,  // 12: yandex.cloud.mdb.redis.v1.config.RedisConfig6_0.ClientOutputBufferLimit.soft_seconds:type_name -> google.protobuf.Int64Value
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_init() }
func file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_init() {
	if File_yandex_cloud_mdb_redis_v1_config_redis6_0_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc), len(file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_redis_v1_config_redis6_0_proto = out.File
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_goTypes = nil
	file_yandex_cloud_mdb_redis_v1_config_redis6_0_proto_depIdxs = nil
}
