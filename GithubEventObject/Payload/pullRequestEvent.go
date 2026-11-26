package payload

import "encoding/json"

type PullRequestEventPayload struct {
	Action      string            `json:"action"`
	Number      int64             `json:"number"`
	PullRequest json.RawMessage   `json:"pull_request"`
	Assignee    *json.RawMessage  `json:"assignee,omitempty"`
	Assignees   []json.RawMessage `json:"assignees,omitempty"`
	Label       *json.RawMessage  `json:"label,omitempty"`
	Labels      []json.RawMessage `json:"labels,omitempty"`
}
