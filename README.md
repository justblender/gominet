# GoMiNET
Basic Minecraft server library written on Golang and based on Taylor Blau's project over at [ttaylorr/minecraft](https://github.com/ttaylorr/minecraft).

## Installation:
`go get -t github.com/justblender/gominet`

## Creating your own basic server:
```go
package main

import (
	"errors"
	"fmt"
	"reflect"
	"github.com/justblender/gominet"
	"github.com/justblender/gominet/protocol"
	"github.com/justblender/gominet/protocol/packet"
)

func main() {
	server := gominet.NewServer("127.0.0.1", 25565, handlePackets)
	server.ListenAndServe()
}

func handlePackets(conn *protocol.Connection, packet packet.Holder) error {
	switch conn.State {
	case protocol.Handshake:
		handshake, ok := packet.(packet.Handshake)
		if !ok {
			return errors.New(fmt.Sprintf("Expected handshake, received: %s", reflect.TypeOf(packet)))
		}

		conn.Protocol = uint16(handshake.ProtocolVersion)
		conn.State = protocol.State(uint8(handshake.NextState))

	default:
		// Do your own thing here now
		return errors.New("Not implemented yet")
	}

	return nil
}
```