package main

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
)

const (
	pattern           = "(?i)^(wolframalpha|wa|calc|ca|math|convert)"
	wolframKeyEnvName = "WOLFRAMKEY"
	urlWolfram = "https://api.wolframalpha.com/v2/query?input=%s&appid=%s&output=JSON"
)

type (
	wolframAlpha struct {
		plugin.BasicCommand
		key string
	}

	wolframResponse struct {
		QueryResult struct {
			Success bool `json:"success"`
			Error   bool `json:"error"`
			Pods    []struct {
				Title   string `json:"title"`
				Primary bool `json:"primary,omitempty"`
				SubPods []struct {
					Plaintext string `json:"plaintext"`
				} `json:"subpods"`
			} `json:"pods"`
		} `json:"queryresult"`
	}
)

var (
	errParseResult = errors.New("error while parsing the result")
)

func (w *wolframAlpha) Start(specs bot.Specs) error {
	key, ok := specs.Get(wolframKeyEnvName)
	if !ok {
		return fmt.Errorf("config %s not found", wolframKeyEnvName)
	}

	w.key = key.(string)

	return nil
}

func (w *wolframAlpha) Name() string {
	return "wolframAlpha"
}

func (w *wolframAlpha) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return w.HandleEvent(event, botUser, w.matcher, w.command)
}

func (w *wolframAlpha) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (w *wolframAlpha) command(text string) (string, error) {
	args := strings.Split(strings.TrimSpace(text), " ")
	if len(args) < 2 {
		return "", nil
	}

	query := url.QueryEscape(strings.Join(args[1:], " "))

	data := &wolframResponse{}
	u := fmt.Sprintf(urlWolfram, query, w.key)
	err := plugin.GetJSON(u, data)
	if err != nil {
		return "", err
	}

	if data.QueryResult.Error || !data.QueryResult.Success {
		return "", errParseResult
	}

	if len(data.QueryResult.Pods) < 1 {
		return "No answer :(", nil
	}

	var title string
	var pods []string
	for _, pod := range data.QueryResult.Pods {
		if !pod.Primary {
			continue
		}

		title = pod.Title

		for _, subPod := range pod.SubPods {
			text := subPod.Plaintext
			if text != "" {
				pods = append(pods, text)
			}
		}
	}

	return fmt.Sprintf("%s: %s", title, strings.Join(pods, ", ")), nil
}

var CustomPlugin wolframAlpha
