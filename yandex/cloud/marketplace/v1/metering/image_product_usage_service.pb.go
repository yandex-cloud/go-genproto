// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/marketplace/v1/metering/image_product_usage_service.proto

package metering

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type WriteImageProductUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Checks whether you have the access required for the emit usage.
	ValidateOnly bool `protobuf:"varint,1,opt,name=validate_only,json=validateOnly,proto3" json:"validate_only,omitempty"`
	// Marketplace Product's ID.
	ProductId string `protobuf:"bytes,2,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	// List of product usage records (up to 25 pet request).
	UsageRecords []*UsageRecord `protobuf:"bytes,3,rep,name=usage_records,json=usageRecords,proto3" json:"usage_records,omitempty"`
}

func (x *WriteImageProductUsageRequest) Reset() {
	*x = WriteImageProductUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteImageProductUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteImageProductUsageRequest) ProtoMessage() {}

func (x *WriteImageProductUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteImageProductUsageRequest.ProtoReflect.Descriptor instead.
func (*WriteImageProductUsageRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescGZIP(), []int{0}
}

func (x *WriteImageProductUsageRequest) GetValidateOnly() bool {
	if x != nil {
		return x.ValidateOnly
	}
	return false
}

func (x *WriteImageProductUsageRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *WriteImageProductUsageRequest) GetUsageRecords() []*UsageRecord {
	if x != nil {
		return x.UsageRecords
	}
	return nil
}

type WriteImageProductUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of accepted product usage records.
	Accepted []*AcceptedUsageRecord `protobuf:"bytes,1,rep,name=accepted,proto3" json:"accepted,omitempty"`
	// List of rejected product usage records (with reason).
	Rejected []*RejectedUsageRecord `protobuf:"bytes,2,rep,name=rejected,proto3" json:"rejected,omitempty"`
}

func (x *WriteImageProductUsageResponse) Reset() {
	*x = WriteImageProductUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteImageProductUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteImageProductUsageResponse) ProtoMessage() {}

func (x *WriteImageProductUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteImageProductUsageResponse.ProtoReflect.Descriptor instead.
func (*WriteImageProductUsageResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescGZIP(), []int{1}
}

func (x *WriteImageProductUsageResponse) GetAccepted() []*AcceptedUsageRecord {
	if x != nil {
		return x.Accepted
	}
	return nil
}

func (x *WriteImageProductUsageResponse) GetRejected() []*RejectedUsageRecord {
	if x != nil {
		return x.Rejected
	}
	return nil
}

var File_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDesc = []byte{
	0x0a, 0x46, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x24, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e,
	0x67, 0x2f, 0x75, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x1d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x0a, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x60, 0x0a, 0x0d, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x42, 0x08, 0x82, 0xc8, 0x31, 0x04, 0x31, 0x2d, 0x32, 0x35, 0x52, 0x0c, 0x75, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xce, 0x01, 0x0a, 0x1e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x55, 0x0a,
	0x08, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x39, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x08, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x12, 0x55, 0x0a, 0x08, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65,
	0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x08, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x32, 0xec, 0x01, 0x0a, 0x18,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xcf, 0x01, 0x0a, 0x05, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x12, 0x43, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x44, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3b, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x35, 0x22, 0x30, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x2f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x7d, 0x0a, 0x28, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67,
	0x3b, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescData = file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDesc
)

func file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescData)
	})
	return file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDescData
}

var file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_goTypes = []interface{}{
	(*WriteImageProductUsageRequest)(nil),  // 0: yandex.cloud.marketplace.v1.metering.WriteImageProductUsageRequest
	(*WriteImageProductUsageResponse)(nil), // 1: yandex.cloud.marketplace.v1.metering.WriteImageProductUsageResponse
	(*UsageRecord)(nil),                    // 2: yandex.cloud.marketplace.v1.metering.UsageRecord
	(*AcceptedUsageRecord)(nil),            // 3: yandex.cloud.marketplace.v1.metering.AcceptedUsageRecord
	(*RejectedUsageRecord)(nil),            // 4: yandex.cloud.marketplace.v1.metering.RejectedUsageRecord
}
var file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.marketplace.v1.metering.WriteImageProductUsageRequest.usage_records:type_name -> yandex.cloud.marketplace.v1.metering.UsageRecord
	3, // 1: yandex.cloud.marketplace.v1.metering.WriteImageProductUsageResponse.accepted:type_name -> yandex.cloud.marketplace.v1.metering.AcceptedUsageRecord
	4, // 2: yandex.cloud.marketplace.v1.metering.WriteImageProductUsageResponse.rejected:type_name -> yandex.cloud.marketplace.v1.metering.RejectedUsageRecord
	0, // 3: yandex.cloud.marketplace.v1.metering.ImageProductUsageService.Write:input_type -> yandex.cloud.marketplace.v1.metering.WriteImageProductUsageRequest
	1, // 4: yandex.cloud.marketplace.v1.metering.ImageProductUsageService.Write:output_type -> yandex.cloud.marketplace.v1.metering.WriteImageProductUsageResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_init() }
func file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_init() {
	if File_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto != nil {
		return
	}
	file_yandex_cloud_marketplace_v1_metering_usage_record_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteImageProductUsageRequest); i {
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
		file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteImageProductUsageResponse); i {
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
			RawDescriptor: file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto = out.File
	file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_rawDesc = nil
	file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_goTypes = nil
	file_yandex_cloud_marketplace_v1_metering_image_product_usage_service_proto_depIdxs = nil
}
