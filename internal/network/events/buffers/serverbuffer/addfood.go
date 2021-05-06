package serverbuffer

import (
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
	"tamagotchi/internal/util"
)

func AddFoodBuffer(id uint32, x, y uint16) []byte {
	buffer := util.EncodeU32(id)
	buffer = append(buffer, util.EncodeU16(x)...)
	buffer = append(buffer, util.EncodeU16(y)...)

	return append(
		header.Header{
			Type:   events.AddFood,
			Length: 8,
		}.Buffer(),
		buffer...,
	)
}
