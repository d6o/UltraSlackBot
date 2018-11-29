package afterShip

import "github.com/spf13/cobra"

const (
	exampleGet = `
		# Get a tracking
		get CDFBCCD75CB44A`
)

func newAfterShipGetCommand(af *afterShip) *cobra.Command {
	c := &cobra.Command{
		Use:     "get TRACKING_NUMBER",
		Short:   "Get a tracking on AfterShip",
		Example: exampleGet,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := af.get(args[0])
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
	}

	return c
}
