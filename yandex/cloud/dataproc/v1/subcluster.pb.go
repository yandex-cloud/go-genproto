// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/dataproc/v1/subcluster.proto

package dataproc

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type Role int32

const (
	Role_ROLE_UNSPECIFIED Role = 0
	// The subcluster fulfills the master role.
	//
	// Master can run the following services, depending on the requested components:
	// * HDFS: Namenode, Secondary Namenode
	// * YARN: ResourceManager, Timeline Server
	// * HBase Master
	// * Hive: Server, Metastore, HCatalog
	// * Spark History Server
	// * Zeppelin
	// * ZooKeeper
	Role_MASTERNODE Role = 1
	// The subcluster is a DATANODE in a Yandex Data Processing cluster.
	//
	// DATANODE can run the following services, depending on the requested components:
	// * HDFS DataNode
	// * YARN NodeManager
	// * HBase RegionServer
	// * Spark libraries
	Role_DATANODE Role = 2
	// The subcluster is a COMPUTENODE in a Yandex Data Processing cluster.
	//
	// COMPUTENODE can run the following services, depending on the requested components:
	// * YARN NodeManager
	// * Spark libraries
	Role_COMPUTENODE Role = 3
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "ROLE_UNSPECIFIED",
		1: "MASTERNODE",
		2: "DATANODE",
		3: "COMPUTENODE",
	}
	Role_value = map[string]int32{
		"ROLE_UNSPECIFIED": 0,
		"MASTERNODE":       1,
		"DATANODE":         2,
		"COMPUTENODE":      3,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_dataproc_v1_subcluster_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_yandex_cloud_dataproc_v1_subcluster_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescGZIP(), []int{0}
}

type AutoscalingConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Upper limit for total instance subcluster count.
	MaxHostsCount int64 `protobuf:"varint,1,opt,name=max_hosts_count,json=maxHostsCount,proto3" json:"max_hosts_count,omitempty"`
	// Preemptible instances are stopped at least once every 24 hours, and can be stopped at any time
	// if their resources are needed by Compute.
	// For more information, see [Preemptible Virtual Machines](/docs/compute/concepts/preemptible-vm).
	Preemptible bool `protobuf:"varint,2,opt,name=preemptible,proto3" json:"preemptible,omitempty"`
	// Time in seconds allotted for averaging metrics.
	MeasurementDuration *durationpb.Duration `protobuf:"bytes,3,opt,name=measurement_duration,json=measurementDuration,proto3" json:"measurement_duration,omitempty"`
	// The warmup time of the instance in seconds. During this time,
	// traffic is sent to the instance, but instance metrics are not collected.
	WarmupDuration *durationpb.Duration `protobuf:"bytes,4,opt,name=warmup_duration,json=warmupDuration,proto3" json:"warmup_duration,omitempty"`
	// Minimum amount of time in seconds allotted for monitoring before
	// Instance Groups can reduce the number of instances in the group.
	// During this time, the group size doesn't decrease, even if the new metric values
	// indicate that it should.
	StabilizationDuration *durationpb.Duration `protobuf:"bytes,5,opt,name=stabilization_duration,json=stabilizationDuration,proto3" json:"stabilization_duration,omitempty"`
	// Defines an autoscaling rule based on the average CPU utilization of the instance group.
	CpuUtilizationTarget float64 `protobuf:"fixed64,6,opt,name=cpu_utilization_target,json=cpuUtilizationTarget,proto3" json:"cpu_utilization_target,omitempty"`
	// Timeout to gracefully decommission nodes during downscaling. In seconds. Default value: 120
	DecommissionTimeout int64 `protobuf:"varint,7,opt,name=decommission_timeout,json=decommissionTimeout,proto3" json:"decommission_timeout,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *AutoscalingConfig) Reset() {
	*x = AutoscalingConfig{}
	mi := &file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutoscalingConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoscalingConfig) ProtoMessage() {}

func (x *AutoscalingConfig) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoscalingConfig.ProtoReflect.Descriptor instead.
func (*AutoscalingConfig) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescGZIP(), []int{0}
}

func (x *AutoscalingConfig) GetMaxHostsCount() int64 {
	if x != nil {
		return x.MaxHostsCount
	}
	return 0
}

func (x *AutoscalingConfig) GetPreemptible() bool {
	if x != nil {
		return x.Preemptible
	}
	return false
}

func (x *AutoscalingConfig) GetMeasurementDuration() *durationpb.Duration {
	if x != nil {
		return x.MeasurementDuration
	}
	return nil
}

func (x *AutoscalingConfig) GetWarmupDuration() *durationpb.Duration {
	if x != nil {
		return x.WarmupDuration
	}
	return nil
}

func (x *AutoscalingConfig) GetStabilizationDuration() *durationpb.Duration {
	if x != nil {
		return x.StabilizationDuration
	}
	return nil
}

func (x *AutoscalingConfig) GetCpuUtilizationTarget() float64 {
	if x != nil {
		return x.CpuUtilizationTarget
	}
	return 0
}

func (x *AutoscalingConfig) GetDecommissionTimeout() int64 {
	if x != nil {
		return x.DecommissionTimeout
	}
	return 0
}

// A Yandex Data Processing subcluster. For details about the concept, see [documentation](/docs/data-proc/concepts/).
type Subcluster struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the subcluster. Generated at creation time.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the Yandex Data Processing cluster that the subcluster belongs to.
	ClusterId string `protobuf:"bytes,2,opt,name=cluster_id,json=clusterId,proto3" json:"cluster_id,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the subcluster. The name is unique within the cluster.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Role that is fulfilled by hosts of the subcluster.
	Role Role `protobuf:"varint,5,opt,name=role,proto3,enum=yandex.cloud.dataproc.v1.Role" json:"role,omitempty"`
	// Resources allocated for each host in the subcluster.
	Resources *Resources `protobuf:"bytes,6,opt,name=resources,proto3" json:"resources,omitempty"`
	// ID of the VPC subnet used for hosts in the subcluster.
	SubnetId string `protobuf:"bytes,7,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// Number of hosts in the subcluster.
	HostsCount int64 `protobuf:"varint,8,opt,name=hosts_count,json=hostsCount,proto3" json:"hosts_count,omitempty"`
	// Assign public ip addresses for all hosts in subcluter.
	AssignPublicIp bool `protobuf:"varint,9,opt,name=assign_public_ip,json=assignPublicIp,proto3" json:"assign_public_ip,omitempty"`
	// Configuration for instance group based subclusters
	AutoscalingConfig *AutoscalingConfig `protobuf:"bytes,10,opt,name=autoscaling_config,json=autoscalingConfig,proto3" json:"autoscaling_config,omitempty"`
	// ID of Compute Instance Group for autoscaling subclusters
	InstanceGroupId string `protobuf:"bytes,11,opt,name=instance_group_id,json=instanceGroupId,proto3" json:"instance_group_id,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *Subcluster) Reset() {
	*x = Subcluster{}
	mi := &file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Subcluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subcluster) ProtoMessage() {}

func (x *Subcluster) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subcluster.ProtoReflect.Descriptor instead.
func (*Subcluster) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescGZIP(), []int{1}
}

func (x *Subcluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subcluster) GetClusterId() string {
	if x != nil {
		return x.ClusterId
	}
	return ""
}

func (x *Subcluster) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Subcluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subcluster) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

func (x *Subcluster) GetResources() *Resources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *Subcluster) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *Subcluster) GetHostsCount() int64 {
	if x != nil {
		return x.HostsCount
	}
	return 0
}

func (x *Subcluster) GetAssignPublicIp() bool {
	if x != nil {
		return x.AssignPublicIp
	}
	return false
}

func (x *Subcluster) GetAutoscalingConfig() *AutoscalingConfig {
	if x != nil {
		return x.AutoscalingConfig
	}
	return nil
}

func (x *Subcluster) GetInstanceGroupId() string {
	if x != nil {
		return x.InstanceGroupId
	}
	return ""
}

