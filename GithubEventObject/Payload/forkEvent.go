package payload

import "encoding/json"

type ForkEventPayload struct {
	Action string          `json:"action"`
	Forkee json.RawMessage `json:"forkee"`
}
