// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/mdb/postgresql/v1/database_service.proto

package postgresql

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

type GetDatabaseRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster that the database belongs to.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the PostgreSQL Database resource to return.
	// To get the name of the database use a [DatabaseService.List] request.
	DatabaseName  string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetDatabaseRequest) Reset() {
	*x = GetDatabaseRequest{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDatabaseRequest) ProtoMessage() {}

func (x *GetDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDatabaseRequest.ProtoReflect.Descriptor instead.
func (*GetDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *GetDatabaseRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type ListDatabasesRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster to list databases in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size], the service returns a [ListDatabasesResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results, Set [page_token] to the [ListDatabasesResponse.next_page_token]
	// returned by the previous list request.
	PageToken     string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDatabasesRequest) Reset() {
	*x = ListDatabasesRequest{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatabasesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabasesRequest) ProtoMessage() {}

func (x *ListDatabasesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabasesRequest.ProtoReflect.Descriptor instead.
func (*ListDatabasesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListDatabasesRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *ListDatabasesRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListDatabasesRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListDatabasesResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// List of PostgreSQL Database resources.
	Databases []*Database `protobuf:"bytes,1,rep,name=databases,proto3" json:"databases,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListDatabasesRequest.page_size], use the [next_page_token] as the value
	// for the [ListDatabasesRequest.page_token] parameter in the next list request. Each subsequent
	// list request will have its own [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListDatabasesResponse) Reset() {
	*x = ListDatabasesResponse{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListDatabasesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListDatabasesResponse) ProtoMessage() {}

func (x *ListDatabasesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListDatabasesResponse.ProtoReflect.Descriptor instead.
func (*ListDatabasesResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListDatabasesResponse) GetDatabases() []*Database {
	if x != nil {
		return x.Databases
	}
	return nil
}

func (x *ListDatabasesResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type CreateDatabaseRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster to create a database in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Configuration of the database to create.
	DatabaseSpec  *DatabaseSpec `protobuf:"bytes,2,opt,name=database_spec,json=databaseSpec,proto3" json:"database_spec,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDatabaseRequest) Reset() {
	*x = CreateDatabaseRequest{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseRequest) ProtoMessage() {}

func (x *CreateDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseRequest.ProtoReflect.Descriptor instead.
func (*CreateDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateDatabaseRequest) GetDatabaseSpec() *DatabaseSpec {
	if x != nil {
		return x.DatabaseSpec
	}
	return nil
}

type CreateDatabaseMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster where a database is being created.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the PostgreSQL database that is being created.
	DatabaseName  string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateDatabaseMetadata) Reset() {
	*x = CreateDatabaseMetadata{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateDatabaseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDatabaseMetadata) ProtoMessage() {}

func (x *CreateDatabaseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDatabaseMetadata.ProtoReflect.Descriptor instead.
func (*CreateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDatabaseMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *CreateDatabaseMetadata) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type UpdateDatabaseRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster to update a database in.
	// To get the cluster ID use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the database to update.
	// To get the name of the database use a [DatabaseService.List] request.
	DatabaseName string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	// Optional. New name of the database.
	NewDatabaseName string `protobuf:"bytes,5,opt,name=new_database_name,json=newDatabaseName,proto3" json:"new_database_name,omitempty"`
	// Field mask that specifies which fields of the Database resource should be updated.
	UpdateMask *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	// PostgreSQL extensions that should be enabled for the database.
	//
	// If the field is sent, the list of enabled extensions is rewritten entirely.
	// Therefore, to disable an active extension you should simply send the list omitting this extension.
	Extensions []*Extension `protobuf:"bytes,4,rep,name=extensions,proto3" json:"extensions,omitempty"`
	// Deletion Protection inhibits deletion of the database
	//
	// Default value: `unspecified` (inherits cluster's deletion_protection)
	DeletionProtection *wrapperspb.BoolValue `protobuf:"bytes,6,opt,name=deletion_protection,json=deletionProtection,proto3" json:"deletion_protection,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *UpdateDatabaseRequest) Reset() {
	*x = UpdateDatabaseRequest{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseRequest) ProtoMessage() {}

func (x *UpdateDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseRequest.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetNewDatabaseName() string {
	if x != nil {
		return x.NewDatabaseName
	}
	return ""
}

func (x *UpdateDatabaseRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

func (x *UpdateDatabaseRequest) GetExtensions() []*Extension {
	if x != nil {
		return x.Extensions
	}
	return nil
}

func (x *UpdateDatabaseRequest) GetDeletionProtection() *wrapperspb.BoolValue {
	if x != nil {
		return x.DeletionProtection
	}
	return nil
}

type UpdateDatabaseMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster where a database is being updated.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the PostgreSQL database that is being updated.
	DatabaseName  string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDatabaseMetadata) Reset() {
	*x = UpdateDatabaseMetadata{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDatabaseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDatabaseMetadata) ProtoMessage() {}

func (x *UpdateDatabaseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDatabaseMetadata.ProtoReflect.Descriptor instead.
func (*UpdateDatabaseMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateDatabaseMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *UpdateDatabaseMetadata) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type DeleteDatabaseRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster to delete a database in.
	// To get the cluster ID, use a [ClusterService.List] request.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the database to delete.
	// To get the name of the database, use a [DatabaseService.List] request.
	DatabaseName  string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDatabaseRequest) Reset() {
	*x = DeleteDatabaseRequest{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDatabaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseRequest) ProtoMessage() {}

func (x *DeleteDatabaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseRequest.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteDatabaseRequest) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteDatabaseRequest) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

type DeleteDatabaseMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the PostgreSQL cluster where a database is being deleted.
	ClusterId string `protobuf:"bytes,1,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Name of the PostgreSQL database that is being deleted.
	DatabaseName  string `protobuf:"bytes,2,opt,name=database_name,json=databaseName,proto3" json:"database_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteDatabaseMetadata) Reset() {
	*x = DeleteDatabaseMetadata{}
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteDatabaseMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteDatabaseMetadata) ProtoMessage() {}

func (x *DeleteDatabaseMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteDatabaseMetadata.ProtoReflect.Descriptor instead.
func (*DeleteDatabaseMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteDatabaseMetadata) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *DeleteDatabaseMetadata) GetDatabaseName() string {
	if x != nil {
		return x.DatabaseName
	}
	return ""
}

var File_yandex_cloud_mdb_postgresql_v1_database_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDesc = "" +
	"\n" +
	"5yandex/cloud/mdb/postgresql/v1/database_service.proto\x12\x1eyandex.cloud.mdb.postgresql.v1\x1a\x1cgoogle/api/annotations.proto\x1a google/protobuf/field_mask.proto\x1a\x1egoogle/protobuf/wrappers.proto\x1a yandex/cloud/api/operation.proto\x1a\x1dyandex/cloud/validation.proto\x1a&yandex/cloud/operation/operation.proto\x1a-yandex/cloud/mdb/postgresql/v1/database.proto\"\x86\x01\n" +
	"\x12GetDatabaseRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12C\n" +
	"\rdatabase_name\x18\x02 \x01(\tB\x1e\xe8\xc71\x01\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\fdatabaseName\"\x96\x01\n" +
	"\x14ListDatabasesRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12'\n" +
	"\tpage_size\x18\x02 \x01(\x03B\n" +
	"\xfa\xc71\x06<=1000R\bpageSize\x12(\n" +
	"\n" +
	"page_token\x18\x03 \x01(\tB\t\x8a\xc81\x05<=100R\tpageToken\"\x87\x01\n" +
	"\x15ListDatabasesResponse\x12F\n" +
	"\tdatabases\x18\x01 \x03(\v2(.yandex.cloud.mdb.postgresql.v1.DatabaseR\tdatabases\x12&\n" +
	"\x0fnext_page_token\x18\x02 \x01(\tR\rnextPageToken\"\x9d\x01\n" +
	"\x15CreateDatabaseRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12W\n" +
	"\rdatabase_spec\x18\x02 \x01(\v2,.yandex.cloud.mdb.postgresql.v1.DatabaseSpecB\x04\xe8\xc71\x01R\fdatabaseSpec\"\\\n" +
	"\x16CreateDatabaseMetadata\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12#\n" +
	"\rdatabase_name\x18\x02 \x01(\tR\fdatabaseName\"\xaa\x03\n" +
	"\x15UpdateDatabaseRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12C\n" +
	"\rdatabase_name\x18\x02 \x01(\tB\x1e\xe8\xc71\x01\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\fdatabaseName\x12J\n" +
	"\x11new_database_name\x18\x05 \x01(\tB\x1e\xe8\xc71\x00\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\x0fnewDatabaseName\x12;\n" +
	"\vupdate_mask\x18\x03 \x01(\v2\x1a.google.protobuf.FieldMaskR\n" +
	"updateMask\x12I\n" +
	"\n" +
	"extensions\x18\x04 \x03(\v2).yandex.cloud.mdb.postgresql.v1.ExtensionR\n" +
	"extensions\x12K\n" +
	"\x13deletion_protection\x18\x06 \x01(\v2\x1a.google.protobuf.BoolValueR\x12deletionProtection\"\\\n" +
	"\x16UpdateDatabaseMetadata\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12#\n" +
	"\rdatabase_name\x18\x02 \x01(\tR\fdatabaseName\"\x89\x01\n" +
	"\x15DeleteDatabaseRequest\x12+\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tB\f\xe8\xc71\x01\x8a\xc81\x04<=50R\tclusterId\x12C\n" +
	"\rdatabase_name\x18\x02 \x01(\tB\x1e\xe8\xc71\x01\xf2\xc71\x0e[a-zA-Z0-9_-]*\x8a\xc81\x04<=63R\fdatabaseName\"\\\n" +
	"\x16DeleteDatabaseMetadata\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x01 \x01(\tR\tclusterId\x12#\n" +
	"\rdatabase_name\x18\x02 \x01(\tR\fdatabaseName2\x92\b\n" +
	"\x0fDatabaseService\x12\xb3\x01\n" +
	"\x03Get\x122.yandex.cloud.mdb.postgresql.v1.GetDatabaseRequest\x1a(.yandex.cloud.mdb.postgresql.v1.Database\"N\x82\xd3\xe4\x93\x02H\x12F/managed-postgresql/v1/clusters/{cluster_id}/databases/{database_name}\x12\xb3\x01\n" +
	"\x04List\x124.yandex.cloud.mdb.postgresql.v1.ListDatabasesRequest\x1a5.yandex.cloud.mdb.postgresql.v1.ListDatabasesResponse\">\x82\xd3\xe4\x93\x028\x126/managed-postgresql/v1/clusters/{cluster_id}/databases\x12\xcb\x01\n" +
	"\x06Create\x125.yandex.cloud.mdb.postgresql.v1.CreateDatabaseRequest\x1a!.yandex.cloud.operation.Operation\"g\xb2\xd2*\"\n" +
	"\x16CreateDatabaseMetadata\x12\bDatabase\x82\xd3\xe4\x93\x02;:\x01*\"6/managed-postgresql/v1/clusters/{cluster_id}/databases\x12\xdb\x01\n" +
	"\x06Update\x125.yandex.cloud.mdb.postgresql.v1.UpdateDatabaseRequest\x1a!.yandex.cloud.operation.Operation\"w\xb2\xd2*\"\n" +
	"\x16UpdateDatabaseMetadata\x12\bDatabase\x82\xd3\xe4\x93\x02K:\x01*2F/managed-postgresql/v1/clusters/{cluster_id}/databases/{database_name}\x12\xe6\x01\n" +
	"\x06Delete\x125.yandex.cloud.mdb.postgresql.v1.DeleteDatabaseRequest\x1a!.yandex.cloud.operation.Operation\"\x81\x01\xb2\xd2*/\n" +
	"\x16DeleteDatabaseMetadata\x12\x15google.protobuf.Empty\x82\xd3\xe4\x93\x02H*F/managed-postgresql/v1/clusters/{cluster_id}/databases/{database_name}Bs\n" +
	"\"yandex.cloud.api.mdb.postgresql.v1ZMgithub.com/yandex-cloud/go-genproto/yandex/cloud/mdb/postgresql/v1;postgresqlb\x06proto3"

var (
	file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescData []byte
)

func file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDesc), len(file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDesc)))
	})
	return file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDescData
}

var file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_mdb_postgresql_v1_database_service_proto_goTypes = []any{
	(*GetDatabaseRequest)(nil),     // 0: yandex.cloud.mdb.postgresql.v1.GetDatabaseRequest
	(*ListDatabasesRequest)(nil),   // 1: yandex.cloud.mdb.postgresql.v1.ListDatabasesRequest
	(*ListDatabasesResponse)(nil),  // 2: yandex.cloud.mdb.postgresql.v1.ListDatabasesResponse
	(*CreateDatabaseRequest)(nil),  // 3: yandex.cloud.mdb.postgresql.v1.CreateDatabaseRequest
	(*CreateDatabaseMetadata)(nil), // 4: yandex.cloud.mdb.postgresql.v1.CreateDatabaseMetadata
	(*UpdateDatabaseRequest)(nil),  // 5: yandex.cloud.mdb.postgresql.v1.UpdateDatabaseRequest
	(*UpdateDatabaseMetadata)(nil), // 6: yandex.cloud.mdb.postgresql.v1.UpdateDatabaseMetadata
	(*DeleteDatabaseRequest)(nil),  // 7: yandex.cloud.mdb.postgresql.v1.DeleteDatabaseRequest
	(*DeleteDatabaseMetadata)(nil), // 8: yandex.cloud.mdb.postgresql.v1.DeleteDatabaseMetadata
	(*Database)(nil),               // 9: yandex.cloud.mdb.postgresql.v1.Database
	(*DatabaseSpec)(nil),           // 10: yandex.cloud.mdb.postgresql.v1.DatabaseSpec
	(*fieldmaskpb.FieldMask)(nil),  // 11: google.protobuf.FieldMask
	(*Extension)(nil),              // 12: yandex.cloud.mdb.postgresql.v1.Extension
	(*wrapperspb.BoolValue)(nil),   // 13: google.protobuf.BoolValue
	(*operation.Operation)(nil),    // 14: yandex.cloud.operation.Operation
}
var file_yandex_cloud_mdb_postgresql_v1_database_service_proto_depIdxs = []int32{
	9,  // 0: yandex.cloud.mdb.postgresql.v1.ListDatabasesResponse.databases:type_name -> yandex.cloud.mdb.postgresql.v1.Database
	10, // 1: yandex.cloud.mdb.postgresql.v1.CreateDatabaseRequest.database_spec:type_name -> yandex.cloud.mdb.postgresql.v1.DatabaseSpec
	11, // 2: yandex.cloud.mdb.postgresql.v1.UpdateDatabaseRequest.update_mask:type_name -> google.protobuf.FieldMask
	12, // 3: yandex.cloud.mdb.postgresql.v1.UpdateDatabaseRequest.extensions:type_name -> yandex.cloud.mdb.postgresql.v1.Extension
	13, // 4: yandex.cloud.mdb.postgresql.v1.UpdateDatabaseRequest.deletion_protection:type_name -> google.protobuf.BoolValue
	0,  // 5: yandex.cloud.mdb.postgresql.v1.DatabaseService.Get:input_type -> yandex.cloud.mdb.postgresql.v1.GetDatabaseRequest
	1,  // 6: yandex.cloud.mdb.postgresql.v1.DatabaseService.List:input_type -> yandex.cloud.mdb.postgresql.v1.ListDatabasesRequest
	3,  // 7: yandex.cloud.mdb.postgresql.v1.DatabaseService.Create:input_type -> yandex.cloud.mdb.postgresql.v1.CreateDatabaseRequest
	5,  // 8: yandex.cloud.mdb.postgresql.v1.DatabaseService.Update:input_type -> yandex.cloud.mdb.postgresql.v1.UpdateDatabaseRequest
	7,  // 9: yandex.cloud.mdb.postgresql.v1.DatabaseService.Delete:input_type -> yandex.cloud.mdb.postgresql.v1.DeleteDatabaseRequest
	9,  // 10: yandex.cloud.mdb.postgresql.v1.DatabaseService.Get:output_type -> yandex.cloud.mdb.postgresql.v1.Database
	2,  // 11: yandex.cloud.mdb.postgresql.v1.DatabaseService.List:output_type -> yandex.cloud.mdb.postgresql.v1.ListDatabasesResponse
	14, // 12: yandex.cloud.mdb.postgresql.v1.DatabaseService.Create:output_type -> yandex.cloud.operation.Operation
	14, // 13: yandex.cloud.mdb.postgresql.v1.DatabaseService.Update:output_type -> yandex.cloud.operation.Operation
	14, // 14: yandex.cloud.mdb.postgresql.v1.DatabaseService.Delete:output_type -> yandex.cloud.operation.Operation
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_mdb_postgresql_v1_database_service_proto_init() }
func file_yandex_cloud_mdb_postgresql_v1_database_service_proto_init() {
	if File_yandex_cloud_mdb_postgresql_v1_database_service_proto != nil {
		return
	}
	file_yandex_cloud_mdb_postgresql_v1_database_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDesc), len(file_yandex_cloud_mdb_postgresql_v1_database_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_mdb_postgresql_v1_database_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_mdb_postgresql_v1_database_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_mdb_postgresql_v1_database_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_mdb_postgresql_v1_database_service_proto = out.File
	file_yandex_cloud_mdb_postgresql_v1_database_service_proto_goTypes = nil
	file_yandex_cloud_mdb_postgresql_v1_database_service_proto_depIdxs = nil
}
