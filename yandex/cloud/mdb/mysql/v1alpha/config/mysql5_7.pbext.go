// Code generated by protoc-gen-goext. DO NOT EDIT.

package mysql

import (
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
)

func (m *MysqlConfig5_7) SetInnodbBufferPoolSize(v *wrappers.Int64Value) {
	m.InnodbBufferPoolSize = v
}

func (m *MysqlConfig5_7) SetMaxConnections(v *wrappers.Int64Value) {
	m.MaxConnections = v
}

func (m *MysqlConfig5_7) SetLongQueryTime(v *wrappers.DoubleValue) {
	m.LongQueryTime = v
}

func (m *MysqlConfigSet5_7) SetEffectiveConfig(v *MysqlConfig5_7) {
	m.EffectiveConfig = v
}

func (m *MysqlConfigSet5_7) SetUserConfig(v *MysqlConfig5_7) {
	m.UserConfig = v
}

func (m *MysqlConfigSet5_7) SetDefaultConfig(v *MysqlConfig5_7) {
	m.DefaultConfig = v
}
