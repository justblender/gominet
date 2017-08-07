package server

import (
	"errors"
	"fmt"
	"net"
	"io"

	"github.com/justblender/minecraft/protocol"
	"github.com/justblender/minecraft/protocol/packet"
)

var NoHandlerException = errors.New("No packet handler has been specified")

type Server struct {
	address 	string
	listener	net.Listener
	handler	 	Handler
}

type Handler func(protocol.Connection, packet.Holder)

func NewServer(host string, port int, handler Handler) *Server {
	return &Server{address: fmt.Sprintf("%s:%d", host, port), handler: handler}
}

func (server *Server) SetHandler(handler Handler) {
	server.handler = handler
}

func (server *Server) Serve() (err error) {
	server.listener, err = net.Listen("tcp", server.address)
	if err != nil {
		return err
	}

	if server.handler == nil {
		return NoHandlerException
	}

	for {
		client, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error occurred while accepting a connection: " + err.Error())
			continue
		}

		fmt.Println("Incoming connection from " + client.RemoteAddr().String())
		go server.handleConnection(protocol.NewConnection(client))
	}
}

func (server *Server) handleConnection(conn *protocol.Connection) {
	defer conn.Close()

	for {
		t, err := conn.Next()
		if err != nil {
			if err == io.EOF {
				continue
			}

			fmt.Println("Error while handling packets: " + err.Error())
			break
		}

		// Run the handler
		server.handler(*conn, t)
	}
}