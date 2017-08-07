package types

import (
	"io"
	"github.com/justblender/minecraft/protocol"
)

type Byte byte

func (_ Byte) Decode(r io.Reader) (interface{}, error) {
	b, err := protocol.ReadInt8(r)
	return Byte(b), err
}

func (b Byte) Encode(w io.Writer) error {
	return protocol.WriteInt8(w, int8(b))
}

type UByte uint8

func (_ UByte) Decode(r io.Reader) (interface{}, error) {
	b, err := protocol.ReadUint8(r)
	return Byte(b), err
}

func (b UByte) Encode(w io.Writer) error {
	return protocol.WriteUint8(w, uint8(b))
}

type ByteArray []byte

func (_ ByteArray) Decode(r io.Reader) (interface{}, error) {
	l, err := protocol.ReadVarInt(r)
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
	err := protocol.WriteVarInt(w, len(b))
	if err != nil {
		return err
	}

	_, err = w.Write(b)
	return err
}
