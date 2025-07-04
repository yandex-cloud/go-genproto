// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *ClickhouseConfig) SetLogLevel(v ClickhouseConfig_LogLevel) {
	m.LogLevel = v
}

func (m *ClickhouseConfig) SetMergeTree(v *ClickhouseConfig_MergeTree) {
	m.MergeTree = v
}

func (m *ClickhouseConfig) SetCompression(v []*ClickhouseConfig_Compression) {
	m.Compression = v
}

func (m *ClickhouseConfig) SetDictionaries(v []*ClickhouseConfig_ExternalDictionary) {
	m.Dictionaries = v
}

func (m *ClickhouseConfig) SetGraphiteRollup(v []*ClickhouseConfig_GraphiteRollup) {
	m.GraphiteRollup = v
}

func (m *ClickhouseConfig) SetKafka(v *ClickhouseConfig_Kafka) {
	m.Kafka = v
}

func (m *ClickhouseConfig) SetKafkaTopics(v []*ClickhouseConfig_KafkaTopic) {
	m.KafkaTopics = v
}

func (m *ClickhouseConfig) SetRabbitmq(v *ClickhouseConfig_Rabbitmq) {
	m.Rabbitmq = v
}

func (m *ClickhouseConfig) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *ClickhouseConfig) SetMaxConcurrentQueries(v *wrapperspb.Int64Value) {
	m.MaxConcurrentQueries = v
}

func (m *ClickhouseConfig) SetKeepAliveTimeout(v *wrapperspb.Int64Value) {
	m.KeepAliveTimeout = v
}

func (m *ClickhouseConfig) SetUncompressedCacheSize(v *wrapperspb.Int64Value) {
	m.UncompressedCacheSize = v
}

func (m *ClickhouseConfig) SetMarkCacheSize(v *wrapperspb.Int64Value) {
	m.MarkCacheSize = v
}

func (m *ClickhouseConfig) SetMaxTableSizeToDrop(v *wrapperspb.Int64Value) {
	m.MaxTableSizeToDrop = v
}

func (m *ClickhouseConfig) SetMaxPartitionSizeToDrop(v *wrapperspb.Int64Value) {
	m.MaxPartitionSizeToDrop = v
}

func (m *ClickhouseConfig) SetBuiltinDictionariesReloadInterval(v *wrapperspb.Int64Value) {
	m.BuiltinDictionariesReloadInterval = v
}

func (m *ClickhouseConfig) SetTimezone(v string) {
	m.Timezone = v
}

func (m *ClickhouseConfig) SetGeobaseEnabled(v *wrapperspb.BoolValue) {
	m.GeobaseEnabled = v
}

func (m *ClickhouseConfig) SetGeobaseUri(v string) {
	m.GeobaseUri = v
}

func (m *ClickhouseConfig) SetQueryLogRetentionSize(v *wrapperspb.Int64Value) {
	m.QueryLogRetentionSize = v
}

func (m *ClickhouseConfig) SetQueryLogRetentionTime(v *wrapperspb.Int64Value) {
	m.QueryLogRetentionTime = v
}

func (m *ClickhouseConfig) SetQueryThreadLogEnabled(v *wrapperspb.BoolValue) {
	m.QueryThreadLogEnabled = v
}

func (m *ClickhouseConfig) SetQueryThreadLogRetentionSize(v *wrapperspb.Int64Value) {
	m.QueryThreadLogRetentionSize = v
}

func (m *ClickhouseConfig) SetQueryThreadLogRetentionTime(v *wrapperspb.Int64Value) {
	m.QueryThreadLogRetentionTime = v
}

func (m *ClickhouseConfig) SetPartLogRetentionSize(v *wrapperspb.Int64Value) {
	m.PartLogRetentionSize = v
}

func (m *ClickhouseConfig) SetPartLogRetentionTime(v *wrapperspb.Int64Value) {
	m.PartLogRetentionTime = v
}

func (m *ClickhouseConfig) SetMetricLogEnabled(v *wrapperspb.BoolValue) {
	m.MetricLogEnabled = v
}

func (m *ClickhouseConfig) SetMetricLogRetentionSize(v *wrapperspb.Int64Value) {
	m.MetricLogRetentionSize = v
}

