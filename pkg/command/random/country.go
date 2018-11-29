package random

import (
	"errors"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

const (
	exampleCountry = `
		# Generate a random country
		!random country`
)

type (
	country struct {
		fullName     bool
		twoLetters   bool
		threeLetters bool
	}
)

func newRandomCountryCommand() *cobra.Command {
	co := newCountry()
	c := &cobra.Command{
		Use:     "country",
		Short:   "Generate a random country",
		Example: exampleCountry,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			country, err := co.generate()
			if err != nil {
				country = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(country))
		},
	}

	c.Flags().BoolVarP(&co.fullName, "fullname", "f", true, "Print a country with full text representation")
	c.Flags().BoolVarP(&co.twoLetters, "two", "2", false, "Print a country using ISO 3166-1 alpha-2")
	c.Flags().BoolVarP(&co.threeLetters, "three", "3", false, "Print a country using ISO 3166-1 alpha-3")

	return c
}

func newCountry() *country {
	return &country{}
}

func (c *country) generate() (string, error) {
	if c.hasIllegalFlags() {
		return "", errors.New("--fullname --three and --two are mutually exclusives")
	}
	if c.twoLetters {
		return randomdata.Country(randomdata.TwoCharCountry), nil
	}
	if c.threeLetters {
		return randomdata.Country(randomdata.ThreeCharCountry), nil
	}

	return randomdata.Country(randomdata.FullCountry), nil
}

func (c *country) hasIllegalFlags() bool {
	if c.fullName == c.threeLetters {
		return c.fullName
	}

	return c.twoLetters
}
