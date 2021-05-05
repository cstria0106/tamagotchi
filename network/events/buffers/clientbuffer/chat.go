package clientbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
)

func ChatBuffer(content string) []byte {
	return append(
		header.Header{Type: events.Chat, Length: uint32(len(content))}.Buffer(),
		[]byte(content)...,
	)
}
