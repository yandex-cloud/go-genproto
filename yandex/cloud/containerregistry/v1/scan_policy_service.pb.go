// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/containerregistry/v1/scan_policy_service.proto

package containerregistry

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetScanPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the scan policy.
	ScanPolicyId string `protobuf:"bytes,1,opt,name=scan_policy_id,json=scanPolicyId,proto3" json:"scan_policy_id,omitempty"`
}

func (x *GetScanPolicyRequest) Reset() {
	*x = GetScanPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScanPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScanPolicyRequest) ProtoMessage() {}

func (x *GetScanPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScanPolicyRequest.ProtoReflect.Descriptor instead.
func (*GetScanPolicyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetScanPolicyRequest) GetScanPolicyId() string {
	if x != nil {
		return x.ScanPolicyId
	}
	return ""
}

type GetScanPolicyByRegistryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the registry with scan policy.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
}

func (x *GetScanPolicyByRegistryRequest) Reset() {
	*x = GetScanPolicyByRegistryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetScanPolicyByRegistryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetScanPolicyByRegistryRequest) ProtoMessage() {}

func (x *GetScanPolicyByRegistryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetScanPolicyByRegistryRequest.ProtoReflect.Descriptor instead.
func (*GetScanPolicyByRegistryRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetScanPolicyByRegistryRequest) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

type CreateScanPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the scan policy registry.
	RegistryId string `protobuf:"bytes,1,opt,name=registry_id,json=registryId,proto3" json:"registry_id,omitempty"`
	// Name of the scan policy.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the scan policy.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Rules of the scan policy.
	Rules *ScanRules `protobuf:"bytes,4,opt,name=rules,proto3" json:"rules,omitempty"`
}

func (x *CreateScanPolicyRequest) Reset() {
	*x = CreateScanPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScanPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScanPolicyRequest) ProtoMessage() {}

func (x *CreateScanPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScanPolicyRequest.ProtoReflect.Descriptor instead.
func (*CreateScanPolicyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateScanPolicyRequest) GetRegistryId() string {
	if x != nil {
		return x.RegistryId
	}
	return ""
}

func (x *CreateScanPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateScanPolicyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateScanPolicyRequest) GetRules() *ScanRules {
	if x != nil {
		return x.Rules
	}
	return nil
}

type UpdateScanPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the scan policy.
	ScanPolicyId string `protobuf:"bytes,1,opt,name=scan_policy_id,json=scanPolicyId,proto3" json:"scan_policy_id,omitempty"`
	// Field mask that specifies which fields of the scan policy resource are going to be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// Name of the scan policy.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the scan policy.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Rules of the scan policy.
	Rules *ScanRules `protobuf:"bytes,5,opt,name=rules,proto3" json:"rules,omitempty"`
}

func (x *UpdateScanPolicyRequest) Reset() {
	*x = UpdateScanPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScanPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScanPolicyRequest) ProtoMessage() {}

func (x *UpdateScanPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScanPolicyRequest.ProtoReflect.Descriptor instead.
func (*UpdateScanPolicyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateScanPolicyRequest) GetScanPolicyId() string {
	if x != nil {
		return x.ScanPolicyId
	}
	return ""
}

func (x *UpdateScanPolicyRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateScanPolicyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateScanPolicyRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateScanPolicyRequest) GetRules() *ScanRules {
	if x != nil {
		return x.Rules
	}
	return nil
}

type DeleteScanPolicyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the scan policy.
	ScanPolicyId string `protobuf:"bytes,1,opt,name=scan_policy_id,json=scanPolicyId,proto3" json:"scan_policy_id,omitempty"`
}

func (x *DeleteScanPolicyRequest) Reset() {
	*x = DeleteScanPolicyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScanPolicyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScanPolicyRequest) ProtoMessage() {}

func (x *DeleteScanPolicyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScanPolicyRequest.ProtoReflect.Descriptor instead.
func (*DeleteScanPolicyRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteScanPolicyRequest) GetScanPolicyId() string {
	if x != nil {
		return x.ScanPolicyId
	}
	return ""
}

type CreateScanPolicyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of created scan policy resource.
	ScanPolicyId string `protobuf:"bytes,1,opt,name=scan_policy_id,json=scanPolicyId,proto3" json:"scan_policy_id,omitempty"`
}

