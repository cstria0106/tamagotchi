package serverbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
	"tamagotchi/util"
)

func CleanPooBuffer(uid uint32) []byte {
	return append(header.Header{
		Type:   events.CleanPoo,
		Length: 4,
	}.Buffer(), util.EncodeU32(uid)...)
}
