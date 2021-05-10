package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lucasb-eyer/go-colorful"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func Cursor() *game.System {
	return &game.System{
		Init: func(g *game.Game) error {
			ebiten.SetCursorMode(ebiten.CursorModeHidden)

			g.WithComponents([]game.ComponentID{component.DRAWABLE, component.CURSOR}, func(components []*game.Component) {
				drawable := components[0].Data.(*component.Drawable)
				drawable.Color = colorful.FastHappyColor()
			})

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
