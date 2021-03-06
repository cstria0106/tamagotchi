package system

import (
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/component"
	"github.com/cstria0106/tamagotchi/cmd/tamagotchi/engine"
	"github.com/cstria0106/tamagotchi/internal/game"
	"github.com/hajimehoshi/ebiten/v2"
)

var Draw = engine.NewSystem(&engine.SystemOptions{
	Draw: func(g game.Game, screen *ebiten.Image) {
		g.WithComponents([]game.ComponentID{component.DRAWABLE}, func(components []*game.Component) {
			drawable := components[0].Data.(*component.Drawable)

			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(float64(drawable.X), float64(drawable.Y))

			colorR, colorG, colorB, _ := drawable.Color.RGBA()
			options.ColorM.Translate(float64(colorR)/0xffff, float64(colorG)/0xffff, float64(colorB)/0xffff, 0)

			screen.DrawImage(drawable.Image, &options)
		})
	},
})
