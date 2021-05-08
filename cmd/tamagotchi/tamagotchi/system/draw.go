package system

import (
	"github.com/hajimehoshi/ebiten/v2"
	"tamagotchi/cmd/tamagotchi/game"
	"tamagotchi/cmd/tamagotchi/tamagotchi/component"
)

func Draw() *game.System {
	return &game.System{
		ID: DRAW,
		Draw: func(g *game.Game, screen *ebiten.Image) {
			g.WithComponents([]game.ComponentID{component.DRAWABLE}, func(components []*game.Component) {
				drawable := components[0].Data.(*component.Drawable)

				options := ebiten.DrawImageOptions{}
				options.GeoM.Translate(float64(drawable.X), float64(drawable.Y))

				screen.DrawImage(drawable.Image, &options)
			})
		},
	}
}
