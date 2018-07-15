package lastfm

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/disiqueira/ultraslackbot/pkg/command"
)

const (
	urlLastFM = "https://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=%s&api_key=%s&format=json&limit=1"

	example = `
		# Last music by maef_5
		!lastfm maef_5`

	answerFormat = "%s is listening to \"%s\" by %s from the album %s."
)

type (
	lastfm struct {
		key string
	}

	lastFMResponse struct {
		RecentTracks struct {
			Track []struct {
				Album struct {
					Text string `json:"#text"`
				} `json:"album"`
				Artist struct {
					Text string `json:"#text"`
				} `json:"artist"`
				Name       string `json:"name"`
			} `json:"track"`
		} `json:"recenttracks"`
	}
)

func NewLastFMCommand(key string) *cobra.Command {
	w := newLastFM(key)

	c := &cobra.Command{
		Use:     "lastfm",
		Short:   "Find the last song played by an user",
		Example: example,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := w.lastSongByUser(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
		Aliases: []string{"lfm", "song"},
	}

	return c
}

func newLastFM(key string) *lastfm {
	return &lastfm{
		key:key,
	}
}

func (l *lastfm) lastSongByUser(user string) (string, error) {
	data := &lastFMResponse{}
	u := fmt.Sprintf(urlLastFM, user, l.key)
	err := command.GetJSON(u, data)
	if err != nil {
		return "", err
	}

	if len(data.RecentTracks.Track) < 1 {
		return "History not found.", nil
	}

	lastTrack := data.RecentTracks.Track[0]

	return fmt.Sprintf(answerFormat, user, lastTrack.Name, lastTrack.Artist.Text, lastTrack.Album.Text), nil
}
