// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: yandex/cloud/ai/assistants/v1/searchindex/common.proto

package searchindex

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
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

// Normalization strategy for relevance scores from different indices
type NormalizationStrategy int32

const (
	NormalizationStrategy_NORMALIZATION_STRATEGY_UNSPECIFIED NormalizationStrategy = 0
	// https://en.wikipedia.org/wiki/Feature_scaling#Rescaling_(min-max_normalization)
	NormalizationStrategy_MIN_MAX NormalizationStrategy = 1
	// https://en.wikipedia.org/wiki/Cosine_similarity#L2-normalized_Euclidean_distance
	NormalizationStrategy_L2 NormalizationStrategy = 2
)

// Enum value maps for NormalizationStrategy.
var (
	NormalizationStrategy_name = map[int32]string{
		0: "NORMALIZATION_STRATEGY_UNSPECIFIED",
		1: "MIN_MAX",
		2: "L2",
	}
	NormalizationStrategy_value = map[string]int32{
		"NORMALIZATION_STRATEGY_UNSPECIFIED": 0,
		"MIN_MAX":                            1,
		"L2":                                 2,
	}
)

func (x NormalizationStrategy) Enum() *NormalizationStrategy {
	p := new(NormalizationStrategy)
	*p = x
	return p
}

func (x NormalizationStrategy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NormalizationStrategy) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_enumTypes[0].Descriptor()
}

func (NormalizationStrategy) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_enumTypes[0]
}

func (x NormalizationStrategy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NormalizationStrategy.Descriptor instead.
func (NormalizationStrategy) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{0}
}

type MeanCombinationStrategy_MeanEvaluationTechnique int32

const (
	MeanCombinationStrategy_MEAN_EVALUATION_TECHNIQUE_UNSPECIFIED MeanCombinationStrategy_MeanEvaluationTechnique = 0
	// https://en.wikipedia.org/wiki/Arithmetic_mean
	MeanCombinationStrategy_ARITHMETIC MeanCombinationStrategy_MeanEvaluationTechnique = 1
	// https://en.wikipedia.org/wiki/Geometric_mean
	MeanCombinationStrategy_GEOMETRIC MeanCombinationStrategy_MeanEvaluationTechnique = 2
	// https://en.wikipedia.org/wiki/Harmonic_mean
	MeanCombinationStrategy_HARMONIC MeanCombinationStrategy_MeanEvaluationTechnique = 3
)

// Enum value maps for MeanCombinationStrategy_MeanEvaluationTechnique.
var (
	MeanCombinationStrategy_MeanEvaluationTechnique_name = map[int32]string{
		0: "MEAN_EVALUATION_TECHNIQUE_UNSPECIFIED",
		1: "ARITHMETIC",
		2: "GEOMETRIC",
		3: "HARMONIC",
	}
	MeanCombinationStrategy_MeanEvaluationTechnique_value = map[string]int32{
		"MEAN_EVALUATION_TECHNIQUE_UNSPECIFIED": 0,
		"ARITHMETIC":                            1,
		"GEOMETRIC":                             2,
		"HARMONIC":                              3,
	}
)

func (x MeanCombinationStrategy_MeanEvaluationTechnique) Enum() *MeanCombinationStrategy_MeanEvaluationTechnique {
	p := new(MeanCombinationStrategy_MeanEvaluationTechnique)
	*p = x
	return p
}

func (x MeanCombinationStrategy_MeanEvaluationTechnique) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MeanCombinationStrategy_MeanEvaluationTechnique) Descriptor() protoreflect.EnumDescriptor {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_enumTypes[1].Descriptor()
}

func (MeanCombinationStrategy_MeanEvaluationTechnique) Type() protoreflect.EnumType {
	return &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_enumTypes[1]
}

func (x MeanCombinationStrategy_MeanEvaluationTechnique) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MeanCombinationStrategy_MeanEvaluationTechnique.Descriptor instead.
func (MeanCombinationStrategy_MeanEvaluationTechnique) EnumDescriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{2, 0}
}

