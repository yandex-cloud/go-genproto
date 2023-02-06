// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: yandex/cloud/ai/vision/v2/image_classifier.proto

package vision

import (
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

type ClassifierSpecification_ClassificationType int32

const (
	ClassifierSpecification_CLASSIFICATION_TYPE_UNSPECIFIED ClassifierSpecification_ClassificationType = 0
	ClassifierSpecification_MULTI_LABEL                     ClassifierSpecification_ClassificationType = 1
	ClassifierSpecification_MULTI_CLASS                     ClassifierSpecification_ClassificationType = 2
)

// Enum value maps for ClassifierSpecification_ClassificationType.
var (
	ClassifierSpecification_ClassificationType_name = map[int32]string{
		0: "CLASSIFICATION_TYPE_UNSPECIFIED",
		1: "MULTI_LABEL",
		2: "MULTI_CLASS",
	}
	ClassifierSpecification_ClassificationType_value = map[string]int32{
		"CLASSIFICATION_TYPE_UNSPECIFIED": 0,
		"MULTI_LABEL":                     1,
		"MULTI_CLASS":                     2,
	}
)

func (x ClassifierSpecification_ClassificationType) Enum() *ClassifierSpecification_ClassificationType {
	p := new(ClassifierSpecification_ClassificationType)
	*p = x
	return p
}

func (x ClassifierSpecification_ClassificationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ClassifierSpecification_ClassificationType) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_enumTypes[0].Descriptor()
}

func (ClassifierSpecification_ClassificationType) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_vision_v2_image_classifier_proto_enumTypes[0]
}

func (x ClassifierSpecification_ClassificationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ClassifierSpecification_ClassificationType.Descriptor instead.
func (ClassifierSpecification_ClassificationType) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP(), []int{2, 0}
}

