package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"tamagotchi/internal/data/version"
	"tamagotchi/internal/util/versionutil"
)

type Server struct {
	version          *version.Version
	listener         net.Listener
	connections      map[string]*Connection
	connectionsMutex sync.Mutex
	game             *Game
}

func Create(port uint16) *Server {
	ver, err := versionutil.GetLocalVersion()
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		log.Fatal(err)
	}

	server := Server{
		version:          ver,
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

func StartServer(port uint16) {
	server := Create(port)
	server.Serve()
}
