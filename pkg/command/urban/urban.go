package urban

import (
	"fmt"
	"strings"
	"encoding/json"
	"net/http"

	"github.com/spf13/cobra"
)

const (
	searchURL       = "http://api.urbandictionary.com/v0/define?term=%s"
	referer         = "http://m.urbandictionary.com"
	noResult		= "no_results"

	example = `
		# Search for a definition on Urban Dictionary
		!urban LOL

		# Search for a definition on Urban Dictionary and skip the first one
		!urban LOL --skip 1`
)

type (
	urban struct {
		skip  int
	}

	urbanResponse struct {
		ResultType string   `json:"result_type"`
		List       []struct {
			Definition  string    `json:"definition"`
			Permalink   string    `json:"permalink"`
		} `json:"list"`
	}
)

func NewUrbanCommand() *cobra.Command {
	w := newUrban()

	c := &cobra.Command{
		Use:     "urban",
		Short:   "Search for a definition on Urban Dictionary",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := w.search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
		Aliases: []string{"ud"},
	}

	c.Flags().IntVarP(&w.skip, "skip","s", 0, "How many images should be skipped")

	return c
}

func newUrban() *urban {
	return &urban{}
}

func (u *urban) search(q string) (string, error) {
	uri := fmt.Sprintf(searchURL, q)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req.Header.Set("referer", referer)
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	var data urbanResponse
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return "", err
	}

	if data.ResultType == noResult || len(data.List) < 1 {
		return "", fmt.Errorf("no definition for %s found", q)
	}

	skip := 0
	for _, item := range data.List {
		if len(item.Definition) == 0 {
			continue
		}
		if u.skip > skip {
			skip++
			continue
		}
		return item.Definition, nil
	}

	return "", fmt.Errorf("no definition for %s found", q)
}

func (u *urban) reset() {
	u.skip = 0
}
