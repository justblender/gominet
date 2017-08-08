package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Float float32

func (_ Float) Decode(r io.Reader) (interface{}, error) {
	f, err := util.ReadInt64(r)
	return Float(f), err
}

func (f Float) Encode(w io.Writer) error {
	return util.WriteFloat32(w, float32(f))
}
