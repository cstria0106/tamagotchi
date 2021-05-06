package clientbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
)

func FeedBuffer(x, y uint16) []byte {
	return append(header.Header{
		Type:   events.Feed,
		Length: 4,
	}.Buffer(),
		append(util.EncodeU16(x), util.EncodeU16(y)...)...,
	)
}
