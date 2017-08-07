package packet

import (
	"github.com/justblender/minecraft/chat"
	"github.com/justblender/minecraft/protocol/types"
)

type LoginStart struct {
	Username types.String
}

func (_ LoginStart) ID() int { return 0x00 }

type LoginSuccess struct {
	UUID     types.String
	Username types.String
}

func (_ LoginSuccess) ID() int { return 0x02 }

type LoginDisconnect struct {
	Chat chat.TextComponent
}

func (_ LoginDisconnect) ID() int { return 0x00 }
