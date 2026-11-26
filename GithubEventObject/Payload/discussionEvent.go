package payload

import "encoding/json"

type DiscussionEventPayload struct {
	Action  string          `json:"action"`
	Comment json.RawMessage `json:"comment"`
}
