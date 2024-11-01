// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlConfig17_1C) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig17_1C) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig17_1C) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig17_1C) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig17_1C) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig17_1C) SetMaintenanceWorkMem(v *wrapperspb.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumWorkMem(v *wrapperspb.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig17_1C) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig17_1C) SetVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig17_1C) SetVacuumCostPageHit(v *wrapperspb.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig17_1C) SetVacuumCostPageMiss(v *wrapperspb.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig17_1C) SetVacuumCostPageDirty(v *wrapperspb.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig17_1C) SetVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig17_1C) SetBgwriterDelay(v *wrapperspb.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig17_1C) SetBgwriterLruMaxpages(v *wrapperspb.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig17_1C) SetBgwriterLruMultiplier(v *wrapperspb.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig17_1C) SetBgwriterFlushAfter(v *wrapperspb.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig17_1C) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig17_1C) SetWalLevel(v PostgresqlConfig17_1C_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig17_1C) SetSynchronousCommit(v PostgresqlConfig17_1C_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig17_1C) SetCheckpointTimeout(v *wrapperspb.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig17_1C) SetCheckpointCompletionTarget(v *wrapperspb.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig17_1C) SetCheckpointFlushAfter(v *wrapperspb.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig17_1C) SetMaxWalSize(v *wrapperspb.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig17_1C) SetMinWalSize(v *wrapperspb.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig17_1C) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig17_1C) SetDefaultStatisticsTarget(v *wrapperspb.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig17_1C) SetConstraintExclusion(v PostgresqlConfig17_1C_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig17_1C) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig17_1C) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig17_1C) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig17_1C) SetDebugParallelQuery(v PostgresqlConfig17_1C_DebugParallelQuery) {
	m.DebugParallelQuery = v
}

func (m *PostgresqlConfig17_1C) SetClientMinMessages(v PostgresqlConfig17_1C_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig17_1C) SetLogMinMessages(v PostgresqlConfig17_1C_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig17_1C) SetLogMinErrorStatement(v PostgresqlConfig17_1C_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig17_1C) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig17_1C) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig17_1C) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig17_1C) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig17_1C) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig17_1C) SetLogErrorVerbosity(v PostgresqlConfig17_1C_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig17_1C) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig17_1C) SetLogStatement(v PostgresqlConfig17_1C_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig17_1C) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig17_1C) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig17_1C) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig17_1C) SetDefaultTransactionIsolation(v PostgresqlConfig17_1C_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig17_1C) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig17_1C) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig17_1C) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig17_1C) SetByteaOutput(v PostgresqlConfig17_1C_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig17_1C) SetXmlbinary(v PostgresqlConfig17_1C_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig17_1C) SetXmloption(v PostgresqlConfig17_1C_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig17_1C) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig17_1C) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig17_1C) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig17_1C) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig17_1C) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig17_1C) SetBackslashQuote(v PostgresqlConfig17_1C_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig17_1C) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig17_1C) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig17_1C) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig17_1C) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig17_1C) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig17_1C) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig17_1C) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig17_1C) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig17_1C) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig17_1C) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumMaxWorkers(v *wrapperspb.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumNaptime(v *wrapperspb.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig17_1C) SetArchiveTimeout(v *wrapperspb.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig17_1C) SetTrackActivityQuerySize(v *wrapperspb.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig17_1C) SetOnlineAnalyzeEnable(v *wrapperspb.BoolValue) {
	m.OnlineAnalyzeEnable = v
}

func (m *PostgresqlConfig17_1C) SetEnableBitmapscan(v *wrapperspb.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig17_1C) SetEnableHashagg(v *wrapperspb.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig17_1C) SetEnableHashjoin(v *wrapperspb.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig17_1C) SetEnableIndexscan(v *wrapperspb.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig17_1C) SetEnableIndexonlyscan(v *wrapperspb.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig17_1C) SetEnableMaterial(v *wrapperspb.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig17_1C) SetEnableMergejoin(v *wrapperspb.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig17_1C) SetEnableNestloop(v *wrapperspb.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig17_1C) SetEnableSeqscan(v *wrapperspb.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig17_1C) SetEnableSort(v *wrapperspb.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig17_1C) SetEnableTidscan(v *wrapperspb.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig17_1C) SetMaxWorkerProcesses(v *wrapperspb.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig17_1C) SetMaxParallelWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig17_1C) SetMaxParallelWorkersPerGather(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumVacuumScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumAnalyzeScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig17_1C) SetDefaultTransactionReadOnly(v *wrapperspb.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig17_1C) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig17_1C) SetEnableParallelAppend(v *wrapperspb.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig17_1C) SetEnableParallelHash(v *wrapperspb.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig17_1C) SetEnablePartitionPruning(v *wrapperspb.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig17_1C) SetEnablePartitionwiseAggregate(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig17_1C) SetEnablePartitionwiseJoin(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig17_1C) SetJit(v *wrapperspb.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig17_1C) SetMaxParallelMaintenanceWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig17_1C) SetParallelLeaderParticipation(v *wrapperspb.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig17_1C) SetLogTransactionSampleRate(v *wrapperspb.DoubleValue) {
	m.LogTransactionSampleRate = v
}

