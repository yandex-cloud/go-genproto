// Code generated by protoc-gen-goext. DO NOT EDIT.

package text_classification

func (m *TextClassificationRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *TextClassificationRequest) SetText(v string) {
	m.Text = v
}

func (m *TextClassificationResponse) SetPredictions(v []*ClassificationLabel) {
	m.Predictions = v
}

func (m *TextClassificationResponse) SetModelVersion(v string) {
	m.ModelVersion = v
}

func (m *TextClassificationResponse) SetInputTokens(v int64) {
	m.InputTokens = v
}

func (m *FewShotTextClassificationRequest) SetModelUri(v string) {
	m.ModelUri = v
}

func (m *FewShotTextClassificationRequest) SetTaskDescription(v string) {
	m.TaskDescription = v
}

func (m *FewShotTextClassificationRequest) SetLabels(v []string) {
	m.Labels = v
}

func (m *FewShotTextClassificationRequest) SetText(v string) {
	m.Text = v
}

func (m *FewShotTextClassificationRequest) SetSamples(v []*ClassificationSample) {
	m.Samples = v
}

func (m *FewShotTextClassificationResponse) SetPredictions(v []*ClassificationLabel) {
	m.Predictions = v
}

func (m *FewShotTextClassificationResponse) SetModelVersion(v string) {
	m.ModelVersion = v
}

func (m *FewShotTextClassificationResponse) SetInputTokens(v int64) {
	m.InputTokens = v
}
