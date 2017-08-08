package types

import (
	"io"
	"github.com/justblender/gominet/protocol"
)

type Double float64

func (_ Double) Decode(r io.Reader) (interface{}, error) {
	f, err := protocol.ReadFloat64(r)
	return Long(f), err
}

func (d Double) Encode(w io.Writer) error {
	return protocol.WriteFloat64(w, float64(d))
}
