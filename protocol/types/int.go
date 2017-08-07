package types

import (
	"io"
	"github.com/justblender/minecraft/protocol"
)

type Int int32

func (_ Int) Decode(r io.Reader) (interface{}, error) {
	i, err := protocol.ReadInt32(r)
	return Int(i), err
}

func (i Int) Encode(w io.Writer) error {
	return protocol.WriteInt32(w, int32(i))
}
