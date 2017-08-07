package types

import (
	"io"
	"github.com/justblender/minecraft/protocol"
)

type String string

func (_ String) Decode(r io.Reader) (interface{}, error) {
	s, err := protocol.ReadString(r)
	return String(s), err
}

func (s String) Encode(w io.Writer) error {
	return protocol.WriteString(w, string(s))
}
