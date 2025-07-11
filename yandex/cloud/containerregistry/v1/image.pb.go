// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/containerregistry/v1/image.proto

package containerregistry

import (
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

// An Image resource. For more information, see [Docker image](/docs/container-registry/concepts/docker-image).
type Image struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Output only. ID of the Docker image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the Docker image.
	// The name is unique within the registry.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Content-addressable identifier of the Docker image.
	Digest string `protobuf:"bytes,3,opt,name=digest,proto3" json:"digest,omitempty"`
	// Compressed size of the Docker image, specified in bytes.
	CompressedSize int64 `protobuf:"varint,4,opt,name=compressed_size,json=compressedSize,proto3" json:"compressed_size,omitempty"`
	// Configuration of the Docker image.
	Config *Blob `protobuf:"bytes,5,opt,name=config,proto3" json:"config,omitempty"`
	// Layers of the Docker image.
	Layers []*Blob `protobuf:"bytes,6,rep,name=layers,proto3" json:"layers,omitempty"`
	// Tags of the Docker image.
	//
	// Each tag is unique within the repository.
	Tags []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	// Output only. Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Image) Reset() {
	*x = Image{}
	mi := &file_yandex_cloud_containerregistry_v1_image_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_containerregistry_v1_image_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_containerregistry_v1_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Image) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Image) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *Image) GetCompressedSize() int64 {
	if x != nil {
		return x.CompressedSize
	}
	return 0
}

func (x *Image) GetConfig() *Blob {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *Image) GetLayers() []*Blob {
	if x != nil {
		return x.Layers
	}
	return nil
}

func (x *Image) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Image) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_yandex_cloud_containerregistry_v1_image_proto protoreflect.FileDescriptor

const file_yandex_cloud_containerregistry_v1_image_proto_rawDesc = "" +
	"\n" +
	"-yandex/cloud/containerregistry/v1/image.proto\x12!yandex.cloud.containerregistry.v1\x1a,yandex/cloud/containerregistry/v1/blob.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xbd\x02\n" +
	"\x05Image\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x16\n" +
	"\x06digest\x18\x03 \x01(\tR\x06digest\x12'\n" +
	"\x0fcompressed_size\x18\x04 \x01(\x03R\x0ecompressedSize\x12?\n" +
	"\x06config\x18\x05 \x01(\v2'.yandex.cloud.containerregistry.v1.BlobR\x06config\x12?\n" +
	"\x06layers\x18\x06 \x03(\v2'.yandex.cloud.containerregistry.v1.BlobR\x06layers\x12\x12\n" +
	"\x04tags\x18\a \x03(\tR\x04tags\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAtB\x80\x01\n" +
	"%yandex.cloud.api.containerregistry.v1ZWgithub.com/yandex-cloud/go-genproto/yandex/cloud/containerregistry/v1;containerregistryb\x06proto3"

var (
	file_yandex_cloud_containerregistry_v1_image_proto_rawDescOnce sync.Once
	file_yandex_cloud_containerregistry_v1_image_proto_rawDescData []byte
)

func file_yandex_cloud_containerregistry_v1_image_proto_rawDescGZIP() []byte {
	file_yandex_cloud_containerregistry_v1_image_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_containerregistry_v1_image_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_containerregistry_v1_image_proto_rawDesc), len(file_yandex_cloud_containerregistry_v1_image_proto_rawDesc)))
	})
	return file_yandex_cloud_containerregistry_v1_image_proto_rawDescData
}

var file_yandex_cloud_containerregistry_v1_image_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_containerregistry_v1_image_proto_goTypes = []any{
	(*Image)(nil),                 // 0: yandex.cloud.containerregistry.v1.Image
	(*Blob)(nil),                  // 1: yandex.cloud.containerregistry.v1.Blob
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_yandex_cloud_containerregistry_v1_image_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.containerregistry.v1.Image.config:type_name -> yandex.cloud.containerregistry.v1.Blob
	1, // 1: yandex.cloud.containerregistry.v1.Image.layers:type_name -> yandex.cloud.containerregistry.v1.Blob
	2, // 2: yandex.cloud.containerregistry.v1.Image.created_at:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_containerregistry_v1_image_proto_init() }
func file_yandex_cloud_containerregistry_v1_image_proto_init() {
	if File_yandex_cloud_containerregistry_v1_image_proto != nil {
		return
	}
	file_yandex_cloud_containerregistry_v1_blob_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_containerregistry_v1_image_proto_rawDesc), len(file_yandex_cloud_containerregistry_v1_image_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_containerregistry_v1_image_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_containerregistry_v1_image_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_containerregistry_v1_image_proto_msgTypes,
	}.Build()
	File_yandex_cloud_containerregistry_v1_image_proto = out.File
	file_yandex_cloud_containerregistry_v1_image_proto_goTypes = nil
	file_yandex_cloud_containerregistry_v1_image_proto_depIdxs = nil
}
