package main

import (
	"strings"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/plugins/calendar/pkg"
	"github.com/disiqueira/ultraslackbot/plugins/calendar/cmd"
	"log"
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
	if !strings.HasPrefix(messageEvent.Text(), "Calendar") {
		return nil,nil
	}
	log.Print("Calendar begin")
	args := strings.Split(removeSpaces(messageEvent.Text()), " ")
	args = args[1:len(args)]

	retStr, err := argParser.ParseCmd(args)
	log.Print("Calendar return")
	// Finally print the collected string
	if err != nil {
		log.Print("Calendar error")
		retStr = err.Error()
	}
	// fmt.Print(retStr)
	/*var outMessages []slack.Message
	retMsgs := strings.Split(retStr, "\n")
	for _, msg := range retMsgs {
		outMessages = append(outMessages, slack.NewMessage(msg, messageEvent.Channel(), botUser.ID()))
	}*/
	log.Print("Calendar before message")
	outMessages := []slack.Message {
		slack.NewMessage(retStr, messageEvent.Channel(), botUser.ID()),
	}
	log.Print("Calendar after message")
	return outMessages, nil
}

func removeSpaces(str string) string {
	return str//strings.Join(strings.Fields(str), " ")
}

var CustomPlugin calendarPlugin
