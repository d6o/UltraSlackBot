package choose

import (
	"math/rand"

	"github.com/spf13/cobra"
)

const (
	example = `
		# Choose between two foods
		!choose pizza feijoada

		# Choose between three options
		!choose opt1 opt2 opt3`
)

func NewChooseCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "choose OPT1 OPT2 [OPT3...]",
		Short:   "Choose",
		Example: example,
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte(args[rand.Intn(len(args))]))
		},
	}

	return c
}
