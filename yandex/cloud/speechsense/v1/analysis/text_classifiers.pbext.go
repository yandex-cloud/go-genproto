// Code generated by protoc-gen-goext. DO NOT EDIT.

package speechsense

import (
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func (m *TextClassifiers) SetClassificationResult(v []*ClassificationResult) {
	m.ClassificationResult = v
}

func (m *ClassificationResult) SetClassifier(v string) {
	m.Classifier = v
}

func (m *ClassificationResult) SetClassifierStatistics(v []*ClassifierStatistics) {
	m.ClassifierStatistics = v
}

func (m *ClassifierStatistics) SetChannelNumber(v *wrapperspb.Int64Value) {
	m.ChannelNumber = v
}

func (m *ClassifierStatistics) SetTotalCount(v int64) {
	m.TotalCount = v
}

func (m *ClassifierStatistics) SetHistograms(v []*Histogram) {
	m.Histograms = v
}

func (m *Histogram) SetCountValues(v []int64) {
	m.CountValues = v
}