// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/redis/v1/config/redis.proto

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

type RedisConfig_MaxmemoryPolicy int32

const (
	RedisConfig_MAXMEMORY_POLICY_UNSPECIFIED RedisConfig_MaxmemoryPolicy = 0
	// Try to remove less recently used (LRU) keys with `expire set`.
	RedisConfig_VOLATILE_LRU RedisConfig_MaxmemoryPolicy = 1
	// Remove less recently used (LRU) keys.
	RedisConfig_ALLKEYS_LRU RedisConfig_MaxmemoryPolicy = 2
	// Try to remove least frequently used (LFU) keys with `expire set`.
	RedisConfig_VOLATILE_LFU RedisConfig_MaxmemoryPolicy = 3
	// Remove least frequently used (LFU) keys.
	RedisConfig_ALLKEYS_LFU RedisConfig_MaxmemoryPolicy = 4
	// Try to remove keys with `expire set` randomly.
	RedisConfig_VOLATILE_RANDOM RedisConfig_MaxmemoryPolicy = 5
	// Remove keys randomly.
	RedisConfig_ALLKEYS_RANDOM RedisConfig_MaxmemoryPolicy = 6
	// Try to remove less recently used (LRU) keys with `expire set`
	// and shorter TTL first.
	RedisConfig_VOLATILE_TTL RedisConfig_MaxmemoryPolicy = 7
	// Return errors when memory limit was reached and commands could require
	// more memory to be used.
	RedisConfig_NOEVICTION RedisConfig_MaxmemoryPolicy = 8
)

