package seedr

import (
	"github.com/spf13/cobra"
	"strconv"
)

type (
	runFunc func(cmd *cobra.Command, args []string)
)

func NewSeedrFolderCommand(m *seedrManager) *cobra.Command {
	c := &cobra.Command{
		Use:     "folder FOLDER_ID",
		Short:   "List all files in a folder",
		Example: example,
		Args:    cobra.ExactArgs(1),
		Run:     runFolder(m),
	}

	return c
}

func runFolder(m *seedrManager) runFunc {
	return func(cmd *cobra.Command, args []string) {
		id, err  := strconv.Atoi(args[0])
		if err != nil {
			cmd.OutOrStdout().Write([]byte(err.Error()))
			return
		}

		f, err := m.Folder(id)
		if err != nil {
			f = err.Error()
		}
		cmd.OutOrStdout().Write([]byte(f))
	}
}