func (x *CreateScanPolicyMetadata) Reset() {
	*x = CreateScanPolicyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScanPolicyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScanPolicyMetadata) ProtoMessage() {}

func (x *CreateScanPolicyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScanPolicyMetadata.ProtoReflect.Descriptor instead.
func (*CreateScanPolicyMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{5}
}

func (x *CreateScanPolicyMetadata) GetScanPolicyId() string {
	if x != nil {
		return x.ScanPolicyId
	}
	return ""
}

type UpdateScanPolicyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the scan policy resource that is updated.
	ScanPolicyId string `protobuf:"bytes,1,opt,name=scan_policy_id,json=scanPolicyId,proto3" json:"scan_policy_id,omitempty"`
}

func (x *UpdateScanPolicyMetadata) Reset() {
	*x = UpdateScanPolicyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateScanPolicyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateScanPolicyMetadata) ProtoMessage() {}

func (x *UpdateScanPolicyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateScanPolicyMetadata.ProtoReflect.Descriptor instead.
func (*UpdateScanPolicyMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateScanPolicyMetadata) GetScanPolicyId() string {
	if x != nil {
		return x.ScanPolicyId
	}
	return ""
}

type DeleteScanPolicyMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the scan policy resource that is deleted.
	ScanPolicyId string `protobuf:"bytes,1,opt,name=scan_policy_id,json=scanPolicyId,proto3" json:"scan_policy_id,omitempty"`
}

func (x *DeleteScanPolicyMetadata) Reset() {
	*x = DeleteScanPolicyMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteScanPolicyMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteScanPolicyMetadata) ProtoMessage() {}

func (x *DeleteScanPolicyMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteScanPolicyMetadata.ProtoReflect.Descriptor instead.
func (*DeleteScanPolicyMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteScanPolicyMetadata) GetScanPolicyId() string {
	if x != nil {
		return x.ScanPolicyId
	}
	return ""
}

var File_yandex_cloud_containerregistry_v1_scan_policy_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4a,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0c, 0x73, 0x63,
	0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x1e, 0x47, 0x65,
	0x74, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x79, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52,
	0x0a, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x22, 0xf0, 0x01, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7,
	0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xf2, 0xc7, 0x31, 0x1d, 0x7c, 0x5b, 0x61, 0x2d, 0x7a, 0x5d,
	0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31, 0x2c, 0x36, 0x31, 0x7d, 0x5b,
	0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32, 0x35, 0x36, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x61, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x22, 0xb2,
	0x02, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x0e, 0x73, 0x63,
	0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30,
	0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73, 0x6b, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x12, 0x35, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x21, 0xf2, 0xc7, 0x31, 0x1d, 0x7c,
	0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x5b, 0x2d, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x31,
	0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a, 0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x32,
	0x35, 0x36, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x42, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x05, 0x72, 0x75,
	0x6c, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x61,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32,
	0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x49, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x24,
	0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x24, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x24, 0x0a, 0x0e, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x63, 0x61, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x49, 0x64, 0x32, 0x81, 0x08, 0x0a, 0x11, 0x53, 0x63, 0x61,
	0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xab,
	0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x37, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63,
	0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x22, 0x3c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x12, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x61,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0xc7, 0x01, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x41,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x42, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x22, 0x44, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x3e, 0x12, 0x3c, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x62, 0x79, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0xc1, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x58, 0xb2, 0xd2, 0x2a, 0x26, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63,
	0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x28, 0x3a, 0x01, 0x2a, 0x22, 0x23, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x63,
	0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x12, 0xd2, 0x01, 0x0a, 0x06, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x69, 0xb2, 0xd2, 0x2a, 0x26, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x0a, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x39, 0x3a, 0x01, 0x2a, 0x32, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x73, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b,
	0x73, 0x63, 0x61, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x12,
	0xda, 0x01, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x3a, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x71, 0xb2, 0xd2, 0x2a, 0x31, 0x0a,
	0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x36, 0x2a, 0x34, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x2d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x63, 0x61, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2f, 0x7b, 0x73, 0x63, 0x61,
	0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5f, 0x69, 0x64, 0x7d, 0x42, 0x80, 0x01, 0x0a,
	0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescData = file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDesc
)

func file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescData)
	})
	return file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDescData
}

