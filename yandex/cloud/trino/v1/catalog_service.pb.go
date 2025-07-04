// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/trino/v1/catalog_service.proto

package trino

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
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetCatalogRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino Cluster resource which contains the requested catalog.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the Trino Catalog resource.
	CatalogId     string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCatalogRequest) Reset() {
	*x = GetCatalogRequest{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCatalogRequest) ProtoMessage() {}

func (x *GetCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCatalogRequest.ProtoReflect.Descriptor instead.
func (*GetCatalogRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetCatalogRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetCatalogRequest) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

type ListCatalogsRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the cluster to list Trino Catalogs in.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListCatalogsResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, set [page_token] to the [ListCatalogsResponse.next_page_token]
	// returned by the previous list request.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	// A filter expression that filters resources listed in the response.
	// The expression must specify:
	// 1. The field name. Currently you can only use filtering with the [Catalog.name] field.
	// 2. An `=` operator.
	// 3. The value in double quotes (`"`). Must be 1-63 characters long and match the regular expression `[a-zA-Z0-9_-]+`.
	Filter        string `protobuf:"bytes,4,opt,name=filter,proto3" json:"filter,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCatalogsRequest) Reset() {
	*x = ListCatalogsRequest{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCatalogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsRequest) ProtoMessage() {}

func (x *ListCatalogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsRequest.ProtoReflect.Descriptor instead.
func (*ListCatalogsRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListCatalogsRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ListCatalogsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCatalogsRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListCatalogsRequest) GetFilter() string {
	if x != nil {
		return x.Filter
	}
	return ""
}

type ListCatalogsResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of Trino Catalog resources.
	Catalogs []*Catalog `protobuf:"bytes,1,rep,name=catalogs,proto3" json:"catalogs,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListCatalogsRequest.page_size], use the [next_page_token] as the value
	// for the [ListCatalogsRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCatalogsResponse) Reset() {
	*x = ListCatalogsResponse{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCatalogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCatalogsResponse) ProtoMessage() {}

func (x *ListCatalogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCatalogsResponse.ProtoReflect.Descriptor instead.
func (*ListCatalogsResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListCatalogsResponse) GetCatalogs() []*Catalog {
	if x != nil {
		return x.Catalogs
	}
	return nil
}

func (x *ListCatalogsResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateCatalogRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino Cluster where the catalog should be created.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Specification of the catalog to be created.
	Catalog       *CatalogSpec `protobuf:"bytes,2,opt,name=catalog,proto3" json:"catalog,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCatalogRequest) Reset() {
	*x = CreateCatalogRequest{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogRequest) ProtoMessage() {}

func (x *CreateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogRequest.ProtoReflect.Descriptor instead.
func (*CreateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCatalogRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateCatalogRequest) GetCatalog() *CatalogSpec {
	if x != nil {
		return x.Catalog
	}
	return nil
}

type CreateCatalogMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino cluster where a catalog is being created.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the catalog that is being created.
	CatalogId     string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCatalogMetadata) Reset() {
	*x = CreateCatalogMetadata{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCatalogMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCatalogMetadata) ProtoMessage() {}

func (x *CreateCatalogMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCatalogMetadata.ProtoReflect.Descriptor instead.
func (*CreateCatalogMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCatalogMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateCatalogMetadata) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

type UpdateCatalogRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino Cluster that contains the catalog to update.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the catalog to update.
	CatalogId string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	// Field mask that specifies which fields of the catalog should be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// New values for the specified fields.
	Catalog       *CatalogUpdateSpec `protobuf:"bytes,4,opt,name=catalog,proto3" json:"catalog,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCatalogRequest) Reset() {
	*x = UpdateCatalogRequest{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogRequest) ProtoMessage() {}

func (x *UpdateCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogRequest.ProtoReflect.Descriptor instead.
func (*UpdateCatalogRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCatalogRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpdateCatalogRequest) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

func (x *UpdateCatalogRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateCatalogRequest) GetCatalog() *CatalogUpdateSpec {
	if x != nil {
		return x.Catalog
	}
	return nil
}

type UpdateCatalogMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino cluster where a catalog is being updated.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the catalog that is being updated.
	CatalogId     string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCatalogMetadata) Reset() {
	*x = UpdateCatalogMetadata{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCatalogMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCatalogMetadata) ProtoMessage() {}

func (x *UpdateCatalogMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCatalogMetadata.ProtoReflect.Descriptor instead.
func (*UpdateCatalogMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCatalogMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpdateCatalogMetadata) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

type DeleteCatalogRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino Cluster resource which contains the requested catalog.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the Trino Catalog resource.
	CatalogId     string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCatalogRequest) Reset() {
	*x = DeleteCatalogRequest{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCatalogRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogRequest) ProtoMessage() {}

func (x *DeleteCatalogRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogRequest.ProtoReflect.Descriptor instead.
func (*DeleteCatalogRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCatalogRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteCatalogRequest) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

type DeleteCatalogMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the Trino cluster where a catalog is being deleted.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// ID of the catalog that is being deleted.
	CatalogId     string `protobuf:"bytes,2,opt,name=catalog_id,json=catalogId,proto3" json:"catalog_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCatalogMetadata) Reset() {
	*x = DeleteCatalogMetadata{}
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCatalogMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCatalogMetadata) ProtoMessage() {}

func (x *DeleteCatalogMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCatalogMetadata.ProtoReflect.Descriptor instead.
func (*DeleteCatalogMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCatalogMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteCatalogMetadata) GetCatalogId() string {
	if x != nil {
		return x.CatalogId
	}
	return ""
}

var File_yandex_cloud_trino_v1_catalog_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_trino_v1_catalog_service_proto_rawDesc = "" +
	"\n" +
	"+yandex/cloud/trino/v1/catalog_service.proto\x12\x15yandex.cloud.trino.v1\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\x1a yandex/cloud/api/operation.proto\x1a&yandex/cloud/operation/operation.proto\x1a#yandex/cloud/trino/v1/catalog.proto\x1a\x1dyandex/cloud/validation.proto\"m\n" +
	"\x11GetCatalogRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12+\n" +
	"\n" +
	"catalog_id\x18\x02 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tcatalogId\"\xb9\x01\n" +
	"\x13ListCatalogsRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12'\n" +
	"\tpage_size\x18\x02 \x01(\x03B\n" +
	"\xfa\xc71\x06<=1000R\bpageSize\x12(\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\t\x8a\xc81\x05<=100R\tpageToken\x12\"\n" +
	"\x06filter\x18\x04 \x01(\tB\n" +
	"\x8a\xc81\x06<=1000R\x06filter\"z\n" +
	"\x14ListCatalogsResponse\x12:\n" +
	"\bcatalogs\x18\x01 \x03(\v2\x1e.yandex.cloud.trino.v1.CatalogR\bcatalogs\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\x87\x01\n" +
	"\x14CreateCatalogRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12B\n" +
	"\acatalog\x18\x02 \x01(\v2\".yandex.cloud.trino.v1.CatalogSpecB\x04\xe8\xc71\x01R\acatalog\"U\n" +
	"\x15CreateCatalogMetadata\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12\x1d\n" +
	"\n" +
	"catalog_id\x18\x02 \x01(\tR\tcatalogId\"\xfd\x01\n" +
	"\x14UpdateCatalogRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12+\n" +
	"\n" +
	"catalog_id\x18\x02 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tcatalogId\x12A\n" +
	"\vupdate_mask\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskB\x04\xe8\xc71\x01R\n" +
	"updateMask\x12H\n" +
	"\acatalog\x18\x04 \x01(\v2(.yandex.cloud.trino.v1.CatalogUpdateSpecB\x04\xe8\xc71\x01R\acatalog\"U\n" +
	"\x15UpdateCatalogMetadata\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12\x1d\n" +
	"\n" +
	"catalog_id\x18\x02 \x01(\tR\tcatalogId\"p\n" +
	"\x14DeleteCatalogRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12+\n" +
	"\n" +
	"catalog_id\x18\x02 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tcatalogId\"U\n" +
	"\x15DeleteCatalogMetadata\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12\x1d\n" +
	"\n" +
	"catalog_id\x18\x02 \x01(\tR\tcatalogId2\x9e\a\n" +
	"\x0eCatalogService\x12\x96\x01\n" +
	"\x03Get\x12(.yandex.cloud.trino.v1.GetCatalogRequest\x1a\x1e.yandex.cloud.trino.v1.Catalog\"E\x82\xd3\xe4\x93\x02?\x12=/managed-trino/v1/clusters/{cluster_id}/catalogs/{catalog_id}\x12\x99\x01\n" +
	"\x04List\x12*.yandex.cloud.trino.v1.ListCatalogsRequest\x1a+.yandex.cloud.trino.v1.ListCatalogsResponse\"8\x82\xd3\xe4\x93\x022\x120/managed-trino/v1/clusters/{cluster_id}/catalogs\x12\xb9\x01\n" +
	"\x06Create\x12+.yandex.cloud.trino.v1.CreateCatalogRequest\x1a!.yandex.cloud.operation.Operation\"_\xb2\xd2* \n" +
	"\x15CreateCatalogMetadata\x12\aCatalog\x82\xd3\xe4\x93\x025:\x01*\"0/managed-trino/v1/clusters/{cluster_id}/catalogs\x12\xc6\x01\n" +
	"\x06Update\x12+.yandex.cloud.trino.v1.UpdateCatalogRequest\x1a!.yandex.cloud.operation.Operation\"l\xb2\xd2* \n" +
	"\x15UpdateCatalogMetadata\x12\aCatalog\x82\xd3\xe4\x93\x02B:\x01*2=/managed-trino/v1/clusters/{cluster_id}/catalogs/{catalog_id}\x12\xd1\x01\n" +
	"\x06Delete\x12+.yandex.cloud.trino.v1.DeleteCatalogRequest\x1a!.yandex.cloud.operation.Operation\"w\xb2\xd2*.\n" +
	"\x15DeleteCatalogMetadata\x12\x15google.protobuf.Empty\x82\xd3\xe4\x93\x02?*=/managed-trino/v1/clusters/{cluster_id}/catalogs/{catalog_id}B\\\n" +
	"\x19yandex.cloud.api.trino.v1Z?github.com/yandex-cloud/go-genproto/yandex/cloud/trino/v1;trinob\x06proto3"

var (
	file_yandex_cloud_trino_v1_catalog_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_trino_v1_catalog_service_proto_rawDescData []byte
)

func file_yandex_cloud_trino_v1_catalog_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_trino_v1_catalog_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_trino_v1_catalog_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_trino_v1_catalog_service_proto_rawDesc), len(file_yandex_cloud_trino_v1_catalog_service_proto_rawDesc)))
	})
	return file_yandex_cloud_trino_v1_catalog_service_proto_rawDescData
}

var file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_trino_v1_catalog_service_proto_goTypes = []any{
	(*GetCatalogRequest)(nil),     // 0: yandex.cloud.trino.v1.GetCatalogRequest
	(*ListCatalogsRequest)(nil),   // 1: yandex.cloud.trino.v1.ListCatalogsRequest
	(*ListCatalogsResponse)(nil),  // 2: yandex.cloud.trino.v1.ListCatalogsResponse
	(*CreateCatalogRequest)(nil),  // 3: yandex.cloud.trino.v1.CreateCatalogRequest
	(*CreateCatalogMetadata)(nil), // 4: yandex.cloud.trino.v1.CreateCatalogMetadata
	(*UpdateCatalogRequest)(nil),  // 5: yandex.cloud.trino.v1.UpdateCatalogRequest
	(*UpdateCatalogMetadata)(nil), // 6: yandex.cloud.trino.v1.UpdateCatalogMetadata
	(*DeleteCatalogRequest)(nil),  // 7: yandex.cloud.trino.v1.DeleteCatalogRequest
	(*DeleteCatalogMetadata)(nil), // 8: yandex.cloud.trino.v1.DeleteCatalogMetadata
	(*Catalog)(nil),               // 9: yandex.cloud.trino.v1.Catalog
	(*CatalogSpec)(nil),           // 10: yandex.cloud.trino.v1.CatalogSpec
	(*fieldmaskpb.FieldMask)(nil), // 11: google.protobuf.FieldMask
	(*CatalogUpdateSpec)(nil),     // 12: yandex.cloud.trino.v1.CatalogUpdateSpec
	(*operation.Operation)(nil),   // 13: yandex.cloud.operation.Operation
}
var file_yandex_cloud_trino_v1_catalog_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.trino.v1.ListCatalogsResponse.catalogs:type_name -> yandex.cloud.trino.v1.Catalog
	10, // 1: yandex.cloud.trino.v1.CreateCatalogRequest.catalog:type_name -> yandex.cloud.trino.v1.CatalogSpec
	11, // 2: yandex.cloud.trino.v1.UpdateCatalogRequest.update_mask:type_name -> google.protobuf.FieldMask
	12, // 3: yandex.cloud.trino.v1.UpdateCatalogRequest.catalog:type_name -> yandex.cloud.trino.v1.CatalogUpdateSpec
	0,  // 4: yandex.cloud.trino.v1.CatalogService.Get:input_type -> yandex.cloud.trino.v1.GetCatalogRequest
	1,  // 5: yandex.cloud.trino.v1.CatalogService.List:input_type -> yandex.cloud.trino.v1.ListCatalogsRequest
	3,  // 6: yandex.cloud.trino.v1.CatalogService.Create:input_type -> yandex.cloud.trino.v1.CreateCatalogRequest
	5,  // 7: yandex.cloud.trino.v1.CatalogService.Update:input_type -> yandex.cloud.trino.v1.UpdateCatalogRequest
	7,  // 8: yandex.cloud.trino.v1.CatalogService.Delete:input_type -> yandex.cloud.trino.v1.DeleteCatalogRequest
	9,  // 9: yandex.cloud.trino.v1.CatalogService.Get:output_type -> yandex.cloud.trino.v1.Catalog
	2,  // 10: yandex.cloud.trino.v1.CatalogService.List:output_type -> yandex.cloud.trino.v1.ListCatalogsResponse
	13, // 11: yandex.cloud.trino.v1.CatalogService.Create:output_type -> yandex.cloud.operation.Operation
	13, // 12: yandex.cloud.trino.v1.CatalogService.Update:output_type -> yandex.cloud.operation.Operation
	13, // 13: yandex.cloud.trino.v1.CatalogService.Delete:output_type -> yandex.cloud.operation.Operation
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_trino_v1_catalog_service_proto_init() }
func file_yandex_cloud_trino_v1_catalog_service_proto_init() {
	if File_yandex_cloud_trino_v1_catalog_service_proto != nil {
		return
	}
	file_yandex_cloud_trino_v1_catalog_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_trino_v1_catalog_service_proto_rawDesc), len(file_yandex_cloud_trino_v1_catalog_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_trino_v1_catalog_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_trino_v1_catalog_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_trino_v1_catalog_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_trino_v1_catalog_service_proto = out.File
	file_yandex_cloud_trino_v1_catalog_service_proto_goTypes = nil
	file_yandex_cloud_trino_v1_catalog_service_proto_depIdxs = nil
}
