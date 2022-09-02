// Code generated by protoc-gen-goext. DO NOT EDIT.

package postgresql

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *PostgresqlConfig11) SetMaxConnections(v *wrapperspb.Int64Value) {
	m.MaxConnections = v
}

func (m *PostgresqlConfig11) SetSharedBuffers(v *wrapperspb.Int64Value) {
	m.SharedBuffers = v
}

func (m *PostgresqlConfig11) SetTempBuffers(v *wrapperspb.Int64Value) {
	m.TempBuffers = v
}

func (m *PostgresqlConfig11) SetMaxPreparedTransactions(v *wrapperspb.Int64Value) {
	m.MaxPreparedTransactions = v
}

func (m *PostgresqlConfig11) SetWorkMem(v *wrapperspb.Int64Value) {
	m.WorkMem = v
}

func (m *PostgresqlConfig11) SetMaintenanceWorkMem(v *wrapperspb.Int64Value) {
	m.MaintenanceWorkMem = v
}

func (m *PostgresqlConfig11) SetAutovacuumWorkMem(v *wrapperspb.Int64Value) {
	m.AutovacuumWorkMem = v
}

func (m *PostgresqlConfig11) SetTempFileLimit(v *wrapperspb.Int64Value) {
	m.TempFileLimit = v
}

func (m *PostgresqlConfig11) SetVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.VacuumCostDelay = v
}

func (m *PostgresqlConfig11) SetVacuumCostPageHit(v *wrapperspb.Int64Value) {
	m.VacuumCostPageHit = v
}

func (m *PostgresqlConfig11) SetVacuumCostPageMiss(v *wrapperspb.Int64Value) {
	m.VacuumCostPageMiss = v
}

func (m *PostgresqlConfig11) SetVacuumCostPageDirty(v *wrapperspb.Int64Value) {
	m.VacuumCostPageDirty = v
}

func (m *PostgresqlConfig11) SetVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.VacuumCostLimit = v
}

func (m *PostgresqlConfig11) SetBgwriterDelay(v *wrapperspb.Int64Value) {
	m.BgwriterDelay = v
}

func (m *PostgresqlConfig11) SetBgwriterLruMaxpages(v *wrapperspb.Int64Value) {
	m.BgwriterLruMaxpages = v
}

func (m *PostgresqlConfig11) SetBgwriterLruMultiplier(v *wrapperspb.DoubleValue) {
	m.BgwriterLruMultiplier = v
}

func (m *PostgresqlConfig11) SetBgwriterFlushAfter(v *wrapperspb.Int64Value) {
	m.BgwriterFlushAfter = v
}

func (m *PostgresqlConfig11) SetBackendFlushAfter(v *wrapperspb.Int64Value) {
	m.BackendFlushAfter = v
}

func (m *PostgresqlConfig11) SetOldSnapshotThreshold(v *wrapperspb.Int64Value) {
	m.OldSnapshotThreshold = v
}

func (m *PostgresqlConfig11) SetWalLevel(v PostgresqlConfig11_WalLevel) {
	m.WalLevel = v
}

func (m *PostgresqlConfig11) SetSynchronousCommit(v PostgresqlConfig11_SynchronousCommit) {
	m.SynchronousCommit = v
}

func (m *PostgresqlConfig11) SetCheckpointTimeout(v *wrapperspb.Int64Value) {
	m.CheckpointTimeout = v
}

func (m *PostgresqlConfig11) SetCheckpointCompletionTarget(v *wrapperspb.DoubleValue) {
	m.CheckpointCompletionTarget = v
}

func (m *PostgresqlConfig11) SetCheckpointFlushAfter(v *wrapperspb.Int64Value) {
	m.CheckpointFlushAfter = v
}

func (m *PostgresqlConfig11) SetMaxWalSize(v *wrapperspb.Int64Value) {
	m.MaxWalSize = v
}

func (m *PostgresqlConfig11) SetMinWalSize(v *wrapperspb.Int64Value) {
	m.MinWalSize = v
}

