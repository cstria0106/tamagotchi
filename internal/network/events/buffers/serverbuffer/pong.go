package serverbuffer

import (
	"tamagotchi/internal/data/version"
	"tamagotchi/internal/network/events"
	"tamagotchi/internal/network/header"
)

func PongBuffer(version *version.Version) []byte {
	buffer, length := version.Buffer()
	return append(header.Header{Type: events.Pong, Length: length}.Buffer(), buffer...)
}
