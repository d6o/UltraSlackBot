package hello

import (
	"github.com/spf13/cobra"
)

const (
	example = `
		# Say Hello!
		!hello`
)

func NewHelloCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "hello",
		Short:   "Say hello",
		Example: example,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.OutOrStdout().Write([]byte("Hello!"))
		},
		Aliases: []string{"hey", "hi"},
	}

	return c
}
