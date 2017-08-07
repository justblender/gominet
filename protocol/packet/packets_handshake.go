package packet

import "github.com/justblender/minecraft/protocol/types"

type Handshake struct {
	ProtocolVersion types.Varint
	ServerAddress   types.String
	ServerPort      types.Short
	NextState       types.Varint
}

func (_ Handshake) ID() int { return 0x00 }
