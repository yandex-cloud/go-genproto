// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/ai/vision/v1/primitives.proto

package vision

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Polygon struct {
	// The bounding polygon vertices.
	Vertices             []*Vertex `protobuf:"bytes,1,rep,name=vertices,proto3" json:"vertices,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Polygon) Reset()         { *m = Polygon{} }
func (m *Polygon) String() string { return proto.CompactTextString(m) }
func (*Polygon) ProtoMessage()    {}
func (*Polygon) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f9f7b59bd69f434, []int{0}
}

func (m *Polygon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Polygon.Unmarshal(m, b)
}
func (m *Polygon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Polygon.Marshal(b, m, deterministic)
}
func (m *Polygon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Polygon.Merge(m, src)
}
func (m *Polygon) XXX_Size() int {
	return xxx_messageInfo_Polygon.Size(m)
}
func (m *Polygon) XXX_DiscardUnknown() {
	xxx_messageInfo_Polygon.DiscardUnknown(m)
}

var xxx_messageInfo_Polygon proto.InternalMessageInfo

func (m *Polygon) GetVertices() []*Vertex {
	if m != nil {
		return m.Vertices
	}
	return nil
}

type Vertex struct {
	// X coordinate in pixels.
	X int64 `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	// Y coordinate in pixels.
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Vertex) Reset()         { *m = Vertex{} }
func (m *Vertex) String() string { return proto.CompactTextString(m) }
func (*Vertex) ProtoMessage()    {}
func (*Vertex) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f9f7b59bd69f434, []int{1}
}

func (m *Vertex) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Vertex.Unmarshal(m, b)
}
func (m *Vertex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Vertex.Marshal(b, m, deterministic)
}
func (m *Vertex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Vertex.Merge(m, src)
}
func (m *Vertex) XXX_Size() int {
	return xxx_messageInfo_Vertex.Size(m)
}
func (m *Vertex) XXX_DiscardUnknown() {
	xxx_messageInfo_Vertex.DiscardUnknown(m)
}

var xxx_messageInfo_Vertex proto.InternalMessageInfo

func (m *Vertex) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *Vertex) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*Polygon)(nil), "yandex.cloud.ai.vision.v1.Polygon")
	proto.RegisterType((*Vertex)(nil), "yandex.cloud.ai.vision.v1.Vertex")
}

func init() {
	proto.RegisterFile("yandex/cloud/ai/vision/v1/primitives.proto", fileDescriptor_9f9f7b59bd69f434)
}

var fileDescriptor_9f9f7b59bd69f434 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xaa, 0x4c, 0xcc, 0x4b,
	0x49, 0xad, 0xd0, 0x4f, 0xce, 0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0xcc, 0xd4, 0x2f, 0xcb, 0x2c, 0xce,
	0xcc, 0xcf, 0xd3, 0x2f, 0x33, 0xd4, 0x2f, 0x28, 0xca, 0xcc, 0xcd, 0x2c, 0xc9, 0x2c, 0x4b, 0x2d,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x84, 0xa8, 0xd5, 0x03, 0xab, 0xd5, 0x4b, 0xcc,
	0xd4, 0x83, 0xa8, 0xd5, 0x2b, 0x33, 0x54, 0xf2, 0xe0, 0x62, 0x0f, 0xc8, 0xcf, 0xa9, 0x4c, 0xcf,
	0xcf, 0x13, 0xb2, 0xe5, 0xe2, 0x28, 0x4b, 0x2d, 0x2a, 0xc9, 0x4c, 0x4e, 0x2d, 0x96, 0x60, 0x54,
	0x60, 0xd6, 0xe0, 0x36, 0x52, 0xd4, 0xc3, 0xa9, 0x51, 0x2f, 0x2c, 0xb5, 0xa8, 0x24, 0xb5, 0x22,
	0x08, 0xae, 0x45, 0x49, 0x85, 0x8b, 0x0d, 0x22, 0x26, 0xc4, 0xc3, 0xc5, 0x58, 0x21, 0xc1, 0xa8,
	0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0x08, 0xe6, 0x55, 0x4a, 0x30, 0x41, 0x78, 0x95, 0x4e, 0x6e, 0x51,
	0x2e, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9, 0xf9, 0xb9, 0xfa, 0x10, 0xe3, 0x75, 0x21,
	0x7e, 0x48, 0xcf, 0xd7, 0x4d, 0x4f, 0xcd, 0x03, 0xbb, 0x58, 0x1f, 0xa7, 0xe7, 0xac, 0x21, 0xac,
	0x24, 0x36, 0xb0, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x43, 0xdd, 0xb7, 0x07,
	0x01, 0x00, 0x00,
}
