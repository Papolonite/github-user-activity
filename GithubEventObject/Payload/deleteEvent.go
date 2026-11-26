package payload

type DeleteEventPayload struct {
	Ref        string `json:"ref"`
	RefType    string `json:"ref_type"`
	FullRef    string `json:"full_ref"`
	PusherType string `json:"pusher_type"`
}
