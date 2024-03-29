// Code generated by protoc-gen-goext. DO NOT EDIT.

package datasphere

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *GetFolderBudgetRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *GetFolderBudgetResponse) SetUnitBalance(v *wrapperspb.Int64Value) {
	m.UnitBalance = v
}

func (m *GetFolderBudgetResponse) SetMaxUnitsPerHour(v *wrapperspb.Int64Value) {
	m.MaxUnitsPerHour = v
}

func (m *GetFolderBudgetResponse) SetMaxUnitsPerExecution(v *wrapperspb.Int64Value) {
	m.MaxUnitsPerExecution = v
}

func (m *SetFolderBudgetRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *SetFolderBudgetRequest) SetSetMask(v *fieldmaskpb.FieldMask) {
	m.SetMask = v
}

func (m *SetFolderBudgetRequest) SetUnitBalance(v *wrapperspb.Int64Value) {
	m.UnitBalance = v
}

func (m *SetFolderBudgetRequest) SetMaxUnitsPerHour(v *wrapperspb.Int64Value) {
	m.MaxUnitsPerHour = v
}

func (m *SetFolderBudgetRequest) SetMaxUnitsPerExecution(v *wrapperspb.Int64Value) {
	m.MaxUnitsPerExecution = v
}
