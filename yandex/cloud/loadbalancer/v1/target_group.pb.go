// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/loadbalancer/v1/target_group.proto

package loadbalancer

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

// A TargetGroup resource. For more information, see [Target groups and resources](/docs/load-balancer/target-resources).
type TargetGroup struct {
	// Output only. ID of the target group.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the folder that the target group belongs to.
	FolderId string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Output only. Creation timestamp in [RFC3339](https://www.ietf.org/rfc/rfc3339.txt) text format.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Name of the target group.
	// The name is unique within the folder. 3-63 characters long.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the target group. 0-256 characters long.
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	// Resource labels as `` key:value `` pairs. Мaximum of 64 per resource.
	Labels map[string]string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// ID of the region where the target group resides.
	RegionId string `protobuf:"bytes,7,opt,name=region_id,json=regionId,proto3" json:"region_id,omitempty"`
	// A list of targets in the target group.
	Targets              []*Target `protobuf:"bytes,9,rep,name=targets,proto3" json:"targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *TargetGroup) Reset()         { *m = TargetGroup{} }
func (m *TargetGroup) String() string { return proto.CompactTextString(m) }
func (*TargetGroup) ProtoMessage()    {}
func (*TargetGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb82306a182641a6, []int{0}
}

func (m *TargetGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TargetGroup.Unmarshal(m, b)
}
func (m *TargetGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TargetGroup.Marshal(b, m, deterministic)
}
func (m *TargetGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TargetGroup.Merge(m, src)
}
func (m *TargetGroup) XXX_Size() int {
	return xxx_messageInfo_TargetGroup.Size(m)
}
func (m *TargetGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_TargetGroup.DiscardUnknown(m)
}

var xxx_messageInfo_TargetGroup proto.InternalMessageInfo

func (m *TargetGroup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TargetGroup) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *TargetGroup) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *TargetGroup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TargetGroup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TargetGroup) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TargetGroup) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *TargetGroup) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

// A Target resource. For more information, see [Target groups and resources](/docs/load-balancer/target-resources).
type Target struct {
	// ID of the subnet that targets are connected to.
	// All targets in the target group must be connected to the same subnet within a single availability zone.
	SubnetId string `protobuf:"bytes,1,opt,name=subnet_id,json=subnetId,proto3" json:"subnet_id,omitempty"`
	// IP address of the target.
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb82306a182641a6, []int{1}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetSubnetId() string {
	if m != nil {
		return m.SubnetId
	}
	return ""
}

func (m *Target) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*TargetGroup)(nil), "yandex.cloud.loadbalancer.v1.TargetGroup")
	proto.RegisterMapType((map[string]string)(nil), "yandex.cloud.loadbalancer.v1.TargetGroup.LabelsEntry")
	proto.RegisterType((*Target)(nil), "yandex.cloud.loadbalancer.v1.Target")
}

func init() {
	proto.RegisterFile("yandex/cloud/loadbalancer/v1/target_group.proto", fileDescriptor_eb82306a182641a6)
}

var fileDescriptor_eb82306a182641a6 = []byte{
	// 415 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x4d, 0x8b, 0x14, 0x31,
	0x10, 0x65, 0x3e, 0x76, 0x66, 0xba, 0x1a, 0x44, 0x82, 0x87, 0x30, 0x2a, 0x0e, 0x8b, 0xc2, 0x5c,
	0x36, 0xed, 0xac, 0x2c, 0xb8, 0x7e, 0x81, 0x0b, 0x22, 0x03, 0x2e, 0x42, 0xb3, 0x27, 0x2f, 0x43,
	0xba, 0x53, 0x1b, 0x83, 0x99, 0x64, 0x48, 0xa7, 0x1b, 0xe7, 0x2f, 0x78, 0xf4, 0x57, 0xf9, 0xb3,
	0xa4, 0x93, 0x6e, 0x68, 0x2f, 0x83, 0xb7, 0x54, 0x55, 0x5e, 0xbd, 0xf7, 0xaa, 0x0a, 0xb2, 0x23,
	0x37, 0x02, 0x7f, 0x66, 0xa5, 0xb6, 0xb5, 0xc8, 0xb4, 0xe5, 0xa2, 0xe0, 0x9a, 0x9b, 0x12, 0x5d,
	0xd6, 0x6c, 0x32, 0xcf, 0x9d, 0x44, 0xbf, 0x93, 0xce, 0xd6, 0x07, 0x76, 0x70, 0xd6, 0x5b, 0xf2,
	0x24, 0x02, 0x58, 0x00, 0xb0, 0x21, 0x80, 0x35, 0x9b, 0xe5, 0x33, 0x69, 0xad, 0xd4, 0x98, 0x85,
	0xbf, 0x45, 0x7d, 0x9f, 0x79, 0xb5, 0xc7, 0xca, 0xf3, 0x7d, 0x07, 0x5f, 0x3e, 0xfd, 0x87, 0xaf,
	0xe1, 0x5a, 0x09, 0xee, 0x95, 0x35, 0xb1, 0x7c, 0xfe, 0x7b, 0x02, 0xe9, 0x5d, 0x20, 0xfd, 0xdc,
	0x72, 0x92, 0x07, 0x30, 0x56, 0x82, 0x8e, 0x56, 0xa3, 0x75, 0x92, 0x8f, 0x95, 0x20, 0x8f, 0x21,
	0xb9, 0xb7, 0x5a, 0xa0, 0xdb, 0x29, 0x41, 0xc7, 0x21, 0xbd, 0x88, 0x89, 0xad, 0x20, 0xd7, 0x00,
	0xa5, 0x43, 0xee, 0x51, 0xec, 0xb8, 0xa7, 0x93, 0xd5, 0x68, 0x9d, 0x5e, 0x2e, 0x59, 0x54, 0xc4,
	0x7a, 0x45, 0xec, 0xae, 0x57, 0x94, 0x27, 0xdd, 0xef, 0x8f, 0x9e, 0x10, 0x98, 0x1a, 0xbe, 0x47,
	0x3a, 0x0d, 0x2d, 0xc3, 0x9b, 0xac, 0x20, 0x15, 0x58, 0x95, 0x4e, 0x1d, 0x5a, 0x81, 0xf4, 0x2c,
	0x94, 0x86, 0x29, 0x72, 0x0b, 0x33, 0xcd, 0x0b, 0xd4, 0x15, 0x9d, 0xad, 0x26, 0xeb, 0xf4, 0xf2,
	0x8a, 0x9d, 0x1a, 0x0e, 0x1b, 0x18, 0x63, 0x5f, 0x02, 0xee, 0x93, 0xf1, 0xee, 0x98, 0x77, 0x4d,
	0x5a, 0x73, 0x0e, 0xa5, 0xb2, 0xa6, 0x35, 0x37, 0x8f, 0xe6, 0x62, 0x62, 0x2b, 0xc8, 0x07, 0x98,
	0xc7, 0x6d, 0x54, 0x34, 0x09, 0x64, 0xcf, 0xff, 0x87, 0x2c, 0xef, 0x41, 0xcb, 0x6b, 0x48, 0x07,
	0x9c, 0xe4, 0x21, 0x4c, 0x7e, 0xe0, 0xb1, 0x9b, 0x6c, 0xfb, 0x24, 0x8f, 0xe0, 0xac, 0xe1, 0xba,
	0xc6, 0x6e, 0xac, 0x31, 0x78, 0x33, 0x7e, 0x3d, 0x3a, 0xdf, 0xc2, 0x2c, 0x76, 0x23, 0x2f, 0x20,
	0xa9, 0xea, 0xc2, 0xa0, 0xdf, 0xf5, 0x5b, 0xb9, 0x59, 0xfc, 0xfa, 0xb3, 0x99, 0xbe, 0x7b, 0x7f,
	0xf5, 0x32, 0x5f, 0xc4, 0xd2, 0x56, 0x10, 0x0a, 0x73, 0x2e, 0x84, 0xc3, 0xaa, 0xea, 0x9a, 0xf5,
	0xe1, 0xcd, 0xd7, 0x6f, 0xb7, 0x52, 0xf9, 0xef, 0x75, 0xc1, 0x4a, 0xbb, 0xef, 0x6e, 0xef, 0x22,
	0xde, 0x82, 0xb4, 0x17, 0x12, 0x4d, 0x58, 0xd3, 0xc9, 0xa3, 0x7c, 0x3b, 0x8c, 0x8b, 0x59, 0x00,
	0xbc, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x82, 0x5e, 0x67, 0x5c, 0xc8, 0x02, 0x00, 0x00,
}
