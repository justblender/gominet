package protocol

type State uint8

const (
	Handshake State = iota
	Status
	Login
	Play
)
