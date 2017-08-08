package packet

import "bytes"

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
