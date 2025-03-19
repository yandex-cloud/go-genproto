// Code generated by protoc-gen-goext. DO NOT EDIT.

package spqr

import (
	timeofday "google.golang.org/genproto/googleapis/type/timeofday"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *Cluster) SetId(v string) {
	m.Id = v
}

func (m *Cluster) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Cluster) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Cluster) SetName(v string) {
	m.Name = v
}

func (m *Cluster) SetDescription(v string) {
	m.Description = v
}

func (m *Cluster) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Cluster) SetEnvironment(v Cluster_Environment) {
	m.Environment = v
}

func (m *Cluster) SetMonitoring(v []*Monitoring) {
	m.Monitoring = v
}

func (m *Cluster) SetConfig(v *ClusterConfig) {
	m.Config = v
}

func (m *Cluster) SetNetworkId(v string) {
	m.NetworkId = v
}

func (m *Cluster) SetHealth(v Cluster_Health) {
	m.Health = v
}

func (m *Cluster) SetStatus(v Cluster_Status) {
	m.Status = v
}

func (m *Cluster) SetMaintenanceWindow(v *MaintenanceWindow) {
	m.MaintenanceWindow = v
}

func (m *Cluster) SetPlannedOperation(v *MaintenanceOperation) {
	m.PlannedOperation = v
}

func (m *Cluster) SetSecurityGroupIds(v []string) {
	m.SecurityGroupIds = v
}

func (m *Cluster) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Monitoring) SetName(v string) {
	m.Name = v
}

func (m *Monitoring) SetDescription(v string) {
	m.Description = v
}

func (m *Monitoring) SetLink(v string) {
	m.Link = v
}

func (m *ClusterConfig) SetSpqrConfig(v *SPQRConfig) {
	m.SpqrConfig = v
}

func (m *ClusterConfig) SetBackupWindowStart(v *timeofday.TimeOfDay) {
	m.BackupWindowStart = v
}

func (m *ClusterConfig) SetBackupRetainPeriodDays(v *wrapperspb.Int64Value) {
	m.BackupRetainPeriodDays = v
}

func (m *ClusterConfig) SetAccess(v *Access) {
	m.Access = v
}

func (m *ClusterConfig) SetSoxAudit(v *wrapperspb.BoolValue) {
	m.SoxAudit = v
}

func (m *Access) SetDataLens(v bool) {
	m.DataLens = v
}

func (m *Access) SetWebSql(v bool) {
	m.WebSql = v
}

func (m *Access) SetDataTransfer(v bool) {
	m.DataTransfer = v
}

func (m *Access) SetServerless(v bool) {
	m.Serverless = v
}
