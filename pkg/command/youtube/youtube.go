package youtube

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
	"google.golang.org/api/googleapi/transport"
	yt "google.golang.org/api/youtube/v3"
)

const (
	example = `
		# Post a video of dogs
		!youtube dogs

		# Post a video of a hot dog
		!youtube hot dog

		# Post two videos of hot dog
		!youtube hot dog --total 2

		# Post the second video of a hot dog
		!youtube hot dog --skip 1

		# Post the third and fourth video of a hot dog
		!youtube hot dog --skip 2 --total 2`
)

type (
	youtube struct {
		youtubeService *yt.Service
		total          int
		skip           int
	}
)

func NewYoutubeCommand(key string) *cobra.Command {
	//@TODO(disiqueira): Propagate and handle the error on bootstrap.
	y, _ := newYoutube(key)

	c := &cobra.Command{
		Use:     "youtube",
		Short:   "Search videos on Youtube",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := y.search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
			y.reset()
		},
		Aliases: []string{"video", "yt", "y"},
	}

	c.Flags().IntVarP(&y.total, "total", "t", 1, "How many results will be returned")
	c.Flags().IntVarP(&y.skip, "skip", "s", 0, "How many results should be skipped")

	return c
}

func newYoutube(key string) (*youtube, error) {
	client := &http.Client{
		Transport: &transport.APIKey{Key: key},
	}

	service, err := yt.New(client)
	if err != nil {
		return nil, fmt.Errorf("error creating new YouTube client: %v", err)
	}

	y := &youtube{}
	y.youtubeService = service

	return y, nil
}

func (y *youtube) search(q string) (string, error) {
	call := y.youtubeService.Search.List("id,snippet").
		Q(q).Type("video").
		MaxResults(int64(y.total + y.skip))
	response, err := call.Do()
	if err != nil {
		return "", fmt.Errorf("error making search API call: %v", err)
	}

	if response.PageInfo.TotalResults <= 0 {
		return fmt.Sprintf("Query: %s Err: No results found.", q), nil
	}

	var msgList []string
	skip := 0
	for _, item := range response.Items {
		if item.Id.VideoId != "" {
			if y.skip > skip {
				skip++
				continue
			}
			msgList = append(msgList, fmt.Sprintf("%s - http://youtu.be/%s", item.Snippet.Title, item.Id.VideoId))
		}
		if len(msgList) == y.total {
			break
		}
	}

	return strings.Join(msgList, "\n"), nil
}

func (y *youtube) reset() {
	y.skip = 0
	y.total = 1
}
