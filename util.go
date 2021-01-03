package esresponse

import (
	"encoding/json"
	"io"
)

func decode(reader io.Reader, v interface{}) error {
	return json.NewDecoder(reader).Decode(v)
}
