// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/ai/translate/v2/translation_service.proto

package translate

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/yandex-cloud/go-genproto/yandex/cloud"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type TranslateRequest_Format int32

const (
	TranslateRequest_FORMAT_UNSPECIFIED TranslateRequest_Format = 0
	// Text without markup. Default value.
	TranslateRequest_PLAIN_TEXT TranslateRequest_Format = 1
	// Text in the HTML format.
	TranslateRequest_HTML TranslateRequest_Format = 2
)

var TranslateRequest_Format_name = map[int32]string{
	0: "FORMAT_UNSPECIFIED",
	1: "PLAIN_TEXT",
	2: "HTML",
}

var TranslateRequest_Format_value = map[string]int32{
	"FORMAT_UNSPECIFIED": 0,
	"PLAIN_TEXT":         1,
	"HTML":               2,
}

func (x TranslateRequest_Format) String() string {
	return proto.EnumName(TranslateRequest_Format_name, int32(x))
}

func (TranslateRequest_Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{0, 0}
}

type TranslateRequest struct {
	// The text language to translate from.
	// Specified in [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) format (for example, `` ru ``).
	//
	// Required for translating with glossary.
	SourceLanguageCode string `protobuf:"bytes,1,opt,name=source_language_code,json=sourceLanguageCode,proto3" json:"source_language_code,omitempty"`
	// The target language to translate the text.
	// Specified in [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) format (for example, `` en ``).
	TargetLanguageCode string `protobuf:"bytes,2,opt,name=target_language_code,json=targetLanguageCode,proto3" json:"target_language_code,omitempty"`
	// Format of the text.
	Format TranslateRequest_Format `protobuf:"varint,3,opt,name=format,proto3,enum=yandex.cloud.ai.translate.v2.TranslateRequest_Format" json:"format,omitempty"`
	// Array of the strings to translate.
	// The maximum total length of all strings is 10000 characters.
	Texts []string `protobuf:"bytes,4,rep,name=texts,proto3" json:"texts,omitempty"`
	// ID of the folder to which you have access.
	// Required for authorization with a user account (see [yandex.cloud.iam.v1.UserAccount] resource).
	// Don't specify this field if you make the request on behalf of a service account.
	FolderId string `protobuf:"bytes,5,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	// Do not specify this field, custom models are not supported yet.
	Model string `protobuf:"bytes,6,opt,name=model,proto3" json:"model,omitempty"`
	// Glossary to be applied for the translation. For more information, see [Glossaries](/docs/translate/concepts/glossary).
	GlossaryConfig       *TranslateGlossaryConfig `protobuf:"bytes,7,opt,name=glossary_config,json=glossaryConfig,proto3" json:"glossary_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *TranslateRequest) Reset()         { *m = TranslateRequest{} }
func (m *TranslateRequest) String() string { return proto.CompactTextString(m) }
func (*TranslateRequest) ProtoMessage()    {}
func (*TranslateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{0}
}

func (m *TranslateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateRequest.Unmarshal(m, b)
}
func (m *TranslateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateRequest.Marshal(b, m, deterministic)
}
func (m *TranslateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateRequest.Merge(m, src)
}
func (m *TranslateRequest) XXX_Size() int {
	return xxx_messageInfo_TranslateRequest.Size(m)
}
func (m *TranslateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateRequest proto.InternalMessageInfo

func (m *TranslateRequest) GetSourceLanguageCode() string {
	if m != nil {
		return m.SourceLanguageCode
	}
	return ""
}

func (m *TranslateRequest) GetTargetLanguageCode() string {
	if m != nil {
		return m.TargetLanguageCode
	}
	return ""
}

func (m *TranslateRequest) GetFormat() TranslateRequest_Format {
	if m != nil {
		return m.Format
	}
	return TranslateRequest_FORMAT_UNSPECIFIED
}

func (m *TranslateRequest) GetTexts() []string {
	if m != nil {
		return m.Texts
	}
	return nil
}

func (m *TranslateRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *TranslateRequest) GetModel() string {
	if m != nil {
		return m.Model
	}
	return ""
}

func (m *TranslateRequest) GetGlossaryConfig() *TranslateGlossaryConfig {
	if m != nil {
		return m.GlossaryConfig
	}
	return nil
}

type TranslateGlossaryConfig struct {
	// Types that are valid to be assigned to GlossarySource:
	//	*TranslateGlossaryConfig_GlossaryData
	GlossarySource       isTranslateGlossaryConfig_GlossarySource `protobuf_oneof:"glossary_source"`
	XXX_NoUnkeyedLiteral struct{}                                 `json:"-"`
	XXX_unrecognized     []byte                                   `json:"-"`
	XXX_sizecache        int32                                    `json:"-"`
}

func (m *TranslateGlossaryConfig) Reset()         { *m = TranslateGlossaryConfig{} }
func (m *TranslateGlossaryConfig) String() string { return proto.CompactTextString(m) }
func (*TranslateGlossaryConfig) ProtoMessage()    {}
func (*TranslateGlossaryConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{1}
}

func (m *TranslateGlossaryConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateGlossaryConfig.Unmarshal(m, b)
}
func (m *TranslateGlossaryConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateGlossaryConfig.Marshal(b, m, deterministic)
}
func (m *TranslateGlossaryConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateGlossaryConfig.Merge(m, src)
}
func (m *TranslateGlossaryConfig) XXX_Size() int {
	return xxx_messageInfo_TranslateGlossaryConfig.Size(m)
}
func (m *TranslateGlossaryConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateGlossaryConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateGlossaryConfig proto.InternalMessageInfo

type isTranslateGlossaryConfig_GlossarySource interface {
	isTranslateGlossaryConfig_GlossarySource()
}

type TranslateGlossaryConfig_GlossaryData struct {
	GlossaryData *GlossaryData `protobuf:"bytes,1,opt,name=glossary_data,json=glossaryData,proto3,oneof"`
}

func (*TranslateGlossaryConfig_GlossaryData) isTranslateGlossaryConfig_GlossarySource() {}

func (m *TranslateGlossaryConfig) GetGlossarySource() isTranslateGlossaryConfig_GlossarySource {
	if m != nil {
		return m.GlossarySource
	}
	return nil
}

func (m *TranslateGlossaryConfig) GetGlossaryData() *GlossaryData {
	if x, ok := m.GetGlossarySource().(*TranslateGlossaryConfig_GlossaryData); ok {
		return x.GlossaryData
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TranslateGlossaryConfig) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TranslateGlossaryConfig_GlossaryData)(nil),
	}
}

type GlossaryData struct {
	// Array of text pairs.
	//
	// The maximum total length of all source texts is 10000 characters.
	// The maximum total length of all translated texts is 10000 characters.
	GlossaryPairs        []*GlossaryPair `protobuf:"bytes,1,rep,name=glossary_pairs,json=glossaryPairs,proto3" json:"glossary_pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GlossaryData) Reset()         { *m = GlossaryData{} }
func (m *GlossaryData) String() string { return proto.CompactTextString(m) }
func (*GlossaryData) ProtoMessage()    {}
func (*GlossaryData) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{2}
}

func (m *GlossaryData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlossaryData.Unmarshal(m, b)
}
func (m *GlossaryData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlossaryData.Marshal(b, m, deterministic)
}
func (m *GlossaryData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlossaryData.Merge(m, src)
}
func (m *GlossaryData) XXX_Size() int {
	return xxx_messageInfo_GlossaryData.Size(m)
}
func (m *GlossaryData) XXX_DiscardUnknown() {
	xxx_messageInfo_GlossaryData.DiscardUnknown(m)
}

var xxx_messageInfo_GlossaryData proto.InternalMessageInfo

func (m *GlossaryData) GetGlossaryPairs() []*GlossaryPair {
	if m != nil {
		return m.GlossaryPairs
	}
	return nil
}

type GlossaryPair struct {
	// Text in the source language.
	SourceText string `protobuf:"bytes,1,opt,name=source_text,json=sourceText,proto3" json:"source_text,omitempty"`
	// Text in the target language.
	TranslatedText       string   `protobuf:"bytes,2,opt,name=translated_text,json=translatedText,proto3" json:"translated_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GlossaryPair) Reset()         { *m = GlossaryPair{} }
func (m *GlossaryPair) String() string { return proto.CompactTextString(m) }
func (*GlossaryPair) ProtoMessage()    {}
func (*GlossaryPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{3}
}

func (m *GlossaryPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GlossaryPair.Unmarshal(m, b)
}
func (m *GlossaryPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GlossaryPair.Marshal(b, m, deterministic)
}
func (m *GlossaryPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GlossaryPair.Merge(m, src)
}
func (m *GlossaryPair) XXX_Size() int {
	return xxx_messageInfo_GlossaryPair.Size(m)
}
func (m *GlossaryPair) XXX_DiscardUnknown() {
	xxx_messageInfo_GlossaryPair.DiscardUnknown(m)
}

var xxx_messageInfo_GlossaryPair proto.InternalMessageInfo

func (m *GlossaryPair) GetSourceText() string {
	if m != nil {
		return m.SourceText
	}
	return ""
}

func (m *GlossaryPair) GetTranslatedText() string {
	if m != nil {
		return m.TranslatedText
	}
	return ""
}

type TranslateResponse struct {
	// Array of the translations.
	Translations         []*TranslatedText `protobuf:"bytes,1,rep,name=translations,proto3" json:"translations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TranslateResponse) Reset()         { *m = TranslateResponse{} }
func (m *TranslateResponse) String() string { return proto.CompactTextString(m) }
func (*TranslateResponse) ProtoMessage()    {}
func (*TranslateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{4}
}

func (m *TranslateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TranslateResponse.Unmarshal(m, b)
}
func (m *TranslateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TranslateResponse.Marshal(b, m, deterministic)
}
func (m *TranslateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TranslateResponse.Merge(m, src)
}
func (m *TranslateResponse) XXX_Size() int {
	return xxx_messageInfo_TranslateResponse.Size(m)
}
func (m *TranslateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TranslateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TranslateResponse proto.InternalMessageInfo

func (m *TranslateResponse) GetTranslations() []*TranslatedText {
	if m != nil {
		return m.Translations
	}
	return nil
}

type DetectLanguageRequest struct {
	// The text to detect the language for.
	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	// List of the most likely languages. These languages will be given preference when detecting the text language.
	// Specified in [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) format (for example, `` ru ``).
	//
	// To get the list of supported languages, use a [TranslationService.ListLanguages] request.
	LanguageCodeHints []string `protobuf:"bytes,2,rep,name=language_code_hints,json=languageCodeHints,proto3" json:"language_code_hints,omitempty"`
	// ID of the folder to which you have access.
	// Required for authorization with a user account (see [yandex.cloud.iam.v1.UserAccount] resource).
	// Don't specify this field if you make the request on behalf of a service account.
	FolderId             string   `protobuf:"bytes,3,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectLanguageRequest) Reset()         { *m = DetectLanguageRequest{} }
func (m *DetectLanguageRequest) String() string { return proto.CompactTextString(m) }
func (*DetectLanguageRequest) ProtoMessage()    {}
func (*DetectLanguageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{5}
}

func (m *DetectLanguageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectLanguageRequest.Unmarshal(m, b)
}
func (m *DetectLanguageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectLanguageRequest.Marshal(b, m, deterministic)
}
func (m *DetectLanguageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectLanguageRequest.Merge(m, src)
}
func (m *DetectLanguageRequest) XXX_Size() int {
	return xxx_messageInfo_DetectLanguageRequest.Size(m)
}
func (m *DetectLanguageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectLanguageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetectLanguageRequest proto.InternalMessageInfo

func (m *DetectLanguageRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *DetectLanguageRequest) GetLanguageCodeHints() []string {
	if m != nil {
		return m.LanguageCodeHints
	}
	return nil
}

func (m *DetectLanguageRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

type DetectLanguageResponse struct {
	// The text language in [ISO 639-1](https://en.wikipedia.org/wiki/ISO_639-1) format (for example, `` ru ``).
	//
	// To get the language name, use a [TranslationService.ListLanguages] request.
	LanguageCode         string   `protobuf:"bytes,1,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DetectLanguageResponse) Reset()         { *m = DetectLanguageResponse{} }
func (m *DetectLanguageResponse) String() string { return proto.CompactTextString(m) }
func (*DetectLanguageResponse) ProtoMessage()    {}
func (*DetectLanguageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{6}
}

func (m *DetectLanguageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectLanguageResponse.Unmarshal(m, b)
}
func (m *DetectLanguageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectLanguageResponse.Marshal(b, m, deterministic)
}
func (m *DetectLanguageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectLanguageResponse.Merge(m, src)
}
func (m *DetectLanguageResponse) XXX_Size() int {
	return xxx_messageInfo_DetectLanguageResponse.Size(m)
}
func (m *DetectLanguageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectLanguageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetectLanguageResponse proto.InternalMessageInfo

func (m *DetectLanguageResponse) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

type ListLanguagesRequest struct {
	// ID of the folder to which you have access.
	// Required for authorization with a user account (see [yandex.cloud.iam.v1.UserAccount] resource).
	// Don't specify this field if you make the request on behalf of a service account.
	FolderId             string   `protobuf:"bytes,1,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListLanguagesRequest) Reset()         { *m = ListLanguagesRequest{} }
func (m *ListLanguagesRequest) String() string { return proto.CompactTextString(m) }
func (*ListLanguagesRequest) ProtoMessage()    {}
func (*ListLanguagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{7}
}

func (m *ListLanguagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLanguagesRequest.Unmarshal(m, b)
}
func (m *ListLanguagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLanguagesRequest.Marshal(b, m, deterministic)
}
func (m *ListLanguagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLanguagesRequest.Merge(m, src)
}
func (m *ListLanguagesRequest) XXX_Size() int {
	return xxx_messageInfo_ListLanguagesRequest.Size(m)
}
func (m *ListLanguagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLanguagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListLanguagesRequest proto.InternalMessageInfo

func (m *ListLanguagesRequest) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

type ListLanguagesResponse struct {
	// List of supported languages.
	Languages            []*Language `protobuf:"bytes,1,rep,name=languages,proto3" json:"languages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ListLanguagesResponse) Reset()         { *m = ListLanguagesResponse{} }
func (m *ListLanguagesResponse) String() string { return proto.CompactTextString(m) }
func (*ListLanguagesResponse) ProtoMessage()    {}
func (*ListLanguagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07212bd16e7bdb8a, []int{8}
}

func (m *ListLanguagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLanguagesResponse.Unmarshal(m, b)
}
func (m *ListLanguagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLanguagesResponse.Marshal(b, m, deterministic)
}
func (m *ListLanguagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLanguagesResponse.Merge(m, src)
}
func (m *ListLanguagesResponse) XXX_Size() int {
	return xxx_messageInfo_ListLanguagesResponse.Size(m)
}
func (m *ListLanguagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLanguagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListLanguagesResponse proto.InternalMessageInfo

func (m *ListLanguagesResponse) GetLanguages() []*Language {
	if m != nil {
		return m.Languages
	}
	return nil
}

func init() {
	proto.RegisterEnum("yandex.cloud.ai.translate.v2.TranslateRequest_Format", TranslateRequest_Format_name, TranslateRequest_Format_value)
	proto.RegisterType((*TranslateRequest)(nil), "yandex.cloud.ai.translate.v2.TranslateRequest")
	proto.RegisterType((*TranslateGlossaryConfig)(nil), "yandex.cloud.ai.translate.v2.TranslateGlossaryConfig")
	proto.RegisterType((*GlossaryData)(nil), "yandex.cloud.ai.translate.v2.GlossaryData")
	proto.RegisterType((*GlossaryPair)(nil), "yandex.cloud.ai.translate.v2.GlossaryPair")
	proto.RegisterType((*TranslateResponse)(nil), "yandex.cloud.ai.translate.v2.TranslateResponse")
	proto.RegisterType((*DetectLanguageRequest)(nil), "yandex.cloud.ai.translate.v2.DetectLanguageRequest")
	proto.RegisterType((*DetectLanguageResponse)(nil), "yandex.cloud.ai.translate.v2.DetectLanguageResponse")
	proto.RegisterType((*ListLanguagesRequest)(nil), "yandex.cloud.ai.translate.v2.ListLanguagesRequest")
	proto.RegisterType((*ListLanguagesResponse)(nil), "yandex.cloud.ai.translate.v2.ListLanguagesResponse")
}

func init() {
	proto.RegisterFile("yandex/cloud/ai/translate/v2/translation_service.proto", fileDescriptor_07212bd16e7bdb8a)
}

var fileDescriptor_07212bd16e7bdb8a = []byte{
	// 828 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x41, 0x6f, 0xe3, 0x44,
	0x14, 0x5e, 0x27, 0x69, 0xb6, 0x79, 0x6d, 0xd3, 0xee, 0xd0, 0xdd, 0x9a, 0xa8, 0x40, 0x64, 0xb4,
	0xa8, 0xaa, 0x88, 0x1d, 0x27, 0x14, 0x89, 0xd2, 0x80, 0x36, 0x4d, 0x4b, 0x83, 0xd2, 0x25, 0x78,
	0x83, 0x84, 0x40, 0x60, 0xcd, 0xc6, 0x53, 0xaf, 0x85, 0xeb, 0x09, 0xf6, 0x24, 0xea, 0x5e, 0x23,
	0x71, 0xd9, 0x13, 0x12, 0x57, 0x6e, 0xfc, 0x11, 0x38, 0x91, 0x3b, 0x57, 0x8e, 0x20, 0xf1, 0x33,
	0x50, 0x66, 0x6c, 0xc7, 0x8e, 0xd2, 0x6c, 0x7a, 0x9c, 0x79, 0xdf, 0x37, 0xdf, 0xf7, 0xe6, 0xbd,
	0x37, 0x03, 0x1f, 0xbe, 0xc4, 0x9e, 0x45, 0x6e, 0xb4, 0xbe, 0x4b, 0x87, 0x96, 0x86, 0x1d, 0x8d,
	0xf9, 0xd8, 0x0b, 0x5c, 0xcc, 0x88, 0x36, 0xaa, 0xc5, 0x0b, 0x87, 0x7a, 0x66, 0x40, 0xfc, 0x91,
	0xd3, 0x27, 0xea, 0xc0, 0xa7, 0x8c, 0xa2, 0x7d, 0xc1, 0x53, 0x39, 0x4f, 0xc5, 0x8e, 0x1a, 0xf3,
	0xd4, 0x51, 0xad, 0xb4, 0x6f, 0x53, 0x6a, 0xbb, 0x44, 0xc3, 0x03, 0x47, 0xc3, 0x9e, 0x47, 0x19,
	0x3f, 0x22, 0x10, 0xdc, 0xd2, 0x5b, 0x29, 0xcd, 0x11, 0x76, 0x1d, 0x8b, 0xc7, 0xc3, 0xb0, 0xba,
	0xaa, 0x25, 0x81, 0x57, 0xfe, 0xce, 0xc2, 0x4e, 0x2f, 0x82, 0x18, 0xe4, 0xc7, 0x21, 0x09, 0x18,
	0xfa, 0x08, 0x76, 0x03, 0x3a, 0xf4, 0xfb, 0xc4, 0x74, 0xb1, 0x67, 0x0f, 0xb1, 0x4d, 0xcc, 0x3e,
	0xb5, 0x88, 0x2c, 0x95, 0xa5, 0x83, 0x42, 0xf3, 0xfe, 0xab, 0x89, 0x9e, 0x3d, 0x69, 0xd4, 0x0d,
	0x24, 0x40, 0x9d, 0x10, 0x73, 0x4a, 0x2d, 0x82, 0x1a, 0xb0, 0xcb, 0xb0, 0x6f, 0x13, 0x36, 0x47,
	0xcd, 0x70, 0xea, 0xc6, 0x7f, 0x7f, 0xea, 0x52, 0x4c, 0x17, 0xc0, 0x14, 0xfd, 0x12, 0xf2, 0x57,
	0xd4, 0xbf, 0xc6, 0x4c, 0xce, 0x96, 0xa5, 0x83, 0x62, 0xed, 0x48, 0x5d, 0x76, 0x55, 0xea, 0xbc,
	0x73, 0xf5, 0x9c, 0x93, 0x8d, 0xf0, 0x10, 0xb4, 0x0f, 0x6b, 0x8c, 0xdc, 0xb0, 0x40, 0xce, 0x95,
	0xb3, 0x07, 0x85, 0x66, 0x7e, 0x3c, 0xd1, 0x33, 0x9f, 0x54, 0x0d, 0xb1, 0x89, 0x1e, 0x43, 0xe1,
	0x8a, 0xba, 0x16, 0xf1, 0x4d, 0xc7, 0x92, 0xd7, 0xb8, 0xc1, 0xf5, 0x57, 0x13, 0x3d, 0x77, 0xd2,
	0x38, 0xaa, 0x1a, 0xeb, 0x22, 0xd4, 0xb6, 0xd0, 0xdb, 0xb0, 0x76, 0x4d, 0x2d, 0xe2, 0xca, 0xf9,
	0x39, 0x88, 0xd8, 0x46, 0xdf, 0xc3, 0xb6, 0xed, 0xd2, 0x20, 0xc0, 0xfe, 0x4b, 0xb3, 0x4f, 0xbd,
	0x2b, 0xc7, 0x96, 0xef, 0x97, 0xa5, 0x83, 0x8d, 0x95, 0xcd, 0x7f, 0x16, 0xb2, 0x4f, 0x39, 0xd9,
	0x28, 0xda, 0xa9, 0xb5, 0x72, 0x0c, 0x79, 0x91, 0x16, 0x7a, 0x04, 0xe8, 0xfc, 0x0b, 0xe3, 0xf2,
	0x49, 0xcf, 0xfc, 0xea, 0xe9, 0xb3, 0xee, 0xd9, 0x69, 0xfb, 0xbc, 0x7d, 0xd6, 0xda, 0xb9, 0x87,
	0x8a, 0x00, 0xdd, 0xce, 0x93, 0xf6, 0x53, 0xb3, 0x77, 0xf6, 0x75, 0x6f, 0x47, 0x42, 0xeb, 0x90,
	0xbb, 0xe8, 0x5d, 0x76, 0x76, 0x32, 0xca, 0x4f, 0x12, 0xec, 0xdd, 0xa2, 0x83, 0xbe, 0x84, 0xad,
	0xd8, 0xb7, 0x85, 0x19, 0xe6, 0xe5, 0xdd, 0xa8, 0x1d, 0x2e, 0x77, 0x1d, 0x1d, 0xd2, 0xc2, 0x0c,
	0x5f, 0xdc, 0x33, 0x36, 0xed, 0xc4, 0xba, 0xb9, 0x97, 0xb8, 0x0a, 0xd1, 0x1c, 0x28, 0xf7, 0xfb,
	0x1f, 0xba, 0xa4, 0xfc, 0x00, 0x9b, 0x49, 0x22, 0xfa, 0x16, 0xe2, 0x2c, 0xcd, 0x01, 0x76, 0xfc,
	0x40, 0x96, 0xca, 0xd9, 0xd5, 0xc5, 0xbb, 0xd8, 0xf1, 0x9b, 0xeb, 0xe3, 0x89, 0x9e, 0xd3, 0x2b,
	0x47, 0x55, 0x23, 0xce, 0x63, 0xba, 0x1f, 0x28, 0xd6, 0x4c, 0x6c, 0xba, 0x81, 0x1e, 0xc3, 0x46,
	0xd8, 0xce, 0xd3, 0xba, 0x87, 0x5d, 0x9c, 0x9b, 0xb6, 0xa2, 0x01, 0x22, 0xd0, 0x23, 0x37, 0x0c,
	0x55, 0x60, 0x3b, 0x16, 0xb3, 0x04, 0x34, 0x93, 0x80, 0x16, 0x67, 0xc1, 0x29, 0x5c, 0x21, 0xf0,
	0x20, 0xd1, 0x7e, 0xc1, 0x80, 0x7a, 0x01, 0x41, 0x5d, 0xd8, 0x4c, 0xcc, 0x58, 0x94, 0xd5, 0xfb,
	0x2b, 0x36, 0x02, 0x3f, 0xd8, 0x48, 0x9d, 0xa0, 0xfc, 0x26, 0xc1, 0xc3, 0x16, 0x61, 0xa4, 0x1f,
	0x0f, 0x4a, 0x34, 0xa5, 0x0a, 0xe4, 0x12, 0xf9, 0x14, 0xc3, 0xd1, 0xca, 0x9f, 0x34, 0xf4, 0x6a,
	0xb5, 0x6a, 0xf0, 0x18, 0xfa, 0x14, 0xde, 0x48, 0xcd, 0xa1, 0xf9, 0xc2, 0xf1, 0x58, 0x20, 0x67,
	0xf8, 0x38, 0x6c, 0x8f, 0x79, 0x27, 0xeb, 0xd5, 0x68, 0x22, 0x1f, 0xb8, 0x89, 0x59, 0xbc, 0x98,
	0x22, 0xd3, 0x33, 0x92, 0xbd, 0x6d, 0x46, 0x94, 0x06, 0x3c, 0x9a, 0x37, 0x19, 0xde, 0xc8, 0xbb,
	0xb0, 0xb5, 0xe0, 0x11, 0x31, 0x36, 0x93, 0x52, 0x4a, 0x03, 0x76, 0x3b, 0x4e, 0x10, 0x93, 0x83,
	0x28, 0xc5, 0x94, 0xba, 0x74, 0xab, 0xfa, 0x77, 0xf0, 0x70, 0x8e, 0x1e, 0x8a, 0xb7, 0xa0, 0x10,
	0xe9, 0x44, 0xb5, 0x78, 0x6f, 0x79, 0x2d, 0x62, 0xff, 0x33, 0x62, 0xed, 0xdf, 0x2c, 0xa0, 0xde,
	0xac, 0x26, 0xcf, 0xc4, 0x5b, 0x8e, 0x7e, 0x96, 0xa0, 0x10, 0x97, 0x0e, 0xa9, 0x77, 0x7b, 0xa9,
	0x4a, 0xda, 0xca, 0x78, 0x91, 0x8b, 0xa2, 0x8c, 0xff, 0xfa, 0xe7, 0x97, 0xcc, 0xbe, 0xb2, 0xb7,
	0xf8, 0x49, 0x27, 0xc7, 0xd2, 0x21, 0xfa, 0x55, 0x82, 0x62, 0xba, 0x0e, 0xa8, 0xbe, 0x5c, 0x67,
	0x61, 0x6b, 0x95, 0x3e, 0xb8, 0x1b, 0x29, 0x74, 0xf8, 0x0e, 0x77, 0xf8, 0xa6, 0xb2, 0x9b, 0x76,
	0x68, 0x71, 0x74, 0x68, 0x6f, 0x2b, 0x55, 0x28, 0x54, 0x7b, 0x4d, 0x35, 0x16, 0x34, 0x45, 0xa9,
	0x7e, 0x27, 0xce, 0xf2, 0xdb, 0x8b, 0x8b, 0x7c, 0x2c, 0x1d, 0x36, 0x3b, 0xdf, 0x7c, 0x6e, 0x3b,
	0xec, 0xc5, 0xf0, 0xb9, 0xda, 0xa7, 0xd7, 0x9a, 0x10, 0xa9, 0x88, 0x8f, 0xd4, 0xa6, 0x15, 0x9b,
	0x78, 0xfc, 0xcb, 0xd4, 0x96, 0xfd, 0xb0, 0x1f, 0xc7, 0x8b, 0xe7, 0x79, 0x8e, 0xae, 0xff, 0x1f,
	0x00, 0x00, 0xff, 0xff, 0x9c, 0x71, 0x17, 0xd1, 0x25, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TranslationServiceClient is the client API for TranslationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TranslationServiceClient interface {
	// Translates the text to the specified language.
	Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error)
	// Detects the language of the text.
	DetectLanguage(ctx context.Context, in *DetectLanguageRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error)
	// Retrieves the list of supported languages.
	ListLanguages(ctx context.Context, in *ListLanguagesRequest, opts ...grpc.CallOption) (*ListLanguagesResponse, error)
}

type translationServiceClient struct {
	cc *grpc.ClientConn
}

func NewTranslationServiceClient(cc *grpc.ClientConn) TranslationServiceClient {
	return &translationServiceClient{cc}
}

func (c *translationServiceClient) Translate(ctx context.Context, in *TranslateRequest, opts ...grpc.CallOption) (*TranslateResponse, error) {
	out := new(TranslateResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ai.translate.v2.TranslationService/Translate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationServiceClient) DetectLanguage(ctx context.Context, in *DetectLanguageRequest, opts ...grpc.CallOption) (*DetectLanguageResponse, error) {
	out := new(DetectLanguageResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ai.translate.v2.TranslationService/DetectLanguage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *translationServiceClient) ListLanguages(ctx context.Context, in *ListLanguagesRequest, opts ...grpc.CallOption) (*ListLanguagesResponse, error) {
	out := new(ListLanguagesResponse)
	err := c.cc.Invoke(ctx, "/yandex.cloud.ai.translate.v2.TranslationService/ListLanguages", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TranslationServiceServer is the server API for TranslationService service.
type TranslationServiceServer interface {
	// Translates the text to the specified language.
	Translate(context.Context, *TranslateRequest) (*TranslateResponse, error)
	// Detects the language of the text.
	DetectLanguage(context.Context, *DetectLanguageRequest) (*DetectLanguageResponse, error)
	// Retrieves the list of supported languages.
	ListLanguages(context.Context, *ListLanguagesRequest) (*ListLanguagesResponse, error)
}

// UnimplementedTranslationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTranslationServiceServer struct {
}

func (*UnimplementedTranslationServiceServer) Translate(ctx context.Context, req *TranslateRequest) (*TranslateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Translate not implemented")
}
func (*UnimplementedTranslationServiceServer) DetectLanguage(ctx context.Context, req *DetectLanguageRequest) (*DetectLanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetectLanguage not implemented")
}
func (*UnimplementedTranslationServiceServer) ListLanguages(ctx context.Context, req *ListLanguagesRequest) (*ListLanguagesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLanguages not implemented")
}

func RegisterTranslationServiceServer(s *grpc.Server, srv TranslationServiceServer) {
	s.RegisterService(&_TranslationService_serviceDesc, srv)
}

func _TranslationService_Translate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranslateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).Translate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ai.translate.v2.TranslationService/Translate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).Translate(ctx, req.(*TranslateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationService_DetectLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectLanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).DetectLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ai.translate.v2.TranslationService/DetectLanguage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).DetectLanguage(ctx, req.(*DetectLanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TranslationService_ListLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLanguagesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TranslationServiceServer).ListLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yandex.cloud.ai.translate.v2.TranslationService/ListLanguages",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TranslationServiceServer).ListLanguages(ctx, req.(*ListLanguagesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TranslationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yandex.cloud.ai.translate.v2.TranslationService",
	HandlerType: (*TranslationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Translate",
			Handler:    _TranslationService_Translate_Handler,
		},
		{
			MethodName: "DetectLanguage",
			Handler:    _TranslationService_DetectLanguage_Handler,
		},
		{
			MethodName: "ListLanguages",
			Handler:    _TranslationService_ListLanguages_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yandex/cloud/ai/translate/v2/translation_service.proto",
}
