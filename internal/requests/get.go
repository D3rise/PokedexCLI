package requests

import (
	"io"
	"net/http"
)

func Get(url string) (res *http.Response, body []byte, err error) {
	res, err = http.Get(url)
	if err != nil {
		return nil, nil, err
	}

	body, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	err = res.Body.Close()
	if err != nil {
		return nil, nil, err
	}

	return res, body, nil
}
