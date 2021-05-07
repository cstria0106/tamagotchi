package component

import (
	"github.com/Tarliton/collision2d"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/segmentio/ksuid"
	"tamagotchi/cmd/tamagotchi/game"
)

var ClickableUID = ksuid.New()

type Clickable struct {
	game.BaseComponent
	X, Y, W, H                    int16
	Hover, Down, JustDown, JustUp bool
	polygon                       collision2d.Polygon
}

func (c *Clickable) Init(_ game.Entity) error {
	c.polygon = collision2d.NewBox(collision2d.Vector{X: float64(c.X), Y: float64(c.Y)}, float64(c.W), float64(c.H)).ToPolygon()
	return nil
}

func (c *Clickable) GetComponentUID() ksuid.KSUID {
	return ClickableUID
}

func (c *Clickable) PreUpdate() error {
	c.Hover = false
	c.Down = false
	c.JustDown = false
	c.JustUp = false

	x, y := ebiten.CursorPosition()

	if collision2d.PointInPolygon(collision2d.Vector{X: float64(x), Y: float64(y)}, c.polygon) {
		c.Hover = true
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			c.Down = true
			if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
				c.JustDown = true
			}
		} else if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
			c.JustUp = true
		}
	}

	return nil
}

func (c *Clickable) Clone() game.Component {
	return &Clickable{X: c.X, Y: c.Y, W: c.W, H: c.H}
}
