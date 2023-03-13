// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.17.3
// source: yandex/cloud/datasphere/v1/project_data_service.proto

package datasphere

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type FileMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Project resource associated with the file.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// File path.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// File size in bytes.
	SizeBytes int64 `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
}

func (x *FileMetadata) Reset() {
	*x = FileMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileMetadata) ProtoMessage() {}

func (x *FileMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileMetadata.ProtoReflect.Descriptor instead.
func (*FileMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescGZIP(), []int{0}
}

func (x *FileMetadata) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *FileMetadata) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *FileMetadata) GetSizeBytes() int64 {
	if x != nil {
		return x.SizeBytes
	}
	return 0
}

type UploadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*UploadFileRequest_Metadata
	//	*UploadFileRequest_Chunk
	Message isUploadFileRequest_Message `protobuf_oneof:"message"`
}

func (x *UploadFileRequest) Reset() {
	*x = UploadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRequest) ProtoMessage() {}

func (x *UploadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRequest.ProtoReflect.Descriptor instead.
func (*UploadFileRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescGZIP(), []int{1}
}

func (m *UploadFileRequest) GetMessage() isUploadFileRequest_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *UploadFileRequest) GetMetadata() *FileMetadata {
	if x, ok := x.GetMessage().(*UploadFileRequest_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *UploadFileRequest) GetChunk() []byte {
	if x, ok := x.GetMessage().(*UploadFileRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isUploadFileRequest_Message interface {
	isUploadFileRequest_Message()
}

type UploadFileRequest_Metadata struct {
	// Metadata of the file to upload.
	Metadata *FileMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type UploadFileRequest_Chunk struct {
	// Byte chunk of the file to upload.
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*UploadFileRequest_Metadata) isUploadFileRequest_Message() {}

func (*UploadFileRequest_Chunk) isUploadFileRequest_Message() {}

type UploadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Metadata of the uploaded file.
	Metadata *FileMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *UploadFileResponse) Reset() {
	*x = UploadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UploadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileResponse) ProtoMessage() {}

func (x *UploadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileResponse.ProtoReflect.Descriptor instead.
func (*UploadFileResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileResponse) GetMetadata() *FileMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type DownloadFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the Project resource to download the file from.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Path of the file to download.
	FilePath string `protobuf:"bytes,2,opt,name=file_path,json=filePath,proto3" json:"file_path,omitempty"`
}

func (x *DownloadFileRequest) Reset() {
	*x = DownloadFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileRequest) ProtoMessage() {}

func (x *DownloadFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileRequest.ProtoReflect.Descriptor instead.
func (*DownloadFileRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescGZIP(), []int{3}
}

func (x *DownloadFileRequest) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *DownloadFileRequest) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type DownloadFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Message:
	//
	//	*DownloadFileResponse_Metadata
	//	*DownloadFileResponse_Chunk
	Message isDownloadFileResponse_Message `protobuf_oneof:"message"`
}

func (x *DownloadFileResponse) Reset() {
	*x = DownloadFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DownloadFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DownloadFileResponse) ProtoMessage() {}

func (x *DownloadFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DownloadFileResponse.ProtoReflect.Descriptor instead.
func (*DownloadFileResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescGZIP(), []int{4}
}

func (m *DownloadFileResponse) GetMessage() isDownloadFileResponse_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (x *DownloadFileResponse) GetMetadata() *FileMetadata {
	if x, ok := x.GetMessage().(*DownloadFileResponse_Metadata); ok {
		return x.Metadata
	}
	return nil
}

func (x *DownloadFileResponse) GetChunk() []byte {
	if x, ok := x.GetMessage().(*DownloadFileResponse_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isDownloadFileResponse_Message interface {
	isDownloadFileResponse_Message()
}

type DownloadFileResponse_Metadata struct {
	// Metadata of the downloaded file.
	Metadata *FileMetadata `protobuf:"bytes,1,opt,name=metadata,proto3,oneof"`
}

type DownloadFileResponse_Chunk struct {
	// Byte chunk of the downloaded file.
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*DownloadFileResponse_Metadata) isDownloadFileResponse_Message() {}

func (*DownloadFileResponse_Chunk) isDownloadFileResponse_Message() {}

var File_yandex_cloud_datasphere_v1_project_data_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x6f, 0x0a, 0x0c, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x05,
	0x3c, 0x3d, 0x32, 0x30, 0x30, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x69, 0x7a, 0x65, 0x5f, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x7a, 0x65, 0x42, 0x79,
	0x74, 0x65, 0x73, 0x22, 0x7e, 0x0a, 0x11, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70,
	0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48,
	0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x66, 0x0a, 0x13, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xe8, 0xc7, 0x31, 0x01,
	0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x30, 0x30, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x81, 0x01, 0x0a, 0x14, 0x44, 0x6f, 0x77, 0x6e,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e,
	0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x42, 0x09, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xf8, 0x01, 0x0a, 0x12,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x6d, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x12, 0x73, 0x0a, 0x0c, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x6b, 0x0a, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68,
	0x65, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescData = file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDesc
)

func file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescData)
	})
	return file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDescData
}

var file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_datasphere_v1_project_data_service_proto_goTypes = []interface{}{
	(*FileMetadata)(nil),         // 0: yandex.cloud.datasphere.v1.FileMetadata
	(*UploadFileRequest)(nil),    // 1: yandex.cloud.datasphere.v1.UploadFileRequest
	(*UploadFileResponse)(nil),   // 2: yandex.cloud.datasphere.v1.UploadFileResponse
	(*DownloadFileRequest)(nil),  // 3: yandex.cloud.datasphere.v1.DownloadFileRequest
	(*DownloadFileResponse)(nil), // 4: yandex.cloud.datasphere.v1.DownloadFileResponse
}
var file_yandex_cloud_datasphere_v1_project_data_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.datasphere.v1.UploadFileRequest.metadata:type_name -> yandex.cloud.datasphere.v1.FileMetadata
	0, // 1: yandex.cloud.datasphere.v1.UploadFileResponse.metadata:type_name -> yandex.cloud.datasphere.v1.FileMetadata
	0, // 2: yandex.cloud.datasphere.v1.DownloadFileResponse.metadata:type_name -> yandex.cloud.datasphere.v1.FileMetadata
	1, // 3: yandex.cloud.datasphere.v1.ProjectDataService.UploadFile:input_type -> yandex.cloud.datasphere.v1.UploadFileRequest
	3, // 4: yandex.cloud.datasphere.v1.ProjectDataService.DownloadFile:input_type -> yandex.cloud.datasphere.v1.DownloadFileRequest
	2, // 5: yandex.cloud.datasphere.v1.ProjectDataService.UploadFile:output_type -> yandex.cloud.datasphere.v1.UploadFileResponse
	4, // 6: yandex.cloud.datasphere.v1.ProjectDataService.DownloadFile:output_type -> yandex.cloud.datasphere.v1.DownloadFileResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v1_project_data_service_proto_init() }
func file_yandex_cloud_datasphere_v1_project_data_service_proto_init() {
	if File_yandex_cloud_datasphere_v1_project_data_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileMetadata); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UploadFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DownloadFileResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*UploadFileRequest_Metadata)(nil),
		(*UploadFileRequest_Chunk)(nil),
	}
	file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*DownloadFileResponse_Metadata)(nil),
		(*DownloadFileResponse_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_datasphere_v1_project_data_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v1_project_data_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datasphere_v1_project_data_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v1_project_data_service_proto = out.File
	file_yandex_cloud_datasphere_v1_project_data_service_proto_rawDesc = nil
	file_yandex_cloud_datasphere_v1_project_data_service_proto_goTypes = nil
	file_yandex_cloud_datasphere_v1_project_data_service_proto_depIdxs = nil
}
