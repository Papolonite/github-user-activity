package payload

import "encoding/json"

type MemberEventPayload struct {
	Action  string          `json:"action"`
	Comment json.RawMessage `json:"comment"`
}
