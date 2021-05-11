package component

import (
	"tamagotchi/internal/game"
)

type Cursor struct{}

func NewCursor() *game.Component {
	return game.NewComponent(CURSOR, &Cursor{})
}
