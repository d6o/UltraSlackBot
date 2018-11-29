package wikipedia

import (
	"fmt"
	"strings"

	"errors"
	"net/url"

	"github.com/disiqueira/ultraslackbot/pkg/command"
	"github.com/spf13/cobra"
)

const (
	urlWikipedia = "https://%s.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&titles=%s"

	example = `
		# Search for an article on Wikipedia
		!wikipedia Brazil`
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
	}

	wikipediaResponse struct {
		Query struct {
			Pages map[string]page `json:"pages"`
		} `json:"query"`
	}

	page struct {
		Extract string `json:"extract"`
		Title   string `json:"title"`
	}
)

func NewWikipediaCommand() *cobra.Command {
	w := newWikipedia()

	c := &cobra.Command{
		Use:     "wikipedia",
		Short:   "Ask a question to Wikipedia|Alpha",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := w.Search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
		Aliases: []string{"wk", "wiki"},
	}

	return c
}

func newWikipedia() *wikipedia {
	return &wikipedia{}
}

func (w *wikipedia) Search(q string) (string, error) {
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
	err := command.GetJSON(u, data)
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
