// Code generated by protoc-gen-goext. DO NOT EDIT.

package clickhouse

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *User) SetName(v string) {
	m.Name = v
}

func (m *User) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *User) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *User) SetSettings(v *UserSettings) {
	m.Settings = v
}

func (m *User) SetQuotas(v []*UserQuota) {
	m.Quotas = v
}

func (m *User) SetConnectionManager(v *ConnectionManager) {
	m.ConnectionManager = v
}

func (m *Permission) SetDatabaseName(v string) {
	m.DatabaseName = v
}

func (m *ConnectionManager) SetConnectionId(v string) {
	m.ConnectionId = v
}

func (m *UserSpec) SetName(v string) {
	m.Name = v
}

func (m *UserSpec) SetPassword(v string) {
	m.Password = v
}

func (m *UserSpec) SetPermissions(v []*Permission) {
	m.Permissions = v
}

func (m *UserSpec) SetSettings(v *UserSettings) {
	m.Settings = v
}

func (m *UserSpec) SetQuotas(v []*UserQuota) {
	m.Quotas = v
}

func (m *UserSpec) SetGeneratePassword(v *wrapperspb.BoolValue) {
	m.GeneratePassword = v
}

func (m *UserSettings) SetReadonly(v *wrapperspb.Int64Value) {
	m.Readonly = v
}

func (m *UserSettings) SetAllowDdl(v *wrapperspb.BoolValue) {
	m.AllowDdl = v
}

func (m *UserSettings) SetAllowIntrospectionFunctions(v *wrapperspb.BoolValue) {
	m.AllowIntrospectionFunctions = v
}

func (m *UserSettings) SetConnectTimeout(v *wrapperspb.Int64Value) {
	m.ConnectTimeout = v
}

func (m *UserSettings) SetConnectTimeoutWithFailover(v *wrapperspb.Int64Value) {
	m.ConnectTimeoutWithFailover = v
}

func (m *UserSettings) SetReceiveTimeout(v *wrapperspb.Int64Value) {
	m.ReceiveTimeout = v
}

func (m *UserSettings) SetSendTimeout(v *wrapperspb.Int64Value) {
	m.SendTimeout = v
}

func (m *UserSettings) SetTimeoutBeforeCheckingExecutionSpeed(v *wrapperspb.Int64Value) {
	m.TimeoutBeforeCheckingExecutionSpeed = v
}

func (m *UserSettings) SetInsertQuorum(v *wrapperspb.Int64Value) {
	m.InsertQuorum = v
}

func (m *UserSettings) SetInsertQuorumTimeout(v *wrapperspb.Int64Value) {
	m.InsertQuorumTimeout = v
}

func (m *UserSettings) SetInsertQuorumParallel(v *wrapperspb.BoolValue) {
	m.InsertQuorumParallel = v
}

func (m *UserSettings) SetInsertNullAsDefault(v *wrapperspb.BoolValue) {
	m.InsertNullAsDefault = v
}

func (m *UserSettings) SetSelectSequentialConsistency(v *wrapperspb.BoolValue) {
	m.SelectSequentialConsistency = v
}

func (m *UserSettings) SetDeduplicateBlocksInDependentMaterializedViews(v *wrapperspb.BoolValue) {
	m.DeduplicateBlocksInDependentMaterializedViews = v
}

func (m *UserSettings) SetReplicationAlterPartitionsSync(v *wrapperspb.Int64Value) {
	m.ReplicationAlterPartitionsSync = v
}

func (m *UserSettings) SetMaxReplicaDelayForDistributedQueries(v *wrapperspb.Int64Value) {
	m.MaxReplicaDelayForDistributedQueries = v
}

func (m *UserSettings) SetFallbackToStaleReplicasForDistributedQueries(v *wrapperspb.BoolValue) {
	m.FallbackToStaleReplicasForDistributedQueries = v
}

func (m *UserSettings) SetDistributedProductMode(v UserSettings_DistributedProductMode) {
	m.DistributedProductMode = v
}

