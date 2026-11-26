package payload

type PushEventPayload struct {
	RepositoryID int64  `json:"repository_id"`
	PushID       int64  `json:"push_id"`
	Ref          string `json:"ref"`
	Head         string `json:"head"`
	Before       string `json:"before"`
}
