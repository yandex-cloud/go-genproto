// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Transcription) SetPhrases(v []*Phrase) {
	m.Phrases = v
}

func (m *Transcription) SetAlgorithmsMetadata(v []*AlgorithmMetadata) {
	m.AlgorithmsMetadata = v
}

func (m *Phrase) SetChannelNumber(v int64) {
	m.ChannelNumber = v
}

func (m *Phrase) SetStartTimeMs(v int64) {
	m.StartTimeMs = v
}

func (m *Phrase) SetEndTimeMs(v int64) {
	m.EndTimeMs = v
}

func (m *Phrase) SetPhrase(v *PhraseText) {
	m.Phrase = v
}

func (m *Phrase) SetStatistics(v *PhraseStatistics) {
	m.Statistics = v
}

func (m *Phrase) SetClassifiers(v []*RecognitionClassifierResult) {
	m.Classifiers = v
}

func (m *PhraseText) SetText(v string) {
	m.Text = v
}

func (m *PhraseText) SetLanguage(v string) {
	m.Language = v
}

func (m *PhraseText) SetNormalizedText(v string) {
	m.NormalizedText = v
}

func (m *PhraseText) SetWords(v []*Word) {
	m.Words = v
}

func (m *Word) SetWord(v string) {
	m.Word = v
}

func (m *Word) SetStartTimeMs(v int64) {
	m.StartTimeMs = v
}

func (m *Word) SetEndTimeMs(v int64) {
	m.EndTimeMs = v
}

func (m *AlgorithmMetadata) SetCreatedTaskDate(v *timestamppb.Timestamp) {
	m.CreatedTaskDate = v
}

func (m *AlgorithmMetadata) SetCompletedTaskDate(v *timestamppb.Timestamp) {
	m.CompletedTaskDate = v
}

func (m *AlgorithmMetadata) SetError(v *Error) {
	m.Error = v
}

func (m *AlgorithmMetadata) SetTraceId(v string) {
	m.TraceId = v
}

func (m *AlgorithmMetadata) SetName(v string) {
	m.Name = v
}

func (m *Error) SetCode(v string) {
	m.Code = v
}

func (m *Error) SetMessage(v string) {
	m.Message = v
}

func (m *PhraseStatistics) SetStatistics(v *UtteranceStatistics) {
	m.Statistics = v
}
