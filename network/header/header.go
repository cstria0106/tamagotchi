package header

import (
	"tamagotchi/network/events"
	"tamagotchi/util"
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
