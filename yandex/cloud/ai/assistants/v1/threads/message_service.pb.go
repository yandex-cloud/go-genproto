// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/threads/message_service.proto

package threads

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
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

// Request to create a new message in a specific thread.
type CreateMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thread to which the message will be added.
	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	// Author of the message, containing details about the message's creator.
	// If not provided, the default author ID specified in the corresponding thread will be used.
	Author *Author `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	// Set of key-value pairs to label the message.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Content of the message.
	Content *MessageContent `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *CreateMessageRequest) Reset() {
	*x = CreateMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMessageRequest) ProtoMessage() {}

func (x *CreateMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMessageRequest.ProtoReflect.Descriptor instead.
func (*CreateMessageRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMessageRequest) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *CreateMessageRequest) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *CreateMessageRequest) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *CreateMessageRequest) GetContent() *MessageContent {
	if x != nil {
		return x.Content
	}
	return nil
}

// Request message for retrieving a message from a thread.
type GetMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thread that contains the message.
	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	// ID of the message to retrieve.
	MessageId string `protobuf:"bytes,2,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
}

func (x *GetMessageRequest) Reset() {
	*x = GetMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMessageRequest) ProtoMessage() {}

func (x *GetMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMessageRequest.ProtoReflect.Descriptor instead.
func (*GetMessageRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetMessageRequest) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *GetMessageRequest) GetMessageId() string {
	if x != nil {
		return x.MessageId
	}
	return ""
}

// Request message for listing messages in a specific thread.
type ListMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the thread whose messages will be listed.
	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
}

func (x *ListMessagesRequest) Reset() {
	*x = ListMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMessagesRequest) ProtoMessage() {}

func (x *ListMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMessagesRequest.ProtoReflect.Descriptor instead.
func (*ListMessagesRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListMessagesRequest) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

var File_yandex_cloud_ai_assistants_v1_threads_message_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x73, 0x1a, 0x33, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x02, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x12, 0x45, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x5f, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x55, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x5b, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x09,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x13, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x21, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x49, 0x64, 0x32, 0xe4, 0x03, 0x0a, 0x0e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x99, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x12, 0x3b, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69,
	0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x3a, 0x01, 0x2a, 0x22, 0x17, 0x2f, 0x61, 0x73, 0x73, 0x69,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x9d, 0x01, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x38, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x26, 0x12, 0x24, 0x2f, 0x61,
	0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x7b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69,
	0x64, 0x7d, 0x12, 0x95, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3a, 0x2e, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65,
	0x61, 0x64, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12,
	0x17, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x30, 0x01, 0x42, 0x7e, 0x0a, 0x29, 0x79, 0x61,
	0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61,
	0x69, 0x2e, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x61, 0x73, 0x73,
	0x69, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x3b, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescData = file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDesc
)

func file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescData)
	})
	return file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_goTypes = []any{
	(*CreateMessageRequest)(nil), // 0: yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest
	(*GetMessageRequest)(nil),    // 1: yandex.cloud.ai.assistants.v1.threads.GetMessageRequest
	(*ListMessagesRequest)(nil),  // 2: yandex.cloud.ai.assistants.v1.threads.ListMessagesRequest
	nil,                          // 3: yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest.LabelsEntry
	(*Author)(nil),               // 4: yandex.cloud.ai.assistants.v1.threads.Author
	(*MessageContent)(nil),       // 5: yandex.cloud.ai.assistants.v1.threads.MessageContent
	(*Message)(nil),              // 6: yandex.cloud.ai.assistants.v1.threads.Message
}
var file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_depIdxs = []int32{
	4, // 0: yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest.author:type_name -> yandex.cloud.ai.assistants.v1.threads.Author
	3, // 1: yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest.labels:type_name -> yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest.LabelsEntry
	5, // 2: yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest.content:type_name -> yandex.cloud.ai.assistants.v1.threads.MessageContent
	0, // 3: yandex.cloud.ai.assistants.v1.threads.MessageService.Create:input_type -> yandex.cloud.ai.assistants.v1.threads.CreateMessageRequest
	1, // 4: yandex.cloud.ai.assistants.v1.threads.MessageService.Get:input_type -> yandex.cloud.ai.assistants.v1.threads.GetMessageRequest
	2, // 5: yandex.cloud.ai.assistants.v1.threads.MessageService.List:input_type -> yandex.cloud.ai.assistants.v1.threads.ListMessagesRequest
	6, // 6: yandex.cloud.ai.assistants.v1.threads.MessageService.Create:output_type -> yandex.cloud.ai.assistants.v1.threads.Message
	6, // 7: yandex.cloud.ai.assistants.v1.threads.MessageService.Get:output_type -> yandex.cloud.ai.assistants.v1.threads.Message
	6, // 8: yandex.cloud.ai.assistants.v1.threads.MessageService.List:output_type -> yandex.cloud.ai.assistants.v1.threads.Message
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_init() }
func file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_threads_message_service_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_threads_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateMessageRequest); i {
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
		file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetMessageRequest); i {
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
		file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ListMessagesRequest); i {
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
			RawDescriptor: file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_threads_message_service_proto = out.File
	file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_rawDesc = nil
	file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_threads_message_service_proto_depIdxs = nil
}
