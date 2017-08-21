package codecs

import (
	"io"
	"errors"
	"encoding/json"
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

type JSON struct {
	V interface{}
}

func (_ JSON) Decode(r io.Reader) (interface{}, error) {
	return nil, errors.New("not yet implemented")
}

func (j JSON) Encode(w io.Writer) error {
	data, err := json.Marshal(j.V)
	if err != nil {
		return err
	}

	str := String(string(data))
	return str.Encode(w)
}

type VarInt int

func (_ VarInt) Decode(r io.Reader) (interface{}, error) {
	v, err := util.ReadVarInt(r)
	return VarInt(v), err
}

func (v VarInt) Encode(w io.Writer) error {
	return util.WriteVarInt(w, int(v))
}

type Boolean bool

func (_ Boolean) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadBool(r)
	return Boolean(l), err
}

func (b Boolean) Encode(w io.Writer) error {
	return util.WriteBool(w, bool(b))
}

type Byte byte

func (_ Byte) Decode(r io.Reader) (interface{}, error) {
	b, err := util.ReadInt8(r)
	return Byte(b), err
}

func (b Byte) Encode(w io.Writer) error {
	return util.WriteInt8(w, int8(b))
}

type UnsignedByte uint8

func (_ UnsignedByte) Decode(r io.Reader) (interface{}, error) {
	b, err := util.ReadUint8(r)
	return UnsignedByte(b), err
}

func (b UnsignedByte) Encode(w io.Writer) error {
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

type Short int16

func (_ Short) Decode(r io.Reader) (interface{}, error) {
	s, err := util.ReadInt16(r)
	return Short(s), err
}

func (s Short) Encode(w io.Writer) error {
	return util.WriteInt16(w, int16(s))
}

type UnsignedShort uint16

func (_ UnsignedShort) Decode(r io.Reader) (interface{}, error) {
	s, err := util.ReadUint16(r)
	return UnsignedShort(s), err
}

func (s UnsignedShort) Encode(w io.Writer) error {
	return util.WriteUint16(w, uint16(s))
}

type Int int32

func (_ Int) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadInt32(r)
	return Int(i), err
}

func (i Int) Encode(w io.Writer) error {
	return util.WriteInt32(w, int32(i))
}

type UnsignedInt int32

func (_ UnsignedInt) Decode(r io.Reader) (interface{}, error) {
	i, err := util.ReadUint32(r)
	return UnsignedInt(i), err
}

func (i UnsignedInt) Encode(w io.Writer) error {
	return util.WriteUint32(w, uint32(i))
}

type Long int64

func (_ Long) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadInt64(r)
	return Long(l), err
}

func (l Long) Encode(w io.Writer) error {
	return util.WriteInt64(w, int64(l))
}

type UnsignedLong uint64

func (_ UnsignedLong) Decode(r io.Reader) (interface{}, error) {
	l, err := util.ReadUint64(r)
	return UnsignedLong(l), err
}

func (l UnsignedLong) Encode(w io.Writer) error {
	return util.WriteUint64(w, uint64(l))
}

type Float float32

func (_ Float) Decode(r io.Reader) (interface{}, error) {
	f, err := util.ReadInt64(r)
	return Float(f), err
}

func (f Float) Encode(w io.Writer) error {
	return util.WriteFloat32(w, float32(f))
}

type Double float64

func (_ Double) Decode(r io.Reader) (interface{}, error) {
	f, err := util.ReadFloat64(r)
	return Double(f), err
}

func (d Double) Encode(w io.Writer) error {
	return util.WriteFloat64(w, float64(d))
}
