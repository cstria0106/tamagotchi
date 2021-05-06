package main

import (
	"log"
	"tamagotchi/cmd/tamagotchi_server/gamestate"
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/events/buffers/serverbuffer"
	"tamagotchi/internal/util"
)

func (s *Server) startHandleEvents(c *Connection) {
	for e := range c.EventChan {
		err := s.handleEvent(c, e)

		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) handleEvent(c *Connection, event *events.Event) error {
	log.Println("Got", event.Type.String(), "from", c.Conn.RemoteAddr().String())

	switch event.Type {
	case events.Ping:
		s.sendTo(c, serverbuffer.PongBuffer(s.version))
	case events.Feed:
		s.game.addFood(gamestate.Point{
			X: util.DecodeU16(event.Payload[0:2]),
			Y: util.DecodeU16(event.Payload[2:4]),
		})
	case events.Clean:
		s.game.clean(util.DecodeU32(event.Payload))
	}

	return nil
}
