package yts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://yts.am/api/v2/list_movies.json?"
)

type (
	YTS struct {
		httpClient *http.Client
	}
)

func New() *YTS {
	return &YTS{
		httpClient: &http.Client{},
	}
}

func (s *YTS) SetHTTPClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

type ListResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		Movies []struct {
			TitleLong        string `json:"title_long"`
			MediumCoverImage string `json:"medium_cover_image"`
			Torrents         []struct {
				URL string `json:"url"`
			} `json:"torrents"`
		} `json:"movies"`
	} `json:"data"`
}

func (s *YTS) Search(q string) (*ListResponse, error) {
	result := &ListResponse{}
	v := url.Values{}
	v.Add("query_term", q)
	return result, s.Get(v.Encode(), result)
}

func (s *YTS) Get(endpoint string, v interface{}) error {
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error making GET request to %s. StatusCode: %d", resp.Request.URL, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