func (m *PostgresqlConfig11) SetMaxStandbyStreamingDelay(v *wrapperspb.Int64Value) {
	m.MaxStandbyStreamingDelay = v
}

func (m *PostgresqlConfig11) SetDefaultStatisticsTarget(v *wrapperspb.Int64Value) {
	m.DefaultStatisticsTarget = v
}

func (m *PostgresqlConfig11) SetConstraintExclusion(v PostgresqlConfig11_ConstraintExclusion) {
	m.ConstraintExclusion = v
}

func (m *PostgresqlConfig11) SetCursorTupleFraction(v *wrapperspb.DoubleValue) {
	m.CursorTupleFraction = v
}

func (m *PostgresqlConfig11) SetFromCollapseLimit(v *wrapperspb.Int64Value) {
	m.FromCollapseLimit = v
}

func (m *PostgresqlConfig11) SetJoinCollapseLimit(v *wrapperspb.Int64Value) {
	m.JoinCollapseLimit = v
}

func (m *PostgresqlConfig11) SetForceParallelMode(v PostgresqlConfig11_ForceParallelMode) {
	m.ForceParallelMode = v
}

func (m *PostgresqlConfig11) SetClientMinMessages(v PostgresqlConfig11_LogLevel) {
	m.ClientMinMessages = v
}

func (m *PostgresqlConfig11) SetLogMinMessages(v PostgresqlConfig11_LogLevel) {
	m.LogMinMessages = v
}

func (m *PostgresqlConfig11) SetLogMinErrorStatement(v PostgresqlConfig11_LogLevel) {
	m.LogMinErrorStatement = v
}

func (m *PostgresqlConfig11) SetLogMinDurationStatement(v *wrapperspb.Int64Value) {
	m.LogMinDurationStatement = v
}

func (m *PostgresqlConfig11) SetLogCheckpoints(v *wrapperspb.BoolValue) {
	m.LogCheckpoints = v
}

func (m *PostgresqlConfig11) SetLogConnections(v *wrapperspb.BoolValue) {
	m.LogConnections = v
}

func (m *PostgresqlConfig11) SetLogDisconnections(v *wrapperspb.BoolValue) {
	m.LogDisconnections = v
}

func (m *PostgresqlConfig11) SetLogDuration(v *wrapperspb.BoolValue) {
	m.LogDuration = v
}

func (m *PostgresqlConfig11) SetLogErrorVerbosity(v PostgresqlConfig11_LogErrorVerbosity) {
	m.LogErrorVerbosity = v
}

func (m *PostgresqlConfig11) SetLogLockWaits(v *wrapperspb.BoolValue) {
	m.LogLockWaits = v
}

func (m *PostgresqlConfig11) SetLogStatement(v PostgresqlConfig11_LogStatement) {
	m.LogStatement = v
}

func (m *PostgresqlConfig11) SetLogTempFiles(v *wrapperspb.Int64Value) {
	m.LogTempFiles = v
}

func (m *PostgresqlConfig11) SetSearchPath(v string) {
	m.SearchPath = v
}

func (m *PostgresqlConfig11) SetRowSecurity(v *wrapperspb.BoolValue) {
	m.RowSecurity = v
}

func (m *PostgresqlConfig11) SetDefaultTransactionIsolation(v PostgresqlConfig11_TransactionIsolation) {
	m.DefaultTransactionIsolation = v
}

func (m *PostgresqlConfig11) SetStatementTimeout(v *wrapperspb.Int64Value) {
	m.StatementTimeout = v
}

func (m *PostgresqlConfig11) SetLockTimeout(v *wrapperspb.Int64Value) {
	m.LockTimeout = v
}

func (m *PostgresqlConfig11) SetIdleInTransactionSessionTimeout(v *wrapperspb.Int64Value) {
	m.IdleInTransactionSessionTimeout = v
}

func (m *PostgresqlConfig11) SetByteaOutput(v PostgresqlConfig11_ByteaOutput) {
	m.ByteaOutput = v
}

