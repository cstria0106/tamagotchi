package component

import "tamagotchi/cmd/tamagotchi/game"

const (
	CHARACTER game.ComponentID = iota
	CURSOR
	DRAWABLE
	MOUSEEVENT
	TWEEN
)
