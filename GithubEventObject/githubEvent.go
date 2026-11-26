package event

import (
	"encoding/json"
	"fmt"
	"time"

	payload "github.com/Papolonite/github-user-activity/GithubEventObject/Payload"
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

type GithubEvent struct {
	ID        int64           `json:"id"`
	Type      EventType       `json:"type"`
	Actor     User            `json:"actor"`
	Repo      Repository      `json:"repo"`
	Payload   json.RawMessage `json:"payload"`
	Public    bool            `json:"public"`
	CreatedAt time.Time       `json:"created_at"`
	Org       *Organization   `json:"org,omitempty"`
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

func (ge *GithubEvent) writeComment() (message string, err error) {
	switch ge.Type {
	case CommitCommentEvent:
		var p payload.CommitCommentEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = fmt.Sprintf("%s Comment on %s", p.Action, ge.Repo.Name)
	case CreateEvent:
		var p payload.CreateEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "Created new {REF_TYPE}{REF} [ON {REPO_NAME} WHEN REF_TYPE IS NOT REPO]"
	case DeleteEvent:
		var p payload.DeleteEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "Deleted {REF_TYPE}{REF} ON {REPO_NAME}"
	case DiscussionEvent:
		var p payload.DiscussionEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = fmt.Sprintf("%s Discussion on %s", p.Action, ge.repo.Name)
	case ForkEvent:
		var p payload.ForkEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{ACTION} Repository {ORIGINAL_REPO} into {REPO_NAME}"
	case GollumEvent:
		var p payload.GollumEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		// Need to iterate each page
		message = "{action} Page {PAGE_NAME}"
	case IssueCommentEvent:
		var p payload.IssueCommentEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{action} comment on issue {issue_name} on {REPO_NAME}"
	case IssuesEvent:
		var p payload.IssueEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{action} issue {issue_name} on {REPO_NAME}"
	case MemberEvent:
		var p payload.MemberEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{action} member {member_name} on {REPO_NAME}"
	case PublicEvent:
		message = fmt.Sprintf("Update %s into public", ge.Repo.Name)
	case PullRequestEvent:
		var p payload.PullRequestEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{Action} Pull Request {Pull Request Name} on {REPO_NAME}"
	case PullRequestReviewEvent:
		var p payload.PullRequestEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{Action} Pull Request Review on {Pull Request Name} on {Repo_Name}"
	case PullRequestReviewCommentEvent:
		var p payload.PullRequestReviewCommentEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{Action} Comment on Pull Request Review on {Pull Request Name} on {REPO_NAME}"
	case PushEvent:
		var p payload.PushEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "Pushed {amount_of_commit} into branch {ref}"
	case ReleaseEvent:
		var p payload.ReleaseEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{Action} {release_name} on repository {REPO_NAME}"
	case WatchEvent:
		var p payload.WatchEventPayload
		if err := json.Unmarshal(ge.Payload, &p); err != nil {
			return message, err
		}
		message = "{Action} Repository {REPO_NAME}"
		message = fmt.Sprintf("%s Repository %s", p.Action, ge.Repo.Name)
	default:
		message = "Event Not Recognized"
	}

	return
}
