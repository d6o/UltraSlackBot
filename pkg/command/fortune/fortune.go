package fortune

import (
	"github.com/disiqueira/ultraslackbot/pkg/command"
	"github.com/spf13/cobra"
)

const (
	example = `
		# Get a fortune quote
		!fortune`
	url = "http://www.yerkee.com/api/fortune"
)

type (
	fortuneResponse struct {
		Fortune string `json:"fortune"`
	}
)

func NewFortuneCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "fortune",
		Short:   "Get a Fortune quote",
		Example: example,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			data := &fortuneResponse{}
			if err := command.GetJSON(url, data); err != nil {
				cmd.OutOrStdout().Write([]byte(err.Error()))
				return
			}
			cmd.OutOrStdout().Write([]byte(data.Fortune))
		},
	}

	return c
}
