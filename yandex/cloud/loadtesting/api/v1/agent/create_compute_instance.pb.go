// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/api/v1/agent/create_compute_instance.proto

package agent

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/compute/v1"
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

type CreateComputeInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Resource labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,4,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the availability zone where the instance resides.
	// To get a list of available zones, use the [yandex.cloud.compute.v1.ZoneService.List] request
	ZoneId string `protobuf:"bytes,5,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// Computing resources of the instance, such as the amount of memory and number of cores.
	// To get a list of available values, see [Levels of core performance](/docs/compute/concepts/performance-levels).
	ResourcesSpec *v1.ResourcesSpec `protobuf:"bytes,7,opt,name=resources_spec,json=resourcesSpec,proto3" json:"resources_spec,omitempty"`
	// The metadata `key:value` pairs that will be assigned to this instance. This includes custom metadata and predefined keys.
	// The total size of all keys and values must be less than 512 KB.
	//
	// Values are free-form strings, and only have meaning as interpreted by the programs which configure the instance.
	// The values must be 256 KB or less.
	//
	// For example, you may use the metadata in order to provide your public SSH key to the instance.
	// For more information, see [Metadata](/docs/compute/concepts/vm-metadata).
	Metadata map[string]string `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Boot disk to attach to the instance.
	BootDiskSpec *v1.AttachedDiskSpec `protobuf:"bytes,9,opt,name=boot_disk_spec,json=bootDiskSpec,proto3" json:"boot_disk_spec,omitempty"`
	// Network configuration for the instance. Specifies how the network interface is configured
	// to interact with other services on the internal network and on the internet.
	// Currently only one network interface is supported per instance.
	NetworkInterfaceSpecs []*v1.NetworkInterfaceSpec `protobuf:"bytes,11,rep,name=network_interface_specs,json=networkInterfaceSpecs,proto3" json:"network_interface_specs,omitempty"`
	// ID of the service account to use for [authentication inside the instance](/docs/compute/operations/vm-connect/auth-inside-vm).
	// To get the service account ID, use a [yandex.cloud.iam.v1.ServiceAccountService.List] request.
	ServiceAccountId string `protobuf:"bytes,14,opt,name=service_account_id,json=serviceAccountId,proto3" json:"service_account_id,omitempty"`
	// ID of the [Compute VM platform](/docs/compute/concepts/vm-platforms) on which the agent will be created.
	// Default value: "standard-v2"
	PlatformId string `protobuf:"bytes,15,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
}

func (x *CreateComputeInstance) Reset() {
	*x = CreateComputeInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateComputeInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateComputeInstance) ProtoMessage() {}

func (x *CreateComputeInstance) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateComputeInstance.ProtoReflect.Descriptor instead.
func (*CreateComputeInstance) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescGZIP(), []int{0}
}

func (x *CreateComputeInstance) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateComputeInstance) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *CreateComputeInstance) GetResourcesSpec() *v1.ResourcesSpec {
	if x != nil {
		return x.ResourcesSpec
	}
	return nil
}

func (x *CreateComputeInstance) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *CreateComputeInstance) GetBootDiskSpec() *v1.AttachedDiskSpec {
	if x != nil {
		return x.BootDiskSpec
	}
	return nil
}

func (x *CreateComputeInstance) GetNetworkInterfaceSpecs() []*v1.NetworkInterfaceSpec {
	if x != nil {
		return x.NetworkInterfaceSpecs
	}
	return nil
}

func (x *CreateComputeInstance) GetServiceAccountId() string {
	if x != nil {
		return x.ServiceAccountId
	}
	return ""
}

func (x *CreateComputeInstance) GetPlatformId() string {
	if x != nil {
		return x.PlatformId
	}
	return ""
}

var File_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDesc = []byte{
	0x0a, 0x43, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x1a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75,
	0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc1, 0x06, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0xa5, 0x01, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x43, 0xf2, 0xc7, 0x31, 0x0f, 0x5b, 0x2d, 0x5f, 0x2e, 0x2f, 0x5c, 0x40, 0x30, 0x2d, 0x39,
	0x61, 0x2d, 0x7a, 0x5d, 0x2a, 0x82, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x36, 0x34, 0x8a, 0xc8, 0x31,
	0x04, 0x3c, 0x3d, 0x36, 0x33, 0xb2, 0xc8, 0x31, 0x1c, 0x12, 0x14, 0x5b, 0x61, 0x2d, 0x7a, 0x5d,
	0x5b, 0x2d, 0x5f, 0x2e, 0x2f, 0x5c, 0x40, 0x30, 0x2d, 0x39, 0x61, 0x2d, 0x7a, 0x5d, 0x2a, 0x1a,
	0x04, 0x31, 0x2d, 0x36, 0x33, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x25, 0x0a,
	0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c,
	0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x06, 0x7a, 0x6f,
	0x6e, 0x65, 0x49, 0x64, 0x12, 0x53, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x75, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x66, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74,
	0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x55, 0x0a, 0x0e, 0x62, 0x6f, 0x6f, 0x74, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x65, 0x64, 0x44, 0x69, 0x73, 0x6b,
	0x53, 0x70, 0x65, 0x63, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x74,
	0x44, 0x69, 0x73, 0x6b, 0x53, 0x70, 0x65, 0x63, 0x12, 0x6c, 0x0a, 0x17, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x66, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x42, 0x05, 0x82, 0xc8, 0x31, 0x01, 0x31, 0x52,
	0x15, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x53, 0x70, 0x65, 0x63, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x49, 0x64, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x4a, 0x04, 0x08,
	0x0a, 0x10, 0x0b, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0e, 0x4a, 0x04, 0x08, 0x06, 0x10, 0x07, 0x42,
	0x7c, 0x0a, 0x29, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5a, 0x4f, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescData = file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDesc
)

func file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDescData
}

var file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_goTypes = []any{
	(*CreateComputeInstance)(nil),   // 0: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance
	nil,                             // 1: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.LabelsEntry
	nil,                             // 2: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.MetadataEntry
	(*v1.ResourcesSpec)(nil),        // 3: yandex.cloud.compute.v1.ResourcesSpec
	(*v1.AttachedDiskSpec)(nil),     // 4: yandex.cloud.compute.v1.AttachedDiskSpec
	(*v1.NetworkInterfaceSpec)(nil), // 5: yandex.cloud.compute.v1.NetworkInterfaceSpec
}
var file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.labels:type_name -> yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.LabelsEntry
	3, // 1: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.resources_spec:type_name -> yandex.cloud.compute.v1.ResourcesSpec
	2, // 2: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.metadata:type_name -> yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.MetadataEntry
	4, // 3: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.boot_disk_spec:type_name -> yandex.cloud.compute.v1.AttachedDiskSpec
	5, // 4: yandex.cloud.loadtesting.api.v1.agent.CreateComputeInstance.network_interface_specs:type_name -> yandex.cloud.compute.v1.NetworkInterfaceSpec
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_init() }
func file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_init() {
	if File_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateComputeInstance); i {
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
			RawDescriptor: file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto = out.File
	file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_goTypes = nil
	file_yandex_cloud_loadtesting_api_v1_agent_create_compute_instance_proto_depIdxs = nil
}
