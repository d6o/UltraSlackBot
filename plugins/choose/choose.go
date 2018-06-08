package main

import (
	"strings"
	"math/rand"

	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

type (
	choose struct { }
)

func (c *choose) Start(specs bot.Specs) error {
	return nil
}

func (c *choose) Name() string {
	return "choose"
}

func (c *choose) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}
	return c.handleMessageEvent(msg, botUser)
}

func (c *choose) handleMessageEvent(messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
	args := strings.Split(strings.TrimSpace(messageEvent.Text()), " ")
	if len(args) < 3 || args[0] != c.Name() {
		return nil, nil
	}

	options := args[1:]

	rightOne := options[rand.Intn(len(options))]

	outMessages := []slack.Message{
		slack.NewMessage(rightOne, messageEvent.Channel(), botUser),
		}

	return outMessages, nil
}

var CustomPlugin choose
