// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/baremetal/v1alpha/server.proto

package baremetal

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Server status.
type Server_Status int32

const (
	// Unspecified server status.
	Server_STATUS_UNSPECIFIED Server_Status = 0
	// Server is waiting for to be allocated from the hardware pool.
	Server_PROVISIONING Server_Status = 1
	// Server is being stopped.
	Server_STOPPING Server_Status = 3
	// Server has been stopped.
	Server_STOPPED Server_Status = 4
	// Server is being started.
	Server_STARTING Server_Status = 5
	// Server is being restarted.
	Server_RESTARTING Server_Status = 6
	// Server encountered a problem and cannot operate.
	Server_ERROR Server_Status = 7
	// Server is being deleted.
	Server_DELETING Server_Status = 8
	// Server operating system is being reinstalled.
	Server_REINSTALLING Server_Status = 9
	// Server is being updated.
	Server_UPDATING Server_Status = 10
	// Server has been quarantined
	Server_QUARANTINED Server_Status = 12
	// Server is running normaly
	Server_RUNNING Server_Status = 14
)

// Enum value maps for Server_Status.
var (
	Server_Status_name = map[int32]string{
		0:  "STATUS_UNSPECIFIED",
		1:  "PROVISIONING",
		3:  "STOPPING",
		4:  "STOPPED",
		5:  "STARTING",
		6:  "RESTARTING",
		7:  "ERROR",
		8:  "DELETING",
		9:  "REINSTALLING",
		10: "UPDATING",
		12: "QUARANTINED",
		14: "RUNNING",
	}
	Server_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"PROVISIONING":       1,
		"STOPPING":           3,
		"STOPPED":            4,
		"STARTING":           5,
		"RESTARTING":         6,
		"ERROR":              7,
		"DELETING":           8,
		"REINSTALLING":       9,
		"UPDATING":           10,
		"QUARANTINED":        12,
		"RUNNING":            14,
	}
)

func (x Server_Status) Enum() *Server_Status {
	p := new(Server_Status)
	*p = x
	return p
}

func (x Server_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Server_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_enumTypes[0].Descriptor()
}

func (Server_Status) Type() protoreflect.EnumType {
	return &file_yandex_cloud_baremetal_v1alpha_server_proto_enumTypes[0]
}

func (x Server_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Server_Status.Descriptor instead.
func (Server_Status) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP(), []int{0, 0}
}