func (m *ClickhouseConfig) SetMetricLogRetentionTime(v *wrapperspb.Int64Value) {
	m.MetricLogRetentionTime = v
}

func (m *ClickhouseConfig) SetTraceLogEnabled(v *wrapperspb.BoolValue) {
	m.TraceLogEnabled = v
}

func (m *ClickhouseConfig) SetTraceLogRetentionSize(v *wrapperspb.Int64Value) {
	m.TraceLogRetentionSize = v
}

func (m *ClickhouseConfig) SetTraceLogRetentionTime(v *wrapperspb.Int64Value) {
	m.TraceLogRetentionTime = v
}

func (m *ClickhouseConfig) SetTextLogEnabled(v *wrapperspb.BoolValue) {
	m.TextLogEnabled = v
}

func (m *ClickhouseConfig) SetTextLogRetentionSize(v *wrapperspb.Int64Value) {
	m.TextLogRetentionSize = v
}

func (m *ClickhouseConfig) SetTextLogRetentionTime(v *wrapperspb.Int64Value) {
	m.TextLogRetentionTime = v
}

func (m *ClickhouseConfig) SetTextLogLevel(v ClickhouseConfig_LogLevel) {
	m.TextLogLevel = v
}

func (m *ClickhouseConfig) SetOpentelemetrySpanLogEnabled(v *wrapperspb.BoolValue) {
	m.OpentelemetrySpanLogEnabled = v
}

func (m *ClickhouseConfig) SetOpentelemetrySpanLogRetentionSize(v *wrapperspb.Int64Value) {
	m.OpentelemetrySpanLogRetentionSize = v
}

func (m *ClickhouseConfig) SetOpentelemetrySpanLogRetentionTime(v *wrapperspb.Int64Value) {
	m.OpentelemetrySpanLogRetentionTime = v
}

func (m *ClickhouseConfig) SetQueryViewsLogEnabled(v *wrapperspb.BoolValue) {
	m.QueryViewsLogEnabled = v
}

func (m *ClickhouseConfig) SetQueryViewsLogRetentionSize(v *wrapperspb.Int64Value) {
	m.QueryViewsLogRetentionSize = v
}

func (m *ClickhouseConfig) SetQueryViewsLogRetentionTime(v *wrapperspb.Int64Value) {
	m.QueryViewsLogRetentionTime = v
}

func (m *ClickhouseConfig) SetAsynchronousMetricLogEnabled(v *wrapperspb.BoolValue) {
	m.AsynchronousMetricLogEnabled = v
}

func (m *ClickhouseConfig) SetAsynchronousMetricLogRetentionSize(v *wrapperspb.Int64Value) {
	m.AsynchronousMetricLogRetentionSize = v
}

func (m *ClickhouseConfig) SetAsynchronousMetricLogRetentionTime(v *wrapperspb.Int64Value) {
	m.AsynchronousMetricLogRetentionTime = v
}

func (m *ClickhouseConfig) SetSessionLogEnabled(v *wrapperspb.BoolValue) {
	m.SessionLogEnabled = v
}

func (m *ClickhouseConfig) SetSessionLogRetentionSize(v *wrapperspb.Int64Value) {
	m.SessionLogRetentionSize = v
}

func (m *ClickhouseConfig) SetSessionLogRetentionTime(v *wrapperspb.Int64Value) {
	m.SessionLogRetentionTime = v
}

func (m *ClickhouseConfig) SetZookeeperLogEnabled(v *wrapperspb.BoolValue) {
	m.ZookeeperLogEnabled = v
}

func (m *ClickhouseConfig) SetZookeeperLogRetentionSize(v *wrapperspb.Int64Value) {
	m.ZookeeperLogRetentionSize = v
}

func (m *ClickhouseConfig) SetZookeeperLogRetentionTime(v *wrapperspb.Int64Value) {
	m.ZookeeperLogRetentionTime = v
}

func (m *ClickhouseConfig) SetAsynchronousInsertLogEnabled(v *wrapperspb.BoolValue) {
	m.AsynchronousInsertLogEnabled = v
}

func (m *ClickhouseConfig) SetAsynchronousInsertLogRetentionSize(v *wrapperspb.Int64Value) {
	m.AsynchronousInsertLogRetentionSize = v
}

