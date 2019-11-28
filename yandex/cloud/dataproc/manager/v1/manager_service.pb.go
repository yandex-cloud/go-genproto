// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/dataproc/manager/v1/manager_service.proto

package dataproc_manager

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HbaseNodeInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Requests             int64    `protobuf:"varint,2,opt,name=requests,proto3" json:"requests,omitempty"`
	HeapSizeMb           int64    `protobuf:"varint,3,opt,name=heap_size_mb,json=heapSizeMb,proto3" json:"heap_size_mb,omitempty"`
	MaxHeapSizeMb        int64    `protobuf:"varint,4,opt,name=max_heap_size_mb,json=maxHeapSizeMb,proto3" json:"max_heap_size_mb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HbaseNodeInfo) Reset()         { *m = HbaseNodeInfo{} }
func (m *HbaseNodeInfo) String() string { return proto.CompactTextString(m) }
func (*HbaseNodeInfo) ProtoMessage()    {}
func (*HbaseNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{0}
}

func (m *HbaseNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HbaseNodeInfo.Unmarshal(m, b)
}
func (m *HbaseNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HbaseNodeInfo.Marshal(b, m, deterministic)
}
func (m *HbaseNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HbaseNodeInfo.Merge(m, src)
}
func (m *HbaseNodeInfo) XXX_Size() int {
	return xxx_messageInfo_HbaseNodeInfo.Size(m)
}
func (m *HbaseNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HbaseNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HbaseNodeInfo proto.InternalMessageInfo

func (m *HbaseNodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HbaseNodeInfo) GetRequests() int64 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *HbaseNodeInfo) GetHeapSizeMb() int64 {
	if m != nil {
		return m.HeapSizeMb
	}
	return 0
}

func (m *HbaseNodeInfo) GetMaxHeapSizeMb() int64 {
	if m != nil {
		return m.MaxHeapSizeMb
	}
	return 0
}

type HbaseInfo struct {
	Available            bool             `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Regions              int64            `protobuf:"varint,2,opt,name=regions,proto3" json:"regions,omitempty"`
	Requests             int64            `protobuf:"varint,3,opt,name=requests,proto3" json:"requests,omitempty"`
	AverageLoad          float64          `protobuf:"fixed64,4,opt,name=average_load,json=averageLoad,proto3" json:"average_load,omitempty"`
	LiveNodes            []*HbaseNodeInfo `protobuf:"bytes,5,rep,name=live_nodes,json=liveNodes,proto3" json:"live_nodes,omitempty"`
	DeadNodes            []*HbaseNodeInfo `protobuf:"bytes,6,rep,name=dead_nodes,json=deadNodes,proto3" json:"dead_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *HbaseInfo) Reset()         { *m = HbaseInfo{} }
func (m *HbaseInfo) String() string { return proto.CompactTextString(m) }
func (*HbaseInfo) ProtoMessage()    {}
func (*HbaseInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{1}
}

func (m *HbaseInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HbaseInfo.Unmarshal(m, b)
}
func (m *HbaseInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HbaseInfo.Marshal(b, m, deterministic)
}
func (m *HbaseInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HbaseInfo.Merge(m, src)
}
func (m *HbaseInfo) XXX_Size() int {
	return xxx_messageInfo_HbaseInfo.Size(m)
}
func (m *HbaseInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HbaseInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HbaseInfo proto.InternalMessageInfo

func (m *HbaseInfo) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *HbaseInfo) GetRegions() int64 {
	if m != nil {
		return m.Regions
	}
	return 0
}

func (m *HbaseInfo) GetRequests() int64 {
	if m != nil {
		return m.Requests
	}
	return 0
}

func (m *HbaseInfo) GetAverageLoad() float64 {
	if m != nil {
		return m.AverageLoad
	}
	return 0
}

func (m *HbaseInfo) GetLiveNodes() []*HbaseNodeInfo {
	if m != nil {
		return m.LiveNodes
	}
	return nil
}

func (m *HbaseInfo) GetDeadNodes() []*HbaseNodeInfo {
	if m != nil {
		return m.DeadNodes
	}
	return nil
}

type HDFSNodeInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Used                 int64    `protobuf:"varint,2,opt,name=used,proto3" json:"used,omitempty"`
	Remaining            int64    `protobuf:"varint,3,opt,name=remaining,proto3" json:"remaining,omitempty"`
	Capacity             int64    `protobuf:"varint,4,opt,name=capacity,proto3" json:"capacity,omitempty"`
	NumBlocks            int64    `protobuf:"varint,5,opt,name=num_blocks,json=numBlocks,proto3" json:"num_blocks,omitempty"`
	State                string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HDFSNodeInfo) Reset()         { *m = HDFSNodeInfo{} }
func (m *HDFSNodeInfo) String() string { return proto.CompactTextString(m) }
func (*HDFSNodeInfo) ProtoMessage()    {}
func (*HDFSNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{2}
}

func (m *HDFSNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HDFSNodeInfo.Unmarshal(m, b)
}
func (m *HDFSNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HDFSNodeInfo.Marshal(b, m, deterministic)
}
func (m *HDFSNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HDFSNodeInfo.Merge(m, src)
}
func (m *HDFSNodeInfo) XXX_Size() int {
	return xxx_messageInfo_HDFSNodeInfo.Size(m)
}
func (m *HDFSNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HDFSNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HDFSNodeInfo proto.InternalMessageInfo

func (m *HDFSNodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HDFSNodeInfo) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *HDFSNodeInfo) GetRemaining() int64 {
	if m != nil {
		return m.Remaining
	}
	return 0
}

func (m *HDFSNodeInfo) GetCapacity() int64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *HDFSNodeInfo) GetNumBlocks() int64 {
	if m != nil {
		return m.NumBlocks
	}
	return 0
}

func (m *HDFSNodeInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type HDFSInfo struct {
	Available               bool            `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	PercentRemaining        float64         `protobuf:"fixed64,2,opt,name=percent_remaining,json=percentRemaining,proto3" json:"percent_remaining,omitempty"`
	Used                    int64           `protobuf:"varint,3,opt,name=used,proto3" json:"used,omitempty"`
	Free                    int64           `protobuf:"varint,4,opt,name=free,proto3" json:"free,omitempty"`
	TotalBlocks             int64           `protobuf:"varint,5,opt,name=total_blocks,json=totalBlocks,proto3" json:"total_blocks,omitempty"`
	MissingBlocks           int64           `protobuf:"varint,6,opt,name=missing_blocks,json=missingBlocks,proto3" json:"missing_blocks,omitempty"`
	MissingBlocksReplicaOne int64           `protobuf:"varint,7,opt,name=missing_blocks_replica_one,json=missingBlocksReplicaOne,proto3" json:"missing_blocks_replica_one,omitempty"`
	LiveNodes               []*HDFSNodeInfo `protobuf:"bytes,8,rep,name=live_nodes,json=liveNodes,proto3" json:"live_nodes,omitempty"`
	DeadNodes               []*HDFSNodeInfo `protobuf:"bytes,9,rep,name=dead_nodes,json=deadNodes,proto3" json:"dead_nodes,omitempty"`
	DecomNodes              []*HDFSNodeInfo `protobuf:"bytes,10,rep,name=decom_nodes,json=decomNodes,proto3" json:"decom_nodes,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}        `json:"-"`
	XXX_unrecognized        []byte          `json:"-"`
	XXX_sizecache           int32           `json:"-"`
}

func (m *HDFSInfo) Reset()         { *m = HDFSInfo{} }
func (m *HDFSInfo) String() string { return proto.CompactTextString(m) }
func (*HDFSInfo) ProtoMessage()    {}
func (*HDFSInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{3}
}

func (m *HDFSInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HDFSInfo.Unmarshal(m, b)
}
func (m *HDFSInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HDFSInfo.Marshal(b, m, deterministic)
}
func (m *HDFSInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HDFSInfo.Merge(m, src)
}
func (m *HDFSInfo) XXX_Size() int {
	return xxx_messageInfo_HDFSInfo.Size(m)
}
func (m *HDFSInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HDFSInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HDFSInfo proto.InternalMessageInfo

func (m *HDFSInfo) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *HDFSInfo) GetPercentRemaining() float64 {
	if m != nil {
		return m.PercentRemaining
	}
	return 0
}

func (m *HDFSInfo) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *HDFSInfo) GetFree() int64 {
	if m != nil {
		return m.Free
	}
	return 0
}

func (m *HDFSInfo) GetTotalBlocks() int64 {
	if m != nil {
		return m.TotalBlocks
	}
	return 0
}

func (m *HDFSInfo) GetMissingBlocks() int64 {
	if m != nil {
		return m.MissingBlocks
	}
	return 0
}

func (m *HDFSInfo) GetMissingBlocksReplicaOne() int64 {
	if m != nil {
		return m.MissingBlocksReplicaOne
	}
	return 0
}

func (m *HDFSInfo) GetLiveNodes() []*HDFSNodeInfo {
	if m != nil {
		return m.LiveNodes
	}
	return nil
}

func (m *HDFSInfo) GetDeadNodes() []*HDFSNodeInfo {
	if m != nil {
		return m.DeadNodes
	}
	return nil
}

func (m *HDFSInfo) GetDecomNodes() []*HDFSNodeInfo {
	if m != nil {
		return m.DecomNodes
	}
	return nil
}

type HiveInfo struct {
	Available            bool     `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	QueriesSucceeded     int64    `protobuf:"varint,2,opt,name=queries_succeeded,json=queriesSucceeded,proto3" json:"queries_succeeded,omitempty"`
	QueriesFailed        int64    `protobuf:"varint,3,opt,name=queries_failed,json=queriesFailed,proto3" json:"queries_failed,omitempty"`
	QueriesExecuting     int64    `protobuf:"varint,4,opt,name=queries_executing,json=queriesExecuting,proto3" json:"queries_executing,omitempty"`
	SessionsOpen         int64    `protobuf:"varint,5,opt,name=sessions_open,json=sessionsOpen,proto3" json:"sessions_open,omitempty"`
	SessionsActive       int64    `protobuf:"varint,6,opt,name=sessions_active,json=sessionsActive,proto3" json:"sessions_active,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HiveInfo) Reset()         { *m = HiveInfo{} }
func (m *HiveInfo) String() string { return proto.CompactTextString(m) }
func (*HiveInfo) ProtoMessage()    {}
func (*HiveInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{4}
}

func (m *HiveInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HiveInfo.Unmarshal(m, b)
}
func (m *HiveInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HiveInfo.Marshal(b, m, deterministic)
}
func (m *HiveInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HiveInfo.Merge(m, src)
}
func (m *HiveInfo) XXX_Size() int {
	return xxx_messageInfo_HiveInfo.Size(m)
}
func (m *HiveInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_HiveInfo.DiscardUnknown(m)
}

var xxx_messageInfo_HiveInfo proto.InternalMessageInfo

func (m *HiveInfo) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *HiveInfo) GetQueriesSucceeded() int64 {
	if m != nil {
		return m.QueriesSucceeded
	}
	return 0
}

func (m *HiveInfo) GetQueriesFailed() int64 {
	if m != nil {
		return m.QueriesFailed
	}
	return 0
}

func (m *HiveInfo) GetQueriesExecuting() int64 {
	if m != nil {
		return m.QueriesExecuting
	}
	return 0
}

func (m *HiveInfo) GetSessionsOpen() int64 {
	if m != nil {
		return m.SessionsOpen
	}
	return 0
}

func (m *HiveInfo) GetSessionsActive() int64 {
	if m != nil {
		return m.SessionsActive
	}
	return 0
}

type YarnNodeInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State                string   `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	NumContainers        int64    `protobuf:"varint,3,opt,name=num_containers,json=numContainers,proto3" json:"num_containers,omitempty"`
	UsedMemoryMb         int64    `protobuf:"varint,4,opt,name=used_memory_mb,json=usedMemoryMb,proto3" json:"used_memory_mb,omitempty"`
	AvailableMemoryMb    int64    `protobuf:"varint,5,opt,name=available_memory_mb,json=availableMemoryMb,proto3" json:"available_memory_mb,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *YarnNodeInfo) Reset()         { *m = YarnNodeInfo{} }
func (m *YarnNodeInfo) String() string { return proto.CompactTextString(m) }
func (*YarnNodeInfo) ProtoMessage()    {}
func (*YarnNodeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{5}
}

func (m *YarnNodeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YarnNodeInfo.Unmarshal(m, b)
}
func (m *YarnNodeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YarnNodeInfo.Marshal(b, m, deterministic)
}
func (m *YarnNodeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YarnNodeInfo.Merge(m, src)
}
func (m *YarnNodeInfo) XXX_Size() int {
	return xxx_messageInfo_YarnNodeInfo.Size(m)
}
func (m *YarnNodeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_YarnNodeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_YarnNodeInfo proto.InternalMessageInfo

func (m *YarnNodeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *YarnNodeInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *YarnNodeInfo) GetNumContainers() int64 {
	if m != nil {
		return m.NumContainers
	}
	return 0
}

func (m *YarnNodeInfo) GetUsedMemoryMb() int64 {
	if m != nil {
		return m.UsedMemoryMb
	}
	return 0
}

func (m *YarnNodeInfo) GetAvailableMemoryMb() int64 {
	if m != nil {
		return m.AvailableMemoryMb
	}
	return 0
}

type YarnInfo struct {
	Available            bool            `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	LiveNodes            []*YarnNodeInfo `protobuf:"bytes,2,rep,name=live_nodes,json=liveNodes,proto3" json:"live_nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *YarnInfo) Reset()         { *m = YarnInfo{} }
func (m *YarnInfo) String() string { return proto.CompactTextString(m) }
func (*YarnInfo) ProtoMessage()    {}
func (*YarnInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{6}
}

func (m *YarnInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_YarnInfo.Unmarshal(m, b)
}
func (m *YarnInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_YarnInfo.Marshal(b, m, deterministic)
}
func (m *YarnInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_YarnInfo.Merge(m, src)
}
func (m *YarnInfo) XXX_Size() int {
	return xxx_messageInfo_YarnInfo.Size(m)
}
func (m *YarnInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_YarnInfo.DiscardUnknown(m)
}

var xxx_messageInfo_YarnInfo proto.InternalMessageInfo

func (m *YarnInfo) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *YarnInfo) GetLiveNodes() []*YarnNodeInfo {
	if m != nil {
		return m.LiveNodes
	}
	return nil
}

type ZookeeperInfo struct {
	Alive                bool     `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZookeeperInfo) Reset()         { *m = ZookeeperInfo{} }
