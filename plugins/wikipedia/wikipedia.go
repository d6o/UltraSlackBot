package main

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
	"errors"
)

const (
	pattern    = "(?i)^(wikipedia|wiki)"
	urlWikipedia = "https://%s.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s"
)

var (
	langs = []string{
		"en",
		"pt",
		"es",
	}

	errPageNotFound = errors.New("page not found")
)

type (
	wikipedia struct {
		plugin.BasicCommand
	}

	wikipediaResponse struct {
		Query         struct {
			Pages map[string]page `json:"pages"`
		} `json:"query"`
	}

	page struct {
		Extract string `json:"extract"`
		Title   string `json:"title"`
	}
)

func (w *wikipedia) Start(specs bot.Specs) error {
	return nil
}

func (w *wikipedia) Name() string {
	return "wikipedia"
}

func (w *wikipedia) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return w.HandleEvent(event, botUser, w.matcher, w.command)
}

func (w *wikipedia) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (w *wikipedia) command(text string) (string, error) {
	args := strings.Split(strings.TrimSpace(text), " ")
	if len(args) < 2 {
		return "", nil
	}

	q := strings.Join(args[1:], " ")
	title, extract, err := w.searchAllLangs(q)
	return fmt.Sprintf("%s: %s", title, extract), err
}

func (w *wikipedia) searchAllLangs(query string) (string, string, error) {
	for _, lang := range langs {
		title, extract, err := w.search(lang, query)
		if err == nil {
			return title, extract, nil
		}
	}
	return "", "", errPageNotFound
}

func (w *wikipedia) search(lang, query string) (string, string, error) {
	query = url.QueryEscape(query)

	data := &wikipediaResponse{}
	u := fmt.Sprintf(urlWikipedia, lang, query)
	err := plugin.GetJSON(u, data)
	if err != nil {
		return "", "", err
	}

	var title, extract string
	for id, page := range data.Query.Pages {
		if id == "-1" {
			return "", "", errPageNotFound
		}
		title = page.Title
		extract = page.Extract
	}

	if extract == "" {
		return "", "", errPageNotFound
	}
	return title, extract, nil
}

var CustomPlugin wikipedia