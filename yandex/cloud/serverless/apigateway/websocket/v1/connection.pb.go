// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/serverless/apigateway/websocket/v1/connection.proto

package websocket

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

type Connection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the connection.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the API Gateway.
	GatewayId string `protobuf:"bytes,2,opt,name=gateway_id,json=gatewayId,proto3" json:"gateway_id,omitempty"`
	// The information about the caller making the request to API Gateway.
	Identity *Identity `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	// The timestamp at which connection was established.
	ConnectedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=connected_at,json=connectedAt,proto3" json:"connected_at,omitempty"`
	// The timestamp at which connection was last accessed.
	LastActiveAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=last_active_at,json=lastActiveAt,proto3" json:"last_active_at,omitempty"`
}

func (x *Connection) Reset() {
	*x = Connection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Connection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Connection) ProtoMessage() {}

func (x *Connection) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Connection.ProtoReflect.Descriptor instead.
func (*Connection) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescGZIP(), []int{0}
}

func (x *Connection) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Connection) GetGatewayId() string {
	if x != nil {
		return x.GatewayId
	}
	return ""
}

func (x *Connection) GetIdentity() *Identity {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *Connection) GetConnectedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ConnectedAt
	}
	return nil
}

func (x *Connection) GetLastActiveAt() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActiveAt
	}
	return nil
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The source IP address of the caller making the request to API Gateway.
	SourceIp string `protobuf:"bytes,1,opt,name=source_ip,json=sourceIp,proto3" json:"source_ip,omitempty"`
	// The User Agent of the caller making the request to API Gateway.
	UserAgent string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent,proto3" json:"user_agent,omitempty"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescGZIP(), []int{1}
}

func (x *Identity) GetSourceIp() string {
	if x != nil {
		return x.SourceIp
	}
	return ""
}

func (x *Identity) GetUserAgent() string {
	if x != nil {
		return x.UserAgent
	}
	return ""
}

var File_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto protoreflect.FileDescriptor

var file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDesc = []byte{
	0x0a, 0x40, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x49, 0x64, 0x12, 0x55, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x39, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2e, 0x61,
	0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3d, 0x0a, 0x0c, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x40, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x41, 0x74, 0x22, 0x46, 0x0a, 0x08, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x42, 0x94, 0x01, 0x0a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x6c, 0x65,
	0x73, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x77, 0x65,
	0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x5d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x6c, 0x65, 0x73, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2f, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescOnce sync.Once
	file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescData = file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDesc
)

func file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescGZIP() []byte {
	file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescData)
	})
	return file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDescData
}

var file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_goTypes = []interface{}{
	(*Connection)(nil),            // 0: yandex.cloud.serverless.apigateway.websocket.v1.Connection
	(*Identity)(nil),              // 1: yandex.cloud.serverless.apigateway.websocket.v1.Identity
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.serverless.apigateway.websocket.v1.Connection.identity:type_name -> yandex.cloud.serverless.apigateway.websocket.v1.Identity
	2, // 1: yandex.cloud.serverless.apigateway.websocket.v1.Connection.connected_at:type_name -> google.protobuf.Timestamp
	2, // 2: yandex.cloud.serverless.apigateway.websocket.v1.Connection.last_active_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_init() }
func file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_init() {
	if File_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Connection); i {
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
		file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
			RawDescriptor: file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_msgTypes,
	}.Build()
	File_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto = out.File
	file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_rawDesc = nil
	file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_goTypes = nil
	file_yandex_cloud_serverless_apigateway_websocket_v1_connection_proto_depIdxs = nil
}
