package random

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/spf13/cobra"
)

const (
	exampleNumber = `
		# Generate a random number
		!random number`
)

type (
	number struct {
		min int
		max int
	}
)

func newRandomNumberCommand() *cobra.Command {
	n := newNumber()
	c := &cobra.Command{
		Use:     "number",
		Short:   "Generate a random number",
		Example: exampleNumber,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			val := fmt.Sprintf("%d", randomdata.Number(n.min, n.max))
			cmd.OutOrStdout().Write([]byte(val))
		},
	}

	c.Flags().IntVarP(&n.min, "min", "m", 0, "Minimum value")
	c.Flags().IntVarP(&n.max, "max", "a", 100, "Maximum value")

	return c
}

func newNumber() *number {
	return &number{}
}

func (n *number) reset() {
	n.min = 0
	n.max = 100
}
