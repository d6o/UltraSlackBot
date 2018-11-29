package random

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

const (
	exampleAddress = `
		# Generate a random address
		!random address`
)

func newRandomAddressCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "address",
		Short:   "Generate a random address",
		Example: exampleAddress,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(randomdata.Address()))
		},
	}

	return c
}
