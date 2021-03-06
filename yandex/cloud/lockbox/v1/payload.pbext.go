// Code generated by protoc-gen-goext. DO NOT EDIT.

package lockbox

func (m *Payload) SetVersionId(v string) {
	m.VersionId = v
}

func (m *Payload) SetEntries(v []*Payload_Entry) {
	m.Entries = v
}

type Payload_Entry_Value = isPayload_Entry_Value

func (m *Payload_Entry) SetValue(v Payload_Entry_Value) {
	m.Value = v
}

func (m *Payload_Entry) SetKey(v string) {
	m.Key = v
}

func (m *Payload_Entry) SetTextValue(v string) {
	m.Value = &Payload_Entry_TextValue{
		TextValue: v,
	}
}

func (m *Payload_Entry) SetBinaryValue(v []byte) {
	m.Value = &Payload_Entry_BinaryValue{
		BinaryValue: v,
	}
}
