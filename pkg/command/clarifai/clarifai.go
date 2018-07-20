package clarifai

import (
	"fmt"
	"strings"

	"encoding/json"
	"errors"
	"github.com/parnurzeal/gorequest"
	"github.com/spf13/cobra"
	"time"
)

const (
	urlClarifai = "https://api.clarifai.com/v2/models/%s/outputs"
	clarifaiOK  = "Ok"

	example = `
		# Ask a question
		!clarifai Who is the president of Brazil?

		# Ask another question
		!clarifai What is the distance between the Earth and the Moon?

		# Solve math problems
		!calc 2+2`
)

type (
	Concepts map[string]float64

	clarifai struct {
		key   string
		model string

		min float64

		total int
		skip  int
	}

	clarifaiRequest struct {
		Inputs []clarifaiInput `json:"inputs"`
	}

	clarifaiInput struct {
		Data clarifaiData `json:"data"`
	}

	clarifaiData struct {
		Image clarifaiImage `json:"image"`
	}

	clarifaiImage struct {
		URL string `json:"url"`
	}

	clarifaiResponse struct {
		Status struct {
			Code        int    `json:"code"`
			Description string `json:"description"`
		} `json:"status"`
		Outputs []struct {
			ID     string `json:"id"`
			Status struct {
				Code        int    `json:"code"`
				Description string `json:"description"`
			} `json:"status"`
			CreatedAt time.Time `json:"created_at"`
			Model     struct {
				ID         string    `json:"id"`
				Name       string    `json:"name"`
				CreatedAt  time.Time `json:"created_at"`
				AppID      string    `json:"app_id"`
				OutputInfo struct {
					Message string `json:"message"`
					Type    string `json:"type"`
					TypeExt string `json:"type_ext"`
				} `json:"output_info"`
				ModelVersion struct {
					ID        string    `json:"id"`
					CreatedAt time.Time `json:"created_at"`
					Status    struct {
						Code        int    `json:"code"`
						Description string `json:"description"`
					} `json:"status"`
				} `json:"model_version"`
				DisplayName string `json:"display_name"`
			} `json:"model"`
			Input struct {
				ID   string `json:"id"`
				Data struct {
					Image struct {
						URL string `json:"url"`
					} `json:"image"`
				} `json:"data"`
			} `json:"input"`
			Data struct {
				Concepts []struct {
					ID    string  `json:"id"`
					Name  string  `json:"name"`
					Value float64 `json:"value"`
					AppID string  `json:"app_id"`
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
			r, err := cla.Analyze(args)
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

func (c *clarifai) Analyze(urls []string) (string, error) {
	conceptList, err := c.analyze(urls)
	if err != nil {
		return "", err
	}

	var msgList []string
	skip := 0
	for url, concepts := range conceptList {
		if c.skip > skip {
			skip++
			continue
		}

		msg := url + "\n" + concepts.String()
		msgList = append(msgList, msg)

		if len(msgList) == c.total {
			break
		}
	}

	return strings.Join(msgList, "\n"), nil
}

func (c *clarifai) analyze(urls []string) (map[string]Concepts, error) {
	body, err := c.makeRequest(c.createRequest(urls))
	if err != nil {
		return nil, err
	}

	r := clarifaiResponse{}
	if err := json.Unmarshal([]byte(body), &r); err != nil {
		return nil, err
	}

	return c.readConcepts(&r)
}

func (c *clarifai) readConcepts(response *clarifaiResponse) (map[string]Concepts, error) {
	if response.Status.Description != clarifaiOK {
		return nil, errors.New("no concepts found")
	}

	con := map[string]Concepts{}
	for _, outputs := range response.Outputs {
		for _, concept := range outputs.Data.Concepts {

			if _, ok := con[outputs.Input.Data.Image.URL]; !ok {
				con[outputs.Input.Data.Image.URL] = Concepts{}
			}

			if concept.Value > c.min {
				con[outputs.Input.Data.Image.URL][concept.Name] = concept.Value
			}
		}
	}
	return con, nil
}

func (c *clarifai) createRequest(urls []string) *clarifaiRequest {
	r := &clarifaiRequest{}
	for _, u := range urls {
		r.Inputs = append(r.Inputs, clarifaiInput{
			Data: clarifaiData{
				Image: clarifaiImage{
					URL: u,
				},
			},
		})
	}
	return r
}

func (c *clarifai) makeRequest(r *clarifaiRequest) (string, error) {
	_, body, errs := gorequest.New().Post(urlClarifai).
		Set("Authorization", "Key "+c.key).
		Send(r).
		End()

	if errs != nil {
		return "", errs[0]
	}

	return body, nil
}

func (co Concepts) String() string {
	var final []string

	for key, value := range co {
		final = append(final, fmt.Sprintf("%s %s(%6.4f)", final, key, value))
	}

	return strings.Join(final, " ")
}
