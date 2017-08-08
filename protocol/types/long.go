package types

import (
	"io"
	"github.com/justblender/gominet/protocol"
)

type Long int64

func (_ Long) Decode(r io.Reader) (interface{}, error) {
	l, err := protocol.ReadInt64(r)
	return Long(l), err
}

func (l Long) Encode(w io.Writer) error {
	return protocol.WriteInt64(w, int64(l))
}

type ULong uint64

func (_ ULong) Decode(r io.Reader) (interface{}, error) {
	l, err := protocol.ReadUint64(r)
	return Long(l), err
}

func (l ULong) Encode(w io.Writer) error {
	return protocol.WriteUint64(w, uint64(l))
}
