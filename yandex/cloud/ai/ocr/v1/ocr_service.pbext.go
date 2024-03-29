// Code generated by protoc-gen-goext. DO NOT EDIT.

package ocr

type RecognizeTextRequest_Source = isRecognizeTextRequest_Source

func (m *RecognizeTextRequest) SetSource(v RecognizeTextRequest_Source) {
	m.Source = v
}

func (m *RecognizeTextRequest) SetContent(v []byte) {
	m.Source = &RecognizeTextRequest_Content{
		Content: v,
	}
}

func (m *RecognizeTextRequest) SetMimeType(v string) {
	m.MimeType = v
}

func (m *RecognizeTextRequest) SetLanguageCodes(v []string) {
	m.LanguageCodes = v
}

func (m *RecognizeTextRequest) SetModel(v string) {
	m.Model = v
}

func (m *RecognizeTextResponse) SetTextAnnotation(v *TextAnnotation) {
	m.TextAnnotation = v
}

func (m *RecognizeTextResponse) SetPage(v int64) {
	m.Page = v
}

func (m *GetRecognitionRequest) SetOperationId(v string) {
	m.OperationId = v
}
