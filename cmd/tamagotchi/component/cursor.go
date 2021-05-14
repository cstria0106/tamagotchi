package component

import (
	"github.com/cstria0106/tamagotchi/internal/game"
)

type Cursor struct{}

func NewCursor() *game.Component {
	return game.NewComponent(CURSOR, &Cursor{})
}
