package server

import (
	"net"
	"tamagotchi/network/events"
)

type UserConnection struct {
	Addr         net.Addr
	EventChannel chan events.Event
}
