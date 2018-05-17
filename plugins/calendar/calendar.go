package main

import (
	"strings"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"./pkg"
	"./cmd"
)

const (
	name = "Calendar"
)

type (
	calendarPlugin struct { }
)

var myCalendar = pkg.CalendarImp{}
var argParser = cmd.ArgParser{Calendar: &myCalendar}



func (c *calendarPlugin) Name() string {
	return name
}

func (c *calendarPlugin) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}
	return c.handleMessageEvent(msg, botUser)
}

func (c *calendarPlugin) handleMessageEvent(messageEvent slack.Message, botUser bot.UserInfo) ([]slack.Message, error) {
		args := strings.Split(strings.TrimSpace(messageEvent.Text()), " ")

	retStr, err := argParser.ParseCmd(args)

	// Finally print the collected string
	if err != nil {
		retStr = err.Error()
	}
	outMessages := []slack.Message{
		slack.NewMessage(retStr, messageEvent.Channel(), botUser.ID()),
	}

	return outMessages, nil
}

var CustomPlugin calendarPlugin
