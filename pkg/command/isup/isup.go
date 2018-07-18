package isup

import (
	"errors"
	"github.com/spf13/cobra"
	"net/http"
	"net/url"
	"time"
	"fmt"
)

const (
	downFormat = "%s is down: %s"
	upFormat = "%s is up!"

	example = `
		# Check if google is available to the bot server
		!isup google.com`
)

type (
	checker struct {
		client *http.Client
	}
)

func NewIsUpCommand() *cobra.Command {
	client := &http.Client{
		Timeout: time.Second * 3,
	}
	checker := newChecker(client)
	c := &cobra.Command{
		Use:     "isup",
		Short:   "Verify if a page is up or down",
		Example: example,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			msg := fmt.Sprintf(upFormat, args[0])
			if err := checker.isUp(args[0]); err != nil {
				msg = fmt.Sprintf(downFormat, args[0], err.Error())
			}
			cmd.OutOrStdout().Write([]byte(msg))
		},
	}

	return c
}

func newChecker(c *http.Client) *checker {
	return &checker{
		client:c,
	}
}

func (c *checker) isUp(link string) error {
	_, err := c.client.Get(link)
	return err
}
