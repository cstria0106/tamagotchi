package serverbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
	"github.com/cstria0106/tamagotchi/internal/util"
)

func CleanPooBuffer(uid uint32) []byte {
	return append(header.Header{
		Type:   events.CleanPoo,
		Length: 4,
	}.Buffer(), util.EncodeU32(uid)...)
}
