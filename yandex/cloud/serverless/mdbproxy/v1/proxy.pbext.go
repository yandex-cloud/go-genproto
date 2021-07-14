// Code generated by protoc-gen-goext. DO NOT EDIT.

package proxy

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
)

func (m *Proxy) SetId(v string) {
	m.Id = v
}

func (m *Proxy) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Proxy) SetCreatedAt(v *timestamp.Timestamp) {
	m.CreatedAt = v
}

func (m *Proxy) SetName(v string) {
	m.Name = v
}

func (m *Proxy) SetDescription(v string) {
	m.Description = v
}

func (m *Proxy) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Proxy) SetTarget(v *Target) {
	m.Target = v
}

type Target_Mdb = isTarget_Mdb

func (m *Target) SetMdb(v Target_Mdb) {
	m.Mdb = v
}

func (m *Target) SetPostgresql(v *Target_PostgreSQL) {
	m.Mdb = &Target_Postgresql{
		Postgresql: v,
	}
}

func (m *Target_PostgreSQL) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *Target_PostgreSQL) SetUser(v string) {
	m.User = v
}

func (m *Target_PostgreSQL) SetPassword(v string) {
	m.Password = v
}

func (m *Target_PostgreSQL) SetDb(v string) {
	m.Db = v
}

func (m *Target_PostgreSQL) SetEndpoint(v string) {
	m.Endpoint = v
}
