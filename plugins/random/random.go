package main

import (
	"math/rand"
	"strconv"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
)

const (
	maxValueDefault = 100
)

type (
	random struct{}
)

func (r *random) Start(specs bot.Specs) error {
	return nil
}

func (r *random) Name() string {
	return "random"
}

func (r *random) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}
	return r.handleMessageEvent(msg, botUser)
}

func (r *random) handleMessageEvent(messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
	args := strings.Split(strings.TrimSpace(messageEvent.Text()), " ")
	if args[0] != r.Name() {
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
