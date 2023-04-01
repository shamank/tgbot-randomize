package commands

import "github.com/shamank/tgbot-randomize/app/internal/messages"

type Commands struct {
	messages *messages.Messages
}

func NewCommands(messages *messages.Messages) *Commands {
	return &Commands{
		messages: messages,
	}
}
