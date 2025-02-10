// Code generated by protoc-gen-goext. DO NOT EDIT.

package fomo

import (
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
)

func (m *DescribeDatasetRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *DescribeDatasetResponse) SetDataset(v *DatasetInfo) {
	m.Dataset = v
}

func (m *ValidateDatasetRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *ValidateDatasetMetadata) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *ValidateDatasetMetadata) SetValidRows(v int64) {
	m.ValidRows = v
}

func (m *ValidateDatasetMetadata) SetProcessedRows(v int64) {
	m.ProcessedRows = v
}

func (m *ValidateDatasetMetadata) SetTotalRows(v int64) {
	m.TotalRows = v
}

func (m *ValidateDatasetResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *ValidateDatasetResponse) SetIsValid(v bool) {
	m.IsValid = v
}

func (m *ValidateDatasetResponse) SetErrors(v []*ValidationError) {
	m.Errors = v
}

func (m *DeleteDatasetRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *CreateDatasetRequest) SetName(v string) {
	m.Name = v
}

func (m *CreateDatasetRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *CreateDatasetRequest) SetDescription(v string) {
	m.Description = v
}

func (m *CreateDatasetRequest) SetMetadata(v string) {
	m.Metadata = v
}

func (m *CreateDatasetRequest) SetTaskType(v string) {
	m.TaskType = v
}

func (m *CreateDatasetRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *CreateDatasetRequest) SetUploadFormat(v string) {
	m.UploadFormat = v
}

func (m *CreateDatasetRequest) SetAllowDataLog(v bool) {
	m.AllowDataLog = v
}

func (m *CreateDatasetResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *CreateDatasetResponse) SetDataset(v *DatasetInfo) {
	m.Dataset = v
}

func (m *UpdateDatasetRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *UpdateDatasetRequest) SetUpdateMask(v *fieldmaskpb.FieldMask) {
	m.UpdateMask = v
}

func (m *UpdateDatasetRequest) SetName(v string) {
	m.Name = v
}

func (m *UpdateDatasetRequest) SetDescription(v string) {
	m.Description = v
}

func (m *UpdateDatasetRequest) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *UpdateDatasetResponse) SetDataset(v *DatasetInfo) {
	m.Dataset = v
}

func (m *GetUploadDraftUrlRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *GetUploadDraftUrlRequest) SetSizeBytes(v int64) {
	m.SizeBytes = v
}

func (m *GetUploadDraftUrlResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *GetUploadDraftUrlResponse) SetUploadUrl(v string) {
	m.UploadUrl = v
}

func (m *StartMultipartUploadDraftRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *StartMultipartUploadDraftRequest) SetSizeBytes(v int64) {
	m.SizeBytes = v
}

func (m *StartMultipartUploadDraftRequest) SetParts(v int64) {
	m.Parts = v
}

func (m *StartMultipartUploadDraftResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *StartMultipartUploadDraftResponse) SetMultipartUploadUrls(v []string) {
	m.MultipartUploadUrls = v
}

func (m *UploadedPartInfo) SetPartNum(v int64) {
	m.PartNum = v
}

func (m *UploadedPartInfo) SetEtag(v string) {
	m.Etag = v
}

func (m *FinishMultipartUploadDraftRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *FinishMultipartUploadDraftRequest) SetUploadedParts(v []*UploadedPartInfo) {
	m.UploadedParts = v
}

func (m *FinishMultipartUploadDraftResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *ListDatasetsRequest) SetFolderId(v string) {
	m.FolderId = v
}

func (m *ListDatasetsRequest) SetStatus(v DatasetInfo_Status) {
	m.Status = v
}

func (m *ListDatasetsRequest) SetDatasetNamePattern(v string) {
	m.DatasetNamePattern = v
}

func (m *ListDatasetsRequest) SetTaskTypeFilter(v string) {
	m.TaskTypeFilter = v
}

func (m *ListDatasetsRequest) SetPageSize(v int64) {
	m.PageSize = v
}

func (m *ListDatasetsRequest) SetPageToken(v string) {
	m.PageToken = v
}

func (m *ListDatasetsResponse) SetDatasets(v []*DatasetInfo) {
	m.Datasets = v
}

func (m *ListDatasetsResponse) SetNextPageToken(v string) {
	m.NextPageToken = v
}

func (m *ListUploadFormatsRequest) SetTaskType(v string) {
	m.TaskType = v
}

func (m *ListUploadFormatsResponse) SetFormats(v []string) {
	m.Formats = v
}

func (m *ListUploadSchemasRequest) SetTaskType(v string) {
	m.TaskType = v
}

func (m *ListUploadSchemasResponse) SetSchemas(v []*DatasetUploadSchema) {
	m.Schemas = v
}

func (m *ListTypesResponse) SetTypes(v []string) {
	m.Types = v
}

func (m *GetDatasetPreviewRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *GetDatasetPreviewResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *GetDatasetPreviewResponse) SetPreviewLines(v []string) {
	m.PreviewLines = v
}

func (m *GetDownloadUrlsRequest) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *GetDownloadUrlsResponse) SetDatasetId(v string) {
	m.DatasetId = v
}

func (m *GetDownloadUrlsResponse) SetDownloadUrls(v []*DatasetFileDownloadUrl) {
	m.DownloadUrls = v
}

func (m *ListOperationsIdsRequest) SetDatasetId(v []string) {
	m.DatasetId = v
}

func (m *ListOperationsIdsResponse) SetDatasetIdToOperationId(v map[string]string) {
	m.DatasetIdToOperationId = v
}