func (m *UserSettings) SetDistributedAggregationMemoryEfficient(v *wrapperspb.BoolValue) {
	m.DistributedAggregationMemoryEfficient = v
}

func (m *UserSettings) SetDistributedDdlTaskTimeout(v *wrapperspb.Int64Value) {
	m.DistributedDdlTaskTimeout = v
}

func (m *UserSettings) SetSkipUnavailableShards(v *wrapperspb.BoolValue) {
	m.SkipUnavailableShards = v
}

func (m *UserSettings) SetCompileExpressions(v *wrapperspb.BoolValue) {
	m.CompileExpressions = v
}

func (m *UserSettings) SetMinCountToCompileExpression(v *wrapperspb.Int64Value) {
	m.MinCountToCompileExpression = v
}

func (m *UserSettings) SetMaxBlockSize(v *wrapperspb.Int64Value) {
	m.MaxBlockSize = v
}

func (m *UserSettings) SetMinInsertBlockSizeRows(v *wrapperspb.Int64Value) {
	m.MinInsertBlockSizeRows = v
}

func (m *UserSettings) SetMinInsertBlockSizeBytes(v *wrapperspb.Int64Value) {
	m.MinInsertBlockSizeBytes = v
}

func (m *UserSettings) SetMaxInsertBlockSize(v *wrapperspb.Int64Value) {
	m.MaxInsertBlockSize = v
}

func (m *UserSettings) SetMinBytesToUseDirectIo(v *wrapperspb.Int64Value) {
	m.MinBytesToUseDirectIo = v
}

func (m *UserSettings) SetUseUncompressedCache(v *wrapperspb.BoolValue) {
	m.UseUncompressedCache = v
}

func (m *UserSettings) SetMergeTreeMaxRowsToUseCache(v *wrapperspb.Int64Value) {
	m.MergeTreeMaxRowsToUseCache = v
}

func (m *UserSettings) SetMergeTreeMaxBytesToUseCache(v *wrapperspb.Int64Value) {
	m.MergeTreeMaxBytesToUseCache = v
}

func (m *UserSettings) SetMergeTreeMinRowsForConcurrentRead(v *wrapperspb.Int64Value) {
	m.MergeTreeMinRowsForConcurrentRead = v
}

func (m *UserSettings) SetMergeTreeMinBytesForConcurrentRead(v *wrapperspb.Int64Value) {
	m.MergeTreeMinBytesForConcurrentRead = v
}

func (m *UserSettings) SetMaxBytesBeforeExternalGroupBy(v *wrapperspb.Int64Value) {
	m.MaxBytesBeforeExternalGroupBy = v
}

func (m *UserSettings) SetMaxBytesBeforeExternalSort(v *wrapperspb.Int64Value) {
	m.MaxBytesBeforeExternalSort = v
}

func (m *UserSettings) SetGroupByTwoLevelThreshold(v *wrapperspb.Int64Value) {
	m.GroupByTwoLevelThreshold = v
}

func (m *UserSettings) SetGroupByTwoLevelThresholdBytes(v *wrapperspb.Int64Value) {
	m.GroupByTwoLevelThresholdBytes = v
}

func (m *UserSettings) SetPriority(v *wrapperspb.Int64Value) {
	m.Priority = v
}

func (m *UserSettings) SetMaxThreads(v *wrapperspb.Int64Value) {
	m.MaxThreads = v
}

func (m *UserSettings) SetMaxMemoryUsage(v *wrapperspb.Int64Value) {
	m.MaxMemoryUsage = v
}

func (m *UserSettings) SetMaxMemoryUsageForUser(v *wrapperspb.Int64Value) {
	m.MaxMemoryUsageForUser = v
}

func (m *UserSettings) SetMaxNetworkBandwidth(v *wrapperspb.Int64Value) {
	m.MaxNetworkBandwidth = v
}

func (m *UserSettings) SetMaxNetworkBandwidthForUser(v *wrapperspb.Int64Value) {
	m.MaxNetworkBandwidthForUser = v
}

