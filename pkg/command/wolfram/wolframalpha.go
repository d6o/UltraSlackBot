package wolfram

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"net/url"
	"github.com/disiqueira/ultraslackbot/pkg/command"
	"errors"
)

const (
	urlWolfram = "https://api.wolframalpha.com/v2/query?input=%s&appid=%s&output=JSON"

	example = `
		# Ask a question
		!wolfram Who is the president of Brazil?

		# Ask another question
		!wolfram What is the distance between the Earth and the Moon?

		# Solve math problems
		!calc 2+2`
)

type (
	wolfram struct {
		key string
	}

	wolframResponse struct {
		QueryResult struct {
			Success bool `json:"success"`
			Error   bool `json:"error"`
			Pods    []struct {
				Title   string `json:"title"`
				Primary bool `json:"primary,omitempty"`
				SubPods []struct {
					Plaintext string `json:"plaintext"`
				} `json:"subpods"`
			} `json:"pods"`
		} `json:"queryresult"`
	}
)

func NewWolframCommand(key string) *cobra.Command {
	w := newWolfram(key)

	c := &cobra.Command{
		Use:     "wolfram",
		Short:   "Ask a question to Wolfram|Alpha",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := w.ask(strings.Join(args, " "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
		Aliases: []string{"wa", "calc", "ask"},
	}

	return c
}

func newWolfram(key string) *wolfram {
	return &wolfram{
		key:key,
	}
}

func (w *wolfram) ask(q string) (string, error) {
	query := url.QueryEscape(q)

	data := &wolframResponse{}
	u := fmt.Sprintf(urlWolfram, query, w.key)
	err := command.GetJSON(u, data)
	if err != nil {
		return "", err
	}

	if data.QueryResult.Error || !data.QueryResult.Success {
		return "", errors.New("error while parsing the result")
	}

	if len(data.QueryResult.Pods) < 1 {
		return "", errors.New("no answer :(")
	}

	var title string
	var pods []string
	for _, pod := range data.QueryResult.Pods {
		if !pod.Primary {
			continue
		}

		title = pod.Title

		for _, subPod := range pod.SubPods {
			text := subPod.Plaintext
			if text != "" {
				pods = append(pods, text)
			}
		}
	}

	return fmt.Sprintf("%s: %s", title, strings.Join(pods, ", ")), nil
}
