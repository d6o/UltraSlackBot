package pokequest

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

type (
	start struct {
		pokeData *pokeData
	}
)

func newStartCommand(pokeData *pokeData) *cobra.Command {
	h := newStart(pokeData)
	c := &cobra.Command{
		Use:     "start",
		Short:   "Start a Poke Quest",
		Example: "start",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			result, err := h.start()
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
		},
		Aliases: []string{"start"},
	}

	return c
}

func newStart(pokeData *pokeData) *start {
	return &start{
		pokeData: pokeData,
	}
}

func (s *start) start() (string, error) {
	s.pokeData.current = rand.Intn(s.pokeData.limit)
	s.pokeData.tips = map[int]string{}
	fmt.Println(s.pokeData.current)
	return "Secret Pokemon I choose you! Use '!pokemon tip' to get tips about your Pokemon.", nil
}
