// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/mdb/mysql/v1alpha/backup.proto

package mysql

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/validation"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A MySQL backup. For more information, see
// the [documentation](/docs/managed-mysql/concepts/backup).
type Backup struct {
	// ID of the backup.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the backup belongs to.
	FolderId  string               `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ID of the MySQL cluster that the backup was created for.
	SourceClusterId string `protobuf:"bytes,4,opt,name=source_cluster_id,json=sourceClusterId,proto3" json:"source_cluster_id,omitempty"`
	// Time when the backup operation was started.
	StartedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Backup) Reset()         { *m = Backup{} }
func (m *Backup) String() string { return proto.CompactTextString(m) }
func (*Backup) ProtoMessage()    {}
func (*Backup) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1a911ab513ac83b, []int{0}
}

func (m *Backup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Backup.Unmarshal(m, b)
}
func (m *Backup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Backup.Marshal(b, m, deterministic)
}
func (m *Backup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Backup.Merge(m, src)
}
func (m *Backup) XXX_Size() int {
	return xxx_messageInfo_Backup.Size(m)
}
func (m *Backup) XXX_DiscardUnknown() {
	xxx_messageInfo_Backup.DiscardUnknown(m)
}

var xxx_messageInfo_Backup proto.InternalMessageInfo

func (m *Backup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Backup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *Backup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Backup) GetSourceClusterId() string {
	if m != nil {
		return m.SourceClusterId
	}
	return ""
}

func (m *Backup) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*Backup)(nil), "yandex.cloud.mdb.mysql.v1alpha.Backup")
}

func init() {
	proto.RegisterFile("yandex/cloud/mdb/mysql/v1alpha/backup.proto", fileDescriptor_c1a911ab513ac83b)
}

var fileDescriptor_c1a911ab513ac83b = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x95, 0x50, 0x2a, 0x62, 0x06, 0x44, 0xc4, 0x10, 0x05, 0x01, 0x15, 0x53, 0x05, 0xaa,
	0xad, 0xc2, 0x84, 0x98, 0x1a, 0x16, 0x60, 0x8c, 0x98, 0x58, 0x22, 0xff, 0x6b, 0x6a, 0x61, 0xc7,
	0x21, 0x39, 0x57, 0xf4, 0x09, 0x79, 0x0a, 0x76, 0x1e, 0x03, 0x61, 0xa7, 0x03, 0x0b, 0x1d, 0xef,
	0xee, 0xbb, 0xdf, 0x27, 0xdd, 0xa1, 0xeb, 0x0d, 0x6d, 0x84, 0xfc, 0x20, 0x5c, 0x5b, 0x27, 0x88,
	0x11, 0x8c, 0x98, 0x4d, 0xff, 0xae, 0xc9, 0x7a, 0x4e, 0x75, 0xbb, 0xa2, 0x84, 0x51, 0xfe, 0xe6,
	0x5a, 0xdc, 0x76, 0x16, 0x6c, 0x7a, 0x1e, 0x60, 0xec, 0x61, 0x6c, 0x04, 0xc3, 0x1e, 0xc6, 0x03,
	0x9c, 0x9f, 0xfd, 0x09, 0x5b, 0x53, 0xad, 0x04, 0x05, 0x65, 0x9b, 0xb0, 0x9e, 0x5f, 0xd4, 0xd6,
	0xd6, 0x5a, 0x12, 0x5f, 0x31, 0xb7, 0x24, 0xa0, 0x8c, 0xec, 0x81, 0x9a, 0x21, 0xff, 0xf2, 0x2b,
	0x42, 0xe3, 0xc2, 0x0b, 0xd3, 0x13, 0x14, 0x2b, 0x91, 0x45, 0x93, 0x68, 0x9a, 0x14, 0xa3, 0xef,
	0xcf, 0x79, 0x54, 0xc6, 0x4a, 0xa4, 0xa7, 0x28, 0x59, 0x5a, 0x2d, 0x64, 0x57, 0x29, 0x91, 0xc5,
	0xbf, 0xc3, 0xf2, 0x20, 0x34, 0x9e, 0x44, 0x7a, 0x87, 0x10, 0xef, 0x24, 0x05, 0x29, 0x2a, 0x0a,
	0xd9, 0xde, 0x24, 0x9a, 0x1e, 0xde, 0xe4, 0x38, 0x38, 0xf1, 0xd6, 0x89, 0x5f, 0xb6, 0xce, 0x32,
	0x19, 0xe8, 0x05, 0xa4, 0x57, 0xe8, 0xb8, 0xb7, 0xae, 0xe3, 0xb2, 0xe2, 0xda, 0xf5, 0x10, 0xf2,
	0x47, 0x3e, 0xff, 0x28, 0x0c, 0x1e, 0x42, 0x3f, 0x68, 0x7a, 0xa0, 0xdd, 0xa0, 0xd9, 0xdf, 0xad,
	0x19, 0xe8, 0x05, 0x14, 0xcf, 0xaf, 0x8f, 0xb5, 0x82, 0x95, 0x63, 0x98, 0x5b, 0x43, 0xc2, 0xb1,
	0x66, 0xe1, 0x58, 0xb5, 0x9d, 0xd5, 0xb2, 0xf1, 0xeb, 0xe4, 0xff, 0x97, 0xdc, 0xfb, 0x8a, 0x8d,
	0x3d, 0x7b, 0xfb, 0x13, 0x00, 0x00, 0xff, 0xff, 0x47, 0xbf, 0x98, 0x15, 0xc1, 0x01, 0x00, 0x00,
}
