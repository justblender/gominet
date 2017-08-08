package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Short int16

func (_ Short) Decode(r io.Reader) (interface{}, error) {
	s, err := util.ReadInt16(r)
	return Short(s), err
}

func (s Short) Encode(w io.Writer) error {
	return util.WriteInt16(w, int16(s))
}

type UShort uint16

func (_ UShort) Decode(r io.Reader) (interface{}, error) {
	s, err := util.ReadUint16(r)
	return Short(s), err
}

func (s UShort) Encode(w io.Writer) error {
	return util.WriteUint16(w, uint16(s))
}
