// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/ai/vision/v1/face_detection.proto

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

type FaceAnnotation struct {
	// An array of detected faces for the specified image.
	Faces                []*Face  `protobuf:"bytes,1,rep,name=faces,proto3" json:"faces,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FaceAnnotation) Reset()         { *m = FaceAnnotation{} }
func (m *FaceAnnotation) String() string { return proto.CompactTextString(m) }
func (*FaceAnnotation) ProtoMessage()    {}
func (*FaceAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_88a95cd8b6371094, []int{0}
}

func (m *FaceAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FaceAnnotation.Unmarshal(m, b)
}
func (m *FaceAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FaceAnnotation.Marshal(b, m, deterministic)
}
func (m *FaceAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FaceAnnotation.Merge(m, src)
}
func (m *FaceAnnotation) XXX_Size() int {
	return xxx_messageInfo_FaceAnnotation.Size(m)
}
func (m *FaceAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_FaceAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_FaceAnnotation proto.InternalMessageInfo

func (m *FaceAnnotation) GetFaces() []*Face {
	if m != nil {
		return m.Faces
	}
	return nil
}

type Face struct {
	// Area on the image where the face is located.
	BoundingBox          *Polygon `protobuf:"bytes,1,opt,name=bounding_box,json=boundingBox,proto3" json:"bounding_box,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Face) Reset()         { *m = Face{} }
func (m *Face) String() string { return proto.CompactTextString(m) }
func (*Face) ProtoMessage()    {}
func (*Face) Descriptor() ([]byte, []int) {
	return fileDescriptor_88a95cd8b6371094, []int{1}
}

func (m *Face) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Face.Unmarshal(m, b)
}
func (m *Face) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Face.Marshal(b, m, deterministic)
}
func (m *Face) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Face.Merge(m, src)
}
func (m *Face) XXX_Size() int {
	return xxx_messageInfo_Face.Size(m)
}
func (m *Face) XXX_DiscardUnknown() {
	xxx_messageInfo_Face.DiscardUnknown(m)
}

var xxx_messageInfo_Face proto.InternalMessageInfo

func (m *Face) GetBoundingBox() *Polygon {
	if m != nil {
		return m.BoundingBox
	}
	return nil
}

func init() {
	proto.RegisterType((*FaceAnnotation)(nil), "yandex.cloud.ai.vision.v1.FaceAnnotation")
	proto.RegisterType((*Face)(nil), "yandex.cloud.ai.vision.v1.Face")
}

func init() {
	proto.RegisterFile("yandex/cloud/ai/vision/v1/face_detection.proto", fileDescriptor_88a95cd8b6371094)
}

var fileDescriptor_88a95cd8b6371094 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0xbd, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0x39, 0xfc, 0x28, 0xf6, 0xc4, 0xe2, 0xaa, 0x68, 0x63, 0xb8, 0x2a, 0x08, 0x99, 0x25,
	0x11, 0x2b, 0x2b, 0x83, 0xc6, 0x4a, 0x90, 0x94, 0x36, 0x61, 0xbf, 0x5c, 0x07, 0x92, 0x99, 0x70,
	0xb7, 0xb7, 0xdc, 0xfd, 0xf7, 0xb2, 0xb7, 0x58, 0x5e, 0xba, 0x5d, 0xf8, 0xbd, 0xdf, 0x9b, 0x27,
	0x60, 0x50, 0x64, 0x5d, 0x2f, 0xcd, 0x81, 0x3b, 0x2b, 0x15, 0xca, 0x88, 0x2d, 0x32, 0xc9, 0xb8,
	0x92, 0x3f, 0xca, 0xb8, 0xbd, 0x75, 0xc1, 0x99, 0x80, 0x4c, 0x70, 0x6a, 0x38, 0x70, 0x75, 0x97,
	0x79, 0x18, 0x79, 0x50, 0x08, 0x99, 0x87, 0xb8, 0xba, 0x7f, 0x9c, 0x56, 0x9d, 0x1a, 0x3c, 0x62,
	0xc0, 0xe8, 0xda, 0xac, 0xa9, 0x3f, 0xc4, 0xed, 0x56, 0x19, 0xf7, 0x4a, 0xc4, 0x41, 0x25, 0x7d,
	0xf5, 0x2c, 0xae, 0x52, 0x61, 0x3b, 0x2b, 0xe6, 0x17, 0x8b, 0x72, 0xfd, 0x00, 0x93, 0x45, 0x90,
	0x92, 0xbb, 0x4c, 0xd7, 0x9f, 0xe2, 0x32, 0x7d, 0xab, 0x77, 0x71, 0xa3, 0xb9, 0x23, 0x8b, 0xe4,
	0xf7, 0x9a, 0xfb, 0x59, 0x31, 0x2f, 0x16, 0xe5, 0xba, 0x3e, 0x63, 0xf9, 0xe2, 0xc3, 0xe0, 0x99,
	0x76, 0xe5, 0x7f, 0x6e, 0xc3, 0xfd, 0x66, 0xfb, 0xfd, 0xe6, 0x31, 0xfc, 0x76, 0x1a, 0x0c, 0x1f,
	0x65, 0x0e, 0x2f, 0xf3, 0x20, 0xcf, 0x4b, 0xef, 0x68, 0x3c, 0x5f, 0x4e, 0x2e, 0x7d, 0xc9, 0x2f,
	0x7d, 0x3d, 0x72, 0x4f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x1a, 0xe1, 0xfa, 0x3d, 0x5f, 0x01,
	0x00, 0x00,
}