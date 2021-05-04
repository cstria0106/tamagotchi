package clientbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
)

func CleanBuffer() []byte {
	return header.Header{Type: events.Clean}.Buffer()
}
