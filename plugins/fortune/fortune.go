package main

import (
	"regexp"
	"fmt"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern = "(?i)^(fortune)"
	msgPrefix = "Fortune: %s"
	fortuneURL = "http://www.yerkee.com/api/fortune"
)

type (
	fortune             struct{
		plugin.BasicCommand
	}
	fortuneResponse struct {
		Fortune string `json:"fortune"`
	}
)

func (f *fortune) Start(specs bot.Specs) error {
	return nil
}

func (f *fortune) Name() string {
	return "fortune"
}

func (f *fortune) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return f.HandleEvent(event, botUser, f.matcher, f.command)
}

func (f *fortune) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (f *fortune) command(text string) (string, error) {
	data := &fortuneResponse{}
	err := plugin.GetJSON(fortuneURL, data)
	if err != nil {
		return "", err
	}

	if len(data.Fortune) == 0 {
		return "", nil
	}

	return fmt.Sprintf(msgPrefix, data.Fortune), nil
}

var CustomPlugin fortune
