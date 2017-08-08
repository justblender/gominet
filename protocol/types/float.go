package types

import (
	"io"
	"github.com/justblender/gominet/protocol"
)

type Float float32

func (_ Float) Decode(r io.Reader) (interface{}, error) {
	f, err := protocol.ReadFloat32(r)
	return Float(f), err
}

func (f Float) Encode(w io.Writer) error {
	return protocol.WriteFloat32(w, float32(f))
}
