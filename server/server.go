package server

import (
	"encoding/binary"
	"log"
	"net"
	"tamagotchi/network"
)

type UserConnection struct {
	Addr net.Addr
}

type Server struct {
	connection *net.UDPConn
}

type Payload struct {
	network.Payload
	addr net.Addr
}

func CreateServer() *Server {
	addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:27775")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	server := Server{
		connection: conn,
	}
	return &server
}

func (s *Server) Serve() {
	for {
		s.handleConnection()
	}
}

func (s *Server) handleConnection() {
	buffer := make([]byte, 6)

	_, addr, err := s.connection.ReadFrom(buffer)
	if err != nil {
		log.Println(err)
		return
	}

	action := network.Action(binary.BigEndian.Uint16(buffer[0:2]))
	length := binary.BigEndian.Uint32(buffer[2:6])

	payload := Payload{
		network.Payload{
			Action:     action,
			DataLength: length,
			Data:       nil,
		},
		addr,
	}

	if length != 0 {
		buffer = make([]byte, length)
		receivedLength, _, err := s.connection.ReadFrom(buffer)

		if err != nil {
			log.Println(err)
			goto handle
		}

		if uint32(receivedLength) != length {
			log.Printf("received length(%d) is mismatching with metadata length(%d)\n", receivedLength, length)
			goto handle
		}

		payload.Data = buffer[:length]
	}

handle:
	s.HandlePayload(&payload)
}

func (s *Server) HandlePayload(payload *Payload) {
	log.Printf("%s #%d\n", payload.Action.ToString(), payload.DataLength)
	if payload.DataLength > 0 {
		log.Printf("%s\n", string(payload.Data))
	}

	if payload.Action == network.ActionPing {
		_, err := s.connection.WriteTo(network.BuildPayload(network.ActionPong, nil).ToBuffer(), payload.addr)
		if err != nil {
			log.Println("failed to send:", err)
		}
	}
}

func StartServer() {
	server := CreateServer()
	server.Serve()
}
