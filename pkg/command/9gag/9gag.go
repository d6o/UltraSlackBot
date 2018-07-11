package ninegag

import (
	"errors"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
)

const (
	randomURL = "https://9gag.com/random"
	example = `
		# Get a random 9gag page
		!9gag`
)

func New9gagCommand() *cobra.Command {
	c := &cobra.Command{
		Use:     "9gag",
		Short:   "Get a random 9gag page",
		Example: example,
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			result, err := nineGag()
			if err != nil {
				result = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(result))
		},
	}

	return c
}

func nineGag() (string, error) {
	redirectNotAllowed := errors.New("redirect")
	redirectedURL := ""

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectedURL = req.URL.String()
		return redirectNotAllowed
	}

	_, err := client.Get(randomURL)
	if urlError, ok := err.(*url.Error); !ok || urlError.Err != redirectNotAllowed {
		return "", err
	}
	return redirectedURL, nil
}
