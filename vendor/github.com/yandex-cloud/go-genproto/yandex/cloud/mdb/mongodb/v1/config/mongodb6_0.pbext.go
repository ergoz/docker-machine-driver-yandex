// Code generated by protoc-gen-goext. DO NOT EDIT.

package mongodb

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *MongodConfig6_0) SetStorage(v *MongodConfig6_0_Storage) {
	m.Storage = v
}

func (m *MongodConfig6_0) SetOperationProfiling(v *MongodConfig6_0_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongodConfig6_0) SetNet(v *MongodConfig6_0_Network) {
	m.Net = v
}

func (m *MongodConfig6_0_Storage) SetWiredTiger(v *MongodConfig6_0_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongodConfig6_0_Storage) SetJournal(v *MongodConfig6_0_Storage_Journal) {
	m.Journal = v
}

func (m *MongodConfig6_0_Storage_WiredTiger) SetEngineConfig(v *MongodConfig6_0_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongodConfig6_0_Storage_WiredTiger) SetCollectionConfig(v *MongodConfig6_0_Storage_WiredTiger_CollectionConfig) {
	m.CollectionConfig = v
}

func (m *MongodConfig6_0_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongodConfig6_0_Storage_WiredTiger_CollectionConfig) SetBlockCompressor(v MongodConfig6_0_Storage_WiredTiger_CollectionConfig_Compressor) {
	m.BlockCompressor = v
}

func (m *MongodConfig6_0_Storage_Journal) SetCommitInterval(v *wrapperspb.Int64Value) {
	m.CommitInterval = v
}

func (m *MongodConfig6_0_OperationProfiling) SetMode(v MongodConfig6_0_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongodConfig6_0_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongodConfig6_0_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongoCfgConfig6_0) SetStorage(v *MongoCfgConfig6_0_Storage) {
	m.Storage = v
}

func (m *MongoCfgConfig6_0) SetOperationProfiling(v *MongoCfgConfig6_0_OperationProfiling) {
	m.OperationProfiling = v
}

func (m *MongoCfgConfig6_0) SetNet(v *MongoCfgConfig6_0_Network) {
	m.Net = v
}

func (m *MongoCfgConfig6_0_Storage) SetWiredTiger(v *MongoCfgConfig6_0_Storage_WiredTiger) {
	m.WiredTiger = v
}

func (m *MongoCfgConfig6_0_Storage_WiredTiger) SetEngineConfig(v *MongoCfgConfig6_0_Storage_WiredTiger_EngineConfig) {
	m.EngineConfig = v
}

func (m *MongoCfgConfig6_0_Storage_WiredTiger_EngineConfig) SetCacheSizeGb(v *wrapperspb.DoubleValue) {
	m.CacheSizeGb = v
}

func (m *MongoCfgConfig6_0_OperationProfiling) SetMode(v MongoCfgConfig6_0_OperationProfiling_Mode) {
	m.Mode = v
}

func (m *MongoCfgConfig6_0_OperationProfiling) SetSlowOpThreshold(v *wrapperspb.Int64Value) {
	m.SlowOpThreshold = v
}

func (m *MongoCfgConfig6_0_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongosConfig6_0) SetNet(v *MongosConfig6_0_Network) {
	m.Net = v
}

func (m *MongosConfig6_0_Network) SetMaxIncomingConnections(v *wrapperspb.Int64Value) {
	m.MaxIncomingConnections = v
}

func (m *MongodConfigSet6_0) SetEffectiveConfig(v *MongodConfig6_0) {
	m.EffectiveConfig = v
}

func (m *MongodConfigSet6_0) SetUserConfig(v *MongodConfig6_0) {
	m.UserConfig = v
}

func (m *MongodConfigSet6_0) SetDefaultConfig(v *MongodConfig6_0) {
	m.DefaultConfig = v
}

func (m *MongoCfgConfigSet6_0) SetEffectiveConfig(v *MongoCfgConfig6_0) {
	m.EffectiveConfig = v
}

func (m *MongoCfgConfigSet6_0) SetUserConfig(v *MongoCfgConfig6_0) {
	m.UserConfig = v
}

func (m *MongoCfgConfigSet6_0) SetDefaultConfig(v *MongoCfgConfig6_0) {
	m.DefaultConfig = v
}

func (m *MongosConfigSet6_0) SetEffectiveConfig(v *MongosConfig6_0) {
	m.EffectiveConfig = v
}

func (m *MongosConfigSet6_0) SetUserConfig(v *MongosConfig6_0) {
	m.UserConfig = v
}

func (m *MongosConfigSet6_0) SetDefaultConfig(v *MongosConfig6_0) {
	m.DefaultConfig = v
}