// Defines a chunking strategy where chunks are created with a fixed maximum chunk size and an overlap between consecutive chunks.
type StaticChunkingStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The maximum number of tokens allowed in a single chunk.
	// Constraints: must be within the range [100, 2048].
	// Default value: 800
	MaxChunkSizeTokens int64 `protobuf:"varint,1,opt,name=max_chunk_size_tokens,json=maxChunkSizeTokens,proto3" json:"max_chunk_size_tokens,omitempty"`
	// The number of tokens that should overlap between consecutive chunks.
	// This allows for some context from the previous chunk to be included in the next chunk.
	// Constraints: must be less than or equal to half of `max_chunk_size_tokens`.
	// Default value: 400
	ChunkOverlapTokens int64 `protobuf:"varint,2,opt,name=chunk_overlap_tokens,json=chunkOverlapTokens,proto3" json:"chunk_overlap_tokens,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *StaticChunkingStrategy) Reset() {
	*x = StaticChunkingStrategy{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StaticChunkingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticChunkingStrategy) ProtoMessage() {}

func (x *StaticChunkingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticChunkingStrategy.ProtoReflect.Descriptor instead.
func (*StaticChunkingStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{0}
}

func (x *StaticChunkingStrategy) GetMaxChunkSizeTokens() int64 {
	if x != nil {
		return x.MaxChunkSizeTokens
	}
	return 0
}

func (x *StaticChunkingStrategy) GetChunkOverlapTokens() int64 {
	if x != nil {
		return x.ChunkOverlapTokens
	}
	return 0
}

// Defines a general strategy for chunking text into smaller segments.
// Currently, only StaticChunkingStrategy is supported.
type ChunkingStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Strategy:
	//
	//	*ChunkingStrategy_StaticStrategy
	Strategy      isChunkingStrategy_Strategy `protobuf_oneof:"Strategy"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChunkingStrategy) Reset() {
	*x = ChunkingStrategy{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkingStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkingStrategy) ProtoMessage() {}

func (x *ChunkingStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkingStrategy.ProtoReflect.Descriptor instead.
func (*ChunkingStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{1}
}

func (x *ChunkingStrategy) GetStrategy() isChunkingStrategy_Strategy {
	if x != nil {
		return x.Strategy
	}
	return nil
}

func (x *ChunkingStrategy) GetStaticStrategy() *StaticChunkingStrategy {
	if x != nil {
		if x, ok := x.Strategy.(*ChunkingStrategy_StaticStrategy); ok {
			return x.StaticStrategy
		}
	}
	return nil
}

type isChunkingStrategy_Strategy interface {
	isChunkingStrategy_Strategy()
}

type ChunkingStrategy_StaticStrategy struct {
	StaticStrategy *StaticChunkingStrategy `protobuf:"bytes,1,opt,name=static_strategy,json=staticStrategy,proto3,oneof"`
}

func (*ChunkingStrategy_StaticStrategy) isChunkingStrategy_Strategy() {}

type MeanCombinationStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Technique for averaging relevance scores from different indices. Default is ARITHMETIC
	MeanEvaluationTechnique MeanCombinationStrategy_MeanEvaluationTechnique `protobuf:"varint,1,opt,name=mean_evaluation_technique,json=meanEvaluationTechnique,proto3,enum=yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy_MeanEvaluationTechnique" json:"mean_evaluation_technique,omitempty"`
	// Weights used for evaluating the weighted mean of relevance scores. The sum of the values must equal 1.0
	// If not provided, all scores are given equal weight
	Weights       []float64 `protobuf:"fixed64,2,rep,packed,name=weights,proto3" json:"weights,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MeanCombinationStrategy) Reset() {
	*x = MeanCombinationStrategy{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MeanCombinationStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeanCombinationStrategy) ProtoMessage() {}

func (x *MeanCombinationStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeanCombinationStrategy.ProtoReflect.Descriptor instead.
func (*MeanCombinationStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{2}
}

func (x *MeanCombinationStrategy) GetMeanEvaluationTechnique() MeanCombinationStrategy_MeanEvaluationTechnique {
	if x != nil {
		return x.MeanEvaluationTechnique
	}
	return MeanCombinationStrategy_MEAN_EVALUATION_TECHNIQUE_UNSPECIFIED
}

func (x *MeanCombinationStrategy) GetWeights() []float64 {
	if x != nil {
		return x.Weights
	}
	return nil
}

// https://plg.uwaterloo.ca/~gvcormac/cormacksigir09-rrf.pdf
type ReciprocalRankFusionCombinationStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The parameter k for RRFscore. Default is 60
	K             *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=k,proto3" json:"k,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReciprocalRankFusionCombinationStrategy) Reset() {
	*x = ReciprocalRankFusionCombinationStrategy{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReciprocalRankFusionCombinationStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReciprocalRankFusionCombinationStrategy) ProtoMessage() {}

func (x *ReciprocalRankFusionCombinationStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReciprocalRankFusionCombinationStrategy.ProtoReflect.Descriptor instead.
func (*ReciprocalRankFusionCombinationStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{3}
}

func (x *ReciprocalRankFusionCombinationStrategy) GetK() *wrapperspb.Int64Value {
	if x != nil {
		return x.K
	}
	return nil
}

// Combination strategy for merging rankings from different indices
type CombinationStrategy struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Strategy:
	//
	//	*CombinationStrategy_MeanCombination
	//	*CombinationStrategy_RrfCombination
	Strategy      isCombinationStrategy_Strategy `protobuf_oneof:"Strategy"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CombinationStrategy) Reset() {
	*x = CombinationStrategy{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CombinationStrategy) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CombinationStrategy) ProtoMessage() {}

func (x *CombinationStrategy) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CombinationStrategy.ProtoReflect.Descriptor instead.
func (*CombinationStrategy) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{4}
}

