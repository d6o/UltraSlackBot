package main

import (
	"regexp"
	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern     = "(?i)\\b(chuck|norris)\\b"
	catFactsURL = "https://api.icndb.com/jokes/random"
)

type (
	chuckNorris            struct{
		plugin.BasicCommand
	}

	chuckNorrisResponse struct {
		Value struct {
			Joke       string        `json:"joke"`
		} `json:"value"`
	}
)

func (c *chuckNorris) Start() error {
	return nil
}

func (c *chuckNorris) Name() string {
	return "chuckNorris"
}

func (c *chuckNorris) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *chuckNorris) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *chuckNorris) command(text string) (string, error) {
	data := &chuckNorrisResponse{}
	err := plugin.GetJSON(catFactsURL, data)
	if err != nil {
		return "", err
	}

	return data.Value.Joke, nil
}

var CustomPlugin chuckNorris
