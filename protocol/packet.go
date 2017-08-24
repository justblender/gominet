package protocol

import (
	"bytes"
	"errors"
	"reflect"
	"github.com/justblender/gominet/protocol/packet"
)

type Packet struct {
	ID 			int
	Direction 	Direction
	Data 		bytes.Buffer
}

type Direction int

const (
	Serverbound Direction = iota
	Clientbound
)

var (
	UnknownPacketType = errors.New("unknown packet type")
	InvalidPacketLength = errors.New("received packet is below zero or above maximum size")
)

var packets = map[Direction]map[State]map[int]reflect.Type{
	Serverbound: {
		Handshake: {
			0x00: reflect.TypeOf(packet.Handshake{}),
		},
		Status: {
			0x00: reflect.TypeOf(packet.StatusRequest{}),
			0x01: reflect.TypeOf(packet.StatusPing{}),
		},
		Login: {
			0x00: reflect.TypeOf(packet.LoginStart{}),
		},
	},

	Clientbound: {
		Status: {
			0x00: reflect.TypeOf(packet.StatusResponse{}),
			0x01: reflect.TypeOf(packet.StatusPong{}),
		},
		Login: {
			0x00: reflect.TypeOf(packet.LoginDisconnect{}),
			0x02: reflect.TypeOf(packet.LoginSuccess{}),
		},
		Play: {
			0x1F: reflect.TypeOf(packet.PlayKeepAlive{}),
			0x0F: reflect.TypeOf(packet.PlayChatMessage{}),
			0x23: reflect.TypeOf(packet.PlayJoinGame{}),
			0x43: reflect.TypeOf(packet.PlaySpawnPosition{}),
			0x2E: reflect.TypeOf(packet.PlayPositionAndLook{}),
		},
	},
}