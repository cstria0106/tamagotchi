package serverbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
)

func PongBuffer() []byte {
	return header.Header{Type: events.Pong}.Buffer()
}
