package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"net/http"
	"encoding/json"
	"time"
)

const (
	pattern         = "(?i)\\b(urban|u)\\b"
	urbanKeyEnvName = "LASTFMKEY"
	searchURL       = "http://api.urbandictionary.com/v0/define?term=%s"
	referer         = "http://m.urbandictionary.com"
	noResult		= "no_results"
)

type (
	urban struct {
		plugin.BasicCommand
		key string
	}

	urbanResponse struct {
		Tags       []string `json:"tags"`
		ResultType string   `json:"result_type"`
		List       []struct {
			Definition  string    `json:"definition"`
			Permalink   string    `json:"permalink"`
			ThumbsUp    int       `json:"thumbs_up"`
			Author      string    `json:"author"`
			Word        string    `json:"word"`
			Defid       int       `json:"defid"`
			CurrentVote string    `json:"current_vote"`
			WrittenOn   time.Time `json:"written_on"`
			Example     string    `json:"example"`
			ThumbsDown  int       `json:"thumbs_down"`
		} `json:"list"`
		Sounds []string `json:"sounds"`
	}
)

func (l *urban) Start(specs bot.Specs) error {
	key, ok := specs.Get(urbanKeyEnvName)
	if !ok {
		return fmt.Errorf("config %s not found", urbanKeyEnvName)
	}
	l.key = key.(string)

	return nil
}

func (l *urban) Name() string {
	return "urban"
}

func (l *urban) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return l.HandleEvent(event, botUser, l.matcher, l.command)
}

func (l *urban) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (l *urban) command(text string) (string, error) {
	args := strings.Split(strings.TrimSpace(text), " ")
	if len(args) < 2 {
		return "", nil
	}

	q := url.QueryEscape(strings.Join(args[1:], " "))
	u := fmt.Sprintf(searchURL, q)

	client := &http.Client{}
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("referer", referer)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var data urbanResponse
	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return "", err
	}

	if data.ResultType == noResult || len(data.List) < 1 {
		return fmt.Sprintf("No definition for %s found.", q), err
	}

	return data.List[0].Definition, nil
}

var CustomPlugin urban
