// Code generated by protoc-gen-goext. DO NOT EDIT.

package containers

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Container) SetId(v string) {
	m.Id = v
}

func (m *Container) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Container) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Container) SetName(v string) {
	m.Name = v
}

func (m *Container) SetDescription(v string) {
	m.Description = v
}

func (m *Container) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Container) SetUrl(v string) {
	m.Url = v
}

func (m *Container) SetStatus(v Container_Status) {
	m.Status = v
}

func (m *Revision) SetId(v string) {
	m.Id = v
}

func (m *Revision) SetContainerId(v string) {
	m.ContainerId = v
}

func (m *Revision) SetDescription(v string) {
	m.Description = v
}

func (m *Revision) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Revision) SetImage(v *Image) {
	m.Image = v
}

func (m *Revision) SetResources(v *Resources) {
	m.Resources = v
}

func (m *Revision) SetExecutionTimeout(v *durationpb.Duration) {
	m.ExecutionTimeout = v
}

func (m *Revision) SetConcurrency(v int64) {
	m.Concurrency = v
}

func (m *Revision) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *Revision) SetStatus(v Revision_Status) {
	m.Status = v
}

func (m *Image) SetImageUrl(v string) {
	m.ImageUrl = v
}

func (m *Image) SetImageDigest(v string) {
	m.ImageDigest = v
}

func (m *Image) SetCommand(v *Command) {
	m.Command = v
}

func (m *Image) SetArgs(v *Args) {
	m.Args = v
}

func (m *Image) SetEnvironment(v map[string]string) {
	m.Environment = v
}

func (m *Image) SetWorkingDir(v string) {
	m.WorkingDir = v
}

func (m *Command) SetCommand(v []string) {
	m.Command = v
}

func (m *Args) SetArgs(v []string) {
	m.Args = v
}

func (m *Resources) SetMemory(v int64) {
	m.Memory = v
}

func (m *Resources) SetCores(v int64) {
	m.Cores = v
}

func (m *Resources) SetCoreFraction(v int64) {
	m.CoreFraction = v
}
