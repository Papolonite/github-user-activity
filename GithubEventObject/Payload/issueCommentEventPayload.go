package payload

import "encoding/json"

type IssueCommentEventPayload struct {
	Action  string          `json:"action"`
	Comment json.RawMessage `json:"comment"`
	Issue   json.RawMessage `json:"issue"`
}
