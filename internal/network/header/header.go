package header

import (
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/util"
)

type Header struct {
	Type   events.EventType
	Length uint32
}

func (h Header) Buffer() []byte {
	return append(util.EncodeU16(uint16(h.Type)), util.EncodeU32(h.Length)...)
}

func FromBuffer(buffer []byte) Header {
	return Header{
		Type:   events.EventType(util.DecodeU16(buffer[0:2])),
		Length: util.DecodeU32(buffer[2:6]),
	}
}
