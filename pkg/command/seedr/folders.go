package seedr

import (
	"github.com/spf13/cobra"
)

func NewSeedrFoldersCommand(m *seedrManager) *cobra.Command {
	c := &cobra.Command{
		Use:     "folders",
		Short:   "List all folders",
		Example: example,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			f, err := m.Folders()
			if err != nil {
				f = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(f))
		},
	}

	return c
}
