// Code generated by protoc-gen-goext. DO NOT EDIT.

package spqr

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

func (m *User) SetGrants(v []string) {
	m.Grants = v
}

func (m *Permission) SetDatabaseName(v string) {
	m.DatabaseName = v
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

func (m *UserSpec) SetGrants(v []string) {
	m.Grants = v
}

func (m *UserSettings) SetConnectionLimit(v *wrapperspb.Int64Value) {
	m.ConnectionLimit = v
}

func (m *UserSettings) SetConnectionRetries(v *wrapperspb.Int64Value) {
	m.ConnectionRetries = v
}
