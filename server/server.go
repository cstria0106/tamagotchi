package server

import (
	"log"
	"net"
	"sync"
)

type Server struct {
	listener         net.Listener
	connections      map[string]*Connection
	connectionsMutex sync.Mutex
	game             *Game
}

func Create() *Server {
	listener, err := net.Listen("tcp", "0.0.0.0:27775")

	if err != nil {
		log.Fatal(err)
	}

	server := Server{
		listener:         listener,
		connections:      map[string]*Connection{},
		connectionsMutex: sync.Mutex{},
	}

	server.game = CreateGame(&server)

	return &server
}

func (s *Server) Serve() {
	go s.startHandleConnections()
	go s.game.Start()

	<-make(chan interface{})
}

func StartServer() {
	server := Create()
	server.Serve()
}