func (m *UserSettings) SetMaxPartitionsPerInsertBlock(v *wrapperspb.Int64Value) {
	m.MaxPartitionsPerInsertBlock = v
}

func (m *UserSettings) SetMaxConcurrentQueriesForUser(v *wrapperspb.Int64Value) {
	m.MaxConcurrentQueriesForUser = v
}

func (m *UserSettings) SetForceIndexByDate(v *wrapperspb.BoolValue) {
	m.ForceIndexByDate = v
}

func (m *UserSettings) SetForcePrimaryKey(v *wrapperspb.BoolValue) {
	m.ForcePrimaryKey = v
}

func (m *UserSettings) SetMaxRowsToRead(v *wrapperspb.Int64Value) {
	m.MaxRowsToRead = v
}

func (m *UserSettings) SetMaxBytesToRead(v *wrapperspb.Int64Value) {
	m.MaxBytesToRead = v
}

func (m *UserSettings) SetReadOverflowMode(v UserSettings_OverflowMode) {
	m.ReadOverflowMode = v
}

func (m *UserSettings) SetMaxRowsToGroupBy(v *wrapperspb.Int64Value) {
	m.MaxRowsToGroupBy = v
}

func (m *UserSettings) SetGroupByOverflowMode(v UserSettings_GroupByOverflowMode) {
	m.GroupByOverflowMode = v
}

func (m *UserSettings) SetMaxRowsToSort(v *wrapperspb.Int64Value) {
	m.MaxRowsToSort = v
}

func (m *UserSettings) SetMaxBytesToSort(v *wrapperspb.Int64Value) {
	m.MaxBytesToSort = v
}

func (m *UserSettings) SetSortOverflowMode(v UserSettings_OverflowMode) {
	m.SortOverflowMode = v
}

func (m *UserSettings) SetMaxResultRows(v *wrapperspb.Int64Value) {
	m.MaxResultRows = v
}

func (m *UserSettings) SetMaxResultBytes(v *wrapperspb.Int64Value) {
	m.MaxResultBytes = v
}

func (m *UserSettings) SetResultOverflowMode(v UserSettings_OverflowMode) {
	m.ResultOverflowMode = v
}

func (m *UserSettings) SetMaxRowsInDistinct(v *wrapperspb.Int64Value) {
	m.MaxRowsInDistinct = v
}

func (m *UserSettings) SetMaxBytesInDistinct(v *wrapperspb.Int64Value) {
	m.MaxBytesInDistinct = v
}

func (m *UserSettings) SetDistinctOverflowMode(v UserSettings_OverflowMode) {
	m.DistinctOverflowMode = v
}

func (m *UserSettings) SetMaxRowsToTransfer(v *wrapperspb.Int64Value) {
	m.MaxRowsToTransfer = v
}

func (m *UserSettings) SetMaxBytesToTransfer(v *wrapperspb.Int64Value) {
	m.MaxBytesToTransfer = v
}

func (m *UserSettings) SetTransferOverflowMode(v UserSettings_OverflowMode) {
	m.TransferOverflowMode = v
}

func (m *UserSettings) SetMaxExecutionTime(v *wrapperspb.Int64Value) {
	m.MaxExecutionTime = v
}

func (m *UserSettings) SetTimeoutOverflowMode(v UserSettings_OverflowMode) {
	m.TimeoutOverflowMode = v
}

func (m *UserSettings) SetMaxRowsInSet(v *wrapperspb.Int64Value) {
	m.MaxRowsInSet = v
}

func (m *UserSettings) SetMaxBytesInSet(v *wrapperspb.Int64Value) {
	m.MaxBytesInSet = v
}

func (m *UserSettings) SetSetOverflowMode(v UserSettings_OverflowMode) {
	m.SetOverflowMode = v
}

func (m *UserSettings) SetMaxRowsInJoin(v *wrapperspb.Int64Value) {
	m.MaxRowsInJoin = v
}

