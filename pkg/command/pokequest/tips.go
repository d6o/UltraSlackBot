package pokequest

import (
	"fmt"
	"github.com/spf13/cobra"
)

func newTipsCommand(pokeData *pokeData) *cobra.Command {
	c := &cobra.Command{
		Use:     "tips",
		Short:   "Get a new tips",
		Example: "tips",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			for num, tip := range pokeData.tips {
				cmd.OutOrStdout().Write([]byte(fmt.Sprintf("%d - %s\n", num, tip)))
			}
		},
		Aliases: []string{"tips"},
	}

	return c
}