// Enum value maps for RedisConfig_MaxmemoryPolicy.
var (
	RedisConfig_MaxmemoryPolicy_name = map[int32]string{
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
	RedisConfig_MaxmemoryPolicy_value = map[string]int32{
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

func (x RedisConfig_MaxmemoryPolicy) Enum() *RedisConfig_MaxmemoryPolicy {
	p := new(RedisConfig_MaxmemoryPolicy)
	*p = x
	return p
}

func (x RedisConfig_MaxmemoryPolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RedisConfig_MaxmemoryPolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_mdb_redis_v1_config_redis_proto_enumTypes[0].Descriptor()
}

func (RedisConfig_MaxmemoryPolicy) Type() protoreflect.EnumType {
	return &file_yandex_cloud_mdb_redis_v1_config_redis_proto_enumTypes[0]
}

func (x RedisConfig_MaxmemoryPolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RedisConfig_MaxmemoryPolicy.Descriptor instead.
func (RedisConfig_MaxmemoryPolicy) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescGZIP(), []int{0, 0}
}

// Fields and structure of `RedisConfig` reflects Redis configuration file
// parameters.
type RedisConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Redis key eviction policy for a dataset that reaches maximum memory,
	// available to the host. Redis maxmemory setting depends on Managed
	// Service for Redis [host class](/docs/managed-redis/concepts/instance-types).
	//
	// All policies are described in detail in [Redis documentation](https://redis.io/topics/lru-cache).
	MaxmemoryPolicy RedisConfig_MaxmemoryPolicy `protobuf:"varint,1,opt,name=maxmemory_policy,json=maxmemoryPolicy,proto3,enum=yandex.cloud.mdb.redis.v1.config.RedisConfig_MaxmemoryPolicy" json:"maxmemory_policy,omitempty"`
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
	ClientOutputBufferLimitPubsub *RedisConfig_ClientOutputBufferLimit `protobuf:"bytes,8,opt,name=client_output_buffer_limit_pubsub,json=clientOutputBufferLimitPubsub,proto3" json:"client_output_buffer_limit_pubsub,omitempty"`
	// Redis connection output buffers limits for clients.
	ClientOutputBufferLimitNormal *RedisConfig_ClientOutputBufferLimit `protobuf:"bytes,9,opt,name=client_output_buffer_limit_normal,json=clientOutputBufferLimitNormal,proto3" json:"client_output_buffer_limit_normal,omitempty"`
	// Redis maxmemory percent
	MaxmemoryPercent *wrapperspb.Int64Value `protobuf:"bytes,10,opt,name=maxmemory_percent,json=maxmemoryPercent,proto3" json:"maxmemory_percent,omitempty"`
	// Maximum time in milliseconds for Lua scripts, 0 - disabled mechanism
	LuaTimeLimit *wrapperspb.Int64Value `protobuf:"bytes,11,opt,name=lua_time_limit,json=luaTimeLimit,proto3" json:"lua_time_limit,omitempty"`
	// Replication backlog size as a percentage of flavor maxmemory
	ReplBacklogSizePercent *wrapperspb.Int64Value `protobuf:"bytes,12,opt,name=repl_backlog_size_percent,json=replBacklogSizePercent,proto3" json:"repl_backlog_size_percent,omitempty"`
	// Controls whether all hash slots must be covered by nodes
	ClusterRequireFullCoverage *wrapperspb.BoolValue `protobuf:"bytes,13,opt,name=cluster_require_full_coverage,json=clusterRequireFullCoverage,proto3" json:"cluster_require_full_coverage,omitempty"`
	// Allows read operations when cluster is down
	ClusterAllowReadsWhenDown *wrapperspb.BoolValue `protobuf:"bytes,14,opt,name=cluster_allow_reads_when_down,json=clusterAllowReadsWhenDown,proto3" json:"cluster_allow_reads_when_down,omitempty"`
	// Permits Pub/Sub shard operations when cluster is down
	ClusterAllowPubsubshardWhenDown *wrapperspb.BoolValue `protobuf:"bytes,15,opt,name=cluster_allow_pubsubshard_when_down,json=clusterAllowPubsubshardWhenDown,proto3" json:"cluster_allow_pubsubshard_when_down,omitempty"`
	// The time, in minutes, that must elapse in order for the key counter to be divided by two (or decremented if it has a value less <= 10)
	LfuDecayTime *wrapperspb.Int64Value `protobuf:"bytes,16,opt,name=lfu_decay_time,json=lfuDecayTime,proto3" json:"lfu_decay_time,omitempty"`
	// Determines how the frequency counter represents key hits.
	LfuLogFactor *wrapperspb.Int64Value `protobuf:"bytes,17,opt,name=lfu_log_factor,json=lfuLogFactor,proto3" json:"lfu_log_factor,omitempty"`
	// Allows to turn before switchover in RDSync
	TurnBeforeSwitchover *wrapperspb.BoolValue `protobuf:"bytes,18,opt,name=turn_before_switchover,json=turnBeforeSwitchover,proto3" json:"turn_before_switchover,omitempty"`
	// Allows some data to be lost in favor of faster switchover/restart
	AllowDataLoss *wrapperspb.BoolValue `protobuf:"bytes,19,opt,name=allow_data_loss,json=allowDataLoss,proto3" json:"allow_data_loss,omitempty"`
	// Use JIT for lua scripts and functions
	UseLuajit *wrapperspb.BoolValue `protobuf:"bytes,20,opt,name=use_luajit,json=useLuajit,proto3" json:"use_luajit,omitempty"`
	// Allow redis to use io-threads
	IoThreadsAllowed *wrapperspb.BoolValue `protobuf:"bytes,21,opt,name=io_threads_allowed,json=ioThreadsAllowed,proto3" json:"io_threads_allowed,omitempty"`
	// Controls max number of entries in zset before conversion from memory-efficient listpack to CPU-efficient hash table and skiplist
	ZsetMaxListpackEntries *wrapperspb.Int64Value `protobuf:"bytes,22,opt,name=zset_max_listpack_entries,json=zsetMaxListpackEntries,proto3" json:"zset_max_listpack_entries,omitempty"`
	// AOF maximum size as a percentage of disk available
	AofMaxSizePercent *wrapperspb.Int64Value `protobuf:"bytes,23,opt,name=aof_max_size_percent,json=aofMaxSizePercent,proto3" json:"aof_max_size_percent,omitempty"`
	// Enable active (online) memory defragmentation
	Activedefrag  *wrapperspb.BoolValue `protobuf:"bytes,24,opt,name=activedefrag,proto3" json:"activedefrag,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RedisConfig) Reset() {
	*x = RedisConfig{}
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedisConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfig) ProtoMessage() {}

func (x *RedisConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfig.ProtoReflect.Descriptor instead.
func (*RedisConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescGZIP(), []int{0}
}

func (x *RedisConfig) GetMaxmemoryPolicy() RedisConfig_MaxmemoryPolicy {
	if x != nil {
		return x.MaxmemoryPolicy
	}
	return RedisConfig_MAXMEMORY_POLICY_UNSPECIFIED
}

func (x *RedisConfig) GetTimeout() *wrapperspb.Int64Value {
	if x != nil {
		return x.Timeout
	}
	return nil
}

func (x *RedisConfig) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RedisConfig) GetDatabases() *wrapperspb.Int64Value {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *RedisConfig) GetSlowlogLogSlowerThan() *wrapperspb.Int64Value {
	if x != nil {
		return x.SlowlogLogSlowerThan
	}
	return nil
}

func (x *RedisConfig) GetSlowlogMaxLen() *wrapperspb.Int64Value {
	if x != nil {
		return x.SlowlogMaxLen
	}
	return nil
}

func (x *RedisConfig) GetNotifyKeyspaceEvents() string {
	if x != nil {
		return x.NotifyKeyspaceEvents
	}
	return ""
}

func (x *RedisConfig) GetClientOutputBufferLimitPubsub() *RedisConfig_ClientOutputBufferLimit {
	if x != nil {
		return x.ClientOutputBufferLimitPubsub
	}
	return nil
}

func (x *RedisConfig) GetClientOutputBufferLimitNormal() *RedisConfig_ClientOutputBufferLimit {
	if x != nil {
		return x.ClientOutputBufferLimitNormal
	}
	return nil
}

func (x *RedisConfig) GetMaxmemoryPercent() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxmemoryPercent
	}
	return nil
}

func (x *RedisConfig) GetLuaTimeLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.LuaTimeLimit
	}
	return nil
}

func (x *RedisConfig) GetReplBacklogSizePercent() *wrapperspb.Int64Value {
	if x != nil {
		return x.ReplBacklogSizePercent
	}
	return nil
}

func (x *RedisConfig) GetClusterRequireFullCoverage() *wrapperspb.BoolValue {
	if x != nil {
		return x.ClusterRequireFullCoverage
	}
	return nil
}

func (x *RedisConfig) GetClusterAllowReadsWhenDown() *wrapperspb.BoolValue {
	if x != nil {
		return x.ClusterAllowReadsWhenDown
	}
	return nil
}

func (x *RedisConfig) GetClusterAllowPubsubshardWhenDown() *wrapperspb.BoolValue {
	if x != nil {
		return x.ClusterAllowPubsubshardWhenDown
	}
	return nil
}

func (x *RedisConfig) GetLfuDecayTime() *wrapperspb.Int64Value {
	if x != nil {
		return x.LfuDecayTime
	}
	return nil
}

func (x *RedisConfig) GetLfuLogFactor() *wrapperspb.Int64Value {
	if x != nil {
		return x.LfuLogFactor
	}
	return nil
}

func (x *RedisConfig) GetTurnBeforeSwitchover() *wrapperspb.BoolValue {
	if x != nil {
		return x.TurnBeforeSwitchover
	}
	return nil
}

func (x *RedisConfig) GetAllowDataLoss() *wrapperspb.BoolValue {
	if x != nil {
		return x.AllowDataLoss
	}
	return nil
}

func (x *RedisConfig) GetUseLuajit() *wrapperspb.BoolValue {
	if x != nil {
		return x.UseLuajit
	}
	return nil
}

func (x *RedisConfig) GetIoThreadsAllowed() *wrapperspb.BoolValue {
	if x != nil {
		return x.IoThreadsAllowed
	}
	return nil
}

func (x *RedisConfig) GetZsetMaxListpackEntries() *wrapperspb.Int64Value {
	if x != nil {
		return x.ZsetMaxListpackEntries
	}
	return nil
}

func (x *RedisConfig) GetAofMaxSizePercent() *wrapperspb.Int64Value {
	if x != nil {
		return x.AofMaxSizePercent
	}
	return nil
}

func (x *RedisConfig) GetActivedefrag() *wrapperspb.BoolValue {
	if x != nil {
		return x.Activedefrag
	}
	return nil
}

type RedisConfigSet struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Effective settings for a Redis cluster (a combination of settings
	// defined in [user_config] and [default_config]).
	EffectiveConfig *RedisConfig `protobuf:"bytes,1,opt,name=effective_config,json=effectiveConfig,proto3" json:"effective_config,omitempty"`
	// User-defined settings for a Redis cluster.
	UserConfig *RedisConfig `protobuf:"bytes,2,opt,name=user_config,json=userConfig,proto3" json:"user_config,omitempty"`
	// Default configuration for a Redis cluster.
	DefaultConfig *RedisConfig `protobuf:"bytes,3,opt,name=default_config,json=defaultConfig,proto3" json:"default_config,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RedisConfigSet) Reset() {
	*x = RedisConfigSet{}
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedisConfigSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfigSet) ProtoMessage() {}

func (x *RedisConfigSet) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfigSet.ProtoReflect.Descriptor instead.
func (*RedisConfigSet) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescGZIP(), []int{1}
}

func (x *RedisConfigSet) GetEffectiveConfig() *RedisConfig {
	if x != nil {
		return x.EffectiveConfig
	}
	return nil
}

func (x *RedisConfigSet) GetUserConfig() *RedisConfig {
	if x != nil {
		return x.UserConfig
	}
	return nil
}

func (x *RedisConfigSet) GetDefaultConfig() *RedisConfig {
	if x != nil {
		return x.DefaultConfig
	}
	return nil
}

type RedisConfig_ClientOutputBufferLimit struct {
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

func (x *RedisConfig_ClientOutputBufferLimit) Reset() {
	*x = RedisConfig_ClientOutputBufferLimit{}
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RedisConfig_ClientOutputBufferLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisConfig_ClientOutputBufferLimit) ProtoMessage() {}

func (x *RedisConfig_ClientOutputBufferLimit) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisConfig_ClientOutputBufferLimit.ProtoReflect.Descriptor instead.
func (*RedisConfig_ClientOutputBufferLimit) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescGZIP(), []int{0, 0}
}

func (x *RedisConfig_ClientOutputBufferLimit) GetHardLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.HardLimit
	}
	return nil
}

