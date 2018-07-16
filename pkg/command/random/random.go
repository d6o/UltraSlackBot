package random

import (
	"github.com/spf13/cobra"
)

func NewRandomCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "random",
		Short:   "Generate random data",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	c.AddCommand(newRandomNameCommand())
	c.AddCommand(newRandomTitleCommand())
	c.AddCommand(newRandomEmailCommand())
	c.AddCommand(newRandomStreetCommand())
	c.AddCommand(newRandomCurrencyCommand())
	c.AddCommand(newRandomAddressCommand())
	c.AddCommand(newRandomCityCommand())
	c.AddCommand(newRandomCountryCommand())
	c.AddCommand(newRandomNumberCommand())

	return c
}