var file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_goTypes = []any{
	(*GetScanPolicyRequest)(nil),           // 0: yandex.cloud.containerregistry.v1.GetScanPolicyRequest
	(*GetScanPolicyByRegistryRequest)(nil), // 1: yandex.cloud.containerregistry.v1.GetScanPolicyByRegistryRequest
	(*CreateScanPolicyRequest)(nil),        // 2: yandex.cloud.containerregistry.v1.CreateScanPolicyRequest
	(*UpdateScanPolicyRequest)(nil),        // 3: yandex.cloud.containerregistry.v1.UpdateScanPolicyRequest
	(*DeleteScanPolicyRequest)(nil),        // 4: yandex.cloud.containerregistry.v1.DeleteScanPolicyRequest
	(*CreateScanPolicyMetadata)(nil),       // 5: yandex.cloud.containerregistry.v1.CreateScanPolicyMetadata
	(*UpdateScanPolicyMetadata)(nil),       // 6: yandex.cloud.containerregistry.v1.UpdateScanPolicyMetadata
	(*DeleteScanPolicyMetadata)(nil),       // 7: yandex.cloud.containerregistry.v1.DeleteScanPolicyMetadata
	(*ScanRules)(nil),                      // 8: yandex.cloud.containerregistry.v1.ScanRules
	(*fieldmaskpb.FieldMask)(nil),          // 9: google.protobuf.FieldMask
	(*ScanPolicy)(nil),                     // 10: yandex.cloud.containerregistry.v1.ScanPolicy
	(*operation.Operation)(nil),            // 11: yandex.cloud.operation.Operation
}
var file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_depIdxs = []int32{
	8,  // 0: yandex.cloud.containerregistry.v1.CreateScanPolicyRequest.rules:type_name -> yandex.cloud.containerregistry.v1.ScanRules
	9,  // 1: yandex.cloud.containerregistry.v1.UpdateScanPolicyRequest.update_mask:type_name -> google.protobuf.FieldMask
	8,  // 2: yandex.cloud.containerregistry.v1.UpdateScanPolicyRequest.rules:type_name -> yandex.cloud.containerregistry.v1.ScanRules
	0,  // 3: yandex.cloud.containerregistry.v1.ScanPolicyService.Get:input_type -> yandex.cloud.containerregistry.v1.GetScanPolicyRequest
	1,  // 4: yandex.cloud.containerregistry.v1.ScanPolicyService.GetByRegistry:input_type -> yandex.cloud.containerregistry.v1.GetScanPolicyByRegistryRequest
	2,  // 5: yandex.cloud.containerregistry.v1.ScanPolicyService.Create:input_type -> yandex.cloud.containerregistry.v1.CreateScanPolicyRequest
	3,  // 6: yandex.cloud.containerregistry.v1.ScanPolicyService.Update:input_type -> yandex.cloud.containerregistry.v1.UpdateScanPolicyRequest
	4,  // 7: yandex.cloud.containerregistry.v1.ScanPolicyService.Delete:input_type -> yandex.cloud.containerregistry.v1.DeleteScanPolicyRequest
	10, // 8: yandex.cloud.containerregistry.v1.ScanPolicyService.Get:output_type -> yandex.cloud.containerregistry.v1.ScanPolicy
	10, // 9: yandex.cloud.containerregistry.v1.ScanPolicyService.GetByRegistry:output_type -> yandex.cloud.containerregistry.v1.ScanPolicy
	11, // 10: yandex.cloud.containerregistry.v1.ScanPolicyService.Create:output_type -> yandex.cloud.operation.Operation
	11, // 11: yandex.cloud.containerregistry.v1.ScanPolicyService.Update:output_type -> yandex.cloud.operation.Operation
	11, // 12: yandex.cloud.containerregistry.v1.ScanPolicyService.Delete:output_type -> yandex.cloud.operation.Operation
	8,  // [8:13] is the sub-list for method output_type
	3,  // [3:8] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_init() }
func file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_init() {
	if File_yandex_cloud_containerregistry_v1_scan_policy_service_proto != nil {
		return
	}
	file_yandex_cloud_containerregistry_v1_scan_policy_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetScanPolicyRequest); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetScanPolicyByRegistryRequest); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateScanPolicyRequest); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateScanPolicyRequest); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteScanPolicyRequest); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CreateScanPolicyMetadata); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateScanPolicyMetadata); i {
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
		file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteScanPolicyMetadata); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_containerregistry_v1_scan_policy_service_proto = out.File
	file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_rawDesc = nil
	file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_scan_policy_service_proto_depIdxs = nil
}
