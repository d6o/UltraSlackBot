package afterShip

import "github.com/spf13/cobra"

const (
	exampleList = `
		# List a tracking
		list CDFBCCD75CB44A`
)

func newAfterShipListCommand(af *afterShip) *cobra.Command {
	c := &cobra.Command{
		Use:     "list",
		Short:   "List a tracking on AfterShip",
		Example: exampleList,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			r, err := af.list()
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
	}

	return c
}

