// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.12
// source: yandex/cloud/billing/v1/customer_service.proto

package billing

import (
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud/api"
	operation "github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
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

type ListCustomersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the reseller.
	ResellerId string `protobuf:"bytes,1,opt,name=reseller_id,json=resellerId,proto3" json:"reseller_id,omitempty"`
	// The maximum number of results per page to return. If the number of available
	// results is larger than [page_size],
	// the service returns a [ListCustomersResponse.next_page_token]
	// that can be used to get the next page of results in subsequent list requests.
	PageSize int64 `protobuf:"varint,3,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Page token. To get the next page of results,
	// set [page_token] to the [ListCustomersResponse.next_page_token]
	// returned by a previous list request.
	PageToken string `protobuf:"bytes,4,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListCustomersRequest) Reset() {
	*x = ListCustomersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCustomersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersRequest) ProtoMessage() {}

func (x *ListCustomersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersRequest.ProtoReflect.Descriptor instead.
func (*ListCustomersRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListCustomersRequest) GetResellerId() string {
	if x != nil {
		return x.ResellerId
	}
	return ""
}

func (x *ListCustomersRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListCustomersRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

type ListCustomersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of customers.
	Customers []*Customer `protobuf:"bytes,1,rep,name=customers,proto3" json:"customers,omitempty"`
	// This token allows you to get the next page of results for list requests. If the number of results
	// is larger than [ListCustomersRequest.page_size], use
	// [next_page_token] as the value
	// for the [ListCustomersRequest.page_token] query parameter
	// in the next list request. Each subsequent list request will have its own
	// [next_page_token] to continue paging through the results.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListCustomersResponse) Reset() {
	*x = ListCustomersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCustomersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCustomersResponse) ProtoMessage() {}

func (x *ListCustomersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCustomersResponse.ProtoReflect.Descriptor instead.
func (*ListCustomersResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListCustomersResponse) GetCustomers() []*Customer {
	if x != nil {
		return x.Customers
	}
	return nil
}

func (x *ListCustomersResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type InviteCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Billing account ID of the reseller that the customer will be associated with.
	ResellerId string `protobuf:"bytes,1,opt,name=reseller_id,json=resellerId,proto3" json:"reseller_id,omitempty"`
	// Name of the customer.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Customer email to send invitation to.
	InvitationEmail string `protobuf:"bytes,3,opt,name=invitation_email,json=invitationEmail,proto3" json:"invitation_email,omitempty"`
	// Person of the customer.
	Person *CustomerPerson `protobuf:"bytes,4,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *InviteCustomerRequest) Reset() {
	*x = InviteCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteCustomerRequest) ProtoMessage() {}

func (x *InviteCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteCustomerRequest.ProtoReflect.Descriptor instead.
func (*InviteCustomerRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{2}
}

func (x *InviteCustomerRequest) GetResellerId() string {
	if x != nil {
		return x.ResellerId
	}
	return ""
}

func (x *InviteCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InviteCustomerRequest) GetInvitationEmail() string {
	if x != nil {
		return x.InvitationEmail
	}
	return ""
}

func (x *InviteCustomerRequest) GetPerson() *CustomerPerson {
	if x != nil {
		return x.Person
	}
	return nil
}

type CreateResellerServedCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the reseller that customer will be associated with.
	//
	// Value must match either one of the two regular expressions:
	// `^([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}|[0-9a-f]{32})$`
	// or `^[a-z][-a-zA-Z0-9.]{0,48}[a-zA-Z0-9]$`.
	ResellerId string `protobuf:"bytes,1,opt,name=reseller_id,json=resellerId,proto3" json:"reseller_id,omitempty"`
	// Name of the customer.
	//
	// String length is not limited.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Person of the customer.
	Person *CustomerPerson `protobuf:"bytes,3,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *CreateResellerServedCustomerRequest) Reset() {
	*x = CreateResellerServedCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResellerServedCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResellerServedCustomerRequest) ProtoMessage() {}

func (x *CreateResellerServedCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResellerServedCustomerRequest.ProtoReflect.Descriptor instead.
func (*CreateResellerServedCustomerRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResellerServedCustomerRequest) GetResellerId() string {
	if x != nil {
		return x.ResellerId
	}
	return ""
}

func (x *CreateResellerServedCustomerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateResellerServedCustomerRequest) GetPerson() *CustomerPerson {
	if x != nil {
		return x.Person
	}
	return nil
}

type ActivateCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the customer.
	// To get the customer ID, use [CustomerService.List] request.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *ActivateCustomerRequest) Reset() {
	*x = ActivateCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivateCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivateCustomerRequest) ProtoMessage() {}

func (x *ActivateCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivateCustomerRequest.ProtoReflect.Descriptor instead.
func (*ActivateCustomerRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{4}
}

func (x *ActivateCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type SuspendCustomerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the customer.
	// To get the customer ID, use [CustomerService.List] request.
	CustomerId string `protobuf:"bytes,1,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *SuspendCustomerRequest) Reset() {
	*x = SuspendCustomerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuspendCustomerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuspendCustomerRequest) ProtoMessage() {}

func (x *SuspendCustomerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuspendCustomerRequest.ProtoReflect.Descriptor instead.
func (*SuspendCustomerRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{5}
}

func (x *SuspendCustomerRequest) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

type CustomerMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the reseller.
	ResellerId string `protobuf:"bytes,1,opt,name=reseller_id,json=resellerId,proto3" json:"reseller_id,omitempty"`
	// ID of the customer.
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId,proto3" json:"customer_id,omitempty"`
}

func (x *CustomerMetadata) Reset() {
	*x = CustomerMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerMetadata) ProtoMessage() {}

func (x *CustomerMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerMetadata.ProtoReflect.Descriptor instead.
func (*CustomerMetadata) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP(), []int{6}
}

func (x *CustomerMetadata) GetResellerId() string {
	if x != nil {
		return x.ResellerId
	}
	return ""
}

func (x *CustomerMetadata) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

var File_yandex_cloud_billing_v1_customer_service_proto protoreflect.FileDescriptor

var file_yandex_cloud_billing_v1_customer_service_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x17, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x14, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04,
	0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x27, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0xc7, 0x31, 0x06, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x30, 0x52,
	0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0x8a,
	0xc8, 0x31, 0x05, 0x3c, 0x3d, 0x31, 0x30, 0x30, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0x80, 0x01, 0x0a, 0x15, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e,
	0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xd8, 0x01, 0x0a,
	0x15, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31,
	0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x6c,
	0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x2f, 0x0a, 0x10, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52,
	0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x45, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52,
	0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0xb5, 0x01, 0x0a, 0x23, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x2d, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d,
	0x35, 0x30, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x04, 0xe8, 0xc7,
	0x31, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x42, 0x04, 0xe8, 0xc7, 0x31, 0x01, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22,
	0x48, 0x0a, 0x17, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8, 0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x16, 0x53, 0x75, 0x73,
	0x70, 0x65, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0c, 0xe8, 0xc7, 0x31, 0x01, 0x8a, 0xc8,
	0x31, 0x04, 0x3c, 0x3d, 0x35, 0x30, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72,
	0x49, 0x64, 0x22, 0x54, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x32, 0x85, 0x07, 0x0a, 0x0f, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x84, 0x01, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2d, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x12, 0xa4, 0x01, 0x0a, 0x06, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0x2e,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69,
	0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x47, 0xb2, 0xd2, 0x2a, 0x1c, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x3a, 0x01, 0x2a, 0x22, 0x1c, 0x2f, 0x62,
	0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x73, 0x3a, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x12, 0xd6, 0x01, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x12, 0x3c, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0xb2, 0xd2, 0x2a, 0x1c, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x08, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x37, 0x3a, 0x01, 0x2a, 0x22,
	0x32, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x3a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0xb5, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x12, 0x30, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0xb2, 0xd2, 0x2a, 0x1c, 0x0a, 0x10, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x08, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2e, 0x22, 0x2c, 0x2f,
	0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x7d, 0x3a, 0x61, 0x63, 0x74, 0x69, 0x76, 0x61, 0x74, 0x65, 0x12, 0xb2, 0x01, 0x0a, 0x07,
	0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x12, 0x2f, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65,
	0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x53, 0xb2, 0xd2, 0x2a,
	0x1c, 0x0a, 0x10, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x12, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x2d, 0x22, 0x2b, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x42, 0x62, 0x0a, 0x1b, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x5a,
	0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64,
	0x65, 0x78, 0x2d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_billing_v1_customer_service_proto_rawDescOnce sync.Once
	file_yandex_cloud_billing_v1_customer_service_proto_rawDescData = file_yandex_cloud_billing_v1_customer_service_proto_rawDesc
)

func file_yandex_cloud_billing_v1_customer_service_proto_rawDescGZIP() []byte {
	file_yandex_cloud_billing_v1_customer_service_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_billing_v1_customer_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_billing_v1_customer_service_proto_rawDescData)
	})
	return file_yandex_cloud_billing_v1_customer_service_proto_rawDescData
}

