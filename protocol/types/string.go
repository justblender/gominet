package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type String string

func (_ String) Decode(r io.Reader) (interface{}, error) {
	s, err := util.ReadString(r)
	return String(s), err
}

func (s String) Encode(w io.Writer) error {
	return util.WriteString(w, string(s))
}
