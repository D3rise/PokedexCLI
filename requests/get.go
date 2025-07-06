package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetAndUnmarshal[T any](url string) (T, error) {
	res, err := http.Get(url)
	if err != nil {
		return *new(T), err
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return *new(T), err
	}

	if res.StatusCode > 299 {
		return *new(T), fmt.Errorf("errornous status code from GET request: %d", res.StatusCode)
	}

	var result T

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
