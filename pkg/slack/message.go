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
		user User
	}

	user struct {
		name string
		id string
	}
)

func NewMessage(text, channel string, user User) Message {
	return &message{
		text:text,
		channel:channel,
		user:user,
	}
}

func (m message) Text() string {
	return m.text
}

func (m message) Channel() string {
	return m.channel
}

func (m message) User() User {
	return m.user
}

func (m message) Args() []string {
	return strings.Split(strings.TrimSpace(m.Text()), spaceCharacter)
}

func (m message) Command() string {
	return m.Args()[0]
}

func NewUser(id, name string) User {
	return &user{
		id:id,
		name:name,
	}
}

func (u user) ID() string {
	return u.id
}

func (u user) Name() string {
	return u.name
}

func EventToMessage(event Event) (Message, error) {
	switch ev := event.Data().(type) {
	case *slack.MessageEvent:
		return NewMessage(ev.Text, ev.Channel, NewUser(ev.User, ev.Username)), nil
	}
	return nil, errors.New("event is not a messageEvent")
}
