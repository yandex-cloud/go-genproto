// Code generated by protoc-gen-goext. DO NOT EDIT.

package fomo

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

type BatchCompletionRequest_StructuredOutput = isBatchCompletionRequest_StructuredOutput

func (m *BatchCompletionRequest) SetStructuredOutput(v BatchCompletionRequest_StructuredOutput) {
	m.StructuredOutput = v
}

func (m *BatchCompletionRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *BatchCompletionRequest) SetSourceDatasetId(v string) {
	m.SourceDatasetId = v
}

func (m *BatchCompletionRequest) SetCompletionOptions(v *CompletionOptions) {
	m.CompletionOptions = v
}

func (m *BatchCompletionRequest) SetDataLoggingEnabled(v bool) {
	m.DataLoggingEnabled = v
}

func (m *BatchCompletionRequest) SetJsonObject(v bool) {
	m.StructuredOutput = &BatchCompletionRequest_JsonObject{
		JsonObject: v,
	}
}

func (m *BatchCompletionRequest) SetJsonSchema(v *JsonSchema) {
	m.StructuredOutput = &BatchCompletionRequest_JsonSchema{
		JsonSchema: v,
	}
}

func (m *CompletionOptions) SetTemperature(v *wrapperspb.DoubleValue) {
	m.Temperature = v
}

func (m *CompletionOptions) SetMaxTokens(v *wrapperspb.Int64Value) {
	m.MaxTokens = v
}

func (m *CompletionOptions) SetReasoningOptions(v *ReasoningOptions) {
	m.ReasoningOptions = v
}

func (m *ReasoningOptions) SetMode(v ReasoningOptions_ReasoningMode) {
	m.Mode = v
}

func (m *JsonSchema) SetSchema(v *structpb.Struct) {
	m.Schema = v
}
