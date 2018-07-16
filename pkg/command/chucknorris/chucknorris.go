package chucknorris

import (
	"github.com/spf13/cobra"

	"github.com/disiqueira/ultraslackbot/pkg/command"
)

const (
	urlChuckNorrisFact = "https://api.icndb.com/jokes/random"
	exampleChuckNorrisFact = `
		# Get a random fact about Chuck Norris
		!chuckNorris`
)

type (
	chuckNorrisFactResponse struct {
		Value struct {
			Joke       string        `json:"joke"`
		} `json:"value"`
	}
)

func NewChuckNorrisFactCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "chucknorris",
		Short:   "Get a random fact about Chuck Norris",
		Example: exampleChuckNorrisFact,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			data := &chuckNorrisFactResponse{}
			if err := command.GetJSON(urlChuckNorrisFact, data); err != nil {
				cmd.OutOrStdout().Write([]byte(err.Error()))
				return
			}
			cmd.OutOrStdout().Write([]byte(data.Value.Joke))
		},
		Aliases: []string{"chuck", "cn"},
	}

	return c
}
