package main

import (
	"log"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/plugins/calendar/cmd"
	"github.com/disiqueira/ultraslackbot/plugins/calendar/pkg"
)

const (
	name = "Calendar"
)

type (
	calendarPlugin struct{}
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
	if !strings.HasPrefix(messageEvent.Text(), name) {
		return nil, nil
	}
	log.Print("Calendar begin")
	args := strings.Split(removeSpaces(messageEvent.Text()), " ")[1:]
	retStr, err := argParser.ParseCmd(args)
	if err != nil {
		log.Print("Calendar error ")
		retStr = err.Error()
	}
	outMessages := []slack.Message{
		slack.NewMessage("```"+retStr+"```", messageEvent.Channel(), botUser.ID()),
	}
	log.Print("Calendar end")
	return outMessages, nil
}

func removeSpaces(str string) string {
	return strings.Join(strings.Fields(str), " ")
}

var CustomPlugin calendarPlugin
