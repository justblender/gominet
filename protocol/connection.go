package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"
	"sync"
	"github.com/justblender/gominet/protocol/packet"
	"github.com/justblender/gominet/protocol/types"
	"github.com/justblender/gominet/util"
)

var UnknownPacketError = errors.New("unknown packet type")

type Connection struct {
	smu 		sync.RWMutex
	rw  		io.ReadWriteCloser

	State 		State
	Protocol 	uint16
}

func NewConnection(conn net.Conn) *Connection {
	return &Connection{rw: conn}
}

func (c *Connection) Next() (packet.Holder, error) {
	p, err := c.packet()
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

func (c *Connection) packet() (*packet.Packet, error) {
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
	holder := c.getHolderType(p)
	if holder == nil {
		return nil, UnknownPacketError
	}

	inst := reflect.New(holder).Elem()

	for i := 0; i < inst.NumField(); i++ {
		f := inst.Field(i)

		codec, ok := f.Interface().(Codec)
		if !ok {
			continue
		}

		v, err := codec.Decode(&p.Data)
		if err != nil {
			return nil, err
		}

		f.Set(reflect.ValueOf(v))
	}

	return inst.Interface().(packet.Holder), nil
}

func (c *Connection) encode(h packet.Holder) (*bytes.Buffer, error) {
	out := new(bytes.Buffer)
	util.WriteVarInt(out, h.ID())

	v := reflect.ValueOf(h)

	for i := 0; i < v.NumField(); i++ {
		codec, ok := v.Field(i).Interface().(Codec)
		if !ok {
			// XXX(taylor): special-casing
			codec = types.JSON{V: v.Field(i).Interface()}
		}

		if err := codec.Encode(out); err != nil {
			return out, err
		}
	}

	return out, nil
}

func (c *Connection) getHolderType(p *packet.Packet) reflect.Type {
	c.smu.RLock()
	defer c.smu.RUnlock()

	return GetPacket(p.Direction, c.State, p.ID)
}