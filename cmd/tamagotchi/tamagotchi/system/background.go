package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"tamagotchi/cmd/tamagotchi/game"
)

func Background() *game.System {
	return &game.System{
		ID: BACKGROUND,
		Draw: func(_ *game.Game, screen *ebiten.Image) {
			screen.Fill(image.White)
		},
	}
}