// A Server resource.
type Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the cloud that the server belongs to.
	CloudId string `protobuf:"bytes,2,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	// ID of the folder that the server belongs to.
	FolderId string `protobuf:"bytes,3,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Name of the server.
	// The name is unique within the folder.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the server.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// ID of the availability zone where the server is resides.
	ZoneId string `protobuf:"bytes,6,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	// ID of the hardware pool that the server belongs to.
	HardwarePoolId string `protobuf:"bytes,7,opt,name=hardware_pool_id,json=hardwarePoolId,proto3" json:"hardware_pool_id,omitempty"`
	// Status of the server.
	Status Server_Status `protobuf:"varint,9,opt,name=status,proto3,enum=yandex.cloud.baremetal.v1alpha.Server_Status" json:"status,omitempty"`
	// Operating system specific settings of the server. Optional, will be empty if the server is
	// provisioned without an operating system.
	OsSettings *OsSettings `protobuf:"bytes,10,opt,name=os_settings,json=osSettings,proto3" json:"os_settings,omitempty"`
	// Array of network interfaces that are attached to the instance.
	NetworkInterfaces []*NetworkInterface `protobuf:"bytes,18,rep,name=network_interfaces,json=networkInterfaces,proto3" json:"network_interfaces,omitempty"`
	// ID of the configuration that was used to create the server.
	ConfigurationId string `protobuf:"bytes,20,opt,name=configuration_id,json=configurationId,proto3" json:"configuration_id,omitempty"`
	// Array of disks that are attached to the server.
	Disks []*Disk `protobuf:"bytes,21,rep,name=disks,proto3" json:"disks,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Resource labels as `key:value` pairs.
	Labels map[string]string `protobuf:"bytes,200,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Server) Reset() {
	*x = Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Server) ProtoMessage() {}

func (x *Server) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Server.ProtoReflect.Descriptor instead.
func (*Server) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP(), []int{0}
}

func (x *Server) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Server) GetCloudId() string {
	if x != nil {
		return x.CloudId
	}
	return ""
}

func (x *Server) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *Server) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Server) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Server) GetZoneId() string {
	if x != nil {
		return x.ZoneId
	}
	return ""
}

func (x *Server) GetHardwarePoolId() string {
	if x != nil {
		return x.HardwarePoolId
	}
	return ""
}

func (x *Server) GetStatus() Server_Status {
	if x != nil {
		return x.Status
	}
	return Server_STATUS_UNSPECIFIED
}

func (x *Server) GetOsSettings() *OsSettings {
	if x != nil {
		return x.OsSettings
	}
	return nil
}

func (x *Server) GetNetworkInterfaces() []*NetworkInterface {
	if x != nil {
		return x.NetworkInterfaces
	}
	return nil
}

func (x *Server) GetConfigurationId() string {
	if x != nil {
		return x.ConfigurationId
	}
	return ""
}

func (x *Server) GetDisks() []*Disk {
	if x != nil {
		return x.Disks
	}
	return nil
}

func (x *Server) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Server) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type NetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the network interface.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// MAC address that is assigned to the network interface.
	MacAddress string `protobuf:"bytes,2,opt,name=mac_address,json=macAddress,proto3" json:"mac_address,omitempty"`
	// IPv4 address that is assigned to the server for this network interface.
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// Subnet that the network interface belongs to.
	//
	// Types that are assignable to Subnet:
	//
	//	*NetworkInterface_PrivateSubnet
	//	*NetworkInterface_PublicSubnet
	Subnet isNetworkInterface_Subnet `protobuf_oneof:"subnet"`
}

func (x *NetworkInterface) Reset() {
	*x = NetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInterface) ProtoMessage() {}

func (x *NetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInterface.ProtoReflect.Descriptor instead.
func (*NetworkInterface) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkInterface) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NetworkInterface) GetMacAddress() string {
	if x != nil {
		return x.MacAddress
	}
	return ""
}

func (x *NetworkInterface) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (m *NetworkInterface) GetSubnet() isNetworkInterface_Subnet {
	if m != nil {
		return m.Subnet
	}
	return nil
}

func (x *NetworkInterface) GetPrivateSubnet() *PrivateSubnetNetworkInterface {
	if x, ok := x.GetSubnet().(*NetworkInterface_PrivateSubnet); ok {
		return x.PrivateSubnet
	}
	return nil
}

func (x *NetworkInterface) GetPublicSubnet() *PublicSubnetNetworkInterface {
	if x, ok := x.GetSubnet().(*NetworkInterface_PublicSubnet); ok {
		return x.PublicSubnet
	}
	return nil
}

type isNetworkInterface_Subnet interface {
	isNetworkInterface_Subnet()
}

type NetworkInterface_PrivateSubnet struct {
	// Private subnet.
	PrivateSubnet *PrivateSubnetNetworkInterface `protobuf:"bytes,7,opt,name=private_subnet,json=privateSubnet,proto3,oneof"`
}

type NetworkInterface_PublicSubnet struct {
	// Public subnet.
	PublicSubnet *PublicSubnetNetworkInterface `protobuf:"bytes,8,opt,name=public_subnet,json=publicSubnet,proto3,oneof"`
}

func (*NetworkInterface_PrivateSubnet) isNetworkInterface_Subnet() {}

func (*NetworkInterface_PublicSubnet) isNetworkInterface_Subnet() {}

type PrivateSubnetNetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the private subnet.
	PrivateSubnetId string `protobuf:"bytes,1,opt,name=private_subnet_id,json=privateSubnetId,proto3" json:"private_subnet_id,omitempty"`
}

func (x *PrivateSubnetNetworkInterface) Reset() {
	*x = PrivateSubnetNetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrivateSubnetNetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrivateSubnetNetworkInterface) ProtoMessage() {}

func (x *PrivateSubnetNetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrivateSubnetNetworkInterface.ProtoReflect.Descriptor instead.
func (*PrivateSubnetNetworkInterface) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP(), []int{2}
}

func (x *PrivateSubnetNetworkInterface) GetPrivateSubnetId() string {
	if x != nil {
		return x.PrivateSubnetId
	}
	return ""
}

type PublicSubnetNetworkInterface struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the public subnet.
	//
	// A new ephemeral public subnet will be created if not specified.
	PublicSubnetId string `protobuf:"bytes,1,opt,name=public_subnet_id,json=publicSubnetId,proto3" json:"public_subnet_id,omitempty"`
}

func (x *PublicSubnetNetworkInterface) Reset() {
	*x = PublicSubnetNetworkInterface{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicSubnetNetworkInterface) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicSubnetNetworkInterface) ProtoMessage() {}

func (x *PublicSubnetNetworkInterface) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicSubnetNetworkInterface.ProtoReflect.Descriptor instead.
func (*PublicSubnetNetworkInterface) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP(), []int{3}
}

func (x *PublicSubnetNetworkInterface) GetPublicSubnetId() string {
	if x != nil {
		return x.PublicSubnetId
	}
	return ""
}

type OsSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the image that the server was created from.
	ImageId string `protobuf:"bytes,1,opt,name=image_id,json=imageId,proto3" json:"image_id,omitempty"`
	// Public SSH key of the server.
	SshPublicKey string `protobuf:"bytes,2,opt,name=ssh_public_key,json=sshPublicKey,proto3" json:"ssh_public_key,omitempty"`
	// List of storages.
	Storages []*Storage `protobuf:"bytes,3,rep,name=storages,proto3" json:"storages,omitempty"`
}

func (x *OsSettings) Reset() {
	*x = OsSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OsSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OsSettings) ProtoMessage() {}

func (x *OsSettings) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OsSettings.ProtoReflect.Descriptor instead.
func (*OsSettings) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP(), []int{4}
}

func (x *OsSettings) GetImageId() string {
	if x != nil {
		return x.ImageId
	}
	return ""
}

func (x *OsSettings) GetSshPublicKey() string {
	if x != nil {
		return x.SshPublicKey
	}
	return ""
}

func (x *OsSettings) GetStorages() []*Storage {
	if x != nil {
		return x.Storages
	}
	return nil
}

var File_yandex_cloud_baremetal_v1alpha_server_proto protoreflect.FileDescriptor

var file_yandex_cloud_baremetal_v1alpha_server_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65,
	0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x72,
	0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x64,
	0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61,
	0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x08, 0x0a, 0x06, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x7a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x10, 0x68, 0x61, 0x72,
	0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x4b, 0x0a, 0x0b, 0x6f, 0x73,
	0x5f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x4f, 0x73, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0a, 0x6f, 0x73, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x5f, 0x0a, 0x12, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x18, 0x12, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x66, 0x61, 0x63, 0x65, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x18, 0x15, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2e, 0x44, 0x69, 0x73, 0x6b, 0x52, 0x05, 0x64, 0x69, 0x73, 0x6b, 0x73, 0x12,
	0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x64, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x4b, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0xc8, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d,
	0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0xd4, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a,
	0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x53, 0x49,
	0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x4f, 0x50, 0x50,
	0x49, 0x4e, 0x47, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x4f, 0x50, 0x50, 0x45, 0x44,
	0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x05,
	0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x53, 0x54, 0x41, 0x52, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x06,
	0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x45, 0x4c, 0x45, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x08, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x45, 0x49,
	0x4e, 0x53, 0x54, 0x41, 0x4c, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x55,
	0x50, 0x44, 0x41, 0x54, 0x49, 0x4e, 0x47, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x51, 0x55, 0x41,
	0x52, 0x41, 0x4e, 0x54, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x55,
	0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x0e, 0x22, 0x04, 0x08, 0x02, 0x10, 0x02, 0x22, 0x04, 0x08,
	0x0b, 0x10, 0x0b, 0x22, 0x04, 0x08, 0x0d, 0x10, 0x0d, 0x4a, 0x04, 0x08, 0x08, 0x10, 0x09, 0x4a,
	0x04, 0x08, 0x0b, 0x10, 0x0c, 0x4a, 0x04, 0x08, 0x0c, 0x10, 0x0d, 0x4a, 0x04, 0x08, 0x0f, 0x10,
	0x12, 0x4a, 0x04, 0x08, 0x13, 0x10, 0x14, 0x4a, 0x04, 0x08, 0x16, 0x10, 0x17, 0x4a, 0x04, 0x08,
	0x17, 0x10, 0x18, 0x4a, 0x04, 0x08, 0x0d, 0x10, 0x0e, 0x4a, 0x04, 0x08, 0x0e, 0x10, 0x0f, 0x4a,
	0x04, 0x08, 0x18, 0x10, 0x64, 0x4a, 0x05, 0x08, 0x65, 0x10, 0xc8, 0x01, 0x22, 0xbf, 0x02, 0x0a,
	0x10, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x63, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x61, 0x63, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x66, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3d, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74,
	0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x63, 0x0a, 0x0d, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x3c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x42, 0x08,
	0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x07, 0x22, 0x4b,
	0x0a, 0x1d, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12,
	0x2a, 0x0a, 0x11, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x48, 0x0a, 0x1c, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x4f, 0x73, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0e, 0x73, 0x73, 0x68, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x73, 0x68, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x52, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x73, 0x42, 0x72, 0x0a, 0x22, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62,
	0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x5a, 0x4c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x3b, 0x62, 0x61, 0x72, 0x65, 0x6d, 0x65, 0x74, 0x61, 0x6c, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescOnce sync.Once
	file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescData = file_yandex_cloud_baremetal_v1alpha_server_proto_rawDesc
)

func file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescGZIP() []byte {
	file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescData)
	})
	return file_yandex_cloud_baremetal_v1alpha_server_proto_rawDescData
}

var file_yandex_cloud_baremetal_v1alpha_server_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_yandex_cloud_baremetal_v1alpha_server_proto_goTypes = []any{
	(Server_Status)(0),                    // 0: yandex.cloud.baremetal.v1alpha.Server.Status
	(*Server)(nil),                        // 1: yandex.cloud.baremetal.v1alpha.Server
	(*NetworkInterface)(nil),              // 2: yandex.cloud.baremetal.v1alpha.NetworkInterface
	(*PrivateSubnetNetworkInterface)(nil), // 3: yandex.cloud.baremetal.v1alpha.PrivateSubnetNetworkInterface
	(*PublicSubnetNetworkInterface)(nil),  // 4: yandex.cloud.baremetal.v1alpha.PublicSubnetNetworkInterface
	(*OsSettings)(nil),                    // 5: yandex.cloud.baremetal.v1alpha.OsSettings
	nil,                                   // 6: yandex.cloud.baremetal.v1alpha.Server.LabelsEntry
	(*Disk)(nil),                          // 7: yandex.cloud.baremetal.v1alpha.Disk
	(*timestamppb.Timestamp)(nil),         // 8: google.protobuf.Timestamp
	(*Storage)(nil),                       // 9: yandex.cloud.baremetal.v1alpha.Storage
}
var file_yandex_cloud_baremetal_v1alpha_server_proto_depIdxs = []int32{
	0, // 0: yandex.cloud.baremetal.v1alpha.Server.status:type_name -> yandex.cloud.baremetal.v1alpha.Server.Status
	5, // 1: yandex.cloud.baremetal.v1alpha.Server.os_settings:type_name -> yandex.cloud.baremetal.v1alpha.OsSettings
	2, // 2: yandex.cloud.baremetal.v1alpha.Server.network_interfaces:type_name -> yandex.cloud.baremetal.v1alpha.NetworkInterface
	7, // 3: yandex.cloud.baremetal.v1alpha.Server.disks:type_name -> yandex.cloud.baremetal.v1alpha.Disk
	8, // 4: yandex.cloud.baremetal.v1alpha.Server.created_at:type_name -> google.protobuf.Timestamp
	6, // 5: yandex.cloud.baremetal.v1alpha.Server.labels:type_name -> yandex.cloud.baremetal.v1alpha.Server.LabelsEntry
	3, // 6: yandex.cloud.baremetal.v1alpha.NetworkInterface.private_subnet:type_name -> yandex.cloud.baremetal.v1alpha.PrivateSubnetNetworkInterface
	4, // 7: yandex.cloud.baremetal.v1alpha.NetworkInterface.public_subnet:type_name -> yandex.cloud.baremetal.v1alpha.PublicSubnetNetworkInterface
	9, // 8: yandex.cloud.baremetal.v1alpha.OsSettings.storages:type_name -> yandex.cloud.baremetal.v1alpha.Storage
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_baremetal_v1alpha_server_proto_init() }
func file_yandex_cloud_baremetal_v1alpha_server_proto_init() {
	if File_yandex_cloud_baremetal_v1alpha_server_proto != nil {
		return
	}
	file_yandex_cloud_baremetal_v1alpha_disk_proto_init()
	file_yandex_cloud_baremetal_v1alpha_storage_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Server); i {
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
		file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*NetworkInterface); i {
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
		file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PrivateSubnetNetworkInterface); i {
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
		file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*PublicSubnetNetworkInterface); i {
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
		file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*OsSettings); i {
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
	file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes[1].OneofWrappers = []any{
		(*NetworkInterface_PrivateSubnet)(nil),
		(*NetworkInterface_PublicSubnet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_yandex_cloud_baremetal_v1alpha_server_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_baremetal_v1alpha_server_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_baremetal_v1alpha_server_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_baremetal_v1alpha_server_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_baremetal_v1alpha_server_proto_msgTypes,
	}.Build()
	File_yandex_cloud_baremetal_v1alpha_server_proto = out.File
	file_yandex_cloud_baremetal_v1alpha_server_proto_rawDesc = nil
	file_yandex_cloud_baremetal_v1alpha_server_proto_goTypes = nil
	file_yandex_cloud_baremetal_v1alpha_server_proto_depIdxs = nil
}
