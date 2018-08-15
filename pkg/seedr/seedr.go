package seedr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	baseURL = "https://www.seedr.cc/rest/"
)

var (
	errRedirectNotAllowed = errors.New("redirect")
)

type (
	Seedr struct {
		username, password string
		httpClient         *http.Client
	}

	HTTPRequester interface {
		Do(*http.Request) (*http.Response, error)
	}
)

func New(username, password string) *Seedr {
	return &Seedr{
		username:   username,
		password:   password,
		httpClient: &http.Client{},
	}
}

func (s *Seedr) SetHTTPClient(httpClient *http.Client) {
	s.httpClient = httpClient
}

type FoldersResponse struct {
	SpaceMax  int64         `json:"space_max"`
	SpaceUsed int64         `json:"space_used"`
	Code      int           `json:"code"`
	Timestamp string        `json:"timestamp"`
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	ParentID  int           `json:"parent_id"`
	Torrents  []interface{} `json:"torrents"`
	Folders   []struct {
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Size       int64  `json:"size"`
		LastUpdate string `json:"last_update"`
	} `json:"folders"`
	Files  []interface{} `json:"files"`
	Result bool          `json:"result"`
}

func (s *Seedr) Folders() (*FoldersResponse, error) {
	result := &FoldersResponse{}
	return result, s.Get("folder", result)
}

type FolderResponse struct {
	SpaceMax  int64         `json:"space_max"`
	SpaceUsed int64         `json:"space_used"`
	Code      int           `json:"code"`
	Timestamp string        `json:"timestamp"`
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	ParentID  int           `json:"parent_id"`
	Torrents  []interface{} `json:"torrents"`
	Folders   []interface{} `json:"folders"`
	Files     []struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Size           int    `json:"size"`
		Hash           string `json:"hash"`
		LastUpdate     string `json:"last_update"`
		StreamAudio    bool   `json:"stream_audio"`
		StreamVideo    bool   `json:"stream_video"`
		VideoConverted string `json:"video_converted,omitempty"`
	} `json:"files"`
	Result bool `json:"result"`
}

func (s *Seedr) Folder(id int) (*FolderResponse, error) {
	result := &FolderResponse{}
	endpoint := fmt.Sprintf("folder/%d", id)
	return result, s.Get(endpoint, result)
}

func (s *Seedr) HLS(id int) (string, error) {
	endpoint := fmt.Sprintf("file/%d/hls", id)
	return s.FinalURL(endpoint)
}

func (s *Seedr) FinalURL(endpoint string) (string, error) {
	redirectedURL := ""

	s.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectedURL = req.URL.String()
		return errRedirectNotAllowed
	}

	_, err := s.httpClient.Get(baseURL + endpoint)
	if urlError, ok := err.(*url.Error); !ok || urlError.Err != errRedirectNotAllowed {
		return "", err
	}
	return redirectedURL, nil
}

func (s *Seedr) Get(endpoint string, v interface{}) error {
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(s.username, s.password)

	s.httpClient.CheckRedirect = nil
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
