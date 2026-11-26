package payload

import "encoding/json"

type ReleaseEventPayload struct {
	Action  string          `json:"action"`
	Release json.RawMessage `json:"release"`
}
