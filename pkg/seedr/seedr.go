package seedr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
	req, err := http.NewRequest("GET", baseURL+endpoint, nil)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(s.username, s.password)
	redirectedURL := ""

	s.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		redirectedURL = req.URL.String()
		return errRedirectNotAllowed
	}

	_, err = s.httpClient.Do(req)
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

type DownloadResponse struct {
	Result        bool   `json:"result"`
	Code          int    `json:"code,omitempty"`
	UserTorrentID int    `json:"user_torrent_id,omitempty"`
	Title         string `json:"title,omitempty"`
	TorrentHash   string `json:"torrent_hash,omitempty"`
	Error         string `json:"error,omitempty"`
}

func (s *Seedr) Download(u string) (*DownloadResponse, error) {
	result := &DownloadResponse{}
	values := &url.Values{}
	values.Add("url", u)
	return result, s.Post("torrent/url", values, result)
}

func (s *Seedr) Post(endpoint string, values *url.Values, v interface{}) error {
	req, err := http.NewRequest("POST", baseURL+endpoint, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(s.username, s.password)

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error making POST request. Status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
