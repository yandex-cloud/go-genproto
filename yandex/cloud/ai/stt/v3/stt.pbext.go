// Code generated by protoc-gen-goext. DO NOT EDIT.

package stt

func (m *TextNormalizationOptions) SetTextNormalization(v TextNormalizationOptions_TextNormalization) {
	m.TextNormalization = v
}

func (m *TextNormalizationOptions) SetProfanityFilter(v bool) {
	m.ProfanityFilter = v
}

func (m *TextNormalizationOptions) SetLiteratureText(v bool) {
	m.LiteratureText = v
}

func (m *TextNormalizationOptions) SetPhoneFormattingMode(v TextNormalizationOptions_PhoneFormattingMode) {
	m.PhoneFormattingMode = v
}

func (m *DefaultEouClassifier) SetType(v DefaultEouClassifier_EouSensitivity) {
	m.Type = v
}

func (m *DefaultEouClassifier) SetMaxPauseBetweenWordsHintMs(v int64) {
	m.MaxPauseBetweenWordsHintMs = v
}

type EouClassifierOptions_Classifier = isEouClassifierOptions_Classifier

func (m *EouClassifierOptions) SetClassifier(v EouClassifierOptions_Classifier) {
	m.Classifier = v
}

func (m *EouClassifierOptions) SetDefaultClassifier(v *DefaultEouClassifier) {
	m.Classifier = &EouClassifierOptions_DefaultClassifier{
		DefaultClassifier: v,
	}
}

func (m *EouClassifierOptions) SetExternalClassifier(v *ExternalEouClassifier) {
	m.Classifier = &EouClassifierOptions_ExternalClassifier{
		ExternalClassifier: v,
	}
}

func (m *RawAudio) SetAudioEncoding(v RawAudio_AudioEncoding) {
	m.AudioEncoding = v
}

func (m *RawAudio) SetSampleRateHertz(v int64) {
	m.SampleRateHertz = v
}

func (m *RawAudio) SetAudioChannelCount(v int64) {
	m.AudioChannelCount = v
}

func (m *ContainerAudio) SetContainerAudioType(v ContainerAudio_ContainerAudioType) {
	m.ContainerAudioType = v
}

type AudioFormatOptions_AudioFormat = isAudioFormatOptions_AudioFormat

func (m *AudioFormatOptions) SetAudioFormat(v AudioFormatOptions_AudioFormat) {
	m.AudioFormat = v
}

func (m *AudioFormatOptions) SetRawAudio(v *RawAudio) {
	m.AudioFormat = &AudioFormatOptions_RawAudio{
		RawAudio: v,
	}
}

func (m *AudioFormatOptions) SetContainerAudio(v *ContainerAudio) {
	m.AudioFormat = &AudioFormatOptions_ContainerAudio{
		ContainerAudio: v,
	}
}

func (m *LanguageRestrictionOptions) SetRestrictionType(v LanguageRestrictionOptions_LanguageRestrictionType) {
	m.RestrictionType = v
}

func (m *LanguageRestrictionOptions) SetLanguageCode(v []string) {
	m.LanguageCode = v
}

func (m *RecognitionModelOptions) SetModel(v string) {
	m.Model = v
}

func (m *RecognitionModelOptions) SetAudioFormat(v *AudioFormatOptions) {
	m.AudioFormat = v
}

func (m *RecognitionModelOptions) SetTextNormalization(v *TextNormalizationOptions) {
	m.TextNormalization = v
}

func (m *RecognitionModelOptions) SetLanguageRestriction(v *LanguageRestrictionOptions) {
	m.LanguageRestriction = v
}

func (m *RecognitionModelOptions) SetAudioProcessingType(v RecognitionModelOptions_AudioProcessingType) {
	m.AudioProcessingType = v
}

func (m *StreamingOptions) SetRecognitionModel(v *RecognitionModelOptions) {
	m.RecognitionModel = v
}

func (m *StreamingOptions) SetEouClassifier(v *EouClassifierOptions) {
	m.EouClassifier = v
}

func (m *AudioChunk) SetData(v []byte) {
	m.Data = v
}

func (m *SilenceChunk) SetDurationMs(v int64) {
	m.DurationMs = v
}

type StreamingRequest_Event = isStreamingRequest_Event

func (m *StreamingRequest) SetEvent(v StreamingRequest_Event) {
	m.Event = v
}

func (m *StreamingRequest) SetSessionOptions(v *StreamingOptions) {
	m.Event = &StreamingRequest_SessionOptions{
		SessionOptions: v,
	}
}

