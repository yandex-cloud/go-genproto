// Code generated by protoc-gen-goext. DO NOT EDIT.

package datasphere

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *JobParameters) SetInputFiles(v []*File) {
	m.InputFiles = v
}

func (m *JobParameters) SetOutputFiles(v []*FileDesc) {
	m.OutputFiles = v
}

func (m *JobParameters) SetS3MountIds(v []string) {
	m.S3MountIds = v
}

func (m *JobParameters) SetDatasetIds(v []string) {
	m.DatasetIds = v
}

func (m *JobParameters) SetCmd(v string) {
	m.Cmd = v
}

func (m *JobParameters) SetEnv(v *Environment) {
	m.Env = v
}

func (m *JobParameters) SetAttachProjectDisk(v bool) {
	m.AttachProjectDisk = v
}

func (m *JobParameters) SetCloudInstanceType(v *CloudInstanceType) {
	m.CloudInstanceType = v
}

func (m *CloudInstanceType) SetName(v string) {
	m.Name = v
}

func (m *File) SetDesc(v *FileDesc) {
	m.Desc = v
}

func (m *File) SetSha256(v string) {
	m.Sha256 = v
}

func (m *File) SetSizeBytes(v int64) {
	m.SizeBytes = v
}

func (m *StorageFile) SetFile(v *File) {
	m.File = v
}

func (m *StorageFile) SetUrl(v string) {
	m.Url = v
}

func (m *FileDesc) SetPath(v string) {
	m.Path = v
}

func (m *FileDesc) SetVar(v string) {
	m.Var = v
}

type Environment_DockerImage = isEnvironment_DockerImage

func (m *Environment) SetDockerImage(v Environment_DockerImage) {
	m.DockerImage = v
}

func (m *Environment) SetVars(v map[string]string) {
	m.Vars = v
}

func (m *Environment) SetDockerImageResourceId(v string) {
	m.DockerImage = &Environment_DockerImageResourceId{
		DockerImageResourceId: v,
	}
}

func (m *Environment) SetDockerImageSpec(v *DockerImageSpec) {
	m.DockerImage = &Environment_DockerImageSpec{
		DockerImageSpec: v,
	}
}

func (m *Environment) SetPythonEnv(v *PythonEnv) {
	m.PythonEnv = v
}

type DockerImageSpec_Password = isDockerImageSpec_Password

func (m *DockerImageSpec) SetPassword(v DockerImageSpec_Password) {
	m.Password = v
}

func (m *DockerImageSpec) SetImageUrl(v string) {
	m.ImageUrl = v
}

func (m *DockerImageSpec) SetUsername(v string) {
	m.Username = v
}

func (m *DockerImageSpec) SetPasswordPlainText(v string) {
	m.Password = &DockerImageSpec_PasswordPlainText{
		PasswordPlainText: v,
	}
}

func (m *DockerImageSpec) SetPasswordDsSecretName(v string) {
	m.Password = &DockerImageSpec_PasswordDsSecretName{
		PasswordDsSecretName: v,
	}
}

func (m *PythonEnv) SetCondaYaml(v string) {
	m.CondaYaml = v
}

func (m *PythonEnv) SetLocalModules(v []*File) {
	m.LocalModules = v
}

func (m *Job) SetId(v string) {
	m.Id = v
}

func (m *Job) SetName(v string) {
	m.Name = v
}

func (m *Job) SetDesc(v string) {
	m.Desc = v
}

func (m *Job) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Job) SetFinishedAt(v *timestamppb.Timestamp) {
	m.FinishedAt = v
}

func (m *Job) SetStatus(v JobStatus) {
	m.Status = v
}

func (m *Job) SetConfig(v string) {
	m.Config = v
}

func (m *Job) SetCreatedById(v string) {
	m.CreatedById = v
}

func (m *Job) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *Job) SetJobParameters(v *JobParameters) {
	m.JobParameters = v
}

func (m *Job) SetDataExpiresAt(v *timestamppb.Timestamp) {
	m.DataExpiresAt = v
}

func (m *Job) SetDataCleared(v bool) {
	m.DataCleared = v
}

func (m *Job) SetOutputFiles(v []*File) {
	m.OutputFiles = v
}

func (m *JobResult) SetReturnCode(v int64) {
	m.ReturnCode = v
}