func (x *RedisConfig_ClientOutputBufferLimit) GetSoftLimit() *wrapperspb.Int64Value {
	if x != nil {
		return x.SoftLimit
	}
	return nil
}

func (x *RedisConfig_ClientOutputBufferLimit) GetSoftSeconds() *wrapperspb.Int64Value {
	if x != nil {
		return x.SoftSeconds
	}
	return nil
}

var File_yandex_cloud_mdb_redis_v1_config_redis_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDesc = "" +
	"\n" +
	",yandex/cloud/mdb/redis/v1/config/redis.proto\x12 yandex.cloud.mdb.redis.v1.config\x1a\x1egoogle/protobuf/wrappers.proto\x1a\x1dyandex/cloud/validation.proto\"\xee\x13\n" +
	"\vRedisConfig\x12h\n" +
	"\x10maxmemory_policy\x18\x01 \x01(\x0e2=.yandex.cloud.mdb.redis.v1.config.RedisConfig.MaxmemoryPolicyR\x0fmaxmemoryPolicy\x125\n" +
	"\atimeout\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\atimeout\x12B\n" +
	"\bpassword\x18\x03 \x01(\tB&\xf2\xc71\"[a-zA-Z0-9@=+?*.,!&#$^<>_-]{8,128}R\bpassword\x12A\n" +
	"\tdatabases\x18\x04 \x01(\v2\x1b.google.protobuf.Int64ValueB\x06\xfa\xc71\x02>0R\tdatabases\x12[\n" +
	"\x17slowlog_log_slower_than\x18\x05 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\x14slowlogLogSlowerThan\x12L\n" +
	"\x0fslowlog_max_len\x18\x06 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\rslowlogMaxLen\x12O\n" +
	"\x16notify_keyspace_events\x18\a \x01(\tB\x19\xf2\xc71\x15[KEg$lshzxeAtm]{0,13}R\x14notifyKeyspaceEvents\x12\x8f\x01\n" +
	"!client_output_buffer_limit_pubsub\x18\b \x01(\v2E.yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimitR\x1dclientOutputBufferLimitPubsub\x12\x8f\x01\n" +
	"!client_output_buffer_limit_normal\x18\t \x01(\v2E.yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimitR\x1dclientOutputBufferLimitNormal\x12R\n" +
	"\x11maxmemory_percent\x18\n" +
	" \x01(\v2\x1b.google.protobuf.Int64ValueB\b\xfa\xc71\x041-75R\x10maxmemoryPercent\x12J\n" +
	"\x0elua_time_limit\x18\v \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\fluaTimeLimit\x12^\n" +
	"\x19repl_backlog_size_percent\x18\f \x01(\v2\x1b.google.protobuf.Int64ValueB\x06\xfa\xc71\x02>0R\x16replBacklogSizePercent\x12]\n" +
	"\x1dcluster_require_full_coverage\x18\r \x01(\v2\x1a.google.protobuf.BoolValueR\x1aclusterRequireFullCoverage\x12\\\n" +
	"\x1dcluster_allow_reads_when_down\x18\x0e \x01(\v2\x1a.google.protobuf.BoolValueR\x19clusterAllowReadsWhenDown\x12h\n" +
	"#cluster_allow_pubsubshard_when_down\x18\x0f \x01(\v2\x1a.google.protobuf.BoolValueR\x1fclusterAllowPubsubshardWhenDown\x12J\n" +
	"\x0elfu_decay_time\x18\x10 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\flfuDecayTime\x12J\n" +
	"\x0elfu_log_factor\x18\x11 \x01(\v2\x1b.google.protobuf.Int64ValueB\a\xfa\xc71\x03>=0R\flfuLogFactor\x12P\n" +
	"\x16turn_before_switchover\x18\x12 \x01(\v2\x1a.google.protobuf.BoolValueR\x14turnBeforeSwitchover\x12B\n" +
	"\x0fallow_data_loss\x18\x13 \x01(\v2\x1a.google.protobuf.BoolValueR\rallowDataLoss\x129\n" +
	"\n" +
	"use_luajit\x18\x14 \x01(\v2\x1a.google.protobuf.BoolValueR\tuseLuajit\x12H\n" +
	"\x12io_threads_allowed\x18\x15 \x01(\v2\x1a.google.protobuf.BoolValueR\x10ioThreadsAllowed\x12c\n" +
	"\x19zset_max_listpack_entries\x18\x16 \x01(\v2\x1b.google.protobuf.Int64ValueB\v\xfa\xc71\a32-2048R\x16zsetMaxListpackEntries\x12V\n" +
	"\x14aof_max_size_percent\x18\x17 \x01(\v2\x1b.google.protobuf.Int64ValueB\b\xfa\xc71\x041-99R\x11aofMaxSizePercent\x12>\n" +
	"\factivedefrag\x18\x18 \x01(\v2\x1a.google.protobuf.BoolValueR\factivedefrag\x1a\xf8\x01\n" +
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
	"NOEVICTION\x10\b\"\x90\x02\n" +
	"\x0eRedisConfigSet\x12X\n" +
	"\x10effective_config\x18\x01 \x01(\v2-.yandex.cloud.mdb.redis.v1.config.RedisConfigR\x0feffectiveConfig\x12N\n" +
	"\vuser_config\x18\x02 \x01(\v2-.yandex.cloud.mdb.redis.v1.config.RedisConfigR\n" +
	"userConfig\x12T\n" +
	"\x0edefault_config\x18\x03 \x01(\v2-.yandex.cloud.mdb.redis.v1.config.RedisConfigR\rdefaultConfigBr\n" +
	"$yandex.cloud.api.mdb.redis.v1.configZJgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1/config;redisb\x06proto3"

