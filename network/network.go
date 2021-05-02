package network

import "encoding/binary"

type Action uint16

const (
	ActionPing Action = 1 + iota
	ActionPong
	ActionFeed
	ActionClean
)

func (a *Action) ToString() string {
	actionStringMap := map[Action]string{
		ActionPing:  "PING",
		ActionPong:  "PONG",
		ActionFeed:  "FEED",
		ActionClean: "CLEAN",
	}

	return actionStringMap[*a]
}

type Payload struct {
	Action     Action
	DataLength uint32
	Data       []byte
}

func BuildPayload(action Action, data []byte) *Payload {
	if data == nil {
		return &Payload{
			Action:     action,
			DataLength: 0,
			Data:       nil,
		}
	}

	return &Payload{
		Action:     action,
		DataLength: uint32(len(data)),
		Data:       data,
	}
}

func (p *Payload) ToBuffer() []byte {
	size := 6 + p.DataLength
	buffer := make([]byte, size)

	binary.BigEndian.PutUint16(buffer[0:2], uint16(p.Action))
	binary.BigEndian.PutUint32(buffer[2:6], p.DataLength)

	if p.DataLength != 0 {
		copy(buffer[6:size], p.Data)
	}

	return buffer[:]
}
