package docs

import (
	"fmt"
	"github.com/disiqueira/ultraslackbot/pkg/command"
	"github.com/spf13/cobra"
	"net/url"
	"strings"
)

const (
	urlGoDocs       = "https://api.godoc.org/search?q=%s"
	errFormatGoDocs = "no result return while searching for: %s"
	formatGoDocs 	= "%s - http://godoc.org/%s (Stars: %d Imports: %d): %s"
)

type (
	goDocs struct {
		total int
		skip  int
	}

	goDocsResponse struct {
		Results []struct {
			Name        string `json:"name"`
			Path        string `json:"path"`
			ImportCount int    `json:"import_count"`
			Synopsis    string `json:"synopsis,omitempty"`
			Stars       int    `json:"stars,omitempty"`
		} `json:"results"`
	}
)

func newDocsGoCommand() *cobra.Command {
	g := newGoDocs()

	c := &cobra.Command{
		Use:     "docs",
		Short:   "Search for docs",
		Example: "# Search for Go docs for fmt \n!docs go fmt",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := g.search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
			g.reset()
		},
	}
	c.Flags().IntVarP(&g.total, "total", "t", 1, "How many results will be returned")
	c.Flags().IntVarP(&g.skip, "skip", "s", 0, "How many results should be skipped")

	return c
}

func newGoDocs() *goDocs {
	return &goDocs{}
}

func (g *goDocs) search(q string) (string, error) {
	query := url.QueryEscape(q)
	u := fmt.Sprintf(urlGoDocs, query)

	data := &goDocsResponse{}
	err := command.GetJSON(u, data)
	if err != nil {
		return "", err
	}

	if len(data.Results) == 0 {
		return "", fmt.Errorf(errFormatGoDocs, q)
	}

	var msgList []string
	skip := 0
	for _, item := range data.Results {
		if g.skip > skip {
			skip++
			continue
		}
		msgList = append(msgList, fmt.Sprintf(formatGoDocs, item.Name, item.Path, item.Stars, item.ImportCount, item.Synopsis))
		if len(msgList) == g.total {
			break
		}
	}

	return strings.Join(msgList, " "), nil
}

func (g *goDocs) reset() {
	g.skip = 0
	g.total = 1
}
