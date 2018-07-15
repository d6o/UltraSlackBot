package lenny

import (
	"github.com/spf13/cobra"
)

const (
	original = "( ͡° ͜ʖ ͡°)"
)

func NewLennyCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "lenny",
		Short:   "Get a Lenny quote",
		Example: "# Get a lenny\n!lenny",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(original))
		},
	}

	c.AddCommand(newRandomCommand())

	return c
}