func (m *StreamingRequest) SetChunk(v *AudioChunk) {
	m.Event = &StreamingRequest_Chunk{
		Chunk: v,
	}
}

func (m *StreamingRequest) SetSilenceChunk(v *SilenceChunk) {
	m.Event = &StreamingRequest_SilenceChunk{
		SilenceChunk: v,
	}
}

func (m *StreamingRequest) SetEou(v *Eou) {
	m.Event = &StreamingRequest_Eou{
		Eou: v,
	}
}

func (m *Word) SetText(v string) {
	m.Text = v
}

func (m *Word) SetStartTimeMs(v int64) {
	m.StartTimeMs = v
}

func (m *Word) SetEndTimeMs(v int64) {
	m.EndTimeMs = v
}

func (m *LanguageEstimation) SetLanguageCode(v string) {
	m.LanguageCode = v
}

func (m *LanguageEstimation) SetProbability(v float64) {
	m.Probability = v
}

func (m *Alternative) SetWords(v []*Word) {
	m.Words = v
}

func (m *Alternative) SetText(v string) {
	m.Text = v
}

func (m *Alternative) SetStartTimeMs(v int64) {
	m.StartTimeMs = v
}

func (m *Alternative) SetEndTimeMs(v int64) {
	m.EndTimeMs = v
}

func (m *Alternative) SetConfidence(v float64) {
	m.Confidence = v
}

func (m *Alternative) SetLanguages(v []*LanguageEstimation) {
	m.Languages = v
}

func (m *EouUpdate) SetTimeMs(v int64) {
	m.TimeMs = v
}

func (m *AlternativeUpdate) SetAlternatives(v []*Alternative) {
	m.Alternatives = v
}

func (m *AlternativeUpdate) SetChannelTag(v string) {
	m.ChannelTag = v
}

func (m *AudioCursors) SetReceivedDataMs(v int64) {
	m.ReceivedDataMs = v
}

func (m *AudioCursors) SetResetTimeMs(v int64) {
	m.ResetTimeMs = v
}

func (m *AudioCursors) SetPartialTimeMs(v int64) {
	m.PartialTimeMs = v
}

func (m *AudioCursors) SetFinalTimeMs(v int64) {
	m.FinalTimeMs = v
}

func (m *AudioCursors) SetFinalIndex(v int64) {
	m.FinalIndex = v
}

func (m *AudioCursors) SetEouTimeMs(v int64) {
	m.EouTimeMs = v
}

type FinalRefinement_Type = isFinalRefinement_Type

func (m *FinalRefinement) SetType(v FinalRefinement_Type) {
	m.Type = v
}

func (m *FinalRefinement) SetFinalIndex(v int64) {
	m.FinalIndex = v
}

func (m *FinalRefinement) SetNormalizedText(v *AlternativeUpdate) {
	m.Type = &FinalRefinement_NormalizedText{
		NormalizedText: v,
	}
}

func (m *StatusCode) SetCodeType(v CodeType) {
	m.CodeType = v
}

func (m *StatusCode) SetMessage(v string) {
	m.Message = v
}

func (m *SessionUuid) SetUuid(v string) {
	m.Uuid = v
}

func (m *SessionUuid) SetUserRequestId(v string) {
	m.UserRequestId = v
}

type StreamingResponse_Event = isStreamingResponse_Event

func (m *StreamingResponse) SetEvent(v StreamingResponse_Event) {
	m.Event = v
}

func (m *StreamingResponse) SetSessionUuid(v *SessionUuid) {
	m.SessionUuid = v
}

func (m *StreamingResponse) SetAudioCursors(v *AudioCursors) {
	m.AudioCursors = v
}

func (m *StreamingResponse) SetResponseWallTimeMs(v int64) {
	m.ResponseWallTimeMs = v
}

func (m *StreamingResponse) SetPartial(v *AlternativeUpdate) {
	m.Event = &StreamingResponse_Partial{
		Partial: v,
	}
}

func (m *StreamingResponse) SetFinal(v *AlternativeUpdate) {
	m.Event = &StreamingResponse_Final{
		Final: v,
	}
}

func (m *StreamingResponse) SetEouUpdate(v *EouUpdate) {
	m.Event = &StreamingResponse_EouUpdate{
		EouUpdate: v,
	}
}

func (m *StreamingResponse) SetFinalRefinement(v *FinalRefinement) {
	m.Event = &StreamingResponse_FinalRefinement{
		FinalRefinement: v,
	}
}

func (m *StreamingResponse) SetStatusCode(v *StatusCode) {
	m.Event = &StreamingResponse_StatusCode{
		StatusCode: v,
	}
}
