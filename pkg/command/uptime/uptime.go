package uptime

import (
	"time"

	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

func NewUptimeCommand() *cobra.Command {
	initTime := time.Now()
	c := &cobra.Command{
		Use:     "uptime",
		Short:   "Get a random Uptime page",
		Example: "uptime",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(humanize.Time(initTime)))
		},
	}

	return c
}
