package protocol

import (
	"io"
	"net"
	"fmt"
	"bytes"
	"reflect"
	"errors"
	"github.com/justblender/gominet/util"
	"github.com/justblender/gominet/protocol/packet"
	"github.com/justblender/gominet/protocol/codecs"
)

type State uint8

const (
	Handshake State = iota
	Status
	Login
	Play
)

type Connection struct {
	rw  		io.ReadWriteCloser

	State 		State
	Protocol 	uint16
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{rw: conn}
}

func (c *Connection) Next() (packet.Holder, error) {
	p, err := c.read()
	if err != nil {
		return nil, err
	}

	return c.decode(p)
}

func (c *Connection) Write(h packet.Holder) (int, error) {
	data, err := c.encode(h)
	if err != nil {
		return -1, err
	}

	err = util.WriteVarInt(c.rw, data.Len())
	if err != nil {
		return -1, err
	}

	n, err := data.WriteTo(c.rw)
	if err != nil {
		return int(n), err
	}

	return int(n), nil
}

func (c *Connection) Close() error {
	if closer, ok := c.rw.(io.Closer); ok {
		return closer.Close()
	}

	return nil
}

func (c *Connection) read() (*packet.Packet, error) {
	length, err := util.ReadVarInt(c.rw)
	if err != nil {
		return nil, err
	}

	if length < 0 {
		err = errors.New(fmt.Sprintf("Decode, Packet length is below zero: %d", length))
		return nil, err
	}

	if length > 1048576 { // 2^(21-1)
		err = errors.New(fmt.Sprintf("Decode, Packet length is above maximum: %d", length))
		return nil, err
	}

	payload := make([]byte, length)
	_, err = io.ReadFull(c.rw, payload)

	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(payload)
	id, err := util.ReadVarInt(buffer)

	if err != nil {
		return nil, err
	}

	return &packet.Packet{
		ID:        id,
		Direction: packet.Serverbound,
		Data:      *buffer,
	}, nil
}

func (c *Connection) decode(p *packet.Packet) (packet.Holder, error) {
	holder := GetPacket(p.Direction, c.State, p.ID)
	if holder == nil {
		return nil, UnknownPacketType
	}

	inst := reflect.New(holder).Elem()

	for i := 0; i < inst.NumField(); i++ {
		field := inst.Field(i)

		codec, ok := field.Interface().(codecs.Codec)
		if !ok {
			return nil, codecs.UnknownCodecType
		}

		value, err := codec.Decode(&p.Data)
		if err != nil {
			return nil, err
		}

		field.Set(reflect.ValueOf(value))
	}

	return inst.Interface().(packet.Holder), nil
}

func (c *Connection) encode(h packet.Holder) (*bytes.Buffer, error) {
	buffer := new(bytes.Buffer)
	util.WriteVarInt(buffer, h.ID())

	value := reflect.ValueOf(h)

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)

		codec, ok := field.Interface().(codecs.Codec)
		if !ok {
			if field.Kind() == reflect.Struct {
				codec = codecs.JSON{V: value.Field(i).Interface()}
			} else {
				return nil, codecs.UnknownCodecType
			}
		}

		if err := codec.Encode(buffer); err != nil {
			return nil, err
		}
	}

	return buffer, nil
}