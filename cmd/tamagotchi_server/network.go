package main

import (
	"fmt"
	"io"
	"log"
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
)

func (s *Server) startHandleConnections() {
	for {
		err := s.handleConnection()

		if err != nil {
			log.Println(err)
		}
	}
}

func (s *Server) handleConnection() error {
	conn, err := s.listener.Accept()
	if err != nil {
		return err
	}

	addr := conn.RemoteAddr()
	connection := s.getConnection(addr)

	if connection == nil {
		connection = s.addConnectionFromConn(conn)
	}

	go s.startHandlePackets(connection)
	go s.startHandleEvents(connection)
	return nil
}

func (s *Server) startHandlePackets(connection *Connection) {
	for {
		err := s.handlePacket(connection)

		if err != nil {
			if err == io.EOF {
				s.removeConnection(connection.Conn.RemoteAddr())
			} else {
				log.Println(err)
			}
			break
		}
	}
}

func (s *Server) handlePacket(connection *Connection) error {
	headerBuffer := [6]byte{}
	_, err := connection.Conn.Read(headerBuffer[:])

	if err != nil {
		return err
	}

	receivedHeader := header.FromBuffer(headerBuffer[:])

	if !receivedHeader.Type.ValidatePayloadLength(receivedHeader.Length) {
		return fmt.Errorf("received metadata length(%d) is not valid for type(%s)",
			receivedHeader.Length,
			receivedHeader.Type.String(),
		)
	}

	payloadBuffer := make([]byte, receivedHeader.Length)

	if receivedHeader.Length > 0 {
		length, err := connection.Conn.Read(payloadBuffer)

		if err != nil {
			return err
		}

		if uint32(length) != receivedHeader.Length {
			return fmt.Errorf("received payload buffer length(%d) is mismatching with metadata length(%d)\n", length, receivedHeader.Length)
		}
	}

	event := &events.Event{
		Type:    receivedHeader.Type,
		Payload: payloadBuffer,
	}

	connection.EventChan <- event
	return nil
}

func (s *Server) sendToAll(buffer []byte) {
	log.Println("Sent", header.FromBuffer(buffer[:6]).Type.String(), "to all")

	for _, connection := range s.connections {
		s.sendTo(connection, buffer)
	}
}

func (s *Server) sendTo(connection *Connection, buffer []byte) {
	log.Println("Sent", header.FromBuffer(buffer[:6]).Type.String(), "to", connection.Conn.RemoteAddr().String())

	_, err := connection.Conn.Write(buffer[0:6])
	if err != nil {
		log.Println(err)
	}

	if util.DecodeU32(buffer[2:6]) != 0 {
		_, err := connection.Conn.Write(buffer[6:])
		if err != nil {
			log.Println(err)
		}
	}
}
