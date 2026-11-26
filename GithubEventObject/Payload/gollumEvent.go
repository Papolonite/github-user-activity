package payload

type GollumEventPayload struct {
	Pages []GollumPage `json:"pages"`
}

type GollumPage struct {
	PageName string  `json:"page_name"`
	Title    string  `json:"title"`
	Summary  *string `json:"summary"`
	Action   string  `json:"action"`
	SHA      string  `json:"sha"`
	HTMLURL  string  `json:"html_url"`
}
