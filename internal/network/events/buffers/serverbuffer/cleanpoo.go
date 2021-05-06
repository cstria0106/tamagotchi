package serverbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
)

func CleanPooBuffer(uid uint32) []byte {
	return append(header.Header{
		Type:   events.CleanPoo,
		Length: 4,
	}.Buffer(), util.EncodeU32(uid)...)
}
