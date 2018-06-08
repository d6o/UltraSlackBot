package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"

	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

const (
	pattern          = "(?i)^(youtube|you|yt)"
	googleKeyEnvName = "GOOGLEKEY"
	videoURL         = "http://youtu.be/%s"
)

type (
	yt struct {
		plugin.BasicCommand
		youtubeService *youtube.Service
	}
)

func (y *yt) Start(specs bot.Specs) error {
	key, ok := specs.Get(googleKeyEnvName)
	if !ok {
		return fmt.Errorf("config %s not found", googleKeyEnvName)
	}

	client := &http.Client{
		Transport: &transport.APIKey{Key: key.(string)},
	}

	service, err := youtube.New(client)
	if err != nil {
		return fmt.Errorf("error creating new YouTube client: %v", err)
	}

	y.youtubeService = service

	return nil
}

func (y *yt) Name() string {
	return "yt"
}

func (y *yt) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return y.HandleEvent(event, botUser, y.matcher, y.command)
}

func (y *yt) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (y *yt) command(text string) (string, error) {
	args := strings.Split(strings.TrimSpace(text), " ")
	if len(args) < 2 {
		return "", nil
	}

	text = strings.Join(args[1:], " ")

	call := y.youtubeService.Search.List("id,snippet").
		Q(text).Type("video").
		MaxResults(1)
	response, err := call.Do()
	if err != nil {
		return "", fmt.Errorf("error making search API call: %v", err)
	}

	if response.PageInfo.TotalResults <= 0 {
		return fmt.Sprintf("Query: %s Err: No results found.", text), nil
	}

	url := ""
	description := ""
	for _, item := range response.Items {
		if item.Id.VideoId != "" {
			url = fmt.Sprintf(videoURL, item.Id.VideoId)
			description = item.Snippet.Title
			break
		}
	}

	return fmt.Sprintf("%s - %s", description, url), nil
}

var CustomPlugin yt
