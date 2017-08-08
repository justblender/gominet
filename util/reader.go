package util

import (
	"fmt"
	"io"
	"math"
	"errors"
	"encoding/binary"
	"github.com/satori/go.uuid"
)

func ReadString(reader io.Reader) (val string, err error) {
	length, err := ReadVarInt(reader)
	if err != nil {
		return
	}
	if length < 0 {
		err = errors.New(fmt.Sprintf("Decode, String length is below zero: %d", length))
		return
	}
	if length > 1048576 { // 2^(21-1)
		err = errors.New(fmt.Sprintf("Decode, String length is above maximum: %d", length))
		return
	}
	bytes := make([]byte, length)
	_, err = reader.Read(bytes)
	if err != nil {
		return
	}
	val = string(bytes)
	return
}

func ReadVarInt(reader io.Reader) (result int, err error) {
	var bytes byte = 0
	var b byte

	for {
		b, err = ReadUint8(reader)
		if err != nil {
			return
		}
		result |= int(uint(b&0x7F) << uint(bytes*7))
		bytes++
		if bytes > 5 {
			err = errors.New("Decode, VarInt is too long")
			return
		}
		if (b & 0x80) == 0x80 {
			continue
		}
		break
	}

	return
}

func ReadUUID(reader io.Reader) (result uuid.UUID, err error) {
	bytes := make([]byte, 16)
	_, err = reader.Read(bytes)
	if err != nil {
		return
	}
	result, _ = uuid.FromBytes(bytes)
	return
}

func ReadBool(reader io.Reader) (val bool, err error) {
	uval, err := ReadUint8(reader)
	if err != nil {
		return
	}
	val = uval != 0
	return
}

func ReadInt8(reader io.Reader) (val int8, err error) {
	uval, err := ReadUint8(reader)
	val = int8(uval)
	return
}

func ReadUint8(reader io.Reader) (val uint8, err error) {
	var protocol [1]byte
	_, err = reader.Read(protocol[:1])
	val = protocol[0]
	return
}

func ReadInt16(reader io.Reader) (val int16, err error) {
	uval, err := ReadUint16(reader)
	val = int16(uval)
	return
}

func ReadUint16(reader io.Reader) (val uint16, err error) {
	var protocol [2]byte
	_, err = reader.Read(protocol[:2])
	val = binary.BigEndian.Uint16(protocol[:2])
	return
}

func ReadInt32(reader io.Reader) (val int32, err error) {
	uval, err := ReadUint32(reader)
	val = int32(uval)
	return
}

func ReadUint32(reader io.Reader) (val uint32, err error) {
	var protocol [4]byte
	_, err = reader.Read(protocol[:4])
	val = binary.BigEndian.Uint32(protocol[:4])
	return
}

func ReadInt64(reader io.Reader) (val int64, err error) {
	uval, err := ReadUint64(reader)
	val = int64(uval)
	return
}

func ReadUint64(reader io.Reader) (val uint64, err error) {
	var protocol [8]byte
	_, err = reader.Read(protocol[:8])
	val = binary.BigEndian.Uint64(protocol[:8])
	return
}

func ReadFloat32(reader io.Reader) (val float32, err error) {
	ival, err := ReadUint32(reader)
	val = math.Float32frombits(ival)
	return
}

func ReadFloat64(reader io.Reader) (val float64, err error) {
	ival, err := ReadUint64(reader)
	val = math.Float64frombits(ival)
	return
}