func (x *CombinationStrategy) GetStrategy() isCombinationStrategy_Strategy {
	if x != nil {
		return x.Strategy
	}
	return nil
}

func (x *CombinationStrategy) GetMeanCombination() *MeanCombinationStrategy {
	if x != nil {
		if x, ok := x.Strategy.(*CombinationStrategy_MeanCombination); ok {
			return x.MeanCombination
		}
	}
	return nil
}

func (x *CombinationStrategy) GetRrfCombination() *ReciprocalRankFusionCombinationStrategy {
	if x != nil {
		if x, ok := x.Strategy.(*CombinationStrategy_RrfCombination); ok {
			return x.RrfCombination
		}
	}
	return nil
}

type isCombinationStrategy_Strategy interface {
	isCombinationStrategy_Strategy()
}

type CombinationStrategy_MeanCombination struct {
	MeanCombination *MeanCombinationStrategy `protobuf:"bytes,1,opt,name=mean_combination,json=meanCombination,proto3,oneof"`
}

type CombinationStrategy_RrfCombination struct {
	RrfCombination *ReciprocalRankFusionCombinationStrategy `protobuf:"bytes,2,opt,name=rrf_combination,json=rrfCombination,proto3,oneof"`
}

func (*CombinationStrategy_MeanCombination) isCombinationStrategy_Strategy() {}

func (*CombinationStrategy_RrfCombination) isCombinationStrategy_Strategy() {}

