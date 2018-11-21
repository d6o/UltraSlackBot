package pokequest

import (
	"fmt"
	"github.com/spf13/cobra"
	"math/rand"
	"strings"
)

type (
	tip struct {
		pokeData *pokeData
	}
)

func newTipCommand(pokeData *pokeData) *cobra.Command {
	h := newTip(pokeData)
	c := &cobra.Command{
		Use:     "tip",
		Short:   "Get a new tip",
		Example: "tip",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			result, err := h.tip()
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
		},
		Aliases: []string{"tip"},
	}

	return c
}

func newTip(pokeData *pokeData) *tip {
	return &tip{
		pokeData: pokeData,
	}
}

func (s *tip) tip() (string, error) {
	if s.pokeData.current == 0 {
		return "", fmt.Errorf("use '!pokequest start' to start a new game")
	}

	pokemon, err := s.pokeData.api.Pokemon(s.pokeData.current)
	if err != nil {
		return "", err
	}

	characteristics, _ := s.pokeData.api.Characteristic(s.pokeData.current)
	species, _ := s.pokeData.api.Species(s.pokeData.current)

	switch s.nextTip() {
	case 1:
		return fmt.Sprintf("The secret Pokemon is %.1f meters tall", float32(pokemon.Height)/10), nil
	case 2:
		return fmt.Sprintf("The secret Pokemon weighs %.1f kilos", float32(pokemon.Weight)/10), nil
	case 3:
		return fmt.Sprintf("If you defeat this Pokemon you'll get %d xp", pokemon.BaseExperience), nil
	case 4:
		return fmt.Sprintf("Types: %s ", strings.Join(pokemon.AllTypes(), ", ")), nil
	case 5:
		return fmt.Sprintf("Stats: %s ", strings.Join(pokemon.AllStats(), ", ")), nil
	case 6:
		return fmt.Sprintf("Descriptions: %s ", strings.Join(characteristics.AllDescriptions(), ", ")), nil
	case 7:
		return fmt.Sprintf("Base Happiness: %d ", species.BaseHappiness), nil
	case 8:
		if species.HasGenderDifferences {
			return "This Pokemon has male and female gender", nil
		}
		return "This is a genderless Pokemon", nil
	case 9:
		if species.IsBaby {
			return "This a baby Pokemon", nil
		}
		return "This an adult Pokemon", nil
	case 10:
		return fmt.Sprintf("This a %s Pokemon", species.Color.Name), nil
	case 11:
		return fmt.Sprintf("Egg-type: %s", strings.Join(species.AllEggGroups(), ", ")), nil
	case 12:
		return fmt.Sprintf("Generation: %s", species.Generation.Name), nil
	case 13:
		return fmt.Sprintf("Habitat: %s", species.Habitat.Name), nil
	case 14:
		return fmt.Sprintf("Growth Rate: %s", species.GrowthRate.Name), nil
	case 15:
		return fmt.Sprintf("Shape: %s", species.Shape.Name), nil
	case 0:
		if species.EvolvesFromSpecies.Name != "" {
			return "This Pokemon is an evolution.", nil
		}

		return "This is the first state of this Pokemon.", nil
	default:
		return "No more tips for you.", nil
	}

	return "oops error :(", nil
}

func (s *tip) nextTip() int {
	for {
		i := rand.Intn(16)
		if !s.pokeData.tips[i] {
			s.pokeData.tips[i] = true
			return i
		}
		if len(s.pokeData.tips) > 15 {
			return -1
		}
	}
}
