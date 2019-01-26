package pokequest

import (
	"github.com/disiqueira/ultraslackbot/pkg/pokeapi"
	"github.com/spf13/cobra"
)

type (
	pokeData struct {
		api     pokeapi.Poke
		current int
		limit   int
		tips    map[int]string
		points  int
	}
)

const (
	genLimit = 151
)

func NewPokeQuestCommand() *cobra.Command {
	pokeClient := pokeapi.New()
	//pokeClient.SetHTTPClient(NewHTTPCache(&http.Client{}))
	data := &pokeData{
		api:   pokeClient,
		limit: genLimit,
		tips:  map[int]string{},
	}

	c := &cobra.Command{
		Use:   "pokequest START",
		Short: "pokequest",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Aliases: []string{"poke", "pq", "pokemon"},
	}

	c.AddCommand(newStartCommand(data))
	c.AddCommand(newTipCommand(data))
	c.AddCommand(newGuessCommand(data))
	c.AddCommand(newTipsCommand(data))

	return c
}
