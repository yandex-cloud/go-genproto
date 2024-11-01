// Code generated by protoc-gen-goext. DO NOT EDIT.

package foundation_models

import (
	structpb "google.golang.org/protobuf/types/known/structpb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *CompletionOptions) SetStream(v bool) {
	m.Stream = v
}

func (m *CompletionOptions) SetTemperature(v *wrapperspb.DoubleValue) {
	m.Temperature = v
}

func (m *CompletionOptions) SetMaxTokens(v *wrapperspb.Int64Value) {
	m.MaxTokens = v
}

type Message_Content = isMessage_Content

func (m *Message) SetContent(v Message_Content) {
	m.Content = v
}

func (m *Message) SetRole(v string) {
	m.Role = v
}

func (m *Message) SetText(v string) {
	m.Content = &Message_Text{
		Text: v,
	}
}

func (m *Message) SetToolCallList(v *ToolCallList) {
	m.Content = &Message_ToolCallList{
		ToolCallList: v,
	}
}

func (m *Message) SetToolResultList(v *ToolResultList) {
	m.Content = &Message_ToolResultList{
		ToolResultList: v,
	}
}

func (m *ContentUsage) SetInputTextTokens(v int64) {
	m.InputTextTokens = v
}

func (m *ContentUsage) SetCompletionTokens(v int64) {
	m.CompletionTokens = v
}

func (m *ContentUsage) SetTotalTokens(v int64) {
	m.TotalTokens = v
}

func (m *Alternative) SetMessage(v *Message) {
	m.Message = v
}

func (m *Alternative) SetStatus(v Alternative_AlternativeStatus) {
	m.Status = v
}

func (m *Token) SetId(v int64) {
	m.Id = v
}

func (m *Token) SetText(v string) {
	m.Text = v
}

func (m *Token) SetSpecial(v bool) {
	m.Special = v
}

type Tool_ToolType = isTool_ToolType

func (m *Tool) SetToolType(v Tool_ToolType) {
	m.ToolType = v
}

func (m *Tool) SetFunction(v *FunctionTool) {
	m.ToolType = &Tool_Function{
		Function: v,
	}
}

func (m *FunctionTool) SetName(v string) {
	m.Name = v
}

func (m *FunctionTool) SetDescription(v string) {
	m.Description = v
}

func (m *FunctionTool) SetParameters(v *structpb.Struct) {
	m.Parameters = v
}

type ToolCall_ToolCallType = isToolCall_ToolCallType

func (m *ToolCall) SetToolCallType(v ToolCall_ToolCallType) {
	m.ToolCallType = v
}

func (m *ToolCall) SetFunctionCall(v *FunctionCall) {
	m.ToolCallType = &ToolCall_FunctionCall{
		FunctionCall: v,
	}
}

func (m *FunctionCall) SetName(v string) {
	m.Name = v
}

func (m *FunctionCall) SetArguments(v *structpb.Struct) {
	m.Arguments = v
}

func (m *ToolCallList) SetToolCalls(v []*ToolCall) {
	m.ToolCalls = v
}

type ToolResult_ToolResultType = isToolResult_ToolResultType

func (m *ToolResult) SetToolResultType(v ToolResult_ToolResultType) {
	m.ToolResultType = v
}

func (m *ToolResult) SetFunctionResult(v *FunctionResult) {
	m.ToolResultType = &ToolResult_FunctionResult{
		FunctionResult: v,
	}
}

type FunctionResult_ContentType = isFunctionResult_ContentType

func (m *FunctionResult) SetContentType(v FunctionResult_ContentType) {
	m.ContentType = v
}

func (m *FunctionResult) SetName(v string) {
	m.Name = v
}

func (m *FunctionResult) SetContent(v string) {
	m.ContentType = &FunctionResult_Content{
		Content: v,
	}
}

func (m *ToolResultList) SetToolResults(v []*ToolResult) {
	m.ToolResults = v
}