var file_yandex_cloud_billing_v1_customer_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_yandex_cloud_billing_v1_customer_service_proto_goTypes = []any{
	(*ListCustomersRequest)(nil),                // 0: yandex.cloud.billing.v1.ListCustomersRequest
	(*ListCustomersResponse)(nil),               // 1: yandex.cloud.billing.v1.ListCustomersResponse
	(*InviteCustomerRequest)(nil),               // 2: yandex.cloud.billing.v1.InviteCustomerRequest
	(*CreateResellerServedCustomerRequest)(nil), // 3: yandex.cloud.billing.v1.CreateResellerServedCustomerRequest
	(*ActivateCustomerRequest)(nil),             // 4: yandex.cloud.billing.v1.ActivateCustomerRequest
	(*SuspendCustomerRequest)(nil),              // 5: yandex.cloud.billing.v1.SuspendCustomerRequest
	(*CustomerMetadata)(nil),                    // 6: yandex.cloud.billing.v1.CustomerMetadata
	(*Customer)(nil),                            // 7: yandex.cloud.billing.v1.Customer
	(*CustomerPerson)(nil),                      // 8: yandex.cloud.billing.v1.CustomerPerson
	(*operation.Operation)(nil),                 // 9: yandex.cloud.operation.Operation
}
var file_yandex_cloud_billing_v1_customer_service_proto_depIdxs = []int32{
	7, // 0: yandex.cloud.billing.v1.ListCustomersResponse.customers:type_name -> yandex.cloud.billing.v1.Customer
	8, // 1: yandex.cloud.billing.v1.InviteCustomerRequest.person:type_name -> yandex.cloud.billing.v1.CustomerPerson
	8, // 2: yandex.cloud.billing.v1.CreateResellerServedCustomerRequest.person:type_name -> yandex.cloud.billing.v1.CustomerPerson
	0, // 3: yandex.cloud.billing.v1.CustomerService.List:input_type -> yandex.cloud.billing.v1.ListCustomersRequest
	2, // 4: yandex.cloud.billing.v1.CustomerService.Invite:input_type -> yandex.cloud.billing.v1.InviteCustomerRequest
	3, // 5: yandex.cloud.billing.v1.CustomerService.CreateResellerServed:input_type -> yandex.cloud.billing.v1.CreateResellerServedCustomerRequest
	4, // 6: yandex.cloud.billing.v1.CustomerService.Activate:input_type -> yandex.cloud.billing.v1.ActivateCustomerRequest
	5, // 7: yandex.cloud.billing.v1.CustomerService.Suspend:input_type -> yandex.cloud.billing.v1.SuspendCustomerRequest
	1, // 8: yandex.cloud.billing.v1.CustomerService.List:output_type -> yandex.cloud.billing.v1.ListCustomersResponse
	9, // 9: yandex.cloud.billing.v1.CustomerService.Invite:output_type -> yandex.cloud.operation.Operation
	9, // 10: yandex.cloud.billing.v1.CustomerService.CreateResellerServed:output_type -> yandex.cloud.operation.Operation
	9, // 11: yandex.cloud.billing.v1.CustomerService.Activate:output_type -> yandex.cloud.operation.Operation
	9, // 12: yandex.cloud.billing.v1.CustomerService.Suspend:output_type -> yandex.cloud.operation.Operation
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_yandex_cloud_billing_v1_customer_service_proto_init() }
func file_yandex_cloud_billing_v1_customer_service_proto_init() {
	if File_yandex_cloud_billing_v1_customer_service_proto != nil {
		return
	}
	file_yandex_cloud_billing_v1_customer_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ListCustomersRequest); i {
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
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ListCustomersResponse); i {
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
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*InviteCustomerRequest); i {
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
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateResellerServedCustomerRequest); i {
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
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ActivateCustomerRequest); i {
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
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SuspendCustomerRequest); i {
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
		file_yandex_cloud_billing_v1_customer_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*CustomerMetadata); i {
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
			RawDescriptor: file_yandex_cloud_billing_v1_customer_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_yandex_cloud_billing_v1_customer_service_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_billing_v1_customer_service_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_billing_v1_customer_service_proto_msgTypes,
	}.Build()
	File_yandex_cloud_billing_v1_customer_service_proto = out.File
	file_yandex_cloud_billing_v1_customer_service_proto_rawDesc = nil
	file_yandex_cloud_billing_v1_customer_service_proto_goTypes = nil
	file_yandex_cloud_billing_v1_customer_service_proto_depIdxs = nil
}
