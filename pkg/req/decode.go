package req

import (
	"encoding/json"
	"io"
)

func Decode[T any](body io.Reader) (T, error) {
	var t T
	err := json.NewDecoder(body).Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}