// Configuration for the NgramTokenizer, which splits text into overlapping character sequences (n-grams) of specified lengths.
//
// Example:
// Input text: `hello`
// min_gram = 2, max_gram = 3
//
// Generated tokens:
// * For n = 2 (2-character n-grams): `he`, `el`, `ll`, `lo`
// * For n = 3 (3-character n-grams): `hel`, `ell`, `llo`
//
// Final tokens: `[he, el, ll, lo, hel, ell, llo]`
type NgramTokenizer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Minimum length of characters in a gram. Defaults to 3
	MinGram *wrapperspb.Int64Value `protobuf:"bytes,1,opt,name=min_gram,json=minGram,proto3" json:"min_gram,omitempty"`
	// Maximum length of characters in a gram. Defaults to 4
	MaxGram       *wrapperspb.Int64Value `protobuf:"bytes,2,opt,name=max_gram,json=maxGram,proto3" json:"max_gram,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NgramTokenizer) Reset() {
	*x = NgramTokenizer{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NgramTokenizer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NgramTokenizer) ProtoMessage() {}

func (x *NgramTokenizer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NgramTokenizer.ProtoReflect.Descriptor instead.
func (*NgramTokenizer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{5}
}

func (x *NgramTokenizer) GetMinGram() *wrapperspb.Int64Value {
	if x != nil {
		return x.MinGram
	}
	return nil
}

func (x *NgramTokenizer) GetMaxGram() *wrapperspb.Int64Value {
	if x != nil {
		return x.MaxGram
	}
	return nil
}

// A standard tokenizer that splits text on word boundaries and removes punctuation.
// It follows the Unicode Text Segmentation rules as specified in Unicode Standard Annex #29.
//
// Example:
// Input text: `Hello, world! How are you?`
// Output tokens: `[Hello, world, How, are, you]`
type StandardTokenizer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StandardTokenizer) Reset() {
	*x = StandardTokenizer{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StandardTokenizer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardTokenizer) ProtoMessage() {}

func (x *StandardTokenizer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardTokenizer.ProtoReflect.Descriptor instead.
func (*StandardTokenizer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{6}
}

// A standard analyzer that uses StandardTokenizer.
type StandardAnalyzer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StandardAnalyzer) Reset() {
	*x = StandardAnalyzer{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StandardAnalyzer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StandardAnalyzer) ProtoMessage() {}

func (x *StandardAnalyzer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StandardAnalyzer.ProtoReflect.Descriptor instead.
func (*StandardAnalyzer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{7}
}

// A specialized analyzer that uses Yandex's lemmatization technology to reduce words to their base forms.
// Particularly effective for Russian and other Slavic languages, handling their complex morphology.
// For more information, see:
// https://yandex.cloud/en/docs/tutorials/dataplatform/opensearch-yandex-lemmer
type YandexLemmerAnalyzer struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *YandexLemmerAnalyzer) Reset() {
	*x = YandexLemmerAnalyzer{}
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *YandexLemmerAnalyzer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*YandexLemmerAnalyzer) ProtoMessage() {}

func (x *YandexLemmerAnalyzer) ProtoReflect() protoreflect.Message {
	mi := &file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use YandexLemmerAnalyzer.ProtoReflect.Descriptor instead.
func (*YandexLemmerAnalyzer) Descriptor() ([]byte, []int) {
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP(), []int{8}
}

var File_yandex_cloud_ai_assistants_v1_searchindex_common_proto protoreflect.FileDescriptor

const file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDesc = "" +
	"\n" +
	"6yandex/cloud/ai/assistants/v1/searchindex/common.proto\x12)yandex.cloud.ai.assistants.v1.searchindex\x1a\x1egoogle/protobuf/wrappers.proto\"}\n" +
	"\x16StaticChunkingStrategy\x121\n" +
	"\x15max_chunk_size_tokens\x18\x01 \x01(\x03R\x12maxChunkSizeTokens\x120\n" +
	"\x14chunk_overlap_tokens\x18\x02 \x01(\x03R\x12chunkOverlapTokens\"\x8c\x01\n" +
	"\x10ChunkingStrategy\x12l\n" +
	"\x0fstatic_strategy\x18\x01 \x01(\v2A.yandex.cloud.ai.assistants.v1.searchindex.StaticChunkingStrategyH\x00R\x0estaticStrategyB\n" +
	"\n" +
	"\bStrategy\"\xbf\x02\n" +
	"\x17MeanCombinationStrategy\x12\x96\x01\n" +
	"\x19mean_evaluation_technique\x18\x01 \x01(\x0e2Z.yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy.MeanEvaluationTechniqueR\x17meanEvaluationTechnique\x12\x18\n" +
	"\aweights\x18\x02 \x03(\x01R\aweights\"q\n" +
	"\x17MeanEvaluationTechnique\x12)\n" +
	"%MEAN_EVALUATION_TECHNIQUE_UNSPECIFIED\x10\x00\x12\x0e\n" +
	"\n" +
	"ARITHMETIC\x10\x01\x12\r\n" +
	"\tGEOMETRIC\x10\x02\x12\f\n" +
	"\bHARMONIC\x10\x03\"T\n" +
	"'ReciprocalRankFusionCombinationStrategy\x12)\n" +
	"\x01k\x18\x01 \x01(\v2\x1b.google.protobuf.Int64ValueR\x01k\"\x91\x02\n" +
	"\x13CombinationStrategy\x12o\n" +
	"\x10mean_combination\x18\x01 \x01(\v2B.yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategyH\x00R\x0fmeanCombination\x12}\n" +
	"\x0frrf_combination\x18\x02 \x01(\v2R.yandex.cloud.ai.assistants.v1.searchindex.ReciprocalRankFusionCombinationStrategyH\x00R\x0errfCombinationB\n" +
	"\n" +
	"\bStrategy\"\x80\x01\n" +
	"\x0eNgramTokenizer\x126\n" +
	"\bmin_gram\x18\x01 \x01(\v2\x1b.google.protobuf.Int64ValueR\aminGram\x126\n" +
	"\bmax_gram\x18\x02 \x01(\v2\x1b.google.protobuf.Int64ValueR\amaxGram\"\x13\n" +
	"\x11StandardTokenizer\"\x12\n" +
	"\x10StandardAnalyzer\"\x16\n" +
	"\x14YandexLemmerAnalyzer*T\n" +
	"\x15NormalizationStrategy\x12&\n" +
	"\"NORMALIZATION_STRATEGY_UNSPECIFIED\x10\x00\x12\v\n" +
	"\aMIN_MAX\x10\x01\x12\x06\n" +
	"\x02L2\x10\x02B\x8a\x01\n" +
	"-yandex.cloud.api.ai.assistants.v1.searchindexZYgithub.com/yandex-cloud/go-genproto/yandex/cloud/ai/assistants/v1/searchindex;searchindexb\x06proto3"

var (
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescOnce sync.Once
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescData []byte
)

func file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescGZIP() []byte {
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescOnce.Do(func() {
		file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDesc)))
	})
	return file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDescData
}

var file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_goTypes = []any{
	(NormalizationStrategy)(0),                           // 0: yandex.cloud.ai.assistants.v1.searchindex.NormalizationStrategy
	(MeanCombinationStrategy_MeanEvaluationTechnique)(0), // 1: yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy.MeanEvaluationTechnique
	(*StaticChunkingStrategy)(nil),                       // 2: yandex.cloud.ai.assistants.v1.searchindex.StaticChunkingStrategy
	(*ChunkingStrategy)(nil),                             // 3: yandex.cloud.ai.assistants.v1.searchindex.ChunkingStrategy
	(*MeanCombinationStrategy)(nil),                      // 4: yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy
	(*ReciprocalRankFusionCombinationStrategy)(nil),      // 5: yandex.cloud.ai.assistants.v1.searchindex.ReciprocalRankFusionCombinationStrategy
	(*CombinationStrategy)(nil),                          // 6: yandex.cloud.ai.assistants.v1.searchindex.CombinationStrategy
	(*NgramTokenizer)(nil),                               // 7: yandex.cloud.ai.assistants.v1.searchindex.NgramTokenizer
	(*StandardTokenizer)(nil),                            // 8: yandex.cloud.ai.assistants.v1.searchindex.StandardTokenizer
	(*StandardAnalyzer)(nil),                             // 9: yandex.cloud.ai.assistants.v1.searchindex.StandardAnalyzer
	(*YandexLemmerAnalyzer)(nil),                         // 10: yandex.cloud.ai.assistants.v1.searchindex.YandexLemmerAnalyzer
	(*wrapperspb.Int64Value)(nil),                        // 11: google.protobuf.Int64Value
}
var file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_depIdxs = []int32{
	2,  // 0: yandex.cloud.ai.assistants.v1.searchindex.ChunkingStrategy.static_strategy:type_name -> yandex.cloud.ai.assistants.v1.searchindex.StaticChunkingStrategy
	1,  // 1: yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy.mean_evaluation_technique:type_name -> yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy.MeanEvaluationTechnique
	11, // 2: yandex.cloud.ai.assistants.v1.searchindex.ReciprocalRankFusionCombinationStrategy.k:type_name -> google.protobuf.Int64Value
	4,  // 3: yandex.cloud.ai.assistants.v1.searchindex.CombinationStrategy.mean_combination:type_name -> yandex.cloud.ai.assistants.v1.searchindex.MeanCombinationStrategy
	5,  // 4: yandex.cloud.ai.assistants.v1.searchindex.CombinationStrategy.rrf_combination:type_name -> yandex.cloud.ai.assistants.v1.searchindex.ReciprocalRankFusionCombinationStrategy
	11, // 5: yandex.cloud.ai.assistants.v1.searchindex.NgramTokenizer.min_gram:type_name -> google.protobuf.Int64Value
	11, // 6: yandex.cloud.ai.assistants.v1.searchindex.NgramTokenizer.max_gram:type_name -> google.protobuf.Int64Value
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_init() }
func file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_init() {
	if File_yandex_cloud_ai_assistants_v1_searchindex_common_proto != nil {
		return
	}
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[1].OneofWrappers = []any{
		(*ChunkingStrategy_StaticStrategy)(nil),
	}
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes[4].OneofWrappers = []any{
		(*CombinationStrategy_MeanCombination)(nil),
		(*CombinationStrategy_RrfCombination)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDesc), len(file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_goTypes,
		DependencyIndexes: file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_depIdxs,
		EnumInfos:         file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_enumTypes,
		MessageInfos:      file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_msgTypes,
	}.Build()
	File_yandex_cloud_ai_assistants_v1_searchindex_common_proto = out.File
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_goTypes = nil
	file_yandex_cloud_ai_assistants_v1_searchindex_common_proto_depIdxs = nil
}
