package component

import (
	"github.com/cstria0106/tamagotchi/internal/game"
)

type Character struct {
	X       int16
	Y       int16
	LastX   int16
	LastY   int16
	TargetX int16
	TargetY int16
	Moving  bool
}

func NewCharacter(x, y int16) *game.Component {
	return game.NewComponent(
		CHARACTER, &Character{
			X:       x,
			Y:       y,
			LastX:   x,
			LastY:   y,
			TargetX: x,
			TargetY: y,
		},
	)
}

func (c *Character) MoveTo(x int16, y int16) {
	c.LastX = c.X
	c.LastY = c.Y
	c.TargetX = x
	c.TargetY = y
	c.Moving = true
}
