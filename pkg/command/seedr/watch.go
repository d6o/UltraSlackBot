package seedr

import (
	"strconv"

	"github.com/spf13/cobra"
)

func NewSeedrWatchCommand(m *seedrManager) *cobra.Command {
	c := &cobra.Command{
		Use:     "watch FILE_ID",
		Short:   "Get the Stream URL to watch a file",
		Example: example,
		Args:    cobra.ExactArgs(1),
		Run:     runWatch(m),
	}

	return c
}

func runWatch(m *seedrManager) runFunc {
	return func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.OutOrStdout().Write([]byte(err.Error()))
			return
		}

		f, err := m.HLS(id)
		if err != nil {
			f = err.Error()
		}
		cmd.OutOrStdout().Write([]byte(f))
	}
}
