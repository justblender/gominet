package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Bool bool

func (_ Bool) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadBool(r)
	return Bool(l), err
}

func (b Bool) Encode(w io.Writer) error {
	return util.WriteBool(w, bool(b))
}
