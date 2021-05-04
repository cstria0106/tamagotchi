package clientbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
)

func PingBuffer() []byte {
	return header.Header{Type: events.Ping}.Buffer()
}
