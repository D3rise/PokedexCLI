package requests

import "encoding/json"

func UnmarshalBody[T any](body []byte) (T, error) {
	var result T

	if err := json.Unmarshal(body, &result); err != nil {
		return result, err
	}

	return result, nil
}
