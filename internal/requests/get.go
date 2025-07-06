package requests

import (
	"fmt"
	"io"
	"net/http"
)

func Get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	err = res.Body.Close()
	if err != nil {
		return nil, err
	}

	if res.StatusCode > 299 {
		return []byte{}, fmt.Errorf("errornous status code from GET request: %d", res.StatusCode)
	}

	return body, nil
}
