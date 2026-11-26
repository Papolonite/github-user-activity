package event

import "time"

type EventType string

const (
	CommitCommentEvent            EventType = "CommitCommentEvent"
	CreateEvent                   EventType = "CreateEvent"
	DeleteEvent                   EventType = "DeleteEvent"
	DiscussionEvent               EventType = "DiscussionEvent"
	ForkEvent                     EventType = "ForkEvent"
	GollumEvent                   EventType = "GollumEvent"
	IssueCommentEvent             EventType = "IssueCommentEvent"
	IssuesEvent                   EventType = "IssuesEvent"
	MemberEvent                   EventType = "MemberEvent"
	PublicEvent                   EventType = "PublicEvent"
	PullRequestEvent              EventType = "PullRequestEvent"
	PullRequestReviewEvent        EventType = "PullRequestReviewEvent"
	PullRequestReviewCommentEvent EventType = "PullRequestReviewCommentEvent"
	PushEvent                     EventType = "PushEvent"
	ReleaseEvent                  EventType = "ReleaseEvent"
	WatchEvent                    EventType = "WatchEvent"
)

type GithubEvent[T any] struct {
	ID        int64         `json:"id"`
	Type      EventType     `json:"type"`
	Actor     User          `json:"actor"`
	Repo      Repository    `json:"repo"`
	Payload   T             `json:"payload"`
	Public    bool          `json:"public"`
	CreatedAt time.Time     `json:"created_at"`
	Org       *Organization `json:"org,omitempty"`
}

type User struct {
	ID           int64  `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type Repository struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Organization struct {
	ID         int64  `json:"id"`
	Login      string `json:"login"`
	GravatarID string `json:"gravatar_id"`
	URL        string `json:"url"`
	AvatarURL  string `json:"avatar_url"`
}
