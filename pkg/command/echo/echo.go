package echo

import (
	"strings"

	"github.com/CrowdSurge/banner"
	"github.com/spf13/cobra"
)

func NewEchoCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "echo",
		Short:   "Write a quote in a banner",
		Example: "# Write a Hello World in a banner\n!echo Hello World",
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			text := "```" + banner.PrintS(strings.Join(args, " ")) + "```"
			cmd.OutOrStdout().Write([]byte(text))
		},
	}
}
