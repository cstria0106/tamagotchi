package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func Cursor() *game.System {
	return &game.System{
		ID: CURSOR,
		Init: func(_ *game.Game) error {
			ebiten.SetCursorMode(ebiten.CursorModeHidden)
			return nil
		},
		Update: func(g *game.Game) error {
			g.WithComponents([]game.ComponentID{component.DRAWABLE, component.CURSOR}, func(components []*game.Component) {
				x, y := ebiten.CursorPosition()
				drawable := components[0].Data.(*component.Drawable)
				drawable.X = int16(x)
				drawable.Y = int16(y)
			})

			return nil
		},
	}
}