func (m *PostgresqlConfig11) SetXmlbinary(v PostgresqlConfig11_XmlBinary) {
	m.Xmlbinary = v
}

func (m *PostgresqlConfig11) SetXmloption(v PostgresqlConfig11_XmlOption) {
	m.Xmloption = v
}

func (m *PostgresqlConfig11) SetGinPendingListLimit(v *wrapperspb.Int64Value) {
	m.GinPendingListLimit = v
}

func (m *PostgresqlConfig11) SetDeadlockTimeout(v *wrapperspb.Int64Value) {
	m.DeadlockTimeout = v
}

func (m *PostgresqlConfig11) SetMaxLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxLocksPerTransaction = v
}

func (m *PostgresqlConfig11) SetMaxPredLocksPerTransaction(v *wrapperspb.Int64Value) {
	m.MaxPredLocksPerTransaction = v
}

func (m *PostgresqlConfig11) SetArrayNulls(v *wrapperspb.BoolValue) {
	m.ArrayNulls = v
}

func (m *PostgresqlConfig11) SetBackslashQuote(v PostgresqlConfig11_BackslashQuote) {
	m.BackslashQuote = v
}

func (m *PostgresqlConfig11) SetDefaultWithOids(v *wrapperspb.BoolValue) {
	m.DefaultWithOids = v
}

func (m *PostgresqlConfig11) SetEscapeStringWarning(v *wrapperspb.BoolValue) {
	m.EscapeStringWarning = v
}

func (m *PostgresqlConfig11) SetLoCompatPrivileges(v *wrapperspb.BoolValue) {
	m.LoCompatPrivileges = v
}

func (m *PostgresqlConfig11) SetOperatorPrecedenceWarning(v *wrapperspb.BoolValue) {
	m.OperatorPrecedenceWarning = v
}

func (m *PostgresqlConfig11) SetQuoteAllIdentifiers(v *wrapperspb.BoolValue) {
	m.QuoteAllIdentifiers = v
}

func (m *PostgresqlConfig11) SetStandardConformingStrings(v *wrapperspb.BoolValue) {
	m.StandardConformingStrings = v
}

func (m *PostgresqlConfig11) SetSynchronizeSeqscans(v *wrapperspb.BoolValue) {
	m.SynchronizeSeqscans = v
}

func (m *PostgresqlConfig11) SetTransformNullEquals(v *wrapperspb.BoolValue) {
	m.TransformNullEquals = v
}

func (m *PostgresqlConfig11) SetExitOnError(v *wrapperspb.BoolValue) {
	m.ExitOnError = v
}

func (m *PostgresqlConfig11) SetSeqPageCost(v *wrapperspb.DoubleValue) {
	m.SeqPageCost = v
}

func (m *PostgresqlConfig11) SetRandomPageCost(v *wrapperspb.DoubleValue) {
	m.RandomPageCost = v
}

func (m *PostgresqlConfig11) SetAutovacuumMaxWorkers(v *wrapperspb.Int64Value) {
	m.AutovacuumMaxWorkers = v
}

func (m *PostgresqlConfig11) SetAutovacuumVacuumCostDelay(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostDelay = v
}

func (m *PostgresqlConfig11) SetAutovacuumVacuumCostLimit(v *wrapperspb.Int64Value) {
	m.AutovacuumVacuumCostLimit = v
}

func (m *PostgresqlConfig11) SetAutovacuumNaptime(v *wrapperspb.Int64Value) {
	m.AutovacuumNaptime = v
}

func (m *PostgresqlConfig11) SetArchiveTimeout(v *wrapperspb.Int64Value) {
	m.ArchiveTimeout = v
}

func (m *PostgresqlConfig11) SetTrackActivityQuerySize(v *wrapperspb.Int64Value) {
	m.TrackActivityQuerySize = v
}

func (m *PostgresqlConfig11) SetEnableBitmapscan(v *wrapperspb.BoolValue) {
	m.EnableBitmapscan = v
}

