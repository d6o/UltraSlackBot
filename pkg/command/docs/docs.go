package docs

import "github.com/spf13/cobra"

func NewDocsCommand() *cobra.Command {
	c := &cobra.Command{
		Use:   "docs",
		Short: "Search for docs",
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	c.AddCommand(newDocsGoCommand())

	return c
}
