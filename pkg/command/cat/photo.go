package cat

import (
	"net/http"

	"github.com/spf13/cobra"
)

const (
	urlCatPhoto = "http://thecatapi.com/api/images/get?format=src&type=jpg"
	examurlCatPhoto = `
		# Get a random cat photo
		!cat photo`
)

func newCatPhotoCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "photo",
		Short:   "Get a random cat photo",
		Example: examurlCatPhoto,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := http.Get(urlCatPhoto)
			if err != nil {
				cmd.OutOrStdout().Write([]byte(err.Error()))
				return
			}
			cmd.OutOrStdout().Write([]byte(res.Request.URL.String()))
		},
	}

	return c
}
