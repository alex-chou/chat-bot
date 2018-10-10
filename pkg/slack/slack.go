package slack

import (
	"github.com/nlopes/slack"
)

// Slack is an interface of the functions used by the server.
type Slack interface {
}

// New creates a new Slack client.
func New(token string, options ...slack.Option) Slack {
	return slack.New(token, options...)
}
