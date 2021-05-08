package component

import (
	"tamagotchi/cmd/tamagotchi/game"
)

type Cursor struct{}

func NewCursor() *game.Component {
	return game.NewComponent(CURSOR, &Cursor{})
}
