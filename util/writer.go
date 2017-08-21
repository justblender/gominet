package util

import (
	"io"
	"math"
	"encoding/binary"
)

func WriteString(writer io.Writer, val string) (err error) {
	bytes := []byte(val)
	err = WriteVarInt(writer, len(bytes))
	if err != nil {
		return
	}

	_, err = writer.Write(bytes)
	return err
}

func WriteVarInt(writer io.Writer, val int) (err error) {
	for val >= 0x80 {
		err = WriteUint8(writer, byte(val)|0x80)
		if err != nil {
			return
		}
		val >>= 7
	}
	err = WriteUint8(writer, byte(val))
	return
}

func WriteBool(writer io.Writer, val bool) (err error) {
	if val {
		err = WriteUint8(writer, 1)
	} else {
		err = WriteUint8(writer, 0)
	}
	return
}

func WriteInt8(writer io.Writer, val int8) (err error) {
	err = WriteUint8(writer, uint8(val))
	return
}

func WriteUint8(writer io.Writer, val uint8) (err error) {
	var protocol [1]byte
	protocol[0] = val
	_, err = writer.Write(protocol[:1])
	return
}

func WriteInt16(writer io.Writer, val int16) (err error) {
	err = WriteUint16(writer, uint16(val))
	return
}

func WriteUint16(writer io.Writer, val uint16) (err error) {
	var protocol [2]byte
	binary.BigEndian.PutUint16(protocol[:2], val)
	_, err = writer.Write(protocol[:2])
	return
}

func WriteInt32(writer io.Writer, val int32) (err error) {
	err = WriteUint32(writer, uint32(val))
	return
}

func WriteUint32(writer io.Writer, val uint32) (err error) {
	var protocol [4]byte
	binary.BigEndian.PutUint32(protocol[:4], val)
	_, err = writer.Write(protocol[:4])
	return
}

func WriteInt64(writer io.Writer, val int64) (err error) {
	err = WriteUint64(writer, uint64(val))
	return
}

func WriteUint64(writer io.Writer, val uint64) (err error) {
	var protocol [8]byte
	binary.BigEndian.PutUint64(protocol[:8], val)
	_, err = writer.Write(protocol[:8])
	return
}

func WriteFloat32(writer io.Writer, val float32) (err error) {
	return WriteUint32(writer, math.Float32bits(val))
}

func WriteFloat64(writer io.Writer, val float64) (err error) {
	return WriteUint64(writer, math.Float64bits(val))
}


