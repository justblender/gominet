package types

import (
	"io"
	"github.com/justblender/gominet/util"
)

type Byte byte

func (_ Byte) Decode(r io.Reader) (interface{}, error) {
	b, err := util.ReadInt8(r)
	return Byte(b), err
}

func (b Byte) Encode(w io.Writer) error {
	return util.WriteInt8(w, int8(b))
}

type UByte uint8

func (_ UByte) Decode(r io.Reader) (interface{}, error) {
	b, err := util.ReadUint8(r)
	return Byte(b), err
}

func (b UByte) Encode(w io.Writer) error {
	return util.WriteUint8(w, uint8(b))
}

type ByteArray []byte

func (_ ByteArray) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadVarInt(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, int(l))
	_, err = io.ReadFull(r, buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (b ByteArray) Encode(w io.Writer) error {
	err := util.WriteVarInt(w, len(b))
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}
