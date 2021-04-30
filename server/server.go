package server

import "net"

type Server struct {
	connections    []*net.Conn
	connectionCount int
}

func CreateServer(maxConnections int) *Server {
	server := Server{
		connections:    make([]*net.Conn, maxConnections),
		connectionCount: 0,
	}

	return &server
}

func (s *Server) Serve() {

}
