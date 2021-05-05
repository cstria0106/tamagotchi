package clientbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
	"tamagotchi/util"
)

func CleanBuffer(id uint32) []byte {
	return append(
		header.Header{Type: events.Clean, Length: 4}.Buffer(),
		util.EncodeU32(id)...,
	)
}
