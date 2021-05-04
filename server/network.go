package server

import (
	"log"
	"tamagotchi/network/events"
	"tamagotchi/network/header"
	"tamagotchi/util"
)

func (s *Server) startHandlePackets() {
	for {
		s.handlePackets()
	}
}

func (s *Server) handlePackets() {
	buffer := make([]byte, 6)

	_, addr, err := s.connection.ReadFrom(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	receivedHeader := header.FromBuffer(buffer[0:6])

	if receivedHeader.Length != 0 {
		buffer = make([]byte, receivedHeader.Length)
		receivedLength, _, err := s.connection.ReadFrom(buffer)

		if err != nil {
			log.Println(err)
			goto handle
		}

		if uint32(receivedLength) != receivedHeader.Length {
			log.Printf("received length(%d) is mismatching with metadata length(%d)\n", receivedLength, receivedHeader.Length)
			goto handle
		}
	}

handle:
	s.userConnectionsMutex.Lock()

	key := addr.String()
	if connection, ok := s.userConnections[key]; !ok {
		connection = &UserConnection{
			addr,
			make(chan events.Event, 10),
		}

		s.userConnections[key] = connection
	}

	s.userConnections[key].EventChannel <- events.Event{
		Type:    receivedHeader.Type,
		Payload: buffer,
	}

	s.userConnectionsMutex.Unlock()
}

func (s *Server) sendToAll(buffer []byte) {
	log.Println("Sent", header.FromBuffer(buffer[:6]).Type.String(), "to all")

	for _, connection := range s.userConnections {
		s.sendTo(connection, buffer)
	}
}

func (s *Server) sendTo(connection *UserConnection, buffer []byte) {
	_, err := s.connection.WriteTo(buffer[:6], connection.Addr)
	if err != nil {
		log.Println(err)
	}

	if util.DecodeU32(buffer[2:6]) != 0 {
		_, err = s.connection.WriteTo(buffer[6:], connection.Addr)
		if err != nil {
			log.Println(err)
		}
	}
}
