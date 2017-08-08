package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Double float64

func (_ Double) Decode(r io.Reader) (interface{}, error) {
	f, err := util.ReadFloat64(r)
	return Long(f), err
}

func (d Double) Encode(w io.Writer) error {
	return util.WriteFloat64(w, float64(d))
}
