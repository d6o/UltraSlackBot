package google

import (
	"net/url"

	"github.com/disiqueira/ultraslackbot/pkg/command"
)

const (
	searchURL = "https://www.googleapis.com/customsearch/v1"
)

type (
	google struct {
		cx    string
		key   string
		total int
		skip  int
	}

	googleResponse struct {
		Items []item `json:"items"`
	}

	item struct {
		Snippet string `json:"snippet"`
		Title   string `json:"title"`
		Link    string `json:"link"`
	}
)

func (g *google) reset() {
	g.skip = 0
	g.total = 1
}

func (g *google) search(text string, params map[string]string) ([]item, error) {
	var gisURL *url.URL
	gisURL, err := url.Parse(searchURL)
	if err != nil {
		return nil, err
	}

	parameters := url.Values{}

	for n, v := range params {
		parameters.Add(n, v)
	}

	parameters.Add("key", g.key)
	parameters.Add("cx", g.cx)
	parameters.Add("q", text)
	gisURL.RawQuery = parameters.Encode()

	data := &googleResponse{}
	if err := command.GetJSON(gisURL.String(), data); err != nil {
		return nil, err
	}

	var msgList []item
	skip := 0
	for _, item := range data.Items {
		if len(item.Link) > 0 {
			if g.skip > skip {
				skip++
				continue
			}
			msgList = append(msgList, item)
		}
		if len(msgList) == g.total {
			break
		}
	}

	return msgList, nil
}
