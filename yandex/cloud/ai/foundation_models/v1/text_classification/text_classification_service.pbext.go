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
