package types

import (
	"io"
	"github.com/justblender/minecraft/protocol"
)

type Varint int

func (_ Varint) Decode(r io.Reader) (interface{}, error) {
	v, err := protocol.ReadVarInt(r)
	return Varint(v), err
}

func (v Varint) Encode(w io.Writer) error {
	return protocol.WriteVarInt(w, int(v))
}
