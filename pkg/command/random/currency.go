package random

import (
	"github.com/spf13/cobra"
	"github.com/Pallinder/go-randomdata"
)

const (
	exampleCurrency = `
		# Generate a random currency
		!random currency`
)

func newRandomCurrencyCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "currency",
		Short:   "Generate a random currency",
		Example: exampleCurrency,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(randomdata.Currency()))
		},
	}

	return c
}
