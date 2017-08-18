package gominet

import (
	"fmt"
	"log"
	"io"
	"net"
	"errors"
	"github.com/justblender/gominet/protocol/packet"
	"github.com/justblender/gominet/protocol"
)

var NoHandlerException = errors.New("No packet handler has been specified")

type Server struct {
	host 		string
	port 		int

	listener	net.Listener
	handler	 	Handler
}

type Handler func(*protocol.Connection, packet.Holder) error

func NewServer(host string, port int, handler Handler) *Server {
	return &Server{host: host, port: port, handler: handler}
}

func (server *Server) SetHandler(handler Handler) {
	server.handler = handler
}

func (server *Server) ListenAndServe() (err error) {
	if server.handler == nil {
		return NoHandlerException
	}

	server.listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", server.host, server.port))
	if err != nil {
		return
	}

	for {
		client, err := server.listener.Accept()
		if err != nil {
			log.Printf("Error occurred while accepting a connection: %v\n", err)
			continue
		}

		log.Printf("Incoming connection from %s\n", client.RemoteAddr().String())
		go server.handleConnection(protocol.NewConnection(client))
	}
}

func (server *Server) handleConnection(conn *protocol.Connection) {
	defer conn.Close()

	for {
		holder, err := conn.Next()
		if err != nil {
			if err == io.EOF || err == protocol.UnknownPacketType {
				continue
			}

			log.Printf("Error occurred while reading packet: %v\n", err)
			break
		}

		if err = server.handler(conn, holder); err != nil {
			log.Printf("Error occurred while handling packet: %v\n", err)
			break
		}
	}
}