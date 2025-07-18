// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/gitlab/v1/maintenance.proto

package gitlab

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type MaintenanceOperation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The description of the operation.
	Info string `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	// Delay time for the maintenance operation.
	DelayedUntil *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=delayed_until,json=delayedUntil,proto3" json:"delayed_until,omitempty"`
	// Time of the last maintenance window.
	LatestMaintenanceTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=latest_maintenance_time,json=latestMaintenanceTime,proto3" json:"latest_maintenance_time,omitempty"`
	// Time of the next maintenance window.
	NextMaintenanceWindowTime *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=next_maintenance_window_time,json=nextMaintenanceWindowTime,proto3" json:"next_maintenance_window_time,omitempty"`
	unknownFields             protoimpl.UnknownFields
	sizeCache                 protoimpl.SizeCache
}

func (x *MaintenanceOperation) Reset() {
	*x = MaintenanceOperation{}
	mi := &file_yandex_cloud_gitlab_v1_maintenance_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MaintenanceOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaintenanceOperation) ProtoMessage() {}

func (x *MaintenanceOperation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_gitlab_v1_maintenance_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaintenanceOperation.ProtoReflect.Descriptor instead.
func (*MaintenanceOperation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescGZIP(), []int{0}
}

func (x *MaintenanceOperation) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *MaintenanceOperation) GetDelayedUntil() *timestamppb.Timestamp {
	if x != nil {
		return x.DelayedUntil
	}
	return nil
}

func (x *MaintenanceOperation) GetLatestMaintenanceTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LatestMaintenanceTime
	}
	return nil
}

func (x *MaintenanceOperation) GetNextMaintenanceWindowTime() *timestamppb.Timestamp {
	if x != nil {
		return x.NextMaintenanceWindowTime
	}
	return nil
}

var File_yandex_cloud_gitlab_v1_maintenance_proto protoreflect.FileDescriptor

const file_yandex_cloud_gitlab_v1_maintenance_proto_rawDesc = "" +
	"\n" +
	"(yandex/cloud/gitlab/v1/maintenance.proto\x12\x16yandex.cloud.gitlab.v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1dyandex/cloud/validation.proto\"\xa7\x02\n" +
	"\x14MaintenanceOperation\x12\x1d\n" +
	"\x04info\x18\x01 \x01(\tB\t\x8a\xc81\x05<=256R\x04info\x12?\n" +
	"\rdelayed_until\x18\x02 \x01(\v2\x1a.google.protobuf.TimestampR\fdelayedUntil\x12R\n" +
	"\x17latest_maintenance_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\x15latestMaintenanceTime\x12[\n" +
	"\x1cnext_maintenance_window_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\x19nextMaintenanceWindowTimeBc\n" +
	"\x1ayandex.cloud.api.gitlab.v1B\x02GMZAgithub.com/yandex-cloud/go-genproto/yandex/cloud/gitlab/v1;gitlabb\x06proto3"

var (
	file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescOnce sync.Once
	file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescData []byte
)

func file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescGZIP() []byte {
	file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_gitlab_v1_maintenance_proto_rawDesc), len(file_yandex_cloud_gitlab_v1_maintenance_proto_rawDesc)))
	})
	return file_yandex_cloud_gitlab_v1_maintenance_proto_rawDescData
}

var file_yandex_cloud_gitlab_v1_maintenance_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_gitlab_v1_maintenance_proto_goTypes = []any{
	(*MaintenanceOperation)(nil),  // 0: yandex.cloud.gitlab.v1.MaintenanceOperation
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_gitlab_v1_maintenance_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.gitlab.v1.MaintenanceOperation.delayed_until:type_name -> google.protobuf.Timestamp
	1, // 1: yandex.cloud.gitlab.v1.MaintenanceOperation.latest_maintenance_time:type_name -> google.protobuf.Timestamp
	1, // 2: yandex.cloud.gitlab.v1.MaintenanceOperation.next_maintenance_window_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_gitlab_v1_maintenance_proto_init() }
func file_yandex_cloud_gitlab_v1_maintenance_proto_init() {
	if File_yandex_cloud_gitlab_v1_maintenance_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_gitlab_v1_maintenance_proto_rawDesc), len(file_yandex_cloud_gitlab_v1_maintenance_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_gitlab_v1_maintenance_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_gitlab_v1_maintenance_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_gitlab_v1_maintenance_proto_msgTypes,
	}.Build()
	File_yandex_cloud_gitlab_v1_maintenance_proto = out.File
	file_yandex_cloud_gitlab_v1_maintenance_proto_goTypes = nil
	file_yandex_cloud_gitlab_v1_maintenance_proto_depIdxs = nil
}
