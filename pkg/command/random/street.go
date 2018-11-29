package random

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

const (
	exampleStreet = `
		# Generate a random street
		!random street`
)

func newRandomStreetCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "street",
		Short:   "Generate a random street",
		Example: exampleStreet,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(randomdata.Street()))
		},
	}

	return c
}
