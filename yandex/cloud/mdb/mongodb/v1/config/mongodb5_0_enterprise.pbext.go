// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *MongodConfig5_0Enterprise) SetStorage(v *MongodConfig5_0Enterprise_Storage) {
	m.Storage = v
}

func (m *MongodConfig5_0Enterprise) SetOperationProfiling(v *MongodConfig5_0Enterprise_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongodConfig5_0Enterprise) SetNet(v *MongodConfig5_0Enterprise_Network) {
	m.Net = v
}

func (m *MongodConfig5_0Enterprise) SetSecurity(v *MongodConfig5_0Enterprise_Security) {
	m.Security = v
}

func (m *MongodConfig5_0Enterprise) SetAuditLog(v *MongodConfig5_0Enterprise_AuditLog) {
	m.AuditLog = v
}

func (m *MongodConfig5_0Enterprise) SetSetParameter(v *MongodConfig5_0Enterprise_SetParameter) {
	m.SetParameter = v
}

func (m *MongodConfig5_0Enterprise_Storage) SetWiredTiger(v *MongodConfig5_0Enterprise_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongodConfig5_0Enterprise_Storage) SetJournal(v *MongodConfig5_0Enterprise_Storage_Journal) {
	m.Journal = v
}

func (m *MongodConfig5_0Enterprise_Storage_WiredTiger) SetEngineConfig(v *MongodConfig5_0Enterprise_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongodConfig5_0Enterprise_Storage_WiredTiger) SetCollectionConfig(v *MongodConfig5_0Enterprise_Storage_WiredTiger_CollectionConfig) {
	m.CollectionConfig = v
}

func (m *MongodConfig5_0Enterprise_Storage_WiredTiger) SetIndexConfig(v *MongodConfig5_0Enterprise_Storage_WiredTiger_IndexConfig) {
	m.IndexConfig = v
}

func (m *MongodConfig5_0Enterprise_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongodConfig5_0Enterprise_Storage_WiredTiger_CollectionConfig) SetBlockCompressor(v MongodConfig5_0Enterprise_Storage_WiredTiger_CollectionConfig_Compressor) {
	m.BlockCompressor = v
}

func (m *MongodConfig5_0Enterprise_Storage_WiredTiger_IndexConfig) SetPrefixCompression(v *wrapperspb.BoolValue) {
	m.PrefixCompression = v
}

func (m *MongodConfig5_0Enterprise_Storage_Journal) SetCommitInterval(v *wrapperspb.Int64Value) {
	m.CommitInterval = v
}

func (m *MongodConfig5_0Enterprise_OperationProfiling) SetMode(v MongodConfig5_0Enterprise_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongodConfig5_0Enterprise_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongodConfig5_0Enterprise_OperationProfiling) SetSlowOpSampleRate(v *wrapperspb.DoubleValue) {
	m.SlowOpSampleRate = v
}

func (m *MongodConfig5_0Enterprise_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongodConfig5_0Enterprise_Network) SetCompression(v *MongodConfig5_0Enterprise_Network_Compression) {
	m.Compression = v
}

func (m *MongodConfig5_0Enterprise_Network_Compression) SetCompressors(v []MongodConfig5_0Enterprise_Network_Compression_Compressor) {
	m.Compressors = v
}

func (m *MongodConfig5_0Enterprise_Security) SetEnableEncryption(v *wrapperspb.BoolValue) {
	m.EnableEncryption = v
}

func (m *MongodConfig5_0Enterprise_Security) SetKmip(v *MongodConfig5_0Enterprise_Security_KMIP) {
	m.Kmip = v
}

func (m *MongodConfig5_0Enterprise_Security_KMIP) SetServerName(v string) {
	m.ServerName = v
}

func (m *MongodConfig5_0Enterprise_Security_KMIP) SetPort(v *wrapperspb.Int64Value) {
	m.Port = v
}

func (m *MongodConfig5_0Enterprise_Security_KMIP) SetServerCa(v string) {
	m.ServerCa = v
}

func (m *MongodConfig5_0Enterprise_Security_KMIP) SetClientCertificate(v string) {
	m.ClientCertificate = v
}

