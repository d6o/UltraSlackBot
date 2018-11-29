package cat

import (
	"github.com/spf13/cobra"

	"github.com/disiqueira/ultraslackbot/pkg/command"
)

const (
	urlCatFact     = "http://catfact.ninja/fact"
	exampleCatFact = `
		# Get a random CatFact page
		!cat fact`
)

type (
	catFactResponse struct {
		Fact string `json:"fact"`
	}
)

func newCatFactCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "fact",
		Short:   "Get a random fact about cats",
		Example: exampleCatFact,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			data := &catFactResponse{}
			if err := command.GetJSON(urlCatFact, data); err != nil {
				cmd.OutOrStdout().Write([]byte(err.Error()))
				return
			}
			cmd.OutOrStdout().Write([]byte(data.Fact))
		},
	}

	return c
}
