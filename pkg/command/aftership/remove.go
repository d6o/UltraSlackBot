package afterShip

import "github.com/spf13/cobra"

const (
	exampleRemove = `
		# Remove a tracking
		remove CDFBCCD75CB44A`
)

func newAfterShipRemoveCommand(af *afterShip) *cobra.Command {
	c := &cobra.Command{
		Use:     "remove TRACKING_NUMBER",
		Short:   "Remove a tracking on AfterShip",
		Example: exampleRemove,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			msg, err := af.remove(args[0])
			if err != nil {
				msg = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(msg))
		},
	}

	return c
}
