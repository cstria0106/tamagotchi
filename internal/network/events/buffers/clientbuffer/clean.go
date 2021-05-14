package clientbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
	"github.com/cstria0106/tamagotchi/internal/util"
)

func CleanBuffer(id uint32) []byte {
	return append(
		header.Header{Type: events.Clean, Length: 4}.Buffer(),
		util.EncodeU32(id)...,
	)
}
