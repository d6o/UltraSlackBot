package google

import (
	"strings"

	"github.com/spf13/cobra"
	"fmt"
)

const (
	searchLong = `
		search make a custom search on google custom search
		and return some of the results.`

	searchExample = `
		# Post the first result for dogs
		!search dogs

		# Post the first result for hot dog
		!search hot dog

		# Post two results for hot dog
		!search hot dog --total 2

		# Post the second result for hot dog
		!search hot dog --skip 1

		# Post the third and fourth results for hot dog
		!search hot dog --skip 2 --total 2`
)

type (
	googleSearch struct {
		google
	}
)

func NewGoogleSearchCommand(key, cx string) *cobra.Command {
	s := newGoogleSearch(key, cx)

	c := &cobra.Command{
		Use:     "search QUERY",
		Short:   "search search for searchs using Google",
		Long:    searchLong,
		Example: searchExample,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := s.Search(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
			s.reset()
		},
		Aliases: []string{"google", "gse", "busca", "g"},
	}

	c.Flags().IntVarP(&s.total, "total", "t",1, "How many results will be returned")
	c.Flags().IntVarP(&s.skip, "skip","s", 0, "How many results should be skipped")

	return c
}

func newGoogleSearch(key, cx string) *googleSearch {
	return &googleSearch{
		google{
			key: key,
			cx : cx,
		},
	}
}

func (gs *googleSearch) Search(q string) (string, error) {
	r, err := gs.search(q, nil)
	if err != nil {
		return "", err
	}

	var msgList []string
	for _, item := range r {
		msgList = append(msgList, fmt.Sprintf("%s - %s", item.Title, item.Link))
	}

	return strings.Join(msgList, "\n"), nil
}