func (m *UserSettings) SetMaxBytesInJoin(v *wrapperspb.Int64Value) {
	m.MaxBytesInJoin = v
}

func (m *UserSettings) SetJoinOverflowMode(v UserSettings_OverflowMode) {
	m.JoinOverflowMode = v
}

func (m *UserSettings) SetJoinAlgorithm(v []UserSettings_JoinAlgorithm) {
	m.JoinAlgorithm = v
}

func (m *UserSettings) SetAnyJoinDistinctRightTableKeys(v *wrapperspb.BoolValue) {
	m.AnyJoinDistinctRightTableKeys = v
}

func (m *UserSettings) SetMaxColumnsToRead(v *wrapperspb.Int64Value) {
	m.MaxColumnsToRead = v
}

func (m *UserSettings) SetMaxTemporaryColumns(v *wrapperspb.Int64Value) {
	m.MaxTemporaryColumns = v
}

func (m *UserSettings) SetMaxTemporaryNonConstColumns(v *wrapperspb.Int64Value) {
	m.MaxTemporaryNonConstColumns = v
}

func (m *UserSettings) SetMaxQuerySize(v *wrapperspb.Int64Value) {
	m.MaxQuerySize = v
}

func (m *UserSettings) SetMaxAstDepth(v *wrapperspb.Int64Value) {
	m.MaxAstDepth = v
}

func (m *UserSettings) SetMaxAstElements(v *wrapperspb.Int64Value) {
	m.MaxAstElements = v
}

func (m *UserSettings) SetMaxExpandedAstElements(v *wrapperspb.Int64Value) {
	m.MaxExpandedAstElements = v
}

func (m *UserSettings) SetMinExecutionSpeed(v *wrapperspb.Int64Value) {
	m.MinExecutionSpeed = v
}

func (m *UserSettings) SetMinExecutionSpeedBytes(v *wrapperspb.Int64Value) {
	m.MinExecutionSpeedBytes = v
}

func (m *UserSettings) SetCountDistinctImplementation(v UserSettings_CountDistinctImplementation) {
	m.CountDistinctImplementation = v
}

func (m *UserSettings) SetInputFormatValuesInterpretExpressions(v *wrapperspb.BoolValue) {
	m.InputFormatValuesInterpretExpressions = v
}

func (m *UserSettings) SetInputFormatDefaultsForOmittedFields(v *wrapperspb.BoolValue) {
	m.InputFormatDefaultsForOmittedFields = v
}

func (m *UserSettings) SetInputFormatNullAsDefault(v *wrapperspb.BoolValue) {
	m.InputFormatNullAsDefault = v
}

func (m *UserSettings) SetDateTimeInputFormat(v UserSettings_DateTimeInputFormat) {
	m.DateTimeInputFormat = v
}

func (m *UserSettings) SetInputFormatWithNamesUseHeader(v *wrapperspb.BoolValue) {
	m.InputFormatWithNamesUseHeader = v
}

func (m *UserSettings) SetOutputFormatJsonQuote_64BitIntegers(v *wrapperspb.BoolValue) {
	m.OutputFormatJsonQuote_64BitIntegers = v
}

func (m *UserSettings) SetOutputFormatJsonQuoteDenormals(v *wrapperspb.BoolValue) {
	m.OutputFormatJsonQuoteDenormals = v
}

func (m *UserSettings) SetDateTimeOutputFormat(v UserSettings_DateTimeOutputFormat) {
	m.DateTimeOutputFormat = v
}

func (m *UserSettings) SetLowCardinalityAllowInNativeFormat(v *wrapperspb.BoolValue) {
	m.LowCardinalityAllowInNativeFormat = v
}

func (m *UserSettings) SetAllowSuspiciousLowCardinalityTypes(v *wrapperspb.BoolValue) {
	m.AllowSuspiciousLowCardinalityTypes = v
}

func (m *UserSettings) SetEmptyResultForAggregationByEmptySet(v *wrapperspb.BoolValue) {
	m.EmptyResultForAggregationByEmptySet = v
}

