package payload

import "encoding/json"

type CommitCommentEventPayload struct {
	Action  string          `json:"action"`
	Comment json.RawMessage `json:"comment"`
}

func (g *GithubEvent[CommitCommentEventPayload]) createComment() {
	// This handler is intentionally empty because commit comment events need no additional processing.
}
