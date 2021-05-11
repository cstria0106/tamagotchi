package component

import "tamagotchi/internal/game"

const (
	CHARACTER game.ComponentID = iota
	CURSOR
	DRAWABLE
	MOUSEEVENT
	TWEEN
	NETWORK
)
