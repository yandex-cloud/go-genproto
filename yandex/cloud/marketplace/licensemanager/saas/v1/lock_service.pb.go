// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/marketplace/licensemanager/saas/v1/lock_service.proto

package licensemanager

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/marketplace/licensemanager/v1"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type EnsureLockRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Signed JWT token which contains information about subscription.
	InstanceToken string `protobuf:"bytes,1,opt,name=instance_token,json=instanceToken,proto3" json:"instance_token,omitempty"`
	// ID of the resource to which the subscription will be locked.
	ResourceId    string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnsureLockRequest) Reset() {
	*x = EnsureLockRequest{}
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnsureLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsureLockRequest) ProtoMessage() {}

func (x *EnsureLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsureLockRequest.ProtoReflect.Descriptor instead.
func (*EnsureLockRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescGZIP(), []int{0}
}

func (x *EnsureLockRequest) GetInstanceToken() string {
	if x != nil {
		return x.InstanceToken
	}
	return ""
}

func (x *EnsureLockRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type EnsureLockMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the subscription lock.
	LockId        string `protobuf:"bytes,1,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EnsureLockMetadata) Reset() {
	*x = EnsureLockMetadata{}
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EnsureLockMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnsureLockMetadata) ProtoMessage() {}

func (x *EnsureLockMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnsureLockMetadata.ProtoReflect.Descriptor instead.
func (*EnsureLockMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescGZIP(), []int{1}
}

func (x *EnsureLockMetadata) GetLockId() string {
	if x != nil {
		return x.LockId
	}
	return ""
}

type GetLockRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the subscription lock.
	LockId        string `protobuf:"bytes,1,opt,name=lock_id,json=lockId,proto3" json:"lock_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLockRequest) Reset() {
	*x = GetLockRequest{}
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockRequest) ProtoMessage() {}

func (x *GetLockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockRequest.ProtoReflect.Descriptor instead.
func (*GetLockRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLockRequest) GetLockId() string {
	if x != nil {
		return x.LockId
	}
	return ""
}

type GetLockByResourceIDRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the resource to with subscription is locked.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// ID of the subscription
	InstanceId    string `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLockByResourceIDRequest) Reset() {
	*x = GetLockByResourceIDRequest{}
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLockByResourceIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLockByResourceIDRequest) ProtoMessage() {}

func (x *GetLockByResourceIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLockByResourceIDRequest.ProtoReflect.Descriptor instead.
func (*GetLockByResourceIDRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetLockByResourceIDRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *GetLockByResourceIDRequest) GetInstanceId() string {
	if x != nil {
		return x.InstanceId
	}
	return ""
}

var File_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto protoreflect.FileDescriptor

const file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDesc = "" +
	"\n" +
	"Byandex/cloud/marketplace/licensemanager/saas/v1/lock_service.proto\x12/yandex.cloud.marketplace.licensemanager.saas.v1\x1a\x1cgoogle/api/annotations.proto\x1a yandex/cloud/api/operation.proto\x1a5yandex/cloud/marketplace/licensemanager/v1/lock.proto\x1a&yandex/cloud/operation/operation.proto\x1a\x1dyandex/cloud/validation.proto\"g\n" +
	"\x11EnsureLockRequest\x12+\n" +
	"\x0einstance_token\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\rinstanceToken\x12%\n" +
	"\vresource_id\x18\x02 \x01(\tB\x04\xe8\xc71\x01R\n" +
	"resourceId\"-\n" +
	"\x12EnsureLockMetadata\x12\x17\n" +
	"\alock_id\x18\x01 \x01(\tR\x06lockId\"/\n" +
	"\x0eGetLockRequest\x12\x1d\n" +
	"\alock_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\x06lockId\"j\n" +
	"\x1aGetLockByResourceIDRequest\x12%\n" +
	"\vresource_id\x18\x01 \x01(\tB\x04\xe8\xc71\x01R\n" +
	"resourceId\x12%\n" +
	"\vinstance_id\x18\x02 \x01(\tB\x04\xe8\xc71\x01R\n" +
	"instanceId2\xfd\x04\n" +
	"\vLockService\x12\xdd\x01\n" +
	"\x06Ensure\x12B.yandex.cloud.marketplace.licensemanager.saas.v1.EnsureLockRequest\x1a!.yandex.cloud.operation.Operation\"l\xb2\xd2*,\n" +
	"\x12EnsureLockMetadata\x12\x16licensemanager.v1.Lock\x82\xd3\xe4\x93\x026:\x01*\"1/marketplace/license-manager/saas/v1/locks/ensure\x12\xb6\x01\n" +
	"\x03Get\x12?.yandex.cloud.marketplace.licensemanager.saas.v1.GetLockRequest\x1a0.yandex.cloud.marketplace.licensemanager.v1.Lock\"<\x82\xd3\xe4\x93\x026\x124/marketplace/license-manager/saas/v1/locks/{lock_id}\x12\xd4\x01\n" +
	"\x0fGetByResourceID\x12K.yandex.cloud.marketplace.licensemanager.saas.v1.GetLockByResourceIDRequest\x1a0.yandex.cloud.marketplace.licensemanager.v1.Lock\"B\x82\xd3\xe4\x93\x02<\x12:/marketplace/license-manager/saas/v1/locks:getByResourceIDB\x99\x01\n" +
	"3yandex.cloud.api.marketplace.licensemanager.saas.v1Zbgithub.com/yandex-cloud/go-genproto/yandex/cloud/marketplace/licensemanager/saas/v1;licensemanagerb\x06proto3"

var (
	file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescData []byte
)

func file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDesc), len(file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDesc)))
	})
	return file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDescData
}

var file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_goTypes = []any{
	(*EnsureLockRequest)(nil),          // 0: yandex.cloud.marketplace.licensemanager.saas.v1.EnsureLockRequest
	(*EnsureLockMetadata)(nil),         // 1: yandex.cloud.marketplace.licensemanager.saas.v1.EnsureLockMetadata
	(*GetLockRequest)(nil),             // 2: yandex.cloud.marketplace.licensemanager.saas.v1.GetLockRequest
	(*GetLockByResourceIDRequest)(nil), // 3: yandex.cloud.marketplace.licensemanager.saas.v1.GetLockByResourceIDRequest
	(*operation.Operation)(nil),        // 4: yandex.cloud.operation.Operation
	(*v1.Lock)(nil),                    // 5: yandex.cloud.marketplace.licensemanager.v1.Lock
}
var file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.marketplace.licensemanager.saas.v1.LockService.Ensure:input_type -> yandex.cloud.marketplace.licensemanager.saas.v1.EnsureLockRequest
	2, // 1: yandex.cloud.marketplace.licensemanager.saas.v1.LockService.Get:input_type -> yandex.cloud.marketplace.licensemanager.saas.v1.GetLockRequest
	3, // 2: yandex.cloud.marketplace.licensemanager.saas.v1.LockService.GetByResourceID:input_type -> yandex.cloud.marketplace.licensemanager.saas.v1.GetLockByResourceIDRequest
	4, // 3: yandex.cloud.marketplace.licensemanager.saas.v1.LockService.Ensure:output_type -> yandex.cloud.operation.Operation
	5, // 4: yandex.cloud.marketplace.licensemanager.saas.v1.LockService.Get:output_type -> yandex.cloud.marketplace.licensemanager.v1.Lock
	5, // 5: yandex.cloud.marketplace.licensemanager.saas.v1.LockService.GetByResourceID:output_type -> yandex.cloud.marketplace.licensemanager.v1.Lock
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_init() }
func file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_init() {
	if File_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDesc), len(file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto = out.File
	file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_goTypes = nil
	file_yandex_cloud_marketplace_licensemanager_saas_v1_lock_service_proto_depIdxs = nil
}