func (m *ClickhouseConfig) SetAsynchronousInsertLogRetentionTime(v *wrapperspb.Int64Value) {
	m.AsynchronousInsertLogRetentionTime = v
}

func (m *ClickhouseConfig) SetProcessorsProfileLogEnabled(v *wrapperspb.BoolValue) {
	m.ProcessorsProfileLogEnabled = v
}

func (m *ClickhouseConfig) SetProcessorsProfileLogRetentionSize(v *wrapperspb.Int64Value) {
	m.ProcessorsProfileLogRetentionSize = v
}

func (m *ClickhouseConfig) SetProcessorsProfileLogRetentionTime(v *wrapperspb.Int64Value) {
	m.ProcessorsProfileLogRetentionTime = v
}

func (m *ClickhouseConfig) SetErrorLogEnabled(v *wrapperspb.BoolValue) {
	m.ErrorLogEnabled = v
}

func (m *ClickhouseConfig) SetErrorLogRetentionSize(v *wrapperspb.Int64Value) {
	m.ErrorLogRetentionSize = v
}

func (m *ClickhouseConfig) SetErrorLogRetentionTime(v *wrapperspb.Int64Value) {
	m.ErrorLogRetentionTime = v
}

func (m *ClickhouseConfig) SetAccessControlImprovements(v *ClickhouseConfig_AccessControlImprovements) {
	m.AccessControlImprovements = v
}

func (m *ClickhouseConfig) SetBackgroundPoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundPoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundMergesMutationsConcurrencyRatio(v *wrapperspb.Int64Value) {
	m.BackgroundMergesMutationsConcurrencyRatio = v
}

func (m *ClickhouseConfig) SetBackgroundSchedulePoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundSchedulePoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundFetchesPoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundFetchesPoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundMovePoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundMovePoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundDistributedSchedulePoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundDistributedSchedulePoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundBufferFlushSchedulePoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundBufferFlushSchedulePoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundMessageBrokerSchedulePoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundMessageBrokerSchedulePoolSize = v
}

func (m *ClickhouseConfig) SetBackgroundCommonPoolSize(v *wrapperspb.Int64Value) {
	m.BackgroundCommonPoolSize = v
}

func (m *ClickhouseConfig) SetDefaultDatabase(v *wrapperspb.StringValue) {
	m.DefaultDatabase = v
}

func (m *ClickhouseConfig) SetTotalMemoryProfilerStep(v *wrapperspb.Int64Value) {
	m.TotalMemoryProfilerStep = v
}

func (m *ClickhouseConfig) SetTotalMemoryTrackerSampleProbability(v *wrapperspb.DoubleValue) {
	m.TotalMemoryTrackerSampleProbability = v
}

func (m *ClickhouseConfig) SetQueryMaskingRules(v []*ClickhouseConfig_QueryMaskingRule) {
	m.QueryMaskingRules = v
}

func (m *ClickhouseConfig) SetDictionariesLazyLoad(v *wrapperspb.BoolValue) {
	m.DictionariesLazyLoad = v
}

func (m *ClickhouseConfig) SetQueryCache(v *ClickhouseConfig_QueryCache) {
	m.QueryCache = v
}

func (m *ClickhouseConfig) SetJdbcBridge(v *ClickhouseConfig_JdbcBridge) {
	m.JdbcBridge = v
}

func (m *ClickhouseConfig_MergeTree) SetReplicatedDeduplicationWindow(v *wrapperspb.Int64Value) {
	m.ReplicatedDeduplicationWindow = v
}

func (m *ClickhouseConfig_MergeTree) SetReplicatedDeduplicationWindowSeconds(v *wrapperspb.Int64Value) {
	m.ReplicatedDeduplicationWindowSeconds = v
}

