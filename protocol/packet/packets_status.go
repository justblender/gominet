package packet

import (
	"github.com/justblender/gominet/chat"
	"github.com/justblender/gominet/protocol/types"
)

type StatusRequest struct{}

func (_ StatusRequest) ID() int { return 0x00 }

type StatusResponse struct {
	Status struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`

		Players struct {
			Max    int `json:"max"`
			Online int `json:"online"`
		} `json:"players"`

		Description chat.TextComponent `json:"description"`
	}
}

func (_ StatusResponse) ID() int { return 0x00 }

type StatusPing struct {
	Payload types.Long
}

func (_ StatusPing) ID() int { return 0x01 }

type StatusPong struct {
	Payload types.Long
}

func (_ StatusPong) ID() int { return 0x01 }
