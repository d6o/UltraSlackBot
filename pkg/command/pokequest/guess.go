package pokequest

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"strings"
)

type (
	guess struct {
		pokeData *pokeData
	}
)

func newGuessCommand(pokeData *pokeData) *cobra.Command {
	h := newGuess(pokeData)
	c := &cobra.Command{
		Use:     "guess",
		Short:   "Guess a Poke Quest",
		Example: "guess",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			result, err := h.guess(args[0])
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
		},
		Aliases: []string{"guess"},
	}

	return c
}

func newGuess(pokeData *pokeData) *guess {
	return &guess{
		pokeData: pokeData,
	}
}

func (s *guess) guess(guess string) (string, error) {
	resp, err := s.pokeData.api.Pokemon(s.pokeData.current)
	if err != nil {
		return "", err
	}

	guess = strings.TrimSpace(strings.ToLower(guess))

	if resp.Name != guess && strconv.Itoa(s.pokeData.current) != guess {
		return fmt.Sprintf("Sorry but %s is not the right answer.", guess), nil
	}

	return resp.Sprites.FrontDefault, nil
}