func (m *ClickhouseConfig_MergeTree) SetPartsToDelayInsert(v *wrapperspb.Int64Value) {
	m.PartsToDelayInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetPartsToThrowInsert(v *wrapperspb.Int64Value) {
	m.PartsToThrowInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetInactivePartsToDelayInsert(v *wrapperspb.Int64Value) {
	m.InactivePartsToDelayInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetInactivePartsToThrowInsert(v *wrapperspb.Int64Value) {
	m.InactivePartsToThrowInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxReplicatedMergesInQueue(v *wrapperspb.Int64Value) {
	m.MaxReplicatedMergesInQueue = v
}

func (m *ClickhouseConfig_MergeTree) SetNumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge(v *wrapperspb.Int64Value) {
	m.NumberOfFreeEntriesInPoolToLowerMaxSizeOfMerge = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxBytesToMergeAtMinSpaceInPool(v *wrapperspb.Int64Value) {
	m.MaxBytesToMergeAtMinSpaceInPool = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxBytesToMergeAtMaxSpaceInPool(v *wrapperspb.Int64Value) {
	m.MaxBytesToMergeAtMaxSpaceInPool = v
}

func (m *ClickhouseConfig_MergeTree) SetMinBytesForWidePart(v *wrapperspb.Int64Value) {
	m.MinBytesForWidePart = v
}

func (m *ClickhouseConfig_MergeTree) SetMinRowsForWidePart(v *wrapperspb.Int64Value) {
	m.MinRowsForWidePart = v
}

func (m *ClickhouseConfig_MergeTree) SetTtlOnlyDropParts(v *wrapperspb.BoolValue) {
	m.TtlOnlyDropParts = v
}

func (m *ClickhouseConfig_MergeTree) SetAllowRemoteFsZeroCopyReplication(v *wrapperspb.BoolValue) {
	m.AllowRemoteFsZeroCopyReplication = v
}

func (m *ClickhouseConfig_MergeTree) SetMergeWithTtlTimeout(v *wrapperspb.Int64Value) {
	m.MergeWithTtlTimeout = v
}

func (m *ClickhouseConfig_MergeTree) SetMergeWithRecompressionTtlTimeout(v *wrapperspb.Int64Value) {
	m.MergeWithRecompressionTtlTimeout = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxPartsInTotal(v *wrapperspb.Int64Value) {
	m.MaxPartsInTotal = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxNumberOfMergesWithTtlInPool(v *wrapperspb.Int64Value) {
	m.MaxNumberOfMergesWithTtlInPool = v
}

func (m *ClickhouseConfig_MergeTree) SetCleanupDelayPeriod(v *wrapperspb.Int64Value) {
	m.CleanupDelayPeriod = v
}

func (m *ClickhouseConfig_MergeTree) SetNumberOfFreeEntriesInPoolToExecuteMutation(v *wrapperspb.Int64Value) {
	m.NumberOfFreeEntriesInPoolToExecuteMutation = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxAvgPartSizeForTooManyParts(v *wrapperspb.Int64Value) {
	m.MaxAvgPartSizeForTooManyParts = v
}

func (m *ClickhouseConfig_MergeTree) SetMinAgeToForceMergeSeconds(v *wrapperspb.Int64Value) {
	m.MinAgeToForceMergeSeconds = v
}

func (m *ClickhouseConfig_MergeTree) SetMinAgeToForceMergeOnPartitionOnly(v *wrapperspb.BoolValue) {
	m.MinAgeToForceMergeOnPartitionOnly = v
}

func (m *ClickhouseConfig_MergeTree) SetMergeSelectingSleepMs(v *wrapperspb.Int64Value) {
	m.MergeSelectingSleepMs = v
}

func (m *ClickhouseConfig_MergeTree) SetMergeMaxBlockSize(v *wrapperspb.Int64Value) {
	m.MergeMaxBlockSize = v
}

func (m *ClickhouseConfig_MergeTree) SetCheckSampleColumnIsCorrect(v *wrapperspb.BoolValue) {
	m.CheckSampleColumnIsCorrect = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxMergeSelectingSleepMs(v *wrapperspb.Int64Value) {
	m.MaxMergeSelectingSleepMs = v
}

func (m *ClickhouseConfig_MergeTree) SetMaxCleanupDelayPeriod(v *wrapperspb.Int64Value) {
	m.MaxCleanupDelayPeriod = v
}

func (m *ClickhouseConfig_MergeTree) SetDeduplicateMergeProjectionMode(v ClickhouseConfig_MergeTree_DeduplicateMergeProjectionMode) {
	m.DeduplicateMergeProjectionMode = v
}

func (m *ClickhouseConfig_MergeTree) SetLightweightMutationProjectionMode(v ClickhouseConfig_MergeTree_LightweightMutationProjectionMode) {
	m.LightweightMutationProjectionMode = v
}

func (m *ClickhouseConfig_MergeTree) SetMaterializeTtlRecalculateOnly(v *wrapperspb.BoolValue) {
	m.MaterializeTtlRecalculateOnly = v
}

func (m *ClickhouseConfig_MergeTree) SetFsyncAfterInsert(v *wrapperspb.BoolValue) {
	m.FsyncAfterInsert = v
}

func (m *ClickhouseConfig_MergeTree) SetFsyncPartDirectory(v *wrapperspb.BoolValue) {
	m.FsyncPartDirectory = v
}

func (m *ClickhouseConfig_MergeTree) SetMinCompressedBytesToFsyncAfterFetch(v *wrapperspb.Int64Value) {
	m.MinCompressedBytesToFsyncAfterFetch = v
}

func (m *ClickhouseConfig_MergeTree) SetMinCompressedBytesToFsyncAfterMerge(v *wrapperspb.Int64Value) {
	m.MinCompressedBytesToFsyncAfterMerge = v
}

func (m *ClickhouseConfig_MergeTree) SetMinRowsToFsyncAfterMerge(v *wrapperspb.Int64Value) {
	m.MinRowsToFsyncAfterMerge = v
}

func (m *ClickhouseConfig_Kafka) SetSecurityProtocol(v ClickhouseConfig_Kafka_SecurityProtocol) {
	m.SecurityProtocol = v
}

func (m *ClickhouseConfig_Kafka) SetSaslMechanism(v ClickhouseConfig_Kafka_SaslMechanism) {
	m.SaslMechanism = v
}

func (m *ClickhouseConfig_Kafka) SetSaslUsername(v string) {
	m.SaslUsername = v
}

func (m *ClickhouseConfig_Kafka) SetSaslPassword(v string) {
	m.SaslPassword = v
}

func (m *ClickhouseConfig_Kafka) SetEnableSslCertificateVerification(v *wrapperspb.BoolValue) {
	m.EnableSslCertificateVerification = v
}

func (m *ClickhouseConfig_Kafka) SetMaxPollIntervalMs(v *wrapperspb.Int64Value) {
	m.MaxPollIntervalMs = v
}

func (m *ClickhouseConfig_Kafka) SetSessionTimeoutMs(v *wrapperspb.Int64Value) {
	m.SessionTimeoutMs = v
}

func (m *ClickhouseConfig_Kafka) SetDebug(v ClickhouseConfig_Kafka_Debug) {
	m.Debug = v
}

func (m *ClickhouseConfig_Kafka) SetAutoOffsetReset(v ClickhouseConfig_Kafka_AutoOffsetReset) {
	m.AutoOffsetReset = v
}

func (m *ClickhouseConfig_KafkaTopic) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_KafkaTopic) SetSettings(v *ClickhouseConfig_Kafka) {
	m.Settings = v
}

func (m *ClickhouseConfig_Rabbitmq) SetUsername(v string) {
	m.Username = v
}

func (m *ClickhouseConfig_Rabbitmq) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_Rabbitmq) SetVhost(v string) {
	m.Vhost = v
}

func (m *ClickhouseConfig_Compression) SetMethod(v ClickhouseConfig_Compression_Method) {
	m.Method = v
}

func (m *ClickhouseConfig_Compression) SetMinPartSize(v int64) {
	m.MinPartSize = v
}

func (m *ClickhouseConfig_Compression) SetMinPartSizeRatio(v float64) {
	m.MinPartSizeRatio = v
}

func (m *ClickhouseConfig_Compression) SetLevel(v *wrapperspb.Int64Value) {
	m.Level = v
}

type ClickhouseConfig_ExternalDictionary_Lifetime = isClickhouseConfig_ExternalDictionary_Lifetime

func (m *ClickhouseConfig_ExternalDictionary) SetLifetime(v ClickhouseConfig_ExternalDictionary_Lifetime) {
	m.Lifetime = v
}

type ClickhouseConfig_ExternalDictionary_Source = isClickhouseConfig_ExternalDictionary_Source

func (m *ClickhouseConfig_ExternalDictionary) SetSource(v ClickhouseConfig_ExternalDictionary_Source) {
	m.Source = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetStructure(v *ClickhouseConfig_ExternalDictionary_Structure) {
	m.Structure = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetLayout(v *ClickhouseConfig_ExternalDictionary_Layout) {
	m.Layout = v
}

func (m *ClickhouseConfig_ExternalDictionary) SetFixedLifetime(v int64) {
	m.Lifetime = &ClickhouseConfig_ExternalDictionary_FixedLifetime{
		FixedLifetime: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetLifetimeRange(v *ClickhouseConfig_ExternalDictionary_Range) {
	m.Lifetime = &ClickhouseConfig_ExternalDictionary_LifetimeRange{
		LifetimeRange: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetHttpSource(v *ClickhouseConfig_ExternalDictionary_HttpSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_HttpSource_{
		HttpSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetMysqlSource(v *ClickhouseConfig_ExternalDictionary_MysqlSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_MysqlSource_{
		MysqlSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetClickhouseSource(v *ClickhouseConfig_ExternalDictionary_ClickhouseSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_ClickhouseSource_{
		ClickhouseSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetMongodbSource(v *ClickhouseConfig_ExternalDictionary_MongodbSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_MongodbSource_{
		MongodbSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary) SetPostgresqlSource(v *ClickhouseConfig_ExternalDictionary_PostgresqlSource) {
	m.Source = &ClickhouseConfig_ExternalDictionary_PostgresqlSource_{
		PostgresqlSource: v,
	}
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource) SetUrl(v string) {
	m.Url = v
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource) SetFormat(v string) {
	m.Format = v
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource) SetHeaders(v []*ClickhouseConfig_ExternalDictionary_HttpSource_Header) {
	m.Headers = v
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource_Header) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary_HttpSource_Header) SetValue(v string) {
	m.Value = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetTable(v string) {
	m.Table = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetReplicas(v []*ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) {
	m.Replicas = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetWhere(v string) {
	m.Where = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetInvalidateQuery(v string) {
	m.InvalidateQuery = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetCloseConnection(v *wrapperspb.BoolValue) {
	m.CloseConnection = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource) SetShareConnection(v *wrapperspb.BoolValue) {
	m.ShareConnection = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetPriority(v int64) {
	m.Priority = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_MysqlSource_Replica) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetTable(v string) {
	m.Table = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetWhere(v string) {
	m.Where = v
}

func (m *ClickhouseConfig_ExternalDictionary_ClickhouseSource) SetSecure(v *wrapperspb.BoolValue) {
	m.Secure = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetCollection(v string) {
	m.Collection = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_MongodbSource) SetOptions(v string) {
	m.Options = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetDb(v string) {
	m.Db = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetTable(v string) {
	m.Table = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetHosts(v []string) {
	m.Hosts = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetPort(v int64) {
	m.Port = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetUser(v string) {
	m.User = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetPassword(v string) {
	m.Password = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetInvalidateQuery(v string) {
	m.InvalidateQuery = v
}

func (m *ClickhouseConfig_ExternalDictionary_PostgresqlSource) SetSslMode(v ClickhouseConfig_ExternalDictionary_PostgresqlSource_SslMode) {
	m.SslMode = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetId(v *ClickhouseConfig_ExternalDictionary_Structure_Id) {
	m.Id = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetKey(v *ClickhouseConfig_ExternalDictionary_Structure_Key) {
	m.Key = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetRangeMin(v *ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.RangeMin = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetRangeMax(v *ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.RangeMax = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure) SetAttributes(v []*ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.Attributes = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetType(v string) {
	m.Type = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetNullValue(v string) {
	m.NullValue = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetExpression(v string) {
	m.Expression = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetHierarchical(v bool) {
	m.Hierarchical = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Attribute) SetInjective(v bool) {
	m.Injective = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Id) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_ExternalDictionary_Structure_Key) SetAttributes(v []*ClickhouseConfig_ExternalDictionary_Structure_Attribute) {
	m.Attributes = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetType(v ClickhouseConfig_ExternalDictionary_Layout_Type) {
	m.Type = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetSizeInCells(v int64) {
	m.SizeInCells = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetAllowReadExpiredKeys(v *wrapperspb.BoolValue) {
	m.AllowReadExpiredKeys = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetMaxUpdateQueueSize(v int64) {
	m.MaxUpdateQueueSize = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetUpdateQueuePushTimeoutMilliseconds(v int64) {
	m.UpdateQueuePushTimeoutMilliseconds = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetQueryWaitTimeoutMilliseconds(v int64) {
	m.QueryWaitTimeoutMilliseconds = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetMaxThreadsForUpdates(v int64) {
	m.MaxThreadsForUpdates = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetInitialArraySize(v int64) {
	m.InitialArraySize = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetMaxArraySize(v int64) {
	m.MaxArraySize = v
}

func (m *ClickhouseConfig_ExternalDictionary_Layout) SetAccessToKeyFromAttributes(v *wrapperspb.BoolValue) {
	m.AccessToKeyFromAttributes = v
}

func (m *ClickhouseConfig_ExternalDictionary_Range) SetMin(v int64) {
	m.Min = v
}

func (m *ClickhouseConfig_ExternalDictionary_Range) SetMax(v int64) {
	m.Max = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetPatterns(v []*ClickhouseConfig_GraphiteRollup_Pattern) {
	m.Patterns = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetPathColumnName(v string) {
	m.PathColumnName = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetTimeColumnName(v string) {
	m.TimeColumnName = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetValueColumnName(v string) {
	m.ValueColumnName = v
}

func (m *ClickhouseConfig_GraphiteRollup) SetVersionColumnName(v string) {
	m.VersionColumnName = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern) SetRegexp(v string) {
	m.Regexp = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern) SetFunction(v string) {
	m.Function = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern) SetRetention(v []*ClickhouseConfig_GraphiteRollup_Pattern_Retention) {
	m.Retention = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern_Retention) SetAge(v int64) {
	m.Age = v
}

func (m *ClickhouseConfig_GraphiteRollup_Pattern_Retention) SetPrecision(v int64) {
	m.Precision = v
}

func (m *ClickhouseConfig_QueryMaskingRule) SetName(v string) {
	m.Name = v
}

func (m *ClickhouseConfig_QueryMaskingRule) SetRegexp(v string) {
	m.Regexp = v
}

func (m *ClickhouseConfig_QueryMaskingRule) SetReplace(v string) {
	m.Replace = v
}

func (m *ClickhouseConfig_QueryCache) SetMaxSizeInBytes(v *wrapperspb.Int64Value) {
	m.MaxSizeInBytes = v
}

func (m *ClickhouseConfig_QueryCache) SetMaxEntries(v *wrapperspb.Int64Value) {
	m.MaxEntries = v
}

func (m *ClickhouseConfig_QueryCache) SetMaxEntrySizeInBytes(v *wrapperspb.Int64Value) {
	m.MaxEntrySizeInBytes = v
}

func (m *ClickhouseConfig_QueryCache) SetMaxEntrySizeInRows(v *wrapperspb.Int64Value) {
	m.MaxEntrySizeInRows = v
}

func (m *ClickhouseConfig_JdbcBridge) SetHost(v string) {
	m.Host = v
}

func (m *ClickhouseConfig_JdbcBridge) SetPort(v *wrapperspb.Int64Value) {
	m.Port = v
}

func (m *ClickhouseConfig_AccessControlImprovements) SetSelectFromSystemDbRequiresGrant(v *wrapperspb.BoolValue) {
	m.SelectFromSystemDbRequiresGrant = v
}

func (m *ClickhouseConfig_AccessControlImprovements) SetSelectFromInformationSchemaRequiresGrant(v *wrapperspb.BoolValue) {
	m.SelectFromInformationSchemaRequiresGrant = v
}

func (m *ClickhouseConfigSet) SetEffectiveConfig(v *ClickhouseConfig) {
	m.EffectiveConfig = v
}

func (m *ClickhouseConfigSet) SetUserConfig(v *ClickhouseConfig) {
	m.UserConfig = v
}

func (m *ClickhouseConfigSet) SetDefaultConfig(v *ClickhouseConfig) {
	m.DefaultConfig = v
}
