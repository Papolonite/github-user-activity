package payload

import "encoding/json"

type PullRequestReviewCommentEventPayload struct {
	Action      string          `json:"action"`
	PullRequest json.RawMessage `json:"pull_request"`
	Comment     json.RawMessage `json:"comment"`
}
