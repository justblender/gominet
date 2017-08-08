package protocol

import (
	"reflect"
	"github.com/justblender/gominet/protocol/packet"
)

var Packets = map[packet.Direction]map[State]map[int]reflect.Type{
	packet.Serverbound: {
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

	packet.Clientbound: {
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

func GetPacket(d packet.Direction, s State, id int) reflect.Type {
	return Packets[d][s][id]
}
