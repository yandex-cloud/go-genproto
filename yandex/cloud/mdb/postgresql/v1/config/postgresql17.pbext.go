// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlConfig17) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig17) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig17) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig17) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig17) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig17) SetMaintenanceWorkMem(v *wrapperspb.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig17) SetAutovacuumWorkMem(v *wrapperspb.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig17) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig17) SetVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig17) SetVacuumCostPageHit(v *wrapperspb.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig17) SetVacuumCostPageMiss(v *wrapperspb.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig17) SetVacuumCostPageDirty(v *wrapperspb.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig17) SetVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig17) SetBgwriterDelay(v *wrapperspb.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig17) SetBgwriterLruMaxpages(v *wrapperspb.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig17) SetBgwriterLruMultiplier(v *wrapperspb.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig17) SetBgwriterFlushAfter(v *wrapperspb.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig17) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig17) SetWalLevel(v PostgresqlConfig17_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig17) SetSynchronousCommit(v PostgresqlConfig17_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig17) SetCheckpointTimeout(v *wrapperspb.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig17) SetCheckpointCompletionTarget(v *wrapperspb.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig17) SetCheckpointFlushAfter(v *wrapperspb.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig17) SetMaxWalSize(v *wrapperspb.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig17) SetMinWalSize(v *wrapperspb.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig17) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig17) SetDefaultStatisticsTarget(v *wrapperspb.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig17) SetConstraintExclusion(v PostgresqlConfig17_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig17) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig17) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig17) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig17) SetDebugParallelQuery(v PostgresqlConfig17_DebugParallelQuery) {
	m.DebugParallelQuery = v
}

func (m *PostgresqlConfig17) SetClientMinMessages(v PostgresqlConfig17_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig17) SetLogMinMessages(v PostgresqlConfig17_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig17) SetLogMinErrorStatement(v PostgresqlConfig17_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig17) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig17) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig17) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig17) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig17) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig17) SetLogErrorVerbosity(v PostgresqlConfig17_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig17) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig17) SetLogStatement(v PostgresqlConfig17_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig17) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig17) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig17) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig17) SetDefaultTransactionIsolation(v PostgresqlConfig17_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig17) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig17) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig17) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig17) SetByteaOutput(v PostgresqlConfig17_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig17) SetXmlbinary(v PostgresqlConfig17_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig17) SetXmloption(v PostgresqlConfig17_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig17) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig17) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig17) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig17) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig17) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig17) SetBackslashQuote(v PostgresqlConfig17_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig17) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig17) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig17) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig17) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig17) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig17) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig17) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig17) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig17) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig17) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig17) SetAutovacuumMaxWorkers(v *wrapperspb.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig17) SetAutovacuumVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig17) SetAutovacuumVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig17) SetAutovacuumNaptime(v *wrapperspb.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig17) SetArchiveTimeout(v *wrapperspb.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig17) SetTrackActivityQuerySize(v *wrapperspb.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig17) SetEnableBitmapscan(v *wrapperspb.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig17) SetEnableHashagg(v *wrapperspb.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig17) SetEnableHashjoin(v *wrapperspb.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig17) SetEnableIndexscan(v *wrapperspb.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig17) SetEnableIndexonlyscan(v *wrapperspb.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig17) SetEnableMaterial(v *wrapperspb.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig17) SetEnableMergejoin(v *wrapperspb.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig17) SetEnableNestloop(v *wrapperspb.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig17) SetEnableSeqscan(v *wrapperspb.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig17) SetEnableSort(v *wrapperspb.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig17) SetEnableTidscan(v *wrapperspb.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig17) SetMaxWorkerProcesses(v *wrapperspb.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig17) SetMaxParallelWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig17) SetMaxParallelWorkersPerGather(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig17) SetAutovacuumVacuumScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig17) SetAutovacuumAnalyzeScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig17) SetDefaultTransactionReadOnly(v *wrapperspb.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig17) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig17) SetEnableParallelAppend(v *wrapperspb.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig17) SetEnableParallelHash(v *wrapperspb.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig17) SetEnablePartitionPruning(v *wrapperspb.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig17) SetEnablePartitionwiseAggregate(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig17) SetEnablePartitionwiseJoin(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig17) SetJit(v *wrapperspb.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig17) SetMaxParallelMaintenanceWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig17) SetParallelLeaderParticipation(v *wrapperspb.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig17) SetLogTransactionSampleRate(v *wrapperspb.DoubleValue) {
	m.LogTransactionSampleRate = v
}