func (m *PostgresqlConfig11) SetEnableHashagg(v *wrapperspb.BoolValue) {
	m.EnableHashagg = v
}

func (m *PostgresqlConfig11) SetEnableHashjoin(v *wrapperspb.BoolValue) {
	m.EnableHashjoin = v
}

func (m *PostgresqlConfig11) SetEnableIndexscan(v *wrapperspb.BoolValue) {
	m.EnableIndexscan = v
}

func (m *PostgresqlConfig11) SetEnableIndexonlyscan(v *wrapperspb.BoolValue) {
	m.EnableIndexonlyscan = v
}

func (m *PostgresqlConfig11) SetEnableMaterial(v *wrapperspb.BoolValue) {
	m.EnableMaterial = v
}

func (m *PostgresqlConfig11) SetEnableMergejoin(v *wrapperspb.BoolValue) {
	m.EnableMergejoin = v
}

func (m *PostgresqlConfig11) SetEnableNestloop(v *wrapperspb.BoolValue) {
	m.EnableNestloop = v
}

func (m *PostgresqlConfig11) SetEnableSeqscan(v *wrapperspb.BoolValue) {
	m.EnableSeqscan = v
}

func (m *PostgresqlConfig11) SetEnableSort(v *wrapperspb.BoolValue) {
	m.EnableSort = v
}

func (m *PostgresqlConfig11) SetEnableTidscan(v *wrapperspb.BoolValue) {
	m.EnableTidscan = v
}

func (m *PostgresqlConfig11) SetMaxWorkerProcesses(v *wrapperspb.Int64Value) {
	m.MaxWorkerProcesses = v
}

func (m *PostgresqlConfig11) SetMaxParallelWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkers = v
}

func (m *PostgresqlConfig11) SetMaxParallelWorkersPerGather(v *wrapperspb.Int64Value) {
	m.MaxParallelWorkersPerGather = v
}

func (m *PostgresqlConfig11) SetAutovacuumVacuumScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumVacuumScaleFactor = v
}

func (m *PostgresqlConfig11) SetAutovacuumAnalyzeScaleFactor(v *wrapperspb.DoubleValue) {
	m.AutovacuumAnalyzeScaleFactor = v
}

func (m *PostgresqlConfig11) SetDefaultTransactionReadOnly(v *wrapperspb.BoolValue) {
	m.DefaultTransactionReadOnly = v
}

func (m *PostgresqlConfig11) SetTimezone(v string) {
	m.Timezone = v
}

func (m *PostgresqlConfig11) SetEnableParallelAppend(v *wrapperspb.BoolValue) {
	m.EnableParallelAppend = v
}

func (m *PostgresqlConfig11) SetEnableParallelHash(v *wrapperspb.BoolValue) {
	m.EnableParallelHash = v
}

func (m *PostgresqlConfig11) SetEnablePartitionPruning(v *wrapperspb.BoolValue) {
	m.EnablePartitionPruning = v
}

func (m *PostgresqlConfig11) SetEnablePartitionwiseAggregate(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseAggregate = v
}

func (m *PostgresqlConfig11) SetEnablePartitionwiseJoin(v *wrapperspb.BoolValue) {
	m.EnablePartitionwiseJoin = v
}

func (m *PostgresqlConfig11) SetJit(v *wrapperspb.BoolValue) {
	m.Jit = v
}

func (m *PostgresqlConfig11) SetMaxParallelMaintenanceWorkers(v *wrapperspb.Int64Value) {
	m.MaxParallelMaintenanceWorkers = v
}

func (m *PostgresqlConfig11) SetParallelLeaderParticipation(v *wrapperspb.BoolValue) {
	m.ParallelLeaderParticipation = v
}

func (m *PostgresqlConfig11) SetVacuumCleanupIndexScaleFactor(v *wrapperspb.DoubleValue) {
	m.VacuumCleanupIndexScaleFactor = v
}

