package main

import (
	"regexp"
	"fmt"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern   = "(?i)\\b(cat|gato|miau|meow|garfield|lolcat)[s|z]{0,1}\\b"
	msgPrefix = "I love cats! Here's a fact: %s"
	catFactURL = "http://catfact.ninja/fact"
)

type (
	catFact             struct{
		plugin.BasicCommand
	}
	catFactResponse struct {
		Fact   string `json:"fact"`
	}
)

func (c *catFact) Name() string {
	return "catFact"
}

func (c *catFact) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *catFact) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *catFact) command(text string) (string, error) {
	data := &catFactResponse{}
	err := plugin.GetJSON(catFactURL, data)
	if err != nil {
		return "", err
	}

	if len(data.Fact) == 0 {
		return "", nil
	}

	return fmt.Sprintf(msgPrefix, data.Fact), nil
}

var CustomPlugin catFact
