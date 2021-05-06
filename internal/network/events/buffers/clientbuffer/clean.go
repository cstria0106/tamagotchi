package clientbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
)

func CleanBuffer(id uint32) []byte {
	return append(
		header.Header{Type: events.Clean, Length: 4}.Buffer(),
		util.EncodeU32(id)...,
	)
}
