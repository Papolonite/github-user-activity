package payload

import "encoding/json"

type CommitCommentEventPayload struct {
	Action  string          `json:"action"`
	Comment json.RawMessage `json:"comment"`
}
