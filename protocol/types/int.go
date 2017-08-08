package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Int int32

func (_ Int) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadInt32(r)
	return Int(i), err
}

func (i Int) Encode(w io.Writer) error {
	return util.WriteInt32(w, int32(i))
}
