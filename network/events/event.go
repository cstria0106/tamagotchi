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
	ClientEnd
)

const (
	AddFood = ClientEnd + iota
	CharacterMove
	AddPoo
	CleanPoo
	CharacterEat
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
}

func (a EventType) String() string {
	return eventStringMap[a]
}
