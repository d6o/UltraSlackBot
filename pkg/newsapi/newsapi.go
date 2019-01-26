package newsapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	urlHeadlines = "https://newsapi.org/v2/top-headlines"
)

type (
	NewsAPI struct {
		httpClient HTTP
		key        string
	}

	HTTP interface {
		Do(req *http.Request) (*http.Response, error)
	}

	HeadlinesResponse struct {
		Status       string    `json:"status"`
		TotalResults int       `json:"totalResults"`
		Articles     []Article `json:"articles"`
	}

	Source struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Article struct {
		Source      Source      `json:"source"`
		Author      interface{} `json:"author"`
		Title       string      `json:"title"`
		Description string      `json:"description"`
		URL         string      `json:"url"`
		URLToImage  string      `json:"urlToImage"`
		PublishedAt time.Time   `json:"publishedAt"`
		Content     string      `json:"content"`
	}
)

func New(key string) *NewsAPI {
	return &NewsAPI{
		httpClient: http.DefaultClient,
		key:        key,
	}
}

func (n *NewsAPI) Headlines(country, category, q string, pageSize, page int) (*HeadlinesResponse, error) {
	params := &url.Values{}
	params.Add("country", country)
	params.Add("category", category)
	params.Add("q", q)
	params.Add("pageSize", fmt.Sprintf("%d", pageSize))
	params.Add("page", fmt.Sprintf("%d", page))
	response := &HeadlinesResponse{}
	return response, n.Get(urlHeadlines, params, response)
}

func (n *NewsAPI) Get(url string, params *url.Values, v interface{}) error {
	params.Add("apiKey", n.key)
	url = fmt.Sprintf("%s?%s", url, params.Encode())
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	resp, err := n.httpClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error making GET request to %s. StatusCode: %d", resp.Request.URL, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
