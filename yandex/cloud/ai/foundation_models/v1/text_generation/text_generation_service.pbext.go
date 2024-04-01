// Code generated by protoc-gen-goext. DO NOT EDIT.

package foundation_models

import (
	v1 "github.com/yandex-cloud/go-genproto/yandex/cloud/ai/foundation_models/v1"
)

func (m *CompletionRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *CompletionRequest) SetCompletionOptions(v *v1.CompletionOptions) {
	m.CompletionOptions = v
}

func (m *CompletionRequest) SetMessages(v []*v1.Message) {
	m.Messages = v
}

func (m *CompletionResponse) SetAlternatives(v []*v1.Alternative) {
	m.Alternatives = v
}

func (m *CompletionResponse) SetUsage(v *v1.ContentUsage) {
	m.Usage = v
}

func (m *CompletionResponse) SetModelVersion(v string) {
	m.ModelVersion = v
}

func (m *TokenizeRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *TokenizeRequest) SetText(v string) {
	m.Text = v
}

func (m *TokenizeResponse) SetTokens(v []*v1.Token) {
	m.Tokens = v
}

func (m *TokenizeResponse) SetModelVersion(v string) {
	m.ModelVersion = v
}