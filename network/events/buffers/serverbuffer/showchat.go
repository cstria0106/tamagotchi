package serverbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
)

func ShowChatBuffer(content string) []byte {
	return append(
		header.Header{Type: events.ShowChat, Length: uint32(len(content))}.Buffer(),
		[]byte(content)...,
	)
}
