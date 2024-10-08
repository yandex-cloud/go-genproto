// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/loadtesting/agent/v1/trail_service.proto

package agent

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateTrailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ComputeInstanceId string   `protobuf:"bytes,1,opt,name=compute_instance_id,json=computeInstanceId,proto3" json:"compute_instance_id,omitempty"`
	Data              []*Trail `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	JobId             string   `protobuf:"bytes,3,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	AgentInstanceId   string   `protobuf:"bytes,4,opt,name=agent_instance_id,json=agentInstanceId,proto3" json:"agent_instance_id,omitempty"`
}

func (x *CreateTrailRequest) Reset() {
	*x = CreateTrailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrailRequest) ProtoMessage() {}

func (x *CreateTrailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrailRequest.ProtoReflect.Descriptor instead.
func (*CreateTrailRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTrailRequest) GetComputeInstanceId() string {
	if x != nil {
		return x.ComputeInstanceId
	}
	return ""
}

func (x *CreateTrailRequest) GetData() []*Trail {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *CreateTrailRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *CreateTrailRequest) GetAgentInstanceId() string {
	if x != nil {
		return x.AgentInstanceId
	}
	return ""
}

type Trail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Overall       int64              `protobuf:"varint,1,opt,name=overall,proto3" json:"overall,omitempty"`
	CaseId        string             `protobuf:"bytes,2,opt,name=case_id,json=caseId,proto3" json:"case_id,omitempty"`
	Time          string             `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Reqps         int64              `protobuf:"varint,4,opt,name=reqps,proto3" json:"reqps,omitempty"`
	Resps         int64              `protobuf:"varint,5,opt,name=resps,proto3" json:"resps,omitempty"`
	Expect        float64            `protobuf:"fixed64,6,opt,name=expect,proto3" json:"expect,omitempty"`
	Input         int64              `protobuf:"varint,7,opt,name=input,proto3" json:"input,omitempty"`
	Output        int64              `protobuf:"varint,8,opt,name=output,proto3" json:"output,omitempty"`
	ConnectTime   float64            `protobuf:"fixed64,9,opt,name=connect_time,json=connectTime,proto3" json:"connect_time,omitempty"`
	SendTime      float64            `protobuf:"fixed64,10,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`
	Latency       float64            `protobuf:"fixed64,11,opt,name=latency,proto3" json:"latency,omitempty"`
	ReceiveTime   float64            `protobuf:"fixed64,12,opt,name=receive_time,json=receiveTime,proto3" json:"receive_time,omitempty"`
	Threads       int64              `protobuf:"varint,13,opt,name=threads,proto3" json:"threads,omitempty"`
	Q50           float64            `protobuf:"fixed64,14,opt,name=q50,proto3" json:"q50,omitempty"`
	Q75           float64            `protobuf:"fixed64,15,opt,name=q75,proto3" json:"q75,omitempty"`
	Q80           float64            `protobuf:"fixed64,16,opt,name=q80,proto3" json:"q80,omitempty"`
	Q85           float64            `protobuf:"fixed64,17,opt,name=q85,proto3" json:"q85,omitempty"`
	Q90           float64            `protobuf:"fixed64,18,opt,name=q90,proto3" json:"q90,omitempty"`
	Q95           float64            `protobuf:"fixed64,19,opt,name=q95,proto3" json:"q95,omitempty"`
	Q98           float64            `protobuf:"fixed64,20,opt,name=q98,proto3" json:"q98,omitempty"`
	Q99           float64            `protobuf:"fixed64,21,opt,name=q99,proto3" json:"q99,omitempty"`
	Q100          float64            `protobuf:"fixed64,22,opt,name=q100,proto3" json:"q100,omitempty"`
	HttpCodes     []*Trail_Codes     `protobuf:"bytes,23,rep,name=http_codes,json=httpCodes,proto3" json:"http_codes,omitempty"`
	NetCodes      []*Trail_Codes     `protobuf:"bytes,24,rep,name=net_codes,json=netCodes,proto3" json:"net_codes,omitempty"`
	TimeIntervals []*Trail_Intervals `protobuf:"bytes,25,rep,name=time_intervals,json=timeIntervals,proto3" json:"time_intervals,omitempty"`
}

func (x *Trail) Reset() {
	*x = Trail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trail) ProtoMessage() {}

func (x *Trail) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trail.ProtoReflect.Descriptor instead.
func (*Trail) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescGZIP(), []int{1}
}

