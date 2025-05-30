// Code generated by protoc-gen-goext. DO NOT EDIT.

package tts

type AudioContent_AudioSource = isAudioContent_AudioSource

func (m *AudioContent) SetAudioSource(v AudioContent_AudioSource) {
	m.AudioSource = v
}

func (m *AudioContent) SetContent(v []byte) {
	m.AudioSource = &AudioContent_Content{
		Content: v,
	}
}

func (m *AudioContent) SetAudioSpec(v *AudioFormatOptions) {
	m.AudioSpec = v
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

func (m *RawAudio) SetAudioEncoding(v RawAudio_AudioEncoding) {
	m.AudioEncoding = v
}

func (m *RawAudio) SetSampleRateHertz(v int64) {
	m.SampleRateHertz = v
}

func (m *ContainerAudio) SetContainerAudioType(v ContainerAudio_ContainerAudioType) {
	m.ContainerAudioType = v
}

func (m *TextVariable) SetVariableName(v string) {
	m.VariableName = v
}

func (m *TextVariable) SetVariableValue(v string) {
	m.VariableValue = v
}

func (m *AudioVariable) SetVariableName(v string) {
	m.VariableName = v
}

func (m *AudioVariable) SetVariableStartMs(v int64) {
	m.VariableStartMs = v
}

func (m *AudioVariable) SetVariableLengthMs(v int64) {
	m.VariableLengthMs = v
}

func (m *UtteranceSynthesisResponse) SetAudioChunk(v *AudioChunk) {
	m.AudioChunk = v
}

func (m *UtteranceSynthesisResponse) SetTextChunk(v *TextChunk) {
	m.TextChunk = v
}

func (m *UtteranceSynthesisResponse) SetStartMs(v int64) {
	m.StartMs = v
}

func (m *UtteranceSynthesisResponse) SetLengthMs(v int64) {
	m.LengthMs = v
}

func (m *AudioTemplate) SetAudio(v *AudioContent) {
	m.Audio = v
}

func (m *AudioTemplate) SetTextTemplate(v *TextTemplate) {
	m.TextTemplate = v
}

func (m *AudioTemplate) SetVariables(v []*AudioVariable) {
	m.Variables = v
}

func (m *AudioChunk) SetData(v []byte) {
	m.Data = v
}

func (m *TextChunk) SetText(v string) {
	m.Text = v
}

func (m *TextTemplate) SetTextTemplate(v string) {
	m.TextTemplate = v
}

func (m *TextTemplate) SetVariables(v []*TextVariable) {
	m.Variables = v
}

func (m *DurationHint) SetPolicy(v DurationHint_DurationHintPolicy) {
	m.Policy = v
}

func (m *DurationHint) SetDurationMs(v int64) {
	m.DurationMs = v
}

type Hints_Hint = isHints_Hint

func (m *Hints) SetHint(v Hints_Hint) {
	m.Hint = v
}

func (m *Hints) SetVoice(v string) {
	m.Hint = &Hints_Voice{
		Voice: v,
	}
}

func (m *Hints) SetAudioTemplate(v *AudioTemplate) {
	m.Hint = &Hints_AudioTemplate{
		AudioTemplate: v,
	}
}

func (m *Hints) SetSpeed(v float64) {
	m.Hint = &Hints_Speed{
		Speed: v,
	}
}

func (m *Hints) SetVolume(v float64) {
	m.Hint = &Hints_Volume{
		Volume: v,
	}
}

func (m *Hints) SetRole(v string) {
	m.Hint = &Hints_Role{
		Role: v,
	}
}

func (m *Hints) SetPitchShift(v float64) {
	m.Hint = &Hints_PitchShift{
		PitchShift: v,
	}
}

func (m *Hints) SetDuration(v *DurationHint) {
	m.Hint = &Hints_Duration{
		Duration: v,
	}
}

type UtteranceSynthesisRequest_Utterance = isUtteranceSynthesisRequest_Utterance

func (m *UtteranceSynthesisRequest) SetUtterance(v UtteranceSynthesisRequest_Utterance) {
	m.Utterance = v
}

func (m *UtteranceSynthesisRequest) SetModel(v string) {
	m.Model = v
}

func (m *UtteranceSynthesisRequest) SetText(v string) {
	m.Utterance = &UtteranceSynthesisRequest_Text{
		Text: v,
	}
}

func (m *UtteranceSynthesisRequest) SetTextTemplate(v *TextTemplate) {
	m.Utterance = &UtteranceSynthesisRequest_TextTemplate{
		TextTemplate: v,
	}
}

func (m *UtteranceSynthesisRequest) SetHints(v []*Hints) {
	m.Hints = v
}

func (m *UtteranceSynthesisRequest) SetOutputAudioSpec(v *AudioFormatOptions) {
	m.OutputAudioSpec = v
}

func (m *UtteranceSynthesisRequest) SetLoudnessNormalizationType(v UtteranceSynthesisRequest_LoudnessNormalizationType) {
	m.LoudnessNormalizationType = v
}

func (m *UtteranceSynthesisRequest) SetUnsafeMode(v bool) {
	m.UnsafeMode = v
}

func (m *SynthesisOptions) SetModel(v string) {
	m.Model = v
}

func (m *SynthesisOptions) SetVoice(v string) {
	m.Voice = v
}

func (m *SynthesisOptions) SetRole(v string) {
	m.Role = v
}

func (m *SynthesisOptions) SetSpeed(v float64) {
	m.Speed = v
}

func (m *SynthesisOptions) SetVolume(v float64) {
	m.Volume = v
}

func (m *SynthesisOptions) SetPitchShift(v float64) {
	m.PitchShift = v
}

func (m *SynthesisOptions) SetOutputAudioSpec(v *AudioFormatOptions) {
	m.OutputAudioSpec = v
}

func (m *SynthesisOptions) SetLoudnessNormalizationType(v LoudnessNormalizationType) {
	m.LoudnessNormalizationType = v
}

func (m *SynthesisInput) SetText(v string) {
	m.Text = v
}

type StreamSynthesisRequest_Event = isStreamSynthesisRequest_Event

func (m *StreamSynthesisRequest) SetEvent(v StreamSynthesisRequest_Event) {
	m.Event = v
}

func (m *StreamSynthesisRequest) SetOptions(v *SynthesisOptions) {
	m.Event = &StreamSynthesisRequest_Options{
		Options: v,
	}
}

func (m *StreamSynthesisRequest) SetSynthesisInput(v *SynthesisInput) {
	m.Event = &StreamSynthesisRequest_SynthesisInput{
		SynthesisInput: v,
	}
}

func (m *StreamSynthesisRequest) SetForceSynthesis(v *ForceSynthesisEvent) {
	m.Event = &StreamSynthesisRequest_ForceSynthesis{
		ForceSynthesis: v,
	}
}

func (m *StreamSynthesisResponse) SetAudioChunk(v *AudioChunk) {
	m.AudioChunk = v
}

func (m *StreamSynthesisResponse) SetTextChunk(v *TextChunk) {
	m.TextChunk = v
}

func (m *StreamSynthesisResponse) SetStartMs(v int64) {
	m.StartMs = v
}

func (m *StreamSynthesisResponse) SetLengthMs(v int64) {
	m.LengthMs = v
}
