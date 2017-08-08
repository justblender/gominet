package gominet

import (
	"fmt"
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

			fmt.Println("Error while reading packet: " + err.Error())
			break
		}

		err = server.handler(conn, t)
		// TODO: I don't know what I should do here, I'll close connection for now
		if err != nil {
			fmt.Println("Error while handing packet: " + err.Error())
			break
		}
	}
}