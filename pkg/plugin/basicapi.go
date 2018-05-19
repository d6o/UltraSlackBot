package plugin

import (
	"regexp"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern   = "(?i)\\b(basic|api)\\b"
)

type (
	BasicCommand struct { }

	Matcher func() *regexp.Regexp
	Command func(text string) (string, error)
)

func (c *BasicCommand) Name() string {
	return "BasicCommand"
}

func (c *BasicCommand) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *BasicCommand) HandleEvent(event slack.Event, botUser bot.UserInfo, matcher Matcher, command Command) ([]slack.Message, error) {
	msg, err := slack.EventToMessage(event)
	if err != nil {
		return nil, nil
	}

	if msg.User().ID() == botUser.ID() {
		return nil, nil
	}

	if !matcher().MatchString(msg.Text()) {
		return nil, nil
	}

	textResponse, err := command(msg.Text())
	if err != nil {
		return nil, err
	}

	outMessages := []slack.Message{
		slack.NewMessage(textResponse, msg.Channel(), botUser),
	}

	return outMessages, nil
}

func (c *BasicCommand) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *BasicCommand) command(text string) (string, error) {
	return "Hello! :)", nil
}
