package payload

import "encoding/json"

type IssueEventPayload struct {
	Action    string            `json:"action"`
	Issue     json.RawMessage   `json:"issue"`
	Assignee  *json.RawMessage  `json:"assignee,omitempty"`
	Assignees []json.RawMessage `json:"assignees,omitempty"`
	Label     *json.RawMessage  `json:"label,omitempty"`
	Labels    []json.RawMessage `json:"labels,omitempty"`
}