var (
	file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDesc), len(file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDescData
}

var file_yandex_cloud_mdb_redis_v1_config_redis_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_mdb_redis_v1_config_redis_proto_goTypes = []any{
	(RedisConfig_MaxmemoryPolicy)(0),            // 0: yandex.cloud.mdb.redis.v1.config.RedisConfig.MaxmemoryPolicy
	(*RedisConfig)(nil),                         // 1: yandex.cloud.mdb.redis.v1.config.RedisConfig
	(*RedisConfigSet)(nil),                      // 2: yandex.cloud.mdb.redis.v1.config.RedisConfigSet
	(*RedisConfig_ClientOutputBufferLimit)(nil), // 3: yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimit
	(*wrapperspb.Int64Value)(nil),               // 4: google.protobuf.Int64Value
	(*wrapperspb.BoolValue)(nil),                // 5: google.protobuf.BoolValue
}
var file_yandex_cloud_mdb_redis_v1_config_redis_proto_depIdxs = []int32{
	0,  // 0: yandex.cloud.mdb.redis.v1.config.RedisConfig.maxmemory_policy:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig.MaxmemoryPolicy
	4,  // 1: yandex.cloud.mdb.redis.v1.config.RedisConfig.timeout:type_name -> google.protobuf.Int64Value
	4,  // 2: yandex.cloud.mdb.redis.v1.config.RedisConfig.databases:type_name -> google.protobuf.Int64Value
	4,  // 3: yandex.cloud.mdb.redis.v1.config.RedisConfig.slowlog_log_slower_than:type_name -> google.protobuf.Int64Value
	4,  // 4: yandex.cloud.mdb.redis.v1.config.RedisConfig.slowlog_max_len:type_name -> google.protobuf.Int64Value
	3,  // 5: yandex.cloud.mdb.redis.v1.config.RedisConfig.client_output_buffer_limit_pubsub:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimit
	3,  // 6: yandex.cloud.mdb.redis.v1.config.RedisConfig.client_output_buffer_limit_normal:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimit
	4,  // 7: yandex.cloud.mdb.redis.v1.config.RedisConfig.maxmemory_percent:type_name -> google.protobuf.Int64Value
	4,  // 8: yandex.cloud.mdb.redis.v1.config.RedisConfig.lua_time_limit:type_name -> google.protobuf.Int64Value
	4,  // 9: yandex.cloud.mdb.redis.v1.config.RedisConfig.repl_backlog_size_percent:type_name -> google.protobuf.Int64Value
	5,  // 10: yandex.cloud.mdb.redis.v1.config.RedisConfig.cluster_require_full_coverage:type_name -> google.protobuf.BoolValue
	5,  // 11: yandex.cloud.mdb.redis.v1.config.RedisConfig.cluster_allow_reads_when_down:type_name -> google.protobuf.BoolValue
	5,  // 12: yandex.cloud.mdb.redis.v1.config.RedisConfig.cluster_allow_pubsubshard_when_down:type_name -> google.protobuf.BoolValue
	4,  // 13: yandex.cloud.mdb.redis.v1.config.RedisConfig.lfu_decay_time:type_name -> google.protobuf.Int64Value
	4,  // 14: yandex.cloud.mdb.redis.v1.config.RedisConfig.lfu_log_factor:type_name -> google.protobuf.Int64Value
	5,  // 15: yandex.cloud.mdb.redis.v1.config.RedisConfig.turn_before_switchover:type_name -> google.protobuf.BoolValue
	5,  // 16: yandex.cloud.mdb.redis.v1.config.RedisConfig.allow_data_loss:type_name -> google.protobuf.BoolValue
	5,  // 17: yandex.cloud.mdb.redis.v1.config.RedisConfig.use_luajit:type_name -> google.protobuf.BoolValue
	5,  // 18: yandex.cloud.mdb.redis.v1.config.RedisConfig.io_threads_allowed:type_name -> google.protobuf.BoolValue
	4,  // 19: yandex.cloud.mdb.redis.v1.config.RedisConfig.zset_max_listpack_entries:type_name -> google.protobuf.Int64Value
	4,  // 20: yandex.cloud.mdb.redis.v1.config.RedisConfig.aof_max_size_percent:type_name -> google.protobuf.Int64Value
	5,  // 21: yandex.cloud.mdb.redis.v1.config.RedisConfig.activedefrag:type_name -> google.protobuf.BoolValue
	1,  // 22: yandex.cloud.mdb.redis.v1.config.RedisConfigSet.effective_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig
	1,  // 23: yandex.cloud.mdb.redis.v1.config.RedisConfigSet.user_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig
	1,  // 24: yandex.cloud.mdb.redis.v1.config.RedisConfigSet.default_config:type_name -> yandex.cloud.mdb.redis.v1.config.RedisConfig
	4,  // 25: yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimit.hard_limit:type_name -> google.protobuf.Int64Value
	4,  // 26: yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimit.soft_limit:type_name -> google.protobuf.Int64Value
	4,  // 27: yandex.cloud.mdb.redis.v1.config.RedisConfig.ClientOutputBufferLimit.soft_seconds:type_name -> google.protobuf.Int64Value
	28, // [28:28] is the sub-list for method output_type
	28, // [28:28] is the sub-list for method input_type
	28, // [28:28] is the sub-list for extension type_name
	28, // [28:28] is the sub-list for extension extendee
	0,  // [0:28] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_redis_v1_config_redis_proto_init() }
func file_yandex_cloud_mdb_redis_v1_config_redis_proto_init() {
	if File_yandex_cloud_mdb_redis_v1_config_redis_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDesc), len(file_yandex_cloud_mdb_redis_v1_config_redis_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_mdb_redis_v1_config_redis_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_redis_v1_config_redis_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_mdb_redis_v1_config_redis_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_mdb_redis_v1_config_redis_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_redis_v1_config_redis_proto = out.File
	file_yandex_cloud_mdb_redis_v1_config_redis_proto_goTypes = nil
	file_yandex_cloud_mdb_redis_v1_config_redis_proto_depIdxs = nil
}
