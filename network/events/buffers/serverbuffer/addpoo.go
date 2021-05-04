package serverbuffer

import (
	"tamagotchi/network/events"
	"tamagotchi/network/header"
	"tamagotchi/util"
)

func AddPooBuffer(id uint32, x, y uint16) []byte {
	buffer := util.EncodeU32(id)
	buffer = append(buffer, util.EncodeU16(x)...)
	buffer = append(buffer, util.EncodeU16(y)...)

	return append(
		header.Header{
			Type:   events.AddPoo,
			Length: 8,
		}.Buffer(),
		buffer...,
	)
}
