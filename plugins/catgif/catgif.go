package main

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern   = "(?i)\\b(cat|gato|miau|meow|garfield|lolcat)[s|z]{0,1}\\b"
	msgPrefix = "I love cats! Here's a gif: %s"
	catGifURL = "http://thecatapi.com/api/images/get?format=src&type=gif"
)

type (
	catGif struct {
		plugin.BasicCommand
	}
)

func (c *catGif) Name() string {
	return "catGif"
}

func (c *catGif) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *catGif) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *catGif) command(text string) (string, error) {
	res, err := http.Get(catGifURL)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf(msgPrefix, res.Request.URL.String()), nil
}

var CustomPlugin catGif
