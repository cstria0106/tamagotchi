package clientbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
)

func ChatBuffer(content string) []byte {
	return append(
		header.Header{Type: events.Chat, Length: uint32(len(content))}.Buffer(),
		[]byte(content)...,
	)
}
