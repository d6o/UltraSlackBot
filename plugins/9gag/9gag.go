package main

import (
	"errors"
	"net/http"
	"net/url"
	"regexp"

	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"github.com/disiqueira/ultraslackbot/pkg/bot"
)

const (
	pattern = "(?i)\\b(9gag|ninegag)\\b"
	randomURL = "http://9gag.com/random"
)

type (
	choose struct {
		plugin.BasicCommand
	}
)

func (c *choose) Start(specs bot.Specs) error {
	return nil
}

func (c *choose) Name() string {
	return "9gag"
}

func (c *choose) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return c.HandleEvent(event, botUser, c.matcher, c.command)
}

func (c *choose) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (c *choose) command(text string) (string, error) {
	redirectNotAllowed := errors.New("redirect")
	redirectedURL := ""

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectedURL = req.URL.String()
		return redirectNotAllowed
	}

	_, err := client.Get(randomURL)
	if urlError, ok := err.(*url.Error); !ok || urlError.Err != redirectNotAllowed {
		return "", err
	}
	return redirectedURL, nil
}

var CustomPlugin choose
