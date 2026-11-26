package payload

import "encoding/json"

type PullRequestReviewEventPayload struct {
	Action      string          `json:"action"`
	PullRequest json.RawMessage `json:"pull_request"`
	Review      json.RawMessage `json:"review"`
}