func (m *UserSettings) SetHttpConnectionTimeout(v *wrapperspb.Int64Value) {
	m.HttpConnectionTimeout = v
}

func (m *UserSettings) SetHttpReceiveTimeout(v *wrapperspb.Int64Value) {
	m.HttpReceiveTimeout = v
}

func (m *UserSettings) SetHttpSendTimeout(v *wrapperspb.Int64Value) {
	m.HttpSendTimeout = v
}

func (m *UserSettings) SetEnableHttpCompression(v *wrapperspb.BoolValue) {
	m.EnableHttpCompression = v
}

func (m *UserSettings) SetSendProgressInHttpHeaders(v *wrapperspb.BoolValue) {
	m.SendProgressInHttpHeaders = v
}

func (m *UserSettings) SetHttpHeadersProgressInterval(v *wrapperspb.Int64Value) {
	m.HttpHeadersProgressInterval = v
}

func (m *UserSettings) SetAddHttpCorsHeader(v *wrapperspb.BoolValue) {
	m.AddHttpCorsHeader = v
}

func (m *UserSettings) SetCancelHttpReadonlyQueriesOnClientClose(v *wrapperspb.BoolValue) {
	m.CancelHttpReadonlyQueriesOnClientClose = v
}

func (m *UserSettings) SetMaxHttpGetRedirects(v *wrapperspb.Int64Value) {
	m.MaxHttpGetRedirects = v
}

func (m *UserSettings) SetHttpMaxFieldNameSize(v *wrapperspb.Int64Value) {
	m.HttpMaxFieldNameSize = v
}

func (m *UserSettings) SetHttpMaxFieldValueSize(v *wrapperspb.Int64Value) {
	m.HttpMaxFieldValueSize = v
}

func (m *UserSettings) SetJoinedSubqueryRequiresAlias(v *wrapperspb.BoolValue) {
	m.JoinedSubqueryRequiresAlias = v
}

func (m *UserSettings) SetJoinUseNulls(v *wrapperspb.BoolValue) {
	m.JoinUseNulls = v
}

func (m *UserSettings) SetTransformNullIn(v *wrapperspb.BoolValue) {
	m.TransformNullIn = v
}

func (m *UserSettings) SetQuotaMode(v UserSettings_QuotaMode) {
	m.QuotaMode = v
}

func (m *UserSettings) SetFlattenNested(v *wrapperspb.BoolValue) {
	m.FlattenNested = v
}

func (m *UserSettings) SetFormatRegexp(v string) {
	m.FormatRegexp = v
}

func (m *UserSettings) SetFormatRegexpEscapingRule(v UserSettings_FormatRegexpEscapingRule) {
	m.FormatRegexpEscapingRule = v
}

func (m *UserSettings) SetFormatRegexpSkipUnmatched(v *wrapperspb.BoolValue) {
	m.FormatRegexpSkipUnmatched = v
}

func (m *UserSettings) SetAsyncInsert(v *wrapperspb.BoolValue) {
	m.AsyncInsert = v
}

func (m *UserSettings) SetAsyncInsertThreads(v *wrapperspb.Int64Value) {
	m.AsyncInsertThreads = v
}

func (m *UserSettings) SetWaitForAsyncInsert(v *wrapperspb.BoolValue) {
	m.WaitForAsyncInsert = v
}

func (m *UserSettings) SetWaitForAsyncInsertTimeout(v *wrapperspb.Int64Value) {
	m.WaitForAsyncInsertTimeout = v
}

func (m *UserSettings) SetAsyncInsertMaxDataSize(v *wrapperspb.Int64Value) {
	m.AsyncInsertMaxDataSize = v
}

func (m *UserSettings) SetAsyncInsertBusyTimeout(v *wrapperspb.Int64Value) {
	m.AsyncInsertBusyTimeout = v
}

func (m *UserSettings) SetAsyncInsertUseAdaptiveBusyTimeout(v *wrapperspb.BoolValue) {
	m.AsyncInsertUseAdaptiveBusyTimeout = v
}

