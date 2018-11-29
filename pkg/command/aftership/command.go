package afterShip

import (
	"github.com/spf13/cobra"
)

const (
	example = `
		# Ask a question
		!afterShip Who is the president of Brazil?

		# Ask another question
		!afterShip What is the distance between the Earth and the Moon?

		# Solve math problems
		!calc 2+2`
)

func NewAfterShipCommand(key string) *cobra.Command {
	as := newAfterShip(key)

	c := &cobra.Command{
		Use:     "packages",
		Short:   "Ask a question to AfterShip",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
		Aliases: []string{"track", "aftership"},
	}

	c.AddCommand(newAfterShipCreateCommand(as))
	c.AddCommand(newAfterShipGetCommand(as))
	c.AddCommand(newAfterShipListCommand(as))
	c.AddCommand(newAfterShipRemoveCommand(as))

	return c
}
