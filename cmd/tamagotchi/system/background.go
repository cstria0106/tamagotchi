package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"tamagotchi/cmd/tamagotchi/engine"
	"tamagotchi/internal/game"
)

var Background = engine.NewSystem(&engine.SystemOptions{
	Draw: func(_ game.Game, screen *ebiten.Image) {
		screen.Fill(image.White)
	},
})
