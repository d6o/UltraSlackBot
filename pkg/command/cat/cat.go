package cat

import (
	"github.com/spf13/cobra"
)

func NewCatCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "cat",
		Short: "Everything cat related",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	c.AddCommand(newCatFactCommand())
	c.AddCommand(newCatGifCommand())
	c.AddCommand(newCatPhotoCommand())

	return c
}
