package clientbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
)

func PingBuffer() []byte {
	return header.Header{Type: events.Ping}.Buffer()
}
