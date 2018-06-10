package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/bot"
	"github.com/disiqueira/ultraslackbot/pkg/plugin"
	"github.com/disiqueira/ultraslackbot/pkg/slack"
)

const (
	pattern          = "(?i)\\b(lastFM|lfm)\\b"
	lastFMKeyEnvName = "LASTFMKEY"
	searchURL        = "https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1"
	answerFormat = "%s is listening to \"%s\" by %s from the album %s."
	historyNotFound = "History not found."
)

type (
	lastFM struct {
		plugin.BasicCommand
		key string
	}

	lastFMResponse struct {
		RecentTracks struct {
			Track []struct {
				Album struct {
					Text string `json:"#text"`
					Mbid string `json:"mbid"`
				} `json:"album"`
				Artist struct {
					Text string `json:"#text"`
					Mbid string `json:"mbid"`
				} `json:"artist"`
				Name       string `json:"name"`
			} `json:"track"`
		} `json:"recenttracks"`
	}
)

func (l *lastFM) Start(specs bot.Specs) error {
	key, ok := specs.Get(lastFMKeyEnvName)
	if !ok {
		return fmt.Errorf("config %s not found", lastFMKeyEnvName)
	}
	l.key = key.(string)

	return nil
}

func (l *lastFM) Name() string {
	return "lastFM"
}

func (l *lastFM) Execute(event slack.Event, botUser bot.UserInfo) ([]slack.Message, error) {
	return l.HandleEvent(event, botUser, l.matcher, l.command)
}

func (l *lastFM) matcher() *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

func (l *lastFM) command(text string) (string, error) {
	args := strings.Split(strings.TrimSpace(text), " ")
	if len(args) < 2 {
		return "", nil
	}

	user := args[1]
	data := &lastFMResponse{}
	u := fmt.Sprintf(searchURL, user, l.key)
	err := plugin.GetJSON(u, data)
	if err != nil {
		return "", err
	}

	if len(data.RecentTracks.Track) < 1 {
		return historyNotFound, nil
	}

	lastTrack := data.RecentTracks.Track[0]

	return fmt.Sprintf(answerFormat, user, lastTrack.Name, lastTrack.Artist.Text, lastTrack.Album.Text), nil
}

var CustomPlugin lastFM