func (m *PostgresqlConfig11) SetEffectiveIoConcurrency(v *wrapperspb.Int64Value) {
	m.EffectiveIoConcurrency = v
}

func (m *PostgresqlConfig11) SetEffectiveCacheSize(v *wrapperspb.Int64Value) {
	m.EffectiveCacheSize = v
}

func (m *PostgresqlConfig11) SetSharedPreloadLibraries(v []PostgresqlConfig11_SharedPreloadLibraries) {
	m.SharedPreloadLibraries = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogMinDuration(v *wrapperspb.Int64Value) {
	m.AutoExplainLogMinDuration = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogAnalyze(v *wrapperspb.BoolValue) {
	m.AutoExplainLogAnalyze = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogBuffers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogBuffers = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogTiming(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTiming = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogTriggers(v *wrapperspb.BoolValue) {
	m.AutoExplainLogTriggers = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogVerbose(v *wrapperspb.BoolValue) {
	m.AutoExplainLogVerbose = v
}

func (m *PostgresqlConfig11) SetAutoExplainLogNestedStatements(v *wrapperspb.BoolValue) {
	m.AutoExplainLogNestedStatements = v
}

func (m *PostgresqlConfig11) SetAutoExplainSampleRate(v *wrapperspb.DoubleValue) {
	m.AutoExplainSampleRate = v
}

func (m *PostgresqlConfig11) SetPgHintPlanEnableHint(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHint = v
}

func (m *PostgresqlConfig11) SetPgHintPlanEnableHintTable(v *wrapperspb.BoolValue) {
	m.PgHintPlanEnableHintTable = v
}

func (m *PostgresqlConfig11) SetPgHintPlanDebugPrint(v PostgresqlConfig11_PgHintPlanDebugPrint) {
	m.PgHintPlanDebugPrint = v
}

func (m *PostgresqlConfig11) SetPgHintPlanMessageLevel(v PostgresqlConfig11_LogLevel) {
	m.PgHintPlanMessageLevel = v
}

func (m *PostgresqlConfig11) SetPgQualstatsEnabled(v *wrapperspb.BoolValue) {
	m.PgQualstatsEnabled = v
}

func (m *PostgresqlConfig11) SetPgQualstatsTrackConstants(v *wrapperspb.BoolValue) {
	m.PgQualstatsTrackConstants = v
}

func (m *PostgresqlConfig11) SetPgQualstatsMax(v *wrapperspb.Int64Value) {
	m.PgQualstatsMax = v
}

func (m *PostgresqlConfig11) SetPgQualstatsResolveOids(v *wrapperspb.BoolValue) {
	m.PgQualstatsResolveOids = v
}

func (m *PostgresqlConfig11) SetPgQualstatsSampleRate(v *wrapperspb.DoubleValue) {
	m.PgQualstatsSampleRate = v
}

func (m *PostgresqlConfig11) SetMaxStackDepth(v *wrapperspb.Int64Value) {
	m.MaxStackDepth = v
}

func (m *PostgresqlConfig11) SetGeqo(v *wrapperspb.BoolValue) {
	m.Geqo = v
}

func (m *PostgresqlConfig11) SetGeqoThreshold(v *wrapperspb.Int64Value) {
	m.GeqoThreshold = v
}

func (m *PostgresqlConfig11) SetGeqoEffort(v *wrapperspb.Int64Value) {
	m.GeqoEffort = v
}

func (m *PostgresqlConfig11) SetGeqoSeed(v *wrapperspb.DoubleValue) {
	m.GeqoSeed = v
}

func (m *PostgresqlConfigSet11) SetEffectiveConfig(v *PostgresqlConfig11) {
	m.EffectiveConfig = v
}

func (m *PostgresqlConfigSet11) SetUserConfig(v *PostgresqlConfig11) {
	m.UserConfig = v
}

func (m *PostgresqlConfigSet11) SetDefaultConfig(v *PostgresqlConfig11) {
	m.DefaultConfig = v
}
