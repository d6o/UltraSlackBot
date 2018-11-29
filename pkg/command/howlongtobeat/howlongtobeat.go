package howlongtobeat

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

const (
	hltbURL = "https://howlongtobeat.com/search_main.php?page=1"
	example = `
		# How long takes to beat Life is Strange
		!howlongtobeat Life is Strange

		# How long takes to beat Life is Strange
		!hltb Life is Strange

		# How long takes to beat the three more popular Crash games
		!howlongtobeat Crash Bandicoot --total 3

		# How long takes to beat the second more popular Crash game
		!howlongtobeat Crash Bandicoot --skip 1

		# How long takes to beat the second and third more popular Crash games
		!howlongtobeat Crash Bandicoot --skip 1 --total 2`
)

type (
	hltb struct {
		total int
		skip  int
	}
)

func NewHLTBCommand() *cobra.Command {
	h := newHLTB()
	c := &cobra.Command{
		Use:     "howlongtobeat",
		Short:   "Search how long takes to beat a game",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := h.search(strings.Join(args, " "))
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
			h.reset()
		},
		Aliases: []string{"hltb"},
	}

	c.Flags().IntVarP(&h.total, "total", "t", 1, "Total of results to print")
	c.Flags().IntVarP(&h.skip, "skip", "s", 0, "How many results should be skipped")

	return c
}

func newHLTB() *hltb {
	return &hltb{}
}

func (h *hltb) search(text string) (string, error) {
	resp, err := http.PostForm(hltbURL,
		url.Values{
			"queryString": {text},
			"t":           {"games"},
			"sorthead":    {"popular"},
		})
	if err != nil {
		return "", err
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	list := map[string]map[string]string{}
	var typeList []string
	var timeList []string

	// @TODO: (disiqueira) This really needs a refactor.
	// Find the review items
	doc.Find("li").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".text_white").Text()
		list[name] = map[string]string{}

		for _, div := range s.Find(".shadow_text").Nodes {
			if len(strings.TrimSpace(div.FirstChild.Data)) < 1 {
				continue
			}
			typeList = append(typeList, div.FirstChild.Data)
		}

		for _, div := range s.Find(".center").Nodes {
			if len(strings.TrimSpace(div.FirstChild.Data)) < 1 {
				continue
			}
			timeList = append(timeList, div.FirstChild.Data)
		}

		for i := range typeList {
			list[name][typeList[i]] = timeList[i]
		}
	})

	var msgList []string
	skip := 0
	for name, game := range list {
		if h.skip > skip {
			skip++
			continue
		}

		msg := name + "\n"
		for t, v := range game {
			msg += t + ": " + v + "\n"
		}
		msgList = append(msgList, msg)

		if len(msgList) == h.total {
			break
		}
	}

	return "```" + strings.Join(msgList, "\n") + "```", nil
}

func (h *hltb) reset() {
	h.skip = 0
	h.total = 1
}
