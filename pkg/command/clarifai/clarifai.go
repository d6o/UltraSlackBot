package clarifai

import (
	"errors"
	"bytes"
	"net/http"
	"fmt"
	"strings"
	"encoding/json"

	"github.com/spf13/cobra"
	"mvdan.cc/xurls"
	"io/ioutil"
)

const (
	urlClarifai = "https://api.clarifai.com/v2/models/%s/outputs"
	requestClarifai = "{\"inputs\":[{\"data\":{\"image\":{\"url\":\"%s\"}}}]}"
	okClarifai  = "Ok"

	example = `
		# Ask a question
		!clarifai Who is the president of Brazil?

		# Ask another question
		!clarifai What is the distance between the Earth and the Moon?

		# Solve math problems
		!calc 2+2`
)

type (
	clarifai struct {
		key   string
		model string

		min float64

		total int
		skip  int
	}

	clarifaiResponse struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Outputs []struct {
			Data struct {
				Concepts []struct {
					Name  string  `json:"name"`
					Value float64 `json:"value"`
				} `json:"concepts"`
			} `json:"data"`
		} `json:"outputs"`
	}
)

func NewClarifaiCommand(key, model string) *cobra.Command {
	cla := newClarifai(key, model)

	c := &cobra.Command{
		Use:     "clarifai",
		Short:   "Ask a question to Clarifai|Alpha",
		Example: example,
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			r, err := cla.Analyze(strings.Join(args," "))
			if err != nil {
				r = err.Error()
			}
			cmd.OutOrStdout().Write([]byte(r))
		},
		Aliases: []string{"wa", "calc", "ask"},
	}

	c.Flags().Float64VarP(&cla.min, "min-value", "m", 0.3, "Minimum value to print the concept")
	c.Flags().IntVarP(&cla.total, "total", "t", 1, "Total of results to print")
	c.Flags().IntVarP(&cla.skip, "skip", "s", 0, "How many results should be skipped")

	return c
}

func newClarifai(key, model string) *clarifai {
	return &clarifai{
		key:   key,
		model: model,
	}
}

func (c *clarifai) Analyze(url string) (string, error) {
	urls := xurls.Strict.FindAllString(url, 1)
	if len(urls) == 0 {
		return "", errors.New("no valid url found")
	}

	jsonStr := fmt.Sprintf(requestClarifai, urls[0])
	url = fmt.Sprintf(urlClarifai, c.model)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonStr)))
	req.Header.Set("Authorization", "Key "+c.key)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	fmt.Println(string(body))

	data := clarifaiResponse{}
	json.NewDecoder(req.Body).Decode(data)

	if data.Status.Description != okClarifai {
		return "", fmt.Errorf("invalid response from Clarifai: Code: %d Description: %s", data.Status.Code, data.Status.Description)
	}

	var msgList []string
	skip := 0
	for _, concept := range data.Outputs {
		if len(concept.Data.Concepts) == 0 {
			continue
		}

		for _, con := range concept.Data.Concepts {
			if con.Value < c.min {
				continue
			}
			if c.skip > skip {
				skip++
				continue
			}
			msgList = append(msgList, fmt.Sprintf("%s(%6.4f)", con.Name, con.Value))

			if len(msgList) == c.total {
				break
			}
		}
	}

	return strings.Join(msgList, "\n"), nil

}

