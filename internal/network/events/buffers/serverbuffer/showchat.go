package serverbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
)

func ShowChatBuffer(content string) []byte {
	return append(
		header.Header{Type: events.ShowChat, Length: uint32(len(content))}.Buffer(),
		[]byte(content)...,
	)
}