// Description of single label
type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Label name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// human readable description of label
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP(), []int{0}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// Image annotation for specific label
type ClassAnnotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// list of annotated labels
	Label *Label `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	// confidence for each label
	Confidence float64 `protobuf:"fixed64,2,opt,name=confidence,proto3" json:"confidence,omitempty"`
}

func (x *ClassAnnotation) Reset() {
	*x = ClassAnnotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassAnnotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassAnnotation) ProtoMessage() {}

func (x *ClassAnnotation) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassAnnotation.ProtoReflect.Descriptor instead.
func (*ClassAnnotation) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP(), []int{1}
}

func (x *ClassAnnotation) GetLabel() *Label {
	if x != nil {
		return x.Label
	}
	return nil
}

func (x *ClassAnnotation) GetConfidence() float64 {
	if x != nil {
		return x.Confidence
	}
	return 0
}

// Specification of model used for annotation
type ClassifierSpecification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of labels, annotated by service
	Labels []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// type of annotation: exclusive (multi-class) or non-exclusive (multi-label)
	ClassificationType ClassifierSpecification_ClassificationType `protobuf:"varint,2,opt,name=classification_type,json=classificationType,proto3,enum=yandex.cloud.ai.vision.v2.ClassifierSpecification_ClassificationType" json:"classification_type,omitempty"`
}

func (x *ClassifierSpecification) Reset() {
	*x = ClassifierSpecification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassifierSpecification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassifierSpecification) ProtoMessage() {}

func (x *ClassifierSpecification) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClassifierSpecification.ProtoReflect.Descriptor instead.
func (*ClassifierSpecification) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP(), []int{2}
}

func (x *ClassifierSpecification) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ClassifierSpecification) GetClassificationType() ClassifierSpecification_ClassificationType {
	if x != nil {
		return x.ClassificationType
	}
	return ClassifierSpecification_CLASSIFICATION_TYPE_UNSPECIFIED
}

type AnnotationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// internal service requestId
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// class specification
	ClassifierSpecification *ClassifierSpecification `protobuf:"bytes,2,opt,name=classifier_specification,json=classifierSpecification,proto3" json:"classifier_specification,omitempty"`
	// annotations for each class
	Annotations []*ClassAnnotation `protobuf:"bytes,3,rep,name=annotations,proto3" json:"annotations,omitempty"`
}

func (x *AnnotationResponse) Reset() {
	*x = AnnotationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationResponse) ProtoMessage() {}

func (x *AnnotationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationResponse.ProtoReflect.Descriptor instead.
func (*AnnotationResponse) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP(), []int{3}
}

func (x *AnnotationResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *AnnotationResponse) GetClassifierSpecification() *ClassifierSpecification {
	if x != nil {
		return x.ClassifierSpecification
	}
	return nil
}

func (x *AnnotationResponse) GetAnnotations() []*ClassAnnotation {
	if x != nil {
		return x.Annotations
	}
	return nil
}

// request for annotation
type AnnotationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// image to annotate
	Image *Image `protobuf:"bytes,1,opt,name=image,proto3" json:"image,omitempty"`
}

func (x *AnnotationRequest) Reset() {
	*x = AnnotationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnnotationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnnotationRequest) ProtoMessage() {}

func (x *AnnotationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnnotationRequest.ProtoReflect.Descriptor instead.
func (*AnnotationRequest) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP(), []int{4}
}

func (x *AnnotationRequest) GetImage() *Image {
	if x != nil {
		return x.Image
	}
	return nil
}

var File_yandex_cloud_ai_vision_v2_image_classifier_proto protoreflect.FileDescriptor

var file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDesc = []byte{
	0x0a, 0x30, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61,
	0x69, 0x2f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x19, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x1a, 0x25, 0x79,
	0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3d, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x69, 0x0a, 0x0f, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x22, 0xa8,
	0x02, 0x0a, 0x17, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x76, 0x0a, 0x13, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x45, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x12, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5b, 0x0a, 0x12,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x49, 0x46, 0x49, 0x43, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x55, 0x4c, 0x54, 0x49,
	0x5f, 0x4c, 0x41, 0x42, 0x45, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4d, 0x55, 0x4c, 0x54,
	0x49, 0x5f, 0x43, 0x4c, 0x41, 0x53, 0x53, 0x10, 0x02, 0x22, 0xf0, 0x01, 0x0a, 0x12, 0x41, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x6d, 0x0a, 0x18, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x5f, 0x73, 0x70,
	0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x32, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x17, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c,
	0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4b, 0x0a, 0x11,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e,
	0x61, 0x69, 0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x42, 0x65, 0x0a, 0x1d, 0x79, 0x61, 0x6e,
	0x64, 0x65, 0x78, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x69,
	0x2e, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2d, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x69, 0x2f,
	0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x3b, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescData = file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDesc
)

func file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescData = protoimpl.X.CompressGZIP(file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescData)
	})
	return file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDescData
}

var file_yandex_cloud_ai_vision_v2_image_classifier_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_yandex_cloud_ai_vision_v2_image_classifier_proto_goTypes = []interface{}{
	(ClassifierSpecification_ClassificationType)(0), // 0: yandex.cloud.ai.vision.v2.ClassifierSpecification.ClassificationType
	(*Label)(nil),                   // 1: yandex.cloud.ai.vision.v2.Label
	(*ClassAnnotation)(nil),         // 2: yandex.cloud.ai.vision.v2.ClassAnnotation
	(*ClassifierSpecification)(nil), // 3: yandex.cloud.ai.vision.v2.ClassifierSpecification
	(*AnnotationResponse)(nil),      // 4: yandex.cloud.ai.vision.v2.AnnotationResponse
	(*AnnotationRequest)(nil),       // 5: yandex.cloud.ai.vision.v2.AnnotationRequest
	(*Image)(nil),                   // 6: yandex.cloud.ai.vision.v2.Image
}
var file_yandex_cloud_ai_vision_v2_image_classifier_proto_depIdxs = []int32{
	1, // 0: yandex.cloud.ai.vision.v2.ClassAnnotation.label:type_name -> yandex.cloud.ai.vision.v2.Label
	1, // 1: yandex.cloud.ai.vision.v2.ClassifierSpecification.labels:type_name -> yandex.cloud.ai.vision.v2.Label
	0, // 2: yandex.cloud.ai.vision.v2.ClassifierSpecification.classification_type:type_name -> yandex.cloud.ai.vision.v2.ClassifierSpecification.ClassificationType
	3, // 3: yandex.cloud.ai.vision.v2.AnnotationResponse.classifier_specification:type_name -> yandex.cloud.ai.vision.v2.ClassifierSpecification
	2, // 4: yandex.cloud.ai.vision.v2.AnnotationResponse.annotations:type_name -> yandex.cloud.ai.vision.v2.ClassAnnotation
	6, // 5: yandex.cloud.ai.vision.v2.AnnotationRequest.image:type_name -> yandex.cloud.ai.vision.v2.Image
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_vision_v2_image_classifier_proto_init() }
func file_yandex_cloud_ai_vision_v2_image_classifier_proto_init() {
	if File_yandex_cloud_ai_vision_v2_image_classifier_proto != nil {
		return
	}
	file_yandex_cloud_ai_vision_v2_image_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
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
		file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassAnnotation); i {
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
		file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassifierSpecification); i {
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
		file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationResponse); i {
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
		file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnnotationRequest); i {
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
			RawDescriptor: file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_vision_v2_image_classifier_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_vision_v2_image_classifier_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_vision_v2_image_classifier_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_vision_v2_image_classifier_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_vision_v2_image_classifier_proto = out.File
	file_yandex_cloud_ai_vision_v2_image_classifier_proto_rawDesc = nil
	file_yandex_cloud_ai_vision_v2_image_classifier_proto_goTypes = nil
	file_yandex_cloud_ai_vision_v2_image_classifier_proto_depIdxs = nil
}