func (m *UserSettings) SetMemoryProfilerStep(v *wrapperspb.Int64Value) {
	m.MemoryProfilerStep = v
}

func (m *UserSettings) SetMemoryProfilerSampleProbability(v *wrapperspb.DoubleValue) {
	m.MemoryProfilerSampleProbability = v
}

func (m *UserSettings) SetMaxFinalThreads(v *wrapperspb.Int64Value) {
	m.MaxFinalThreads = v
}

func (m *UserSettings) SetInputFormatParallelParsing(v *wrapperspb.BoolValue) {
	m.InputFormatParallelParsing = v
}

func (m *UserSettings) SetInputFormatImportNestedJson(v *wrapperspb.BoolValue) {
	m.InputFormatImportNestedJson = v
}

func (m *UserSettings) SetFormatAvroSchemaRegistryUrl(v string) {
	m.FormatAvroSchemaRegistryUrl = v
}

func (m *UserSettings) SetDataTypeDefaultNullable(v *wrapperspb.BoolValue) {
	m.DataTypeDefaultNullable = v
}

func (m *UserSettings) SetLocalFilesystemReadMethod(v UserSettings_LocalFilesystemReadMethod) {
	m.LocalFilesystemReadMethod = v
}

func (m *UserSettings) SetMaxReadBufferSize(v *wrapperspb.Int64Value) {
	m.MaxReadBufferSize = v
}

func (m *UserSettings) SetInsertKeeperMaxRetries(v *wrapperspb.Int64Value) {
	m.InsertKeeperMaxRetries = v
}

func (m *UserSettings) SetMaxTemporaryDataOnDiskSizeForUser(v *wrapperspb.Int64Value) {
	m.MaxTemporaryDataOnDiskSizeForUser = v
}

func (m *UserSettings) SetMaxTemporaryDataOnDiskSizeForQuery(v *wrapperspb.Int64Value) {
	m.MaxTemporaryDataOnDiskSizeForQuery = v
}

func (m *UserSettings) SetMaxParserDepth(v *wrapperspb.Int64Value) {
	m.MaxParserDepth = v
}

func (m *UserSettings) SetRemoteFilesystemReadMethod(v UserSettings_RemoteFilesystemReadMethod) {
	m.RemoteFilesystemReadMethod = v
}

func (m *UserSettings) SetMemoryOvercommitRatioDenominator(v *wrapperspb.Int64Value) {
	m.MemoryOvercommitRatioDenominator = v
}

func (m *UserSettings) SetMemoryOvercommitRatioDenominatorForUser(v *wrapperspb.Int64Value) {
	m.MemoryOvercommitRatioDenominatorForUser = v
}

func (m *UserSettings) SetMemoryUsageOvercommitMaxWaitMicroseconds(v *wrapperspb.Int64Value) {
	m.MemoryUsageOvercommitMaxWaitMicroseconds = v
}

func (m *UserSettings) SetLogQueryThreads(v *wrapperspb.BoolValue) {
	m.LogQueryThreads = v
}

func (m *UserSettings) SetLogQueryViews(v *wrapperspb.BoolValue) {
	m.LogQueryViews = v
}

func (m *UserSettings) SetLogQueriesProbability(v *wrapperspb.DoubleValue) {
	m.LogQueriesProbability = v
}

func (m *UserSettings) SetLogProcessorsProfiles(v *wrapperspb.BoolValue) {
	m.LogProcessorsProfiles = v
}

func (m *UserSettings) SetUseQueryCache(v *wrapperspb.BoolValue) {
	m.UseQueryCache = v
}

func (m *UserSettings) SetEnableReadsFromQueryCache(v *wrapperspb.BoolValue) {
	m.EnableReadsFromQueryCache = v
}

func (m *UserSettings) SetEnableWritesToQueryCache(v *wrapperspb.BoolValue) {
	m.EnableWritesToQueryCache = v
}

