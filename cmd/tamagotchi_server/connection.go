package main

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"net"
)

type Connection struct {
	Conn      net.Conn
	EventChan chan *events.Event
}

func (s *Server) getConnection(addr net.Addr) *Connection {
	s.connectionsMutex.Lock()

	if connection, ok := s.connections[addr.String()]; ok {
		s.connectionsMutex.Unlock()
		return connection
	}

	s.connectionsMutex.Unlock()
	return nil
}

func (s *Server) removeConnection(addr net.Addr) {
	connection := s.getConnection(addr)
	close(connection.EventChan)

	s.connectionsMutex.Lock()
	delete(s.connections, addr.String())
	s.connectionsMutex.Unlock()
}

func (s *Server) addConnectionFromConn(conn net.Conn) *Connection {
	connection := &Connection{
		Conn:      conn,
		EventChan: make(chan *events.Event, 10),
	}

	s.connectionsMutex.Lock()
	s.connections[conn.RemoteAddr().String()] = connection
	s.connectionsMutex.Unlock()

	return connection
}
