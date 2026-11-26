package event

import (
	"time"
)

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

// Output:
// - Pushed 3 commits to kamranahmedse/developer-roadmap
// - Opened a new issue in kamranahmedse/developer-roadmap
// - Starred kamranahmedse/developer-roadmap
// - ...

func (ge *GithubEvent[T]) writeComment() string {
	var message string
	switch ge.Type {
	case CommitCommentEvent:
		message = "{Action} Comment on {REPO_NAME}"
	case CreateEvent:
		message = "Created new {REF_TYPE}{REF} [ON {REPO_NAME} WHEN REF_TYPE IS NOT REPO]"
	case DeleteEvent:
		message = "Deleted {REF_TYPE}{REF} ON {REPO_NAME}"
	case DiscussionEvent:
		message = "{Action} Discussion on {REPO_NAME}"
	case ForkEvent:
		message = "{ACTION} Repository {ORIGINAL_REPO} into {REPO_NAME}"
	case GollumEvent:
		// Need to iterate each page
		message = "{action} Page {PAGE_NAME}"
	case IssueCommentEvent:
		message = "{action} comment on issue {issue_name} on {REPO_NAME}"
	case IssuesEvent:
		message = "{action} issue {issue_name} on {REPO_NAME}"
	case MemberEvent:
		message = "{action} member {member_name} on {REPO_NAME}"
	case PublicEvent:
		message = "update {repository} into public"
	case PullRequestEvent:
		message = "{Action} Pull Request {Pull Request Name} on {REPO_NAME}"
	case PullRequestReviewEvent:
		message = "{Action} Pull Request Review on {Pull Request Name} on {Repo_Name}"
	case PullRequestReviewCommentEvent:
		message = "{Action} Comment on Pull Request Review on {Pull Request Name} on {REPO_NAME}"
	case PushEvent:
		message = "Pushed {amount_of_commit} into branch {ref}"
	case ReleaseEvent:
		message = "{Action} {release_name} on repository {REPO_NAME}"
	case WatchEvent:
		message = "{Action} Repository {REPO_NAME}"
	default:
		message = "Event Not Recognized"
	}

	return message
}
