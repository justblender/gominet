package types

import (
	"encoding/json"
	"errors"
	"io"
)

type JSON struct {
	V interface{}
}

func (_ JSON) Decode(r io.Reader) (interface{}, error) {
	return nil, errors.New("not yet implemented")
}

func (j JSON) Encode(w io.Writer) error {
	data, err := json.Marshal(j.V)
	if err != nil {
		return err
	}

	str := String(string(data))
	return str.Encode(w)
}
