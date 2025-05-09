// Code generated by protoc-gen-goext. DO NOT EDIT.

package eventrouter

import (
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func (m *Connector) SetId(v string) {
	m.Id = v
}

func (m *Connector) SetBusId(v string) {
	m.BusId = v
}

func (m *Connector) SetFolderId(v string) {
	m.FolderId = v
}

func (m *Connector) SetCloudId(v string) {
	m.CloudId = v
}

func (m *Connector) SetCreatedAt(v *timestamppb.Timestamp) {
	m.CreatedAt = v
}

func (m *Connector) SetName(v string) {
	m.Name = v
}

func (m *Connector) SetDescription(v string) {
	m.Description = v
}

func (m *Connector) SetLabels(v map[string]string) {
	m.Labels = v
}

func (m *Connector) SetSource(v *Source) {
	m.Source = v
}

func (m *Connector) SetDeletionProtection(v bool) {
	m.DeletionProtection = v
}

func (m *Connector) SetStatus(v Connector_Status) {
	m.Status = v
}

type Source_Source = isSource_Source

func (m *Source) SetSource(v Source_Source) {
	m.Source = v
}

func (m *Source) SetDataStream(v *DataStream) {
	m.Source = &Source_DataStream{
		DataStream: v,
	}
}

func (m *Source) SetMessageQueue(v *MessageQueue) {
	m.Source = &Source_MessageQueue{
		MessageQueue: v,
	}
}

func (m *Source) SetTimer(v *Timer) {
	m.Source = &Source_Timer{
		Timer: v,
	}
}

func (m *Source) SetEventServiceSource(v *EventServiceSource) {
	m.Source = &Source_EventServiceSource{
		EventServiceSource: v,
	}
}

func (m *DataStream) SetDatabase(v string) {
	m.Database = v
}

func (m *DataStream) SetStreamName(v string) {
	m.StreamName = v
}

func (m *DataStream) SetConsumer(v string) {
	m.Consumer = v
}

func (m *DataStream) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *MessageQueue) SetQueueArn(v string) {
	m.QueueArn = v
}

func (m *MessageQueue) SetServiceAccountId(v string) {
	m.ServiceAccountId = v
}

func (m *MessageQueue) SetVisibilityTimeout(v *durationpb.Duration) {
	m.VisibilityTimeout = v
}

func (m *MessageQueue) SetBatchSize(v int64) {
	m.BatchSize = v
}

func (m *MessageQueue) SetPollingTimeout(v *durationpb.Duration) {
	m.PollingTimeout = v
}

func (m *Timer) SetCronExpression(v string) {
	m.CronExpression = v
}

func (m *Timer) SetTimeZone(v string) {
	m.TimeZone = v
}

func (m *Timer) SetPayload(v string) {
	m.Payload = v
}
