package seedr

import (
	"github.com/disiqueira/ultraslackbot/pkg/seedr"
	"github.com/spf13/cobra"
)

const (
	example = `# TODO`
)

func NewSeedrCommand(username, password string) *cobra.Command {
	s := seedr.New(username, password)
	m := newSeedrManager(s)

	c := &cobra.Command{
		Use:     "seedr",
		Short:   "Download and watch movies",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := m.Folders()
			if err != nil {
				f = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(f))
		},
	}

	c.AddCommand(NewSeedrFoldersCommand(m))
	c.AddCommand(NewSeedrFolderCommand(m))
	c.AddCommand(NewSeedrWatchCommand(m))

	return c
}