func (m *ZookeeperInfo) String() string { return proto.CompactTextString(m) }
func (*ZookeeperInfo) ProtoMessage()    {}
func (*ZookeeperInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{7}
}

func (m *ZookeeperInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZookeeperInfo.Unmarshal(m, b)
}
func (m *ZookeeperInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZookeeperInfo.Marshal(b, m, deterministic)
}
func (m *ZookeeperInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZookeeperInfo.Merge(m, src)
}
func (m *ZookeeperInfo) XXX_Size() int {
	return xxx_messageInfo_ZookeeperInfo.Size(m)
}
func (m *ZookeeperInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ZookeeperInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ZookeeperInfo proto.InternalMessageInfo

func (m *ZookeeperInfo) GetAlive() bool {
	if m != nil {
		return m.Alive
	}
	return false
}

type OozieInfo struct {
	Alive                bool     `protobuf:"varint,1,opt,name=alive,proto3" json:"alive,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OozieInfo) Reset()         { *m = OozieInfo{} }
func (m *OozieInfo) String() string { return proto.CompactTextString(m) }
func (*OozieInfo) ProtoMessage()    {}
func (*OozieInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{8}
}

func (m *OozieInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OozieInfo.Unmarshal(m, b)
}
func (m *OozieInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OozieInfo.Marshal(b, m, deterministic)
}
func (m *OozieInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OozieInfo.Merge(m, src)
}
func (m *OozieInfo) XXX_Size() int {
	return xxx_messageInfo_OozieInfo.Size(m)
}
func (m *OozieInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OozieInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OozieInfo proto.InternalMessageInfo

func (m *OozieInfo) GetAlive() bool {
	if m != nil {
		return m.Alive
	}
	return false
}

type Info struct {
	Hdfs                 *HDFSInfo      `protobuf:"bytes,1,opt,name=hdfs,proto3" json:"hdfs,omitempty"`
	Yarn                 *YarnInfo      `protobuf:"bytes,2,opt,name=yarn,proto3" json:"yarn,omitempty"`
	Hive                 *HiveInfo      `protobuf:"bytes,3,opt,name=hive,proto3" json:"hive,omitempty"`
	Zookeeper            *ZookeeperInfo `protobuf:"bytes,4,opt,name=zookeeper,proto3" json:"zookeeper,omitempty"`
	Hbase                *HbaseInfo     `protobuf:"bytes,5,opt,name=hbase,proto3" json:"hbase,omitempty"`
	Oozie                *OozieInfo     `protobuf:"bytes,6,opt,name=oozie,proto3" json:"oozie,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{9}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetHdfs() *HDFSInfo {
	if m != nil {
		return m.Hdfs
	}
	return nil
}

func (m *Info) GetYarn() *YarnInfo {
	if m != nil {
		return m.Yarn
	}
	return nil
}

func (m *Info) GetHive() *HiveInfo {
	if m != nil {
		return m.Hive
	}
	return nil
}

func (m *Info) GetZookeeper() *ZookeeperInfo {
	if m != nil {
		return m.Zookeeper
	}
	return nil
}

func (m *Info) GetHbase() *HbaseInfo {
	if m != nil {
		return m.Hbase
	}
	return nil
}

func (m *Info) GetOozie() *OozieInfo {
	if m != nil {
		return m.Oozie
	}
	return nil
}

// The request message containing the host status report.
type ReportRequest struct {
	Cid                  string   `protobuf:"bytes,1,opt,name=cid,proto3" json:"cid,omitempty"`
	TopologyRevision     int64    `protobuf:"varint,2,opt,name=topology_revision,json=topologyRevision,proto3" json:"topology_revision,omitempty"`
	Info                 *Info    `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{10}
}

func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (m *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(m, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetCid() string {
	if m != nil {
		return m.Cid
	}
	return ""
}

func (m *ReportRequest) GetTopologyRevision() int64 {
	if m != nil {
		return m.TopologyRevision
	}
	return 0
}

func (m *ReportRequest) GetInfo() *Info {
	if m != nil {
		return m.Info
	}
	return nil
}

// The response message containing the agent commands to apply on host.
type ReportReply struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportReply) Reset()         { *m = ReportReply{} }
func (m *ReportReply) String() string { return proto.CompactTextString(m) }
func (*ReportReply) ProtoMessage()    {}
func (*ReportReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c143361743b1395, []int{11}
}

func (m *ReportReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportReply.Unmarshal(m, b)
}
func (m *ReportReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportReply.Marshal(b, m, deterministic)
}
func (m *ReportReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportReply.Merge(m, src)
}
func (m *ReportReply) XXX_Size() int {
	return xxx_messageInfo_ReportReply.Size(m)
}
func (m *ReportReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportReply.DiscardUnknown(m)
}

var xxx_messageInfo_ReportReply proto.InternalMessageInfo

func init() {
	proto.RegisterType((*HbaseNodeInfo)(nil), "yandex.cloud.dataproc.manager.v1.HbaseNodeInfo")
	proto.RegisterType((*HbaseInfo)(nil), "yandex.cloud.dataproc.manager.v1.HbaseInfo")
	proto.RegisterType((*HDFSNodeInfo)(nil), "yandex.cloud.dataproc.manager.v1.HDFSNodeInfo")
	proto.RegisterType((*HDFSInfo)(nil), "yandex.cloud.dataproc.manager.v1.HDFSInfo")
	proto.RegisterType((*HiveInfo)(nil), "yandex.cloud.dataproc.manager.v1.HiveInfo")
	proto.RegisterType((*YarnNodeInfo)(nil), "yandex.cloud.dataproc.manager.v1.YarnNodeInfo")
	proto.RegisterType((*YarnInfo)(nil), "yandex.cloud.dataproc.manager.v1.YarnInfo")
	proto.RegisterType((*ZookeeperInfo)(nil), "yandex.cloud.dataproc.manager.v1.ZookeeperInfo")
	proto.RegisterType((*OozieInfo)(nil), "yandex.cloud.dataproc.manager.v1.OozieInfo")
	proto.RegisterType((*Info)(nil), "yandex.cloud.dataproc.manager.v1.Info")
	proto.RegisterType((*ReportRequest)(nil), "yandex.cloud.dataproc.manager.v1.ReportRequest")
	proto.RegisterType((*ReportReply)(nil), "yandex.cloud.dataproc.manager.v1.ReportReply")
}

func init() {
	proto.RegisterFile("yandex/cloud/dataproc/manager/v1/manager_service.proto", fileDescriptor_5c143361743b1395)
}

var fileDescriptor_5c143361743b1395 = []byte{
	// 972 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0x26, 0x3f, 0x36, 0x4d, 0x5e, 0x92, 0x65, 0x3b, 0x54, 0x10, 0xad, 0x40, 0x4a, 0x4d, 0x4b,
	0x57, 0x54, 0xeb, 0xa8, 0x8b, 0xc4, 0x81, 0x4a, 0x48, 0x2d, 0xa5, 0x5a, 0x24, 0xd2, 0x95, 0xbc,
	0xe2, 0x40, 0x2f, 0xd6, 0xc4, 0x7e, 0xc9, 0x0e, 0xb5, 0x67, 0xbc, 0xfe, 0x11, 0x36, 0x2b, 0x71,
	0xe3, 0xc4, 0xbf, 0xc0, 0x0d, 0x89, 0x3f, 0x81, 0xbf, 0x8c, 0x03, 0x57, 0x34, 0xcf, 0x33, 0x4e,
	0x8c, 0xd0, 0xae, 0xb5, 0x37, 0xcf, 0xf7, 0xde, 0xf7, 0xfc, 0xcd, 0xfb, 0xde, 0xd8, 0x03, 0x5f,
	0x6e, 0xb8, 0x0c, 0xf1, 0x6a, 0x16, 0x44, 0xaa, 0x08, 0x67, 0x21, 0xcf, 0x79, 0x92, 0xaa, 0x60,
	0x16, 0x73, 0xc9, 0x57, 0x98, 0xce, 0xd6, 0xcf, 0xec, 0xa3, 0x9f, 0x61, 0xba, 0x16, 0x01, 0xba,
	0x49, 0xaa, 0x72, 0xc5, 0xa6, 0x25, 0xcf, 0x25, 0x9e, 0x6b, 0x79, 0xae, 0x49, 0x76, 0xd7, 0xcf,
	0x9c, 0xdf, 0x5a, 0x30, 0x3e, 0x5d, 0xf0, 0x0c, 0xdf, 0xa8, 0x10, 0xbf, 0x93, 0x4b, 0xc5, 0x18,
	0x74, 0x25, 0x8f, 0x71, 0xd2, 0x9a, 0xb6, 0x8e, 0x06, 0x1e, 0x3d, 0xb3, 0x43, 0xe8, 0xa7, 0x78,
	0x59, 0x60, 0x96, 0x67, 0x93, 0xf6, 0xb4, 0x75, 0xd4, 0xf1, 0xaa, 0x35, 0x9b, 0xc2, 0xe8, 0x02,
	0x79, 0xe2, 0x67, 0xe2, 0x1a, 0xfd, 0x78, 0x31, 0xe9, 0x50, 0x1c, 0x34, 0x76, 0x2e, 0xae, 0x71,
	0xbe, 0x60, 0x4f, 0xe0, 0x20, 0xe6, 0x57, 0x7e, 0x2d, 0xab, 0x4b, 0x59, 0xe3, 0x98, 0x5f, 0x9d,
	0x56, 0x89, 0xce, 0x1f, 0x6d, 0x18, 0x90, 0x18, 0x12, 0xf2, 0x31, 0x0c, 0xf8, 0x9a, 0x8b, 0x88,
	0x2f, 0xa2, 0x52, 0x4d, 0xdf, 0xdb, 0x02, 0x6c, 0x02, 0xf7, 0x52, 0x5c, 0x09, 0x25, 0xad, 0x22,
	0xbb, 0xac, 0x89, 0xed, 0xfc, 0x47, 0xec, 0x43, 0x18, 0xf1, 0x35, 0xa6, 0x7c, 0x85, 0x7e, 0xa4,
	0x78, 0x48, 0x32, 0x5a, 0xde, 0xd0, 0x60, 0xdf, 0x2b, 0x1e, 0xb2, 0x37, 0x00, 0x91, 0x58, 0xa3,
	0x2f, 0x55, 0x88, 0xd9, 0x64, 0x6f, 0xda, 0x39, 0x1a, 0x9e, 0xcc, 0xdc, 0xdb, 0x1a, 0xe9, 0xd6,
	0x9a, 0xe8, 0x0d, 0x74, 0x09, 0xbd, 0xca, 0x74, 0xbd, 0x10, 0x79, 0x68, 0xea, 0xf5, 0xee, 0x58,
	0x4f, 0x97, 0xa0, 0x7a, 0xce, 0x9f, 0x2d, 0x18, 0x9d, 0xbe, 0x7a, 0x7d, 0x7e, 0xa3, 0x61, 0x0c,
	0xba, 0x45, 0x86, 0xa1, 0x69, 0x0d, 0x3d, 0xeb, 0x7e, 0xa6, 0x18, 0x73, 0x21, 0x85, 0x5c, 0x99,
	0xc6, 0x6c, 0x01, 0xdd, 0xb5, 0x80, 0x27, 0x3c, 0x10, 0xf9, 0xc6, 0x98, 0x53, 0xad, 0xd9, 0x27,
	0x00, 0xb2, 0x88, 0xfd, 0x45, 0xa4, 0x82, 0x77, 0xba, 0x25, 0x44, 0x95, 0x45, 0xfc, 0x92, 0x00,
	0xf6, 0x00, 0xf6, 0xb2, 0x9c, 0xe7, 0x38, 0xe9, 0x91, 0x82, 0x72, 0xe1, 0xfc, 0xdd, 0x81, 0xbe,
	0xd6, 0xd9, 0xc0, 0xcb, 0xa7, 0x70, 0x3f, 0xc1, 0x34, 0x40, 0x99, 0xfb, 0x5b, 0x85, 0x6d, 0xb2,
	0xe6, 0xc0, 0x04, 0xbc, 0x4a, 0xa8, 0xdd, 0x5a, 0x67, 0x67, 0x6b, 0x0c, 0xba, 0xcb, 0x14, 0xd1,
	0x08, 0xa7, 0x67, 0x6d, 0x75, 0xae, 0x72, 0x1e, 0xd5, 0x65, 0x0f, 0x09, 0x33, 0xc2, 0x1f, 0xc3,
	0x7e, 0x2c, 0xb2, 0x4c, 0xc8, 0x95, 0x4d, 0xea, 0x99, 0xb1, 0x2c, 0x51, 0x93, 0xf6, 0x1c, 0x0e,
	0xeb, 0x69, 0x7e, 0x8a, 0x49, 0x24, 0x02, 0xee, 0x2b, 0x89, 0x93, 0x7b, 0x44, 0xf9, 0xa8, 0x46,
	0xf1, 0xca, 0xf8, 0x99, 0x44, 0x36, 0xaf, 0x8d, 0x53, 0x9f, 0xec, 0x77, 0x1b, 0xd8, 0xbf, 0xe3,
	0xf0, 0xee, 0x34, 0xcd, 0x6b, 0xd3, 0x34, 0xb8, 0x5b, 0xb9, 0x6a, 0x98, 0xd8, 0x19, 0x0c, 0x43,
	0x0c, 0x54, 0x6c, 0xea, 0xc1, 0x9d, 0xea, 0x01, 0x95, 0x28, 0xa7, 0xf3, 0x9f, 0x16, 0xf4, 0x4f,
	0xc5, 0x1a, 0x9b, 0xb9, 0x7e, 0x59, 0x60, 0x2a, 0x30, 0xf3, 0xb3, 0x22, 0x08, 0x10, 0xc3, 0x6a,
	0x60, 0x0f, 0x4c, 0xe0, 0xdc, 0xe2, 0xda, 0x2a, 0x9b, 0xbc, 0xe4, 0x22, 0xaa, 0xfc, 0x1f, 0x1b,
	0xf4, 0x35, 0x81, 0xbb, 0x35, 0xf1, 0x0a, 0x83, 0x22, 0xd7, 0x93, 0xd4, 0xad, 0xd5, 0xfc, 0xd6,
	0xe2, 0xec, 0x53, 0x18, 0x67, 0x98, 0x65, 0xfa, 0xa3, 0xe1, 0xab, 0x04, 0xa5, 0x19, 0x91, 0x91,
	0x05, 0xcf, 0x12, 0x94, 0xec, 0x09, 0xbc, 0x5f, 0x25, 0xf1, 0x20, 0x17, 0x6b, 0x34, 0x43, 0xb2,
	0x6f, 0xe1, 0x17, 0x84, 0x3a, 0x7f, 0xb5, 0x60, 0xf4, 0x23, 0x4f, 0xe5, 0x8d, 0xe7, 0xb2, 0x3a,
	0x2a, 0xed, 0x9d, 0xa3, 0xa2, 0x37, 0xa7, 0xcf, 0x57, 0xa0, 0x64, 0xce, 0x85, 0xc4, 0xd4, 0x7e,
	0xb7, 0xc6, 0xb2, 0x88, 0xbf, 0xa9, 0x40, 0xf6, 0x08, 0xf6, 0xf5, 0xb4, 0xfb, 0x31, 0xc6, 0x2a,
	0xdd, 0x6c, 0xbf, 0xa2, 0x23, 0x8d, 0xce, 0x09, 0x9c, 0x2f, 0x98, 0x0b, 0x1f, 0x54, 0x3d, 0xde,
	0x49, 0x2d, 0xf7, 0x76, 0xbf, 0x0a, 0xd9, 0x7c, 0xe7, 0x67, 0xe8, 0x6b, 0xd9, 0x0d, 0x0c, 0xab,
	0x8f, 0x72, 0xbb, 0xe9, 0xac, 0xec, 0x36, 0x65, 0x67, 0x94, 0x9d, 0xc7, 0x30, 0x7e, 0xab, 0xd4,
	0x3b, 0xc4, 0x04, 0x53, 0x7a, 0xfb, 0x03, 0xd8, 0xe3, 0x3a, 0x6c, 0xde, 0x5c, 0x2e, 0x9c, 0x87,
	0x30, 0x38, 0x53, 0xd7, 0x02, 0x6f, 0x48, 0xf9, 0xbd, 0x03, 0x5d, 0x0a, 0x7f, 0x0d, 0xdd, 0x8b,
	0x70, 0x99, 0x51, 0x74, 0x78, 0xf2, 0x79, 0xb3, 0x39, 0x26, 0x5d, 0xc4, 0xd3, 0xfc, 0x0d, 0x4f,
	0x25, 0xb9, 0xd3, 0x88, 0x6f, 0x3b, 0xe7, 0x11, 0x8f, 0xde, 0xaf, 0xd5, 0x75, 0x1a, 0xbf, 0xdf,
	0x1c, 0x15, 0x8f, 0x78, 0x6c, 0x0e, 0x83, 0x6b, 0xdb, 0x12, 0x32, 0xb7, 0xd1, 0xaf, 0xa2, 0xd6,
	0x45, 0x6f, 0x5b, 0x81, 0xbd, 0x80, 0xbd, 0x0b, 0xfd, 0x1b, 0x21, 0xf3, 0x87, 0x27, 0x4f, 0x1b,
	0xfe, 0x75, 0xa8, 0x4c, 0xc9, 0xd4, 0x25, 0x94, 0xee, 0x3e, 0x0d, 0x7d, 0xa3, 0x12, 0x95, 0x59,
	0x5e, 0xc9, 0xa4, 0x2b, 0x86, 0x87, 0x89, 0x4a, 0x73, 0xaf, 0xfc, 0x0d, 0xb3, 0x03, 0xe8, 0x04,
	0x22, 0x34, 0x07, 0x43, 0x3f, 0xea, 0x73, 0x9b, 0xab, 0x44, 0x45, 0x6a, 0xb5, 0xf1, 0x53, 0x5c,
	0x0b, 0x7d, 0xb0, 0xec, 0xb7, 0xc0, 0x06, 0x3c, 0x83, 0xb3, 0xaf, 0xa0, 0x2b, 0xe4, 0x52, 0x99,
	0x2e, 0x7f, 0x76, 0xbb, 0xa4, 0xb2, 0xc3, 0x9a, 0xe3, 0x8c, 0x61, 0x68, 0xb5, 0x24, 0xd1, 0xe6,
	0xe4, 0xd7, 0x16, 0x7c, 0xf8, 0xca, 0x30, 0xe6, 0x25, 0xe1, 0xbc, 0xbc, 0x41, 0xb1, 0x9f, 0xa0,
	0x57, 0x66, 0xb2, 0x06, 0x16, 0xd4, 0xf6, 0x77, 0x78, 0xdc, 0x9c, 0x90, 0x44, 0x1b, 0xe7, 0xbd,
	0x97, 0xbf, 0xc0, 0xa3, 0x1a, 0x83, 0x27, 0xe2, 0xff, 0x58, 0x6f, 0x7f, 0x58, 0x89, 0xfc, 0xa2,
	0x58, 0xb8, 0x81, 0x8a, 0x67, 0x25, 0xe1, 0xb8, 0xbc, 0x12, 0xae, 0xd4, 0xf1, 0x0a, 0x25, 0x5d,
	0xfa, 0x66, 0xb7, 0xdd, 0x15, 0x9f, 0x5b, 0xcc, 0x37, 0xd8, 0xa2, 0x47, 0xc4, 0x2f, 0xfe, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0x82, 0x4f, 0x91, 0x8b, 0x67, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DataprocManagerServiceClient is the client API for DataprocManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DataprocManagerServiceClient interface {
	// Sends a status report from a host
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error)
}

type dataprocManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewDataprocManagerServiceClient(cc *grpc.ClientConn) DataprocManagerServiceClient {
	return &dataprocManagerServiceClient{cc}
}

func (c *dataprocManagerServiceClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportReply, error) {
	out := new(ReportReply)
	err := c.cc.Invoke(ctx, "/yandex.cloud.dataproc.manager.v1.DataprocManagerService/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataprocManagerServiceServer is the server API for DataprocManagerService service.
type DataprocManagerServiceServer interface {
	// Sends a status report from a host
	Report(context.Context, *ReportRequest) (*ReportReply, error)
}

// UnimplementedDataprocManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDataprocManagerServiceServer struct {
}

func (*UnimplementedDataprocManagerServiceServer) Report(ctx context.Context, req *ReportRequest) (*ReportReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterDataprocManagerServiceServer(s *grpc.Server, srv DataprocManagerServiceServer) {
	s.RegisterService(&_DataprocManagerService_serviceDesc, srv)
}

func _DataprocManagerService_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataprocManagerServiceServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.dataproc.manager.v1.DataprocManagerService/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataprocManagerServiceServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DataprocManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.dataproc.manager.v1.DataprocManagerService",
	HandlerType: (*DataprocManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _DataprocManagerService_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/dataproc/manager/v1/manager_service.proto",
}
