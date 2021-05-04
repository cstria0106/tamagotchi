package serverbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
	"tamagotchi/util"
)

func CharacterEatBuffer(uid uint32) []byte {
	return append(header.Header{
		Type:   events.CharacterEat,
		Length: 4,
	}.Buffer(), util.EncodeU32(uid)...)
}