package backend

import "github.com/alex-chou/chat-bot/pkg/slack"

// Backend handles upstream client calls.
type Backend interface {
}

type backend struct {
	slack slack.Slack
}

// NewBackend create a new backend to use.
func NewBackend(slack slack.Slack) Backend {
	return &backend{
		slack: slack,
	}
}
