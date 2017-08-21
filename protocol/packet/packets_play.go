package packet

import (
	"github.com/justblender/gominet/chat"
	"github.com/justblender/gominet/protocol/codecs"
)

type PlayKeepAlive struct {
	AliveId 	codecs.VarInt
}

func (_ PlayKeepAlive) ID() int { return 0x1F }

type PlayChatMessage struct {
	Chat     	chat.TextComponent
	Position 	codecs.Byte
}

func (_ PlayChatMessage) ID() int { return 0x0F }

type PlayJoinGame struct {
	EntityId   	codecs.Int
	Gamemode   	codecs.UnsignedByte
	Dimension  	codecs.Int
	Difficulty 	codecs.UnsignedByte
	MaxPlayers 	codecs.UnsignedByte
	LevelType  	codecs.String
	Debug      	codecs.Boolean
}

func (_ PlayJoinGame) ID() int { return 0x23 }

type PlaySpawnPosition struct {
	Location	codecs.Long
}

func (_ PlaySpawnPosition) ID() int { return 0x43 }

type PlayPositionAndLook struct {
	X     		codecs.Double
	Y     		codecs.Double
	Z     		codecs.Double
	Yaw   		codecs.Float
	Pitch 		codecs.Float
	Flags 		codecs.Byte
	Data 		codecs.VarInt
}

func (_ PlayPositionAndLook) ID() int { return 0x2E }
