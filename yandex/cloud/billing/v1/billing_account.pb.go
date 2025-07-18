// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/billing/v1/billing_account.proto

package billing

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

// A BillingAccount resource. For more information, see [BillingAccount](/docs/billing/concepts/billing-account).
type BillingAccount struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// ID of the billing account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the billing account.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Creation timestamp.
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// ISO 3166-1 alpha-2 country code of the billing account.
	CountryCode string `protobuf:"bytes,4,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	// Currency of the billing account.
	// Can be one of the following:
	// * `RUB`
	// * `USD`
	// * `KZT`
	Currency string `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	// Represents whether corresponding billable objects can be used or not.
	Active bool `protobuf:"varint,6,opt,name=active,proto3" json:"active,omitempty"`
	// Current balance of the billing account.
	Balance       string `protobuf:"bytes,7,opt,name=balance,proto3" json:"balance,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BillingAccount) Reset() {
	*x = BillingAccount{}
	mi := &file_yandex_cloud_billing_v1_billing_account_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BillingAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BillingAccount) ProtoMessage() {}

func (x *BillingAccount) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_billing_v1_billing_account_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BillingAccount.ProtoReflect.Descriptor instead.
func (*BillingAccount) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_billing_v1_billing_account_proto_rawDescGZIP(), []int{0}
}

func (x *BillingAccount) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BillingAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BillingAccount) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *BillingAccount) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *BillingAccount) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *BillingAccount) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *BillingAccount) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

var File_yandex_cloud_billing_v1_billing_account_proto protoreflect.FileDescriptor

const file_yandex_cloud_billing_v1_billing_account_proto_rawDesc = "" +
	"\n" +
	"-yandex/cloud/billing/v1/billing_account.proto\x12\x17yandex.cloud.billing.v1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xe0\x01\n" +
	"\x0eBillingAccount\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12!\n" +
	"\fcountry_code\x18\x04 \x01(\tR\vcountryCode\x12\x1a\n" +
	"\bcurrency\x18\x05 \x01(\tR\bcurrency\x12\x16\n" +
	"\x06active\x18\x06 \x01(\bR\x06active\x12\x18\n" +
	"\abalance\x18\a \x01(\tR\abalanceBb\n" +
	"\x1byandex.cloud.api.billing.v1ZCgithub.com/yandex-cloud/go-genproto/yandex/cloud/billing/v1;billingb\x06proto3"

var (
	file_yandex_cloud_billing_v1_billing_account_proto_rawDescOnce sync.Once
	file_yandex_cloud_billing_v1_billing_account_proto_rawDescData []byte
)

func file_yandex_cloud_billing_v1_billing_account_proto_rawDescGZIP() []byte {
	file_yandex_cloud_billing_v1_billing_account_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_billing_v1_billing_account_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_billing_v1_billing_account_proto_rawDesc), len(file_yandex_cloud_billing_v1_billing_account_proto_rawDesc)))
	})
	return file_yandex_cloud_billing_v1_billing_account_proto_rawDescData
}

var file_yandex_cloud_billing_v1_billing_account_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_yandex_cloud_billing_v1_billing_account_proto_goTypes = []any{
	(*BillingAccount)(nil),        // 0: yandex.cloud.billing.v1.BillingAccount
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_yandex_cloud_billing_v1_billing_account_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.billing.v1.BillingAccount.created_at:type_name -> google.protobuf.Timestamp
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_yandex_cloud_billing_v1_billing_account_proto_init() }
func file_yandex_cloud_billing_v1_billing_account_proto_init() {
	if File_yandex_cloud_billing_v1_billing_account_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_billing_v1_billing_account_proto_rawDesc), len(file_yandex_cloud_billing_v1_billing_account_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_billing_v1_billing_account_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_billing_v1_billing_account_proto_depIdxs,
		MessageInfos:      file_yandex_cloud_billing_v1_billing_account_proto_msgTypes,
	}.Build()
	File_yandex_cloud_billing_v1_billing_account_proto = out.File
	file_yandex_cloud_billing_v1_billing_account_proto_goTypes = nil
	file_yandex_cloud_billing_v1_billing_account_proto_depIdxs = nil
}
