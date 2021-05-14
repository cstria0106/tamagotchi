package serverbuffer

import (
	"github.com/cstria0106/tamagotchi/internal/data/version"
	"github.com/cstria0106/tamagotchi/internal/network/events"
	"github.com/cstria0106/tamagotchi/internal/network/header"
)

func PongBuffer(version *version.Version) []byte {
	buffer, length := version.Buffer()
	return append(header.Header{Type: events.Pong, Length: length}.Buffer(), buffer...)
}
