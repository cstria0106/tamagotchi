package events

type Event struct {
	Type    EventType
	Payload []byte
}

type EventType uint16

const (
	Ping EventType = 1 + iota
	Pong
	Feed
	Clean
	Chat
	ClientEnd
)

const (
	AddFood = ClientEnd + iota
	CharacterMove
	AddPoo
	CleanPoo
	CharacterEat
	ShowChat
)

var eventStringMap = map[EventType]string{
	Ping:          "Ping",
	Pong:          "Pong",
	Feed:          "Feed",
	Clean:         "Clean",
	CharacterEat:  "CharacterEat",
	CharacterMove: "CharacterMove",
	AddPoo:        "AddPoo",
	AddFood:       "AddFood",
	CleanPoo:      "CleanPoo",
	Chat:          "Chat",
	ShowChat:      "ShowChat",
}

var eventPayloadLengthMap = map[EventType]uint32{
	Ping:  0,
	Pong:  0,
	Feed:  4,
	Clean: 4,
}

func (a EventType) String() string {
	return eventStringMap[a]
}

func (e EventType) ValidatePayloadLength(length uint32) bool {
	if required, ok := eventPayloadLengthMap[e]; ok {
		return required <= length
	}

	return true
}
