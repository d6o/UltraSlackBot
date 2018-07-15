package shrug

import (
	"github.com/spf13/cobra"
)

const (
	original = "¯\\_(ツ)_/¯"
)

func NewShrugCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "shrug",
		Short:   "Get a Shrug quote",
		Example: "# Get a shrug\n!shrug",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(original))
		},
	}

	return c
}
