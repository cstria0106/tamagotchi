package system

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/engine"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

var Background = engine.NewSystem(&engine.SystemOptions{
	Draw: func(_ game.Game, screen *ebiten.Image) {
		screen.Fill(image.White)
	},
})
