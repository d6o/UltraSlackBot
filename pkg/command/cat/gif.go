package cat

import (
	"net/http"

	"github.com/spf13/cobra"
)

const (
	urlCatGif = "http://thecatapi.com/api/images/get?format=src&type=gif"
	exampleCatGif = `
		# Get a random cat gif
		!cat gif`
)

func newCatGifCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "gif",
		Short:   "Get a random cat gif",
		Example: exampleCatGif,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := http.Get(urlCatGif)
			if err != nil {
				cmd.OutOrStdout().Write([]byte(err.Error()))
				return
			}
			cmd.OutOrStdout().Write([]byte(res.Request.URL.String()))
		},
	}

	return c
}
