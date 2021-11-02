// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: yandex/cloud/datasphere/v1/node_service.proto

package datasphere

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NodeExecutionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the folder that will be matched with Node ACL.
	FolderId string `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// ID of the Node to perform request on.
	NodeId string `protobuf:"bytes,2,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Input data for the execution.
	Input *structpb.Struct `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
}

func (x *NodeExecutionRequest) Reset() {
	*x = NodeExecutionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeExecutionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeExecutionRequest) ProtoMessage() {}

func (x *NodeExecutionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeExecutionRequest.ProtoReflect.Descriptor instead.
func (*NodeExecutionRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_node_service_proto_rawDescGZIP(), []int{0}
}

func (x *NodeExecutionRequest) GetFolderId() string {
	if x != nil {
		return x.FolderId
	}
	return ""
}

func (x *NodeExecutionRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *NodeExecutionRequest) GetInput() *structpb.Struct {
	if x != nil {
		return x.Input
	}
	return nil
}

type NodeExecutionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Result of the execution.
	Output *structpb.Struct `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *NodeExecutionResponse) Reset() {
	*x = NodeExecutionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeExecutionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeExecutionResponse) ProtoMessage() {}

func (x *NodeExecutionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeExecutionResponse.ProtoReflect.Descriptor instead.
func (*NodeExecutionResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_datasphere_v1_node_service_proto_rawDescGZIP(), []int{1}
}

func (x *NodeExecutionResponse) GetOutput() *structpb.Struct {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_yandex_cloud_datasphere_v1_node_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_datasphere_v1_node_service_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64,
	0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x1a, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x14, 0x4e, 0x6f, 0x64, 0x65, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e,
	0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x22, 0x48, 0x0a, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x32, 0xb1,
	0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xa1,
	0x01, 0x0a, 0x07, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x12, 0x30, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70,
	0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x22, 0x26, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70,
	0x68, 0x65, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2f, 0x7b, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x3a,
	0x01, 0x2a, 0x42, 0x6b, 0x0a, 0x1e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72,
	0x65, 0x2e, 0x76, 0x31, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f,
	0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72,
	0x65, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x61, 0x74, 0x61, 0x73, 0x70, 0x68, 0x65, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_datasphere_v1_node_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_datasphere_v1_node_service_proto_rawDescData = file_yandex_cloud_datasphere_v1_node_service_proto_rawDesc
)

func file_yandex_cloud_datasphere_v1_node_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_datasphere_v1_node_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_datasphere_v1_node_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_datasphere_v1_node_service_proto_rawDescData)
	})
	return file_yandex_cloud_datasphere_v1_node_service_proto_rawDescData
}

var file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_datasphere_v1_node_service_proto_goTypes = []interface{}{
	(*NodeExecutionRequest)(nil),  // 0: yandex.cloud.datasphere.v1.NodeExecutionRequest
	(*NodeExecutionResponse)(nil), // 1: yandex.cloud.datasphere.v1.NodeExecutionResponse
	(*structpb.Struct)(nil),       // 2: google.protobuf.Struct
}
var file_yandex_cloud_datasphere_v1_node_service_proto_depIdxs = []int32{
	2, // 0: yandex.cloud.datasphere.v1.NodeExecutionRequest.input:type_name -> google.protobuf.Struct
	2, // 1: yandex.cloud.datasphere.v1.NodeExecutionResponse.output:type_name -> google.protobuf.Struct
	0, // 2: yandex.cloud.datasphere.v1.NodeService.Execute:input_type -> yandex.cloud.datasphere.v1.NodeExecutionRequest
	1, // 3: yandex.cloud.datasphere.v1.NodeService.Execute:output_type -> yandex.cloud.datasphere.v1.NodeExecutionResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yandex_cloud_datasphere_v1_node_service_proto_init() }
func file_yandex_cloud_datasphere_v1_node_service_proto_init() {
	if File_yandex_cloud_datasphere_v1_node_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeExecutionRequest); i {
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
		file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeExecutionResponse); i {
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
			RawDescriptor: file_yandex_cloud_datasphere_v1_node_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_datasphere_v1_node_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_datasphere_v1_node_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_datasphere_v1_node_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_datasphere_v1_node_service_proto = out.File
	file_yandex_cloud_datasphere_v1_node_service_proto_rawDesc = nil
	file_yandex_cloud_datasphere_v1_node_service_proto_goTypes = nil
	file_yandex_cloud_datasphere_v1_node_service_proto_depIdxs = nil
}
