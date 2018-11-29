package pokequest

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
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
		s.pokeData.points -= 2
		return fmt.Sprintf("Sorry but %s is not the right answer.", guess), nil
	}

	finalPoints := s.pokeData.points
	if finalPoints < 0 {
		finalPoints = 0
	}

	msg := fmt.Sprintf("You are the best Pokemon Master ever! Keep up your training \n *You got %d Points!! (beta)* %s", finalPoints, resp.Sprites.FrontDefault)
	return msg, nil
}
