package main

import (
	"strings"
	"math/rand"

	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"strconv"
)

const (
	maxValueDefault = 100
)

type (
	random struct { }
)

func (c *random) Name() string {
	return "random"
}

func (c *random) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}
	return c.handleMessageEvent(msg, botUser)
}

func (c *random) handleMessageEvent(messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
	args := strings.Split(strings.TrimSpace(messageEvent.Text()), " ")
	if args[0] != c.Name() {
		return nil, nil
	}

	maxNumber := maxValueDefault
	if len(args) >= 2 {
		val, err := strconv.Atoi(args[1])
		if err == nil {
			maxNumber = val
		}
	}

	outMessages := []slack.Message{
		slack.NewMessage(strconv.Itoa(rand.Intn(maxNumber)), messageEvent.Channel(), botUser),
		}

	return outMessages, nil
}

var CustomPlugin random
