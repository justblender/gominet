package packet

import "github.com/justblender/gominet/protocol/codecs"

type Handshake struct {
	ProtocolVersion codecs.VarInt
	ServerAddress   codecs.String
	ServerPort      codecs.UnsignedShort
	NextState       codecs.VarInt
}

func (_ Handshake) ID() int { return 0x00 }
