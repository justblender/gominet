package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Varint int

func (_ Varint) Decode(r io.Reader) (interface{}, error) {
	v, err := util.ReadVarInt(r)
	return Varint(v), err
}

func (v Varint) Encode(w io.Writer) error {
	return util.WriteVarInt(w, int(v))
}