func (m *MongodConfig5_0Enterprise_Security_KMIP) SetKeyIdentifier(v string) {
	m.KeyIdentifier = v
}

func (m *MongodConfig5_0Enterprise_AuditLog) SetFilter(v string) {
	m.Filter = v
}

func (m *MongodConfig5_0Enterprise_AuditLog) SetRuntimeConfiguration(v *wrapperspb.BoolValue) {
	m.RuntimeConfiguration = v
}

func (m *MongodConfig5_0Enterprise_SetParameter) SetAuditAuthorizationSuccess(v *wrapperspb.BoolValue) {
	m.AuditAuthorizationSuccess = v
}

func (m *MongodConfig5_0Enterprise_SetParameter) SetEnableFlowControl(v *wrapperspb.BoolValue) {
	m.EnableFlowControl = v
}

func (m *MongodConfig5_0Enterprise_SetParameter) SetMinSnapshotHistoryWindowInSeconds(v *wrapperspb.Int64Value) {
	m.MinSnapshotHistoryWindowInSeconds = v
}

func (m *MongoCfgConfig5_0Enterprise) SetStorage(v *MongoCfgConfig5_0Enterprise_Storage) {
	m.Storage = v
}

func (m *MongoCfgConfig5_0Enterprise) SetOperationProfiling(v *MongoCfgConfig5_0Enterprise_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongoCfgConfig5_0Enterprise) SetNet(v *MongoCfgConfig5_0Enterprise_Network) {
	m.Net = v
}

func (m *MongoCfgConfig5_0Enterprise_Storage) SetWiredTiger(v *MongoCfgConfig5_0Enterprise_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongoCfgConfig5_0Enterprise_Storage_WiredTiger) SetEngineConfig(v *MongoCfgConfig5_0Enterprise_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongoCfgConfig5_0Enterprise_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongoCfgConfig5_0Enterprise_OperationProfiling) SetMode(v MongoCfgConfig5_0Enterprise_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongoCfgConfig5_0Enterprise_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongoCfgConfig5_0Enterprise_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongosConfig5_0Enterprise) SetNet(v *MongosConfig5_0Enterprise_Network) {
	m.Net = v
}

func (m *MongosConfig5_0Enterprise_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongosConfig5_0Enterprise_Network) SetCompression(v *MongosConfig5_0Enterprise_Network_Compression) {
	m.Compression = v
}

func (m *MongosConfig5_0Enterprise_Network_Compression) SetCompressors(v []MongosConfig5_0Enterprise_Network_Compression_Compressor) {
	m.Compressors = v
}

func (m *MongodConfigSet5_0Enterprise) SetEffectiveConfig(v *MongodConfig5_0Enterprise) {
	m.EffectiveConfig = v
}

func (m *MongodConfigSet5_0Enterprise) SetUserConfig(v *MongodConfig5_0Enterprise) {
	m.UserConfig = v
}

func (m *MongodConfigSet5_0Enterprise) SetDefaultConfig(v *MongodConfig5_0Enterprise) {
	m.DefaultConfig = v
}

func (m *MongoCfgConfigSet5_0Enterprise) SetEffectiveConfig(v *MongoCfgConfig5_0Enterprise) {
	m.EffectiveConfig = v
}

func (m *MongoCfgConfigSet5_0Enterprise) SetUserConfig(v *MongoCfgConfig5_0Enterprise) {
	m.UserConfig = v
}

func (m *MongoCfgConfigSet5_0Enterprise) SetDefaultConfig(v *MongoCfgConfig5_0Enterprise) {
	m.DefaultConfig = v
}

func (m *MongosConfigSet5_0Enterprise) SetEffectiveConfig(v *MongosConfig5_0Enterprise) {
	m.EffectiveConfig = v
}

func (m *MongosConfigSet5_0Enterprise) SetUserConfig(v *MongosConfig5_0Enterprise) {
	m.UserConfig = v
}

func (m *MongosConfigSet5_0Enterprise) SetDefaultConfig(v *MongosConfig5_0Enterprise) {
	m.DefaultConfig = v
}
