package command

import (
	"encoding/json"
	"net/http"
)

func GetJSON(url string, obj interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if err := json.NewDecoder(response.Body).Decode(obj); err != nil {
		return err
	}

	return nil
}
