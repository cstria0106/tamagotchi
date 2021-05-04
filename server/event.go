package server

import (
	"log"
	"tamagotchi/network/events"
	"tamagotchi/network/events/buffers/serverbuffer"
	"tamagotchi/server/gamestate"
	"tamagotchi/util"
)

func (s *Server) startHandleEvents() {
	for {
		s.handleEvents()
	}
}

func (s *Server) handleEvents() {
	s.userConnectionsMutex.Lock()
	for _, connection := range s.userConnections {
		select {
		case event := <-connection.EventChannel:
			s.handleEvent(connection, event)
		default:
			continue
		}
	}
	s.userConnectionsMutex.Unlock()
}

func (s *Server) handleEvent(connection *UserConnection, event events.Event) {
	log.Println("got", event.Type.String())

	switch event.Type {
	case events.Ping:
		s.sendTo(connection, serverbuffer.PongBuffer())
	case events.Feed:
		s.game.addFood(gamestate.Point{
			X: util.DecodeU16(event.Payload[0:2]),
			Y: util.DecodeU16(event.Payload[2:4]),
		})
	case events.Clean:
		s.game.clean(util.DecodeU32(event.Payload))
	}
}
