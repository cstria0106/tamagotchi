package system

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"tamagotchi/cmd/tamagotchi/component"
	"tamagotchi/cmd/tamagotchi/engine"
	"tamagotchi/internal/game"
)

var Mouse = engine.NewSystem(&engine.SystemOptions{
	Update: func(g game.Game) error {
		g.WithComponents([]game.ComponentID{component.MOUSEEVENT}, func(components []*game.Component) {
			clickable := components[0].Data.(*component.MouseEvent)

			clickable.Hover = false
			clickable.Down = false
			clickable.JustDown = false
			clickable.JustUp = false

			x, y := ebiten.CursorPosition()

			if collision2d.PointInPolygon(collision2d.Vector{X: float64(x), Y: float64(y)}, clickable.Polygon) {
				clickable.Hover = true
				if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
					clickable.Down = true
					if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
						clickable.JustDown = true
					}
				} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
					clickable.JustUp = true
				}
			}
		})

		return nil
	},
})