func (x *Trail) GetOverall() int64 {
	if x != nil {
		return x.Overall
	}
	return 0
}

func (x *Trail) GetCaseId() string {
	if x != nil {
		return x.CaseId
	}
	return ""
}

func (x *Trail) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *Trail) GetReqps() int64 {
	if x != nil {
		return x.Reqps
	}
	return 0
}

func (x *Trail) GetResps() int64 {
	if x != nil {
		return x.Resps
	}
	return 0
}

func (x *Trail) GetExpect() float64 {
	if x != nil {
		return x.Expect
	}
	return 0
}

func (x *Trail) GetInput() int64 {
	if x != nil {
		return x.Input
	}
	return 0
}

func (x *Trail) GetOutput() int64 {
	if x != nil {
		return x.Output
	}
	return 0
}

func (x *Trail) GetConnectTime() float64 {
	if x != nil {
		return x.ConnectTime
	}
	return 0
}

func (x *Trail) GetSendTime() float64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *Trail) GetLatency() float64 {
	if x != nil {
		return x.Latency
	}
	return 0
}

func (x *Trail) GetReceiveTime() float64 {
	if x != nil {
		return x.ReceiveTime
	}
	return 0
}

func (x *Trail) GetThreads() int64 {
	if x != nil {
		return x.Threads
	}
	return 0
}

func (x *Trail) GetQ50() float64 {
	if x != nil {
		return x.Q50
	}
	return 0
}

func (x *Trail) GetQ75() float64 {
	if x != nil {
		return x.Q75
	}
	return 0
}

func (x *Trail) GetQ80() float64 {
	if x != nil {
		return x.Q80
	}
	return 0
}

func (x *Trail) GetQ85() float64 {
	if x != nil {
		return x.Q85
	}
	return 0
}

func (x *Trail) GetQ90() float64 {
	if x != nil {
		return x.Q90
	}
	return 0
}

func (x *Trail) GetQ95() float64 {
	if x != nil {
		return x.Q95
	}
	return 0
}

func (x *Trail) GetQ98() float64 {
	if x != nil {
		return x.Q98
	}
	return 0
}

func (x *Trail) GetQ99() float64 {
	if x != nil {
		return x.Q99
	}
	return 0
}

func (x *Trail) GetQ100() float64 {
	if x != nil {
		return x.Q100
	}
	return 0
}

func (x *Trail) GetHttpCodes() []*Trail_Codes {
	if x != nil {
		return x.HttpCodes
	}
	return nil
}

func (x *Trail) GetNetCodes() []*Trail_Codes {
	if x != nil {
		return x.NetCodes
	}
	return nil
}

func (x *Trail) GetTimeIntervals() []*Trail_Intervals {
	if x != nil {
		return x.TimeIntervals
	}
	return nil
}

type CreateTrailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrailId string `protobuf:"bytes,1,opt,name=trail_id,json=trailId,proto3" json:"trail_id,omitempty"`
	Code    int64  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *CreateTrailResponse) Reset() {
	*x = CreateTrailResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTrailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTrailResponse) ProtoMessage() {}

func (x *CreateTrailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTrailResponse.ProtoReflect.Descriptor instead.
func (*CreateTrailResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTrailResponse) GetTrailId() string {
	if x != nil {
		return x.TrailId
	}
	return ""
}

func (x *CreateTrailResponse) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Trail_Codes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  int64 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Trail_Codes) Reset() {
	*x = Trail_Codes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trail_Codes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trail_Codes) ProtoMessage() {}

