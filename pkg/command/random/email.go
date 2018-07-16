package random

import (
	"github.com/spf13/cobra"
	"github.com/Pallinder/go-randomdata"
)

const (
	exampleEmail = `
		# Generate a random email
		!random email`
)

func newRandomEmailCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "email",
		Short:   "Generate a random email",
		Example: exampleEmail,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(randomdata.Email()))
		},
	}

	return c
}
