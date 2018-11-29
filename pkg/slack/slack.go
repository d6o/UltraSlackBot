package slack

import (
	"fmt"

	slackClient "github.com/nlopes/slack"
)

type (
	Slack struct {
		token     string
		api       *slackClient.Client
		rtm       *slackClient.RTM
		eventList chan Event
	}

	Message interface {
		Text() string
		Channel() string
		User() User
	}

	User interface {
		ID() string
		Name() string
	}

	Event struct {
		name EventName
		data interface{}
	}

	EventName string

	UserInfo struct {
		id   string
		name string
	}
)

const (
	eventBufferSize = 1024
)

func New(token string) *Slack {
	return &Slack{
		token: token,
	}
}

func (s *Slack) Listen() chan Event {
	s.api = slackClient.New(s.token)
	s.eventList = make(chan Event, eventBufferSize)

	go s.manageMessages()

	return s.eventList
}

func (s *Slack) manageMessages() {
	s.rtm = s.api.NewRTM()
	go s.rtm.ManageConnection()
	for rtmEvent := range s.rtm.IncomingEvents {
		s.eventList <- Event{
			name: EventName(rtmEvent.Type),
			data: rtmEvent.Data,
		}
	}
}

func (s *Slack) Send(msg Message) {
	s.rtm.SendMessage(s.rtm.NewOutgoingMessage(msg.Text(), msg.Channel()))
}

func (s *Slack) UserInfo() (*UserInfo, error) {
	info, err := s.api.AuthTest()
	if err != nil {
		return nil, fmt.Errorf("AuthTest err: %s", err.Error())
	}
	userInfo := &UserInfo{
		id:   info.UserID,
		name: info.User,
	}
	return userInfo, nil
}

func (e *Event) Name() EventName {
	return e.name
}

func (e *Event) Data() interface{} {
	return e.data
}

func (u UserInfo) ID() string {
	return u.id
}

func (u UserInfo) Name() string {
	return u.name
}
