package serverbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
	"tamagotchi/util"
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
