// Code generated by protoc-gen-goext. DO NOT EDIT.

package endpoint

func (m *PostgresObjectTransferSettings) SetSequence(v ObjectTransferStage) {
	m.Sequence = v
}

func (m *PostgresObjectTransferSettings) SetSequenceOwnedBy(v ObjectTransferStage) {
	m.SequenceOwnedBy = v
}

func (m *PostgresObjectTransferSettings) SetTable(v ObjectTransferStage) {
	m.Table = v
}

func (m *PostgresObjectTransferSettings) SetPrimaryKey(v ObjectTransferStage) {
	m.PrimaryKey = v
}

func (m *PostgresObjectTransferSettings) SetFkConstraint(v ObjectTransferStage) {
	m.FkConstraint = v
}

func (m *PostgresObjectTransferSettings) SetDefaultValues(v ObjectTransferStage) {
	m.DefaultValues = v
}

func (m *PostgresObjectTransferSettings) SetConstraint(v ObjectTransferStage) {
	m.Constraint = v
}

func (m *PostgresObjectTransferSettings) SetIndex(v ObjectTransferStage) {
	m.Index = v
}

func (m *PostgresObjectTransferSettings) SetView(v ObjectTransferStage) {
	m.View = v
}

func (m *PostgresObjectTransferSettings) SetFunction(v ObjectTransferStage) {
	m.Function = v
}

func (m *PostgresObjectTransferSettings) SetTrigger(v ObjectTransferStage) {
	m.Trigger = v
}

func (m *PostgresObjectTransferSettings) SetType(v ObjectTransferStage) {
	m.Type = v
}

func (m *PostgresObjectTransferSettings) SetRule(v ObjectTransferStage) {
	m.Rule = v
}

func (m *PostgresObjectTransferSettings) SetCollation(v ObjectTransferStage) {
	m.Collation = v
}

func (m *PostgresObjectTransferSettings) SetPolicy(v ObjectTransferStage) {
	m.Policy = v
}

func (m *PostgresObjectTransferSettings) SetCast(v ObjectTransferStage) {
	m.Cast = v
}

func (m *OnPremisePostgres) SetHosts(v []string) {
	m.Hosts = v
}

func (m *OnPremisePostgres) SetPort(v int64) {
	m.Port = v
}

func (m *OnPremisePostgres) SetTlsMode(v *TLSMode) {
	m.TlsMode = v
}

func (m *OnPremisePostgres) SetSubnetId(v string) {
	m.SubnetId = v
}

type PostgresConnection_Connection = isPostgresConnection_Connection

func (m *PostgresConnection) SetConnection(v PostgresConnection_Connection) {
	m.Connection = v
}

func (m *PostgresConnection) SetMdbClusterId(v string) {
	m.Connection = &PostgresConnection_MdbClusterId{
		MdbClusterId: v,
	}
}

func (m *PostgresConnection) SetOnPremise(v *OnPremisePostgres) {
	m.Connection = &PostgresConnection_OnPremise{
		OnPremise: v,
	}
}

func (m *PostgresSource) SetConnection(v *PostgresConnection) {
	m.Connection = v
}

func (m *PostgresSource) SetDatabase(v string) {
	m.Database = v
}

func (m *PostgresSource) SetUser(v string) {
	m.User = v
}

func (m *PostgresSource) SetPassword(v *Secret) {
	m.Password = v
}

func (m *PostgresSource) SetIncludeTables(v []string) {
	m.IncludeTables = v
}

func (m *PostgresSource) SetExcludeTables(v []string) {
	m.ExcludeTables = v
}

func (m *PostgresSource) SetSlotByteLagLimit(v int64) {
	m.SlotByteLagLimit = v
}

func (m *PostgresSource) SetServiceSchema(v string) {
	m.ServiceSchema = v
}

func (m *PostgresSource) SetObjectTransferSettings(v *PostgresObjectTransferSettings) {
	m.ObjectTransferSettings = v
}

func (m *PostgresTarget) SetConnection(v *PostgresConnection) {
	m.Connection = v
}

func (m *PostgresTarget) SetDatabase(v string) {
	m.Database = v
}

func (m *PostgresTarget) SetUser(v string) {
	m.User = v
}

func (m *PostgresTarget) SetPassword(v *Secret) {
	m.Password = v
}