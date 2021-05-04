package server

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	connection           *net.UDPConn
	userConnections      map[string]*UserConnection
	userConnectionsMutex sync.Mutex
	game                 *Game
}

func Create() *Server {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:27775")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := Server{
		connection:           conn,
		userConnections:      map[string]*UserConnection{},
		userConnectionsMutex: sync.Mutex{},
	}

	server.game = CreateGame(&server)

	return &server
}

func (s *Server) Serve() {
	go s.startHandlePackets()
	go s.startHandleEvents()
	go s.game.Start()

	<-make(chan interface{})
}

func StartServer() {
	server := Create()
	server.Serve()
}
