package pokequest

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/spf13/cobra"
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
			num, tip, err := h.tip()
			if err != nil {
				tip = err.Error()
			}
			if num >= 0 {
				pokeData.tips[num] = tip
				pokeData.points--
			}
			cmd.OutOrStdout().Write([]byte(fmt.Sprintf("*Points: %d*\n\n", pokeData.points)))
			cmd.OutOrStdout().Write([]byte(fmt.Sprintf("%d - %s", num, tip)))
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

func (s *tip) tip() (int, string, error) {
	if s.pokeData.current == 0 {
		return 0, "", fmt.Errorf("use '!pokequest start' to start a new game")
	}

	pokemon, err := s.pokeData.api.Pokemon(s.pokeData.current)
	if err != nil {
		return 0, "", err
	}

	characteristics, _ := s.pokeData.api.Characteristic(s.pokeData.current)
	species, _ := s.pokeData.api.Species(s.pokeData.current)

	tip := s.nextTip()
	switch tip {
	case 1:
		return tip, fmt.Sprintf("The Pokemon is %.1f meters tall", float32(pokemon.Height)/10), nil
	case 2:
		return tip, fmt.Sprintf("The Pokemon weighs %.1f kilos", float32(pokemon.Weight)/10), nil
	case 3:
		return tip, fmt.Sprintf("If you defeat this Pokemon you'll get %d xp", pokemon.BaseExperience), nil
	case 4:
		return tip, fmt.Sprintf("This is a %s Pokemon", strings.Join(pokemon.AllTypes(), ", ")), nil
	case 5:
		return tip, fmt.Sprintf("This is a strong one, look at this stats: %s ", strings.Join(pokemon.AllStats(), ", ")), nil
	case 6:
		return tip, fmt.Sprintf("The best way to describe the Pokemon your looking for: %s ", strings.Join(characteristics.AllDescriptions(), ", ")), nil
	case 7:
		return tip, fmt.Sprintf("You got the happiest Pokemon on Earth: %d ", species.BaseHappiness), nil
	case 8:
		if species.HasGenderDifferences {
			return tip, "This Pokemon has gender differences", nil
		}
		return tip, "This Pokemon does not has gender differences", nil
	case 9:
		if species.IsBaby {
			return tip, "This a baby Pokemon", nil
		}
		return tip, "This an adult Pokemon", nil
	case 10:
		return tip, fmt.Sprintf("This is a beutiful %s Pokemon", species.Color.Name), nil
	case 11:
		return tip, fmt.Sprintf("If you wanna bread this Pokemon, start looking for some %s eggs", strings.Join(species.AllEggGroups(), ", ")), nil
	case 12:
		return tip, fmt.Sprintf("This Pokemon belongs to the BEST generation: %s", species.Generation.Name), nil
	case 13:
		return tip, fmt.Sprintf("Prepare your bagpack, to find your Pokemon you need to get to some %s", species.Habitat.Name), nil
	case 14:
		return tip, fmt.Sprintf("I don't know if this is gonna help you, but this is the Growth rate of your Pokemon: %s", species.GrowthRate.Name), nil
	case 15:
		return tip, fmt.Sprintf("We can categorize every Pokemon based on his shape, the Pokemon you're looking for has a %s shape", species.Shape.Name), nil
	case 0:
		if species.EvolvesFromSpecies.Name != "" {
			return tip, "This Pokemon is an evolution.", nil
		}

		return tip, "This is the first state of this Pokemon.", nil
	default:
		return -1, "No more tips for you.", nil
	}

	return -1, "oops error :(", nil
}

func (s *tip) nextTip() int {
	for {
		i := rand.Intn(16)
		if s.pokeData.tips[i] == "" {
			return i
		}
		if len(s.pokeData.tips) > 15 {
			return -1
		}
	}
}
