package clientbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
	"github.com/cstria0106/tamagotchi/internal/util"
)

func FeedBuffer(x, y uint16) []byte {
	return append(header.Header{
		Type:   events.Feed,
		Length: 4,
	}.Buffer(),
		append(util.EncodeU16(x), util.EncodeU16(y)...)...,
	)
}