func (m *PostgresqlConfig17_1C) SetPlanCacheMode(v PostgresqlConfig17_1C_PlanCacheMode) {
	m.PlanCacheMode = v
}

func (m *PostgresqlConfig17_1C) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig17_1C) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig17_1C) SetSharedPreloadLibraries(v []PostgresqlConfig17_1C_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogMinDuration(v *wrapperspb.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogAnalyze(v *wrapperspb.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogBuffers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogTiming(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogTriggers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogVerbose(v *wrapperspb.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogNestedStatements(v *wrapperspb.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainSampleRate(v *wrapperspb.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig17_1C) SetPgHintPlanEnableHint(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig17_1C) SetPgHintPlanEnableHintTable(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig17_1C) SetPgHintPlanDebugPrint(v PostgresqlConfig17_1C_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig17_1C) SetPgHintPlanMessageLevel(v PostgresqlConfig17_1C_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfig17_1C) SetHashMemMultiplier(v *wrapperspb.DoubleValue) {
	m.HashMemMultiplier = v
}

func (m *PostgresqlConfig17_1C) SetLogicalDecodingWorkMem(v *wrapperspb.Int64Value) {
	m.LogicalDecodingWorkMem = v
}

func (m *PostgresqlConfig17_1C) SetMaintenanceIoConcurrency(v *wrapperspb.Int64Value) {
	m.MaintenanceIoConcurrency = v
}

func (m *PostgresqlConfig17_1C) SetMaxSlotWalKeepSize(v *wrapperspb.Int64Value) {
	m.MaxSlotWalKeepSize = v
}

func (m *PostgresqlConfig17_1C) SetWalKeepSize(v *wrapperspb.Int64Value) {
	m.WalKeepSize = v
}

func (m *PostgresqlConfig17_1C) SetEnableIncrementalSort(v *wrapperspb.BoolValue) {
	m.EnableIncrementalSort = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumVacuumInsertThreshold(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumInsertThreshold = v
}

func (m *PostgresqlConfig17_1C) SetAutovacuumVacuumInsertScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumInsertScaleFactor = v
}

func (m *PostgresqlConfig17_1C) SetLogMinDurationSample(v *wrapperspb.Int64Value) {
	m.LogMinDurationSample = v
}

func (m *PostgresqlConfig17_1C) SetLogStatementSampleRate(v *wrapperspb.DoubleValue) {
	m.LogStatementSampleRate = v
}

func (m *PostgresqlConfig17_1C) SetLogParameterMaxLength(v *wrapperspb.Int64Value) {
	m.LogParameterMaxLength = v
}

func (m *PostgresqlConfig17_1C) SetLogParameterMaxLengthOnError(v *wrapperspb.Int64Value) {
	m.LogParameterMaxLengthOnError = v
}

func (m *PostgresqlConfig17_1C) SetClientConnectionCheckInterval(v *wrapperspb.Int64Value) {
	m.ClientConnectionCheckInterval = v
}

func (m *PostgresqlConfig17_1C) SetEnableAsyncAppend(v *wrapperspb.BoolValue) {
	m.EnableAsyncAppend = v
}

func (m *PostgresqlConfig17_1C) SetEnableGathermerge(v *wrapperspb.BoolValue) {
	m.EnableGathermerge = v
}

func (m *PostgresqlConfig17_1C) SetEnableMemoize(v *wrapperspb.BoolValue) {
	m.EnableMemoize = v
}

func (m *PostgresqlConfig17_1C) SetLogRecoveryConflictWaits(v *wrapperspb.BoolValue) {
	m.LogRecoveryConflictWaits = v
}

func (m *PostgresqlConfig17_1C) SetVacuumFailsafeAge(v *wrapperspb.Int64Value) {
	m.VacuumFailsafeAge = v
}

func (m *PostgresqlConfig17_1C) SetVacuumMultixactFailsafeAge(v *wrapperspb.Int64Value) {
	m.VacuumMultixactFailsafeAge = v
}

func (m *PostgresqlConfig17_1C) SetPgQualstatsEnabled(v *wrapperspb.BoolValue) {
	m.PgQualstatsEnabled = v
}

func (m *PostgresqlConfig17_1C) SetPgQualstatsTrackConstants(v *wrapperspb.BoolValue) {
	m.PgQualstatsTrackConstants = v
}

func (m *PostgresqlConfig17_1C) SetPgQualstatsMax(v *wrapperspb.Int64Value) {
	m.PgQualstatsMax = v
}

func (m *PostgresqlConfig17_1C) SetPgQualstatsResolveOids(v *wrapperspb.BoolValue) {
	m.PgQualstatsResolveOids = v
}

func (m *PostgresqlConfig17_1C) SetPgQualstatsSampleRate(v *wrapperspb.DoubleValue) {
	m.PgQualstatsSampleRate = v
}

func (m *PostgresqlConfig17_1C) SetPlantunerFixEmptyTable(v *wrapperspb.BoolValue) {
	m.PlantunerFixEmptyTable = v
}

func (m *PostgresqlConfig17_1C) SetMaxStackDepth(v *wrapperspb.Int64Value) {
	m.MaxStackDepth = v
}

func (m *PostgresqlConfig17_1C) SetEnableGroupByReordering(v *wrapperspb.BoolValue) {
	m.EnableGroupByReordering = v
}

func (m *PostgresqlConfig17_1C) SetGeqo(v *wrapperspb.BoolValue) {
	m.Geqo = v
}

func (m *PostgresqlConfig17_1C) SetGeqoThreshold(v *wrapperspb.Int64Value) {
	m.GeqoThreshold = v
}

func (m *PostgresqlConfig17_1C) SetGeqoEffort(v *wrapperspb.Int64Value) {
	m.GeqoEffort = v
}

func (m *PostgresqlConfig17_1C) SetGeqoPoolSize(v *wrapperspb.Int64Value) {
	m.GeqoPoolSize = v
}

func (m *PostgresqlConfig17_1C) SetGeqoGenerations(v *wrapperspb.Int64Value) {
	m.GeqoGenerations = v
}

func (m *PostgresqlConfig17_1C) SetGeqoSelectionBias(v *wrapperspb.DoubleValue) {
	m.GeqoSelectionBias = v
}

func (m *PostgresqlConfig17_1C) SetGeqoSeed(v *wrapperspb.DoubleValue) {
	m.GeqoSeed = v
}

func (m *PostgresqlConfig17_1C) SetPgTrgmSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmSimilarityThreshold = v
}

func (m *PostgresqlConfig17_1C) SetPgTrgmWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmWordSimilarityThreshold = v
}

func (m *PostgresqlConfig17_1C) SetPgTrgmStrictWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmStrictWordSimilarityThreshold = v
}

func (m *PostgresqlConfig17_1C) SetMaxStandbyArchiveDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyArchiveDelay = v
}

func (m *PostgresqlConfig17_1C) SetSessionDurationTimeout(v *wrapperspb.Int64Value) {
	m.SessionDurationTimeout = v
}

func (m *PostgresqlConfig17_1C) SetLogReplicationCommands(v *wrapperspb.BoolValue) {
	m.LogReplicationCommands = v
}

func (m *PostgresqlConfig17_1C) SetLogAutovacuumMinDuration(v *wrapperspb.Int64Value) {
	m.LogAutovacuumMinDuration = v
}

func (m *PostgresqlConfig17_1C) SetPasswordEncryption(v PostgresqlConfig17_1C_PasswordEncryption) {
	m.PasswordEncryption = v
}

func (m *PostgresqlConfig17_1C) SetAutoExplainLogFormat(v PostgresqlConfig17_1C_AutoExplainLogFormat) {
	m.AutoExplainLogFormat = v
}

func (m *PostgresqlConfig17_1C) SetTrackCommitTimestamp(v *wrapperspb.BoolValue) {
	m.TrackCommitTimestamp = v
}

func (m *PostgresqlConfig17_1C) SetMaxLogicalReplicationWorkers(v *wrapperspb.Int64Value) {
	m.MaxLogicalReplicationWorkers = v
}

func (m *PostgresqlConfig17_1C) SetMaxWalSenders(v *wrapperspb.Int64Value) {
	m.MaxWalSenders = v
}

func (m *PostgresqlConfig17_1C) SetMaxReplicationSlots(v *wrapperspb.Int64Value) {
	m.MaxReplicationSlots = v
}

func (m *PostgresqlConfigSet17_1C) SetEffectiveConfig(v *PostgresqlConfig17_1C) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet17_1C) SetUserConfig(v *PostgresqlConfig17_1C) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet17_1C) SetDefaultConfig(v *PostgresqlConfig17_1C) {
	m.DefaultConfig = v
}
