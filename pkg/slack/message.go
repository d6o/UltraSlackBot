package slack

import (
	"strings"
	"github.com/nlopes/slack"
	"errors"
)

const (
	spaceCharacter = " "
)

type (
	message struct {
		text string
		channel string
		userID string
	}
)

func NewMessage(text, channel, userID string) Message {
	return &message{
		text:text,
		channel:channel,
		userID:userID,
	}
}

func (m message) Text() string {
	return m.text
}

func (m message) Channel() string {
	return m.channel
}

func (m message) UserID() string {
	return m.userID
}

func (m message) Args() []string {
	return strings.Split(strings.TrimSpace(m.Text()), spaceCharacter)
}

func (m message) Command() string {
	return m.Args()[0]
}

func EventToMessage(event Event) (Message, error) {
	switch ev := event.Data().(type) {
	case *slack.MessageEvent:
		return NewMessage(ev.Text, ev.Channel, ev.User), nil
	}
	return nil, errors.New("event is not a messageEvent")
}
