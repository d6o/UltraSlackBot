package uuid

import (
	"github.com/beevik/guid"
	"github.com/spf13/cobra"
)

func NewUUIDCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "uuid",
		Short:   "Generate an UUID",
		Example: "# Generate a new UUID\n!uuid",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(guid.NewString()))
		},
	}
	return c
}
