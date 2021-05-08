package component

import (
	"tamagotchi/cmd/tamagotchi/game"
)

type Character struct {
	X float64
	Y float64
}

func NewCharacter() *game.Component {
	return game.NewComponent(
		CHARACTER, &Character{},
	)
}