// A Yandex Data Processing host. For details about the concept, see [documentation](/docs/data-proc/concepts/).
type Host struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the Yandex Data Processing host. The host name is assigned by Yandex Data Processing at creation time
	// and cannot be changed. The name is generated to be unique across all Yandex Data Processing
	// hosts that exist on the platform, as it defines the FQDN of the host.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// ID of the Yandex Data Processing subcluster that the host belongs to.
	SubclusterId string `protobuf:"bytes,2,opt,name=subcluster_id,json=subclusterId,proto3" json:"subcluster_id,omitempty"`
	// Status code of the aggregated health of the host.
	Health Health `protobuf:"varint,3,opt,name=health,proto3,enum=yandex.cloud.dataproc.v1.Health" json:"health,omitempty"`
	// ID of the Compute virtual machine that is used as the Yandex Data Processing host.
	ComputeInstanceId string `protobuf:"bytes,4,opt,name=compute_instance_id,json=computeInstanceId,proto3" json:"compute_instance_id,omitempty"`
	// Role of the host in the cluster.
	Role          Role `protobuf:"varint,5,opt,name=role,proto3,enum=yandex.cloud.dataproc.v1.Role" json:"role,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Host) Reset() {
	*x = Host{}
	mi := &file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Host) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Host) ProtoMessage() {}

func (x *Host) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Host.ProtoReflect.Descriptor instead.
func (*Host) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescGZIP(), []int{2}
}

func (x *Host) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Host) GetSubclusterId() string {
	if x != nil {
		return x.SubclusterId
	}
	return ""
}

func (x *Host) GetHealth() Health {
	if x != nil {
		return x.Health
	}
	return Health_HEALTH_UNKNOWN
}

func (x *Host) GetComputeInstanceId() string {
	if x != nil {
		return x.ComputeInstanceId
	}
	return ""
}

func (x *Host) GetRole() Role {
	if x != nil {
		return x.Role
	}
	return Role_ROLE_UNSPECIFIED
}

var File_yandex_cloud_dataproc_v1_subcluster_proto protoreflect.FileDescriptor

const file_yandex_cloud_dataproc_v1_subcluster_proto_rawDesc = "" +
	"\n" +
	")yandex/cloud/dataproc/v1/subcluster.proto\x12\x18yandex.cloud.dataproc.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a%yandex/cloud/dataproc/v1/common.proto\x1a\x1dyandex/cloud/validation.proto\x1a\x1egoogle/protobuf/duration.proto\"\xf4\x03\n" +
	"\x11AutoscalingConfig\x121\n" +
	"\x0fmax_hosts_count\x18\x01 \x01(\x03B\t\xfa\xc71\x051-100R\rmaxHostsCount\x12 \n" +
	"\vpreemptible\x18\x02 \x01(\bR\vpreemptible\x12\\\n" +
	"\x14measurement_duration\x18\x03 \x01(\v2\x19.google.protobuf.DurationB\x0e\xe8\xc71\x01\xfa\xc71\x061m-10mR\x13measurementDuration\x12M\n" +
	"\x0fwarmup_duration\x18\x04 \x01(\v2\x19.google.protobuf.DurationB\t\xfa\xc71\x05<=10mR\x0ewarmupDuration\x12\\\n" +
	"\x16stabilization_duration\x18\x05 \x01(\v2\x19.google.protobuf.DurationB\n" +
	"\xfa\xc71\x061m-30mR\x15stabilizationDuration\x12?\n" +
	"\x16cpu_utilization_target\x18\x06 \x01(\x01B\t\xfa\xc71\x050-100R\x14cpuUtilizationTarget\x12>\n" +
	"\x14decommission_timeout\x18\a \x01(\x03B\v\xfa\xc71\a0-86400R\x13decommissionTimeout\"\xfb\x03\n" +
	"\n" +
	"Subcluster\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n" +
	"\n" +
	"cluster_id\x18\x02 \x01(\tR\tclusterId\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12\x1c\n" +
	"\x04name\x18\x04 \x01(\tB\b\x8a\xc81\x041-63R\x04name\x122\n" +
	"\x04role\x18\x05 \x01(\x0e2\x1e.yandex.cloud.dataproc.v1.RoleR\x04role\x12A\n" +
	"\tresources\x18\x06 \x01(\v2#.yandex.cloud.dataproc.v1.ResourcesR\tresources\x12\x1b\n" +
	"\tsubnet_id\x18\a \x01(\tR\bsubnetId\x12\x1f\n" +
	"\vhosts_count\x18\b \x01(\x03R\n" +
	"hostsCount\x12(\n" +
	"\x10assign_public_ip\x18\t \x01(\bR\x0eassignPublicIp\x12Z\n" +
	"\x12autoscaling_config\x18\n" +
	" \x01(\v2+.yandex.cloud.dataproc.v1.AutoscalingConfigR\x11autoscalingConfig\x12*\n" +
	"\x11instance_group_id\x18\v \x01(\tR\x0finstanceGroupId\"\xdd\x01\n" +
	"\x04Host\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12#\n" +
	"\rsubcluster_id\x18\x02 \x01(\tR\fsubclusterId\x128\n" +
	"\x06health\x18\x03 \x01(\x0e2 .yandex.cloud.dataproc.v1.HealthR\x06health\x12.\n" +
	"\x13compute_instance_id\x18\x04 \x01(\tR\x11computeInstanceId\x122\n" +
	"\x04role\x18\x05 \x01(\x0e2\x1e.yandex.cloud.dataproc.v1.RoleR\x04role*K\n" +
	"\x04Role\x12\x14\n" +
	"\x10ROLE_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"MASTERNODE\x10\x01\x12\f\n" +
	"\bDATANODE\x10\x02\x12\x0f\n" +
	"\vCOMPUTENODE\x10\x03Be\n" +
	"\x1cyandex.cloud.api.dataproc.v1ZEgithub.com/yandex-cloud/go-genproto/yandex/cloud/dataproc/v1;dataprocb\x06proto3"

var (
	file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescOnce sync.Once
	file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescData []byte
)

func file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescGZIP() []byte {
	file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_dataproc_v1_subcluster_proto_rawDesc), len(file_yandex_cloud_dataproc_v1_subcluster_proto_rawDesc)))
	})
	return file_yandex_cloud_dataproc_v1_subcluster_proto_rawDescData
}

var file_yandex_cloud_dataproc_v1_subcluster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_yandex_cloud_dataproc_v1_subcluster_proto_goTypes = []any{
	(Role)(0),                     // 0: yandex.cloud.dataproc.v1.Role
	(*AutoscalingConfig)(nil),     // 1: yandex.cloud.dataproc.v1.AutoscalingConfig
	(*Subcluster)(nil),            // 2: yandex.cloud.dataproc.v1.Subcluster
	(*Host)(nil),                  // 3: yandex.cloud.dataproc.v1.Host
	(*durationpb.Duration)(nil),   // 4: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*Resources)(nil),             // 6: yandex.cloud.dataproc.v1.Resources
	(Health)(0),                   // 7: yandex.cloud.dataproc.v1.Health
}
var file_yandex_cloud_dataproc_v1_subcluster_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.dataproc.v1.AutoscalingConfig.measurement_duration:type_name -> google.protobuf.Duration
	4, // 1: yandex.cloud.dataproc.v1.AutoscalingConfig.warmup_duration:type_name -> google.protobuf.Duration
	4, // 2: yandex.cloud.dataproc.v1.AutoscalingConfig.stabilization_duration:type_name -> google.protobuf.Duration
	5, // 3: yandex.cloud.dataproc.v1.Subcluster.created_at:type_name -> google.protobuf.Timestamp
	0, // 4: yandex.cloud.dataproc.v1.Subcluster.role:type_name -> yandex.cloud.dataproc.v1.Role
	6, // 5: yandex.cloud.dataproc.v1.Subcluster.resources:type_name -> yandex.cloud.dataproc.v1.Resources
	1, // 6: yandex.cloud.dataproc.v1.Subcluster.autoscaling_config:type_name -> yandex.cloud.dataproc.v1.AutoscalingConfig
	7, // 7: yandex.cloud.dataproc.v1.Host.health:type_name -> yandex.cloud.dataproc.v1.Health
	0, // 8: yandex.cloud.dataproc.v1.Host.role:type_name -> yandex.cloud.dataproc.v1.Role
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_yandex_cloud_dataproc_v1_subcluster_proto_init() }
func file_yandex_cloud_dataproc_v1_subcluster_proto_init() {
	if File_yandex_cloud_dataproc_v1_subcluster_proto != nil {
		return
	}
	file_yandex_cloud_dataproc_v1_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_dataproc_v1_subcluster_proto_rawDesc), len(file_yandex_cloud_dataproc_v1_subcluster_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_dataproc_v1_subcluster_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_dataproc_v1_subcluster_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_dataproc_v1_subcluster_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_dataproc_v1_subcluster_proto_msgTypes,
	}.Build()
	File_yandex_cloud_dataproc_v1_subcluster_proto = out.File
	file_yandex_cloud_dataproc_v1_subcluster_proto_goTypes = nil
	file_yandex_cloud_dataproc_v1_subcluster_proto_depIdxs = nil
}
