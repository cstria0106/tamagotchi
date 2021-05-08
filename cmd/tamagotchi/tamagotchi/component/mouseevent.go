package component

import (
	"github.com/Tarliton/collision2d"
	"tamagotchi/cmd/tamagotchi/game"
)

type MouseEvent struct {
	Hover, Down, JustDown, JustUp bool
	Polygon                       collision2d.Polygon
}

func NewClickable(x, y, w, h int16) *game.Component {
	return game.NewComponent(MOUSEEVENT, &MouseEvent{
		Polygon: collision2d.NewBox(
			collision2d.Vector{X: float64(x), Y: float64(y)}, float64(w), float64(h),
		).ToPolygon(),
	})
}

func (c *MouseEvent) Move(x, y int16) {
	c.Polygon.Pos = c.Polygon.Pos.Add(collision2d.Vector{X: float64(x), Y: float64(y)})
}

func (c *MouseEvent) MoveTo(x, y int16) {
	c.Polygon.Pos = collision2d.Vector{X: float64(x), Y: float64(y)}
}