func (x *Trail_Codes) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trail_Codes.ProtoReflect.Descriptor instead.
func (*Trail_Codes) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Trail_Codes) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Trail_Codes) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type Trail_Intervals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To    float64 `protobuf:"fixed64,1,opt,name=to,proto3" json:"to,omitempty"`
	Count int64   `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *Trail_Intervals) Reset() {
	*x = Trail_Intervals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Trail_Intervals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trail_Intervals) ProtoMessage() {}

func (x *Trail_Intervals) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trail_Intervals.ProtoReflect.Descriptor instead.
func (*Trail_Intervals) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Trail_Intervals) GetTo() float64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *Trail_Intervals) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_yandex_cloud_loadtesting_agent_v1_trail_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDesc = []byte{
	0x0a, 0x35, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6c,
	0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2e, 0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x63, 0x6f,
	0x6d, 0x70, 0x75, 0x74, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x3c, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a,
	0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a,
	0x6f, 0x62, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x64,
	0x22, 0xd8, 0x06, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x76,
	0x65, 0x72, 0x61, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x76, 0x65,
	0x72, 0x61, 0x6c, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x71, 0x70, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x72, 0x65, 0x71, 0x70, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x73, 0x70, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x72, 0x65, 0x73, 0x70, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x65,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x75, 0x74,
	0x70, 0x75, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x6c, 0x61, 0x74, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x21, 0x0a,
	0x0c, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x35,
	0x30, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x71, 0x35, 0x30, 0x12, 0x10, 0x0a, 0x03,
	0x71, 0x37, 0x35, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x71, 0x37, 0x35, 0x12, 0x10,
	0x0a, 0x03, 0x71, 0x38, 0x30, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x71, 0x38, 0x30,
	0x12, 0x10, 0x0a, 0x03, 0x71, 0x38, 0x35, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x71,
	0x38, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x39, 0x30, 0x18, 0x12, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x03, 0x71, 0x39, 0x30, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x39, 0x35, 0x18, 0x13, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x03, 0x71, 0x39, 0x35, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x39, 0x38, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x03, 0x71, 0x39, 0x38, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x39, 0x39, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03, 0x71, 0x39, 0x39, 0x12, 0x12, 0x0a, 0x04, 0x71, 0x31,
	0x30, 0x30, 0x18, 0x16, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x71, 0x31, 0x30, 0x30, 0x12, 0x4d,
	0x0a, 0x0a, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x17, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x73, 0x52, 0x09, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x4b, 0x0a,
	0x09, 0x6e, 0x65, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x18, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x73,
	0x52, 0x08, 0x6e, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x0e, 0x74, 0x69,
	0x6d, 0x65, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x18, 0x19, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x73, 0x52, 0x0d, 0x74, 0x69, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x73, 0x1a, 0x31, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x31, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x61, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x44, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x32, 0xb1, 0x01, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0xa0, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x35, 0x2e,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61,
	0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x72, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x74, 0x0a, 0x25, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x4b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x6c, 0x6f, 0x61, 0x64, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescData = file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDesc
)

func file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescData)
	})
	return file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDescData
}

var file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_goTypes = []any{
	(*CreateTrailRequest)(nil),  // 0: yandex.cloud.loadtesting.agent.v1.CreateTrailRequest
	(*Trail)(nil),               // 1: yandex.cloud.loadtesting.agent.v1.Trail
	(*CreateTrailResponse)(nil), // 2: yandex.cloud.loadtesting.agent.v1.CreateTrailResponse
	(*Trail_Codes)(nil),         // 3: yandex.cloud.loadtesting.agent.v1.Trail.Codes
	(*Trail_Intervals)(nil),     // 4: yandex.cloud.loadtesting.agent.v1.Trail.Intervals
}
var file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.loadtesting.agent.v1.CreateTrailRequest.data:type_name -> yandex.cloud.loadtesting.agent.v1.Trail
	3, // 1: yandex.cloud.loadtesting.agent.v1.Trail.http_codes:type_name -> yandex.cloud.loadtesting.agent.v1.Trail.Codes
	3, // 2: yandex.cloud.loadtesting.agent.v1.Trail.net_codes:type_name -> yandex.cloud.loadtesting.agent.v1.Trail.Codes
	4, // 3: yandex.cloud.loadtesting.agent.v1.Trail.time_intervals:type_name -> yandex.cloud.loadtesting.agent.v1.Trail.Intervals
	0, // 4: yandex.cloud.loadtesting.agent.v1.TrailService.Create:input_type -> yandex.cloud.loadtesting.agent.v1.CreateTrailRequest
	2, // 5: yandex.cloud.loadtesting.agent.v1.TrailService.Create:output_type -> yandex.cloud.loadtesting.agent.v1.CreateTrailResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_init() }
func file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_init() {
	if File_yandex_cloud_loadtesting_agent_v1_trail_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTrailRequest); i {
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
		file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Trail); i {
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
		file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTrailResponse); i {
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
		file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Trail_Codes); i {
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
		file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*Trail_Intervals); i {
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
			RawDescriptor: file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_loadtesting_agent_v1_trail_service_proto = out.File
	file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_rawDesc = nil
	file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_goTypes = nil
	file_yandex_cloud_loadtesting_agent_v1_trail_service_proto_depIdxs = nil
}
