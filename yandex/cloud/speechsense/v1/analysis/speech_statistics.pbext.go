// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

func (m *SpeechStatistics) SetTotalSimultaneousSpeechDurationSeconds(v int64) {
	m.TotalSimultaneousSpeechDurationSeconds = v
}

func (m *SpeechStatistics) SetTotalSimultaneousSpeechDurationMs(v int64) {
	m.TotalSimultaneousSpeechDurationMs = v
}

func (m *SpeechStatistics) SetTotalSimultaneousSpeechRatio(v float64) {
	m.TotalSimultaneousSpeechRatio = v
}

func (m *SpeechStatistics) SetSimultaneousSpeechDurationEstimation(v *DescriptiveStatistics) {
	m.SimultaneousSpeechDurationEstimation = v
}