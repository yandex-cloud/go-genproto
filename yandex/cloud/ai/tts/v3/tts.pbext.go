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

func (m *TextTemplate) SetTextTemplate(v string) {
	m.TextTemplate = v
}

func (m *TextTemplate) SetVariables(v []*TextVariable) {
	m.Variables = v
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