func (m *PostgresqlConfig17) SetPlanCacheMode(v PostgresqlConfig17_PlanCacheMode) {
	m.PlanCacheMode = v
}

func (m *PostgresqlConfig17) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig17) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig17) SetSharedPreloadLibraries(v []PostgresqlConfig17_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogMinDuration(v *wrapperspb.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogAnalyze(v *wrapperspb.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogBuffers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogTiming(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogTriggers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogVerbose(v *wrapperspb.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogNestedStatements(v *wrapperspb.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig17) SetAutoExplainSampleRate(v *wrapperspb.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig17) SetPgHintPlanEnableHint(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig17) SetPgHintPlanEnableHintTable(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig17) SetPgHintPlanDebugPrint(v PostgresqlConfig17_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig17) SetPgHintPlanMessageLevel(v PostgresqlConfig17_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfig17) SetHashMemMultiplier(v *wrapperspb.DoubleValue) {
	m.HashMemMultiplier = v
}

func (m *PostgresqlConfig17) SetLogicalDecodingWorkMem(v *wrapperspb.Int64Value) {
	m.LogicalDecodingWorkMem = v
}

func (m *PostgresqlConfig17) SetMaintenanceIoConcurrency(v *wrapperspb.Int64Value) {
	m.MaintenanceIoConcurrency = v
}

func (m *PostgresqlConfig17) SetMaxSlotWalKeepSize(v *wrapperspb.Int64Value) {
	m.MaxSlotWalKeepSize = v
}

func (m *PostgresqlConfig17) SetWalKeepSize(v *wrapperspb.Int64Value) {
	m.WalKeepSize = v
}

func (m *PostgresqlConfig17) SetEnableIncrementalSort(v *wrapperspb.BoolValue) {
	m.EnableIncrementalSort = v
}

func (m *PostgresqlConfig17) SetAutovacuumVacuumInsertThreshold(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumInsertThreshold = v
}

func (m *PostgresqlConfig17) SetAutovacuumVacuumInsertScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumInsertScaleFactor = v
}

func (m *PostgresqlConfig17) SetLogMinDurationSample(v *wrapperspb.Int64Value) {
	m.LogMinDurationSample = v
}

func (m *PostgresqlConfig17) SetLogStatementSampleRate(v *wrapperspb.DoubleValue) {
	m.LogStatementSampleRate = v
}

func (m *PostgresqlConfig17) SetLogParameterMaxLength(v *wrapperspb.Int64Value) {
	m.LogParameterMaxLength = v
}

func (m *PostgresqlConfig17) SetLogParameterMaxLengthOnError(v *wrapperspb.Int64Value) {
	m.LogParameterMaxLengthOnError = v
}

func (m *PostgresqlConfig17) SetClientConnectionCheckInterval(v *wrapperspb.Int64Value) {
	m.ClientConnectionCheckInterval = v
}

func (m *PostgresqlConfig17) SetEnableAsyncAppend(v *wrapperspb.BoolValue) {
	m.EnableAsyncAppend = v
}

func (m *PostgresqlConfig17) SetEnableGathermerge(v *wrapperspb.BoolValue) {
	m.EnableGathermerge = v
}

func (m *PostgresqlConfig17) SetEnableMemoize(v *wrapperspb.BoolValue) {
	m.EnableMemoize = v
}

func (m *PostgresqlConfig17) SetLogRecoveryConflictWaits(v *wrapperspb.BoolValue) {
	m.LogRecoveryConflictWaits = v
}

func (m *PostgresqlConfig17) SetVacuumFailsafeAge(v *wrapperspb.Int64Value) {
	m.VacuumFailsafeAge = v
}

func (m *PostgresqlConfig17) SetVacuumMultixactFailsafeAge(v *wrapperspb.Int64Value) {
	m.VacuumMultixactFailsafeAge = v
}

func (m *PostgresqlConfig17) SetPgQualstatsEnabled(v *wrapperspb.BoolValue) {
	m.PgQualstatsEnabled = v
}

func (m *PostgresqlConfig17) SetPgQualstatsTrackConstants(v *wrapperspb.BoolValue) {
	m.PgQualstatsTrackConstants = v
}

func (m *PostgresqlConfig17) SetPgQualstatsMax(v *wrapperspb.Int64Value) {
	m.PgQualstatsMax = v
}

func (m *PostgresqlConfig17) SetPgQualstatsResolveOids(v *wrapperspb.BoolValue) {
	m.PgQualstatsResolveOids = v
}

func (m *PostgresqlConfig17) SetPgQualstatsSampleRate(v *wrapperspb.DoubleValue) {
	m.PgQualstatsSampleRate = v
}

func (m *PostgresqlConfig17) SetMaxStackDepth(v *wrapperspb.Int64Value) {
	m.MaxStackDepth = v
}

func (m *PostgresqlConfig17) SetEnableGroupByReordering(v *wrapperspb.BoolValue) {
	m.EnableGroupByReordering = v
}

func (m *PostgresqlConfig17) SetGeqo(v *wrapperspb.BoolValue) {
	m.Geqo = v
}

func (m *PostgresqlConfig17) SetGeqoThreshold(v *wrapperspb.Int64Value) {
	m.GeqoThreshold = v
}

func (m *PostgresqlConfig17) SetGeqoEffort(v *wrapperspb.Int64Value) {
	m.GeqoEffort = v
}

func (m *PostgresqlConfig17) SetGeqoPoolSize(v *wrapperspb.Int64Value) {
	m.GeqoPoolSize = v
}

func (m *PostgresqlConfig17) SetGeqoGenerations(v *wrapperspb.Int64Value) {
	m.GeqoGenerations = v
}

func (m *PostgresqlConfig17) SetGeqoSelectionBias(v *wrapperspb.DoubleValue) {
	m.GeqoSelectionBias = v
}

func (m *PostgresqlConfig17) SetGeqoSeed(v *wrapperspb.DoubleValue) {
	m.GeqoSeed = v
}

func (m *PostgresqlConfig17) SetPgTrgmSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmSimilarityThreshold = v
}

func (m *PostgresqlConfig17) SetPgTrgmWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmWordSimilarityThreshold = v
}

func (m *PostgresqlConfig17) SetPgTrgmStrictWordSimilarityThreshold(v *wrapperspb.DoubleValue) {
	m.PgTrgmStrictWordSimilarityThreshold = v
}

func (m *PostgresqlConfig17) SetMaxStandbyArchiveDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyArchiveDelay = v
}

func (m *PostgresqlConfig17) SetSessionDurationTimeout(v *wrapperspb.Int64Value) {
	m.SessionDurationTimeout = v
}

func (m *PostgresqlConfig17) SetLogReplicationCommands(v *wrapperspb.BoolValue) {
	m.LogReplicationCommands = v
}

func (m *PostgresqlConfig17) SetLogAutovacuumMinDuration(v *wrapperspb.Int64Value) {
	m.LogAutovacuumMinDuration = v
}

func (m *PostgresqlConfig17) SetPasswordEncryption(v PostgresqlConfig17_PasswordEncryption) {
	m.PasswordEncryption = v
}

func (m *PostgresqlConfig17) SetAutoExplainLogFormat(v PostgresqlConfig17_AutoExplainLogFormat) {
	m.AutoExplainLogFormat = v
}

func (m *PostgresqlConfig17) SetTrackCommitTimestamp(v *wrapperspb.BoolValue) {
	m.TrackCommitTimestamp = v
}

func (m *PostgresqlConfig17) SetMaxLogicalReplicationWorkers(v *wrapperspb.Int64Value) {
	m.MaxLogicalReplicationWorkers = v
}

func (m *PostgresqlConfig17) SetMaxWalSenders(v *wrapperspb.Int64Value) {
	m.MaxWalSenders = v
}

func (m *PostgresqlConfig17) SetMaxReplicationSlots(v *wrapperspb.Int64Value) {
	m.MaxReplicationSlots = v
}

func (m *PostgresqlConfigSet17) SetEffectiveConfig(v *PostgresqlConfig17) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet17) SetUserConfig(v *PostgresqlConfig17) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet17) SetDefaultConfig(v *PostgresqlConfig17) {
	m.DefaultConfig = v
}