// Code generated by protoc-gen-goext. DO NOT EDIT.

package dataproc_manager

func (m *ListJobsRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *ListJobsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListJobsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListJobsRequest) SetFilter(v string) {
	m.Filter = v
}

func (m *ListJobsResponse) SetJobs(v []*Job) {
	m.Jobs = v
}

func (m *ListJobsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateJobStatusRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateJobStatusRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *UpdateJobStatusRequest) SetStatus(v Job_Status) {
	m.Status = v
}

func (m *UpdateJobStatusRequest) SetApplicationInfo(v *ApplicationInfo) {
	m.ApplicationInfo = v
}

func (m *ListSupportJobsResponse) SetJobs(v []*SupportJob) {
	m.Jobs = v
}

func (m *ListSupportJobsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *UpdateSupportJobStatusRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *UpdateSupportJobStatusRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *UpdateSupportJobStatusRequest) SetStatus(v SupportJob_Status) {
	m.Status = v
}

func (m *SaveSupportJobLogRequest) SetClusterId(v string) {
	m.ClusterId = v
}

func (m *SaveSupportJobLogRequest) SetJobId(v string) {
	m.JobId = v
}

func (m *SaveSupportJobLogRequest) SetOutput(v string) {
	m.Output = v
}
