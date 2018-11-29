package lenny

import (
	"fmt"
	"strings"

	"github.com/disiqueira/ultraslackbot/pkg/command"
	"github.com/spf13/cobra"
)

const (
	lennyURL = "https://api.lenny.today/v1/random?limit=%d"
	maxTotal = 5
	example  = `
		# Get a random Lenny
		!lenny random

		# Get 5 random Lenny
		!lenny random --total 5`
)

type (
	lenny struct {
		total int
	}

	lennyResponse []struct {
		Face string `json:"face"`
	}
)

func newRandomCommand() *cobra.Command {
	h := newLenny()
	c := &cobra.Command{
		Use:     "random",
		Short:   "Get random lennys",
		Example: example,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			result, err := h.search()
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
			h.reset()
		},
		Aliases: []string{"lenny"},
	}

	c.Flags().IntVarP(&h.total, "total", "t", 1, "Total of results to print")

	return c
}

func newLenny() *lenny {
	return &lenny{}
}

func (h *lenny) search() (string, error) {
	if h.total > maxTotal {
		h.total = maxTotal
	}
	data := lennyResponse{}
	url := fmt.Sprintf(lennyURL, h.total)
	if err := command.GetJSON(url, &data); err != nil {
		return "", err
	}

	var faces []string
	for _, f := range data {
		faces = append(faces, f.Face)
	}

	return strings.Join(faces, "\n"), nil

}

func (h *lenny) reset() {
	h.total = 1
}
