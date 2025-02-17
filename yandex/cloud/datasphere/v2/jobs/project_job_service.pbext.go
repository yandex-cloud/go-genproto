// Code generated by protoc-gen-goext. DO NOT EDIT.

package datasphere

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *CreateProjectJobRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *CreateProjectJobRequest) SetJobParameters(v *JobParameters) {
	m.JobParameters = v
}

func (m *CreateProjectJobRequest) SetConfig(v string) {
	m.Config = v
}

func (m *CreateProjectJobRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateProjectJobRequest) SetDesc(v string) {
	m.Desc = v
}

func (m *CreateProjectJobRequest) SetDataTtl(v *durationpb.Duration) {
	m.DataTtl = v
}

func (m *CreateProjectJobMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *CreateProjectJobMetadata) SetJobId(v string) {
	m.JobId = v
}

func (m *CreateProjectJobResponse) SetJobId(v string) {
	m.JobId = v
}

func (m *CreateProjectJobResponse) SetUploadFiles(v []*StorageFile) {
	m.UploadFiles = v
}

func (m *CloneProjectJobRequest) SetSourceJobId(v string) {
	m.SourceJobId = v
}

func (m *CloneProjectJobRequest) SetJobParametersOverrides(v *JobParameters) {
	m.JobParametersOverrides = v
}

func (m *CloneProjectJobRequest) SetName(v string) {
	m.Name = v
}

func (m *CloneProjectJobRequest) SetDesc(v string) {
	m.Desc = v
}

func (m *CloneProjectJobRequest) SetDataTtl(v *durationpb.Duration) {
	m.DataTtl = v
}

func (m *CloneProjectJobResponse) SetJobId(v string) {
	m.JobId = v
}

func (m *CloneProjectJobResponse) SetUploadFiles(v []*StorageFile) {
	m.UploadFiles = v
}

func (m *CloneProjectJobMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *CloneProjectJobMetadata) SetJobId(v string) {
	m.JobId = v
}

func (m *ExecuteProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *ExecuteProjectJobResponse) SetOutputFiles(v []*StorageFile) {
	m.OutputFiles = v
}

func (m *ExecuteProjectJobResponse) SetOutputFilesErrors(v []*FileUploadError) {
	m.OutputFilesErrors = v
}

func (m *ExecuteProjectJobResponse) SetOutputDatasets(v []*OutputDataset) {
	m.OutputDatasets = v
}

func (m *ExecuteProjectJobResponse) SetResult(v *JobResult) {
	m.Result = v
}

func (m *ExecuteProjectJobMetadata) SetJob(v *Job) {
	m.Job = v
}

func (m *ExecuteProjectJobMetadata) SetProgress(v *JobProgress) {
	m.Progress = v
}

func (m *ExecuteProjectJobMetadata) SetMetadata(v *JobMetadata) {
	m.Metadata = v
}

func (m *CancelProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *CancelProjectJobRequest) SetReason(v string) {
	m.Reason = v
}

func (m *CancelProjectJobRequest) SetGraceful(v bool) {
	m.Graceful = v
}

func (m *ReadProjectJobStdLogsRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *ReadProjectJobStdLogsRequest) SetOffset(v int64) {
	m.Offset = v
}

func (m *ReadProjectJobStdLogsResponse) SetLogs(v []*StdLog) {
	m.Logs = v
}

func (m *ReadProjectJobStdLogsResponse) SetOffset(v int64) {
	m.Offset = v
}

func (m *ReadProjectJobLogsRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *ReadProjectJobLogsRequest) SetOffset(v int64) {
	m.Offset = v
}

func (m *ReadProjectJobLogsResponse) SetLogs(v []*LogMessage) {
	m.Logs = v
}

func (m *ReadProjectJobLogsResponse) SetOffset(v int64) {
	m.Offset = v
}

func (m *DownloadProjectJobFilesRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *DownloadProjectJobFilesRequest) SetFiles(v []*File) {
	m.Files = v
}

func (m *DownloadProjectJobFilesResponse) SetDownloadFiles(v []*StorageFile) {
	m.DownloadFiles = v
}

func (m *ListProjectJobRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *ListProjectJobRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListProjectJobRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListProjectJobRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListProjectJobResponse) SetJobs(v []*Job) {
	m.Jobs = v
}

func (m *ListProjectJobResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *GetProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteProjectJobRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteProjectJobMetadata) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteProjectJobDataRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteProjectJobDataMetadata) SetJobId(v string) {
	m.JobId = v
}

func (m *DeleteAllProjectJobDataRequest) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *DeleteAllProjectJobDataMetadata) SetProjectId(v string) {
	m.ProjectId = v
}

func (m *SetProjectJobDataTtlRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *SetProjectJobDataTtlRequest) SetTtl(v *durationpb.Duration) {
	m.Ttl = v
}

func (m *StdLog) SetContent(v []byte) {
	m.Content = v
}

func (m *StdLog) SetType(v StdLog_Type) {
	m.Type = v
}

type LogMessage_Source = isLogMessage_Source

func (m *LogMessage) SetSource(v LogMessage_Source) {
	m.Source = v
}

func (m *LogMessage) SetContent(v []byte) {
	m.Content = v
}

func (m *LogMessage) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *LogMessage) SetStandardStream(v StandardStream) {
	m.Source = &LogMessage_StandardStream{
		StandardStream: v,
	}
}

func (m *LogMessage) SetFilePath(v string) {
	m.Source = &LogMessage_FilePath{
		FilePath: v,
	}
}
