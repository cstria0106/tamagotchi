package serverbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
)

func CharacterMoveBuffer(x, y uint16) []byte {
	return append(
		header.Header{
			Type:   events.CharacterMove,
			Length: 4,
		}.Buffer(),
		append(util.EncodeU16(x), util.EncodeU16(y)...)...,
	)
}
