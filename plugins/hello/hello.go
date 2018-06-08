package main

import (
	"strings"
	"fmt"

	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

type (
	hello struct { }
)


func (h *hello) Start(specs bot.Specs) error {
	return nil
}

func (h *hello) Name() string {
	return "hello"
}

func (h *hello) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}
	return h.handleMessageEvent(msg, botUser)
}

func (h *hello) handleMessageEvent(messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
	args := strings.Split(strings.TrimSpace(messageEvent.Text()), " ")
	if args[0] != h.Name() {
		return nil, nil
	}

	outMessages := []slack.Message{
		slack.NewMessage(fmt.Sprintf("Hello %s", messageEvent.User().Name()), messageEvent.Channel(), botUser),
	}

	return outMessages, nil
}

var CustomPlugin hello

