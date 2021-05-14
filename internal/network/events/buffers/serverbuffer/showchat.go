package serverbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
)

func ShowChatBuffer(content string) []byte {
	return append(
		header.Header{Type: events.ShowChat, Length: uint32(len(content))}.Buffer(),
		[]byte(content)...,
	)
}
