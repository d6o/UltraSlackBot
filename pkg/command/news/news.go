package news

import (
	"github.com/disiqueira/ultraslackbot/pkg/newsapi"
	"strings"

	"fmt"

	"github.com/spf13/cobra"
)

const (
	newsLong = `
		Live top and breaking headlines for a country, specific category in a country, 
		single source, or multiple sources. You can also search with keywords. Articles 
		are sorted by the earliest date published first`

	newsExample = `
		# Get the latest articles about "Corinthians" 
		!news Corinthians

		# Get the latest articles about country (2-letter ISO 3166-1)
		!news --country br

		# Get the latest business articles (business entertainment general health science sports technology) 
		!news --category business

		# Get the second item from a search
		!news Corinthians --size 1

		# Get the third and fourth items from a search
		!news Corinthians --size 2 --page 2`
)

type (
	news struct {
		newsapi  *newsapi.NewsAPI
		page     int
		size     int
		country  string
		category string
		q        string
	}
)

func NewNewsCommand(key string) *cobra.Command {
	news := &news{
		newsapi: newsapi.New(key),
	}

	c := &cobra.Command{
		Use:     "news QUERY",
		Short:   "search for top breaking news",
		Long:    newsLong,
		Example: newsExample,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := news.Search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			_, _ = cmd.OutOrStdout().Write([]byte(r))
			news.Reset()
		},
	}

	c.Flags().IntVarP(&news.page, "page", "p", 0, "Which page do you want")
	c.Flags().IntVarP(&news.size, "size", "s", 1, "How many articles per page")
	c.Flags().StringVarP(&news.country, "country", "c", "", "Filter by country")
	c.Flags().StringVarP(&news.category, "category", "a", "", "Filter by category")

	return c
}

func (n *news) Search(q string) (string, error) {
	r, err := n.newsapi.Headlines(n.country, n.category, q, n.size, n.page)
	if err != nil {
		return "", err
	}
	if r.Status != "ok" {
		return r.Status, nil
	}
	if r.TotalResults <= 0 {
		return "No results", nil
	}
	msg := ""
	for _, article := range r.Articles {
		msg += fmt.Sprintf("%s (%s)\n", article.URL, article.PublishedAt.Format("02-01-2006"))
	}
	return msg, nil
}

func (n *news) Reset() {
	n.page = 0
	n.size = 1
	n.category = ""
	n.country = ""
}
