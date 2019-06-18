// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/ai/vision/v1/classification.proto

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

type ClassAnnotation struct {
	// Properties extracted by a specified model.
	//
	// For example, if you ask to evaluate the image quality,
	// the service could return such properties as `good` and `bad`.
	Properties           []*Property `protobuf:"bytes,1,rep,name=properties,proto3" json:"properties,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ClassAnnotation) Reset()         { *m = ClassAnnotation{} }
func (m *ClassAnnotation) String() string { return proto.CompactTextString(m) }
func (*ClassAnnotation) ProtoMessage()    {}
func (*ClassAnnotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b6a3b28cb9191ec, []int{0}
}

func (m *ClassAnnotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClassAnnotation.Unmarshal(m, b)
}
func (m *ClassAnnotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClassAnnotation.Marshal(b, m, deterministic)
}
func (m *ClassAnnotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClassAnnotation.Merge(m, src)
}
func (m *ClassAnnotation) XXX_Size() int {
	return xxx_messageInfo_ClassAnnotation.Size(m)
}
func (m *ClassAnnotation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClassAnnotation.DiscardUnknown(m)
}

var xxx_messageInfo_ClassAnnotation proto.InternalMessageInfo

func (m *ClassAnnotation) GetProperties() []*Property {
	if m != nil {
		return m.Properties
	}
	return nil
}

type Property struct {
	// Property name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Probability of the property, from 0 to 1.
	Probability          float64  `protobuf:"fixed64,2,opt,name=probability,proto3" json:"probability,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Property) Reset()         { *m = Property{} }
func (m *Property) String() string { return proto.CompactTextString(m) }
func (*Property) ProtoMessage()    {}
func (*Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b6a3b28cb9191ec, []int{1}
}

func (m *Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Property.Unmarshal(m, b)
}
func (m *Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Property.Marshal(b, m, deterministic)
}
func (m *Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Property.Merge(m, src)
}
func (m *Property) XXX_Size() int {
	return xxx_messageInfo_Property.Size(m)
}
func (m *Property) XXX_DiscardUnknown() {
	xxx_messageInfo_Property.DiscardUnknown(m)
}

var xxx_messageInfo_Property proto.InternalMessageInfo

func (m *Property) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Property) GetProbability() float64 {
	if m != nil {
		return m.Probability
	}
	return 0
}

func init() {
	proto.RegisterType((*ClassAnnotation)(nil), "yandex.cloud.ai.vision.v1.ClassAnnotation")
	proto.RegisterType((*Property)(nil), "yandex.cloud.ai.vision.v1.Property")
}

func init() {
	proto.RegisterFile("yandex/cloud/ai/vision/v1/classification.proto", fileDescriptor_4b6a3b28cb9191ec)
}

var fileDescriptor_4b6a3b28cb9191ec = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x89, 0x8a, 0x68, 0x6e, 0x10, 0x32, 0xd5, 0x2d, 0x9c, 0x4b, 0x97, 0x7b, 0xe1, 0x74,
	0x74, 0x51, 0x4f, 0x9c, 0xa5, 0x83, 0x83, 0x5b, 0x92, 0x8b, 0xf5, 0x41, 0x2f, 0x2f, 0x24, 0x69,
	0xb1, 0xff, 0xbd, 0x98, 0x20, 0x74, 0xe9, 0xf6, 0xf1, 0xf1, 0xfb, 0x25, 0x7c, 0x8f, 0xc3, 0xac,
	0xfd, 0xd1, 0xfd, 0x28, 0x3b, 0xd0, 0x78, 0x54, 0x1a, 0xd5, 0x84, 0x09, 0xc9, 0xab, 0x69, 0xaf,
	0xec, 0xa0, 0x53, 0xc2, 0x2f, 0xb4, 0x3a, 0x23, 0x79, 0x08, 0x91, 0x32, 0x89, 0xdb, 0xca, 0x43,
	0xe1, 0x41, 0x23, 0x54, 0x1e, 0xa6, 0xfd, 0xf6, 0x83, 0xdf, 0x1c, 0xfe, 0x94, 0x67, 0xef, 0x29,
	0x17, 0x47, 0x1c, 0x38, 0x0f, 0x91, 0x82, 0x8b, 0x19, 0x5d, 0x6a, 0x98, 0x3c, 0x6f, 0x37, 0xf7,
	0x77, 0xb0, 0xfa, 0x04, 0xbc, 0x57, 0x78, 0xee, 0x16, 0xda, 0xf6, 0x89, 0x5f, 0xfd, 0xf7, 0x42,
	0xf0, 0x0b, 0xaf, 0x4f, 0xae, 0x61, 0x92, 0xb5, 0xd7, 0x5d, 0xc9, 0x42, 0xf2, 0x4d, 0x88, 0x64,
	0xb4, 0xc1, 0x01, 0xf3, 0xdc, 0x9c, 0x49, 0xd6, 0xb2, 0x6e, 0x59, 0xbd, 0xbc, 0x7d, 0xbe, 0xf6,
	0x98, 0xbf, 0x47, 0x03, 0x96, 0x4e, 0xaa, 0x7e, 0xbf, 0xab, 0x8b, 0x7b, 0xda, 0xf5, 0xce, 0x97,
	0x6d, 0x6a, 0xf5, 0x14, 0x8f, 0x35, 0x99, 0xcb, 0xc2, 0x3d, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0x93, 0x55, 0x94, 0x35, 0x01, 0x00, 0x00,
}