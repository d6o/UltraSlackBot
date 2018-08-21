package watch

import (
	"github.com/disiqueira/ultraslackbot/pkg/seedr"
	"github.com/disiqueira/ultraslackbot/pkg/yts"
	"github.com/spf13/cobra"
	"strings"
)

const (
	example = `# TODO`
)

func NewWatchCommand(username, password string) *cobra.Command {
	s := seedr.New(username, password)
	y := yts.New()
	m := newSeedrManager(s, y)

	c := &cobra.Command{
		Use:     "watch",
		Short:   "Download and watch movies",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			f, err := m.Watch(strings.Join(args, " "))
			if err != nil {
				f = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(f))
		},
	}

	return c
}
