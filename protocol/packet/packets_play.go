package packet

import (
	"github.com/justblender/gominet/chat"
	"github.com/justblender/gominet/protocol/types"
)

type PlayKeepAlive struct {
	AliveId types.Varint
}

func (_ PlayKeepAlive) ID() int { return 0x1F }

type PlayChatMessage struct {
	Chat     chat.TextComponent
	Position types.Byte
}

func (_ PlayChatMessage) ID() int { return 0x0F }

type PlayJoinGame struct {
	EntityId   types.Int
	Gamemode   types.UByte
	Dimension  types.Int
	Difficulty types.UByte
	MaxPlayers types.UByte
	LevelType  types.String
	Debug      types.Bool
}

func (_ PlayJoinGame) ID() int { return 0x23 }

type PlaySpawnPosition struct {
	Location types.Long
}

func (_ PlaySpawnPosition) ID() int { return 0x43 }

type PlayPositionAndLook struct {
	X     types.Double
	Y     types.Double
	Z     types.Double
	Yaw   types.Float
	Pitch types.Float
	Flags types.Byte
	Hueta types.Varint
}

func (_ PlayPositionAndLook) ID() int { return 0x2E }
