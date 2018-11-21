package pokequest

import (
	"fmt"
	"net/http"
)

type (
	httpCache struct {
		cache  map[string]*http.Response
		client *http.Client
	}
)

func NewHTTPCache(client *http.Client) *httpCache {
	return &httpCache{
		cache:  map[string]*http.Response{},
		client: client,
	}
}

func (h *httpCache) Do(req *http.Request) (*http.Response, error) {
	resp, ok := h.cache[req.URL.String()]
	if !ok {
		resp, err := h.client.Do(req)
		if err != nil {
			return nil, err
		}
		h.cache[req.URL.String()] = resp
		fmt.Println("CACHED!!")
		return resp, nil
	}

	return resp, nil
}
