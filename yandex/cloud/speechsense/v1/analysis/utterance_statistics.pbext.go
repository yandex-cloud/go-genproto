// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

func (m *UtteranceStatistics) SetSpeakerTag(v string) {
	m.SpeakerTag = v
}

func (m *UtteranceStatistics) SetSpeechBoundaries(v *AudioSegmentBoundaries) {
	m.SpeechBoundaries = v
}

func (m *UtteranceStatistics) SetTotalSpeechMs(v int64) {
	m.TotalSpeechMs = v
}

func (m *UtteranceStatistics) SetSpeechRatio(v float64) {
	m.SpeechRatio = v
}

func (m *UtteranceStatistics) SetTotalSilenceMs(v int64) {
	m.TotalSilenceMs = v
}

func (m *UtteranceStatistics) SetSilenceRatio(v float64) {
	m.SilenceRatio = v
}

func (m *UtteranceStatistics) SetWordsCount(v int64) {
	m.WordsCount = v
}

func (m *UtteranceStatistics) SetLettersCount(v int64) {
	m.LettersCount = v
}

func (m *UtteranceStatistics) SetWordsPerSecond(v *DescriptiveStatistics) {
	m.WordsPerSecond = v
}

func (m *UtteranceStatistics) SetLettersPerSecond(v *DescriptiveStatistics) {
	m.LettersPerSecond = v
}
