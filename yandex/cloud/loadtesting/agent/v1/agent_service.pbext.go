// Code generated by protoc-gen-goext. DO NOT EDIT.

package agent

func (m *ClaimAgentStatusRequest) SetAgentInstanceId(v string) {
	m.AgentInstanceId = v
}

func (m *ClaimAgentStatusRequest) SetStatus(v ClaimAgentStatusRequest_Status) {
	m.Status = v
}

func (m *ClaimAgentStatusRequest) SetStatusMessage(v string) {
	m.StatusMessage = v
}

func (m *ClaimAgentStatusResponse) SetCode(v int64) {
	m.Code = v
}
