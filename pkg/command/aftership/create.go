package afterShip

import "github.com/spf13/cobra"

const (
	exampleCreate = `
		# Create a tracking
		create CDFBCCD75CB44A`
)

func newAfterShipCreateCommand(af *afterShip) *cobra.Command {
	c := &cobra.Command{
		Use:     "create TRACKING_NUMBER",
		Short:   "Create a tracking on AfterShip",
		Example: exampleCreate,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := af.create(args[0])
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
	}

	return c
}

