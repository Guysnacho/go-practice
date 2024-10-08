package util

import (
	"encoding/json"
	"io"
)

type SongRequest struct {
	name  *string
	query *string
}

func ParseBody(body *io.ReadCloser) (SongRequest, error) {
	decoder := json.NewDecoder(*body)
	t := SongRequest{}
	err := decoder.Decode(&t)
	return t, err
}
