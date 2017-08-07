package types

import (
	"io"
	"github.com/justblender/minecraft/protocol"
)

type Bool bool

func (_ Bool) Decode(r io.Reader) (interface{}, error) {
	l, err := protocol.ReadBool(r)
	return Bool(l), err
}

func (b Bool) Encode(w io.Writer) error {
	return protocol.WriteBool(w, bool(b))
}
