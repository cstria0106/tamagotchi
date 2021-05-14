package serverbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
	"github.com/cstria0106/tamagotchi/internal/util"
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
