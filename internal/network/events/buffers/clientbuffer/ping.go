package clientbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
)

func PingBuffer() []byte {
	return header.Header{Type: events.Ping}.Buffer()
}