func (m *UserSettings) SetQueryCacheMinQueryRuns(v *wrapperspb.Int64Value) {
	m.QueryCacheMinQueryRuns = v
}

func (m *UserSettings) SetQueryCacheMinQueryDuration(v *wrapperspb.Int64Value) {
	m.QueryCacheMinQueryDuration = v
}

func (m *UserSettings) SetQueryCacheTtl(v *wrapperspb.Int64Value) {
	m.QueryCacheTtl = v
}

func (m *UserSettings) SetQueryCacheMaxEntries(v *wrapperspb.Int64Value) {
	m.QueryCacheMaxEntries = v
}

func (m *UserSettings) SetQueryCacheMaxSizeInBytes(v *wrapperspb.Int64Value) {
	m.QueryCacheMaxSizeInBytes = v
}

func (m *UserSettings) SetQueryCacheTag(v string) {
	m.QueryCacheTag = v
}

func (m *UserSettings) SetQueryCacheShareBetweenUsers(v *wrapperspb.BoolValue) {
	m.QueryCacheShareBetweenUsers = v
}

func (m *UserSettings) SetQueryCacheNondeterministicFunctionHandling(v UserSettings_QueryCacheNondeterministicFunctionHandling) {
	m.QueryCacheNondeterministicFunctionHandling = v
}

func (m *UserSettings) SetQueryCacheSystemTableHandling(v UserSettings_QueryCacheSystemTableHandling) {
	m.QueryCacheSystemTableHandling = v
}

func (m *UserSettings) SetMaxInsertThreads(v *wrapperspb.Int64Value) {
	m.MaxInsertThreads = v
}

func (m *UserSettings) SetUseHedgedRequests(v *wrapperspb.BoolValue) {
	m.UseHedgedRequests = v
}

func (m *UserSettings) SetIdleConnectionTimeout(v *wrapperspb.Int64Value) {
	m.IdleConnectionTimeout = v
}

func (m *UserSettings) SetHedgedConnectionTimeoutMs(v *wrapperspb.Int64Value) {
	m.HedgedConnectionTimeoutMs = v
}

func (m *UserSettings) SetLoadBalancing(v UserSettings_LoadBalancing) {
	m.LoadBalancing = v
}

func (m *UserSettings) SetPreferLocalhostReplica(v *wrapperspb.BoolValue) {
	m.PreferLocalhostReplica = v
}

func (m *UserSettings) SetDoNotMergeAcrossPartitionsSelectFinal(v *wrapperspb.BoolValue) {
	m.DoNotMergeAcrossPartitionsSelectFinal = v
}

func (m *UserSettings) SetIgnoreMaterializedViewsWithDroppedTargetTable(v *wrapperspb.BoolValue) {
	m.IgnoreMaterializedViewsWithDroppedTargetTable = v
}

func (m *UserSettings) SetEnableAnalyzer(v *wrapperspb.BoolValue) {
	m.EnableAnalyzer = v
}

func (m *UserSettings) SetCompile(v *wrapperspb.BoolValue) {
	m.Compile = v
}

func (m *UserSettings) SetMinCountToCompile(v *wrapperspb.Int64Value) {
	m.MinCountToCompile = v
}

func (m *UserSettings) SetAsyncInsertStaleTimeout(v *wrapperspb.Int64Value) {
	m.AsyncInsertStaleTimeout = v
}

func (m *UserQuota) SetIntervalDuration(v *wrapperspb.Int64Value) {
	m.IntervalDuration = v
}

func (m *UserQuota) SetQueries(v *wrapperspb.Int64Value) {
	m.Queries = v
}

func (m *UserQuota) SetErrors(v *wrapperspb.Int64Value) {
	m.Errors = v
}

func (m *UserQuota) SetResultRows(v *wrapperspb.Int64Value) {
	m.ResultRows = v
}

func (m *UserQuota) SetReadRows(v *wrapperspb.Int64Value) {
	m.ReadRows = v
}

func (m *UserQuota) SetExecutionTime(v *wrapperspb.Int64Value) {
	m.ExecutionTime = v
}
