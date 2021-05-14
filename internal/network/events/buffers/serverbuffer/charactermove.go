package serverbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
	"github.com/cstria0106/tamagotchi/internal/util"
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
