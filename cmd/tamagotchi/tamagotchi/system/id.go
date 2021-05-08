package system

import "tamagotchi/cmd/tamagotchi/game"

const (
	BACKGROUND = game.MERGED + 1 + iota
	CHARACTER
	CURSOR
	DRAW
	FPS_COUNTER
	MOUSE
	TWEEN
)
